package main

import (
	"fmt"
	"log"
	"log/slog"
	"strconv"
	"strings"

	"github.com/go-clang/bootstrap/clang"
)

type Parser struct {
	path string
	mod  *Module
	tu   clang.TranslationUnit
}

func (p *Parser) ParseFile(indent string, path string) {
	p.path = path

	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(
		path,
		[]string{
			"-fparse-all-comments",
			"-I/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/Kernel.framework/Headers/",
		},
		nil,
		clang.TranslationUnit_IncludeBriefCommentsInCodeCompletion|clang.TranslationUnit_DetailedPreprocessingRecord,
	)
	defer tu.Dispose()

	diagnostics := tu.Diagnostics()
	for _, d := range diagnostics {
		fmt.Println("PROBLEM:", d.Spelling())
	}

	p.tu = tu

	tu.
		TranslationUnitCursor().
		Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
			p.parseTopLevel(indent, cursor)

			return clang.ChildVisit_Continue
		})
}

func (p *Parser) parseTopLevel(indent string, c clang.Cursor) {
	loc, _, _, _ := c.Location().FileLocation()
	if loc.Name() != p.path {
		return
	}

	switch c.Kind() {
	case clang.Cursor_TypedefDecl:
		p.parseTypedef(indent, c)
	//
	//case clang.Cursor_VarDecl:
	//	log.Println("vardecl", "name", c.Spelling())
	//	log.Println(" ", c.Type().Spelling())
	//
	case clang.Cursor_FunctionDecl:
		p.parseFunction(indent, c)

	case clang.Cursor_MacroDefinition:
		if c.IsMacroFunctionLike() {
			return
		}

		name := c.Spelling()

		log.Println("macro def", "name", name)

	case clang.Cursor_EnumDecl:
		p.parseEnum(indent, c, false)

	case clang.Cursor_StructDecl:
		p.parseStruct(indent, c, false)

	case clang.Cursor_MacroExpansion:
		name := c.Spelling()

		switch name {
		case "VMA_CALL_PRE", "VMA_CALL_POST", "VMA_NULLABLE", "VMA_NOT_NULL", "VMA_NULLABLE_NON_DISPATCHABLE":
			// uninteresting

		case "VK_DEFINE_HANDLE", "VK_DEFINE_NON_DISPATCHABLE_HANDLE":
			parts := getMacroOriginal(c)

			if len(parts) != 4 {
				panic("unexpected macro definition")
			}

			item := parts[2]

			slog.Info("Found handle", "name", item)

			p.mod.Handles[item] = &Handle{
				Name:            item,
				NonDispatchable: name == "VK_DEFINE_NON_DISPATCHABLE_HANDLE",
			}

		default:
			log.Println("MACRO EXPANSION", c.Spelling())

		}

	case clang.Cursor_UnionDecl:
		log.Println("UNION EXPANSION", c.Spelling())

	case clang.Cursor_InclusionDirective:
		//		ignore

	default:
		log.Panicln("Unexpected top level", "kind", c.Kind().String())
	}
}

func (p *Parser) parseTypedef(indent string, c clang.Cursor) {
	name := c.Spelling()
	indent = fmt.Sprintf("%v[%v]", indent, name)

	log.Println("typedef", "name", c.Spelling())

}

func (p *Parser) parseStruct(indent string, c clang.Cursor, typedef bool) {
	log.Println("struct", "name", c.Spelling())

	name := c.Spelling()
	indent = fmt.Sprintf("%v[%v]", indent, name)

	if strings.HasSuffix(name, "_T") {
		return
	}

	comment := processComment(c.RawCommentText())

	s := &Struct{
		Name:    name,
		Comment: comment,
	}

	p.mod.Structs[name] = s

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_FieldDecl {
			name := cursor.Spelling()
			if name == "" {
				log.Fatal("no field name")
			}

			cmt := processComment(cursor.RawCommentText())

			fIndent := fmt.Sprintf("%v[%v]", indent, name)

			if val := cursor.FieldDeclBitWidth(); val != -1 {
				panic("bitfield not supported: " + strconv.Itoa(int(val)))
			}

			ty := p.parseType(fIndent, cursor.Type())

			s.Fields = append(s.Fields, &Field{
				Name:    name,
				Type:    ty,
				Comment: cmt,
			})
		}

		return clang.ChildVisit_Continue
	})
}

func (p *Parser) parseEnum(indent string, c clang.Cursor, typedef bool) {
	//log.Println("enum", "name", c.Spelling())

	name := c.Spelling()

	slog.Info("Parsing enum", "name", name)

	indent = fmt.Sprintf("%v[%v]", indent, name)

	if strings.Contains(name, " ") {
		panic("enum name contains spaces " + name)
	}

	comment := processComment(c.RawCommentText())

	slog.Info("comment", "comment", comment)

	if _, ok := p.mod.Enums[name]; ok {
		panic("duplicate enum name: " + name)
	}

	enum := &Enum{
		Name: name,
		//Typedefd: typedef,
		Comment: comment,
	}

	p.mod.Enums[name] = enum

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_EnumConstantDecl {
			panic("unknown enum type " + cursor.Kind().String())
		}

		comment := processComment(cursor.RawCommentText())

		enum.Constants = append(enum.Constants, &Constant{
			Name:    cursor.Spelling(),
			Comment: comment,
		})

		slog.Info("value", "name", cursor.Spelling(), "comment", comment)

		return clang.ChildVisit_Continue
	})
}

func (p *Parser) parseFunction(indent string, c clang.Cursor) {
	name := c.Spelling()

	log.Println("Parsing function", name)
	indent = fmt.Sprintf("%v[%v]", indent, name)

	comment := processComment(c.RawCommentText())

	result := p.parseType(fmt.Sprintf("%v[return]", indent), c.ResultType())

	var params []*Param

	for i := 0; i < int(c.NumArguments()); i++ {
		arg := c.Argument(uint32(i))

		if arg.Kind() != clang.Cursor_ParmDecl {
			log.Panicln(indent, "Argument not of parmdecl type", arg.Kind())
		}

		name := arg.Spelling()
		if name == "" {
			log.Println(indent, "no param name")

			name = fmt.Sprintf("param%v", i)
		}

		aIndent := fmt.Sprintf("%v[%v]", indent, name)

		ty := p.parseType(aIndent, arg.Type())

		params = append(params, &Param{
			Name: name,
			Type: ty,
		})
	}

	if c.IsVariadic() {
		panic("variadic not supported")
	}

	p.mod.Functions[name] = &Function{
		Name:       name,
		Args:       params,
		ReturnType: result,
		Comment:    comment,
	}
}

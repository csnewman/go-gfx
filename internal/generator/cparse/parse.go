package cparse

import (
	"fmt"
	"generator/repo"
	"log"
	"log/slog"
	"strconv"
	"strings"

	"github.com/go-clang/bootstrap/clang"
)

type Parser struct {
	path string
	repo *repo.Repo
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

			converted := &repo.Type{
				Name:       item,
				MappedName: convertEnumName(item),
				Category:   repo.TypeCategoryHandle,
			}

			if name == "VK_DEFINE_NON_DISPATCHABLE_HANDLE" {
				converted.Category = repo.TypeCategoryHandleNonDispatchable
			}

			p.repo.Types[item] = converted

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

	s := &repo.Type{
		Name:         name,
		MappedName:   convertStructName(name),
		Category:     repo.TypeCategoryStruct,
		Comment:      comment,
		StructFields: make(map[string]*repo.Field),
	}

	p.repo.Types[name] = s

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
			ty.Name = name
			ty.MappedName = strings.ToUpper(name[:1]) + name[1:]
			ty.Comment = cmt

			s.StructFields[name] = ty
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

	enum := &repo.Type{
		Name:       name,
		MappedName: convertEnumName(name),
		Category:   repo.TypeCategoryEnum,
		Comment:    comment,
		EnumValues: make(map[string]*repo.EnumValue),
	}

	if strings.Contains(name, "FlagBits") {
		convName := strings.ReplaceAll(name, "FlagBits", "Flags")
		enum.Name = convName
		enum.MappedName = convertEnumName(convName)
		enum.Category = repo.TypeCategoryBitmask
		enum.BitmaskWidth = 32
		enum.Aliases = map[string]*repo.Alias{
			name: {
				Name:         name,
				MappedName:   convertBitmaskName(name),
				NoGeneration: true,
			},
		}
	}

	if _, ok := p.repo.Types[enum.Name]; ok {
		panic("duplicate enum name: " + name)
	}

	p.repo.Types[enum.Name] = enum

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_EnumConstantDecl {
			panic("unknown enum type " + cursor.Kind().String())
		}

		comment := processComment(cursor.RawCommentText())
		name := cursor.Spelling()

		enum.EnumValues[name] = &repo.EnumValue{
			Name:       name,
			MappedName: convertEnumItem(name),
			Comment:    comment,
		}

		slog.Info("value", "name", cursor.Spelling(), "comment", comment)

		return clang.ChildVisit_Continue
	})
}

func convertEnumItem(in string) string {
	name, ok := strings.CutPrefix(in, "VMA_")
	if !ok {
		panic(fmt.Sprintf("item does not have prefix: %s", in))
	}

	return name
}

func (p *Parser) parseFunction(indent string, c clang.Cursor) {
	name := c.Spelling()

	log.Println("Parsing function", name)
	indent = fmt.Sprintf("%v[%v]", indent, name)

	comment := processComment(c.RawCommentText())

	result := p.parseType(fmt.Sprintf("%v[return]", indent), c.ResultType())

	if result.Category == repo.FieldCategoryDirect && result.Type == "void" {
		result = nil
	}

	var members []*repo.Field

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
		ty.Name = name
		ty.MappedName = name

		members = append(members, ty)
	}

	if c.IsVariadic() {
		panic("variadic not supported")
	}

	p.repo.Functions[name] = &repo.Function{
		Name:       name,
		MappedName: convertFuncName(name),
		Aliases:    nil,
		ReturnType: result,
		Members:    members,
		Comment:    comment,
	}
}

func convertFuncName(in string) string {
	name, ok := strings.CutPrefix(in, "vma")
	if !ok {
		panic(fmt.Sprintf("cmd does not have prefix: %s", in))
	}

	return name
}

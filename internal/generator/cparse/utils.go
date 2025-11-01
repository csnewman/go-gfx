package cparse

import (
	"fmt"
	"generator/repo"
	"log"
	"strings"
	"unicode"

	"github.com/go-clang/bootstrap/clang"
)

func processComment(in string) string {
	txt := in
	txt = strings.ReplaceAll(txt, "\r\n", "\n")
	txt = strings.TrimSpace(txt)
	txt = strings.TrimPrefix(txt, "/**\n")
	txt = strings.TrimSuffix(txt, "*/")

	txt = strings.TrimRightFunc(txt, unicode.IsSpace)

	var rebuilt []string

	for _, s := range strings.Split(txt, "\n") {
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "* ")
		s = strings.TrimPrefix(s, "/// ")
		s = strings.TrimPrefix(s, "///")
		s = strings.TrimPrefix(s, "// ")
		s = strings.TrimPrefix(s, "//")

		if strings.HasPrefix(s, "/**") {
			rebuilt = nil
			s = strings.TrimPrefix(s, "/** ")
			s = strings.TrimPrefix(s, "/**")

			if strings.TrimSpace(s) == "" {
				continue
			}
		}

		if strings.HasPrefix(s, "/*") {
			rebuilt = nil
			s = strings.TrimPrefix(s, "/* ")
			s = strings.TrimPrefix(s, "/*")

			if strings.TrimSpace(s) == "" {
				continue
			}
		}

		if strings.HasPrefix(s, "@defgroup") || strings.HasPrefix(s, "@ingroup") ||
			strings.HasPrefix(s, "@addtogroup") || strings.HasPrefix(s, "@}") {
			continue
		}

		if strings.HasPrefix(s, "@{") {
			rebuilt = nil
			continue
		}

		if s == "*" {
			s = ""
		}

		// Double space to enter verbatim mode
		rebuilt = append(rebuilt, fmt.Sprintf("  %v", s))
	}

	txt = strings.Join(rebuilt, "\n")
	txt = strings.TrimRightFunc(txt, unicode.IsSpace)

	if strings.Count(txt, "\n") == 0 {
		txt = strings.TrimLeftFunc(txt, unicode.IsSpace)
		txt = strings.TrimPrefix(txt, "<")
		txt = fmt.Sprintf("  %v", txt)
	}

	if strings.TrimSpace(txt) == "" {
		return ""
	}

	return txt
}

func getMacroOriginal(c clang.Cursor) []string {
	toks := c.TranslationUnit().Tokenize(c.Extent())
	out := make([]string, 0, len(toks))

	for _, tok := range toks {
		spel := c.TranslationUnit().TokenSpelling(tok)

		out = append(out, spel)
	}

	return out
}

func (p *Parser) parseType(indent string, t clang.Type) *repo.Field {
	switch t.Kind() {
	case clang.Type_Void:
		log.Println(indent, "Parsing type", t.Spelling(), "as void")

		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     "void",
		}
	case clang.Type_Int, clang.Type_UInt, clang.Type_Long, clang.Type_ULong, clang.Type_UChar, clang.Type_Char_S,
		clang.Type_Float, clang.Type_Double:
		log.Println(indent, "Parsing type", t.Spelling(), "as ident")
		name := t.CanonicalType().Spelling()
		name = strings.TrimPrefix(name, "const ")

		if strings.Contains(name, " ") {
			panic(name + " contains spaces")
		}

		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     name,
		}
	case clang.Type_Pointer:
		log.Println(indent, "Parsing type", t.Spelling(), "as pointer")

		inner := p.parseType(indent, t.PointeeType())

		switch inner.Category {
		case repo.FieldCategoryDirect:
			return &repo.Field{
				Category: repo.FieldCategoryPointer,
				Type:     inner.Type,
			}
		case repo.FieldCategoryPointer:
			return &repo.Field{
				Category: repo.FieldCategoryPointer2,
				Type:     inner.Type,
			}
		default:
			panic("unexpected category " + string(inner.Category))

		}

	case clang.Type_Elaborated:
		log.Println(indent, "Parsing type", t.Spelling(), "as elaborated")

		name := t.Spelling()
		name = strings.TrimPrefix(name, "const ")

		if strings.Contains(name, " ") {
			panic(name + " contains spaces")
		}

		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     name,
		}
	case clang.Type_ConstantArray:
		log.Println(indent, "Parsing type", t.Spelling(), "as const array")
		inner := p.parseType(indent, t.ArrayElementType())
		size := t.ArraySize()

		switch inner.Category {
		case repo.FieldCategoryDirect:
			return &repo.Field{
				Category:  repo.FieldCategoryArray,
				Type:      inner.Type,
				ArraySize: int(size),
			}
		default:
			panic("unexpected category " + string(inner.Category))
		}
	default:
		log.Println(indent, "Parsing type", t.Spelling(), "as ???")
		log.Panicln(indent, "Unknown type", t.Kind().Spelling())
		return nil
	}

	return nil
}

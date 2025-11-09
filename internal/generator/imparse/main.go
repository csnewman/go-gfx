package imparse

import (
	"encoding/json"
	"generator/repo"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type EnumValue struct {
	Name string `json:"name"`
}

type StructMember struct {
	Name         string `json:"name"`
	TemplateType string `json:"template_type"`
	Type         string `json:"type"`
	Size         int    `json:"size"`
	BitFieldSize string `json:"bitfield"`
}

type StructAndEnums struct {
	Enums     map[string][]*EnumValue    `json:"enums"`
	EnumTypes map[string]string          `json:"enumtypes"`
	Locations map[string]string          `json:"locations"`
	Structs   map[string][]*StructMember `json:"structs"`
	NonPOD    map[string]bool            `json:"nonPOD"`
}

type DefArg struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Definition struct {
	Location    string    `json:"location"`
	Args        []*DefArg `json:"argsT"`
	Ret         string    `json:"ret"`
	Templated   bool      `json:"templated"`
	Constructor bool      `json:"constructor"`
	StaticFunc  bool      `json:"is_static_function"`
	STName      string    `json:"stname"`
}

func Parse(path string) *repo.Repo {
	out := &repo.Repo{
		Types:     make(map[string]*repo.Type),
		Functions: make(map[string]*repo.Function),
	}

	out.Types["ImFontAtlasRectId"] = &repo.Type{
		Name:       "ImFontAtlasRectId",
		MappedName: "FontAtlasRectId",
		Category:   repo.TypeCategoryHandleNonDispatchable,
		HandleSize: 32,
	}

	structEnums := readJson[*StructAndEnums](filepath.Join(path, "structs_and_enums.json"))

	convertEnums(out, structEnums)
	convertStructs(out, structEnums)

	defs := readJson[map[string][]*Definition](filepath.Join(path, "definitions.json"))

	convertDefinitions(out, structEnums, defs)

	out.PostProcess()

	return out
}

func convertDefinitions(out *repo.Repo, structEnums *StructAndEnums, defs map[string][]*Definition) {
	for name, inner := range defs {
		var def *Definition

		for _, cdef := range inner {
			if strings.Contains(cdef.Location, "internal") {
				continue
			}

			def = cdef
		}

		if def == nil {
			continue
		}

		if def.Templated {
			continue
		}

		fun := &repo.Function{
			Name:       name,
			MappedName: name,
			Aliases:    nil,
			ReturnType: nil,
			Members:    nil,
			Comment:    "",
		}

		if def.STName == "" {
			var (
				foundIG bool
				foundIM bool
			)

			fun.MappedName, foundIG = strings.CutPrefix(fun.MappedName, "ig")
			fun.MappedName, foundIM = strings.CutPrefix(fun.MappedName, "Im")

			if !foundIM && !foundIG {
				panic("missing prefix: " + name)
			}
		}

		for _, arg := range def.Args {
			ty := parseType(structEnums, arg.Type)

			ty.Name = arg.Name
			ty.MappedName = arg.Name

			fun.Members = append(fun.Members, ty)
		}

		if def.STName != "" && !def.Constructor && !def.StaticFunc {
			foundSelf := false

			for _, member := range fun.Members {
				if member.Category == repo.FieldCategoryPointer && member.Type == def.STName && member.Name == "self" {
					foundSelf = true
					member.IsSelf = true

					break
				}
			}

			if !foundSelf {
				panic("missing self: " + name)
			}

			var ok bool

			fun.MappedName, ok = strings.CutPrefix(fun.MappedName, def.STName+"_")
			if !ok {
				panic("missing self prefix: " + fun.MappedName)
			}
		}

		if def.Constructor {
			var ok bool

			fun.MappedName, ok = strings.CutPrefix(fun.MappedName, def.STName+"_"+def.STName)
			if !ok {
				panic("missing stname prefix: " + fun.MappedName)
			}

			fun.MappedName = convertStructName(def.STName) + "New" + fun.MappedName

			fun.ReturnType = &repo.Field{
				Category: repo.FieldCategoryPointer,
				Type:     def.STName,
			}

		} else if def.Ret != "void" {
			fun.ReturnType = parseType(structEnums, def.Ret)
		}

		fun.MappedName = strings.ToUpper(fun.MappedName[:1]) + fun.MappedName[1:]

		out.Functions[name] = fun

		if def.STName != "" {
			ty, ok := out.Types[def.STName]
			if !ok {
				panic("missing type: " + def.STName)
			}

			if ty.Category != repo.TypeCategoryStruct {
				panic("wrong type: " + def.STName)
			}

			for _, field := range ty.StructFields {
				if strings.EqualFold("set"+field.MappedName, fun.MappedName) {
					field.SuppressSet = true
				}
			}
		}
	}
}

func convertEnums(out *repo.Repo, structEnums *StructAndEnums) {
	for name, vals := range structEnums.Enums {
		slog.Info("Parsing enum", "name", name)

		cleaned, foundPrefix := strings.CutPrefix(name, "Im")
		if !foundPrefix {
			panic("missing prefix: " + name)
		}

		cleaned = strings.TrimSuffix(cleaned, "_")

		width := 32

		if mappping, ok := structEnums.EnumTypes[name]; ok {
			switch mappping {
			case "int":
				width = 32
			case "ImU8":
				width = 8
			default:
				panic("unknown enum type mapping " + mappping)
			}
		}

		ty := &repo.Type{
			Name:       name,
			MappedName: cleaned,
			Category:   repo.TypeCategoryEnum,
			EnumValues: make(map[string]*repo.EnumValue),
		}

		if strings.Contains(cleaned, "Flags") {
			ty.Category = repo.TypeCategoryBitmask
			ty.BitmaskWidth = width
		} else {
			ty.EnumSize = width
		}

		out.Types[name] = ty

		for _, val := range vals {
			cleaned, foundPrefix := strings.CutPrefix(val.Name, "Im")
			if !foundPrefix {
				panic("missing prefix: " + val.Name)
			}

			cleaned = strings.TrimSuffix(cleaned, "_")

			ty.EnumValues[val.Name] = &repo.EnumValue{
				Name:       val.Name,
				MappedName: cleaned,
			}
		}
	}
}

func convertStructName(name string) string {
	cleaned, foundPrefix := strings.CutPrefix(name, "Im")
	if !foundPrefix {
		panic("missing prefix: " + name)
	}

	cleaned = strings.TrimSuffix(cleaned, "_")

	return cleaned
}

func convertStructs(out *repo.Repo, structEnums *StructAndEnums) {
	for name, vals := range structEnums.Structs {
		slog.Info("Parsing struct", "name", name)

		if strings.HasPrefix(name, "stbrp") {
			continue
		}

		cleaned := convertStructName(name)

		_, nonPOD := structEnums.NonPOD[name]

		ty := &repo.Type{
			Name:          name,
			MappedName:    cleaned,
			Category:      repo.TypeCategoryStruct,
			StructFields:  make(map[string]*repo.Field),
			SuppressAlloc: nonPOD,
		}

		out.Types[name] = ty

		for _, val := range vals {
			convertStructItem(ty, structEnums, val)
		}

		_ = vals
	}
}

var nativeMapping = map[string]string{
	"float":          "float",
	"unsigned short": "uint16_t",
	"ImU32":          "uint32_t",
	"int":            "int32_t",
	"bool":           "bool",
	"char":           "char",
	"unsigned char":  "char",
	"void":           "void",

	"ImFontAtlasRectId": "ImFontAtlasRectId",
}

func convertStructItem(ty *repo.Type, enums *StructAndEnums, val *StructMember) {
	mappedName := strings.TrimPrefix(val.Name, "_")

	if len(mappedName) > 1 {
		mappedName = strings.ToUpper(mappedName[:1]) + mappedName[1:]
	} else {
		mappedName = strings.ToUpper(mappedName)
	}

	var out *repo.Field

	if val.BitFieldSize != "" || val.Size != 0 || val.TemplateType != "" {
		out = &repo.Field{
			Category: repo.FieldCategoryUnsupported,
		}
	} else {
		out = parseType(enums, val.Type)
	}

	out.Name = val.Name
	out.MappedName = mappedName

	ty.StructFields[val.Name] = out
}

func parseType(enums *StructAndEnums, in string) *repo.Field {
	slog.Info("Mapping type", "in", in)

	in = strings.TrimPrefix(in, "const ")

	if mapping, ok := nativeMapping[in]; ok {
		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     mapping,
		}
	}

	if _, ok := enums.Structs[in]; ok {
		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     in,
		}
	}

	if _, ok := enums.Enums[in]; ok {
		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     in,
		}
	}

	if _, ok := enums.Enums[in+"_"]; ok {
		return &repo.Field{
			Category: repo.FieldCategoryDirect,
			Type:     in + "_",
		}
	}

	if before, ok := strings.CutSuffix(in, "*"); ok {
		inner := parseType(enums, before)

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
		case repo.FieldCategoryUnsupported:
			return &repo.Field{
				Category: repo.FieldCategoryUnsupported,
			}
		default:
			panic("unexpected inner type: " + inner.Category)
		}
	}

	return &repo.Field{
		Category: repo.FieldCategoryUnsupported,
	}
}

func readJson[T any](path string) T {
	var out T

	enc, err := os.ReadFile(path)
	if err != nil {
		panic(err)

	}

	if err := json.Unmarshal(enc, &out); err != nil {
		panic(err)
	}

	return out
}

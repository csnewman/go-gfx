package main

import (
	"fmt"
	"go/token"
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
)

const ffiPath = "github.com/csnewman/go-gfx/ffi"

func generate(reg *Registry) {
	oEnums := jen.NewFile("vk")
	oBitmask := jen.NewFile("vk")

	oStructs := jen.NewFile("vk")
	oStructs.CgoPreamble(`#include "vulkan.h"`)

	_ = oBitmask

	sortedTypes := slices.Sorted(maps.Keys(reg.Types))

	for _, name := range sortedTypes {
		ty := reg.Types[name]

		if ty.Feature == "" {
			continue
		}

		switch ty.Category {
		case CategoryEnum:
			generateEnumType(reg, oEnums, ty)

		case CategoryBitmask:

		case CategoryStruct:
			generateStructType(reg, oStructs, ty)
		default:
			panic(fmt.Sprintf("unknown type %v", ty.Category))
		}
	}

	if err := oEnums.Save("../../vk/enums.go"); err != nil {
		panic(err)
	}

	if err := oStructs.Save("../../vk/structs.go"); err != nil {
		panic(err)
	}
}

type ToStrEntry struct {
	Local string
	Name  string
}

func convertEnumName(in string) string {
	name, ok := strings.CutPrefix(in, "Vk")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	return name
}

func generateEnumType(reg *Registry, o *jen.File, ty *Type) {
	name := convertEnumName(ty.Name)

	mapping, hasMapping := reg.Enums[ty.Name]
	if hasMapping && mapping.Type == "bitmask" {
		slog.Info("Skipping enum as suspected bitmask", "name", ty.Name)

		return
	}

	slog.Info("Generating enum", "name", ty.Name)

	o.Commentf("%s wraps the enum %s.", name, ty.Name)

	o.Type().Id(name).Id("int32")

	slices.Sort(ty.Aliases)

	for _, alias := range ty.Aliases {
		aliasName, ok := strings.CutPrefix(alias, "Vk")
		if !ok {
			panic(fmt.Sprintf("enum alias does not have prefix: %s", alias))
		}

		o.Commentf("%s wraps the enum %s. An alias for %s", aliasName, alias, name)

		o.Type().Id(aliasName).Op("=").Id(name)
	}

	var entries []ToStrEntry

	if hasMapping {
		o.Const().DefsFunc(func(group *jen.Group) {
			for _, v := range mapping.Values {
				if v.Name == "" {
					panic(fmt.Sprintf("enum does not have a name: %s", ty.Name))
				}

				itemName, ok := strings.CutPrefix(v.Name, "VK_")
				if !ok {
					panic(fmt.Sprintf("enum item does not have prefix: %s", ty.Name))
				}

				if v.Alias != "" {
					aliasName, ok := strings.CutPrefix(v.Alias, "VK_")
					if !ok {
						panic(fmt.Sprintf("enum alias does not have prefix: %s", ty.Name))
					}

					group.Commentf("%s wraps %s.", itemName, v.Name)

					if v.Deprecated != "" {
						group.Comment("")
						group.Commentf("Deprecated: Use %s instead.", aliasName)
					}

					group.Id(itemName).Id(name).Op("=").Id(aliasName)

					continue
				}

				if v.Value == "" {
					panic(fmt.Sprintf("enum does not have a value: %s", ty.Name))
				}

				group.Commentf("%s wraps %s.", itemName, v.Name)
				group.Id(itemName).Id(name).Op("=").Id(v.Value)

				entries = append(entries, ToStrEntry{
					Local: itemName,
					Name:  v.Name,
				})
			}
		})
	}

	o.Func().
		ParamsFunc(func(group *jen.Group) {
			group.Id("e").Id(name)
		}).Id("String").
		Params().
		Id("string").
		BlockFunc(func(group *jen.Group) {
			if len(entries) == 0 {
				group.Return().
					Qual("fmt", "Sprintf").
					Call(
						jen.Lit(fmt.Sprintf("%s(%%d)", ty.Name)),
						jen.Id("e"),
					)

				return
			}

			group.Switch().Id("e").BlockFunc(func(group *jen.Group) {
				for _, entry := range entries {
					group.Case(jen.Id(entry.Local)).Line().Return(jen.Lit(entry.Name))
				}
				group.Default().List().Return().
					Qual("fmt", "Sprintf").
					Call(
						jen.Lit(fmt.Sprintf("%s(%%d)", ty.Name)),
						jen.Id("e"),
					)
			})
		})
	o.Line()
}

func convertStructName(in string) string {
	name, ok := strings.CutPrefix(in, "Vk")
	if !ok {
		panic(fmt.Sprintf("struct does not have prefix: %s", in))
	}

	return name
}

func generateStructType(reg *Registry, o *jen.File, ty *Type) {
	name := convertStructName(ty.Name)

	slog.Info("Generating struct", "name", ty.Name)
	o.Commentf("%s wraps %s.", name, ty.Name)

	cName := "C." + ty.Name

	o.Type().Id(name).StructFunc(func(g *jen.Group) {
		g.Id("ptr").Id("*" + cName)
	})

	o.Commentf("%s is a null pointer.", name+"Nil")
	o.Var().Id(name + "Nil").Id(name)

	sizeOfName := name + "SizeOf"

	o.Commentf("%s is the byte size of %s.", sizeOfName, ty.Name)
	o.Const().Id(sizeOfName).Op("=").Id("int").Call(jen.Id("C.sizeof_" + ty.Name))

	o.Commentf("%s converts a raw pointer to a %s.", name+"FromPtr", name)
	o.Func().Id(name + "FromPtr").
		Params(jen.Id("ptr").Qual("unsafe", "Pointer")).
		Id(name).
		Block(
			jen.Return(jen.Id(name).Values(
				jen.Id("ptr").Op(":").Parens(jen.Id("*" + cName)).Params(jen.Id("ptr"))),
			),
		)

	o.Commentf("%s allocates a continuous block of %s.", name+"Alloc", ty.Name)
	o.Func().Id(name+"Alloc").
		Types(jen.Id("T").Qual(ffiPath, "Allocator")).
		Params(jen.Id("alloc").Id("T"), jen.Id("count").Id("int")).
		Id(name).
		Block(
			jen.Id("ptr").Op(":=").Id("alloc").Dot("Allocate").Call(jen.Id(sizeOfName).Op("*").Id("count")),

			jen.Return(jen.Id(name).Values(
				jen.Id("ptr").Op(":").Parens(jen.Id("*"+cName)).Params(jen.Id("ptr"))),
			),
		)

	o.Comment("Raw returns a raw pointer to the struct.")
	o.Func().Params(jen.Id("p").Id(name)).Id("Raw").Params().Qual("unsafe", "Pointer").Block(
		jen.Return(jen.Qual("unsafe", "Pointer").Call(jen.Id("p").Dot("ptr"))),
	)

	o.Comment("Offset returns an offset pointer by the size of the struct and the provided multiple.")
	o.Func().Params(jen.Id("p").Id(name)).Id("Offset").
		Params(jen.Id("offset").Id("int")).
		Id(name).
		Block(
			jen.Id("ptr").Op(":=").Qual("unsafe", "Add").Call(
				jen.Qual("unsafe", "Pointer").Call(jen.Id("p").Dot("ptr")),
				jen.Id("offset").Op("*").Id(sizeOfName),
			),

			jen.Return(jen.Id(name).Values(
				jen.Id("ptr").Op(":").Parens(jen.Id("*"+cName)).Params(jen.Id("ptr"))),
			),
		)

	for _, field := range ty.Fields {
		generateStructField(reg, o, name, ty, field)
	}
}

var nativeTypes = map[string]string{
	"size_t":   "uintptr",
	"uint8_t":  "uint8",
	"int8_t":   "int8",
	"uint16_t": "uint16",
	"int16_t":  "int16",
	"uint32_t": "uint32",
	"int32_t":  "int32",
	"uint64_t": "uint64",
	"int64_t":  "int64",
}

func generateStructField(reg *Registry, o *jen.File, tyName string, ty *Type, field StructField) {
	fieldName := strings.ToUpper(field.Name[:1]) + field.Name[1:]

	var (
		mappedType jen.Code
		getBody    []jen.Code
		setBody    []jen.Code

		generateDefault bool
	)

	cFieldName := field.Name

	if token.IsKeyword(cFieldName) {
		cFieldName = "_" + cFieldName
	}

	switch field.Category {
	case FieldCategoryDirect:
		if mapping, ok := nativeTypes[field.Type]; ok {
			mappedType = jen.Id(mapping)
			generateDefault = true

			break
		}

		fieldType, ok := reg.Types[field.Type]
		if !ok {
			o.Commentf("%s.%s is unsupported: unknown type %s.", tyName, field.Name, field.Type)
			o.Line()

			return
		}

		switch fieldType.Category {
		case CategoryEnum:
			mapping, hasMapping := reg.Enums[fieldType.Name]
			if hasMapping && mapping.Type == "bitmask" {
				o.Commentf("%s.%s is unsupported: bitmask.", tyName, field.Name, field.Type)
				o.Line()

				return
			}

			enumName := convertEnumName(fieldType.Name)
			mappedType = jen.Id(enumName)
			generateDefault = true

		//case CategoryBitmask, CategoryStruct:

		default:
			o.Commentf("%s.%s is unsupported: direct category %s.", tyName, field.Name, fieldType.Category)
			o.Line()

			return
		}

	case FieldCategoryPointer:
		if field.Type == "void" {
			mappedType = jen.Qual("unsafe", "Pointer")

			getBody = []jen.Code{
				jen.Return(
					jen.Add(mappedType).Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Id("value"),
			}

			break
		}

		if field.Type == "char" {
			mappedType = jen.Qual(ffiPath, "CString")

			getBody = []jen.Code{
				jen.Return(jen.Qual(ffiPath, "CStringFromPtr").Parens(
					jen.Parens(jen.Qual("unsafe", "Pointer")).Parens(
						jen.Id("p").Dot("ptr").Dot(cFieldName))),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").
					Parens(jen.Id("*C.char")).Parens(
					jen.Id("value").Dot("Raw").Params(),
				),
			}
			break
		}

		fieldType, ok := reg.Types[field.Type]
		if ok && fieldType.Category == CategoryStruct {
			structName := convertStructName(fieldType.Name)
			mappedType = jen.Id(structName)

			getBody = []jen.Code{
				jen.Return(jen.Id(structName).Values(
					jen.Id("ptr").Op(":").Id("p").Dot("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Id("value").Dot("ptr"),
			}

			break
		}

		o.Commentf("%s.%s is unsupported: category %s.", tyName, field.Name, field.Category)
		o.Line()

		return

	//case FieldCategoryPointer, FieldCategoryPointer2, FieldCategoryUnsupported:
	default:
		o.Commentf("%s.%s is unsupported: category %s.", tyName, field.Name, field.Category)
		o.Line()

		return
	}

	if generateDefault {
		getBody = []jen.Code{
			jen.Return(
				jen.Add(mappedType).Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
			),
		}

		setBody = []jen.Code{
			jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Parens(jen.Id("C." + field.Type)).Params(jen.Id("value")),
		}
	}

	o.Line()
	o.Commentf("Get%s returns the value in %s.", fieldName, field.Name)
	o.Func().Params(jen.Id("p").Id(tyName)).Id("Get" + fieldName).
		Params().
		Add(mappedType).
		Block(getBody...)

	o.Line()
	o.Commentf("Set%s sets the value in %s.", fieldName, field.Name)
	o.Func().Params(jen.Id("p").Id(tyName)).Id("Set" + fieldName).
		Params(jen.Id("value").Add(mappedType)).
		Block(setBody...)
}

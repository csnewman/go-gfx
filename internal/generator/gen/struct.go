package gen

import (
	"generator/repo"
	"go/token"
	"log/slog"
	"maps"
	"slices"

	"github.com/dave/jennifer/jen"
)

func (g *Generator) generateStructType(ty *repo.Type) {
	slog.Info("Generating struct", "name", ty.Name)

	cName := "C." + ty.Name

	g.OStructs.Commentf("%s wraps struct %s.", ty.MappedName, ty.Name)
	g.OStructs.Type().Id(ty.MappedName).StructFunc(func(g *jen.Group) {
		g.Id("ptr").Id("*" + cName)
	})

	g.OStructs.Commentf("%sNil is a null pointer.", ty.MappedName)
	g.OStructs.Var().Id(ty.MappedName + "Nil").Id(ty.MappedName)

	sizeOfName := ty.MappedName + "SizeOf"

	g.OStructs.Commentf("%s is the byte size of %s.", sizeOfName, ty.Name)
	g.OStructs.Const().Id(sizeOfName).Op("=").Id("int").Call(jen.Id("C.sizeof_" + ty.Name))

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		if alias.NoGeneration {
			continue
		}

		g.OStructs.Commentf("%s wraps struct %s. An alias for %s.", alias.MappedName, alias.Name, ty.MappedName)
		g.OStructs.Type().Id(alias.MappedName).Op("=").Id(ty.MappedName)

		g.OStructs.Commentf("%sNil is a null pointer.", alias.MappedName)
		g.OStructs.Var().Id(alias.MappedName + "Nil").Id(alias.MappedName)

		g.OStructs.Commentf("%sSizeOf is the byte size of %s.", alias.MappedName, alias.MappedName)
		g.OStructs.Const().Id(alias.MappedName + "SizeOf").Op("=").Id(sizeOfName)
	}

	g.OStructs.Commentf("%sFromPtr converts a raw pointer to a %s.", ty.MappedName, ty.MappedName)
	g.OStructs.Func().Id(ty.MappedName + "FromPtr").
		Params(jen.Id("ptr").Qual("unsafe", "Pointer")).
		Id(ty.MappedName).
		Block(
			jen.Return(jen.Id(ty.MappedName).Values(
				jen.Id("ptr").Op(":").Parens(jen.Id("*" + cName)).Params(jen.Id("ptr"))),
			),
		)

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		if alias.NoGeneration {
			continue
		}

		g.OStructs.Commentf("%sFromPtr converts a raw pointer to a %s.", alias.MappedName, alias.MappedName)
		g.OStructs.Func().Id(alias.MappedName + "FromPtr").
			Params(jen.Id("ptr").Qual("unsafe", "Pointer")).
			Id(alias.MappedName).
			Block(
				jen.Return(jen.Id(ty.MappedName + "FromPtr").Params(jen.Id("ptr"))),
			)
	}

	g.OStructs.Commentf("%sAlloc allocates a continuous block of %s.", ty.MappedName, ty.MappedName)
	g.OStructs.Func().Id(ty.MappedName+"Alloc").
		Params(jen.Id("alloc").Qual(ffiPath, "Allocator"), jen.Id("count").Id("int")).
		Id(ty.MappedName).
		Block(
			jen.Id("ptr").Op(":=").Id("alloc").Dot("Allocate").Call(jen.Id(sizeOfName).Op("*").Id("count")),

			jen.Return(jen.Id(ty.MappedName).Values(
				jen.Id("ptr").Op(":").Parens(jen.Id("*"+cName)).Params(jen.Id("ptr"))),
			),
		)

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		if alias.NoGeneration {
			continue
		}

		g.OStructs.Commentf("%sAlloc allocates a continuous block of %s.", alias.MappedName, alias.MappedName)
		g.OStructs.Func().Id(alias.MappedName+"Alloc").
			Params(jen.Id("alloc").Qual(ffiPath, "Allocator"), jen.Id("count").Id("int")).
			Id(alias.MappedName).
			Block(
				jen.Return(jen.Id(ty.MappedName+"Alloc").Params(jen.Id("alloc"), jen.Id("count"))),
			)
	}

	g.OStructs.Comment("Raw returns a raw pointer to the struct.")
	g.OStructs.Func().
		Params(jen.Id("p").Id(ty.MappedName)).
		Id("Raw").
		Params().
		Qual("unsafe", "Pointer").
		Block(
			jen.Return(jen.Qual("unsafe", "Pointer").Call(jen.Id("p").Dot("ptr"))),
		)

	g.OStructs.Comment("Offset returns an offset pointer by the size of the struct and the provided multiple.")
	g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Offset").
		Params(jen.Id("offset").Id("int")).
		Id(ty.MappedName).
		Block(
			jen.Id("ptr").Op(":=").Qual("unsafe", "Add").Call(
				jen.Qual("unsafe", "Pointer").Call(jen.Id("p").Dot("ptr")),
				jen.Id("offset").Op("*").Id(sizeOfName),
			),

			jen.Return(jen.Id(ty.MappedName).Values(
				jen.Id("ptr").Op(":").Parens(jen.Id("*"+cName)).Params(jen.Id("ptr"))),
			),
		)

	for _, name := range slices.Sorted(maps.Keys(ty.StructFields)) {
		field := ty.StructFields[name]

		g.generateStructField(ty, field)
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
	"float":    "float32",

	"VkDeviceSize":    "DeviceSize",
	"VkDeviceAddress": "DeviceAddress",
	"VkSampleMask":    "SampleMask",
}

func (g *Generator) generateStructField(ty *repo.Type, field *repo.Field) {
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
	case repo.FieldCategoryDirect:
		if mapping, ok := nativeTypes[field.Type]; ok {
			mappedType = jen.Id(mapping)
			generateDefault = true

			break
		}

		if field.Type == "VkBool32" {
			mappedType = jen.Id("bool")

			getBody = []jen.Code{
				jen.Return(
					jen.Id("p").Dot("ptr").Dot(cFieldName).Op("!=").Id("0"),
				),
			}

			setBody = []jen.Code{
				jen.If(jen.Id("value")).Block(
					jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Id("C.VkBool32").Parens(jen.Id("1")),
				).Else().Block(
					jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Id("C.VkBool32").Parens(jen.Id("0")),
				),
			}

			break
		}

		if field.Type == "HINSTANCE" || field.Type == "HWND" {
			mappedType = jen.Qual("unsafe", "Pointer")
			//
			//if fieldType.NonDispatchable {
			//	generateDefault = true
			//} else {
			getBody = []jen.Code{
				jen.Return(
					jen.Qual("unsafe", "Pointer").Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Parens(
					jen.Id("C." + field.Type)).Params(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))),
			}
			//}

			break
		}

		fieldType, ok := g.Repo.LookupType(field.Type)
		if !ok {
			g.OStructs.Commentf("%s.%s is unsupported: unknown type %s.", ty.MappedName, field.Name, field.Type)
			g.OStructs.Line()

			return
		}

		switch fieldType.Category {
		case repo.TypeCategoryEnum, repo.TypeCategoryBitmask, repo.TypeCategoryHandleNonDispatchable:
			mappedType = jen.Id(fieldType.MappedName)
			generateDefault = true

		case repo.TypeCategoryHandle:
			mappedType = jen.Id(fieldType.MappedName)

			getBody = []jen.Code{
				jen.Return(
					jen.Add(mappedType).Params(
						jen.Qual("unsafe", "Pointer").Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Parens(
					jen.Id("C." + field.Type)).Params(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))),
			}

		case repo.TypeCategoryStruct:
			mappedType = jen.Id(fieldType.MappedName)

			offsetName := "offsetof_" + ty.Name + "_" + field.Name

			if _, ok := g.StructGeneratedOffsets[offsetName]; !ok {
				g.StructGeneratedOffsets[offsetName] = struct{}{}

				g.OStructs.CgoPreamble("static const int " + offsetName + " = offsetof(" + ty.Name + ", " + field.Name + ");")
			}

			g.OStructs.Line()
			g.OStructs.Commentf("Ref%s returns pointer to the %s field.", field.MappedName, field.Name)
			g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Ref" + field.MappedName).
				Params().
				Add(mappedType).
				Block(
					jen.Return(jen.Id(fieldType.MappedName).Values(
						jen.Id("ptr").Op(":").
							Params(jen.Id("*C." + fieldType.Name)).
							Params(
								jen.Qual("unsafe", "Add").
									Call(
										jen.Qual("unsafe", "Pointer").Call(jen.Id("p").Dot("ptr")),
										jen.Id("uintptr").Call(jen.Qual("C", offsetName)),
									),
							),
					)),
				)

			return

		default:
			g.OStructs.Commentf("%s.%s is unsupported: direct category %s.", ty.MappedName, field.Name, fieldType.Category)
			g.OStructs.Line()

			return
		}

	case repo.FieldCategoryPointer:
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

		if field.Type == "CAMetalLayer" {
			mappedType = jen.Qual("unsafe", "Pointer")

			getBody = []jen.Code{
				jen.Return(
					jen.Add(mappedType).Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Id("value"),
			}

			//setBody = []jen.Code{
			//	jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Call(jen.Id("*C." + field.Type)).Call(jen.Id("value")),
			//}

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

		if prim, ok := nativeTypes[field.Type]; ok {
			mappedType = jen.Qual(ffiPath, "Ref").Types(jen.Id(prim))

			getBody = []jen.Code{
				jen.Return(
					jen.Qual(ffiPath, "RefFromPtr").Types(jen.Id(prim)).Call(
						jen.Qual("unsafe", "Pointer").Call(
							jen.Id("p").Dot("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("*C." + field.Type)).
					Call(
						jen.Id("value").Dot("Raw").Call(),
					),
			}

			break
		}

		fieldType, ok := g.Repo.LookupType(field.Type)
		if !ok {
			g.OStructs.Commentf("%s.%s is unsupported: category %s -> ??.", ty.MappedName, field.Name, field.Category)
			g.OStructs.Line()

			return
		}

		switch fieldType.Category {
		case repo.TypeCategoryStruct:
			mappedType = jen.Id(fieldType.MappedName)

			getBody = []jen.Code{
				jen.Return(jen.Id(fieldType.MappedName).Values(
					jen.Id("ptr").Op(":").Id("p").Dot("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Id("value").Dot("ptr"),
			}
		case repo.TypeCategoryEnum, repo.TypeCategoryHandle, repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryBitmask:
			mappedType = jen.Qual(ffiPath, "Ref").Types(jen.Id(fieldType.MappedName))

			getBody = []jen.Code{
				jen.Return(
					jen.Qual(ffiPath, "RefFromPtr").Types(jen.Id(fieldType.MappedName)).Call(
						jen.Qual("unsafe", "Pointer").Call(
							jen.Id("p").Dot("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("*C." + field.Type)).
					Call(
						jen.Id("value").Dot("Raw").Call(),
					),
			}

		default:
			g.OStructs.Commentf("%s.%s is unsupported: category pointer %s -> %s.", ty.MappedName, field.Name, fieldType.Category, fieldType.Name)
			g.OStructs.Line()

			return

		}
	case repo.FieldCategoryPointer2:
		if field.Type == "char" {
			mappedType = jen.Qual(ffiPath, "Ref").Types(jen.Qual(ffiPath, "CString"))

			getBody = []jen.Code{
				jen.Return(
					jen.Qual(ffiPath, "RefFromPtr").Types(jen.Qual(ffiPath, "CString")).Call(
						jen.Qual("unsafe", "Pointer").Call(
							jen.Id("p").Dot("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("**C." + field.Type)).
					Call(
						jen.Id("value").Dot("Raw").Call(),
					),
			}

			break
		}

		g.OStructs.Commentf("%s.%s is unsupported: category %s.", ty.MappedName, field.Name, field.Category)
		g.OStructs.Line()

		return

	default:
		g.OStructs.Commentf("%s.%s is unsupported: category %s.", ty.MappedName, field.Name, field.Category)
		g.OStructs.Line()

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

	g.OStructs.Line()
	g.OStructs.Commentf("Get%s returns the value in %s.", field.MappedName, field.Name)
	g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Get" + field.MappedName).
		Params().
		Add(mappedType).
		Block(getBody...)

	g.OStructs.Line()
	g.OStructs.Commentf("Set%s sets the value in %s.", field.MappedName, field.Name)
	g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Set" + field.MappedName).
		Params(jen.Id("value").Add(mappedType)).
		Block(setBody...)
}

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

	g.OStructs.Commentf("%s wraps struct %s.", ty.MappedName, ty.Name)

	if ty.Comment != "" {
		g.OStructs.Comment(ty.Comment)
	}

	g.OStructs.Type().Id(ty.MappedName).Id("uintptr")

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

	if ty.SuppressAlloc {
		g.OStructs.Commentf("%s allocator is suppressed.", ty.Name)
		g.OStructs.Line()
	} else {
		g.OStructs.Commentf("%sAlloc allocates a continuous block of %s.", ty.MappedName, ty.MappedName)
		g.OStructs.Func().Id(ty.MappedName+"Alloc").
			Params(jen.Id("alloc").Qual(ffiPath, "Allocator"), jen.Id("count").Id("int")).
			Id(ty.MappedName).
			Block(
				jen.Id("ptr").Op(":=").Id("alloc").Dot("Allocate").Call(jen.Id(sizeOfName).Op("*").Id("count")),

				jen.Return(jen.Id(ty.MappedName).Call(jen.Id("ptr"))),
			)
	}

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

	g.OStructs.Comment("Offset returns an offset pointer by the size of the struct and the provided multiple.")
	g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Offset").
		Params(jen.Id("offset").Id("int")).
		Id(ty.MappedName).
		Block(
			jen.Return(jen.Id("p").Op("+").Id(ty.MappedName).Call(jen.Id("offset").Op("*").Id(sizeOfName))),
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
	"bool":     "bool",
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
					jen.Id("ptr").Dot(cFieldName).Op("!=").Id("0"),
				),
			}

			setBody = []jen.Code{
				jen.If(jen.Id("value")).Block(
					jen.Id("ptr").Dot(cFieldName).Op("=").Id("C.VkBool32").Parens(jen.Id("1")),
				).Else().Block(
					jen.Id("ptr").Dot(cFieldName).Op("=").Id("C.VkBool32").Parens(jen.Id("0")),
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
					jen.Qual("unsafe", "Pointer").Params(jen.Id("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").Parens(
					jen.Id("C." + field.Type)).Params(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))),
			}
			//}

			break
		}

		fieldType, fieldPkg, ok := g.LookupType(field.Type)
		if !ok {
			g.OStructs.Commentf("%s.%s is unsupported: unknown type %s.", ty.MappedName, field.Name, field.Type)
			g.OStructs.Line()

			return
		}

		switch fieldType.Category {
		case repo.TypeCategoryEnum, repo.TypeCategoryBitmask, repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryDirect:
			mappedType = jen.Qual(fieldPkg, fieldType.MappedName)
			generateDefault = true

		case repo.TypeCategoryHandle:
			mappedType = jen.Qual(fieldPkg, fieldType.MappedName)

			getBody = []jen.Code{
				jen.Return(
					jen.Add(mappedType).Params(
						jen.Qual("unsafe", "Pointer").Params(jen.Id("ptr").Dot(cFieldName)),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").Parens(
					jen.Id("C." + field.Type)).Params(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))),
			}

		case repo.TypeCategoryStruct:
			mappedType = jen.Qual(fieldPkg, fieldType.MappedName)

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
					jen.Return(
						jen.Qual(fieldPkg, fieldType.MappedName).Params(
							jen.Id("p").Op("+").Id(ty.MappedName).Call(jen.Qual("C", offsetName)),
						),
					),
				)

			return

		default:
			g.OStructs.Commentf("%s.%s is unsupported: direct category %s.", ty.MappedName, field.Name, fieldType.Category)
			g.OStructs.Line()

			return
		}

	case repo.FieldCategoryPointer:
		if field.Type == "void" {
			mappedType = jen.Id("uintptr")

			getBody = []jen.Code{
				jen.Return(
					jen.Add(mappedType).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id("ptr").Dot(cFieldName))),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").Qual("unsafe", "Pointer").Call(jen.Id("value")),
			}

			break
		}

		if field.Type == "CAMetalLayer" {
			mappedType = jen.Qual("unsafe", "Pointer")

			getBody = []jen.Code{
				jen.Return(
					jen.Add(mappedType).Params(jen.Id("ptr").Dot(cFieldName)),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").Id("value"),
			}

			break
		}

		if field.Type == "char" {
			mappedType = jen.Qual(ffiPath, "CString")

			getBody = []jen.Code{
				jen.Return(jen.Qual(ffiPath, "CStringFromPtr").Parens(
					jen.Parens(jen.Qual("unsafe", "Pointer")).Parens(
						jen.Id("ptr").Dot(cFieldName))),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").
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
							jen.Id("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("*C." + field.Type)).
					Call(
						jen.Id("value").Dot("Raw").Call(),
					),
			}

			break
		}

		fieldType, fieldPkg, ok := g.LookupType(field.Type)
		if !ok {
			g.OStructs.Commentf("%s.%s is unsupported: category %s -> ??.", ty.MappedName, field.Name, field.Category)
			g.OStructs.Line()

			return
		}

		switch fieldType.Category {
		case repo.TypeCategoryStruct:
			mappedType = jen.Qual(fieldPkg, fieldType.MappedName)

			getBody = []jen.Code{
				jen.Return(jen.Qual(fieldPkg, fieldType.MappedName).Params(
					jen.Qual("unsafe", "Pointer").Call(jen.Id("ptr").Dot(cFieldName))),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("*C." + field.Type)).
					Call(
						jen.Qual("unsafe", "Pointer").Call(jen.Id("value")),
					),
			}
		case repo.TypeCategoryEnum, repo.TypeCategoryHandle, repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryBitmask, repo.TypeCategoryDirect:
			mappedType = jen.Qual(ffiPath, "Ref").Types(jen.Qual(fieldPkg, fieldType.MappedName))

			getBody = []jen.Code{
				jen.Return(
					jen.Qual(ffiPath, "RefFromPtr").Types(jen.Qual(fieldPkg, fieldType.MappedName)).Call(
						jen.Qual("unsafe", "Pointer").Call(
							jen.Id("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").
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
							jen.Id("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("**C." + field.Type)).
					Call(
						jen.Id("value").Dot("Raw").Call(),
					),
			}

			break
		}

		fieldType, fieldPkg, ok := g.LookupType(field.Type)
		if !ok {
			g.OStructs.Commentf("%s.%s is unsupported: category %s -> ??.", ty.MappedName, field.Name, field.Category)
			g.OStructs.Line()

			return
		}

		switch fieldType.Category {
		case repo.TypeCategoryStruct:
			mappedType = jen.Qual(ffiPath, "Ref").Types(jen.Qual(fieldPkg, fieldType.MappedName))

			getBody = []jen.Code{
				jen.Return(
					jen.Qual(ffiPath, "RefFromPtr").Types(jen.Qual(fieldPkg, fieldType.MappedName)).Call(
						jen.Qual("unsafe", "Pointer").Call(
							jen.Id("ptr").Dot(cFieldName),
						),
					),
				),
			}

			setBody = []jen.Code{
				jen.Id("ptr").Dot(cFieldName).Op("=").
					Call(jen.Id("**C." + field.Type)).
					Call(
						jen.Id("value").Dot("Raw").Call(),
					),
			}

		default:
			g.OStructs.Commentf("%s.%s is unsupported: category pointer2 %s -> %s.", ty.MappedName, field.Name, fieldType.Category, fieldType.Name)
			g.OStructs.Line()

			return
		}
	default:
		g.OStructs.Commentf("%s.%s is unsupported: category %s.", ty.MappedName, field.Name, field.Category)
		g.OStructs.Line()

		return
	}

	if generateDefault {
		getBody = []jen.Code{
			jen.Return(
				jen.Add(mappedType).Params(jen.Id("ptr").Dot(cFieldName)),
			),
		}

		setBody = []jen.Code{
			jen.Id("ptr").Dot(cFieldName).Op("=").Parens(jen.Id("C." + field.Type)).Params(jen.Id("value")),
		}
	}

	g.OStructs.Line()

	if field.SuppressGet {
		g.OStructs.Commentf("%s.%s setter is suppressed.", ty.MappedName, field.Name)
		g.OStructs.Line()
	} else {
		var stmt []jen.Code

		stmt = append(stmt, jen.Id("ptr").Op(":=").Call(jen.Id("*C."+ty.Name)).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id("p"))))
		stmt = append(stmt, getBody...)

		g.OStructs.Commentf("Get%s returns the value in %s.", field.MappedName, field.Name)
		g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Get" + field.MappedName).
			Params().
			Add(mappedType).
			Block(stmt...)
	}

	g.OStructs.Line()

	if field.SuppressSet {
		g.OStructs.Commentf("%s.%s getter is suppressed.", ty.MappedName, field.Name)
		g.OStructs.Line()
	} else {
		var stmt []jen.Code

		stmt = append(stmt, jen.Id("ptr").Op(":=").Call(jen.Id("*C."+ty.Name)).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id("p"))))
		stmt = append(stmt, setBody...)

		g.OStructs.Commentf("Set%s sets the value in %s.", field.MappedName, field.Name)
		g.OStructs.Func().Params(jen.Id("p").Id(ty.MappedName)).Id("Set" + field.MappedName).
			Params(jen.Id("value").Add(mappedType)).
			Block(stmt...)
	}
}

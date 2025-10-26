package main

import (
	"fmt"
	"generator/repo"
	"go/token"
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
)

const (
	ffiPath = "github.com/csnewman/go-gfx/ffi"
	vkPath  = "github.com/csnewman/go-gfx/vk"
)

type Generator struct {
	Mod              *Module
	VK               *repo.Repo
	VMA              *repo.Repo
	GeneratedOffsets map[string]struct{}
}

func (g *Generator) generate() {
	oEnums := jen.NewFile("vma")
	oEnums.CgoPreamble(`#include "vma.h"`)

	oBitmask := jen.NewFile("vma")
	oBitmask.CgoPreamble(`#include "vma.h"`)

	for _, name := range slices.Sorted(maps.Keys(g.Mod.Enums)) {
		e := g.Mod.Enums[name]

		if strings.Contains(name, "FlagBits") {
			g.generateBitmaskType(oBitmask, e)
		} else {
			// Same impl for now.
			g.generateBitmaskType(oEnums, e)
		}
	}

	oHandles := jen.NewFile("vma")
	oHandles.CgoPreamble(`#include "vma.h"`)

	for _, name := range slices.Sorted(maps.Keys(g.Mod.Handles)) {
		ty := g.Mod.Handles[name]

		g.generateHandleType(oHandles, ty)
	}

	oStructs := jen.NewFile("vma")
	oStructs.CgoPreamble(`#include "vma.h"`)

	for _, name := range slices.Sorted(maps.Keys(g.Mod.Structs)) {
		ty := g.Mod.Structs[name]

		g.generateStructType(oStructs, ty)
	}

	oFuncs := jen.NewFile("vma")
	oFuncs.CgoPreamble(`#include "vma.h"`)

	for _, name := range slices.Sorted(maps.Keys(g.Mod.Functions)) {
		ty := g.Mod.Functions[name]

		g.generateFunction(oFuncs, ty)
	}

	if err := oEnums.Save("../../vma/enums.gen.go"); err != nil {
		panic(err)
	}

	if err := oBitmask.Save("../../vma/bitmasks.gen.go"); err != nil {
		panic(err)
	}

	if err := oHandles.Save("../../vma/handles.gen.go"); err != nil {
		panic(err)
	}

	if err := oStructs.Save("../../vma/structs.gen.go"); err != nil {
		panic(err)
	}

	if err := oFuncs.Save("../../vma/funcs.gen.go"); err != nil {
		panic(err)
	}
}

func convertEnumName(in string) string {
	name, ok := strings.CutPrefix(in, "Vma")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	// XXX: bit of a hack! Should properly map these.
	name = strings.ReplaceAll(name, "FlagBits", "Flags")

	return name
}

func (g *Generator) generateBitmaskType(o *jen.File, ty *Enum) {
	name := convertEnumName(ty.Name)

	slog.Info("Generating bitmask", "name", ty.Name)

	o.Commentf("%s wraps the bitmask %s.", name, ty.Name)

	if ty.Comment != "" {
		o.Comment(ty.Comment)
	}

	goType := "int32"

	o.Type().Id(name).Id(goType)

	o.Const().DefsFunc(func(group *jen.Group) {
		for _, v := range ty.Constants {
			if v.Name == "" {
				panic(fmt.Sprintf("bitmask does not have a name: %s", ty.Name))
			}

			itemName, ok := strings.CutPrefix(v.Name, "VMA_")
			if !ok {
				panic(fmt.Sprintf("bitmask item does not have prefix: %s", ty.Name))
			}

			group.Commentf("%s wraps %s.", itemName, v.Name)

			if v.Comment != "" {
				group.Comment(v.Comment)
			}

			group.Id(itemName).Id(name).Op("=").Qual("C", v.Name)
		}
	})

	o.Func().
		ParamsFunc(func(group *jen.Group) {
			group.Id("e").Id(name)
		}).Id("String").
		Params().
		Id("string").
		BlockFunc(func(group *jen.Group) {
			group.Return().
				Qual("fmt", "Sprintf").
				Call(
					jen.Lit(fmt.Sprintf("%s(%%b)", ty.Name)),
					jen.Id("e"),
				)
		})
	o.Line()
}

func convertStructName(in string) string {
	name, ok := strings.CutPrefix(in, "Vma")
	if !ok {
		panic(fmt.Sprintf("struct does not have prefix: %s", in))
	}

	return name
}

func (g *Generator) generateHandleType(o *jen.File, ty *Handle) {
	name := convertStructName(ty.Name)

	slog.Info("Generating handle", "name", ty.Name)
	o.Commentf("%s wraps %s.", name, ty.Name)

	cType := "uintptr"

	if ty.NonDispatchable {
		cType = "uint64"
	}

	o.Type().Id(name).Id(cType)

	o.Commentf("%s is a null pointer.", name+"Nil")
	o.Var().Id(name + "Nil").Id(name)
}

func (g *Generator) generateStructType(o *jen.File, ty *Struct) {
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
		Params(jen.Id("alloc").Qual(ffiPath, "Allocator"), jen.Id("count").Id("int")).
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
		g.generateStructField(o, name, ty, field)
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
}

func (g *Generator) generateStructField(o *jen.File, tyName string, ty *Struct, field *Field) {
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

	switch field.Type.Category {
	case TypeCategoryDirect:
		if mapping, ok := nativeTypes[field.Type.Name]; ok {
			mappedType = jen.Id(mapping)
			generateDefault = true

			break
		}

		if field.Type.Name == "VkBool32" {
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

		if mapping, ok := g.VK.Types[field.Type.Name]; ok {
			switch mapping.Category {
			case repo.TypeCategoryDirect, repo.TypeCategoryEnum, repo.TypeCategoryBitmask, repo.TypeCategoryHandleNonDispatchable:
				enumName := mapping.MappedName
				mappedType = jen.Qual(vkPath, enumName)
				generateDefault = true

			case repo.TypeCategoryHandle:
				enumName := mapping.MappedName
				mappedType = jen.Qual(vkPath, enumName)

				getBody = []jen.Code{
					jen.Return(
						jen.Add(mappedType).Params(
							jen.Qual("unsafe", "Pointer").Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
						),
					),
				}

				setBody = []jen.Code{
					jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Parens(
						jen.Id("C." + field.Type.Name)).Params(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))),
				}

			default:
				o.Commentf("%s.%s is unsupported: direct category vk %s.", tyName, field.Name, mapping.Category)
				o.Line()

				return

			}

			break
		}

		if mapping, ok := g.VMA.Types[field.Type.Name]; ok {
			switch mapping.Category {
			case repo.TypeCategoryDirect, repo.TypeCategoryEnum, repo.TypeCategoryBitmask, repo.TypeCategoryHandleNonDispatchable:
				enumName := mapping.MappedName
				mappedType = jen.Id(enumName)
				generateDefault = true

			case repo.TypeCategoryHandle:
				enumName := mapping.MappedName
				mappedType = jen.Id(enumName)

				getBody = []jen.Code{
					jen.Return(
						jen.Add(mappedType).Params(
							jen.Qual("unsafe", "Pointer").Params(jen.Id("p").Dot("ptr").Dot(cFieldName)),
						),
					),
				}

				setBody = []jen.Code{
					jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Parens(
						jen.Id("C." + field.Type.Name)).Params(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))),
				}

			case repo.TypeCategoryStruct:
				structName := mapping.MappedName
				mappedType = jen.Id(structName)

				offsetName := "offsetof_" + ty.Name + "_" + field.Name

				if _, ok := g.GeneratedOffsets[offsetName]; !ok {
					g.GeneratedOffsets[offsetName] = struct{}{}

					o.CgoPreamble("static const int " + offsetName + " = offsetof(" + ty.Name + ", " + field.Name + ");")
				}
				o.Line()
				o.Commentf("Ref%s returns pointer to the %s field.", fieldName, field.Name)
				o.Func().Params(jen.Id("p").Id(tyName)).Id("Ref" + fieldName).
					Params().
					Add(mappedType).
					Block(
						jen.Return(jen.Id(structName).Values(
							jen.Id("ptr").Op(":").
								Params(jen.Id("*C." + field.Type.Name)).
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
				o.Commentf("%s.%s is unsupported: direct category vma %s.", tyName, field.Name, mapping.Category)
				o.Line()

				return

			}

			break
		}

		o.Commentf("%s.%s is unsupported: direct category unknown type %s.", tyName, field.Name, field.Type.Name)
		o.Line()

		return

	case TypeCategoryPointer:
		if field.Type.Name == "void" {
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

		if field.Type.Name == "char" {
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

		if mapping, ok := g.VK.Types[field.Type.Name]; ok {
			switch mapping.Category {
			case repo.TypeCategoryDirect, repo.TypeCategoryEnum, repo.TypeCategoryBitmask, repo.TypeCategoryHandleNonDispatchable:
				innerType := jen.Qual(vkPath, mapping.MappedName)
				mappedType = jen.Qual(ffiPath, "Ref").Types(innerType)

				getBody = []jen.Code{
					jen.Return(
						jen.Qual(ffiPath, "RefFromPtr").Types(innerType).Call(
							jen.Qual("unsafe", "Pointer").Call(
								jen.Id("p").Dot("ptr").Dot(cFieldName),
							),
						),
					),
				}

				setBody = []jen.Code{
					jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").
						Call(jen.Id("*C." + field.Type.Name)).
						Call(
							jen.Id("value").Dot("Raw").Call(),
						),
				}

				break

			default:
				o.Commentf("%s.%s is unsupported: pointer category vk %s.", tyName, field.Name, mapping.Category)
				o.Line()

				return

			}

			break
		}

		o.Commentf("%s.%s is unsupported: pointer category unknown type %s.", tyName, field.Name, field.Type.Name)
		o.Line()

		return

	default:
		o.Commentf("%s.%s is unsupported: category %s.", tyName, field.Name, field.Type.Category)
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
			jen.Id("p").Dot("ptr").Dot(cFieldName).Op("=").Parens(jen.Id("C." + field.Type.Name)).Params(jen.Id("value")),
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

func convertFuncName(in string) string {
	name, ok := strings.CutPrefix(in, "vma")
	if !ok {
		panic(fmt.Sprintf("cmd does not have prefix: %s", in))
	}

	return name
}

func (g *Generator) generateFunction(o *jen.File, ty *Function) {
	funcName := convertFuncName(ty.Name)

	o.Line()

	var (
		tmpRetVar     = "ret"
		retMappedType jen.Code
		postCall      []jen.Code
	)

	if ty.ReturnType.Name == "void" {
		tmpRetVar = ""
	} else if mapping, ok := nativeTypes[ty.ReturnType.Name]; ok {
		retMappedType = jen.Id(mapping)
		postCall = append(postCall, jen.Return(jen.Id(mapping).Call(jen.Id(tmpRetVar))))
	} else if ty.ReturnType.Name == "VkBool32" {
		retMappedType = jen.Id("bool")
		postCall = append(
			postCall,
			jen.If(jen.Id(tmpRetVar).Op(">").Id("0")).
				Block(jen.Return(jen.True())).
				Else().
				Block(jen.Return(jen.False())),
		)
	} else if vkType, ok := g.VK.Types[ty.ReturnType.Name]; ok {
		switch vkType.Category {
		case repo.TypeCategoryEnum:
			retMappedType = jen.Qual(vkPath, vkType.MappedName)
			postCall = append(postCall, jen.Return(jen.Add(retMappedType).Call(jen.Id(tmpRetVar))))
		default:
			o.Commentf("%s is unsupported: result: unknown category %s.", ty.Name, vkType.Category)
			o.Line()

			return
		}
	} else {
		o.Commentf("%s is unsupported: result: unknown type %s.", ty.Name, ty.ReturnType.Name)
		o.Line()

		return
	}

	var retTypeBlock []jen.Code

	if retMappedType != nil {
		retTypeBlock = append(retTypeBlock, retMappedType)
	}

	var (
		paramsBlock []jen.Code
		callBlock   []jen.Code
		preCall     []jen.Code
	)

	for _, param := range ty.Args {
		safeName := param.Name

		if token.IsKeyword(safeName) {
			safeName += "_"
		}

		switch param.Type.Category {
		case TypeCategoryDirect:

			if mapping, ok := nativeTypes[param.Type.Name]; ok {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Id(mapping))
				callBlock = append(callBlock, jen.Qual("C", param.Type.Name).Call(jen.Id(safeName)))

				break
			}

			if mapping, ok := g.VMA.Types[param.Type.Name]; ok {
				switch mapping.Category {
				case repo.TypeCategoryHandle:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Id(mapping.MappedName))
					callBlock = append(callBlock, jen.Qual("C", param.Type.Name).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id(safeName))))

				case repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryDirect:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Id(mapping.MappedName))
					callBlock = append(callBlock, jen.Qual("C", param.Type.Name).Call(jen.Id(safeName)))

				default:
					o.Commentf("%s.%s is unsupported: param: unknown direct vma category %s.", ty.Name, param.Name, mapping.Category)
					o.Line()

					return
				}

				break
			}

			if mapping, ok := g.VK.Types[param.Type.Name]; ok {
				switch mapping.Category {
				case repo.TypeCategoryHandle:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(vkPath, mapping.MappedName))
					callBlock = append(callBlock, jen.Qual("C", param.Type.Name).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id(safeName))))

				case repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryDirect:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(vkPath, mapping.MappedName))
					callBlock = append(callBlock, jen.Qual("C", param.Type.Name).Call(jen.Id(safeName)))

				default:
					o.Commentf("%s.%s is unsupported: param: unknown direct vk category %s.", ty.Name, param.Name, mapping.Category)
					o.Line()

					return
				}

				break
			}

			o.Commentf("%s.%s is unsupported: param: unknown direct type %s.", ty.Name, param.Name, param.Type.Name)
			o.Line()

			return

		case TypeCategoryPointer:
			if param.Type.Name == "void" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual("unsafe", "Pointer"))
				callBlock = append(callBlock, jen.Id(safeName))

				break
			}

			if param.Type.Name == "char" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "CString"))

				callBlock = append(callBlock, jen.Call(jen.Id("*C.char")).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

				break
			}

			if mapping, ok := g.VMA.Types[param.Type.Name]; ok {
				switch mapping.Category {

				case repo.TypeCategoryStruct:

					paramsBlock = append(paramsBlock, jen.Id(safeName).Id(mapping.MappedName))
					callBlock = append(callBlock, jen.Id(safeName).Dot("ptr"))

				case repo.TypeCategoryEnum, repo.TypeCategoryHandle, repo.TypeCategoryBitmask:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Id(mapping.MappedName)))

					callBlock = append(callBlock, jen.Call(jen.Id("*C."+param.Type.Name)).
						Call(
							jen.Id(safeName).Dot("Raw").Call(),
						),
					)

				default:
					o.Commentf("%s.%s is unsupported: param: unknown pointer vma category %s.", ty.Name, param.Name, mapping.Category)
					o.Line()

					return
				}

				break
			}

			if mapping, ok := g.VK.Types[param.Type.Name]; ok {
				switch mapping.Category {

				case repo.TypeCategoryStruct:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(vkPath, mapping.MappedName))
					callBlock = append(callBlock, jen.Call(jen.Id("*C."+param.Type.Name)).
						Call(
							jen.Id(safeName).Dot("Raw").Call(),
						),
					)

				case repo.TypeCategoryEnum, repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryBitmask:
					paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Qual(vkPath, mapping.MappedName)))

					callBlock = append(callBlock, jen.Call(jen.Id("*C."+param.Type.Name)).
						Call(
							jen.Id(safeName).Dot("Raw").Call(),
						),
					)

				default:
					o.Commentf("%s.%s is unsupported: param: unknown pointer vk category %s.", ty.Name, param.Name, mapping.Category)
					o.Line()

					return
				}

				break
			}

			o.Commentf("%s.%s is unsupported: param: unknown pointer type %s.", ty.Name, param.Name, param.Type.Name)
			o.Line()

			return

		case TypeCategoryPointer2:
			if param.Type.Name == "void" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Qual("unsafe", "Pointer")))
				callBlock = append(callBlock, jen.Call(jen.Id("*unsafe.Pointer")).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

				break
			}

			o.Commentf("%s.%s is unsupported: param: unknown pointer2 type %s.", ty.Name, param.Name, param.Type.Name)
			o.Line()

			return

		default:
			o.Commentf("%s.%s is unsupported: param: category %s.", ty.Name, param.Name, param.Type.Category)
			o.Line()

			return
		}

	}

	//_ = paramsC
	//_ = paramCNames

	var body []jen.Code

	body = append(body, preCall...)

	if tmpRetVar != "" {
		body = append(
			body,
			jen.Id(tmpRetVar).Op(":=").Qual("C", ty.Name).Call(callBlock...),
			jen.Line(),
		)
	} else {
		body = append(body, jen.Qual("C", ty.Name).Call(callBlock...))
	}

	body = append(body, postCall...)

	o.Commentf("%s wraps %s.", funcName, ty.Name)
	o.Func().Id(funcName).Params(paramsBlock...).Add(retTypeBlock...).Block(
		body...,
	)
}

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

	oHandles := jen.NewFile("vk")
	oHandles.CgoPreamble(`#include "vulkan.h"`)

	oCommands := jen.NewFile("vk")
	oCommands.CgoPreamble(`#include "vulkan.h"`)

	sortedTypes := slices.Sorted(maps.Keys(reg.Types))

	generatedOffsets := make(map[string]struct{})

	for _, name := range sortedTypes {
		ty := reg.Types[name]

		if ty.Feature == "" {
			continue
		}

		switch ty.Category {
		case CategoryEnum:
			generateEnumType(reg, oEnums, ty)

		case CategoryBitmask:
			generateBitmaskType(reg, oBitmask, ty)

		case CategoryStruct:
			generateStructType(reg, generatedOffsets, oStructs, ty)

		case CategoryHandle:
			generateHandleType(reg, oHandles, ty)

		default:
			panic(fmt.Sprintf("unknown type %v", ty.Category))
		}
	}

	oCommands.Comment("Load attempts to load all vulkan functions.")
	oCommands.Func().Id("Load").Params(jen.Id("loader").Qual("unsafe", "Pointer")).Block(
		jen.Qual("C", "gfx_vkInit").Call(jen.Id("loader")),
	)

	sortedCmds := slices.Sorted(maps.Keys(reg.Commands))

	generatedCmds := make(map[string]struct{})

	for _, name := range sortedCmds {
		cmd := reg.Commands[name]

		if cmd.Feature == "" {
			continue
		}

		generateCommand(reg, oCommands, cmd, generatedCmds)
	}

	var loadBody strings.Builder

	for _, name := range slices.Sorted(maps.Keys(generatedCmds)) {
		loadBody.WriteString("    ptr_")
		loadBody.WriteString(name)
		loadBody.WriteString(" = (PFN_")
		loadBody.WriteString(name)
		loadBody.WriteString(") ptr_vkGetInstanceProcAddr(context, \"")
		loadBody.WriteString(name)
		loadBody.WriteString("\");\n")
	}

	oCommands.CgoPreamble(fmt.Sprintf(`PFN_vkGetInstanceProcAddr ptr_vkGetInstanceProcAddr;

void gfx_vkInit(void* loader) {
    ptr_vkGetInstanceProcAddr = (PFN_vkGetInstanceProcAddr) loader;
    void* context = NULL;

%s}`, loadBody.String()))

	if err := oEnums.Save("../../vk/enums.gen.go"); err != nil {
		panic(err)
	}

	if err := oBitmask.Save("../../vk/bitmasks.gen.go"); err != nil {
		panic(err)
	}

	if err := oStructs.Save("../../vk/structs.gen.go"); err != nil {
		panic(err)
	}

	if err := oHandles.Save("../../vk/handles.gen.go"); err != nil {
		panic(err)
	}

	if err := oCommands.Save("../../vk/commands.gen.go"); err != nil {
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

	// XXX: bit of a hack! Should properly map these.
	if removed, ok := strings.CutSuffix(name, "FlagBits"); ok {
		return removed + "Flags"
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

		o.Commentf("%s wraps the enum %s. An alias for %s.", aliasName, alias, name)

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

func generateBitmaskType(reg *Registry, o *jen.File, ty *Type) {
	name := convertEnumName(ty.Name)

	mapping, hasMapping := reg.Enums[ty.Requires]

	slog.Info("Generating bitmask", "name", ty.Name)

	o.Commentf("%s wraps the bitmask %s.", name, ty.Name)

	var goType string

	switch ty.BitmaskWidth {
	case 32:
		goType = "int32"
	case 64:
		goType = "int64"
	default:
		panic(fmt.Sprintf("bitmask bit width %d", ty.BitmaskWidth))
	}

	o.Type().Id(name).Id(goType)

	slices.Sort(ty.Aliases)

	for _, alias := range ty.Aliases {
		aliasName, ok := strings.CutPrefix(alias, "Vk")
		if !ok {
			panic(fmt.Sprintf("bitmask alias does not have prefix: %s", alias))
		}

		o.Commentf("%s wraps the bitmask %s. An alias for %s.", aliasName, alias, name)

		o.Type().Id(aliasName).Op("=").Id(name)
	}

	if hasMapping && len(mapping.Values) > 0 {
		o.Const().DefsFunc(func(group *jen.Group) {
			for _, v := range mapping.Values {
				if v.Name == "" {
					panic(fmt.Sprintf("bitmask does not have a name: %s", ty.Name))
				}

				itemName, ok := strings.CutPrefix(v.Name, "VK_")
				if !ok {
					panic(fmt.Sprintf("bitmask item does not have prefix: %s", ty.Name))
				}

				if v.Alias != "" {
					aliasName, ok := strings.CutPrefix(v.Alias, "VK_")
					if !ok {
						panic(fmt.Sprintf("bitmask alias does not have prefix: %s", ty.Name))
					}

					group.Commentf("%s wraps %s.", itemName, v.Name)

					if v.Deprecated != "" {
						group.Comment("")
						group.Commentf("Deprecated: Use %s instead.", aliasName)
					}

					group.Id(itemName).Id(name).Op("=").Id(aliasName)

					continue
				}

				var value string

				if v.Value != "" {
					value = v.Value
				} else if v.HasBitPos {
					value = fmt.Sprintf("1<<%d", v.BitPos)
				} else {
					panic(fmt.Sprintf("bitmask does not have a value: %s", ty.Name))
				}

				group.Commentf("%s wraps %s.", itemName, v.Name)
				group.Id(itemName).Id(name).Op("=").Id(value)
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
	name, ok := strings.CutPrefix(in, "Vk")
	if !ok {
		panic(fmt.Sprintf("struct does not have prefix: %s", in))
	}

	return name
}

func generateStructType(reg *Registry, generatedOffsets map[string]struct{}, o *jen.File, ty *Type) {
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
		generateStructField(reg, generatedOffsets, o, name, ty, field)
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

func generateStructField(reg *Registry, generatedOffsets map[string]struct{}, o *jen.File, tyName string, ty *Type, field StructField) {
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

		fieldType, ok := reg.Types[field.Type]
		if !ok {
			o.Commentf("%s.%s is unsupported: unknown type %s.", tyName, field.Name, field.Type)
			o.Line()

			return
		}

		switch fieldType.Category {
		case CategoryEnum, CategoryBitmask:
			enumName := convertEnumName(fieldType.Name)
			mappedType = jen.Id(enumName)
			generateDefault = true

		case CategoryHandle:
			enumName := convertStructName(fieldType.Name)
			mappedType = jen.Id(enumName)

			if fieldType.NonDispatchable {
				generateDefault = true
			} else {
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
			}

		case CategoryStruct:
			structName := convertStructName(fieldType.Name)
			mappedType = jen.Id(structName)

			offsetName := "offsetof_" + ty.Name + "_" + field.Name

			if _, ok := generatedOffsets[offsetName]; !ok {
				generatedOffsets[offsetName] = struct{}{}

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

		fieldType, ok := reg.Types[field.Type]
		if !ok {
			o.Commentf("%s.%s is unsupported: category %s -> ??.", tyName, field.Name, field.Category)
			o.Line()

			return
		}

		switch fieldType.Category {
		case CategoryStruct:
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
		case CategoryEnum, CategoryHandle, CategoryBitmask:
			enumName := convertEnumName(fieldType.Name)

			mappedType = jen.Qual(ffiPath, "Ref").Types(jen.Id(enumName))

			getBody = []jen.Code{
				jen.Return(
					jen.Qual(ffiPath, "RefFromPtr").Types(jen.Id(enumName)).Call(
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
			o.Commentf("%s.%s is unsupported: category pointer %s -> %s.", tyName, field.Name, fieldType.Category, fieldType.Name)
			o.Line()

			return

		}
	case FieldCategoryPointer2:
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

		o.Commentf("%s.%s is unsupported: category %s.", tyName, field.Name, field.Category)
		o.Line()

		return

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

func generateHandleType(reg *Registry, o *jen.File, ty *Type) {
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

func convertCommandName(in string) string {
	name, ok := strings.CutPrefix(in, "vk")
	if !ok {
		panic(fmt.Sprintf("cmd does not have prefix: %s", in))
	}

	return name
}

func generateCommand(reg *Registry, o *jen.File, cmd *Command, generatedCmds map[string]struct{}) {
	cmdName := convertCommandName(cmd.Name)

	o.Line()

	var (
		tmpRetVar     = "ret"
		retMappedType jen.Code
		retCType      = cmd.ReturnType
		postCall      []jen.Code
	)

	if cmd.ReturnType == "void" {
		tmpRetVar = ""
	} else if mapping, ok := nativeTypes[cmd.ReturnType]; ok {
		retMappedType = jen.Id(mapping)
		postCall = append(postCall, jen.Return(jen.Id(mapping).Call(jen.Id(tmpRetVar))))
	} else if cmd.ReturnType == "VkBool32" {
		retMappedType = jen.Id("bool")
		postCall = append(postCall, jen.Return(jen.Id("bool").Call(jen.Id(tmpRetVar))))
	} else if fieldType, ok := reg.Types[cmd.ReturnType]; ok {
		switch fieldType.Category {
		case CategoryEnum:
			retMappedType = jen.Id(convertEnumName(fieldType.Name))
			postCall = append(postCall, jen.Return(jen.Add(retMappedType).Call(jen.Id(tmpRetVar))))
		default:
			o.Commentf("%s is unsupported: unknown category %s.", cmd.Name, fieldType.Category)
			o.Line()

			return
		}
	} else {
		o.Commentf("%s is unsupported: unknown type %s.", cmd.Name, cmd.ReturnType)
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
	)

	var (
		paramsC     []string
		paramCNames []string

		preCall []jen.Code
	)

	for _, member := range cmd.Members {
		safeName := member.Name

		if token.IsKeyword(safeName) {
			safeName += "_"
		}

		paramCNames = append(paramCNames, member.Name)

		switch member.Category {
		case MemberCategoryDirect:
			paramsC = append(paramsC, member.Type+" "+member.Name)

			if mapping, ok := nativeTypes[member.Type]; ok {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Id(mapping))
				callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Id(safeName)))

				break
			}

			if member.Type == "VkBool32" {
				tmpName := "tmp_" + member.Name

				paramsBlock = append(paramsBlock, jen.Id(safeName).Id("bool"))

				preCall = append(
					preCall,
					jen.Id(tmpName).Op(":=").Id("0"),
					jen.Line(),
					jen.If(jen.Id(safeName)).Block(
						jen.Id(tmpName).Op("=").Id("1"),
					),
					jen.Line(),
				)

				callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Id(tmpName)))

				break
			}

			memberType, ok := reg.Types[member.Type]
			if !ok {
				o.Commentf("%s.%s is unsupported: unknown type %s.", cmd.Name, member.Name, member.Type)
				o.Line()

				return
			}

			switch memberType.Category {
			case CategoryEnum, CategoryBitmask:
				enumName := convertEnumName(memberType.Name)

				paramsBlock = append(paramsBlock, jen.Id(safeName).Id(enumName))
				callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Id(safeName)))

			case CategoryHandle:
				handleName := convertStructName(memberType.Name)
				paramsBlock = append(paramsBlock, jen.Id(safeName).Id(handleName))

				if memberType.NonDispatchable {
					callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Id(safeName)))
				} else {
					callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id(safeName))))
				}

			default:
				o.Commentf("%s.%s is unsupported: direct category %s.", cmd.Name, member.Name, memberType.Category)
				o.Line()

				return

			}
		case MemberCategoryPointer:
			paramsC = append(paramsC, member.Type+"* "+member.Name)

			if member.Type == "void" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual("unsafe", "Pointer"))
				callBlock = append(callBlock, jen.Id(safeName))

				break
			}

			if member.Type == "char" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "CString"))

				callBlock = append(callBlock, jen.Call(jen.Id("*C.char")).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

				break
			}

			if prim, ok := nativeTypes[member.Type]; ok {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Id(prim)))

				callBlock = append(callBlock, jen.Call(jen.Id("*C."+member.Type)).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

				break
			}

			memberType, ok := reg.Types[member.Type]
			if !ok {
				o.Commentf("%s.%s is unsupported: category %s -> ?? %s.", cmd.Name, member.Name, member.Category, member.Type)
				o.Line()

				return
			}

			switch memberType.Category {
			case CategoryStruct:
				structName := convertStructName(memberType.Name)

				paramsBlock = append(paramsBlock, jen.Id(safeName).Id(structName))
				callBlock = append(callBlock, jen.Id(safeName).Dot("ptr"))
			case CategoryEnum, CategoryHandle, CategoryBitmask:
				enumName := convertEnumName(memberType.Name)

				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Id(enumName)))

				callBlock = append(callBlock, jen.Call(jen.Id("*C."+member.Type)).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

			default:
				o.Commentf("%s.%s is unsupported: category %s -> %s.", cmd.Name, member.Name, member.Category, memberType.Category)
				o.Line()

				return
			}

		default:
			o.Commentf("%s.%s is unsupported: category %s.", cmd.Name, member.Name, member.Category)
			o.Line()

			return
		}
	}

	o.CgoPreamble(fmt.Sprintf(`PFN_%s ptr_%s;
VKAPI_ATTR %s VKAPI_CALL %s(%s) {
	return ptr_%s(%s);
}`, cmd.Name, cmd.Name, retCType, cmd.Name, strings.Join(paramsC, ", "), cmd.Name, strings.Join(paramCNames, ", ")))

	var body []jen.Code

	body = append(body, preCall...)

	if tmpRetVar != "" {
		body = append(
			body,
			jen.Id(tmpRetVar).Op(":=").Qual("C", cmd.Name).Call(callBlock...),
			jen.Line(),
		)
	} else {
		body = append(body, jen.Qual("C", cmd.Name).Call(callBlock...))
	}

	body = append(body, postCall...)

	o.Commentf("%s wraps %s.", cmdName, cmd.Name)
	o.Func().Id(cmdName).Params(paramsBlock...).Add(retTypeBlock...).Block(
		body...,
	)

	generatedCmds[cmd.Name] = struct{}{}

}

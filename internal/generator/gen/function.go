package gen

import (
	"generator/repo"
	"go/token"

	"github.com/dave/jennifer/jen"
)

func (g *Generator) generateFunction(fun *repo.Function) {
	g.OFunctions.Line()

	var (
		tmpRetVar     = "ret"
		retMappedType jen.Code
		postCall      []jen.Code
	)

	if fun.ReturnType == nil {
		tmpRetVar = ""
	} else {
		if fun.ReturnType.Category != repo.FieldCategoryDirect {
			g.OFunctions.Commentf("%s is unsupported: unknown return category %s.", fun.Name, fun.ReturnType.Category)
			g.OFunctions.Line()

			return
		}

		if mapping, ok := nativeTypes[fun.ReturnType.Type]; ok {
			retMappedType = jen.Id(mapping)
			postCall = append(postCall, jen.Return(jen.Id(mapping).Call(jen.Id(tmpRetVar))))
		} else if fun.ReturnType.Type == "VkBool32" {
			retMappedType = jen.Id("bool")
			postCall = append(
				postCall,
				jen.If(jen.Id(tmpRetVar).Op(">").Id("0")).
					Block(jen.Return(jen.True())).
					Else().
					Block(jen.Return(jen.False())),
			)
		} else if fun.ReturnType.Type == "PFN_vkVoidFunction" {
			retMappedType = jen.Qual("unsafe", "Pointer")
			postCall = append(postCall, jen.Return(jen.Qual("unsafe", "Pointer").Call(jen.Id(tmpRetVar))))
		} else if fieldType, pkgName, ok := g.LookupType(fun.ReturnType.Type); ok {
			switch fieldType.Category {
			case repo.TypeCategoryEnum:
				retMappedType = jen.Qual(pkgName, fieldType.MappedName)
				postCall = append(postCall, jen.Return(jen.Add(retMappedType).Call(jen.Id(tmpRetVar))))
			default:
				g.OFunctions.Commentf("%s is unsupported: unknown category %s.", fun.Name, fieldType.Category)
				g.OFunctions.Line()

				return
			}
		} else {
			g.OFunctions.Commentf("%s is unsupported: unknown return type %s.", fun.Name, fun.ReturnType.Type)
			g.OFunctions.Line()
			return
		}
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

	for _, member := range fun.Members {
		safeName := member.MappedName

		if token.IsKeyword(safeName) {
			safeName += "_"
		}

		switch member.Category {
		case repo.FieldCategoryDirect:
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

			memberType, memberPkg, ok := g.LookupType(member.Type)
			if !ok {
				g.OFunctions.Commentf("%s.%s is unsupported: unknown member direct type %s.", fun.Name, member.Name, member.Type)
				g.OFunctions.Line()

				return
			}

			switch memberType.Category {
			case repo.TypeCategoryEnum, repo.TypeCategoryBitmask, repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryDirect:
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(memberPkg, memberType.MappedName))
				callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Id(safeName)))

			case repo.TypeCategoryHandle:
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(memberPkg, memberType.MappedName))
				callBlock = append(callBlock, jen.Qual("C", member.Type).Call(jen.Qual("unsafe", "Pointer").Call(jen.Id(safeName))))

			default:
				g.OFunctions.Commentf("%s.%s is unsupported: direct category %s.", fun.Name, member.Name, memberType.Category)
				g.OFunctions.Line()

				return

			}
		case repo.FieldCategoryPointer:
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

			memberType, memberPkg, ok := g.LookupType(member.Type)
			if !ok {
				g.OFunctions.Commentf("%s.%s is unsupported: category %s -> ?? %s.", fun.Name, member.Name, member.Category, member.Type)
				g.OFunctions.Line()

				return
			}

			switch memberType.Category {
			case repo.TypeCategoryStruct:
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(memberPkg, memberType.MappedName))

				callBlock = append(callBlock, jen.Call(jen.Id("*C."+member.Type)).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

			case repo.TypeCategoryEnum, repo.TypeCategoryHandle, repo.TypeCategoryHandleNonDispatchable, repo.TypeCategoryBitmask, repo.TypeCategoryDirect:
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Qual(memberPkg, memberType.MappedName)))

				callBlock = append(callBlock, jen.Call(jen.Id("*C."+member.Type)).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

			default:
				g.OFunctions.Commentf("%s.%s is unsupported: category %s -> %s.", fun.Name, member.Name, member.Category, memberType.Category)
				g.OFunctions.Line()

				return
			}

		case repo.FieldCategoryPointer2:
			if member.Type == "void" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Qual("unsafe", "Pointer")))
				callBlock = append(callBlock, jen.Call(jen.Id("*unsafe.Pointer")).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

				break
			} else if member.Type == "char" {
				paramsBlock = append(paramsBlock, jen.Id(safeName).Qual(ffiPath, "Ref").Types(jen.Qual(ffiPath, "CString")))
				callBlock = append(callBlock, jen.Call(jen.Id("**C.char")).
					Call(
						jen.Id(safeName).Dot("Raw").Call(),
					),
				)

				break
			}

			g.OFunctions.Commentf("%s.%s is unsupported: category %s type %s.", fun.Name, member.Name, member.Category, member.Type)
			g.OFunctions.Line()

			return

		default:
			g.OFunctions.Commentf("%s.%s is unsupported: category %s.", fun.Name, member.Name, member.Category)
			g.OFunctions.Line()

			return
		}
	}

	var body []jen.Code

	body = append(body, preCall...)

	if tmpRetVar != "" {
		body = append(
			body,
			jen.Id(tmpRetVar).Op(":=").Qual("C", fun.Name).Call(callBlock...),
			jen.Line(),
		)
	} else {
		body = append(body, jen.Qual("C", fun.Name).Call(callBlock...))
	}

	body = append(body, postCall...)

	g.OFunctions.Commentf("%s wraps %s.", fun.MappedName, fun.Name)

	if fun.Comment != "" {
		g.OFunctions.Comment(fun.Comment)
	}

	g.OFunctions.Func().Id(fun.MappedName).Params(paramsBlock...).Add(retTypeBlock...).Block(
		body...,
	)
}

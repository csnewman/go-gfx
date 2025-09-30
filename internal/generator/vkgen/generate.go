package main

import (
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
)

func generate(reg *Registry) {
	oEnums := jen.NewFile("vk")
	oBitmask := jen.NewFile("vk")

	_ = oEnums
	_ = oBitmask

	sortedTypes := slices.Sorted(maps.Keys(reg.Types))

	for _, name := range sortedTypes {
		ty := reg.Types[name]

		switch ty.Category {
		case CategoryEnum:
			generateEnumType(reg, oEnums, ty)

		case CategoryBitmask:

		default:
			panic(fmt.Sprintf("unknown type %v", ty.Category))
		}
	}

	if err := oEnums.Save("../../vk/enums.go"); err != nil {
		panic(err)
	}
}

type ToStrEntry struct {
	Local string
	Name  string
}

func generateEnumType(reg *Registry, o *jen.File, ty *Type) {
	name, ok := strings.CutPrefix(ty.Name, "Vk")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", ty.Name))
	}

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

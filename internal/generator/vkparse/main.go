package vkparse

import (
	"fmt"
	"generator/repo"
	"log/slog"
	"maps"
	"slices"
	"strings"
)

func Parse(path string) *repo.Repo {
	r, err := parse(path)
	if err != nil {
		panic(err)
	}

	slog.Info("got", "r", r)

	out := &repo.Repo{
		Types: map[string]*repo.Type{
			"VkDeviceSize": {
				Name:       "VkDeviceSize",
				MappedName: "DeviceSize",
				Category:   repo.TypeCategoryDirect,
			},
			"VkDeviceAddress": {
				Name:       "VkDeviceAddress",
				MappedName: "DeviceAddress",
				Category:   repo.TypeCategoryDirect,
			},
			"VkSampleMask": {
				Name:       "VkSampleMask",
				MappedName: "SampleMask",
				Category:   repo.TypeCategoryDirect,
			},
		},
		Functions: make(map[string]*repo.Function),
	}

	for name, ty := range r.Types {
		if ty.Feature == "" {
			continue
		}

		converted := &repo.Type{
			Name:    name,
			Aliases: make(map[string]*repo.Alias),

			BitmaskWidth: ty.BitmaskWidth,
		}

		var nameFunc func(string) string

		switch ty.Category {
		case CategoryEnum:
			mapping, hasMapping := r.Enums[ty.Name]
			if hasMapping && mapping.Type == "bitmask" {
				continue
			}

			converted.Category = repo.TypeCategoryEnum
			converted.MappedName = convertEnumName(name)
			nameFunc = convertEnumName

			if hasMapping {
				convertEnumValues(converted, mapping)
			}
		case CategoryBitmask:
			converted.Category = repo.TypeCategoryBitmask
			converted.MappedName = convertEnumName(name)
			nameFunc = convertEnumName

			if mapping, hasMapping := r.Enums[ty.Requires]; hasMapping {
				if mapping.Type != "bitmask" {
					panic("bitmask enum must have type bitmask")
				}

				convertEnumValues(converted, mapping)

				converted.Aliases[ty.Requires] = &repo.Alias{
					Name:         ty.Requires,
					MappedName:   convertBitmaskName(ty.Requires),
					NoGeneration: true,
				}
			}

		case CategoryStruct:
			converted.Category = repo.TypeCategoryStruct
			converted.MappedName = convertStructName(name)
			nameFunc = convertStructName
			converted.StructFields = make(map[string]*repo.Field)

			for _, field := range ty.Fields {
				converted.StructFields[field.Name] = &repo.Field{
					Name:       field.Name,
					MappedName: strings.ToUpper(field.Name[:1]) + field.Name[1:],
					Type:       field.Type,
					Category:   repo.FieldCategory(field.Category),
				}
			}
		case CategoryHandle:
			converted.MappedName = convertStructName(name)
			nameFunc = convertStructName

			if ty.NonDispatchable {
				converted.Category = repo.TypeCategoryHandleNonDispatchable
			} else {
				converted.Category = repo.TypeCategoryHandle
			}
		default:
			panic(fmt.Sprintf("unknown type %v", ty.Category))
		}

		for _, alias := range ty.Aliases {
			converted.Aliases[alias] = &repo.Alias{
				Name:       alias,
				MappedName: nameFunc(alias),
			}
		}

		out.Types[name] = converted
	}

	for _, name := range slices.Sorted(maps.Keys(r.Commands)) {
		cmd := r.Commands[name]

		if cmd.Feature == "" {
			continue
		}

		converted := &repo.Function{
			Name:       cmd.Name,
			MappedName: convertCommandName(cmd.Name),
			Aliases:    make(map[string]*repo.Alias),
		}

		for _, alias := range cmd.Aliases {
			converted.Aliases[alias] = &repo.Alias{
				Name:       alias,
				MappedName: convertCommandName(alias),
			}
		}

		if cmd.ReturnType != "" && cmd.ReturnType != "void" {
			converted.ReturnType = &repo.Field{
				Category: repo.FieldCategoryDirect,
				Type:     cmd.ReturnType,
			}
		}

		for _, member := range cmd.Members {
			converted.Members = append(converted.Members, &repo.Field{
				Name:       member.Name,
				MappedName: member.Name,
				Category:   repo.FieldCategory(member.Category),
				Type:       member.Type,
			})
		}

		out.Functions[cmd.Name] = converted
	}

	out.PostProcess()

	return out
}

func convertEnumValues(converted *repo.Type, mapping *Enum) {
	if len(mapping.Values) == 0 {
		return
	}

	converted.EnumValues = make(map[string]*repo.EnumValue)

	aliases := make(map[string]string)

	for _, value := range mapping.Values {
		if value.Alias != "" {
			aliases[value.Name] = value.Alias

			continue
		}

		value := &repo.EnumValue{
			Name:       value.Name,
			MappedName: convertEnumItem(value.Name),
		}

		converted.EnumValues[value.Name] = value
	}

	for {
		changed := false

		for from, to := range aliases {
			if mapped, ok := aliases[to]; ok {
				aliases[from] = mapped
				changed = true
			}
		}

		if !changed {
			break
		}
	}

	for from, to := range aliases {
		value, ok := converted.EnumValues[to]
		if !ok {
			panic(fmt.Sprintf("enum value for alias %q not found for %q", to, from))
		}

		if value.Aliases == nil {
			value.Aliases = make(map[string]*repo.Alias)
		}

		value.Aliases[from] = &repo.Alias{
			Name:       from,
			MappedName: convertEnumItem(from),
		}
	}
}

func convertEnumName(in string) string {
	name, ok := strings.CutPrefix(in, "Vk")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	// XXX: bit of a hack! Should properly map these.
	name = strings.ReplaceAll(name, "FlagBits", "Flags")

	return name
}

func convertBitmaskName(in string) string {
	name, ok := strings.CutPrefix(in, "Vk")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	return name
}

func convertEnumItem(in string) string {
	name, ok := strings.CutPrefix(in, "VK_")
	if !ok {
		panic(fmt.Sprintf("item does not have prefix: %s", in))
	}

	return name
}

func convertStructName(in string) string {
	name, ok := strings.CutPrefix(in, "Vk")
	if !ok {
		panic(fmt.Sprintf("struct does not have prefix: %s", in))
	}

	return name
}

func convertCommandName(in string) string {
	name, ok := strings.CutPrefix(in, "vk")
	if !ok {
		panic(fmt.Sprintf("cmd does not have prefix: %s", in))
	}

	return name
}

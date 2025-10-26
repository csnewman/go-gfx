package main

import (
	"encoding/json"
	"fmt"
	"generator/repo"
	"log/slog"
	"os"
)

func main() {
	r, err := parse("../../thirdparty/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		panic(err)
	}

	slog.Info("got", "r", r)

	generate(r)

	out := &repo.Repo{
		Types: map[string]*repo.Type{
			"VkDeviceSize": {
				Name:       "VkDeviceSize",
				MappedName: "DeviceSize",
				Category:   repo.TypeCategoryDirect,
			},
		},
	}

	for name, ty := range r.Types {
		if ty.Feature == "" {
			continue
		}

		switch ty.Category {
		case CategoryEnum:
			mapping, hasMapping := r.Enums[ty.Name]
			if hasMapping && mapping.Type == "bitmask" {
				continue
			}

			out.Types[name] = &repo.Type{
				Name:       name,
				MappedName: convertEnumName(name),
				Category:   repo.TypeCategoryEnum,
			}

		case CategoryBitmask:
			out.Types[name] = &repo.Type{
				Name:       name,
				MappedName: convertEnumName(name),
				Category:   repo.TypeCategoryBitmask,
			}

		case CategoryStruct:
			out.Types[name] = &repo.Type{
				Name:       name,
				MappedName: convertStructName(name),
				Category:   repo.TypeCategoryStruct,
			}

		case CategoryHandle:
			if ty.NonDispatchable {
				out.Types[name] = &repo.Type{
					Name:       name,
					MappedName: convertStructName(name),
					Category:   repo.TypeCategoryHandleNonDispatchable,
				}
			} else {
				out.Types[name] = &repo.Type{
					Name:       name,
					MappedName: convertStructName(name),
					Category:   repo.TypeCategoryHandle,
				}
			}

		default:
			panic(fmt.Sprintf("unknown type %v", ty.Category))
		}
	}

	encRepo, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("./vulkan-repo.json", encRepo, 0666); err != nil {
		panic(err)
	}
}

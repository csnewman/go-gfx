package main

import (
	"generator/repo"
	"strings"
)

func main() {

	p := &Parser{
		mod: &Module{
			Enums:     make(map[string]*Enum),
			Handles:   make(map[string]*Handle),
			Structs:   make(map[string]*Struct),
			Functions: make(map[string]*Function),
		},
	}

	p.ParseFile("", "../../thirdparty/VulkanMemoryAllocator/include/vk_mem_alloc.h")

	vkRepo := repo.Load("vulkan-repo.json")

	rOut := &repo.Repo{
		Types: map[string]*repo.Type{},
	}

	for name := range p.mod.Enums {
		if strings.Contains(name, "FlagBits") {
			convName := strings.ReplaceAll(name, "FlagBits", "Flags")

			rOut.Types[name] = &repo.Type{
				Name:       convName,
				MappedName: convertEnumName(name),
				Category:   repo.TypeCategoryBitmask,
			}

			rOut.Types[convName] = &repo.Type{
				Name:       convName,
				MappedName: convertEnumName(name),
				Category:   repo.TypeCategoryBitmask,
			}
		} else {
			rOut.Types[name] = &repo.Type{
				Name:       name,
				MappedName: convertEnumName(name),
				Category:   repo.TypeCategoryEnum,
			}
		}
	}

	for name, ty := range p.mod.Handles {
		if ty.NonDispatchable {
			rOut.Types[name] = &repo.Type{
				Name:       name,
				MappedName: convertStructName(name),
				Category:   repo.TypeCategoryHandleNonDispatchable,
			}
		} else {
			rOut.Types[name] = &repo.Type{
				Name:       name,
				MappedName: convertStructName(name),
				Category:   repo.TypeCategoryHandle,
			}
		}
	}

	for name := range p.mod.Structs {
		rOut.Types[name] = &repo.Type{
			Name:       name,
			MappedName: convertStructName(name),
			Category:   repo.TypeCategoryStruct,
		}
	}

	g := &Generator{
		Mod:              p.mod,
		VK:               vkRepo,
		VMA:              rOut,
		GeneratedOffsets: make(map[string]struct{}),
	}

	g.generate()

}

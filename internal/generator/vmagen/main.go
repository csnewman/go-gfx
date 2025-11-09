package main

import (
	"fmt"
	"generator/cparse"
	"generator/gen"
	"generator/repo"
	"strings"
)

func main() {

	parsed := cparse.Parse(&cparse.ParseConfig{
		Path: "../../thirdparty/VulkanMemoryAllocator/include/vk_mem_alloc.h",

		ConvertEnumName: func(in string) *cparse.ConvertEnumOp {
			name, ok := strings.CutPrefix(in, "Vma")
			if !ok {
				panic(fmt.Sprintf("enum does not have prefix: %s", in))
			}

			// XXX: bit of a hack! Should properly map these.
			name = strings.ReplaceAll(name, "FlagBits", "Flags")

			return &cparse.ConvertEnumOp{
				Name: name,
			}
		},

		ConvertEnumItemName: func(in string) *cparse.ConvertEnumItemOp {
			name, ok := strings.CutPrefix(in, "VMA_")
			if !ok {
				panic(fmt.Sprintf("item does not have prefix: %s", in))
			}

			return &cparse.ConvertEnumItemOp{
				Name: name,
			}
		},

		ConvertStructName: func(in string) *cparse.ConvertStructOp {
			name, ok := strings.CutPrefix(in, "Vma")
			if !ok {
				panic(fmt.Sprintf("struct does not have prefix: %s", in))
			}

			return &cparse.ConvertStructOp{
				Name: name,
			}
		},

		ConvertFuncName: func(in string) *cparse.ConvertFuncOp {
			name, ok := strings.CutPrefix(in, "vma")
			if !ok {
				panic(fmt.Sprintf("cmd does not have prefix: %s", in))
			}

			return &cparse.ConvertFuncOp{
				Name: name,
			}
		},
	})

	parsed.Write("./vma-repo.json")

	vkRepo := repo.Load("vulkan-repo.json")

	gen.Generate(&gen.Config{
		Repo:      parsed,
		PKGName:   "vma",
		Path:      "../../vma",
		CPreamble: `#include "vma.h"`,

		Deps: map[string]*repo.Repo{
			"github.com/csnewman/go-gfx/vk": vkRepo,
		},
	})
}

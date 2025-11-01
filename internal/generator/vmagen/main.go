package main

import (
	"generator/cparse"
	"generator/gen"
	"generator/repo"
)

func main() {

	parsed := cparse.Parse(&cparse.ParseConfig{
		Path: "../../thirdparty/VulkanMemoryAllocator/include/vk_mem_alloc.h",
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

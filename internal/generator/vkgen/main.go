package main

import (
	"generator/gen"
	"generator/vkparse"
)

func main() {
	parsed := vkparse.Parse("../../thirdparty/Vulkan-Headers/registry/vk.xml")

	parsed.Write("./vulkan-repo.json")

	gen.Generate(&gen.Config{
		Repo:      parsed,
		PKGName:   "vk",
		Path:      "../../vk/",
		CPreamble: `#include "vulkan.h"`,
	})

	generateLoader(parsed)
}

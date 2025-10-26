package main

import (
	"encoding/json"
	"generator/gen"
	"generator/vkparse"
	"os"
)

func main() {
	parsed := vkparse.Parse("../../thirdparty/Vulkan-Headers/registry/vk.xml")

	encRepo, err := json.MarshalIndent(parsed, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("./vulkan-repo.json", encRepo, 0666); err != nil {
		panic(err)
	}

	gen.Generate(&gen.Config{
		Repo:      parsed,
		PKGName:   "vk",
		Path:      "../../vk/",
		CPreamble: `#include "vulkan.h"`,
	})

	generateLoader(parsed)
}

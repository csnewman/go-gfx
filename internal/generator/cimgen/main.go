package main

import (
	"generator/gen"
	"generator/imparse"
)

func main() {
	parsed := imparse.Parse("../../thirdparty/cimgui/output/")

	parsed.Write("./imgui-repo.json")

	gen.Generate(&gen.Config{
		Repo:      parsed,
		PKGName:   "imgui",
		Path:      "../../imgui",
		CPreamble: `#include "imgui_wrapper.h"`,
	})
}

package main

import (
	"log/slog"
)

func main() {
	r, err := parse("../../thirdparty/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		panic(err)
	}

	slog.Info("got", "r", r)
}

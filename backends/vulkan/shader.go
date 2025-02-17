package vulkan

import (
	"github.com/csnewman/go-gfx/gfx"
	"unsafe"
)

/*
#include "vulkan.h"
*/
import "C"

type Shader struct {
	shader C.VkShaderModule
}

func (g *Graphics) CreateShader(cfg gfx.ShaderConfig) (gfx.Shader, error) {
	var createInfo C.VkShaderModuleCreateInfo
	createInfo.sType = C.VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
	createInfo.codeSize = C.size_t(len(cfg.SPIRV))
	createInfo.pCode = (*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(cfg.SPIRV)))

	var shaderModule C.VkShaderModule

	if err := mapError(C.vkCreateShaderModule(g.device, &createInfo, nil, &shaderModule)); err != nil {
		return nil, err
	}

	return &Shader{
		shader: shaderModule,
	}, nil
}

type ShaderFunction struct {
	shader   *Shader
	function string
}

func (s *Shader) Function(name string) (gfx.ShaderFunction, error) {
	return &ShaderFunction{
		shader:   s,
		function: name,
	}, nil
}

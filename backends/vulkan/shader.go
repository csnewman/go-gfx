package vulkan

import (
	"unsafe"

	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
)

type Shader struct {
	shader vk.ShaderModule
}

func (g *Graphics) CreateShader(cfg gfx.ShaderConfig) (gfx.Shader, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	createInfo := vk.ShaderModuleCreateInfoAlloc(arena, 1)
	createInfo.SetSType(vk.STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO)
	createInfo.SetCodeSize(uintptr(len(cfg.SPIRV)))
	createInfo.SetPCode(ffi.RefFromPtr[uint32](unsafe.Pointer(unsafe.SliceData(cfg.SPIRV))))

	moduleRef := ffi.RefAlloc[vk.ShaderModule](arena, 1)

	if err := mapError(vk.CreateShaderModule(g.device, createInfo, vk.AllocationCallbacksNil, moduleRef)); err != nil {
		return nil, err
	}

	return &Shader{
		shader: moduleRef.Get(),
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

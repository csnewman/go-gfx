package vulkan

import (
	"fmt"

	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
)

func (g *Graphics) createDescriptors(arena *ffi.Arena) error {
	maxBindings := 10000 // TODO: replace

	poolSizes := vk.DescriptorPoolSizeAlloc(arena, 2)

	poolSizes.SetType(vk.DESCRIPTOR_TYPE_SAMPLER)
	poolSizes.SetDescriptorCount(uint32(maxBindings))

	size2 := poolSizes.Offset(1)
	size2.SetType(vk.DESCRIPTOR_TYPE_SAMPLED_IMAGE)
	size2.SetDescriptorCount(uint32(maxBindings))

	poolInfo := vk.DescriptorPoolCreateInfoAlloc(arena, 1)
	poolInfo.SetSType(vk.STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO)
	poolInfo.SetFlags(vk.DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT)
	poolInfo.SetMaxSets(2)
	poolInfo.SetPoolSizeCount(2)
	poolInfo.SetPPoolSizes(poolSizes)

	poolRef := ffi.RefAlloc[vk.DescriptorPool](arena, 1)

	if err := mapError(vk.CreateDescriptorPool(g.device, poolInfo, vk.AllocationCallbacksNil, poolRef)); err != nil {
		return fmt.Errorf("failed to create pool: %s", err)
	}

	pool := poolRef.Get()

	var err error

	g.samplerSet, g.samplerLayout, err = g.allocateDescriptorSet(
		arena,
		pool,
		vk.DESCRIPTOR_TYPE_SAMPLER,
		maxBindings,
	)
	if err != nil {
		return fmt.Errorf("failed to create set: %s", err)
	}

	g.textureSet, g.textureLayout, err = g.allocateDescriptorSet(
		arena,
		pool,
		vk.DESCRIPTOR_TYPE_SAMPLED_IMAGE,
		maxBindings,
	)
	if err != nil {
		return fmt.Errorf("failed to create set: %s", err)
	}

	pipelineLayoutInfo := vk.PipelineLayoutCreateInfoAlloc(arena, 1)
	pipelineLayoutInfo.SetSType(vk.STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO)
	pipelineLayoutInfo.SetSetLayoutCount(2)
	pipelineLayoutInfo.SetPSetLayouts(ffi.RefFromValues(arena,
		g.samplerLayout,
		g.textureLayout,
	))

	push := vk.PushConstantRangeAlloc(arena, 1)
	push.SetStageFlags(vk.SHADER_STAGE_VERTEX_BIT | vk.SHADER_STAGE_FRAGMENT_BIT)
	push.SetOffset(0)
	push.SetSize(8 * 4) // 4 ptrs

	pipelineLayoutInfo.SetPushConstantRangeCount(1)
	pipelineLayoutInfo.SetPPushConstantRanges(push)

	layoutRef := ffi.RefAlloc[vk.PipelineLayout](arena, 1)

	if err := mapError(vk.CreatePipelineLayout(g.device, pipelineLayoutInfo, vk.AllocationCallbacksNil, layoutRef)); err != nil {
		return fmt.Errorf("failed to create pipeline layout: %s", err)
	}

	g.pipelineLayout = layoutRef.Get()

	g.pipelineSets = []vk.DescriptorSet{
		g.samplerSet, g.textureSet,
	}

	return nil
}

func (g *Graphics) allocateDescriptorSet(
	arena *ffi.Arena,
	pool vk.DescriptorPool,
	desType vk.DescriptorType,
	maxBindings int,
) (vk.DescriptorSet, vk.DescriptorSetLayout, error) {
	bindFlags := vk.DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT |
		vk.DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT |
		vk.DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT

	extendedInfo := vk.DescriptorSetLayoutBindingFlagsCreateInfoAlloc(arena, 1)
	extendedInfo.SetSType(vk.STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO)
	extendedInfo.SetBindingCount(1)
	extendedInfo.SetPBindingFlags(ffi.RefFromValues(arena, bindFlags))

	info := vk.DescriptorSetLayoutCreateInfoAlloc(arena, 1)
	info.SetSType(vk.STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO)
	info.SetFlags(vk.DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT)
	info.SetPNext(uintptr(extendedInfo))

	binding := vk.DescriptorSetLayoutBindingAlloc(arena, 1)
	binding.SetDescriptorType(desType)
	binding.SetDescriptorCount(uint32(maxBindings))
	binding.SetStageFlags(vk.SHADER_STAGE_ALL)

	info.SetBindingCount(1)
	info.SetPBindings(binding)

	layoutRef := ffi.RefAlloc[vk.DescriptorSetLayout](arena, 1)

	if err := mapError(vk.CreateDescriptorSetLayout(g.device, info, vk.AllocationCallbacksNil, layoutRef)); err != nil {
		return vk.DescriptorSetNil, vk.DescriptorSetLayoutNil, err
	}

	layout := layoutRef.Get()

	countInfo := vk.DescriptorSetVariableDescriptorCountAllocateInfoAlloc(arena, 1)
	countInfo.SetSType(vk.STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO)
	countInfo.SetDescriptorSetCount(1)
	countInfo.SetPDescriptorCounts(ffi.RefFromValues(arena, uint32(maxBindings-1)))

	alloc := vk.DescriptorSetAllocateInfoAlloc(arena, 1)
	alloc.SetSType(vk.STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO)
	alloc.SetDescriptorPool(pool)
	alloc.SetDescriptorSetCount(1)
	alloc.SetPSetLayouts(layoutRef)
	alloc.SetPNext(uintptr(countInfo))

	setRef := ffi.RefAlloc[vk.DescriptorSet](arena, 1)

	if err := mapError(vk.AllocateDescriptorSets(g.device, alloc, setRef)); err != nil {
		return vk.DescriptorSetNil, vk.DescriptorSetLayoutNil, err
	}

	return setRef.Get(), layout, nil
}

type Sampler struct {
	id      uint32
	sampler vk.Sampler
}

func (g *Graphics) CreateSampler(des gfx.SamplerDescriptor) (gfx.Sampler, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	samplerInfo := vk.SamplerCreateInfoAlloc(arena, 1)
	samplerInfo.SetSType(vk.STRUCTURE_TYPE_SAMPLER_CREATE_INFO)
	samplerInfo.SetMagFilter(vk.FILTER_LINEAR)
	samplerInfo.SetMinFilter(vk.FILTER_LINEAR)
	samplerInfo.SetAddressModeU(vk.SAMPLER_ADDRESS_MODE_REPEAT)
	samplerInfo.SetAddressModeV(vk.SAMPLER_ADDRESS_MODE_REPEAT)
	samplerInfo.SetAddressModeW(vk.SAMPLER_ADDRESS_MODE_REPEAT)
	//samplerInfo.anisotropyEnable = C.VK_TRUE
	//samplerInfo.maxAnisotropy = properties.limits.maxSamplerAnisotropy
	samplerInfo.SetBorderColor(vk.BORDER_COLOR_INT_OPAQUE_BLACK)
	samplerInfo.SetUnnormalizedCoordinates(false)
	samplerInfo.SetCompareEnable(false)
	samplerInfo.SetCompareOp(vk.COMPARE_OP_ALWAYS)
	samplerInfo.SetMipmapMode(vk.SAMPLER_MIPMAP_MODE_LINEAR)

	samplerRef := ffi.RefAlloc[vk.Sampler](arena, 1)

	if err := mapError(vk.CreateSampler(g.device, samplerInfo, vk.AllocationCallbacksNil, samplerRef)); err != nil {
		return nil, err
	}

	sampler := samplerRef.Get()

	id, ok := g.samplerIDs.NextClear(g.lastSamplerID)
	if !ok {
		id = g.lastSamplerID + 1
	}

	g.samplerIDs.Set(id)

	g.lastSamplerID = max(g.lastSamplerID, id)

	write := vk.WriteDescriptorSetAlloc(arena, 1)
	write.SetSType(vk.STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET)
	write.SetDstSet(g.samplerSet)
	write.SetDstBinding(0)

	write.SetDstArrayElement(uint32(id))
	write.SetDescriptorCount(1)
	write.SetDescriptorType(vk.DESCRIPTOR_TYPE_SAMPLER)

	info := vk.DescriptorImageInfoAlloc(arena, 1)
	info.SetSampler(sampler)
	write.SetPImageInfo(info)

	vk.UpdateDescriptorSets(g.device, 1, write, 0, vk.CopyDescriptorSetNil)

	return &Sampler{
		id:      uint32(id),
		sampler: sampler,
	}, nil
}

func (s *Sampler) ID() uint32 {
	return s.id
}

func (g *Graphics) allocTexture(arena *ffi.Arena, view vk.ImageView) uint32 {
	id, ok := g.textureIDs.NextClear(g.lastTextureID)
	if !ok {
		id = g.lastTextureID + 1
	}

	g.textureIDs.Set(id)

	g.lastTextureID = max(g.lastTextureID, id)

	write := vk.WriteDescriptorSetAlloc(arena, 1)
	write.SetSType(vk.STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET)
	write.SetDstSet(g.textureSet)
	write.SetDstBinding(0)

	write.SetDstArrayElement(uint32(id))
	write.SetDescriptorCount(1)
	write.SetDescriptorType(vk.DESCRIPTOR_TYPE_SAMPLED_IMAGE)

	info := vk.DescriptorImageInfoAlloc(arena, 1)
	info.SetImageView(view)
	info.SetImageLayout(vk.IMAGE_LAYOUT_GENERAL)

	write.SetPImageInfo(info)

	vk.UpdateDescriptorSets(g.device, 1, write, 0, vk.CopyDescriptorSetNil)

	return uint32(id)
}

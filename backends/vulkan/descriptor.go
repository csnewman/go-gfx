package vulkan

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/csnewman/go-gfx/gfx"
)

/*
#include "vulkan.h"
*/
import "C"

func (g *Graphics) createDescriptors() error {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	maxBindings := 10000 // TODO: replace

	poolSizes := make([]C.VkDescriptorPoolSize, 2)
	poolSizes[0]._type = C.VK_DESCRIPTOR_TYPE_SAMPLER
	poolSizes[0].descriptorCount = C.uint32_t(maxBindings)

	poolSizes[1]._type = C.VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
	poolSizes[1].descriptorCount = C.uint32_t(maxBindings)

	var poolInfo C.VkDescriptorPoolCreateInfo
	poolInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
	poolInfo.flags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT
	poolInfo.maxSets = 2

	poolInfo.poolSizeCount = C.uint32_t(len(poolSizes))
	poolInfo.pPoolSizes = unsafe.SliceData(poolSizes)
	pinner.Pin(poolInfo.pPoolSizes)

	var pool C.VkDescriptorPool

	if err := mapError(C.vkCreateDescriptorPool(g.device, &poolInfo, nil, &pool)); err != nil {
		return fmt.Errorf("failed to create pool: %s", err)
	}

	var err error

	g.samplerSet, g.samplerLayout, err = g.allocateDescriptorSet(
		pinner,
		pool,
		C.VK_DESCRIPTOR_TYPE_SAMPLER,
		maxBindings,
	)
	if err != nil {
		return fmt.Errorf("failed to create set: %s", err)
	}

	g.textureSet, g.textureLayout, err = g.allocateDescriptorSet(
		pinner,
		pool,
		C.VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE,
		maxBindings,
	)
	if err != nil {
		return fmt.Errorf("failed to create set: %s", err)
	}

	var pipelineLayoutInfo C.VkPipelineLayoutCreateInfo
	pipelineLayoutInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO

	sets := make([]C.VkDescriptorSetLayout, 2)
	sets[0] = g.samplerLayout
	sets[1] = g.textureLayout

	pipelineLayoutInfo.setLayoutCount = C.uint32_t(len(sets))
	pipelineLayoutInfo.pSetLayouts = unsafe.SliceData(sets)
	pinner.Pin(pipelineLayoutInfo.pSetLayouts)

	var push C.VkPushConstantRange
	push.stageFlags = C.VK_SHADER_STAGE_VERTEX_BIT | C.VK_SHADER_STAGE_FRAGMENT_BIT
	push.offset = C.uint32_t(0)
	push.size = C.uint32_t(8 * 4) // 4 ptrs

	pipelineLayoutInfo.pushConstantRangeCount = C.uint32_t(1)
	pipelineLayoutInfo.pPushConstantRanges = &push
	pinner.Pin(pipelineLayoutInfo.pPushConstantRanges)

	var pipelineLayout C.VkPipelineLayout

	if err := mapError(C.vkCreatePipelineLayout(g.device, &pipelineLayoutInfo, nil, &pipelineLayout)); err != nil {
		return fmt.Errorf("failed to create pipeline layout: %s", err)
	}

	g.pipelineLayout = pipelineLayout

	g.pipelineSets = []C.VkDescriptorSet{
		g.samplerSet, g.textureSet,
	}

	return nil
}

func (g *Graphics) allocateDescriptorSet(
	pinner *runtime.Pinner,
	pool C.VkDescriptorPool,
	desType C.VkDescriptorType,
	maxBindings int,
) (C.VkDescriptorSet, C.VkDescriptorSetLayout, error) {
	bindFlags := C.VkDescriptorBindingFlags(C.VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT |
		C.VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT |
		C.VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT)

	var extendedInfo C.VkDescriptorSetLayoutBindingFlagsCreateInfo
	extendedInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO
	extendedInfo.bindingCount = 1
	extendedInfo.pBindingFlags = &bindFlags
	pinner.Pin(extendedInfo.pBindingFlags)

	var info C.VkDescriptorSetLayoutCreateInfo
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
	info.flags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT

	info.pNext = unsafe.Pointer(&extendedInfo)
	pinner.Pin(info.pNext)

	var binding C.VkDescriptorSetLayoutBinding
	binding.descriptorType = desType
	binding.descriptorCount = C.uint32_t(maxBindings)
	binding.stageFlags = C.VK_SHADER_STAGE_ALL

	info.bindingCount = C.uint32_t(1)
	info.pBindings = &binding
	pinner.Pin(info.pBindings)

	var layout C.VkDescriptorSetLayout

	if err := mapError(C.vkCreateDescriptorSetLayout(g.device, &info, nil, &layout)); err != nil {
		return nil, nil, err
	}

	var countInfo C.VkDescriptorSetVariableDescriptorCountAllocateInfoEXT
	countInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT

	descriptorCount := C.uint32_t(maxBindings - 1)
	countInfo.descriptorSetCount = 1
	countInfo.pDescriptorCounts = &descriptorCount
	pinner.Pin(countInfo.pDescriptorCounts)

	var alloc C.VkDescriptorSetAllocateInfo
	alloc.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
	alloc.descriptorPool = pool
	alloc.descriptorSetCount = 1
	alloc.pSetLayouts = &layout
	pinner.Pin(alloc.pSetLayouts)

	alloc.pNext = unsafe.Pointer(&countInfo)
	pinner.Pin(alloc.pNext)

	var set C.VkDescriptorSet

	if err := mapError(C.vkAllocateDescriptorSets(g.device, &alloc, &set)); err != nil {
		return nil, nil, err
	}

	return set, layout, nil
}

type Sampler struct {
	id      uint32
	sampler C.VkSampler
}

func (g *Graphics) CreateSampler(des gfx.SamplerDescriptor) (gfx.Sampler, error) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var samplerInfo C.VkSamplerCreateInfo
	samplerInfo.sType = C.VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO
	samplerInfo.magFilter = C.VK_FILTER_LINEAR
	samplerInfo.minFilter = C.VK_FILTER_LINEAR
	samplerInfo.addressModeU = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.addressModeV = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.addressModeW = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	//samplerInfo.anisotropyEnable = C.VK_TRUE
	//samplerInfo.maxAnisotropy = properties.limits.maxSamplerAnisotropy
	samplerInfo.borderColor = C.VK_BORDER_COLOR_INT_OPAQUE_BLACK
	samplerInfo.unnormalizedCoordinates = C.VK_FALSE
	samplerInfo.compareEnable = C.VK_FALSE
	samplerInfo.compareOp = C.VK_COMPARE_OP_ALWAYS
	samplerInfo.mipmapMode = C.VK_SAMPLER_MIPMAP_MODE_LINEAR

	var sampler C.VkSampler

	if err := mapError(C.vkCreateSampler(g.device, &samplerInfo, nil, &sampler)); err != nil {
		return nil, err
	}

	id, ok := g.samplerIDs.NextClear(g.lastSamplerID)
	if !ok {
		id = g.lastSamplerID + 1
	}

	g.samplerIDs.Set(id)

	g.lastSamplerID = id

	var write C.VkWriteDescriptorSet
	write.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	write.dstSet = g.samplerSet
	write.dstBinding = 0

	write.dstArrayElement = C.uint32_t(id)
	write.descriptorCount = 1
	write.descriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLER

	var info C.VkDescriptorImageInfo
	info.sampler = sampler

	write.pImageInfo = &info
	pinner.Pin(write.pImageInfo)

	C.vkUpdateDescriptorSets(g.device, 1, &write, 0, nil)

	return &Sampler{
		id:      uint32(id),
		sampler: sampler,
	}, nil
}

func (s *Sampler) ID() uint32 {
	return s.id
}

func (g *Graphics) allocTexture(pinner *runtime.Pinner, view C.VkImageView) uint32 {
	id, ok := g.textureIDs.NextClear(g.lastTextureID)
	if !ok {
		id = g.lastTextureID + 1
	}

	g.textureIDs.Set(id)

	g.lastTextureID = id

	var write C.VkWriteDescriptorSet
	write.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	write.dstSet = g.textureSet
	write.dstBinding = 0

	write.dstArrayElement = C.uint32_t(id)
	write.descriptorCount = 1
	write.descriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE

	var info C.VkDescriptorImageInfo
	info.imageView = view
	info.imageLayout = C.VK_IMAGE_LAYOUT_GENERAL

	write.pImageInfo = &info
	pinner.Pin(write.pImageInfo)

	C.vkUpdateDescriptorSets(g.device, 1, &write, 0, nil)

	return uint32(id)
}

package gfx

import (
	"runtime"
	"unsafe"
)

/*
#include "vulkan.h"
*/
import "C"

func (g *Graphics) createDescriptors() error {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	maxBindings := 50 // TODO: replace

	poolSizes := make([]C.VkDescriptorPoolSize, 1)
	poolSizes[0]._type = C.VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
	poolSizes[0].descriptorCount = C.uint32_t(maxBindings) // TODO: fix

	var poolInfo C.VkDescriptorPoolCreateInfo
	poolInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
	poolInfo.flags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT
	poolInfo.maxSets = 1

	poolInfo.poolSizeCount = C.uint32_t(len(poolSizes))
	poolInfo.pPoolSizes = unsafe.SliceData(poolSizes)
	pinner.Pin(poolInfo.pPoolSizes)

	pools := make([]C.VkDescriptorPool, len(poolSizes))

	if err := mapError(C.vkCreateDescriptorPool(g.device, &poolInfo, nil, unsafe.SliceData(pools))); err != nil {
		return err
	}

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
	binding.descriptorType = C.VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
	binding.descriptorCount = C.uint32_t(maxBindings)
	binding.stageFlags = C.VK_SHADER_STAGE_ALL

	info.bindingCount = C.uint32_t(1)
	info.pBindings = &binding
	pinner.Pin(info.pBindings)

	var layout C.VkDescriptorSetLayout

	if err := mapError(C.vkCreateDescriptorSetLayout(g.device, &info, nil, &layout)); err != nil {
		return err
	}

	var countInfo C.VkDescriptorSetVariableDescriptorCountAllocateInfoEXT
	countInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT

	descriptorCount := C.uint32_t(maxBindings - 1) // TODO: check
	countInfo.descriptorSetCount = 1
	countInfo.pDescriptorCounts = &descriptorCount
	pinner.Pin(countInfo.pDescriptorCounts)

	var alloc C.VkDescriptorSetAllocateInfo
	alloc.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
	alloc.descriptorPool = pools[0]
	alloc.descriptorSetCount = 1
	alloc.pSetLayouts = &layout
	pinner.Pin(alloc.pSetLayouts)

	alloc.pNext = unsafe.Pointer(&countInfo)
	pinner.Pin(alloc.pNext)

	var set C.VkDescriptorSet

	if err := mapError(C.vkAllocateDescriptorSets(g.device, &alloc, &set)); err != nil {
		return err
	}

	g.textureLayout = layout
	g.textureSet = set

	var pipelineLayoutInfo C.VkPipelineLayoutCreateInfo
	pipelineLayoutInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	pipelineLayoutInfo.setLayoutCount = 0
	pipelineLayoutInfo.pushConstantRangeCount = 0

	sets := make([]C.VkDescriptorSetLayout, 1)
	sets[0] = g.textureLayout

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
		return err
	}

	g.pipelineLayout = pipelineLayout

	return nil
}

func (g *Graphics) SetDescriptor(img *Image) error {
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

	var textureSampler C.VkSampler

	if err := mapError(C.vkCreateSampler(g.device, &samplerInfo, nil, &textureSampler)); err != nil {
		return err
	}

	var write C.VkWriteDescriptorSet
	write.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	write.dstSet = g.textureSet
	write.dstBinding = 0

	write.dstArrayElement = 0
	write.descriptorCount = 1
	write.descriptorType = C.VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER

	var info C.VkDescriptorImageInfo

	info.sampler = textureSampler
	info.imageView = img.view
	info.imageLayout = C.VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL

	write.pImageInfo = &info
	pinner.Pin(write.pImageInfo)

	C.vkUpdateDescriptorSets(g.device, 1, &write, 0, nil)

	return nil
}

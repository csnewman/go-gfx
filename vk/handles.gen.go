package vk

// #include "vulkan.h"
import "C"

// Buffer wraps VkBuffer.
type Buffer uint64

// BufferNil is a null pointer.
var BufferNil Buffer

// BufferView wraps VkBufferView.
type BufferView uint64

// BufferViewNil is a null pointer.
var BufferViewNil BufferView

// CommandBuffer wraps VkCommandBuffer.
type CommandBuffer uintptr

// CommandBufferNil is a null pointer.
var CommandBufferNil CommandBuffer

// CommandPool wraps VkCommandPool.
type CommandPool uint64

// CommandPoolNil is a null pointer.
var CommandPoolNil CommandPool

// DescriptorPool wraps VkDescriptorPool.
type DescriptorPool uint64

// DescriptorPoolNil is a null pointer.
var DescriptorPoolNil DescriptorPool

// DescriptorSet wraps VkDescriptorSet.
type DescriptorSet uint64

// DescriptorSetNil is a null pointer.
var DescriptorSetNil DescriptorSet

// DescriptorSetLayout wraps VkDescriptorSetLayout.
type DescriptorSetLayout uint64

// DescriptorSetLayoutNil is a null pointer.
var DescriptorSetLayoutNil DescriptorSetLayout

// DescriptorUpdateTemplate wraps VkDescriptorUpdateTemplate.
type DescriptorUpdateTemplate uint64

// DescriptorUpdateTemplateNil is a null pointer.
var DescriptorUpdateTemplateNil DescriptorUpdateTemplate

// Device wraps VkDevice.
type Device uintptr

// DeviceNil is a null pointer.
var DeviceNil Device

// DeviceMemory wraps VkDeviceMemory.
type DeviceMemory uint64

// DeviceMemoryNil is a null pointer.
var DeviceMemoryNil DeviceMemory

// Event wraps VkEvent.
type Event uint64

// EventNil is a null pointer.
var EventNil Event

// Fence wraps VkFence.
type Fence uint64

// FenceNil is a null pointer.
var FenceNil Fence

// Framebuffer wraps VkFramebuffer.
type Framebuffer uint64

// FramebufferNil is a null pointer.
var FramebufferNil Framebuffer

// Image wraps VkImage.
type Image uint64

// ImageNil is a null pointer.
var ImageNil Image

// ImageView wraps VkImageView.
type ImageView uint64

// ImageViewNil is a null pointer.
var ImageViewNil ImageView

// Instance wraps VkInstance.
type Instance uintptr

// InstanceNil is a null pointer.
var InstanceNil Instance

// PhysicalDevice wraps VkPhysicalDevice.
type PhysicalDevice uintptr

// PhysicalDeviceNil is a null pointer.
var PhysicalDeviceNil PhysicalDevice

// Pipeline wraps VkPipeline.
type Pipeline uint64

// PipelineNil is a null pointer.
var PipelineNil Pipeline

// PipelineCache wraps VkPipelineCache.
type PipelineCache uint64

// PipelineCacheNil is a null pointer.
var PipelineCacheNil PipelineCache

// PipelineLayout wraps VkPipelineLayout.
type PipelineLayout uint64

// PipelineLayoutNil is a null pointer.
var PipelineLayoutNil PipelineLayout

// PrivateDataSlot wraps VkPrivateDataSlot.
type PrivateDataSlot uint64

// PrivateDataSlotNil is a null pointer.
var PrivateDataSlotNil PrivateDataSlot

// QueryPool wraps VkQueryPool.
type QueryPool uint64

// QueryPoolNil is a null pointer.
var QueryPoolNil QueryPool

// Queue wraps VkQueue.
type Queue uintptr

// QueueNil is a null pointer.
var QueueNil Queue

// RenderPass wraps VkRenderPass.
type RenderPass uint64

// RenderPassNil is a null pointer.
var RenderPassNil RenderPass

// Sampler wraps VkSampler.
type Sampler uint64

// SamplerNil is a null pointer.
var SamplerNil Sampler

// SamplerYcbcrConversion wraps VkSamplerYcbcrConversion.
type SamplerYcbcrConversion uint64

// SamplerYcbcrConversionNil is a null pointer.
var SamplerYcbcrConversionNil SamplerYcbcrConversion

// Semaphore wraps VkSemaphore.
type Semaphore uint64

// SemaphoreNil is a null pointer.
var SemaphoreNil Semaphore

// ShaderModule wraps VkShaderModule.
type ShaderModule uint64

// ShaderModuleNil is a null pointer.
var ShaderModuleNil ShaderModule

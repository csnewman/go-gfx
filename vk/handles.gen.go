package vk

// #include "vulkan.h"
import "C"

// AccelerationStructureKHR wraps the handle VkAccelerationStructureKHR.
type AccelerationStructureKHR uint64

// AccelerationStructureKHRNil is a null pointer.
var AccelerationStructureKHRNil AccelerationStructureKHR

// AccelerationStructureNV wraps the handle VkAccelerationStructureNV.
type AccelerationStructureNV uint64

// AccelerationStructureNVNil is a null pointer.
var AccelerationStructureNVNil AccelerationStructureNV

// Buffer wraps the handle VkBuffer.
type Buffer uint64

// BufferNil is a null pointer.
var BufferNil Buffer

// BufferView wraps the handle VkBufferView.
type BufferView uint64

// BufferViewNil is a null pointer.
var BufferViewNil BufferView

// CommandBuffer wraps the handle VkCommandBuffer.
type CommandBuffer uintptr

// CommandBufferNil is a null pointer.
var CommandBufferNil CommandBuffer

// CommandPool wraps the handle VkCommandPool.
type CommandPool uint64

// CommandPoolNil is a null pointer.
var CommandPoolNil CommandPool

// CuFunctionNVX wraps the handle VkCuFunctionNVX.
type CuFunctionNVX uint64

// CuFunctionNVXNil is a null pointer.
var CuFunctionNVXNil CuFunctionNVX

// CuModuleNVX wraps the handle VkCuModuleNVX.
type CuModuleNVX uint64

// CuModuleNVXNil is a null pointer.
var CuModuleNVXNil CuModuleNVX

// DataGraphPipelineSessionARM wraps the handle VkDataGraphPipelineSessionARM.
type DataGraphPipelineSessionARM uint64

// DataGraphPipelineSessionARMNil is a null pointer.
var DataGraphPipelineSessionARMNil DataGraphPipelineSessionARM

// DebugReportCallbackEXT wraps the handle VkDebugReportCallbackEXT.
type DebugReportCallbackEXT uint64

// DebugReportCallbackEXTNil is a null pointer.
var DebugReportCallbackEXTNil DebugReportCallbackEXT

// DebugUtilsMessengerEXT wraps the handle VkDebugUtilsMessengerEXT.
type DebugUtilsMessengerEXT uint64

// DebugUtilsMessengerEXTNil is a null pointer.
var DebugUtilsMessengerEXTNil DebugUtilsMessengerEXT

// DeferredOperationKHR wraps the handle VkDeferredOperationKHR.
type DeferredOperationKHR uint64

// DeferredOperationKHRNil is a null pointer.
var DeferredOperationKHRNil DeferredOperationKHR

// DescriptorPool wraps the handle VkDescriptorPool.
type DescriptorPool uint64

// DescriptorPoolNil is a null pointer.
var DescriptorPoolNil DescriptorPool

// DescriptorSet wraps the handle VkDescriptorSet.
type DescriptorSet uint64

// DescriptorSetNil is a null pointer.
var DescriptorSetNil DescriptorSet

// DescriptorSetLayout wraps the handle VkDescriptorSetLayout.
type DescriptorSetLayout uint64

// DescriptorSetLayoutNil is a null pointer.
var DescriptorSetLayoutNil DescriptorSetLayout

// DescriptorUpdateTemplate wraps the handle VkDescriptorUpdateTemplate.
type DescriptorUpdateTemplate uint64

// DescriptorUpdateTemplateKHR wraps the handle VkDescriptorUpdateTemplateKHR. An alias for DescriptorUpdateTemplate.
type DescriptorUpdateTemplateKHR = DescriptorUpdateTemplate

// DescriptorUpdateTemplateNil is a null pointer.
var DescriptorUpdateTemplateNil DescriptorUpdateTemplate

// DescriptorUpdateTemplateKHRNil is a null pointer.
var DescriptorUpdateTemplateKHRNil DescriptorUpdateTemplateKHR

// Device wraps the handle VkDevice.
type Device uintptr

// DeviceNil is a null pointer.
var DeviceNil Device

// DeviceMemory wraps the handle VkDeviceMemory.
type DeviceMemory uint64

// DeviceMemoryNil is a null pointer.
var DeviceMemoryNil DeviceMemory

// DisplayKHR wraps the handle VkDisplayKHR.
type DisplayKHR uint64

// DisplayKHRNil is a null pointer.
var DisplayKHRNil DisplayKHR

// DisplayModeKHR wraps the handle VkDisplayModeKHR.
type DisplayModeKHR uint64

// DisplayModeKHRNil is a null pointer.
var DisplayModeKHRNil DisplayModeKHR

// Event wraps the handle VkEvent.
type Event uint64

// EventNil is a null pointer.
var EventNil Event

// ExternalComputeQueueNV wraps the handle VkExternalComputeQueueNV.
type ExternalComputeQueueNV uintptr

// ExternalComputeQueueNVNil is a null pointer.
var ExternalComputeQueueNVNil ExternalComputeQueueNV

// Fence wraps the handle VkFence.
type Fence uint64

// FenceNil is a null pointer.
var FenceNil Fence

// Framebuffer wraps the handle VkFramebuffer.
type Framebuffer uint64

// FramebufferNil is a null pointer.
var FramebufferNil Framebuffer

// Image wraps the handle VkImage.
type Image uint64

// ImageNil is a null pointer.
var ImageNil Image

// ImageView wraps the handle VkImageView.
type ImageView uint64

// ImageViewNil is a null pointer.
var ImageViewNil ImageView

// IndirectCommandsLayoutEXT wraps the handle VkIndirectCommandsLayoutEXT.
type IndirectCommandsLayoutEXT uint64

// IndirectCommandsLayoutEXTNil is a null pointer.
var IndirectCommandsLayoutEXTNil IndirectCommandsLayoutEXT

// IndirectCommandsLayoutNV wraps the handle VkIndirectCommandsLayoutNV.
type IndirectCommandsLayoutNV uint64

// IndirectCommandsLayoutNVNil is a null pointer.
var IndirectCommandsLayoutNVNil IndirectCommandsLayoutNV

// IndirectExecutionSetEXT wraps the handle VkIndirectExecutionSetEXT.
type IndirectExecutionSetEXT uint64

// IndirectExecutionSetEXTNil is a null pointer.
var IndirectExecutionSetEXTNil IndirectExecutionSetEXT

// Instance wraps the handle VkInstance.
type Instance uintptr

// InstanceNil is a null pointer.
var InstanceNil Instance

// MicromapEXT wraps the handle VkMicromapEXT.
type MicromapEXT uint64

// MicromapEXTNil is a null pointer.
var MicromapEXTNil MicromapEXT

// OpticalFlowSessionNV wraps the handle VkOpticalFlowSessionNV.
type OpticalFlowSessionNV uint64

// OpticalFlowSessionNVNil is a null pointer.
var OpticalFlowSessionNVNil OpticalFlowSessionNV

// PerformanceConfigurationINTEL wraps the handle VkPerformanceConfigurationINTEL.
type PerformanceConfigurationINTEL uint64

// PerformanceConfigurationINTELNil is a null pointer.
var PerformanceConfigurationINTELNil PerformanceConfigurationINTEL

// PhysicalDevice wraps the handle VkPhysicalDevice.
type PhysicalDevice uintptr

// PhysicalDeviceNil is a null pointer.
var PhysicalDeviceNil PhysicalDevice

// Pipeline wraps the handle VkPipeline.
type Pipeline uint64

// PipelineNil is a null pointer.
var PipelineNil Pipeline

// PipelineBinaryKHR wraps the handle VkPipelineBinaryKHR.
type PipelineBinaryKHR uint64

// PipelineBinaryKHRNil is a null pointer.
var PipelineBinaryKHRNil PipelineBinaryKHR

// PipelineCache wraps the handle VkPipelineCache.
type PipelineCache uint64

// PipelineCacheNil is a null pointer.
var PipelineCacheNil PipelineCache

// PipelineLayout wraps the handle VkPipelineLayout.
type PipelineLayout uint64

// PipelineLayoutNil is a null pointer.
var PipelineLayoutNil PipelineLayout

// PrivateDataSlot wraps the handle VkPrivateDataSlot.
type PrivateDataSlot uint64

// PrivateDataSlotEXT wraps the handle VkPrivateDataSlotEXT. An alias for PrivateDataSlot.
type PrivateDataSlotEXT = PrivateDataSlot

// PrivateDataSlotNil is a null pointer.
var PrivateDataSlotNil PrivateDataSlot

// PrivateDataSlotEXTNil is a null pointer.
var PrivateDataSlotEXTNil PrivateDataSlotEXT

// QueryPool wraps the handle VkQueryPool.
type QueryPool uint64

// QueryPoolNil is a null pointer.
var QueryPoolNil QueryPool

// Queue wraps the handle VkQueue.
type Queue uintptr

// QueueNil is a null pointer.
var QueueNil Queue

// RenderPass wraps the handle VkRenderPass.
type RenderPass uint64

// RenderPassNil is a null pointer.
var RenderPassNil RenderPass

// Sampler wraps the handle VkSampler.
type Sampler uint64

// SamplerNil is a null pointer.
var SamplerNil Sampler

// SamplerYcbcrConversion wraps the handle VkSamplerYcbcrConversion.
type SamplerYcbcrConversion uint64

// SamplerYcbcrConversionKHR wraps the handle VkSamplerYcbcrConversionKHR. An alias for SamplerYcbcrConversion.
type SamplerYcbcrConversionKHR = SamplerYcbcrConversion

// SamplerYcbcrConversionNil is a null pointer.
var SamplerYcbcrConversionNil SamplerYcbcrConversion

// SamplerYcbcrConversionKHRNil is a null pointer.
var SamplerYcbcrConversionKHRNil SamplerYcbcrConversionKHR

// Semaphore wraps the handle VkSemaphore.
type Semaphore uint64

// SemaphoreNil is a null pointer.
var SemaphoreNil Semaphore

// ShaderEXT wraps the handle VkShaderEXT.
type ShaderEXT uint64

// ShaderEXTNil is a null pointer.
var ShaderEXTNil ShaderEXT

// ShaderModule wraps the handle VkShaderModule.
type ShaderModule uint64

// ShaderModuleNil is a null pointer.
var ShaderModuleNil ShaderModule

// SurfaceKHR wraps the handle VkSurfaceKHR.
type SurfaceKHR uint64

// SurfaceKHRNil is a null pointer.
var SurfaceKHRNil SurfaceKHR

// SwapchainKHR wraps the handle VkSwapchainKHR.
type SwapchainKHR uint64

// SwapchainKHRNil is a null pointer.
var SwapchainKHRNil SwapchainKHR

// TensorARM wraps the handle VkTensorARM.
type TensorARM uint64

// TensorARMNil is a null pointer.
var TensorARMNil TensorARM

// TensorViewARM wraps the handle VkTensorViewARM.
type TensorViewARM uint64

// TensorViewARMNil is a null pointer.
var TensorViewARMNil TensorViewARM

// ValidationCacheEXT wraps the handle VkValidationCacheEXT.
type ValidationCacheEXT uint64

// ValidationCacheEXTNil is a null pointer.
var ValidationCacheEXTNil ValidationCacheEXT

// VideoSessionKHR wraps the handle VkVideoSessionKHR.
type VideoSessionKHR uint64

// VideoSessionKHRNil is a null pointer.
var VideoSessionKHRNil VideoSessionKHR

// VideoSessionParametersKHR wraps the handle VkVideoSessionParametersKHR.
type VideoSessionParametersKHR uint64

// VideoSessionParametersKHRNil is a null pointer.
var VideoSessionParametersKHRNil VideoSessionParametersKHR

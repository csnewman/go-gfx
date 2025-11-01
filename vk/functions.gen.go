package vk

import (
	ffi "github.com/csnewman/go-gfx/ffi"
	"unsafe"
)

// #include "vulkan.h"
import "C"

// AcquireDrmDisplayEXT wraps vkAcquireDrmDisplayEXT.
func AcquireDrmDisplayEXT(physicalDevice PhysicalDevice, drmFd int32, display DisplayKHR) Result {
	ret := C.vkAcquireDrmDisplayEXT(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.int32_t(drmFd), C.VkDisplayKHR(display))

	return Result(ret)
}

// AcquireFullScreenExclusiveModeEXT wraps vkAcquireFullScreenExclusiveModeEXT.
func AcquireFullScreenExclusiveModeEXT(device Device, swapchain SwapchainKHR) Result {
	ret := C.vkAcquireFullScreenExclusiveModeEXT(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain))

	return Result(ret)
}

// AcquireNextImage2KHR wraps vkAcquireNextImage2KHR.
func AcquireNextImage2KHR(device Device, pAcquireInfo AcquireNextImageInfoKHR, pImageIndex ffi.Ref[uint32]) Result {
	ret := C.vkAcquireNextImage2KHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkAcquireNextImageInfoKHR)(pAcquireInfo.Raw()), (*C.uint32_t)(pImageIndex.Raw()))

	return Result(ret)
}

// AcquireNextImageKHR wraps vkAcquireNextImageKHR.
func AcquireNextImageKHR(device Device, swapchain SwapchainKHR, timeout uint64, semaphore Semaphore, fence Fence, pImageIndex ffi.Ref[uint32]) Result {
	ret := C.vkAcquireNextImageKHR(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), C.uint64_t(timeout), C.VkSemaphore(semaphore), C.VkFence(fence), (*C.uint32_t)(pImageIndex.Raw()))

	return Result(ret)
}

// AcquirePerformanceConfigurationINTEL wraps vkAcquirePerformanceConfigurationINTEL.
func AcquirePerformanceConfigurationINTEL(device Device, pAcquireInfo PerformanceConfigurationAcquireInfoINTEL, pConfiguration ffi.Ref[PerformanceConfigurationINTEL]) Result {
	ret := C.vkAcquirePerformanceConfigurationINTEL(C.VkDevice(unsafe.Pointer(device)), (*C.VkPerformanceConfigurationAcquireInfoINTEL)(pAcquireInfo.Raw()), (*C.VkPerformanceConfigurationINTEL)(pConfiguration.Raw()))

	return Result(ret)
}

// AcquireProfilingLockKHR wraps vkAcquireProfilingLockKHR.
func AcquireProfilingLockKHR(device Device, pInfo AcquireProfilingLockInfoKHR) Result {
	ret := C.vkAcquireProfilingLockKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkAcquireProfilingLockInfoKHR)(pInfo.Raw()))

	return Result(ret)
}

// AcquireWinrtDisplayNV wraps vkAcquireWinrtDisplayNV.
func AcquireWinrtDisplayNV(physicalDevice PhysicalDevice, display DisplayKHR) Result {
	ret := C.vkAcquireWinrtDisplayNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkDisplayKHR(display))

	return Result(ret)
}

// AllocateCommandBuffers wraps vkAllocateCommandBuffers.
func AllocateCommandBuffers(device Device, pAllocateInfo CommandBufferAllocateInfo, pCommandBuffers ffi.Ref[CommandBuffer]) Result {
	ret := C.vkAllocateCommandBuffers(C.VkDevice(unsafe.Pointer(device)), (*C.VkCommandBufferAllocateInfo)(pAllocateInfo.Raw()), (*C.VkCommandBuffer)(pCommandBuffers.Raw()))

	return Result(ret)
}

// AllocateDescriptorSets wraps vkAllocateDescriptorSets.
func AllocateDescriptorSets(device Device, pAllocateInfo DescriptorSetAllocateInfo, pDescriptorSets ffi.Ref[DescriptorSet]) Result {
	ret := C.vkAllocateDescriptorSets(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorSetAllocateInfo)(pAllocateInfo.Raw()), (*C.VkDescriptorSet)(pDescriptorSets.Raw()))

	return Result(ret)
}

// AllocateMemory wraps vkAllocateMemory.
func AllocateMemory(device Device, pAllocateInfo MemoryAllocateInfo, pAllocator AllocationCallbacks, pMemory ffi.Ref[DeviceMemory]) Result {
	ret := C.vkAllocateMemory(C.VkDevice(unsafe.Pointer(device)), (*C.VkMemoryAllocateInfo)(pAllocateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDeviceMemory)(pMemory.Raw()))

	return Result(ret)
}

// AntiLagUpdateAMD wraps vkAntiLagUpdateAMD.
func AntiLagUpdateAMD(device Device, pData AntiLagDataAMD) {
	C.vkAntiLagUpdateAMD(C.VkDevice(unsafe.Pointer(device)), (*C.VkAntiLagDataAMD)(pData.Raw()))
}

// BeginCommandBuffer wraps vkBeginCommandBuffer.
func BeginCommandBuffer(commandBuffer CommandBuffer, pBeginInfo CommandBufferBeginInfo) Result {
	ret := C.vkBeginCommandBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCommandBufferBeginInfo)(pBeginInfo.Raw()))

	return Result(ret)
}

// BindAccelerationStructureMemoryNV wraps vkBindAccelerationStructureMemoryNV.
func BindAccelerationStructureMemoryNV(device Device, bindInfoCount uint32, pBindInfos BindAccelerationStructureMemoryInfoNV) Result {
	ret := C.vkBindAccelerationStructureMemoryNV(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), (*C.VkBindAccelerationStructureMemoryInfoNV)(pBindInfos.Raw()))

	return Result(ret)
}

// BindBufferMemory wraps vkBindBufferMemory.
func BindBufferMemory(device Device, buffer Buffer, memory DeviceMemory, memoryOffset DeviceSize) Result {
	ret := C.vkBindBufferMemory(C.VkDevice(unsafe.Pointer(device)), C.VkBuffer(buffer), C.VkDeviceMemory(memory), C.VkDeviceSize(memoryOffset))

	return Result(ret)
}

// BindBufferMemory2 wraps vkBindBufferMemory2.
func BindBufferMemory2(device Device, bindInfoCount uint32, pBindInfos BindBufferMemoryInfo) Result {
	ret := C.vkBindBufferMemory2(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), (*C.VkBindBufferMemoryInfo)(pBindInfos.Raw()))

	return Result(ret)
}

// BindDataGraphPipelineSessionMemoryARM wraps vkBindDataGraphPipelineSessionMemoryARM.
func BindDataGraphPipelineSessionMemoryARM(device Device, bindInfoCount uint32, pBindInfos BindDataGraphPipelineSessionMemoryInfoARM) Result {
	ret := C.vkBindDataGraphPipelineSessionMemoryARM(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), (*C.VkBindDataGraphPipelineSessionMemoryInfoARM)(pBindInfos.Raw()))

	return Result(ret)
}

// BindImageMemory wraps vkBindImageMemory.
func BindImageMemory(device Device, image Image, memory DeviceMemory, memoryOffset DeviceSize) Result {
	ret := C.vkBindImageMemory(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), C.VkDeviceMemory(memory), C.VkDeviceSize(memoryOffset))

	return Result(ret)
}

// BindImageMemory2 wraps vkBindImageMemory2.
func BindImageMemory2(device Device, bindInfoCount uint32, pBindInfos BindImageMemoryInfo) Result {
	ret := C.vkBindImageMemory2(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), (*C.VkBindImageMemoryInfo)(pBindInfos.Raw()))

	return Result(ret)
}

// BindOpticalFlowSessionImageNV wraps vkBindOpticalFlowSessionImageNV.
func BindOpticalFlowSessionImageNV(device Device, session OpticalFlowSessionNV, bindingPoint OpticalFlowSessionBindingPointNV, view ImageView, layout ImageLayout) Result {
	ret := C.vkBindOpticalFlowSessionImageNV(C.VkDevice(unsafe.Pointer(device)), C.VkOpticalFlowSessionNV(session), C.VkOpticalFlowSessionBindingPointNV(bindingPoint), C.VkImageView(view), C.VkImageLayout(layout))

	return Result(ret)
}

// BindTensorMemoryARM wraps vkBindTensorMemoryARM.
func BindTensorMemoryARM(device Device, bindInfoCount uint32, pBindInfos BindTensorMemoryInfoARM) Result {
	ret := C.vkBindTensorMemoryARM(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), (*C.VkBindTensorMemoryInfoARM)(pBindInfos.Raw()))

	return Result(ret)
}

// BindVideoSessionMemoryKHR wraps vkBindVideoSessionMemoryKHR.
func BindVideoSessionMemoryKHR(device Device, videoSession VideoSessionKHR, bindSessionMemoryInfoCount uint32, pBindSessionMemoryInfos BindVideoSessionMemoryInfoKHR) Result {
	ret := C.vkBindVideoSessionMemoryKHR(C.VkDevice(unsafe.Pointer(device)), C.VkVideoSessionKHR(videoSession), C.uint32_t(bindSessionMemoryInfoCount), (*C.VkBindVideoSessionMemoryInfoKHR)(pBindSessionMemoryInfos.Raw()))

	return Result(ret)
}

// vkBuildAccelerationStructuresKHR.ppBuildRangeInfos is unsupported: category pointer2 type VkAccelerationStructureBuildRangeInfoKHR.

// BuildMicromapsEXT wraps vkBuildMicromapsEXT.
func BuildMicromapsEXT(device Device, deferredOperation DeferredOperationKHR, infoCount uint32, pInfos MicromapBuildInfoEXT) Result {
	ret := C.vkBuildMicromapsEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), C.uint32_t(infoCount), (*C.VkMicromapBuildInfoEXT)(pInfos.Raw()))

	return Result(ret)
}

// CmdBeginConditionalRenderingEXT wraps vkCmdBeginConditionalRenderingEXT.
func CmdBeginConditionalRenderingEXT(commandBuffer CommandBuffer, pConditionalRenderingBegin ConditionalRenderingBeginInfoEXT) {
	C.vkCmdBeginConditionalRenderingEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkConditionalRenderingBeginInfoEXT)(pConditionalRenderingBegin.Raw()))
}

// CmdBeginDebugUtilsLabelEXT wraps vkCmdBeginDebugUtilsLabelEXT.
func CmdBeginDebugUtilsLabelEXT(commandBuffer CommandBuffer, pLabelInfo DebugUtilsLabelEXT) {
	C.vkCmdBeginDebugUtilsLabelEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDebugUtilsLabelEXT)(pLabelInfo.Raw()))
}

// CmdBeginPerTileExecutionQCOM wraps vkCmdBeginPerTileExecutionQCOM.
func CmdBeginPerTileExecutionQCOM(commandBuffer CommandBuffer, pPerTileBeginInfo PerTileBeginInfoQCOM) {
	C.vkCmdBeginPerTileExecutionQCOM(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPerTileBeginInfoQCOM)(pPerTileBeginInfo.Raw()))
}

// CmdBeginQuery wraps vkCmdBeginQuery.
func CmdBeginQuery(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, flags QueryControlFlags) {
	C.vkCmdBeginQuery(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(query), C.VkQueryControlFlags(flags))
}

// CmdBeginQueryIndexedEXT wraps vkCmdBeginQueryIndexedEXT.
func CmdBeginQueryIndexedEXT(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, flags QueryControlFlags, index uint32) {
	C.vkCmdBeginQueryIndexedEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(query), C.VkQueryControlFlags(flags), C.uint32_t(index))
}

// CmdBeginRenderPass wraps vkCmdBeginRenderPass.
func CmdBeginRenderPass(commandBuffer CommandBuffer, pRenderPassBegin RenderPassBeginInfo, contents SubpassContents) {
	C.vkCmdBeginRenderPass(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkRenderPassBeginInfo)(pRenderPassBegin.Raw()), C.VkSubpassContents(contents))
}

// CmdBeginRenderPass2 wraps vkCmdBeginRenderPass2.
func CmdBeginRenderPass2(commandBuffer CommandBuffer, pRenderPassBegin RenderPassBeginInfo, pSubpassBeginInfo SubpassBeginInfo) {
	C.vkCmdBeginRenderPass2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkRenderPassBeginInfo)(pRenderPassBegin.Raw()), (*C.VkSubpassBeginInfo)(pSubpassBeginInfo.Raw()))
}

// CmdBeginRendering wraps vkCmdBeginRendering.
func CmdBeginRendering(commandBuffer CommandBuffer, pRenderingInfo RenderingInfo) {
	C.vkCmdBeginRendering(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkRenderingInfo)(pRenderingInfo.Raw()))
}

// CmdBeginTransformFeedbackEXT wraps vkCmdBeginTransformFeedbackEXT.
func CmdBeginTransformFeedbackEXT(commandBuffer CommandBuffer, firstCounterBuffer uint32, counterBufferCount uint32, pCounterBuffers ffi.Ref[Buffer], pCounterBufferOffsets ffi.Ref[DeviceSize]) {
	C.vkCmdBeginTransformFeedbackEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstCounterBuffer), C.uint32_t(counterBufferCount), (*C.VkBuffer)(pCounterBuffers.Raw()), (*C.VkDeviceSize)(pCounterBufferOffsets.Raw()))
}

// CmdBeginVideoCodingKHR wraps vkCmdBeginVideoCodingKHR.
func CmdBeginVideoCodingKHR(commandBuffer CommandBuffer, pBeginInfo VideoBeginCodingInfoKHR) {
	C.vkCmdBeginVideoCodingKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkVideoBeginCodingInfoKHR)(pBeginInfo.Raw()))
}

// CmdBindDescriptorBufferEmbeddedSamplers2EXT wraps vkCmdBindDescriptorBufferEmbeddedSamplers2EXT.
func CmdBindDescriptorBufferEmbeddedSamplers2EXT(commandBuffer CommandBuffer, pBindDescriptorBufferEmbeddedSamplersInfo BindDescriptorBufferEmbeddedSamplersInfoEXT) {
	C.vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkBindDescriptorBufferEmbeddedSamplersInfoEXT)(pBindDescriptorBufferEmbeddedSamplersInfo.Raw()))
}

// CmdBindDescriptorBufferEmbeddedSamplersEXT wraps vkCmdBindDescriptorBufferEmbeddedSamplersEXT.
func CmdBindDescriptorBufferEmbeddedSamplersEXT(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, set uint32) {
	C.vkCmdBindDescriptorBufferEmbeddedSamplersEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipelineLayout(layout), C.uint32_t(set))
}

// CmdBindDescriptorBuffersEXT wraps vkCmdBindDescriptorBuffersEXT.
func CmdBindDescriptorBuffersEXT(commandBuffer CommandBuffer, bufferCount uint32, pBindingInfos DescriptorBufferBindingInfoEXT) {
	C.vkCmdBindDescriptorBuffersEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(bufferCount), (*C.VkDescriptorBufferBindingInfoEXT)(pBindingInfos.Raw()))
}

// CmdBindDescriptorSets wraps vkCmdBindDescriptorSets.
func CmdBindDescriptorSets(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, firstSet uint32, descriptorSetCount uint32, pDescriptorSets ffi.Ref[DescriptorSet], dynamicOffsetCount uint32, pDynamicOffsets ffi.Ref[uint32]) {
	C.vkCmdBindDescriptorSets(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipelineLayout(layout), C.uint32_t(firstSet), C.uint32_t(descriptorSetCount), (*C.VkDescriptorSet)(pDescriptorSets.Raw()), C.uint32_t(dynamicOffsetCount), (*C.uint32_t)(pDynamicOffsets.Raw()))
}

// CmdBindDescriptorSets2 wraps vkCmdBindDescriptorSets2.
func CmdBindDescriptorSets2(commandBuffer CommandBuffer, pBindDescriptorSetsInfo BindDescriptorSetsInfo) {
	C.vkCmdBindDescriptorSets2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkBindDescriptorSetsInfo)(pBindDescriptorSetsInfo.Raw()))
}

// CmdBindIndexBuffer wraps vkCmdBindIndexBuffer.
func CmdBindIndexBuffer(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, indexType IndexType) {
	C.vkCmdBindIndexBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkIndexType(indexType))
}

// CmdBindIndexBuffer2 wraps vkCmdBindIndexBuffer2.
func CmdBindIndexBuffer2(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, size DeviceSize, indexType IndexType) {
	C.vkCmdBindIndexBuffer2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkDeviceSize(size), C.VkIndexType(indexType))
}

// CmdBindInvocationMaskHUAWEI wraps vkCmdBindInvocationMaskHUAWEI.
func CmdBindInvocationMaskHUAWEI(commandBuffer CommandBuffer, imageView ImageView, imageLayout ImageLayout) {
	C.vkCmdBindInvocationMaskHUAWEI(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImageView(imageView), C.VkImageLayout(imageLayout))
}

// CmdBindPipeline wraps vkCmdBindPipeline.
func CmdBindPipeline(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	C.vkCmdBindPipeline(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipeline(pipeline))
}

// CmdBindPipelineShaderGroupNV wraps vkCmdBindPipelineShaderGroupNV.
func CmdBindPipelineShaderGroupNV(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline, groupIndex uint32) {
	C.vkCmdBindPipelineShaderGroupNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipeline(pipeline), C.uint32_t(groupIndex))
}

// CmdBindShadersEXT wraps vkCmdBindShadersEXT.
func CmdBindShadersEXT(commandBuffer CommandBuffer, stageCount uint32, pStages ffi.Ref[ShaderStageFlags], pShaders ffi.Ref[ShaderEXT]) {
	C.vkCmdBindShadersEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(stageCount), (*C.VkShaderStageFlagBits)(pStages.Raw()), (*C.VkShaderEXT)(pShaders.Raw()))
}

// CmdBindShadingRateImageNV wraps vkCmdBindShadingRateImageNV.
func CmdBindShadingRateImageNV(commandBuffer CommandBuffer, imageView ImageView, imageLayout ImageLayout) {
	C.vkCmdBindShadingRateImageNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImageView(imageView), C.VkImageLayout(imageLayout))
}

// CmdBindTileMemoryQCOM wraps vkCmdBindTileMemoryQCOM.
func CmdBindTileMemoryQCOM(commandBuffer CommandBuffer, pTileMemoryBindInfo TileMemoryBindInfoQCOM) {
	C.vkCmdBindTileMemoryQCOM(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkTileMemoryBindInfoQCOM)(pTileMemoryBindInfo.Raw()))
}

// CmdBindTransformFeedbackBuffersEXT wraps vkCmdBindTransformFeedbackBuffersEXT.
func CmdBindTransformFeedbackBuffersEXT(commandBuffer CommandBuffer, firstBinding uint32, bindingCount uint32, pBuffers ffi.Ref[Buffer], pOffsets ffi.Ref[DeviceSize], pSizes ffi.Ref[DeviceSize]) {
	C.vkCmdBindTransformFeedbackBuffersEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstBinding), C.uint32_t(bindingCount), (*C.VkBuffer)(pBuffers.Raw()), (*C.VkDeviceSize)(pOffsets.Raw()), (*C.VkDeviceSize)(pSizes.Raw()))
}

// CmdBindVertexBuffers wraps vkCmdBindVertexBuffers.
func CmdBindVertexBuffers(commandBuffer CommandBuffer, firstBinding uint32, bindingCount uint32, pBuffers ffi.Ref[Buffer], pOffsets ffi.Ref[DeviceSize]) {
	C.vkCmdBindVertexBuffers(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstBinding), C.uint32_t(bindingCount), (*C.VkBuffer)(pBuffers.Raw()), (*C.VkDeviceSize)(pOffsets.Raw()))
}

// CmdBindVertexBuffers2 wraps vkCmdBindVertexBuffers2.
func CmdBindVertexBuffers2(commandBuffer CommandBuffer, firstBinding uint32, bindingCount uint32, pBuffers ffi.Ref[Buffer], pOffsets ffi.Ref[DeviceSize], pSizes ffi.Ref[DeviceSize], pStrides ffi.Ref[DeviceSize]) {
	C.vkCmdBindVertexBuffers2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstBinding), C.uint32_t(bindingCount), (*C.VkBuffer)(pBuffers.Raw()), (*C.VkDeviceSize)(pOffsets.Raw()), (*C.VkDeviceSize)(pSizes.Raw()), (*C.VkDeviceSize)(pStrides.Raw()))
}

// CmdBlitImage wraps vkCmdBlitImage.
func CmdBlitImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions ImageBlit, filter Filter) {
	C.vkCmdBlitImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), (*C.VkImageBlit)(pRegions.Raw()), C.VkFilter(filter))
}

// CmdBlitImage2 wraps vkCmdBlitImage2.
func CmdBlitImage2(commandBuffer CommandBuffer, pBlitImageInfo BlitImageInfo2) {
	C.vkCmdBlitImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkBlitImageInfo2)(pBlitImageInfo.Raw()))
}

// CmdBuildAccelerationStructureNV wraps vkCmdBuildAccelerationStructureNV.
func CmdBuildAccelerationStructureNV(commandBuffer CommandBuffer, pInfo AccelerationStructureInfoNV, instanceData Buffer, instanceOffset DeviceSize, update bool, dst AccelerationStructureNV, src AccelerationStructureNV, scratch Buffer, scratchOffset DeviceSize) {
	tmp_update := 0

	if update {
		tmp_update = 1
	}

	C.vkCmdBuildAccelerationStructureNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkAccelerationStructureInfoNV)(pInfo.Raw()), C.VkBuffer(instanceData), C.VkDeviceSize(instanceOffset), C.VkBool32(tmp_update), C.VkAccelerationStructureNV(dst), C.VkAccelerationStructureNV(src), C.VkBuffer(scratch), C.VkDeviceSize(scratchOffset))
}

// vkCmdBuildAccelerationStructuresIndirectKHR.ppMaxPrimitiveCounts is unsupported: category pointer2 type uint32_t.

// vkCmdBuildAccelerationStructuresKHR.ppBuildRangeInfos is unsupported: category pointer2 type VkAccelerationStructureBuildRangeInfoKHR.

// CmdBuildClusterAccelerationStructureIndirectNV wraps vkCmdBuildClusterAccelerationStructureIndirectNV.
func CmdBuildClusterAccelerationStructureIndirectNV(commandBuffer CommandBuffer, pCommandInfos ClusterAccelerationStructureCommandsInfoNV) {
	C.vkCmdBuildClusterAccelerationStructureIndirectNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkClusterAccelerationStructureCommandsInfoNV)(pCommandInfos.Raw()))
}

// CmdBuildMicromapsEXT wraps vkCmdBuildMicromapsEXT.
func CmdBuildMicromapsEXT(commandBuffer CommandBuffer, infoCount uint32, pInfos MicromapBuildInfoEXT) {
	C.vkCmdBuildMicromapsEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(infoCount), (*C.VkMicromapBuildInfoEXT)(pInfos.Raw()))
}

// CmdBuildPartitionedAccelerationStructuresNV wraps vkCmdBuildPartitionedAccelerationStructuresNV.
func CmdBuildPartitionedAccelerationStructuresNV(commandBuffer CommandBuffer, pBuildInfo BuildPartitionedAccelerationStructureInfoNV) {
	C.vkCmdBuildPartitionedAccelerationStructuresNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkBuildPartitionedAccelerationStructureInfoNV)(pBuildInfo.Raw()))
}

// CmdClearAttachments wraps vkCmdClearAttachments.
func CmdClearAttachments(commandBuffer CommandBuffer, attachmentCount uint32, pAttachments ClearAttachment, rectCount uint32, pRects ClearRect) {
	C.vkCmdClearAttachments(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(attachmentCount), (*C.VkClearAttachment)(pAttachments.Raw()), C.uint32_t(rectCount), (*C.VkClearRect)(pRects.Raw()))
}

// vkCmdClearColorImage.pColor is unsupported: category pointer -> ?? VkClearColorValue.

// CmdClearDepthStencilImage wraps vkCmdClearDepthStencilImage.
func CmdClearDepthStencilImage(commandBuffer CommandBuffer, image Image, imageLayout ImageLayout, pDepthStencil ClearDepthStencilValue, rangeCount uint32, pRanges ImageSubresourceRange) {
	C.vkCmdClearDepthStencilImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(image), C.VkImageLayout(imageLayout), (*C.VkClearDepthStencilValue)(pDepthStencil.Raw()), C.uint32_t(rangeCount), (*C.VkImageSubresourceRange)(pRanges.Raw()))
}

// CmdControlVideoCodingKHR wraps vkCmdControlVideoCodingKHR.
func CmdControlVideoCodingKHR(commandBuffer CommandBuffer, pCodingControlInfo VideoCodingControlInfoKHR) {
	C.vkCmdControlVideoCodingKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkVideoCodingControlInfoKHR)(pCodingControlInfo.Raw()))
}

// CmdConvertCooperativeVectorMatrixNV wraps vkCmdConvertCooperativeVectorMatrixNV.
func CmdConvertCooperativeVectorMatrixNV(commandBuffer CommandBuffer, infoCount uint32, pInfos ConvertCooperativeVectorMatrixInfoNV) {
	C.vkCmdConvertCooperativeVectorMatrixNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(infoCount), (*C.VkConvertCooperativeVectorMatrixInfoNV)(pInfos.Raw()))
}

// CmdCopyAccelerationStructureKHR wraps vkCmdCopyAccelerationStructureKHR.
func CmdCopyAccelerationStructureKHR(commandBuffer CommandBuffer, pInfo CopyAccelerationStructureInfoKHR) {
	C.vkCmdCopyAccelerationStructureKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyAccelerationStructureInfoKHR)(pInfo.Raw()))
}

// CmdCopyAccelerationStructureNV wraps vkCmdCopyAccelerationStructureNV.
func CmdCopyAccelerationStructureNV(commandBuffer CommandBuffer, dst AccelerationStructureNV, src AccelerationStructureNV, mode CopyAccelerationStructureModeKHR) {
	C.vkCmdCopyAccelerationStructureNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkAccelerationStructureNV(dst), C.VkAccelerationStructureNV(src), C.VkCopyAccelerationStructureModeKHR(mode))
}

// CmdCopyAccelerationStructureToMemoryKHR wraps vkCmdCopyAccelerationStructureToMemoryKHR.
func CmdCopyAccelerationStructureToMemoryKHR(commandBuffer CommandBuffer, pInfo CopyAccelerationStructureToMemoryInfoKHR) {
	C.vkCmdCopyAccelerationStructureToMemoryKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyAccelerationStructureToMemoryInfoKHR)(pInfo.Raw()))
}

// CmdCopyBuffer wraps vkCmdCopyBuffer.
func CmdCopyBuffer(commandBuffer CommandBuffer, srcBuffer Buffer, dstBuffer Buffer, regionCount uint32, pRegions BufferCopy) {
	C.vkCmdCopyBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(srcBuffer), C.VkBuffer(dstBuffer), C.uint32_t(regionCount), (*C.VkBufferCopy)(pRegions.Raw()))
}

// CmdCopyBuffer2 wraps vkCmdCopyBuffer2.
func CmdCopyBuffer2(commandBuffer CommandBuffer, pCopyBufferInfo CopyBufferInfo2) {
	C.vkCmdCopyBuffer2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyBufferInfo2)(pCopyBufferInfo.Raw()))
}

// CmdCopyBufferToImage wraps vkCmdCopyBufferToImage.
func CmdCopyBufferToImage(commandBuffer CommandBuffer, srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions BufferImageCopy) {
	C.vkCmdCopyBufferToImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(srcBuffer), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), (*C.VkBufferImageCopy)(pRegions.Raw()))
}

// CmdCopyBufferToImage2 wraps vkCmdCopyBufferToImage2.
func CmdCopyBufferToImage2(commandBuffer CommandBuffer, pCopyBufferToImageInfo CopyBufferToImageInfo2) {
	C.vkCmdCopyBufferToImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyBufferToImageInfo2)(pCopyBufferToImageInfo.Raw()))
}

// CmdCopyImage wraps vkCmdCopyImage.
func CmdCopyImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions ImageCopy) {
	C.vkCmdCopyImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), (*C.VkImageCopy)(pRegions.Raw()))
}

// CmdCopyImage2 wraps vkCmdCopyImage2.
func CmdCopyImage2(commandBuffer CommandBuffer, pCopyImageInfo CopyImageInfo2) {
	C.vkCmdCopyImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyImageInfo2)(pCopyImageInfo.Raw()))
}

// CmdCopyImageToBuffer wraps vkCmdCopyImageToBuffer.
func CmdCopyImageToBuffer(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regionCount uint32, pRegions BufferImageCopy) {
	C.vkCmdCopyImageToBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkBuffer(dstBuffer), C.uint32_t(regionCount), (*C.VkBufferImageCopy)(pRegions.Raw()))
}

// CmdCopyImageToBuffer2 wraps vkCmdCopyImageToBuffer2.
func CmdCopyImageToBuffer2(commandBuffer CommandBuffer, pCopyImageToBufferInfo CopyImageToBufferInfo2) {
	C.vkCmdCopyImageToBuffer2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyImageToBufferInfo2)(pCopyImageToBufferInfo.Raw()))
}

// CmdCopyMemoryIndirectKHR wraps vkCmdCopyMemoryIndirectKHR.
func CmdCopyMemoryIndirectKHR(commandBuffer CommandBuffer, pCopyMemoryIndirectInfo CopyMemoryIndirectInfoKHR) {
	C.vkCmdCopyMemoryIndirectKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyMemoryIndirectInfoKHR)(pCopyMemoryIndirectInfo.Raw()))
}

// CmdCopyMemoryIndirectNV wraps vkCmdCopyMemoryIndirectNV.
func CmdCopyMemoryIndirectNV(commandBuffer CommandBuffer, copyBufferAddress DeviceAddress, copyCount uint32, stride uint32) {
	C.vkCmdCopyMemoryIndirectNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDeviceAddress(copyBufferAddress), C.uint32_t(copyCount), C.uint32_t(stride))
}

// CmdCopyMemoryToAccelerationStructureKHR wraps vkCmdCopyMemoryToAccelerationStructureKHR.
func CmdCopyMemoryToAccelerationStructureKHR(commandBuffer CommandBuffer, pInfo CopyMemoryToAccelerationStructureInfoKHR) {
	C.vkCmdCopyMemoryToAccelerationStructureKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyMemoryToAccelerationStructureInfoKHR)(pInfo.Raw()))
}

// CmdCopyMemoryToImageIndirectKHR wraps vkCmdCopyMemoryToImageIndirectKHR.
func CmdCopyMemoryToImageIndirectKHR(commandBuffer CommandBuffer, pCopyMemoryToImageIndirectInfo CopyMemoryToImageIndirectInfoKHR) {
	C.vkCmdCopyMemoryToImageIndirectKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyMemoryToImageIndirectInfoKHR)(pCopyMemoryToImageIndirectInfo.Raw()))
}

// CmdCopyMemoryToImageIndirectNV wraps vkCmdCopyMemoryToImageIndirectNV.
func CmdCopyMemoryToImageIndirectNV(commandBuffer CommandBuffer, copyBufferAddress DeviceAddress, copyCount uint32, stride uint32, dstImage Image, dstImageLayout ImageLayout, pImageSubresources ImageSubresourceLayers) {
	C.vkCmdCopyMemoryToImageIndirectNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDeviceAddress(copyBufferAddress), C.uint32_t(copyCount), C.uint32_t(stride), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), (*C.VkImageSubresourceLayers)(pImageSubresources.Raw()))
}

// CmdCopyMemoryToMicromapEXT wraps vkCmdCopyMemoryToMicromapEXT.
func CmdCopyMemoryToMicromapEXT(commandBuffer CommandBuffer, pInfo CopyMemoryToMicromapInfoEXT) {
	C.vkCmdCopyMemoryToMicromapEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyMemoryToMicromapInfoEXT)(pInfo.Raw()))
}

// CmdCopyMicromapEXT wraps vkCmdCopyMicromapEXT.
func CmdCopyMicromapEXT(commandBuffer CommandBuffer, pInfo CopyMicromapInfoEXT) {
	C.vkCmdCopyMicromapEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyMicromapInfoEXT)(pInfo.Raw()))
}

// CmdCopyMicromapToMemoryEXT wraps vkCmdCopyMicromapToMemoryEXT.
func CmdCopyMicromapToMemoryEXT(commandBuffer CommandBuffer, pInfo CopyMicromapToMemoryInfoEXT) {
	C.vkCmdCopyMicromapToMemoryEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyMicromapToMemoryInfoEXT)(pInfo.Raw()))
}

// CmdCopyQueryPoolResults wraps vkCmdCopyQueryPoolResults.
func CmdCopyQueryPoolResults(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery uint32, queryCount uint32, dstBuffer Buffer, dstOffset DeviceSize, stride DeviceSize, flags QueryResultFlags) {
	C.vkCmdCopyQueryPoolResults(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(firstQuery), C.uint32_t(queryCount), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.VkDeviceSize(stride), C.VkQueryResultFlags(flags))
}

// CmdCopyTensorARM wraps vkCmdCopyTensorARM.
func CmdCopyTensorARM(commandBuffer CommandBuffer, pCopyTensorInfo CopyTensorInfoARM) {
	C.vkCmdCopyTensorARM(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCopyTensorInfoARM)(pCopyTensorInfo.Raw()))
}

// CmdCuLaunchKernelNVX wraps vkCmdCuLaunchKernelNVX.
func CmdCuLaunchKernelNVX(commandBuffer CommandBuffer, pLaunchInfo CuLaunchInfoNVX) {
	C.vkCmdCuLaunchKernelNVX(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkCuLaunchInfoNVX)(pLaunchInfo.Raw()))
}

// CmdDebugMarkerBeginEXT wraps vkCmdDebugMarkerBeginEXT.
func CmdDebugMarkerBeginEXT(commandBuffer CommandBuffer, pMarkerInfo DebugMarkerMarkerInfoEXT) {
	C.vkCmdDebugMarkerBeginEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDebugMarkerMarkerInfoEXT)(pMarkerInfo.Raw()))
}

// CmdDebugMarkerEndEXT wraps vkCmdDebugMarkerEndEXT.
func CmdDebugMarkerEndEXT(commandBuffer CommandBuffer) {
	C.vkCmdDebugMarkerEndEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdDebugMarkerInsertEXT wraps vkCmdDebugMarkerInsertEXT.
func CmdDebugMarkerInsertEXT(commandBuffer CommandBuffer, pMarkerInfo DebugMarkerMarkerInfoEXT) {
	C.vkCmdDebugMarkerInsertEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDebugMarkerMarkerInfoEXT)(pMarkerInfo.Raw()))
}

// CmdDecodeVideoKHR wraps vkCmdDecodeVideoKHR.
func CmdDecodeVideoKHR(commandBuffer CommandBuffer, pDecodeInfo VideoDecodeInfoKHR) {
	C.vkCmdDecodeVideoKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkVideoDecodeInfoKHR)(pDecodeInfo.Raw()))
}

// CmdDecompressMemoryIndirectCountNV wraps vkCmdDecompressMemoryIndirectCountNV.
func CmdDecompressMemoryIndirectCountNV(commandBuffer CommandBuffer, indirectCommandsAddress DeviceAddress, indirectCommandsCountAddress DeviceAddress, stride uint32) {
	C.vkCmdDecompressMemoryIndirectCountNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDeviceAddress(indirectCommandsAddress), C.VkDeviceAddress(indirectCommandsCountAddress), C.uint32_t(stride))
}

// CmdDecompressMemoryNV wraps vkCmdDecompressMemoryNV.
func CmdDecompressMemoryNV(commandBuffer CommandBuffer, decompressRegionCount uint32, pDecompressMemoryRegions DecompressMemoryRegionNV) {
	C.vkCmdDecompressMemoryNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(decompressRegionCount), (*C.VkDecompressMemoryRegionNV)(pDecompressMemoryRegions.Raw()))
}

// CmdDispatch wraps vkCmdDispatch.
func CmdDispatch(commandBuffer CommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.vkCmdDispatch(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(groupCountX), C.uint32_t(groupCountY), C.uint32_t(groupCountZ))
}

// CmdDispatchBase wraps vkCmdDispatchBase.
func CmdDispatchBase(commandBuffer CommandBuffer, baseGroupX uint32, baseGroupY uint32, baseGroupZ uint32, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.vkCmdDispatchBase(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(baseGroupX), C.uint32_t(baseGroupY), C.uint32_t(baseGroupZ), C.uint32_t(groupCountX), C.uint32_t(groupCountY), C.uint32_t(groupCountZ))
}

// CmdDispatchDataGraphARM wraps vkCmdDispatchDataGraphARM.
func CmdDispatchDataGraphARM(commandBuffer CommandBuffer, session DataGraphPipelineSessionARM, pInfo DataGraphPipelineDispatchInfoARM) {
	C.vkCmdDispatchDataGraphARM(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDataGraphPipelineSessionARM(session), (*C.VkDataGraphPipelineDispatchInfoARM)(pInfo.Raw()))
}

// CmdDispatchIndirect wraps vkCmdDispatchIndirect.
func CmdDispatchIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize) {
	C.vkCmdDispatchIndirect(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset))
}

// CmdDispatchTileQCOM wraps vkCmdDispatchTileQCOM.
func CmdDispatchTileQCOM(commandBuffer CommandBuffer, pDispatchTileInfo DispatchTileInfoQCOM) {
	C.vkCmdDispatchTileQCOM(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDispatchTileInfoQCOM)(pDispatchTileInfo.Raw()))
}

// CmdDraw wraps vkCmdDraw.
func CmdDraw(commandBuffer CommandBuffer, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	C.vkCmdDraw(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

// CmdDrawClusterHUAWEI wraps vkCmdDrawClusterHUAWEI.
func CmdDrawClusterHUAWEI(commandBuffer CommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.vkCmdDrawClusterHUAWEI(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(groupCountX), C.uint32_t(groupCountY), C.uint32_t(groupCountZ))
}

// CmdDrawClusterIndirectHUAWEI wraps vkCmdDrawClusterIndirectHUAWEI.
func CmdDrawClusterIndirectHUAWEI(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize) {
	C.vkCmdDrawClusterIndirectHUAWEI(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset))
}

// CmdDrawIndexed wraps vkCmdDrawIndexed.
func CmdDrawIndexed(commandBuffer CommandBuffer, indexCount uint32, instanceCount uint32, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	C.vkCmdDrawIndexed(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(indexCount), C.uint32_t(instanceCount), C.uint32_t(firstIndex), C.int32_t(vertexOffset), C.uint32_t(firstInstance))
}

// CmdDrawIndexedIndirect wraps vkCmdDrawIndexedIndirect.
func CmdDrawIndexedIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount uint32, stride uint32) {
	C.vkCmdDrawIndexedIndirect(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.uint32_t(drawCount), C.uint32_t(stride))
}

// CmdDrawIndexedIndirectCount wraps vkCmdDrawIndexedIndirectCount.
func CmdDrawIndexedIndirectCount(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	C.vkCmdDrawIndexedIndirectCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkBuffer(countBuffer), C.VkDeviceSize(countBufferOffset), C.uint32_t(maxDrawCount), C.uint32_t(stride))
}

// CmdDrawIndirect wraps vkCmdDrawIndirect.
func CmdDrawIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount uint32, stride uint32) {
	C.vkCmdDrawIndirect(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.uint32_t(drawCount), C.uint32_t(stride))
}

// CmdDrawIndirectByteCountEXT wraps vkCmdDrawIndirectByteCountEXT.
func CmdDrawIndirectByteCountEXT(commandBuffer CommandBuffer, instanceCount uint32, firstInstance uint32, counterBuffer Buffer, counterBufferOffset DeviceSize, counterOffset uint32, vertexStride uint32) {
	C.vkCmdDrawIndirectByteCountEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(instanceCount), C.uint32_t(firstInstance), C.VkBuffer(counterBuffer), C.VkDeviceSize(counterBufferOffset), C.uint32_t(counterOffset), C.uint32_t(vertexStride))
}

// CmdDrawIndirectCount wraps vkCmdDrawIndirectCount.
func CmdDrawIndirectCount(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	C.vkCmdDrawIndirectCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkBuffer(countBuffer), C.VkDeviceSize(countBufferOffset), C.uint32_t(maxDrawCount), C.uint32_t(stride))
}

// CmdDrawMeshTasksEXT wraps vkCmdDrawMeshTasksEXT.
func CmdDrawMeshTasksEXT(commandBuffer CommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.vkCmdDrawMeshTasksEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(groupCountX), C.uint32_t(groupCountY), C.uint32_t(groupCountZ))
}

// CmdDrawMeshTasksIndirectCountEXT wraps vkCmdDrawMeshTasksIndirectCountEXT.
func CmdDrawMeshTasksIndirectCountEXT(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	C.vkCmdDrawMeshTasksIndirectCountEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkBuffer(countBuffer), C.VkDeviceSize(countBufferOffset), C.uint32_t(maxDrawCount), C.uint32_t(stride))
}

// CmdDrawMeshTasksIndirectCountNV wraps vkCmdDrawMeshTasksIndirectCountNV.
func CmdDrawMeshTasksIndirectCountNV(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	C.vkCmdDrawMeshTasksIndirectCountNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkBuffer(countBuffer), C.VkDeviceSize(countBufferOffset), C.uint32_t(maxDrawCount), C.uint32_t(stride))
}

// CmdDrawMeshTasksIndirectEXT wraps vkCmdDrawMeshTasksIndirectEXT.
func CmdDrawMeshTasksIndirectEXT(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount uint32, stride uint32) {
	C.vkCmdDrawMeshTasksIndirectEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.uint32_t(drawCount), C.uint32_t(stride))
}

// CmdDrawMeshTasksIndirectNV wraps vkCmdDrawMeshTasksIndirectNV.
func CmdDrawMeshTasksIndirectNV(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount uint32, stride uint32) {
	C.vkCmdDrawMeshTasksIndirectNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.uint32_t(drawCount), C.uint32_t(stride))
}

// CmdDrawMeshTasksNV wraps vkCmdDrawMeshTasksNV.
func CmdDrawMeshTasksNV(commandBuffer CommandBuffer, taskCount uint32, firstTask uint32) {
	C.vkCmdDrawMeshTasksNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(taskCount), C.uint32_t(firstTask))
}

// CmdDrawMultiEXT wraps vkCmdDrawMultiEXT.
func CmdDrawMultiEXT(commandBuffer CommandBuffer, drawCount uint32, pVertexInfo MultiDrawInfoEXT, instanceCount uint32, firstInstance uint32, stride uint32) {
	C.vkCmdDrawMultiEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(drawCount), (*C.VkMultiDrawInfoEXT)(pVertexInfo.Raw()), C.uint32_t(instanceCount), C.uint32_t(firstInstance), C.uint32_t(stride))
}

// CmdDrawMultiIndexedEXT wraps vkCmdDrawMultiIndexedEXT.
func CmdDrawMultiIndexedEXT(commandBuffer CommandBuffer, drawCount uint32, pIndexInfo MultiDrawIndexedInfoEXT, instanceCount uint32, firstInstance uint32, stride uint32, pVertexOffset ffi.Ref[int32]) {
	C.vkCmdDrawMultiIndexedEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(drawCount), (*C.VkMultiDrawIndexedInfoEXT)(pIndexInfo.Raw()), C.uint32_t(instanceCount), C.uint32_t(firstInstance), C.uint32_t(stride), (*C.int32_t)(pVertexOffset.Raw()))
}

// CmdEncodeVideoKHR wraps vkCmdEncodeVideoKHR.
func CmdEncodeVideoKHR(commandBuffer CommandBuffer, pEncodeInfo VideoEncodeInfoKHR) {
	C.vkCmdEncodeVideoKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkVideoEncodeInfoKHR)(pEncodeInfo.Raw()))
}

// CmdEndConditionalRenderingEXT wraps vkCmdEndConditionalRenderingEXT.
func CmdEndConditionalRenderingEXT(commandBuffer CommandBuffer) {
	C.vkCmdEndConditionalRenderingEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdEndDebugUtilsLabelEXT wraps vkCmdEndDebugUtilsLabelEXT.
func CmdEndDebugUtilsLabelEXT(commandBuffer CommandBuffer) {
	C.vkCmdEndDebugUtilsLabelEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdEndPerTileExecutionQCOM wraps vkCmdEndPerTileExecutionQCOM.
func CmdEndPerTileExecutionQCOM(commandBuffer CommandBuffer, pPerTileEndInfo PerTileEndInfoQCOM) {
	C.vkCmdEndPerTileExecutionQCOM(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPerTileEndInfoQCOM)(pPerTileEndInfo.Raw()))
}

// CmdEndQuery wraps vkCmdEndQuery.
func CmdEndQuery(commandBuffer CommandBuffer, queryPool QueryPool, query uint32) {
	C.vkCmdEndQuery(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(query))
}

// CmdEndQueryIndexedEXT wraps vkCmdEndQueryIndexedEXT.
func CmdEndQueryIndexedEXT(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, index uint32) {
	C.vkCmdEndQueryIndexedEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(query), C.uint32_t(index))
}

// CmdEndRenderPass wraps vkCmdEndRenderPass.
func CmdEndRenderPass(commandBuffer CommandBuffer) {
	C.vkCmdEndRenderPass(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdEndRenderPass2 wraps vkCmdEndRenderPass2.
func CmdEndRenderPass2(commandBuffer CommandBuffer, pSubpassEndInfo SubpassEndInfo) {
	C.vkCmdEndRenderPass2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkSubpassEndInfo)(pSubpassEndInfo.Raw()))
}

// CmdEndRendering wraps vkCmdEndRendering.
func CmdEndRendering(commandBuffer CommandBuffer) {
	C.vkCmdEndRendering(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdEndRendering2EXT wraps vkCmdEndRendering2EXT.
func CmdEndRendering2EXT(commandBuffer CommandBuffer, pRenderingEndInfo RenderingEndInfoEXT) {
	C.vkCmdEndRendering2EXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkRenderingEndInfoEXT)(pRenderingEndInfo.Raw()))
}

// CmdEndTransformFeedbackEXT wraps vkCmdEndTransformFeedbackEXT.
func CmdEndTransformFeedbackEXT(commandBuffer CommandBuffer, firstCounterBuffer uint32, counterBufferCount uint32, pCounterBuffers ffi.Ref[Buffer], pCounterBufferOffsets ffi.Ref[DeviceSize]) {
	C.vkCmdEndTransformFeedbackEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstCounterBuffer), C.uint32_t(counterBufferCount), (*C.VkBuffer)(pCounterBuffers.Raw()), (*C.VkDeviceSize)(pCounterBufferOffsets.Raw()))
}

// CmdEndVideoCodingKHR wraps vkCmdEndVideoCodingKHR.
func CmdEndVideoCodingKHR(commandBuffer CommandBuffer, pEndCodingInfo VideoEndCodingInfoKHR) {
	C.vkCmdEndVideoCodingKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkVideoEndCodingInfoKHR)(pEndCodingInfo.Raw()))
}

// CmdExecuteCommands wraps vkCmdExecuteCommands.
func CmdExecuteCommands(commandBuffer CommandBuffer, commandBufferCount uint32, pCommandBuffers ffi.Ref[CommandBuffer]) {
	C.vkCmdExecuteCommands(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(commandBufferCount), (*C.VkCommandBuffer)(pCommandBuffers.Raw()))
}

// CmdExecuteGeneratedCommandsEXT wraps vkCmdExecuteGeneratedCommandsEXT.
func CmdExecuteGeneratedCommandsEXT(commandBuffer CommandBuffer, isPreprocessed bool, pGeneratedCommandsInfo GeneratedCommandsInfoEXT) {
	tmp_isPreprocessed := 0

	if isPreprocessed {
		tmp_isPreprocessed = 1
	}

	C.vkCmdExecuteGeneratedCommandsEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_isPreprocessed), (*C.VkGeneratedCommandsInfoEXT)(pGeneratedCommandsInfo.Raw()))
}

// CmdExecuteGeneratedCommandsNV wraps vkCmdExecuteGeneratedCommandsNV.
func CmdExecuteGeneratedCommandsNV(commandBuffer CommandBuffer, isPreprocessed bool, pGeneratedCommandsInfo GeneratedCommandsInfoNV) {
	tmp_isPreprocessed := 0

	if isPreprocessed {
		tmp_isPreprocessed = 1
	}

	C.vkCmdExecuteGeneratedCommandsNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_isPreprocessed), (*C.VkGeneratedCommandsInfoNV)(pGeneratedCommandsInfo.Raw()))
}

// CmdFillBuffer wraps vkCmdFillBuffer.
func CmdFillBuffer(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset DeviceSize, size DeviceSize, data uint32) {
	C.vkCmdFillBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.VkDeviceSize(size), C.uint32_t(data))
}

// CmdInsertDebugUtilsLabelEXT wraps vkCmdInsertDebugUtilsLabelEXT.
func CmdInsertDebugUtilsLabelEXT(commandBuffer CommandBuffer, pLabelInfo DebugUtilsLabelEXT) {
	C.vkCmdInsertDebugUtilsLabelEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDebugUtilsLabelEXT)(pLabelInfo.Raw()))
}

// CmdNextSubpass wraps vkCmdNextSubpass.
func CmdNextSubpass(commandBuffer CommandBuffer, contents SubpassContents) {
	C.vkCmdNextSubpass(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkSubpassContents(contents))
}

// CmdNextSubpass2 wraps vkCmdNextSubpass2.
func CmdNextSubpass2(commandBuffer CommandBuffer, pSubpassBeginInfo SubpassBeginInfo, pSubpassEndInfo SubpassEndInfo) {
	C.vkCmdNextSubpass2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkSubpassBeginInfo)(pSubpassBeginInfo.Raw()), (*C.VkSubpassEndInfo)(pSubpassEndInfo.Raw()))
}

// CmdOpticalFlowExecuteNV wraps vkCmdOpticalFlowExecuteNV.
func CmdOpticalFlowExecuteNV(commandBuffer CommandBuffer, session OpticalFlowSessionNV, pExecuteInfo OpticalFlowExecuteInfoNV) {
	C.vkCmdOpticalFlowExecuteNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkOpticalFlowSessionNV(session), (*C.VkOpticalFlowExecuteInfoNV)(pExecuteInfo.Raw()))
}

// CmdPipelineBarrier wraps vkCmdPipelineBarrier.
func CmdPipelineBarrier(commandBuffer CommandBuffer, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, dependencyFlags DependencyFlags, memoryBarrierCount uint32, pMemoryBarriers MemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers BufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers ImageMemoryBarrier) {
	C.vkCmdPipelineBarrier(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlags(srcStageMask), C.VkPipelineStageFlags(dstStageMask), C.VkDependencyFlags(dependencyFlags), C.uint32_t(memoryBarrierCount), (*C.VkMemoryBarrier)(pMemoryBarriers.Raw()), C.uint32_t(bufferMemoryBarrierCount), (*C.VkBufferMemoryBarrier)(pBufferMemoryBarriers.Raw()), C.uint32_t(imageMemoryBarrierCount), (*C.VkImageMemoryBarrier)(pImageMemoryBarriers.Raw()))
}

// CmdPipelineBarrier2 wraps vkCmdPipelineBarrier2.
func CmdPipelineBarrier2(commandBuffer CommandBuffer, pDependencyInfo DependencyInfo) {
	C.vkCmdPipelineBarrier2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDependencyInfo)(pDependencyInfo.Raw()))
}

// CmdPreprocessGeneratedCommandsEXT wraps vkCmdPreprocessGeneratedCommandsEXT.
func CmdPreprocessGeneratedCommandsEXT(commandBuffer CommandBuffer, pGeneratedCommandsInfo GeneratedCommandsInfoEXT, stateCommandBuffer CommandBuffer) {
	C.vkCmdPreprocessGeneratedCommandsEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkGeneratedCommandsInfoEXT)(pGeneratedCommandsInfo.Raw()), C.VkCommandBuffer(unsafe.Pointer(stateCommandBuffer)))
}

// CmdPreprocessGeneratedCommandsNV wraps vkCmdPreprocessGeneratedCommandsNV.
func CmdPreprocessGeneratedCommandsNV(commandBuffer CommandBuffer, pGeneratedCommandsInfo GeneratedCommandsInfoNV) {
	C.vkCmdPreprocessGeneratedCommandsNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkGeneratedCommandsInfoNV)(pGeneratedCommandsInfo.Raw()))
}

// CmdPushConstants wraps vkCmdPushConstants.
func CmdPushConstants(commandBuffer CommandBuffer, layout PipelineLayout, stageFlags ShaderStageFlags, offset uint32, size uint32, pValues unsafe.Pointer) {
	C.vkCmdPushConstants(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineLayout(layout), C.VkShaderStageFlags(stageFlags), C.uint32_t(offset), C.uint32_t(size), pValues)
}

// CmdPushConstants2 wraps vkCmdPushConstants2.
func CmdPushConstants2(commandBuffer CommandBuffer, pPushConstantsInfo PushConstantsInfo) {
	C.vkCmdPushConstants2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPushConstantsInfo)(pPushConstantsInfo.Raw()))
}

// CmdPushDescriptorSet wraps vkCmdPushDescriptorSet.
func CmdPushDescriptorSet(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, set uint32, descriptorWriteCount uint32, pDescriptorWrites WriteDescriptorSet) {
	C.vkCmdPushDescriptorSet(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipelineLayout(layout), C.uint32_t(set), C.uint32_t(descriptorWriteCount), (*C.VkWriteDescriptorSet)(pDescriptorWrites.Raw()))
}

// CmdPushDescriptorSet2 wraps vkCmdPushDescriptorSet2.
func CmdPushDescriptorSet2(commandBuffer CommandBuffer, pPushDescriptorSetInfo PushDescriptorSetInfo) {
	C.vkCmdPushDescriptorSet2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPushDescriptorSetInfo)(pPushDescriptorSetInfo.Raw()))
}

// CmdPushDescriptorSetWithTemplate wraps vkCmdPushDescriptorSetWithTemplate.
func CmdPushDescriptorSetWithTemplate(commandBuffer CommandBuffer, descriptorUpdateTemplate DescriptorUpdateTemplate, layout PipelineLayout, set uint32, pData unsafe.Pointer) {
	C.vkCmdPushDescriptorSetWithTemplate(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate), C.VkPipelineLayout(layout), C.uint32_t(set), pData)
}

// CmdPushDescriptorSetWithTemplate2 wraps vkCmdPushDescriptorSetWithTemplate2.
func CmdPushDescriptorSetWithTemplate2(commandBuffer CommandBuffer, pPushDescriptorSetWithTemplateInfo PushDescriptorSetWithTemplateInfo) {
	C.vkCmdPushDescriptorSetWithTemplate2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPushDescriptorSetWithTemplateInfo)(pPushDescriptorSetWithTemplateInfo.Raw()))
}

// CmdResetEvent wraps vkCmdResetEvent.
func CmdResetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	C.vkCmdResetEvent(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkEvent(event), C.VkPipelineStageFlags(stageMask))
}

// CmdResetEvent2 wraps vkCmdResetEvent2.
func CmdResetEvent2(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags2) {
	C.vkCmdResetEvent2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkEvent(event), C.VkPipelineStageFlags2(stageMask))
}

// CmdResetQueryPool wraps vkCmdResetQueryPool.
func CmdResetQueryPool(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery uint32, queryCount uint32) {
	C.vkCmdResetQueryPool(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(firstQuery), C.uint32_t(queryCount))
}

// CmdResolveImage wraps vkCmdResolveImage.
func CmdResolveImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions ImageResolve) {
	C.vkCmdResolveImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), (*C.VkImageResolve)(pRegions.Raw()))
}

// CmdResolveImage2 wraps vkCmdResolveImage2.
func CmdResolveImage2(commandBuffer CommandBuffer, pResolveImageInfo ResolveImageInfo2) {
	C.vkCmdResolveImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkResolveImageInfo2)(pResolveImageInfo.Raw()))
}

// CmdSetAlphaToCoverageEnableEXT wraps vkCmdSetAlphaToCoverageEnableEXT.
func CmdSetAlphaToCoverageEnableEXT(commandBuffer CommandBuffer, alphaToCoverageEnable bool) {
	tmp_alphaToCoverageEnable := 0

	if alphaToCoverageEnable {
		tmp_alphaToCoverageEnable = 1
	}

	C.vkCmdSetAlphaToCoverageEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_alphaToCoverageEnable))
}

// CmdSetAlphaToOneEnableEXT wraps vkCmdSetAlphaToOneEnableEXT.
func CmdSetAlphaToOneEnableEXT(commandBuffer CommandBuffer, alphaToOneEnable bool) {
	tmp_alphaToOneEnable := 0

	if alphaToOneEnable {
		tmp_alphaToOneEnable = 1
	}

	C.vkCmdSetAlphaToOneEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_alphaToOneEnable))
}

// CmdSetAttachmentFeedbackLoopEnableEXT wraps vkCmdSetAttachmentFeedbackLoopEnableEXT.
func CmdSetAttachmentFeedbackLoopEnableEXT(commandBuffer CommandBuffer, aspectMask ImageAspectFlags) {
	C.vkCmdSetAttachmentFeedbackLoopEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImageAspectFlags(aspectMask))
}

// vkCmdSetBlendConstants.blendConstants is unsupported: category unsupported.

// CmdSetCheckpointNV wraps vkCmdSetCheckpointNV.
func CmdSetCheckpointNV(commandBuffer CommandBuffer, pCheckpointMarker unsafe.Pointer) {
	C.vkCmdSetCheckpointNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pCheckpointMarker)
}

// CmdSetCoarseSampleOrderNV wraps vkCmdSetCoarseSampleOrderNV.
func CmdSetCoarseSampleOrderNV(commandBuffer CommandBuffer, sampleOrderType CoarseSampleOrderTypeNV, customSampleOrderCount uint32, pCustomSampleOrders CoarseSampleOrderCustomNV) {
	C.vkCmdSetCoarseSampleOrderNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCoarseSampleOrderTypeNV(sampleOrderType), C.uint32_t(customSampleOrderCount), (*C.VkCoarseSampleOrderCustomNV)(pCustomSampleOrders.Raw()))
}

// CmdSetColorBlendAdvancedEXT wraps vkCmdSetColorBlendAdvancedEXT.
func CmdSetColorBlendAdvancedEXT(commandBuffer CommandBuffer, firstAttachment uint32, attachmentCount uint32, pColorBlendAdvanced ColorBlendAdvancedEXT) {
	C.vkCmdSetColorBlendAdvancedEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstAttachment), C.uint32_t(attachmentCount), (*C.VkColorBlendAdvancedEXT)(pColorBlendAdvanced.Raw()))
}

// vkCmdSetColorBlendEnableEXT.pColorBlendEnables is unsupported: category pointer -> ?? VkBool32.

// CmdSetColorBlendEquationEXT wraps vkCmdSetColorBlendEquationEXT.
func CmdSetColorBlendEquationEXT(commandBuffer CommandBuffer, firstAttachment uint32, attachmentCount uint32, pColorBlendEquations ColorBlendEquationEXT) {
	C.vkCmdSetColorBlendEquationEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstAttachment), C.uint32_t(attachmentCount), (*C.VkColorBlendEquationEXT)(pColorBlendEquations.Raw()))
}

// vkCmdSetColorWriteEnableEXT.pColorWriteEnables is unsupported: category pointer -> ?? VkBool32.

// CmdSetColorWriteMaskEXT wraps vkCmdSetColorWriteMaskEXT.
func CmdSetColorWriteMaskEXT(commandBuffer CommandBuffer, firstAttachment uint32, attachmentCount uint32, pColorWriteMasks ffi.Ref[ColorComponentFlags]) {
	C.vkCmdSetColorWriteMaskEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstAttachment), C.uint32_t(attachmentCount), (*C.VkColorComponentFlags)(pColorWriteMasks.Raw()))
}

// CmdSetConservativeRasterizationModeEXT wraps vkCmdSetConservativeRasterizationModeEXT.
func CmdSetConservativeRasterizationModeEXT(commandBuffer CommandBuffer, conservativeRasterizationMode ConservativeRasterizationModeEXT) {
	C.vkCmdSetConservativeRasterizationModeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkConservativeRasterizationModeEXT(conservativeRasterizationMode))
}

// CmdSetCoverageModulationModeNV wraps vkCmdSetCoverageModulationModeNV.
func CmdSetCoverageModulationModeNV(commandBuffer CommandBuffer, coverageModulationMode CoverageModulationModeNV) {
	C.vkCmdSetCoverageModulationModeNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCoverageModulationModeNV(coverageModulationMode))
}

// CmdSetCoverageModulationTableEnableNV wraps vkCmdSetCoverageModulationTableEnableNV.
func CmdSetCoverageModulationTableEnableNV(commandBuffer CommandBuffer, coverageModulationTableEnable bool) {
	tmp_coverageModulationTableEnable := 0

	if coverageModulationTableEnable {
		tmp_coverageModulationTableEnable = 1
	}

	C.vkCmdSetCoverageModulationTableEnableNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_coverageModulationTableEnable))
}

// CmdSetCoverageModulationTableNV wraps vkCmdSetCoverageModulationTableNV.
func CmdSetCoverageModulationTableNV(commandBuffer CommandBuffer, coverageModulationTableCount uint32, pCoverageModulationTable ffi.Ref[float32]) {
	C.vkCmdSetCoverageModulationTableNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(coverageModulationTableCount), (*C.float)(pCoverageModulationTable.Raw()))
}

// CmdSetCoverageReductionModeNV wraps vkCmdSetCoverageReductionModeNV.
func CmdSetCoverageReductionModeNV(commandBuffer CommandBuffer, coverageReductionMode CoverageReductionModeNV) {
	C.vkCmdSetCoverageReductionModeNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCoverageReductionModeNV(coverageReductionMode))
}

// CmdSetCoverageToColorEnableNV wraps vkCmdSetCoverageToColorEnableNV.
func CmdSetCoverageToColorEnableNV(commandBuffer CommandBuffer, coverageToColorEnable bool) {
	tmp_coverageToColorEnable := 0

	if coverageToColorEnable {
		tmp_coverageToColorEnable = 1
	}

	C.vkCmdSetCoverageToColorEnableNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_coverageToColorEnable))
}

// CmdSetCoverageToColorLocationNV wraps vkCmdSetCoverageToColorLocationNV.
func CmdSetCoverageToColorLocationNV(commandBuffer CommandBuffer, coverageToColorLocation uint32) {
	C.vkCmdSetCoverageToColorLocationNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(coverageToColorLocation))
}

// CmdSetCullMode wraps vkCmdSetCullMode.
func CmdSetCullMode(commandBuffer CommandBuffer, cullMode CullModeFlags) {
	C.vkCmdSetCullMode(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCullModeFlags(cullMode))
}

// CmdSetDepthBias wraps vkCmdSetDepthBias.
func CmdSetDepthBias(commandBuffer CommandBuffer, depthBiasConstantFactor float32, depthBiasClamp float32, depthBiasSlopeFactor float32) {
	C.vkCmdSetDepthBias(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.float(depthBiasConstantFactor), C.float(depthBiasClamp), C.float(depthBiasSlopeFactor))
}

// CmdSetDepthBias2EXT wraps vkCmdSetDepthBias2EXT.
func CmdSetDepthBias2EXT(commandBuffer CommandBuffer, pDepthBiasInfo DepthBiasInfoEXT) {
	C.vkCmdSetDepthBias2EXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkDepthBiasInfoEXT)(pDepthBiasInfo.Raw()))
}

// CmdSetDepthBiasEnable wraps vkCmdSetDepthBiasEnable.
func CmdSetDepthBiasEnable(commandBuffer CommandBuffer, depthBiasEnable bool) {
	tmp_depthBiasEnable := 0

	if depthBiasEnable {
		tmp_depthBiasEnable = 1
	}

	C.vkCmdSetDepthBiasEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_depthBiasEnable))
}

// CmdSetDepthBounds wraps vkCmdSetDepthBounds.
func CmdSetDepthBounds(commandBuffer CommandBuffer, minDepthBounds float32, maxDepthBounds float32) {
	C.vkCmdSetDepthBounds(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.float(minDepthBounds), C.float(maxDepthBounds))
}

// CmdSetDepthBoundsTestEnable wraps vkCmdSetDepthBoundsTestEnable.
func CmdSetDepthBoundsTestEnable(commandBuffer CommandBuffer, depthBoundsTestEnable bool) {
	tmp_depthBoundsTestEnable := 0

	if depthBoundsTestEnable {
		tmp_depthBoundsTestEnable = 1
	}

	C.vkCmdSetDepthBoundsTestEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_depthBoundsTestEnable))
}

// CmdSetDepthClampEnableEXT wraps vkCmdSetDepthClampEnableEXT.
func CmdSetDepthClampEnableEXT(commandBuffer CommandBuffer, depthClampEnable bool) {
	tmp_depthClampEnable := 0

	if depthClampEnable {
		tmp_depthClampEnable = 1
	}

	C.vkCmdSetDepthClampEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_depthClampEnable))
}

// CmdSetDepthClampRangeEXT wraps vkCmdSetDepthClampRangeEXT.
func CmdSetDepthClampRangeEXT(commandBuffer CommandBuffer, depthClampMode DepthClampModeEXT, pDepthClampRange DepthClampRangeEXT) {
	C.vkCmdSetDepthClampRangeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDepthClampModeEXT(depthClampMode), (*C.VkDepthClampRangeEXT)(pDepthClampRange.Raw()))
}

// CmdSetDepthClipEnableEXT wraps vkCmdSetDepthClipEnableEXT.
func CmdSetDepthClipEnableEXT(commandBuffer CommandBuffer, depthClipEnable bool) {
	tmp_depthClipEnable := 0

	if depthClipEnable {
		tmp_depthClipEnable = 1
	}

	C.vkCmdSetDepthClipEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_depthClipEnable))
}

// CmdSetDepthClipNegativeOneToOneEXT wraps vkCmdSetDepthClipNegativeOneToOneEXT.
func CmdSetDepthClipNegativeOneToOneEXT(commandBuffer CommandBuffer, negativeOneToOne bool) {
	tmp_negativeOneToOne := 0

	if negativeOneToOne {
		tmp_negativeOneToOne = 1
	}

	C.vkCmdSetDepthClipNegativeOneToOneEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_negativeOneToOne))
}

// CmdSetDepthCompareOp wraps vkCmdSetDepthCompareOp.
func CmdSetDepthCompareOp(commandBuffer CommandBuffer, depthCompareOp CompareOp) {
	C.vkCmdSetDepthCompareOp(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCompareOp(depthCompareOp))
}

// CmdSetDepthTestEnable wraps vkCmdSetDepthTestEnable.
func CmdSetDepthTestEnable(commandBuffer CommandBuffer, depthTestEnable bool) {
	tmp_depthTestEnable := 0

	if depthTestEnable {
		tmp_depthTestEnable = 1
	}

	C.vkCmdSetDepthTestEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_depthTestEnable))
}

// CmdSetDepthWriteEnable wraps vkCmdSetDepthWriteEnable.
func CmdSetDepthWriteEnable(commandBuffer CommandBuffer, depthWriteEnable bool) {
	tmp_depthWriteEnable := 0

	if depthWriteEnable {
		tmp_depthWriteEnable = 1
	}

	C.vkCmdSetDepthWriteEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_depthWriteEnable))
}

// CmdSetDescriptorBufferOffsets2EXT wraps vkCmdSetDescriptorBufferOffsets2EXT.
func CmdSetDescriptorBufferOffsets2EXT(commandBuffer CommandBuffer, pSetDescriptorBufferOffsetsInfo SetDescriptorBufferOffsetsInfoEXT) {
	C.vkCmdSetDescriptorBufferOffsets2EXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkSetDescriptorBufferOffsetsInfoEXT)(pSetDescriptorBufferOffsetsInfo.Raw()))
}

// CmdSetDescriptorBufferOffsetsEXT wraps vkCmdSetDescriptorBufferOffsetsEXT.
func CmdSetDescriptorBufferOffsetsEXT(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, firstSet uint32, setCount uint32, pBufferIndices ffi.Ref[uint32], pOffsets ffi.Ref[DeviceSize]) {
	C.vkCmdSetDescriptorBufferOffsetsEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipelineLayout(layout), C.uint32_t(firstSet), C.uint32_t(setCount), (*C.uint32_t)(pBufferIndices.Raw()), (*C.VkDeviceSize)(pOffsets.Raw()))
}

// CmdSetDeviceMask wraps vkCmdSetDeviceMask.
func CmdSetDeviceMask(commandBuffer CommandBuffer, deviceMask uint32) {
	C.vkCmdSetDeviceMask(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(deviceMask))
}

// CmdSetDiscardRectangleEXT wraps vkCmdSetDiscardRectangleEXT.
func CmdSetDiscardRectangleEXT(commandBuffer CommandBuffer, firstDiscardRectangle uint32, discardRectangleCount uint32, pDiscardRectangles Rect2D) {
	C.vkCmdSetDiscardRectangleEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstDiscardRectangle), C.uint32_t(discardRectangleCount), (*C.VkRect2D)(pDiscardRectangles.Raw()))
}

// CmdSetDiscardRectangleEnableEXT wraps vkCmdSetDiscardRectangleEnableEXT.
func CmdSetDiscardRectangleEnableEXT(commandBuffer CommandBuffer, discardRectangleEnable bool) {
	tmp_discardRectangleEnable := 0

	if discardRectangleEnable {
		tmp_discardRectangleEnable = 1
	}

	C.vkCmdSetDiscardRectangleEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_discardRectangleEnable))
}

// CmdSetDiscardRectangleModeEXT wraps vkCmdSetDiscardRectangleModeEXT.
func CmdSetDiscardRectangleModeEXT(commandBuffer CommandBuffer, discardRectangleMode DiscardRectangleModeEXT) {
	C.vkCmdSetDiscardRectangleModeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDiscardRectangleModeEXT(discardRectangleMode))
}

// CmdSetEvent wraps vkCmdSetEvent.
func CmdSetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	C.vkCmdSetEvent(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkEvent(event), C.VkPipelineStageFlags(stageMask))
}

// CmdSetEvent2 wraps vkCmdSetEvent2.
func CmdSetEvent2(commandBuffer CommandBuffer, event Event, pDependencyInfo DependencyInfo) {
	C.vkCmdSetEvent2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkEvent(event), (*C.VkDependencyInfo)(pDependencyInfo.Raw()))
}

// vkCmdSetExclusiveScissorEnableNV.pExclusiveScissorEnables is unsupported: category pointer -> ?? VkBool32.

// CmdSetExclusiveScissorNV wraps vkCmdSetExclusiveScissorNV.
func CmdSetExclusiveScissorNV(commandBuffer CommandBuffer, firstExclusiveScissor uint32, exclusiveScissorCount uint32, pExclusiveScissors Rect2D) {
	C.vkCmdSetExclusiveScissorNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstExclusiveScissor), C.uint32_t(exclusiveScissorCount), (*C.VkRect2D)(pExclusiveScissors.Raw()))
}

// CmdSetExtraPrimitiveOverestimationSizeEXT wraps vkCmdSetExtraPrimitiveOverestimationSizeEXT.
func CmdSetExtraPrimitiveOverestimationSizeEXT(commandBuffer CommandBuffer, extraPrimitiveOverestimationSize float32) {
	C.vkCmdSetExtraPrimitiveOverestimationSizeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.float(extraPrimitiveOverestimationSize))
}

// vkCmdSetFragmentShadingRateEnumNV.combinerOps is unsupported: category unsupported.

// vkCmdSetFragmentShadingRateKHR.combinerOps is unsupported: category unsupported.

// CmdSetFrontFace wraps vkCmdSetFrontFace.
func CmdSetFrontFace(commandBuffer CommandBuffer, frontFace FrontFace) {
	C.vkCmdSetFrontFace(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkFrontFace(frontFace))
}

// CmdSetLineRasterizationModeEXT wraps vkCmdSetLineRasterizationModeEXT.
func CmdSetLineRasterizationModeEXT(commandBuffer CommandBuffer, lineRasterizationMode LineRasterizationMode) {
	C.vkCmdSetLineRasterizationModeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkLineRasterizationModeEXT(lineRasterizationMode))
}

// CmdSetLineStipple wraps vkCmdSetLineStipple.
func CmdSetLineStipple(commandBuffer CommandBuffer, lineStippleFactor uint32, lineStipplePattern uint16) {
	C.vkCmdSetLineStipple(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(lineStippleFactor), C.uint16_t(lineStipplePattern))
}

// CmdSetLineStippleEnableEXT wraps vkCmdSetLineStippleEnableEXT.
func CmdSetLineStippleEnableEXT(commandBuffer CommandBuffer, stippledLineEnable bool) {
	tmp_stippledLineEnable := 0

	if stippledLineEnable {
		tmp_stippledLineEnable = 1
	}

	C.vkCmdSetLineStippleEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_stippledLineEnable))
}

// CmdSetLineWidth wraps vkCmdSetLineWidth.
func CmdSetLineWidth(commandBuffer CommandBuffer, lineWidth float32) {
	C.vkCmdSetLineWidth(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.float(lineWidth))
}

// CmdSetLogicOpEXT wraps vkCmdSetLogicOpEXT.
func CmdSetLogicOpEXT(commandBuffer CommandBuffer, logicOp LogicOp) {
	C.vkCmdSetLogicOpEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkLogicOp(logicOp))
}

// CmdSetLogicOpEnableEXT wraps vkCmdSetLogicOpEnableEXT.
func CmdSetLogicOpEnableEXT(commandBuffer CommandBuffer, logicOpEnable bool) {
	tmp_logicOpEnable := 0

	if logicOpEnable {
		tmp_logicOpEnable = 1
	}

	C.vkCmdSetLogicOpEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_logicOpEnable))
}

// CmdSetPatchControlPointsEXT wraps vkCmdSetPatchControlPointsEXT.
func CmdSetPatchControlPointsEXT(commandBuffer CommandBuffer, patchControlPoints uint32) {
	C.vkCmdSetPatchControlPointsEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(patchControlPoints))
}

// CmdSetPerformanceMarkerINTEL wraps vkCmdSetPerformanceMarkerINTEL.
func CmdSetPerformanceMarkerINTEL(commandBuffer CommandBuffer, pMarkerInfo PerformanceMarkerInfoINTEL) Result {
	ret := C.vkCmdSetPerformanceMarkerINTEL(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPerformanceMarkerInfoINTEL)(pMarkerInfo.Raw()))

	return Result(ret)
}

// CmdSetPerformanceOverrideINTEL wraps vkCmdSetPerformanceOverrideINTEL.
func CmdSetPerformanceOverrideINTEL(commandBuffer CommandBuffer, pOverrideInfo PerformanceOverrideInfoINTEL) Result {
	ret := C.vkCmdSetPerformanceOverrideINTEL(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPerformanceOverrideInfoINTEL)(pOverrideInfo.Raw()))

	return Result(ret)
}

// CmdSetPerformanceStreamMarkerINTEL wraps vkCmdSetPerformanceStreamMarkerINTEL.
func CmdSetPerformanceStreamMarkerINTEL(commandBuffer CommandBuffer, pMarkerInfo PerformanceStreamMarkerInfoINTEL) Result {
	ret := C.vkCmdSetPerformanceStreamMarkerINTEL(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkPerformanceStreamMarkerInfoINTEL)(pMarkerInfo.Raw()))

	return Result(ret)
}

// CmdSetPolygonModeEXT wraps vkCmdSetPolygonModeEXT.
func CmdSetPolygonModeEXT(commandBuffer CommandBuffer, polygonMode PolygonMode) {
	C.vkCmdSetPolygonModeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPolygonMode(polygonMode))
}

// CmdSetPrimitiveRestartEnable wraps vkCmdSetPrimitiveRestartEnable.
func CmdSetPrimitiveRestartEnable(commandBuffer CommandBuffer, primitiveRestartEnable bool) {
	tmp_primitiveRestartEnable := 0

	if primitiveRestartEnable {
		tmp_primitiveRestartEnable = 1
	}

	C.vkCmdSetPrimitiveRestartEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_primitiveRestartEnable))
}

// CmdSetPrimitiveTopology wraps vkCmdSetPrimitiveTopology.
func CmdSetPrimitiveTopology(commandBuffer CommandBuffer, primitiveTopology PrimitiveTopology) {
	C.vkCmdSetPrimitiveTopology(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPrimitiveTopology(primitiveTopology))
}

// CmdSetProvokingVertexModeEXT wraps vkCmdSetProvokingVertexModeEXT.
func CmdSetProvokingVertexModeEXT(commandBuffer CommandBuffer, provokingVertexMode ProvokingVertexModeEXT) {
	C.vkCmdSetProvokingVertexModeEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkProvokingVertexModeEXT(provokingVertexMode))
}

// CmdSetRasterizationSamplesEXT wraps vkCmdSetRasterizationSamplesEXT.
func CmdSetRasterizationSamplesEXT(commandBuffer CommandBuffer, rasterizationSamples SampleCountFlags) {
	C.vkCmdSetRasterizationSamplesEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkSampleCountFlagBits(rasterizationSamples))
}

// CmdSetRasterizationStreamEXT wraps vkCmdSetRasterizationStreamEXT.
func CmdSetRasterizationStreamEXT(commandBuffer CommandBuffer, rasterizationStream uint32) {
	C.vkCmdSetRasterizationStreamEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(rasterizationStream))
}

// CmdSetRasterizerDiscardEnable wraps vkCmdSetRasterizerDiscardEnable.
func CmdSetRasterizerDiscardEnable(commandBuffer CommandBuffer, rasterizerDiscardEnable bool) {
	tmp_rasterizerDiscardEnable := 0

	if rasterizerDiscardEnable {
		tmp_rasterizerDiscardEnable = 1
	}

	C.vkCmdSetRasterizerDiscardEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_rasterizerDiscardEnable))
}

// CmdSetRayTracingPipelineStackSizeKHR wraps vkCmdSetRayTracingPipelineStackSizeKHR.
func CmdSetRayTracingPipelineStackSizeKHR(commandBuffer CommandBuffer, pipelineStackSize uint32) {
	C.vkCmdSetRayTracingPipelineStackSizeKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(pipelineStackSize))
}

// CmdSetRenderingAttachmentLocations wraps vkCmdSetRenderingAttachmentLocations.
func CmdSetRenderingAttachmentLocations(commandBuffer CommandBuffer, pLocationInfo RenderingAttachmentLocationInfo) {
	C.vkCmdSetRenderingAttachmentLocations(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkRenderingAttachmentLocationInfo)(pLocationInfo.Raw()))
}

// CmdSetRenderingInputAttachmentIndices wraps vkCmdSetRenderingInputAttachmentIndices.
func CmdSetRenderingInputAttachmentIndices(commandBuffer CommandBuffer, pInputAttachmentIndexInfo RenderingInputAttachmentIndexInfo) {
	C.vkCmdSetRenderingInputAttachmentIndices(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkRenderingInputAttachmentIndexInfo)(pInputAttachmentIndexInfo.Raw()))
}

// CmdSetRepresentativeFragmentTestEnableNV wraps vkCmdSetRepresentativeFragmentTestEnableNV.
func CmdSetRepresentativeFragmentTestEnableNV(commandBuffer CommandBuffer, representativeFragmentTestEnable bool) {
	tmp_representativeFragmentTestEnable := 0

	if representativeFragmentTestEnable {
		tmp_representativeFragmentTestEnable = 1
	}

	C.vkCmdSetRepresentativeFragmentTestEnableNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_representativeFragmentTestEnable))
}

// CmdSetSampleLocationsEXT wraps vkCmdSetSampleLocationsEXT.
func CmdSetSampleLocationsEXT(commandBuffer CommandBuffer, pSampleLocationsInfo SampleLocationsInfoEXT) {
	C.vkCmdSetSampleLocationsEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkSampleLocationsInfoEXT)(pSampleLocationsInfo.Raw()))
}

// CmdSetSampleLocationsEnableEXT wraps vkCmdSetSampleLocationsEnableEXT.
func CmdSetSampleLocationsEnableEXT(commandBuffer CommandBuffer, sampleLocationsEnable bool) {
	tmp_sampleLocationsEnable := 0

	if sampleLocationsEnable {
		tmp_sampleLocationsEnable = 1
	}

	C.vkCmdSetSampleLocationsEnableEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_sampleLocationsEnable))
}

// CmdSetSampleMaskEXT wraps vkCmdSetSampleMaskEXT.
func CmdSetSampleMaskEXT(commandBuffer CommandBuffer, samples SampleCountFlags, pSampleMask ffi.Ref[SampleMask]) {
	C.vkCmdSetSampleMaskEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkSampleCountFlagBits(samples), (*C.VkSampleMask)(pSampleMask.Raw()))
}

// CmdSetScissor wraps vkCmdSetScissor.
func CmdSetScissor(commandBuffer CommandBuffer, firstScissor uint32, scissorCount uint32, pScissors Rect2D) {
	C.vkCmdSetScissor(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstScissor), C.uint32_t(scissorCount), (*C.VkRect2D)(pScissors.Raw()))
}

// CmdSetScissorWithCount wraps vkCmdSetScissorWithCount.
func CmdSetScissorWithCount(commandBuffer CommandBuffer, scissorCount uint32, pScissors Rect2D) {
	C.vkCmdSetScissorWithCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(scissorCount), (*C.VkRect2D)(pScissors.Raw()))
}

// CmdSetShadingRateImageEnableNV wraps vkCmdSetShadingRateImageEnableNV.
func CmdSetShadingRateImageEnableNV(commandBuffer CommandBuffer, shadingRateImageEnable bool) {
	tmp_shadingRateImageEnable := 0

	if shadingRateImageEnable {
		tmp_shadingRateImageEnable = 1
	}

	C.vkCmdSetShadingRateImageEnableNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_shadingRateImageEnable))
}

// CmdSetStencilCompareMask wraps vkCmdSetStencilCompareMask.
func CmdSetStencilCompareMask(commandBuffer CommandBuffer, faceMask StencilFaceFlags, compareMask uint32) {
	C.vkCmdSetStencilCompareMask(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkStencilFaceFlags(faceMask), C.uint32_t(compareMask))
}

// CmdSetStencilOp wraps vkCmdSetStencilOp.
func CmdSetStencilOp(commandBuffer CommandBuffer, faceMask StencilFaceFlags, failOp StencilOp, passOp StencilOp, depthFailOp StencilOp, compareOp CompareOp) {
	C.vkCmdSetStencilOp(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkStencilFaceFlags(faceMask), C.VkStencilOp(failOp), C.VkStencilOp(passOp), C.VkStencilOp(depthFailOp), C.VkCompareOp(compareOp))
}

// CmdSetStencilReference wraps vkCmdSetStencilReference.
func CmdSetStencilReference(commandBuffer CommandBuffer, faceMask StencilFaceFlags, reference uint32) {
	C.vkCmdSetStencilReference(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkStencilFaceFlags(faceMask), C.uint32_t(reference))
}

// CmdSetStencilTestEnable wraps vkCmdSetStencilTestEnable.
func CmdSetStencilTestEnable(commandBuffer CommandBuffer, stencilTestEnable bool) {
	tmp_stencilTestEnable := 0

	if stencilTestEnable {
		tmp_stencilTestEnable = 1
	}

	C.vkCmdSetStencilTestEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_stencilTestEnable))
}

// CmdSetStencilWriteMask wraps vkCmdSetStencilWriteMask.
func CmdSetStencilWriteMask(commandBuffer CommandBuffer, faceMask StencilFaceFlags, writeMask uint32) {
	C.vkCmdSetStencilWriteMask(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkStencilFaceFlags(faceMask), C.uint32_t(writeMask))
}

// CmdSetTessellationDomainOriginEXT wraps vkCmdSetTessellationDomainOriginEXT.
func CmdSetTessellationDomainOriginEXT(commandBuffer CommandBuffer, domainOrigin TessellationDomainOrigin) {
	C.vkCmdSetTessellationDomainOriginEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkTessellationDomainOrigin(domainOrigin))
}

// CmdSetVertexInputEXT wraps vkCmdSetVertexInputEXT.
func CmdSetVertexInputEXT(commandBuffer CommandBuffer, vertexBindingDescriptionCount uint32, pVertexBindingDescriptions VertexInputBindingDescription2EXT, vertexAttributeDescriptionCount uint32, pVertexAttributeDescriptions VertexInputAttributeDescription2EXT) {
	C.vkCmdSetVertexInputEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(vertexBindingDescriptionCount), (*C.VkVertexInputBindingDescription2EXT)(pVertexBindingDescriptions.Raw()), C.uint32_t(vertexAttributeDescriptionCount), (*C.VkVertexInputAttributeDescription2EXT)(pVertexAttributeDescriptions.Raw()))
}

// CmdSetViewport wraps vkCmdSetViewport.
func CmdSetViewport(commandBuffer CommandBuffer, firstViewport uint32, viewportCount uint32, pViewports Viewport) {
	C.vkCmdSetViewport(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstViewport), C.uint32_t(viewportCount), (*C.VkViewport)(pViewports.Raw()))
}

// CmdSetViewportShadingRatePaletteNV wraps vkCmdSetViewportShadingRatePaletteNV.
func CmdSetViewportShadingRatePaletteNV(commandBuffer CommandBuffer, firstViewport uint32, viewportCount uint32, pShadingRatePalettes ShadingRatePaletteNV) {
	C.vkCmdSetViewportShadingRatePaletteNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstViewport), C.uint32_t(viewportCount), (*C.VkShadingRatePaletteNV)(pShadingRatePalettes.Raw()))
}

// CmdSetViewportSwizzleNV wraps vkCmdSetViewportSwizzleNV.
func CmdSetViewportSwizzleNV(commandBuffer CommandBuffer, firstViewport uint32, viewportCount uint32, pViewportSwizzles ViewportSwizzleNV) {
	C.vkCmdSetViewportSwizzleNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstViewport), C.uint32_t(viewportCount), (*C.VkViewportSwizzleNV)(pViewportSwizzles.Raw()))
}

// CmdSetViewportWScalingEnableNV wraps vkCmdSetViewportWScalingEnableNV.
func CmdSetViewportWScalingEnableNV(commandBuffer CommandBuffer, viewportWScalingEnable bool) {
	tmp_viewportWScalingEnable := 0

	if viewportWScalingEnable {
		tmp_viewportWScalingEnable = 1
	}

	C.vkCmdSetViewportWScalingEnableNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_viewportWScalingEnable))
}

// CmdSetViewportWScalingNV wraps vkCmdSetViewportWScalingNV.
func CmdSetViewportWScalingNV(commandBuffer CommandBuffer, firstViewport uint32, viewportCount uint32, pViewportWScalings ViewportWScalingNV) {
	C.vkCmdSetViewportWScalingNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstViewport), C.uint32_t(viewportCount), (*C.VkViewportWScalingNV)(pViewportWScalings.Raw()))
}

// CmdSetViewportWithCount wraps vkCmdSetViewportWithCount.
func CmdSetViewportWithCount(commandBuffer CommandBuffer, viewportCount uint32, pViewports Viewport) {
	C.vkCmdSetViewportWithCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(viewportCount), (*C.VkViewport)(pViewports.Raw()))
}

// CmdSubpassShadingHUAWEI wraps vkCmdSubpassShadingHUAWEI.
func CmdSubpassShadingHUAWEI(commandBuffer CommandBuffer) {
	C.vkCmdSubpassShadingHUAWEI(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdTraceRaysIndirect2KHR wraps vkCmdTraceRaysIndirect2KHR.
func CmdTraceRaysIndirect2KHR(commandBuffer CommandBuffer, indirectDeviceAddress DeviceAddress) {
	C.vkCmdTraceRaysIndirect2KHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDeviceAddress(indirectDeviceAddress))
}

// CmdTraceRaysIndirectKHR wraps vkCmdTraceRaysIndirectKHR.
func CmdTraceRaysIndirectKHR(commandBuffer CommandBuffer, pRaygenShaderBindingTable StridedDeviceAddressRegionKHR, pMissShaderBindingTable StridedDeviceAddressRegionKHR, pHitShaderBindingTable StridedDeviceAddressRegionKHR, pCallableShaderBindingTable StridedDeviceAddressRegionKHR, indirectDeviceAddress DeviceAddress) {
	C.vkCmdTraceRaysIndirectKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkStridedDeviceAddressRegionKHR)(pRaygenShaderBindingTable.Raw()), (*C.VkStridedDeviceAddressRegionKHR)(pMissShaderBindingTable.Raw()), (*C.VkStridedDeviceAddressRegionKHR)(pHitShaderBindingTable.Raw()), (*C.VkStridedDeviceAddressRegionKHR)(pCallableShaderBindingTable.Raw()), C.VkDeviceAddress(indirectDeviceAddress))
}

// CmdTraceRaysKHR wraps vkCmdTraceRaysKHR.
func CmdTraceRaysKHR(commandBuffer CommandBuffer, pRaygenShaderBindingTable StridedDeviceAddressRegionKHR, pMissShaderBindingTable StridedDeviceAddressRegionKHR, pHitShaderBindingTable StridedDeviceAddressRegionKHR, pCallableShaderBindingTable StridedDeviceAddressRegionKHR, width uint32, height uint32, depth uint32) {
	C.vkCmdTraceRaysKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), (*C.VkStridedDeviceAddressRegionKHR)(pRaygenShaderBindingTable.Raw()), (*C.VkStridedDeviceAddressRegionKHR)(pMissShaderBindingTable.Raw()), (*C.VkStridedDeviceAddressRegionKHR)(pHitShaderBindingTable.Raw()), (*C.VkStridedDeviceAddressRegionKHR)(pCallableShaderBindingTable.Raw()), C.uint32_t(width), C.uint32_t(height), C.uint32_t(depth))
}

// CmdTraceRaysNV wraps vkCmdTraceRaysNV.
func CmdTraceRaysNV(commandBuffer CommandBuffer, raygenShaderBindingTableBuffer Buffer, raygenShaderBindingOffset DeviceSize, missShaderBindingTableBuffer Buffer, missShaderBindingOffset DeviceSize, missShaderBindingStride DeviceSize, hitShaderBindingTableBuffer Buffer, hitShaderBindingOffset DeviceSize, hitShaderBindingStride DeviceSize, callableShaderBindingTableBuffer Buffer, callableShaderBindingOffset DeviceSize, callableShaderBindingStride DeviceSize, width uint32, height uint32, depth uint32) {
	C.vkCmdTraceRaysNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(raygenShaderBindingTableBuffer), C.VkDeviceSize(raygenShaderBindingOffset), C.VkBuffer(missShaderBindingTableBuffer), C.VkDeviceSize(missShaderBindingOffset), C.VkDeviceSize(missShaderBindingStride), C.VkBuffer(hitShaderBindingTableBuffer), C.VkDeviceSize(hitShaderBindingOffset), C.VkDeviceSize(hitShaderBindingStride), C.VkBuffer(callableShaderBindingTableBuffer), C.VkDeviceSize(callableShaderBindingOffset), C.VkDeviceSize(callableShaderBindingStride), C.uint32_t(width), C.uint32_t(height), C.uint32_t(depth))
}

// CmdUpdateBuffer wraps vkCmdUpdateBuffer.
func CmdUpdateBuffer(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset DeviceSize, dataSize DeviceSize, pData unsafe.Pointer) {
	C.vkCmdUpdateBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.VkDeviceSize(dataSize), pData)
}

// CmdUpdatePipelineIndirectBufferNV wraps vkCmdUpdatePipelineIndirectBufferNV.
func CmdUpdatePipelineIndirectBufferNV(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	C.vkCmdUpdatePipelineIndirectBufferNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipeline(pipeline))
}

// CmdWaitEvents wraps vkCmdWaitEvents.
func CmdWaitEvents(commandBuffer CommandBuffer, eventCount uint32, pEvents ffi.Ref[Event], srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarrierCount uint32, pMemoryBarriers MemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers BufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers ImageMemoryBarrier) {
	C.vkCmdWaitEvents(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(eventCount), (*C.VkEvent)(pEvents.Raw()), C.VkPipelineStageFlags(srcStageMask), C.VkPipelineStageFlags(dstStageMask), C.uint32_t(memoryBarrierCount), (*C.VkMemoryBarrier)(pMemoryBarriers.Raw()), C.uint32_t(bufferMemoryBarrierCount), (*C.VkBufferMemoryBarrier)(pBufferMemoryBarriers.Raw()), C.uint32_t(imageMemoryBarrierCount), (*C.VkImageMemoryBarrier)(pImageMemoryBarriers.Raw()))
}

// CmdWaitEvents2 wraps vkCmdWaitEvents2.
func CmdWaitEvents2(commandBuffer CommandBuffer, eventCount uint32, pEvents ffi.Ref[Event], pDependencyInfos DependencyInfo) {
	C.vkCmdWaitEvents2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(eventCount), (*C.VkEvent)(pEvents.Raw()), (*C.VkDependencyInfo)(pDependencyInfos.Raw()))
}

// CmdWriteAccelerationStructuresPropertiesKHR wraps vkCmdWriteAccelerationStructuresPropertiesKHR.
func CmdWriteAccelerationStructuresPropertiesKHR(commandBuffer CommandBuffer, accelerationStructureCount uint32, pAccelerationStructures ffi.Ref[AccelerationStructureKHR], queryType QueryType, queryPool QueryPool, firstQuery uint32) {
	C.vkCmdWriteAccelerationStructuresPropertiesKHR(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(accelerationStructureCount), (*C.VkAccelerationStructureKHR)(pAccelerationStructures.Raw()), C.VkQueryType(queryType), C.VkQueryPool(queryPool), C.uint32_t(firstQuery))
}

// CmdWriteAccelerationStructuresPropertiesNV wraps vkCmdWriteAccelerationStructuresPropertiesNV.
func CmdWriteAccelerationStructuresPropertiesNV(commandBuffer CommandBuffer, accelerationStructureCount uint32, pAccelerationStructures ffi.Ref[AccelerationStructureNV], queryType QueryType, queryPool QueryPool, firstQuery uint32) {
	C.vkCmdWriteAccelerationStructuresPropertiesNV(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(accelerationStructureCount), (*C.VkAccelerationStructureNV)(pAccelerationStructures.Raw()), C.VkQueryType(queryType), C.VkQueryPool(queryPool), C.uint32_t(firstQuery))
}

// CmdWriteBufferMarker2AMD wraps vkCmdWriteBufferMarker2AMD.
func CmdWriteBufferMarker2AMD(commandBuffer CommandBuffer, stage PipelineStageFlags2, dstBuffer Buffer, dstOffset DeviceSize, marker uint32) {
	C.vkCmdWriteBufferMarker2AMD(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlags2(stage), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.uint32_t(marker))
}

// CmdWriteBufferMarkerAMD wraps vkCmdWriteBufferMarkerAMD.
func CmdWriteBufferMarkerAMD(commandBuffer CommandBuffer, pipelineStage PipelineStageFlags, dstBuffer Buffer, dstOffset DeviceSize, marker uint32) {
	C.vkCmdWriteBufferMarkerAMD(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlagBits(pipelineStage), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.uint32_t(marker))
}

// CmdWriteMicromapsPropertiesEXT wraps vkCmdWriteMicromapsPropertiesEXT.
func CmdWriteMicromapsPropertiesEXT(commandBuffer CommandBuffer, micromapCount uint32, pMicromaps ffi.Ref[MicromapEXT], queryType QueryType, queryPool QueryPool, firstQuery uint32) {
	C.vkCmdWriteMicromapsPropertiesEXT(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(micromapCount), (*C.VkMicromapEXT)(pMicromaps.Raw()), C.VkQueryType(queryType), C.VkQueryPool(queryPool), C.uint32_t(firstQuery))
}

// CmdWriteTimestamp wraps vkCmdWriteTimestamp.
func CmdWriteTimestamp(commandBuffer CommandBuffer, pipelineStage PipelineStageFlags, queryPool QueryPool, query uint32) {
	C.vkCmdWriteTimestamp(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlagBits(pipelineStage), C.VkQueryPool(queryPool), C.uint32_t(query))
}

// CmdWriteTimestamp2 wraps vkCmdWriteTimestamp2.
func CmdWriteTimestamp2(commandBuffer CommandBuffer, stage PipelineStageFlags2, queryPool QueryPool, query uint32) {
	C.vkCmdWriteTimestamp2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlags2(stage), C.VkQueryPool(queryPool), C.uint32_t(query))
}

// CompileDeferredNV wraps vkCompileDeferredNV.
func CompileDeferredNV(device Device, pipeline Pipeline, shader uint32) Result {
	ret := C.vkCompileDeferredNV(C.VkDevice(unsafe.Pointer(device)), C.VkPipeline(pipeline), C.uint32_t(shader))

	return Result(ret)
}

// ConvertCooperativeVectorMatrixNV wraps vkConvertCooperativeVectorMatrixNV.
func ConvertCooperativeVectorMatrixNV(device Device, pInfo ConvertCooperativeVectorMatrixInfoNV) Result {
	ret := C.vkConvertCooperativeVectorMatrixNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkConvertCooperativeVectorMatrixInfoNV)(pInfo.Raw()))

	return Result(ret)
}

// CopyAccelerationStructureKHR wraps vkCopyAccelerationStructureKHR.
func CopyAccelerationStructureKHR(device Device, deferredOperation DeferredOperationKHR, pInfo CopyAccelerationStructureInfoKHR) Result {
	ret := C.vkCopyAccelerationStructureKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), (*C.VkCopyAccelerationStructureInfoKHR)(pInfo.Raw()))

	return Result(ret)
}

// CopyAccelerationStructureToMemoryKHR wraps vkCopyAccelerationStructureToMemoryKHR.
func CopyAccelerationStructureToMemoryKHR(device Device, deferredOperation DeferredOperationKHR, pInfo CopyAccelerationStructureToMemoryInfoKHR) Result {
	ret := C.vkCopyAccelerationStructureToMemoryKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), (*C.VkCopyAccelerationStructureToMemoryInfoKHR)(pInfo.Raw()))

	return Result(ret)
}

// CopyImageToImage wraps vkCopyImageToImage.
func CopyImageToImage(device Device, pCopyImageToImageInfo CopyImageToImageInfo) Result {
	ret := C.vkCopyImageToImage(C.VkDevice(unsafe.Pointer(device)), (*C.VkCopyImageToImageInfo)(pCopyImageToImageInfo.Raw()))

	return Result(ret)
}

// CopyImageToMemory wraps vkCopyImageToMemory.
func CopyImageToMemory(device Device, pCopyImageToMemoryInfo CopyImageToMemoryInfo) Result {
	ret := C.vkCopyImageToMemory(C.VkDevice(unsafe.Pointer(device)), (*C.VkCopyImageToMemoryInfo)(pCopyImageToMemoryInfo.Raw()))

	return Result(ret)
}

// CopyMemoryToAccelerationStructureKHR wraps vkCopyMemoryToAccelerationStructureKHR.
func CopyMemoryToAccelerationStructureKHR(device Device, deferredOperation DeferredOperationKHR, pInfo CopyMemoryToAccelerationStructureInfoKHR) Result {
	ret := C.vkCopyMemoryToAccelerationStructureKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), (*C.VkCopyMemoryToAccelerationStructureInfoKHR)(pInfo.Raw()))

	return Result(ret)
}

// CopyMemoryToImage wraps vkCopyMemoryToImage.
func CopyMemoryToImage(device Device, pCopyMemoryToImageInfo CopyMemoryToImageInfo) Result {
	ret := C.vkCopyMemoryToImage(C.VkDevice(unsafe.Pointer(device)), (*C.VkCopyMemoryToImageInfo)(pCopyMemoryToImageInfo.Raw()))

	return Result(ret)
}

// CopyMemoryToMicromapEXT wraps vkCopyMemoryToMicromapEXT.
func CopyMemoryToMicromapEXT(device Device, deferredOperation DeferredOperationKHR, pInfo CopyMemoryToMicromapInfoEXT) Result {
	ret := C.vkCopyMemoryToMicromapEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), (*C.VkCopyMemoryToMicromapInfoEXT)(pInfo.Raw()))

	return Result(ret)
}

// CopyMicromapEXT wraps vkCopyMicromapEXT.
func CopyMicromapEXT(device Device, deferredOperation DeferredOperationKHR, pInfo CopyMicromapInfoEXT) Result {
	ret := C.vkCopyMicromapEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), (*C.VkCopyMicromapInfoEXT)(pInfo.Raw()))

	return Result(ret)
}

// CopyMicromapToMemoryEXT wraps vkCopyMicromapToMemoryEXT.
func CopyMicromapToMemoryEXT(device Device, deferredOperation DeferredOperationKHR, pInfo CopyMicromapToMemoryInfoEXT) Result {
	ret := C.vkCopyMicromapToMemoryEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), (*C.VkCopyMicromapToMemoryInfoEXT)(pInfo.Raw()))

	return Result(ret)
}

// CreateAccelerationStructureKHR wraps vkCreateAccelerationStructureKHR.
func CreateAccelerationStructureKHR(device Device, pCreateInfo AccelerationStructureCreateInfoKHR, pAllocator AllocationCallbacks, pAccelerationStructure ffi.Ref[AccelerationStructureKHR]) Result {
	ret := C.vkCreateAccelerationStructureKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkAccelerationStructureCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkAccelerationStructureKHR)(pAccelerationStructure.Raw()))

	return Result(ret)
}

// CreateAccelerationStructureNV wraps vkCreateAccelerationStructureNV.
func CreateAccelerationStructureNV(device Device, pCreateInfo AccelerationStructureCreateInfoNV, pAllocator AllocationCallbacks, pAccelerationStructure ffi.Ref[AccelerationStructureNV]) Result {
	ret := C.vkCreateAccelerationStructureNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkAccelerationStructureCreateInfoNV)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkAccelerationStructureNV)(pAccelerationStructure.Raw()))

	return Result(ret)
}

// CreateBuffer wraps vkCreateBuffer.
func CreateBuffer(device Device, pCreateInfo BufferCreateInfo, pAllocator AllocationCallbacks, pBuffer ffi.Ref[Buffer]) Result {
	ret := C.vkCreateBuffer(C.VkDevice(unsafe.Pointer(device)), (*C.VkBufferCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkBuffer)(pBuffer.Raw()))

	return Result(ret)
}

// CreateBufferView wraps vkCreateBufferView.
func CreateBufferView(device Device, pCreateInfo BufferViewCreateInfo, pAllocator AllocationCallbacks, pView ffi.Ref[BufferView]) Result {
	ret := C.vkCreateBufferView(C.VkDevice(unsafe.Pointer(device)), (*C.VkBufferViewCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkBufferView)(pView.Raw()))

	return Result(ret)
}

// CreateCommandPool wraps vkCreateCommandPool.
func CreateCommandPool(device Device, pCreateInfo CommandPoolCreateInfo, pAllocator AllocationCallbacks, pCommandPool ffi.Ref[CommandPool]) Result {
	ret := C.vkCreateCommandPool(C.VkDevice(unsafe.Pointer(device)), (*C.VkCommandPoolCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkCommandPool)(pCommandPool.Raw()))

	return Result(ret)
}

// CreateComputePipelines wraps vkCreateComputePipelines.
func CreateComputePipelines(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos ComputePipelineCreateInfo, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateComputePipelines(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), (*C.VkComputePipelineCreateInfo)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateCuFunctionNVX wraps vkCreateCuFunctionNVX.
func CreateCuFunctionNVX(device Device, pCreateInfo CuFunctionCreateInfoNVX, pAllocator AllocationCallbacks, pFunction ffi.Ref[CuFunctionNVX]) Result {
	ret := C.vkCreateCuFunctionNVX(C.VkDevice(unsafe.Pointer(device)), (*C.VkCuFunctionCreateInfoNVX)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkCuFunctionNVX)(pFunction.Raw()))

	return Result(ret)
}

// CreateCuModuleNVX wraps vkCreateCuModuleNVX.
func CreateCuModuleNVX(device Device, pCreateInfo CuModuleCreateInfoNVX, pAllocator AllocationCallbacks, pModule ffi.Ref[CuModuleNVX]) Result {
	ret := C.vkCreateCuModuleNVX(C.VkDevice(unsafe.Pointer(device)), (*C.VkCuModuleCreateInfoNVX)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkCuModuleNVX)(pModule.Raw()))

	return Result(ret)
}

// CreateDataGraphPipelineSessionARM wraps vkCreateDataGraphPipelineSessionARM.
func CreateDataGraphPipelineSessionARM(device Device, pCreateInfo DataGraphPipelineSessionCreateInfoARM, pAllocator AllocationCallbacks, pSession ffi.Ref[DataGraphPipelineSessionARM]) Result {
	ret := C.vkCreateDataGraphPipelineSessionARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkDataGraphPipelineSessionCreateInfoARM)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDataGraphPipelineSessionARM)(pSession.Raw()))

	return Result(ret)
}

// CreateDataGraphPipelinesARM wraps vkCreateDataGraphPipelinesARM.
func CreateDataGraphPipelinesARM(device Device, deferredOperation DeferredOperationKHR, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos DataGraphPipelineCreateInfoARM, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateDataGraphPipelinesARM(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), (*C.VkDataGraphPipelineCreateInfoARM)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateDebugReportCallbackEXT wraps vkCreateDebugReportCallbackEXT.
func CreateDebugReportCallbackEXT(instance Instance, pCreateInfo DebugReportCallbackCreateInfoEXT, pAllocator AllocationCallbacks, pCallback ffi.Ref[DebugReportCallbackEXT]) Result {
	ret := C.vkCreateDebugReportCallbackEXT(C.VkInstance(unsafe.Pointer(instance)), (*C.VkDebugReportCallbackCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDebugReportCallbackEXT)(pCallback.Raw()))

	return Result(ret)
}

// CreateDebugUtilsMessengerEXT wraps vkCreateDebugUtilsMessengerEXT.
func CreateDebugUtilsMessengerEXT(instance Instance, pCreateInfo DebugUtilsMessengerCreateInfoEXT, pAllocator AllocationCallbacks, pMessenger ffi.Ref[DebugUtilsMessengerEXT]) Result {
	ret := C.vkCreateDebugUtilsMessengerEXT(C.VkInstance(unsafe.Pointer(instance)), (*C.VkDebugUtilsMessengerCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDebugUtilsMessengerEXT)(pMessenger.Raw()))

	return Result(ret)
}

// CreateDeferredOperationKHR wraps vkCreateDeferredOperationKHR.
func CreateDeferredOperationKHR(device Device, pAllocator AllocationCallbacks, pDeferredOperation ffi.Ref[DeferredOperationKHR]) Result {
	ret := C.vkCreateDeferredOperationKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDeferredOperationKHR)(pDeferredOperation.Raw()))

	return Result(ret)
}

// CreateDescriptorPool wraps vkCreateDescriptorPool.
func CreateDescriptorPool(device Device, pCreateInfo DescriptorPoolCreateInfo, pAllocator AllocationCallbacks, pDescriptorPool ffi.Ref[DescriptorPool]) Result {
	ret := C.vkCreateDescriptorPool(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorPoolCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDescriptorPool)(pDescriptorPool.Raw()))

	return Result(ret)
}

// CreateDescriptorSetLayout wraps vkCreateDescriptorSetLayout.
func CreateDescriptorSetLayout(device Device, pCreateInfo DescriptorSetLayoutCreateInfo, pAllocator AllocationCallbacks, pSetLayout ffi.Ref[DescriptorSetLayout]) Result {
	ret := C.vkCreateDescriptorSetLayout(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDescriptorSetLayout)(pSetLayout.Raw()))

	return Result(ret)
}

// CreateDescriptorUpdateTemplate wraps vkCreateDescriptorUpdateTemplate.
func CreateDescriptorUpdateTemplate(device Device, pCreateInfo DescriptorUpdateTemplateCreateInfo, pAllocator AllocationCallbacks, pDescriptorUpdateTemplate ffi.Ref[DescriptorUpdateTemplate]) Result {
	ret := C.vkCreateDescriptorUpdateTemplate(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorUpdateTemplateCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDescriptorUpdateTemplate)(pDescriptorUpdateTemplate.Raw()))

	return Result(ret)
}

// CreateDevice wraps vkCreateDevice.
func CreateDevice(physicalDevice PhysicalDevice, pCreateInfo DeviceCreateInfo, pAllocator AllocationCallbacks, pDevice ffi.Ref[Device]) Result {
	ret := C.vkCreateDevice(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkDeviceCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDevice)(pDevice.Raw()))

	return Result(ret)
}

// CreateDisplayModeKHR wraps vkCreateDisplayModeKHR.
func CreateDisplayModeKHR(physicalDevice PhysicalDevice, display DisplayKHR, pCreateInfo DisplayModeCreateInfoKHR, pAllocator AllocationCallbacks, pMode ffi.Ref[DisplayModeKHR]) Result {
	ret := C.vkCreateDisplayModeKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkDisplayKHR(display), (*C.VkDisplayModeCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkDisplayModeKHR)(pMode.Raw()))

	return Result(ret)
}

// CreateDisplayPlaneSurfaceKHR wraps vkCreateDisplayPlaneSurfaceKHR.
func CreateDisplayPlaneSurfaceKHR(instance Instance, pCreateInfo DisplaySurfaceCreateInfoKHR, pAllocator AllocationCallbacks, pSurface ffi.Ref[SurfaceKHR]) Result {
	ret := C.vkCreateDisplayPlaneSurfaceKHR(C.VkInstance(unsafe.Pointer(instance)), (*C.VkDisplaySurfaceCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSurfaceKHR)(pSurface.Raw()))

	return Result(ret)
}

// CreateEvent wraps vkCreateEvent.
func CreateEvent(device Device, pCreateInfo EventCreateInfo, pAllocator AllocationCallbacks, pEvent ffi.Ref[Event]) Result {
	ret := C.vkCreateEvent(C.VkDevice(unsafe.Pointer(device)), (*C.VkEventCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkEvent)(pEvent.Raw()))

	return Result(ret)
}

// CreateExternalComputeQueueNV wraps vkCreateExternalComputeQueueNV.
func CreateExternalComputeQueueNV(device Device, pCreateInfo ExternalComputeQueueCreateInfoNV, pAllocator AllocationCallbacks, pExternalQueue ffi.Ref[ExternalComputeQueueNV]) Result {
	ret := C.vkCreateExternalComputeQueueNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkExternalComputeQueueCreateInfoNV)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkExternalComputeQueueNV)(pExternalQueue.Raw()))

	return Result(ret)
}

// CreateFence wraps vkCreateFence.
func CreateFence(device Device, pCreateInfo FenceCreateInfo, pAllocator AllocationCallbacks, pFence ffi.Ref[Fence]) Result {
	ret := C.vkCreateFence(C.VkDevice(unsafe.Pointer(device)), (*C.VkFenceCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkFence)(pFence.Raw()))

	return Result(ret)
}

// CreateFramebuffer wraps vkCreateFramebuffer.
func CreateFramebuffer(device Device, pCreateInfo FramebufferCreateInfo, pAllocator AllocationCallbacks, pFramebuffer ffi.Ref[Framebuffer]) Result {
	ret := C.vkCreateFramebuffer(C.VkDevice(unsafe.Pointer(device)), (*C.VkFramebufferCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkFramebuffer)(pFramebuffer.Raw()))

	return Result(ret)
}

// CreateGraphicsPipelines wraps vkCreateGraphicsPipelines.
func CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos GraphicsPipelineCreateInfo, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateGraphicsPipelines(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), (*C.VkGraphicsPipelineCreateInfo)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateHeadlessSurfaceEXT wraps vkCreateHeadlessSurfaceEXT.
func CreateHeadlessSurfaceEXT(instance Instance, pCreateInfo HeadlessSurfaceCreateInfoEXT, pAllocator AllocationCallbacks, pSurface ffi.Ref[SurfaceKHR]) Result {
	ret := C.vkCreateHeadlessSurfaceEXT(C.VkInstance(unsafe.Pointer(instance)), (*C.VkHeadlessSurfaceCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSurfaceKHR)(pSurface.Raw()))

	return Result(ret)
}

// CreateImage wraps vkCreateImage.
func CreateImage(device Device, pCreateInfo ImageCreateInfo, pAllocator AllocationCallbacks, pImage ffi.Ref[Image]) Result {
	ret := C.vkCreateImage(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkImage)(pImage.Raw()))

	return Result(ret)
}

// CreateImageView wraps vkCreateImageView.
func CreateImageView(device Device, pCreateInfo ImageViewCreateInfo, pAllocator AllocationCallbacks, pView ffi.Ref[ImageView]) Result {
	ret := C.vkCreateImageView(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageViewCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkImageView)(pView.Raw()))

	return Result(ret)
}

// CreateIndirectCommandsLayoutEXT wraps vkCreateIndirectCommandsLayoutEXT.
func CreateIndirectCommandsLayoutEXT(device Device, pCreateInfo IndirectCommandsLayoutCreateInfoEXT, pAllocator AllocationCallbacks, pIndirectCommandsLayout ffi.Ref[IndirectCommandsLayoutEXT]) Result {
	ret := C.vkCreateIndirectCommandsLayoutEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkIndirectCommandsLayoutCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkIndirectCommandsLayoutEXT)(pIndirectCommandsLayout.Raw()))

	return Result(ret)
}

// CreateIndirectCommandsLayoutNV wraps vkCreateIndirectCommandsLayoutNV.
func CreateIndirectCommandsLayoutNV(device Device, pCreateInfo IndirectCommandsLayoutCreateInfoNV, pAllocator AllocationCallbacks, pIndirectCommandsLayout ffi.Ref[IndirectCommandsLayoutNV]) Result {
	ret := C.vkCreateIndirectCommandsLayoutNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkIndirectCommandsLayoutCreateInfoNV)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkIndirectCommandsLayoutNV)(pIndirectCommandsLayout.Raw()))

	return Result(ret)
}

// CreateIndirectExecutionSetEXT wraps vkCreateIndirectExecutionSetEXT.
func CreateIndirectExecutionSetEXT(device Device, pCreateInfo IndirectExecutionSetCreateInfoEXT, pAllocator AllocationCallbacks, pIndirectExecutionSet ffi.Ref[IndirectExecutionSetEXT]) Result {
	ret := C.vkCreateIndirectExecutionSetEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkIndirectExecutionSetCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkIndirectExecutionSetEXT)(pIndirectExecutionSet.Raw()))

	return Result(ret)
}

// CreateInstance wraps vkCreateInstance.
func CreateInstance(pCreateInfo InstanceCreateInfo, pAllocator AllocationCallbacks, pInstance ffi.Ref[Instance]) Result {
	ret := C.vkCreateInstance((*C.VkInstanceCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkInstance)(pInstance.Raw()))

	return Result(ret)
}

// CreateMetalSurfaceEXT wraps vkCreateMetalSurfaceEXT.
func CreateMetalSurfaceEXT(instance Instance, pCreateInfo MetalSurfaceCreateInfoEXT, pAllocator AllocationCallbacks, pSurface ffi.Ref[SurfaceKHR]) Result {
	ret := C.vkCreateMetalSurfaceEXT(C.VkInstance(unsafe.Pointer(instance)), (*C.VkMetalSurfaceCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSurfaceKHR)(pSurface.Raw()))

	return Result(ret)
}

// CreateMicromapEXT wraps vkCreateMicromapEXT.
func CreateMicromapEXT(device Device, pCreateInfo MicromapCreateInfoEXT, pAllocator AllocationCallbacks, pMicromap ffi.Ref[MicromapEXT]) Result {
	ret := C.vkCreateMicromapEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkMicromapCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkMicromapEXT)(pMicromap.Raw()))

	return Result(ret)
}

// CreateOpticalFlowSessionNV wraps vkCreateOpticalFlowSessionNV.
func CreateOpticalFlowSessionNV(device Device, pCreateInfo OpticalFlowSessionCreateInfoNV, pAllocator AllocationCallbacks, pSession ffi.Ref[OpticalFlowSessionNV]) Result {
	ret := C.vkCreateOpticalFlowSessionNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkOpticalFlowSessionCreateInfoNV)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkOpticalFlowSessionNV)(pSession.Raw()))

	return Result(ret)
}

// CreatePipelineBinariesKHR wraps vkCreatePipelineBinariesKHR.
func CreatePipelineBinariesKHR(device Device, pCreateInfo PipelineBinaryCreateInfoKHR, pAllocator AllocationCallbacks, pBinaries PipelineBinaryHandlesInfoKHR) Result {
	ret := C.vkCreatePipelineBinariesKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineBinaryCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipelineBinaryHandlesInfoKHR)(pBinaries.Raw()))

	return Result(ret)
}

// CreatePipelineCache wraps vkCreatePipelineCache.
func CreatePipelineCache(device Device, pCreateInfo PipelineCacheCreateInfo, pAllocator AllocationCallbacks, pPipelineCache ffi.Ref[PipelineCache]) Result {
	ret := C.vkCreatePipelineCache(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineCacheCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipelineCache)(pPipelineCache.Raw()))

	return Result(ret)
}

// CreatePipelineLayout wraps vkCreatePipelineLayout.
func CreatePipelineLayout(device Device, pCreateInfo PipelineLayoutCreateInfo, pAllocator AllocationCallbacks, pPipelineLayout ffi.Ref[PipelineLayout]) Result {
	ret := C.vkCreatePipelineLayout(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineLayoutCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipelineLayout)(pPipelineLayout.Raw()))

	return Result(ret)
}

// CreatePrivateDataSlot wraps vkCreatePrivateDataSlot.
func CreatePrivateDataSlot(device Device, pCreateInfo PrivateDataSlotCreateInfo, pAllocator AllocationCallbacks, pPrivateDataSlot ffi.Ref[PrivateDataSlot]) Result {
	ret := C.vkCreatePrivateDataSlot(C.VkDevice(unsafe.Pointer(device)), (*C.VkPrivateDataSlotCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPrivateDataSlot)(pPrivateDataSlot.Raw()))

	return Result(ret)
}

// CreateQueryPool wraps vkCreateQueryPool.
func CreateQueryPool(device Device, pCreateInfo QueryPoolCreateInfo, pAllocator AllocationCallbacks, pQueryPool ffi.Ref[QueryPool]) Result {
	ret := C.vkCreateQueryPool(C.VkDevice(unsafe.Pointer(device)), (*C.VkQueryPoolCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkQueryPool)(pQueryPool.Raw()))

	return Result(ret)
}

// CreateRayTracingPipelinesKHR wraps vkCreateRayTracingPipelinesKHR.
func CreateRayTracingPipelinesKHR(device Device, deferredOperation DeferredOperationKHR, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos RayTracingPipelineCreateInfoKHR, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateRayTracingPipelinesKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(deferredOperation), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), (*C.VkRayTracingPipelineCreateInfoKHR)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateRayTracingPipelinesNV wraps vkCreateRayTracingPipelinesNV.
func CreateRayTracingPipelinesNV(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos RayTracingPipelineCreateInfoNV, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateRayTracingPipelinesNV(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), (*C.VkRayTracingPipelineCreateInfoNV)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateRenderPass wraps vkCreateRenderPass.
func CreateRenderPass(device Device, pCreateInfo RenderPassCreateInfo, pAllocator AllocationCallbacks, pRenderPass ffi.Ref[RenderPass]) Result {
	ret := C.vkCreateRenderPass(C.VkDevice(unsafe.Pointer(device)), (*C.VkRenderPassCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkRenderPass)(pRenderPass.Raw()))

	return Result(ret)
}

// CreateRenderPass2 wraps vkCreateRenderPass2.
func CreateRenderPass2(device Device, pCreateInfo RenderPassCreateInfo2, pAllocator AllocationCallbacks, pRenderPass ffi.Ref[RenderPass]) Result {
	ret := C.vkCreateRenderPass2(C.VkDevice(unsafe.Pointer(device)), (*C.VkRenderPassCreateInfo2)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkRenderPass)(pRenderPass.Raw()))

	return Result(ret)
}

// CreateSampler wraps vkCreateSampler.
func CreateSampler(device Device, pCreateInfo SamplerCreateInfo, pAllocator AllocationCallbacks, pSampler ffi.Ref[Sampler]) Result {
	ret := C.vkCreateSampler(C.VkDevice(unsafe.Pointer(device)), (*C.VkSamplerCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSampler)(pSampler.Raw()))

	return Result(ret)
}

// CreateSamplerYcbcrConversion wraps vkCreateSamplerYcbcrConversion.
func CreateSamplerYcbcrConversion(device Device, pCreateInfo SamplerYcbcrConversionCreateInfo, pAllocator AllocationCallbacks, pYcbcrConversion ffi.Ref[SamplerYcbcrConversion]) Result {
	ret := C.vkCreateSamplerYcbcrConversion(C.VkDevice(unsafe.Pointer(device)), (*C.VkSamplerYcbcrConversionCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSamplerYcbcrConversion)(pYcbcrConversion.Raw()))

	return Result(ret)
}

// CreateSemaphore wraps vkCreateSemaphore.
func CreateSemaphore(device Device, pCreateInfo SemaphoreCreateInfo, pAllocator AllocationCallbacks, pSemaphore ffi.Ref[Semaphore]) Result {
	ret := C.vkCreateSemaphore(C.VkDevice(unsafe.Pointer(device)), (*C.VkSemaphoreCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSemaphore)(pSemaphore.Raw()))

	return Result(ret)
}

// CreateShaderModule wraps vkCreateShaderModule.
func CreateShaderModule(device Device, pCreateInfo ShaderModuleCreateInfo, pAllocator AllocationCallbacks, pShaderModule ffi.Ref[ShaderModule]) Result {
	ret := C.vkCreateShaderModule(C.VkDevice(unsafe.Pointer(device)), (*C.VkShaderModuleCreateInfo)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkShaderModule)(pShaderModule.Raw()))

	return Result(ret)
}

// CreateShadersEXT wraps vkCreateShadersEXT.
func CreateShadersEXT(device Device, createInfoCount uint32, pCreateInfos ShaderCreateInfoEXT, pAllocator AllocationCallbacks, pShaders ffi.Ref[ShaderEXT]) Result {
	ret := C.vkCreateShadersEXT(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(createInfoCount), (*C.VkShaderCreateInfoEXT)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkShaderEXT)(pShaders.Raw()))

	return Result(ret)
}

// CreateSharedSwapchainsKHR wraps vkCreateSharedSwapchainsKHR.
func CreateSharedSwapchainsKHR(device Device, swapchainCount uint32, pCreateInfos SwapchainCreateInfoKHR, pAllocator AllocationCallbacks, pSwapchains ffi.Ref[SwapchainKHR]) Result {
	ret := C.vkCreateSharedSwapchainsKHR(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(swapchainCount), (*C.VkSwapchainCreateInfoKHR)(pCreateInfos.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSwapchainKHR)(pSwapchains.Raw()))

	return Result(ret)
}

// CreateSwapchainKHR wraps vkCreateSwapchainKHR.
func CreateSwapchainKHR(device Device, pCreateInfo SwapchainCreateInfoKHR, pAllocator AllocationCallbacks, pSwapchain ffi.Ref[SwapchainKHR]) Result {
	ret := C.vkCreateSwapchainKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkSwapchainCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSwapchainKHR)(pSwapchain.Raw()))

	return Result(ret)
}

// CreateTensorARM wraps vkCreateTensorARM.
func CreateTensorARM(device Device, pCreateInfo TensorCreateInfoARM, pAllocator AllocationCallbacks, pTensor ffi.Ref[TensorARM]) Result {
	ret := C.vkCreateTensorARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkTensorCreateInfoARM)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkTensorARM)(pTensor.Raw()))

	return Result(ret)
}

// CreateTensorViewARM wraps vkCreateTensorViewARM.
func CreateTensorViewARM(device Device, pCreateInfo TensorViewCreateInfoARM, pAllocator AllocationCallbacks, pView ffi.Ref[TensorViewARM]) Result {
	ret := C.vkCreateTensorViewARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkTensorViewCreateInfoARM)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkTensorViewARM)(pView.Raw()))

	return Result(ret)
}

// CreateValidationCacheEXT wraps vkCreateValidationCacheEXT.
func CreateValidationCacheEXT(device Device, pCreateInfo ValidationCacheCreateInfoEXT, pAllocator AllocationCallbacks, pValidationCache ffi.Ref[ValidationCacheEXT]) Result {
	ret := C.vkCreateValidationCacheEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkValidationCacheCreateInfoEXT)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkValidationCacheEXT)(pValidationCache.Raw()))

	return Result(ret)
}

// CreateVideoSessionKHR wraps vkCreateVideoSessionKHR.
func CreateVideoSessionKHR(device Device, pCreateInfo VideoSessionCreateInfoKHR, pAllocator AllocationCallbacks, pVideoSession ffi.Ref[VideoSessionKHR]) Result {
	ret := C.vkCreateVideoSessionKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkVideoSessionCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkVideoSessionKHR)(pVideoSession.Raw()))

	return Result(ret)
}

// CreateVideoSessionParametersKHR wraps vkCreateVideoSessionParametersKHR.
func CreateVideoSessionParametersKHR(device Device, pCreateInfo VideoSessionParametersCreateInfoKHR, pAllocator AllocationCallbacks, pVideoSessionParameters ffi.Ref[VideoSessionParametersKHR]) Result {
	ret := C.vkCreateVideoSessionParametersKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkVideoSessionParametersCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkVideoSessionParametersKHR)(pVideoSessionParameters.Raw()))

	return Result(ret)
}

// CreateWin32SurfaceKHR wraps vkCreateWin32SurfaceKHR.
func CreateWin32SurfaceKHR(instance Instance, pCreateInfo Win32SurfaceCreateInfoKHR, pAllocator AllocationCallbacks, pSurface ffi.Ref[SurfaceKHR]) Result {
	ret := C.vkCreateWin32SurfaceKHR(C.VkInstance(unsafe.Pointer(instance)), (*C.VkWin32SurfaceCreateInfoKHR)(pCreateInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkSurfaceKHR)(pSurface.Raw()))

	return Result(ret)
}

// DebugMarkerSetObjectNameEXT wraps vkDebugMarkerSetObjectNameEXT.
func DebugMarkerSetObjectNameEXT(device Device, pNameInfo DebugMarkerObjectNameInfoEXT) Result {
	ret := C.vkDebugMarkerSetObjectNameEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDebugMarkerObjectNameInfoEXT)(pNameInfo.Raw()))

	return Result(ret)
}

// DebugMarkerSetObjectTagEXT wraps vkDebugMarkerSetObjectTagEXT.
func DebugMarkerSetObjectTagEXT(device Device, pTagInfo DebugMarkerObjectTagInfoEXT) Result {
	ret := C.vkDebugMarkerSetObjectTagEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDebugMarkerObjectTagInfoEXT)(pTagInfo.Raw()))

	return Result(ret)
}

// DebugReportMessageEXT wraps vkDebugReportMessageEXT.
func DebugReportMessageEXT(instance Instance, flags DebugReportFlagsEXT, objectType DebugReportObjectTypeEXT, object uint64, location uintptr, messageCode int32, pLayerPrefix ffi.CString, pMessage ffi.CString) {
	C.vkDebugReportMessageEXT(C.VkInstance(unsafe.Pointer(instance)), C.VkDebugReportFlagsEXT(flags), C.VkDebugReportObjectTypeEXT(objectType), C.uint64_t(object), C.size_t(location), C.int32_t(messageCode), (*C.char)(pLayerPrefix.Raw()), (*C.char)(pMessage.Raw()))
}

// DeferredOperationJoinKHR wraps vkDeferredOperationJoinKHR.
func DeferredOperationJoinKHR(device Device, operation DeferredOperationKHR) Result {
	ret := C.vkDeferredOperationJoinKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(operation))

	return Result(ret)
}

// DestroyAccelerationStructureKHR wraps vkDestroyAccelerationStructureKHR.
func DestroyAccelerationStructureKHR(device Device, accelerationStructure AccelerationStructureKHR, pAllocator AllocationCallbacks) {
	C.vkDestroyAccelerationStructureKHR(C.VkDevice(unsafe.Pointer(device)), C.VkAccelerationStructureKHR(accelerationStructure), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyAccelerationStructureNV wraps vkDestroyAccelerationStructureNV.
func DestroyAccelerationStructureNV(device Device, accelerationStructure AccelerationStructureNV, pAllocator AllocationCallbacks) {
	C.vkDestroyAccelerationStructureNV(C.VkDevice(unsafe.Pointer(device)), C.VkAccelerationStructureNV(accelerationStructure), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyBuffer wraps vkDestroyBuffer.
func DestroyBuffer(device Device, buffer Buffer, pAllocator AllocationCallbacks) {
	C.vkDestroyBuffer(C.VkDevice(unsafe.Pointer(device)), C.VkBuffer(buffer), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyBufferView wraps vkDestroyBufferView.
func DestroyBufferView(device Device, bufferView BufferView, pAllocator AllocationCallbacks) {
	C.vkDestroyBufferView(C.VkDevice(unsafe.Pointer(device)), C.VkBufferView(bufferView), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyCommandPool wraps vkDestroyCommandPool.
func DestroyCommandPool(device Device, commandPool CommandPool, pAllocator AllocationCallbacks) {
	C.vkDestroyCommandPool(C.VkDevice(unsafe.Pointer(device)), C.VkCommandPool(commandPool), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyCuFunctionNVX wraps vkDestroyCuFunctionNVX.
func DestroyCuFunctionNVX(device Device, function CuFunctionNVX, pAllocator AllocationCallbacks) {
	C.vkDestroyCuFunctionNVX(C.VkDevice(unsafe.Pointer(device)), C.VkCuFunctionNVX(function), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyCuModuleNVX wraps vkDestroyCuModuleNVX.
func DestroyCuModuleNVX(device Device, module CuModuleNVX, pAllocator AllocationCallbacks) {
	C.vkDestroyCuModuleNVX(C.VkDevice(unsafe.Pointer(device)), C.VkCuModuleNVX(module), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDataGraphPipelineSessionARM wraps vkDestroyDataGraphPipelineSessionARM.
func DestroyDataGraphPipelineSessionARM(device Device, session DataGraphPipelineSessionARM, pAllocator AllocationCallbacks) {
	C.vkDestroyDataGraphPipelineSessionARM(C.VkDevice(unsafe.Pointer(device)), C.VkDataGraphPipelineSessionARM(session), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDebugReportCallbackEXT wraps vkDestroyDebugReportCallbackEXT.
func DestroyDebugReportCallbackEXT(instance Instance, callback DebugReportCallbackEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyDebugReportCallbackEXT(C.VkInstance(unsafe.Pointer(instance)), C.VkDebugReportCallbackEXT(callback), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDebugUtilsMessengerEXT wraps vkDestroyDebugUtilsMessengerEXT.
func DestroyDebugUtilsMessengerEXT(instance Instance, messenger DebugUtilsMessengerEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyDebugUtilsMessengerEXT(C.VkInstance(unsafe.Pointer(instance)), C.VkDebugUtilsMessengerEXT(messenger), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDeferredOperationKHR wraps vkDestroyDeferredOperationKHR.
func DestroyDeferredOperationKHR(device Device, operation DeferredOperationKHR, pAllocator AllocationCallbacks) {
	C.vkDestroyDeferredOperationKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(operation), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDescriptorPool wraps vkDestroyDescriptorPool.
func DestroyDescriptorPool(device Device, descriptorPool DescriptorPool, pAllocator AllocationCallbacks) {
	C.vkDestroyDescriptorPool(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorPool(descriptorPool), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDescriptorSetLayout wraps vkDestroyDescriptorSetLayout.
func DestroyDescriptorSetLayout(device Device, descriptorSetLayout DescriptorSetLayout, pAllocator AllocationCallbacks) {
	C.vkDestroyDescriptorSetLayout(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSetLayout(descriptorSetLayout), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDescriptorUpdateTemplate wraps vkDestroyDescriptorUpdateTemplate.
func DestroyDescriptorUpdateTemplate(device Device, descriptorUpdateTemplate DescriptorUpdateTemplate, pAllocator AllocationCallbacks) {
	C.vkDestroyDescriptorUpdateTemplate(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyDevice wraps vkDestroyDevice.
func DestroyDevice(device Device, pAllocator AllocationCallbacks) {
	C.vkDestroyDevice(C.VkDevice(unsafe.Pointer(device)), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyEvent wraps vkDestroyEvent.
func DestroyEvent(device Device, event Event, pAllocator AllocationCallbacks) {
	C.vkDestroyEvent(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyExternalComputeQueueNV wraps vkDestroyExternalComputeQueueNV.
func DestroyExternalComputeQueueNV(device Device, externalQueue ExternalComputeQueueNV, pAllocator AllocationCallbacks) {
	C.vkDestroyExternalComputeQueueNV(C.VkDevice(unsafe.Pointer(device)), C.VkExternalComputeQueueNV(unsafe.Pointer(externalQueue)), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyFence wraps vkDestroyFence.
func DestroyFence(device Device, fence Fence, pAllocator AllocationCallbacks) {
	C.vkDestroyFence(C.VkDevice(unsafe.Pointer(device)), C.VkFence(fence), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyFramebuffer wraps vkDestroyFramebuffer.
func DestroyFramebuffer(device Device, framebuffer Framebuffer, pAllocator AllocationCallbacks) {
	C.vkDestroyFramebuffer(C.VkDevice(unsafe.Pointer(device)), C.VkFramebuffer(framebuffer), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyImage wraps vkDestroyImage.
func DestroyImage(device Device, image Image, pAllocator AllocationCallbacks) {
	C.vkDestroyImage(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyImageView wraps vkDestroyImageView.
func DestroyImageView(device Device, imageView ImageView, pAllocator AllocationCallbacks) {
	C.vkDestroyImageView(C.VkDevice(unsafe.Pointer(device)), C.VkImageView(imageView), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyIndirectCommandsLayoutEXT wraps vkDestroyIndirectCommandsLayoutEXT.
func DestroyIndirectCommandsLayoutEXT(device Device, indirectCommandsLayout IndirectCommandsLayoutEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyIndirectCommandsLayoutEXT(C.VkDevice(unsafe.Pointer(device)), C.VkIndirectCommandsLayoutEXT(indirectCommandsLayout), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyIndirectCommandsLayoutNV wraps vkDestroyIndirectCommandsLayoutNV.
func DestroyIndirectCommandsLayoutNV(device Device, indirectCommandsLayout IndirectCommandsLayoutNV, pAllocator AllocationCallbacks) {
	C.vkDestroyIndirectCommandsLayoutNV(C.VkDevice(unsafe.Pointer(device)), C.VkIndirectCommandsLayoutNV(indirectCommandsLayout), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyIndirectExecutionSetEXT wraps vkDestroyIndirectExecutionSetEXT.
func DestroyIndirectExecutionSetEXT(device Device, indirectExecutionSet IndirectExecutionSetEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyIndirectExecutionSetEXT(C.VkDevice(unsafe.Pointer(device)), C.VkIndirectExecutionSetEXT(indirectExecutionSet), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyInstance wraps vkDestroyInstance.
func DestroyInstance(instance Instance, pAllocator AllocationCallbacks) {
	C.vkDestroyInstance(C.VkInstance(unsafe.Pointer(instance)), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyMicromapEXT wraps vkDestroyMicromapEXT.
func DestroyMicromapEXT(device Device, micromap MicromapEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyMicromapEXT(C.VkDevice(unsafe.Pointer(device)), C.VkMicromapEXT(micromap), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyOpticalFlowSessionNV wraps vkDestroyOpticalFlowSessionNV.
func DestroyOpticalFlowSessionNV(device Device, session OpticalFlowSessionNV, pAllocator AllocationCallbacks) {
	C.vkDestroyOpticalFlowSessionNV(C.VkDevice(unsafe.Pointer(device)), C.VkOpticalFlowSessionNV(session), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyPipeline wraps vkDestroyPipeline.
func DestroyPipeline(device Device, pipeline Pipeline, pAllocator AllocationCallbacks) {
	C.vkDestroyPipeline(C.VkDevice(unsafe.Pointer(device)), C.VkPipeline(pipeline), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyPipelineBinaryKHR wraps vkDestroyPipelineBinaryKHR.
func DestroyPipelineBinaryKHR(device Device, pipelineBinary PipelineBinaryKHR, pAllocator AllocationCallbacks) {
	C.vkDestroyPipelineBinaryKHR(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineBinaryKHR(pipelineBinary), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyPipelineCache wraps vkDestroyPipelineCache.
func DestroyPipelineCache(device Device, pipelineCache PipelineCache, pAllocator AllocationCallbacks) {
	C.vkDestroyPipelineCache(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyPipelineLayout wraps vkDestroyPipelineLayout.
func DestroyPipelineLayout(device Device, pipelineLayout PipelineLayout, pAllocator AllocationCallbacks) {
	C.vkDestroyPipelineLayout(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineLayout(pipelineLayout), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyPrivateDataSlot wraps vkDestroyPrivateDataSlot.
func DestroyPrivateDataSlot(device Device, privateDataSlot PrivateDataSlot, pAllocator AllocationCallbacks) {
	C.vkDestroyPrivateDataSlot(C.VkDevice(unsafe.Pointer(device)), C.VkPrivateDataSlot(privateDataSlot), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyQueryPool wraps vkDestroyQueryPool.
func DestroyQueryPool(device Device, queryPool QueryPool, pAllocator AllocationCallbacks) {
	C.vkDestroyQueryPool(C.VkDevice(unsafe.Pointer(device)), C.VkQueryPool(queryPool), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyRenderPass wraps vkDestroyRenderPass.
func DestroyRenderPass(device Device, renderPass RenderPass, pAllocator AllocationCallbacks) {
	C.vkDestroyRenderPass(C.VkDevice(unsafe.Pointer(device)), C.VkRenderPass(renderPass), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroySampler wraps vkDestroySampler.
func DestroySampler(device Device, sampler Sampler, pAllocator AllocationCallbacks) {
	C.vkDestroySampler(C.VkDevice(unsafe.Pointer(device)), C.VkSampler(sampler), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroySamplerYcbcrConversion wraps vkDestroySamplerYcbcrConversion.
func DestroySamplerYcbcrConversion(device Device, ycbcrConversion SamplerYcbcrConversion, pAllocator AllocationCallbacks) {
	C.vkDestroySamplerYcbcrConversion(C.VkDevice(unsafe.Pointer(device)), C.VkSamplerYcbcrConversion(ycbcrConversion), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroySemaphore wraps vkDestroySemaphore.
func DestroySemaphore(device Device, semaphore Semaphore, pAllocator AllocationCallbacks) {
	C.vkDestroySemaphore(C.VkDevice(unsafe.Pointer(device)), C.VkSemaphore(semaphore), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyShaderEXT wraps vkDestroyShaderEXT.
func DestroyShaderEXT(device Device, shader ShaderEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyShaderEXT(C.VkDevice(unsafe.Pointer(device)), C.VkShaderEXT(shader), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyShaderModule wraps vkDestroyShaderModule.
func DestroyShaderModule(device Device, shaderModule ShaderModule, pAllocator AllocationCallbacks) {
	C.vkDestroyShaderModule(C.VkDevice(unsafe.Pointer(device)), C.VkShaderModule(shaderModule), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroySurfaceKHR wraps vkDestroySurfaceKHR.
func DestroySurfaceKHR(instance Instance, surface SurfaceKHR, pAllocator AllocationCallbacks) {
	C.vkDestroySurfaceKHR(C.VkInstance(unsafe.Pointer(instance)), C.VkSurfaceKHR(surface), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroySwapchainKHR wraps vkDestroySwapchainKHR.
func DestroySwapchainKHR(device Device, swapchain SwapchainKHR, pAllocator AllocationCallbacks) {
	C.vkDestroySwapchainKHR(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyTensorARM wraps vkDestroyTensorARM.
func DestroyTensorARM(device Device, tensor TensorARM, pAllocator AllocationCallbacks) {
	C.vkDestroyTensorARM(C.VkDevice(unsafe.Pointer(device)), C.VkTensorARM(tensor), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyTensorViewARM wraps vkDestroyTensorViewARM.
func DestroyTensorViewARM(device Device, tensorView TensorViewARM, pAllocator AllocationCallbacks) {
	C.vkDestroyTensorViewARM(C.VkDevice(unsafe.Pointer(device)), C.VkTensorViewARM(tensorView), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyValidationCacheEXT wraps vkDestroyValidationCacheEXT.
func DestroyValidationCacheEXT(device Device, validationCache ValidationCacheEXT, pAllocator AllocationCallbacks) {
	C.vkDestroyValidationCacheEXT(C.VkDevice(unsafe.Pointer(device)), C.VkValidationCacheEXT(validationCache), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyVideoSessionKHR wraps vkDestroyVideoSessionKHR.
func DestroyVideoSessionKHR(device Device, videoSession VideoSessionKHR, pAllocator AllocationCallbacks) {
	C.vkDestroyVideoSessionKHR(C.VkDevice(unsafe.Pointer(device)), C.VkVideoSessionKHR(videoSession), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DestroyVideoSessionParametersKHR wraps vkDestroyVideoSessionParametersKHR.
func DestroyVideoSessionParametersKHR(device Device, videoSessionParameters VideoSessionParametersKHR, pAllocator AllocationCallbacks) {
	C.vkDestroyVideoSessionParametersKHR(C.VkDevice(unsafe.Pointer(device)), C.VkVideoSessionParametersKHR(videoSessionParameters), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// DeviceWaitIdle wraps vkDeviceWaitIdle.
func DeviceWaitIdle(device Device) Result {
	ret := C.vkDeviceWaitIdle(C.VkDevice(unsafe.Pointer(device)))

	return Result(ret)
}

// DisplayPowerControlEXT wraps vkDisplayPowerControlEXT.
func DisplayPowerControlEXT(device Device, display DisplayKHR, pDisplayPowerInfo DisplayPowerInfoEXT) Result {
	ret := C.vkDisplayPowerControlEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDisplayKHR(display), (*C.VkDisplayPowerInfoEXT)(pDisplayPowerInfo.Raw()))

	return Result(ret)
}

// EndCommandBuffer wraps vkEndCommandBuffer.
func EndCommandBuffer(commandBuffer CommandBuffer) Result {
	ret := C.vkEndCommandBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))

	return Result(ret)
}

// EnumerateDeviceExtensionProperties wraps vkEnumerateDeviceExtensionProperties.
func EnumerateDeviceExtensionProperties(physicalDevice PhysicalDevice, pLayerName ffi.CString, pPropertyCount ffi.Ref[uint32], pProperties ExtensionProperties) Result {
	ret := C.vkEnumerateDeviceExtensionProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.char)(pLayerName.Raw()), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkExtensionProperties)(pProperties.Raw()))

	return Result(ret)
}

// EnumerateDeviceLayerProperties wraps vkEnumerateDeviceLayerProperties.
func EnumerateDeviceLayerProperties(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties LayerProperties) Result {
	ret := C.vkEnumerateDeviceLayerProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkLayerProperties)(pProperties.Raw()))

	return Result(ret)
}

// EnumerateInstanceExtensionProperties wraps vkEnumerateInstanceExtensionProperties.
func EnumerateInstanceExtensionProperties(pLayerName ffi.CString, pPropertyCount ffi.Ref[uint32], pProperties ExtensionProperties) Result {
	ret := C.vkEnumerateInstanceExtensionProperties((*C.char)(pLayerName.Raw()), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkExtensionProperties)(pProperties.Raw()))

	return Result(ret)
}

// EnumerateInstanceLayerProperties wraps vkEnumerateInstanceLayerProperties.
func EnumerateInstanceLayerProperties(pPropertyCount ffi.Ref[uint32], pProperties LayerProperties) Result {
	ret := C.vkEnumerateInstanceLayerProperties((*C.uint32_t)(pPropertyCount.Raw()), (*C.VkLayerProperties)(pProperties.Raw()))

	return Result(ret)
}

// EnumerateInstanceVersion wraps vkEnumerateInstanceVersion.
func EnumerateInstanceVersion(pApiVersion ffi.Ref[uint32]) Result {
	ret := C.vkEnumerateInstanceVersion((*C.uint32_t)(pApiVersion.Raw()))

	return Result(ret)
}

// EnumeratePhysicalDeviceGroups wraps vkEnumeratePhysicalDeviceGroups.
func EnumeratePhysicalDeviceGroups(instance Instance, pPhysicalDeviceGroupCount ffi.Ref[uint32], pPhysicalDeviceGroupProperties PhysicalDeviceGroupProperties) Result {
	ret := C.vkEnumeratePhysicalDeviceGroups(C.VkInstance(unsafe.Pointer(instance)), (*C.uint32_t)(pPhysicalDeviceGroupCount.Raw()), (*C.VkPhysicalDeviceGroupProperties)(pPhysicalDeviceGroupProperties.Raw()))

	return Result(ret)
}

// EnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR wraps vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.
func EnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(physicalDevice PhysicalDevice, queueFamilyIndex uint32, pCounterCount ffi.Ref[uint32], pCounters PerformanceCounterKHR, pCounterDescriptions PerformanceCounterDescriptionKHR) Result {
	ret := C.vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.uint32_t(queueFamilyIndex), (*C.uint32_t)(pCounterCount.Raw()), (*C.VkPerformanceCounterKHR)(pCounters.Raw()), (*C.VkPerformanceCounterDescriptionKHR)(pCounterDescriptions.Raw()))

	return Result(ret)
}

// EnumeratePhysicalDevices wraps vkEnumeratePhysicalDevices.
func EnumeratePhysicalDevices(instance Instance, pPhysicalDeviceCount ffi.Ref[uint32], pPhysicalDevices ffi.Ref[PhysicalDevice]) Result {
	ret := C.vkEnumeratePhysicalDevices(C.VkInstance(unsafe.Pointer(instance)), (*C.uint32_t)(pPhysicalDeviceCount.Raw()), (*C.VkPhysicalDevice)(pPhysicalDevices.Raw()))

	return Result(ret)
}

// ExportMetalObjectsEXT wraps vkExportMetalObjectsEXT.
func ExportMetalObjectsEXT(device Device, pMetalObjectsInfo ExportMetalObjectsInfoEXT) {
	C.vkExportMetalObjectsEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkExportMetalObjectsInfoEXT)(pMetalObjectsInfo.Raw()))
}

// FlushMappedMemoryRanges wraps vkFlushMappedMemoryRanges.
func FlushMappedMemoryRanges(device Device, memoryRangeCount uint32, pMemoryRanges MappedMemoryRange) Result {
	ret := C.vkFlushMappedMemoryRanges(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(memoryRangeCount), (*C.VkMappedMemoryRange)(pMemoryRanges.Raw()))

	return Result(ret)
}

// FreeCommandBuffers wraps vkFreeCommandBuffers.
func FreeCommandBuffers(device Device, commandPool CommandPool, commandBufferCount uint32, pCommandBuffers ffi.Ref[CommandBuffer]) {
	C.vkFreeCommandBuffers(C.VkDevice(unsafe.Pointer(device)), C.VkCommandPool(commandPool), C.uint32_t(commandBufferCount), (*C.VkCommandBuffer)(pCommandBuffers.Raw()))
}

// FreeDescriptorSets wraps vkFreeDescriptorSets.
func FreeDescriptorSets(device Device, descriptorPool DescriptorPool, descriptorSetCount uint32, pDescriptorSets ffi.Ref[DescriptorSet]) Result {
	ret := C.vkFreeDescriptorSets(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorPool(descriptorPool), C.uint32_t(descriptorSetCount), (*C.VkDescriptorSet)(pDescriptorSets.Raw()))

	return Result(ret)
}

// FreeMemory wraps vkFreeMemory.
func FreeMemory(device Device, memory DeviceMemory, pAllocator AllocationCallbacks) {
	C.vkFreeMemory(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory), (*C.VkAllocationCallbacks)(pAllocator.Raw()))
}

// GetAccelerationStructureBuildSizesKHR wraps vkGetAccelerationStructureBuildSizesKHR.
func GetAccelerationStructureBuildSizesKHR(device Device, buildType AccelerationStructureBuildTypeKHR, pBuildInfo AccelerationStructureBuildGeometryInfoKHR, pMaxPrimitiveCounts ffi.Ref[uint32], pSizeInfo AccelerationStructureBuildSizesInfoKHR) {
	C.vkGetAccelerationStructureBuildSizesKHR(C.VkDevice(unsafe.Pointer(device)), C.VkAccelerationStructureBuildTypeKHR(buildType), (*C.VkAccelerationStructureBuildGeometryInfoKHR)(pBuildInfo.Raw()), (*C.uint32_t)(pMaxPrimitiveCounts.Raw()), (*C.VkAccelerationStructureBuildSizesInfoKHR)(pSizeInfo.Raw()))
}

// vkGetAccelerationStructureDeviceAddressKHR is unsupported: unknown category direct.

// GetAccelerationStructureHandleNV wraps vkGetAccelerationStructureHandleNV.
func GetAccelerationStructureHandleNV(device Device, accelerationStructure AccelerationStructureNV, dataSize uintptr, pData unsafe.Pointer) Result {
	ret := C.vkGetAccelerationStructureHandleNV(C.VkDevice(unsafe.Pointer(device)), C.VkAccelerationStructureNV(accelerationStructure), C.size_t(dataSize), pData)

	return Result(ret)
}

// GetAccelerationStructureMemoryRequirementsNV wraps vkGetAccelerationStructureMemoryRequirementsNV.
func GetAccelerationStructureMemoryRequirementsNV(device Device, pInfo AccelerationStructureMemoryRequirementsInfoNV, pMemoryRequirements MemoryRequirements2) {
	C.vkGetAccelerationStructureMemoryRequirementsNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkAccelerationStructureMemoryRequirementsInfoNV)(pInfo.Raw()), (*C.VkMemoryRequirements2KHR)(pMemoryRequirements.Raw()))
}

// GetAccelerationStructureOpaqueCaptureDescriptorDataEXT wraps vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT.
func GetAccelerationStructureOpaqueCaptureDescriptorDataEXT(device Device, pInfo AccelerationStructureCaptureDescriptorDataInfoEXT, pData unsafe.Pointer) Result {
	ret := C.vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkAccelerationStructureCaptureDescriptorDataInfoEXT)(pInfo.Raw()), pData)

	return Result(ret)
}

// vkGetBufferDeviceAddress is unsupported: unknown category direct.

// GetBufferMemoryRequirements wraps vkGetBufferMemoryRequirements.
func GetBufferMemoryRequirements(device Device, buffer Buffer, pMemoryRequirements MemoryRequirements) {
	C.vkGetBufferMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), C.VkBuffer(buffer), (*C.VkMemoryRequirements)(pMemoryRequirements.Raw()))
}

// GetBufferMemoryRequirements2 wraps vkGetBufferMemoryRequirements2.
func GetBufferMemoryRequirements2(device Device, pInfo BufferMemoryRequirementsInfo2, pMemoryRequirements MemoryRequirements2) {
	C.vkGetBufferMemoryRequirements2(C.VkDevice(unsafe.Pointer(device)), (*C.VkBufferMemoryRequirementsInfo2)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetBufferOpaqueCaptureAddress wraps vkGetBufferOpaqueCaptureAddress.
func GetBufferOpaqueCaptureAddress(device Device, pInfo BufferDeviceAddressInfo) uint64 {
	ret := C.vkGetBufferOpaqueCaptureAddress(C.VkDevice(unsafe.Pointer(device)), (*C.VkBufferDeviceAddressInfo)(pInfo.Raw()))

	return uint64(ret)
}

// GetBufferOpaqueCaptureDescriptorDataEXT wraps vkGetBufferOpaqueCaptureDescriptorDataEXT.
func GetBufferOpaqueCaptureDescriptorDataEXT(device Device, pInfo BufferCaptureDescriptorDataInfoEXT, pData unsafe.Pointer) Result {
	ret := C.vkGetBufferOpaqueCaptureDescriptorDataEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkBufferCaptureDescriptorDataInfoEXT)(pInfo.Raw()), pData)

	return Result(ret)
}

// GetCalibratedTimestampsKHR wraps vkGetCalibratedTimestampsKHR.
func GetCalibratedTimestampsKHR(device Device, timestampCount uint32, pTimestampInfos CalibratedTimestampInfoKHR, pTimestamps ffi.Ref[uint64], pMaxDeviation ffi.Ref[uint64]) Result {
	ret := C.vkGetCalibratedTimestampsKHR(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(timestampCount), (*C.VkCalibratedTimestampInfoKHR)(pTimestampInfos.Raw()), (*C.uint64_t)(pTimestamps.Raw()), (*C.uint64_t)(pMaxDeviation.Raw()))

	return Result(ret)
}

// GetClusterAccelerationStructureBuildSizesNV wraps vkGetClusterAccelerationStructureBuildSizesNV.
func GetClusterAccelerationStructureBuildSizesNV(device Device, pInfo ClusterAccelerationStructureInputInfoNV, pSizeInfo AccelerationStructureBuildSizesInfoKHR) {
	C.vkGetClusterAccelerationStructureBuildSizesNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkClusterAccelerationStructureInputInfoNV)(pInfo.Raw()), (*C.VkAccelerationStructureBuildSizesInfoKHR)(pSizeInfo.Raw()))
}

// GetDataGraphPipelineAvailablePropertiesARM wraps vkGetDataGraphPipelineAvailablePropertiesARM.
func GetDataGraphPipelineAvailablePropertiesARM(device Device, pPipelineInfo DataGraphPipelineInfoARM, pPropertiesCount ffi.Ref[uint32], pProperties ffi.Ref[DataGraphPipelinePropertyARM]) Result {
	ret := C.vkGetDataGraphPipelineAvailablePropertiesARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkDataGraphPipelineInfoARM)(pPipelineInfo.Raw()), (*C.uint32_t)(pPropertiesCount.Raw()), (*C.VkDataGraphPipelinePropertyARM)(pProperties.Raw()))

	return Result(ret)
}

// GetDataGraphPipelinePropertiesARM wraps vkGetDataGraphPipelinePropertiesARM.
func GetDataGraphPipelinePropertiesARM(device Device, pPipelineInfo DataGraphPipelineInfoARM, propertiesCount uint32, pProperties DataGraphPipelinePropertyQueryResultARM) Result {
	ret := C.vkGetDataGraphPipelinePropertiesARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkDataGraphPipelineInfoARM)(pPipelineInfo.Raw()), C.uint32_t(propertiesCount), (*C.VkDataGraphPipelinePropertyQueryResultARM)(pProperties.Raw()))

	return Result(ret)
}

// GetDataGraphPipelineSessionBindPointRequirementsARM wraps vkGetDataGraphPipelineSessionBindPointRequirementsARM.
func GetDataGraphPipelineSessionBindPointRequirementsARM(device Device, pInfo DataGraphPipelineSessionBindPointRequirementsInfoARM, pBindPointRequirementCount ffi.Ref[uint32], pBindPointRequirements DataGraphPipelineSessionBindPointRequirementARM) Result {
	ret := C.vkGetDataGraphPipelineSessionBindPointRequirementsARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkDataGraphPipelineSessionBindPointRequirementsInfoARM)(pInfo.Raw()), (*C.uint32_t)(pBindPointRequirementCount.Raw()), (*C.VkDataGraphPipelineSessionBindPointRequirementARM)(pBindPointRequirements.Raw()))

	return Result(ret)
}

// GetDataGraphPipelineSessionMemoryRequirementsARM wraps vkGetDataGraphPipelineSessionMemoryRequirementsARM.
func GetDataGraphPipelineSessionMemoryRequirementsARM(device Device, pInfo DataGraphPipelineSessionMemoryRequirementsInfoARM, pMemoryRequirements MemoryRequirements2) {
	C.vkGetDataGraphPipelineSessionMemoryRequirementsARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkDataGraphPipelineSessionMemoryRequirementsInfoARM)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetDeferredOperationMaxConcurrencyKHR wraps vkGetDeferredOperationMaxConcurrencyKHR.
func GetDeferredOperationMaxConcurrencyKHR(device Device, operation DeferredOperationKHR) uint32 {
	ret := C.vkGetDeferredOperationMaxConcurrencyKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(operation))

	return uint32(ret)
}

// GetDeferredOperationResultKHR wraps vkGetDeferredOperationResultKHR.
func GetDeferredOperationResultKHR(device Device, operation DeferredOperationKHR) Result {
	ret := C.vkGetDeferredOperationResultKHR(C.VkDevice(unsafe.Pointer(device)), C.VkDeferredOperationKHR(operation))

	return Result(ret)
}

// GetDescriptorEXT wraps vkGetDescriptorEXT.
func GetDescriptorEXT(device Device, pDescriptorInfo DescriptorGetInfoEXT, dataSize uintptr, pDescriptor unsafe.Pointer) {
	C.vkGetDescriptorEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorGetInfoEXT)(pDescriptorInfo.Raw()), C.size_t(dataSize), pDescriptor)
}

// GetDescriptorSetHostMappingVALVE wraps vkGetDescriptorSetHostMappingVALVE.
func GetDescriptorSetHostMappingVALVE(device Device, descriptorSet DescriptorSet, ppData ffi.Ref[unsafe.Pointer]) {
	C.vkGetDescriptorSetHostMappingVALVE(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSet(descriptorSet), (*unsafe.Pointer)(ppData.Raw()))
}

// GetDescriptorSetLayoutBindingOffsetEXT wraps vkGetDescriptorSetLayoutBindingOffsetEXT.
func GetDescriptorSetLayoutBindingOffsetEXT(device Device, layout DescriptorSetLayout, binding uint32, pOffset ffi.Ref[DeviceSize]) {
	C.vkGetDescriptorSetLayoutBindingOffsetEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSetLayout(layout), C.uint32_t(binding), (*C.VkDeviceSize)(pOffset.Raw()))
}

// GetDescriptorSetLayoutHostMappingInfoVALVE wraps vkGetDescriptorSetLayoutHostMappingInfoVALVE.
func GetDescriptorSetLayoutHostMappingInfoVALVE(device Device, pBindingReference DescriptorSetBindingReferenceVALVE, pHostMapping DescriptorSetLayoutHostMappingInfoVALVE) {
	C.vkGetDescriptorSetLayoutHostMappingInfoVALVE(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorSetBindingReferenceVALVE)(pBindingReference.Raw()), (*C.VkDescriptorSetLayoutHostMappingInfoVALVE)(pHostMapping.Raw()))
}

// GetDescriptorSetLayoutSizeEXT wraps vkGetDescriptorSetLayoutSizeEXT.
func GetDescriptorSetLayoutSizeEXT(device Device, layout DescriptorSetLayout, pLayoutSizeInBytes ffi.Ref[DeviceSize]) {
	C.vkGetDescriptorSetLayoutSizeEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSetLayout(layout), (*C.VkDeviceSize)(pLayoutSizeInBytes.Raw()))
}

// GetDescriptorSetLayoutSupport wraps vkGetDescriptorSetLayoutSupport.
func GetDescriptorSetLayoutSupport(device Device, pCreateInfo DescriptorSetLayoutCreateInfo, pSupport DescriptorSetLayoutSupport) {
	C.vkGetDescriptorSetLayoutSupport(C.VkDevice(unsafe.Pointer(device)), (*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo.Raw()), (*C.VkDescriptorSetLayoutSupport)(pSupport.Raw()))
}

// GetDeviceAccelerationStructureCompatibilityKHR wraps vkGetDeviceAccelerationStructureCompatibilityKHR.
func GetDeviceAccelerationStructureCompatibilityKHR(device Device, pVersionInfo AccelerationStructureVersionInfoKHR, pCompatibility ffi.Ref[AccelerationStructureCompatibilityKHR]) {
	C.vkGetDeviceAccelerationStructureCompatibilityKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkAccelerationStructureVersionInfoKHR)(pVersionInfo.Raw()), (*C.VkAccelerationStructureCompatibilityKHR)(pCompatibility.Raw()))
}

// GetDeviceBufferMemoryRequirements wraps vkGetDeviceBufferMemoryRequirements.
func GetDeviceBufferMemoryRequirements(device Device, pInfo DeviceBufferMemoryRequirements, pMemoryRequirements MemoryRequirements2) {
	C.vkGetDeviceBufferMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceBufferMemoryRequirements)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetDeviceFaultInfoEXT wraps vkGetDeviceFaultInfoEXT.
func GetDeviceFaultInfoEXT(device Device, pFaultCounts DeviceFaultCountsEXT, pFaultInfo DeviceFaultInfoEXT) Result {
	ret := C.vkGetDeviceFaultInfoEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceFaultCountsEXT)(pFaultCounts.Raw()), (*C.VkDeviceFaultInfoEXT)(pFaultInfo.Raw()))

	return Result(ret)
}

// GetDeviceGroupPeerMemoryFeatures wraps vkGetDeviceGroupPeerMemoryFeatures.
func GetDeviceGroupPeerMemoryFeatures(device Device, heapIndex uint32, localDeviceIndex uint32, remoteDeviceIndex uint32, pPeerMemoryFeatures ffi.Ref[PeerMemoryFeatureFlags]) {
	C.vkGetDeviceGroupPeerMemoryFeatures(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(heapIndex), C.uint32_t(localDeviceIndex), C.uint32_t(remoteDeviceIndex), (*C.VkPeerMemoryFeatureFlags)(pPeerMemoryFeatures.Raw()))
}

// GetDeviceGroupPresentCapabilitiesKHR wraps vkGetDeviceGroupPresentCapabilitiesKHR.
func GetDeviceGroupPresentCapabilitiesKHR(device Device, pDeviceGroupPresentCapabilities DeviceGroupPresentCapabilitiesKHR) Result {
	ret := C.vkGetDeviceGroupPresentCapabilitiesKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceGroupPresentCapabilitiesKHR)(pDeviceGroupPresentCapabilities.Raw()))

	return Result(ret)
}

// GetDeviceGroupSurfacePresentModes2EXT wraps vkGetDeviceGroupSurfacePresentModes2EXT.
func GetDeviceGroupSurfacePresentModes2EXT(device Device, pSurfaceInfo PhysicalDeviceSurfaceInfo2KHR, pModes ffi.Ref[DeviceGroupPresentModeFlagsKHR]) Result {
	ret := C.vkGetDeviceGroupSurfacePresentModes2EXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkPhysicalDeviceSurfaceInfo2KHR)(pSurfaceInfo.Raw()), (*C.VkDeviceGroupPresentModeFlagsKHR)(pModes.Raw()))

	return Result(ret)
}

// GetDeviceGroupSurfacePresentModesKHR wraps vkGetDeviceGroupSurfacePresentModesKHR.
func GetDeviceGroupSurfacePresentModesKHR(device Device, surface SurfaceKHR, pModes ffi.Ref[DeviceGroupPresentModeFlagsKHR]) Result {
	ret := C.vkGetDeviceGroupSurfacePresentModesKHR(C.VkDevice(unsafe.Pointer(device)), C.VkSurfaceKHR(surface), (*C.VkDeviceGroupPresentModeFlagsKHR)(pModes.Raw()))

	return Result(ret)
}

// GetDeviceImageMemoryRequirements wraps vkGetDeviceImageMemoryRequirements.
func GetDeviceImageMemoryRequirements(device Device, pInfo DeviceImageMemoryRequirements, pMemoryRequirements MemoryRequirements2) {
	C.vkGetDeviceImageMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceImageMemoryRequirements)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetDeviceImageSparseMemoryRequirements wraps vkGetDeviceImageSparseMemoryRequirements.
func GetDeviceImageSparseMemoryRequirements(device Device, pInfo DeviceImageMemoryRequirements, pSparseMemoryRequirementCount ffi.Ref[uint32], pSparseMemoryRequirements SparseImageMemoryRequirements2) {
	C.vkGetDeviceImageSparseMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceImageMemoryRequirements)(pInfo.Raw()), (*C.uint32_t)(pSparseMemoryRequirementCount.Raw()), (*C.VkSparseImageMemoryRequirements2)(pSparseMemoryRequirements.Raw()))
}

// GetDeviceImageSubresourceLayout wraps vkGetDeviceImageSubresourceLayout.
func GetDeviceImageSubresourceLayout(device Device, pInfo DeviceImageSubresourceInfo, pLayout SubresourceLayout2) {
	C.vkGetDeviceImageSubresourceLayout(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceImageSubresourceInfo)(pInfo.Raw()), (*C.VkSubresourceLayout2)(pLayout.Raw()))
}

// GetDeviceMemoryCommitment wraps vkGetDeviceMemoryCommitment.
func GetDeviceMemoryCommitment(device Device, memory DeviceMemory, pCommittedMemoryInBytes ffi.Ref[DeviceSize]) {
	C.vkGetDeviceMemoryCommitment(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory), (*C.VkDeviceSize)(pCommittedMemoryInBytes.Raw()))
}

// GetDeviceMemoryOpaqueCaptureAddress wraps vkGetDeviceMemoryOpaqueCaptureAddress.
func GetDeviceMemoryOpaqueCaptureAddress(device Device, pInfo DeviceMemoryOpaqueCaptureAddressInfo) uint64 {
	ret := C.vkGetDeviceMemoryOpaqueCaptureAddress(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceMemoryOpaqueCaptureAddressInfo)(pInfo.Raw()))

	return uint64(ret)
}

// GetDeviceMicromapCompatibilityEXT wraps vkGetDeviceMicromapCompatibilityEXT.
func GetDeviceMicromapCompatibilityEXT(device Device, pVersionInfo MicromapVersionInfoEXT, pCompatibility ffi.Ref[AccelerationStructureCompatibilityKHR]) {
	C.vkGetDeviceMicromapCompatibilityEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkMicromapVersionInfoEXT)(pVersionInfo.Raw()), (*C.VkAccelerationStructureCompatibilityKHR)(pCompatibility.Raw()))
}

// GetDeviceProcAddr wraps vkGetDeviceProcAddr.
func GetDeviceProcAddr(device Device, pName ffi.CString) unsafe.Pointer {
	ret := C.vkGetDeviceProcAddr(C.VkDevice(unsafe.Pointer(device)), (*C.char)(pName.Raw()))

	return unsafe.Pointer(ret)
}

// GetDeviceQueue wraps vkGetDeviceQueue.
func GetDeviceQueue(device Device, queueFamilyIndex uint32, queueIndex uint32, pQueue ffi.Ref[Queue]) {
	C.vkGetDeviceQueue(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), (*C.VkQueue)(pQueue.Raw()))
}

// GetDeviceQueue2 wraps vkGetDeviceQueue2.
func GetDeviceQueue2(device Device, pQueueInfo DeviceQueueInfo2, pQueue ffi.Ref[Queue]) {
	C.vkGetDeviceQueue2(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceQueueInfo2)(pQueueInfo.Raw()), (*C.VkQueue)(pQueue.Raw()))
}

// GetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI wraps vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI.
func GetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(device Device, renderpass RenderPass, pMaxWorkgroupSize Extent2D) Result {
	ret := C.vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(C.VkDevice(unsafe.Pointer(device)), C.VkRenderPass(renderpass), (*C.VkExtent2D)(pMaxWorkgroupSize.Raw()))

	return Result(ret)
}

// GetDeviceTensorMemoryRequirementsARM wraps vkGetDeviceTensorMemoryRequirementsARM.
func GetDeviceTensorMemoryRequirementsARM(device Device, pInfo DeviceTensorMemoryRequirementsARM, pMemoryRequirements MemoryRequirements2) {
	C.vkGetDeviceTensorMemoryRequirementsARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceTensorMemoryRequirementsARM)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetDisplayModeProperties2KHR wraps vkGetDisplayModeProperties2KHR.
func GetDisplayModeProperties2KHR(physicalDevice PhysicalDevice, display DisplayKHR, pPropertyCount ffi.Ref[uint32], pProperties DisplayModeProperties2KHR) Result {
	ret := C.vkGetDisplayModeProperties2KHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkDisplayKHR(display), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkDisplayModeProperties2KHR)(pProperties.Raw()))

	return Result(ret)
}

// GetDisplayModePropertiesKHR wraps vkGetDisplayModePropertiesKHR.
func GetDisplayModePropertiesKHR(physicalDevice PhysicalDevice, display DisplayKHR, pPropertyCount ffi.Ref[uint32], pProperties DisplayModePropertiesKHR) Result {
	ret := C.vkGetDisplayModePropertiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkDisplayKHR(display), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkDisplayModePropertiesKHR)(pProperties.Raw()))

	return Result(ret)
}

// GetDisplayPlaneCapabilities2KHR wraps vkGetDisplayPlaneCapabilities2KHR.
func GetDisplayPlaneCapabilities2KHR(physicalDevice PhysicalDevice, pDisplayPlaneInfo DisplayPlaneInfo2KHR, pCapabilities DisplayPlaneCapabilities2KHR) Result {
	ret := C.vkGetDisplayPlaneCapabilities2KHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkDisplayPlaneInfo2KHR)(pDisplayPlaneInfo.Raw()), (*C.VkDisplayPlaneCapabilities2KHR)(pCapabilities.Raw()))

	return Result(ret)
}

// GetDisplayPlaneCapabilitiesKHR wraps vkGetDisplayPlaneCapabilitiesKHR.
func GetDisplayPlaneCapabilitiesKHR(physicalDevice PhysicalDevice, mode DisplayModeKHR, planeIndex uint32, pCapabilities DisplayPlaneCapabilitiesKHR) Result {
	ret := C.vkGetDisplayPlaneCapabilitiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkDisplayModeKHR(mode), C.uint32_t(planeIndex), (*C.VkDisplayPlaneCapabilitiesKHR)(pCapabilities.Raw()))

	return Result(ret)
}

// GetDisplayPlaneSupportedDisplaysKHR wraps vkGetDisplayPlaneSupportedDisplaysKHR.
func GetDisplayPlaneSupportedDisplaysKHR(physicalDevice PhysicalDevice, planeIndex uint32, pDisplayCount ffi.Ref[uint32], pDisplays ffi.Ref[DisplayKHR]) Result {
	ret := C.vkGetDisplayPlaneSupportedDisplaysKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.uint32_t(planeIndex), (*C.uint32_t)(pDisplayCount.Raw()), (*C.VkDisplayKHR)(pDisplays.Raw()))

	return Result(ret)
}

// GetDrmDisplayEXT wraps vkGetDrmDisplayEXT.
func GetDrmDisplayEXT(physicalDevice PhysicalDevice, drmFd int32, connectorId uint32, display ffi.Ref[DisplayKHR]) Result {
	ret := C.vkGetDrmDisplayEXT(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.int32_t(drmFd), C.uint32_t(connectorId), (*C.VkDisplayKHR)(display.Raw()))

	return Result(ret)
}

// GetDynamicRenderingTilePropertiesQCOM wraps vkGetDynamicRenderingTilePropertiesQCOM.
func GetDynamicRenderingTilePropertiesQCOM(device Device, pRenderingInfo RenderingInfo, pProperties TilePropertiesQCOM) Result {
	ret := C.vkGetDynamicRenderingTilePropertiesQCOM(C.VkDevice(unsafe.Pointer(device)), (*C.VkRenderingInfo)(pRenderingInfo.Raw()), (*C.VkTilePropertiesQCOM)(pProperties.Raw()))

	return Result(ret)
}

// GetEncodedVideoSessionParametersKHR wraps vkGetEncodedVideoSessionParametersKHR.
func GetEncodedVideoSessionParametersKHR(device Device, pVideoSessionParametersInfo VideoEncodeSessionParametersGetInfoKHR, pFeedbackInfo VideoEncodeSessionParametersFeedbackInfoKHR, pDataSize ffi.Ref[uintptr], pData unsafe.Pointer) Result {
	ret := C.vkGetEncodedVideoSessionParametersKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkVideoEncodeSessionParametersGetInfoKHR)(pVideoSessionParametersInfo.Raw()), (*C.VkVideoEncodeSessionParametersFeedbackInfoKHR)(pFeedbackInfo.Raw()), (*C.size_t)(pDataSize.Raw()), pData)

	return Result(ret)
}

// GetEventStatus wraps vkGetEventStatus.
func GetEventStatus(device Device, event Event) Result {
	ret := C.vkGetEventStatus(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event))

	return Result(ret)
}

// GetExternalComputeQueueDataNV wraps vkGetExternalComputeQueueDataNV.
func GetExternalComputeQueueDataNV(externalQueue ExternalComputeQueueNV, params ExternalComputeQueueDataParamsNV, pData unsafe.Pointer) {
	C.vkGetExternalComputeQueueDataNV(C.VkExternalComputeQueueNV(unsafe.Pointer(externalQueue)), (*C.VkExternalComputeQueueDataParamsNV)(params.Raw()), pData)
}

// vkGetFenceFdKHR.pFd is unsupported: category pointer -> ?? int.

// GetFenceStatus wraps vkGetFenceStatus.
func GetFenceStatus(device Device, fence Fence) Result {
	ret := C.vkGetFenceStatus(C.VkDevice(unsafe.Pointer(device)), C.VkFence(fence))

	return Result(ret)
}

// vkGetFenceWin32HandleKHR.pHandle is unsupported: category pointer -> ?? HANDLE.

// GetFramebufferTilePropertiesQCOM wraps vkGetFramebufferTilePropertiesQCOM.
func GetFramebufferTilePropertiesQCOM(device Device, framebuffer Framebuffer, pPropertiesCount ffi.Ref[uint32], pProperties TilePropertiesQCOM) Result {
	ret := C.vkGetFramebufferTilePropertiesQCOM(C.VkDevice(unsafe.Pointer(device)), C.VkFramebuffer(framebuffer), (*C.uint32_t)(pPropertiesCount.Raw()), (*C.VkTilePropertiesQCOM)(pProperties.Raw()))

	return Result(ret)
}

// GetGeneratedCommandsMemoryRequirementsEXT wraps vkGetGeneratedCommandsMemoryRequirementsEXT.
func GetGeneratedCommandsMemoryRequirementsEXT(device Device, pInfo GeneratedCommandsMemoryRequirementsInfoEXT, pMemoryRequirements MemoryRequirements2) {
	C.vkGetGeneratedCommandsMemoryRequirementsEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkGeneratedCommandsMemoryRequirementsInfoEXT)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetGeneratedCommandsMemoryRequirementsNV wraps vkGetGeneratedCommandsMemoryRequirementsNV.
func GetGeneratedCommandsMemoryRequirementsNV(device Device, pInfo GeneratedCommandsMemoryRequirementsInfoNV, pMemoryRequirements MemoryRequirements2) {
	C.vkGetGeneratedCommandsMemoryRequirementsNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkGeneratedCommandsMemoryRequirementsInfoNV)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetImageDrmFormatModifierPropertiesEXT wraps vkGetImageDrmFormatModifierPropertiesEXT.
func GetImageDrmFormatModifierPropertiesEXT(device Device, image Image, pProperties ImageDrmFormatModifierPropertiesEXT) Result {
	ret := C.vkGetImageDrmFormatModifierPropertiesEXT(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.VkImageDrmFormatModifierPropertiesEXT)(pProperties.Raw()))

	return Result(ret)
}

// GetImageMemoryRequirements wraps vkGetImageMemoryRequirements.
func GetImageMemoryRequirements(device Device, image Image, pMemoryRequirements MemoryRequirements) {
	C.vkGetImageMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.VkMemoryRequirements)(pMemoryRequirements.Raw()))
}

// GetImageMemoryRequirements2 wraps vkGetImageMemoryRequirements2.
func GetImageMemoryRequirements2(device Device, pInfo ImageMemoryRequirementsInfo2, pMemoryRequirements MemoryRequirements2) {
	C.vkGetImageMemoryRequirements2(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageMemoryRequirementsInfo2)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetImageOpaqueCaptureDescriptorDataEXT wraps vkGetImageOpaqueCaptureDescriptorDataEXT.
func GetImageOpaqueCaptureDescriptorDataEXT(device Device, pInfo ImageCaptureDescriptorDataInfoEXT, pData unsafe.Pointer) Result {
	ret := C.vkGetImageOpaqueCaptureDescriptorDataEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageCaptureDescriptorDataInfoEXT)(pInfo.Raw()), pData)

	return Result(ret)
}

// GetImageSparseMemoryRequirements wraps vkGetImageSparseMemoryRequirements.
func GetImageSparseMemoryRequirements(device Device, image Image, pSparseMemoryRequirementCount ffi.Ref[uint32], pSparseMemoryRequirements SparseImageMemoryRequirements) {
	C.vkGetImageSparseMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.uint32_t)(pSparseMemoryRequirementCount.Raw()), (*C.VkSparseImageMemoryRequirements)(pSparseMemoryRequirements.Raw()))
}

// GetImageSparseMemoryRequirements2 wraps vkGetImageSparseMemoryRequirements2.
func GetImageSparseMemoryRequirements2(device Device, pInfo ImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount ffi.Ref[uint32], pSparseMemoryRequirements SparseImageMemoryRequirements2) {
	C.vkGetImageSparseMemoryRequirements2(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageSparseMemoryRequirementsInfo2)(pInfo.Raw()), (*C.uint32_t)(pSparseMemoryRequirementCount.Raw()), (*C.VkSparseImageMemoryRequirements2)(pSparseMemoryRequirements.Raw()))
}

// GetImageSubresourceLayout wraps vkGetImageSubresourceLayout.
func GetImageSubresourceLayout(device Device, image Image, pSubresource ImageSubresource, pLayout SubresourceLayout) {
	C.vkGetImageSubresourceLayout(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.VkImageSubresource)(pSubresource.Raw()), (*C.VkSubresourceLayout)(pLayout.Raw()))
}

// GetImageSubresourceLayout2 wraps vkGetImageSubresourceLayout2.
func GetImageSubresourceLayout2(device Device, image Image, pSubresource ImageSubresource2, pLayout SubresourceLayout2) {
	C.vkGetImageSubresourceLayout2(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.VkImageSubresource2)(pSubresource.Raw()), (*C.VkSubresourceLayout2)(pLayout.Raw()))
}

// GetImageViewAddressNVX wraps vkGetImageViewAddressNVX.
func GetImageViewAddressNVX(device Device, imageView ImageView, pProperties ImageViewAddressPropertiesNVX) Result {
	ret := C.vkGetImageViewAddressNVX(C.VkDevice(unsafe.Pointer(device)), C.VkImageView(imageView), (*C.VkImageViewAddressPropertiesNVX)(pProperties.Raw()))

	return Result(ret)
}

// GetImageViewHandle64NVX wraps vkGetImageViewHandle64NVX.
func GetImageViewHandle64NVX(device Device, pInfo ImageViewHandleInfoNVX) uint64 {
	ret := C.vkGetImageViewHandle64NVX(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageViewHandleInfoNVX)(pInfo.Raw()))

	return uint64(ret)
}

// GetImageViewHandleNVX wraps vkGetImageViewHandleNVX.
func GetImageViewHandleNVX(device Device, pInfo ImageViewHandleInfoNVX) uint32 {
	ret := C.vkGetImageViewHandleNVX(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageViewHandleInfoNVX)(pInfo.Raw()))

	return uint32(ret)
}

// GetImageViewOpaqueCaptureDescriptorDataEXT wraps vkGetImageViewOpaqueCaptureDescriptorDataEXT.
func GetImageViewOpaqueCaptureDescriptorDataEXT(device Device, pInfo ImageViewCaptureDescriptorDataInfoEXT, pData unsafe.Pointer) Result {
	ret := C.vkGetImageViewOpaqueCaptureDescriptorDataEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkImageViewCaptureDescriptorDataInfoEXT)(pInfo.Raw()), pData)

	return Result(ret)
}

// GetInstanceProcAddr wraps vkGetInstanceProcAddr.
func GetInstanceProcAddr(instance Instance, pName ffi.CString) unsafe.Pointer {
	ret := C.vkGetInstanceProcAddr(C.VkInstance(unsafe.Pointer(instance)), (*C.char)(pName.Raw()))

	return unsafe.Pointer(ret)
}

// GetLatencyTimingsNV wraps vkGetLatencyTimingsNV.
func GetLatencyTimingsNV(device Device, swapchain SwapchainKHR, pLatencyMarkerInfo GetLatencyMarkerInfoNV) {
	C.vkGetLatencyTimingsNV(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkGetLatencyMarkerInfoNV)(pLatencyMarkerInfo.Raw()))
}

// vkGetMemoryFdKHR.pFd is unsupported: category pointer -> ?? int.

// vkGetMemoryFdPropertiesKHR.fd is unsupported: unknown member direct type int.

// GetMemoryHostPointerPropertiesEXT wraps vkGetMemoryHostPointerPropertiesEXT.
func GetMemoryHostPointerPropertiesEXT(device Device, handleType ExternalMemoryHandleTypeFlags, pHostPointer unsafe.Pointer, pMemoryHostPointerProperties MemoryHostPointerPropertiesEXT) Result {
	ret := C.vkGetMemoryHostPointerPropertiesEXT(C.VkDevice(unsafe.Pointer(device)), C.VkExternalMemoryHandleTypeFlagBits(handleType), pHostPointer, (*C.VkMemoryHostPointerPropertiesEXT)(pMemoryHostPointerProperties.Raw()))

	return Result(ret)
}

// GetMemoryMetalHandleEXT wraps vkGetMemoryMetalHandleEXT.
func GetMemoryMetalHandleEXT(device Device, pGetMetalHandleInfo MemoryGetMetalHandleInfoEXT, pHandle ffi.Ref[unsafe.Pointer]) Result {
	ret := C.vkGetMemoryMetalHandleEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkMemoryGetMetalHandleInfoEXT)(pGetMetalHandleInfo.Raw()), (*unsafe.Pointer)(pHandle.Raw()))

	return Result(ret)
}

// GetMemoryMetalHandlePropertiesEXT wraps vkGetMemoryMetalHandlePropertiesEXT.
func GetMemoryMetalHandlePropertiesEXT(device Device, handleType ExternalMemoryHandleTypeFlags, pHandle unsafe.Pointer, pMemoryMetalHandleProperties MemoryMetalHandlePropertiesEXT) Result {
	ret := C.vkGetMemoryMetalHandlePropertiesEXT(C.VkDevice(unsafe.Pointer(device)), C.VkExternalMemoryHandleTypeFlagBits(handleType), pHandle, (*C.VkMemoryMetalHandlePropertiesEXT)(pMemoryMetalHandleProperties.Raw()))

	return Result(ret)
}

// vkGetMemoryRemoteAddressNV.pAddress is unsupported: category pointer -> ?? VkRemoteAddressNV.

// vkGetMemoryWin32HandleKHR.pHandle is unsupported: category pointer -> ?? HANDLE.

// vkGetMemoryWin32HandleNV.pHandle is unsupported: category pointer -> ?? HANDLE.

// vkGetMemoryWin32HandlePropertiesKHR.handle is unsupported: unknown member direct type HANDLE.

// GetMicromapBuildSizesEXT wraps vkGetMicromapBuildSizesEXT.
func GetMicromapBuildSizesEXT(device Device, buildType AccelerationStructureBuildTypeKHR, pBuildInfo MicromapBuildInfoEXT, pSizeInfo MicromapBuildSizesInfoEXT) {
	C.vkGetMicromapBuildSizesEXT(C.VkDevice(unsafe.Pointer(device)), C.VkAccelerationStructureBuildTypeKHR(buildType), (*C.VkMicromapBuildInfoEXT)(pBuildInfo.Raw()), (*C.VkMicromapBuildSizesInfoEXT)(pSizeInfo.Raw()))
}

// GetPartitionedAccelerationStructuresBuildSizesNV wraps vkGetPartitionedAccelerationStructuresBuildSizesNV.
func GetPartitionedAccelerationStructuresBuildSizesNV(device Device, pInfo PartitionedAccelerationStructureInstancesInputNV, pSizeInfo AccelerationStructureBuildSizesInfoKHR) {
	C.vkGetPartitionedAccelerationStructuresBuildSizesNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkPartitionedAccelerationStructureInstancesInputNV)(pInfo.Raw()), (*C.VkAccelerationStructureBuildSizesInfoKHR)(pSizeInfo.Raw()))
}

// GetPastPresentationTimingGOOGLE wraps vkGetPastPresentationTimingGOOGLE.
func GetPastPresentationTimingGOOGLE(device Device, swapchain SwapchainKHR, pPresentationTimingCount ffi.Ref[uint32], pPresentationTimings PastPresentationTimingGOOGLE) Result {
	ret := C.vkGetPastPresentationTimingGOOGLE(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.uint32_t)(pPresentationTimingCount.Raw()), (*C.VkPastPresentationTimingGOOGLE)(pPresentationTimings.Raw()))

	return Result(ret)
}

// GetPerformanceParameterINTEL wraps vkGetPerformanceParameterINTEL.
func GetPerformanceParameterINTEL(device Device, parameter PerformanceParameterTypeINTEL, pValue PerformanceValueINTEL) Result {
	ret := C.vkGetPerformanceParameterINTEL(C.VkDevice(unsafe.Pointer(device)), C.VkPerformanceParameterTypeINTEL(parameter), (*C.VkPerformanceValueINTEL)(pValue.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceCalibrateableTimeDomainsKHR wraps vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.
func GetPhysicalDeviceCalibrateableTimeDomainsKHR(physicalDevice PhysicalDevice, pTimeDomainCount ffi.Ref[uint32], pTimeDomains ffi.Ref[TimeDomainKHR]) Result {
	ret := C.vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pTimeDomainCount.Raw()), (*C.VkTimeDomainKHR)(pTimeDomains.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV wraps vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV.
func GetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties CooperativeMatrixFlexibleDimensionsPropertiesNV) Result {
	ret := C.vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkCooperativeMatrixFlexibleDimensionsPropertiesNV)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceCooperativeMatrixPropertiesKHR wraps vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR.
func GetPhysicalDeviceCooperativeMatrixPropertiesKHR(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties CooperativeMatrixPropertiesKHR) Result {
	ret := C.vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkCooperativeMatrixPropertiesKHR)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceCooperativeMatrixPropertiesNV wraps vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.
func GetPhysicalDeviceCooperativeMatrixPropertiesNV(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties CooperativeMatrixPropertiesNV) Result {
	ret := C.vkGetPhysicalDeviceCooperativeMatrixPropertiesNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkCooperativeMatrixPropertiesNV)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceCooperativeVectorPropertiesNV wraps vkGetPhysicalDeviceCooperativeVectorPropertiesNV.
func GetPhysicalDeviceCooperativeVectorPropertiesNV(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties CooperativeVectorPropertiesNV) Result {
	ret := C.vkGetPhysicalDeviceCooperativeVectorPropertiesNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkCooperativeVectorPropertiesNV)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceDisplayPlaneProperties2KHR wraps vkGetPhysicalDeviceDisplayPlaneProperties2KHR.
func GetPhysicalDeviceDisplayPlaneProperties2KHR(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties DisplayPlaneProperties2KHR) Result {
	ret := C.vkGetPhysicalDeviceDisplayPlaneProperties2KHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkDisplayPlaneProperties2KHR)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceDisplayPlanePropertiesKHR wraps vkGetPhysicalDeviceDisplayPlanePropertiesKHR.
func GetPhysicalDeviceDisplayPlanePropertiesKHR(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties DisplayPlanePropertiesKHR) Result {
	ret := C.vkGetPhysicalDeviceDisplayPlanePropertiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkDisplayPlanePropertiesKHR)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceDisplayProperties2KHR wraps vkGetPhysicalDeviceDisplayProperties2KHR.
func GetPhysicalDeviceDisplayProperties2KHR(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties DisplayProperties2KHR) Result {
	ret := C.vkGetPhysicalDeviceDisplayProperties2KHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkDisplayProperties2KHR)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceDisplayPropertiesKHR wraps vkGetPhysicalDeviceDisplayPropertiesKHR.
func GetPhysicalDeviceDisplayPropertiesKHR(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties DisplayPropertiesKHR) Result {
	ret := C.vkGetPhysicalDeviceDisplayPropertiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkDisplayPropertiesKHR)(pProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceExternalBufferProperties wraps vkGetPhysicalDeviceExternalBufferProperties.
func GetPhysicalDeviceExternalBufferProperties(physicalDevice PhysicalDevice, pExternalBufferInfo PhysicalDeviceExternalBufferInfo, pExternalBufferProperties ExternalBufferProperties) {
	C.vkGetPhysicalDeviceExternalBufferProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceExternalBufferInfo)(pExternalBufferInfo.Raw()), (*C.VkExternalBufferProperties)(pExternalBufferProperties.Raw()))
}

// GetPhysicalDeviceExternalFenceProperties wraps vkGetPhysicalDeviceExternalFenceProperties.
func GetPhysicalDeviceExternalFenceProperties(physicalDevice PhysicalDevice, pExternalFenceInfo PhysicalDeviceExternalFenceInfo, pExternalFenceProperties ExternalFenceProperties) {
	C.vkGetPhysicalDeviceExternalFenceProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceExternalFenceInfo)(pExternalFenceInfo.Raw()), (*C.VkExternalFenceProperties)(pExternalFenceProperties.Raw()))
}

// GetPhysicalDeviceExternalImageFormatPropertiesNV wraps vkGetPhysicalDeviceExternalImageFormatPropertiesNV.
func GetPhysicalDeviceExternalImageFormatPropertiesNV(physicalDevice PhysicalDevice, format Format, type_ ImageType, tiling ImageTiling, usage ImageUsageFlags, flags ImageCreateFlags, externalHandleType ExternalMemoryHandleTypeFlagsNV, pExternalImageFormatProperties ExternalImageFormatPropertiesNV) Result {
	ret := C.vkGetPhysicalDeviceExternalImageFormatPropertiesNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), C.VkImageType(type_), C.VkImageTiling(tiling), C.VkImageUsageFlags(usage), C.VkImageCreateFlags(flags), C.VkExternalMemoryHandleTypeFlagsNV(externalHandleType), (*C.VkExternalImageFormatPropertiesNV)(pExternalImageFormatProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceExternalSemaphoreProperties wraps vkGetPhysicalDeviceExternalSemaphoreProperties.
func GetPhysicalDeviceExternalSemaphoreProperties(physicalDevice PhysicalDevice, pExternalSemaphoreInfo PhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties ExternalSemaphoreProperties) {
	C.vkGetPhysicalDeviceExternalSemaphoreProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceExternalSemaphoreInfo)(pExternalSemaphoreInfo.Raw()), (*C.VkExternalSemaphoreProperties)(pExternalSemaphoreProperties.Raw()))
}

// GetPhysicalDeviceExternalTensorPropertiesARM wraps vkGetPhysicalDeviceExternalTensorPropertiesARM.
func GetPhysicalDeviceExternalTensorPropertiesARM(physicalDevice PhysicalDevice, pExternalTensorInfo PhysicalDeviceExternalTensorInfoARM, pExternalTensorProperties ExternalTensorPropertiesARM) {
	C.vkGetPhysicalDeviceExternalTensorPropertiesARM(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceExternalTensorInfoARM)(pExternalTensorInfo.Raw()), (*C.VkExternalTensorPropertiesARM)(pExternalTensorProperties.Raw()))
}

// GetPhysicalDeviceFeatures wraps vkGetPhysicalDeviceFeatures.
func GetPhysicalDeviceFeatures(physicalDevice PhysicalDevice, pFeatures PhysicalDeviceFeatures) {
	C.vkGetPhysicalDeviceFeatures(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceFeatures)(pFeatures.Raw()))
}

// GetPhysicalDeviceFeatures2 wraps vkGetPhysicalDeviceFeatures2.
func GetPhysicalDeviceFeatures2(physicalDevice PhysicalDevice, pFeatures PhysicalDeviceFeatures2) {
	C.vkGetPhysicalDeviceFeatures2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceFeatures2)(pFeatures.Raw()))
}

// GetPhysicalDeviceFormatProperties wraps vkGetPhysicalDeviceFormatProperties.
func GetPhysicalDeviceFormatProperties(physicalDevice PhysicalDevice, format Format, pFormatProperties FormatProperties) {
	C.vkGetPhysicalDeviceFormatProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), (*C.VkFormatProperties)(pFormatProperties.Raw()))
}

// GetPhysicalDeviceFormatProperties2 wraps vkGetPhysicalDeviceFormatProperties2.
func GetPhysicalDeviceFormatProperties2(physicalDevice PhysicalDevice, format Format, pFormatProperties FormatProperties2) {
	C.vkGetPhysicalDeviceFormatProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), (*C.VkFormatProperties2)(pFormatProperties.Raw()))
}

// GetPhysicalDeviceFragmentShadingRatesKHR wraps vkGetPhysicalDeviceFragmentShadingRatesKHR.
func GetPhysicalDeviceFragmentShadingRatesKHR(physicalDevice PhysicalDevice, pFragmentShadingRateCount ffi.Ref[uint32], pFragmentShadingRates PhysicalDeviceFragmentShadingRateKHR) Result {
	ret := C.vkGetPhysicalDeviceFragmentShadingRatesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pFragmentShadingRateCount.Raw()), (*C.VkPhysicalDeviceFragmentShadingRateKHR)(pFragmentShadingRates.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceImageFormatProperties wraps vkGetPhysicalDeviceImageFormatProperties.
func GetPhysicalDeviceImageFormatProperties(physicalDevice PhysicalDevice, format Format, type_ ImageType, tiling ImageTiling, usage ImageUsageFlags, flags ImageCreateFlags, pImageFormatProperties ImageFormatProperties) Result {
	ret := C.vkGetPhysicalDeviceImageFormatProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), C.VkImageType(type_), C.VkImageTiling(tiling), C.VkImageUsageFlags(usage), C.VkImageCreateFlags(flags), (*C.VkImageFormatProperties)(pImageFormatProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceImageFormatProperties2 wraps vkGetPhysicalDeviceImageFormatProperties2.
func GetPhysicalDeviceImageFormatProperties2(physicalDevice PhysicalDevice, pImageFormatInfo PhysicalDeviceImageFormatInfo2, pImageFormatProperties ImageFormatProperties2) Result {
	ret := C.vkGetPhysicalDeviceImageFormatProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceImageFormatInfo2)(pImageFormatInfo.Raw()), (*C.VkImageFormatProperties2)(pImageFormatProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceMemoryProperties wraps vkGetPhysicalDeviceMemoryProperties.
func GetPhysicalDeviceMemoryProperties(physicalDevice PhysicalDevice, pMemoryProperties PhysicalDeviceMemoryProperties) {
	C.vkGetPhysicalDeviceMemoryProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceMemoryProperties)(pMemoryProperties.Raw()))
}

// GetPhysicalDeviceMemoryProperties2 wraps vkGetPhysicalDeviceMemoryProperties2.
func GetPhysicalDeviceMemoryProperties2(physicalDevice PhysicalDevice, pMemoryProperties PhysicalDeviceMemoryProperties2) {
	C.vkGetPhysicalDeviceMemoryProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceMemoryProperties2)(pMemoryProperties.Raw()))
}

// GetPhysicalDeviceMultisamplePropertiesEXT wraps vkGetPhysicalDeviceMultisamplePropertiesEXT.
func GetPhysicalDeviceMultisamplePropertiesEXT(physicalDevice PhysicalDevice, samples SampleCountFlags, pMultisampleProperties MultisamplePropertiesEXT) {
	C.vkGetPhysicalDeviceMultisamplePropertiesEXT(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkSampleCountFlagBits(samples), (*C.VkMultisamplePropertiesEXT)(pMultisampleProperties.Raw()))
}

// GetPhysicalDeviceOpticalFlowImageFormatsNV wraps vkGetPhysicalDeviceOpticalFlowImageFormatsNV.
func GetPhysicalDeviceOpticalFlowImageFormatsNV(physicalDevice PhysicalDevice, pOpticalFlowImageFormatInfo OpticalFlowImageFormatInfoNV, pFormatCount ffi.Ref[uint32], pImageFormatProperties OpticalFlowImageFormatPropertiesNV) Result {
	ret := C.vkGetPhysicalDeviceOpticalFlowImageFormatsNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkOpticalFlowImageFormatInfoNV)(pOpticalFlowImageFormatInfo.Raw()), (*C.uint32_t)(pFormatCount.Raw()), (*C.VkOpticalFlowImageFormatPropertiesNV)(pImageFormatProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDevicePresentRectanglesKHR wraps vkGetPhysicalDevicePresentRectanglesKHR.
func GetPhysicalDevicePresentRectanglesKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, pRectCount ffi.Ref[uint32], pRects Rect2D) Result {
	ret := C.vkGetPhysicalDevicePresentRectanglesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkSurfaceKHR(surface), (*C.uint32_t)(pRectCount.Raw()), (*C.VkRect2D)(pRects.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceProperties wraps vkGetPhysicalDeviceProperties.
func GetPhysicalDeviceProperties(physicalDevice PhysicalDevice, pProperties PhysicalDeviceProperties) {
	C.vkGetPhysicalDeviceProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceProperties)(pProperties.Raw()))
}

// GetPhysicalDeviceProperties2 wraps vkGetPhysicalDeviceProperties2.
func GetPhysicalDeviceProperties2(physicalDevice PhysicalDevice, pProperties PhysicalDeviceProperties2) {
	C.vkGetPhysicalDeviceProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceProperties2)(pProperties.Raw()))
}

// GetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM wraps vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.
func GetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM(physicalDevice PhysicalDevice, pQueueFamilyDataGraphProcessingEngineInfo PhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM, pQueueFamilyDataGraphProcessingEngineProperties QueueFamilyDataGraphProcessingEnginePropertiesARM) {
	C.vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM)(pQueueFamilyDataGraphProcessingEngineInfo.Raw()), (*C.VkQueueFamilyDataGraphProcessingEnginePropertiesARM)(pQueueFamilyDataGraphProcessingEngineProperties.Raw()))
}

// GetPhysicalDeviceQueueFamilyDataGraphPropertiesARM wraps vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.
func GetPhysicalDeviceQueueFamilyDataGraphPropertiesARM(physicalDevice PhysicalDevice, queueFamilyIndex uint32, pQueueFamilyDataGraphPropertyCount ffi.Ref[uint32], pQueueFamilyDataGraphProperties QueueFamilyDataGraphPropertiesARM) Result {
	ret := C.vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.uint32_t(queueFamilyIndex), (*C.uint32_t)(pQueueFamilyDataGraphPropertyCount.Raw()), (*C.VkQueueFamilyDataGraphPropertiesARM)(pQueueFamilyDataGraphProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR wraps vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.
func GetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(physicalDevice PhysicalDevice, pPerformanceQueryCreateInfo QueryPoolPerformanceCreateInfoKHR, pNumPasses ffi.Ref[uint32]) {
	C.vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkQueryPoolPerformanceCreateInfoKHR)(pPerformanceQueryCreateInfo.Raw()), (*C.uint32_t)(pNumPasses.Raw()))
}

// GetPhysicalDeviceQueueFamilyProperties wraps vkGetPhysicalDeviceQueueFamilyProperties.
func GetPhysicalDeviceQueueFamilyProperties(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount ffi.Ref[uint32], pQueueFamilyProperties QueueFamilyProperties) {
	C.vkGetPhysicalDeviceQueueFamilyProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pQueueFamilyPropertyCount.Raw()), (*C.VkQueueFamilyProperties)(pQueueFamilyProperties.Raw()))
}

// GetPhysicalDeviceQueueFamilyProperties2 wraps vkGetPhysicalDeviceQueueFamilyProperties2.
func GetPhysicalDeviceQueueFamilyProperties2(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount ffi.Ref[uint32], pQueueFamilyProperties QueueFamilyProperties2) {
	C.vkGetPhysicalDeviceQueueFamilyProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pQueueFamilyPropertyCount.Raw()), (*C.VkQueueFamilyProperties2)(pQueueFamilyProperties.Raw()))
}

// GetPhysicalDeviceSparseImageFormatProperties wraps vkGetPhysicalDeviceSparseImageFormatProperties.
func GetPhysicalDeviceSparseImageFormatProperties(physicalDevice PhysicalDevice, format Format, type_ ImageType, samples SampleCountFlags, usage ImageUsageFlags, tiling ImageTiling, pPropertyCount ffi.Ref[uint32], pProperties SparseImageFormatProperties) {
	C.vkGetPhysicalDeviceSparseImageFormatProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), C.VkImageType(type_), C.VkSampleCountFlagBits(samples), C.VkImageUsageFlags(usage), C.VkImageTiling(tiling), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkSparseImageFormatProperties)(pProperties.Raw()))
}

// GetPhysicalDeviceSparseImageFormatProperties2 wraps vkGetPhysicalDeviceSparseImageFormatProperties2.
func GetPhysicalDeviceSparseImageFormatProperties2(physicalDevice PhysicalDevice, pFormatInfo PhysicalDeviceSparseImageFormatInfo2, pPropertyCount ffi.Ref[uint32], pProperties SparseImageFormatProperties2) {
	C.vkGetPhysicalDeviceSparseImageFormatProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceSparseImageFormatInfo2)(pFormatInfo.Raw()), (*C.uint32_t)(pPropertyCount.Raw()), (*C.VkSparseImageFormatProperties2)(pProperties.Raw()))
}

// GetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV wraps vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.
func GetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(physicalDevice PhysicalDevice, pCombinationCount ffi.Ref[uint32], pCombinations FramebufferMixedSamplesCombinationNV) Result {
	ret := C.vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pCombinationCount.Raw()), (*C.VkFramebufferMixedSamplesCombinationNV)(pCombinations.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfaceCapabilities2EXT wraps vkGetPhysicalDeviceSurfaceCapabilities2EXT.
func GetPhysicalDeviceSurfaceCapabilities2EXT(physicalDevice PhysicalDevice, surface SurfaceKHR, pSurfaceCapabilities SurfaceCapabilities2EXT) Result {
	ret := C.vkGetPhysicalDeviceSurfaceCapabilities2EXT(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkSurfaceKHR(surface), (*C.VkSurfaceCapabilities2EXT)(pSurfaceCapabilities.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfaceCapabilities2KHR wraps vkGetPhysicalDeviceSurfaceCapabilities2KHR.
func GetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice PhysicalDevice, pSurfaceInfo PhysicalDeviceSurfaceInfo2KHR, pSurfaceCapabilities SurfaceCapabilities2KHR) Result {
	ret := C.vkGetPhysicalDeviceSurfaceCapabilities2KHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceSurfaceInfo2KHR)(pSurfaceInfo.Raw()), (*C.VkSurfaceCapabilities2KHR)(pSurfaceCapabilities.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfaceCapabilitiesKHR wraps vkGetPhysicalDeviceSurfaceCapabilitiesKHR.
func GetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, pSurfaceCapabilities SurfaceCapabilitiesKHR) Result {
	ret := C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkSurfaceKHR(surface), (*C.VkSurfaceCapabilitiesKHR)(pSurfaceCapabilities.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfaceFormats2KHR wraps vkGetPhysicalDeviceSurfaceFormats2KHR.
func GetPhysicalDeviceSurfaceFormats2KHR(physicalDevice PhysicalDevice, pSurfaceInfo PhysicalDeviceSurfaceInfo2KHR, pSurfaceFormatCount ffi.Ref[uint32], pSurfaceFormats SurfaceFormat2KHR) Result {
	ret := C.vkGetPhysicalDeviceSurfaceFormats2KHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceSurfaceInfo2KHR)(pSurfaceInfo.Raw()), (*C.uint32_t)(pSurfaceFormatCount.Raw()), (*C.VkSurfaceFormat2KHR)(pSurfaceFormats.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfaceFormatsKHR wraps vkGetPhysicalDeviceSurfaceFormatsKHR.
func GetPhysicalDeviceSurfaceFormatsKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, pSurfaceFormatCount ffi.Ref[uint32], pSurfaceFormats SurfaceFormatKHR) Result {
	ret := C.vkGetPhysicalDeviceSurfaceFormatsKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkSurfaceKHR(surface), (*C.uint32_t)(pSurfaceFormatCount.Raw()), (*C.VkSurfaceFormatKHR)(pSurfaceFormats.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfacePresentModes2EXT wraps vkGetPhysicalDeviceSurfacePresentModes2EXT.
func GetPhysicalDeviceSurfacePresentModes2EXT(physicalDevice PhysicalDevice, pSurfaceInfo PhysicalDeviceSurfaceInfo2KHR, pPresentModeCount ffi.Ref[uint32], pPresentModes ffi.Ref[PresentModeKHR]) Result {
	ret := C.vkGetPhysicalDeviceSurfacePresentModes2EXT(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceSurfaceInfo2KHR)(pSurfaceInfo.Raw()), (*C.uint32_t)(pPresentModeCount.Raw()), (*C.VkPresentModeKHR)(pPresentModes.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceSurfacePresentModesKHR wraps vkGetPhysicalDeviceSurfacePresentModesKHR.
func GetPhysicalDeviceSurfacePresentModesKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, pPresentModeCount ffi.Ref[uint32], pPresentModes ffi.Ref[PresentModeKHR]) Result {
	ret := C.vkGetPhysicalDeviceSurfacePresentModesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkSurfaceKHR(surface), (*C.uint32_t)(pPresentModeCount.Raw()), (*C.VkPresentModeKHR)(pPresentModes.Raw()))

	return Result(ret)
}

// vkGetPhysicalDeviceSurfaceSupportKHR.pSupported is unsupported: category pointer -> ?? VkBool32.

// GetPhysicalDeviceToolProperties wraps vkGetPhysicalDeviceToolProperties.
func GetPhysicalDeviceToolProperties(physicalDevice PhysicalDevice, pToolCount ffi.Ref[uint32], pToolProperties PhysicalDeviceToolProperties) Result {
	ret := C.vkGetPhysicalDeviceToolProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pToolCount.Raw()), (*C.VkPhysicalDeviceToolProperties)(pToolProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceVideoCapabilitiesKHR wraps vkGetPhysicalDeviceVideoCapabilitiesKHR.
func GetPhysicalDeviceVideoCapabilitiesKHR(physicalDevice PhysicalDevice, pVideoProfile VideoProfileInfoKHR, pCapabilities VideoCapabilitiesKHR) Result {
	ret := C.vkGetPhysicalDeviceVideoCapabilitiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkVideoProfileInfoKHR)(pVideoProfile.Raw()), (*C.VkVideoCapabilitiesKHR)(pCapabilities.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR wraps vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.
func GetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR(physicalDevice PhysicalDevice, pQualityLevelInfo PhysicalDeviceVideoEncodeQualityLevelInfoKHR, pQualityLevelProperties VideoEncodeQualityLevelPropertiesKHR) Result {
	ret := C.vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR)(pQualityLevelInfo.Raw()), (*C.VkVideoEncodeQualityLevelPropertiesKHR)(pQualityLevelProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceVideoFormatPropertiesKHR wraps vkGetPhysicalDeviceVideoFormatPropertiesKHR.
func GetPhysicalDeviceVideoFormatPropertiesKHR(physicalDevice PhysicalDevice, pVideoFormatInfo PhysicalDeviceVideoFormatInfoKHR, pVideoFormatPropertyCount ffi.Ref[uint32], pVideoFormatProperties VideoFormatPropertiesKHR) Result {
	ret := C.vkGetPhysicalDeviceVideoFormatPropertiesKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.VkPhysicalDeviceVideoFormatInfoKHR)(pVideoFormatInfo.Raw()), (*C.uint32_t)(pVideoFormatPropertyCount.Raw()), (*C.VkVideoFormatPropertiesKHR)(pVideoFormatProperties.Raw()))

	return Result(ret)
}

// GetPhysicalDeviceWin32PresentationSupportKHR wraps vkGetPhysicalDeviceWin32PresentationSupportKHR.
func GetPhysicalDeviceWin32PresentationSupportKHR(physicalDevice PhysicalDevice, queueFamilyIndex uint32) bool {
	ret := C.vkGetPhysicalDeviceWin32PresentationSupportKHR(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.uint32_t(queueFamilyIndex))

	if ret > 0 {
		return true
	} else {
		return false
	}
}

// GetPipelineBinaryDataKHR wraps vkGetPipelineBinaryDataKHR.
func GetPipelineBinaryDataKHR(device Device, pInfo PipelineBinaryDataInfoKHR, pPipelineBinaryKey PipelineBinaryKeyKHR, pPipelineBinaryDataSize ffi.Ref[uintptr], pPipelineBinaryData unsafe.Pointer) Result {
	ret := C.vkGetPipelineBinaryDataKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineBinaryDataInfoKHR)(pInfo.Raw()), (*C.VkPipelineBinaryKeyKHR)(pPipelineBinaryKey.Raw()), (*C.size_t)(pPipelineBinaryDataSize.Raw()), pPipelineBinaryData)

	return Result(ret)
}

// GetPipelineCacheData wraps vkGetPipelineCacheData.
func GetPipelineCacheData(device Device, pipelineCache PipelineCache, pDataSize ffi.Ref[uintptr], pData unsafe.Pointer) Result {
	ret := C.vkGetPipelineCacheData(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), (*C.size_t)(pDataSize.Raw()), pData)

	return Result(ret)
}

// GetPipelineExecutableInternalRepresentationsKHR wraps vkGetPipelineExecutableInternalRepresentationsKHR.
func GetPipelineExecutableInternalRepresentationsKHR(device Device, pExecutableInfo PipelineExecutableInfoKHR, pInternalRepresentationCount ffi.Ref[uint32], pInternalRepresentations PipelineExecutableInternalRepresentationKHR) Result {
	ret := C.vkGetPipelineExecutableInternalRepresentationsKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineExecutableInfoKHR)(pExecutableInfo.Raw()), (*C.uint32_t)(pInternalRepresentationCount.Raw()), (*C.VkPipelineExecutableInternalRepresentationKHR)(pInternalRepresentations.Raw()))

	return Result(ret)
}

// GetPipelineExecutablePropertiesKHR wraps vkGetPipelineExecutablePropertiesKHR.
func GetPipelineExecutablePropertiesKHR(device Device, pPipelineInfo PipelineInfoKHR, pExecutableCount ffi.Ref[uint32], pProperties PipelineExecutablePropertiesKHR) Result {
	ret := C.vkGetPipelineExecutablePropertiesKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineInfoKHR)(pPipelineInfo.Raw()), (*C.uint32_t)(pExecutableCount.Raw()), (*C.VkPipelineExecutablePropertiesKHR)(pProperties.Raw()))

	return Result(ret)
}

// GetPipelineExecutableStatisticsKHR wraps vkGetPipelineExecutableStatisticsKHR.
func GetPipelineExecutableStatisticsKHR(device Device, pExecutableInfo PipelineExecutableInfoKHR, pStatisticCount ffi.Ref[uint32], pStatistics PipelineExecutableStatisticKHR) Result {
	ret := C.vkGetPipelineExecutableStatisticsKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineExecutableInfoKHR)(pExecutableInfo.Raw()), (*C.uint32_t)(pStatisticCount.Raw()), (*C.VkPipelineExecutableStatisticKHR)(pStatistics.Raw()))

	return Result(ret)
}

// vkGetPipelineIndirectDeviceAddressNV is unsupported: unknown category direct.

// GetPipelineIndirectMemoryRequirementsNV wraps vkGetPipelineIndirectMemoryRequirementsNV.
func GetPipelineIndirectMemoryRequirementsNV(device Device, pCreateInfo ComputePipelineCreateInfo, pMemoryRequirements MemoryRequirements2) {
	C.vkGetPipelineIndirectMemoryRequirementsNV(C.VkDevice(unsafe.Pointer(device)), (*C.VkComputePipelineCreateInfo)(pCreateInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetPipelineKeyKHR wraps vkGetPipelineKeyKHR.
func GetPipelineKeyKHR(device Device, pPipelineCreateInfo PipelineCreateInfoKHR, pPipelineKey PipelineBinaryKeyKHR) Result {
	ret := C.vkGetPipelineKeyKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineCreateInfoKHR)(pPipelineCreateInfo.Raw()), (*C.VkPipelineBinaryKeyKHR)(pPipelineKey.Raw()))

	return Result(ret)
}

// GetPipelinePropertiesEXT wraps vkGetPipelinePropertiesEXT.
func GetPipelinePropertiesEXT(device Device, pPipelineInfo PipelineInfoKHR, pPipelineProperties BaseOutStructure) Result {
	ret := C.vkGetPipelinePropertiesEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkPipelineInfoEXT)(pPipelineInfo.Raw()), (*C.VkBaseOutStructure)(pPipelineProperties.Raw()))

	return Result(ret)
}

// GetPrivateData wraps vkGetPrivateData.
func GetPrivateData(device Device, objectType ObjectType, objectHandle uint64, privateDataSlot PrivateDataSlot, pData ffi.Ref[uint64]) {
	C.vkGetPrivateData(C.VkDevice(unsafe.Pointer(device)), C.VkObjectType(objectType), C.uint64_t(objectHandle), C.VkPrivateDataSlot(privateDataSlot), (*C.uint64_t)(pData.Raw()))
}

// GetQueryPoolResults wraps vkGetQueryPoolResults.
func GetQueryPoolResults(device Device, queryPool QueryPool, firstQuery uint32, queryCount uint32, dataSize uintptr, pData unsafe.Pointer, stride DeviceSize, flags QueryResultFlags) Result {
	ret := C.vkGetQueryPoolResults(C.VkDevice(unsafe.Pointer(device)), C.VkQueryPool(queryPool), C.uint32_t(firstQuery), C.uint32_t(queryCount), C.size_t(dataSize), pData, C.VkDeviceSize(stride), C.VkQueryResultFlags(flags))

	return Result(ret)
}

// GetQueueCheckpointData2NV wraps vkGetQueueCheckpointData2NV.
func GetQueueCheckpointData2NV(queue Queue, pCheckpointDataCount ffi.Ref[uint32], pCheckpointData CheckpointData2NV) {
	C.vkGetQueueCheckpointData2NV(C.VkQueue(unsafe.Pointer(queue)), (*C.uint32_t)(pCheckpointDataCount.Raw()), (*C.VkCheckpointData2NV)(pCheckpointData.Raw()))
}

// GetQueueCheckpointDataNV wraps vkGetQueueCheckpointDataNV.
func GetQueueCheckpointDataNV(queue Queue, pCheckpointDataCount ffi.Ref[uint32], pCheckpointData CheckpointDataNV) {
	C.vkGetQueueCheckpointDataNV(C.VkQueue(unsafe.Pointer(queue)), (*C.uint32_t)(pCheckpointDataCount.Raw()), (*C.VkCheckpointDataNV)(pCheckpointData.Raw()))
}

// GetRayTracingCaptureReplayShaderGroupHandlesKHR wraps vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.
func GetRayTracingCaptureReplayShaderGroupHandlesKHR(device Device, pipeline Pipeline, firstGroup uint32, groupCount uint32, dataSize uintptr, pData unsafe.Pointer) Result {
	ret := C.vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(C.VkDevice(unsafe.Pointer(device)), C.VkPipeline(pipeline), C.uint32_t(firstGroup), C.uint32_t(groupCount), C.size_t(dataSize), pData)

	return Result(ret)
}

// GetRayTracingShaderGroupHandlesKHR wraps vkGetRayTracingShaderGroupHandlesKHR.
func GetRayTracingShaderGroupHandlesKHR(device Device, pipeline Pipeline, firstGroup uint32, groupCount uint32, dataSize uintptr, pData unsafe.Pointer) Result {
	ret := C.vkGetRayTracingShaderGroupHandlesKHR(C.VkDevice(unsafe.Pointer(device)), C.VkPipeline(pipeline), C.uint32_t(firstGroup), C.uint32_t(groupCount), C.size_t(dataSize), pData)

	return Result(ret)
}

// vkGetRayTracingShaderGroupStackSizeKHR is unsupported: unknown category direct.

// GetRefreshCycleDurationGOOGLE wraps vkGetRefreshCycleDurationGOOGLE.
func GetRefreshCycleDurationGOOGLE(device Device, swapchain SwapchainKHR, pDisplayTimingProperties RefreshCycleDurationGOOGLE) Result {
	ret := C.vkGetRefreshCycleDurationGOOGLE(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkRefreshCycleDurationGOOGLE)(pDisplayTimingProperties.Raw()))

	return Result(ret)
}

// GetRenderAreaGranularity wraps vkGetRenderAreaGranularity.
func GetRenderAreaGranularity(device Device, renderPass RenderPass, pGranularity Extent2D) {
	C.vkGetRenderAreaGranularity(C.VkDevice(unsafe.Pointer(device)), C.VkRenderPass(renderPass), (*C.VkExtent2D)(pGranularity.Raw()))
}

// GetRenderingAreaGranularity wraps vkGetRenderingAreaGranularity.
func GetRenderingAreaGranularity(device Device, pRenderingAreaInfo RenderingAreaInfo, pGranularity Extent2D) {
	C.vkGetRenderingAreaGranularity(C.VkDevice(unsafe.Pointer(device)), (*C.VkRenderingAreaInfo)(pRenderingAreaInfo.Raw()), (*C.VkExtent2D)(pGranularity.Raw()))
}

// GetSamplerOpaqueCaptureDescriptorDataEXT wraps vkGetSamplerOpaqueCaptureDescriptorDataEXT.
func GetSamplerOpaqueCaptureDescriptorDataEXT(device Device, pInfo SamplerCaptureDescriptorDataInfoEXT, pData unsafe.Pointer) Result {
	ret := C.vkGetSamplerOpaqueCaptureDescriptorDataEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkSamplerCaptureDescriptorDataInfoEXT)(pInfo.Raw()), pData)

	return Result(ret)
}

// GetSemaphoreCounterValue wraps vkGetSemaphoreCounterValue.
func GetSemaphoreCounterValue(device Device, semaphore Semaphore, pValue ffi.Ref[uint64]) Result {
	ret := C.vkGetSemaphoreCounterValue(C.VkDevice(unsafe.Pointer(device)), C.VkSemaphore(semaphore), (*C.uint64_t)(pValue.Raw()))

	return Result(ret)
}

// vkGetSemaphoreFdKHR.pFd is unsupported: category pointer -> ?? int.

// vkGetSemaphoreWin32HandleKHR.pHandle is unsupported: category pointer -> ?? HANDLE.

// GetShaderBinaryDataEXT wraps vkGetShaderBinaryDataEXT.
func GetShaderBinaryDataEXT(device Device, shader ShaderEXT, pDataSize ffi.Ref[uintptr], pData unsafe.Pointer) Result {
	ret := C.vkGetShaderBinaryDataEXT(C.VkDevice(unsafe.Pointer(device)), C.VkShaderEXT(shader), (*C.size_t)(pDataSize.Raw()), pData)

	return Result(ret)
}

// GetShaderInfoAMD wraps vkGetShaderInfoAMD.
func GetShaderInfoAMD(device Device, pipeline Pipeline, shaderStage ShaderStageFlags, infoType ShaderInfoTypeAMD, pInfoSize ffi.Ref[uintptr], pInfo unsafe.Pointer) Result {
	ret := C.vkGetShaderInfoAMD(C.VkDevice(unsafe.Pointer(device)), C.VkPipeline(pipeline), C.VkShaderStageFlagBits(shaderStage), C.VkShaderInfoTypeAMD(infoType), (*C.size_t)(pInfoSize.Raw()), pInfo)

	return Result(ret)
}

// GetShaderModuleCreateInfoIdentifierEXT wraps vkGetShaderModuleCreateInfoIdentifierEXT.
func GetShaderModuleCreateInfoIdentifierEXT(device Device, pCreateInfo ShaderModuleCreateInfo, pIdentifier ShaderModuleIdentifierEXT) {
	C.vkGetShaderModuleCreateInfoIdentifierEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkShaderModuleCreateInfo)(pCreateInfo.Raw()), (*C.VkShaderModuleIdentifierEXT)(pIdentifier.Raw()))
}

// GetShaderModuleIdentifierEXT wraps vkGetShaderModuleIdentifierEXT.
func GetShaderModuleIdentifierEXT(device Device, shaderModule ShaderModule, pIdentifier ShaderModuleIdentifierEXT) {
	C.vkGetShaderModuleIdentifierEXT(C.VkDevice(unsafe.Pointer(device)), C.VkShaderModule(shaderModule), (*C.VkShaderModuleIdentifierEXT)(pIdentifier.Raw()))
}

// GetSwapchainCounterEXT wraps vkGetSwapchainCounterEXT.
func GetSwapchainCounterEXT(device Device, swapchain SwapchainKHR, counter SurfaceCounterFlagsEXT, pCounterValue ffi.Ref[uint64]) Result {
	ret := C.vkGetSwapchainCounterEXT(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), C.VkSurfaceCounterFlagBitsEXT(counter), (*C.uint64_t)(pCounterValue.Raw()))

	return Result(ret)
}

// GetSwapchainImagesKHR wraps vkGetSwapchainImagesKHR.
func GetSwapchainImagesKHR(device Device, swapchain SwapchainKHR, pSwapchainImageCount ffi.Ref[uint32], pSwapchainImages ffi.Ref[Image]) Result {
	ret := C.vkGetSwapchainImagesKHR(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.uint32_t)(pSwapchainImageCount.Raw()), (*C.VkImage)(pSwapchainImages.Raw()))

	return Result(ret)
}

// GetSwapchainStatusKHR wraps vkGetSwapchainStatusKHR.
func GetSwapchainStatusKHR(device Device, swapchain SwapchainKHR) Result {
	ret := C.vkGetSwapchainStatusKHR(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain))

	return Result(ret)
}

// GetTensorMemoryRequirementsARM wraps vkGetTensorMemoryRequirementsARM.
func GetTensorMemoryRequirementsARM(device Device, pInfo TensorMemoryRequirementsInfoARM, pMemoryRequirements MemoryRequirements2) {
	C.vkGetTensorMemoryRequirementsARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkTensorMemoryRequirementsInfoARM)(pInfo.Raw()), (*C.VkMemoryRequirements2)(pMemoryRequirements.Raw()))
}

// GetTensorOpaqueCaptureDescriptorDataARM wraps vkGetTensorOpaqueCaptureDescriptorDataARM.
func GetTensorOpaqueCaptureDescriptorDataARM(device Device, pInfo TensorCaptureDescriptorDataInfoARM, pData unsafe.Pointer) Result {
	ret := C.vkGetTensorOpaqueCaptureDescriptorDataARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkTensorCaptureDescriptorDataInfoARM)(pInfo.Raw()), pData)

	return Result(ret)
}

// GetTensorViewOpaqueCaptureDescriptorDataARM wraps vkGetTensorViewOpaqueCaptureDescriptorDataARM.
func GetTensorViewOpaqueCaptureDescriptorDataARM(device Device, pInfo TensorViewCaptureDescriptorDataInfoARM, pData unsafe.Pointer) Result {
	ret := C.vkGetTensorViewOpaqueCaptureDescriptorDataARM(C.VkDevice(unsafe.Pointer(device)), (*C.VkTensorViewCaptureDescriptorDataInfoARM)(pInfo.Raw()), pData)

	return Result(ret)
}

// GetValidationCacheDataEXT wraps vkGetValidationCacheDataEXT.
func GetValidationCacheDataEXT(device Device, validationCache ValidationCacheEXT, pDataSize ffi.Ref[uintptr], pData unsafe.Pointer) Result {
	ret := C.vkGetValidationCacheDataEXT(C.VkDevice(unsafe.Pointer(device)), C.VkValidationCacheEXT(validationCache), (*C.size_t)(pDataSize.Raw()), pData)

	return Result(ret)
}

// GetVideoSessionMemoryRequirementsKHR wraps vkGetVideoSessionMemoryRequirementsKHR.
func GetVideoSessionMemoryRequirementsKHR(device Device, videoSession VideoSessionKHR, pMemoryRequirementsCount ffi.Ref[uint32], pMemoryRequirements VideoSessionMemoryRequirementsKHR) Result {
	ret := C.vkGetVideoSessionMemoryRequirementsKHR(C.VkDevice(unsafe.Pointer(device)), C.VkVideoSessionKHR(videoSession), (*C.uint32_t)(pMemoryRequirementsCount.Raw()), (*C.VkVideoSessionMemoryRequirementsKHR)(pMemoryRequirements.Raw()))

	return Result(ret)
}

// GetWinrtDisplayNV wraps vkGetWinrtDisplayNV.
func GetWinrtDisplayNV(physicalDevice PhysicalDevice, deviceRelativeId uint32, pDisplay ffi.Ref[DisplayKHR]) Result {
	ret := C.vkGetWinrtDisplayNV(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.uint32_t(deviceRelativeId), (*C.VkDisplayKHR)(pDisplay.Raw()))

	return Result(ret)
}

// ImportFenceFdKHR wraps vkImportFenceFdKHR.
func ImportFenceFdKHR(device Device, pImportFenceFdInfo ImportFenceFdInfoKHR) Result {
	ret := C.vkImportFenceFdKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkImportFenceFdInfoKHR)(pImportFenceFdInfo.Raw()))

	return Result(ret)
}

// ImportFenceWin32HandleKHR wraps vkImportFenceWin32HandleKHR.
func ImportFenceWin32HandleKHR(device Device, pImportFenceWin32HandleInfo ImportFenceWin32HandleInfoKHR) Result {
	ret := C.vkImportFenceWin32HandleKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkImportFenceWin32HandleInfoKHR)(pImportFenceWin32HandleInfo.Raw()))

	return Result(ret)
}

// ImportSemaphoreFdKHR wraps vkImportSemaphoreFdKHR.
func ImportSemaphoreFdKHR(device Device, pImportSemaphoreFdInfo ImportSemaphoreFdInfoKHR) Result {
	ret := C.vkImportSemaphoreFdKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkImportSemaphoreFdInfoKHR)(pImportSemaphoreFdInfo.Raw()))

	return Result(ret)
}

// ImportSemaphoreWin32HandleKHR wraps vkImportSemaphoreWin32HandleKHR.
func ImportSemaphoreWin32HandleKHR(device Device, pImportSemaphoreWin32HandleInfo ImportSemaphoreWin32HandleInfoKHR) Result {
	ret := C.vkImportSemaphoreWin32HandleKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkImportSemaphoreWin32HandleInfoKHR)(pImportSemaphoreWin32HandleInfo.Raw()))

	return Result(ret)
}

// InitializePerformanceApiINTEL wraps vkInitializePerformanceApiINTEL.
func InitializePerformanceApiINTEL(device Device, pInitializeInfo InitializePerformanceApiInfoINTEL) Result {
	ret := C.vkInitializePerformanceApiINTEL(C.VkDevice(unsafe.Pointer(device)), (*C.VkInitializePerformanceApiInfoINTEL)(pInitializeInfo.Raw()))

	return Result(ret)
}

// InvalidateMappedMemoryRanges wraps vkInvalidateMappedMemoryRanges.
func InvalidateMappedMemoryRanges(device Device, memoryRangeCount uint32, pMemoryRanges MappedMemoryRange) Result {
	ret := C.vkInvalidateMappedMemoryRanges(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(memoryRangeCount), (*C.VkMappedMemoryRange)(pMemoryRanges.Raw()))

	return Result(ret)
}

// LatencySleepNV wraps vkLatencySleepNV.
func LatencySleepNV(device Device, swapchain SwapchainKHR, pSleepInfo LatencySleepInfoNV) Result {
	ret := C.vkLatencySleepNV(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkLatencySleepInfoNV)(pSleepInfo.Raw()))

	return Result(ret)
}

// MapMemory wraps vkMapMemory.
func MapMemory(device Device, memory DeviceMemory, offset DeviceSize, size DeviceSize, flags MemoryMapFlags, ppData ffi.Ref[unsafe.Pointer]) Result {
	ret := C.vkMapMemory(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory), C.VkDeviceSize(offset), C.VkDeviceSize(size), C.VkMemoryMapFlags(flags), (*unsafe.Pointer)(ppData.Raw()))

	return Result(ret)
}

// MapMemory2 wraps vkMapMemory2.
func MapMemory2(device Device, pMemoryMapInfo MemoryMapInfo, ppData ffi.Ref[unsafe.Pointer]) Result {
	ret := C.vkMapMemory2(C.VkDevice(unsafe.Pointer(device)), (*C.VkMemoryMapInfo)(pMemoryMapInfo.Raw()), (*unsafe.Pointer)(ppData.Raw()))

	return Result(ret)
}

// MergePipelineCaches wraps vkMergePipelineCaches.
func MergePipelineCaches(device Device, dstCache PipelineCache, srcCacheCount uint32, pSrcCaches ffi.Ref[PipelineCache]) Result {
	ret := C.vkMergePipelineCaches(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(dstCache), C.uint32_t(srcCacheCount), (*C.VkPipelineCache)(pSrcCaches.Raw()))

	return Result(ret)
}

// MergeValidationCachesEXT wraps vkMergeValidationCachesEXT.
func MergeValidationCachesEXT(device Device, dstCache ValidationCacheEXT, srcCacheCount uint32, pSrcCaches ffi.Ref[ValidationCacheEXT]) Result {
	ret := C.vkMergeValidationCachesEXT(C.VkDevice(unsafe.Pointer(device)), C.VkValidationCacheEXT(dstCache), C.uint32_t(srcCacheCount), (*C.VkValidationCacheEXT)(pSrcCaches.Raw()))

	return Result(ret)
}

// QueueBeginDebugUtilsLabelEXT wraps vkQueueBeginDebugUtilsLabelEXT.
func QueueBeginDebugUtilsLabelEXT(queue Queue, pLabelInfo DebugUtilsLabelEXT) {
	C.vkQueueBeginDebugUtilsLabelEXT(C.VkQueue(unsafe.Pointer(queue)), (*C.VkDebugUtilsLabelEXT)(pLabelInfo.Raw()))
}

// QueueBindSparse wraps vkQueueBindSparse.
func QueueBindSparse(queue Queue, bindInfoCount uint32, pBindInfo BindSparseInfo, fence Fence) Result {
	ret := C.vkQueueBindSparse(C.VkQueue(unsafe.Pointer(queue)), C.uint32_t(bindInfoCount), (*C.VkBindSparseInfo)(pBindInfo.Raw()), C.VkFence(fence))

	return Result(ret)
}

// QueueEndDebugUtilsLabelEXT wraps vkQueueEndDebugUtilsLabelEXT.
func QueueEndDebugUtilsLabelEXT(queue Queue) {
	C.vkQueueEndDebugUtilsLabelEXT(C.VkQueue(unsafe.Pointer(queue)))
}

// QueueInsertDebugUtilsLabelEXT wraps vkQueueInsertDebugUtilsLabelEXT.
func QueueInsertDebugUtilsLabelEXT(queue Queue, pLabelInfo DebugUtilsLabelEXT) {
	C.vkQueueInsertDebugUtilsLabelEXT(C.VkQueue(unsafe.Pointer(queue)), (*C.VkDebugUtilsLabelEXT)(pLabelInfo.Raw()))
}

// QueueNotifyOutOfBandNV wraps vkQueueNotifyOutOfBandNV.
func QueueNotifyOutOfBandNV(queue Queue, pQueueTypeInfo OutOfBandQueueTypeInfoNV) {
	C.vkQueueNotifyOutOfBandNV(C.VkQueue(unsafe.Pointer(queue)), (*C.VkOutOfBandQueueTypeInfoNV)(pQueueTypeInfo.Raw()))
}

// QueuePresentKHR wraps vkQueuePresentKHR.
func QueuePresentKHR(queue Queue, pPresentInfo PresentInfoKHR) Result {
	ret := C.vkQueuePresentKHR(C.VkQueue(unsafe.Pointer(queue)), (*C.VkPresentInfoKHR)(pPresentInfo.Raw()))

	return Result(ret)
}

// QueueSetPerformanceConfigurationINTEL wraps vkQueueSetPerformanceConfigurationINTEL.
func QueueSetPerformanceConfigurationINTEL(queue Queue, configuration PerformanceConfigurationINTEL) Result {
	ret := C.vkQueueSetPerformanceConfigurationINTEL(C.VkQueue(unsafe.Pointer(queue)), C.VkPerformanceConfigurationINTEL(configuration))

	return Result(ret)
}

// QueueSubmit wraps vkQueueSubmit.
func QueueSubmit(queue Queue, submitCount uint32, pSubmits SubmitInfo, fence Fence) Result {
	ret := C.vkQueueSubmit(C.VkQueue(unsafe.Pointer(queue)), C.uint32_t(submitCount), (*C.VkSubmitInfo)(pSubmits.Raw()), C.VkFence(fence))

	return Result(ret)
}

// QueueSubmit2 wraps vkQueueSubmit2.
func QueueSubmit2(queue Queue, submitCount uint32, pSubmits SubmitInfo2, fence Fence) Result {
	ret := C.vkQueueSubmit2(C.VkQueue(unsafe.Pointer(queue)), C.uint32_t(submitCount), (*C.VkSubmitInfo2)(pSubmits.Raw()), C.VkFence(fence))

	return Result(ret)
}

// QueueWaitIdle wraps vkQueueWaitIdle.
func QueueWaitIdle(queue Queue) Result {
	ret := C.vkQueueWaitIdle(C.VkQueue(unsafe.Pointer(queue)))

	return Result(ret)
}

// RegisterDeviceEventEXT wraps vkRegisterDeviceEventEXT.
func RegisterDeviceEventEXT(device Device, pDeviceEventInfo DeviceEventInfoEXT, pAllocator AllocationCallbacks, pFence ffi.Ref[Fence]) Result {
	ret := C.vkRegisterDeviceEventEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDeviceEventInfoEXT)(pDeviceEventInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkFence)(pFence.Raw()))

	return Result(ret)
}

// RegisterDisplayEventEXT wraps vkRegisterDisplayEventEXT.
func RegisterDisplayEventEXT(device Device, display DisplayKHR, pDisplayEventInfo DisplayEventInfoEXT, pAllocator AllocationCallbacks, pFence ffi.Ref[Fence]) Result {
	ret := C.vkRegisterDisplayEventEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDisplayKHR(display), (*C.VkDisplayEventInfoEXT)(pDisplayEventInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()), (*C.VkFence)(pFence.Raw()))

	return Result(ret)
}

// ReleaseCapturedPipelineDataKHR wraps vkReleaseCapturedPipelineDataKHR.
func ReleaseCapturedPipelineDataKHR(device Device, pInfo ReleaseCapturedPipelineDataInfoKHR, pAllocator AllocationCallbacks) Result {
	ret := C.vkReleaseCapturedPipelineDataKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkReleaseCapturedPipelineDataInfoKHR)(pInfo.Raw()), (*C.VkAllocationCallbacks)(pAllocator.Raw()))

	return Result(ret)
}

// ReleaseDisplayEXT wraps vkReleaseDisplayEXT.
func ReleaseDisplayEXT(physicalDevice PhysicalDevice, display DisplayKHR) Result {
	ret := C.vkReleaseDisplayEXT(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkDisplayKHR(display))

	return Result(ret)
}

// ReleaseFullScreenExclusiveModeEXT wraps vkReleaseFullScreenExclusiveModeEXT.
func ReleaseFullScreenExclusiveModeEXT(device Device, swapchain SwapchainKHR) Result {
	ret := C.vkReleaseFullScreenExclusiveModeEXT(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain))

	return Result(ret)
}

// ReleasePerformanceConfigurationINTEL wraps vkReleasePerformanceConfigurationINTEL.
func ReleasePerformanceConfigurationINTEL(device Device, configuration PerformanceConfigurationINTEL) Result {
	ret := C.vkReleasePerformanceConfigurationINTEL(C.VkDevice(unsafe.Pointer(device)), C.VkPerformanceConfigurationINTEL(configuration))

	return Result(ret)
}

// ReleaseProfilingLockKHR wraps vkReleaseProfilingLockKHR.
func ReleaseProfilingLockKHR(device Device) {
	C.vkReleaseProfilingLockKHR(C.VkDevice(unsafe.Pointer(device)))
}

// ReleaseSwapchainImagesKHR wraps vkReleaseSwapchainImagesKHR.
func ReleaseSwapchainImagesKHR(device Device, pReleaseInfo ReleaseSwapchainImagesInfoKHR) Result {
	ret := C.vkReleaseSwapchainImagesKHR(C.VkDevice(unsafe.Pointer(device)), (*C.VkReleaseSwapchainImagesInfoKHR)(pReleaseInfo.Raw()))

	return Result(ret)
}

// ResetCommandBuffer wraps vkResetCommandBuffer.
func ResetCommandBuffer(commandBuffer CommandBuffer, flags CommandBufferResetFlags) Result {
	ret := C.vkResetCommandBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCommandBufferResetFlags(flags))

	return Result(ret)
}

// ResetCommandPool wraps vkResetCommandPool.
func ResetCommandPool(device Device, commandPool CommandPool, flags CommandPoolResetFlags) Result {
	ret := C.vkResetCommandPool(C.VkDevice(unsafe.Pointer(device)), C.VkCommandPool(commandPool), C.VkCommandPoolResetFlags(flags))

	return Result(ret)
}

// ResetDescriptorPool wraps vkResetDescriptorPool.
func ResetDescriptorPool(device Device, descriptorPool DescriptorPool, flags DescriptorPoolResetFlags) Result {
	ret := C.vkResetDescriptorPool(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorPool(descriptorPool), C.VkDescriptorPoolResetFlags(flags))

	return Result(ret)
}

// ResetEvent wraps vkResetEvent.
func ResetEvent(device Device, event Event) Result {
	ret := C.vkResetEvent(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event))

	return Result(ret)
}

// ResetFences wraps vkResetFences.
func ResetFences(device Device, fenceCount uint32, pFences ffi.Ref[Fence]) Result {
	ret := C.vkResetFences(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(fenceCount), (*C.VkFence)(pFences.Raw()))

	return Result(ret)
}

// ResetQueryPool wraps vkResetQueryPool.
func ResetQueryPool(device Device, queryPool QueryPool, firstQuery uint32, queryCount uint32) {
	C.vkResetQueryPool(C.VkDevice(unsafe.Pointer(device)), C.VkQueryPool(queryPool), C.uint32_t(firstQuery), C.uint32_t(queryCount))
}

// SetDebugUtilsObjectNameEXT wraps vkSetDebugUtilsObjectNameEXT.
func SetDebugUtilsObjectNameEXT(device Device, pNameInfo DebugUtilsObjectNameInfoEXT) Result {
	ret := C.vkSetDebugUtilsObjectNameEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDebugUtilsObjectNameInfoEXT)(pNameInfo.Raw()))

	return Result(ret)
}

// SetDebugUtilsObjectTagEXT wraps vkSetDebugUtilsObjectTagEXT.
func SetDebugUtilsObjectTagEXT(device Device, pTagInfo DebugUtilsObjectTagInfoEXT) Result {
	ret := C.vkSetDebugUtilsObjectTagEXT(C.VkDevice(unsafe.Pointer(device)), (*C.VkDebugUtilsObjectTagInfoEXT)(pTagInfo.Raw()))

	return Result(ret)
}

// SetDeviceMemoryPriorityEXT wraps vkSetDeviceMemoryPriorityEXT.
func SetDeviceMemoryPriorityEXT(device Device, memory DeviceMemory, priority float32) {
	C.vkSetDeviceMemoryPriorityEXT(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory), C.float(priority))
}

// SetEvent wraps vkSetEvent.
func SetEvent(device Device, event Event) Result {
	ret := C.vkSetEvent(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event))

	return Result(ret)
}

// SetHdrMetadataEXT wraps vkSetHdrMetadataEXT.
func SetHdrMetadataEXT(device Device, swapchainCount uint32, pSwapchains ffi.Ref[SwapchainKHR], pMetadata HdrMetadataEXT) {
	C.vkSetHdrMetadataEXT(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(swapchainCount), (*C.VkSwapchainKHR)(pSwapchains.Raw()), (*C.VkHdrMetadataEXT)(pMetadata.Raw()))
}

// SetLatencyMarkerNV wraps vkSetLatencyMarkerNV.
func SetLatencyMarkerNV(device Device, swapchain SwapchainKHR, pLatencyMarkerInfo SetLatencyMarkerInfoNV) {
	C.vkSetLatencyMarkerNV(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkSetLatencyMarkerInfoNV)(pLatencyMarkerInfo.Raw()))
}

// SetLatencySleepModeNV wraps vkSetLatencySleepModeNV.
func SetLatencySleepModeNV(device Device, swapchain SwapchainKHR, pSleepModeInfo LatencySleepModeInfoNV) Result {
	ret := C.vkSetLatencySleepModeNV(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkLatencySleepModeInfoNV)(pSleepModeInfo.Raw()))

	return Result(ret)
}

// SetLocalDimmingAMD wraps vkSetLocalDimmingAMD.
func SetLocalDimmingAMD(device Device, swapChain SwapchainKHR, localDimmingEnable bool) {
	tmp_localDimmingEnable := 0

	if localDimmingEnable {
		tmp_localDimmingEnable = 1
	}

	C.vkSetLocalDimmingAMD(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapChain), C.VkBool32(tmp_localDimmingEnable))
}

// SetPrivateData wraps vkSetPrivateData.
func SetPrivateData(device Device, objectType ObjectType, objectHandle uint64, privateDataSlot PrivateDataSlot, data uint64) Result {
	ret := C.vkSetPrivateData(C.VkDevice(unsafe.Pointer(device)), C.VkObjectType(objectType), C.uint64_t(objectHandle), C.VkPrivateDataSlot(privateDataSlot), C.uint64_t(data))

	return Result(ret)
}

// SignalSemaphore wraps vkSignalSemaphore.
func SignalSemaphore(device Device, pSignalInfo SemaphoreSignalInfo) Result {
	ret := C.vkSignalSemaphore(C.VkDevice(unsafe.Pointer(device)), (*C.VkSemaphoreSignalInfo)(pSignalInfo.Raw()))

	return Result(ret)
}

// SubmitDebugUtilsMessageEXT wraps vkSubmitDebugUtilsMessageEXT.
func SubmitDebugUtilsMessageEXT(instance Instance, messageSeverity DebugUtilsMessageSeverityFlagsEXT, messageTypes DebugUtilsMessageTypeFlagsEXT, pCallbackData DebugUtilsMessengerCallbackDataEXT) {
	C.vkSubmitDebugUtilsMessageEXT(C.VkInstance(unsafe.Pointer(instance)), C.VkDebugUtilsMessageSeverityFlagBitsEXT(messageSeverity), C.VkDebugUtilsMessageTypeFlagsEXT(messageTypes), (*C.VkDebugUtilsMessengerCallbackDataEXT)(pCallbackData.Raw()))
}

// TransitionImageLayout wraps vkTransitionImageLayout.
func TransitionImageLayout(device Device, transitionCount uint32, pTransitions HostImageLayoutTransitionInfo) Result {
	ret := C.vkTransitionImageLayout(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(transitionCount), (*C.VkHostImageLayoutTransitionInfo)(pTransitions.Raw()))

	return Result(ret)
}

// TrimCommandPool wraps vkTrimCommandPool.
func TrimCommandPool(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
	C.vkTrimCommandPool(C.VkDevice(unsafe.Pointer(device)), C.VkCommandPool(commandPool), C.VkCommandPoolTrimFlags(flags))
}

// UninitializePerformanceApiINTEL wraps vkUninitializePerformanceApiINTEL.
func UninitializePerformanceApiINTEL(device Device) {
	C.vkUninitializePerformanceApiINTEL(C.VkDevice(unsafe.Pointer(device)))
}

// UnmapMemory wraps vkUnmapMemory.
func UnmapMemory(device Device, memory DeviceMemory) {
	C.vkUnmapMemory(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory))
}

// UnmapMemory2 wraps vkUnmapMemory2.
func UnmapMemory2(device Device, pMemoryUnmapInfo MemoryUnmapInfo) Result {
	ret := C.vkUnmapMemory2(C.VkDevice(unsafe.Pointer(device)), (*C.VkMemoryUnmapInfo)(pMemoryUnmapInfo.Raw()))

	return Result(ret)
}

// UpdateDescriptorSetWithTemplate wraps vkUpdateDescriptorSetWithTemplate.
func UpdateDescriptorSetWithTemplate(device Device, descriptorSet DescriptorSet, descriptorUpdateTemplate DescriptorUpdateTemplate, pData unsafe.Pointer) {
	C.vkUpdateDescriptorSetWithTemplate(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSet(descriptorSet), C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate), pData)
}

// UpdateDescriptorSets wraps vkUpdateDescriptorSets.
func UpdateDescriptorSets(device Device, descriptorWriteCount uint32, pDescriptorWrites WriteDescriptorSet, descriptorCopyCount uint32, pDescriptorCopies CopyDescriptorSet) {
	C.vkUpdateDescriptorSets(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(descriptorWriteCount), (*C.VkWriteDescriptorSet)(pDescriptorWrites.Raw()), C.uint32_t(descriptorCopyCount), (*C.VkCopyDescriptorSet)(pDescriptorCopies.Raw()))
}

// UpdateIndirectExecutionSetPipelineEXT wraps vkUpdateIndirectExecutionSetPipelineEXT.
func UpdateIndirectExecutionSetPipelineEXT(device Device, indirectExecutionSet IndirectExecutionSetEXT, executionSetWriteCount uint32, pExecutionSetWrites WriteIndirectExecutionSetPipelineEXT) {
	C.vkUpdateIndirectExecutionSetPipelineEXT(C.VkDevice(unsafe.Pointer(device)), C.VkIndirectExecutionSetEXT(indirectExecutionSet), C.uint32_t(executionSetWriteCount), (*C.VkWriteIndirectExecutionSetPipelineEXT)(pExecutionSetWrites.Raw()))
}

// UpdateIndirectExecutionSetShaderEXT wraps vkUpdateIndirectExecutionSetShaderEXT.
func UpdateIndirectExecutionSetShaderEXT(device Device, indirectExecutionSet IndirectExecutionSetEXT, executionSetWriteCount uint32, pExecutionSetWrites WriteIndirectExecutionSetShaderEXT) {
	C.vkUpdateIndirectExecutionSetShaderEXT(C.VkDevice(unsafe.Pointer(device)), C.VkIndirectExecutionSetEXT(indirectExecutionSet), C.uint32_t(executionSetWriteCount), (*C.VkWriteIndirectExecutionSetShaderEXT)(pExecutionSetWrites.Raw()))
}

// UpdateVideoSessionParametersKHR wraps vkUpdateVideoSessionParametersKHR.
func UpdateVideoSessionParametersKHR(device Device, videoSessionParameters VideoSessionParametersKHR, pUpdateInfo VideoSessionParametersUpdateInfoKHR) Result {
	ret := C.vkUpdateVideoSessionParametersKHR(C.VkDevice(unsafe.Pointer(device)), C.VkVideoSessionParametersKHR(videoSessionParameters), (*C.VkVideoSessionParametersUpdateInfoKHR)(pUpdateInfo.Raw()))

	return Result(ret)
}

// WaitForFences wraps vkWaitForFences.
func WaitForFences(device Device, fenceCount uint32, pFences ffi.Ref[Fence], waitAll bool, timeout uint64) Result {
	tmp_waitAll := 0

	if waitAll {
		tmp_waitAll = 1
	}

	ret := C.vkWaitForFences(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(fenceCount), (*C.VkFence)(pFences.Raw()), C.VkBool32(tmp_waitAll), C.uint64_t(timeout))

	return Result(ret)
}

// WaitForPresent2KHR wraps vkWaitForPresent2KHR.
func WaitForPresent2KHR(device Device, swapchain SwapchainKHR, pPresentWait2Info PresentWait2InfoKHR) Result {
	ret := C.vkWaitForPresent2KHR(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), (*C.VkPresentWait2InfoKHR)(pPresentWait2Info.Raw()))

	return Result(ret)
}

// WaitForPresentKHR wraps vkWaitForPresentKHR.
func WaitForPresentKHR(device Device, swapchain SwapchainKHR, presentId uint64, timeout uint64) Result {
	ret := C.vkWaitForPresentKHR(C.VkDevice(unsafe.Pointer(device)), C.VkSwapchainKHR(swapchain), C.uint64_t(presentId), C.uint64_t(timeout))

	return Result(ret)
}

// WaitSemaphores wraps vkWaitSemaphores.
func WaitSemaphores(device Device, pWaitInfo SemaphoreWaitInfo, timeout uint64) Result {
	ret := C.vkWaitSemaphores(C.VkDevice(unsafe.Pointer(device)), (*C.VkSemaphoreWaitInfo)(pWaitInfo.Raw()), C.uint64_t(timeout))

	return Result(ret)
}

// WriteAccelerationStructuresPropertiesKHR wraps vkWriteAccelerationStructuresPropertiesKHR.
func WriteAccelerationStructuresPropertiesKHR(device Device, accelerationStructureCount uint32, pAccelerationStructures ffi.Ref[AccelerationStructureKHR], queryType QueryType, dataSize uintptr, pData unsafe.Pointer, stride uintptr) Result {
	ret := C.vkWriteAccelerationStructuresPropertiesKHR(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(accelerationStructureCount), (*C.VkAccelerationStructureKHR)(pAccelerationStructures.Raw()), C.VkQueryType(queryType), C.size_t(dataSize), pData, C.size_t(stride))

	return Result(ret)
}

// WriteMicromapsPropertiesEXT wraps vkWriteMicromapsPropertiesEXT.
func WriteMicromapsPropertiesEXT(device Device, micromapCount uint32, pMicromaps ffi.Ref[MicromapEXT], queryType QueryType, dataSize uintptr, pData unsafe.Pointer, stride uintptr) Result {
	ret := C.vkWriteMicromapsPropertiesEXT(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(micromapCount), (*C.VkMicromapEXT)(pMicromaps.Raw()), C.VkQueryType(queryType), C.size_t(dataSize), pData, C.size_t(stride))

	return Result(ret)
}

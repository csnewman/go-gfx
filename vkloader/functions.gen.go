package vkloader

import (
	vk "github.com/csnewman/go-gfx/vk"
	"unsafe"
)

/*
#define VK_NO_PROTOTYPES 1
#include "../vk/vulkan.h"
*/
/*
PFN_vkAcquireDrmDisplayEXT ptr_vkAcquireDrmDisplayEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquireDrmDisplayEXT(VkPhysicalDevice physicalDevice, int32_t drmFd, VkDisplayKHR display) {
	return ptr_vkAcquireDrmDisplayEXT(physicalDevice, drmFd, display);
}
*/
/*
PFN_vkAcquireFullScreenExclusiveModeEXT ptr_vkAcquireFullScreenExclusiveModeEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquireFullScreenExclusiveModeEXT(VkDevice device, VkSwapchainKHR swapchain) {
	return ptr_vkAcquireFullScreenExclusiveModeEXT(device, swapchain);
}
*/
/*
PFN_vkAcquireNextImage2KHR ptr_vkAcquireNextImage2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquireNextImage2KHR(VkDevice device, VkAcquireNextImageInfoKHR* pAcquireInfo, uint32_t* pImageIndex) {
	return ptr_vkAcquireNextImage2KHR(device, pAcquireInfo, pImageIndex);
}
*/
/*
PFN_vkAcquireNextImageKHR ptr_vkAcquireNextImageKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquireNextImageKHR(VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex) {
	return ptr_vkAcquireNextImageKHR(device, swapchain, timeout, semaphore, fence, pImageIndex);
}
*/
/*
PFN_vkAcquirePerformanceConfigurationINTEL ptr_vkAcquirePerformanceConfigurationINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquirePerformanceConfigurationINTEL(VkDevice device, VkPerformanceConfigurationAcquireInfoINTEL* pAcquireInfo, VkPerformanceConfigurationINTEL* pConfiguration) {
	return ptr_vkAcquirePerformanceConfigurationINTEL(device, pAcquireInfo, pConfiguration);
}
*/
/*
PFN_vkAcquireProfilingLockKHR ptr_vkAcquireProfilingLockKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquireProfilingLockKHR(VkDevice device, VkAcquireProfilingLockInfoKHR* pInfo) {
	return ptr_vkAcquireProfilingLockKHR(device, pInfo);
}
*/
/*
PFN_vkAcquireWinrtDisplayNV ptr_vkAcquireWinrtDisplayNV;
VKAPI_ATTR VkResult VKAPI_CALL vkAcquireWinrtDisplayNV(VkPhysicalDevice physicalDevice, VkDisplayKHR display) {
	return ptr_vkAcquireWinrtDisplayNV(physicalDevice, display);
}
*/
/*
PFN_vkAllocateCommandBuffers ptr_vkAllocateCommandBuffers;
VKAPI_ATTR VkResult VKAPI_CALL vkAllocateCommandBuffers(VkDevice device, VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers) {
	return ptr_vkAllocateCommandBuffers(device, pAllocateInfo, pCommandBuffers);
}
*/
/*
PFN_vkAllocateDescriptorSets ptr_vkAllocateDescriptorSets;
VKAPI_ATTR VkResult VKAPI_CALL vkAllocateDescriptorSets(VkDevice device, VkDescriptorSetAllocateInfo* pAllocateInfo, VkDescriptorSet* pDescriptorSets) {
	return ptr_vkAllocateDescriptorSets(device, pAllocateInfo, pDescriptorSets);
}
*/
/*
PFN_vkAllocateMemory ptr_vkAllocateMemory;
VKAPI_ATTR VkResult VKAPI_CALL vkAllocateMemory(VkDevice device, VkMemoryAllocateInfo* pAllocateInfo, VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory) {
	return ptr_vkAllocateMemory(device, pAllocateInfo, pAllocator, pMemory);
}
*/
/*
PFN_vkAntiLagUpdateAMD ptr_vkAntiLagUpdateAMD;
VKAPI_ATTR void VKAPI_CALL vkAntiLagUpdateAMD(VkDevice device, VkAntiLagDataAMD* pData) {
	ptr_vkAntiLagUpdateAMD(device, pData);
}
*/
/*
PFN_vkBeginCommandBuffer ptr_vkBeginCommandBuffer;
VKAPI_ATTR VkResult VKAPI_CALL vkBeginCommandBuffer(VkCommandBuffer commandBuffer, VkCommandBufferBeginInfo* pBeginInfo) {
	return ptr_vkBeginCommandBuffer(commandBuffer, pBeginInfo);
}
*/
/*
PFN_vkBindAccelerationStructureMemoryNV ptr_vkBindAccelerationStructureMemoryNV;
VKAPI_ATTR VkResult VKAPI_CALL vkBindAccelerationStructureMemoryNV(VkDevice device, uint32_t bindInfoCount, VkBindAccelerationStructureMemoryInfoNV* pBindInfos) {
	return ptr_vkBindAccelerationStructureMemoryNV(device, bindInfoCount, pBindInfos);
}
*/
/*
PFN_vkBindBufferMemory ptr_vkBindBufferMemory;
VKAPI_ATTR VkResult VKAPI_CALL vkBindBufferMemory(VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return ptr_vkBindBufferMemory(device, buffer, memory, memoryOffset);
}
*/
/*
PFN_vkBindBufferMemory2 ptr_vkBindBufferMemory2;
VKAPI_ATTR VkResult VKAPI_CALL vkBindBufferMemory2(VkDevice device, uint32_t bindInfoCount, VkBindBufferMemoryInfo* pBindInfos) {
	return ptr_vkBindBufferMemory2(device, bindInfoCount, pBindInfos);
}
*/
/*
PFN_vkBindDataGraphPipelineSessionMemoryARM ptr_vkBindDataGraphPipelineSessionMemoryARM;
VKAPI_ATTR VkResult VKAPI_CALL vkBindDataGraphPipelineSessionMemoryARM(VkDevice device, uint32_t bindInfoCount, VkBindDataGraphPipelineSessionMemoryInfoARM* pBindInfos) {
	return ptr_vkBindDataGraphPipelineSessionMemoryARM(device, bindInfoCount, pBindInfos);
}
*/
/*
PFN_vkBindImageMemory ptr_vkBindImageMemory;
VKAPI_ATTR VkResult VKAPI_CALL vkBindImageMemory(VkDevice device, VkImage image, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
	return ptr_vkBindImageMemory(device, image, memory, memoryOffset);
}
*/
/*
PFN_vkBindImageMemory2 ptr_vkBindImageMemory2;
VKAPI_ATTR VkResult VKAPI_CALL vkBindImageMemory2(VkDevice device, uint32_t bindInfoCount, VkBindImageMemoryInfo* pBindInfos) {
	return ptr_vkBindImageMemory2(device, bindInfoCount, pBindInfos);
}
*/
/*
PFN_vkBindOpticalFlowSessionImageNV ptr_vkBindOpticalFlowSessionImageNV;
VKAPI_ATTR VkResult VKAPI_CALL vkBindOpticalFlowSessionImageNV(VkDevice device, VkOpticalFlowSessionNV session, VkOpticalFlowSessionBindingPointNV bindingPoint, VkImageView view, VkImageLayout layout) {
	return ptr_vkBindOpticalFlowSessionImageNV(device, session, bindingPoint, view, layout);
}
*/
/*
PFN_vkBindTensorMemoryARM ptr_vkBindTensorMemoryARM;
VKAPI_ATTR VkResult VKAPI_CALL vkBindTensorMemoryARM(VkDevice device, uint32_t bindInfoCount, VkBindTensorMemoryInfoARM* pBindInfos) {
	return ptr_vkBindTensorMemoryARM(device, bindInfoCount, pBindInfos);
}
*/
/*
PFN_vkBindVideoSessionMemoryKHR ptr_vkBindVideoSessionMemoryKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkBindVideoSessionMemoryKHR(VkDevice device, VkVideoSessionKHR videoSession, uint32_t bindSessionMemoryInfoCount, VkBindVideoSessionMemoryInfoKHR* pBindSessionMemoryInfos) {
	return ptr_vkBindVideoSessionMemoryKHR(device, videoSession, bindSessionMemoryInfoCount, pBindSessionMemoryInfos);
}
*/
/*
PFN_vkBuildAccelerationStructuresKHR ptr_vkBuildAccelerationStructuresKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkBuildAccelerationStructuresKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, uint32_t infoCount, VkAccelerationStructureBuildGeometryInfoKHR* pInfos, VkAccelerationStructureBuildRangeInfoKHR** ppBuildRangeInfos) {
	return ptr_vkBuildAccelerationStructuresKHR(device, deferredOperation, infoCount, pInfos, ppBuildRangeInfos);
}
*/
/*
PFN_vkBuildMicromapsEXT ptr_vkBuildMicromapsEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkBuildMicromapsEXT(VkDevice device, VkDeferredOperationKHR deferredOperation, uint32_t infoCount, VkMicromapBuildInfoEXT* pInfos) {
	return ptr_vkBuildMicromapsEXT(device, deferredOperation, infoCount, pInfos);
}
*/
/*
PFN_vkCmdBeginConditionalRenderingEXT ptr_vkCmdBeginConditionalRenderingEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginConditionalRenderingEXT(VkCommandBuffer commandBuffer, VkConditionalRenderingBeginInfoEXT* pConditionalRenderingBegin) {
	ptr_vkCmdBeginConditionalRenderingEXT(commandBuffer, pConditionalRenderingBegin);
}
*/
/*
PFN_vkCmdBeginDebugUtilsLabelEXT ptr_vkCmdBeginDebugUtilsLabelEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginDebugUtilsLabelEXT(VkCommandBuffer commandBuffer, VkDebugUtilsLabelEXT* pLabelInfo) {
	ptr_vkCmdBeginDebugUtilsLabelEXT(commandBuffer, pLabelInfo);
}
*/
/*
PFN_vkCmdBeginPerTileExecutionQCOM ptr_vkCmdBeginPerTileExecutionQCOM;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginPerTileExecutionQCOM(VkCommandBuffer commandBuffer, VkPerTileBeginInfoQCOM* pPerTileBeginInfo) {
	ptr_vkCmdBeginPerTileExecutionQCOM(commandBuffer, pPerTileBeginInfo);
}
*/
/*
PFN_vkCmdBeginQuery ptr_vkCmdBeginQuery;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags) {
	ptr_vkCmdBeginQuery(commandBuffer, queryPool, query, flags);
}
*/
/*
PFN_vkCmdBeginQueryIndexedEXT ptr_vkCmdBeginQueryIndexedEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginQueryIndexedEXT(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags, uint32_t index) {
	ptr_vkCmdBeginQueryIndexedEXT(commandBuffer, queryPool, query, flags, index);
}
*/
/*
PFN_vkCmdBeginRenderPass ptr_vkCmdBeginRenderPass;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginRenderPass(VkCommandBuffer commandBuffer, VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
	ptr_vkCmdBeginRenderPass(commandBuffer, pRenderPassBegin, contents);
}
*/
/*
PFN_vkCmdBeginRenderPass2 ptr_vkCmdBeginRenderPass2;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginRenderPass2(VkCommandBuffer commandBuffer, VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassBeginInfo* pSubpassBeginInfo) {
	ptr_vkCmdBeginRenderPass2(commandBuffer, pRenderPassBegin, pSubpassBeginInfo);
}
*/
/*
PFN_vkCmdBeginRendering ptr_vkCmdBeginRendering;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginRendering(VkCommandBuffer commandBuffer, VkRenderingInfo* pRenderingInfo) {
	ptr_vkCmdBeginRendering(commandBuffer, pRenderingInfo);
}
*/
/*
PFN_vkCmdBeginTransformFeedbackEXT ptr_vkCmdBeginTransformFeedbackEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginTransformFeedbackEXT(VkCommandBuffer commandBuffer, uint32_t firstCounterBuffer, uint32_t counterBufferCount, VkBuffer* pCounterBuffers, VkDeviceSize* pCounterBufferOffsets) {
	ptr_vkCmdBeginTransformFeedbackEXT(commandBuffer, firstCounterBuffer, counterBufferCount, pCounterBuffers, pCounterBufferOffsets);
}
*/
/*
PFN_vkCmdBeginVideoCodingKHR ptr_vkCmdBeginVideoCodingKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginVideoCodingKHR(VkCommandBuffer commandBuffer, VkVideoBeginCodingInfoKHR* pBeginInfo) {
	ptr_vkCmdBeginVideoCodingKHR(commandBuffer, pBeginInfo);
}
*/
/*
PFN_vkCmdBindDescriptorBufferEmbeddedSamplers2EXT ptr_vkCmdBindDescriptorBufferEmbeddedSamplers2EXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(VkCommandBuffer commandBuffer, VkBindDescriptorBufferEmbeddedSamplersInfoEXT* pBindDescriptorBufferEmbeddedSamplersInfo) {
	ptr_vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(commandBuffer, pBindDescriptorBufferEmbeddedSamplersInfo);
}
*/
/*
PFN_vkCmdBindDescriptorBufferEmbeddedSamplersEXT ptr_vkCmdBindDescriptorBufferEmbeddedSamplersEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorBufferEmbeddedSamplersEXT(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t set) {
	ptr_vkCmdBindDescriptorBufferEmbeddedSamplersEXT(commandBuffer, pipelineBindPoint, layout, set);
}
*/
/*
PFN_vkCmdBindDescriptorBuffersEXT ptr_vkCmdBindDescriptorBuffersEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorBuffersEXT(VkCommandBuffer commandBuffer, uint32_t bufferCount, VkDescriptorBufferBindingInfoEXT* pBindingInfos) {
	ptr_vkCmdBindDescriptorBuffersEXT(commandBuffer, bufferCount, pBindingInfos);
}
*/
/*
PFN_vkCmdBindDescriptorSets ptr_vkCmdBindDescriptorSets;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorSets(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t descriptorSetCount, VkDescriptorSet* pDescriptorSets, uint32_t dynamicOffsetCount, uint32_t* pDynamicOffsets) {
	ptr_vkCmdBindDescriptorSets(commandBuffer, pipelineBindPoint, layout, firstSet, descriptorSetCount, pDescriptorSets, dynamicOffsetCount, pDynamicOffsets);
}
*/
/*
PFN_vkCmdBindDescriptorSets2 ptr_vkCmdBindDescriptorSets2;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorSets2(VkCommandBuffer commandBuffer, VkBindDescriptorSetsInfo* pBindDescriptorSetsInfo) {
	ptr_vkCmdBindDescriptorSets2(commandBuffer, pBindDescriptorSetsInfo);
}
*/
/*
PFN_vkCmdBindIndexBuffer ptr_vkCmdBindIndexBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdBindIndexBuffer(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType) {
	ptr_vkCmdBindIndexBuffer(commandBuffer, buffer, offset, indexType);
}
*/
/*
PFN_vkCmdBindIndexBuffer2 ptr_vkCmdBindIndexBuffer2;
VKAPI_ATTR void VKAPI_CALL vkCmdBindIndexBuffer2(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkDeviceSize size, VkIndexType indexType) {
	ptr_vkCmdBindIndexBuffer2(commandBuffer, buffer, offset, size, indexType);
}
*/
/*
PFN_vkCmdBindInvocationMaskHUAWEI ptr_vkCmdBindInvocationMaskHUAWEI;
VKAPI_ATTR void VKAPI_CALL vkCmdBindInvocationMaskHUAWEI(VkCommandBuffer commandBuffer, VkImageView imageView, VkImageLayout imageLayout) {
	ptr_vkCmdBindInvocationMaskHUAWEI(commandBuffer, imageView, imageLayout);
}
*/
/*
PFN_vkCmdBindPipeline ptr_vkCmdBindPipeline;
VKAPI_ATTR void VKAPI_CALL vkCmdBindPipeline(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
	ptr_vkCmdBindPipeline(commandBuffer, pipelineBindPoint, pipeline);
}
*/
/*
PFN_vkCmdBindPipelineShaderGroupNV ptr_vkCmdBindPipelineShaderGroupNV;
VKAPI_ATTR void VKAPI_CALL vkCmdBindPipelineShaderGroupNV(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline, uint32_t groupIndex) {
	ptr_vkCmdBindPipelineShaderGroupNV(commandBuffer, pipelineBindPoint, pipeline, groupIndex);
}
*/
/*
PFN_vkCmdBindShadersEXT ptr_vkCmdBindShadersEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBindShadersEXT(VkCommandBuffer commandBuffer, uint32_t stageCount, VkShaderStageFlagBits* pStages, VkShaderEXT* pShaders) {
	ptr_vkCmdBindShadersEXT(commandBuffer, stageCount, pStages, pShaders);
}
*/
/*
PFN_vkCmdBindShadingRateImageNV ptr_vkCmdBindShadingRateImageNV;
VKAPI_ATTR void VKAPI_CALL vkCmdBindShadingRateImageNV(VkCommandBuffer commandBuffer, VkImageView imageView, VkImageLayout imageLayout) {
	ptr_vkCmdBindShadingRateImageNV(commandBuffer, imageView, imageLayout);
}
*/
/*
PFN_vkCmdBindTileMemoryQCOM ptr_vkCmdBindTileMemoryQCOM;
VKAPI_ATTR void VKAPI_CALL vkCmdBindTileMemoryQCOM(VkCommandBuffer commandBuffer, VkTileMemoryBindInfoQCOM* pTileMemoryBindInfo) {
	ptr_vkCmdBindTileMemoryQCOM(commandBuffer, pTileMemoryBindInfo);
}
*/
/*
PFN_vkCmdBindTransformFeedbackBuffersEXT ptr_vkCmdBindTransformFeedbackBuffersEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBindTransformFeedbackBuffersEXT(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, VkBuffer* pBuffers, VkDeviceSize* pOffsets, VkDeviceSize* pSizes) {
	ptr_vkCmdBindTransformFeedbackBuffersEXT(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets, pSizes);
}
*/
/*
PFN_vkCmdBindVertexBuffers ptr_vkCmdBindVertexBuffers;
VKAPI_ATTR void VKAPI_CALL vkCmdBindVertexBuffers(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, VkBuffer* pBuffers, VkDeviceSize* pOffsets) {
	ptr_vkCmdBindVertexBuffers(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets);
}
*/
/*
PFN_vkCmdBindVertexBuffers2 ptr_vkCmdBindVertexBuffers2;
VKAPI_ATTR void VKAPI_CALL vkCmdBindVertexBuffers2(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, VkBuffer* pBuffers, VkDeviceSize* pOffsets, VkDeviceSize* pSizes, VkDeviceSize* pStrides) {
	ptr_vkCmdBindVertexBuffers2(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets, pSizes, pStrides);
}
*/
/*
PFN_vkCmdBlitImage ptr_vkCmdBlitImage;
VKAPI_ATTR void VKAPI_CALL vkCmdBlitImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageBlit* pRegions, VkFilter filter) {
	ptr_vkCmdBlitImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions, filter);
}
*/
/*
PFN_vkCmdBlitImage2 ptr_vkCmdBlitImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdBlitImage2(VkCommandBuffer commandBuffer, VkBlitImageInfo2* pBlitImageInfo) {
	ptr_vkCmdBlitImage2(commandBuffer, pBlitImageInfo);
}
*/
/*
PFN_vkCmdBuildAccelerationStructureNV ptr_vkCmdBuildAccelerationStructureNV;
VKAPI_ATTR void VKAPI_CALL vkCmdBuildAccelerationStructureNV(VkCommandBuffer commandBuffer, VkAccelerationStructureInfoNV* pInfo, VkBuffer instanceData, VkDeviceSize instanceOffset, VkBool32 update, VkAccelerationStructureNV dst, VkAccelerationStructureNV src, VkBuffer scratch, VkDeviceSize scratchOffset) {
	ptr_vkCmdBuildAccelerationStructureNV(commandBuffer, pInfo, instanceData, instanceOffset, update, dst, src, scratch, scratchOffset);
}
*/
/*
PFN_vkCmdBuildAccelerationStructuresIndirectKHR ptr_vkCmdBuildAccelerationStructuresIndirectKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdBuildAccelerationStructuresIndirectKHR(VkCommandBuffer commandBuffer, uint32_t infoCount, VkAccelerationStructureBuildGeometryInfoKHR* pInfos, VkDeviceAddress* pIndirectDeviceAddresses, uint32_t* pIndirectStrides, uint32_t** ppMaxPrimitiveCounts) {
	ptr_vkCmdBuildAccelerationStructuresIndirectKHR(commandBuffer, infoCount, pInfos, pIndirectDeviceAddresses, pIndirectStrides, ppMaxPrimitiveCounts);
}
*/
/*
PFN_vkCmdBuildAccelerationStructuresKHR ptr_vkCmdBuildAccelerationStructuresKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdBuildAccelerationStructuresKHR(VkCommandBuffer commandBuffer, uint32_t infoCount, VkAccelerationStructureBuildGeometryInfoKHR* pInfos, VkAccelerationStructureBuildRangeInfoKHR** ppBuildRangeInfos) {
	ptr_vkCmdBuildAccelerationStructuresKHR(commandBuffer, infoCount, pInfos, ppBuildRangeInfos);
}
*/
/*
PFN_vkCmdBuildClusterAccelerationStructureIndirectNV ptr_vkCmdBuildClusterAccelerationStructureIndirectNV;
VKAPI_ATTR void VKAPI_CALL vkCmdBuildClusterAccelerationStructureIndirectNV(VkCommandBuffer commandBuffer, VkClusterAccelerationStructureCommandsInfoNV* pCommandInfos) {
	ptr_vkCmdBuildClusterAccelerationStructureIndirectNV(commandBuffer, pCommandInfos);
}
*/
/*
PFN_vkCmdBuildMicromapsEXT ptr_vkCmdBuildMicromapsEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdBuildMicromapsEXT(VkCommandBuffer commandBuffer, uint32_t infoCount, VkMicromapBuildInfoEXT* pInfos) {
	ptr_vkCmdBuildMicromapsEXT(commandBuffer, infoCount, pInfos);
}
*/
/*
PFN_vkCmdBuildPartitionedAccelerationStructuresNV ptr_vkCmdBuildPartitionedAccelerationStructuresNV;
VKAPI_ATTR void VKAPI_CALL vkCmdBuildPartitionedAccelerationStructuresNV(VkCommandBuffer commandBuffer, VkBuildPartitionedAccelerationStructureInfoNV* pBuildInfo) {
	ptr_vkCmdBuildPartitionedAccelerationStructuresNV(commandBuffer, pBuildInfo);
}
*/
/*
PFN_vkCmdClearAttachments ptr_vkCmdClearAttachments;
VKAPI_ATTR void VKAPI_CALL vkCmdClearAttachments(VkCommandBuffer commandBuffer, uint32_t attachmentCount, VkClearAttachment* pAttachments, uint32_t rectCount, VkClearRect* pRects) {
	ptr_vkCmdClearAttachments(commandBuffer, attachmentCount, pAttachments, rectCount, pRects);
}
*/
/*
PFN_vkCmdClearColorImage ptr_vkCmdClearColorImage;
VKAPI_ATTR void VKAPI_CALL vkCmdClearColorImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, VkClearColorValue* pColor, uint32_t rangeCount, VkImageSubresourceRange* pRanges) {
	ptr_vkCmdClearColorImage(commandBuffer, image, imageLayout, pColor, rangeCount, pRanges);
}
*/
/*
PFN_vkCmdClearDepthStencilImage ptr_vkCmdClearDepthStencilImage;
VKAPI_ATTR void VKAPI_CALL vkCmdClearDepthStencilImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, VkImageSubresourceRange* pRanges) {
	ptr_vkCmdClearDepthStencilImage(commandBuffer, image, imageLayout, pDepthStencil, rangeCount, pRanges);
}
*/
/*
PFN_vkCmdControlVideoCodingKHR ptr_vkCmdControlVideoCodingKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdControlVideoCodingKHR(VkCommandBuffer commandBuffer, VkVideoCodingControlInfoKHR* pCodingControlInfo) {
	ptr_vkCmdControlVideoCodingKHR(commandBuffer, pCodingControlInfo);
}
*/
/*
PFN_vkCmdConvertCooperativeVectorMatrixNV ptr_vkCmdConvertCooperativeVectorMatrixNV;
VKAPI_ATTR void VKAPI_CALL vkCmdConvertCooperativeVectorMatrixNV(VkCommandBuffer commandBuffer, uint32_t infoCount, VkConvertCooperativeVectorMatrixInfoNV* pInfos) {
	ptr_vkCmdConvertCooperativeVectorMatrixNV(commandBuffer, infoCount, pInfos);
}
*/
/*
PFN_vkCmdCopyAccelerationStructureKHR ptr_vkCmdCopyAccelerationStructureKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyAccelerationStructureKHR(VkCommandBuffer commandBuffer, VkCopyAccelerationStructureInfoKHR* pInfo) {
	ptr_vkCmdCopyAccelerationStructureKHR(commandBuffer, pInfo);
}
*/
/*
PFN_vkCmdCopyAccelerationStructureNV ptr_vkCmdCopyAccelerationStructureNV;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyAccelerationStructureNV(VkCommandBuffer commandBuffer, VkAccelerationStructureNV dst, VkAccelerationStructureNV src, VkCopyAccelerationStructureModeKHR mode) {
	ptr_vkCmdCopyAccelerationStructureNV(commandBuffer, dst, src, mode);
}
*/
/*
PFN_vkCmdCopyAccelerationStructureToMemoryKHR ptr_vkCmdCopyAccelerationStructureToMemoryKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyAccelerationStructureToMemoryKHR(VkCommandBuffer commandBuffer, VkCopyAccelerationStructureToMemoryInfoKHR* pInfo) {
	ptr_vkCmdCopyAccelerationStructureToMemoryKHR(commandBuffer, pInfo);
}
*/
/*
PFN_vkCmdCopyBuffer ptr_vkCmdCopyBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBuffer(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkBuffer dstBuffer, uint32_t regionCount, VkBufferCopy* pRegions) {
	ptr_vkCmdCopyBuffer(commandBuffer, srcBuffer, dstBuffer, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyBuffer2 ptr_vkCmdCopyBuffer2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBuffer2(VkCommandBuffer commandBuffer, VkCopyBufferInfo2* pCopyBufferInfo) {
	ptr_vkCmdCopyBuffer2(commandBuffer, pCopyBufferInfo);
}
*/
/*
PFN_vkCmdCopyBufferToImage ptr_vkCmdCopyBufferToImage;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBufferToImage(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkBufferImageCopy* pRegions) {
	ptr_vkCmdCopyBufferToImage(commandBuffer, srcBuffer, dstImage, dstImageLayout, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyBufferToImage2 ptr_vkCmdCopyBufferToImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBufferToImage2(VkCommandBuffer commandBuffer, VkCopyBufferToImageInfo2* pCopyBufferToImageInfo) {
	ptr_vkCmdCopyBufferToImage2(commandBuffer, pCopyBufferToImageInfo);
}
*/
/*
PFN_vkCmdCopyImage ptr_vkCmdCopyImage;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageCopy* pRegions) {
	ptr_vkCmdCopyImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyImage2 ptr_vkCmdCopyImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImage2(VkCommandBuffer commandBuffer, VkCopyImageInfo2* pCopyImageInfo) {
	ptr_vkCmdCopyImage2(commandBuffer, pCopyImageInfo);
}
*/
/*
PFN_vkCmdCopyImageToBuffer ptr_vkCmdCopyImageToBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImageToBuffer(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkBuffer dstBuffer, uint32_t regionCount, VkBufferImageCopy* pRegions) {
	ptr_vkCmdCopyImageToBuffer(commandBuffer, srcImage, srcImageLayout, dstBuffer, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyImageToBuffer2 ptr_vkCmdCopyImageToBuffer2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImageToBuffer2(VkCommandBuffer commandBuffer, VkCopyImageToBufferInfo2* pCopyImageToBufferInfo) {
	ptr_vkCmdCopyImageToBuffer2(commandBuffer, pCopyImageToBufferInfo);
}
*/
/*
PFN_vkCmdCopyMemoryIndirectKHR ptr_vkCmdCopyMemoryIndirectKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMemoryIndirectKHR(VkCommandBuffer commandBuffer, VkCopyMemoryIndirectInfoKHR* pCopyMemoryIndirectInfo) {
	ptr_vkCmdCopyMemoryIndirectKHR(commandBuffer, pCopyMemoryIndirectInfo);
}
*/
/*
PFN_vkCmdCopyMemoryIndirectNV ptr_vkCmdCopyMemoryIndirectNV;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMemoryIndirectNV(VkCommandBuffer commandBuffer, VkDeviceAddress copyBufferAddress, uint32_t copyCount, uint32_t stride) {
	ptr_vkCmdCopyMemoryIndirectNV(commandBuffer, copyBufferAddress, copyCount, stride);
}
*/
/*
PFN_vkCmdCopyMemoryToAccelerationStructureKHR ptr_vkCmdCopyMemoryToAccelerationStructureKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMemoryToAccelerationStructureKHR(VkCommandBuffer commandBuffer, VkCopyMemoryToAccelerationStructureInfoKHR* pInfo) {
	ptr_vkCmdCopyMemoryToAccelerationStructureKHR(commandBuffer, pInfo);
}
*/
/*
PFN_vkCmdCopyMemoryToImageIndirectKHR ptr_vkCmdCopyMemoryToImageIndirectKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMemoryToImageIndirectKHR(VkCommandBuffer commandBuffer, VkCopyMemoryToImageIndirectInfoKHR* pCopyMemoryToImageIndirectInfo) {
	ptr_vkCmdCopyMemoryToImageIndirectKHR(commandBuffer, pCopyMemoryToImageIndirectInfo);
}
*/
/*
PFN_vkCmdCopyMemoryToImageIndirectNV ptr_vkCmdCopyMemoryToImageIndirectNV;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMemoryToImageIndirectNV(VkCommandBuffer commandBuffer, VkDeviceAddress copyBufferAddress, uint32_t copyCount, uint32_t stride, VkImage dstImage, VkImageLayout dstImageLayout, VkImageSubresourceLayers* pImageSubresources) {
	ptr_vkCmdCopyMemoryToImageIndirectNV(commandBuffer, copyBufferAddress, copyCount, stride, dstImage, dstImageLayout, pImageSubresources);
}
*/
/*
PFN_vkCmdCopyMemoryToMicromapEXT ptr_vkCmdCopyMemoryToMicromapEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMemoryToMicromapEXT(VkCommandBuffer commandBuffer, VkCopyMemoryToMicromapInfoEXT* pInfo) {
	ptr_vkCmdCopyMemoryToMicromapEXT(commandBuffer, pInfo);
}
*/
/*
PFN_vkCmdCopyMicromapEXT ptr_vkCmdCopyMicromapEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMicromapEXT(VkCommandBuffer commandBuffer, VkCopyMicromapInfoEXT* pInfo) {
	ptr_vkCmdCopyMicromapEXT(commandBuffer, pInfo);
}
*/
/*
PFN_vkCmdCopyMicromapToMemoryEXT ptr_vkCmdCopyMicromapToMemoryEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyMicromapToMemoryEXT(VkCommandBuffer commandBuffer, VkCopyMicromapToMemoryInfoEXT* pInfo) {
	ptr_vkCmdCopyMicromapToMemoryEXT(commandBuffer, pInfo);
}
*/
/*
PFN_vkCmdCopyQueryPoolResults ptr_vkCmdCopyQueryPoolResults;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyQueryPoolResults(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize stride, VkQueryResultFlags flags) {
	ptr_vkCmdCopyQueryPoolResults(commandBuffer, queryPool, firstQuery, queryCount, dstBuffer, dstOffset, stride, flags);
}
*/
/*
PFN_vkCmdCopyTensorARM ptr_vkCmdCopyTensorARM;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyTensorARM(VkCommandBuffer commandBuffer, VkCopyTensorInfoARM* pCopyTensorInfo) {
	ptr_vkCmdCopyTensorARM(commandBuffer, pCopyTensorInfo);
}
*/
/*
PFN_vkCmdCuLaunchKernelNVX ptr_vkCmdCuLaunchKernelNVX;
VKAPI_ATTR void VKAPI_CALL vkCmdCuLaunchKernelNVX(VkCommandBuffer commandBuffer, VkCuLaunchInfoNVX* pLaunchInfo) {
	ptr_vkCmdCuLaunchKernelNVX(commandBuffer, pLaunchInfo);
}
*/
/*
PFN_vkCmdDebugMarkerBeginEXT ptr_vkCmdDebugMarkerBeginEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDebugMarkerBeginEXT(VkCommandBuffer commandBuffer, VkDebugMarkerMarkerInfoEXT* pMarkerInfo) {
	ptr_vkCmdDebugMarkerBeginEXT(commandBuffer, pMarkerInfo);
}
*/
/*
PFN_vkCmdDebugMarkerEndEXT ptr_vkCmdDebugMarkerEndEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDebugMarkerEndEXT(VkCommandBuffer commandBuffer) {
	ptr_vkCmdDebugMarkerEndEXT(commandBuffer);
}
*/
/*
PFN_vkCmdDebugMarkerInsertEXT ptr_vkCmdDebugMarkerInsertEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDebugMarkerInsertEXT(VkCommandBuffer commandBuffer, VkDebugMarkerMarkerInfoEXT* pMarkerInfo) {
	ptr_vkCmdDebugMarkerInsertEXT(commandBuffer, pMarkerInfo);
}
*/
/*
PFN_vkCmdDecodeVideoKHR ptr_vkCmdDecodeVideoKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdDecodeVideoKHR(VkCommandBuffer commandBuffer, VkVideoDecodeInfoKHR* pDecodeInfo) {
	ptr_vkCmdDecodeVideoKHR(commandBuffer, pDecodeInfo);
}
*/
/*
PFN_vkCmdDecompressMemoryIndirectCountNV ptr_vkCmdDecompressMemoryIndirectCountNV;
VKAPI_ATTR void VKAPI_CALL vkCmdDecompressMemoryIndirectCountNV(VkCommandBuffer commandBuffer, VkDeviceAddress indirectCommandsAddress, VkDeviceAddress indirectCommandsCountAddress, uint32_t stride) {
	ptr_vkCmdDecompressMemoryIndirectCountNV(commandBuffer, indirectCommandsAddress, indirectCommandsCountAddress, stride);
}
*/
/*
PFN_vkCmdDecompressMemoryNV ptr_vkCmdDecompressMemoryNV;
VKAPI_ATTR void VKAPI_CALL vkCmdDecompressMemoryNV(VkCommandBuffer commandBuffer, uint32_t decompressRegionCount, VkDecompressMemoryRegionNV* pDecompressMemoryRegions) {
	ptr_vkCmdDecompressMemoryNV(commandBuffer, decompressRegionCount, pDecompressMemoryRegions);
}
*/
/*
PFN_vkCmdDispatch ptr_vkCmdDispatch;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatch(VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	ptr_vkCmdDispatch(commandBuffer, groupCountX, groupCountY, groupCountZ);
}
*/
/*
PFN_vkCmdDispatchBase ptr_vkCmdDispatchBase;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatchBase(VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	ptr_vkCmdDispatchBase(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}
*/
/*
PFN_vkCmdDispatchDataGraphARM ptr_vkCmdDispatchDataGraphARM;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatchDataGraphARM(VkCommandBuffer commandBuffer, VkDataGraphPipelineSessionARM session, VkDataGraphPipelineDispatchInfoARM* pInfo) {
	ptr_vkCmdDispatchDataGraphARM(commandBuffer, session, pInfo);
}
*/
/*
PFN_vkCmdDispatchIndirect ptr_vkCmdDispatchIndirect;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatchIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset) {
	ptr_vkCmdDispatchIndirect(commandBuffer, buffer, offset);
}
*/
/*
PFN_vkCmdDispatchTileQCOM ptr_vkCmdDispatchTileQCOM;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatchTileQCOM(VkCommandBuffer commandBuffer, VkDispatchTileInfoQCOM* pDispatchTileInfo) {
	ptr_vkCmdDispatchTileQCOM(commandBuffer, pDispatchTileInfo);
}
*/
/*
PFN_vkCmdDraw ptr_vkCmdDraw;
VKAPI_ATTR void VKAPI_CALL vkCmdDraw(VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance) {
	ptr_vkCmdDraw(commandBuffer, vertexCount, instanceCount, firstVertex, firstInstance);
}
*/
/*
PFN_vkCmdDrawClusterHUAWEI ptr_vkCmdDrawClusterHUAWEI;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawClusterHUAWEI(VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	ptr_vkCmdDrawClusterHUAWEI(commandBuffer, groupCountX, groupCountY, groupCountZ);
}
*/
/*
PFN_vkCmdDrawClusterIndirectHUAWEI ptr_vkCmdDrawClusterIndirectHUAWEI;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawClusterIndirectHUAWEI(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset) {
	ptr_vkCmdDrawClusterIndirectHUAWEI(commandBuffer, buffer, offset);
}
*/
/*
PFN_vkCmdDrawIndexed ptr_vkCmdDrawIndexed;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndexed(VkCommandBuffer commandBuffer, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t vertexOffset, uint32_t firstInstance) {
	ptr_vkCmdDrawIndexed(commandBuffer, indexCount, instanceCount, firstIndex, vertexOffset, firstInstance);
}
*/
/*
PFN_vkCmdDrawIndexedIndirect ptr_vkCmdDrawIndexedIndirect;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndexedIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	ptr_vkCmdDrawIndexedIndirect(commandBuffer, buffer, offset, drawCount, stride);
}
*/
/*
PFN_vkCmdDrawIndexedIndirectCount ptr_vkCmdDrawIndexedIndirectCount;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndexedIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	ptr_vkCmdDrawIndexedIndirectCount(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
/*
PFN_vkCmdDrawIndirect ptr_vkCmdDrawIndirect;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	ptr_vkCmdDrawIndirect(commandBuffer, buffer, offset, drawCount, stride);
}
*/
/*
PFN_vkCmdDrawIndirectByteCountEXT ptr_vkCmdDrawIndirectByteCountEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndirectByteCountEXT(VkCommandBuffer commandBuffer, uint32_t instanceCount, uint32_t firstInstance, VkBuffer counterBuffer, VkDeviceSize counterBufferOffset, uint32_t counterOffset, uint32_t vertexStride) {
	ptr_vkCmdDrawIndirectByteCountEXT(commandBuffer, instanceCount, firstInstance, counterBuffer, counterBufferOffset, counterOffset, vertexStride);
}
*/
/*
PFN_vkCmdDrawIndirectCount ptr_vkCmdDrawIndirectCount;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	ptr_vkCmdDrawIndirectCount(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
/*
PFN_vkCmdDrawMeshTasksEXT ptr_vkCmdDrawMeshTasksEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMeshTasksEXT(VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	ptr_vkCmdDrawMeshTasksEXT(commandBuffer, groupCountX, groupCountY, groupCountZ);
}
*/
/*
PFN_vkCmdDrawMeshTasksIndirectCountEXT ptr_vkCmdDrawMeshTasksIndirectCountEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMeshTasksIndirectCountEXT(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	ptr_vkCmdDrawMeshTasksIndirectCountEXT(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
/*
PFN_vkCmdDrawMeshTasksIndirectCountNV ptr_vkCmdDrawMeshTasksIndirectCountNV;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMeshTasksIndirectCountNV(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	ptr_vkCmdDrawMeshTasksIndirectCountNV(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
/*
PFN_vkCmdDrawMeshTasksIndirectEXT ptr_vkCmdDrawMeshTasksIndirectEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMeshTasksIndirectEXT(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	ptr_vkCmdDrawMeshTasksIndirectEXT(commandBuffer, buffer, offset, drawCount, stride);
}
*/
/*
PFN_vkCmdDrawMeshTasksIndirectNV ptr_vkCmdDrawMeshTasksIndirectNV;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMeshTasksIndirectNV(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	ptr_vkCmdDrawMeshTasksIndirectNV(commandBuffer, buffer, offset, drawCount, stride);
}
*/
/*
PFN_vkCmdDrawMeshTasksNV ptr_vkCmdDrawMeshTasksNV;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMeshTasksNV(VkCommandBuffer commandBuffer, uint32_t taskCount, uint32_t firstTask) {
	ptr_vkCmdDrawMeshTasksNV(commandBuffer, taskCount, firstTask);
}
*/
/*
PFN_vkCmdDrawMultiEXT ptr_vkCmdDrawMultiEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMultiEXT(VkCommandBuffer commandBuffer, uint32_t drawCount, VkMultiDrawInfoEXT* pVertexInfo, uint32_t instanceCount, uint32_t firstInstance, uint32_t stride) {
	ptr_vkCmdDrawMultiEXT(commandBuffer, drawCount, pVertexInfo, instanceCount, firstInstance, stride);
}
*/
/*
PFN_vkCmdDrawMultiIndexedEXT ptr_vkCmdDrawMultiIndexedEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawMultiIndexedEXT(VkCommandBuffer commandBuffer, uint32_t drawCount, VkMultiDrawIndexedInfoEXT* pIndexInfo, uint32_t instanceCount, uint32_t firstInstance, uint32_t stride, int32_t* pVertexOffset) {
	ptr_vkCmdDrawMultiIndexedEXT(commandBuffer, drawCount, pIndexInfo, instanceCount, firstInstance, stride, pVertexOffset);
}
*/
/*
PFN_vkCmdEncodeVideoKHR ptr_vkCmdEncodeVideoKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdEncodeVideoKHR(VkCommandBuffer commandBuffer, VkVideoEncodeInfoKHR* pEncodeInfo) {
	ptr_vkCmdEncodeVideoKHR(commandBuffer, pEncodeInfo);
}
*/
/*
PFN_vkCmdEndConditionalRenderingEXT ptr_vkCmdEndConditionalRenderingEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdEndConditionalRenderingEXT(VkCommandBuffer commandBuffer) {
	ptr_vkCmdEndConditionalRenderingEXT(commandBuffer);
}
*/
/*
PFN_vkCmdEndDebugUtilsLabelEXT ptr_vkCmdEndDebugUtilsLabelEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdEndDebugUtilsLabelEXT(VkCommandBuffer commandBuffer) {
	ptr_vkCmdEndDebugUtilsLabelEXT(commandBuffer);
}
*/
/*
PFN_vkCmdEndPerTileExecutionQCOM ptr_vkCmdEndPerTileExecutionQCOM;
VKAPI_ATTR void VKAPI_CALL vkCmdEndPerTileExecutionQCOM(VkCommandBuffer commandBuffer, VkPerTileEndInfoQCOM* pPerTileEndInfo) {
	ptr_vkCmdEndPerTileExecutionQCOM(commandBuffer, pPerTileEndInfo);
}
*/
/*
PFN_vkCmdEndQuery ptr_vkCmdEndQuery;
VKAPI_ATTR void VKAPI_CALL vkCmdEndQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query) {
	ptr_vkCmdEndQuery(commandBuffer, queryPool, query);
}
*/
/*
PFN_vkCmdEndQueryIndexedEXT ptr_vkCmdEndQueryIndexedEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdEndQueryIndexedEXT(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, uint32_t index) {
	ptr_vkCmdEndQueryIndexedEXT(commandBuffer, queryPool, query, index);
}
*/
/*
PFN_vkCmdEndRenderPass ptr_vkCmdEndRenderPass;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRenderPass(VkCommandBuffer commandBuffer) {
	ptr_vkCmdEndRenderPass(commandBuffer);
}
*/
/*
PFN_vkCmdEndRenderPass2 ptr_vkCmdEndRenderPass2;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRenderPass2(VkCommandBuffer commandBuffer, VkSubpassEndInfo* pSubpassEndInfo) {
	ptr_vkCmdEndRenderPass2(commandBuffer, pSubpassEndInfo);
}
*/
/*
PFN_vkCmdEndRendering ptr_vkCmdEndRendering;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRendering(VkCommandBuffer commandBuffer) {
	ptr_vkCmdEndRendering(commandBuffer);
}
*/
/*
PFN_vkCmdEndRendering2EXT ptr_vkCmdEndRendering2EXT;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRendering2EXT(VkCommandBuffer commandBuffer, VkRenderingEndInfoEXT* pRenderingEndInfo) {
	ptr_vkCmdEndRendering2EXT(commandBuffer, pRenderingEndInfo);
}
*/
/*
PFN_vkCmdEndTransformFeedbackEXT ptr_vkCmdEndTransformFeedbackEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdEndTransformFeedbackEXT(VkCommandBuffer commandBuffer, uint32_t firstCounterBuffer, uint32_t counterBufferCount, VkBuffer* pCounterBuffers, VkDeviceSize* pCounterBufferOffsets) {
	ptr_vkCmdEndTransformFeedbackEXT(commandBuffer, firstCounterBuffer, counterBufferCount, pCounterBuffers, pCounterBufferOffsets);
}
*/
/*
PFN_vkCmdEndVideoCodingKHR ptr_vkCmdEndVideoCodingKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdEndVideoCodingKHR(VkCommandBuffer commandBuffer, VkVideoEndCodingInfoKHR* pEndCodingInfo) {
	ptr_vkCmdEndVideoCodingKHR(commandBuffer, pEndCodingInfo);
}
*/
/*
PFN_vkCmdExecuteCommands ptr_vkCmdExecuteCommands;
VKAPI_ATTR void VKAPI_CALL vkCmdExecuteCommands(VkCommandBuffer commandBuffer, uint32_t commandBufferCount, VkCommandBuffer* pCommandBuffers) {
	ptr_vkCmdExecuteCommands(commandBuffer, commandBufferCount, pCommandBuffers);
}
*/
/*
PFN_vkCmdExecuteGeneratedCommandsEXT ptr_vkCmdExecuteGeneratedCommandsEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdExecuteGeneratedCommandsEXT(VkCommandBuffer commandBuffer, VkBool32 isPreprocessed, VkGeneratedCommandsInfoEXT* pGeneratedCommandsInfo) {
	ptr_vkCmdExecuteGeneratedCommandsEXT(commandBuffer, isPreprocessed, pGeneratedCommandsInfo);
}
*/
/*
PFN_vkCmdExecuteGeneratedCommandsNV ptr_vkCmdExecuteGeneratedCommandsNV;
VKAPI_ATTR void VKAPI_CALL vkCmdExecuteGeneratedCommandsNV(VkCommandBuffer commandBuffer, VkBool32 isPreprocessed, VkGeneratedCommandsInfoNV* pGeneratedCommandsInfo) {
	ptr_vkCmdExecuteGeneratedCommandsNV(commandBuffer, isPreprocessed, pGeneratedCommandsInfo);
}
*/
/*
PFN_vkCmdFillBuffer ptr_vkCmdFillBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdFillBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data) {
	ptr_vkCmdFillBuffer(commandBuffer, dstBuffer, dstOffset, size, data);
}
*/
/*
PFN_vkCmdInsertDebugUtilsLabelEXT ptr_vkCmdInsertDebugUtilsLabelEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdInsertDebugUtilsLabelEXT(VkCommandBuffer commandBuffer, VkDebugUtilsLabelEXT* pLabelInfo) {
	ptr_vkCmdInsertDebugUtilsLabelEXT(commandBuffer, pLabelInfo);
}
*/
/*
PFN_vkCmdNextSubpass ptr_vkCmdNextSubpass;
VKAPI_ATTR void VKAPI_CALL vkCmdNextSubpass(VkCommandBuffer commandBuffer, VkSubpassContents contents) {
	ptr_vkCmdNextSubpass(commandBuffer, contents);
}
*/
/*
PFN_vkCmdNextSubpass2 ptr_vkCmdNextSubpass2;
VKAPI_ATTR void VKAPI_CALL vkCmdNextSubpass2(VkCommandBuffer commandBuffer, VkSubpassBeginInfo* pSubpassBeginInfo, VkSubpassEndInfo* pSubpassEndInfo) {
	ptr_vkCmdNextSubpass2(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}
*/
/*
PFN_vkCmdOpticalFlowExecuteNV ptr_vkCmdOpticalFlowExecuteNV;
VKAPI_ATTR void VKAPI_CALL vkCmdOpticalFlowExecuteNV(VkCommandBuffer commandBuffer, VkOpticalFlowSessionNV session, VkOpticalFlowExecuteInfoNV* pExecuteInfo) {
	ptr_vkCmdOpticalFlowExecuteNV(commandBuffer, session, pExecuteInfo);
}
*/
/*
PFN_vkCmdPipelineBarrier ptr_vkCmdPipelineBarrier;
VKAPI_ATTR void VKAPI_CALL vkCmdPipelineBarrier(VkCommandBuffer commandBuffer, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, VkDependencyFlags dependencyFlags, uint32_t memoryBarrierCount, VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, VkImageMemoryBarrier* pImageMemoryBarriers) {
	ptr_vkCmdPipelineBarrier(commandBuffer, srcStageMask, dstStageMask, dependencyFlags, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}
*/
/*
PFN_vkCmdPipelineBarrier2 ptr_vkCmdPipelineBarrier2;
VKAPI_ATTR void VKAPI_CALL vkCmdPipelineBarrier2(VkCommandBuffer commandBuffer, VkDependencyInfo* pDependencyInfo) {
	ptr_vkCmdPipelineBarrier2(commandBuffer, pDependencyInfo);
}
*/
/*
PFN_vkCmdPreprocessGeneratedCommandsEXT ptr_vkCmdPreprocessGeneratedCommandsEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdPreprocessGeneratedCommandsEXT(VkCommandBuffer commandBuffer, VkGeneratedCommandsInfoEXT* pGeneratedCommandsInfo, VkCommandBuffer stateCommandBuffer) {
	ptr_vkCmdPreprocessGeneratedCommandsEXT(commandBuffer, pGeneratedCommandsInfo, stateCommandBuffer);
}
*/
/*
PFN_vkCmdPreprocessGeneratedCommandsNV ptr_vkCmdPreprocessGeneratedCommandsNV;
VKAPI_ATTR void VKAPI_CALL vkCmdPreprocessGeneratedCommandsNV(VkCommandBuffer commandBuffer, VkGeneratedCommandsInfoNV* pGeneratedCommandsInfo) {
	ptr_vkCmdPreprocessGeneratedCommandsNV(commandBuffer, pGeneratedCommandsInfo);
}
*/
/*
PFN_vkCmdPushConstants ptr_vkCmdPushConstants;
VKAPI_ATTR void VKAPI_CALL vkCmdPushConstants(VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, void* pValues) {
	ptr_vkCmdPushConstants(commandBuffer, layout, stageFlags, offset, size, pValues);
}
*/
/*
PFN_vkCmdPushConstants2 ptr_vkCmdPushConstants2;
VKAPI_ATTR void VKAPI_CALL vkCmdPushConstants2(VkCommandBuffer commandBuffer, VkPushConstantsInfo* pPushConstantsInfo) {
	ptr_vkCmdPushConstants2(commandBuffer, pPushConstantsInfo);
}
*/
/*
PFN_vkCmdPushDescriptorSet ptr_vkCmdPushDescriptorSet;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSet(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t set, uint32_t descriptorWriteCount, VkWriteDescriptorSet* pDescriptorWrites) {
	ptr_vkCmdPushDescriptorSet(commandBuffer, pipelineBindPoint, layout, set, descriptorWriteCount, pDescriptorWrites);
}
*/
/*
PFN_vkCmdPushDescriptorSet2 ptr_vkCmdPushDescriptorSet2;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSet2(VkCommandBuffer commandBuffer, VkPushDescriptorSetInfo* pPushDescriptorSetInfo) {
	ptr_vkCmdPushDescriptorSet2(commandBuffer, pPushDescriptorSetInfo);
}
*/
/*
PFN_vkCmdPushDescriptorSetWithTemplate ptr_vkCmdPushDescriptorSetWithTemplate;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSetWithTemplate(VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, void* pData) {
	ptr_vkCmdPushDescriptorSetWithTemplate(commandBuffer, descriptorUpdateTemplate, layout, set, pData);
}
*/
/*
PFN_vkCmdPushDescriptorSetWithTemplate2 ptr_vkCmdPushDescriptorSetWithTemplate2;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSetWithTemplate2(VkCommandBuffer commandBuffer, VkPushDescriptorSetWithTemplateInfo* pPushDescriptorSetWithTemplateInfo) {
	ptr_vkCmdPushDescriptorSetWithTemplate2(commandBuffer, pPushDescriptorSetWithTemplateInfo);
}
*/
/*
PFN_vkCmdResetEvent ptr_vkCmdResetEvent;
VKAPI_ATTR void VKAPI_CALL vkCmdResetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	ptr_vkCmdResetEvent(commandBuffer, event, stageMask);
}
*/
/*
PFN_vkCmdResetEvent2 ptr_vkCmdResetEvent2;
VKAPI_ATTR void VKAPI_CALL vkCmdResetEvent2(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags2 stageMask) {
	ptr_vkCmdResetEvent2(commandBuffer, event, stageMask);
}
*/
/*
PFN_vkCmdResetQueryPool ptr_vkCmdResetQueryPool;
VKAPI_ATTR void VKAPI_CALL vkCmdResetQueryPool(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	ptr_vkCmdResetQueryPool(commandBuffer, queryPool, firstQuery, queryCount);
}
*/
/*
PFN_vkCmdResolveImage ptr_vkCmdResolveImage;
VKAPI_ATTR void VKAPI_CALL vkCmdResolveImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageResolve* pRegions) {
	ptr_vkCmdResolveImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}
*/
/*
PFN_vkCmdResolveImage2 ptr_vkCmdResolveImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdResolveImage2(VkCommandBuffer commandBuffer, VkResolveImageInfo2* pResolveImageInfo) {
	ptr_vkCmdResolveImage2(commandBuffer, pResolveImageInfo);
}
*/
/*
PFN_vkCmdSetAlphaToCoverageEnableEXT ptr_vkCmdSetAlphaToCoverageEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetAlphaToCoverageEnableEXT(VkCommandBuffer commandBuffer, VkBool32 alphaToCoverageEnable) {
	ptr_vkCmdSetAlphaToCoverageEnableEXT(commandBuffer, alphaToCoverageEnable);
}
*/
/*
PFN_vkCmdSetAlphaToOneEnableEXT ptr_vkCmdSetAlphaToOneEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetAlphaToOneEnableEXT(VkCommandBuffer commandBuffer, VkBool32 alphaToOneEnable) {
	ptr_vkCmdSetAlphaToOneEnableEXT(commandBuffer, alphaToOneEnable);
}
*/
/*
PFN_vkCmdSetAttachmentFeedbackLoopEnableEXT ptr_vkCmdSetAttachmentFeedbackLoopEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetAttachmentFeedbackLoopEnableEXT(VkCommandBuffer commandBuffer, VkImageAspectFlags aspectMask) {
	ptr_vkCmdSetAttachmentFeedbackLoopEnableEXT(commandBuffer, aspectMask);
}
*/
// // vkCmdSetBlendConstants.blendConstants is unuspported
/*
PFN_vkCmdSetCheckpointNV ptr_vkCmdSetCheckpointNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCheckpointNV(VkCommandBuffer commandBuffer, void* pCheckpointMarker) {
	ptr_vkCmdSetCheckpointNV(commandBuffer, pCheckpointMarker);
}
*/
/*
PFN_vkCmdSetCoarseSampleOrderNV ptr_vkCmdSetCoarseSampleOrderNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoarseSampleOrderNV(VkCommandBuffer commandBuffer, VkCoarseSampleOrderTypeNV sampleOrderType, uint32_t customSampleOrderCount, VkCoarseSampleOrderCustomNV* pCustomSampleOrders) {
	ptr_vkCmdSetCoarseSampleOrderNV(commandBuffer, sampleOrderType, customSampleOrderCount, pCustomSampleOrders);
}
*/
/*
PFN_vkCmdSetColorBlendAdvancedEXT ptr_vkCmdSetColorBlendAdvancedEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetColorBlendAdvancedEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, VkColorBlendAdvancedEXT* pColorBlendAdvanced) {
	ptr_vkCmdSetColorBlendAdvancedEXT(commandBuffer, firstAttachment, attachmentCount, pColorBlendAdvanced);
}
*/
/*
PFN_vkCmdSetColorBlendEnableEXT ptr_vkCmdSetColorBlendEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetColorBlendEnableEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, VkBool32* pColorBlendEnables) {
	ptr_vkCmdSetColorBlendEnableEXT(commandBuffer, firstAttachment, attachmentCount, pColorBlendEnables);
}
*/
/*
PFN_vkCmdSetColorBlendEquationEXT ptr_vkCmdSetColorBlendEquationEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetColorBlendEquationEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, VkColorBlendEquationEXT* pColorBlendEquations) {
	ptr_vkCmdSetColorBlendEquationEXT(commandBuffer, firstAttachment, attachmentCount, pColorBlendEquations);
}
*/
/*
PFN_vkCmdSetColorWriteEnableEXT ptr_vkCmdSetColorWriteEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetColorWriteEnableEXT(VkCommandBuffer commandBuffer, uint32_t attachmentCount, VkBool32* pColorWriteEnables) {
	ptr_vkCmdSetColorWriteEnableEXT(commandBuffer, attachmentCount, pColorWriteEnables);
}
*/
/*
PFN_vkCmdSetColorWriteMaskEXT ptr_vkCmdSetColorWriteMaskEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetColorWriteMaskEXT(VkCommandBuffer commandBuffer, uint32_t firstAttachment, uint32_t attachmentCount, VkColorComponentFlags* pColorWriteMasks) {
	ptr_vkCmdSetColorWriteMaskEXT(commandBuffer, firstAttachment, attachmentCount, pColorWriteMasks);
}
*/
/*
PFN_vkCmdSetConservativeRasterizationModeEXT ptr_vkCmdSetConservativeRasterizationModeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetConservativeRasterizationModeEXT(VkCommandBuffer commandBuffer, VkConservativeRasterizationModeEXT conservativeRasterizationMode) {
	ptr_vkCmdSetConservativeRasterizationModeEXT(commandBuffer, conservativeRasterizationMode);
}
*/
/*
PFN_vkCmdSetCoverageModulationModeNV ptr_vkCmdSetCoverageModulationModeNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoverageModulationModeNV(VkCommandBuffer commandBuffer, VkCoverageModulationModeNV coverageModulationMode) {
	ptr_vkCmdSetCoverageModulationModeNV(commandBuffer, coverageModulationMode);
}
*/
/*
PFN_vkCmdSetCoverageModulationTableEnableNV ptr_vkCmdSetCoverageModulationTableEnableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoverageModulationTableEnableNV(VkCommandBuffer commandBuffer, VkBool32 coverageModulationTableEnable) {
	ptr_vkCmdSetCoverageModulationTableEnableNV(commandBuffer, coverageModulationTableEnable);
}
*/
/*
PFN_vkCmdSetCoverageModulationTableNV ptr_vkCmdSetCoverageModulationTableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoverageModulationTableNV(VkCommandBuffer commandBuffer, uint32_t coverageModulationTableCount, float* pCoverageModulationTable) {
	ptr_vkCmdSetCoverageModulationTableNV(commandBuffer, coverageModulationTableCount, pCoverageModulationTable);
}
*/
/*
PFN_vkCmdSetCoverageReductionModeNV ptr_vkCmdSetCoverageReductionModeNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoverageReductionModeNV(VkCommandBuffer commandBuffer, VkCoverageReductionModeNV coverageReductionMode) {
	ptr_vkCmdSetCoverageReductionModeNV(commandBuffer, coverageReductionMode);
}
*/
/*
PFN_vkCmdSetCoverageToColorEnableNV ptr_vkCmdSetCoverageToColorEnableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoverageToColorEnableNV(VkCommandBuffer commandBuffer, VkBool32 coverageToColorEnable) {
	ptr_vkCmdSetCoverageToColorEnableNV(commandBuffer, coverageToColorEnable);
}
*/
/*
PFN_vkCmdSetCoverageToColorLocationNV ptr_vkCmdSetCoverageToColorLocationNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCoverageToColorLocationNV(VkCommandBuffer commandBuffer, uint32_t coverageToColorLocation) {
	ptr_vkCmdSetCoverageToColorLocationNV(commandBuffer, coverageToColorLocation);
}
*/
/*
PFN_vkCmdSetCullMode ptr_vkCmdSetCullMode;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCullMode(VkCommandBuffer commandBuffer, VkCullModeFlags cullMode) {
	ptr_vkCmdSetCullMode(commandBuffer, cullMode);
}
*/
/*
PFN_vkCmdSetDepthBias ptr_vkCmdSetDepthBias;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBias(VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor) {
	ptr_vkCmdSetDepthBias(commandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor);
}
*/
/*
PFN_vkCmdSetDepthBias2EXT ptr_vkCmdSetDepthBias2EXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBias2EXT(VkCommandBuffer commandBuffer, VkDepthBiasInfoEXT* pDepthBiasInfo) {
	ptr_vkCmdSetDepthBias2EXT(commandBuffer, pDepthBiasInfo);
}
*/
/*
PFN_vkCmdSetDepthBiasEnable ptr_vkCmdSetDepthBiasEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBiasEnable(VkCommandBuffer commandBuffer, VkBool32 depthBiasEnable) {
	ptr_vkCmdSetDepthBiasEnable(commandBuffer, depthBiasEnable);
}
*/
/*
PFN_vkCmdSetDepthBounds ptr_vkCmdSetDepthBounds;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBounds(VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds) {
	ptr_vkCmdSetDepthBounds(commandBuffer, minDepthBounds, maxDepthBounds);
}
*/
/*
PFN_vkCmdSetDepthBoundsTestEnable ptr_vkCmdSetDepthBoundsTestEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBoundsTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthBoundsTestEnable) {
	ptr_vkCmdSetDepthBoundsTestEnable(commandBuffer, depthBoundsTestEnable);
}
*/
/*
PFN_vkCmdSetDepthClampEnableEXT ptr_vkCmdSetDepthClampEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthClampEnableEXT(VkCommandBuffer commandBuffer, VkBool32 depthClampEnable) {
	ptr_vkCmdSetDepthClampEnableEXT(commandBuffer, depthClampEnable);
}
*/
/*
PFN_vkCmdSetDepthClampRangeEXT ptr_vkCmdSetDepthClampRangeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthClampRangeEXT(VkCommandBuffer commandBuffer, VkDepthClampModeEXT depthClampMode, VkDepthClampRangeEXT* pDepthClampRange) {
	ptr_vkCmdSetDepthClampRangeEXT(commandBuffer, depthClampMode, pDepthClampRange);
}
*/
/*
PFN_vkCmdSetDepthClipEnableEXT ptr_vkCmdSetDepthClipEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthClipEnableEXT(VkCommandBuffer commandBuffer, VkBool32 depthClipEnable) {
	ptr_vkCmdSetDepthClipEnableEXT(commandBuffer, depthClipEnable);
}
*/
/*
PFN_vkCmdSetDepthClipNegativeOneToOneEXT ptr_vkCmdSetDepthClipNegativeOneToOneEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthClipNegativeOneToOneEXT(VkCommandBuffer commandBuffer, VkBool32 negativeOneToOne) {
	ptr_vkCmdSetDepthClipNegativeOneToOneEXT(commandBuffer, negativeOneToOne);
}
*/
/*
PFN_vkCmdSetDepthCompareOp ptr_vkCmdSetDepthCompareOp;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthCompareOp(VkCommandBuffer commandBuffer, VkCompareOp depthCompareOp) {
	ptr_vkCmdSetDepthCompareOp(commandBuffer, depthCompareOp);
}
*/
/*
PFN_vkCmdSetDepthTestEnable ptr_vkCmdSetDepthTestEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthTestEnable) {
	ptr_vkCmdSetDepthTestEnable(commandBuffer, depthTestEnable);
}
*/
/*
PFN_vkCmdSetDepthWriteEnable ptr_vkCmdSetDepthWriteEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthWriteEnable(VkCommandBuffer commandBuffer, VkBool32 depthWriteEnable) {
	ptr_vkCmdSetDepthWriteEnable(commandBuffer, depthWriteEnable);
}
*/
/*
PFN_vkCmdSetDescriptorBufferOffsets2EXT ptr_vkCmdSetDescriptorBufferOffsets2EXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDescriptorBufferOffsets2EXT(VkCommandBuffer commandBuffer, VkSetDescriptorBufferOffsetsInfoEXT* pSetDescriptorBufferOffsetsInfo) {
	ptr_vkCmdSetDescriptorBufferOffsets2EXT(commandBuffer, pSetDescriptorBufferOffsetsInfo);
}
*/
/*
PFN_vkCmdSetDescriptorBufferOffsetsEXT ptr_vkCmdSetDescriptorBufferOffsetsEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDescriptorBufferOffsetsEXT(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t setCount, uint32_t* pBufferIndices, VkDeviceSize* pOffsets) {
	ptr_vkCmdSetDescriptorBufferOffsetsEXT(commandBuffer, pipelineBindPoint, layout, firstSet, setCount, pBufferIndices, pOffsets);
}
*/
/*
PFN_vkCmdSetDeviceMask ptr_vkCmdSetDeviceMask;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDeviceMask(VkCommandBuffer commandBuffer, uint32_t deviceMask) {
	ptr_vkCmdSetDeviceMask(commandBuffer, deviceMask);
}
*/
/*
PFN_vkCmdSetDiscardRectangleEXT ptr_vkCmdSetDiscardRectangleEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDiscardRectangleEXT(VkCommandBuffer commandBuffer, uint32_t firstDiscardRectangle, uint32_t discardRectangleCount, VkRect2D* pDiscardRectangles) {
	ptr_vkCmdSetDiscardRectangleEXT(commandBuffer, firstDiscardRectangle, discardRectangleCount, pDiscardRectangles);
}
*/
/*
PFN_vkCmdSetDiscardRectangleEnableEXT ptr_vkCmdSetDiscardRectangleEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDiscardRectangleEnableEXT(VkCommandBuffer commandBuffer, VkBool32 discardRectangleEnable) {
	ptr_vkCmdSetDiscardRectangleEnableEXT(commandBuffer, discardRectangleEnable);
}
*/
/*
PFN_vkCmdSetDiscardRectangleModeEXT ptr_vkCmdSetDiscardRectangleModeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDiscardRectangleModeEXT(VkCommandBuffer commandBuffer, VkDiscardRectangleModeEXT discardRectangleMode) {
	ptr_vkCmdSetDiscardRectangleModeEXT(commandBuffer, discardRectangleMode);
}
*/
/*
PFN_vkCmdSetEvent ptr_vkCmdSetEvent;
VKAPI_ATTR void VKAPI_CALL vkCmdSetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	ptr_vkCmdSetEvent(commandBuffer, event, stageMask);
}
*/
/*
PFN_vkCmdSetEvent2 ptr_vkCmdSetEvent2;
VKAPI_ATTR void VKAPI_CALL vkCmdSetEvent2(VkCommandBuffer commandBuffer, VkEvent event, VkDependencyInfo* pDependencyInfo) {
	ptr_vkCmdSetEvent2(commandBuffer, event, pDependencyInfo);
}
*/
/*
PFN_vkCmdSetExclusiveScissorEnableNV ptr_vkCmdSetExclusiveScissorEnableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetExclusiveScissorEnableNV(VkCommandBuffer commandBuffer, uint32_t firstExclusiveScissor, uint32_t exclusiveScissorCount, VkBool32* pExclusiveScissorEnables) {
	ptr_vkCmdSetExclusiveScissorEnableNV(commandBuffer, firstExclusiveScissor, exclusiveScissorCount, pExclusiveScissorEnables);
}
*/
/*
PFN_vkCmdSetExclusiveScissorNV ptr_vkCmdSetExclusiveScissorNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetExclusiveScissorNV(VkCommandBuffer commandBuffer, uint32_t firstExclusiveScissor, uint32_t exclusiveScissorCount, VkRect2D* pExclusiveScissors) {
	ptr_vkCmdSetExclusiveScissorNV(commandBuffer, firstExclusiveScissor, exclusiveScissorCount, pExclusiveScissors);
}
*/
/*
PFN_vkCmdSetExtraPrimitiveOverestimationSizeEXT ptr_vkCmdSetExtraPrimitiveOverestimationSizeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetExtraPrimitiveOverestimationSizeEXT(VkCommandBuffer commandBuffer, float extraPrimitiveOverestimationSize) {
	ptr_vkCmdSetExtraPrimitiveOverestimationSizeEXT(commandBuffer, extraPrimitiveOverestimationSize);
}
*/
// // vkCmdSetFragmentShadingRateEnumNV.combinerOps is unuspported
// // vkCmdSetFragmentShadingRateKHR.combinerOps is unuspported
/*
PFN_vkCmdSetFrontFace ptr_vkCmdSetFrontFace;
VKAPI_ATTR void VKAPI_CALL vkCmdSetFrontFace(VkCommandBuffer commandBuffer, VkFrontFace frontFace) {
	ptr_vkCmdSetFrontFace(commandBuffer, frontFace);
}
*/
/*
PFN_vkCmdSetLineRasterizationModeEXT ptr_vkCmdSetLineRasterizationModeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLineRasterizationModeEXT(VkCommandBuffer commandBuffer, VkLineRasterizationModeEXT lineRasterizationMode) {
	ptr_vkCmdSetLineRasterizationModeEXT(commandBuffer, lineRasterizationMode);
}
*/
/*
PFN_vkCmdSetLineStipple ptr_vkCmdSetLineStipple;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLineStipple(VkCommandBuffer commandBuffer, uint32_t lineStippleFactor, uint16_t lineStipplePattern) {
	ptr_vkCmdSetLineStipple(commandBuffer, lineStippleFactor, lineStipplePattern);
}
*/
/*
PFN_vkCmdSetLineStippleEnableEXT ptr_vkCmdSetLineStippleEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLineStippleEnableEXT(VkCommandBuffer commandBuffer, VkBool32 stippledLineEnable) {
	ptr_vkCmdSetLineStippleEnableEXT(commandBuffer, stippledLineEnable);
}
*/
/*
PFN_vkCmdSetLineWidth ptr_vkCmdSetLineWidth;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLineWidth(VkCommandBuffer commandBuffer, float lineWidth) {
	ptr_vkCmdSetLineWidth(commandBuffer, lineWidth);
}
*/
/*
PFN_vkCmdSetLogicOpEXT ptr_vkCmdSetLogicOpEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLogicOpEXT(VkCommandBuffer commandBuffer, VkLogicOp logicOp) {
	ptr_vkCmdSetLogicOpEXT(commandBuffer, logicOp);
}
*/
/*
PFN_vkCmdSetLogicOpEnableEXT ptr_vkCmdSetLogicOpEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLogicOpEnableEXT(VkCommandBuffer commandBuffer, VkBool32 logicOpEnable) {
	ptr_vkCmdSetLogicOpEnableEXT(commandBuffer, logicOpEnable);
}
*/
/*
PFN_vkCmdSetPatchControlPointsEXT ptr_vkCmdSetPatchControlPointsEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetPatchControlPointsEXT(VkCommandBuffer commandBuffer, uint32_t patchControlPoints) {
	ptr_vkCmdSetPatchControlPointsEXT(commandBuffer, patchControlPoints);
}
*/
/*
PFN_vkCmdSetPerformanceMarkerINTEL ptr_vkCmdSetPerformanceMarkerINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkCmdSetPerformanceMarkerINTEL(VkCommandBuffer commandBuffer, VkPerformanceMarkerInfoINTEL* pMarkerInfo) {
	return ptr_vkCmdSetPerformanceMarkerINTEL(commandBuffer, pMarkerInfo);
}
*/
/*
PFN_vkCmdSetPerformanceOverrideINTEL ptr_vkCmdSetPerformanceOverrideINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkCmdSetPerformanceOverrideINTEL(VkCommandBuffer commandBuffer, VkPerformanceOverrideInfoINTEL* pOverrideInfo) {
	return ptr_vkCmdSetPerformanceOverrideINTEL(commandBuffer, pOverrideInfo);
}
*/
/*
PFN_vkCmdSetPerformanceStreamMarkerINTEL ptr_vkCmdSetPerformanceStreamMarkerINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkCmdSetPerformanceStreamMarkerINTEL(VkCommandBuffer commandBuffer, VkPerformanceStreamMarkerInfoINTEL* pMarkerInfo) {
	return ptr_vkCmdSetPerformanceStreamMarkerINTEL(commandBuffer, pMarkerInfo);
}
*/
/*
PFN_vkCmdSetPolygonModeEXT ptr_vkCmdSetPolygonModeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetPolygonModeEXT(VkCommandBuffer commandBuffer, VkPolygonMode polygonMode) {
	ptr_vkCmdSetPolygonModeEXT(commandBuffer, polygonMode);
}
*/
/*
PFN_vkCmdSetPrimitiveRestartEnable ptr_vkCmdSetPrimitiveRestartEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetPrimitiveRestartEnable(VkCommandBuffer commandBuffer, VkBool32 primitiveRestartEnable) {
	ptr_vkCmdSetPrimitiveRestartEnable(commandBuffer, primitiveRestartEnable);
}
*/
/*
PFN_vkCmdSetPrimitiveTopology ptr_vkCmdSetPrimitiveTopology;
VKAPI_ATTR void VKAPI_CALL vkCmdSetPrimitiveTopology(VkCommandBuffer commandBuffer, VkPrimitiveTopology primitiveTopology) {
	ptr_vkCmdSetPrimitiveTopology(commandBuffer, primitiveTopology);
}
*/
/*
PFN_vkCmdSetProvokingVertexModeEXT ptr_vkCmdSetProvokingVertexModeEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetProvokingVertexModeEXT(VkCommandBuffer commandBuffer, VkProvokingVertexModeEXT provokingVertexMode) {
	ptr_vkCmdSetProvokingVertexModeEXT(commandBuffer, provokingVertexMode);
}
*/
/*
PFN_vkCmdSetRasterizationSamplesEXT ptr_vkCmdSetRasterizationSamplesEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRasterizationSamplesEXT(VkCommandBuffer commandBuffer, VkSampleCountFlagBits rasterizationSamples) {
	ptr_vkCmdSetRasterizationSamplesEXT(commandBuffer, rasterizationSamples);
}
*/
/*
PFN_vkCmdSetRasterizationStreamEXT ptr_vkCmdSetRasterizationStreamEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRasterizationStreamEXT(VkCommandBuffer commandBuffer, uint32_t rasterizationStream) {
	ptr_vkCmdSetRasterizationStreamEXT(commandBuffer, rasterizationStream);
}
*/
/*
PFN_vkCmdSetRasterizerDiscardEnable ptr_vkCmdSetRasterizerDiscardEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRasterizerDiscardEnable(VkCommandBuffer commandBuffer, VkBool32 rasterizerDiscardEnable) {
	ptr_vkCmdSetRasterizerDiscardEnable(commandBuffer, rasterizerDiscardEnable);
}
*/
/*
PFN_vkCmdSetRayTracingPipelineStackSizeKHR ptr_vkCmdSetRayTracingPipelineStackSizeKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRayTracingPipelineStackSizeKHR(VkCommandBuffer commandBuffer, uint32_t pipelineStackSize) {
	ptr_vkCmdSetRayTracingPipelineStackSizeKHR(commandBuffer, pipelineStackSize);
}
*/
/*
PFN_vkCmdSetRenderingAttachmentLocations ptr_vkCmdSetRenderingAttachmentLocations;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRenderingAttachmentLocations(VkCommandBuffer commandBuffer, VkRenderingAttachmentLocationInfo* pLocationInfo) {
	ptr_vkCmdSetRenderingAttachmentLocations(commandBuffer, pLocationInfo);
}
*/
/*
PFN_vkCmdSetRenderingInputAttachmentIndices ptr_vkCmdSetRenderingInputAttachmentIndices;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRenderingInputAttachmentIndices(VkCommandBuffer commandBuffer, VkRenderingInputAttachmentIndexInfo* pInputAttachmentIndexInfo) {
	ptr_vkCmdSetRenderingInputAttachmentIndices(commandBuffer, pInputAttachmentIndexInfo);
}
*/
/*
PFN_vkCmdSetRepresentativeFragmentTestEnableNV ptr_vkCmdSetRepresentativeFragmentTestEnableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRepresentativeFragmentTestEnableNV(VkCommandBuffer commandBuffer, VkBool32 representativeFragmentTestEnable) {
	ptr_vkCmdSetRepresentativeFragmentTestEnableNV(commandBuffer, representativeFragmentTestEnable);
}
*/
/*
PFN_vkCmdSetSampleLocationsEXT ptr_vkCmdSetSampleLocationsEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetSampleLocationsEXT(VkCommandBuffer commandBuffer, VkSampleLocationsInfoEXT* pSampleLocationsInfo) {
	ptr_vkCmdSetSampleLocationsEXT(commandBuffer, pSampleLocationsInfo);
}
*/
/*
PFN_vkCmdSetSampleLocationsEnableEXT ptr_vkCmdSetSampleLocationsEnableEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetSampleLocationsEnableEXT(VkCommandBuffer commandBuffer, VkBool32 sampleLocationsEnable) {
	ptr_vkCmdSetSampleLocationsEnableEXT(commandBuffer, sampleLocationsEnable);
}
*/
/*
PFN_vkCmdSetSampleMaskEXT ptr_vkCmdSetSampleMaskEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetSampleMaskEXT(VkCommandBuffer commandBuffer, VkSampleCountFlagBits samples, VkSampleMask* pSampleMask) {
	ptr_vkCmdSetSampleMaskEXT(commandBuffer, samples, pSampleMask);
}
*/
/*
PFN_vkCmdSetScissor ptr_vkCmdSetScissor;
VKAPI_ATTR void VKAPI_CALL vkCmdSetScissor(VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, VkRect2D* pScissors) {
	ptr_vkCmdSetScissor(commandBuffer, firstScissor, scissorCount, pScissors);
}
*/
/*
PFN_vkCmdSetScissorWithCount ptr_vkCmdSetScissorWithCount;
VKAPI_ATTR void VKAPI_CALL vkCmdSetScissorWithCount(VkCommandBuffer commandBuffer, uint32_t scissorCount, VkRect2D* pScissors) {
	ptr_vkCmdSetScissorWithCount(commandBuffer, scissorCount, pScissors);
}
*/
/*
PFN_vkCmdSetShadingRateImageEnableNV ptr_vkCmdSetShadingRateImageEnableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetShadingRateImageEnableNV(VkCommandBuffer commandBuffer, VkBool32 shadingRateImageEnable) {
	ptr_vkCmdSetShadingRateImageEnableNV(commandBuffer, shadingRateImageEnable);
}
*/
/*
PFN_vkCmdSetStencilCompareMask ptr_vkCmdSetStencilCompareMask;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilCompareMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t compareMask) {
	ptr_vkCmdSetStencilCompareMask(commandBuffer, faceMask, compareMask);
}
*/
/*
PFN_vkCmdSetStencilOp ptr_vkCmdSetStencilOp;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilOp(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, VkStencilOp failOp, VkStencilOp passOp, VkStencilOp depthFailOp, VkCompareOp compareOp) {
	ptr_vkCmdSetStencilOp(commandBuffer, faceMask, failOp, passOp, depthFailOp, compareOp);
}
*/
/*
PFN_vkCmdSetStencilReference ptr_vkCmdSetStencilReference;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilReference(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t reference) {
	ptr_vkCmdSetStencilReference(commandBuffer, faceMask, reference);
}
*/
/*
PFN_vkCmdSetStencilTestEnable ptr_vkCmdSetStencilTestEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilTestEnable(VkCommandBuffer commandBuffer, VkBool32 stencilTestEnable) {
	ptr_vkCmdSetStencilTestEnable(commandBuffer, stencilTestEnable);
}
*/
/*
PFN_vkCmdSetStencilWriteMask ptr_vkCmdSetStencilWriteMask;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilWriteMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t writeMask) {
	ptr_vkCmdSetStencilWriteMask(commandBuffer, faceMask, writeMask);
}
*/
/*
PFN_vkCmdSetTessellationDomainOriginEXT ptr_vkCmdSetTessellationDomainOriginEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetTessellationDomainOriginEXT(VkCommandBuffer commandBuffer, VkTessellationDomainOrigin domainOrigin) {
	ptr_vkCmdSetTessellationDomainOriginEXT(commandBuffer, domainOrigin);
}
*/
/*
PFN_vkCmdSetVertexInputEXT ptr_vkCmdSetVertexInputEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdSetVertexInputEXT(VkCommandBuffer commandBuffer, uint32_t vertexBindingDescriptionCount, VkVertexInputBindingDescription2EXT* pVertexBindingDescriptions, uint32_t vertexAttributeDescriptionCount, VkVertexInputAttributeDescription2EXT* pVertexAttributeDescriptions) {
	ptr_vkCmdSetVertexInputEXT(commandBuffer, vertexBindingDescriptionCount, pVertexBindingDescriptions, vertexAttributeDescriptionCount, pVertexAttributeDescriptions);
}
*/
/*
PFN_vkCmdSetViewport ptr_vkCmdSetViewport;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewport(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, VkViewport* pViewports) {
	ptr_vkCmdSetViewport(commandBuffer, firstViewport, viewportCount, pViewports);
}
*/
/*
PFN_vkCmdSetViewportShadingRatePaletteNV ptr_vkCmdSetViewportShadingRatePaletteNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewportShadingRatePaletteNV(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, VkShadingRatePaletteNV* pShadingRatePalettes) {
	ptr_vkCmdSetViewportShadingRatePaletteNV(commandBuffer, firstViewport, viewportCount, pShadingRatePalettes);
}
*/
/*
PFN_vkCmdSetViewportSwizzleNV ptr_vkCmdSetViewportSwizzleNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewportSwizzleNV(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, VkViewportSwizzleNV* pViewportSwizzles) {
	ptr_vkCmdSetViewportSwizzleNV(commandBuffer, firstViewport, viewportCount, pViewportSwizzles);
}
*/
/*
PFN_vkCmdSetViewportWScalingEnableNV ptr_vkCmdSetViewportWScalingEnableNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewportWScalingEnableNV(VkCommandBuffer commandBuffer, VkBool32 viewportWScalingEnable) {
	ptr_vkCmdSetViewportWScalingEnableNV(commandBuffer, viewportWScalingEnable);
}
*/
/*
PFN_vkCmdSetViewportWScalingNV ptr_vkCmdSetViewportWScalingNV;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewportWScalingNV(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, VkViewportWScalingNV* pViewportWScalings) {
	ptr_vkCmdSetViewportWScalingNV(commandBuffer, firstViewport, viewportCount, pViewportWScalings);
}
*/
/*
PFN_vkCmdSetViewportWithCount ptr_vkCmdSetViewportWithCount;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewportWithCount(VkCommandBuffer commandBuffer, uint32_t viewportCount, VkViewport* pViewports) {
	ptr_vkCmdSetViewportWithCount(commandBuffer, viewportCount, pViewports);
}
*/
/*
PFN_vkCmdSubpassShadingHUAWEI ptr_vkCmdSubpassShadingHUAWEI;
VKAPI_ATTR void VKAPI_CALL vkCmdSubpassShadingHUAWEI(VkCommandBuffer commandBuffer) {
	ptr_vkCmdSubpassShadingHUAWEI(commandBuffer);
}
*/
/*
PFN_vkCmdTraceRaysIndirect2KHR ptr_vkCmdTraceRaysIndirect2KHR;
VKAPI_ATTR void VKAPI_CALL vkCmdTraceRaysIndirect2KHR(VkCommandBuffer commandBuffer, VkDeviceAddress indirectDeviceAddress) {
	ptr_vkCmdTraceRaysIndirect2KHR(commandBuffer, indirectDeviceAddress);
}
*/
/*
PFN_vkCmdTraceRaysIndirectKHR ptr_vkCmdTraceRaysIndirectKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdTraceRaysIndirectKHR(VkCommandBuffer commandBuffer, VkStridedDeviceAddressRegionKHR* pRaygenShaderBindingTable, VkStridedDeviceAddressRegionKHR* pMissShaderBindingTable, VkStridedDeviceAddressRegionKHR* pHitShaderBindingTable, VkStridedDeviceAddressRegionKHR* pCallableShaderBindingTable, VkDeviceAddress indirectDeviceAddress) {
	ptr_vkCmdTraceRaysIndirectKHR(commandBuffer, pRaygenShaderBindingTable, pMissShaderBindingTable, pHitShaderBindingTable, pCallableShaderBindingTable, indirectDeviceAddress);
}
*/
/*
PFN_vkCmdTraceRaysKHR ptr_vkCmdTraceRaysKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdTraceRaysKHR(VkCommandBuffer commandBuffer, VkStridedDeviceAddressRegionKHR* pRaygenShaderBindingTable, VkStridedDeviceAddressRegionKHR* pMissShaderBindingTable, VkStridedDeviceAddressRegionKHR* pHitShaderBindingTable, VkStridedDeviceAddressRegionKHR* pCallableShaderBindingTable, uint32_t width, uint32_t height, uint32_t depth) {
	ptr_vkCmdTraceRaysKHR(commandBuffer, pRaygenShaderBindingTable, pMissShaderBindingTable, pHitShaderBindingTable, pCallableShaderBindingTable, width, height, depth);
}
*/
/*
PFN_vkCmdTraceRaysNV ptr_vkCmdTraceRaysNV;
VKAPI_ATTR void VKAPI_CALL vkCmdTraceRaysNV(VkCommandBuffer commandBuffer, VkBuffer raygenShaderBindingTableBuffer, VkDeviceSize raygenShaderBindingOffset, VkBuffer missShaderBindingTableBuffer, VkDeviceSize missShaderBindingOffset, VkDeviceSize missShaderBindingStride, VkBuffer hitShaderBindingTableBuffer, VkDeviceSize hitShaderBindingOffset, VkDeviceSize hitShaderBindingStride, VkBuffer callableShaderBindingTableBuffer, VkDeviceSize callableShaderBindingOffset, VkDeviceSize callableShaderBindingStride, uint32_t width, uint32_t height, uint32_t depth) {
	ptr_vkCmdTraceRaysNV(commandBuffer, raygenShaderBindingTableBuffer, raygenShaderBindingOffset, missShaderBindingTableBuffer, missShaderBindingOffset, missShaderBindingStride, hitShaderBindingTableBuffer, hitShaderBindingOffset, hitShaderBindingStride, callableShaderBindingTableBuffer, callableShaderBindingOffset, callableShaderBindingStride, width, height, depth);
}
*/
/*
PFN_vkCmdUpdateBuffer ptr_vkCmdUpdateBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdUpdateBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize dataSize, void* pData) {
	ptr_vkCmdUpdateBuffer(commandBuffer, dstBuffer, dstOffset, dataSize, pData);
}
*/
/*
PFN_vkCmdUpdatePipelineIndirectBufferNV ptr_vkCmdUpdatePipelineIndirectBufferNV;
VKAPI_ATTR void VKAPI_CALL vkCmdUpdatePipelineIndirectBufferNV(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
	ptr_vkCmdUpdatePipelineIndirectBufferNV(commandBuffer, pipelineBindPoint, pipeline);
}
*/
/*
PFN_vkCmdWaitEvents ptr_vkCmdWaitEvents;
VKAPI_ATTR void VKAPI_CALL vkCmdWaitEvents(VkCommandBuffer commandBuffer, uint32_t eventCount, VkEvent* pEvents, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, uint32_t memoryBarrierCount, VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, VkImageMemoryBarrier* pImageMemoryBarriers) {
	ptr_vkCmdWaitEvents(commandBuffer, eventCount, pEvents, srcStageMask, dstStageMask, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}
*/
/*
PFN_vkCmdWaitEvents2 ptr_vkCmdWaitEvents2;
VKAPI_ATTR void VKAPI_CALL vkCmdWaitEvents2(VkCommandBuffer commandBuffer, uint32_t eventCount, VkEvent* pEvents, VkDependencyInfo* pDependencyInfos) {
	ptr_vkCmdWaitEvents2(commandBuffer, eventCount, pEvents, pDependencyInfos);
}
*/
/*
PFN_vkCmdWriteAccelerationStructuresPropertiesKHR ptr_vkCmdWriteAccelerationStructuresPropertiesKHR;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteAccelerationStructuresPropertiesKHR(VkCommandBuffer commandBuffer, uint32_t accelerationStructureCount, VkAccelerationStructureKHR* pAccelerationStructures, VkQueryType queryType, VkQueryPool queryPool, uint32_t firstQuery) {
	ptr_vkCmdWriteAccelerationStructuresPropertiesKHR(commandBuffer, accelerationStructureCount, pAccelerationStructures, queryType, queryPool, firstQuery);
}
*/
/*
PFN_vkCmdWriteAccelerationStructuresPropertiesNV ptr_vkCmdWriteAccelerationStructuresPropertiesNV;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteAccelerationStructuresPropertiesNV(VkCommandBuffer commandBuffer, uint32_t accelerationStructureCount, VkAccelerationStructureNV* pAccelerationStructures, VkQueryType queryType, VkQueryPool queryPool, uint32_t firstQuery) {
	ptr_vkCmdWriteAccelerationStructuresPropertiesNV(commandBuffer, accelerationStructureCount, pAccelerationStructures, queryType, queryPool, firstQuery);
}
*/
/*
PFN_vkCmdWriteBufferMarker2AMD ptr_vkCmdWriteBufferMarker2AMD;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteBufferMarker2AMD(VkCommandBuffer commandBuffer, VkPipelineStageFlags2 stage, VkBuffer dstBuffer, VkDeviceSize dstOffset, uint32_t marker) {
	ptr_vkCmdWriteBufferMarker2AMD(commandBuffer, stage, dstBuffer, dstOffset, marker);
}
*/
/*
PFN_vkCmdWriteBufferMarkerAMD ptr_vkCmdWriteBufferMarkerAMD;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteBufferMarkerAMD(VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkBuffer dstBuffer, VkDeviceSize dstOffset, uint32_t marker) {
	ptr_vkCmdWriteBufferMarkerAMD(commandBuffer, pipelineStage, dstBuffer, dstOffset, marker);
}
*/
/*
PFN_vkCmdWriteMicromapsPropertiesEXT ptr_vkCmdWriteMicromapsPropertiesEXT;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteMicromapsPropertiesEXT(VkCommandBuffer commandBuffer, uint32_t micromapCount, VkMicromapEXT* pMicromaps, VkQueryType queryType, VkQueryPool queryPool, uint32_t firstQuery) {
	ptr_vkCmdWriteMicromapsPropertiesEXT(commandBuffer, micromapCount, pMicromaps, queryType, queryPool, firstQuery);
}
*/
/*
PFN_vkCmdWriteTimestamp ptr_vkCmdWriteTimestamp;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteTimestamp(VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkQueryPool queryPool, uint32_t query) {
	ptr_vkCmdWriteTimestamp(commandBuffer, pipelineStage, queryPool, query);
}
*/
/*
PFN_vkCmdWriteTimestamp2 ptr_vkCmdWriteTimestamp2;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteTimestamp2(VkCommandBuffer commandBuffer, VkPipelineStageFlags2 stage, VkQueryPool queryPool, uint32_t query) {
	ptr_vkCmdWriteTimestamp2(commandBuffer, stage, queryPool, query);
}
*/
/*
PFN_vkCompileDeferredNV ptr_vkCompileDeferredNV;
VKAPI_ATTR VkResult VKAPI_CALL vkCompileDeferredNV(VkDevice device, VkPipeline pipeline, uint32_t shader) {
	return ptr_vkCompileDeferredNV(device, pipeline, shader);
}
*/
/*
PFN_vkConvertCooperativeVectorMatrixNV ptr_vkConvertCooperativeVectorMatrixNV;
VKAPI_ATTR VkResult VKAPI_CALL vkConvertCooperativeVectorMatrixNV(VkDevice device, VkConvertCooperativeVectorMatrixInfoNV* pInfo) {
	return ptr_vkConvertCooperativeVectorMatrixNV(device, pInfo);
}
*/
/*
PFN_vkCopyAccelerationStructureKHR ptr_vkCopyAccelerationStructureKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyAccelerationStructureKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyAccelerationStructureInfoKHR* pInfo) {
	return ptr_vkCopyAccelerationStructureKHR(device, deferredOperation, pInfo);
}
*/
/*
PFN_vkCopyAccelerationStructureToMemoryKHR ptr_vkCopyAccelerationStructureToMemoryKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyAccelerationStructureToMemoryKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyAccelerationStructureToMemoryInfoKHR* pInfo) {
	return ptr_vkCopyAccelerationStructureToMemoryKHR(device, deferredOperation, pInfo);
}
*/
/*
PFN_vkCopyImageToImage ptr_vkCopyImageToImage;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyImageToImage(VkDevice device, VkCopyImageToImageInfo* pCopyImageToImageInfo) {
	return ptr_vkCopyImageToImage(device, pCopyImageToImageInfo);
}
*/
/*
PFN_vkCopyImageToMemory ptr_vkCopyImageToMemory;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyImageToMemory(VkDevice device, VkCopyImageToMemoryInfo* pCopyImageToMemoryInfo) {
	return ptr_vkCopyImageToMemory(device, pCopyImageToMemoryInfo);
}
*/
/*
PFN_vkCopyMemoryToAccelerationStructureKHR ptr_vkCopyMemoryToAccelerationStructureKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyMemoryToAccelerationStructureKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyMemoryToAccelerationStructureInfoKHR* pInfo) {
	return ptr_vkCopyMemoryToAccelerationStructureKHR(device, deferredOperation, pInfo);
}
*/
/*
PFN_vkCopyMemoryToImage ptr_vkCopyMemoryToImage;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyMemoryToImage(VkDevice device, VkCopyMemoryToImageInfo* pCopyMemoryToImageInfo) {
	return ptr_vkCopyMemoryToImage(device, pCopyMemoryToImageInfo);
}
*/
/*
PFN_vkCopyMemoryToMicromapEXT ptr_vkCopyMemoryToMicromapEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyMemoryToMicromapEXT(VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyMemoryToMicromapInfoEXT* pInfo) {
	return ptr_vkCopyMemoryToMicromapEXT(device, deferredOperation, pInfo);
}
*/
/*
PFN_vkCopyMicromapEXT ptr_vkCopyMicromapEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyMicromapEXT(VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyMicromapInfoEXT* pInfo) {
	return ptr_vkCopyMicromapEXT(device, deferredOperation, pInfo);
}
*/
/*
PFN_vkCopyMicromapToMemoryEXT ptr_vkCopyMicromapToMemoryEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyMicromapToMemoryEXT(VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyMicromapToMemoryInfoEXT* pInfo) {
	return ptr_vkCopyMicromapToMemoryEXT(device, deferredOperation, pInfo);
}
*/
/*
PFN_vkCreateAccelerationStructureKHR ptr_vkCreateAccelerationStructureKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateAccelerationStructureKHR(VkDevice device, VkAccelerationStructureCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkAccelerationStructureKHR* pAccelerationStructure) {
	return ptr_vkCreateAccelerationStructureKHR(device, pCreateInfo, pAllocator, pAccelerationStructure);
}
*/
/*
PFN_vkCreateAccelerationStructureNV ptr_vkCreateAccelerationStructureNV;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateAccelerationStructureNV(VkDevice device, VkAccelerationStructureCreateInfoNV* pCreateInfo, VkAllocationCallbacks* pAllocator, VkAccelerationStructureNV* pAccelerationStructure) {
	return ptr_vkCreateAccelerationStructureNV(device, pCreateInfo, pAllocator, pAccelerationStructure);
}
*/
/*
PFN_vkCreateBuffer ptr_vkCreateBuffer;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateBuffer(VkDevice device, VkBufferCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer) {
	return ptr_vkCreateBuffer(device, pCreateInfo, pAllocator, pBuffer);
}
*/
/*
PFN_vkCreateBufferView ptr_vkCreateBufferView;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateBufferView(VkDevice device, VkBufferViewCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkBufferView* pView) {
	return ptr_vkCreateBufferView(device, pCreateInfo, pAllocator, pView);
}
*/
/*
PFN_vkCreateCommandPool ptr_vkCreateCommandPool;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateCommandPool(VkDevice device, VkCommandPoolCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool) {
	return ptr_vkCreateCommandPool(device, pCreateInfo, pAllocator, pCommandPool);
}
*/
/*
PFN_vkCreateComputePipelines ptr_vkCreateComputePipelines;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateComputePipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkComputePipelineCreateInfo* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return ptr_vkCreateComputePipelines(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
*/
/*
PFN_vkCreateCuFunctionNVX ptr_vkCreateCuFunctionNVX;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateCuFunctionNVX(VkDevice device, VkCuFunctionCreateInfoNVX* pCreateInfo, VkAllocationCallbacks* pAllocator, VkCuFunctionNVX* pFunction) {
	return ptr_vkCreateCuFunctionNVX(device, pCreateInfo, pAllocator, pFunction);
}
*/
/*
PFN_vkCreateCuModuleNVX ptr_vkCreateCuModuleNVX;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateCuModuleNVX(VkDevice device, VkCuModuleCreateInfoNVX* pCreateInfo, VkAllocationCallbacks* pAllocator, VkCuModuleNVX* pModule) {
	return ptr_vkCreateCuModuleNVX(device, pCreateInfo, pAllocator, pModule);
}
*/
/*
PFN_vkCreateDataGraphPipelineSessionARM ptr_vkCreateDataGraphPipelineSessionARM;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDataGraphPipelineSessionARM(VkDevice device, VkDataGraphPipelineSessionCreateInfoARM* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDataGraphPipelineSessionARM* pSession) {
	return ptr_vkCreateDataGraphPipelineSessionARM(device, pCreateInfo, pAllocator, pSession);
}
*/
/*
PFN_vkCreateDataGraphPipelinesARM ptr_vkCreateDataGraphPipelinesARM;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDataGraphPipelinesARM(VkDevice device, VkDeferredOperationKHR deferredOperation, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkDataGraphPipelineCreateInfoARM* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return ptr_vkCreateDataGraphPipelinesARM(device, deferredOperation, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
*/
/*
PFN_vkCreateDebugReportCallbackEXT ptr_vkCreateDebugReportCallbackEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDebugReportCallbackEXT(VkInstance instance, VkDebugReportCallbackCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDebugReportCallbackEXT* pCallback) {
	return ptr_vkCreateDebugReportCallbackEXT(instance, pCreateInfo, pAllocator, pCallback);
}
*/
/*
PFN_vkCreateDebugUtilsMessengerEXT ptr_vkCreateDebugUtilsMessengerEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDebugUtilsMessengerEXT(VkInstance instance, VkDebugUtilsMessengerCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDebugUtilsMessengerEXT* pMessenger) {
	return ptr_vkCreateDebugUtilsMessengerEXT(instance, pCreateInfo, pAllocator, pMessenger);
}
*/
/*
PFN_vkCreateDeferredOperationKHR ptr_vkCreateDeferredOperationKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDeferredOperationKHR(VkDevice device, VkAllocationCallbacks* pAllocator, VkDeferredOperationKHR* pDeferredOperation) {
	return ptr_vkCreateDeferredOperationKHR(device, pAllocator, pDeferredOperation);
}
*/
/*
PFN_vkCreateDescriptorPool ptr_vkCreateDescriptorPool;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDescriptorPool(VkDevice device, VkDescriptorPoolCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDescriptorPool* pDescriptorPool) {
	return ptr_vkCreateDescriptorPool(device, pCreateInfo, pAllocator, pDescriptorPool);
}
*/
/*
PFN_vkCreateDescriptorSetLayout ptr_vkCreateDescriptorSetLayout;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDescriptorSetLayout(VkDevice device, VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDescriptorSetLayout* pSetLayout) {
	return ptr_vkCreateDescriptorSetLayout(device, pCreateInfo, pAllocator, pSetLayout);
}
*/
/*
PFN_vkCreateDescriptorUpdateTemplate ptr_vkCreateDescriptorUpdateTemplate;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDescriptorUpdateTemplate(VkDevice device, VkDescriptorUpdateTemplateCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDescriptorUpdateTemplate* pDescriptorUpdateTemplate) {
	return ptr_vkCreateDescriptorUpdateTemplate(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}
*/
/*
PFN_vkCreateDevice ptr_vkCreateDevice;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDevice(VkPhysicalDevice physicalDevice, VkDeviceCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDevice* pDevice) {
	return ptr_vkCreateDevice(physicalDevice, pCreateInfo, pAllocator, pDevice);
}
*/
/*
PFN_vkCreateDisplayModeKHR ptr_vkCreateDisplayModeKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDisplayModeKHR(VkPhysicalDevice physicalDevice, VkDisplayKHR display, VkDisplayModeCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDisplayModeKHR* pMode) {
	return ptr_vkCreateDisplayModeKHR(physicalDevice, display, pCreateInfo, pAllocator, pMode);
}
*/
/*
PFN_vkCreateDisplayPlaneSurfaceKHR ptr_vkCreateDisplayPlaneSurfaceKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateDisplayPlaneSurfaceKHR(VkInstance instance, VkDisplaySurfaceCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return ptr_vkCreateDisplayPlaneSurfaceKHR(instance, pCreateInfo, pAllocator, pSurface);
}
*/
/*
PFN_vkCreateEvent ptr_vkCreateEvent;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateEvent(VkDevice device, VkEventCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkEvent* pEvent) {
	return ptr_vkCreateEvent(device, pCreateInfo, pAllocator, pEvent);
}
*/
/*
PFN_vkCreateExternalComputeQueueNV ptr_vkCreateExternalComputeQueueNV;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateExternalComputeQueueNV(VkDevice device, VkExternalComputeQueueCreateInfoNV* pCreateInfo, VkAllocationCallbacks* pAllocator, VkExternalComputeQueueNV* pExternalQueue) {
	return ptr_vkCreateExternalComputeQueueNV(device, pCreateInfo, pAllocator, pExternalQueue);
}
*/
/*
PFN_vkCreateFence ptr_vkCreateFence;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateFence(VkDevice device, VkFenceCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return ptr_vkCreateFence(device, pCreateInfo, pAllocator, pFence);
}
*/
/*
PFN_vkCreateFramebuffer ptr_vkCreateFramebuffer;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateFramebuffer(VkDevice device, VkFramebufferCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer) {
	return ptr_vkCreateFramebuffer(device, pCreateInfo, pAllocator, pFramebuffer);
}
*/
/*
PFN_vkCreateGraphicsPipelines ptr_vkCreateGraphicsPipelines;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateGraphicsPipelines(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkGraphicsPipelineCreateInfo* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return ptr_vkCreateGraphicsPipelines(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
*/
/*
PFN_vkCreateHeadlessSurfaceEXT ptr_vkCreateHeadlessSurfaceEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateHeadlessSurfaceEXT(VkInstance instance, VkHeadlessSurfaceCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return ptr_vkCreateHeadlessSurfaceEXT(instance, pCreateInfo, pAllocator, pSurface);
}
*/
/*
PFN_vkCreateImage ptr_vkCreateImage;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateImage(VkDevice device, VkImageCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkImage* pImage) {
	return ptr_vkCreateImage(device, pCreateInfo, pAllocator, pImage);
}
*/
/*
PFN_vkCreateImageView ptr_vkCreateImageView;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateImageView(VkDevice device, VkImageViewCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkImageView* pView) {
	return ptr_vkCreateImageView(device, pCreateInfo, pAllocator, pView);
}
*/
/*
PFN_vkCreateIndirectCommandsLayoutEXT ptr_vkCreateIndirectCommandsLayoutEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateIndirectCommandsLayoutEXT(VkDevice device, VkIndirectCommandsLayoutCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkIndirectCommandsLayoutEXT* pIndirectCommandsLayout) {
	return ptr_vkCreateIndirectCommandsLayoutEXT(device, pCreateInfo, pAllocator, pIndirectCommandsLayout);
}
*/
/*
PFN_vkCreateIndirectCommandsLayoutNV ptr_vkCreateIndirectCommandsLayoutNV;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateIndirectCommandsLayoutNV(VkDevice device, VkIndirectCommandsLayoutCreateInfoNV* pCreateInfo, VkAllocationCallbacks* pAllocator, VkIndirectCommandsLayoutNV* pIndirectCommandsLayout) {
	return ptr_vkCreateIndirectCommandsLayoutNV(device, pCreateInfo, pAllocator, pIndirectCommandsLayout);
}
*/
/*
PFN_vkCreateIndirectExecutionSetEXT ptr_vkCreateIndirectExecutionSetEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateIndirectExecutionSetEXT(VkDevice device, VkIndirectExecutionSetCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkIndirectExecutionSetEXT* pIndirectExecutionSet) {
	return ptr_vkCreateIndirectExecutionSetEXT(device, pCreateInfo, pAllocator, pIndirectExecutionSet);
}
*/
/*
PFN_vkCreateInstance ptr_vkCreateInstance;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateInstance(VkInstanceCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkInstance* pInstance) {
	return ptr_vkCreateInstance(pCreateInfo, pAllocator, pInstance);
}
*/
/*
PFN_vkCreateMetalSurfaceEXT ptr_vkCreateMetalSurfaceEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateMetalSurfaceEXT(VkInstance instance, VkMetalSurfaceCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return ptr_vkCreateMetalSurfaceEXT(instance, pCreateInfo, pAllocator, pSurface);
}
*/
/*
PFN_vkCreateMicromapEXT ptr_vkCreateMicromapEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateMicromapEXT(VkDevice device, VkMicromapCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkMicromapEXT* pMicromap) {
	return ptr_vkCreateMicromapEXT(device, pCreateInfo, pAllocator, pMicromap);
}
*/
/*
PFN_vkCreateOpticalFlowSessionNV ptr_vkCreateOpticalFlowSessionNV;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateOpticalFlowSessionNV(VkDevice device, VkOpticalFlowSessionCreateInfoNV* pCreateInfo, VkAllocationCallbacks* pAllocator, VkOpticalFlowSessionNV* pSession) {
	return ptr_vkCreateOpticalFlowSessionNV(device, pCreateInfo, pAllocator, pSession);
}
*/
/*
PFN_vkCreatePipelineBinariesKHR ptr_vkCreatePipelineBinariesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreatePipelineBinariesKHR(VkDevice device, VkPipelineBinaryCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkPipelineBinaryHandlesInfoKHR* pBinaries) {
	return ptr_vkCreatePipelineBinariesKHR(device, pCreateInfo, pAllocator, pBinaries);
}
*/
/*
PFN_vkCreatePipelineCache ptr_vkCreatePipelineCache;
VKAPI_ATTR VkResult VKAPI_CALL vkCreatePipelineCache(VkDevice device, VkPipelineCacheCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkPipelineCache* pPipelineCache) {
	return ptr_vkCreatePipelineCache(device, pCreateInfo, pAllocator, pPipelineCache);
}
*/
/*
PFN_vkCreatePipelineLayout ptr_vkCreatePipelineLayout;
VKAPI_ATTR VkResult VKAPI_CALL vkCreatePipelineLayout(VkDevice device, VkPipelineLayoutCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout) {
	return ptr_vkCreatePipelineLayout(device, pCreateInfo, pAllocator, pPipelineLayout);
}
*/
/*
PFN_vkCreatePrivateDataSlot ptr_vkCreatePrivateDataSlot;
VKAPI_ATTR VkResult VKAPI_CALL vkCreatePrivateDataSlot(VkDevice device, VkPrivateDataSlotCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkPrivateDataSlot* pPrivateDataSlot) {
	return ptr_vkCreatePrivateDataSlot(device, pCreateInfo, pAllocator, pPrivateDataSlot);
}
*/
/*
PFN_vkCreateQueryPool ptr_vkCreateQueryPool;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateQueryPool(VkDevice device, VkQueryPoolCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkQueryPool* pQueryPool) {
	return ptr_vkCreateQueryPool(device, pCreateInfo, pAllocator, pQueryPool);
}
*/
/*
PFN_vkCreateRayTracingPipelinesKHR ptr_vkCreateRayTracingPipelinesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateRayTracingPipelinesKHR(VkDevice device, VkDeferredOperationKHR deferredOperation, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkRayTracingPipelineCreateInfoKHR* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return ptr_vkCreateRayTracingPipelinesKHR(device, deferredOperation, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
*/
/*
PFN_vkCreateRayTracingPipelinesNV ptr_vkCreateRayTracingPipelinesNV;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateRayTracingPipelinesNV(VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkRayTracingPipelineCreateInfoNV* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
	return ptr_vkCreateRayTracingPipelinesNV(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}
*/
/*
PFN_vkCreateRenderPass ptr_vkCreateRenderPass;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateRenderPass(VkDevice device, VkRenderPassCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return ptr_vkCreateRenderPass(device, pCreateInfo, pAllocator, pRenderPass);
}
*/
/*
PFN_vkCreateRenderPass2 ptr_vkCreateRenderPass2;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateRenderPass2(VkDevice device, VkRenderPassCreateInfo2* pCreateInfo, VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
	return ptr_vkCreateRenderPass2(device, pCreateInfo, pAllocator, pRenderPass);
}
*/
/*
PFN_vkCreateSampler ptr_vkCreateSampler;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateSampler(VkDevice device, VkSamplerCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSampler* pSampler) {
	return ptr_vkCreateSampler(device, pCreateInfo, pAllocator, pSampler);
}
*/
/*
PFN_vkCreateSamplerYcbcrConversion ptr_vkCreateSamplerYcbcrConversion;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateSamplerYcbcrConversion(VkDevice device, VkSamplerYcbcrConversionCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSamplerYcbcrConversion* pYcbcrConversion) {
	return ptr_vkCreateSamplerYcbcrConversion(device, pCreateInfo, pAllocator, pYcbcrConversion);
}
*/
/*
PFN_vkCreateSemaphore ptr_vkCreateSemaphore;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateSemaphore(VkDevice device, VkSemaphoreCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore) {
	return ptr_vkCreateSemaphore(device, pCreateInfo, pAllocator, pSemaphore);
}
*/
/*
PFN_vkCreateShaderModule ptr_vkCreateShaderModule;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateShaderModule(VkDevice device, VkShaderModuleCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule) {
	return ptr_vkCreateShaderModule(device, pCreateInfo, pAllocator, pShaderModule);
}
*/
/*
PFN_vkCreateShadersEXT ptr_vkCreateShadersEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateShadersEXT(VkDevice device, uint32_t createInfoCount, VkShaderCreateInfoEXT* pCreateInfos, VkAllocationCallbacks* pAllocator, VkShaderEXT* pShaders) {
	return ptr_vkCreateShadersEXT(device, createInfoCount, pCreateInfos, pAllocator, pShaders);
}
*/
/*
PFN_vkCreateSharedSwapchainsKHR ptr_vkCreateSharedSwapchainsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateSharedSwapchainsKHR(VkDevice device, uint32_t swapchainCount, VkSwapchainCreateInfoKHR* pCreateInfos, VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchains) {
	return ptr_vkCreateSharedSwapchainsKHR(device, swapchainCount, pCreateInfos, pAllocator, pSwapchains);
}
*/
/*
PFN_vkCreateSwapchainKHR ptr_vkCreateSwapchainKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateSwapchainKHR(VkDevice device, VkSwapchainCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain) {
	return ptr_vkCreateSwapchainKHR(device, pCreateInfo, pAllocator, pSwapchain);
}
*/
/*
PFN_vkCreateTensorARM ptr_vkCreateTensorARM;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateTensorARM(VkDevice device, VkTensorCreateInfoARM* pCreateInfo, VkAllocationCallbacks* pAllocator, VkTensorARM* pTensor) {
	return ptr_vkCreateTensorARM(device, pCreateInfo, pAllocator, pTensor);
}
*/
/*
PFN_vkCreateTensorViewARM ptr_vkCreateTensorViewARM;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateTensorViewARM(VkDevice device, VkTensorViewCreateInfoARM* pCreateInfo, VkAllocationCallbacks* pAllocator, VkTensorViewARM* pView) {
	return ptr_vkCreateTensorViewARM(device, pCreateInfo, pAllocator, pView);
}
*/
/*
PFN_vkCreateValidationCacheEXT ptr_vkCreateValidationCacheEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateValidationCacheEXT(VkDevice device, VkValidationCacheCreateInfoEXT* pCreateInfo, VkAllocationCallbacks* pAllocator, VkValidationCacheEXT* pValidationCache) {
	return ptr_vkCreateValidationCacheEXT(device, pCreateInfo, pAllocator, pValidationCache);
}
*/
/*
PFN_vkCreateVideoSessionKHR ptr_vkCreateVideoSessionKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateVideoSessionKHR(VkDevice device, VkVideoSessionCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkVideoSessionKHR* pVideoSession) {
	return ptr_vkCreateVideoSessionKHR(device, pCreateInfo, pAllocator, pVideoSession);
}
*/
/*
PFN_vkCreateVideoSessionParametersKHR ptr_vkCreateVideoSessionParametersKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateVideoSessionParametersKHR(VkDevice device, VkVideoSessionParametersCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkVideoSessionParametersKHR* pVideoSessionParameters) {
	return ptr_vkCreateVideoSessionParametersKHR(device, pCreateInfo, pAllocator, pVideoSessionParameters);
}
*/
/*
PFN_vkCreateWin32SurfaceKHR ptr_vkCreateWin32SurfaceKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateWin32SurfaceKHR(VkInstance instance, VkWin32SurfaceCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return ptr_vkCreateWin32SurfaceKHR(instance, pCreateInfo, pAllocator, pSurface);
}
*/
/*
PFN_vkDebugMarkerSetObjectNameEXT ptr_vkDebugMarkerSetObjectNameEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkDebugMarkerSetObjectNameEXT(VkDevice device, VkDebugMarkerObjectNameInfoEXT* pNameInfo) {
	return ptr_vkDebugMarkerSetObjectNameEXT(device, pNameInfo);
}
*/
/*
PFN_vkDebugMarkerSetObjectTagEXT ptr_vkDebugMarkerSetObjectTagEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkDebugMarkerSetObjectTagEXT(VkDevice device, VkDebugMarkerObjectTagInfoEXT* pTagInfo) {
	return ptr_vkDebugMarkerSetObjectTagEXT(device, pTagInfo);
}
*/
/*
PFN_vkDebugReportMessageEXT ptr_vkDebugReportMessageEXT;
VKAPI_ATTR void VKAPI_CALL vkDebugReportMessageEXT(VkInstance instance, VkDebugReportFlagsEXT flags, VkDebugReportObjectTypeEXT objectType, uint64_t object, size_t location, int32_t messageCode, char* pLayerPrefix, char* pMessage) {
	ptr_vkDebugReportMessageEXT(instance, flags, objectType, object, location, messageCode, pLayerPrefix, pMessage);
}
*/
/*
PFN_vkDeferredOperationJoinKHR ptr_vkDeferredOperationJoinKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkDeferredOperationJoinKHR(VkDevice device, VkDeferredOperationKHR operation) {
	return ptr_vkDeferredOperationJoinKHR(device, operation);
}
*/
/*
PFN_vkDestroyAccelerationStructureKHR ptr_vkDestroyAccelerationStructureKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroyAccelerationStructureKHR(VkDevice device, VkAccelerationStructureKHR accelerationStructure, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyAccelerationStructureKHR(device, accelerationStructure, pAllocator);
}
*/
/*
PFN_vkDestroyAccelerationStructureNV ptr_vkDestroyAccelerationStructureNV;
VKAPI_ATTR void VKAPI_CALL vkDestroyAccelerationStructureNV(VkDevice device, VkAccelerationStructureNV accelerationStructure, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyAccelerationStructureNV(device, accelerationStructure, pAllocator);
}
*/
/*
PFN_vkDestroyBuffer ptr_vkDestroyBuffer;
VKAPI_ATTR void VKAPI_CALL vkDestroyBuffer(VkDevice device, VkBuffer buffer, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyBuffer(device, buffer, pAllocator);
}
*/
/*
PFN_vkDestroyBufferView ptr_vkDestroyBufferView;
VKAPI_ATTR void VKAPI_CALL vkDestroyBufferView(VkDevice device, VkBufferView bufferView, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyBufferView(device, bufferView, pAllocator);
}
*/
/*
PFN_vkDestroyCommandPool ptr_vkDestroyCommandPool;
VKAPI_ATTR void VKAPI_CALL vkDestroyCommandPool(VkDevice device, VkCommandPool commandPool, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyCommandPool(device, commandPool, pAllocator);
}
*/
/*
PFN_vkDestroyCuFunctionNVX ptr_vkDestroyCuFunctionNVX;
VKAPI_ATTR void VKAPI_CALL vkDestroyCuFunctionNVX(VkDevice device, VkCuFunctionNVX function, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyCuFunctionNVX(device, function, pAllocator);
}
*/
/*
PFN_vkDestroyCuModuleNVX ptr_vkDestroyCuModuleNVX;
VKAPI_ATTR void VKAPI_CALL vkDestroyCuModuleNVX(VkDevice device, VkCuModuleNVX module, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyCuModuleNVX(device, module, pAllocator);
}
*/
/*
PFN_vkDestroyDataGraphPipelineSessionARM ptr_vkDestroyDataGraphPipelineSessionARM;
VKAPI_ATTR void VKAPI_CALL vkDestroyDataGraphPipelineSessionARM(VkDevice device, VkDataGraphPipelineSessionARM session, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDataGraphPipelineSessionARM(device, session, pAllocator);
}
*/
/*
PFN_vkDestroyDebugReportCallbackEXT ptr_vkDestroyDebugReportCallbackEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyDebugReportCallbackEXT(VkInstance instance, VkDebugReportCallbackEXT callback, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDebugReportCallbackEXT(instance, callback, pAllocator);
}
*/
/*
PFN_vkDestroyDebugUtilsMessengerEXT ptr_vkDestroyDebugUtilsMessengerEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyDebugUtilsMessengerEXT(VkInstance instance, VkDebugUtilsMessengerEXT messenger, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDebugUtilsMessengerEXT(instance, messenger, pAllocator);
}
*/
/*
PFN_vkDestroyDeferredOperationKHR ptr_vkDestroyDeferredOperationKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroyDeferredOperationKHR(VkDevice device, VkDeferredOperationKHR operation, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDeferredOperationKHR(device, operation, pAllocator);
}
*/
/*
PFN_vkDestroyDescriptorPool ptr_vkDestroyDescriptorPool;
VKAPI_ATTR void VKAPI_CALL vkDestroyDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDescriptorPool(device, descriptorPool, pAllocator);
}
*/
/*
PFN_vkDestroyDescriptorSetLayout ptr_vkDestroyDescriptorSetLayout;
VKAPI_ATTR void VKAPI_CALL vkDestroyDescriptorSetLayout(VkDevice device, VkDescriptorSetLayout descriptorSetLayout, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDescriptorSetLayout(device, descriptorSetLayout, pAllocator);
}
*/
/*
PFN_vkDestroyDescriptorUpdateTemplate ptr_vkDestroyDescriptorUpdateTemplate;
VKAPI_ATTR void VKAPI_CALL vkDestroyDescriptorUpdateTemplate(VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDescriptorUpdateTemplate(device, descriptorUpdateTemplate, pAllocator);
}
*/
/*
PFN_vkDestroyDevice ptr_vkDestroyDevice;
VKAPI_ATTR void VKAPI_CALL vkDestroyDevice(VkDevice device, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyDevice(device, pAllocator);
}
*/
/*
PFN_vkDestroyEvent ptr_vkDestroyEvent;
VKAPI_ATTR void VKAPI_CALL vkDestroyEvent(VkDevice device, VkEvent event, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyEvent(device, event, pAllocator);
}
*/
/*
PFN_vkDestroyExternalComputeQueueNV ptr_vkDestroyExternalComputeQueueNV;
VKAPI_ATTR void VKAPI_CALL vkDestroyExternalComputeQueueNV(VkDevice device, VkExternalComputeQueueNV externalQueue, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyExternalComputeQueueNV(device, externalQueue, pAllocator);
}
*/
/*
PFN_vkDestroyFence ptr_vkDestroyFence;
VKAPI_ATTR void VKAPI_CALL vkDestroyFence(VkDevice device, VkFence fence, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyFence(device, fence, pAllocator);
}
*/
/*
PFN_vkDestroyFramebuffer ptr_vkDestroyFramebuffer;
VKAPI_ATTR void VKAPI_CALL vkDestroyFramebuffer(VkDevice device, VkFramebuffer framebuffer, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyFramebuffer(device, framebuffer, pAllocator);
}
*/
/*
PFN_vkDestroyImage ptr_vkDestroyImage;
VKAPI_ATTR void VKAPI_CALL vkDestroyImage(VkDevice device, VkImage image, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyImage(device, image, pAllocator);
}
*/
/*
PFN_vkDestroyImageView ptr_vkDestroyImageView;
VKAPI_ATTR void VKAPI_CALL vkDestroyImageView(VkDevice device, VkImageView imageView, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyImageView(device, imageView, pAllocator);
}
*/
/*
PFN_vkDestroyIndirectCommandsLayoutEXT ptr_vkDestroyIndirectCommandsLayoutEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyIndirectCommandsLayoutEXT(VkDevice device, VkIndirectCommandsLayoutEXT indirectCommandsLayout, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyIndirectCommandsLayoutEXT(device, indirectCommandsLayout, pAllocator);
}
*/
/*
PFN_vkDestroyIndirectCommandsLayoutNV ptr_vkDestroyIndirectCommandsLayoutNV;
VKAPI_ATTR void VKAPI_CALL vkDestroyIndirectCommandsLayoutNV(VkDevice device, VkIndirectCommandsLayoutNV indirectCommandsLayout, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyIndirectCommandsLayoutNV(device, indirectCommandsLayout, pAllocator);
}
*/
/*
PFN_vkDestroyIndirectExecutionSetEXT ptr_vkDestroyIndirectExecutionSetEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyIndirectExecutionSetEXT(VkDevice device, VkIndirectExecutionSetEXT indirectExecutionSet, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyIndirectExecutionSetEXT(device, indirectExecutionSet, pAllocator);
}
*/
/*
PFN_vkDestroyInstance ptr_vkDestroyInstance;
VKAPI_ATTR void VKAPI_CALL vkDestroyInstance(VkInstance instance, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyInstance(instance, pAllocator);
}
*/
/*
PFN_vkDestroyMicromapEXT ptr_vkDestroyMicromapEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyMicromapEXT(VkDevice device, VkMicromapEXT micromap, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyMicromapEXT(device, micromap, pAllocator);
}
*/
/*
PFN_vkDestroyOpticalFlowSessionNV ptr_vkDestroyOpticalFlowSessionNV;
VKAPI_ATTR void VKAPI_CALL vkDestroyOpticalFlowSessionNV(VkDevice device, VkOpticalFlowSessionNV session, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyOpticalFlowSessionNV(device, session, pAllocator);
}
*/
/*
PFN_vkDestroyPipeline ptr_vkDestroyPipeline;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipeline(VkDevice device, VkPipeline pipeline, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyPipeline(device, pipeline, pAllocator);
}
*/
/*
PFN_vkDestroyPipelineBinaryKHR ptr_vkDestroyPipelineBinaryKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipelineBinaryKHR(VkDevice device, VkPipelineBinaryKHR pipelineBinary, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyPipelineBinaryKHR(device, pipelineBinary, pAllocator);
}
*/
/*
PFN_vkDestroyPipelineCache ptr_vkDestroyPipelineCache;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipelineCache(VkDevice device, VkPipelineCache pipelineCache, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyPipelineCache(device, pipelineCache, pAllocator);
}
*/
/*
PFN_vkDestroyPipelineLayout ptr_vkDestroyPipelineLayout;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipelineLayout(VkDevice device, VkPipelineLayout pipelineLayout, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyPipelineLayout(device, pipelineLayout, pAllocator);
}
*/
/*
PFN_vkDestroyPrivateDataSlot ptr_vkDestroyPrivateDataSlot;
VKAPI_ATTR void VKAPI_CALL vkDestroyPrivateDataSlot(VkDevice device, VkPrivateDataSlot privateDataSlot, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyPrivateDataSlot(device, privateDataSlot, pAllocator);
}
*/
/*
PFN_vkDestroyQueryPool ptr_vkDestroyQueryPool;
VKAPI_ATTR void VKAPI_CALL vkDestroyQueryPool(VkDevice device, VkQueryPool queryPool, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyQueryPool(device, queryPool, pAllocator);
}
*/
/*
PFN_vkDestroyRenderPass ptr_vkDestroyRenderPass;
VKAPI_ATTR void VKAPI_CALL vkDestroyRenderPass(VkDevice device, VkRenderPass renderPass, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyRenderPass(device, renderPass, pAllocator);
}
*/
/*
PFN_vkDestroySampler ptr_vkDestroySampler;
VKAPI_ATTR void VKAPI_CALL vkDestroySampler(VkDevice device, VkSampler sampler, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroySampler(device, sampler, pAllocator);
}
*/
/*
PFN_vkDestroySamplerYcbcrConversion ptr_vkDestroySamplerYcbcrConversion;
VKAPI_ATTR void VKAPI_CALL vkDestroySamplerYcbcrConversion(VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroySamplerYcbcrConversion(device, ycbcrConversion, pAllocator);
}
*/
/*
PFN_vkDestroySemaphore ptr_vkDestroySemaphore;
VKAPI_ATTR void VKAPI_CALL vkDestroySemaphore(VkDevice device, VkSemaphore semaphore, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroySemaphore(device, semaphore, pAllocator);
}
*/
/*
PFN_vkDestroyShaderEXT ptr_vkDestroyShaderEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyShaderEXT(VkDevice device, VkShaderEXT shader, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyShaderEXT(device, shader, pAllocator);
}
*/
/*
PFN_vkDestroyShaderModule ptr_vkDestroyShaderModule;
VKAPI_ATTR void VKAPI_CALL vkDestroyShaderModule(VkDevice device, VkShaderModule shaderModule, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyShaderModule(device, shaderModule, pAllocator);
}
*/
/*
PFN_vkDestroySurfaceKHR ptr_vkDestroySurfaceKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroySurfaceKHR(VkInstance instance, VkSurfaceKHR surface, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroySurfaceKHR(instance, surface, pAllocator);
}
*/
/*
PFN_vkDestroySwapchainKHR ptr_vkDestroySwapchainKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroySwapchainKHR(VkDevice device, VkSwapchainKHR swapchain, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroySwapchainKHR(device, swapchain, pAllocator);
}
*/
/*
PFN_vkDestroyTensorARM ptr_vkDestroyTensorARM;
VKAPI_ATTR void VKAPI_CALL vkDestroyTensorARM(VkDevice device, VkTensorARM tensor, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyTensorARM(device, tensor, pAllocator);
}
*/
/*
PFN_vkDestroyTensorViewARM ptr_vkDestroyTensorViewARM;
VKAPI_ATTR void VKAPI_CALL vkDestroyTensorViewARM(VkDevice device, VkTensorViewARM tensorView, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyTensorViewARM(device, tensorView, pAllocator);
}
*/
/*
PFN_vkDestroyValidationCacheEXT ptr_vkDestroyValidationCacheEXT;
VKAPI_ATTR void VKAPI_CALL vkDestroyValidationCacheEXT(VkDevice device, VkValidationCacheEXT validationCache, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyValidationCacheEXT(device, validationCache, pAllocator);
}
*/
/*
PFN_vkDestroyVideoSessionKHR ptr_vkDestroyVideoSessionKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroyVideoSessionKHR(VkDevice device, VkVideoSessionKHR videoSession, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyVideoSessionKHR(device, videoSession, pAllocator);
}
*/
/*
PFN_vkDestroyVideoSessionParametersKHR ptr_vkDestroyVideoSessionParametersKHR;
VKAPI_ATTR void VKAPI_CALL vkDestroyVideoSessionParametersKHR(VkDevice device, VkVideoSessionParametersKHR videoSessionParameters, VkAllocationCallbacks* pAllocator) {
	ptr_vkDestroyVideoSessionParametersKHR(device, videoSessionParameters, pAllocator);
}
*/
/*
PFN_vkDeviceWaitIdle ptr_vkDeviceWaitIdle;
VKAPI_ATTR VkResult VKAPI_CALL vkDeviceWaitIdle(VkDevice device) {
	return ptr_vkDeviceWaitIdle(device);
}
*/
/*
PFN_vkDisplayPowerControlEXT ptr_vkDisplayPowerControlEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkDisplayPowerControlEXT(VkDevice device, VkDisplayKHR display, VkDisplayPowerInfoEXT* pDisplayPowerInfo) {
	return ptr_vkDisplayPowerControlEXT(device, display, pDisplayPowerInfo);
}
*/
/*
PFN_vkEndCommandBuffer ptr_vkEndCommandBuffer;
VKAPI_ATTR VkResult VKAPI_CALL vkEndCommandBuffer(VkCommandBuffer commandBuffer) {
	return ptr_vkEndCommandBuffer(commandBuffer);
}
*/
/*
PFN_vkEnumerateDeviceExtensionProperties ptr_vkEnumerateDeviceExtensionProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumerateDeviceExtensionProperties(VkPhysicalDevice physicalDevice, char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return ptr_vkEnumerateDeviceExtensionProperties(physicalDevice, pLayerName, pPropertyCount, pProperties);
}
*/
/*
PFN_vkEnumerateDeviceLayerProperties ptr_vkEnumerateDeviceLayerProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumerateDeviceLayerProperties(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
	return ptr_vkEnumerateDeviceLayerProperties(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkEnumerateInstanceExtensionProperties ptr_vkEnumerateInstanceExtensionProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumerateInstanceExtensionProperties(char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
	return ptr_vkEnumerateInstanceExtensionProperties(pLayerName, pPropertyCount, pProperties);
}
*/
/*
PFN_vkEnumerateInstanceLayerProperties ptr_vkEnumerateInstanceLayerProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumerateInstanceLayerProperties(uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
	return ptr_vkEnumerateInstanceLayerProperties(pPropertyCount, pProperties);
}
*/
/*
PFN_vkEnumerateInstanceVersion ptr_vkEnumerateInstanceVersion;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumerateInstanceVersion(uint32_t* pApiVersion) {
	return ptr_vkEnumerateInstanceVersion(pApiVersion);
}
*/
/*
PFN_vkEnumeratePhysicalDeviceGroups ptr_vkEnumeratePhysicalDeviceGroups;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumeratePhysicalDeviceGroups(VkInstance instance, uint32_t* pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupProperties* pPhysicalDeviceGroupProperties) {
	return ptr_vkEnumeratePhysicalDeviceGroups(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
}
*/
/*
PFN_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR ptr_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, uint32_t* pCounterCount, VkPerformanceCounterKHR* pCounters, VkPerformanceCounterDescriptionKHR* pCounterDescriptions) {
	return ptr_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(physicalDevice, queueFamilyIndex, pCounterCount, pCounters, pCounterDescriptions);
}
*/
/*
PFN_vkEnumeratePhysicalDevices ptr_vkEnumeratePhysicalDevices;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumeratePhysicalDevices(VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices) {
	return ptr_vkEnumeratePhysicalDevices(instance, pPhysicalDeviceCount, pPhysicalDevices);
}
*/
/*
PFN_vkExportMetalObjectsEXT ptr_vkExportMetalObjectsEXT;
VKAPI_ATTR void VKAPI_CALL vkExportMetalObjectsEXT(VkDevice device, VkExportMetalObjectsInfoEXT* pMetalObjectsInfo) {
	ptr_vkExportMetalObjectsEXT(device, pMetalObjectsInfo);
}
*/
/*
PFN_vkFlushMappedMemoryRanges ptr_vkFlushMappedMemoryRanges;
VKAPI_ATTR VkResult VKAPI_CALL vkFlushMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, VkMappedMemoryRange* pMemoryRanges) {
	return ptr_vkFlushMappedMemoryRanges(device, memoryRangeCount, pMemoryRanges);
}
*/
/*
PFN_vkFreeCommandBuffers ptr_vkFreeCommandBuffers;
VKAPI_ATTR void VKAPI_CALL vkFreeCommandBuffers(VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, VkCommandBuffer* pCommandBuffers) {
	ptr_vkFreeCommandBuffers(device, commandPool, commandBufferCount, pCommandBuffers);
}
*/
/*
PFN_vkFreeDescriptorSets ptr_vkFreeDescriptorSets;
VKAPI_ATTR VkResult VKAPI_CALL vkFreeDescriptorSets(VkDevice device, VkDescriptorPool descriptorPool, uint32_t descriptorSetCount, VkDescriptorSet* pDescriptorSets) {
	return ptr_vkFreeDescriptorSets(device, descriptorPool, descriptorSetCount, pDescriptorSets);
}
*/
/*
PFN_vkFreeMemory ptr_vkFreeMemory;
VKAPI_ATTR void VKAPI_CALL vkFreeMemory(VkDevice device, VkDeviceMemory memory, VkAllocationCallbacks* pAllocator) {
	ptr_vkFreeMemory(device, memory, pAllocator);
}
*/
/*
PFN_vkGetAccelerationStructureBuildSizesKHR ptr_vkGetAccelerationStructureBuildSizesKHR;
VKAPI_ATTR void VKAPI_CALL vkGetAccelerationStructureBuildSizesKHR(VkDevice device, VkAccelerationStructureBuildTypeKHR buildType, VkAccelerationStructureBuildGeometryInfoKHR* pBuildInfo, uint32_t* pMaxPrimitiveCounts, VkAccelerationStructureBuildSizesInfoKHR* pSizeInfo) {
	ptr_vkGetAccelerationStructureBuildSizesKHR(device, buildType, pBuildInfo, pMaxPrimitiveCounts, pSizeInfo);
}
*/
/*
PFN_vkGetAccelerationStructureDeviceAddressKHR ptr_vkGetAccelerationStructureDeviceAddressKHR;
VKAPI_ATTR VkDeviceAddress VKAPI_CALL vkGetAccelerationStructureDeviceAddressKHR(VkDevice device, VkAccelerationStructureDeviceAddressInfoKHR* pInfo) {
	return ptr_vkGetAccelerationStructureDeviceAddressKHR(device, pInfo);
}
*/
/*
PFN_vkGetAccelerationStructureHandleNV ptr_vkGetAccelerationStructureHandleNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetAccelerationStructureHandleNV(VkDevice device, VkAccelerationStructureNV accelerationStructure, size_t dataSize, void* pData) {
	return ptr_vkGetAccelerationStructureHandleNV(device, accelerationStructure, dataSize, pData);
}
*/
/*
PFN_vkGetAccelerationStructureMemoryRequirementsNV ptr_vkGetAccelerationStructureMemoryRequirementsNV;
VKAPI_ATTR void VKAPI_CALL vkGetAccelerationStructureMemoryRequirementsNV(VkDevice device, VkAccelerationStructureMemoryRequirementsInfoNV* pInfo, VkMemoryRequirements2KHR* pMemoryRequirements) {
	ptr_vkGetAccelerationStructureMemoryRequirementsNV(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT ptr_vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(VkDevice device, VkAccelerationStructureCaptureDescriptorDataInfoEXT* pInfo, void* pData) {
	return ptr_vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(device, pInfo, pData);
}
*/
/*
PFN_vkGetBufferDeviceAddress ptr_vkGetBufferDeviceAddress;
VKAPI_ATTR VkDeviceAddress VKAPI_CALL vkGetBufferDeviceAddress(VkDevice device, VkBufferDeviceAddressInfo* pInfo) {
	return ptr_vkGetBufferDeviceAddress(device, pInfo);
}
*/
/*
PFN_vkGetBufferMemoryRequirements ptr_vkGetBufferMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetBufferMemoryRequirements(VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements) {
	ptr_vkGetBufferMemoryRequirements(device, buffer, pMemoryRequirements);
}
*/
/*
PFN_vkGetBufferMemoryRequirements2 ptr_vkGetBufferMemoryRequirements2;
VKAPI_ATTR void VKAPI_CALL vkGetBufferMemoryRequirements2(VkDevice device, VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetBufferMemoryRequirements2(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetBufferOpaqueCaptureAddress ptr_vkGetBufferOpaqueCaptureAddress;
VKAPI_ATTR uint64_t VKAPI_CALL vkGetBufferOpaqueCaptureAddress(VkDevice device, VkBufferDeviceAddressInfo* pInfo) {
	return ptr_vkGetBufferOpaqueCaptureAddress(device, pInfo);
}
*/
/*
PFN_vkGetBufferOpaqueCaptureDescriptorDataEXT ptr_vkGetBufferOpaqueCaptureDescriptorDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetBufferOpaqueCaptureDescriptorDataEXT(VkDevice device, VkBufferCaptureDescriptorDataInfoEXT* pInfo, void* pData) {
	return ptr_vkGetBufferOpaqueCaptureDescriptorDataEXT(device, pInfo, pData);
}
*/
/*
PFN_vkGetCalibratedTimestampsKHR ptr_vkGetCalibratedTimestampsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetCalibratedTimestampsKHR(VkDevice device, uint32_t timestampCount, VkCalibratedTimestampInfoKHR* pTimestampInfos, uint64_t* pTimestamps, uint64_t* pMaxDeviation) {
	return ptr_vkGetCalibratedTimestampsKHR(device, timestampCount, pTimestampInfos, pTimestamps, pMaxDeviation);
}
*/
/*
PFN_vkGetClusterAccelerationStructureBuildSizesNV ptr_vkGetClusterAccelerationStructureBuildSizesNV;
VKAPI_ATTR void VKAPI_CALL vkGetClusterAccelerationStructureBuildSizesNV(VkDevice device, VkClusterAccelerationStructureInputInfoNV* pInfo, VkAccelerationStructureBuildSizesInfoKHR* pSizeInfo) {
	ptr_vkGetClusterAccelerationStructureBuildSizesNV(device, pInfo, pSizeInfo);
}
*/
/*
PFN_vkGetDataGraphPipelineAvailablePropertiesARM ptr_vkGetDataGraphPipelineAvailablePropertiesARM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDataGraphPipelineAvailablePropertiesARM(VkDevice device, VkDataGraphPipelineInfoARM* pPipelineInfo, uint32_t* pPropertiesCount, VkDataGraphPipelinePropertyARM* pProperties) {
	return ptr_vkGetDataGraphPipelineAvailablePropertiesARM(device, pPipelineInfo, pPropertiesCount, pProperties);
}
*/
/*
PFN_vkGetDataGraphPipelinePropertiesARM ptr_vkGetDataGraphPipelinePropertiesARM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDataGraphPipelinePropertiesARM(VkDevice device, VkDataGraphPipelineInfoARM* pPipelineInfo, uint32_t propertiesCount, VkDataGraphPipelinePropertyQueryResultARM* pProperties) {
	return ptr_vkGetDataGraphPipelinePropertiesARM(device, pPipelineInfo, propertiesCount, pProperties);
}
*/
/*
PFN_vkGetDataGraphPipelineSessionBindPointRequirementsARM ptr_vkGetDataGraphPipelineSessionBindPointRequirementsARM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDataGraphPipelineSessionBindPointRequirementsARM(VkDevice device, VkDataGraphPipelineSessionBindPointRequirementsInfoARM* pInfo, uint32_t* pBindPointRequirementCount, VkDataGraphPipelineSessionBindPointRequirementARM* pBindPointRequirements) {
	return ptr_vkGetDataGraphPipelineSessionBindPointRequirementsARM(device, pInfo, pBindPointRequirementCount, pBindPointRequirements);
}
*/
/*
PFN_vkGetDataGraphPipelineSessionMemoryRequirementsARM ptr_vkGetDataGraphPipelineSessionMemoryRequirementsARM;
VKAPI_ATTR void VKAPI_CALL vkGetDataGraphPipelineSessionMemoryRequirementsARM(VkDevice device, VkDataGraphPipelineSessionMemoryRequirementsInfoARM* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetDataGraphPipelineSessionMemoryRequirementsARM(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetDeferredOperationMaxConcurrencyKHR ptr_vkGetDeferredOperationMaxConcurrencyKHR;
VKAPI_ATTR uint32_t VKAPI_CALL vkGetDeferredOperationMaxConcurrencyKHR(VkDevice device, VkDeferredOperationKHR operation) {
	return ptr_vkGetDeferredOperationMaxConcurrencyKHR(device, operation);
}
*/
/*
PFN_vkGetDeferredOperationResultKHR ptr_vkGetDeferredOperationResultKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDeferredOperationResultKHR(VkDevice device, VkDeferredOperationKHR operation) {
	return ptr_vkGetDeferredOperationResultKHR(device, operation);
}
*/
/*
PFN_vkGetDescriptorEXT ptr_vkGetDescriptorEXT;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorEXT(VkDevice device, VkDescriptorGetInfoEXT* pDescriptorInfo, size_t dataSize, void* pDescriptor) {
	ptr_vkGetDescriptorEXT(device, pDescriptorInfo, dataSize, pDescriptor);
}
*/
/*
PFN_vkGetDescriptorSetHostMappingVALVE ptr_vkGetDescriptorSetHostMappingVALVE;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorSetHostMappingVALVE(VkDevice device, VkDescriptorSet descriptorSet, void** ppData) {
	ptr_vkGetDescriptorSetHostMappingVALVE(device, descriptorSet, ppData);
}
*/
/*
PFN_vkGetDescriptorSetLayoutBindingOffsetEXT ptr_vkGetDescriptorSetLayoutBindingOffsetEXT;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorSetLayoutBindingOffsetEXT(VkDevice device, VkDescriptorSetLayout layout, uint32_t binding, VkDeviceSize* pOffset) {
	ptr_vkGetDescriptorSetLayoutBindingOffsetEXT(device, layout, binding, pOffset);
}
*/
/*
PFN_vkGetDescriptorSetLayoutHostMappingInfoVALVE ptr_vkGetDescriptorSetLayoutHostMappingInfoVALVE;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorSetLayoutHostMappingInfoVALVE(VkDevice device, VkDescriptorSetBindingReferenceVALVE* pBindingReference, VkDescriptorSetLayoutHostMappingInfoVALVE* pHostMapping) {
	ptr_vkGetDescriptorSetLayoutHostMappingInfoVALVE(device, pBindingReference, pHostMapping);
}
*/
/*
PFN_vkGetDescriptorSetLayoutSizeEXT ptr_vkGetDescriptorSetLayoutSizeEXT;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorSetLayoutSizeEXT(VkDevice device, VkDescriptorSetLayout layout, VkDeviceSize* pLayoutSizeInBytes) {
	ptr_vkGetDescriptorSetLayoutSizeEXT(device, layout, pLayoutSizeInBytes);
}
*/
/*
PFN_vkGetDescriptorSetLayoutSupport ptr_vkGetDescriptorSetLayoutSupport;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorSetLayoutSupport(VkDevice device, VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport) {
	ptr_vkGetDescriptorSetLayoutSupport(device, pCreateInfo, pSupport);
}
*/
/*
PFN_vkGetDeviceAccelerationStructureCompatibilityKHR ptr_vkGetDeviceAccelerationStructureCompatibilityKHR;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceAccelerationStructureCompatibilityKHR(VkDevice device, VkAccelerationStructureVersionInfoKHR* pVersionInfo, VkAccelerationStructureCompatibilityKHR* pCompatibility) {
	ptr_vkGetDeviceAccelerationStructureCompatibilityKHR(device, pVersionInfo, pCompatibility);
}
*/
/*
PFN_vkGetDeviceBufferMemoryRequirements ptr_vkGetDeviceBufferMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceBufferMemoryRequirements(VkDevice device, VkDeviceBufferMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetDeviceBufferMemoryRequirements(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetDeviceFaultInfoEXT ptr_vkGetDeviceFaultInfoEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDeviceFaultInfoEXT(VkDevice device, VkDeviceFaultCountsEXT* pFaultCounts, VkDeviceFaultInfoEXT* pFaultInfo) {
	return ptr_vkGetDeviceFaultInfoEXT(device, pFaultCounts, pFaultInfo);
}
*/
/*
PFN_vkGetDeviceGroupPeerMemoryFeatures ptr_vkGetDeviceGroupPeerMemoryFeatures;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceGroupPeerMemoryFeatures(VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures) {
	ptr_vkGetDeviceGroupPeerMemoryFeatures(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}
*/
/*
PFN_vkGetDeviceGroupPresentCapabilitiesKHR ptr_vkGetDeviceGroupPresentCapabilitiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDeviceGroupPresentCapabilitiesKHR(VkDevice device, VkDeviceGroupPresentCapabilitiesKHR* pDeviceGroupPresentCapabilities) {
	return ptr_vkGetDeviceGroupPresentCapabilitiesKHR(device, pDeviceGroupPresentCapabilities);
}
*/
/*
PFN_vkGetDeviceGroupSurfacePresentModes2EXT ptr_vkGetDeviceGroupSurfacePresentModes2EXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDeviceGroupSurfacePresentModes2EXT(VkDevice device, VkPhysicalDeviceSurfaceInfo2KHR* pSurfaceInfo, VkDeviceGroupPresentModeFlagsKHR* pModes) {
	return ptr_vkGetDeviceGroupSurfacePresentModes2EXT(device, pSurfaceInfo, pModes);
}
*/
/*
PFN_vkGetDeviceGroupSurfacePresentModesKHR ptr_vkGetDeviceGroupSurfacePresentModesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDeviceGroupSurfacePresentModesKHR(VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR* pModes) {
	return ptr_vkGetDeviceGroupSurfacePresentModesKHR(device, surface, pModes);
}
*/
/*
PFN_vkGetDeviceImageMemoryRequirements ptr_vkGetDeviceImageMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceImageMemoryRequirements(VkDevice device, VkDeviceImageMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetDeviceImageMemoryRequirements(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetDeviceImageSparseMemoryRequirements ptr_vkGetDeviceImageSparseMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceImageSparseMemoryRequirements(VkDevice device, VkDeviceImageMemoryRequirements* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	ptr_vkGetDeviceImageSparseMemoryRequirements(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
*/
/*
PFN_vkGetDeviceImageSubresourceLayout ptr_vkGetDeviceImageSubresourceLayout;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceImageSubresourceLayout(VkDevice device, VkDeviceImageSubresourceInfo* pInfo, VkSubresourceLayout2* pLayout) {
	ptr_vkGetDeviceImageSubresourceLayout(device, pInfo, pLayout);
}
*/
/*
PFN_vkGetDeviceMemoryCommitment ptr_vkGetDeviceMemoryCommitment;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceMemoryCommitment(VkDevice device, VkDeviceMemory memory, VkDeviceSize* pCommittedMemoryInBytes) {
	ptr_vkGetDeviceMemoryCommitment(device, memory, pCommittedMemoryInBytes);
}
*/
/*
PFN_vkGetDeviceMemoryOpaqueCaptureAddress ptr_vkGetDeviceMemoryOpaqueCaptureAddress;
VKAPI_ATTR uint64_t VKAPI_CALL vkGetDeviceMemoryOpaqueCaptureAddress(VkDevice device, VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo) {
	return ptr_vkGetDeviceMemoryOpaqueCaptureAddress(device, pInfo);
}
*/
/*
PFN_vkGetDeviceMicromapCompatibilityEXT ptr_vkGetDeviceMicromapCompatibilityEXT;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceMicromapCompatibilityEXT(VkDevice device, VkMicromapVersionInfoEXT* pVersionInfo, VkAccelerationStructureCompatibilityKHR* pCompatibility) {
	ptr_vkGetDeviceMicromapCompatibilityEXT(device, pVersionInfo, pCompatibility);
}
*/
/*
PFN_vkGetDeviceProcAddr ptr_vkGetDeviceProcAddr;
VKAPI_ATTR PFN_vkVoidFunction VKAPI_CALL vkGetDeviceProcAddr(VkDevice device, char* pName) {
	return ptr_vkGetDeviceProcAddr(device, pName);
}
*/
/*
PFN_vkGetDeviceQueue ptr_vkGetDeviceQueue;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceQueue(VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue) {
	ptr_vkGetDeviceQueue(device, queueFamilyIndex, queueIndex, pQueue);
}
*/
/*
PFN_vkGetDeviceQueue2 ptr_vkGetDeviceQueue2;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceQueue2(VkDevice device, VkDeviceQueueInfo2* pQueueInfo, VkQueue* pQueue) {
	ptr_vkGetDeviceQueue2(device, pQueueInfo, pQueue);
}
*/
/*
PFN_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI ptr_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(VkDevice device, VkRenderPass renderpass, VkExtent2D* pMaxWorkgroupSize) {
	return ptr_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(device, renderpass, pMaxWorkgroupSize);
}
*/
/*
PFN_vkGetDeviceTensorMemoryRequirementsARM ptr_vkGetDeviceTensorMemoryRequirementsARM;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceTensorMemoryRequirementsARM(VkDevice device, VkDeviceTensorMemoryRequirementsARM* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetDeviceTensorMemoryRequirementsARM(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetDisplayModeProperties2KHR ptr_vkGetDisplayModeProperties2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDisplayModeProperties2KHR(VkPhysicalDevice physicalDevice, VkDisplayKHR display, uint32_t* pPropertyCount, VkDisplayModeProperties2KHR* pProperties) {
	return ptr_vkGetDisplayModeProperties2KHR(physicalDevice, display, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetDisplayModePropertiesKHR ptr_vkGetDisplayModePropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDisplayModePropertiesKHR(VkPhysicalDevice physicalDevice, VkDisplayKHR display, uint32_t* pPropertyCount, VkDisplayModePropertiesKHR* pProperties) {
	return ptr_vkGetDisplayModePropertiesKHR(physicalDevice, display, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetDisplayPlaneCapabilities2KHR ptr_vkGetDisplayPlaneCapabilities2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDisplayPlaneCapabilities2KHR(VkPhysicalDevice physicalDevice, VkDisplayPlaneInfo2KHR* pDisplayPlaneInfo, VkDisplayPlaneCapabilities2KHR* pCapabilities) {
	return ptr_vkGetDisplayPlaneCapabilities2KHR(physicalDevice, pDisplayPlaneInfo, pCapabilities);
}
*/
/*
PFN_vkGetDisplayPlaneCapabilitiesKHR ptr_vkGetDisplayPlaneCapabilitiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDisplayPlaneCapabilitiesKHR(VkPhysicalDevice physicalDevice, VkDisplayModeKHR mode, uint32_t planeIndex, VkDisplayPlaneCapabilitiesKHR* pCapabilities) {
	return ptr_vkGetDisplayPlaneCapabilitiesKHR(physicalDevice, mode, planeIndex, pCapabilities);
}
*/
/*
PFN_vkGetDisplayPlaneSupportedDisplaysKHR ptr_vkGetDisplayPlaneSupportedDisplaysKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDisplayPlaneSupportedDisplaysKHR(VkPhysicalDevice physicalDevice, uint32_t planeIndex, uint32_t* pDisplayCount, VkDisplayKHR* pDisplays) {
	return ptr_vkGetDisplayPlaneSupportedDisplaysKHR(physicalDevice, planeIndex, pDisplayCount, pDisplays);
}
*/
/*
PFN_vkGetDrmDisplayEXT ptr_vkGetDrmDisplayEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDrmDisplayEXT(VkPhysicalDevice physicalDevice, int32_t drmFd, uint32_t connectorId, VkDisplayKHR* display) {
	return ptr_vkGetDrmDisplayEXT(physicalDevice, drmFd, connectorId, display);
}
*/
/*
PFN_vkGetDynamicRenderingTilePropertiesQCOM ptr_vkGetDynamicRenderingTilePropertiesQCOM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetDynamicRenderingTilePropertiesQCOM(VkDevice device, VkRenderingInfo* pRenderingInfo, VkTilePropertiesQCOM* pProperties) {
	return ptr_vkGetDynamicRenderingTilePropertiesQCOM(device, pRenderingInfo, pProperties);
}
*/
/*
PFN_vkGetEncodedVideoSessionParametersKHR ptr_vkGetEncodedVideoSessionParametersKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetEncodedVideoSessionParametersKHR(VkDevice device, VkVideoEncodeSessionParametersGetInfoKHR* pVideoSessionParametersInfo, VkVideoEncodeSessionParametersFeedbackInfoKHR* pFeedbackInfo, size_t* pDataSize, void* pData) {
	return ptr_vkGetEncodedVideoSessionParametersKHR(device, pVideoSessionParametersInfo, pFeedbackInfo, pDataSize, pData);
}
*/
/*
PFN_vkGetEventStatus ptr_vkGetEventStatus;
VKAPI_ATTR VkResult VKAPI_CALL vkGetEventStatus(VkDevice device, VkEvent event) {
	return ptr_vkGetEventStatus(device, event);
}
*/
/*
PFN_vkGetExternalComputeQueueDataNV ptr_vkGetExternalComputeQueueDataNV;
VKAPI_ATTR void VKAPI_CALL vkGetExternalComputeQueueDataNV(VkExternalComputeQueueNV externalQueue, VkExternalComputeQueueDataParamsNV* params, void* pData) {
	ptr_vkGetExternalComputeQueueDataNV(externalQueue, params, pData);
}
*/
/*
PFN_vkGetFenceFdKHR ptr_vkGetFenceFdKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetFenceFdKHR(VkDevice device, VkFenceGetFdInfoKHR* pGetFdInfo, int* pFd) {
	return ptr_vkGetFenceFdKHR(device, pGetFdInfo, pFd);
}
*/
/*
PFN_vkGetFenceStatus ptr_vkGetFenceStatus;
VKAPI_ATTR VkResult VKAPI_CALL vkGetFenceStatus(VkDevice device, VkFence fence) {
	return ptr_vkGetFenceStatus(device, fence);
}
*/
/*
PFN_vkGetFenceWin32HandleKHR ptr_vkGetFenceWin32HandleKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetFenceWin32HandleKHR(VkDevice device, VkFenceGetWin32HandleInfoKHR* pGetWin32HandleInfo, HANDLE* pHandle) {
	return ptr_vkGetFenceWin32HandleKHR(device, pGetWin32HandleInfo, pHandle);
}
*/
/*
PFN_vkGetFramebufferTilePropertiesQCOM ptr_vkGetFramebufferTilePropertiesQCOM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetFramebufferTilePropertiesQCOM(VkDevice device, VkFramebuffer framebuffer, uint32_t* pPropertiesCount, VkTilePropertiesQCOM* pProperties) {
	return ptr_vkGetFramebufferTilePropertiesQCOM(device, framebuffer, pPropertiesCount, pProperties);
}
*/
/*
PFN_vkGetGeneratedCommandsMemoryRequirementsEXT ptr_vkGetGeneratedCommandsMemoryRequirementsEXT;
VKAPI_ATTR void VKAPI_CALL vkGetGeneratedCommandsMemoryRequirementsEXT(VkDevice device, VkGeneratedCommandsMemoryRequirementsInfoEXT* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetGeneratedCommandsMemoryRequirementsEXT(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetGeneratedCommandsMemoryRequirementsNV ptr_vkGetGeneratedCommandsMemoryRequirementsNV;
VKAPI_ATTR void VKAPI_CALL vkGetGeneratedCommandsMemoryRequirementsNV(VkDevice device, VkGeneratedCommandsMemoryRequirementsInfoNV* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetGeneratedCommandsMemoryRequirementsNV(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetImageDrmFormatModifierPropertiesEXT ptr_vkGetImageDrmFormatModifierPropertiesEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetImageDrmFormatModifierPropertiesEXT(VkDevice device, VkImage image, VkImageDrmFormatModifierPropertiesEXT* pProperties) {
	return ptr_vkGetImageDrmFormatModifierPropertiesEXT(device, image, pProperties);
}
*/
/*
PFN_vkGetImageMemoryRequirements ptr_vkGetImageMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetImageMemoryRequirements(VkDevice device, VkImage image, VkMemoryRequirements* pMemoryRequirements) {
	ptr_vkGetImageMemoryRequirements(device, image, pMemoryRequirements);
}
*/
/*
PFN_vkGetImageMemoryRequirements2 ptr_vkGetImageMemoryRequirements2;
VKAPI_ATTR void VKAPI_CALL vkGetImageMemoryRequirements2(VkDevice device, VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetImageMemoryRequirements2(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetImageOpaqueCaptureDescriptorDataEXT ptr_vkGetImageOpaqueCaptureDescriptorDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetImageOpaqueCaptureDescriptorDataEXT(VkDevice device, VkImageCaptureDescriptorDataInfoEXT* pInfo, void* pData) {
	return ptr_vkGetImageOpaqueCaptureDescriptorDataEXT(device, pInfo, pData);
}
*/
/*
PFN_vkGetImageSparseMemoryRequirements ptr_vkGetImageSparseMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetImageSparseMemoryRequirements(VkDevice device, VkImage image, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements* pSparseMemoryRequirements) {
	ptr_vkGetImageSparseMemoryRequirements(device, image, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
*/
/*
PFN_vkGetImageSparseMemoryRequirements2 ptr_vkGetImageSparseMemoryRequirements2;
VKAPI_ATTR void VKAPI_CALL vkGetImageSparseMemoryRequirements2(VkDevice device, VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	ptr_vkGetImageSparseMemoryRequirements2(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
*/
/*
PFN_vkGetImageSubresourceLayout ptr_vkGetImageSubresourceLayout;
VKAPI_ATTR void VKAPI_CALL vkGetImageSubresourceLayout(VkDevice device, VkImage image, VkImageSubresource* pSubresource, VkSubresourceLayout* pLayout) {
	ptr_vkGetImageSubresourceLayout(device, image, pSubresource, pLayout);
}
*/
/*
PFN_vkGetImageSubresourceLayout2 ptr_vkGetImageSubresourceLayout2;
VKAPI_ATTR void VKAPI_CALL vkGetImageSubresourceLayout2(VkDevice device, VkImage image, VkImageSubresource2* pSubresource, VkSubresourceLayout2* pLayout) {
	ptr_vkGetImageSubresourceLayout2(device, image, pSubresource, pLayout);
}
*/
/*
PFN_vkGetImageViewAddressNVX ptr_vkGetImageViewAddressNVX;
VKAPI_ATTR VkResult VKAPI_CALL vkGetImageViewAddressNVX(VkDevice device, VkImageView imageView, VkImageViewAddressPropertiesNVX* pProperties) {
	return ptr_vkGetImageViewAddressNVX(device, imageView, pProperties);
}
*/
/*
PFN_vkGetImageViewHandle64NVX ptr_vkGetImageViewHandle64NVX;
VKAPI_ATTR uint64_t VKAPI_CALL vkGetImageViewHandle64NVX(VkDevice device, VkImageViewHandleInfoNVX* pInfo) {
	return ptr_vkGetImageViewHandle64NVX(device, pInfo);
}
*/
/*
PFN_vkGetImageViewHandleNVX ptr_vkGetImageViewHandleNVX;
VKAPI_ATTR uint32_t VKAPI_CALL vkGetImageViewHandleNVX(VkDevice device, VkImageViewHandleInfoNVX* pInfo) {
	return ptr_vkGetImageViewHandleNVX(device, pInfo);
}
*/
/*
PFN_vkGetImageViewOpaqueCaptureDescriptorDataEXT ptr_vkGetImageViewOpaqueCaptureDescriptorDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetImageViewOpaqueCaptureDescriptorDataEXT(VkDevice device, VkImageViewCaptureDescriptorDataInfoEXT* pInfo, void* pData) {
	return ptr_vkGetImageViewOpaqueCaptureDescriptorDataEXT(device, pInfo, pData);
}
*/
/*
PFN_vkGetInstanceProcAddr ptr_vkGetInstanceProcAddr;
VKAPI_ATTR PFN_vkVoidFunction VKAPI_CALL vkGetInstanceProcAddr(VkInstance instance, char* pName) {
	return ptr_vkGetInstanceProcAddr(instance, pName);
}
*/
/*
PFN_vkGetLatencyTimingsNV ptr_vkGetLatencyTimingsNV;
VKAPI_ATTR void VKAPI_CALL vkGetLatencyTimingsNV(VkDevice device, VkSwapchainKHR swapchain, VkGetLatencyMarkerInfoNV* pLatencyMarkerInfo) {
	ptr_vkGetLatencyTimingsNV(device, swapchain, pLatencyMarkerInfo);
}
*/
/*
PFN_vkGetMemoryFdKHR ptr_vkGetMemoryFdKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryFdKHR(VkDevice device, VkMemoryGetFdInfoKHR* pGetFdInfo, int* pFd) {
	return ptr_vkGetMemoryFdKHR(device, pGetFdInfo, pFd);
}
*/
/*
PFN_vkGetMemoryFdPropertiesKHR ptr_vkGetMemoryFdPropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryFdPropertiesKHR(VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, int fd, VkMemoryFdPropertiesKHR* pMemoryFdProperties) {
	return ptr_vkGetMemoryFdPropertiesKHR(device, handleType, fd, pMemoryFdProperties);
}
*/
/*
PFN_vkGetMemoryHostPointerPropertiesEXT ptr_vkGetMemoryHostPointerPropertiesEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryHostPointerPropertiesEXT(VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, void* pHostPointer, VkMemoryHostPointerPropertiesEXT* pMemoryHostPointerProperties) {
	return ptr_vkGetMemoryHostPointerPropertiesEXT(device, handleType, pHostPointer, pMemoryHostPointerProperties);
}
*/
/*
PFN_vkGetMemoryMetalHandleEXT ptr_vkGetMemoryMetalHandleEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryMetalHandleEXT(VkDevice device, VkMemoryGetMetalHandleInfoEXT* pGetMetalHandleInfo, void** pHandle) {
	return ptr_vkGetMemoryMetalHandleEXT(device, pGetMetalHandleInfo, pHandle);
}
*/
/*
PFN_vkGetMemoryMetalHandlePropertiesEXT ptr_vkGetMemoryMetalHandlePropertiesEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryMetalHandlePropertiesEXT(VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, void* pHandle, VkMemoryMetalHandlePropertiesEXT* pMemoryMetalHandleProperties) {
	return ptr_vkGetMemoryMetalHandlePropertiesEXT(device, handleType, pHandle, pMemoryMetalHandleProperties);
}
*/
/*
PFN_vkGetMemoryRemoteAddressNV ptr_vkGetMemoryRemoteAddressNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryRemoteAddressNV(VkDevice device, VkMemoryGetRemoteAddressInfoNV* pMemoryGetRemoteAddressInfo, VkRemoteAddressNV* pAddress) {
	return ptr_vkGetMemoryRemoteAddressNV(device, pMemoryGetRemoteAddressInfo, pAddress);
}
*/
/*
PFN_vkGetMemoryWin32HandleKHR ptr_vkGetMemoryWin32HandleKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryWin32HandleKHR(VkDevice device, VkMemoryGetWin32HandleInfoKHR* pGetWin32HandleInfo, HANDLE* pHandle) {
	return ptr_vkGetMemoryWin32HandleKHR(device, pGetWin32HandleInfo, pHandle);
}
*/
/*
PFN_vkGetMemoryWin32HandleNV ptr_vkGetMemoryWin32HandleNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryWin32HandleNV(VkDevice device, VkDeviceMemory memory, VkExternalMemoryHandleTypeFlagsNV handleType, HANDLE* pHandle) {
	return ptr_vkGetMemoryWin32HandleNV(device, memory, handleType, pHandle);
}
*/
/*
PFN_vkGetMemoryWin32HandlePropertiesKHR ptr_vkGetMemoryWin32HandlePropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetMemoryWin32HandlePropertiesKHR(VkDevice device, VkExternalMemoryHandleTypeFlagBits handleType, HANDLE handle, VkMemoryWin32HandlePropertiesKHR* pMemoryWin32HandleProperties) {
	return ptr_vkGetMemoryWin32HandlePropertiesKHR(device, handleType, handle, pMemoryWin32HandleProperties);
}
*/
/*
PFN_vkGetMicromapBuildSizesEXT ptr_vkGetMicromapBuildSizesEXT;
VKAPI_ATTR void VKAPI_CALL vkGetMicromapBuildSizesEXT(VkDevice device, VkAccelerationStructureBuildTypeKHR buildType, VkMicromapBuildInfoEXT* pBuildInfo, VkMicromapBuildSizesInfoEXT* pSizeInfo) {
	ptr_vkGetMicromapBuildSizesEXT(device, buildType, pBuildInfo, pSizeInfo);
}
*/
/*
PFN_vkGetPartitionedAccelerationStructuresBuildSizesNV ptr_vkGetPartitionedAccelerationStructuresBuildSizesNV;
VKAPI_ATTR void VKAPI_CALL vkGetPartitionedAccelerationStructuresBuildSizesNV(VkDevice device, VkPartitionedAccelerationStructureInstancesInputNV* pInfo, VkAccelerationStructureBuildSizesInfoKHR* pSizeInfo) {
	ptr_vkGetPartitionedAccelerationStructuresBuildSizesNV(device, pInfo, pSizeInfo);
}
*/
/*
PFN_vkGetPastPresentationTimingGOOGLE ptr_vkGetPastPresentationTimingGOOGLE;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPastPresentationTimingGOOGLE(VkDevice device, VkSwapchainKHR swapchain, uint32_t* pPresentationTimingCount, VkPastPresentationTimingGOOGLE* pPresentationTimings) {
	return ptr_vkGetPastPresentationTimingGOOGLE(device, swapchain, pPresentationTimingCount, pPresentationTimings);
}
*/
/*
PFN_vkGetPerformanceParameterINTEL ptr_vkGetPerformanceParameterINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPerformanceParameterINTEL(VkDevice device, VkPerformanceParameterTypeINTEL parameter, VkPerformanceValueINTEL* pValue) {
	return ptr_vkGetPerformanceParameterINTEL(device, parameter, pValue);
}
*/
/*
PFN_vkGetPhysicalDeviceCalibrateableTimeDomainsKHR ptr_vkGetPhysicalDeviceCalibrateableTimeDomainsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(VkPhysicalDevice physicalDevice, uint32_t* pTimeDomainCount, VkTimeDomainKHR* pTimeDomains) {
	return ptr_vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(physicalDevice, pTimeDomainCount, pTimeDomains);
}
*/
/*
PFN_vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV ptr_vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkCooperativeMatrixFlexibleDimensionsPropertiesNV* pProperties) {
	return ptr_vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR ptr_vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkCooperativeMatrixPropertiesKHR* pProperties) {
	return ptr_vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV ptr_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceCooperativeMatrixPropertiesNV(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkCooperativeMatrixPropertiesNV* pProperties) {
	return ptr_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceCooperativeVectorPropertiesNV ptr_vkGetPhysicalDeviceCooperativeVectorPropertiesNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceCooperativeVectorPropertiesNV(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkCooperativeVectorPropertiesNV* pProperties) {
	return ptr_vkGetPhysicalDeviceCooperativeVectorPropertiesNV(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR ptr_vkGetPhysicalDeviceDisplayPlaneProperties2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceDisplayPlaneProperties2KHR(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayPlaneProperties2KHR* pProperties) {
	return ptr_vkGetPhysicalDeviceDisplayPlaneProperties2KHR(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR ptr_vkGetPhysicalDeviceDisplayPlanePropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceDisplayPlanePropertiesKHR(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayPlanePropertiesKHR* pProperties) {
	return ptr_vkGetPhysicalDeviceDisplayPlanePropertiesKHR(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceDisplayProperties2KHR ptr_vkGetPhysicalDeviceDisplayProperties2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceDisplayProperties2KHR(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayProperties2KHR* pProperties) {
	return ptr_vkGetPhysicalDeviceDisplayProperties2KHR(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceDisplayPropertiesKHR ptr_vkGetPhysicalDeviceDisplayPropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceDisplayPropertiesKHR(VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkDisplayPropertiesKHR* pProperties) {
	return ptr_vkGetPhysicalDeviceDisplayPropertiesKHR(physicalDevice, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalBufferProperties ptr_vkGetPhysicalDeviceExternalBufferProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalBufferProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties) {
	ptr_vkGetPhysicalDeviceExternalBufferProperties(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalFenceProperties ptr_vkGetPhysicalDeviceExternalFenceProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalFenceProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties) {
	ptr_vkGetPhysicalDeviceExternalFenceProperties(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV ptr_vkGetPhysicalDeviceExternalImageFormatPropertiesNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceExternalImageFormatPropertiesNV(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkExternalMemoryHandleTypeFlagsNV externalHandleType, VkExternalImageFormatPropertiesNV* pExternalImageFormatProperties) {
	return ptr_vkGetPhysicalDeviceExternalImageFormatPropertiesNV(physicalDevice, format, type, tiling, usage, flags, externalHandleType, pExternalImageFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalSemaphoreProperties ptr_vkGetPhysicalDeviceExternalSemaphoreProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalSemaphoreProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties) {
	ptr_vkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalTensorPropertiesARM ptr_vkGetPhysicalDeviceExternalTensorPropertiesARM;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalTensorPropertiesARM(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalTensorInfoARM* pExternalTensorInfo, VkExternalTensorPropertiesARM* pExternalTensorProperties) {
	ptr_vkGetPhysicalDeviceExternalTensorPropertiesARM(physicalDevice, pExternalTensorInfo, pExternalTensorProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceFeatures ptr_vkGetPhysicalDeviceFeatures;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFeatures(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures) {
	ptr_vkGetPhysicalDeviceFeatures(physicalDevice, pFeatures);
}
*/
/*
PFN_vkGetPhysicalDeviceFeatures2 ptr_vkGetPhysicalDeviceFeatures2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFeatures2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures) {
	ptr_vkGetPhysicalDeviceFeatures2(physicalDevice, pFeatures);
}
*/
/*
PFN_vkGetPhysicalDeviceFormatProperties ptr_vkGetPhysicalDeviceFormatProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties* pFormatProperties) {
	ptr_vkGetPhysicalDeviceFormatProperties(physicalDevice, format, pFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceFormatProperties2 ptr_vkGetPhysicalDeviceFormatProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFormatProperties2(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties) {
	ptr_vkGetPhysicalDeviceFormatProperties2(physicalDevice, format, pFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceFragmentShadingRatesKHR ptr_vkGetPhysicalDeviceFragmentShadingRatesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceFragmentShadingRatesKHR(VkPhysicalDevice physicalDevice, uint32_t* pFragmentShadingRateCount, VkPhysicalDeviceFragmentShadingRateKHR* pFragmentShadingRates) {
	return ptr_vkGetPhysicalDeviceFragmentShadingRatesKHR(physicalDevice, pFragmentShadingRateCount, pFragmentShadingRates);
}
*/
/*
PFN_vkGetPhysicalDeviceImageFormatProperties ptr_vkGetPhysicalDeviceImageFormatProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkImageFormatProperties* pImageFormatProperties) {
	return ptr_vkGetPhysicalDeviceImageFormatProperties(physicalDevice, format, type, tiling, usage, flags, pImageFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceImageFormatProperties2 ptr_vkGetPhysicalDeviceImageFormatProperties2;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceImageFormatProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceImageFormatInfo2* pImageFormatInfo, VkImageFormatProperties2* pImageFormatProperties) {
	return ptr_vkGetPhysicalDeviceImageFormatProperties2(physicalDevice, pImageFormatInfo, pImageFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceMemoryProperties ptr_vkGetPhysicalDeviceMemoryProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceMemoryProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties) {
	ptr_vkGetPhysicalDeviceMemoryProperties(physicalDevice, pMemoryProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceMemoryProperties2 ptr_vkGetPhysicalDeviceMemoryProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceMemoryProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties) {
	ptr_vkGetPhysicalDeviceMemoryProperties2(physicalDevice, pMemoryProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceMultisamplePropertiesEXT ptr_vkGetPhysicalDeviceMultisamplePropertiesEXT;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceMultisamplePropertiesEXT(VkPhysicalDevice physicalDevice, VkSampleCountFlagBits samples, VkMultisamplePropertiesEXT* pMultisampleProperties) {
	ptr_vkGetPhysicalDeviceMultisamplePropertiesEXT(physicalDevice, samples, pMultisampleProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceOpticalFlowImageFormatsNV ptr_vkGetPhysicalDeviceOpticalFlowImageFormatsNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceOpticalFlowImageFormatsNV(VkPhysicalDevice physicalDevice, VkOpticalFlowImageFormatInfoNV* pOpticalFlowImageFormatInfo, uint32_t* pFormatCount, VkOpticalFlowImageFormatPropertiesNV* pImageFormatProperties) {
	return ptr_vkGetPhysicalDeviceOpticalFlowImageFormatsNV(physicalDevice, pOpticalFlowImageFormatInfo, pFormatCount, pImageFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDevicePresentRectanglesKHR ptr_vkGetPhysicalDevicePresentRectanglesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDevicePresentRectanglesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pRectCount, VkRect2D* pRects) {
	return ptr_vkGetPhysicalDevicePresentRectanglesKHR(physicalDevice, surface, pRectCount, pRects);
}
*/
/*
PFN_vkGetPhysicalDeviceProperties ptr_vkGetPhysicalDeviceProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties) {
	ptr_vkGetPhysicalDeviceProperties(physicalDevice, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceProperties2 ptr_vkGetPhysicalDeviceProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
	ptr_vkGetPhysicalDeviceProperties2(physicalDevice, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM ptr_vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM(VkPhysicalDevice physicalDevice, VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM* pQueueFamilyDataGraphProcessingEngineInfo, VkQueueFamilyDataGraphProcessingEnginePropertiesARM* pQueueFamilyDataGraphProcessingEngineProperties) {
	ptr_vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM(physicalDevice, pQueueFamilyDataGraphProcessingEngineInfo, pQueueFamilyDataGraphProcessingEngineProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM ptr_vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, uint32_t* pQueueFamilyDataGraphPropertyCount, VkQueueFamilyDataGraphPropertiesARM* pQueueFamilyDataGraphProperties) {
	return ptr_vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM(physicalDevice, queueFamilyIndex, pQueueFamilyDataGraphPropertyCount, pQueueFamilyDataGraphProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR ptr_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(VkPhysicalDevice physicalDevice, VkQueryPoolPerformanceCreateInfoKHR* pPerformanceQueryCreateInfo, uint32_t* pNumPasses) {
	ptr_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(physicalDevice, pPerformanceQueryCreateInfo, pNumPasses);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyProperties ptr_vkGetPhysicalDeviceQueueFamilyProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceQueueFamilyProperties(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties) {
	ptr_vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyProperties2 ptr_vkGetPhysicalDeviceQueueFamilyProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceQueueFamilyProperties2(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties) {
	ptr_vkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceSparseImageFormatProperties ptr_vkGetPhysicalDeviceSparseImageFormatProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceSparseImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkSampleCountFlagBits samples, VkImageUsageFlags usage, VkImageTiling tiling, uint32_t* pPropertyCount, VkSparseImageFormatProperties* pProperties) {
	ptr_vkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice, format, type, samples, usage, tiling, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceSparseImageFormatProperties2 ptr_vkGetPhysicalDeviceSparseImageFormatProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceSparseImageFormatProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties) {
	ptr_vkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV ptr_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(VkPhysicalDevice physicalDevice, uint32_t* pCombinationCount, VkFramebufferMixedSamplesCombinationNV* pCombinations) {
	return ptr_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(physicalDevice, pCombinationCount, pCombinations);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT ptr_vkGetPhysicalDeviceSurfaceCapabilities2EXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfaceCapabilities2EXT(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilities2EXT* pSurfaceCapabilities) {
	return ptr_vkGetPhysicalDeviceSurfaceCapabilities2EXT(physicalDevice, surface, pSurfaceCapabilities);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR ptr_vkGetPhysicalDeviceSurfaceCapabilities2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfaceCapabilities2KHR(VkPhysicalDevice physicalDevice, VkPhysicalDeviceSurfaceInfo2KHR* pSurfaceInfo, VkSurfaceCapabilities2KHR* pSurfaceCapabilities) {
	return ptr_vkGetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice, pSurfaceInfo, pSurfaceCapabilities);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR ptr_vkGetPhysicalDeviceSurfaceCapabilitiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfaceCapabilitiesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return ptr_vkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice, surface, pSurfaceCapabilities);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfaceFormats2KHR ptr_vkGetPhysicalDeviceSurfaceFormats2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfaceFormats2KHR(VkPhysicalDevice physicalDevice, VkPhysicalDeviceSurfaceInfo2KHR* pSurfaceInfo, uint32_t* pSurfaceFormatCount, VkSurfaceFormat2KHR* pSurfaceFormats) {
	return ptr_vkGetPhysicalDeviceSurfaceFormats2KHR(physicalDevice, pSurfaceInfo, pSurfaceFormatCount, pSurfaceFormats);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfaceFormatsKHR ptr_vkGetPhysicalDeviceSurfaceFormatsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfaceFormatsKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return ptr_vkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT ptr_vkGetPhysicalDeviceSurfacePresentModes2EXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfacePresentModes2EXT(VkPhysicalDevice physicalDevice, VkPhysicalDeviceSurfaceInfo2KHR* pSurfaceInfo, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return ptr_vkGetPhysicalDeviceSurfacePresentModes2EXT(physicalDevice, pSurfaceInfo, pPresentModeCount, pPresentModes);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfacePresentModesKHR ptr_vkGetPhysicalDeviceSurfacePresentModesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfacePresentModesKHR(VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return ptr_vkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice, surface, pPresentModeCount, pPresentModes);
}
*/
/*
PFN_vkGetPhysicalDeviceSurfaceSupportKHR ptr_vkGetPhysicalDeviceSurfaceSupportKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceSurfaceSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return ptr_vkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice, queueFamilyIndex, surface, pSupported);
}
*/
/*
PFN_vkGetPhysicalDeviceToolProperties ptr_vkGetPhysicalDeviceToolProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceToolProperties(VkPhysicalDevice physicalDevice, uint32_t* pToolCount, VkPhysicalDeviceToolProperties* pToolProperties) {
	return ptr_vkGetPhysicalDeviceToolProperties(physicalDevice, pToolCount, pToolProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceVideoCapabilitiesKHR ptr_vkGetPhysicalDeviceVideoCapabilitiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceVideoCapabilitiesKHR(VkPhysicalDevice physicalDevice, VkVideoProfileInfoKHR* pVideoProfile, VkVideoCapabilitiesKHR* pCapabilities) {
	return ptr_vkGetPhysicalDeviceVideoCapabilitiesKHR(physicalDevice, pVideoProfile, pCapabilities);
}
*/
/*
PFN_vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR ptr_vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR(VkPhysicalDevice physicalDevice, VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR* pQualityLevelInfo, VkVideoEncodeQualityLevelPropertiesKHR* pQualityLevelProperties) {
	return ptr_vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR(physicalDevice, pQualityLevelInfo, pQualityLevelProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceVideoFormatPropertiesKHR ptr_vkGetPhysicalDeviceVideoFormatPropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceVideoFormatPropertiesKHR(VkPhysicalDevice physicalDevice, VkPhysicalDeviceVideoFormatInfoKHR* pVideoFormatInfo, uint32_t* pVideoFormatPropertyCount, VkVideoFormatPropertiesKHR* pVideoFormatProperties) {
	return ptr_vkGetPhysicalDeviceVideoFormatPropertiesKHR(physicalDevice, pVideoFormatInfo, pVideoFormatPropertyCount, pVideoFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR ptr_vkGetPhysicalDeviceWin32PresentationSupportKHR;
VKAPI_ATTR VkBool32 VKAPI_CALL vkGetPhysicalDeviceWin32PresentationSupportKHR(VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex) {
	return ptr_vkGetPhysicalDeviceWin32PresentationSupportKHR(physicalDevice, queueFamilyIndex);
}
*/
/*
PFN_vkGetPipelineBinaryDataKHR ptr_vkGetPipelineBinaryDataKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineBinaryDataKHR(VkDevice device, VkPipelineBinaryDataInfoKHR* pInfo, VkPipelineBinaryKeyKHR* pPipelineBinaryKey, size_t* pPipelineBinaryDataSize, void* pPipelineBinaryData) {
	return ptr_vkGetPipelineBinaryDataKHR(device, pInfo, pPipelineBinaryKey, pPipelineBinaryDataSize, pPipelineBinaryData);
}
*/
/*
PFN_vkGetPipelineCacheData ptr_vkGetPipelineCacheData;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineCacheData(VkDevice device, VkPipelineCache pipelineCache, size_t* pDataSize, void* pData) {
	return ptr_vkGetPipelineCacheData(device, pipelineCache, pDataSize, pData);
}
*/
/*
PFN_vkGetPipelineExecutableInternalRepresentationsKHR ptr_vkGetPipelineExecutableInternalRepresentationsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineExecutableInternalRepresentationsKHR(VkDevice device, VkPipelineExecutableInfoKHR* pExecutableInfo, uint32_t* pInternalRepresentationCount, VkPipelineExecutableInternalRepresentationKHR* pInternalRepresentations) {
	return ptr_vkGetPipelineExecutableInternalRepresentationsKHR(device, pExecutableInfo, pInternalRepresentationCount, pInternalRepresentations);
}
*/
/*
PFN_vkGetPipelineExecutablePropertiesKHR ptr_vkGetPipelineExecutablePropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineExecutablePropertiesKHR(VkDevice device, VkPipelineInfoKHR* pPipelineInfo, uint32_t* pExecutableCount, VkPipelineExecutablePropertiesKHR* pProperties) {
	return ptr_vkGetPipelineExecutablePropertiesKHR(device, pPipelineInfo, pExecutableCount, pProperties);
}
*/
/*
PFN_vkGetPipelineExecutableStatisticsKHR ptr_vkGetPipelineExecutableStatisticsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineExecutableStatisticsKHR(VkDevice device, VkPipelineExecutableInfoKHR* pExecutableInfo, uint32_t* pStatisticCount, VkPipelineExecutableStatisticKHR* pStatistics) {
	return ptr_vkGetPipelineExecutableStatisticsKHR(device, pExecutableInfo, pStatisticCount, pStatistics);
}
*/
/*
PFN_vkGetPipelineIndirectDeviceAddressNV ptr_vkGetPipelineIndirectDeviceAddressNV;
VKAPI_ATTR VkDeviceAddress VKAPI_CALL vkGetPipelineIndirectDeviceAddressNV(VkDevice device, VkPipelineIndirectDeviceAddressInfoNV* pInfo) {
	return ptr_vkGetPipelineIndirectDeviceAddressNV(device, pInfo);
}
*/
/*
PFN_vkGetPipelineIndirectMemoryRequirementsNV ptr_vkGetPipelineIndirectMemoryRequirementsNV;
VKAPI_ATTR void VKAPI_CALL vkGetPipelineIndirectMemoryRequirementsNV(VkDevice device, VkComputePipelineCreateInfo* pCreateInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetPipelineIndirectMemoryRequirementsNV(device, pCreateInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetPipelineKeyKHR ptr_vkGetPipelineKeyKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineKeyKHR(VkDevice device, VkPipelineCreateInfoKHR* pPipelineCreateInfo, VkPipelineBinaryKeyKHR* pPipelineKey) {
	return ptr_vkGetPipelineKeyKHR(device, pPipelineCreateInfo, pPipelineKey);
}
*/
/*
PFN_vkGetPipelinePropertiesEXT ptr_vkGetPipelinePropertiesEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelinePropertiesEXT(VkDevice device, VkPipelineInfoEXT* pPipelineInfo, VkBaseOutStructure* pPipelineProperties) {
	return ptr_vkGetPipelinePropertiesEXT(device, pPipelineInfo, pPipelineProperties);
}
*/
/*
PFN_vkGetPrivateData ptr_vkGetPrivateData;
VKAPI_ATTR void VKAPI_CALL vkGetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t* pData) {
	ptr_vkGetPrivateData(device, objectType, objectHandle, privateDataSlot, pData);
}
*/
/*
PFN_vkGetQueryPoolResults ptr_vkGetQueryPoolResults;
VKAPI_ATTR VkResult VKAPI_CALL vkGetQueryPoolResults(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, size_t dataSize, void* pData, VkDeviceSize stride, VkQueryResultFlags flags) {
	return ptr_vkGetQueryPoolResults(device, queryPool, firstQuery, queryCount, dataSize, pData, stride, flags);
}
*/
/*
PFN_vkGetQueueCheckpointData2NV ptr_vkGetQueueCheckpointData2NV;
VKAPI_ATTR void VKAPI_CALL vkGetQueueCheckpointData2NV(VkQueue queue, uint32_t* pCheckpointDataCount, VkCheckpointData2NV* pCheckpointData) {
	ptr_vkGetQueueCheckpointData2NV(queue, pCheckpointDataCount, pCheckpointData);
}
*/
/*
PFN_vkGetQueueCheckpointDataNV ptr_vkGetQueueCheckpointDataNV;
VKAPI_ATTR void VKAPI_CALL vkGetQueueCheckpointDataNV(VkQueue queue, uint32_t* pCheckpointDataCount, VkCheckpointDataNV* pCheckpointData) {
	ptr_vkGetQueueCheckpointDataNV(queue, pCheckpointDataCount, pCheckpointData);
}
*/
/*
PFN_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR ptr_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(VkDevice device, VkPipeline pipeline, uint32_t firstGroup, uint32_t groupCount, size_t dataSize, void* pData) {
	return ptr_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(device, pipeline, firstGroup, groupCount, dataSize, pData);
}
*/
/*
PFN_vkGetRayTracingShaderGroupHandlesKHR ptr_vkGetRayTracingShaderGroupHandlesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetRayTracingShaderGroupHandlesKHR(VkDevice device, VkPipeline pipeline, uint32_t firstGroup, uint32_t groupCount, size_t dataSize, void* pData) {
	return ptr_vkGetRayTracingShaderGroupHandlesKHR(device, pipeline, firstGroup, groupCount, dataSize, pData);
}
*/
/*
PFN_vkGetRayTracingShaderGroupStackSizeKHR ptr_vkGetRayTracingShaderGroupStackSizeKHR;
VKAPI_ATTR VkDeviceSize VKAPI_CALL vkGetRayTracingShaderGroupStackSizeKHR(VkDevice device, VkPipeline pipeline, uint32_t group, VkShaderGroupShaderKHR groupShader) {
	return ptr_vkGetRayTracingShaderGroupStackSizeKHR(device, pipeline, group, groupShader);
}
*/
/*
PFN_vkGetRefreshCycleDurationGOOGLE ptr_vkGetRefreshCycleDurationGOOGLE;
VKAPI_ATTR VkResult VKAPI_CALL vkGetRefreshCycleDurationGOOGLE(VkDevice device, VkSwapchainKHR swapchain, VkRefreshCycleDurationGOOGLE* pDisplayTimingProperties) {
	return ptr_vkGetRefreshCycleDurationGOOGLE(device, swapchain, pDisplayTimingProperties);
}
*/
/*
PFN_vkGetRenderAreaGranularity ptr_vkGetRenderAreaGranularity;
VKAPI_ATTR void VKAPI_CALL vkGetRenderAreaGranularity(VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity) {
	ptr_vkGetRenderAreaGranularity(device, renderPass, pGranularity);
}
*/
/*
PFN_vkGetRenderingAreaGranularity ptr_vkGetRenderingAreaGranularity;
VKAPI_ATTR void VKAPI_CALL vkGetRenderingAreaGranularity(VkDevice device, VkRenderingAreaInfo* pRenderingAreaInfo, VkExtent2D* pGranularity) {
	ptr_vkGetRenderingAreaGranularity(device, pRenderingAreaInfo, pGranularity);
}
*/
/*
PFN_vkGetSamplerOpaqueCaptureDescriptorDataEXT ptr_vkGetSamplerOpaqueCaptureDescriptorDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSamplerOpaqueCaptureDescriptorDataEXT(VkDevice device, VkSamplerCaptureDescriptorDataInfoEXT* pInfo, void* pData) {
	return ptr_vkGetSamplerOpaqueCaptureDescriptorDataEXT(device, pInfo, pData);
}
*/
/*
PFN_vkGetSemaphoreCounterValue ptr_vkGetSemaphoreCounterValue;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSemaphoreCounterValue(VkDevice device, VkSemaphore semaphore, uint64_t* pValue) {
	return ptr_vkGetSemaphoreCounterValue(device, semaphore, pValue);
}
*/
/*
PFN_vkGetSemaphoreFdKHR ptr_vkGetSemaphoreFdKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSemaphoreFdKHR(VkDevice device, VkSemaphoreGetFdInfoKHR* pGetFdInfo, int* pFd) {
	return ptr_vkGetSemaphoreFdKHR(device, pGetFdInfo, pFd);
}
*/
/*
PFN_vkGetSemaphoreWin32HandleKHR ptr_vkGetSemaphoreWin32HandleKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSemaphoreWin32HandleKHR(VkDevice device, VkSemaphoreGetWin32HandleInfoKHR* pGetWin32HandleInfo, HANDLE* pHandle) {
	return ptr_vkGetSemaphoreWin32HandleKHR(device, pGetWin32HandleInfo, pHandle);
}
*/
/*
PFN_vkGetShaderBinaryDataEXT ptr_vkGetShaderBinaryDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetShaderBinaryDataEXT(VkDevice device, VkShaderEXT shader, size_t* pDataSize, void* pData) {
	return ptr_vkGetShaderBinaryDataEXT(device, shader, pDataSize, pData);
}
*/
/*
PFN_vkGetShaderInfoAMD ptr_vkGetShaderInfoAMD;
VKAPI_ATTR VkResult VKAPI_CALL vkGetShaderInfoAMD(VkDevice device, VkPipeline pipeline, VkShaderStageFlagBits shaderStage, VkShaderInfoTypeAMD infoType, size_t* pInfoSize, void* pInfo) {
	return ptr_vkGetShaderInfoAMD(device, pipeline, shaderStage, infoType, pInfoSize, pInfo);
}
*/
/*
PFN_vkGetShaderModuleCreateInfoIdentifierEXT ptr_vkGetShaderModuleCreateInfoIdentifierEXT;
VKAPI_ATTR void VKAPI_CALL vkGetShaderModuleCreateInfoIdentifierEXT(VkDevice device, VkShaderModuleCreateInfo* pCreateInfo, VkShaderModuleIdentifierEXT* pIdentifier) {
	ptr_vkGetShaderModuleCreateInfoIdentifierEXT(device, pCreateInfo, pIdentifier);
}
*/
/*
PFN_vkGetShaderModuleIdentifierEXT ptr_vkGetShaderModuleIdentifierEXT;
VKAPI_ATTR void VKAPI_CALL vkGetShaderModuleIdentifierEXT(VkDevice device, VkShaderModule shaderModule, VkShaderModuleIdentifierEXT* pIdentifier) {
	ptr_vkGetShaderModuleIdentifierEXT(device, shaderModule, pIdentifier);
}
*/
/*
PFN_vkGetSwapchainCounterEXT ptr_vkGetSwapchainCounterEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSwapchainCounterEXT(VkDevice device, VkSwapchainKHR swapchain, VkSurfaceCounterFlagBitsEXT counter, uint64_t* pCounterValue) {
	return ptr_vkGetSwapchainCounterEXT(device, swapchain, counter, pCounterValue);
}
*/
/*
PFN_vkGetSwapchainImagesKHR ptr_vkGetSwapchainImagesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSwapchainImagesKHR(VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages) {
	return ptr_vkGetSwapchainImagesKHR(device, swapchain, pSwapchainImageCount, pSwapchainImages);
}
*/
/*
PFN_vkGetSwapchainStatusKHR ptr_vkGetSwapchainStatusKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSwapchainStatusKHR(VkDevice device, VkSwapchainKHR swapchain) {
	return ptr_vkGetSwapchainStatusKHR(device, swapchain);
}
*/
/*
PFN_vkGetTensorMemoryRequirementsARM ptr_vkGetTensorMemoryRequirementsARM;
VKAPI_ATTR void VKAPI_CALL vkGetTensorMemoryRequirementsARM(VkDevice device, VkTensorMemoryRequirementsInfoARM* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	ptr_vkGetTensorMemoryRequirementsARM(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetTensorOpaqueCaptureDescriptorDataARM ptr_vkGetTensorOpaqueCaptureDescriptorDataARM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetTensorOpaqueCaptureDescriptorDataARM(VkDevice device, VkTensorCaptureDescriptorDataInfoARM* pInfo, void* pData) {
	return ptr_vkGetTensorOpaqueCaptureDescriptorDataARM(device, pInfo, pData);
}
*/
/*
PFN_vkGetTensorViewOpaqueCaptureDescriptorDataARM ptr_vkGetTensorViewOpaqueCaptureDescriptorDataARM;
VKAPI_ATTR VkResult VKAPI_CALL vkGetTensorViewOpaqueCaptureDescriptorDataARM(VkDevice device, VkTensorViewCaptureDescriptorDataInfoARM* pInfo, void* pData) {
	return ptr_vkGetTensorViewOpaqueCaptureDescriptorDataARM(device, pInfo, pData);
}
*/
/*
PFN_vkGetValidationCacheDataEXT ptr_vkGetValidationCacheDataEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkGetValidationCacheDataEXT(VkDevice device, VkValidationCacheEXT validationCache, size_t* pDataSize, void* pData) {
	return ptr_vkGetValidationCacheDataEXT(device, validationCache, pDataSize, pData);
}
*/
/*
PFN_vkGetVideoSessionMemoryRequirementsKHR ptr_vkGetVideoSessionMemoryRequirementsKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkGetVideoSessionMemoryRequirementsKHR(VkDevice device, VkVideoSessionKHR videoSession, uint32_t* pMemoryRequirementsCount, VkVideoSessionMemoryRequirementsKHR* pMemoryRequirements) {
	return ptr_vkGetVideoSessionMemoryRequirementsKHR(device, videoSession, pMemoryRequirementsCount, pMemoryRequirements);
}
*/
/*
PFN_vkGetWinrtDisplayNV ptr_vkGetWinrtDisplayNV;
VKAPI_ATTR VkResult VKAPI_CALL vkGetWinrtDisplayNV(VkPhysicalDevice physicalDevice, uint32_t deviceRelativeId, VkDisplayKHR* pDisplay) {
	return ptr_vkGetWinrtDisplayNV(physicalDevice, deviceRelativeId, pDisplay);
}
*/
/*
PFN_vkImportFenceFdKHR ptr_vkImportFenceFdKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkImportFenceFdKHR(VkDevice device, VkImportFenceFdInfoKHR* pImportFenceFdInfo) {
	return ptr_vkImportFenceFdKHR(device, pImportFenceFdInfo);
}
*/
/*
PFN_vkImportFenceWin32HandleKHR ptr_vkImportFenceWin32HandleKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkImportFenceWin32HandleKHR(VkDevice device, VkImportFenceWin32HandleInfoKHR* pImportFenceWin32HandleInfo) {
	return ptr_vkImportFenceWin32HandleKHR(device, pImportFenceWin32HandleInfo);
}
*/
/*
PFN_vkImportSemaphoreFdKHR ptr_vkImportSemaphoreFdKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkImportSemaphoreFdKHR(VkDevice device, VkImportSemaphoreFdInfoKHR* pImportSemaphoreFdInfo) {
	return ptr_vkImportSemaphoreFdKHR(device, pImportSemaphoreFdInfo);
}
*/
/*
PFN_vkImportSemaphoreWin32HandleKHR ptr_vkImportSemaphoreWin32HandleKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkImportSemaphoreWin32HandleKHR(VkDevice device, VkImportSemaphoreWin32HandleInfoKHR* pImportSemaphoreWin32HandleInfo) {
	return ptr_vkImportSemaphoreWin32HandleKHR(device, pImportSemaphoreWin32HandleInfo);
}
*/
/*
PFN_vkInitializePerformanceApiINTEL ptr_vkInitializePerformanceApiINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkInitializePerformanceApiINTEL(VkDevice device, VkInitializePerformanceApiInfoINTEL* pInitializeInfo) {
	return ptr_vkInitializePerformanceApiINTEL(device, pInitializeInfo);
}
*/
/*
PFN_vkInvalidateMappedMemoryRanges ptr_vkInvalidateMappedMemoryRanges;
VKAPI_ATTR VkResult VKAPI_CALL vkInvalidateMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, VkMappedMemoryRange* pMemoryRanges) {
	return ptr_vkInvalidateMappedMemoryRanges(device, memoryRangeCount, pMemoryRanges);
}
*/
/*
PFN_vkLatencySleepNV ptr_vkLatencySleepNV;
VKAPI_ATTR VkResult VKAPI_CALL vkLatencySleepNV(VkDevice device, VkSwapchainKHR swapchain, VkLatencySleepInfoNV* pSleepInfo) {
	return ptr_vkLatencySleepNV(device, swapchain, pSleepInfo);
}
*/
/*
PFN_vkMapMemory ptr_vkMapMemory;
VKAPI_ATTR VkResult VKAPI_CALL vkMapMemory(VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData) {
	return ptr_vkMapMemory(device, memory, offset, size, flags, ppData);
}
*/
/*
PFN_vkMapMemory2 ptr_vkMapMemory2;
VKAPI_ATTR VkResult VKAPI_CALL vkMapMemory2(VkDevice device, VkMemoryMapInfo* pMemoryMapInfo, void** ppData) {
	return ptr_vkMapMemory2(device, pMemoryMapInfo, ppData);
}
*/
/*
PFN_vkMergePipelineCaches ptr_vkMergePipelineCaches;
VKAPI_ATTR VkResult VKAPI_CALL vkMergePipelineCaches(VkDevice device, VkPipelineCache dstCache, uint32_t srcCacheCount, VkPipelineCache* pSrcCaches) {
	return ptr_vkMergePipelineCaches(device, dstCache, srcCacheCount, pSrcCaches);
}
*/
/*
PFN_vkMergeValidationCachesEXT ptr_vkMergeValidationCachesEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkMergeValidationCachesEXT(VkDevice device, VkValidationCacheEXT dstCache, uint32_t srcCacheCount, VkValidationCacheEXT* pSrcCaches) {
	return ptr_vkMergeValidationCachesEXT(device, dstCache, srcCacheCount, pSrcCaches);
}
*/
/*
PFN_vkQueueBeginDebugUtilsLabelEXT ptr_vkQueueBeginDebugUtilsLabelEXT;
VKAPI_ATTR void VKAPI_CALL vkQueueBeginDebugUtilsLabelEXT(VkQueue queue, VkDebugUtilsLabelEXT* pLabelInfo) {
	ptr_vkQueueBeginDebugUtilsLabelEXT(queue, pLabelInfo);
}
*/
/*
PFN_vkQueueBindSparse ptr_vkQueueBindSparse;
VKAPI_ATTR VkResult VKAPI_CALL vkQueueBindSparse(VkQueue queue, uint32_t bindInfoCount, VkBindSparseInfo* pBindInfo, VkFence fence) {
	return ptr_vkQueueBindSparse(queue, bindInfoCount, pBindInfo, fence);
}
*/
/*
PFN_vkQueueEndDebugUtilsLabelEXT ptr_vkQueueEndDebugUtilsLabelEXT;
VKAPI_ATTR void VKAPI_CALL vkQueueEndDebugUtilsLabelEXT(VkQueue queue) {
	ptr_vkQueueEndDebugUtilsLabelEXT(queue);
}
*/
/*
PFN_vkQueueInsertDebugUtilsLabelEXT ptr_vkQueueInsertDebugUtilsLabelEXT;
VKAPI_ATTR void VKAPI_CALL vkQueueInsertDebugUtilsLabelEXT(VkQueue queue, VkDebugUtilsLabelEXT* pLabelInfo) {
	ptr_vkQueueInsertDebugUtilsLabelEXT(queue, pLabelInfo);
}
*/
/*
PFN_vkQueueNotifyOutOfBandNV ptr_vkQueueNotifyOutOfBandNV;
VKAPI_ATTR void VKAPI_CALL vkQueueNotifyOutOfBandNV(VkQueue queue, VkOutOfBandQueueTypeInfoNV* pQueueTypeInfo) {
	ptr_vkQueueNotifyOutOfBandNV(queue, pQueueTypeInfo);
}
*/
/*
PFN_vkQueuePresentKHR ptr_vkQueuePresentKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkQueuePresentKHR(VkQueue queue, VkPresentInfoKHR* pPresentInfo) {
	return ptr_vkQueuePresentKHR(queue, pPresentInfo);
}
*/
/*
PFN_vkQueueSetPerformanceConfigurationINTEL ptr_vkQueueSetPerformanceConfigurationINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkQueueSetPerformanceConfigurationINTEL(VkQueue queue, VkPerformanceConfigurationINTEL configuration) {
	return ptr_vkQueueSetPerformanceConfigurationINTEL(queue, configuration);
}
*/
/*
PFN_vkQueueSubmit ptr_vkQueueSubmit;
VKAPI_ATTR VkResult VKAPI_CALL vkQueueSubmit(VkQueue queue, uint32_t submitCount, VkSubmitInfo* pSubmits, VkFence fence) {
	return ptr_vkQueueSubmit(queue, submitCount, pSubmits, fence);
}
*/
/*
PFN_vkQueueSubmit2 ptr_vkQueueSubmit2;
VKAPI_ATTR VkResult VKAPI_CALL vkQueueSubmit2(VkQueue queue, uint32_t submitCount, VkSubmitInfo2* pSubmits, VkFence fence) {
	return ptr_vkQueueSubmit2(queue, submitCount, pSubmits, fence);
}
*/
/*
PFN_vkQueueWaitIdle ptr_vkQueueWaitIdle;
VKAPI_ATTR VkResult VKAPI_CALL vkQueueWaitIdle(VkQueue queue) {
	return ptr_vkQueueWaitIdle(queue);
}
*/
/*
PFN_vkRegisterDeviceEventEXT ptr_vkRegisterDeviceEventEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkRegisterDeviceEventEXT(VkDevice device, VkDeviceEventInfoEXT* pDeviceEventInfo, VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return ptr_vkRegisterDeviceEventEXT(device, pDeviceEventInfo, pAllocator, pFence);
}
*/
/*
PFN_vkRegisterDisplayEventEXT ptr_vkRegisterDisplayEventEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkRegisterDisplayEventEXT(VkDevice device, VkDisplayKHR display, VkDisplayEventInfoEXT* pDisplayEventInfo, VkAllocationCallbacks* pAllocator, VkFence* pFence) {
	return ptr_vkRegisterDisplayEventEXT(device, display, pDisplayEventInfo, pAllocator, pFence);
}
*/
/*
PFN_vkReleaseCapturedPipelineDataKHR ptr_vkReleaseCapturedPipelineDataKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkReleaseCapturedPipelineDataKHR(VkDevice device, VkReleaseCapturedPipelineDataInfoKHR* pInfo, VkAllocationCallbacks* pAllocator) {
	return ptr_vkReleaseCapturedPipelineDataKHR(device, pInfo, pAllocator);
}
*/
/*
PFN_vkReleaseDisplayEXT ptr_vkReleaseDisplayEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkReleaseDisplayEXT(VkPhysicalDevice physicalDevice, VkDisplayKHR display) {
	return ptr_vkReleaseDisplayEXT(physicalDevice, display);
}
*/
/*
PFN_vkReleaseFullScreenExclusiveModeEXT ptr_vkReleaseFullScreenExclusiveModeEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkReleaseFullScreenExclusiveModeEXT(VkDevice device, VkSwapchainKHR swapchain) {
	return ptr_vkReleaseFullScreenExclusiveModeEXT(device, swapchain);
}
*/
/*
PFN_vkReleasePerformanceConfigurationINTEL ptr_vkReleasePerformanceConfigurationINTEL;
VKAPI_ATTR VkResult VKAPI_CALL vkReleasePerformanceConfigurationINTEL(VkDevice device, VkPerformanceConfigurationINTEL configuration) {
	return ptr_vkReleasePerformanceConfigurationINTEL(device, configuration);
}
*/
/*
PFN_vkReleaseProfilingLockKHR ptr_vkReleaseProfilingLockKHR;
VKAPI_ATTR void VKAPI_CALL vkReleaseProfilingLockKHR(VkDevice device) {
	ptr_vkReleaseProfilingLockKHR(device);
}
*/
/*
PFN_vkReleaseSwapchainImagesKHR ptr_vkReleaseSwapchainImagesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkReleaseSwapchainImagesKHR(VkDevice device, VkReleaseSwapchainImagesInfoKHR* pReleaseInfo) {
	return ptr_vkReleaseSwapchainImagesKHR(device, pReleaseInfo);
}
*/
/*
PFN_vkResetCommandBuffer ptr_vkResetCommandBuffer;
VKAPI_ATTR VkResult VKAPI_CALL vkResetCommandBuffer(VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags) {
	return ptr_vkResetCommandBuffer(commandBuffer, flags);
}
*/
/*
PFN_vkResetCommandPool ptr_vkResetCommandPool;
VKAPI_ATTR VkResult VKAPI_CALL vkResetCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags) {
	return ptr_vkResetCommandPool(device, commandPool, flags);
}
*/
/*
PFN_vkResetDescriptorPool ptr_vkResetDescriptorPool;
VKAPI_ATTR VkResult VKAPI_CALL vkResetDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, VkDescriptorPoolResetFlags flags) {
	return ptr_vkResetDescriptorPool(device, descriptorPool, flags);
}
*/
/*
PFN_vkResetEvent ptr_vkResetEvent;
VKAPI_ATTR VkResult VKAPI_CALL vkResetEvent(VkDevice device, VkEvent event) {
	return ptr_vkResetEvent(device, event);
}
*/
/*
PFN_vkResetFences ptr_vkResetFences;
VKAPI_ATTR VkResult VKAPI_CALL vkResetFences(VkDevice device, uint32_t fenceCount, VkFence* pFences) {
	return ptr_vkResetFences(device, fenceCount, pFences);
}
*/
/*
PFN_vkResetQueryPool ptr_vkResetQueryPool;
VKAPI_ATTR void VKAPI_CALL vkResetQueryPool(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	ptr_vkResetQueryPool(device, queryPool, firstQuery, queryCount);
}
*/
/*
PFN_vkSetDebugUtilsObjectNameEXT ptr_vkSetDebugUtilsObjectNameEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkSetDebugUtilsObjectNameEXT(VkDevice device, VkDebugUtilsObjectNameInfoEXT* pNameInfo) {
	return ptr_vkSetDebugUtilsObjectNameEXT(device, pNameInfo);
}
*/
/*
PFN_vkSetDebugUtilsObjectTagEXT ptr_vkSetDebugUtilsObjectTagEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkSetDebugUtilsObjectTagEXT(VkDevice device, VkDebugUtilsObjectTagInfoEXT* pTagInfo) {
	return ptr_vkSetDebugUtilsObjectTagEXT(device, pTagInfo);
}
*/
/*
PFN_vkSetDeviceMemoryPriorityEXT ptr_vkSetDeviceMemoryPriorityEXT;
VKAPI_ATTR void VKAPI_CALL vkSetDeviceMemoryPriorityEXT(VkDevice device, VkDeviceMemory memory, float priority) {
	ptr_vkSetDeviceMemoryPriorityEXT(device, memory, priority);
}
*/
/*
PFN_vkSetEvent ptr_vkSetEvent;
VKAPI_ATTR VkResult VKAPI_CALL vkSetEvent(VkDevice device, VkEvent event) {
	return ptr_vkSetEvent(device, event);
}
*/
/*
PFN_vkSetHdrMetadataEXT ptr_vkSetHdrMetadataEXT;
VKAPI_ATTR void VKAPI_CALL vkSetHdrMetadataEXT(VkDevice device, uint32_t swapchainCount, VkSwapchainKHR* pSwapchains, VkHdrMetadataEXT* pMetadata) {
	ptr_vkSetHdrMetadataEXT(device, swapchainCount, pSwapchains, pMetadata);
}
*/
/*
PFN_vkSetLatencyMarkerNV ptr_vkSetLatencyMarkerNV;
VKAPI_ATTR void VKAPI_CALL vkSetLatencyMarkerNV(VkDevice device, VkSwapchainKHR swapchain, VkSetLatencyMarkerInfoNV* pLatencyMarkerInfo) {
	ptr_vkSetLatencyMarkerNV(device, swapchain, pLatencyMarkerInfo);
}
*/
/*
PFN_vkSetLatencySleepModeNV ptr_vkSetLatencySleepModeNV;
VKAPI_ATTR VkResult VKAPI_CALL vkSetLatencySleepModeNV(VkDevice device, VkSwapchainKHR swapchain, VkLatencySleepModeInfoNV* pSleepModeInfo) {
	return ptr_vkSetLatencySleepModeNV(device, swapchain, pSleepModeInfo);
}
*/
/*
PFN_vkSetLocalDimmingAMD ptr_vkSetLocalDimmingAMD;
VKAPI_ATTR void VKAPI_CALL vkSetLocalDimmingAMD(VkDevice device, VkSwapchainKHR swapChain, VkBool32 localDimmingEnable) {
	ptr_vkSetLocalDimmingAMD(device, swapChain, localDimmingEnable);
}
*/
/*
PFN_vkSetPrivateData ptr_vkSetPrivateData;
VKAPI_ATTR VkResult VKAPI_CALL vkSetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t data) {
	return ptr_vkSetPrivateData(device, objectType, objectHandle, privateDataSlot, data);
}
*/
/*
PFN_vkSignalSemaphore ptr_vkSignalSemaphore;
VKAPI_ATTR VkResult VKAPI_CALL vkSignalSemaphore(VkDevice device, VkSemaphoreSignalInfo* pSignalInfo) {
	return ptr_vkSignalSemaphore(device, pSignalInfo);
}
*/
/*
PFN_vkSubmitDebugUtilsMessageEXT ptr_vkSubmitDebugUtilsMessageEXT;
VKAPI_ATTR void VKAPI_CALL vkSubmitDebugUtilsMessageEXT(VkInstance instance, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, VkDebugUtilsMessengerCallbackDataEXT* pCallbackData) {
	ptr_vkSubmitDebugUtilsMessageEXT(instance, messageSeverity, messageTypes, pCallbackData);
}
*/
/*
PFN_vkTransitionImageLayout ptr_vkTransitionImageLayout;
VKAPI_ATTR VkResult VKAPI_CALL vkTransitionImageLayout(VkDevice device, uint32_t transitionCount, VkHostImageLayoutTransitionInfo* pTransitions) {
	return ptr_vkTransitionImageLayout(device, transitionCount, pTransitions);
}
*/
/*
PFN_vkTrimCommandPool ptr_vkTrimCommandPool;
VKAPI_ATTR void VKAPI_CALL vkTrimCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
	ptr_vkTrimCommandPool(device, commandPool, flags);
}
*/
/*
PFN_vkUninitializePerformanceApiINTEL ptr_vkUninitializePerformanceApiINTEL;
VKAPI_ATTR void VKAPI_CALL vkUninitializePerformanceApiINTEL(VkDevice device) {
	ptr_vkUninitializePerformanceApiINTEL(device);
}
*/
/*
PFN_vkUnmapMemory ptr_vkUnmapMemory;
VKAPI_ATTR void VKAPI_CALL vkUnmapMemory(VkDevice device, VkDeviceMemory memory) {
	ptr_vkUnmapMemory(device, memory);
}
*/
/*
PFN_vkUnmapMemory2 ptr_vkUnmapMemory2;
VKAPI_ATTR VkResult VKAPI_CALL vkUnmapMemory2(VkDevice device, VkMemoryUnmapInfo* pMemoryUnmapInfo) {
	return ptr_vkUnmapMemory2(device, pMemoryUnmapInfo);
}
*/
/*
PFN_vkUpdateDescriptorSetWithTemplate ptr_vkUpdateDescriptorSetWithTemplate;
VKAPI_ATTR void VKAPI_CALL vkUpdateDescriptorSetWithTemplate(VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplate descriptorUpdateTemplate, void* pData) {
	ptr_vkUpdateDescriptorSetWithTemplate(device, descriptorSet, descriptorUpdateTemplate, pData);
}
*/
/*
PFN_vkUpdateDescriptorSets ptr_vkUpdateDescriptorSets;
VKAPI_ATTR void VKAPI_CALL vkUpdateDescriptorSets(VkDevice device, uint32_t descriptorWriteCount, VkWriteDescriptorSet* pDescriptorWrites, uint32_t descriptorCopyCount, VkCopyDescriptorSet* pDescriptorCopies) {
	ptr_vkUpdateDescriptorSets(device, descriptorWriteCount, pDescriptorWrites, descriptorCopyCount, pDescriptorCopies);
}
*/
/*
PFN_vkUpdateIndirectExecutionSetPipelineEXT ptr_vkUpdateIndirectExecutionSetPipelineEXT;
VKAPI_ATTR void VKAPI_CALL vkUpdateIndirectExecutionSetPipelineEXT(VkDevice device, VkIndirectExecutionSetEXT indirectExecutionSet, uint32_t executionSetWriteCount, VkWriteIndirectExecutionSetPipelineEXT* pExecutionSetWrites) {
	ptr_vkUpdateIndirectExecutionSetPipelineEXT(device, indirectExecutionSet, executionSetWriteCount, pExecutionSetWrites);
}
*/
/*
PFN_vkUpdateIndirectExecutionSetShaderEXT ptr_vkUpdateIndirectExecutionSetShaderEXT;
VKAPI_ATTR void VKAPI_CALL vkUpdateIndirectExecutionSetShaderEXT(VkDevice device, VkIndirectExecutionSetEXT indirectExecutionSet, uint32_t executionSetWriteCount, VkWriteIndirectExecutionSetShaderEXT* pExecutionSetWrites) {
	ptr_vkUpdateIndirectExecutionSetShaderEXT(device, indirectExecutionSet, executionSetWriteCount, pExecutionSetWrites);
}
*/
/*
PFN_vkUpdateVideoSessionParametersKHR ptr_vkUpdateVideoSessionParametersKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkUpdateVideoSessionParametersKHR(VkDevice device, VkVideoSessionParametersKHR videoSessionParameters, VkVideoSessionParametersUpdateInfoKHR* pUpdateInfo) {
	return ptr_vkUpdateVideoSessionParametersKHR(device, videoSessionParameters, pUpdateInfo);
}
*/
/*
PFN_vkWaitForFences ptr_vkWaitForFences;
VKAPI_ATTR VkResult VKAPI_CALL vkWaitForFences(VkDevice device, uint32_t fenceCount, VkFence* pFences, VkBool32 waitAll, uint64_t timeout) {
	return ptr_vkWaitForFences(device, fenceCount, pFences, waitAll, timeout);
}
*/
/*
PFN_vkWaitForPresent2KHR ptr_vkWaitForPresent2KHR;
VKAPI_ATTR VkResult VKAPI_CALL vkWaitForPresent2KHR(VkDevice device, VkSwapchainKHR swapchain, VkPresentWait2InfoKHR* pPresentWait2Info) {
	return ptr_vkWaitForPresent2KHR(device, swapchain, pPresentWait2Info);
}
*/
/*
PFN_vkWaitForPresentKHR ptr_vkWaitForPresentKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkWaitForPresentKHR(VkDevice device, VkSwapchainKHR swapchain, uint64_t presentId, uint64_t timeout) {
	return ptr_vkWaitForPresentKHR(device, swapchain, presentId, timeout);
}
*/
/*
PFN_vkWaitSemaphores ptr_vkWaitSemaphores;
VKAPI_ATTR VkResult VKAPI_CALL vkWaitSemaphores(VkDevice device, VkSemaphoreWaitInfo* pWaitInfo, uint64_t timeout) {
	return ptr_vkWaitSemaphores(device, pWaitInfo, timeout);
}
*/
/*
PFN_vkWriteAccelerationStructuresPropertiesKHR ptr_vkWriteAccelerationStructuresPropertiesKHR;
VKAPI_ATTR VkResult VKAPI_CALL vkWriteAccelerationStructuresPropertiesKHR(VkDevice device, uint32_t accelerationStructureCount, VkAccelerationStructureKHR* pAccelerationStructures, VkQueryType queryType, size_t dataSize, void* pData, size_t stride) {
	return ptr_vkWriteAccelerationStructuresPropertiesKHR(device, accelerationStructureCount, pAccelerationStructures, queryType, dataSize, pData, stride);
}
*/
/*
PFN_vkWriteMicromapsPropertiesEXT ptr_vkWriteMicromapsPropertiesEXT;
VKAPI_ATTR VkResult VKAPI_CALL vkWriteMicromapsPropertiesEXT(VkDevice device, uint32_t micromapCount, VkMicromapEXT* pMicromaps, VkQueryType queryType, size_t dataSize, void* pData, size_t stride) {
	return ptr_vkWriteMicromapsPropertiesEXT(device, micromapCount, pMicromaps, queryType, dataSize, pData, stride);
}
*/
/*
PFN_vkGetInstanceProcAddr ptr_vkGetInstanceProcAddr;

void gfx_vkInit(void* loader) {
    ptr_vkGetInstanceProcAddr = (PFN_vkGetInstanceProcAddr) loader;
    void* context = NULL;

    ptr_vkCreateInstance = (PFN_vkCreateInstance) ptr_vkGetInstanceProcAddr(context, "vkCreateInstance");
    ptr_vkEnumerateInstanceExtensionProperties = (PFN_vkEnumerateInstanceExtensionProperties) ptr_vkGetInstanceProcAddr(context, "vkEnumerateInstanceExtensionProperties");
    ptr_vkEnumerateInstanceLayerProperties = (PFN_vkEnumerateInstanceLayerProperties) ptr_vkGetInstanceProcAddr(context, "vkEnumerateInstanceLayerProperties");
    ptr_vkEnumerateInstanceVersion = (PFN_vkEnumerateInstanceVersion) ptr_vkGetInstanceProcAddr(context, "vkEnumerateInstanceVersion");
}

void gfx_vkInitInstance(VkInstance context) {
    ptr_vkAcquireDrmDisplayEXT = (PFN_vkAcquireDrmDisplayEXT) ptr_vkGetInstanceProcAddr(context, "vkAcquireDrmDisplayEXT");
    ptr_vkAcquireFullScreenExclusiveModeEXT = (PFN_vkAcquireFullScreenExclusiveModeEXT) ptr_vkGetInstanceProcAddr(context, "vkAcquireFullScreenExclusiveModeEXT");
    ptr_vkAcquireNextImage2KHR = (PFN_vkAcquireNextImage2KHR) ptr_vkGetInstanceProcAddr(context, "vkAcquireNextImage2KHR");
    ptr_vkAcquireNextImageKHR = (PFN_vkAcquireNextImageKHR) ptr_vkGetInstanceProcAddr(context, "vkAcquireNextImageKHR");
    ptr_vkAcquirePerformanceConfigurationINTEL = (PFN_vkAcquirePerformanceConfigurationINTEL) ptr_vkGetInstanceProcAddr(context, "vkAcquirePerformanceConfigurationINTEL");
    ptr_vkAcquireProfilingLockKHR = (PFN_vkAcquireProfilingLockKHR) ptr_vkGetInstanceProcAddr(context, "vkAcquireProfilingLockKHR");
    ptr_vkAcquireWinrtDisplayNV = (PFN_vkAcquireWinrtDisplayNV) ptr_vkGetInstanceProcAddr(context, "vkAcquireWinrtDisplayNV");
    ptr_vkAllocateCommandBuffers = (PFN_vkAllocateCommandBuffers) ptr_vkGetInstanceProcAddr(context, "vkAllocateCommandBuffers");
    ptr_vkAllocateDescriptorSets = (PFN_vkAllocateDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkAllocateDescriptorSets");
    ptr_vkAllocateMemory = (PFN_vkAllocateMemory) ptr_vkGetInstanceProcAddr(context, "vkAllocateMemory");
    ptr_vkAntiLagUpdateAMD = (PFN_vkAntiLagUpdateAMD) ptr_vkGetInstanceProcAddr(context, "vkAntiLagUpdateAMD");
    ptr_vkBeginCommandBuffer = (PFN_vkBeginCommandBuffer) ptr_vkGetInstanceProcAddr(context, "vkBeginCommandBuffer");
    ptr_vkBindAccelerationStructureMemoryNV = (PFN_vkBindAccelerationStructureMemoryNV) ptr_vkGetInstanceProcAddr(context, "vkBindAccelerationStructureMemoryNV");
    ptr_vkBindBufferMemory = (PFN_vkBindBufferMemory) ptr_vkGetInstanceProcAddr(context, "vkBindBufferMemory");
    ptr_vkBindBufferMemory2 = (PFN_vkBindBufferMemory2) ptr_vkGetInstanceProcAddr(context, "vkBindBufferMemory2");
    ptr_vkBindDataGraphPipelineSessionMemoryARM = (PFN_vkBindDataGraphPipelineSessionMemoryARM) ptr_vkGetInstanceProcAddr(context, "vkBindDataGraphPipelineSessionMemoryARM");
    ptr_vkBindImageMemory = (PFN_vkBindImageMemory) ptr_vkGetInstanceProcAddr(context, "vkBindImageMemory");
    ptr_vkBindImageMemory2 = (PFN_vkBindImageMemory2) ptr_vkGetInstanceProcAddr(context, "vkBindImageMemory2");
    ptr_vkBindOpticalFlowSessionImageNV = (PFN_vkBindOpticalFlowSessionImageNV) ptr_vkGetInstanceProcAddr(context, "vkBindOpticalFlowSessionImageNV");
    ptr_vkBindTensorMemoryARM = (PFN_vkBindTensorMemoryARM) ptr_vkGetInstanceProcAddr(context, "vkBindTensorMemoryARM");
    ptr_vkBindVideoSessionMemoryKHR = (PFN_vkBindVideoSessionMemoryKHR) ptr_vkGetInstanceProcAddr(context, "vkBindVideoSessionMemoryKHR");
    ptr_vkBuildAccelerationStructuresKHR = (PFN_vkBuildAccelerationStructuresKHR) ptr_vkGetInstanceProcAddr(context, "vkBuildAccelerationStructuresKHR");
    ptr_vkBuildMicromapsEXT = (PFN_vkBuildMicromapsEXT) ptr_vkGetInstanceProcAddr(context, "vkBuildMicromapsEXT");
    ptr_vkCmdBeginConditionalRenderingEXT = (PFN_vkCmdBeginConditionalRenderingEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginConditionalRenderingEXT");
    ptr_vkCmdBeginDebugUtilsLabelEXT = (PFN_vkCmdBeginDebugUtilsLabelEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginDebugUtilsLabelEXT");
    ptr_vkCmdBeginPerTileExecutionQCOM = (PFN_vkCmdBeginPerTileExecutionQCOM) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginPerTileExecutionQCOM");
    ptr_vkCmdBeginQuery = (PFN_vkCmdBeginQuery) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginQuery");
    ptr_vkCmdBeginQueryIndexedEXT = (PFN_vkCmdBeginQueryIndexedEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginQueryIndexedEXT");
    ptr_vkCmdBeginRenderPass = (PFN_vkCmdBeginRenderPass) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginRenderPass");
    ptr_vkCmdBeginRenderPass2 = (PFN_vkCmdBeginRenderPass2) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginRenderPass2");
    ptr_vkCmdBeginRendering = (PFN_vkCmdBeginRendering) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginRendering");
    ptr_vkCmdBeginTransformFeedbackEXT = (PFN_vkCmdBeginTransformFeedbackEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginTransformFeedbackEXT");
    ptr_vkCmdBeginVideoCodingKHR = (PFN_vkCmdBeginVideoCodingKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginVideoCodingKHR");
    ptr_vkCmdBindDescriptorBufferEmbeddedSamplers2EXT = (PFN_vkCmdBindDescriptorBufferEmbeddedSamplers2EXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorBufferEmbeddedSamplers2EXT");
    ptr_vkCmdBindDescriptorBufferEmbeddedSamplersEXT = (PFN_vkCmdBindDescriptorBufferEmbeddedSamplersEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorBufferEmbeddedSamplersEXT");
    ptr_vkCmdBindDescriptorBuffersEXT = (PFN_vkCmdBindDescriptorBuffersEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorBuffersEXT");
    ptr_vkCmdBindDescriptorSets = (PFN_vkCmdBindDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorSets");
    ptr_vkCmdBindDescriptorSets2 = (PFN_vkCmdBindDescriptorSets2) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorSets2");
    ptr_vkCmdBindIndexBuffer = (PFN_vkCmdBindIndexBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdBindIndexBuffer");
    ptr_vkCmdBindIndexBuffer2 = (PFN_vkCmdBindIndexBuffer2) ptr_vkGetInstanceProcAddr(context, "vkCmdBindIndexBuffer2");
    ptr_vkCmdBindInvocationMaskHUAWEI = (PFN_vkCmdBindInvocationMaskHUAWEI) ptr_vkGetInstanceProcAddr(context, "vkCmdBindInvocationMaskHUAWEI");
    ptr_vkCmdBindPipeline = (PFN_vkCmdBindPipeline) ptr_vkGetInstanceProcAddr(context, "vkCmdBindPipeline");
    ptr_vkCmdBindPipelineShaderGroupNV = (PFN_vkCmdBindPipelineShaderGroupNV) ptr_vkGetInstanceProcAddr(context, "vkCmdBindPipelineShaderGroupNV");
    ptr_vkCmdBindShadersEXT = (PFN_vkCmdBindShadersEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBindShadersEXT");
    ptr_vkCmdBindShadingRateImageNV = (PFN_vkCmdBindShadingRateImageNV) ptr_vkGetInstanceProcAddr(context, "vkCmdBindShadingRateImageNV");
    ptr_vkCmdBindTileMemoryQCOM = (PFN_vkCmdBindTileMemoryQCOM) ptr_vkGetInstanceProcAddr(context, "vkCmdBindTileMemoryQCOM");
    ptr_vkCmdBindTransformFeedbackBuffersEXT = (PFN_vkCmdBindTransformFeedbackBuffersEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBindTransformFeedbackBuffersEXT");
    ptr_vkCmdBindVertexBuffers = (PFN_vkCmdBindVertexBuffers) ptr_vkGetInstanceProcAddr(context, "vkCmdBindVertexBuffers");
    ptr_vkCmdBindVertexBuffers2 = (PFN_vkCmdBindVertexBuffers2) ptr_vkGetInstanceProcAddr(context, "vkCmdBindVertexBuffers2");
    ptr_vkCmdBlitImage = (PFN_vkCmdBlitImage) ptr_vkGetInstanceProcAddr(context, "vkCmdBlitImage");
    ptr_vkCmdBlitImage2 = (PFN_vkCmdBlitImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdBlitImage2");
    ptr_vkCmdBuildAccelerationStructureNV = (PFN_vkCmdBuildAccelerationStructureNV) ptr_vkGetInstanceProcAddr(context, "vkCmdBuildAccelerationStructureNV");
    ptr_vkCmdBuildAccelerationStructuresIndirectKHR = (PFN_vkCmdBuildAccelerationStructuresIndirectKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdBuildAccelerationStructuresIndirectKHR");
    ptr_vkCmdBuildAccelerationStructuresKHR = (PFN_vkCmdBuildAccelerationStructuresKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdBuildAccelerationStructuresKHR");
    ptr_vkCmdBuildClusterAccelerationStructureIndirectNV = (PFN_vkCmdBuildClusterAccelerationStructureIndirectNV) ptr_vkGetInstanceProcAddr(context, "vkCmdBuildClusterAccelerationStructureIndirectNV");
    ptr_vkCmdBuildMicromapsEXT = (PFN_vkCmdBuildMicromapsEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdBuildMicromapsEXT");
    ptr_vkCmdBuildPartitionedAccelerationStructuresNV = (PFN_vkCmdBuildPartitionedAccelerationStructuresNV) ptr_vkGetInstanceProcAddr(context, "vkCmdBuildPartitionedAccelerationStructuresNV");
    ptr_vkCmdClearAttachments = (PFN_vkCmdClearAttachments) ptr_vkGetInstanceProcAddr(context, "vkCmdClearAttachments");
    ptr_vkCmdClearColorImage = (PFN_vkCmdClearColorImage) ptr_vkGetInstanceProcAddr(context, "vkCmdClearColorImage");
    ptr_vkCmdClearDepthStencilImage = (PFN_vkCmdClearDepthStencilImage) ptr_vkGetInstanceProcAddr(context, "vkCmdClearDepthStencilImage");
    ptr_vkCmdControlVideoCodingKHR = (PFN_vkCmdControlVideoCodingKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdControlVideoCodingKHR");
    ptr_vkCmdConvertCooperativeVectorMatrixNV = (PFN_vkCmdConvertCooperativeVectorMatrixNV) ptr_vkGetInstanceProcAddr(context, "vkCmdConvertCooperativeVectorMatrixNV");
    ptr_vkCmdCopyAccelerationStructureKHR = (PFN_vkCmdCopyAccelerationStructureKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyAccelerationStructureKHR");
    ptr_vkCmdCopyAccelerationStructureNV = (PFN_vkCmdCopyAccelerationStructureNV) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyAccelerationStructureNV");
    ptr_vkCmdCopyAccelerationStructureToMemoryKHR = (PFN_vkCmdCopyAccelerationStructureToMemoryKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyAccelerationStructureToMemoryKHR");
    ptr_vkCmdCopyBuffer = (PFN_vkCmdCopyBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBuffer");
    ptr_vkCmdCopyBuffer2 = (PFN_vkCmdCopyBuffer2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBuffer2");
    ptr_vkCmdCopyBufferToImage = (PFN_vkCmdCopyBufferToImage) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBufferToImage");
    ptr_vkCmdCopyBufferToImage2 = (PFN_vkCmdCopyBufferToImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBufferToImage2");
    ptr_vkCmdCopyImage = (PFN_vkCmdCopyImage) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImage");
    ptr_vkCmdCopyImage2 = (PFN_vkCmdCopyImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImage2");
    ptr_vkCmdCopyImageToBuffer = (PFN_vkCmdCopyImageToBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImageToBuffer");
    ptr_vkCmdCopyImageToBuffer2 = (PFN_vkCmdCopyImageToBuffer2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImageToBuffer2");
    ptr_vkCmdCopyMemoryIndirectKHR = (PFN_vkCmdCopyMemoryIndirectKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMemoryIndirectKHR");
    ptr_vkCmdCopyMemoryIndirectNV = (PFN_vkCmdCopyMemoryIndirectNV) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMemoryIndirectNV");
    ptr_vkCmdCopyMemoryToAccelerationStructureKHR = (PFN_vkCmdCopyMemoryToAccelerationStructureKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMemoryToAccelerationStructureKHR");
    ptr_vkCmdCopyMemoryToImageIndirectKHR = (PFN_vkCmdCopyMemoryToImageIndirectKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMemoryToImageIndirectKHR");
    ptr_vkCmdCopyMemoryToImageIndirectNV = (PFN_vkCmdCopyMemoryToImageIndirectNV) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMemoryToImageIndirectNV");
    ptr_vkCmdCopyMemoryToMicromapEXT = (PFN_vkCmdCopyMemoryToMicromapEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMemoryToMicromapEXT");
    ptr_vkCmdCopyMicromapEXT = (PFN_vkCmdCopyMicromapEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMicromapEXT");
    ptr_vkCmdCopyMicromapToMemoryEXT = (PFN_vkCmdCopyMicromapToMemoryEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyMicromapToMemoryEXT");
    ptr_vkCmdCopyQueryPoolResults = (PFN_vkCmdCopyQueryPoolResults) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyQueryPoolResults");
    ptr_vkCmdCopyTensorARM = (PFN_vkCmdCopyTensorARM) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyTensorARM");
    ptr_vkCmdCuLaunchKernelNVX = (PFN_vkCmdCuLaunchKernelNVX) ptr_vkGetInstanceProcAddr(context, "vkCmdCuLaunchKernelNVX");
    ptr_vkCmdDebugMarkerBeginEXT = (PFN_vkCmdDebugMarkerBeginEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDebugMarkerBeginEXT");
    ptr_vkCmdDebugMarkerEndEXT = (PFN_vkCmdDebugMarkerEndEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDebugMarkerEndEXT");
    ptr_vkCmdDebugMarkerInsertEXT = (PFN_vkCmdDebugMarkerInsertEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDebugMarkerInsertEXT");
    ptr_vkCmdDecodeVideoKHR = (PFN_vkCmdDecodeVideoKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdDecodeVideoKHR");
    ptr_vkCmdDecompressMemoryIndirectCountNV = (PFN_vkCmdDecompressMemoryIndirectCountNV) ptr_vkGetInstanceProcAddr(context, "vkCmdDecompressMemoryIndirectCountNV");
    ptr_vkCmdDecompressMemoryNV = (PFN_vkCmdDecompressMemoryNV) ptr_vkGetInstanceProcAddr(context, "vkCmdDecompressMemoryNV");
    ptr_vkCmdDispatch = (PFN_vkCmdDispatch) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatch");
    ptr_vkCmdDispatchBase = (PFN_vkCmdDispatchBase) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatchBase");
    ptr_vkCmdDispatchDataGraphARM = (PFN_vkCmdDispatchDataGraphARM) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatchDataGraphARM");
    ptr_vkCmdDispatchIndirect = (PFN_vkCmdDispatchIndirect) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatchIndirect");
    ptr_vkCmdDispatchTileQCOM = (PFN_vkCmdDispatchTileQCOM) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatchTileQCOM");
    ptr_vkCmdDraw = (PFN_vkCmdDraw) ptr_vkGetInstanceProcAddr(context, "vkCmdDraw");
    ptr_vkCmdDrawClusterHUAWEI = (PFN_vkCmdDrawClusterHUAWEI) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawClusterHUAWEI");
    ptr_vkCmdDrawClusterIndirectHUAWEI = (PFN_vkCmdDrawClusterIndirectHUAWEI) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawClusterIndirectHUAWEI");
    ptr_vkCmdDrawIndexed = (PFN_vkCmdDrawIndexed) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndexed");
    ptr_vkCmdDrawIndexedIndirect = (PFN_vkCmdDrawIndexedIndirect) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndexedIndirect");
    ptr_vkCmdDrawIndexedIndirectCount = (PFN_vkCmdDrawIndexedIndirectCount) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndexedIndirectCount");
    ptr_vkCmdDrawIndirect = (PFN_vkCmdDrawIndirect) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndirect");
    ptr_vkCmdDrawIndirectByteCountEXT = (PFN_vkCmdDrawIndirectByteCountEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndirectByteCountEXT");
    ptr_vkCmdDrawIndirectCount = (PFN_vkCmdDrawIndirectCount) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndirectCount");
    ptr_vkCmdDrawMeshTasksEXT = (PFN_vkCmdDrawMeshTasksEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMeshTasksEXT");
    ptr_vkCmdDrawMeshTasksIndirectCountEXT = (PFN_vkCmdDrawMeshTasksIndirectCountEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMeshTasksIndirectCountEXT");
    ptr_vkCmdDrawMeshTasksIndirectCountNV = (PFN_vkCmdDrawMeshTasksIndirectCountNV) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMeshTasksIndirectCountNV");
    ptr_vkCmdDrawMeshTasksIndirectEXT = (PFN_vkCmdDrawMeshTasksIndirectEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMeshTasksIndirectEXT");
    ptr_vkCmdDrawMeshTasksIndirectNV = (PFN_vkCmdDrawMeshTasksIndirectNV) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMeshTasksIndirectNV");
    ptr_vkCmdDrawMeshTasksNV = (PFN_vkCmdDrawMeshTasksNV) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMeshTasksNV");
    ptr_vkCmdDrawMultiEXT = (PFN_vkCmdDrawMultiEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMultiEXT");
    ptr_vkCmdDrawMultiIndexedEXT = (PFN_vkCmdDrawMultiIndexedEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawMultiIndexedEXT");
    ptr_vkCmdEncodeVideoKHR = (PFN_vkCmdEncodeVideoKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdEncodeVideoKHR");
    ptr_vkCmdEndConditionalRenderingEXT = (PFN_vkCmdEndConditionalRenderingEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdEndConditionalRenderingEXT");
    ptr_vkCmdEndDebugUtilsLabelEXT = (PFN_vkCmdEndDebugUtilsLabelEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdEndDebugUtilsLabelEXT");
    ptr_vkCmdEndPerTileExecutionQCOM = (PFN_vkCmdEndPerTileExecutionQCOM) ptr_vkGetInstanceProcAddr(context, "vkCmdEndPerTileExecutionQCOM");
    ptr_vkCmdEndQuery = (PFN_vkCmdEndQuery) ptr_vkGetInstanceProcAddr(context, "vkCmdEndQuery");
    ptr_vkCmdEndQueryIndexedEXT = (PFN_vkCmdEndQueryIndexedEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdEndQueryIndexedEXT");
    ptr_vkCmdEndRenderPass = (PFN_vkCmdEndRenderPass) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRenderPass");
    ptr_vkCmdEndRenderPass2 = (PFN_vkCmdEndRenderPass2) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRenderPass2");
    ptr_vkCmdEndRendering = (PFN_vkCmdEndRendering) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRendering");
    ptr_vkCmdEndRendering2EXT = (PFN_vkCmdEndRendering2EXT) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRendering2EXT");
    ptr_vkCmdEndTransformFeedbackEXT = (PFN_vkCmdEndTransformFeedbackEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdEndTransformFeedbackEXT");
    ptr_vkCmdEndVideoCodingKHR = (PFN_vkCmdEndVideoCodingKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdEndVideoCodingKHR");
    ptr_vkCmdExecuteCommands = (PFN_vkCmdExecuteCommands) ptr_vkGetInstanceProcAddr(context, "vkCmdExecuteCommands");
    ptr_vkCmdExecuteGeneratedCommandsEXT = (PFN_vkCmdExecuteGeneratedCommandsEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdExecuteGeneratedCommandsEXT");
    ptr_vkCmdExecuteGeneratedCommandsNV = (PFN_vkCmdExecuteGeneratedCommandsNV) ptr_vkGetInstanceProcAddr(context, "vkCmdExecuteGeneratedCommandsNV");
    ptr_vkCmdFillBuffer = (PFN_vkCmdFillBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdFillBuffer");
    ptr_vkCmdInsertDebugUtilsLabelEXT = (PFN_vkCmdInsertDebugUtilsLabelEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdInsertDebugUtilsLabelEXT");
    ptr_vkCmdNextSubpass = (PFN_vkCmdNextSubpass) ptr_vkGetInstanceProcAddr(context, "vkCmdNextSubpass");
    ptr_vkCmdNextSubpass2 = (PFN_vkCmdNextSubpass2) ptr_vkGetInstanceProcAddr(context, "vkCmdNextSubpass2");
    ptr_vkCmdOpticalFlowExecuteNV = (PFN_vkCmdOpticalFlowExecuteNV) ptr_vkGetInstanceProcAddr(context, "vkCmdOpticalFlowExecuteNV");
    ptr_vkCmdPipelineBarrier = (PFN_vkCmdPipelineBarrier) ptr_vkGetInstanceProcAddr(context, "vkCmdPipelineBarrier");
    ptr_vkCmdPipelineBarrier2 = (PFN_vkCmdPipelineBarrier2) ptr_vkGetInstanceProcAddr(context, "vkCmdPipelineBarrier2");
    ptr_vkCmdPreprocessGeneratedCommandsEXT = (PFN_vkCmdPreprocessGeneratedCommandsEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdPreprocessGeneratedCommandsEXT");
    ptr_vkCmdPreprocessGeneratedCommandsNV = (PFN_vkCmdPreprocessGeneratedCommandsNV) ptr_vkGetInstanceProcAddr(context, "vkCmdPreprocessGeneratedCommandsNV");
    ptr_vkCmdPushConstants = (PFN_vkCmdPushConstants) ptr_vkGetInstanceProcAddr(context, "vkCmdPushConstants");
    ptr_vkCmdPushConstants2 = (PFN_vkCmdPushConstants2) ptr_vkGetInstanceProcAddr(context, "vkCmdPushConstants2");
    ptr_vkCmdPushDescriptorSet = (PFN_vkCmdPushDescriptorSet) ptr_vkGetInstanceProcAddr(context, "vkCmdPushDescriptorSet");
    ptr_vkCmdPushDescriptorSet2 = (PFN_vkCmdPushDescriptorSet2) ptr_vkGetInstanceProcAddr(context, "vkCmdPushDescriptorSet2");
    ptr_vkCmdPushDescriptorSetWithTemplate = (PFN_vkCmdPushDescriptorSetWithTemplate) ptr_vkGetInstanceProcAddr(context, "vkCmdPushDescriptorSetWithTemplate");
    ptr_vkCmdPushDescriptorSetWithTemplate2 = (PFN_vkCmdPushDescriptorSetWithTemplate2) ptr_vkGetInstanceProcAddr(context, "vkCmdPushDescriptorSetWithTemplate2");
    ptr_vkCmdResetEvent = (PFN_vkCmdResetEvent) ptr_vkGetInstanceProcAddr(context, "vkCmdResetEvent");
    ptr_vkCmdResetEvent2 = (PFN_vkCmdResetEvent2) ptr_vkGetInstanceProcAddr(context, "vkCmdResetEvent2");
    ptr_vkCmdResetQueryPool = (PFN_vkCmdResetQueryPool) ptr_vkGetInstanceProcAddr(context, "vkCmdResetQueryPool");
    ptr_vkCmdResolveImage = (PFN_vkCmdResolveImage) ptr_vkGetInstanceProcAddr(context, "vkCmdResolveImage");
    ptr_vkCmdResolveImage2 = (PFN_vkCmdResolveImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdResolveImage2");
    ptr_vkCmdSetAlphaToCoverageEnableEXT = (PFN_vkCmdSetAlphaToCoverageEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetAlphaToCoverageEnableEXT");
    ptr_vkCmdSetAlphaToOneEnableEXT = (PFN_vkCmdSetAlphaToOneEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetAlphaToOneEnableEXT");
    ptr_vkCmdSetAttachmentFeedbackLoopEnableEXT = (PFN_vkCmdSetAttachmentFeedbackLoopEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetAttachmentFeedbackLoopEnableEXT");
    ptr_vkCmdSetCheckpointNV = (PFN_vkCmdSetCheckpointNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCheckpointNV");
    ptr_vkCmdSetCoarseSampleOrderNV = (PFN_vkCmdSetCoarseSampleOrderNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoarseSampleOrderNV");
    ptr_vkCmdSetColorBlendAdvancedEXT = (PFN_vkCmdSetColorBlendAdvancedEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetColorBlendAdvancedEXT");
    ptr_vkCmdSetColorBlendEnableEXT = (PFN_vkCmdSetColorBlendEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetColorBlendEnableEXT");
    ptr_vkCmdSetColorBlendEquationEXT = (PFN_vkCmdSetColorBlendEquationEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetColorBlendEquationEXT");
    ptr_vkCmdSetColorWriteEnableEXT = (PFN_vkCmdSetColorWriteEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetColorWriteEnableEXT");
    ptr_vkCmdSetColorWriteMaskEXT = (PFN_vkCmdSetColorWriteMaskEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetColorWriteMaskEXT");
    ptr_vkCmdSetConservativeRasterizationModeEXT = (PFN_vkCmdSetConservativeRasterizationModeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetConservativeRasterizationModeEXT");
    ptr_vkCmdSetCoverageModulationModeNV = (PFN_vkCmdSetCoverageModulationModeNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoverageModulationModeNV");
    ptr_vkCmdSetCoverageModulationTableEnableNV = (PFN_vkCmdSetCoverageModulationTableEnableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoverageModulationTableEnableNV");
    ptr_vkCmdSetCoverageModulationTableNV = (PFN_vkCmdSetCoverageModulationTableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoverageModulationTableNV");
    ptr_vkCmdSetCoverageReductionModeNV = (PFN_vkCmdSetCoverageReductionModeNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoverageReductionModeNV");
    ptr_vkCmdSetCoverageToColorEnableNV = (PFN_vkCmdSetCoverageToColorEnableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoverageToColorEnableNV");
    ptr_vkCmdSetCoverageToColorLocationNV = (PFN_vkCmdSetCoverageToColorLocationNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCoverageToColorLocationNV");
    ptr_vkCmdSetCullMode = (PFN_vkCmdSetCullMode) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCullMode");
    ptr_vkCmdSetDepthBias = (PFN_vkCmdSetDepthBias) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBias");
    ptr_vkCmdSetDepthBias2EXT = (PFN_vkCmdSetDepthBias2EXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBias2EXT");
    ptr_vkCmdSetDepthBiasEnable = (PFN_vkCmdSetDepthBiasEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBiasEnable");
    ptr_vkCmdSetDepthBounds = (PFN_vkCmdSetDepthBounds) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBounds");
    ptr_vkCmdSetDepthBoundsTestEnable = (PFN_vkCmdSetDepthBoundsTestEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBoundsTestEnable");
    ptr_vkCmdSetDepthClampEnableEXT = (PFN_vkCmdSetDepthClampEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthClampEnableEXT");
    ptr_vkCmdSetDepthClampRangeEXT = (PFN_vkCmdSetDepthClampRangeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthClampRangeEXT");
    ptr_vkCmdSetDepthClipEnableEXT = (PFN_vkCmdSetDepthClipEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthClipEnableEXT");
    ptr_vkCmdSetDepthClipNegativeOneToOneEXT = (PFN_vkCmdSetDepthClipNegativeOneToOneEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthClipNegativeOneToOneEXT");
    ptr_vkCmdSetDepthCompareOp = (PFN_vkCmdSetDepthCompareOp) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthCompareOp");
    ptr_vkCmdSetDepthTestEnable = (PFN_vkCmdSetDepthTestEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthTestEnable");
    ptr_vkCmdSetDepthWriteEnable = (PFN_vkCmdSetDepthWriteEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthWriteEnable");
    ptr_vkCmdSetDescriptorBufferOffsets2EXT = (PFN_vkCmdSetDescriptorBufferOffsets2EXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDescriptorBufferOffsets2EXT");
    ptr_vkCmdSetDescriptorBufferOffsetsEXT = (PFN_vkCmdSetDescriptorBufferOffsetsEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDescriptorBufferOffsetsEXT");
    ptr_vkCmdSetDeviceMask = (PFN_vkCmdSetDeviceMask) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDeviceMask");
    ptr_vkCmdSetDiscardRectangleEXT = (PFN_vkCmdSetDiscardRectangleEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDiscardRectangleEXT");
    ptr_vkCmdSetDiscardRectangleEnableEXT = (PFN_vkCmdSetDiscardRectangleEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDiscardRectangleEnableEXT");
    ptr_vkCmdSetDiscardRectangleModeEXT = (PFN_vkCmdSetDiscardRectangleModeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDiscardRectangleModeEXT");
    ptr_vkCmdSetEvent = (PFN_vkCmdSetEvent) ptr_vkGetInstanceProcAddr(context, "vkCmdSetEvent");
    ptr_vkCmdSetEvent2 = (PFN_vkCmdSetEvent2) ptr_vkGetInstanceProcAddr(context, "vkCmdSetEvent2");
    ptr_vkCmdSetExclusiveScissorEnableNV = (PFN_vkCmdSetExclusiveScissorEnableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetExclusiveScissorEnableNV");
    ptr_vkCmdSetExclusiveScissorNV = (PFN_vkCmdSetExclusiveScissorNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetExclusiveScissorNV");
    ptr_vkCmdSetExtraPrimitiveOverestimationSizeEXT = (PFN_vkCmdSetExtraPrimitiveOverestimationSizeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetExtraPrimitiveOverestimationSizeEXT");
    ptr_vkCmdSetFrontFace = (PFN_vkCmdSetFrontFace) ptr_vkGetInstanceProcAddr(context, "vkCmdSetFrontFace");
    ptr_vkCmdSetLineRasterizationModeEXT = (PFN_vkCmdSetLineRasterizationModeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLineRasterizationModeEXT");
    ptr_vkCmdSetLineStipple = (PFN_vkCmdSetLineStipple) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLineStipple");
    ptr_vkCmdSetLineStippleEnableEXT = (PFN_vkCmdSetLineStippleEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLineStippleEnableEXT");
    ptr_vkCmdSetLineWidth = (PFN_vkCmdSetLineWidth) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLineWidth");
    ptr_vkCmdSetLogicOpEXT = (PFN_vkCmdSetLogicOpEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLogicOpEXT");
    ptr_vkCmdSetLogicOpEnableEXT = (PFN_vkCmdSetLogicOpEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLogicOpEnableEXT");
    ptr_vkCmdSetPatchControlPointsEXT = (PFN_vkCmdSetPatchControlPointsEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPatchControlPointsEXT");
    ptr_vkCmdSetPerformanceMarkerINTEL = (PFN_vkCmdSetPerformanceMarkerINTEL) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPerformanceMarkerINTEL");
    ptr_vkCmdSetPerformanceOverrideINTEL = (PFN_vkCmdSetPerformanceOverrideINTEL) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPerformanceOverrideINTEL");
    ptr_vkCmdSetPerformanceStreamMarkerINTEL = (PFN_vkCmdSetPerformanceStreamMarkerINTEL) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPerformanceStreamMarkerINTEL");
    ptr_vkCmdSetPolygonModeEXT = (PFN_vkCmdSetPolygonModeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPolygonModeEXT");
    ptr_vkCmdSetPrimitiveRestartEnable = (PFN_vkCmdSetPrimitiveRestartEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPrimitiveRestartEnable");
    ptr_vkCmdSetPrimitiveTopology = (PFN_vkCmdSetPrimitiveTopology) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPrimitiveTopology");
    ptr_vkCmdSetProvokingVertexModeEXT = (PFN_vkCmdSetProvokingVertexModeEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetProvokingVertexModeEXT");
    ptr_vkCmdSetRasterizationSamplesEXT = (PFN_vkCmdSetRasterizationSamplesEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRasterizationSamplesEXT");
    ptr_vkCmdSetRasterizationStreamEXT = (PFN_vkCmdSetRasterizationStreamEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRasterizationStreamEXT");
    ptr_vkCmdSetRasterizerDiscardEnable = (PFN_vkCmdSetRasterizerDiscardEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRasterizerDiscardEnable");
    ptr_vkCmdSetRayTracingPipelineStackSizeKHR = (PFN_vkCmdSetRayTracingPipelineStackSizeKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRayTracingPipelineStackSizeKHR");
    ptr_vkCmdSetRenderingAttachmentLocations = (PFN_vkCmdSetRenderingAttachmentLocations) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRenderingAttachmentLocations");
    ptr_vkCmdSetRenderingInputAttachmentIndices = (PFN_vkCmdSetRenderingInputAttachmentIndices) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRenderingInputAttachmentIndices");
    ptr_vkCmdSetRepresentativeFragmentTestEnableNV = (PFN_vkCmdSetRepresentativeFragmentTestEnableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRepresentativeFragmentTestEnableNV");
    ptr_vkCmdSetSampleLocationsEXT = (PFN_vkCmdSetSampleLocationsEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetSampleLocationsEXT");
    ptr_vkCmdSetSampleLocationsEnableEXT = (PFN_vkCmdSetSampleLocationsEnableEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetSampleLocationsEnableEXT");
    ptr_vkCmdSetSampleMaskEXT = (PFN_vkCmdSetSampleMaskEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetSampleMaskEXT");
    ptr_vkCmdSetScissor = (PFN_vkCmdSetScissor) ptr_vkGetInstanceProcAddr(context, "vkCmdSetScissor");
    ptr_vkCmdSetScissorWithCount = (PFN_vkCmdSetScissorWithCount) ptr_vkGetInstanceProcAddr(context, "vkCmdSetScissorWithCount");
    ptr_vkCmdSetShadingRateImageEnableNV = (PFN_vkCmdSetShadingRateImageEnableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetShadingRateImageEnableNV");
    ptr_vkCmdSetStencilCompareMask = (PFN_vkCmdSetStencilCompareMask) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilCompareMask");
    ptr_vkCmdSetStencilOp = (PFN_vkCmdSetStencilOp) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilOp");
    ptr_vkCmdSetStencilReference = (PFN_vkCmdSetStencilReference) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilReference");
    ptr_vkCmdSetStencilTestEnable = (PFN_vkCmdSetStencilTestEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilTestEnable");
    ptr_vkCmdSetStencilWriteMask = (PFN_vkCmdSetStencilWriteMask) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilWriteMask");
    ptr_vkCmdSetTessellationDomainOriginEXT = (PFN_vkCmdSetTessellationDomainOriginEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetTessellationDomainOriginEXT");
    ptr_vkCmdSetVertexInputEXT = (PFN_vkCmdSetVertexInputEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdSetVertexInputEXT");
    ptr_vkCmdSetViewport = (PFN_vkCmdSetViewport) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewport");
    ptr_vkCmdSetViewportShadingRatePaletteNV = (PFN_vkCmdSetViewportShadingRatePaletteNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewportShadingRatePaletteNV");
    ptr_vkCmdSetViewportSwizzleNV = (PFN_vkCmdSetViewportSwizzleNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewportSwizzleNV");
    ptr_vkCmdSetViewportWScalingEnableNV = (PFN_vkCmdSetViewportWScalingEnableNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewportWScalingEnableNV");
    ptr_vkCmdSetViewportWScalingNV = (PFN_vkCmdSetViewportWScalingNV) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewportWScalingNV");
    ptr_vkCmdSetViewportWithCount = (PFN_vkCmdSetViewportWithCount) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewportWithCount");
    ptr_vkCmdSubpassShadingHUAWEI = (PFN_vkCmdSubpassShadingHUAWEI) ptr_vkGetInstanceProcAddr(context, "vkCmdSubpassShadingHUAWEI");
    ptr_vkCmdTraceRaysIndirect2KHR = (PFN_vkCmdTraceRaysIndirect2KHR) ptr_vkGetInstanceProcAddr(context, "vkCmdTraceRaysIndirect2KHR");
    ptr_vkCmdTraceRaysIndirectKHR = (PFN_vkCmdTraceRaysIndirectKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdTraceRaysIndirectKHR");
    ptr_vkCmdTraceRaysKHR = (PFN_vkCmdTraceRaysKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdTraceRaysKHR");
    ptr_vkCmdTraceRaysNV = (PFN_vkCmdTraceRaysNV) ptr_vkGetInstanceProcAddr(context, "vkCmdTraceRaysNV");
    ptr_vkCmdUpdateBuffer = (PFN_vkCmdUpdateBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdUpdateBuffer");
    ptr_vkCmdUpdatePipelineIndirectBufferNV = (PFN_vkCmdUpdatePipelineIndirectBufferNV) ptr_vkGetInstanceProcAddr(context, "vkCmdUpdatePipelineIndirectBufferNV");
    ptr_vkCmdWaitEvents = (PFN_vkCmdWaitEvents) ptr_vkGetInstanceProcAddr(context, "vkCmdWaitEvents");
    ptr_vkCmdWaitEvents2 = (PFN_vkCmdWaitEvents2) ptr_vkGetInstanceProcAddr(context, "vkCmdWaitEvents2");
    ptr_vkCmdWriteAccelerationStructuresPropertiesKHR = (PFN_vkCmdWriteAccelerationStructuresPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteAccelerationStructuresPropertiesKHR");
    ptr_vkCmdWriteAccelerationStructuresPropertiesNV = (PFN_vkCmdWriteAccelerationStructuresPropertiesNV) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteAccelerationStructuresPropertiesNV");
    ptr_vkCmdWriteBufferMarker2AMD = (PFN_vkCmdWriteBufferMarker2AMD) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteBufferMarker2AMD");
    ptr_vkCmdWriteBufferMarkerAMD = (PFN_vkCmdWriteBufferMarkerAMD) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteBufferMarkerAMD");
    ptr_vkCmdWriteMicromapsPropertiesEXT = (PFN_vkCmdWriteMicromapsPropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteMicromapsPropertiesEXT");
    ptr_vkCmdWriteTimestamp = (PFN_vkCmdWriteTimestamp) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteTimestamp");
    ptr_vkCmdWriteTimestamp2 = (PFN_vkCmdWriteTimestamp2) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteTimestamp2");
    ptr_vkCompileDeferredNV = (PFN_vkCompileDeferredNV) ptr_vkGetInstanceProcAddr(context, "vkCompileDeferredNV");
    ptr_vkConvertCooperativeVectorMatrixNV = (PFN_vkConvertCooperativeVectorMatrixNV) ptr_vkGetInstanceProcAddr(context, "vkConvertCooperativeVectorMatrixNV");
    ptr_vkCopyAccelerationStructureKHR = (PFN_vkCopyAccelerationStructureKHR) ptr_vkGetInstanceProcAddr(context, "vkCopyAccelerationStructureKHR");
    ptr_vkCopyAccelerationStructureToMemoryKHR = (PFN_vkCopyAccelerationStructureToMemoryKHR) ptr_vkGetInstanceProcAddr(context, "vkCopyAccelerationStructureToMemoryKHR");
    ptr_vkCopyImageToImage = (PFN_vkCopyImageToImage) ptr_vkGetInstanceProcAddr(context, "vkCopyImageToImage");
    ptr_vkCopyImageToMemory = (PFN_vkCopyImageToMemory) ptr_vkGetInstanceProcAddr(context, "vkCopyImageToMemory");
    ptr_vkCopyMemoryToAccelerationStructureKHR = (PFN_vkCopyMemoryToAccelerationStructureKHR) ptr_vkGetInstanceProcAddr(context, "vkCopyMemoryToAccelerationStructureKHR");
    ptr_vkCopyMemoryToImage = (PFN_vkCopyMemoryToImage) ptr_vkGetInstanceProcAddr(context, "vkCopyMemoryToImage");
    ptr_vkCopyMemoryToMicromapEXT = (PFN_vkCopyMemoryToMicromapEXT) ptr_vkGetInstanceProcAddr(context, "vkCopyMemoryToMicromapEXT");
    ptr_vkCopyMicromapEXT = (PFN_vkCopyMicromapEXT) ptr_vkGetInstanceProcAddr(context, "vkCopyMicromapEXT");
    ptr_vkCopyMicromapToMemoryEXT = (PFN_vkCopyMicromapToMemoryEXT) ptr_vkGetInstanceProcAddr(context, "vkCopyMicromapToMemoryEXT");
    ptr_vkCreateAccelerationStructureKHR = (PFN_vkCreateAccelerationStructureKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateAccelerationStructureKHR");
    ptr_vkCreateAccelerationStructureNV = (PFN_vkCreateAccelerationStructureNV) ptr_vkGetInstanceProcAddr(context, "vkCreateAccelerationStructureNV");
    ptr_vkCreateBuffer = (PFN_vkCreateBuffer) ptr_vkGetInstanceProcAddr(context, "vkCreateBuffer");
    ptr_vkCreateBufferView = (PFN_vkCreateBufferView) ptr_vkGetInstanceProcAddr(context, "vkCreateBufferView");
    ptr_vkCreateCommandPool = (PFN_vkCreateCommandPool) ptr_vkGetInstanceProcAddr(context, "vkCreateCommandPool");
    ptr_vkCreateComputePipelines = (PFN_vkCreateComputePipelines) ptr_vkGetInstanceProcAddr(context, "vkCreateComputePipelines");
    ptr_vkCreateCuFunctionNVX = (PFN_vkCreateCuFunctionNVX) ptr_vkGetInstanceProcAddr(context, "vkCreateCuFunctionNVX");
    ptr_vkCreateCuModuleNVX = (PFN_vkCreateCuModuleNVX) ptr_vkGetInstanceProcAddr(context, "vkCreateCuModuleNVX");
    ptr_vkCreateDataGraphPipelineSessionARM = (PFN_vkCreateDataGraphPipelineSessionARM) ptr_vkGetInstanceProcAddr(context, "vkCreateDataGraphPipelineSessionARM");
    ptr_vkCreateDataGraphPipelinesARM = (PFN_vkCreateDataGraphPipelinesARM) ptr_vkGetInstanceProcAddr(context, "vkCreateDataGraphPipelinesARM");
    ptr_vkCreateDebugReportCallbackEXT = (PFN_vkCreateDebugReportCallbackEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateDebugReportCallbackEXT");
    ptr_vkCreateDebugUtilsMessengerEXT = (PFN_vkCreateDebugUtilsMessengerEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateDebugUtilsMessengerEXT");
    ptr_vkCreateDeferredOperationKHR = (PFN_vkCreateDeferredOperationKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateDeferredOperationKHR");
    ptr_vkCreateDescriptorPool = (PFN_vkCreateDescriptorPool) ptr_vkGetInstanceProcAddr(context, "vkCreateDescriptorPool");
    ptr_vkCreateDescriptorSetLayout = (PFN_vkCreateDescriptorSetLayout) ptr_vkGetInstanceProcAddr(context, "vkCreateDescriptorSetLayout");
    ptr_vkCreateDescriptorUpdateTemplate = (PFN_vkCreateDescriptorUpdateTemplate) ptr_vkGetInstanceProcAddr(context, "vkCreateDescriptorUpdateTemplate");
    ptr_vkCreateDevice = (PFN_vkCreateDevice) ptr_vkGetInstanceProcAddr(context, "vkCreateDevice");
    ptr_vkCreateDisplayModeKHR = (PFN_vkCreateDisplayModeKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateDisplayModeKHR");
    ptr_vkCreateDisplayPlaneSurfaceKHR = (PFN_vkCreateDisplayPlaneSurfaceKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateDisplayPlaneSurfaceKHR");
    ptr_vkCreateEvent = (PFN_vkCreateEvent) ptr_vkGetInstanceProcAddr(context, "vkCreateEvent");
    ptr_vkCreateExternalComputeQueueNV = (PFN_vkCreateExternalComputeQueueNV) ptr_vkGetInstanceProcAddr(context, "vkCreateExternalComputeQueueNV");
    ptr_vkCreateFence = (PFN_vkCreateFence) ptr_vkGetInstanceProcAddr(context, "vkCreateFence");
    ptr_vkCreateFramebuffer = (PFN_vkCreateFramebuffer) ptr_vkGetInstanceProcAddr(context, "vkCreateFramebuffer");
    ptr_vkCreateGraphicsPipelines = (PFN_vkCreateGraphicsPipelines) ptr_vkGetInstanceProcAddr(context, "vkCreateGraphicsPipelines");
    ptr_vkCreateHeadlessSurfaceEXT = (PFN_vkCreateHeadlessSurfaceEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateHeadlessSurfaceEXT");
    ptr_vkCreateImage = (PFN_vkCreateImage) ptr_vkGetInstanceProcAddr(context, "vkCreateImage");
    ptr_vkCreateImageView = (PFN_vkCreateImageView) ptr_vkGetInstanceProcAddr(context, "vkCreateImageView");
    ptr_vkCreateIndirectCommandsLayoutEXT = (PFN_vkCreateIndirectCommandsLayoutEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateIndirectCommandsLayoutEXT");
    ptr_vkCreateIndirectCommandsLayoutNV = (PFN_vkCreateIndirectCommandsLayoutNV) ptr_vkGetInstanceProcAddr(context, "vkCreateIndirectCommandsLayoutNV");
    ptr_vkCreateIndirectExecutionSetEXT = (PFN_vkCreateIndirectExecutionSetEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateIndirectExecutionSetEXT");
    ptr_vkCreateMetalSurfaceEXT = (PFN_vkCreateMetalSurfaceEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateMetalSurfaceEXT");
    ptr_vkCreateMicromapEXT = (PFN_vkCreateMicromapEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateMicromapEXT");
    ptr_vkCreateOpticalFlowSessionNV = (PFN_vkCreateOpticalFlowSessionNV) ptr_vkGetInstanceProcAddr(context, "vkCreateOpticalFlowSessionNV");
    ptr_vkCreatePipelineBinariesKHR = (PFN_vkCreatePipelineBinariesKHR) ptr_vkGetInstanceProcAddr(context, "vkCreatePipelineBinariesKHR");
    ptr_vkCreatePipelineCache = (PFN_vkCreatePipelineCache) ptr_vkGetInstanceProcAddr(context, "vkCreatePipelineCache");
    ptr_vkCreatePipelineLayout = (PFN_vkCreatePipelineLayout) ptr_vkGetInstanceProcAddr(context, "vkCreatePipelineLayout");
    ptr_vkCreatePrivateDataSlot = (PFN_vkCreatePrivateDataSlot) ptr_vkGetInstanceProcAddr(context, "vkCreatePrivateDataSlot");
    ptr_vkCreateQueryPool = (PFN_vkCreateQueryPool) ptr_vkGetInstanceProcAddr(context, "vkCreateQueryPool");
    ptr_vkCreateRayTracingPipelinesKHR = (PFN_vkCreateRayTracingPipelinesKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateRayTracingPipelinesKHR");
    ptr_vkCreateRayTracingPipelinesNV = (PFN_vkCreateRayTracingPipelinesNV) ptr_vkGetInstanceProcAddr(context, "vkCreateRayTracingPipelinesNV");
    ptr_vkCreateRenderPass = (PFN_vkCreateRenderPass) ptr_vkGetInstanceProcAddr(context, "vkCreateRenderPass");
    ptr_vkCreateRenderPass2 = (PFN_vkCreateRenderPass2) ptr_vkGetInstanceProcAddr(context, "vkCreateRenderPass2");
    ptr_vkCreateSampler = (PFN_vkCreateSampler) ptr_vkGetInstanceProcAddr(context, "vkCreateSampler");
    ptr_vkCreateSamplerYcbcrConversion = (PFN_vkCreateSamplerYcbcrConversion) ptr_vkGetInstanceProcAddr(context, "vkCreateSamplerYcbcrConversion");
    ptr_vkCreateSemaphore = (PFN_vkCreateSemaphore) ptr_vkGetInstanceProcAddr(context, "vkCreateSemaphore");
    ptr_vkCreateShaderModule = (PFN_vkCreateShaderModule) ptr_vkGetInstanceProcAddr(context, "vkCreateShaderModule");
    ptr_vkCreateShadersEXT = (PFN_vkCreateShadersEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateShadersEXT");
    ptr_vkCreateSharedSwapchainsKHR = (PFN_vkCreateSharedSwapchainsKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateSharedSwapchainsKHR");
    ptr_vkCreateSwapchainKHR = (PFN_vkCreateSwapchainKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateSwapchainKHR");
    ptr_vkCreateTensorARM = (PFN_vkCreateTensorARM) ptr_vkGetInstanceProcAddr(context, "vkCreateTensorARM");
    ptr_vkCreateTensorViewARM = (PFN_vkCreateTensorViewARM) ptr_vkGetInstanceProcAddr(context, "vkCreateTensorViewARM");
    ptr_vkCreateValidationCacheEXT = (PFN_vkCreateValidationCacheEXT) ptr_vkGetInstanceProcAddr(context, "vkCreateValidationCacheEXT");
    ptr_vkCreateVideoSessionKHR = (PFN_vkCreateVideoSessionKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateVideoSessionKHR");
    ptr_vkCreateVideoSessionParametersKHR = (PFN_vkCreateVideoSessionParametersKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateVideoSessionParametersKHR");
    ptr_vkCreateWin32SurfaceKHR = (PFN_vkCreateWin32SurfaceKHR) ptr_vkGetInstanceProcAddr(context, "vkCreateWin32SurfaceKHR");
    ptr_vkDebugMarkerSetObjectNameEXT = (PFN_vkDebugMarkerSetObjectNameEXT) ptr_vkGetInstanceProcAddr(context, "vkDebugMarkerSetObjectNameEXT");
    ptr_vkDebugMarkerSetObjectTagEXT = (PFN_vkDebugMarkerSetObjectTagEXT) ptr_vkGetInstanceProcAddr(context, "vkDebugMarkerSetObjectTagEXT");
    ptr_vkDebugReportMessageEXT = (PFN_vkDebugReportMessageEXT) ptr_vkGetInstanceProcAddr(context, "vkDebugReportMessageEXT");
    ptr_vkDeferredOperationJoinKHR = (PFN_vkDeferredOperationJoinKHR) ptr_vkGetInstanceProcAddr(context, "vkDeferredOperationJoinKHR");
    ptr_vkDestroyAccelerationStructureKHR = (PFN_vkDestroyAccelerationStructureKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroyAccelerationStructureKHR");
    ptr_vkDestroyAccelerationStructureNV = (PFN_vkDestroyAccelerationStructureNV) ptr_vkGetInstanceProcAddr(context, "vkDestroyAccelerationStructureNV");
    ptr_vkDestroyBuffer = (PFN_vkDestroyBuffer) ptr_vkGetInstanceProcAddr(context, "vkDestroyBuffer");
    ptr_vkDestroyBufferView = (PFN_vkDestroyBufferView) ptr_vkGetInstanceProcAddr(context, "vkDestroyBufferView");
    ptr_vkDestroyCommandPool = (PFN_vkDestroyCommandPool) ptr_vkGetInstanceProcAddr(context, "vkDestroyCommandPool");
    ptr_vkDestroyCuFunctionNVX = (PFN_vkDestroyCuFunctionNVX) ptr_vkGetInstanceProcAddr(context, "vkDestroyCuFunctionNVX");
    ptr_vkDestroyCuModuleNVX = (PFN_vkDestroyCuModuleNVX) ptr_vkGetInstanceProcAddr(context, "vkDestroyCuModuleNVX");
    ptr_vkDestroyDataGraphPipelineSessionARM = (PFN_vkDestroyDataGraphPipelineSessionARM) ptr_vkGetInstanceProcAddr(context, "vkDestroyDataGraphPipelineSessionARM");
    ptr_vkDestroyDebugReportCallbackEXT = (PFN_vkDestroyDebugReportCallbackEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyDebugReportCallbackEXT");
    ptr_vkDestroyDebugUtilsMessengerEXT = (PFN_vkDestroyDebugUtilsMessengerEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyDebugUtilsMessengerEXT");
    ptr_vkDestroyDeferredOperationKHR = (PFN_vkDestroyDeferredOperationKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroyDeferredOperationKHR");
    ptr_vkDestroyDescriptorPool = (PFN_vkDestroyDescriptorPool) ptr_vkGetInstanceProcAddr(context, "vkDestroyDescriptorPool");
    ptr_vkDestroyDescriptorSetLayout = (PFN_vkDestroyDescriptorSetLayout) ptr_vkGetInstanceProcAddr(context, "vkDestroyDescriptorSetLayout");
    ptr_vkDestroyDescriptorUpdateTemplate = (PFN_vkDestroyDescriptorUpdateTemplate) ptr_vkGetInstanceProcAddr(context, "vkDestroyDescriptorUpdateTemplate");
    ptr_vkDestroyDevice = (PFN_vkDestroyDevice) ptr_vkGetInstanceProcAddr(context, "vkDestroyDevice");
    ptr_vkDestroyEvent = (PFN_vkDestroyEvent) ptr_vkGetInstanceProcAddr(context, "vkDestroyEvent");
    ptr_vkDestroyExternalComputeQueueNV = (PFN_vkDestroyExternalComputeQueueNV) ptr_vkGetInstanceProcAddr(context, "vkDestroyExternalComputeQueueNV");
    ptr_vkDestroyFence = (PFN_vkDestroyFence) ptr_vkGetInstanceProcAddr(context, "vkDestroyFence");
    ptr_vkDestroyFramebuffer = (PFN_vkDestroyFramebuffer) ptr_vkGetInstanceProcAddr(context, "vkDestroyFramebuffer");
    ptr_vkDestroyImage = (PFN_vkDestroyImage) ptr_vkGetInstanceProcAddr(context, "vkDestroyImage");
    ptr_vkDestroyImageView = (PFN_vkDestroyImageView) ptr_vkGetInstanceProcAddr(context, "vkDestroyImageView");
    ptr_vkDestroyIndirectCommandsLayoutEXT = (PFN_vkDestroyIndirectCommandsLayoutEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyIndirectCommandsLayoutEXT");
    ptr_vkDestroyIndirectCommandsLayoutNV = (PFN_vkDestroyIndirectCommandsLayoutNV) ptr_vkGetInstanceProcAddr(context, "vkDestroyIndirectCommandsLayoutNV");
    ptr_vkDestroyIndirectExecutionSetEXT = (PFN_vkDestroyIndirectExecutionSetEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyIndirectExecutionSetEXT");
    ptr_vkDestroyInstance = (PFN_vkDestroyInstance) ptr_vkGetInstanceProcAddr(context, "vkDestroyInstance");
    ptr_vkDestroyMicromapEXT = (PFN_vkDestroyMicromapEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyMicromapEXT");
    ptr_vkDestroyOpticalFlowSessionNV = (PFN_vkDestroyOpticalFlowSessionNV) ptr_vkGetInstanceProcAddr(context, "vkDestroyOpticalFlowSessionNV");
    ptr_vkDestroyPipeline = (PFN_vkDestroyPipeline) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipeline");
    ptr_vkDestroyPipelineBinaryKHR = (PFN_vkDestroyPipelineBinaryKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipelineBinaryKHR");
    ptr_vkDestroyPipelineCache = (PFN_vkDestroyPipelineCache) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipelineCache");
    ptr_vkDestroyPipelineLayout = (PFN_vkDestroyPipelineLayout) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipelineLayout");
    ptr_vkDestroyPrivateDataSlot = (PFN_vkDestroyPrivateDataSlot) ptr_vkGetInstanceProcAddr(context, "vkDestroyPrivateDataSlot");
    ptr_vkDestroyQueryPool = (PFN_vkDestroyQueryPool) ptr_vkGetInstanceProcAddr(context, "vkDestroyQueryPool");
    ptr_vkDestroyRenderPass = (PFN_vkDestroyRenderPass) ptr_vkGetInstanceProcAddr(context, "vkDestroyRenderPass");
    ptr_vkDestroySampler = (PFN_vkDestroySampler) ptr_vkGetInstanceProcAddr(context, "vkDestroySampler");
    ptr_vkDestroySamplerYcbcrConversion = (PFN_vkDestroySamplerYcbcrConversion) ptr_vkGetInstanceProcAddr(context, "vkDestroySamplerYcbcrConversion");
    ptr_vkDestroySemaphore = (PFN_vkDestroySemaphore) ptr_vkGetInstanceProcAddr(context, "vkDestroySemaphore");
    ptr_vkDestroyShaderEXT = (PFN_vkDestroyShaderEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyShaderEXT");
    ptr_vkDestroyShaderModule = (PFN_vkDestroyShaderModule) ptr_vkGetInstanceProcAddr(context, "vkDestroyShaderModule");
    ptr_vkDestroySurfaceKHR = (PFN_vkDestroySurfaceKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroySurfaceKHR");
    ptr_vkDestroySwapchainKHR = (PFN_vkDestroySwapchainKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroySwapchainKHR");
    ptr_vkDestroyTensorARM = (PFN_vkDestroyTensorARM) ptr_vkGetInstanceProcAddr(context, "vkDestroyTensorARM");
    ptr_vkDestroyTensorViewARM = (PFN_vkDestroyTensorViewARM) ptr_vkGetInstanceProcAddr(context, "vkDestroyTensorViewARM");
    ptr_vkDestroyValidationCacheEXT = (PFN_vkDestroyValidationCacheEXT) ptr_vkGetInstanceProcAddr(context, "vkDestroyValidationCacheEXT");
    ptr_vkDestroyVideoSessionKHR = (PFN_vkDestroyVideoSessionKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroyVideoSessionKHR");
    ptr_vkDestroyVideoSessionParametersKHR = (PFN_vkDestroyVideoSessionParametersKHR) ptr_vkGetInstanceProcAddr(context, "vkDestroyVideoSessionParametersKHR");
    ptr_vkDeviceWaitIdle = (PFN_vkDeviceWaitIdle) ptr_vkGetInstanceProcAddr(context, "vkDeviceWaitIdle");
    ptr_vkDisplayPowerControlEXT = (PFN_vkDisplayPowerControlEXT) ptr_vkGetInstanceProcAddr(context, "vkDisplayPowerControlEXT");
    ptr_vkEndCommandBuffer = (PFN_vkEndCommandBuffer) ptr_vkGetInstanceProcAddr(context, "vkEndCommandBuffer");
    ptr_vkEnumerateDeviceExtensionProperties = (PFN_vkEnumerateDeviceExtensionProperties) ptr_vkGetInstanceProcAddr(context, "vkEnumerateDeviceExtensionProperties");
    ptr_vkEnumerateDeviceLayerProperties = (PFN_vkEnumerateDeviceLayerProperties) ptr_vkGetInstanceProcAddr(context, "vkEnumerateDeviceLayerProperties");
    ptr_vkEnumeratePhysicalDeviceGroups = (PFN_vkEnumeratePhysicalDeviceGroups) ptr_vkGetInstanceProcAddr(context, "vkEnumeratePhysicalDeviceGroups");
    ptr_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR = (PFN_vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR) ptr_vkGetInstanceProcAddr(context, "vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR");
    ptr_vkEnumeratePhysicalDevices = (PFN_vkEnumeratePhysicalDevices) ptr_vkGetInstanceProcAddr(context, "vkEnumeratePhysicalDevices");
    ptr_vkExportMetalObjectsEXT = (PFN_vkExportMetalObjectsEXT) ptr_vkGetInstanceProcAddr(context, "vkExportMetalObjectsEXT");
    ptr_vkFlushMappedMemoryRanges = (PFN_vkFlushMappedMemoryRanges) ptr_vkGetInstanceProcAddr(context, "vkFlushMappedMemoryRanges");
    ptr_vkFreeCommandBuffers = (PFN_vkFreeCommandBuffers) ptr_vkGetInstanceProcAddr(context, "vkFreeCommandBuffers");
    ptr_vkFreeDescriptorSets = (PFN_vkFreeDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkFreeDescriptorSets");
    ptr_vkFreeMemory = (PFN_vkFreeMemory) ptr_vkGetInstanceProcAddr(context, "vkFreeMemory");
    ptr_vkGetAccelerationStructureBuildSizesKHR = (PFN_vkGetAccelerationStructureBuildSizesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetAccelerationStructureBuildSizesKHR");
    ptr_vkGetAccelerationStructureDeviceAddressKHR = (PFN_vkGetAccelerationStructureDeviceAddressKHR) ptr_vkGetInstanceProcAddr(context, "vkGetAccelerationStructureDeviceAddressKHR");
    ptr_vkGetAccelerationStructureHandleNV = (PFN_vkGetAccelerationStructureHandleNV) ptr_vkGetInstanceProcAddr(context, "vkGetAccelerationStructureHandleNV");
    ptr_vkGetAccelerationStructureMemoryRequirementsNV = (PFN_vkGetAccelerationStructureMemoryRequirementsNV) ptr_vkGetInstanceProcAddr(context, "vkGetAccelerationStructureMemoryRequirementsNV");
    ptr_vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT = (PFN_vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT");
    ptr_vkGetBufferDeviceAddress = (PFN_vkGetBufferDeviceAddress) ptr_vkGetInstanceProcAddr(context, "vkGetBufferDeviceAddress");
    ptr_vkGetBufferMemoryRequirements = (PFN_vkGetBufferMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetBufferMemoryRequirements");
    ptr_vkGetBufferMemoryRequirements2 = (PFN_vkGetBufferMemoryRequirements2) ptr_vkGetInstanceProcAddr(context, "vkGetBufferMemoryRequirements2");
    ptr_vkGetBufferOpaqueCaptureAddress = (PFN_vkGetBufferOpaqueCaptureAddress) ptr_vkGetInstanceProcAddr(context, "vkGetBufferOpaqueCaptureAddress");
    ptr_vkGetBufferOpaqueCaptureDescriptorDataEXT = (PFN_vkGetBufferOpaqueCaptureDescriptorDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetBufferOpaqueCaptureDescriptorDataEXT");
    ptr_vkGetCalibratedTimestampsKHR = (PFN_vkGetCalibratedTimestampsKHR) ptr_vkGetInstanceProcAddr(context, "vkGetCalibratedTimestampsKHR");
    ptr_vkGetClusterAccelerationStructureBuildSizesNV = (PFN_vkGetClusterAccelerationStructureBuildSizesNV) ptr_vkGetInstanceProcAddr(context, "vkGetClusterAccelerationStructureBuildSizesNV");
    ptr_vkGetDataGraphPipelineAvailablePropertiesARM = (PFN_vkGetDataGraphPipelineAvailablePropertiesARM) ptr_vkGetInstanceProcAddr(context, "vkGetDataGraphPipelineAvailablePropertiesARM");
    ptr_vkGetDataGraphPipelinePropertiesARM = (PFN_vkGetDataGraphPipelinePropertiesARM) ptr_vkGetInstanceProcAddr(context, "vkGetDataGraphPipelinePropertiesARM");
    ptr_vkGetDataGraphPipelineSessionBindPointRequirementsARM = (PFN_vkGetDataGraphPipelineSessionBindPointRequirementsARM) ptr_vkGetInstanceProcAddr(context, "vkGetDataGraphPipelineSessionBindPointRequirementsARM");
    ptr_vkGetDataGraphPipelineSessionMemoryRequirementsARM = (PFN_vkGetDataGraphPipelineSessionMemoryRequirementsARM) ptr_vkGetInstanceProcAddr(context, "vkGetDataGraphPipelineSessionMemoryRequirementsARM");
    ptr_vkGetDeferredOperationMaxConcurrencyKHR = (PFN_vkGetDeferredOperationMaxConcurrencyKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDeferredOperationMaxConcurrencyKHR");
    ptr_vkGetDeferredOperationResultKHR = (PFN_vkGetDeferredOperationResultKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDeferredOperationResultKHR");
    ptr_vkGetDescriptorEXT = (PFN_vkGetDescriptorEXT) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorEXT");
    ptr_vkGetDescriptorSetHostMappingVALVE = (PFN_vkGetDescriptorSetHostMappingVALVE) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorSetHostMappingVALVE");
    ptr_vkGetDescriptorSetLayoutBindingOffsetEXT = (PFN_vkGetDescriptorSetLayoutBindingOffsetEXT) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorSetLayoutBindingOffsetEXT");
    ptr_vkGetDescriptorSetLayoutHostMappingInfoVALVE = (PFN_vkGetDescriptorSetLayoutHostMappingInfoVALVE) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorSetLayoutHostMappingInfoVALVE");
    ptr_vkGetDescriptorSetLayoutSizeEXT = (PFN_vkGetDescriptorSetLayoutSizeEXT) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorSetLayoutSizeEXT");
    ptr_vkGetDescriptorSetLayoutSupport = (PFN_vkGetDescriptorSetLayoutSupport) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorSetLayoutSupport");
    ptr_vkGetDeviceAccelerationStructureCompatibilityKHR = (PFN_vkGetDeviceAccelerationStructureCompatibilityKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceAccelerationStructureCompatibilityKHR");
    ptr_vkGetDeviceBufferMemoryRequirements = (PFN_vkGetDeviceBufferMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceBufferMemoryRequirements");
    ptr_vkGetDeviceFaultInfoEXT = (PFN_vkGetDeviceFaultInfoEXT) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceFaultInfoEXT");
    ptr_vkGetDeviceGroupPeerMemoryFeatures = (PFN_vkGetDeviceGroupPeerMemoryFeatures) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceGroupPeerMemoryFeatures");
    ptr_vkGetDeviceGroupPresentCapabilitiesKHR = (PFN_vkGetDeviceGroupPresentCapabilitiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceGroupPresentCapabilitiesKHR");
    ptr_vkGetDeviceGroupSurfacePresentModes2EXT = (PFN_vkGetDeviceGroupSurfacePresentModes2EXT) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceGroupSurfacePresentModes2EXT");
    ptr_vkGetDeviceGroupSurfacePresentModesKHR = (PFN_vkGetDeviceGroupSurfacePresentModesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceGroupSurfacePresentModesKHR");
    ptr_vkGetDeviceImageMemoryRequirements = (PFN_vkGetDeviceImageMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceImageMemoryRequirements");
    ptr_vkGetDeviceImageSparseMemoryRequirements = (PFN_vkGetDeviceImageSparseMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceImageSparseMemoryRequirements");
    ptr_vkGetDeviceImageSubresourceLayout = (PFN_vkGetDeviceImageSubresourceLayout) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceImageSubresourceLayout");
    ptr_vkGetDeviceMemoryCommitment = (PFN_vkGetDeviceMemoryCommitment) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceMemoryCommitment");
    ptr_vkGetDeviceMemoryOpaqueCaptureAddress = (PFN_vkGetDeviceMemoryOpaqueCaptureAddress) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceMemoryOpaqueCaptureAddress");
    ptr_vkGetDeviceMicromapCompatibilityEXT = (PFN_vkGetDeviceMicromapCompatibilityEXT) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceMicromapCompatibilityEXT");
    ptr_vkGetDeviceProcAddr = (PFN_vkGetDeviceProcAddr) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceProcAddr");
    ptr_vkGetDeviceQueue = (PFN_vkGetDeviceQueue) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceQueue");
    ptr_vkGetDeviceQueue2 = (PFN_vkGetDeviceQueue2) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceQueue2");
    ptr_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI = (PFN_vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI");
    ptr_vkGetDeviceTensorMemoryRequirementsARM = (PFN_vkGetDeviceTensorMemoryRequirementsARM) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceTensorMemoryRequirementsARM");
    ptr_vkGetDisplayModeProperties2KHR = (PFN_vkGetDisplayModeProperties2KHR) ptr_vkGetInstanceProcAddr(context, "vkGetDisplayModeProperties2KHR");
    ptr_vkGetDisplayModePropertiesKHR = (PFN_vkGetDisplayModePropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDisplayModePropertiesKHR");
    ptr_vkGetDisplayPlaneCapabilities2KHR = (PFN_vkGetDisplayPlaneCapabilities2KHR) ptr_vkGetInstanceProcAddr(context, "vkGetDisplayPlaneCapabilities2KHR");
    ptr_vkGetDisplayPlaneCapabilitiesKHR = (PFN_vkGetDisplayPlaneCapabilitiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDisplayPlaneCapabilitiesKHR");
    ptr_vkGetDisplayPlaneSupportedDisplaysKHR = (PFN_vkGetDisplayPlaneSupportedDisplaysKHR) ptr_vkGetInstanceProcAddr(context, "vkGetDisplayPlaneSupportedDisplaysKHR");
    ptr_vkGetDrmDisplayEXT = (PFN_vkGetDrmDisplayEXT) ptr_vkGetInstanceProcAddr(context, "vkGetDrmDisplayEXT");
    ptr_vkGetDynamicRenderingTilePropertiesQCOM = (PFN_vkGetDynamicRenderingTilePropertiesQCOM) ptr_vkGetInstanceProcAddr(context, "vkGetDynamicRenderingTilePropertiesQCOM");
    ptr_vkGetEncodedVideoSessionParametersKHR = (PFN_vkGetEncodedVideoSessionParametersKHR) ptr_vkGetInstanceProcAddr(context, "vkGetEncodedVideoSessionParametersKHR");
    ptr_vkGetEventStatus = (PFN_vkGetEventStatus) ptr_vkGetInstanceProcAddr(context, "vkGetEventStatus");
    ptr_vkGetExternalComputeQueueDataNV = (PFN_vkGetExternalComputeQueueDataNV) ptr_vkGetInstanceProcAddr(context, "vkGetExternalComputeQueueDataNV");
    ptr_vkGetFenceFdKHR = (PFN_vkGetFenceFdKHR) ptr_vkGetInstanceProcAddr(context, "vkGetFenceFdKHR");
    ptr_vkGetFenceStatus = (PFN_vkGetFenceStatus) ptr_vkGetInstanceProcAddr(context, "vkGetFenceStatus");
    ptr_vkGetFenceWin32HandleKHR = (PFN_vkGetFenceWin32HandleKHR) ptr_vkGetInstanceProcAddr(context, "vkGetFenceWin32HandleKHR");
    ptr_vkGetFramebufferTilePropertiesQCOM = (PFN_vkGetFramebufferTilePropertiesQCOM) ptr_vkGetInstanceProcAddr(context, "vkGetFramebufferTilePropertiesQCOM");
    ptr_vkGetGeneratedCommandsMemoryRequirementsEXT = (PFN_vkGetGeneratedCommandsMemoryRequirementsEXT) ptr_vkGetInstanceProcAddr(context, "vkGetGeneratedCommandsMemoryRequirementsEXT");
    ptr_vkGetGeneratedCommandsMemoryRequirementsNV = (PFN_vkGetGeneratedCommandsMemoryRequirementsNV) ptr_vkGetInstanceProcAddr(context, "vkGetGeneratedCommandsMemoryRequirementsNV");
    ptr_vkGetImageDrmFormatModifierPropertiesEXT = (PFN_vkGetImageDrmFormatModifierPropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkGetImageDrmFormatModifierPropertiesEXT");
    ptr_vkGetImageMemoryRequirements = (PFN_vkGetImageMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetImageMemoryRequirements");
    ptr_vkGetImageMemoryRequirements2 = (PFN_vkGetImageMemoryRequirements2) ptr_vkGetInstanceProcAddr(context, "vkGetImageMemoryRequirements2");
    ptr_vkGetImageOpaqueCaptureDescriptorDataEXT = (PFN_vkGetImageOpaqueCaptureDescriptorDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetImageOpaqueCaptureDescriptorDataEXT");
    ptr_vkGetImageSparseMemoryRequirements = (PFN_vkGetImageSparseMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetImageSparseMemoryRequirements");
    ptr_vkGetImageSparseMemoryRequirements2 = (PFN_vkGetImageSparseMemoryRequirements2) ptr_vkGetInstanceProcAddr(context, "vkGetImageSparseMemoryRequirements2");
    ptr_vkGetImageSubresourceLayout = (PFN_vkGetImageSubresourceLayout) ptr_vkGetInstanceProcAddr(context, "vkGetImageSubresourceLayout");
    ptr_vkGetImageSubresourceLayout2 = (PFN_vkGetImageSubresourceLayout2) ptr_vkGetInstanceProcAddr(context, "vkGetImageSubresourceLayout2");
    ptr_vkGetImageViewAddressNVX = (PFN_vkGetImageViewAddressNVX) ptr_vkGetInstanceProcAddr(context, "vkGetImageViewAddressNVX");
    ptr_vkGetImageViewHandle64NVX = (PFN_vkGetImageViewHandle64NVX) ptr_vkGetInstanceProcAddr(context, "vkGetImageViewHandle64NVX");
    ptr_vkGetImageViewHandleNVX = (PFN_vkGetImageViewHandleNVX) ptr_vkGetInstanceProcAddr(context, "vkGetImageViewHandleNVX");
    ptr_vkGetImageViewOpaqueCaptureDescriptorDataEXT = (PFN_vkGetImageViewOpaqueCaptureDescriptorDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetImageViewOpaqueCaptureDescriptorDataEXT");
    ptr_vkGetInstanceProcAddr = (PFN_vkGetInstanceProcAddr) ptr_vkGetInstanceProcAddr(context, "vkGetInstanceProcAddr");
    ptr_vkGetLatencyTimingsNV = (PFN_vkGetLatencyTimingsNV) ptr_vkGetInstanceProcAddr(context, "vkGetLatencyTimingsNV");
    ptr_vkGetMemoryFdKHR = (PFN_vkGetMemoryFdKHR) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryFdKHR");
    ptr_vkGetMemoryFdPropertiesKHR = (PFN_vkGetMemoryFdPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryFdPropertiesKHR");
    ptr_vkGetMemoryHostPointerPropertiesEXT = (PFN_vkGetMemoryHostPointerPropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryHostPointerPropertiesEXT");
    ptr_vkGetMemoryMetalHandleEXT = (PFN_vkGetMemoryMetalHandleEXT) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryMetalHandleEXT");
    ptr_vkGetMemoryMetalHandlePropertiesEXT = (PFN_vkGetMemoryMetalHandlePropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryMetalHandlePropertiesEXT");
    ptr_vkGetMemoryRemoteAddressNV = (PFN_vkGetMemoryRemoteAddressNV) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryRemoteAddressNV");
    ptr_vkGetMemoryWin32HandleKHR = (PFN_vkGetMemoryWin32HandleKHR) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryWin32HandleKHR");
    ptr_vkGetMemoryWin32HandleNV = (PFN_vkGetMemoryWin32HandleNV) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryWin32HandleNV");
    ptr_vkGetMemoryWin32HandlePropertiesKHR = (PFN_vkGetMemoryWin32HandlePropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetMemoryWin32HandlePropertiesKHR");
    ptr_vkGetMicromapBuildSizesEXT = (PFN_vkGetMicromapBuildSizesEXT) ptr_vkGetInstanceProcAddr(context, "vkGetMicromapBuildSizesEXT");
    ptr_vkGetPartitionedAccelerationStructuresBuildSizesNV = (PFN_vkGetPartitionedAccelerationStructuresBuildSizesNV) ptr_vkGetInstanceProcAddr(context, "vkGetPartitionedAccelerationStructuresBuildSizesNV");
    ptr_vkGetPastPresentationTimingGOOGLE = (PFN_vkGetPastPresentationTimingGOOGLE) ptr_vkGetInstanceProcAddr(context, "vkGetPastPresentationTimingGOOGLE");
    ptr_vkGetPerformanceParameterINTEL = (PFN_vkGetPerformanceParameterINTEL) ptr_vkGetInstanceProcAddr(context, "vkGetPerformanceParameterINTEL");
    ptr_vkGetPhysicalDeviceCalibrateableTimeDomainsKHR = (PFN_vkGetPhysicalDeviceCalibrateableTimeDomainsKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceCalibrateableTimeDomainsKHR");
    ptr_vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV = (PFN_vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV");
    ptr_vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR = (PFN_vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR");
    ptr_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV = (PFN_vkGetPhysicalDeviceCooperativeMatrixPropertiesNV) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceCooperativeMatrixPropertiesNV");
    ptr_vkGetPhysicalDeviceCooperativeVectorPropertiesNV = (PFN_vkGetPhysicalDeviceCooperativeVectorPropertiesNV) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceCooperativeVectorPropertiesNV");
    ptr_vkGetPhysicalDeviceDisplayPlaneProperties2KHR = (PFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceDisplayPlaneProperties2KHR");
    ptr_vkGetPhysicalDeviceDisplayPlanePropertiesKHR = (PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceDisplayPlanePropertiesKHR");
    ptr_vkGetPhysicalDeviceDisplayProperties2KHR = (PFN_vkGetPhysicalDeviceDisplayProperties2KHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceDisplayProperties2KHR");
    ptr_vkGetPhysicalDeviceDisplayPropertiesKHR = (PFN_vkGetPhysicalDeviceDisplayPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceDisplayPropertiesKHR");
    ptr_vkGetPhysicalDeviceExternalBufferProperties = (PFN_vkGetPhysicalDeviceExternalBufferProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalBufferProperties");
    ptr_vkGetPhysicalDeviceExternalFenceProperties = (PFN_vkGetPhysicalDeviceExternalFenceProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalFenceProperties");
    ptr_vkGetPhysicalDeviceExternalImageFormatPropertiesNV = (PFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalImageFormatPropertiesNV");
    ptr_vkGetPhysicalDeviceExternalSemaphoreProperties = (PFN_vkGetPhysicalDeviceExternalSemaphoreProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalSemaphoreProperties");
    ptr_vkGetPhysicalDeviceExternalTensorPropertiesARM = (PFN_vkGetPhysicalDeviceExternalTensorPropertiesARM) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalTensorPropertiesARM");
    ptr_vkGetPhysicalDeviceFeatures = (PFN_vkGetPhysicalDeviceFeatures) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFeatures");
    ptr_vkGetPhysicalDeviceFeatures2 = (PFN_vkGetPhysicalDeviceFeatures2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFeatures2");
    ptr_vkGetPhysicalDeviceFormatProperties = (PFN_vkGetPhysicalDeviceFormatProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFormatProperties");
    ptr_vkGetPhysicalDeviceFormatProperties2 = (PFN_vkGetPhysicalDeviceFormatProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFormatProperties2");
    ptr_vkGetPhysicalDeviceFragmentShadingRatesKHR = (PFN_vkGetPhysicalDeviceFragmentShadingRatesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFragmentShadingRatesKHR");
    ptr_vkGetPhysicalDeviceImageFormatProperties = (PFN_vkGetPhysicalDeviceImageFormatProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceImageFormatProperties");
    ptr_vkGetPhysicalDeviceImageFormatProperties2 = (PFN_vkGetPhysicalDeviceImageFormatProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceImageFormatProperties2");
    ptr_vkGetPhysicalDeviceMemoryProperties = (PFN_vkGetPhysicalDeviceMemoryProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceMemoryProperties");
    ptr_vkGetPhysicalDeviceMemoryProperties2 = (PFN_vkGetPhysicalDeviceMemoryProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceMemoryProperties2");
    ptr_vkGetPhysicalDeviceMultisamplePropertiesEXT = (PFN_vkGetPhysicalDeviceMultisamplePropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceMultisamplePropertiesEXT");
    ptr_vkGetPhysicalDeviceOpticalFlowImageFormatsNV = (PFN_vkGetPhysicalDeviceOpticalFlowImageFormatsNV) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceOpticalFlowImageFormatsNV");
    ptr_vkGetPhysicalDevicePresentRectanglesKHR = (PFN_vkGetPhysicalDevicePresentRectanglesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDevicePresentRectanglesKHR");
    ptr_vkGetPhysicalDeviceProperties = (PFN_vkGetPhysicalDeviceProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceProperties");
    ptr_vkGetPhysicalDeviceProperties2 = (PFN_vkGetPhysicalDeviceProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceProperties2");
    ptr_vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM = (PFN_vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM");
    ptr_vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM = (PFN_vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM");
    ptr_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR = (PFN_vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR");
    ptr_vkGetPhysicalDeviceQueueFamilyProperties = (PFN_vkGetPhysicalDeviceQueueFamilyProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyProperties");
    ptr_vkGetPhysicalDeviceQueueFamilyProperties2 = (PFN_vkGetPhysicalDeviceQueueFamilyProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyProperties2");
    ptr_vkGetPhysicalDeviceSparseImageFormatProperties = (PFN_vkGetPhysicalDeviceSparseImageFormatProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSparseImageFormatProperties");
    ptr_vkGetPhysicalDeviceSparseImageFormatProperties2 = (PFN_vkGetPhysicalDeviceSparseImageFormatProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSparseImageFormatProperties2");
    ptr_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV = (PFN_vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV");
    ptr_vkGetPhysicalDeviceSurfaceCapabilities2EXT = (PFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfaceCapabilities2EXT");
    ptr_vkGetPhysicalDeviceSurfaceCapabilities2KHR = (PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfaceCapabilities2KHR");
    ptr_vkGetPhysicalDeviceSurfaceCapabilitiesKHR = (PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfaceCapabilitiesKHR");
    ptr_vkGetPhysicalDeviceSurfaceFormats2KHR = (PFN_vkGetPhysicalDeviceSurfaceFormats2KHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfaceFormats2KHR");
    ptr_vkGetPhysicalDeviceSurfaceFormatsKHR = (PFN_vkGetPhysicalDeviceSurfaceFormatsKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfaceFormatsKHR");
    ptr_vkGetPhysicalDeviceSurfacePresentModes2EXT = (PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfacePresentModes2EXT");
    ptr_vkGetPhysicalDeviceSurfacePresentModesKHR = (PFN_vkGetPhysicalDeviceSurfacePresentModesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfacePresentModesKHR");
    ptr_vkGetPhysicalDeviceSurfaceSupportKHR = (PFN_vkGetPhysicalDeviceSurfaceSupportKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSurfaceSupportKHR");
    ptr_vkGetPhysicalDeviceToolProperties = (PFN_vkGetPhysicalDeviceToolProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceToolProperties");
    ptr_vkGetPhysicalDeviceVideoCapabilitiesKHR = (PFN_vkGetPhysicalDeviceVideoCapabilitiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceVideoCapabilitiesKHR");
    ptr_vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR = (PFN_vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR");
    ptr_vkGetPhysicalDeviceVideoFormatPropertiesKHR = (PFN_vkGetPhysicalDeviceVideoFormatPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceVideoFormatPropertiesKHR");
    ptr_vkGetPhysicalDeviceWin32PresentationSupportKHR = (PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceWin32PresentationSupportKHR");
    ptr_vkGetPipelineBinaryDataKHR = (PFN_vkGetPipelineBinaryDataKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineBinaryDataKHR");
    ptr_vkGetPipelineCacheData = (PFN_vkGetPipelineCacheData) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineCacheData");
    ptr_vkGetPipelineExecutableInternalRepresentationsKHR = (PFN_vkGetPipelineExecutableInternalRepresentationsKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineExecutableInternalRepresentationsKHR");
    ptr_vkGetPipelineExecutablePropertiesKHR = (PFN_vkGetPipelineExecutablePropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineExecutablePropertiesKHR");
    ptr_vkGetPipelineExecutableStatisticsKHR = (PFN_vkGetPipelineExecutableStatisticsKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineExecutableStatisticsKHR");
    ptr_vkGetPipelineIndirectDeviceAddressNV = (PFN_vkGetPipelineIndirectDeviceAddressNV) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineIndirectDeviceAddressNV");
    ptr_vkGetPipelineIndirectMemoryRequirementsNV = (PFN_vkGetPipelineIndirectMemoryRequirementsNV) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineIndirectMemoryRequirementsNV");
    ptr_vkGetPipelineKeyKHR = (PFN_vkGetPipelineKeyKHR) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineKeyKHR");
    ptr_vkGetPipelinePropertiesEXT = (PFN_vkGetPipelinePropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkGetPipelinePropertiesEXT");
    ptr_vkGetPrivateData = (PFN_vkGetPrivateData) ptr_vkGetInstanceProcAddr(context, "vkGetPrivateData");
    ptr_vkGetQueryPoolResults = (PFN_vkGetQueryPoolResults) ptr_vkGetInstanceProcAddr(context, "vkGetQueryPoolResults");
    ptr_vkGetQueueCheckpointData2NV = (PFN_vkGetQueueCheckpointData2NV) ptr_vkGetInstanceProcAddr(context, "vkGetQueueCheckpointData2NV");
    ptr_vkGetQueueCheckpointDataNV = (PFN_vkGetQueueCheckpointDataNV) ptr_vkGetInstanceProcAddr(context, "vkGetQueueCheckpointDataNV");
    ptr_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR = (PFN_vkGetRayTracingCaptureReplayShaderGroupHandlesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetRayTracingCaptureReplayShaderGroupHandlesKHR");
    ptr_vkGetRayTracingShaderGroupHandlesKHR = (PFN_vkGetRayTracingShaderGroupHandlesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetRayTracingShaderGroupHandlesKHR");
    ptr_vkGetRayTracingShaderGroupStackSizeKHR = (PFN_vkGetRayTracingShaderGroupStackSizeKHR) ptr_vkGetInstanceProcAddr(context, "vkGetRayTracingShaderGroupStackSizeKHR");
    ptr_vkGetRefreshCycleDurationGOOGLE = (PFN_vkGetRefreshCycleDurationGOOGLE) ptr_vkGetInstanceProcAddr(context, "vkGetRefreshCycleDurationGOOGLE");
    ptr_vkGetRenderAreaGranularity = (PFN_vkGetRenderAreaGranularity) ptr_vkGetInstanceProcAddr(context, "vkGetRenderAreaGranularity");
    ptr_vkGetRenderingAreaGranularity = (PFN_vkGetRenderingAreaGranularity) ptr_vkGetInstanceProcAddr(context, "vkGetRenderingAreaGranularity");
    ptr_vkGetSamplerOpaqueCaptureDescriptorDataEXT = (PFN_vkGetSamplerOpaqueCaptureDescriptorDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetSamplerOpaqueCaptureDescriptorDataEXT");
    ptr_vkGetSemaphoreCounterValue = (PFN_vkGetSemaphoreCounterValue) ptr_vkGetInstanceProcAddr(context, "vkGetSemaphoreCounterValue");
    ptr_vkGetSemaphoreFdKHR = (PFN_vkGetSemaphoreFdKHR) ptr_vkGetInstanceProcAddr(context, "vkGetSemaphoreFdKHR");
    ptr_vkGetSemaphoreWin32HandleKHR = (PFN_vkGetSemaphoreWin32HandleKHR) ptr_vkGetInstanceProcAddr(context, "vkGetSemaphoreWin32HandleKHR");
    ptr_vkGetShaderBinaryDataEXT = (PFN_vkGetShaderBinaryDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetShaderBinaryDataEXT");
    ptr_vkGetShaderInfoAMD = (PFN_vkGetShaderInfoAMD) ptr_vkGetInstanceProcAddr(context, "vkGetShaderInfoAMD");
    ptr_vkGetShaderModuleCreateInfoIdentifierEXT = (PFN_vkGetShaderModuleCreateInfoIdentifierEXT) ptr_vkGetInstanceProcAddr(context, "vkGetShaderModuleCreateInfoIdentifierEXT");
    ptr_vkGetShaderModuleIdentifierEXT = (PFN_vkGetShaderModuleIdentifierEXT) ptr_vkGetInstanceProcAddr(context, "vkGetShaderModuleIdentifierEXT");
    ptr_vkGetSwapchainCounterEXT = (PFN_vkGetSwapchainCounterEXT) ptr_vkGetInstanceProcAddr(context, "vkGetSwapchainCounterEXT");
    ptr_vkGetSwapchainImagesKHR = (PFN_vkGetSwapchainImagesKHR) ptr_vkGetInstanceProcAddr(context, "vkGetSwapchainImagesKHR");
    ptr_vkGetSwapchainStatusKHR = (PFN_vkGetSwapchainStatusKHR) ptr_vkGetInstanceProcAddr(context, "vkGetSwapchainStatusKHR");
    ptr_vkGetTensorMemoryRequirementsARM = (PFN_vkGetTensorMemoryRequirementsARM) ptr_vkGetInstanceProcAddr(context, "vkGetTensorMemoryRequirementsARM");
    ptr_vkGetTensorOpaqueCaptureDescriptorDataARM = (PFN_vkGetTensorOpaqueCaptureDescriptorDataARM) ptr_vkGetInstanceProcAddr(context, "vkGetTensorOpaqueCaptureDescriptorDataARM");
    ptr_vkGetTensorViewOpaqueCaptureDescriptorDataARM = (PFN_vkGetTensorViewOpaqueCaptureDescriptorDataARM) ptr_vkGetInstanceProcAddr(context, "vkGetTensorViewOpaqueCaptureDescriptorDataARM");
    ptr_vkGetValidationCacheDataEXT = (PFN_vkGetValidationCacheDataEXT) ptr_vkGetInstanceProcAddr(context, "vkGetValidationCacheDataEXT");
    ptr_vkGetVideoSessionMemoryRequirementsKHR = (PFN_vkGetVideoSessionMemoryRequirementsKHR) ptr_vkGetInstanceProcAddr(context, "vkGetVideoSessionMemoryRequirementsKHR");
    ptr_vkGetWinrtDisplayNV = (PFN_vkGetWinrtDisplayNV) ptr_vkGetInstanceProcAddr(context, "vkGetWinrtDisplayNV");
    ptr_vkImportFenceFdKHR = (PFN_vkImportFenceFdKHR) ptr_vkGetInstanceProcAddr(context, "vkImportFenceFdKHR");
    ptr_vkImportFenceWin32HandleKHR = (PFN_vkImportFenceWin32HandleKHR) ptr_vkGetInstanceProcAddr(context, "vkImportFenceWin32HandleKHR");
    ptr_vkImportSemaphoreFdKHR = (PFN_vkImportSemaphoreFdKHR) ptr_vkGetInstanceProcAddr(context, "vkImportSemaphoreFdKHR");
    ptr_vkImportSemaphoreWin32HandleKHR = (PFN_vkImportSemaphoreWin32HandleKHR) ptr_vkGetInstanceProcAddr(context, "vkImportSemaphoreWin32HandleKHR");
    ptr_vkInitializePerformanceApiINTEL = (PFN_vkInitializePerformanceApiINTEL) ptr_vkGetInstanceProcAddr(context, "vkInitializePerformanceApiINTEL");
    ptr_vkInvalidateMappedMemoryRanges = (PFN_vkInvalidateMappedMemoryRanges) ptr_vkGetInstanceProcAddr(context, "vkInvalidateMappedMemoryRanges");
    ptr_vkLatencySleepNV = (PFN_vkLatencySleepNV) ptr_vkGetInstanceProcAddr(context, "vkLatencySleepNV");
    ptr_vkMapMemory = (PFN_vkMapMemory) ptr_vkGetInstanceProcAddr(context, "vkMapMemory");
    ptr_vkMapMemory2 = (PFN_vkMapMemory2) ptr_vkGetInstanceProcAddr(context, "vkMapMemory2");
    ptr_vkMergePipelineCaches = (PFN_vkMergePipelineCaches) ptr_vkGetInstanceProcAddr(context, "vkMergePipelineCaches");
    ptr_vkMergeValidationCachesEXT = (PFN_vkMergeValidationCachesEXT) ptr_vkGetInstanceProcAddr(context, "vkMergeValidationCachesEXT");
    ptr_vkQueueBeginDebugUtilsLabelEXT = (PFN_vkQueueBeginDebugUtilsLabelEXT) ptr_vkGetInstanceProcAddr(context, "vkQueueBeginDebugUtilsLabelEXT");
    ptr_vkQueueBindSparse = (PFN_vkQueueBindSparse) ptr_vkGetInstanceProcAddr(context, "vkQueueBindSparse");
    ptr_vkQueueEndDebugUtilsLabelEXT = (PFN_vkQueueEndDebugUtilsLabelEXT) ptr_vkGetInstanceProcAddr(context, "vkQueueEndDebugUtilsLabelEXT");
    ptr_vkQueueInsertDebugUtilsLabelEXT = (PFN_vkQueueInsertDebugUtilsLabelEXT) ptr_vkGetInstanceProcAddr(context, "vkQueueInsertDebugUtilsLabelEXT");
    ptr_vkQueueNotifyOutOfBandNV = (PFN_vkQueueNotifyOutOfBandNV) ptr_vkGetInstanceProcAddr(context, "vkQueueNotifyOutOfBandNV");
    ptr_vkQueuePresentKHR = (PFN_vkQueuePresentKHR) ptr_vkGetInstanceProcAddr(context, "vkQueuePresentKHR");
    ptr_vkQueueSetPerformanceConfigurationINTEL = (PFN_vkQueueSetPerformanceConfigurationINTEL) ptr_vkGetInstanceProcAddr(context, "vkQueueSetPerformanceConfigurationINTEL");
    ptr_vkQueueSubmit = (PFN_vkQueueSubmit) ptr_vkGetInstanceProcAddr(context, "vkQueueSubmit");
    ptr_vkQueueSubmit2 = (PFN_vkQueueSubmit2) ptr_vkGetInstanceProcAddr(context, "vkQueueSubmit2");
    ptr_vkQueueWaitIdle = (PFN_vkQueueWaitIdle) ptr_vkGetInstanceProcAddr(context, "vkQueueWaitIdle");
    ptr_vkRegisterDeviceEventEXT = (PFN_vkRegisterDeviceEventEXT) ptr_vkGetInstanceProcAddr(context, "vkRegisterDeviceEventEXT");
    ptr_vkRegisterDisplayEventEXT = (PFN_vkRegisterDisplayEventEXT) ptr_vkGetInstanceProcAddr(context, "vkRegisterDisplayEventEXT");
    ptr_vkReleaseCapturedPipelineDataKHR = (PFN_vkReleaseCapturedPipelineDataKHR) ptr_vkGetInstanceProcAddr(context, "vkReleaseCapturedPipelineDataKHR");
    ptr_vkReleaseDisplayEXT = (PFN_vkReleaseDisplayEXT) ptr_vkGetInstanceProcAddr(context, "vkReleaseDisplayEXT");
    ptr_vkReleaseFullScreenExclusiveModeEXT = (PFN_vkReleaseFullScreenExclusiveModeEXT) ptr_vkGetInstanceProcAddr(context, "vkReleaseFullScreenExclusiveModeEXT");
    ptr_vkReleasePerformanceConfigurationINTEL = (PFN_vkReleasePerformanceConfigurationINTEL) ptr_vkGetInstanceProcAddr(context, "vkReleasePerformanceConfigurationINTEL");
    ptr_vkReleaseProfilingLockKHR = (PFN_vkReleaseProfilingLockKHR) ptr_vkGetInstanceProcAddr(context, "vkReleaseProfilingLockKHR");
    ptr_vkReleaseSwapchainImagesKHR = (PFN_vkReleaseSwapchainImagesKHR) ptr_vkGetInstanceProcAddr(context, "vkReleaseSwapchainImagesKHR");
    ptr_vkResetCommandBuffer = (PFN_vkResetCommandBuffer) ptr_vkGetInstanceProcAddr(context, "vkResetCommandBuffer");
    ptr_vkResetCommandPool = (PFN_vkResetCommandPool) ptr_vkGetInstanceProcAddr(context, "vkResetCommandPool");
    ptr_vkResetDescriptorPool = (PFN_vkResetDescriptorPool) ptr_vkGetInstanceProcAddr(context, "vkResetDescriptorPool");
    ptr_vkResetEvent = (PFN_vkResetEvent) ptr_vkGetInstanceProcAddr(context, "vkResetEvent");
    ptr_vkResetFences = (PFN_vkResetFences) ptr_vkGetInstanceProcAddr(context, "vkResetFences");
    ptr_vkResetQueryPool = (PFN_vkResetQueryPool) ptr_vkGetInstanceProcAddr(context, "vkResetQueryPool");
    ptr_vkSetDebugUtilsObjectNameEXT = (PFN_vkSetDebugUtilsObjectNameEXT) ptr_vkGetInstanceProcAddr(context, "vkSetDebugUtilsObjectNameEXT");
    ptr_vkSetDebugUtilsObjectTagEXT = (PFN_vkSetDebugUtilsObjectTagEXT) ptr_vkGetInstanceProcAddr(context, "vkSetDebugUtilsObjectTagEXT");
    ptr_vkSetDeviceMemoryPriorityEXT = (PFN_vkSetDeviceMemoryPriorityEXT) ptr_vkGetInstanceProcAddr(context, "vkSetDeviceMemoryPriorityEXT");
    ptr_vkSetEvent = (PFN_vkSetEvent) ptr_vkGetInstanceProcAddr(context, "vkSetEvent");
    ptr_vkSetHdrMetadataEXT = (PFN_vkSetHdrMetadataEXT) ptr_vkGetInstanceProcAddr(context, "vkSetHdrMetadataEXT");
    ptr_vkSetLatencyMarkerNV = (PFN_vkSetLatencyMarkerNV) ptr_vkGetInstanceProcAddr(context, "vkSetLatencyMarkerNV");
    ptr_vkSetLatencySleepModeNV = (PFN_vkSetLatencySleepModeNV) ptr_vkGetInstanceProcAddr(context, "vkSetLatencySleepModeNV");
    ptr_vkSetLocalDimmingAMD = (PFN_vkSetLocalDimmingAMD) ptr_vkGetInstanceProcAddr(context, "vkSetLocalDimmingAMD");
    ptr_vkSetPrivateData = (PFN_vkSetPrivateData) ptr_vkGetInstanceProcAddr(context, "vkSetPrivateData");
    ptr_vkSignalSemaphore = (PFN_vkSignalSemaphore) ptr_vkGetInstanceProcAddr(context, "vkSignalSemaphore");
    ptr_vkSubmitDebugUtilsMessageEXT = (PFN_vkSubmitDebugUtilsMessageEXT) ptr_vkGetInstanceProcAddr(context, "vkSubmitDebugUtilsMessageEXT");
    ptr_vkTransitionImageLayout = (PFN_vkTransitionImageLayout) ptr_vkGetInstanceProcAddr(context, "vkTransitionImageLayout");
    ptr_vkTrimCommandPool = (PFN_vkTrimCommandPool) ptr_vkGetInstanceProcAddr(context, "vkTrimCommandPool");
    ptr_vkUninitializePerformanceApiINTEL = (PFN_vkUninitializePerformanceApiINTEL) ptr_vkGetInstanceProcAddr(context, "vkUninitializePerformanceApiINTEL");
    ptr_vkUnmapMemory = (PFN_vkUnmapMemory) ptr_vkGetInstanceProcAddr(context, "vkUnmapMemory");
    ptr_vkUnmapMemory2 = (PFN_vkUnmapMemory2) ptr_vkGetInstanceProcAddr(context, "vkUnmapMemory2");
    ptr_vkUpdateDescriptorSetWithTemplate = (PFN_vkUpdateDescriptorSetWithTemplate) ptr_vkGetInstanceProcAddr(context, "vkUpdateDescriptorSetWithTemplate");
    ptr_vkUpdateDescriptorSets = (PFN_vkUpdateDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkUpdateDescriptorSets");
    ptr_vkUpdateIndirectExecutionSetPipelineEXT = (PFN_vkUpdateIndirectExecutionSetPipelineEXT) ptr_vkGetInstanceProcAddr(context, "vkUpdateIndirectExecutionSetPipelineEXT");
    ptr_vkUpdateIndirectExecutionSetShaderEXT = (PFN_vkUpdateIndirectExecutionSetShaderEXT) ptr_vkGetInstanceProcAddr(context, "vkUpdateIndirectExecutionSetShaderEXT");
    ptr_vkUpdateVideoSessionParametersKHR = (PFN_vkUpdateVideoSessionParametersKHR) ptr_vkGetInstanceProcAddr(context, "vkUpdateVideoSessionParametersKHR");
    ptr_vkWaitForFences = (PFN_vkWaitForFences) ptr_vkGetInstanceProcAddr(context, "vkWaitForFences");
    ptr_vkWaitForPresent2KHR = (PFN_vkWaitForPresent2KHR) ptr_vkGetInstanceProcAddr(context, "vkWaitForPresent2KHR");
    ptr_vkWaitForPresentKHR = (PFN_vkWaitForPresentKHR) ptr_vkGetInstanceProcAddr(context, "vkWaitForPresentKHR");
    ptr_vkWaitSemaphores = (PFN_vkWaitSemaphores) ptr_vkGetInstanceProcAddr(context, "vkWaitSemaphores");
    ptr_vkWriteAccelerationStructuresPropertiesKHR = (PFN_vkWriteAccelerationStructuresPropertiesKHR) ptr_vkGetInstanceProcAddr(context, "vkWriteAccelerationStructuresPropertiesKHR");
    ptr_vkWriteMicromapsPropertiesEXT = (PFN_vkWriteMicromapsPropertiesEXT) ptr_vkGetInstanceProcAddr(context, "vkWriteMicromapsPropertiesEXT");
}
*/
import "C"

// Load attempts to load all global vulkan functions.
func Load(loader unsafe.Pointer) {
	C.gfx_vkInit(loader)
}

// LoadInstance attempts to load all instance vulkan functions.
func LoadInstance(instance vk.Instance) {
	C.gfx_vkInitInstance(C.VkInstance(unsafe.Pointer(instance)))
}

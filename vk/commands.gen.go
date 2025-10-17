package vk

import (
	"unsafe"

	ffi "github.com/csnewman/go-gfx/ffi"
)

// #include "vulkan.h"
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
PFN_vkBeginCommandBuffer ptr_vkBeginCommandBuffer;
VKAPI_ATTR VkResult VKAPI_CALL vkBeginCommandBuffer(VkCommandBuffer commandBuffer, VkCommandBufferBeginInfo* pBeginInfo) {
	return ptr_vkBeginCommandBuffer(commandBuffer, pBeginInfo);
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
PFN_vkCmdBeginQuery ptr_vkCmdBeginQuery;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags) {
	return ptr_vkCmdBeginQuery(commandBuffer, queryPool, query, flags);
}
*/
/*
PFN_vkCmdBeginRenderPass ptr_vkCmdBeginRenderPass;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginRenderPass(VkCommandBuffer commandBuffer, VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
	return ptr_vkCmdBeginRenderPass(commandBuffer, pRenderPassBegin, contents);
}
*/
/*
PFN_vkCmdBeginRenderPass2 ptr_vkCmdBeginRenderPass2;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginRenderPass2(VkCommandBuffer commandBuffer, VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassBeginInfo* pSubpassBeginInfo) {
	return ptr_vkCmdBeginRenderPass2(commandBuffer, pRenderPassBegin, pSubpassBeginInfo);
}
*/
/*
PFN_vkCmdBeginRendering ptr_vkCmdBeginRendering;
VKAPI_ATTR void VKAPI_CALL vkCmdBeginRendering(VkCommandBuffer commandBuffer, VkRenderingInfo* pRenderingInfo) {
	return ptr_vkCmdBeginRendering(commandBuffer, pRenderingInfo);
}
*/
/*
PFN_vkCmdBindDescriptorSets ptr_vkCmdBindDescriptorSets;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorSets(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t descriptorSetCount, VkDescriptorSet* pDescriptorSets, uint32_t dynamicOffsetCount, uint32_t* pDynamicOffsets) {
	return ptr_vkCmdBindDescriptorSets(commandBuffer, pipelineBindPoint, layout, firstSet, descriptorSetCount, pDescriptorSets, dynamicOffsetCount, pDynamicOffsets);
}
*/
/*
PFN_vkCmdBindDescriptorSets2 ptr_vkCmdBindDescriptorSets2;
VKAPI_ATTR void VKAPI_CALL vkCmdBindDescriptorSets2(VkCommandBuffer commandBuffer, VkBindDescriptorSetsInfo* pBindDescriptorSetsInfo) {
	return ptr_vkCmdBindDescriptorSets2(commandBuffer, pBindDescriptorSetsInfo);
}
*/
/*
PFN_vkCmdBindIndexBuffer ptr_vkCmdBindIndexBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdBindIndexBuffer(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType) {
	return ptr_vkCmdBindIndexBuffer(commandBuffer, buffer, offset, indexType);
}
*/
/*
PFN_vkCmdBindIndexBuffer2 ptr_vkCmdBindIndexBuffer2;
VKAPI_ATTR void VKAPI_CALL vkCmdBindIndexBuffer2(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkDeviceSize size, VkIndexType indexType) {
	return ptr_vkCmdBindIndexBuffer2(commandBuffer, buffer, offset, size, indexType);
}
*/
/*
PFN_vkCmdBindPipeline ptr_vkCmdBindPipeline;
VKAPI_ATTR void VKAPI_CALL vkCmdBindPipeline(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
	return ptr_vkCmdBindPipeline(commandBuffer, pipelineBindPoint, pipeline);
}
*/
/*
PFN_vkCmdBindVertexBuffers ptr_vkCmdBindVertexBuffers;
VKAPI_ATTR void VKAPI_CALL vkCmdBindVertexBuffers(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, VkBuffer* pBuffers, VkDeviceSize* pOffsets) {
	return ptr_vkCmdBindVertexBuffers(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets);
}
*/
/*
PFN_vkCmdBindVertexBuffers2 ptr_vkCmdBindVertexBuffers2;
VKAPI_ATTR void VKAPI_CALL vkCmdBindVertexBuffers2(VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, VkBuffer* pBuffers, VkDeviceSize* pOffsets, VkDeviceSize* pSizes, VkDeviceSize* pStrides) {
	return ptr_vkCmdBindVertexBuffers2(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets, pSizes, pStrides);
}
*/
/*
PFN_vkCmdBlitImage ptr_vkCmdBlitImage;
VKAPI_ATTR void VKAPI_CALL vkCmdBlitImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageBlit* pRegions, VkFilter filter) {
	return ptr_vkCmdBlitImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions, filter);
}
*/
/*
PFN_vkCmdBlitImage2 ptr_vkCmdBlitImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdBlitImage2(VkCommandBuffer commandBuffer, VkBlitImageInfo2* pBlitImageInfo) {
	return ptr_vkCmdBlitImage2(commandBuffer, pBlitImageInfo);
}
*/
/*
PFN_vkCmdClearAttachments ptr_vkCmdClearAttachments;
VKAPI_ATTR void VKAPI_CALL vkCmdClearAttachments(VkCommandBuffer commandBuffer, uint32_t attachmentCount, VkClearAttachment* pAttachments, uint32_t rectCount, VkClearRect* pRects) {
	return ptr_vkCmdClearAttachments(commandBuffer, attachmentCount, pAttachments, rectCount, pRects);
}
*/
/*
PFN_vkCmdClearDepthStencilImage ptr_vkCmdClearDepthStencilImage;
VKAPI_ATTR void VKAPI_CALL vkCmdClearDepthStencilImage(VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, VkImageSubresourceRange* pRanges) {
	return ptr_vkCmdClearDepthStencilImage(commandBuffer, image, imageLayout, pDepthStencil, rangeCount, pRanges);
}
*/
/*
PFN_vkCmdCopyBuffer ptr_vkCmdCopyBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBuffer(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkBuffer dstBuffer, uint32_t regionCount, VkBufferCopy* pRegions) {
	return ptr_vkCmdCopyBuffer(commandBuffer, srcBuffer, dstBuffer, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyBuffer2 ptr_vkCmdCopyBuffer2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBuffer2(VkCommandBuffer commandBuffer, VkCopyBufferInfo2* pCopyBufferInfo) {
	return ptr_vkCmdCopyBuffer2(commandBuffer, pCopyBufferInfo);
}
*/
/*
PFN_vkCmdCopyBufferToImage ptr_vkCmdCopyBufferToImage;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBufferToImage(VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkBufferImageCopy* pRegions) {
	return ptr_vkCmdCopyBufferToImage(commandBuffer, srcBuffer, dstImage, dstImageLayout, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyBufferToImage2 ptr_vkCmdCopyBufferToImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyBufferToImage2(VkCommandBuffer commandBuffer, VkCopyBufferToImageInfo2* pCopyBufferToImageInfo) {
	return ptr_vkCmdCopyBufferToImage2(commandBuffer, pCopyBufferToImageInfo);
}
*/
/*
PFN_vkCmdCopyImage ptr_vkCmdCopyImage;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageCopy* pRegions) {
	return ptr_vkCmdCopyImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyImage2 ptr_vkCmdCopyImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImage2(VkCommandBuffer commandBuffer, VkCopyImageInfo2* pCopyImageInfo) {
	return ptr_vkCmdCopyImage2(commandBuffer, pCopyImageInfo);
}
*/
/*
PFN_vkCmdCopyImageToBuffer ptr_vkCmdCopyImageToBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImageToBuffer(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkBuffer dstBuffer, uint32_t regionCount, VkBufferImageCopy* pRegions) {
	return ptr_vkCmdCopyImageToBuffer(commandBuffer, srcImage, srcImageLayout, dstBuffer, regionCount, pRegions);
}
*/
/*
PFN_vkCmdCopyImageToBuffer2 ptr_vkCmdCopyImageToBuffer2;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyImageToBuffer2(VkCommandBuffer commandBuffer, VkCopyImageToBufferInfo2* pCopyImageToBufferInfo) {
	return ptr_vkCmdCopyImageToBuffer2(commandBuffer, pCopyImageToBufferInfo);
}
*/
/*
PFN_vkCmdCopyQueryPoolResults ptr_vkCmdCopyQueryPoolResults;
VKAPI_ATTR void VKAPI_CALL vkCmdCopyQueryPoolResults(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize stride, VkQueryResultFlags flags) {
	return ptr_vkCmdCopyQueryPoolResults(commandBuffer, queryPool, firstQuery, queryCount, dstBuffer, dstOffset, stride, flags);
}
*/
/*
PFN_vkCmdDispatch ptr_vkCmdDispatch;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatch(VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	return ptr_vkCmdDispatch(commandBuffer, groupCountX, groupCountY, groupCountZ);
}
*/
/*
PFN_vkCmdDispatchBase ptr_vkCmdDispatchBase;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatchBase(VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	return ptr_vkCmdDispatchBase(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}
*/
/*
PFN_vkCmdDispatchIndirect ptr_vkCmdDispatchIndirect;
VKAPI_ATTR void VKAPI_CALL vkCmdDispatchIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset) {
	return ptr_vkCmdDispatchIndirect(commandBuffer, buffer, offset);
}
*/
/*
PFN_vkCmdDraw ptr_vkCmdDraw;
VKAPI_ATTR void VKAPI_CALL vkCmdDraw(VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance) {
	return ptr_vkCmdDraw(commandBuffer, vertexCount, instanceCount, firstVertex, firstInstance);
}
*/
/*
PFN_vkCmdDrawIndexed ptr_vkCmdDrawIndexed;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndexed(VkCommandBuffer commandBuffer, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t vertexOffset, uint32_t firstInstance) {
	return ptr_vkCmdDrawIndexed(commandBuffer, indexCount, instanceCount, firstIndex, vertexOffset, firstInstance);
}
*/
/*
PFN_vkCmdDrawIndexedIndirect ptr_vkCmdDrawIndexedIndirect;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndexedIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	return ptr_vkCmdDrawIndexedIndirect(commandBuffer, buffer, offset, drawCount, stride);
}
*/
/*
PFN_vkCmdDrawIndexedIndirectCount ptr_vkCmdDrawIndexedIndirectCount;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndexedIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	return ptr_vkCmdDrawIndexedIndirectCount(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
/*
PFN_vkCmdDrawIndirect ptr_vkCmdDrawIndirect;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndirect(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
	return ptr_vkCmdDrawIndirect(commandBuffer, buffer, offset, drawCount, stride);
}
*/
/*
PFN_vkCmdDrawIndirectCount ptr_vkCmdDrawIndirectCount;
VKAPI_ATTR void VKAPI_CALL vkCmdDrawIndirectCount(VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	return ptr_vkCmdDrawIndirectCount(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
/*
PFN_vkCmdEndQuery ptr_vkCmdEndQuery;
VKAPI_ATTR void VKAPI_CALL vkCmdEndQuery(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query) {
	return ptr_vkCmdEndQuery(commandBuffer, queryPool, query);
}
*/
/*
PFN_vkCmdEndRenderPass ptr_vkCmdEndRenderPass;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRenderPass(VkCommandBuffer commandBuffer) {
	return ptr_vkCmdEndRenderPass(commandBuffer);
}
*/
/*
PFN_vkCmdEndRenderPass2 ptr_vkCmdEndRenderPass2;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRenderPass2(VkCommandBuffer commandBuffer, VkSubpassEndInfo* pSubpassEndInfo) {
	return ptr_vkCmdEndRenderPass2(commandBuffer, pSubpassEndInfo);
}
*/
/*
PFN_vkCmdEndRendering ptr_vkCmdEndRendering;
VKAPI_ATTR void VKAPI_CALL vkCmdEndRendering(VkCommandBuffer commandBuffer) {
	return ptr_vkCmdEndRendering(commandBuffer);
}
*/
/*
PFN_vkCmdExecuteCommands ptr_vkCmdExecuteCommands;
VKAPI_ATTR void VKAPI_CALL vkCmdExecuteCommands(VkCommandBuffer commandBuffer, uint32_t commandBufferCount, VkCommandBuffer* pCommandBuffers) {
	return ptr_vkCmdExecuteCommands(commandBuffer, commandBufferCount, pCommandBuffers);
}
*/
/*
PFN_vkCmdFillBuffer ptr_vkCmdFillBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdFillBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data) {
	return ptr_vkCmdFillBuffer(commandBuffer, dstBuffer, dstOffset, size, data);
}
*/
/*
PFN_vkCmdNextSubpass ptr_vkCmdNextSubpass;
VKAPI_ATTR void VKAPI_CALL vkCmdNextSubpass(VkCommandBuffer commandBuffer, VkSubpassContents contents) {
	return ptr_vkCmdNextSubpass(commandBuffer, contents);
}
*/
/*
PFN_vkCmdNextSubpass2 ptr_vkCmdNextSubpass2;
VKAPI_ATTR void VKAPI_CALL vkCmdNextSubpass2(VkCommandBuffer commandBuffer, VkSubpassBeginInfo* pSubpassBeginInfo, VkSubpassEndInfo* pSubpassEndInfo) {
	return ptr_vkCmdNextSubpass2(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}
*/
/*
PFN_vkCmdPipelineBarrier ptr_vkCmdPipelineBarrier;
VKAPI_ATTR void VKAPI_CALL vkCmdPipelineBarrier(VkCommandBuffer commandBuffer, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, VkDependencyFlags dependencyFlags, uint32_t memoryBarrierCount, VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, VkImageMemoryBarrier* pImageMemoryBarriers) {
	return ptr_vkCmdPipelineBarrier(commandBuffer, srcStageMask, dstStageMask, dependencyFlags, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}
*/
/*
PFN_vkCmdPipelineBarrier2 ptr_vkCmdPipelineBarrier2;
VKAPI_ATTR void VKAPI_CALL vkCmdPipelineBarrier2(VkCommandBuffer commandBuffer, VkDependencyInfo* pDependencyInfo) {
	return ptr_vkCmdPipelineBarrier2(commandBuffer, pDependencyInfo);
}
*/
/*
PFN_vkCmdPushConstants ptr_vkCmdPushConstants;
VKAPI_ATTR void VKAPI_CALL vkCmdPushConstants(VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, void* pValues) {
	return ptr_vkCmdPushConstants(commandBuffer, layout, stageFlags, offset, size, pValues);
}
*/
/*
PFN_vkCmdPushConstants2 ptr_vkCmdPushConstants2;
VKAPI_ATTR void VKAPI_CALL vkCmdPushConstants2(VkCommandBuffer commandBuffer, VkPushConstantsInfo* pPushConstantsInfo) {
	return ptr_vkCmdPushConstants2(commandBuffer, pPushConstantsInfo);
}
*/
/*
PFN_vkCmdPushDescriptorSet ptr_vkCmdPushDescriptorSet;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSet(VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t set, uint32_t descriptorWriteCount, VkWriteDescriptorSet* pDescriptorWrites) {
	return ptr_vkCmdPushDescriptorSet(commandBuffer, pipelineBindPoint, layout, set, descriptorWriteCount, pDescriptorWrites);
}
*/
/*
PFN_vkCmdPushDescriptorSet2 ptr_vkCmdPushDescriptorSet2;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSet2(VkCommandBuffer commandBuffer, VkPushDescriptorSetInfo* pPushDescriptorSetInfo) {
	return ptr_vkCmdPushDescriptorSet2(commandBuffer, pPushDescriptorSetInfo);
}
*/
/*
PFN_vkCmdPushDescriptorSetWithTemplate ptr_vkCmdPushDescriptorSetWithTemplate;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSetWithTemplate(VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, void* pData) {
	return ptr_vkCmdPushDescriptorSetWithTemplate(commandBuffer, descriptorUpdateTemplate, layout, set, pData);
}
*/
/*
PFN_vkCmdPushDescriptorSetWithTemplate2 ptr_vkCmdPushDescriptorSetWithTemplate2;
VKAPI_ATTR void VKAPI_CALL vkCmdPushDescriptorSetWithTemplate2(VkCommandBuffer commandBuffer, VkPushDescriptorSetWithTemplateInfo* pPushDescriptorSetWithTemplateInfo) {
	return ptr_vkCmdPushDescriptorSetWithTemplate2(commandBuffer, pPushDescriptorSetWithTemplateInfo);
}
*/
/*
PFN_vkCmdResetEvent ptr_vkCmdResetEvent;
VKAPI_ATTR void VKAPI_CALL vkCmdResetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	return ptr_vkCmdResetEvent(commandBuffer, event, stageMask);
}
*/
/*
PFN_vkCmdResetEvent2 ptr_vkCmdResetEvent2;
VKAPI_ATTR void VKAPI_CALL vkCmdResetEvent2(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags2 stageMask) {
	return ptr_vkCmdResetEvent2(commandBuffer, event, stageMask);
}
*/
/*
PFN_vkCmdResetQueryPool ptr_vkCmdResetQueryPool;
VKAPI_ATTR void VKAPI_CALL vkCmdResetQueryPool(VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	return ptr_vkCmdResetQueryPool(commandBuffer, queryPool, firstQuery, queryCount);
}
*/
/*
PFN_vkCmdResolveImage ptr_vkCmdResolveImage;
VKAPI_ATTR void VKAPI_CALL vkCmdResolveImage(VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageResolve* pRegions) {
	return ptr_vkCmdResolveImage(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}
*/
/*
PFN_vkCmdResolveImage2 ptr_vkCmdResolveImage2;
VKAPI_ATTR void VKAPI_CALL vkCmdResolveImage2(VkCommandBuffer commandBuffer, VkResolveImageInfo2* pResolveImageInfo) {
	return ptr_vkCmdResolveImage2(commandBuffer, pResolveImageInfo);
}
*/
/*
PFN_vkCmdSetCullMode ptr_vkCmdSetCullMode;
VKAPI_ATTR void VKAPI_CALL vkCmdSetCullMode(VkCommandBuffer commandBuffer, VkCullModeFlags cullMode) {
	return ptr_vkCmdSetCullMode(commandBuffer, cullMode);
}
*/
/*
PFN_vkCmdSetDepthBias ptr_vkCmdSetDepthBias;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBias(VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor) {
	return ptr_vkCmdSetDepthBias(commandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor);
}
*/
/*
PFN_vkCmdSetDepthBiasEnable ptr_vkCmdSetDepthBiasEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBiasEnable(VkCommandBuffer commandBuffer, VkBool32 depthBiasEnable) {
	return ptr_vkCmdSetDepthBiasEnable(commandBuffer, depthBiasEnable);
}
*/
/*
PFN_vkCmdSetDepthBounds ptr_vkCmdSetDepthBounds;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBounds(VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds) {
	return ptr_vkCmdSetDepthBounds(commandBuffer, minDepthBounds, maxDepthBounds);
}
*/
/*
PFN_vkCmdSetDepthBoundsTestEnable ptr_vkCmdSetDepthBoundsTestEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthBoundsTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthBoundsTestEnable) {
	return ptr_vkCmdSetDepthBoundsTestEnable(commandBuffer, depthBoundsTestEnable);
}
*/
/*
PFN_vkCmdSetDepthCompareOp ptr_vkCmdSetDepthCompareOp;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthCompareOp(VkCommandBuffer commandBuffer, VkCompareOp depthCompareOp) {
	return ptr_vkCmdSetDepthCompareOp(commandBuffer, depthCompareOp);
}
*/
/*
PFN_vkCmdSetDepthTestEnable ptr_vkCmdSetDepthTestEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthTestEnable(VkCommandBuffer commandBuffer, VkBool32 depthTestEnable) {
	return ptr_vkCmdSetDepthTestEnable(commandBuffer, depthTestEnable);
}
*/
/*
PFN_vkCmdSetDepthWriteEnable ptr_vkCmdSetDepthWriteEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDepthWriteEnable(VkCommandBuffer commandBuffer, VkBool32 depthWriteEnable) {
	return ptr_vkCmdSetDepthWriteEnable(commandBuffer, depthWriteEnable);
}
*/
/*
PFN_vkCmdSetDeviceMask ptr_vkCmdSetDeviceMask;
VKAPI_ATTR void VKAPI_CALL vkCmdSetDeviceMask(VkCommandBuffer commandBuffer, uint32_t deviceMask) {
	return ptr_vkCmdSetDeviceMask(commandBuffer, deviceMask);
}
*/
/*
PFN_vkCmdSetEvent ptr_vkCmdSetEvent;
VKAPI_ATTR void VKAPI_CALL vkCmdSetEvent(VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
	return ptr_vkCmdSetEvent(commandBuffer, event, stageMask);
}
*/
/*
PFN_vkCmdSetEvent2 ptr_vkCmdSetEvent2;
VKAPI_ATTR void VKAPI_CALL vkCmdSetEvent2(VkCommandBuffer commandBuffer, VkEvent event, VkDependencyInfo* pDependencyInfo) {
	return ptr_vkCmdSetEvent2(commandBuffer, event, pDependencyInfo);
}
*/
/*
PFN_vkCmdSetFrontFace ptr_vkCmdSetFrontFace;
VKAPI_ATTR void VKAPI_CALL vkCmdSetFrontFace(VkCommandBuffer commandBuffer, VkFrontFace frontFace) {
	return ptr_vkCmdSetFrontFace(commandBuffer, frontFace);
}
*/
/*
PFN_vkCmdSetLineStipple ptr_vkCmdSetLineStipple;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLineStipple(VkCommandBuffer commandBuffer, uint32_t lineStippleFactor, uint16_t lineStipplePattern) {
	return ptr_vkCmdSetLineStipple(commandBuffer, lineStippleFactor, lineStipplePattern);
}
*/
/*
PFN_vkCmdSetLineWidth ptr_vkCmdSetLineWidth;
VKAPI_ATTR void VKAPI_CALL vkCmdSetLineWidth(VkCommandBuffer commandBuffer, float lineWidth) {
	return ptr_vkCmdSetLineWidth(commandBuffer, lineWidth);
}
*/
/*
PFN_vkCmdSetPrimitiveRestartEnable ptr_vkCmdSetPrimitiveRestartEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetPrimitiveRestartEnable(VkCommandBuffer commandBuffer, VkBool32 primitiveRestartEnable) {
	return ptr_vkCmdSetPrimitiveRestartEnable(commandBuffer, primitiveRestartEnable);
}
*/
/*
PFN_vkCmdSetPrimitiveTopology ptr_vkCmdSetPrimitiveTopology;
VKAPI_ATTR void VKAPI_CALL vkCmdSetPrimitiveTopology(VkCommandBuffer commandBuffer, VkPrimitiveTopology primitiveTopology) {
	return ptr_vkCmdSetPrimitiveTopology(commandBuffer, primitiveTopology);
}
*/
/*
PFN_vkCmdSetRasterizerDiscardEnable ptr_vkCmdSetRasterizerDiscardEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRasterizerDiscardEnable(VkCommandBuffer commandBuffer, VkBool32 rasterizerDiscardEnable) {
	return ptr_vkCmdSetRasterizerDiscardEnable(commandBuffer, rasterizerDiscardEnable);
}
*/
/*
PFN_vkCmdSetRenderingAttachmentLocations ptr_vkCmdSetRenderingAttachmentLocations;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRenderingAttachmentLocations(VkCommandBuffer commandBuffer, VkRenderingAttachmentLocationInfo* pLocationInfo) {
	return ptr_vkCmdSetRenderingAttachmentLocations(commandBuffer, pLocationInfo);
}
*/
/*
PFN_vkCmdSetRenderingInputAttachmentIndices ptr_vkCmdSetRenderingInputAttachmentIndices;
VKAPI_ATTR void VKAPI_CALL vkCmdSetRenderingInputAttachmentIndices(VkCommandBuffer commandBuffer, VkRenderingInputAttachmentIndexInfo* pInputAttachmentIndexInfo) {
	return ptr_vkCmdSetRenderingInputAttachmentIndices(commandBuffer, pInputAttachmentIndexInfo);
}
*/
/*
PFN_vkCmdSetScissor ptr_vkCmdSetScissor;
VKAPI_ATTR void VKAPI_CALL vkCmdSetScissor(VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, VkRect2D* pScissors) {
	return ptr_vkCmdSetScissor(commandBuffer, firstScissor, scissorCount, pScissors);
}
*/
/*
PFN_vkCmdSetScissorWithCount ptr_vkCmdSetScissorWithCount;
VKAPI_ATTR void VKAPI_CALL vkCmdSetScissorWithCount(VkCommandBuffer commandBuffer, uint32_t scissorCount, VkRect2D* pScissors) {
	return ptr_vkCmdSetScissorWithCount(commandBuffer, scissorCount, pScissors);
}
*/
/*
PFN_vkCmdSetStencilCompareMask ptr_vkCmdSetStencilCompareMask;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilCompareMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t compareMask) {
	return ptr_vkCmdSetStencilCompareMask(commandBuffer, faceMask, compareMask);
}
*/
/*
PFN_vkCmdSetStencilOp ptr_vkCmdSetStencilOp;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilOp(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, VkStencilOp failOp, VkStencilOp passOp, VkStencilOp depthFailOp, VkCompareOp compareOp) {
	return ptr_vkCmdSetStencilOp(commandBuffer, faceMask, failOp, passOp, depthFailOp, compareOp);
}
*/
/*
PFN_vkCmdSetStencilReference ptr_vkCmdSetStencilReference;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilReference(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t reference) {
	return ptr_vkCmdSetStencilReference(commandBuffer, faceMask, reference);
}
*/
/*
PFN_vkCmdSetStencilTestEnable ptr_vkCmdSetStencilTestEnable;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilTestEnable(VkCommandBuffer commandBuffer, VkBool32 stencilTestEnable) {
	return ptr_vkCmdSetStencilTestEnable(commandBuffer, stencilTestEnable);
}
*/
/*
PFN_vkCmdSetStencilWriteMask ptr_vkCmdSetStencilWriteMask;
VKAPI_ATTR void VKAPI_CALL vkCmdSetStencilWriteMask(VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t writeMask) {
	return ptr_vkCmdSetStencilWriteMask(commandBuffer, faceMask, writeMask);
}
*/
/*
PFN_vkCmdSetViewport ptr_vkCmdSetViewport;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewport(VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, VkViewport* pViewports) {
	return ptr_vkCmdSetViewport(commandBuffer, firstViewport, viewportCount, pViewports);
}
*/
/*
PFN_vkCmdSetViewportWithCount ptr_vkCmdSetViewportWithCount;
VKAPI_ATTR void VKAPI_CALL vkCmdSetViewportWithCount(VkCommandBuffer commandBuffer, uint32_t viewportCount, VkViewport* pViewports) {
	return ptr_vkCmdSetViewportWithCount(commandBuffer, viewportCount, pViewports);
}
*/
/*
PFN_vkCmdUpdateBuffer ptr_vkCmdUpdateBuffer;
VKAPI_ATTR void VKAPI_CALL vkCmdUpdateBuffer(VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize dataSize, void* pData) {
	return ptr_vkCmdUpdateBuffer(commandBuffer, dstBuffer, dstOffset, dataSize, pData);
}
*/
/*
PFN_vkCmdWaitEvents ptr_vkCmdWaitEvents;
VKAPI_ATTR void VKAPI_CALL vkCmdWaitEvents(VkCommandBuffer commandBuffer, uint32_t eventCount, VkEvent* pEvents, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, uint32_t memoryBarrierCount, VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, VkImageMemoryBarrier* pImageMemoryBarriers) {
	return ptr_vkCmdWaitEvents(commandBuffer, eventCount, pEvents, srcStageMask, dstStageMask, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}
*/
/*
PFN_vkCmdWaitEvents2 ptr_vkCmdWaitEvents2;
VKAPI_ATTR void VKAPI_CALL vkCmdWaitEvents2(VkCommandBuffer commandBuffer, uint32_t eventCount, VkEvent* pEvents, VkDependencyInfo* pDependencyInfos) {
	return ptr_vkCmdWaitEvents2(commandBuffer, eventCount, pEvents, pDependencyInfos);
}
*/
/*
PFN_vkCmdWriteTimestamp ptr_vkCmdWriteTimestamp;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteTimestamp(VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkQueryPool queryPool, uint32_t query) {
	return ptr_vkCmdWriteTimestamp(commandBuffer, pipelineStage, queryPool, query);
}
*/
/*
PFN_vkCmdWriteTimestamp2 ptr_vkCmdWriteTimestamp2;
VKAPI_ATTR void VKAPI_CALL vkCmdWriteTimestamp2(VkCommandBuffer commandBuffer, VkPipelineStageFlags2 stage, VkQueryPool queryPool, uint32_t query) {
	return ptr_vkCmdWriteTimestamp2(commandBuffer, stage, queryPool, query);
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
PFN_vkCopyMemoryToImage ptr_vkCopyMemoryToImage;
VKAPI_ATTR VkResult VKAPI_CALL vkCopyMemoryToImage(VkDevice device, VkCopyMemoryToImageInfo* pCopyMemoryToImageInfo) {
	return ptr_vkCopyMemoryToImage(device, pCopyMemoryToImageInfo);
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
PFN_vkCreateEvent ptr_vkCreateEvent;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateEvent(VkDevice device, VkEventCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkEvent* pEvent) {
	return ptr_vkCreateEvent(device, pCreateInfo, pAllocator, pEvent);
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
PFN_vkCreateInstance ptr_vkCreateInstance;
VKAPI_ATTR VkResult VKAPI_CALL vkCreateInstance(VkInstanceCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkInstance* pInstance) {
	return ptr_vkCreateInstance(pCreateInfo, pAllocator, pInstance);
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
PFN_vkDestroyBuffer ptr_vkDestroyBuffer;
VKAPI_ATTR void VKAPI_CALL vkDestroyBuffer(VkDevice device, VkBuffer buffer, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyBuffer(device, buffer, pAllocator);
}
*/
/*
PFN_vkDestroyBufferView ptr_vkDestroyBufferView;
VKAPI_ATTR void VKAPI_CALL vkDestroyBufferView(VkDevice device, VkBufferView bufferView, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyBufferView(device, bufferView, pAllocator);
}
*/
/*
PFN_vkDestroyCommandPool ptr_vkDestroyCommandPool;
VKAPI_ATTR void VKAPI_CALL vkDestroyCommandPool(VkDevice device, VkCommandPool commandPool, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyCommandPool(device, commandPool, pAllocator);
}
*/
/*
PFN_vkDestroyDescriptorPool ptr_vkDestroyDescriptorPool;
VKAPI_ATTR void VKAPI_CALL vkDestroyDescriptorPool(VkDevice device, VkDescriptorPool descriptorPool, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyDescriptorPool(device, descriptorPool, pAllocator);
}
*/
/*
PFN_vkDestroyDescriptorSetLayout ptr_vkDestroyDescriptorSetLayout;
VKAPI_ATTR void VKAPI_CALL vkDestroyDescriptorSetLayout(VkDevice device, VkDescriptorSetLayout descriptorSetLayout, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyDescriptorSetLayout(device, descriptorSetLayout, pAllocator);
}
*/
/*
PFN_vkDestroyDescriptorUpdateTemplate ptr_vkDestroyDescriptorUpdateTemplate;
VKAPI_ATTR void VKAPI_CALL vkDestroyDescriptorUpdateTemplate(VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyDescriptorUpdateTemplate(device, descriptorUpdateTemplate, pAllocator);
}
*/
/*
PFN_vkDestroyDevice ptr_vkDestroyDevice;
VKAPI_ATTR void VKAPI_CALL vkDestroyDevice(VkDevice device, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyDevice(device, pAllocator);
}
*/
/*
PFN_vkDestroyEvent ptr_vkDestroyEvent;
VKAPI_ATTR void VKAPI_CALL vkDestroyEvent(VkDevice device, VkEvent event, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyEvent(device, event, pAllocator);
}
*/
/*
PFN_vkDestroyFence ptr_vkDestroyFence;
VKAPI_ATTR void VKAPI_CALL vkDestroyFence(VkDevice device, VkFence fence, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyFence(device, fence, pAllocator);
}
*/
/*
PFN_vkDestroyFramebuffer ptr_vkDestroyFramebuffer;
VKAPI_ATTR void VKAPI_CALL vkDestroyFramebuffer(VkDevice device, VkFramebuffer framebuffer, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyFramebuffer(device, framebuffer, pAllocator);
}
*/
/*
PFN_vkDestroyImage ptr_vkDestroyImage;
VKAPI_ATTR void VKAPI_CALL vkDestroyImage(VkDevice device, VkImage image, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyImage(device, image, pAllocator);
}
*/
/*
PFN_vkDestroyImageView ptr_vkDestroyImageView;
VKAPI_ATTR void VKAPI_CALL vkDestroyImageView(VkDevice device, VkImageView imageView, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyImageView(device, imageView, pAllocator);
}
*/
/*
PFN_vkDestroyInstance ptr_vkDestroyInstance;
VKAPI_ATTR void VKAPI_CALL vkDestroyInstance(VkInstance instance, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyInstance(instance, pAllocator);
}
*/
/*
PFN_vkDestroyPipeline ptr_vkDestroyPipeline;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipeline(VkDevice device, VkPipeline pipeline, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyPipeline(device, pipeline, pAllocator);
}
*/
/*
PFN_vkDestroyPipelineCache ptr_vkDestroyPipelineCache;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipelineCache(VkDevice device, VkPipelineCache pipelineCache, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyPipelineCache(device, pipelineCache, pAllocator);
}
*/
/*
PFN_vkDestroyPipelineLayout ptr_vkDestroyPipelineLayout;
VKAPI_ATTR void VKAPI_CALL vkDestroyPipelineLayout(VkDevice device, VkPipelineLayout pipelineLayout, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyPipelineLayout(device, pipelineLayout, pAllocator);
}
*/
/*
PFN_vkDestroyPrivateDataSlot ptr_vkDestroyPrivateDataSlot;
VKAPI_ATTR void VKAPI_CALL vkDestroyPrivateDataSlot(VkDevice device, VkPrivateDataSlot privateDataSlot, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyPrivateDataSlot(device, privateDataSlot, pAllocator);
}
*/
/*
PFN_vkDestroyQueryPool ptr_vkDestroyQueryPool;
VKAPI_ATTR void VKAPI_CALL vkDestroyQueryPool(VkDevice device, VkQueryPool queryPool, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyQueryPool(device, queryPool, pAllocator);
}
*/
/*
PFN_vkDestroyRenderPass ptr_vkDestroyRenderPass;
VKAPI_ATTR void VKAPI_CALL vkDestroyRenderPass(VkDevice device, VkRenderPass renderPass, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyRenderPass(device, renderPass, pAllocator);
}
*/
/*
PFN_vkDestroySampler ptr_vkDestroySampler;
VKAPI_ATTR void VKAPI_CALL vkDestroySampler(VkDevice device, VkSampler sampler, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroySampler(device, sampler, pAllocator);
}
*/
/*
PFN_vkDestroySamplerYcbcrConversion ptr_vkDestroySamplerYcbcrConversion;
VKAPI_ATTR void VKAPI_CALL vkDestroySamplerYcbcrConversion(VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroySamplerYcbcrConversion(device, ycbcrConversion, pAllocator);
}
*/
/*
PFN_vkDestroySemaphore ptr_vkDestroySemaphore;
VKAPI_ATTR void VKAPI_CALL vkDestroySemaphore(VkDevice device, VkSemaphore semaphore, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroySemaphore(device, semaphore, pAllocator);
}
*/
/*
PFN_vkDestroyShaderModule ptr_vkDestroyShaderModule;
VKAPI_ATTR void VKAPI_CALL vkDestroyShaderModule(VkDevice device, VkShaderModule shaderModule, VkAllocationCallbacks* pAllocator) {
	return ptr_vkDestroyShaderModule(device, shaderModule, pAllocator);
}
*/
/*
PFN_vkDeviceWaitIdle ptr_vkDeviceWaitIdle;
VKAPI_ATTR VkResult VKAPI_CALL vkDeviceWaitIdle(VkDevice device) {
	return ptr_vkDeviceWaitIdle(device);
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
PFN_vkEnumeratePhysicalDevices ptr_vkEnumeratePhysicalDevices;
VKAPI_ATTR VkResult VKAPI_CALL vkEnumeratePhysicalDevices(VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices) {
	return ptr_vkEnumeratePhysicalDevices(instance, pPhysicalDeviceCount, pPhysicalDevices);
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
	return ptr_vkFreeCommandBuffers(device, commandPool, commandBufferCount, pCommandBuffers);
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
	return ptr_vkFreeMemory(device, memory, pAllocator);
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
	return ptr_vkGetBufferMemoryRequirements(device, buffer, pMemoryRequirements);
}
*/
/*
PFN_vkGetBufferMemoryRequirements2 ptr_vkGetBufferMemoryRequirements2;
VKAPI_ATTR void VKAPI_CALL vkGetBufferMemoryRequirements2(VkDevice device, VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	return ptr_vkGetBufferMemoryRequirements2(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetBufferOpaqueCaptureAddress ptr_vkGetBufferOpaqueCaptureAddress;
VKAPI_ATTR uint64_t VKAPI_CALL vkGetBufferOpaqueCaptureAddress(VkDevice device, VkBufferDeviceAddressInfo* pInfo) {
	return ptr_vkGetBufferOpaqueCaptureAddress(device, pInfo);
}
*/
/*
PFN_vkGetDescriptorSetLayoutSupport ptr_vkGetDescriptorSetLayoutSupport;
VKAPI_ATTR void VKAPI_CALL vkGetDescriptorSetLayoutSupport(VkDevice device, VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport) {
	return ptr_vkGetDescriptorSetLayoutSupport(device, pCreateInfo, pSupport);
}
*/
/*
PFN_vkGetDeviceBufferMemoryRequirements ptr_vkGetDeviceBufferMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceBufferMemoryRequirements(VkDevice device, VkDeviceBufferMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	return ptr_vkGetDeviceBufferMemoryRequirements(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetDeviceGroupPeerMemoryFeatures ptr_vkGetDeviceGroupPeerMemoryFeatures;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceGroupPeerMemoryFeatures(VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures) {
	return ptr_vkGetDeviceGroupPeerMemoryFeatures(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}
*/
/*
PFN_vkGetDeviceImageMemoryRequirements ptr_vkGetDeviceImageMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceImageMemoryRequirements(VkDevice device, VkDeviceImageMemoryRequirements* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	return ptr_vkGetDeviceImageMemoryRequirements(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetDeviceImageSparseMemoryRequirements ptr_vkGetDeviceImageSparseMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceImageSparseMemoryRequirements(VkDevice device, VkDeviceImageMemoryRequirements* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	return ptr_vkGetDeviceImageSparseMemoryRequirements(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
*/
/*
PFN_vkGetDeviceImageSubresourceLayout ptr_vkGetDeviceImageSubresourceLayout;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceImageSubresourceLayout(VkDevice device, VkDeviceImageSubresourceInfo* pInfo, VkSubresourceLayout2* pLayout) {
	return ptr_vkGetDeviceImageSubresourceLayout(device, pInfo, pLayout);
}
*/
/*
PFN_vkGetDeviceMemoryCommitment ptr_vkGetDeviceMemoryCommitment;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceMemoryCommitment(VkDevice device, VkDeviceMemory memory, VkDeviceSize* pCommittedMemoryInBytes) {
	return ptr_vkGetDeviceMemoryCommitment(device, memory, pCommittedMemoryInBytes);
}
*/
/*
PFN_vkGetDeviceMemoryOpaqueCaptureAddress ptr_vkGetDeviceMemoryOpaqueCaptureAddress;
VKAPI_ATTR uint64_t VKAPI_CALL vkGetDeviceMemoryOpaqueCaptureAddress(VkDevice device, VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo) {
	return ptr_vkGetDeviceMemoryOpaqueCaptureAddress(device, pInfo);
}
*/
/*
PFN_vkGetDeviceQueue ptr_vkGetDeviceQueue;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceQueue(VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue) {
	return ptr_vkGetDeviceQueue(device, queueFamilyIndex, queueIndex, pQueue);
}
*/
/*
PFN_vkGetDeviceQueue2 ptr_vkGetDeviceQueue2;
VKAPI_ATTR void VKAPI_CALL vkGetDeviceQueue2(VkDevice device, VkDeviceQueueInfo2* pQueueInfo, VkQueue* pQueue) {
	return ptr_vkGetDeviceQueue2(device, pQueueInfo, pQueue);
}
*/
/*
PFN_vkGetEventStatus ptr_vkGetEventStatus;
VKAPI_ATTR VkResult VKAPI_CALL vkGetEventStatus(VkDevice device, VkEvent event) {
	return ptr_vkGetEventStatus(device, event);
}
*/
/*
PFN_vkGetFenceStatus ptr_vkGetFenceStatus;
VKAPI_ATTR VkResult VKAPI_CALL vkGetFenceStatus(VkDevice device, VkFence fence) {
	return ptr_vkGetFenceStatus(device, fence);
}
*/
/*
PFN_vkGetImageMemoryRequirements ptr_vkGetImageMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetImageMemoryRequirements(VkDevice device, VkImage image, VkMemoryRequirements* pMemoryRequirements) {
	return ptr_vkGetImageMemoryRequirements(device, image, pMemoryRequirements);
}
*/
/*
PFN_vkGetImageMemoryRequirements2 ptr_vkGetImageMemoryRequirements2;
VKAPI_ATTR void VKAPI_CALL vkGetImageMemoryRequirements2(VkDevice device, VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
	return ptr_vkGetImageMemoryRequirements2(device, pInfo, pMemoryRequirements);
}
*/
/*
PFN_vkGetImageSparseMemoryRequirements ptr_vkGetImageSparseMemoryRequirements;
VKAPI_ATTR void VKAPI_CALL vkGetImageSparseMemoryRequirements(VkDevice device, VkImage image, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements* pSparseMemoryRequirements) {
	return ptr_vkGetImageSparseMemoryRequirements(device, image, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
*/
/*
PFN_vkGetImageSparseMemoryRequirements2 ptr_vkGetImageSparseMemoryRequirements2;
VKAPI_ATTR void VKAPI_CALL vkGetImageSparseMemoryRequirements2(VkDevice device, VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
	return ptr_vkGetImageSparseMemoryRequirements2(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}
*/
/*
PFN_vkGetImageSubresourceLayout ptr_vkGetImageSubresourceLayout;
VKAPI_ATTR void VKAPI_CALL vkGetImageSubresourceLayout(VkDevice device, VkImage image, VkImageSubresource* pSubresource, VkSubresourceLayout* pLayout) {
	return ptr_vkGetImageSubresourceLayout(device, image, pSubresource, pLayout);
}
*/
/*
PFN_vkGetImageSubresourceLayout2 ptr_vkGetImageSubresourceLayout2;
VKAPI_ATTR void VKAPI_CALL vkGetImageSubresourceLayout2(VkDevice device, VkImage image, VkImageSubresource2* pSubresource, VkSubresourceLayout2* pLayout) {
	return ptr_vkGetImageSubresourceLayout2(device, image, pSubresource, pLayout);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalBufferProperties ptr_vkGetPhysicalDeviceExternalBufferProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalBufferProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties) {
	return ptr_vkGetPhysicalDeviceExternalBufferProperties(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalFenceProperties ptr_vkGetPhysicalDeviceExternalFenceProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalFenceProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties) {
	return ptr_vkGetPhysicalDeviceExternalFenceProperties(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceExternalSemaphoreProperties ptr_vkGetPhysicalDeviceExternalSemaphoreProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceExternalSemaphoreProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties) {
	return ptr_vkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceFeatures ptr_vkGetPhysicalDeviceFeatures;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFeatures(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures) {
	return ptr_vkGetPhysicalDeviceFeatures(physicalDevice, pFeatures);
}
*/
/*
PFN_vkGetPhysicalDeviceFeatures2 ptr_vkGetPhysicalDeviceFeatures2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFeatures2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures) {
	return ptr_vkGetPhysicalDeviceFeatures2(physicalDevice, pFeatures);
}
*/
/*
PFN_vkGetPhysicalDeviceFormatProperties ptr_vkGetPhysicalDeviceFormatProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties* pFormatProperties) {
	return ptr_vkGetPhysicalDeviceFormatProperties(physicalDevice, format, pFormatProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceFormatProperties2 ptr_vkGetPhysicalDeviceFormatProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceFormatProperties2(VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties) {
	return ptr_vkGetPhysicalDeviceFormatProperties2(physicalDevice, format, pFormatProperties);
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
	return ptr_vkGetPhysicalDeviceMemoryProperties(physicalDevice, pMemoryProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceMemoryProperties2 ptr_vkGetPhysicalDeviceMemoryProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceMemoryProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties) {
	return ptr_vkGetPhysicalDeviceMemoryProperties2(physicalDevice, pMemoryProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceProperties ptr_vkGetPhysicalDeviceProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceProperties(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties) {
	return ptr_vkGetPhysicalDeviceProperties(physicalDevice, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceProperties2 ptr_vkGetPhysicalDeviceProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
	return ptr_vkGetPhysicalDeviceProperties2(physicalDevice, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyProperties ptr_vkGetPhysicalDeviceQueueFamilyProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceQueueFamilyProperties(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties) {
	return ptr_vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceQueueFamilyProperties2 ptr_vkGetPhysicalDeviceQueueFamilyProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceQueueFamilyProperties2(VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties) {
	return ptr_vkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceSparseImageFormatProperties ptr_vkGetPhysicalDeviceSparseImageFormatProperties;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceSparseImageFormatProperties(VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkSampleCountFlagBits samples, VkImageUsageFlags usage, VkImageTiling tiling, uint32_t* pPropertyCount, VkSparseImageFormatProperties* pProperties) {
	return ptr_vkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice, format, type, samples, usage, tiling, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceSparseImageFormatProperties2 ptr_vkGetPhysicalDeviceSparseImageFormatProperties2;
VKAPI_ATTR void VKAPI_CALL vkGetPhysicalDeviceSparseImageFormatProperties2(VkPhysicalDevice physicalDevice, VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties) {
	return ptr_vkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}
*/
/*
PFN_vkGetPhysicalDeviceToolProperties ptr_vkGetPhysicalDeviceToolProperties;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPhysicalDeviceToolProperties(VkPhysicalDevice physicalDevice, uint32_t* pToolCount, VkPhysicalDeviceToolProperties* pToolProperties) {
	return ptr_vkGetPhysicalDeviceToolProperties(physicalDevice, pToolCount, pToolProperties);
}
*/
/*
PFN_vkGetPipelineCacheData ptr_vkGetPipelineCacheData;
VKAPI_ATTR VkResult VKAPI_CALL vkGetPipelineCacheData(VkDevice device, VkPipelineCache pipelineCache, size_t* pDataSize, void* pData) {
	return ptr_vkGetPipelineCacheData(device, pipelineCache, pDataSize, pData);
}
*/
/*
PFN_vkGetPrivateData ptr_vkGetPrivateData;
VKAPI_ATTR void VKAPI_CALL vkGetPrivateData(VkDevice device, VkObjectType objectType, uint64_t objectHandle, VkPrivateDataSlot privateDataSlot, uint64_t* pData) {
	return ptr_vkGetPrivateData(device, objectType, objectHandle, privateDataSlot, pData);
}
*/
/*
PFN_vkGetQueryPoolResults ptr_vkGetQueryPoolResults;
VKAPI_ATTR VkResult VKAPI_CALL vkGetQueryPoolResults(VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, size_t dataSize, void* pData, VkDeviceSize stride, VkQueryResultFlags flags) {
	return ptr_vkGetQueryPoolResults(device, queryPool, firstQuery, queryCount, dataSize, pData, stride, flags);
}
*/
/*
PFN_vkGetRenderAreaGranularity ptr_vkGetRenderAreaGranularity;
VKAPI_ATTR void VKAPI_CALL vkGetRenderAreaGranularity(VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity) {
	return ptr_vkGetRenderAreaGranularity(device, renderPass, pGranularity);
}
*/
/*
PFN_vkGetRenderingAreaGranularity ptr_vkGetRenderingAreaGranularity;
VKAPI_ATTR void VKAPI_CALL vkGetRenderingAreaGranularity(VkDevice device, VkRenderingAreaInfo* pRenderingAreaInfo, VkExtent2D* pGranularity) {
	return ptr_vkGetRenderingAreaGranularity(device, pRenderingAreaInfo, pGranularity);
}
*/
/*
PFN_vkGetSemaphoreCounterValue ptr_vkGetSemaphoreCounterValue;
VKAPI_ATTR VkResult VKAPI_CALL vkGetSemaphoreCounterValue(VkDevice device, VkSemaphore semaphore, uint64_t* pValue) {
	return ptr_vkGetSemaphoreCounterValue(device, semaphore, pValue);
}
*/
/*
PFN_vkInvalidateMappedMemoryRanges ptr_vkInvalidateMappedMemoryRanges;
VKAPI_ATTR VkResult VKAPI_CALL vkInvalidateMappedMemoryRanges(VkDevice device, uint32_t memoryRangeCount, VkMappedMemoryRange* pMemoryRanges) {
	return ptr_vkInvalidateMappedMemoryRanges(device, memoryRangeCount, pMemoryRanges);
}
*/
/*
PFN_vkMergePipelineCaches ptr_vkMergePipelineCaches;
VKAPI_ATTR VkResult VKAPI_CALL vkMergePipelineCaches(VkDevice device, VkPipelineCache dstCache, uint32_t srcCacheCount, VkPipelineCache* pSrcCaches) {
	return ptr_vkMergePipelineCaches(device, dstCache, srcCacheCount, pSrcCaches);
}
*/
/*
PFN_vkQueueBindSparse ptr_vkQueueBindSparse;
VKAPI_ATTR VkResult VKAPI_CALL vkQueueBindSparse(VkQueue queue, uint32_t bindInfoCount, VkBindSparseInfo* pBindInfo, VkFence fence) {
	return ptr_vkQueueBindSparse(queue, bindInfoCount, pBindInfo, fence);
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
	return ptr_vkResetQueryPool(device, queryPool, firstQuery, queryCount);
}
*/
/*
PFN_vkSetEvent ptr_vkSetEvent;
VKAPI_ATTR VkResult VKAPI_CALL vkSetEvent(VkDevice device, VkEvent event) {
	return ptr_vkSetEvent(device, event);
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
PFN_vkTransitionImageLayout ptr_vkTransitionImageLayout;
VKAPI_ATTR VkResult VKAPI_CALL vkTransitionImageLayout(VkDevice device, uint32_t transitionCount, VkHostImageLayoutTransitionInfo* pTransitions) {
	return ptr_vkTransitionImageLayout(device, transitionCount, pTransitions);
}
*/
/*
PFN_vkTrimCommandPool ptr_vkTrimCommandPool;
VKAPI_ATTR void VKAPI_CALL vkTrimCommandPool(VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
	return ptr_vkTrimCommandPool(device, commandPool, flags);
}
*/
/*
PFN_vkUnmapMemory ptr_vkUnmapMemory;
VKAPI_ATTR void VKAPI_CALL vkUnmapMemory(VkDevice device, VkDeviceMemory memory) {
	return ptr_vkUnmapMemory(device, memory);
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
	return ptr_vkUpdateDescriptorSetWithTemplate(device, descriptorSet, descriptorUpdateTemplate, pData);
}
*/
/*
PFN_vkUpdateDescriptorSets ptr_vkUpdateDescriptorSets;
VKAPI_ATTR void VKAPI_CALL vkUpdateDescriptorSets(VkDevice device, uint32_t descriptorWriteCount, VkWriteDescriptorSet* pDescriptorWrites, uint32_t descriptorCopyCount, VkCopyDescriptorSet* pDescriptorCopies) {
	return ptr_vkUpdateDescriptorSets(device, descriptorWriteCount, pDescriptorWrites, descriptorCopyCount, pDescriptorCopies);
}
*/
/*
PFN_vkWaitForFences ptr_vkWaitForFences;
VKAPI_ATTR VkResult VKAPI_CALL vkWaitForFences(VkDevice device, uint32_t fenceCount, VkFence* pFences, VkBool32 waitAll, uint64_t timeout) {
	return ptr_vkWaitForFences(device, fenceCount, pFences, waitAll, timeout);
}
*/
/*
PFN_vkWaitSemaphores ptr_vkWaitSemaphores;
VKAPI_ATTR VkResult VKAPI_CALL vkWaitSemaphores(VkDevice device, VkSemaphoreWaitInfo* pWaitInfo, uint64_t timeout) {
	return ptr_vkWaitSemaphores(device, pWaitInfo, timeout);
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
    ptr_vkAllocateCommandBuffers = (PFN_vkAllocateCommandBuffers) ptr_vkGetInstanceProcAddr(context, "vkAllocateCommandBuffers");
    ptr_vkAllocateDescriptorSets = (PFN_vkAllocateDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkAllocateDescriptorSets");
    ptr_vkAllocateMemory = (PFN_vkAllocateMemory) ptr_vkGetInstanceProcAddr(context, "vkAllocateMemory");
    ptr_vkBeginCommandBuffer = (PFN_vkBeginCommandBuffer) ptr_vkGetInstanceProcAddr(context, "vkBeginCommandBuffer");
    ptr_vkBindBufferMemory = (PFN_vkBindBufferMemory) ptr_vkGetInstanceProcAddr(context, "vkBindBufferMemory");
    ptr_vkBindBufferMemory2 = (PFN_vkBindBufferMemory2) ptr_vkGetInstanceProcAddr(context, "vkBindBufferMemory2");
    ptr_vkBindImageMemory = (PFN_vkBindImageMemory) ptr_vkGetInstanceProcAddr(context, "vkBindImageMemory");
    ptr_vkBindImageMemory2 = (PFN_vkBindImageMemory2) ptr_vkGetInstanceProcAddr(context, "vkBindImageMemory2");
    ptr_vkCmdBeginQuery = (PFN_vkCmdBeginQuery) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginQuery");
    ptr_vkCmdBeginRenderPass = (PFN_vkCmdBeginRenderPass) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginRenderPass");
    ptr_vkCmdBeginRenderPass2 = (PFN_vkCmdBeginRenderPass2) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginRenderPass2");
    ptr_vkCmdBeginRendering = (PFN_vkCmdBeginRendering) ptr_vkGetInstanceProcAddr(context, "vkCmdBeginRendering");
    ptr_vkCmdBindDescriptorSets = (PFN_vkCmdBindDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorSets");
    ptr_vkCmdBindDescriptorSets2 = (PFN_vkCmdBindDescriptorSets2) ptr_vkGetInstanceProcAddr(context, "vkCmdBindDescriptorSets2");
    ptr_vkCmdBindIndexBuffer = (PFN_vkCmdBindIndexBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdBindIndexBuffer");
    ptr_vkCmdBindIndexBuffer2 = (PFN_vkCmdBindIndexBuffer2) ptr_vkGetInstanceProcAddr(context, "vkCmdBindIndexBuffer2");
    ptr_vkCmdBindPipeline = (PFN_vkCmdBindPipeline) ptr_vkGetInstanceProcAddr(context, "vkCmdBindPipeline");
    ptr_vkCmdBindVertexBuffers = (PFN_vkCmdBindVertexBuffers) ptr_vkGetInstanceProcAddr(context, "vkCmdBindVertexBuffers");
    ptr_vkCmdBindVertexBuffers2 = (PFN_vkCmdBindVertexBuffers2) ptr_vkGetInstanceProcAddr(context, "vkCmdBindVertexBuffers2");
    ptr_vkCmdBlitImage = (PFN_vkCmdBlitImage) ptr_vkGetInstanceProcAddr(context, "vkCmdBlitImage");
    ptr_vkCmdBlitImage2 = (PFN_vkCmdBlitImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdBlitImage2");
    ptr_vkCmdClearAttachments = (PFN_vkCmdClearAttachments) ptr_vkGetInstanceProcAddr(context, "vkCmdClearAttachments");
    ptr_vkCmdClearDepthStencilImage = (PFN_vkCmdClearDepthStencilImage) ptr_vkGetInstanceProcAddr(context, "vkCmdClearDepthStencilImage");
    ptr_vkCmdCopyBuffer = (PFN_vkCmdCopyBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBuffer");
    ptr_vkCmdCopyBuffer2 = (PFN_vkCmdCopyBuffer2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBuffer2");
    ptr_vkCmdCopyBufferToImage = (PFN_vkCmdCopyBufferToImage) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBufferToImage");
    ptr_vkCmdCopyBufferToImage2 = (PFN_vkCmdCopyBufferToImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyBufferToImage2");
    ptr_vkCmdCopyImage = (PFN_vkCmdCopyImage) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImage");
    ptr_vkCmdCopyImage2 = (PFN_vkCmdCopyImage2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImage2");
    ptr_vkCmdCopyImageToBuffer = (PFN_vkCmdCopyImageToBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImageToBuffer");
    ptr_vkCmdCopyImageToBuffer2 = (PFN_vkCmdCopyImageToBuffer2) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyImageToBuffer2");
    ptr_vkCmdCopyQueryPoolResults = (PFN_vkCmdCopyQueryPoolResults) ptr_vkGetInstanceProcAddr(context, "vkCmdCopyQueryPoolResults");
    ptr_vkCmdDispatch = (PFN_vkCmdDispatch) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatch");
    ptr_vkCmdDispatchBase = (PFN_vkCmdDispatchBase) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatchBase");
    ptr_vkCmdDispatchIndirect = (PFN_vkCmdDispatchIndirect) ptr_vkGetInstanceProcAddr(context, "vkCmdDispatchIndirect");
    ptr_vkCmdDraw = (PFN_vkCmdDraw) ptr_vkGetInstanceProcAddr(context, "vkCmdDraw");
    ptr_vkCmdDrawIndexed = (PFN_vkCmdDrawIndexed) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndexed");
    ptr_vkCmdDrawIndexedIndirect = (PFN_vkCmdDrawIndexedIndirect) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndexedIndirect");
    ptr_vkCmdDrawIndexedIndirectCount = (PFN_vkCmdDrawIndexedIndirectCount) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndexedIndirectCount");
    ptr_vkCmdDrawIndirect = (PFN_vkCmdDrawIndirect) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndirect");
    ptr_vkCmdDrawIndirectCount = (PFN_vkCmdDrawIndirectCount) ptr_vkGetInstanceProcAddr(context, "vkCmdDrawIndirectCount");
    ptr_vkCmdEndQuery = (PFN_vkCmdEndQuery) ptr_vkGetInstanceProcAddr(context, "vkCmdEndQuery");
    ptr_vkCmdEndRenderPass = (PFN_vkCmdEndRenderPass) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRenderPass");
    ptr_vkCmdEndRenderPass2 = (PFN_vkCmdEndRenderPass2) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRenderPass2");
    ptr_vkCmdEndRendering = (PFN_vkCmdEndRendering) ptr_vkGetInstanceProcAddr(context, "vkCmdEndRendering");
    ptr_vkCmdExecuteCommands = (PFN_vkCmdExecuteCommands) ptr_vkGetInstanceProcAddr(context, "vkCmdExecuteCommands");
    ptr_vkCmdFillBuffer = (PFN_vkCmdFillBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdFillBuffer");
    ptr_vkCmdNextSubpass = (PFN_vkCmdNextSubpass) ptr_vkGetInstanceProcAddr(context, "vkCmdNextSubpass");
    ptr_vkCmdNextSubpass2 = (PFN_vkCmdNextSubpass2) ptr_vkGetInstanceProcAddr(context, "vkCmdNextSubpass2");
    ptr_vkCmdPipelineBarrier = (PFN_vkCmdPipelineBarrier) ptr_vkGetInstanceProcAddr(context, "vkCmdPipelineBarrier");
    ptr_vkCmdPipelineBarrier2 = (PFN_vkCmdPipelineBarrier2) ptr_vkGetInstanceProcAddr(context, "vkCmdPipelineBarrier2");
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
    ptr_vkCmdSetCullMode = (PFN_vkCmdSetCullMode) ptr_vkGetInstanceProcAddr(context, "vkCmdSetCullMode");
    ptr_vkCmdSetDepthBias = (PFN_vkCmdSetDepthBias) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBias");
    ptr_vkCmdSetDepthBiasEnable = (PFN_vkCmdSetDepthBiasEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBiasEnable");
    ptr_vkCmdSetDepthBounds = (PFN_vkCmdSetDepthBounds) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBounds");
    ptr_vkCmdSetDepthBoundsTestEnable = (PFN_vkCmdSetDepthBoundsTestEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthBoundsTestEnable");
    ptr_vkCmdSetDepthCompareOp = (PFN_vkCmdSetDepthCompareOp) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthCompareOp");
    ptr_vkCmdSetDepthTestEnable = (PFN_vkCmdSetDepthTestEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthTestEnable");
    ptr_vkCmdSetDepthWriteEnable = (PFN_vkCmdSetDepthWriteEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDepthWriteEnable");
    ptr_vkCmdSetDeviceMask = (PFN_vkCmdSetDeviceMask) ptr_vkGetInstanceProcAddr(context, "vkCmdSetDeviceMask");
    ptr_vkCmdSetEvent = (PFN_vkCmdSetEvent) ptr_vkGetInstanceProcAddr(context, "vkCmdSetEvent");
    ptr_vkCmdSetEvent2 = (PFN_vkCmdSetEvent2) ptr_vkGetInstanceProcAddr(context, "vkCmdSetEvent2");
    ptr_vkCmdSetFrontFace = (PFN_vkCmdSetFrontFace) ptr_vkGetInstanceProcAddr(context, "vkCmdSetFrontFace");
    ptr_vkCmdSetLineStipple = (PFN_vkCmdSetLineStipple) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLineStipple");
    ptr_vkCmdSetLineWidth = (PFN_vkCmdSetLineWidth) ptr_vkGetInstanceProcAddr(context, "vkCmdSetLineWidth");
    ptr_vkCmdSetPrimitiveRestartEnable = (PFN_vkCmdSetPrimitiveRestartEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPrimitiveRestartEnable");
    ptr_vkCmdSetPrimitiveTopology = (PFN_vkCmdSetPrimitiveTopology) ptr_vkGetInstanceProcAddr(context, "vkCmdSetPrimitiveTopology");
    ptr_vkCmdSetRasterizerDiscardEnable = (PFN_vkCmdSetRasterizerDiscardEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRasterizerDiscardEnable");
    ptr_vkCmdSetRenderingAttachmentLocations = (PFN_vkCmdSetRenderingAttachmentLocations) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRenderingAttachmentLocations");
    ptr_vkCmdSetRenderingInputAttachmentIndices = (PFN_vkCmdSetRenderingInputAttachmentIndices) ptr_vkGetInstanceProcAddr(context, "vkCmdSetRenderingInputAttachmentIndices");
    ptr_vkCmdSetScissor = (PFN_vkCmdSetScissor) ptr_vkGetInstanceProcAddr(context, "vkCmdSetScissor");
    ptr_vkCmdSetScissorWithCount = (PFN_vkCmdSetScissorWithCount) ptr_vkGetInstanceProcAddr(context, "vkCmdSetScissorWithCount");
    ptr_vkCmdSetStencilCompareMask = (PFN_vkCmdSetStencilCompareMask) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilCompareMask");
    ptr_vkCmdSetStencilOp = (PFN_vkCmdSetStencilOp) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilOp");
    ptr_vkCmdSetStencilReference = (PFN_vkCmdSetStencilReference) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilReference");
    ptr_vkCmdSetStencilTestEnable = (PFN_vkCmdSetStencilTestEnable) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilTestEnable");
    ptr_vkCmdSetStencilWriteMask = (PFN_vkCmdSetStencilWriteMask) ptr_vkGetInstanceProcAddr(context, "vkCmdSetStencilWriteMask");
    ptr_vkCmdSetViewport = (PFN_vkCmdSetViewport) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewport");
    ptr_vkCmdSetViewportWithCount = (PFN_vkCmdSetViewportWithCount) ptr_vkGetInstanceProcAddr(context, "vkCmdSetViewportWithCount");
    ptr_vkCmdUpdateBuffer = (PFN_vkCmdUpdateBuffer) ptr_vkGetInstanceProcAddr(context, "vkCmdUpdateBuffer");
    ptr_vkCmdWaitEvents = (PFN_vkCmdWaitEvents) ptr_vkGetInstanceProcAddr(context, "vkCmdWaitEvents");
    ptr_vkCmdWaitEvents2 = (PFN_vkCmdWaitEvents2) ptr_vkGetInstanceProcAddr(context, "vkCmdWaitEvents2");
    ptr_vkCmdWriteTimestamp = (PFN_vkCmdWriteTimestamp) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteTimestamp");
    ptr_vkCmdWriteTimestamp2 = (PFN_vkCmdWriteTimestamp2) ptr_vkGetInstanceProcAddr(context, "vkCmdWriteTimestamp2");
    ptr_vkCopyImageToImage = (PFN_vkCopyImageToImage) ptr_vkGetInstanceProcAddr(context, "vkCopyImageToImage");
    ptr_vkCopyImageToMemory = (PFN_vkCopyImageToMemory) ptr_vkGetInstanceProcAddr(context, "vkCopyImageToMemory");
    ptr_vkCopyMemoryToImage = (PFN_vkCopyMemoryToImage) ptr_vkGetInstanceProcAddr(context, "vkCopyMemoryToImage");
    ptr_vkCreateBuffer = (PFN_vkCreateBuffer) ptr_vkGetInstanceProcAddr(context, "vkCreateBuffer");
    ptr_vkCreateBufferView = (PFN_vkCreateBufferView) ptr_vkGetInstanceProcAddr(context, "vkCreateBufferView");
    ptr_vkCreateCommandPool = (PFN_vkCreateCommandPool) ptr_vkGetInstanceProcAddr(context, "vkCreateCommandPool");
    ptr_vkCreateComputePipelines = (PFN_vkCreateComputePipelines) ptr_vkGetInstanceProcAddr(context, "vkCreateComputePipelines");
    ptr_vkCreateDescriptorPool = (PFN_vkCreateDescriptorPool) ptr_vkGetInstanceProcAddr(context, "vkCreateDescriptorPool");
    ptr_vkCreateDescriptorSetLayout = (PFN_vkCreateDescriptorSetLayout) ptr_vkGetInstanceProcAddr(context, "vkCreateDescriptorSetLayout");
    ptr_vkCreateDescriptorUpdateTemplate = (PFN_vkCreateDescriptorUpdateTemplate) ptr_vkGetInstanceProcAddr(context, "vkCreateDescriptorUpdateTemplate");
    ptr_vkCreateDevice = (PFN_vkCreateDevice) ptr_vkGetInstanceProcAddr(context, "vkCreateDevice");
    ptr_vkCreateEvent = (PFN_vkCreateEvent) ptr_vkGetInstanceProcAddr(context, "vkCreateEvent");
    ptr_vkCreateFence = (PFN_vkCreateFence) ptr_vkGetInstanceProcAddr(context, "vkCreateFence");
    ptr_vkCreateFramebuffer = (PFN_vkCreateFramebuffer) ptr_vkGetInstanceProcAddr(context, "vkCreateFramebuffer");
    ptr_vkCreateGraphicsPipelines = (PFN_vkCreateGraphicsPipelines) ptr_vkGetInstanceProcAddr(context, "vkCreateGraphicsPipelines");
    ptr_vkCreateImage = (PFN_vkCreateImage) ptr_vkGetInstanceProcAddr(context, "vkCreateImage");
    ptr_vkCreateImageView = (PFN_vkCreateImageView) ptr_vkGetInstanceProcAddr(context, "vkCreateImageView");
    ptr_vkCreatePipelineCache = (PFN_vkCreatePipelineCache) ptr_vkGetInstanceProcAddr(context, "vkCreatePipelineCache");
    ptr_vkCreatePipelineLayout = (PFN_vkCreatePipelineLayout) ptr_vkGetInstanceProcAddr(context, "vkCreatePipelineLayout");
    ptr_vkCreatePrivateDataSlot = (PFN_vkCreatePrivateDataSlot) ptr_vkGetInstanceProcAddr(context, "vkCreatePrivateDataSlot");
    ptr_vkCreateQueryPool = (PFN_vkCreateQueryPool) ptr_vkGetInstanceProcAddr(context, "vkCreateQueryPool");
    ptr_vkCreateRenderPass = (PFN_vkCreateRenderPass) ptr_vkGetInstanceProcAddr(context, "vkCreateRenderPass");
    ptr_vkCreateRenderPass2 = (PFN_vkCreateRenderPass2) ptr_vkGetInstanceProcAddr(context, "vkCreateRenderPass2");
    ptr_vkCreateSampler = (PFN_vkCreateSampler) ptr_vkGetInstanceProcAddr(context, "vkCreateSampler");
    ptr_vkCreateSamplerYcbcrConversion = (PFN_vkCreateSamplerYcbcrConversion) ptr_vkGetInstanceProcAddr(context, "vkCreateSamplerYcbcrConversion");
    ptr_vkCreateSemaphore = (PFN_vkCreateSemaphore) ptr_vkGetInstanceProcAddr(context, "vkCreateSemaphore");
    ptr_vkCreateShaderModule = (PFN_vkCreateShaderModule) ptr_vkGetInstanceProcAddr(context, "vkCreateShaderModule");
    ptr_vkDestroyBuffer = (PFN_vkDestroyBuffer) ptr_vkGetInstanceProcAddr(context, "vkDestroyBuffer");
    ptr_vkDestroyBufferView = (PFN_vkDestroyBufferView) ptr_vkGetInstanceProcAddr(context, "vkDestroyBufferView");
    ptr_vkDestroyCommandPool = (PFN_vkDestroyCommandPool) ptr_vkGetInstanceProcAddr(context, "vkDestroyCommandPool");
    ptr_vkDestroyDescriptorPool = (PFN_vkDestroyDescriptorPool) ptr_vkGetInstanceProcAddr(context, "vkDestroyDescriptorPool");
    ptr_vkDestroyDescriptorSetLayout = (PFN_vkDestroyDescriptorSetLayout) ptr_vkGetInstanceProcAddr(context, "vkDestroyDescriptorSetLayout");
    ptr_vkDestroyDescriptorUpdateTemplate = (PFN_vkDestroyDescriptorUpdateTemplate) ptr_vkGetInstanceProcAddr(context, "vkDestroyDescriptorUpdateTemplate");
    ptr_vkDestroyDevice = (PFN_vkDestroyDevice) ptr_vkGetInstanceProcAddr(context, "vkDestroyDevice");
    ptr_vkDestroyEvent = (PFN_vkDestroyEvent) ptr_vkGetInstanceProcAddr(context, "vkDestroyEvent");
    ptr_vkDestroyFence = (PFN_vkDestroyFence) ptr_vkGetInstanceProcAddr(context, "vkDestroyFence");
    ptr_vkDestroyFramebuffer = (PFN_vkDestroyFramebuffer) ptr_vkGetInstanceProcAddr(context, "vkDestroyFramebuffer");
    ptr_vkDestroyImage = (PFN_vkDestroyImage) ptr_vkGetInstanceProcAddr(context, "vkDestroyImage");
    ptr_vkDestroyImageView = (PFN_vkDestroyImageView) ptr_vkGetInstanceProcAddr(context, "vkDestroyImageView");
    ptr_vkDestroyInstance = (PFN_vkDestroyInstance) ptr_vkGetInstanceProcAddr(context, "vkDestroyInstance");
    ptr_vkDestroyPipeline = (PFN_vkDestroyPipeline) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipeline");
    ptr_vkDestroyPipelineCache = (PFN_vkDestroyPipelineCache) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipelineCache");
    ptr_vkDestroyPipelineLayout = (PFN_vkDestroyPipelineLayout) ptr_vkGetInstanceProcAddr(context, "vkDestroyPipelineLayout");
    ptr_vkDestroyPrivateDataSlot = (PFN_vkDestroyPrivateDataSlot) ptr_vkGetInstanceProcAddr(context, "vkDestroyPrivateDataSlot");
    ptr_vkDestroyQueryPool = (PFN_vkDestroyQueryPool) ptr_vkGetInstanceProcAddr(context, "vkDestroyQueryPool");
    ptr_vkDestroyRenderPass = (PFN_vkDestroyRenderPass) ptr_vkGetInstanceProcAddr(context, "vkDestroyRenderPass");
    ptr_vkDestroySampler = (PFN_vkDestroySampler) ptr_vkGetInstanceProcAddr(context, "vkDestroySampler");
    ptr_vkDestroySamplerYcbcrConversion = (PFN_vkDestroySamplerYcbcrConversion) ptr_vkGetInstanceProcAddr(context, "vkDestroySamplerYcbcrConversion");
    ptr_vkDestroySemaphore = (PFN_vkDestroySemaphore) ptr_vkGetInstanceProcAddr(context, "vkDestroySemaphore");
    ptr_vkDestroyShaderModule = (PFN_vkDestroyShaderModule) ptr_vkGetInstanceProcAddr(context, "vkDestroyShaderModule");
    ptr_vkDeviceWaitIdle = (PFN_vkDeviceWaitIdle) ptr_vkGetInstanceProcAddr(context, "vkDeviceWaitIdle");
    ptr_vkEndCommandBuffer = (PFN_vkEndCommandBuffer) ptr_vkGetInstanceProcAddr(context, "vkEndCommandBuffer");
    ptr_vkEnumerateDeviceExtensionProperties = (PFN_vkEnumerateDeviceExtensionProperties) ptr_vkGetInstanceProcAddr(context, "vkEnumerateDeviceExtensionProperties");
    ptr_vkEnumerateDeviceLayerProperties = (PFN_vkEnumerateDeviceLayerProperties) ptr_vkGetInstanceProcAddr(context, "vkEnumerateDeviceLayerProperties");
    ptr_vkEnumeratePhysicalDeviceGroups = (PFN_vkEnumeratePhysicalDeviceGroups) ptr_vkGetInstanceProcAddr(context, "vkEnumeratePhysicalDeviceGroups");
    ptr_vkEnumeratePhysicalDevices = (PFN_vkEnumeratePhysicalDevices) ptr_vkGetInstanceProcAddr(context, "vkEnumeratePhysicalDevices");
    ptr_vkFlushMappedMemoryRanges = (PFN_vkFlushMappedMemoryRanges) ptr_vkGetInstanceProcAddr(context, "vkFlushMappedMemoryRanges");
    ptr_vkFreeCommandBuffers = (PFN_vkFreeCommandBuffers) ptr_vkGetInstanceProcAddr(context, "vkFreeCommandBuffers");
    ptr_vkFreeDescriptorSets = (PFN_vkFreeDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkFreeDescriptorSets");
    ptr_vkFreeMemory = (PFN_vkFreeMemory) ptr_vkGetInstanceProcAddr(context, "vkFreeMemory");
    ptr_vkGetBufferDeviceAddress = (PFN_vkGetBufferDeviceAddress) ptr_vkGetInstanceProcAddr(context, "vkGetBufferDeviceAddress");
    ptr_vkGetBufferMemoryRequirements = (PFN_vkGetBufferMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetBufferMemoryRequirements");
    ptr_vkGetBufferMemoryRequirements2 = (PFN_vkGetBufferMemoryRequirements2) ptr_vkGetInstanceProcAddr(context, "vkGetBufferMemoryRequirements2");
    ptr_vkGetBufferOpaqueCaptureAddress = (PFN_vkGetBufferOpaqueCaptureAddress) ptr_vkGetInstanceProcAddr(context, "vkGetBufferOpaqueCaptureAddress");
    ptr_vkGetDescriptorSetLayoutSupport = (PFN_vkGetDescriptorSetLayoutSupport) ptr_vkGetInstanceProcAddr(context, "vkGetDescriptorSetLayoutSupport");
    ptr_vkGetDeviceBufferMemoryRequirements = (PFN_vkGetDeviceBufferMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceBufferMemoryRequirements");
    ptr_vkGetDeviceGroupPeerMemoryFeatures = (PFN_vkGetDeviceGroupPeerMemoryFeatures) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceGroupPeerMemoryFeatures");
    ptr_vkGetDeviceImageMemoryRequirements = (PFN_vkGetDeviceImageMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceImageMemoryRequirements");
    ptr_vkGetDeviceImageSparseMemoryRequirements = (PFN_vkGetDeviceImageSparseMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceImageSparseMemoryRequirements");
    ptr_vkGetDeviceImageSubresourceLayout = (PFN_vkGetDeviceImageSubresourceLayout) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceImageSubresourceLayout");
    ptr_vkGetDeviceMemoryCommitment = (PFN_vkGetDeviceMemoryCommitment) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceMemoryCommitment");
    ptr_vkGetDeviceMemoryOpaqueCaptureAddress = (PFN_vkGetDeviceMemoryOpaqueCaptureAddress) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceMemoryOpaqueCaptureAddress");
    ptr_vkGetDeviceQueue = (PFN_vkGetDeviceQueue) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceQueue");
    ptr_vkGetDeviceQueue2 = (PFN_vkGetDeviceQueue2) ptr_vkGetInstanceProcAddr(context, "vkGetDeviceQueue2");
    ptr_vkGetEventStatus = (PFN_vkGetEventStatus) ptr_vkGetInstanceProcAddr(context, "vkGetEventStatus");
    ptr_vkGetFenceStatus = (PFN_vkGetFenceStatus) ptr_vkGetInstanceProcAddr(context, "vkGetFenceStatus");
    ptr_vkGetImageMemoryRequirements = (PFN_vkGetImageMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetImageMemoryRequirements");
    ptr_vkGetImageMemoryRequirements2 = (PFN_vkGetImageMemoryRequirements2) ptr_vkGetInstanceProcAddr(context, "vkGetImageMemoryRequirements2");
    ptr_vkGetImageSparseMemoryRequirements = (PFN_vkGetImageSparseMemoryRequirements) ptr_vkGetInstanceProcAddr(context, "vkGetImageSparseMemoryRequirements");
    ptr_vkGetImageSparseMemoryRequirements2 = (PFN_vkGetImageSparseMemoryRequirements2) ptr_vkGetInstanceProcAddr(context, "vkGetImageSparseMemoryRequirements2");
    ptr_vkGetImageSubresourceLayout = (PFN_vkGetImageSubresourceLayout) ptr_vkGetInstanceProcAddr(context, "vkGetImageSubresourceLayout");
    ptr_vkGetImageSubresourceLayout2 = (PFN_vkGetImageSubresourceLayout2) ptr_vkGetInstanceProcAddr(context, "vkGetImageSubresourceLayout2");
    ptr_vkGetPhysicalDeviceExternalBufferProperties = (PFN_vkGetPhysicalDeviceExternalBufferProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalBufferProperties");
    ptr_vkGetPhysicalDeviceExternalFenceProperties = (PFN_vkGetPhysicalDeviceExternalFenceProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalFenceProperties");
    ptr_vkGetPhysicalDeviceExternalSemaphoreProperties = (PFN_vkGetPhysicalDeviceExternalSemaphoreProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceExternalSemaphoreProperties");
    ptr_vkGetPhysicalDeviceFeatures = (PFN_vkGetPhysicalDeviceFeatures) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFeatures");
    ptr_vkGetPhysicalDeviceFeatures2 = (PFN_vkGetPhysicalDeviceFeatures2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFeatures2");
    ptr_vkGetPhysicalDeviceFormatProperties = (PFN_vkGetPhysicalDeviceFormatProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFormatProperties");
    ptr_vkGetPhysicalDeviceFormatProperties2 = (PFN_vkGetPhysicalDeviceFormatProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceFormatProperties2");
    ptr_vkGetPhysicalDeviceImageFormatProperties = (PFN_vkGetPhysicalDeviceImageFormatProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceImageFormatProperties");
    ptr_vkGetPhysicalDeviceImageFormatProperties2 = (PFN_vkGetPhysicalDeviceImageFormatProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceImageFormatProperties2");
    ptr_vkGetPhysicalDeviceMemoryProperties = (PFN_vkGetPhysicalDeviceMemoryProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceMemoryProperties");
    ptr_vkGetPhysicalDeviceMemoryProperties2 = (PFN_vkGetPhysicalDeviceMemoryProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceMemoryProperties2");
    ptr_vkGetPhysicalDeviceProperties = (PFN_vkGetPhysicalDeviceProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceProperties");
    ptr_vkGetPhysicalDeviceProperties2 = (PFN_vkGetPhysicalDeviceProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceProperties2");
    ptr_vkGetPhysicalDeviceQueueFamilyProperties = (PFN_vkGetPhysicalDeviceQueueFamilyProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyProperties");
    ptr_vkGetPhysicalDeviceQueueFamilyProperties2 = (PFN_vkGetPhysicalDeviceQueueFamilyProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceQueueFamilyProperties2");
    ptr_vkGetPhysicalDeviceSparseImageFormatProperties = (PFN_vkGetPhysicalDeviceSparseImageFormatProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSparseImageFormatProperties");
    ptr_vkGetPhysicalDeviceSparseImageFormatProperties2 = (PFN_vkGetPhysicalDeviceSparseImageFormatProperties2) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceSparseImageFormatProperties2");
    ptr_vkGetPhysicalDeviceToolProperties = (PFN_vkGetPhysicalDeviceToolProperties) ptr_vkGetInstanceProcAddr(context, "vkGetPhysicalDeviceToolProperties");
    ptr_vkGetPipelineCacheData = (PFN_vkGetPipelineCacheData) ptr_vkGetInstanceProcAddr(context, "vkGetPipelineCacheData");
    ptr_vkGetPrivateData = (PFN_vkGetPrivateData) ptr_vkGetInstanceProcAddr(context, "vkGetPrivateData");
    ptr_vkGetQueryPoolResults = (PFN_vkGetQueryPoolResults) ptr_vkGetInstanceProcAddr(context, "vkGetQueryPoolResults");
    ptr_vkGetRenderAreaGranularity = (PFN_vkGetRenderAreaGranularity) ptr_vkGetInstanceProcAddr(context, "vkGetRenderAreaGranularity");
    ptr_vkGetRenderingAreaGranularity = (PFN_vkGetRenderingAreaGranularity) ptr_vkGetInstanceProcAddr(context, "vkGetRenderingAreaGranularity");
    ptr_vkGetSemaphoreCounterValue = (PFN_vkGetSemaphoreCounterValue) ptr_vkGetInstanceProcAddr(context, "vkGetSemaphoreCounterValue");
    ptr_vkInvalidateMappedMemoryRanges = (PFN_vkInvalidateMappedMemoryRanges) ptr_vkGetInstanceProcAddr(context, "vkInvalidateMappedMemoryRanges");
    ptr_vkMergePipelineCaches = (PFN_vkMergePipelineCaches) ptr_vkGetInstanceProcAddr(context, "vkMergePipelineCaches");
    ptr_vkQueueBindSparse = (PFN_vkQueueBindSparse) ptr_vkGetInstanceProcAddr(context, "vkQueueBindSparse");
    ptr_vkQueueSubmit = (PFN_vkQueueSubmit) ptr_vkGetInstanceProcAddr(context, "vkQueueSubmit");
    ptr_vkQueueSubmit2 = (PFN_vkQueueSubmit2) ptr_vkGetInstanceProcAddr(context, "vkQueueSubmit2");
    ptr_vkQueueWaitIdle = (PFN_vkQueueWaitIdle) ptr_vkGetInstanceProcAddr(context, "vkQueueWaitIdle");
    ptr_vkResetCommandBuffer = (PFN_vkResetCommandBuffer) ptr_vkGetInstanceProcAddr(context, "vkResetCommandBuffer");
    ptr_vkResetCommandPool = (PFN_vkResetCommandPool) ptr_vkGetInstanceProcAddr(context, "vkResetCommandPool");
    ptr_vkResetDescriptorPool = (PFN_vkResetDescriptorPool) ptr_vkGetInstanceProcAddr(context, "vkResetDescriptorPool");
    ptr_vkResetEvent = (PFN_vkResetEvent) ptr_vkGetInstanceProcAddr(context, "vkResetEvent");
    ptr_vkResetFences = (PFN_vkResetFences) ptr_vkGetInstanceProcAddr(context, "vkResetFences");
    ptr_vkResetQueryPool = (PFN_vkResetQueryPool) ptr_vkGetInstanceProcAddr(context, "vkResetQueryPool");
    ptr_vkSetEvent = (PFN_vkSetEvent) ptr_vkGetInstanceProcAddr(context, "vkSetEvent");
    ptr_vkSetPrivateData = (PFN_vkSetPrivateData) ptr_vkGetInstanceProcAddr(context, "vkSetPrivateData");
    ptr_vkSignalSemaphore = (PFN_vkSignalSemaphore) ptr_vkGetInstanceProcAddr(context, "vkSignalSemaphore");
    ptr_vkTransitionImageLayout = (PFN_vkTransitionImageLayout) ptr_vkGetInstanceProcAddr(context, "vkTransitionImageLayout");
    ptr_vkTrimCommandPool = (PFN_vkTrimCommandPool) ptr_vkGetInstanceProcAddr(context, "vkTrimCommandPool");
    ptr_vkUnmapMemory = (PFN_vkUnmapMemory) ptr_vkGetInstanceProcAddr(context, "vkUnmapMemory");
    ptr_vkUnmapMemory2 = (PFN_vkUnmapMemory2) ptr_vkGetInstanceProcAddr(context, "vkUnmapMemory2");
    ptr_vkUpdateDescriptorSetWithTemplate = (PFN_vkUpdateDescriptorSetWithTemplate) ptr_vkGetInstanceProcAddr(context, "vkUpdateDescriptorSetWithTemplate");
    ptr_vkUpdateDescriptorSets = (PFN_vkUpdateDescriptorSets) ptr_vkGetInstanceProcAddr(context, "vkUpdateDescriptorSets");
    ptr_vkWaitForFences = (PFN_vkWaitForFences) ptr_vkGetInstanceProcAddr(context, "vkWaitForFences");
    ptr_vkWaitSemaphores = (PFN_vkWaitSemaphores) ptr_vkGetInstanceProcAddr(context, "vkWaitSemaphores");
}
*/
import "C"

// Load attempts to load all global vulkan functions.
func Load(loader unsafe.Pointer) {
	C.gfx_vkInit(loader)
}

// LoadInstance attempts to load all instance vulkan functions.
func LoadInstance(instance Instance) {
	C.gfx_vkInitInstance(C.VkInstance(unsafe.Pointer(instance)))
}

// AllocateCommandBuffers wraps vkAllocateCommandBuffers.
func AllocateCommandBuffers(device Device, pAllocateInfo CommandBufferAllocateInfo, pCommandBuffers ffi.Ref[CommandBuffer]) Result {
	ret := C.vkAllocateCommandBuffers(C.VkDevice(unsafe.Pointer(device)), pAllocateInfo.ptr, (*C.VkCommandBuffer)(pCommandBuffers.Raw()))

	return Result(ret)
}

// AllocateDescriptorSets wraps vkAllocateDescriptorSets.
func AllocateDescriptorSets(device Device, pAllocateInfo DescriptorSetAllocateInfo, pDescriptorSets ffi.Ref[DescriptorSet]) Result {
	ret := C.vkAllocateDescriptorSets(C.VkDevice(unsafe.Pointer(device)), pAllocateInfo.ptr, (*C.VkDescriptorSet)(pDescriptorSets.Raw()))

	return Result(ret)
}

// AllocateMemory wraps vkAllocateMemory.
func AllocateMemory(device Device, pAllocateInfo MemoryAllocateInfo, pAllocator AllocationCallbacks, pMemory ffi.Ref[DeviceMemory]) Result {
	ret := C.vkAllocateMemory(C.VkDevice(unsafe.Pointer(device)), pAllocateInfo.ptr, pAllocator.ptr, (*C.VkDeviceMemory)(pMemory.Raw()))

	return Result(ret)
}

// BeginCommandBuffer wraps vkBeginCommandBuffer.
func BeginCommandBuffer(commandBuffer CommandBuffer, pBeginInfo CommandBufferBeginInfo) Result {
	ret := C.vkBeginCommandBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pBeginInfo.ptr)

	return Result(ret)
}

// BindBufferMemory wraps vkBindBufferMemory.
func BindBufferMemory(device Device, buffer Buffer, memory DeviceMemory, memoryOffset DeviceSize) Result {
	ret := C.vkBindBufferMemory(C.VkDevice(unsafe.Pointer(device)), C.VkBuffer(buffer), C.VkDeviceMemory(memory), C.VkDeviceSize(memoryOffset))

	return Result(ret)
}

// BindBufferMemory2 wraps vkBindBufferMemory2.
func BindBufferMemory2(device Device, bindInfoCount uint32, pBindInfos BindBufferMemoryInfo) Result {
	ret := C.vkBindBufferMemory2(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), pBindInfos.ptr)

	return Result(ret)
}

// BindImageMemory wraps vkBindImageMemory.
func BindImageMemory(device Device, image Image, memory DeviceMemory, memoryOffset DeviceSize) Result {
	ret := C.vkBindImageMemory(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), C.VkDeviceMemory(memory), C.VkDeviceSize(memoryOffset))

	return Result(ret)
}

// BindImageMemory2 wraps vkBindImageMemory2.
func BindImageMemory2(device Device, bindInfoCount uint32, pBindInfos BindImageMemoryInfo) Result {
	ret := C.vkBindImageMemory2(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(bindInfoCount), pBindInfos.ptr)

	return Result(ret)
}

// CmdBeginQuery wraps vkCmdBeginQuery.
func CmdBeginQuery(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, flags QueryControlFlags) {
	C.vkCmdBeginQuery(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(query), C.VkQueryControlFlags(flags))
}

// CmdBeginRenderPass wraps vkCmdBeginRenderPass.
func CmdBeginRenderPass(commandBuffer CommandBuffer, pRenderPassBegin RenderPassBeginInfo, contents SubpassContents) {
	C.vkCmdBeginRenderPass(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pRenderPassBegin.ptr, C.VkSubpassContents(contents))
}

// CmdBeginRenderPass2 wraps vkCmdBeginRenderPass2.
func CmdBeginRenderPass2(commandBuffer CommandBuffer, pRenderPassBegin RenderPassBeginInfo, pSubpassBeginInfo SubpassBeginInfo) {
	C.vkCmdBeginRenderPass2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pRenderPassBegin.ptr, pSubpassBeginInfo.ptr)
}

// CmdBeginRendering wraps vkCmdBeginRendering.
func CmdBeginRendering(commandBuffer CommandBuffer, pRenderingInfo RenderingInfo) {
	C.vkCmdBeginRendering(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pRenderingInfo.ptr)
}

// CmdBindDescriptorSets wraps vkCmdBindDescriptorSets.
func CmdBindDescriptorSets(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, firstSet uint32, descriptorSetCount uint32, pDescriptorSets ffi.Ref[DescriptorSet], dynamicOffsetCount uint32, pDynamicOffsets ffi.Ref[uint32]) {
	C.vkCmdBindDescriptorSets(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipelineLayout(layout), C.uint32_t(firstSet), C.uint32_t(descriptorSetCount), (*C.VkDescriptorSet)(pDescriptorSets.Raw()), C.uint32_t(dynamicOffsetCount), (*C.uint32_t)(pDynamicOffsets.Raw()))
}

// CmdBindDescriptorSets2 wraps vkCmdBindDescriptorSets2.
func CmdBindDescriptorSets2(commandBuffer CommandBuffer, pBindDescriptorSetsInfo BindDescriptorSetsInfo) {
	C.vkCmdBindDescriptorSets2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pBindDescriptorSetsInfo.ptr)
}

// CmdBindIndexBuffer wraps vkCmdBindIndexBuffer.
func CmdBindIndexBuffer(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, indexType IndexType) {
	C.vkCmdBindIndexBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkIndexType(indexType))
}

// CmdBindIndexBuffer2 wraps vkCmdBindIndexBuffer2.
func CmdBindIndexBuffer2(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, size DeviceSize, indexType IndexType) {
	C.vkCmdBindIndexBuffer2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkDeviceSize(size), C.VkIndexType(indexType))
}

// CmdBindPipeline wraps vkCmdBindPipeline.
func CmdBindPipeline(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	C.vkCmdBindPipeline(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipeline(pipeline))
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
	C.vkCmdBlitImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), pRegions.ptr, C.VkFilter(filter))
}

// CmdBlitImage2 wraps vkCmdBlitImage2.
func CmdBlitImage2(commandBuffer CommandBuffer, pBlitImageInfo BlitImageInfo2) {
	C.vkCmdBlitImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pBlitImageInfo.ptr)
}

// CmdClearAttachments wraps vkCmdClearAttachments.
func CmdClearAttachments(commandBuffer CommandBuffer, attachmentCount uint32, pAttachments ClearAttachment, rectCount uint32, pRects ClearRect) {
	C.vkCmdClearAttachments(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(attachmentCount), pAttachments.ptr, C.uint32_t(rectCount), pRects.ptr)
}

// vkCmdClearColorImage.pColor is unsupported: category pointer -> ?? VkClearColorValue.

// CmdClearDepthStencilImage wraps vkCmdClearDepthStencilImage.
func CmdClearDepthStencilImage(commandBuffer CommandBuffer, image Image, imageLayout ImageLayout, pDepthStencil ClearDepthStencilValue, rangeCount uint32, pRanges ImageSubresourceRange) {
	C.vkCmdClearDepthStencilImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(image), C.VkImageLayout(imageLayout), pDepthStencil.ptr, C.uint32_t(rangeCount), pRanges.ptr)
}

// CmdCopyBuffer wraps vkCmdCopyBuffer.
func CmdCopyBuffer(commandBuffer CommandBuffer, srcBuffer Buffer, dstBuffer Buffer, regionCount uint32, pRegions BufferCopy) {
	C.vkCmdCopyBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(srcBuffer), C.VkBuffer(dstBuffer), C.uint32_t(regionCount), pRegions.ptr)
}

// CmdCopyBuffer2 wraps vkCmdCopyBuffer2.
func CmdCopyBuffer2(commandBuffer CommandBuffer, pCopyBufferInfo CopyBufferInfo2) {
	C.vkCmdCopyBuffer2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pCopyBufferInfo.ptr)
}

// CmdCopyBufferToImage wraps vkCmdCopyBufferToImage.
func CmdCopyBufferToImage(commandBuffer CommandBuffer, srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions BufferImageCopy) {
	C.vkCmdCopyBufferToImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(srcBuffer), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), pRegions.ptr)
}

// CmdCopyBufferToImage2 wraps vkCmdCopyBufferToImage2.
func CmdCopyBufferToImage2(commandBuffer CommandBuffer, pCopyBufferToImageInfo CopyBufferToImageInfo2) {
	C.vkCmdCopyBufferToImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pCopyBufferToImageInfo.ptr)
}

// CmdCopyImage wraps vkCmdCopyImage.
func CmdCopyImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions ImageCopy) {
	C.vkCmdCopyImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), pRegions.ptr)
}

// CmdCopyImage2 wraps vkCmdCopyImage2.
func CmdCopyImage2(commandBuffer CommandBuffer, pCopyImageInfo CopyImageInfo2) {
	C.vkCmdCopyImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pCopyImageInfo.ptr)
}

// CmdCopyImageToBuffer wraps vkCmdCopyImageToBuffer.
func CmdCopyImageToBuffer(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regionCount uint32, pRegions BufferImageCopy) {
	C.vkCmdCopyImageToBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkBuffer(dstBuffer), C.uint32_t(regionCount), pRegions.ptr)
}

// CmdCopyImageToBuffer2 wraps vkCmdCopyImageToBuffer2.
func CmdCopyImageToBuffer2(commandBuffer CommandBuffer, pCopyImageToBufferInfo CopyImageToBufferInfo2) {
	C.vkCmdCopyImageToBuffer2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pCopyImageToBufferInfo.ptr)
}

// CmdCopyQueryPoolResults wraps vkCmdCopyQueryPoolResults.
func CmdCopyQueryPoolResults(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery uint32, queryCount uint32, dstBuffer Buffer, dstOffset DeviceSize, stride DeviceSize, flags QueryResultFlags) {
	C.vkCmdCopyQueryPoolResults(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(firstQuery), C.uint32_t(queryCount), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.VkDeviceSize(stride), C.VkQueryResultFlags(flags))
}

// CmdDispatch wraps vkCmdDispatch.
func CmdDispatch(commandBuffer CommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.vkCmdDispatch(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(groupCountX), C.uint32_t(groupCountY), C.uint32_t(groupCountZ))
}

// CmdDispatchBase wraps vkCmdDispatchBase.
func CmdDispatchBase(commandBuffer CommandBuffer, baseGroupX uint32, baseGroupY uint32, baseGroupZ uint32, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.vkCmdDispatchBase(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(baseGroupX), C.uint32_t(baseGroupY), C.uint32_t(baseGroupZ), C.uint32_t(groupCountX), C.uint32_t(groupCountY), C.uint32_t(groupCountZ))
}

// CmdDispatchIndirect wraps vkCmdDispatchIndirect.
func CmdDispatchIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize) {
	C.vkCmdDispatchIndirect(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset))
}

// CmdDraw wraps vkCmdDraw.
func CmdDraw(commandBuffer CommandBuffer, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	C.vkCmdDraw(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
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

// CmdDrawIndirectCount wraps vkCmdDrawIndirectCount.
func CmdDrawIndirectCount(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	C.vkCmdDrawIndirectCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(buffer), C.VkDeviceSize(offset), C.VkBuffer(countBuffer), C.VkDeviceSize(countBufferOffset), C.uint32_t(maxDrawCount), C.uint32_t(stride))
}

// CmdEndQuery wraps vkCmdEndQuery.
func CmdEndQuery(commandBuffer CommandBuffer, queryPool QueryPool, query uint32) {
	C.vkCmdEndQuery(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkQueryPool(queryPool), C.uint32_t(query))
}

// CmdEndRenderPass wraps vkCmdEndRenderPass.
func CmdEndRenderPass(commandBuffer CommandBuffer) {
	C.vkCmdEndRenderPass(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdEndRenderPass2 wraps vkCmdEndRenderPass2.
func CmdEndRenderPass2(commandBuffer CommandBuffer, pSubpassEndInfo SubpassEndInfo) {
	C.vkCmdEndRenderPass2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pSubpassEndInfo.ptr)
}

// CmdEndRendering wraps vkCmdEndRendering.
func CmdEndRendering(commandBuffer CommandBuffer) {
	C.vkCmdEndRendering(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

// CmdExecuteCommands wraps vkCmdExecuteCommands.
func CmdExecuteCommands(commandBuffer CommandBuffer, commandBufferCount uint32, pCommandBuffers ffi.Ref[CommandBuffer]) {
	C.vkCmdExecuteCommands(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(commandBufferCount), (*C.VkCommandBuffer)(pCommandBuffers.Raw()))
}

// CmdFillBuffer wraps vkCmdFillBuffer.
func CmdFillBuffer(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset DeviceSize, size DeviceSize, data uint32) {
	C.vkCmdFillBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.VkDeviceSize(size), C.uint32_t(data))
}

// CmdNextSubpass wraps vkCmdNextSubpass.
func CmdNextSubpass(commandBuffer CommandBuffer, contents SubpassContents) {
	C.vkCmdNextSubpass(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkSubpassContents(contents))
}

// CmdNextSubpass2 wraps vkCmdNextSubpass2.
func CmdNextSubpass2(commandBuffer CommandBuffer, pSubpassBeginInfo SubpassBeginInfo, pSubpassEndInfo SubpassEndInfo) {
	C.vkCmdNextSubpass2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pSubpassBeginInfo.ptr, pSubpassEndInfo.ptr)
}

// CmdPipelineBarrier wraps vkCmdPipelineBarrier.
func CmdPipelineBarrier(commandBuffer CommandBuffer, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, dependencyFlags DependencyFlags, memoryBarrierCount uint32, pMemoryBarriers MemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers BufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers ImageMemoryBarrier) {
	C.vkCmdPipelineBarrier(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlags(srcStageMask), C.VkPipelineStageFlags(dstStageMask), C.VkDependencyFlags(dependencyFlags), C.uint32_t(memoryBarrierCount), pMemoryBarriers.ptr, C.uint32_t(bufferMemoryBarrierCount), pBufferMemoryBarriers.ptr, C.uint32_t(imageMemoryBarrierCount), pImageMemoryBarriers.ptr)
}

// CmdPipelineBarrier2 wraps vkCmdPipelineBarrier2.
func CmdPipelineBarrier2(commandBuffer CommandBuffer, pDependencyInfo DependencyInfo) {
	C.vkCmdPipelineBarrier2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pDependencyInfo.ptr)
}

// CmdPushConstants wraps vkCmdPushConstants.
func CmdPushConstants(commandBuffer CommandBuffer, layout PipelineLayout, stageFlags ShaderStageFlags, offset uint32, size uint32, pValues unsafe.Pointer) {
	C.vkCmdPushConstants(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineLayout(layout), C.VkShaderStageFlags(stageFlags), C.uint32_t(offset), C.uint32_t(size), pValues)
}

// CmdPushConstants2 wraps vkCmdPushConstants2.
func CmdPushConstants2(commandBuffer CommandBuffer, pPushConstantsInfo PushConstantsInfo) {
	C.vkCmdPushConstants2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pPushConstantsInfo.ptr)
}

// CmdPushDescriptorSet wraps vkCmdPushDescriptorSet.
func CmdPushDescriptorSet(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, set uint32, descriptorWriteCount uint32, pDescriptorWrites WriteDescriptorSet) {
	C.vkCmdPushDescriptorSet(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineBindPoint(pipelineBindPoint), C.VkPipelineLayout(layout), C.uint32_t(set), C.uint32_t(descriptorWriteCount), pDescriptorWrites.ptr)
}

// CmdPushDescriptorSet2 wraps vkCmdPushDescriptorSet2.
func CmdPushDescriptorSet2(commandBuffer CommandBuffer, pPushDescriptorSetInfo PushDescriptorSetInfo) {
	C.vkCmdPushDescriptorSet2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pPushDescriptorSetInfo.ptr)
}

// CmdPushDescriptorSetWithTemplate wraps vkCmdPushDescriptorSetWithTemplate.
func CmdPushDescriptorSetWithTemplate(commandBuffer CommandBuffer, descriptorUpdateTemplate DescriptorUpdateTemplate, layout PipelineLayout, set uint32, pData unsafe.Pointer) {
	C.vkCmdPushDescriptorSetWithTemplate(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate), C.VkPipelineLayout(layout), C.uint32_t(set), pData)
}

// CmdPushDescriptorSetWithTemplate2 wraps vkCmdPushDescriptorSetWithTemplate2.
func CmdPushDescriptorSetWithTemplate2(commandBuffer CommandBuffer, pPushDescriptorSetWithTemplateInfo PushDescriptorSetWithTemplateInfo) {
	C.vkCmdPushDescriptorSetWithTemplate2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pPushDescriptorSetWithTemplateInfo.ptr)
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
	C.vkCmdResolveImage(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkImage(srcImage), C.VkImageLayout(srcImageLayout), C.VkImage(dstImage), C.VkImageLayout(dstImageLayout), C.uint32_t(regionCount), pRegions.ptr)
}

// CmdResolveImage2 wraps vkCmdResolveImage2.
func CmdResolveImage2(commandBuffer CommandBuffer, pResolveImageInfo ResolveImageInfo2) {
	C.vkCmdResolveImage2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pResolveImageInfo.ptr)
}

// vkCmdSetBlendConstants.blendConstants is unsupported: category unsupported.

// CmdSetCullMode wraps vkCmdSetCullMode.
func CmdSetCullMode(commandBuffer CommandBuffer, cullMode CullModeFlags) {
	C.vkCmdSetCullMode(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkCullModeFlags(cullMode))
}

// CmdSetDepthBias wraps vkCmdSetDepthBias.
func CmdSetDepthBias(commandBuffer CommandBuffer, depthBiasConstantFactor float32, depthBiasClamp float32, depthBiasSlopeFactor float32) {
	C.vkCmdSetDepthBias(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.float(depthBiasConstantFactor), C.float(depthBiasClamp), C.float(depthBiasSlopeFactor))
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

// CmdSetDeviceMask wraps vkCmdSetDeviceMask.
func CmdSetDeviceMask(commandBuffer CommandBuffer, deviceMask uint32) {
	C.vkCmdSetDeviceMask(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(deviceMask))
}

// CmdSetEvent wraps vkCmdSetEvent.
func CmdSetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	C.vkCmdSetEvent(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkEvent(event), C.VkPipelineStageFlags(stageMask))
}

// CmdSetEvent2 wraps vkCmdSetEvent2.
func CmdSetEvent2(commandBuffer CommandBuffer, event Event, pDependencyInfo DependencyInfo) {
	C.vkCmdSetEvent2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkEvent(event), pDependencyInfo.ptr)
}

// CmdSetFrontFace wraps vkCmdSetFrontFace.
func CmdSetFrontFace(commandBuffer CommandBuffer, frontFace FrontFace) {
	C.vkCmdSetFrontFace(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkFrontFace(frontFace))
}

// CmdSetLineStipple wraps vkCmdSetLineStipple.
func CmdSetLineStipple(commandBuffer CommandBuffer, lineStippleFactor uint32, lineStipplePattern uint16) {
	C.vkCmdSetLineStipple(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(lineStippleFactor), C.uint16_t(lineStipplePattern))
}

// CmdSetLineWidth wraps vkCmdSetLineWidth.
func CmdSetLineWidth(commandBuffer CommandBuffer, lineWidth float32) {
	C.vkCmdSetLineWidth(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.float(lineWidth))
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

// CmdSetRasterizerDiscardEnable wraps vkCmdSetRasterizerDiscardEnable.
func CmdSetRasterizerDiscardEnable(commandBuffer CommandBuffer, rasterizerDiscardEnable bool) {
	tmp_rasterizerDiscardEnable := 0

	if rasterizerDiscardEnable {
		tmp_rasterizerDiscardEnable = 1
	}

	C.vkCmdSetRasterizerDiscardEnable(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBool32(tmp_rasterizerDiscardEnable))
}

// CmdSetRenderingAttachmentLocations wraps vkCmdSetRenderingAttachmentLocations.
func CmdSetRenderingAttachmentLocations(commandBuffer CommandBuffer, pLocationInfo RenderingAttachmentLocationInfo) {
	C.vkCmdSetRenderingAttachmentLocations(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pLocationInfo.ptr)
}

// CmdSetRenderingInputAttachmentIndices wraps vkCmdSetRenderingInputAttachmentIndices.
func CmdSetRenderingInputAttachmentIndices(commandBuffer CommandBuffer, pInputAttachmentIndexInfo RenderingInputAttachmentIndexInfo) {
	C.vkCmdSetRenderingInputAttachmentIndices(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), pInputAttachmentIndexInfo.ptr)
}

// CmdSetScissor wraps vkCmdSetScissor.
func CmdSetScissor(commandBuffer CommandBuffer, firstScissor uint32, scissorCount uint32, pScissors Rect2D) {
	C.vkCmdSetScissor(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstScissor), C.uint32_t(scissorCount), pScissors.ptr)
}

// CmdSetScissorWithCount wraps vkCmdSetScissorWithCount.
func CmdSetScissorWithCount(commandBuffer CommandBuffer, scissorCount uint32, pScissors Rect2D) {
	C.vkCmdSetScissorWithCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(scissorCount), pScissors.ptr)
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

// CmdSetViewport wraps vkCmdSetViewport.
func CmdSetViewport(commandBuffer CommandBuffer, firstViewport uint32, viewportCount uint32, pViewports Viewport) {
	C.vkCmdSetViewport(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(firstViewport), C.uint32_t(viewportCount), pViewports.ptr)
}

// CmdSetViewportWithCount wraps vkCmdSetViewportWithCount.
func CmdSetViewportWithCount(commandBuffer CommandBuffer, viewportCount uint32, pViewports Viewport) {
	C.vkCmdSetViewportWithCount(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(viewportCount), pViewports.ptr)
}

// CmdUpdateBuffer wraps vkCmdUpdateBuffer.
func CmdUpdateBuffer(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset DeviceSize, dataSize DeviceSize, pData unsafe.Pointer) {
	C.vkCmdUpdateBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkBuffer(dstBuffer), C.VkDeviceSize(dstOffset), C.VkDeviceSize(dataSize), pData)
}

// CmdWaitEvents wraps vkCmdWaitEvents.
func CmdWaitEvents(commandBuffer CommandBuffer, eventCount uint32, pEvents ffi.Ref[Event], srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarrierCount uint32, pMemoryBarriers MemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers BufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers ImageMemoryBarrier) {
	C.vkCmdWaitEvents(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(eventCount), (*C.VkEvent)(pEvents.Raw()), C.VkPipelineStageFlags(srcStageMask), C.VkPipelineStageFlags(dstStageMask), C.uint32_t(memoryBarrierCount), pMemoryBarriers.ptr, C.uint32_t(bufferMemoryBarrierCount), pBufferMemoryBarriers.ptr, C.uint32_t(imageMemoryBarrierCount), pImageMemoryBarriers.ptr)
}

// CmdWaitEvents2 wraps vkCmdWaitEvents2.
func CmdWaitEvents2(commandBuffer CommandBuffer, eventCount uint32, pEvents ffi.Ref[Event], pDependencyInfos DependencyInfo) {
	C.vkCmdWaitEvents2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.uint32_t(eventCount), (*C.VkEvent)(pEvents.Raw()), pDependencyInfos.ptr)
}

// CmdWriteTimestamp wraps vkCmdWriteTimestamp.
func CmdWriteTimestamp(commandBuffer CommandBuffer, pipelineStage PipelineStageFlags, queryPool QueryPool, query uint32) {
	C.vkCmdWriteTimestamp(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlagBits(pipelineStage), C.VkQueryPool(queryPool), C.uint32_t(query))
}

// CmdWriteTimestamp2 wraps vkCmdWriteTimestamp2.
func CmdWriteTimestamp2(commandBuffer CommandBuffer, stage PipelineStageFlags2, queryPool QueryPool, query uint32) {
	C.vkCmdWriteTimestamp2(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)), C.VkPipelineStageFlags2(stage), C.VkQueryPool(queryPool), C.uint32_t(query))
}

// CopyImageToImage wraps vkCopyImageToImage.
func CopyImageToImage(device Device, pCopyImageToImageInfo CopyImageToImageInfo) Result {
	ret := C.vkCopyImageToImage(C.VkDevice(unsafe.Pointer(device)), pCopyImageToImageInfo.ptr)

	return Result(ret)
}

// CopyImageToMemory wraps vkCopyImageToMemory.
func CopyImageToMemory(device Device, pCopyImageToMemoryInfo CopyImageToMemoryInfo) Result {
	ret := C.vkCopyImageToMemory(C.VkDevice(unsafe.Pointer(device)), pCopyImageToMemoryInfo.ptr)

	return Result(ret)
}

// CopyMemoryToImage wraps vkCopyMemoryToImage.
func CopyMemoryToImage(device Device, pCopyMemoryToImageInfo CopyMemoryToImageInfo) Result {
	ret := C.vkCopyMemoryToImage(C.VkDevice(unsafe.Pointer(device)), pCopyMemoryToImageInfo.ptr)

	return Result(ret)
}

// CreateBuffer wraps vkCreateBuffer.
func CreateBuffer(device Device, pCreateInfo BufferCreateInfo, pAllocator AllocationCallbacks, pBuffer ffi.Ref[Buffer]) Result {
	ret := C.vkCreateBuffer(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkBuffer)(pBuffer.Raw()))

	return Result(ret)
}

// CreateBufferView wraps vkCreateBufferView.
func CreateBufferView(device Device, pCreateInfo BufferViewCreateInfo, pAllocator AllocationCallbacks, pView ffi.Ref[BufferView]) Result {
	ret := C.vkCreateBufferView(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkBufferView)(pView.Raw()))

	return Result(ret)
}

// CreateCommandPool wraps vkCreateCommandPool.
func CreateCommandPool(device Device, pCreateInfo CommandPoolCreateInfo, pAllocator AllocationCallbacks, pCommandPool ffi.Ref[CommandPool]) Result {
	ret := C.vkCreateCommandPool(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkCommandPool)(pCommandPool.Raw()))

	return Result(ret)
}

// CreateComputePipelines wraps vkCreateComputePipelines.
func CreateComputePipelines(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos ComputePipelineCreateInfo, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateComputePipelines(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), pCreateInfos.ptr, pAllocator.ptr, (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateDescriptorPool wraps vkCreateDescriptorPool.
func CreateDescriptorPool(device Device, pCreateInfo DescriptorPoolCreateInfo, pAllocator AllocationCallbacks, pDescriptorPool ffi.Ref[DescriptorPool]) Result {
	ret := C.vkCreateDescriptorPool(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkDescriptorPool)(pDescriptorPool.Raw()))

	return Result(ret)
}

// CreateDescriptorSetLayout wraps vkCreateDescriptorSetLayout.
func CreateDescriptorSetLayout(device Device, pCreateInfo DescriptorSetLayoutCreateInfo, pAllocator AllocationCallbacks, pSetLayout ffi.Ref[DescriptorSetLayout]) Result {
	ret := C.vkCreateDescriptorSetLayout(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkDescriptorSetLayout)(pSetLayout.Raw()))

	return Result(ret)
}

// CreateDescriptorUpdateTemplate wraps vkCreateDescriptorUpdateTemplate.
func CreateDescriptorUpdateTemplate(device Device, pCreateInfo DescriptorUpdateTemplateCreateInfo, pAllocator AllocationCallbacks, pDescriptorUpdateTemplate ffi.Ref[DescriptorUpdateTemplate]) Result {
	ret := C.vkCreateDescriptorUpdateTemplate(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkDescriptorUpdateTemplate)(pDescriptorUpdateTemplate.Raw()))

	return Result(ret)
}

// CreateDevice wraps vkCreateDevice.
func CreateDevice(physicalDevice PhysicalDevice, pCreateInfo DeviceCreateInfo, pAllocator AllocationCallbacks, pDevice ffi.Ref[Device]) Result {
	ret := C.vkCreateDevice(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkDevice)(pDevice.Raw()))

	return Result(ret)
}

// CreateEvent wraps vkCreateEvent.
func CreateEvent(device Device, pCreateInfo EventCreateInfo, pAllocator AllocationCallbacks, pEvent ffi.Ref[Event]) Result {
	ret := C.vkCreateEvent(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkEvent)(pEvent.Raw()))

	return Result(ret)
}

// CreateFence wraps vkCreateFence.
func CreateFence(device Device, pCreateInfo FenceCreateInfo, pAllocator AllocationCallbacks, pFence ffi.Ref[Fence]) Result {
	ret := C.vkCreateFence(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkFence)(pFence.Raw()))

	return Result(ret)
}

// CreateFramebuffer wraps vkCreateFramebuffer.
func CreateFramebuffer(device Device, pCreateInfo FramebufferCreateInfo, pAllocator AllocationCallbacks, pFramebuffer ffi.Ref[Framebuffer]) Result {
	ret := C.vkCreateFramebuffer(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkFramebuffer)(pFramebuffer.Raw()))

	return Result(ret)
}

// CreateGraphicsPipelines wraps vkCreateGraphicsPipelines.
func CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos GraphicsPipelineCreateInfo, pAllocator AllocationCallbacks, pPipelines ffi.Ref[Pipeline]) Result {
	ret := C.vkCreateGraphicsPipelines(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), C.uint32_t(createInfoCount), pCreateInfos.ptr, pAllocator.ptr, (*C.VkPipeline)(pPipelines.Raw()))

	return Result(ret)
}

// CreateImage wraps vkCreateImage.
func CreateImage(device Device, pCreateInfo ImageCreateInfo, pAllocator AllocationCallbacks, pImage ffi.Ref[Image]) Result {
	ret := C.vkCreateImage(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkImage)(pImage.Raw()))

	return Result(ret)
}

// CreateImageView wraps vkCreateImageView.
func CreateImageView(device Device, pCreateInfo ImageViewCreateInfo, pAllocator AllocationCallbacks, pView ffi.Ref[ImageView]) Result {
	ret := C.vkCreateImageView(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkImageView)(pView.Raw()))

	return Result(ret)
}

// CreateInstance wraps vkCreateInstance.
func CreateInstance(pCreateInfo InstanceCreateInfo, pAllocator AllocationCallbacks, pInstance ffi.Ref[Instance]) Result {
	ret := C.vkCreateInstance(pCreateInfo.ptr, pAllocator.ptr, (*C.VkInstance)(pInstance.Raw()))

	return Result(ret)
}

// CreatePipelineCache wraps vkCreatePipelineCache.
func CreatePipelineCache(device Device, pCreateInfo PipelineCacheCreateInfo, pAllocator AllocationCallbacks, pPipelineCache ffi.Ref[PipelineCache]) Result {
	ret := C.vkCreatePipelineCache(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkPipelineCache)(pPipelineCache.Raw()))

	return Result(ret)
}

// CreatePipelineLayout wraps vkCreatePipelineLayout.
func CreatePipelineLayout(device Device, pCreateInfo PipelineLayoutCreateInfo, pAllocator AllocationCallbacks, pPipelineLayout ffi.Ref[PipelineLayout]) Result {
	ret := C.vkCreatePipelineLayout(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkPipelineLayout)(pPipelineLayout.Raw()))

	return Result(ret)
}

// CreatePrivateDataSlot wraps vkCreatePrivateDataSlot.
func CreatePrivateDataSlot(device Device, pCreateInfo PrivateDataSlotCreateInfo, pAllocator AllocationCallbacks, pPrivateDataSlot ffi.Ref[PrivateDataSlot]) Result {
	ret := C.vkCreatePrivateDataSlot(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkPrivateDataSlot)(pPrivateDataSlot.Raw()))

	return Result(ret)
}

// CreateQueryPool wraps vkCreateQueryPool.
func CreateQueryPool(device Device, pCreateInfo QueryPoolCreateInfo, pAllocator AllocationCallbacks, pQueryPool ffi.Ref[QueryPool]) Result {
	ret := C.vkCreateQueryPool(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkQueryPool)(pQueryPool.Raw()))

	return Result(ret)
}

// CreateRenderPass wraps vkCreateRenderPass.
func CreateRenderPass(device Device, pCreateInfo RenderPassCreateInfo, pAllocator AllocationCallbacks, pRenderPass ffi.Ref[RenderPass]) Result {
	ret := C.vkCreateRenderPass(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkRenderPass)(pRenderPass.Raw()))

	return Result(ret)
}

// CreateRenderPass2 wraps vkCreateRenderPass2.
func CreateRenderPass2(device Device, pCreateInfo RenderPassCreateInfo2, pAllocator AllocationCallbacks, pRenderPass ffi.Ref[RenderPass]) Result {
	ret := C.vkCreateRenderPass2(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkRenderPass)(pRenderPass.Raw()))

	return Result(ret)
}

// CreateSampler wraps vkCreateSampler.
func CreateSampler(device Device, pCreateInfo SamplerCreateInfo, pAllocator AllocationCallbacks, pSampler ffi.Ref[Sampler]) Result {
	ret := C.vkCreateSampler(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkSampler)(pSampler.Raw()))

	return Result(ret)
}

// CreateSamplerYcbcrConversion wraps vkCreateSamplerYcbcrConversion.
func CreateSamplerYcbcrConversion(device Device, pCreateInfo SamplerYcbcrConversionCreateInfo, pAllocator AllocationCallbacks, pYcbcrConversion ffi.Ref[SamplerYcbcrConversion]) Result {
	ret := C.vkCreateSamplerYcbcrConversion(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkSamplerYcbcrConversion)(pYcbcrConversion.Raw()))

	return Result(ret)
}

// CreateSemaphore wraps vkCreateSemaphore.
func CreateSemaphore(device Device, pCreateInfo SemaphoreCreateInfo, pAllocator AllocationCallbacks, pSemaphore ffi.Ref[Semaphore]) Result {
	ret := C.vkCreateSemaphore(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkSemaphore)(pSemaphore.Raw()))

	return Result(ret)
}

// CreateShaderModule wraps vkCreateShaderModule.
func CreateShaderModule(device Device, pCreateInfo ShaderModuleCreateInfo, pAllocator AllocationCallbacks, pShaderModule ffi.Ref[ShaderModule]) Result {
	ret := C.vkCreateShaderModule(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pAllocator.ptr, (*C.VkShaderModule)(pShaderModule.Raw()))

	return Result(ret)
}

// DestroyBuffer wraps vkDestroyBuffer.
func DestroyBuffer(device Device, buffer Buffer, pAllocator AllocationCallbacks) {
	C.vkDestroyBuffer(C.VkDevice(unsafe.Pointer(device)), C.VkBuffer(buffer), pAllocator.ptr)
}

// DestroyBufferView wraps vkDestroyBufferView.
func DestroyBufferView(device Device, bufferView BufferView, pAllocator AllocationCallbacks) {
	C.vkDestroyBufferView(C.VkDevice(unsafe.Pointer(device)), C.VkBufferView(bufferView), pAllocator.ptr)
}

// DestroyCommandPool wraps vkDestroyCommandPool.
func DestroyCommandPool(device Device, commandPool CommandPool, pAllocator AllocationCallbacks) {
	C.vkDestroyCommandPool(C.VkDevice(unsafe.Pointer(device)), C.VkCommandPool(commandPool), pAllocator.ptr)
}

// DestroyDescriptorPool wraps vkDestroyDescriptorPool.
func DestroyDescriptorPool(device Device, descriptorPool DescriptorPool, pAllocator AllocationCallbacks) {
	C.vkDestroyDescriptorPool(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorPool(descriptorPool), pAllocator.ptr)
}

// DestroyDescriptorSetLayout wraps vkDestroyDescriptorSetLayout.
func DestroyDescriptorSetLayout(device Device, descriptorSetLayout DescriptorSetLayout, pAllocator AllocationCallbacks) {
	C.vkDestroyDescriptorSetLayout(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSetLayout(descriptorSetLayout), pAllocator.ptr)
}

// DestroyDescriptorUpdateTemplate wraps vkDestroyDescriptorUpdateTemplate.
func DestroyDescriptorUpdateTemplate(device Device, descriptorUpdateTemplate DescriptorUpdateTemplate, pAllocator AllocationCallbacks) {
	C.vkDestroyDescriptorUpdateTemplate(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate), pAllocator.ptr)
}

// DestroyDevice wraps vkDestroyDevice.
func DestroyDevice(device Device, pAllocator AllocationCallbacks) {
	C.vkDestroyDevice(C.VkDevice(unsafe.Pointer(device)), pAllocator.ptr)
}

// DestroyEvent wraps vkDestroyEvent.
func DestroyEvent(device Device, event Event, pAllocator AllocationCallbacks) {
	C.vkDestroyEvent(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event), pAllocator.ptr)
}

// DestroyFence wraps vkDestroyFence.
func DestroyFence(device Device, fence Fence, pAllocator AllocationCallbacks) {
	C.vkDestroyFence(C.VkDevice(unsafe.Pointer(device)), C.VkFence(fence), pAllocator.ptr)
}

// DestroyFramebuffer wraps vkDestroyFramebuffer.
func DestroyFramebuffer(device Device, framebuffer Framebuffer, pAllocator AllocationCallbacks) {
	C.vkDestroyFramebuffer(C.VkDevice(unsafe.Pointer(device)), C.VkFramebuffer(framebuffer), pAllocator.ptr)
}

// DestroyImage wraps vkDestroyImage.
func DestroyImage(device Device, image Image, pAllocator AllocationCallbacks) {
	C.vkDestroyImage(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), pAllocator.ptr)
}

// DestroyImageView wraps vkDestroyImageView.
func DestroyImageView(device Device, imageView ImageView, pAllocator AllocationCallbacks) {
	C.vkDestroyImageView(C.VkDevice(unsafe.Pointer(device)), C.VkImageView(imageView), pAllocator.ptr)
}

// DestroyInstance wraps vkDestroyInstance.
func DestroyInstance(instance Instance, pAllocator AllocationCallbacks) {
	C.vkDestroyInstance(C.VkInstance(unsafe.Pointer(instance)), pAllocator.ptr)
}

// DestroyPipeline wraps vkDestroyPipeline.
func DestroyPipeline(device Device, pipeline Pipeline, pAllocator AllocationCallbacks) {
	C.vkDestroyPipeline(C.VkDevice(unsafe.Pointer(device)), C.VkPipeline(pipeline), pAllocator.ptr)
}

// DestroyPipelineCache wraps vkDestroyPipelineCache.
func DestroyPipelineCache(device Device, pipelineCache PipelineCache, pAllocator AllocationCallbacks) {
	C.vkDestroyPipelineCache(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), pAllocator.ptr)
}

// DestroyPipelineLayout wraps vkDestroyPipelineLayout.
func DestroyPipelineLayout(device Device, pipelineLayout PipelineLayout, pAllocator AllocationCallbacks) {
	C.vkDestroyPipelineLayout(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineLayout(pipelineLayout), pAllocator.ptr)
}

// DestroyPrivateDataSlot wraps vkDestroyPrivateDataSlot.
func DestroyPrivateDataSlot(device Device, privateDataSlot PrivateDataSlot, pAllocator AllocationCallbacks) {
	C.vkDestroyPrivateDataSlot(C.VkDevice(unsafe.Pointer(device)), C.VkPrivateDataSlot(privateDataSlot), pAllocator.ptr)
}

// DestroyQueryPool wraps vkDestroyQueryPool.
func DestroyQueryPool(device Device, queryPool QueryPool, pAllocator AllocationCallbacks) {
	C.vkDestroyQueryPool(C.VkDevice(unsafe.Pointer(device)), C.VkQueryPool(queryPool), pAllocator.ptr)
}

// DestroyRenderPass wraps vkDestroyRenderPass.
func DestroyRenderPass(device Device, renderPass RenderPass, pAllocator AllocationCallbacks) {
	C.vkDestroyRenderPass(C.VkDevice(unsafe.Pointer(device)), C.VkRenderPass(renderPass), pAllocator.ptr)
}

// DestroySampler wraps vkDestroySampler.
func DestroySampler(device Device, sampler Sampler, pAllocator AllocationCallbacks) {
	C.vkDestroySampler(C.VkDevice(unsafe.Pointer(device)), C.VkSampler(sampler), pAllocator.ptr)
}

// DestroySamplerYcbcrConversion wraps vkDestroySamplerYcbcrConversion.
func DestroySamplerYcbcrConversion(device Device, ycbcrConversion SamplerYcbcrConversion, pAllocator AllocationCallbacks) {
	C.vkDestroySamplerYcbcrConversion(C.VkDevice(unsafe.Pointer(device)), C.VkSamplerYcbcrConversion(ycbcrConversion), pAllocator.ptr)
}

// DestroySemaphore wraps vkDestroySemaphore.
func DestroySemaphore(device Device, semaphore Semaphore, pAllocator AllocationCallbacks) {
	C.vkDestroySemaphore(C.VkDevice(unsafe.Pointer(device)), C.VkSemaphore(semaphore), pAllocator.ptr)
}

// DestroyShaderModule wraps vkDestroyShaderModule.
func DestroyShaderModule(device Device, shaderModule ShaderModule, pAllocator AllocationCallbacks) {
	C.vkDestroyShaderModule(C.VkDevice(unsafe.Pointer(device)), C.VkShaderModule(shaderModule), pAllocator.ptr)
}

// DeviceWaitIdle wraps vkDeviceWaitIdle.
func DeviceWaitIdle(device Device) Result {
	ret := C.vkDeviceWaitIdle(C.VkDevice(unsafe.Pointer(device)))

	return Result(ret)
}

// EndCommandBuffer wraps vkEndCommandBuffer.
func EndCommandBuffer(commandBuffer CommandBuffer) Result {
	ret := C.vkEndCommandBuffer(C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))

	return Result(ret)
}

// EnumerateDeviceExtensionProperties wraps vkEnumerateDeviceExtensionProperties.
func EnumerateDeviceExtensionProperties(physicalDevice PhysicalDevice, pLayerName ffi.CString, pPropertyCount ffi.Ref[uint32], pProperties ExtensionProperties) Result {
	ret := C.vkEnumerateDeviceExtensionProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.char)(pLayerName.Raw()), (*C.uint32_t)(pPropertyCount.Raw()), pProperties.ptr)

	return Result(ret)
}

// EnumerateDeviceLayerProperties wraps vkEnumerateDeviceLayerProperties.
func EnumerateDeviceLayerProperties(physicalDevice PhysicalDevice, pPropertyCount ffi.Ref[uint32], pProperties LayerProperties) Result {
	ret := C.vkEnumerateDeviceLayerProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pPropertyCount.Raw()), pProperties.ptr)

	return Result(ret)
}

// EnumerateInstanceExtensionProperties wraps vkEnumerateInstanceExtensionProperties.
func EnumerateInstanceExtensionProperties(pLayerName ffi.CString, pPropertyCount ffi.Ref[uint32], pProperties ExtensionProperties) Result {
	ret := C.vkEnumerateInstanceExtensionProperties((*C.char)(pLayerName.Raw()), (*C.uint32_t)(pPropertyCount.Raw()), pProperties.ptr)

	return Result(ret)
}

// EnumerateInstanceLayerProperties wraps vkEnumerateInstanceLayerProperties.
func EnumerateInstanceLayerProperties(pPropertyCount ffi.Ref[uint32], pProperties LayerProperties) Result {
	ret := C.vkEnumerateInstanceLayerProperties((*C.uint32_t)(pPropertyCount.Raw()), pProperties.ptr)

	return Result(ret)
}

// EnumerateInstanceVersion wraps vkEnumerateInstanceVersion.
func EnumerateInstanceVersion(pApiVersion ffi.Ref[uint32]) Result {
	ret := C.vkEnumerateInstanceVersion((*C.uint32_t)(pApiVersion.Raw()))

	return Result(ret)
}

// EnumeratePhysicalDeviceGroups wraps vkEnumeratePhysicalDeviceGroups.
func EnumeratePhysicalDeviceGroups(instance Instance, pPhysicalDeviceGroupCount ffi.Ref[uint32], pPhysicalDeviceGroupProperties PhysicalDeviceGroupProperties) Result {
	ret := C.vkEnumeratePhysicalDeviceGroups(C.VkInstance(unsafe.Pointer(instance)), (*C.uint32_t)(pPhysicalDeviceGroupCount.Raw()), pPhysicalDeviceGroupProperties.ptr)

	return Result(ret)
}

// EnumeratePhysicalDevices wraps vkEnumeratePhysicalDevices.
func EnumeratePhysicalDevices(instance Instance, pPhysicalDeviceCount ffi.Ref[uint32], pPhysicalDevices ffi.Ref[PhysicalDevice]) Result {
	ret := C.vkEnumeratePhysicalDevices(C.VkInstance(unsafe.Pointer(instance)), (*C.uint32_t)(pPhysicalDeviceCount.Raw()), (*C.VkPhysicalDevice)(pPhysicalDevices.Raw()))

	return Result(ret)
}

// FlushMappedMemoryRanges wraps vkFlushMappedMemoryRanges.
func FlushMappedMemoryRanges(device Device, memoryRangeCount uint32, pMemoryRanges MappedMemoryRange) Result {
	ret := C.vkFlushMappedMemoryRanges(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(memoryRangeCount), pMemoryRanges.ptr)

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
	C.vkFreeMemory(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory), pAllocator.ptr)
}

// GetBufferDeviceAddress wraps vkGetBufferDeviceAddress.
func GetBufferDeviceAddress(device Device, pInfo BufferDeviceAddressInfo) DeviceAddress {
	ret := C.vkGetBufferDeviceAddress(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr)

	return DeviceAddress(ret)
}

// GetBufferMemoryRequirements wraps vkGetBufferMemoryRequirements.
func GetBufferMemoryRequirements(device Device, buffer Buffer, pMemoryRequirements MemoryRequirements) {
	C.vkGetBufferMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), C.VkBuffer(buffer), pMemoryRequirements.ptr)
}

// GetBufferMemoryRequirements2 wraps vkGetBufferMemoryRequirements2.
func GetBufferMemoryRequirements2(device Device, pInfo BufferMemoryRequirementsInfo2, pMemoryRequirements MemoryRequirements2) {
	C.vkGetBufferMemoryRequirements2(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, pMemoryRequirements.ptr)
}

// GetBufferOpaqueCaptureAddress wraps vkGetBufferOpaqueCaptureAddress.
func GetBufferOpaqueCaptureAddress(device Device, pInfo BufferDeviceAddressInfo) uint64 {
	ret := C.vkGetBufferOpaqueCaptureAddress(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr)

	return uint64(ret)
}

// GetDescriptorSetLayoutSupport wraps vkGetDescriptorSetLayoutSupport.
func GetDescriptorSetLayoutSupport(device Device, pCreateInfo DescriptorSetLayoutCreateInfo, pSupport DescriptorSetLayoutSupport) {
	C.vkGetDescriptorSetLayoutSupport(C.VkDevice(unsafe.Pointer(device)), pCreateInfo.ptr, pSupport.ptr)
}

// GetDeviceBufferMemoryRequirements wraps vkGetDeviceBufferMemoryRequirements.
func GetDeviceBufferMemoryRequirements(device Device, pInfo DeviceBufferMemoryRequirements, pMemoryRequirements MemoryRequirements2) {
	C.vkGetDeviceBufferMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, pMemoryRequirements.ptr)
}

// GetDeviceGroupPeerMemoryFeatures wraps vkGetDeviceGroupPeerMemoryFeatures.
func GetDeviceGroupPeerMemoryFeatures(device Device, heapIndex uint32, localDeviceIndex uint32, remoteDeviceIndex uint32, pPeerMemoryFeatures ffi.Ref[PeerMemoryFeatureFlags]) {
	C.vkGetDeviceGroupPeerMemoryFeatures(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(heapIndex), C.uint32_t(localDeviceIndex), C.uint32_t(remoteDeviceIndex), (*C.VkPeerMemoryFeatureFlags)(pPeerMemoryFeatures.Raw()))
}

// GetDeviceImageMemoryRequirements wraps vkGetDeviceImageMemoryRequirements.
func GetDeviceImageMemoryRequirements(device Device, pInfo DeviceImageMemoryRequirements, pMemoryRequirements MemoryRequirements2) {
	C.vkGetDeviceImageMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, pMemoryRequirements.ptr)
}

// GetDeviceImageSparseMemoryRequirements wraps vkGetDeviceImageSparseMemoryRequirements.
func GetDeviceImageSparseMemoryRequirements(device Device, pInfo DeviceImageMemoryRequirements, pSparseMemoryRequirementCount ffi.Ref[uint32], pSparseMemoryRequirements SparseImageMemoryRequirements2) {
	C.vkGetDeviceImageSparseMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, (*C.uint32_t)(pSparseMemoryRequirementCount.Raw()), pSparseMemoryRequirements.ptr)
}

// GetDeviceImageSubresourceLayout wraps vkGetDeviceImageSubresourceLayout.
func GetDeviceImageSubresourceLayout(device Device, pInfo DeviceImageSubresourceInfo, pLayout SubresourceLayout2) {
	C.vkGetDeviceImageSubresourceLayout(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, pLayout.ptr)
}

// GetDeviceMemoryCommitment wraps vkGetDeviceMemoryCommitment.
func GetDeviceMemoryCommitment(device Device, memory DeviceMemory, pCommittedMemoryInBytes ffi.Ref[DeviceSize]) {
	C.vkGetDeviceMemoryCommitment(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory), (*C.VkDeviceSize)(pCommittedMemoryInBytes.Raw()))
}

// GetDeviceMemoryOpaqueCaptureAddress wraps vkGetDeviceMemoryOpaqueCaptureAddress.
func GetDeviceMemoryOpaqueCaptureAddress(device Device, pInfo DeviceMemoryOpaqueCaptureAddressInfo) uint64 {
	ret := C.vkGetDeviceMemoryOpaqueCaptureAddress(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr)

	return uint64(ret)
}

// vkGetDeviceProcAddr is unsupported: unknown type PFN_vkVoidFunction.

// GetDeviceQueue wraps vkGetDeviceQueue.
func GetDeviceQueue(device Device, queueFamilyIndex uint32, queueIndex uint32, pQueue ffi.Ref[Queue]) {
	C.vkGetDeviceQueue(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), (*C.VkQueue)(pQueue.Raw()))
}

// GetDeviceQueue2 wraps vkGetDeviceQueue2.
func GetDeviceQueue2(device Device, pQueueInfo DeviceQueueInfo2, pQueue ffi.Ref[Queue]) {
	C.vkGetDeviceQueue2(C.VkDevice(unsafe.Pointer(device)), pQueueInfo.ptr, (*C.VkQueue)(pQueue.Raw()))
}

// GetEventStatus wraps vkGetEventStatus.
func GetEventStatus(device Device, event Event) Result {
	ret := C.vkGetEventStatus(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event))

	return Result(ret)
}

// GetFenceStatus wraps vkGetFenceStatus.
func GetFenceStatus(device Device, fence Fence) Result {
	ret := C.vkGetFenceStatus(C.VkDevice(unsafe.Pointer(device)), C.VkFence(fence))

	return Result(ret)
}

// GetImageMemoryRequirements wraps vkGetImageMemoryRequirements.
func GetImageMemoryRequirements(device Device, image Image, pMemoryRequirements MemoryRequirements) {
	C.vkGetImageMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), pMemoryRequirements.ptr)
}

// GetImageMemoryRequirements2 wraps vkGetImageMemoryRequirements2.
func GetImageMemoryRequirements2(device Device, pInfo ImageMemoryRequirementsInfo2, pMemoryRequirements MemoryRequirements2) {
	C.vkGetImageMemoryRequirements2(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, pMemoryRequirements.ptr)
}

// GetImageSparseMemoryRequirements wraps vkGetImageSparseMemoryRequirements.
func GetImageSparseMemoryRequirements(device Device, image Image, pSparseMemoryRequirementCount ffi.Ref[uint32], pSparseMemoryRequirements SparseImageMemoryRequirements) {
	C.vkGetImageSparseMemoryRequirements(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), (*C.uint32_t)(pSparseMemoryRequirementCount.Raw()), pSparseMemoryRequirements.ptr)
}

// GetImageSparseMemoryRequirements2 wraps vkGetImageSparseMemoryRequirements2.
func GetImageSparseMemoryRequirements2(device Device, pInfo ImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount ffi.Ref[uint32], pSparseMemoryRequirements SparseImageMemoryRequirements2) {
	C.vkGetImageSparseMemoryRequirements2(C.VkDevice(unsafe.Pointer(device)), pInfo.ptr, (*C.uint32_t)(pSparseMemoryRequirementCount.Raw()), pSparseMemoryRequirements.ptr)
}

// GetImageSubresourceLayout wraps vkGetImageSubresourceLayout.
func GetImageSubresourceLayout(device Device, image Image, pSubresource ImageSubresource, pLayout SubresourceLayout) {
	C.vkGetImageSubresourceLayout(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), pSubresource.ptr, pLayout.ptr)
}

// GetImageSubresourceLayout2 wraps vkGetImageSubresourceLayout2.
func GetImageSubresourceLayout2(device Device, image Image, pSubresource ImageSubresource2, pLayout SubresourceLayout2) {
	C.vkGetImageSubresourceLayout2(C.VkDevice(unsafe.Pointer(device)), C.VkImage(image), pSubresource.ptr, pLayout.ptr)
}

// vkGetInstanceProcAddr is unsupported: unknown type PFN_vkVoidFunction.

// GetPhysicalDeviceExternalBufferProperties wraps vkGetPhysicalDeviceExternalBufferProperties.
func GetPhysicalDeviceExternalBufferProperties(physicalDevice PhysicalDevice, pExternalBufferInfo PhysicalDeviceExternalBufferInfo, pExternalBufferProperties ExternalBufferProperties) {
	C.vkGetPhysicalDeviceExternalBufferProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pExternalBufferInfo.ptr, pExternalBufferProperties.ptr)
}

// GetPhysicalDeviceExternalFenceProperties wraps vkGetPhysicalDeviceExternalFenceProperties.
func GetPhysicalDeviceExternalFenceProperties(physicalDevice PhysicalDevice, pExternalFenceInfo PhysicalDeviceExternalFenceInfo, pExternalFenceProperties ExternalFenceProperties) {
	C.vkGetPhysicalDeviceExternalFenceProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pExternalFenceInfo.ptr, pExternalFenceProperties.ptr)
}

// GetPhysicalDeviceExternalSemaphoreProperties wraps vkGetPhysicalDeviceExternalSemaphoreProperties.
func GetPhysicalDeviceExternalSemaphoreProperties(physicalDevice PhysicalDevice, pExternalSemaphoreInfo PhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties ExternalSemaphoreProperties) {
	C.vkGetPhysicalDeviceExternalSemaphoreProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pExternalSemaphoreInfo.ptr, pExternalSemaphoreProperties.ptr)
}

// GetPhysicalDeviceFeatures wraps vkGetPhysicalDeviceFeatures.
func GetPhysicalDeviceFeatures(physicalDevice PhysicalDevice, pFeatures PhysicalDeviceFeatures) {
	C.vkGetPhysicalDeviceFeatures(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pFeatures.ptr)
}

// GetPhysicalDeviceFeatures2 wraps vkGetPhysicalDeviceFeatures2.
func GetPhysicalDeviceFeatures2(physicalDevice PhysicalDevice, pFeatures PhysicalDeviceFeatures2) {
	C.vkGetPhysicalDeviceFeatures2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pFeatures.ptr)
}

// GetPhysicalDeviceFormatProperties wraps vkGetPhysicalDeviceFormatProperties.
func GetPhysicalDeviceFormatProperties(physicalDevice PhysicalDevice, format Format, pFormatProperties FormatProperties) {
	C.vkGetPhysicalDeviceFormatProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), pFormatProperties.ptr)
}

// GetPhysicalDeviceFormatProperties2 wraps vkGetPhysicalDeviceFormatProperties2.
func GetPhysicalDeviceFormatProperties2(physicalDevice PhysicalDevice, format Format, pFormatProperties FormatProperties2) {
	C.vkGetPhysicalDeviceFormatProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), pFormatProperties.ptr)
}

// GetPhysicalDeviceImageFormatProperties wraps vkGetPhysicalDeviceImageFormatProperties.
func GetPhysicalDeviceImageFormatProperties(physicalDevice PhysicalDevice, format Format, type_ ImageType, tiling ImageTiling, usage ImageUsageFlags, flags ImageCreateFlags, pImageFormatProperties ImageFormatProperties) Result {
	ret := C.vkGetPhysicalDeviceImageFormatProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), C.VkImageType(type_), C.VkImageTiling(tiling), C.VkImageUsageFlags(usage), C.VkImageCreateFlags(flags), pImageFormatProperties.ptr)

	return Result(ret)
}

// GetPhysicalDeviceImageFormatProperties2 wraps vkGetPhysicalDeviceImageFormatProperties2.
func GetPhysicalDeviceImageFormatProperties2(physicalDevice PhysicalDevice, pImageFormatInfo PhysicalDeviceImageFormatInfo2, pImageFormatProperties ImageFormatProperties2) Result {
	ret := C.vkGetPhysicalDeviceImageFormatProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pImageFormatInfo.ptr, pImageFormatProperties.ptr)

	return Result(ret)
}

// GetPhysicalDeviceMemoryProperties wraps vkGetPhysicalDeviceMemoryProperties.
func GetPhysicalDeviceMemoryProperties(physicalDevice PhysicalDevice, pMemoryProperties PhysicalDeviceMemoryProperties) {
	C.vkGetPhysicalDeviceMemoryProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pMemoryProperties.ptr)
}

// GetPhysicalDeviceMemoryProperties2 wraps vkGetPhysicalDeviceMemoryProperties2.
func GetPhysicalDeviceMemoryProperties2(physicalDevice PhysicalDevice, pMemoryProperties PhysicalDeviceMemoryProperties2) {
	C.vkGetPhysicalDeviceMemoryProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pMemoryProperties.ptr)
}

// GetPhysicalDeviceProperties wraps vkGetPhysicalDeviceProperties.
func GetPhysicalDeviceProperties(physicalDevice PhysicalDevice, pProperties PhysicalDeviceProperties) {
	C.vkGetPhysicalDeviceProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pProperties.ptr)
}

// GetPhysicalDeviceProperties2 wraps vkGetPhysicalDeviceProperties2.
func GetPhysicalDeviceProperties2(physicalDevice PhysicalDevice, pProperties PhysicalDeviceProperties2) {
	C.vkGetPhysicalDeviceProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pProperties.ptr)
}

// GetPhysicalDeviceQueueFamilyProperties wraps vkGetPhysicalDeviceQueueFamilyProperties.
func GetPhysicalDeviceQueueFamilyProperties(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount ffi.Ref[uint32], pQueueFamilyProperties QueueFamilyProperties) {
	C.vkGetPhysicalDeviceQueueFamilyProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pQueueFamilyPropertyCount.Raw()), pQueueFamilyProperties.ptr)
}

// GetPhysicalDeviceQueueFamilyProperties2 wraps vkGetPhysicalDeviceQueueFamilyProperties2.
func GetPhysicalDeviceQueueFamilyProperties2(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount ffi.Ref[uint32], pQueueFamilyProperties QueueFamilyProperties2) {
	C.vkGetPhysicalDeviceQueueFamilyProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pQueueFamilyPropertyCount.Raw()), pQueueFamilyProperties.ptr)
}

// GetPhysicalDeviceSparseImageFormatProperties wraps vkGetPhysicalDeviceSparseImageFormatProperties.
func GetPhysicalDeviceSparseImageFormatProperties(physicalDevice PhysicalDevice, format Format, type_ ImageType, samples SampleCountFlags, usage ImageUsageFlags, tiling ImageTiling, pPropertyCount ffi.Ref[uint32], pProperties SparseImageFormatProperties) {
	C.vkGetPhysicalDeviceSparseImageFormatProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), C.VkFormat(format), C.VkImageType(type_), C.VkSampleCountFlagBits(samples), C.VkImageUsageFlags(usage), C.VkImageTiling(tiling), (*C.uint32_t)(pPropertyCount.Raw()), pProperties.ptr)
}

// GetPhysicalDeviceSparseImageFormatProperties2 wraps vkGetPhysicalDeviceSparseImageFormatProperties2.
func GetPhysicalDeviceSparseImageFormatProperties2(physicalDevice PhysicalDevice, pFormatInfo PhysicalDeviceSparseImageFormatInfo2, pPropertyCount ffi.Ref[uint32], pProperties SparseImageFormatProperties2) {
	C.vkGetPhysicalDeviceSparseImageFormatProperties2(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), pFormatInfo.ptr, (*C.uint32_t)(pPropertyCount.Raw()), pProperties.ptr)
}

// GetPhysicalDeviceToolProperties wraps vkGetPhysicalDeviceToolProperties.
func GetPhysicalDeviceToolProperties(physicalDevice PhysicalDevice, pToolCount ffi.Ref[uint32], pToolProperties PhysicalDeviceToolProperties) Result {
	ret := C.vkGetPhysicalDeviceToolProperties(C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)), (*C.uint32_t)(pToolCount.Raw()), pToolProperties.ptr)

	return Result(ret)
}

// GetPipelineCacheData wraps vkGetPipelineCacheData.
func GetPipelineCacheData(device Device, pipelineCache PipelineCache, pDataSize ffi.Ref[uintptr], pData unsafe.Pointer) Result {
	ret := C.vkGetPipelineCacheData(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(pipelineCache), (*C.size_t)(pDataSize.Raw()), pData)

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

// GetRenderAreaGranularity wraps vkGetRenderAreaGranularity.
func GetRenderAreaGranularity(device Device, renderPass RenderPass, pGranularity Extent2D) {
	C.vkGetRenderAreaGranularity(C.VkDevice(unsafe.Pointer(device)), C.VkRenderPass(renderPass), pGranularity.ptr)
}

// GetRenderingAreaGranularity wraps vkGetRenderingAreaGranularity.
func GetRenderingAreaGranularity(device Device, pRenderingAreaInfo RenderingAreaInfo, pGranularity Extent2D) {
	C.vkGetRenderingAreaGranularity(C.VkDevice(unsafe.Pointer(device)), pRenderingAreaInfo.ptr, pGranularity.ptr)
}

// GetSemaphoreCounterValue wraps vkGetSemaphoreCounterValue.
func GetSemaphoreCounterValue(device Device, semaphore Semaphore, pValue ffi.Ref[uint64]) Result {
	ret := C.vkGetSemaphoreCounterValue(C.VkDevice(unsafe.Pointer(device)), C.VkSemaphore(semaphore), (*C.uint64_t)(pValue.Raw()))

	return Result(ret)
}

// InvalidateMappedMemoryRanges wraps vkInvalidateMappedMemoryRanges.
func InvalidateMappedMemoryRanges(device Device, memoryRangeCount uint32, pMemoryRanges MappedMemoryRange) Result {
	ret := C.vkInvalidateMappedMemoryRanges(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(memoryRangeCount), pMemoryRanges.ptr)

	return Result(ret)
}

// vkMapMemory.ppData is unsupported: category pointer2.

// vkMapMemory2.ppData is unsupported: category pointer2.

// MergePipelineCaches wraps vkMergePipelineCaches.
func MergePipelineCaches(device Device, dstCache PipelineCache, srcCacheCount uint32, pSrcCaches ffi.Ref[PipelineCache]) Result {
	ret := C.vkMergePipelineCaches(C.VkDevice(unsafe.Pointer(device)), C.VkPipelineCache(dstCache), C.uint32_t(srcCacheCount), (*C.VkPipelineCache)(pSrcCaches.Raw()))

	return Result(ret)
}

// QueueBindSparse wraps vkQueueBindSparse.
func QueueBindSparse(queue Queue, bindInfoCount uint32, pBindInfo BindSparseInfo, fence Fence) Result {
	ret := C.vkQueueBindSparse(C.VkQueue(unsafe.Pointer(queue)), C.uint32_t(bindInfoCount), pBindInfo.ptr, C.VkFence(fence))

	return Result(ret)
}

// QueueSubmit wraps vkQueueSubmit.
func QueueSubmit(queue Queue, submitCount uint32, pSubmits SubmitInfo, fence Fence) Result {
	ret := C.vkQueueSubmit(C.VkQueue(unsafe.Pointer(queue)), C.uint32_t(submitCount), pSubmits.ptr, C.VkFence(fence))

	return Result(ret)
}

// QueueSubmit2 wraps vkQueueSubmit2.
func QueueSubmit2(queue Queue, submitCount uint32, pSubmits SubmitInfo2, fence Fence) Result {
	ret := C.vkQueueSubmit2(C.VkQueue(unsafe.Pointer(queue)), C.uint32_t(submitCount), pSubmits.ptr, C.VkFence(fence))

	return Result(ret)
}

// QueueWaitIdle wraps vkQueueWaitIdle.
func QueueWaitIdle(queue Queue) Result {
	ret := C.vkQueueWaitIdle(C.VkQueue(unsafe.Pointer(queue)))

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

// SetEvent wraps vkSetEvent.
func SetEvent(device Device, event Event) Result {
	ret := C.vkSetEvent(C.VkDevice(unsafe.Pointer(device)), C.VkEvent(event))

	return Result(ret)
}

// SetPrivateData wraps vkSetPrivateData.
func SetPrivateData(device Device, objectType ObjectType, objectHandle uint64, privateDataSlot PrivateDataSlot, data uint64) Result {
	ret := C.vkSetPrivateData(C.VkDevice(unsafe.Pointer(device)), C.VkObjectType(objectType), C.uint64_t(objectHandle), C.VkPrivateDataSlot(privateDataSlot), C.uint64_t(data))

	return Result(ret)
}

// SignalSemaphore wraps vkSignalSemaphore.
func SignalSemaphore(device Device, pSignalInfo SemaphoreSignalInfo) Result {
	ret := C.vkSignalSemaphore(C.VkDevice(unsafe.Pointer(device)), pSignalInfo.ptr)

	return Result(ret)
}

// TransitionImageLayout wraps vkTransitionImageLayout.
func TransitionImageLayout(device Device, transitionCount uint32, pTransitions HostImageLayoutTransitionInfo) Result {
	ret := C.vkTransitionImageLayout(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(transitionCount), pTransitions.ptr)

	return Result(ret)
}

// TrimCommandPool wraps vkTrimCommandPool.
func TrimCommandPool(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
	C.vkTrimCommandPool(C.VkDevice(unsafe.Pointer(device)), C.VkCommandPool(commandPool), C.VkCommandPoolTrimFlags(flags))
}

// UnmapMemory wraps vkUnmapMemory.
func UnmapMemory(device Device, memory DeviceMemory) {
	C.vkUnmapMemory(C.VkDevice(unsafe.Pointer(device)), C.VkDeviceMemory(memory))
}

// UnmapMemory2 wraps vkUnmapMemory2.
func UnmapMemory2(device Device, pMemoryUnmapInfo MemoryUnmapInfo) Result {
	ret := C.vkUnmapMemory2(C.VkDevice(unsafe.Pointer(device)), pMemoryUnmapInfo.ptr)

	return Result(ret)
}

// UpdateDescriptorSetWithTemplate wraps vkUpdateDescriptorSetWithTemplate.
func UpdateDescriptorSetWithTemplate(device Device, descriptorSet DescriptorSet, descriptorUpdateTemplate DescriptorUpdateTemplate, pData unsafe.Pointer) {
	C.vkUpdateDescriptorSetWithTemplate(C.VkDevice(unsafe.Pointer(device)), C.VkDescriptorSet(descriptorSet), C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate), pData)
}

// UpdateDescriptorSets wraps vkUpdateDescriptorSets.
func UpdateDescriptorSets(device Device, descriptorWriteCount uint32, pDescriptorWrites WriteDescriptorSet, descriptorCopyCount uint32, pDescriptorCopies CopyDescriptorSet) {
	C.vkUpdateDescriptorSets(C.VkDevice(unsafe.Pointer(device)), C.uint32_t(descriptorWriteCount), pDescriptorWrites.ptr, C.uint32_t(descriptorCopyCount), pDescriptorCopies.ptr)
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

// WaitSemaphores wraps vkWaitSemaphores.
func WaitSemaphores(device Device, pWaitInfo SemaphoreWaitInfo, timeout uint64) Result {
	ret := C.vkWaitSemaphores(C.VkDevice(unsafe.Pointer(device)), pWaitInfo.ptr, C.uint64_t(timeout))

	return Result(ret)
}

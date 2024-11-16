package vulkan

import (
	"github.com/csnewman/go-gfx/hal"
	"runtime"
	"unsafe"
)

/*
#include "vulkan.h"

const VkPipelineStageFlagBits2 GFX_VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT = VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT;
const VkPipelineStageFlagBits2 GFX_VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT = VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT;
const VkPipelineStageFlagBits2 GFX_VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT = VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT;

const VkAccessFlagBits2 GFX_VK_ACCESS_2_NONE = VK_ACCESS_2_NONE;
const VkAccessFlagBits2 GFX_VK_ACCESS_2_MEMORY_READ_BIT = VK_ACCESS_2_MEMORY_READ_BIT;
const VkAccessFlagBits2 GFX_VK_ACCESS_2_MEMORY_WRITE_BIT = VK_ACCESS_2_MEMORY_WRITE_BIT;

VkClearValue gfx_vk_clear_color(float r, float g, float b, float a) {
	VkClearValue col;
	col.color.float32[0] = r;
	col.color.float32[1] = g;
	col.color.float32[2] = b;
	col.color.float32[3] = a;
	return col;
}
*/
import "C"

type CommandBuffer struct {
	graphics      *Graphics
	frame         *SurfaceFrame
	commandBuffer C.VkCommandBuffer
}

func (c *CommandBuffer) Barrier(barrier hal.Barrier) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var depInfo C.VkDependencyInfo
	depInfo.sType = C.VK_STRUCTURE_TYPE_DEPENDENCY_INFO

	var imgBarriers []C.VkImageMemoryBarrier2

	for _, halB := range barrier.Textures {
		var imgBarrier C.VkImageMemoryBarrier2
		imgBarrier.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2

		// TODO: reduce stageMask scope
		imgBarrier.srcStageMask = C.GFX_VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT
		imgBarrier.dstStageMask = C.GFX_VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT

		// TODO: access masks (srcAccessMask, dstAccessMask. for now, layout guesses)
		// TODO: transfers (srcQueueFamilyIndex, dstQueueFamilyIndex)

		switch halB.SrcLayout {
		case hal.TextureLayoutUndefined:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_NONE
		case hal.TextureLayoutAttachment:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT | C.GFX_VK_ACCESS_2_MEMORY_WRITE_BIT
		case hal.TextureLayoutRead:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT
		case hal.TextureLayoutPresent:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT
		default:
			panic("unknown layout")
		}

		switch halB.DstLayout {
		case hal.TextureLayoutUndefined:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
			panic("todo check")
		case hal.TextureLayoutAttachment:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL
			imgBarrier.dstAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT | C.GFX_VK_ACCESS_2_MEMORY_WRITE_BIT
		case hal.TextureLayoutRead:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL
			imgBarrier.dstAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT
		case hal.TextureLayoutPresent:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
			imgBarrier.dstAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT
		default:
			panic("unknown layout")
		}

		t, ok := halB.Texture.(*Texture)
		if !ok {
			panic("unexpected type")
		}

		imgBarrier.image = t.img

		// TODO: subresourceRange (e.g. depth)
		imgBarrier.subresourceRange.aspectMask = C.VK_IMAGE_ASPECT_COLOR_BIT
		imgBarrier.subresourceRange.baseMipLevel = 0
		imgBarrier.subresourceRange.levelCount = C.VK_REMAINING_MIP_LEVELS
		imgBarrier.subresourceRange.baseArrayLayer = 0
		imgBarrier.subresourceRange.layerCount = C.VK_REMAINING_ARRAY_LAYERS

		imgBarriers = append(imgBarriers, imgBarrier)
	}

	depInfo.imageMemoryBarrierCount = C.uint32_t(len(imgBarriers))
	depInfo.pImageMemoryBarriers = unsafe.SliceData(imgBarriers)
	pinner.Pin(depInfo.pImageMemoryBarriers)

	C.gfx_vkCmdPipelineBarrier2KHR(c.commandBuffer, &depInfo)
}

func (f *SurfaceFrame) CreateCommandBuffer() hal.CommandBuffer {
	var allocInfo C.VkCommandBufferAllocateInfo
	allocInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	allocInfo.level = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.commandPool = f.entry.commandPool
	allocInfo.commandBufferCount = 1

	var commandBuffer C.VkCommandBuffer

	if err := mapError(C.gfx_vkAllocateCommandBuffers(f.graphics.device, &allocInfo, &commandBuffer)); err != nil {
		panic(err)
	}

	var beginInfo C.VkCommandBufferBeginInfo
	beginInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	beginInfo.flags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT

	if err := mapError(C.gfx_vkBeginCommandBuffer(commandBuffer, &beginInfo)); err != nil {
		panic(err)
	}

	return &CommandBuffer{
		graphics:      f.graphics,
		frame:         f,
		commandBuffer: commandBuffer,
	}
}

func (c *CommandBuffer) BeginRenderPass(description hal.RenderPassDescriptor) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var cAttachs []C.VkRenderingAttachmentInfo

	for _, c := range description.ColorAttachments {
		tv, ok := c.View.(*TextureView)
		if !ok {
			panic("unexpected view type")
		}

		var colorAttachment C.VkRenderingAttachmentInfo
		colorAttachment.sType = C.VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO
		colorAttachment.imageView = tv.view
		colorAttachment.imageLayout = C.VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL

		if c.Load {
			colorAttachment.loadOp = C.VK_ATTACHMENT_LOAD_OP_LOAD
		} else {
			colorAttachment.loadOp = C.VK_ATTACHMENT_LOAD_OP_CLEAR
		}

		if c.Discard {
			// TODO: change discard, as this may corrupt data?
			colorAttachment.loadOp = C.VK_ATTACHMENT_STORE_OP_NONE
		} else {
			colorAttachment.storeOp = C.VK_ATTACHMENT_STORE_OP_STORE
		}

		colorAttachment.clearValue = C.gfx_vk_clear_color(
			C.float(c.ClearColor.R),
			C.float(c.ClearColor.G),
			C.float(c.ClearColor.B),
			C.float(c.ClearColor.A),
		)

		cAttachs = append(cAttachs, colorAttachment)
	}

	var renderingInfo C.VkRenderingInfo
	renderingInfo.sType = C.VK_STRUCTURE_TYPE_RENDERING_INFO
	renderingInfo.renderArea.offset.x = 0
	renderingInfo.renderArea.offset.y = 0
	renderingInfo.renderArea.extent.width = C.uint32_t(c.frame.img.width)
	renderingInfo.renderArea.extent.height = C.uint32_t(c.frame.img.height)
	renderingInfo.layerCount = 1
	renderingInfo.colorAttachmentCount = C.uint32_t(len(cAttachs))
	renderingInfo.pColorAttachments = unsafe.SliceData(cAttachs)
	pinner.Pin(renderingInfo.pColorAttachments)

	C.gfx_vkCmdBeginRenderingKHR(c.commandBuffer, &renderingInfo)

	var viewport C.VkViewport
	viewport.x = C.float(0)
	viewport.y = C.float(0)
	viewport.width = C.float(c.frame.img.width)
	viewport.height = C.float(c.frame.img.height)
	viewport.minDepth = C.float(0)
	viewport.maxDepth = C.float(1)

	C.gfx_vkCmdSetViewportWithCountEXT(c.commandBuffer, 1, &viewport)

	var scissor C.VkRect2D
	scissor.offset.x = 0
	scissor.offset.y = 0
	scissor.extent.width = C.uint32_t(c.frame.img.width)
	scissor.extent.height = C.uint32_t(c.frame.img.height)

	C.gfx_vkCmdSetScissorWithCountEXT(c.commandBuffer, 1, &scissor)
}

func (c *CommandBuffer) SetRenderPipeline(pipeline hal.RenderPipeline) {
	p, ok := pipeline.(*RenderPipeline)
	if !ok {
		panic("unexpected type")
	}

	C.gfx_vkCmdBindPipeline(c.commandBuffer, C.VK_PIPELINE_BIND_POINT_GRAPHICS, p.pipeline)
}

func (c *CommandBuffer) SetVertexBuffer(data hal.Buffer) {
	//TODO implement me
	panic("implement me")
}

func (c *CommandBuffer) Draw(start int, count int) {
	C.gfx_vkCmdDraw(c.commandBuffer, C.uint32_t(count), 1, C.uint32_t(start), 0)
}

func (c *CommandBuffer) EndRenderPass() {
	C.gfx_vkCmdEndRenderingKHR(c.commandBuffer)
}

func (c *CommandBuffer) Submit() {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	C.gfx_vkEndCommandBuffer(c.commandBuffer)

	var cmdinfo C.VkCommandBufferSubmitInfo
	cmdinfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO
	cmdinfo.pNext = nil
	cmdinfo.commandBuffer = c.commandBuffer
	cmdinfo.deviceMask = 0

	var waitInfo C.VkSemaphoreSubmitInfo
	waitInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	waitInfo.pNext = nil
	waitInfo.semaphore = c.frame.entry.imgSem
	waitInfo.stageMask = C.GFX_VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT
	waitInfo.deviceIndex = 0
	waitInfo.value = 1

	var signalInfo C.VkSemaphoreSubmitInfo
	signalInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	signalInfo.pNext = nil
	signalInfo.semaphore = c.frame.entry.completeSem
	signalInfo.stageMask = C.GFX_VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT
	signalInfo.deviceIndex = 0
	signalInfo.value = 1

	var submitInfo C.VkSubmitInfo2
	submitInfo.sType = C.VK_STRUCTURE_TYPE_SUBMIT_INFO_2
	submitInfo.pNext = nil

	submitInfo.waitSemaphoreInfoCount = 1
	submitInfo.pWaitSemaphoreInfos = &waitInfo
	pinner.Pin(submitInfo.pWaitSemaphoreInfos)

	submitInfo.signalSemaphoreInfoCount = 1
	submitInfo.pSignalSemaphoreInfos = &signalInfo
	pinner.Pin(submitInfo.pSignalSemaphoreInfos)

	submitInfo.commandBufferInfoCount = 1
	submitInfo.pCommandBufferInfos = &cmdinfo
	pinner.Pin(submitInfo.pCommandBufferInfos)

	if err := mapError(C.gfx_vkQueueSubmit2KHR(c.graphics.graphicsQueue, 1, &submitInfo, c.frame.entry.fence)); err != nil {
		panic(err)
	}

	// TODO: free buffer

}

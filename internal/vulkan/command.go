package vulkan

import (
	"github.com/csnewman/go-gfx/hal"
	"runtime"
	"unsafe"
)

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
#include "vulkan.h"

VkClearValue gfx_vk_clear_color(float r, float g, float b, float a) {
	VkClearValue col;
	col.color.float32[0] = r;
	col.color.float32[1] = g;
	col.color.float32[2] = b;
	col.color.float32[3] = a;
	return col;
}

void gfx_vkCmdBeginRenderingKHR(
    VkInstance                                  instance,
	VkCommandBuffer                             commandBuffer,
    const VkRenderingInfo*                      pRenderingInfo
) {
	GFX_VK_EXT_FUNC(vkCmdBeginRenderingKHR, commandBuffer, pRenderingInfo);
}

void gfx_vkCmdEndRenderingKHR(
    VkInstance                                  instance,
	VkCommandBuffer                             commandBuffer
) {
	GFX_VK_EXT_FUNC(vkCmdEndRenderingKHR, commandBuffer);
}

VkResult gfx_vkQueueSubmit2KHR(
    VkInstance                                  instance,
	VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo2*                        pSubmits,
    VkFence                                     fence
) {
	return GFX_VK_EXT_FUNC(vkQueueSubmit2KHR, queue, submitCount, pSubmits, fence);
}

*/
import "C"

type CommandBuffer struct {
	graphics      *Graphics
	frame         *SurfaceFrame
	commandBuffer C.VkCommandBuffer
}

func (f *SurfaceFrame) CreateCommandBuffer() hal.CommandBuffer {
	var allocInfo C.VkCommandBufferAllocateInfo
	allocInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	allocInfo.level = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.commandPool = f.entry.commandPool
	allocInfo.commandBufferCount = 1

	var commandBuffer C.VkCommandBuffer

	if err := mapError(C.vkAllocateCommandBuffers(f.graphics.device, &allocInfo, &commandBuffer)); err != nil {
		panic(err)
	}

	var beginInfo C.VkCommandBufferBeginInfo
	beginInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	beginInfo.flags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT

	if err := mapError(C.vkBeginCommandBuffer(commandBuffer, &beginInfo)); err != nil {
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

	C.gfx_vkCmdBeginRenderingKHR(c.graphics.instance, c.commandBuffer, &renderingInfo)
}

func (c *CommandBuffer) SetRenderPipeline(pipeline hal.RenderPipeline) {
	//TODO implement me
	panic("implement me")
}

func (c *CommandBuffer) SetVertexBuffer(data hal.Buffer) {
	//TODO implement me
	panic("implement me")
}

func (c *CommandBuffer) Draw(start int, count int) {
	//TODO implement me
	panic("implement me")
}

func (c *CommandBuffer) EndRenderPass() {
	C.gfx_vkCmdEndRenderingKHR(c.graphics.instance, c.commandBuffer)
}

func (c *CommandBuffer) Submit() {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	C.vkEndCommandBuffer(c.commandBuffer)

	var cmdinfo C.VkCommandBufferSubmitInfo
	cmdinfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO
	cmdinfo.pNext = nil
	cmdinfo.commandBuffer = c.commandBuffer
	cmdinfo.deviceMask = 0

	var waitInfo C.VkSemaphoreSubmitInfo
	waitInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	waitInfo.pNext = nil
	waitInfo.semaphore = c.frame.entry.imgSem
	waitInfo.stageMask = C.VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT_KHR
	waitInfo.deviceIndex = 0
	waitInfo.value = 1

	var signalInfo C.VkSemaphoreSubmitInfo
	signalInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	signalInfo.pNext = nil
	signalInfo.semaphore = c.frame.entry.completeSem
	signalInfo.stageMask = C.VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT
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

	if err := mapError(C.gfx_vkQueueSubmit2KHR(c.graphics.instance, c.graphics.graphicsQueue, 1, &submitInfo, c.frame.entry.fence)); err != nil {
		panic(err)
	}

	// TODO: free buffer

}
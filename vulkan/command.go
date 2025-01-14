package vulkan

import (
	"github.com/csnewman/go-gfx/gfx"
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
	commandBuffer C.VkCommandBuffer
	pool          C.VkCommandPool
	setsBound     bool
}

func (g *Graphics) CreateCommandBuffer() (*CommandBuffer, error) {
	var allocInfo C.VkCommandBufferAllocateInfo
	allocInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	allocInfo.level = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.commandPool = g.mainCommandPool
	allocInfo.commandBufferCount = 1

	var commandBuffer C.VkCommandBuffer

	if err := mapError(C.vkAllocateCommandBuffers(g.device, &allocInfo, &commandBuffer)); err != nil {
		return nil, err
	}

	var beginInfo C.VkCommandBufferBeginInfo
	beginInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	beginInfo.flags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT

	if err := mapError(C.vkBeginCommandBuffer(commandBuffer, &beginInfo)); err != nil {
		return nil, err
	}

	return &CommandBuffer{
		graphics:      g,
		commandBuffer: commandBuffer,
		pool:          g.mainCommandPool,
	}, nil
}

func (c *CommandBuffer) Close() {
	C.vkFreeCommandBuffers(c.graphics.device, c.pool, 1, &c.commandBuffer)
}

type ImageLayout int

const (
	ImageLayoutUndefined ImageLayout = iota
	ImageLayoutAttachment
	ImageLayoutPresent
	ImageLayoutGeneral
)

type ImageBarrier struct {
	Image     *Image
	SrcLayout ImageLayout
	DstLayout ImageLayout
}

type Barrier struct {
	Images []ImageBarrier
}

func (c *CommandBuffer) Barrier(barrier Barrier) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var depInfo C.VkDependencyInfo
	depInfo.sType = C.VK_STRUCTURE_TYPE_DEPENDENCY_INFO

	var imgBarriers []C.VkImageMemoryBarrier2

	for _, bar := range barrier.Images {
		var imgBarrier C.VkImageMemoryBarrier2
		imgBarrier.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2

		// TODO: reduce stageMask scope
		imgBarrier.srcStageMask = C.GFX_VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT
		imgBarrier.dstStageMask = C.GFX_VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT

		// TODO: access masks (srcAccessMask, dstAccessMask. for now, layout guesses)
		// TODO: transfers (srcQueueFamilyIndex, dstQueueFamilyIndex)

		switch bar.SrcLayout {
		case ImageLayoutUndefined:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_NONE
		case ImageLayoutAttachment:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT | C.GFX_VK_ACCESS_2_MEMORY_WRITE_BIT
		case ImageLayoutPresent:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT
		case ImageLayoutGeneral:
			imgBarrier.oldLayout = C.VK_IMAGE_LAYOUT_GENERAL
			imgBarrier.srcAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT | C.GFX_VK_ACCESS_2_MEMORY_WRITE_BIT
		default:
			panic("unknown layout")
		}

		switch bar.DstLayout {
		case ImageLayoutUndefined:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
			panic("todo check")
		case ImageLayoutAttachment:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL
			imgBarrier.dstAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT | C.GFX_VK_ACCESS_2_MEMORY_WRITE_BIT
		case ImageLayoutPresent:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
			imgBarrier.dstAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT
		case ImageLayoutGeneral:
			imgBarrier.newLayout = C.VK_IMAGE_LAYOUT_GENERAL
			imgBarrier.dstAccessMask = C.GFX_VK_ACCESS_2_MEMORY_READ_BIT | C.GFX_VK_ACCESS_2_MEMORY_WRITE_BIT
		default:
			panic("unknown layout")
		}

		imgBarrier.image = bar.Image.image

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

	C.vkCmdPipelineBarrier2KHR(c.commandBuffer, &depInfo)
}

func (f *SurfaceFrame) CreateCommandBuffer() *CommandBuffer {
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
		commandBuffer: commandBuffer,
	}
}

func (c *CommandBuffer) BeginRenderPass(description gfx.RenderPassDescriptor) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var (
		cAttachs    []C.VkRenderingAttachmentInfo
		frameWidth  int
		frameHeight int
	)

	for i, c := range description.ColorAttachments {
		if i == 0 {
			frameWidth = c.Target.Width()
			frameHeight = c.Target.Height()
		}

		vv := c.Target.(*ImageView)

		var colorAttachment C.VkRenderingAttachmentInfo
		colorAttachment.sType = C.VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO
		colorAttachment.imageView = vv.view
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
	renderingInfo.renderArea.extent.width = C.uint32_t(frameWidth)
	renderingInfo.renderArea.extent.height = C.uint32_t(frameHeight)
	renderingInfo.layerCount = 1
	renderingInfo.colorAttachmentCount = C.uint32_t(len(cAttachs))
	renderingInfo.pColorAttachments = unsafe.SliceData(cAttachs)
	pinner.Pin(renderingInfo.pColorAttachments)

	C.vkCmdBeginRenderingKHR(c.commandBuffer, &renderingInfo)

	var viewport C.VkViewport
	viewport.x = C.float(0)
	viewport.y = C.float(0)
	viewport.width = C.float(frameWidth)
	viewport.height = C.float(frameHeight)
	viewport.minDepth = C.float(0)
	viewport.maxDepth = C.float(1)

	C.vkCmdSetViewportWithCountEXT(c.commandBuffer, 1, &viewport)

	var scissor C.VkRect2D
	scissor.offset.x = 0
	scissor.offset.y = 0
	scissor.extent.width = C.uint32_t(frameWidth)
	scissor.extent.height = C.uint32_t(frameHeight)

	C.vkCmdSetScissorWithCountEXT(c.commandBuffer, 1, &scissor)
}

func (c *CommandBuffer) SetRenderPipeline(pipeline gfx.RenderPipeline) {
	vrp := pipeline.(*RenderPipeline)

	if !c.setsBound {
		C.vkCmdBindDescriptorSets(
			c.commandBuffer,
			C.VK_PIPELINE_BIND_POINT_GRAPHICS,
			c.graphics.pipelineLayout,
			0,
			C.uint32_t(len(c.graphics.pipelineSets)),
			unsafe.SliceData(c.graphics.pipelineSets),
			0,
			nil,
		)

		c.setsBound = true
	}

	C.vkCmdBindPipeline(c.commandBuffer, C.VK_PIPELINE_BIND_POINT_GRAPHICS, vrp.pipeline)
}

func (c *CommandBuffer) SetPushConstants(offset int, size int, data unsafe.Pointer) {
	C.vkCmdPushConstants(
		c.commandBuffer,
		c.graphics.pipelineLayout,
		C.VK_SHADER_STAGE_VERTEX_BIT|C.VK_SHADER_STAGE_FRAGMENT_BIT,
		C.uint32_t(offset),
		C.uint32_t(size),
		data,
	)
}

func (c *CommandBuffer) SetVertexBuffer(binding int, buffer gfx.Buffer, offset int) {
	vb := buffer.(*Buffer)
	buf := vb.buffer
	off := C.VkDeviceSize(offset)

	C.vkCmdBindVertexBuffers(c.commandBuffer, C.uint32_t(binding), C.uint32_t(1), &buf, &off)
}

func (c *CommandBuffer) SetIndexBuffer(buffer gfx.Buffer, offset int) {
	vb := buffer.(*Buffer)
	C.vkCmdBindIndexBuffer(c.commandBuffer, vb.buffer, C.VkDeviceSize(offset), C.VK_INDEX_TYPE_UINT32)
}

func (c *CommandBuffer) Draw(start int, count int) {
	C.vkCmdDraw(c.commandBuffer, C.uint32_t(count), 1, C.uint32_t(start), 0)
}

func (c *CommandBuffer) DrawIndexed(start int, count int, vertexOffset int) {
	C.vkCmdDrawIndexed(c.commandBuffer, C.uint32_t(count), 1, C.uint32_t(start), C.int32_t(vertexOffset), 0)
}

func (c *CommandBuffer) CopyBufferToImage(buffer *Buffer, image *Image) {
	var region C.VkBufferImageCopy
	region.imageSubresource.aspectMask = C.VK_IMAGE_ASPECT_COLOR_BIT
	region.imageSubresource.layerCount = 1
	region.imageExtent.width = C.uint32_t(image.width)
	region.imageExtent.height = C.uint32_t(image.height)
	region.imageExtent.depth = 1

	C.vkCmdCopyBufferToImage(
		c.commandBuffer,
		buffer.buffer,
		image.image,
		C.VK_IMAGE_LAYOUT_GENERAL,
		1,
		&region,
	)
}

func (c *CommandBuffer) EndRenderPass() {
	C.vkCmdEndRenderingKHR(c.commandBuffer)
}

// TODO: merge with submit and expose signaling
func (c *CommandBuffer) SubmitAndWait() error {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	C.vkEndCommandBuffer(c.commandBuffer)

	var cmdinfo C.VkCommandBufferSubmitInfo
	cmdinfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO
	cmdinfo.pNext = nil
	cmdinfo.commandBuffer = c.commandBuffer
	cmdinfo.deviceMask = 0

	var submitInfo C.VkSubmitInfo2
	submitInfo.sType = C.VK_STRUCTURE_TYPE_SUBMIT_INFO_2
	submitInfo.pNext = nil
	submitInfo.waitSemaphoreInfoCount = 0
	submitInfo.signalSemaphoreInfoCount = 0
	submitInfo.commandBufferInfoCount = 1
	submitInfo.pCommandBufferInfos = &cmdinfo
	pinner.Pin(submitInfo.pCommandBufferInfos)

	if err := mapError(C.vkQueueSubmit2KHR(c.graphics.graphicsQueue, 1, &submitInfo, nil)); err != nil {
		return err
	}

	if err := mapError(C.vkDeviceWaitIdle(c.graphics.device)); err != nil {
		return err
	}

	return nil
}

func (c *CommandBuffer) SubmitFrame(frame *SurfaceFrame) error {
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
	waitInfo.semaphore = frame.entry.imgSem
	waitInfo.stageMask = C.GFX_VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT
	waitInfo.deviceIndex = 0
	waitInfo.value = 1

	var signalInfo C.VkSemaphoreSubmitInfo
	signalInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	signalInfo.pNext = nil
	signalInfo.semaphore = frame.entry.completeSem
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

	// TODO: split from frame
	if err := mapError(C.vkQueueSubmit2KHR(c.graphics.graphicsQueue, 1, &submitInfo, frame.entry.fence)); err != nil {
		return err
	}

	// TODO: free buffer

	return nil
}

type CommandEncoder struct {
	buffer *CommandBuffer
}

func (g *Graphics) CreateCommandEncoder() (gfx.CommandEncoder, error) {
	buffer, err := g.CreateCommandBuffer()
	if err != nil {
		return nil, err
	}

	return &CommandEncoder{
		buffer: buffer,
	}, nil
}

func (c *CommandEncoder) InitImage(img gfx.Image) {
	vi := img.(*Image)

	c.buffer.Barrier(Barrier{Images: []ImageBarrier{
		{
			Image:     vi,
			SrcLayout: ImageLayoutUndefined,
			DstLayout: ImageLayoutGeneral,
		},
	}})
}

func (c *CommandEncoder) CopyBufferToImage(buffer gfx.Buffer, image gfx.Image) {
	vb := buffer.(*Buffer)
	vi := image.(*Image)

	c.buffer.CopyBufferToImage(vb, vi)
}

func (c *CommandEncoder) SubmitAndWait() error {
	return c.buffer.SubmitAndWait()
}

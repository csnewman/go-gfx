package vulkan

import (
	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"

	"unsafe"
)

type CommandBuffer struct {
	graphics      *Graphics
	commandBuffer vk.CommandBuffer
	pool          vk.CommandPool
	setsBound     bool
}

func (g *Graphics) CreateCommandBuffer() (*CommandBuffer, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	allocInfo := vk.CommandBufferAllocateInfoAlloc(arena, 1)
	allocInfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO)
	allocInfo.SetLevel(vk.COMMAND_BUFFER_LEVEL_PRIMARY)
	allocInfo.SetCommandPool(g.mainCommandPool)
	allocInfo.SetCommandBufferCount(1)

	cmdRef := ffi.RefAlloc[vk.CommandBuffer](arena, 1)

	if err := mapError(vk.AllocateCommandBuffers(g.device, allocInfo, cmdRef)); err != nil {
		return nil, err
	}

	commandBuffer := cmdRef.Get()

	beginInfo := vk.CommandBufferBeginInfoAlloc(arena, 1)
	beginInfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO)
	beginInfo.SetFlags(vk.COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT)

	if err := mapError(vk.BeginCommandBuffer(commandBuffer, beginInfo)); err != nil {
		return nil, err
	}

	return &CommandBuffer{
		graphics:      g,
		commandBuffer: commandBuffer,
		pool:          g.mainCommandPool,
	}, nil
}

func (c *CommandBuffer) Close() {
	arena := ffi.NewArena()
	defer arena.Close()

	commandRef := ffi.RefAlloc[vk.CommandBuffer](arena, 1)
	commandRef.Set(c.commandBuffer)

	vk.FreeCommandBuffers(c.graphics.device, c.pool, 1, commandRef)
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

func (c *CommandBuffer) Barrier(arena *ffi.Arena, barrier Barrier) {
	depInfo := vk.DependencyInfoAlloc(arena, 1)
	depInfo.SetSType(vk.STRUCTURE_TYPE_DEPENDENCY_INFO)

	barriers := vk.ImageMemoryBarrier2Alloc(arena, len(barrier.Images))

	for i, bar := range barrier.Images {
		imgBarrier := barriers.Offset(i)
		imgBarrier.SetSType(vk.STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2)

		// TODO: reduce stageMask scope
		imgBarrier.SetSrcStageMask(vk.PIPELINE_STAGE_2_ALL_COMMANDS_BIT)
		imgBarrier.SetDstStageMask(vk.PIPELINE_STAGE_2_ALL_COMMANDS_BIT)

		// TODO: access masks (srcAccessMask, dstAccessMask. for now, layout guesses)
		// TODO: transfers (srcQueueFamilyIndex, dstQueueFamilyIndex)

		switch bar.SrcLayout {
		case ImageLayoutUndefined:
			imgBarrier.SetOldLayout(vk.IMAGE_LAYOUT_UNDEFINED)
			imgBarrier.SetSrcAccessMask(vk.ACCESS_2_NONE)
		case ImageLayoutAttachment:
			imgBarrier.SetOldLayout(vk.IMAGE_LAYOUT_ATTACHMENT_OPTIMAL)
			imgBarrier.SetSrcAccessMask(vk.ACCESS_2_MEMORY_READ_BIT | vk.ACCESS_2_MEMORY_WRITE_BIT)
		case ImageLayoutPresent:
			imgBarrier.SetOldLayout(vk.IMAGE_LAYOUT_PRESENT_SRC_KHR)
			imgBarrier.SetSrcAccessMask(vk.ACCESS_2_MEMORY_READ_BIT)
		case ImageLayoutGeneral:
			imgBarrier.SetOldLayout(vk.IMAGE_LAYOUT_GENERAL)
			imgBarrier.SetSrcAccessMask(vk.ACCESS_2_MEMORY_READ_BIT | vk.ACCESS_2_MEMORY_WRITE_BIT)
		default:
			panic("unknown layout")
		}

		switch bar.DstLayout {
		case ImageLayoutUndefined:
			imgBarrier.SetNewLayout(vk.IMAGE_LAYOUT_UNDEFINED)
			panic("todo check")
		case ImageLayoutAttachment:
			imgBarrier.SetNewLayout(vk.IMAGE_LAYOUT_ATTACHMENT_OPTIMAL)
			imgBarrier.SetDstAccessMask(vk.ACCESS_2_MEMORY_READ_BIT | vk.ACCESS_2_MEMORY_WRITE_BIT)
		case ImageLayoutPresent:
			imgBarrier.SetNewLayout(vk.IMAGE_LAYOUT_PRESENT_SRC_KHR)
			imgBarrier.SetDstAccessMask(vk.ACCESS_2_MEMORY_READ_BIT)
		case ImageLayoutGeneral:
			imgBarrier.SetNewLayout(vk.IMAGE_LAYOUT_GENERAL)
			imgBarrier.SetDstAccessMask(vk.ACCESS_2_MEMORY_READ_BIT | vk.ACCESS_2_MEMORY_WRITE_BIT)
		default:
			panic("unknown layout")
		}

		imgBarrier.SetImage(bar.Image.image)

		// TODO: subresourceRange (e.g. depth)
		subresourceRange := imgBarrier.RefSubresourceRange()

		subresourceRange.SetAspectMask(vk.IMAGE_ASPECT_COLOR_BIT)
		subresourceRange.SetBaseMipLevel(0)
		subresourceRange.SetLevelCount(vk.REMAINING_MIP_LEVELS)
		subresourceRange.SetBaseArrayLayer(0)
		subresourceRange.SetLayerCount(vk.REMAINING_ARRAY_LAYERS)
	}

	depInfo.SetImageMemoryBarrierCount(uint32(len(barrier.Images)))
	depInfo.SetPImageMemoryBarriers(barriers)

	vk.CmdPipelineBarrier2(c.commandBuffer, depInfo)
}

func (f *SurfaceFrame) CreateCommandBuffer(arena *ffi.Arena) *CommandBuffer {
	var commandBuffer vk.CommandBuffer

	if f.entry.bufferPos < len(f.entry.buffers) {
		commandBuffer = f.entry.buffers[f.entry.bufferPos]
		f.entry.bufferPos++
	} else {
		allocInfo := vk.CommandBufferAllocateInfoAlloc(arena, 1)
		allocInfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO)
		allocInfo.SetLevel(vk.COMMAND_BUFFER_LEVEL_PRIMARY)
		allocInfo.SetCommandPool(f.entry.commandPool)
		allocInfo.SetCommandBufferCount(1)

		cmdBufRef := ffi.RefAlloc[vk.CommandBuffer](arena, 1)

		if err := mapError(vk.AllocateCommandBuffers(f.graphics.device, allocInfo, cmdBufRef)); err != nil {
			panic(err)
		}

		commandBuffer = cmdBufRef.Get()

		f.entry.buffers = append(f.entry.buffers, commandBuffer)
		f.entry.bufferPos = len(f.entry.buffers)
	}

	beginInfo := vk.CommandBufferBeginInfoAlloc(arena, 1)
	beginInfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO)
	beginInfo.SetFlags(vk.COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT)

	if err := mapError(vk.BeginCommandBuffer(commandBuffer, beginInfo)); err != nil {
		panic(err)
	}

	return &CommandBuffer{
		graphics:      f.graphics,
		commandBuffer: commandBuffer,
	}
}

func (c *CommandBuffer) BeginRenderPass(description gfx.RenderPassDescriptor) {
	arena := ffi.NewArena()
	defer arena.Close()

	var (
		frameWidth  int
		frameHeight int
	)

	cAttachs := vk.RenderingAttachmentInfoAlloc(arena, len(description.ColorAttachments))

	for i, c := range description.ColorAttachments {
		if i == 0 {
			frameWidth = c.Target.Width()
			frameHeight = c.Target.Height()
		}

		vv := c.Target.(*ImageView)

		colorAttachment := cAttachs.Offset(i)
		colorAttachment.SetSType(vk.STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO)
		colorAttachment.SetImageView(vv.view)
		colorAttachment.SetImageLayout(vk.IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL)

		if c.Load {
			colorAttachment.SetLoadOp(vk.ATTACHMENT_LOAD_OP_LOAD)
		} else {
			colorAttachment.SetLoadOp(vk.ATTACHMENT_LOAD_OP_CLEAR)
		}

		if c.Discard {
			// TODO: change discard, as this may corrupt data?
			colorAttachment.SetStoreOp(vk.ATTACHMENT_STORE_OP_NONE)
		} else {
			colorAttachment.SetStoreOp(vk.ATTACHMENT_STORE_OP_STORE)
		}

		colorAttachment.TempSetClearValue(
			float32(c.ClearColor.R),
			float32(c.ClearColor.G),
			float32(c.ClearColor.B),
			float32(c.ClearColor.A),
		)
	}

	renderingInfo := vk.RenderingInfoAlloc(arena, 1)
	renderingInfo.SetSType(vk.STRUCTURE_TYPE_RENDERING_INFO)

	renderArea := renderingInfo.RefRenderArea()

	offset := renderArea.RefOffset()
	offset.SetX(0)
	offset.SetY(0)

	extent := renderArea.RefExtent()
	extent.SetWidth(uint32(frameWidth))
	extent.SetHeight(uint32(frameHeight))

	renderingInfo.SetLayerCount(1)

	renderingInfo.SetColorAttachmentCount(uint32(len(description.ColorAttachments)))
	renderingInfo.SetPColorAttachments(cAttachs)

	vk.CmdBeginRendering(c.commandBuffer, renderingInfo)

	viewport := vk.ViewportAlloc(arena, 1)
	viewport.SetX(0)
	viewport.SetY(0)
	viewport.SetWidth(float32(frameWidth))
	viewport.SetHeight(float32(frameHeight))
	viewport.SetMinDepth(0)
	viewport.SetMaxDepth(1)

	vk.CmdSetViewportWithCount(c.commandBuffer, 1, viewport)

	scissor := vk.Rect2DAlloc(arena, 1)

	offset = scissor.RefOffset()
	offset.SetX(0)
	offset.SetY(0)

	extent = scissor.RefExtent()
	extent.SetWidth(uint32(frameWidth))
	extent.SetHeight(uint32(frameHeight))

	vk.CmdSetScissorWithCount(c.commandBuffer, 1, scissor)
}

func (c *CommandBuffer) SetRenderPipeline(pipeline gfx.RenderPipeline) {
	vrp := pipeline.(*RenderPipeline)

	arena := ffi.NewArena()
	defer arena.Close()

	if !c.setsBound {
		// TODO: remove RefFromPtr
		vk.CmdBindDescriptorSets(
			c.commandBuffer,
			vk.PIPELINE_BIND_POINT_GRAPHICS,
			c.graphics.pipelineLayout,
			0,
			uint32(len(c.graphics.pipelineSets)),
			ffi.RefFromValues(arena, c.graphics.pipelineSets...),
			0,
			ffi.RefNil[uint32](),
		)

		c.setsBound = true
	}

	vk.CmdBindPipeline(c.commandBuffer, vk.PIPELINE_BIND_POINT_GRAPHICS, vrp.pipeline)
}

func (c *CommandBuffer) SetPushConstants(offset int, size int, data unsafe.Pointer) {
	vk.CmdPushConstants(
		c.commandBuffer,
		c.graphics.pipelineLayout,
		vk.SHADER_STAGE_VERTEX_BIT|vk.SHADER_STAGE_FRAGMENT_BIT,
		uint32(offset),
		uint32(size),
		data,
	)
}

func (c *CommandBuffer) SetVertexBuffer(binding int, buffer gfx.Buffer, offset int) {
	vb := buffer.(*Buffer)

	arena := ffi.NewArena()
	defer arena.Close()

	vk.CmdBindVertexBuffers(
		c.commandBuffer,
		uint32(binding),
		uint32(1),
		ffi.RefFromValues(arena, vb.buffer),
		ffi.RefFromValues(arena, vk.DeviceSize(offset)),
	)
}

func (c *CommandBuffer) SetIndexBuffer(buffer gfx.Buffer, offset int) {
	vb := buffer.(*Buffer)
	vk.CmdBindIndexBuffer(c.commandBuffer, vb.buffer, vk.DeviceSize(offset), vk.INDEX_TYPE_UINT32)
}

func (c *CommandBuffer) Draw(start int, count int) {
	vk.CmdDraw(c.commandBuffer, uint32(count), 1, uint32(start), 0)
}

func (c *CommandBuffer) DrawIndexed(start int, count int, vertexOffset int) {
	vk.CmdDrawIndexed(c.commandBuffer, uint32(count), 1, uint32(start), int32(vertexOffset), 0)
}

func (c *CommandBuffer) CopyBufferToImage(buffer *Buffer, image *Image) {
	arena := ffi.NewArena()
	defer arena.Close()

	region := vk.BufferImageCopyAlloc(arena, 1)

	imageSubresource := region.RefImageSubresource()
	imageSubresource.SetAspectMask(vk.IMAGE_ASPECT_COLOR_BIT)
	imageSubresource.SetLayerCount(1)

	imageExtent := region.RefImageExtent()
	imageExtent.SetWidth(uint32(image.width))
	imageExtent.SetHeight(uint32(image.height))
	imageExtent.SetDepth(1)

	vk.CmdCopyBufferToImage(
		c.commandBuffer,
		buffer.buffer,
		image.image,
		vk.IMAGE_LAYOUT_GENERAL,
		1,
		region,
	)
}

func (c *CommandBuffer) EndRenderPass() {
	vk.CmdEndRendering(c.commandBuffer)
}

// TODO: merge with submit and expose signaling
func (c *CommandBuffer) SubmitAndWait() error {
	arena := ffi.NewArena()
	defer arena.Close()

	vk.EndCommandBuffer(c.commandBuffer)

	cmdinfo := vk.CommandBufferSubmitInfoAlloc(arena, 1)
	cmdinfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO)
	cmdinfo.SetCommandBuffer(c.commandBuffer)
	cmdinfo.SetDeviceMask(0)

	submitInfo := vk.SubmitInfo2Alloc(arena, 1)
	submitInfo.SetSType(vk.STRUCTURE_TYPE_SUBMIT_INFO_2)
	submitInfo.SetWaitSemaphoreInfoCount(0)
	submitInfo.SetSignalSemaphoreInfoCount(0)
	submitInfo.SetCommandBufferInfoCount(1)
	submitInfo.SetPCommandBufferInfos(cmdinfo)

	if err := mapError(vk.QueueSubmit2(c.graphics.graphicsQueue, 1, submitInfo, vk.FenceNil)); err != nil {
		return err
	}

	if err := mapError(vk.DeviceWaitIdle(c.graphics.device)); err != nil {
		return err
	}

	return nil
}

func (c *CommandBuffer) SubmitFrame(arena *ffi.Arena, frame *SurfaceFrame) error {
	vk.EndCommandBuffer(c.commandBuffer)

	cmdinfo := vk.CommandBufferSubmitInfoAlloc(arena, 1)
	cmdinfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO)
	cmdinfo.SetCommandBuffer(c.commandBuffer)
	cmdinfo.SetDeviceMask(0)

	waitInfo := vk.SemaphoreSubmitInfoAlloc(arena, 1)
	waitInfo.SetSType(vk.STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO)
	waitInfo.SetSemaphore(frame.entry.imgSem)
	waitInfo.SetStageMask(vk.PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT)
	waitInfo.SetDeviceIndex(0)
	waitInfo.SetValue(1)

	signalInfo := vk.SemaphoreSubmitInfoAlloc(arena, 1)
	signalInfo.SetSType(vk.STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO)
	signalInfo.SetSemaphore(frame.entry.completeSem)
	signalInfo.SetStageMask(vk.PIPELINE_STAGE_2_ALL_GRAPHICS_BIT)
	signalInfo.SetDeviceIndex(0)
	signalInfo.SetValue(1)

	submitInfo := vk.SubmitInfo2Alloc(arena, 1)
	submitInfo.SetSType(vk.STRUCTURE_TYPE_SUBMIT_INFO_2)

	submitInfo.SetWaitSemaphoreInfoCount(1)
	submitInfo.SetPWaitSemaphoreInfos(waitInfo)

	submitInfo.SetSignalSemaphoreInfoCount(1)
	submitInfo.SetPSignalSemaphoreInfos(signalInfo)

	submitInfo.SetCommandBufferInfoCount(1)
	submitInfo.SetPCommandBufferInfos(cmdinfo)

	// TODO: split from frame
	if err := mapError(vk.QueueSubmit2(c.graphics.graphicsQueue, 1, submitInfo, frame.entry.fence)); err != nil {
		return err
	}

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
	arena := ffi.NewArena()
	defer arena.Close()

	vi := img.(*Image)

	c.buffer.Barrier(arena, Barrier{Images: []ImageBarrier{
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

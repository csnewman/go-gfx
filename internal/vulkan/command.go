package vulkan

import "github.com/csnewman/go-gfx/hal"

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
#include "vulkan.h"
*/
import "C"

type CommandBuffer struct {
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

	return &CommandBuffer{}
}

func (c CommandBuffer) BeginRenderPass(description hal.RenderPassDescriptor) {
	//TODO implement me
	panic("implement me")
}

func (c CommandBuffer) SetRenderPipeline(pipeline hal.RenderPipeline) {
	//TODO implement me
	panic("implement me")
}

func (c CommandBuffer) SetVertexBuffer(data hal.Buffer) {
	//TODO implement me
	panic("implement me")
}

func (c CommandBuffer) Draw(start int, count int) {
	//TODO implement me
	panic("implement me")
}

func (c CommandBuffer) EndRenderPass() {
	//TODO implement me
	panic("implement me")
}

func (c CommandBuffer) Submit() {
	//TODO implement me
	panic("implement me")
}

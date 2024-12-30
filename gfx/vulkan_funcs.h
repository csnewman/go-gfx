/* --- VMA --- */

GFX_FUNC(PFN_vkVoidFunction, vkGetDeviceProcAddr,
    (VkDevice, device),
    (const char*, pName)
);

GFX_FUNC(void, vkGetPhysicalDeviceMemoryProperties,
    (VkPhysicalDevice, physicalDevice),
    (VkPhysicalDeviceMemoryProperties*, pMemoryProperties)
);

GFX_FUNC(VkResult, vkAllocateMemory,
    (VkDevice, device),
    (const VkMemoryAllocateInfo*, pAllocateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkDeviceMemory*, pMemory)
);

GFX_FUNC(void, vkFreeMemory,
    (VkDevice, device),
    (VkDeviceMemory, memory),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(VkResult, vkMapMemory,
    (VkDevice, device),
    (VkDeviceMemory, memory),
    (VkDeviceSize, offset),
    (VkDeviceSize, size),
    (VkMemoryMapFlags, flags),
    (void**, ppData)
);

GFX_FUNC(void, vkUnmapMemory,
    (VkDevice, device),
    (VkDeviceMemory, memory)
);

GFX_FUNC(VkResult, vkFlushMappedMemoryRanges,
    (VkDevice, device),
    (uint32_t, memoryRangeCount),
    (const VkMappedMemoryRange*, pMemoryRanges)
);

GFX_FUNC(VkResult, vkInvalidateMappedMemoryRanges,
    (VkDevice, device),
    (uint32_t, memoryRangeCount),
    (const VkMappedMemoryRange*, pMemoryRanges)
);

GFX_FUNC(VkResult, vkBindBufferMemory,
    (VkDevice, device),
    (VkBuffer, buffer),
    (VkDeviceMemory, memory),
    (VkDeviceSize, memoryOffset)
);

GFX_FUNC(VkResult, vkBindImageMemory,
    (VkDevice, device),
    (VkImage, image),
    (VkDeviceMemory, memory),
    (VkDeviceSize, memoryOffset)
);

GFX_FUNC(void, vkGetBufferMemoryRequirements,
    (VkDevice, device),
    (VkBuffer, buffer),
    (VkMemoryRequirements*, pMemoryRequirements)
);

GFX_FUNC(void, vkGetImageMemoryRequirements,
    (VkDevice, device),
    (VkImage, image),
    (VkMemoryRequirements*, pMemoryRequirements)
);

GFX_FUNC(VkResult, vkCreateBuffer,
    (VkDevice, device),
    (const VkBufferCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkBuffer*, pBuffer)
);

GFX_FUNC(void, vkDestroyBuffer,
    (VkDevice, device),
    (VkBuffer, buffer),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(VkResult, vkCreateImage,
    (VkDevice, device),
    (const VkImageCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkImage*, pImage)
);

GFX_FUNC(void, vkDestroyImage,
    (VkDevice, device),
    (VkImage, image),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(void, vkCmdCopyBuffer,
    (VkCommandBuffer, commandBuffer),
    (VkBuffer, srcBuffer),
    (VkBuffer, dstBuffer),
    (uint32_t, regionCount),
    (const VkBufferCopy*, pRegions)
);

GFX_FUNC(void, vkGetBufferMemoryRequirements2,
    (VkDevice, device),
    (const VkBufferMemoryRequirementsInfo2*, pInfo),
    (VkMemoryRequirements2*, pMemoryRequirements)
);

GFX_FUNC(void, vkGetImageMemoryRequirements2,
    (VkDevice, device),
    (const VkImageMemoryRequirementsInfo2*, pInfo),
    (VkMemoryRequirements2*, pMemoryRequirements)
);

GFX_FUNC(VkResult, vkBindBufferMemory2,
    (VkDevice, device),
    (uint32_t, bindInfoCount),
    (const VkBindBufferMemoryInfo*, pBindInfos)
);

GFX_FUNC(VkResult, vkBindImageMemory2,
    (VkDevice, device),
    (uint32_t, bindInfoCount),
    (const VkBindImageMemoryInfo*, pBindInfos)
);

GFX_FUNC(void, vkGetPhysicalDeviceMemoryProperties2,
    (VkPhysicalDevice, physicalDevice),
    (VkPhysicalDeviceMemoryProperties2*, pMemoryProperties)
);

GFX_FUNC(void, vkGetDeviceBufferMemoryRequirements,
    (VkDevice, device),
    (const VkDeviceBufferMemoryRequirements*, pInfo),
    (VkMemoryRequirements2*, pMemoryRequirements)
);

GFX_FUNC(void, vkGetDeviceImageMemoryRequirements,
    (VkDevice, device),
    (const VkDeviceImageMemoryRequirements*, pInfo),
    (VkMemoryRequirements2*, pMemoryRequirements)
);

/* --- instance/device --- */

GFX_FUNC(VkResult, vkCreateInstance,
    (const VkInstanceCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkInstance*, pInstance)
);

GFX_FUNC(VkResult, vkEnumeratePhysicalDevices,
    (VkInstance, instance),
    (uint32_t*, pPhysicalDeviceCount),
    (VkPhysicalDevice*, pPhysicalDevices)
);

GFX_FUNC(void, vkGetPhysicalDeviceProperties,
    (VkPhysicalDevice, physicalDevice),
    (VkPhysicalDeviceProperties*, pProperties)
);

GFX_FUNC(void, vkGetPhysicalDeviceQueueFamilyProperties,
    (VkPhysicalDevice, physicalDevice),
    (uint32_t*, pQueueFamilyPropertyCount),
    (VkQueueFamilyProperties*, pQueueFamilyProperties)
);

GFX_FUNC(VkResult, vkEnumerateDeviceExtensionProperties,
    (VkPhysicalDevice, physicalDevice),
    (const char*, pLayerName),
    (uint32_t*, pPropertyCount),
    (VkExtensionProperties*, pProperties)
);

GFX_FUNC(VkResult, vkCreateDevice,
    (VkPhysicalDevice, physicalDevice),
    (const VkDeviceCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkDevice*, pDevice)
);

GFX_FUNC(void, vkGetDeviceQueue,
    (VkDevice, device),
    (uint32_t, queueFamilyIndex),
    (uint32_t, queueIndex),
    (VkQueue*, pQueue)
);

GFX_FUNC(VkResult, vkCreateDebugUtilsMessengerEXT,
    (VkInstance, instance),
    (const VkDebugUtilsMessengerCreateInfoEXT*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkDebugUtilsMessengerEXT*, pMessenger)
);

GFX_FUNC(VkResult, vkDeviceWaitIdle,
    (VkDevice,device)
);

/* --- surface --- */

GFX_FUNC(VkResult, vkCreateMetalSurfaceEXT,
    (VkInstance, instance),
    (const VkMetalSurfaceCreateInfoEXT*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSurfaceKHR*, pSurface)
);

GFX_FUNC(VkResult, vkCreateWin32SurfaceKHR,
    (VkInstance, instance),
    (const VkWin32SurfaceCreateInfoKHR*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSurfaceKHR*, pSurface)
);

GFX_FUNC(VkResult, vkGetPhysicalDeviceSurfaceCapabilitiesKHR,
    (VkPhysicalDevice, physicalDevice),
    (VkSurfaceKHR, surface),
    (VkSurfaceCapabilitiesKHR*, pSurfaceCapabilities)
);

GFX_FUNC(VkResult, vkGetPhysicalDeviceSurfaceFormatsKHR,
    (VkPhysicalDevice, physicalDevice),
    (VkSurfaceKHR, surface),
    (uint32_t*, pSurfaceFormatCount),
    (VkSurfaceFormatKHR*, pSurfaceFormats)
);

GFX_FUNC(VkResult, vkCreateSwapchainKHR,
    (VkDevice, device),
    (const VkSwapchainCreateInfoKHR*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSwapchainKHR*, pSwapchain)
);

GFX_FUNC(void, vkDestroySwapchainKHR,
    (VkDevice, device),
    (VkSwapchainKHR, swapchain),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(VkResult, vkGetSwapchainImagesKHR,
    (VkDevice, device),
    (VkSwapchainKHR, swapchain),
    (uint32_t*, pSwapchainImageCount),
    (VkImage*, pSwapchainImages)
);

GFX_FUNC(VkResult, vkAcquireNextImageKHR,
    (VkDevice, device),
    (VkSwapchainKHR, swapchain),
    (uint64_t, timeout),
    (VkSemaphore, semaphore),
    (VkFence, fence),
    (uint32_t*, pImageIndex)
);

GFX_FUNC(VkResult, vkQueuePresentKHR,
    (VkQueue, queue),
    (const VkPresentInfoKHR*, pPresentInfo)
);

/* --- sync --- */

GFX_FUNC(VkResult, vkCreateSemaphore,
    (VkDevice, device),
    (const VkSemaphoreCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSemaphore*, pSemaphore)
);

GFX_FUNC(void, vkDestroySemaphore,
    (VkDevice, device),
    (VkSemaphore, semaphore),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(VkResult, vkCreateFence,
    (VkDevice, device),
    (const VkFenceCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkFence*, pFence)
);

GFX_FUNC(void, vkDestroyFence,
    (VkDevice, device),
    (VkFence, fence),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(VkResult, vkWaitForFences,
    (VkDevice, device),
    (uint32_t, fenceCount),
    (const VkFence*, pFences),
    (VkBool32, waitAll),
    (uint64_t, timeout)
);

GFX_FUNC(VkResult, vkResetFences,
    (VkDevice, device),
    (uint32_t, fenceCount),
    (const VkFence*, pFences)
);

/* --- command --- */

GFX_FUNC(VkResult, vkCreateCommandPool,
    (VkDevice, device),
    (const VkCommandPoolCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkCommandPool*, pCommandPool)
);

GFX_FUNC(void, vkDestroyCommandPool,
    (VkDevice, device),
    (VkCommandPool, commandPool),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(VkResult, vkAllocateCommandBuffers,
    (VkDevice, device),
    (const VkCommandBufferAllocateInfo*, pAllocateInfo),
    (VkCommandBuffer*, pCommandBuffers)
);

GFX_FUNC(void, vkFreeCommandBuffers,
    (VkDevice, device),
    (VkCommandPool, commandPool),
    (uint32_t, commandBufferCount),
    (const VkCommandBuffer*, pCommandBuffers)
);

GFX_FUNC(VkResult, vkBeginCommandBuffer,
    (VkCommandBuffer, commandBuffer),
    (const VkCommandBufferBeginInfo*, pBeginInfo)
);

GFX_FUNC(void, vkCmdBindPipeline,
    (VkCommandBuffer, commandBuffer),
    (VkPipelineBindPoint, pipelineBindPoint),
    (VkPipeline, pipeline)
);

GFX_FUNC(void, vkCmdDraw,
    (VkCommandBuffer, commandBuffer),
    (uint32_t, vertexCount),
    (uint32_t, instanceCount),
    (uint32_t, firstVertex),
    (uint32_t, firstInstance)
);

GFX_FUNC(VkResult, vkEndCommandBuffer,
    (VkCommandBuffer, commandBuffer)
);

/* --- graphics --- */

GFX_FUNC(VkResult, vkCreateShaderModule,
    (VkDevice, device),
    (const VkShaderModuleCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkShaderModule*, pShaderModule)
);

GFX_FUNC(VkResult, vkCreatePipelineLayout,
    (VkDevice, device),
    (const VkPipelineLayoutCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkPipelineLayout*, pPipelineLayout)
);

GFX_FUNC(VkResult, vkCreateGraphicsPipelines,
    (VkDevice, device),
    (VkPipelineCache, pipelineCache),
    (uint32_t, createInfoCount),
    (const VkGraphicsPipelineCreateInfo*, pCreateInfos),
    (const VkAllocationCallbacks*, pAllocator),
    (VkPipeline*, pPipelines)
);

GFX_FUNC(void, vkCmdBeginRenderingKHR,
  	(VkCommandBuffer, commandBuffer),
    (const VkRenderingInfo*, pRenderingInfo)
);

GFX_FUNC(void, vkCmdBindVertexBuffers,
    (VkCommandBuffer, commandBuffer),
    (uint32_t, firstBinding),
    (uint32_t, bindingCount),
    (const VkBuffer*, pBuffers),
    (const VkDeviceSize*, pOffsets)
);

GFX_FUNC(void, vkCmdEndRenderingKHR,
	(VkCommandBuffer, commandBuffer)
);

GFX_FUNC(VkResult, vkQueueSubmit2KHR,
    (VkQueue, queue),
    (uint32_t, submitCount),
    (const VkSubmitInfo2*, pSubmits),
    (VkFence, fence)
);

GFX_FUNC(void, vkCmdSetViewportWithCountEXT,
	(VkCommandBuffer, commandBuffer),
    (uint32_t, viewportCount),
    (const VkViewport*, pViewports)
);

GFX_FUNC(void, vkCmdSetScissorWithCountEXT,
	(VkCommandBuffer, commandBuffer),
    (uint32_t, scissorCount),
    (const VkRect2D*, pScissors)
);

GFX_FUNC(void, vkCmdPipelineBarrier2KHR,
	(VkCommandBuffer, commandBuffer),
    (const VkDependencyInfo*, pDependencyInfo)
);

/* --- resources --- */

GFX_FUNC(VkResult, vkCreateImageView,
    (VkDevice, device),
    (const VkImageViewCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkImageView*, pView)
);

GFX_FUNC(void, vkDestroyImageView,
    (VkDevice, device),
    (VkImageView, imageView),
    (const VkAllocationCallbacks*, pAllocator)
);

GFX_FUNC(void, vkCmdCopyBufferToImage,
    (VkCommandBuffer, commandBuffer),
    (VkBuffer, srcBuffer),
    (VkImage, dstImage),
    (VkImageLayout, dstImageLayout),
    (uint32_t, regionCount),
    (const VkBufferImageCopy*, pRegions)
);

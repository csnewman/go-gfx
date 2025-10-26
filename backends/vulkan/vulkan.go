package vulkan

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/bits-and-blooms/bitset"
	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
	"github.com/csnewman/go-gfx/vkloader"
	"github.com/csnewman/go-gfx/vma"
)

const (
	portabilityExtension = "VK_KHR_portability_subset"
	validationLayer      = "VK_LAYER_KHRONOS_validation"
)

var instance *Graphics

func Register() {
	if instance != nil {
		return
	}

	instance = &Graphics{}

	gfx.RegisterGraphics(instance)
}

type Graphics struct {
	logger   *slog.Logger
	platform gfx.VulkanPlatform

	instance        vk.Instance
	debugMessenger  vk.DebugUtilsMessengerEXT
	physicalDevice  vk.PhysicalDevice
	graphicsFamily  int
	device          vk.Device
	graphicsQueue   vk.Queue
	memoryAllocator vma.Allocator
	mainCommandPool vk.CommandPool

	pipelineLayout vk.PipelineLayout
	pipelineSets   []vk.DescriptorSet

	textureLayout vk.DescriptorSetLayout
	textureSet    vk.DescriptorSet
	textureIDs    bitset.BitSet
	lastTextureID uint

	samplerLayout vk.DescriptorSetLayout
	samplerSet    vk.DescriptorSet
	samplerIDs    bitset.BitSet
	lastSamplerID uint
}

func (g *Graphics) Name() string {
	return "vulkan"
}

func (g *Graphics) Priority() int {
	return 0
}

func (g *Graphics) Available(plt gfx.Platform) bool {
	_, ok := plt.(gfx.VulkanPlatform)

	// TODO: check if vulkan is loadable?
	return ok
}

func (g *Graphics) Init(init gfx.GraphicsInit) error {
	g.logger = init.Logger

	var ok bool

	g.platform, ok = init.Platform.(gfx.VulkanPlatform)
	if !ok {
		return gfx.ErrUnsupportedPlatform
	}

	if err := g.platform.LoadVulkan(); err != nil {
		return fmt.Errorf("failed to load vulkan: %w", err)
	}

	arena := ffi.NewArena()
	defer arena.Close()

	if err := g.createInstance(arena); err != nil {
		return err
	}

	device, err := g.selectDevice(arena)
	if err != nil {
		return err
	}

	if err := g.createDevice(arena, device); err != nil {
		return err
	}

	if err := g.createDescriptors(arena); err != nil {
		return err
	}

	return nil
}

func (g *Graphics) createInstance(arena *ffi.Arena) error {
	defer arena.Checkpoint()()

	// TODO: check null
	vkloader.Load(g.platform.VKInstanceProcAddr())

	appInfo := vk.ApplicationInfoAlloc(arena, 1)
	appInfo.SetSType(vk.STRUCTURE_TYPE_APPLICATION_INFO)
	appInfo.SetPApplicationName(ffi.CStringAlloc(arena, "TODO"))
	appInfo.SetApplicationVersion(vk.MakeAPIVersion(0, 0, 1, 0))
	appInfo.SetPEngineName(ffi.CStringAlloc(arena, "go-gfx"))
	appInfo.SetEngineVersion(vk.MakeAPIVersion(0, 0, 1, 0))
	appInfo.SetApiVersion(vk.APIVersion_1_4)

	instInfo := vk.InstanceCreateInfoAlloc(arena, 1)
	instInfo.SetSType(vk.STRUCTURE_TYPE_INSTANCE_CREATE_INFO)
	instInfo.SetPApplicationInfo(appInfo)

	var exts []string

	exts = append(exts, vk.KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME)

	instInfo.SetFlags(vk.INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR)

	exts = append(exts, vk.KHR_SURFACE_EXTENSION_NAME)

	for _, e := range g.platform.RequiredVKExtensions() {
		exts = append(exts, e)
	}

	exts = append(exts, vk.EXT_DEBUG_UTILS_EXTENSION_NAME)

	instInfo.SetEnabledExtensionCount(uint32(len(exts)))
	instInfo.SetPpEnabledExtensionNames(convertStringArray(arena, exts))

	var layers []string

	layers = append(layers, validationLayer)

	instInfo.SetEnabledLayerCount(1)
	instInfo.SetPpEnabledLayerNames(convertStringArray(arena, layers))

	instOut := ffi.RefAlloc[vk.Instance](arena, 1)

	if err := mapError(vk.CreateInstance(instInfo, vk.AllocationCallbacksNil, instOut)); err != nil {
		return err
	}

	g.instance = instOut.Get()

	vkloader.LoadInstance(g.instance)

	vk.Logger = g.logger

	debugInfo := vk.DebugUtilsMessengerCreateInfoEXTAlloc(arena, 1)
	debugInfo.SetSType(vk.STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT)
	debugInfo.SetMessageSeverity(vk.DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT | vk.DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT | vk.DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT)
	debugInfo.SetMessageType(vk.DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT | vk.DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT | vk.DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT)
	debugInfo.TempSetCallback()

	messengerRef := ffi.RefAlloc[vk.DebugUtilsMessengerEXT](arena, 1)

	if err := mapError(vk.CreateDebugUtilsMessengerEXT(g.instance, debugInfo, vk.AllocationCallbacksNil, messengerRef)); err != nil {
		return err
	}

	g.debugMessenger = messengerRef.Get()

	return nil
}

type selectedDevice struct {
	name           string
	score          int
	device         vk.PhysicalDevice
	graphicsFamily int
	portability    bool
}

func (g *Graphics) selectDevice(arena *ffi.Arena) (*selectedDevice, error) {
	defer arena.Checkpoint()()

	devCount := ffi.RefAlloc[uint32](arena, 1)

	if err := mapError(vk.EnumeratePhysicalDevices(g.instance, devCount, ffi.RefNil[vk.PhysicalDevice]())); err != nil {
		return nil, err
	}

	physicalDeviceRef := ffi.RefAlloc[vk.PhysicalDevice](arena, int(devCount.Get()))

	if err := mapError(vk.EnumeratePhysicalDevices(g.instance, devCount, physicalDeviceRef)); err != nil {
		return nil, err
	}

	physicalDevices := physicalDeviceRef.Slice(int(devCount.Get()))

	currentScore := -1

	var bestDevice *selectedDevice

	for _, device := range physicalDevices {
		option, err := g.considerDevice(arena, device)
		if errors.Is(err, gfx.ErrUnsupportedDevice) {
			g.logger.Warn("Device not supported", "err", err)

			continue
		}

		if err != nil {
			return nil, fmt.Errorf("error considering device: %v", err)
		}

		if option.score > currentScore {
			currentScore = option.score
			bestDevice = option
		}
	}

	if currentScore < 0 {
		return nil, gfx.ErrNoSuitableDevice
	}

	return bestDevice, nil
}

func (g *Graphics) considerDevice(arena *ffi.Arena, device vk.PhysicalDevice) (*selectedDevice, error) {
	defer arena.Checkpoint()()

	props := vk.PhysicalDevicePropertiesAlloc(arena, 1)

	vk.GetPhysicalDeviceProperties(device, props)

	// TODO: check if device can present

	// Ensure 1.4 support
	if props.GetApiVersion() < vk.APIVersion_1_4 {
		return nil, fmt.Errorf("%w: no vk1.4 support: %v < %v", gfx.ErrUnsupportedDevice, props.GetApiVersion(), vk.APIVersion_1_4)
	}

	score := 0

	switch props.GetDeviceType() {
	case vk.PHYSICAL_DEVICE_TYPE_OTHER:
		score = 1
	case vk.PHYSICAL_DEVICE_TYPE_CPU:
		score = 2
	case vk.PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU:
		score = 3
	case vk.PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU:
		score = 4
	case vk.PHYSICAL_DEVICE_TYPE_DISCRETE_GPU:
		score = 5
	default:
		return nil, fmt.Errorf("%w: unknown type", gfx.ErrUnsupportedDevice)
	}

	// Check features
	features12 := vk.PhysicalDeviceVulkan12FeaturesAlloc(arena, 1)
	features12.SetSType(vk.STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES)

	feat2 := vk.PhysicalDeviceFeatures2Alloc(arena, 1)
	feat2.SetSType(vk.STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2)
	feat2.SetPNext(features12.Raw())

	vk.GetPhysicalDeviceFeatures2(device, feat2)

	// descriptor features are required
	if !features12.GetDescriptorBindingPartiallyBound() || !features12.GetRuntimeDescriptorArray() {
		return nil, fmt.Errorf("%w: missing features", gfx.ErrUnsupportedDevice)
	}

	queueFamilyCount := ffi.RefAlloc[uint32](arena, 1)

	vk.GetPhysicalDeviceQueueFamilyProperties(device, queueFamilyCount, vk.QueueFamilyPropertiesNil)

	queueFamilies := vk.QueueFamilyPropertiesAlloc(arena, int(queueFamilyCount.Get()))

	vk.GetPhysicalDeviceQueueFamilyProperties(device, queueFamilyCount, queueFamilies)

	graphicsQueue := -1

	for i := range int(queueFamilyCount.Get()) {
		family := queueFamilies.Offset(i)

		if family.GetQueueFlags()&vk.QUEUE_GRAPHICS_BIT != 0 {
			graphicsQueue = i

			break
		}
	}

	if graphicsQueue == -1 {
		return nil, fmt.Errorf("%w: no graphics queue", gfx.ErrUnsupportedDevice)
	}

	extensionCount := ffi.RefAlloc[uint32](arena, 1)

	if err := mapError(vk.EnumerateDeviceExtensionProperties(
		device,
		ffi.CStringNil,
		extensionCount,
		vk.ExtensionPropertiesNil,
	)); err != nil {
		return nil, err
	}

	extensions := vk.ExtensionPropertiesAlloc(arena, int(extensionCount.Get()))

	if err := mapError(vk.EnumerateDeviceExtensionProperties(
		device,
		ffi.CStringNil,
		extensionCount,
		extensions,
	)); err != nil {
		return nil, err
	}

	portability := false

	for i := range int(extensionCount.Get()) {
		ext := extensions.Offset(i)
		name := ext.GetExtensionNameString()

		if name == portabilityExtension {
			portability = true
		}
	}

	// TODO: other scoring tie breakers

	return &selectedDevice{
		name:           props.GetDeviceNameString(),
		score:          score,
		device:         device,
		graphicsFamily: graphicsQueue,
		portability:    portability,
	}, nil
}

func (g *Graphics) createDevice(arena *ffi.Arena, sel *selectedDevice) error {
	defer arena.Checkpoint()()

	g.logger.Info("Create device", "name", sel.name)

	g.physicalDevice = sel.device
	g.graphicsFamily = sel.graphicsFamily

	// now not required?
	//extendedDynamicState := vk.PhysicalDeviceExtendedDynamicStateFeatures
	//extendedDynamicState.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT
	//extendedDynamicState.extendedDynamicState = C.VkBool32(1)
	//extendedDynamicState.pNext = unsafe.Pointer(&synchronization2Features)
	//pinner.Pin(extendedDynamicState.pNext)

	features13 := vk.PhysicalDeviceVulkan13FeaturesAlloc(arena, 1)
	features13.SetSType(vk.STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_3_FEATURES)
	features13.SetSynchronization2(true)
	features13.SetDynamicRendering(true)

	features12 := vk.PhysicalDeviceVulkan12FeaturesAlloc(arena, 1)
	features12.SetSType(vk.STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES)
	features12.SetPNext(features13.Raw())
	features12.SetBufferDeviceAddress(true)
	features12.SetDescriptorIndexing(true)
	//VkBool32           shaderInputAttachmentArrayDynamicIndexing;
	//VkBool32           shaderUniformTexelBufferArrayDynamicIndexing;
	//VkBool32           shaderStorageTexelBufferArrayDynamicIndexing;
	//VkBool32           shaderUniformBufferArrayNonUniformIndexing;
	features12.SetShaderSampledImageArrayNonUniformIndexing(true)
	//VkBool32           shaderStorageBufferArrayNonUniformIndexing;
	//VkBool32           shaderStorageImageArrayNonUniformIndexing;
	//VkBool32           shaderInputAttachmentArrayNonUniformIndexing;
	//VkBool32           shaderUniformTexelBufferArrayNonUniformIndexing;
	//VkBool32           shaderStorageTexelBufferArrayNonUniformIndexing;
	//VkBool32           descriptorBindingUniformBufferUpdateAfterBind;
	//VkBool32           descriptorBindingSampledImageUpdateAfterBind;
	features12.SetDescriptorBindingSampledImageUpdateAfterBind(true)
	//VkBool32           descriptorBindingStorageImageUpdateAfterBind;
	//VkBool32           descriptorBindingStorageBufferUpdateAfterBind;
	//VkBool32           descriptorBindingUniformTexelBufferUpdateAfterBind;
	//VkBool32           descriptorBindingStorageTexelBufferUpdateAfterBind;
	//VkBool32           descriptorBindingUpdateUnusedWhilePending;
	features12.SetDescriptorBindingPartiallyBound(true)
	features12.SetDescriptorBindingVariableDescriptorCount(true)
	features12.SetRuntimeDescriptorArray(true)

	priorities := ffi.RefFromValues[float32](arena, 1.0)

	queueCreateInfo := vk.DeviceQueueCreateInfoAlloc(arena, 1)
	queueCreateInfo.SetSType(vk.STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO)
	queueCreateInfo.SetQueueFamilyIndex(uint32(sel.graphicsFamily))
	queueCreateInfo.SetQueueCount(1)
	queueCreateInfo.SetPQueuePriorities(priorities)

	createInfo := vk.DeviceCreateInfoAlloc(arena, 1)
	createInfo.SetSType(vk.STRUCTURE_TYPE_DEVICE_CREATE_INFO)
	createInfo.SetPNext(features12.Raw())
	createInfo.SetQueueCreateInfoCount(1)
	createInfo.SetPQueueCreateInfos(queueCreateInfo)

	var exts []string

	if sel.portability {
		exts = append(exts, portabilityExtension)
	}

	exts = append(exts, vk.KHR_SWAPCHAIN_EXTENSION_NAME)

	createInfo.SetEnabledExtensionCount(uint32(len(exts)))
	createInfo.SetPpEnabledExtensionNames(convertStringArray(arena, exts))

	deviceRef := ffi.RefAlloc[vk.Device](arena, 1)

	if err := mapError(vk.CreateDevice(sel.device, createInfo, vk.AllocationCallbacksNil, deviceRef)); err != nil {
		return err
	}

	g.device = deviceRef.Get()

	queueRef := ffi.RefAlloc[vk.Queue](arena, 1)

	vk.GetDeviceQueue(g.device, uint32(sel.graphicsFamily), 0, queueRef)

	g.graphicsQueue = queueRef.Get()

	vmaInfo := vma.AllocatorCreateInfoAlloc(arena, 1)
	vmaInfo.SetVulkanApiVersion(vk.APIVersion_1_4)
	vmaInfo.SetPhysicalDevice(sel.device)
	vmaInfo.SetDevice(g.device)
	vmaInfo.SetInstance(g.instance)
	vmaInfo.SetFlags(vma.ALLOCATOR_CREATE_BUFFER_DEVICE_ADDRESS_BIT)

	allocatorRef := ffi.RefAlloc[vma.Allocator](arena, 1)

	if err := mapError(vma.CreateAllocator(vmaInfo, allocatorRef)); err != nil {
		return err
	}

	g.memoryAllocator = allocatorRef.Get()

	commandInfo := vk.CommandPoolCreateInfoAlloc(arena, 1)
	commandInfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO)
	commandInfo.SetQueueFamilyIndex(uint32(g.graphicsFamily))
	commandInfo.SetFlags(vk.COMMAND_POOL_CREATE_TRANSIENT_BIT)

	poolRef := ffi.RefAlloc[vk.CommandPool](arena, 1)

	if err := mapError(vk.CreateCommandPool(g.device, commandInfo, vk.AllocationCallbacksNil, poolRef)); err != nil {
		return err
	}

	g.mainCommandPool = poolRef.Get()

	return nil
}

func mapError(err vk.Result) error {
	if err >= 0 {
		return nil
	}

	switch err {
	case vk.ERROR_INITIALIZATION_FAILED:
		return gfx.ErrInitializationFailed

	case vk.ERROR_LAYER_NOT_PRESENT:
		return fmt.Errorf("%w: layer", gfx.ErrMissingFeature)

	case vk.ERROR_EXTENSION_NOT_PRESENT:
		return fmt.Errorf("%w: extension", gfx.ErrMissingFeature)

	case vk.ERROR_INCOMPATIBLE_DRIVER:
		return gfx.ErrIncompatibleDriver

	default:
		return fmt.Errorf("%w: vulkan %v", gfx.ErrUnexpectedStatus, err)
	}
}

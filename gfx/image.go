package gfx

/*
#include "vulkan.h"
*/
import "C"

type Image struct {
	image      C.VkImage
	view       C.VkImageView
	allocation C.VmaAllocation
	width      int
	height     int
}

type ImageType int

const (
	ImageType1D ImageType = C.VK_IMAGE_TYPE_1D
	ImageType2D ImageType = C.VK_IMAGE_TYPE_2D
	ImageType3D ImageType = C.VK_IMAGE_TYPE_3D
)

type ImageUsage int

const (
	ImageUsageAttachment ImageUsage = 1 << iota
	ImageUsageSampled
	ImageUsageCopySrc
	ImageUsageCopyDst
)

type ImageDescriptor struct {
	Type   ImageType
	Width  int
	Height int
	Depth  int
	Format Format
	Usage  ImageUsage
}

func (g *Graphics) CreateImage(des ImageDescriptor) (*Image, error) {
	switch des.Type {
	case ImageType1D:
		des.Height = 1
		des.Depth = 1
	case ImageType2D:
		des.Depth = 1
	}

	if des.Width <= 0 || des.Height <= 0 || des.Depth <= 0 {
		return nil, ErrInvalidDescriptor
	}

	var imgInfo C.VkImageCreateInfo
	imgInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
	imgInfo.imageType = C.VkImageType(des.Type)
	imgInfo.format = ToFormat(des.Format)
	imgInfo.extent.width = C.uint32_t(des.Width)
	imgInfo.extent.height = C.uint32_t(des.Height)
	imgInfo.extent.depth = C.uint32_t(des.Depth)
	imgInfo.mipLevels = 1
	imgInfo.arrayLayers = 1
	imgInfo.samples = C.VK_SAMPLE_COUNT_1_BIT
	imgInfo.tiling = C.VK_IMAGE_TILING_OPTIMAL

	if des.Usage&ImageUsageCopySrc != 0 {
		imgInfo.usage |= C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	}

	if des.Usage&ImageUsageCopyDst != 0 {
		imgInfo.usage |= C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	}

	if des.Usage&ImageUsageSampled != 0 {
		imgInfo.usage |= C.VK_IMAGE_USAGE_SAMPLED_BIT
	}

	if des.Usage&ImageUsageAttachment != 0 {
		if des.Format.Depth() {
			imgInfo.usage |= C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
		} else {
			imgInfo.usage |= C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
		}
	}

	var allocCreateInfo C.VmaAllocationCreateInfo
	allocCreateInfo.usage = C.VMA_MEMORY_USAGE_AUTO

	// TODO: Flags
	//allocCreateInfo.flags = VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT
	//allocCreateInfo.flags = VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT | VMA_ALLOCATION_CREATE_MAPPED_BIT
	//allocCreateInfo.flags = VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT | VMA_ALLOCATION_CREATE_MAPPED_BIT

	var (
		img   C.VkImage
		alloc C.VmaAllocation
	)

	if err := mapError(C.vmaCreateImage(g.memoryAllocator, &imgInfo, &allocCreateInfo, &img, &alloc, nil)); err != nil {
		return nil, err
	}

	var viewInfo C.VkImageViewCreateInfo
	viewInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
	viewInfo.image = img

	switch des.Type {
	case ImageType1D:
		viewInfo.viewType = C.VK_IMAGE_VIEW_TYPE_1D
	case ImageType2D:
		viewInfo.viewType = C.VK_IMAGE_VIEW_TYPE_2D
	case ImageType3D:
		viewInfo.viewType = C.VK_IMAGE_VIEW_TYPE_3D
	default:
		panic("unexpected image type")
	}

	viewInfo.format = imgInfo.format
	viewInfo.subresourceRange.baseMipLevel = 0
	viewInfo.subresourceRange.levelCount = 1
	viewInfo.subresourceRange.baseArrayLayer = 0
	viewInfo.subresourceRange.layerCount = 1

	if des.Format.Depth() {
		viewInfo.subresourceRange.aspectMask |= C.VK_IMAGE_ASPECT_DEPTH_BIT
	} else {
		viewInfo.subresourceRange.aspectMask |= C.VK_IMAGE_ASPECT_COLOR_BIT
	}

	var view C.VkImageView

	if err := mapError(C.vkCreateImageView(g.device, &viewInfo, nil, &view)); err != nil {
		return nil, err
	}

	return &Image{
		image:      img,
		view:       view,
		allocation: alloc,
		width:      des.Width,
		height:     des.Height,
	}, nil
}

func (i *Image) ImageView() *ImageView {
	return &ImageView{
		view: i.view,
	}
}

type ImageView struct {
	view C.VkImageView
}

type ImageViewer interface {
	ImageView() *ImageView
}

type Format int

func (f Format) Depth() bool {
	// TODO
	return false
}

const (
	FormatBGRA8UNorm Format = iota
	FormatRGBA8UNorm
	FormatRGBA16SFloat
	FormatRGB32SFloat
	FormatRG32SFloat
)

func ToFormat(format Format) C.VkFormat {
	switch format {
	case FormatBGRA8UNorm:
		return C.VK_FORMAT_B8G8R8A8_UNORM
	case FormatRGBA8UNorm:
		return C.VK_FORMAT_R8G8B8A8_UNORM
	case FormatRGBA16SFloat:
		return C.VK_FORMAT_R16G16B16A16_SFLOAT
	case FormatRGB32SFloat:
		return C.VK_FORMAT_R32G32B32_SFLOAT
	case FormatRG32SFloat:
		return C.VK_FORMAT_R32G32_SFLOAT
	default:
		panic("unknown format")
	}
}

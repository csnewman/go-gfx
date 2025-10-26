package vulkan

import (
	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
	"github.com/csnewman/go-gfx/vma"
)

type Image struct {
	image      vk.Image
	viewID     uint32
	view       vk.ImageView
	allocation vma.Allocation
	width      int
	height     int
}

func (g *Graphics) CreateImage(des gfx.ImageDescriptor) (gfx.Image, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	switch des.Type {
	case gfx.ImageType1D:
		des.Height = 1
		des.Depth = 1
	case gfx.ImageType2D:
		des.Depth = 1
	default:
		panic("unsupported image type")
	}

	if des.Width <= 0 || des.Height <= 0 || des.Depth <= 0 {
		return nil, gfx.ErrInvalidDescriptor
	}

	imgInfo := vk.ImageCreateInfoAlloc(arena, 1)
	imgInfo.SetSType(vk.STRUCTURE_TYPE_IMAGE_CREATE_INFO)
	imgInfo.SetImageType(vk.ImageType(des.Type))
	imgInfo.SetFormat(ToFormat(des.Format))

	extent := imgInfo.RefExtent()
	extent.SetWidth(uint32(des.Width))
	extent.SetHeight(uint32(des.Height))
	extent.SetDepth(uint32(des.Depth))

	imgInfo.SetMipLevels(1)
	imgInfo.SetArrayLayers(1)
	imgInfo.SetSamples(vk.SAMPLE_COUNT_1_BIT)
	imgInfo.SetTiling(vk.IMAGE_TILING_OPTIMAL)

	var usage vk.ImageUsageFlags

	if des.Usage&gfx.ImageUsageCopySrc != 0 {
		usage |= vk.IMAGE_USAGE_TRANSFER_SRC_BIT
	}

	if des.Usage&gfx.ImageUsageCopyDst != 0 {
		usage |= vk.IMAGE_USAGE_TRANSFER_DST_BIT
	}

	if des.Usage&gfx.ImageUsageSampled != 0 {
		usage |= vk.IMAGE_USAGE_SAMPLED_BIT
	}

	if des.Usage&gfx.ImageUsageAttachment != 0 {
		if des.Format.Depth() {
			usage |= vk.IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
		} else {
			usage |= vk.IMAGE_USAGE_COLOR_ATTACHMENT_BIT
		}
	}

	imgInfo.SetUsage(usage)

	allocCreateInfo := vma.AllocationCreateInfoAlloc(arena, 1)
	allocCreateInfo.SetUsage(vma.MEMORY_USAGE_AUTO)

	// TODO: Flags
	//allocCreateInfo.flags = VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT
	//allocCreateInfo.flags = VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT | VMA_ALLOCATION_CREATE_MAPPED_BIT
	//allocCreateInfo.flags = VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT | VMA_ALLOCATION_CREATE_MAPPED_BIT

	imgRef := ffi.RefAlloc[vk.Image](arena, 1)
	allocRef := ffi.RefAlloc[vma.Allocation](arena, 1)

	if err := mapError(vma.CreateImage(g.memoryAllocator, imgInfo, allocCreateInfo, imgRef, allocRef, vma.AllocationInfoNil)); err != nil {
		return nil, err
	}

	img := imgRef.Get()
	alloc := allocRef.Get()

	viewInfo := vk.ImageViewCreateInfoAlloc(arena, 1)
	viewInfo.SetSType(vk.STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO)
	viewInfo.SetImage(img)

	switch des.Type {
	case gfx.ImageType1D:
		viewInfo.SetViewType(vk.IMAGE_VIEW_TYPE_1D)
	case gfx.ImageType2D:
		viewInfo.SetViewType(vk.IMAGE_VIEW_TYPE_2D)
	case gfx.ImageType3D:
		viewInfo.SetViewType(vk.IMAGE_VIEW_TYPE_3D)
	default:
		panic("unexpected image type")
	}

	viewInfo.SetFormat(imgInfo.GetFormat())

	subresourceRange := viewInfo.RefSubresourceRange()
	subresourceRange.SetBaseMipLevel(0)
	subresourceRange.SetLevelCount(1)
	subresourceRange.SetBaseArrayLayer(0)
	subresourceRange.SetLayerCount(1)

	if des.Format.Depth() {
		subresourceRange.SetAspectMask(vk.IMAGE_ASPECT_DEPTH_BIT)
	} else {
		subresourceRange.SetAspectMask(vk.IMAGE_ASPECT_COLOR_BIT)
	}

	viewRef := ffi.RefAlloc[vk.ImageView](arena, 1)

	if err := mapError(vk.CreateImageView(g.device, viewInfo, vk.AllocationCallbacksNil, viewRef)); err != nil {
		return nil, err
	}

	view := viewRef.Get()
	viewID := g.allocTexture(arena, view)

	return &Image{
		image:      img,
		viewID:     viewID,
		view:       view,
		allocation: alloc,
		width:      des.Width,
		height:     des.Height,
	}, nil
}

func (i *Image) DefaultView() gfx.ImageView {
	return &ImageView{
		id:     i.viewID,
		view:   i.view,
		width:  i.width,
		height: i.height,
	}
}

type ImageView struct {
	id     uint32
	view   vk.ImageView
	width  int
	height int
}

func (i *ImageView) ID() uint32 {
	return i.id
}

func (i *ImageView) Width() int {
	return i.width
}

func (i *ImageView) Height() int {
	return i.height
}

func ToFormat(format gfx.Format) vk.Format {
	switch format {
	case gfx.FormatBGRA8UNorm:
		return vk.FORMAT_B8G8R8A8_UNORM
	case gfx.FormatRGBA8UNorm:
		return vk.FORMAT_R8G8B8A8_UNORM
	case gfx.FormatRGBA16SFloat:
		return vk.FORMAT_R16G16B16A16_SFLOAT
	case gfx.FormatRGB32SFloat:
		return vk.FORMAT_R32G32B32_SFLOAT
	case gfx.FormatRG32SFloat:
		return vk.FORMAT_R32G32_SFLOAT
	default:
		panic("unknown format")
	}
}

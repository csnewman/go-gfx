package gfx

/*
#include "vulkan.h"
*/
import "C"

type Image struct {
	image  C.VkImage
	view   C.VkImageView
	width  int
	height int
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

const (
	FormatBGRA8UNorm Format = iota
)

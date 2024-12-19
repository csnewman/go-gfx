package gfx

/*
#include "vulkan.h"
*/
import "C"

type Texture struct {
	img C.VkImage
}

type TextureView struct {
	view C.VkImageView
}

type TextureViewable interface {
	TextureView() *TextureView
}

type TextureFormat int

const (
	TextureFormatBGRA8UNorm TextureFormat = iota
)

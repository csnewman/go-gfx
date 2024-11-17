package spvc

import (
	"encoding/binary"
	"errors"
	"unsafe"
)

/*
#cgo CXXFLAGS: -std=c++20 -Wno-deprecated-this-capture -DSPIRV_CROSS_C_API_GLSL -DSPIRV_CROSS_C_API_MSL

#include "spirv_cross_c.h"

void gfx_spvc_error_callback(void*, const char*);

void gfx_spvc_set_errror_callback(spvc_context context, void *userdata) {
	spvc_context_set_error_callback(context, &gfx_spvc_error_callback, userdata);
}

*/
import "C"

type Backend = C.spvc_backend

var BackendMSL Backend = C.SPVC_BACKEND_MSL

type ExecutionModel = C.SpvExecutionModel

var (
	ExecutionModelVertex   ExecutionModel = C.SpvExecutionModelVertex
	ExecutionModelFragment ExecutionModel = C.SpvExecutionModelFragment
)

type MSLPlatform = C.spvc_msl_platform

var (
	MSLPlatformIOS   MSLPlatform = C.SPVC_MSL_PLATFORM_IOS
	MSLPlatformMacOS MSLPlatform = C.SPVC_MSL_PLATFORM_MACOS
)

func Version() (int, int, int) {
	var (
		major C.uint
		minor C.uint
		patch C.uint
	)

	C.spvc_get_version(&major, &minor, &patch)

	return int(major), int(minor), int(patch)
}

var (
	ErrInvalidSPIRV     = errors.New("invalid SPIRV")
	ErrUnsupportedSPIRV = errors.New("unsupported SPIRV")
	ErrOutOfMemory      = errors.New("out of memory")
	ErrInvalidArgument  = errors.New("invalid argument")
)

func mapError(res C.spvc_result) error {
	switch res {
	case C.SPVC_SUCCESS:
		return nil
	case C.SPVC_ERROR_INVALID_SPIRV:
		return ErrInvalidSPIRV
	case C.SPVC_ERROR_UNSUPPORTED_SPIRV:
		return ErrUnsupportedSPIRV
	case C.SPVC_ERROR_OUT_OF_MEMORY:
		return ErrOutOfMemory
	case C.SPVC_ERROR_INVALID_ARGUMENT:
		return ErrInvalidArgument
	default:
		panic("unexpected error")
	}
}

type Context struct {
	context C.spvc_context
}

func NewContext() (*Context, error) {
	var context C.spvc_context

	if err := mapError(C.spvc_context_create(&context)); err != nil {
		return nil, err
	}

	C.gfx_spvc_set_errror_callback(context, nil)

	return &Context{context: context}, nil
}

func (c *Context) Close() {
	C.spvc_context_destroy(c.context)
}

type ParsedIR struct {
	context *Context
	parsed  C.spvc_parsed_ir
}

func (c *Context) ParseSPIRV(data []byte) (*ParsedIR, error) {
	if len(data)%4 != 0 {
		return nil, ErrInvalidSPIRV
	}

	spv := make([]C.SpvId, len(data)/4)

	for i := 0; i < len(spv); i++ {
		spv[i] = C.SpvId(binary.LittleEndian.Uint32(data[i*4:]))
	}

	var parsed C.spvc_parsed_ir

	if err := mapError(C.spvc_context_parse_spirv(
		c.context,
		unsafe.SliceData(spv),
		C.size_t(len(spv)),
		&parsed,
	)); err != nil {
		return nil, err
	}

	return &ParsedIR{
		context: c,
		parsed:  parsed,
	}, nil
}

type Compiler struct {
	compiler C.spvc_compiler
	opts     C.spvc_compiler_options
}

func (i *ParsedIR) CreateCompiler(backend Backend) (*Compiler, error) {
	var compiler C.spvc_compiler

	if err := mapError(C.spvc_context_create_compiler(
		i.context.context,
		backend,
		i.parsed,
		C.SPVC_CAPTURE_MODE_COPY,
		&compiler,
	)); err != nil {
		return nil, err
	}

	var opts C.spvc_compiler_options

	if err := mapError(C.spvc_compiler_create_compiler_options(
		compiler,
		&opts,
	)); err != nil {
		return nil, err
	}

	return &Compiler{
		compiler: compiler,
		opts:     opts,
	}, nil
}

type EntryPoint struct {
	Name  string
	cName *C.char
	Model ExecutionModel
}

func (c *Compiler) EntryPoints() ([]EntryPoint, error) {
	var (
		points *C.spvc_entry_point
		count  C.size_t
	)

	if err := mapError(C.spvc_compiler_get_entry_points(
		c.compiler,
		&points,
		&count,
	)); err != nil {
		return nil, err
	}

	eps := make([]EntryPoint, count)

	pointsSlice := unsafe.Slice(points, int(count))

	for i, point := range pointsSlice {
		eps[i] = EntryPoint{
			Name:  C.GoString(point.name),
			cName: point.name,
			Model: point.execution_model,
		}
	}

	return eps, nil
}

func (c *Compiler) SetEntryPoint(point EntryPoint) error {
	return mapError(C.spvc_compiler_set_entry_point(c.compiler, point.cName, point.Model))
}

func (c *Compiler) GetCleanedEntryPoint(point EntryPoint) string {
	return C.GoString(C.spvc_compiler_get_cleansed_entry_point_name(c.compiler, point.cName, point.Model))
}

func (c *Compiler) setOpt(opt C.spvc_compiler_option, value C.uint) {
	if err := mapError(C.spvc_compiler_options_set_uint(c.opts, opt, value)); err != nil {
		// This should not happen
		panic(err)
	}
}

func (c *Compiler) SetMSLPlatform(value MSLPlatform) {
	c.setOpt(C.SPVC_COMPILER_OPTION_MSL_PLATFORM, C.uint(value))
}

func (c *Compiler) Compile() (string, error) {
	if err := mapError(C.spvc_compiler_install_compiler_options(c.compiler, c.opts)); err != nil {
		return "", err
	}

	var src *C.char

	if err := mapError(C.spvc_compiler_compile(
		c.compiler,
		&src,
	)); err != nil {
		return "", err
	}

	return C.GoString(src), nil
}

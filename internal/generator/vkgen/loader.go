package main

import (
	"fmt"
	"generator/repo"
	"go/token"
	"maps"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
)

const vkPath = "github.com/csnewman/go-gfx/vk"

var globalCommands = map[string]struct{}{
	"vkEnumerateInstanceVersion":             {},
	"vkEnumerateInstanceExtensionProperties": {},
	"vkEnumerateInstanceLayerProperties":     {},
	"vkCreateInstance":                       {},
}

func generateLoader(r *repo.Repo) {
	out := jen.NewFile("vkloader")
	out.CgoPreamble(`#define VK_NO_PROTOTYPES 1
#include "../vk/vulkan.h"`)

	out.Comment("Load attempts to load all global vulkan functions.")
	out.Func().Id("Load").Params(jen.Id("loader").Qual("unsafe", "Pointer")).Block(
		jen.Qual("C", "gfx_vkInit").Call(jen.Id("loader")),
	)
	out.Comment("LoadInstance attempts to load all instance vulkan functions.")
	out.Func().Id("LoadInstance").Params(jen.Id("instance").Qual(vkPath, "Instance")).Block(
		jen.Qual("C", "gfx_vkInitInstance").Call(
			jen.Qual("C", "VkInstance").Params(
				jen.Qual("unsafe", "Pointer").Params(jen.Id("instance")),
			),
		),
	)

	loadBodyGlobal := new(strings.Builder)
	loadBodyInst := new(strings.Builder)
	generatedCmds := make(map[string]struct{})

	for _, name := range slices.Sorted(maps.Keys(r.Functions)) {
		cmd := r.Functions[name]

		generateCommand(out, cmd, generatedCmds)
	}

	for _, name := range slices.Sorted(maps.Keys(generatedCmds)) {
		tgt := loadBodyInst

		if _, ok := globalCommands[name]; ok {
			tgt = loadBodyGlobal
		}

		tgt.WriteString("    ptr_")
		tgt.WriteString(name)
		tgt.WriteString(" = (PFN_")
		tgt.WriteString(name)
		tgt.WriteString(") ptr_vkGetInstanceProcAddr(context, \"")
		tgt.WriteString(name)
		tgt.WriteString("\");\n")
	}

	out.CgoPreamble(fmt.Sprintf(`PFN_vkGetInstanceProcAddr ptr_vkGetInstanceProcAddr;

void gfx_vkInit(void* loader) {
    ptr_vkGetInstanceProcAddr = (PFN_vkGetInstanceProcAddr) loader;
    void* context = NULL;

%s}

void gfx_vkInitInstance(VkInstance context) {
%s}`, loadBodyGlobal.String(), loadBodyInst.String()))

	if err := out.Save("../../vkloader/functions.gen.go"); err != nil {
		panic(err)
	}
}

func generateCommand(o *jen.File, cmd *repo.Function, generated map[string]struct{}) {
	o.Line()

	retCType := "void"
	retString := ""

	if cmd.ReturnType != nil {
		retString = "return "
		switch cmd.ReturnType.Category {
		case repo.FieldCategoryDirect:
			retCType = cmd.ReturnType.Type
		default:
			panic("unknown category: " + string(cmd.ReturnType.Category))
		}
	}

	var (
		paramsC     []string
		paramCNames []string
	)

	for _, member := range cmd.Members {
		safeName := member.Name

		if token.IsKeyword(safeName) {
			safeName += "_"
		}

		paramCNames = append(paramCNames, member.Name)

		switch member.Category {
		case repo.FieldCategoryDirect:
			paramsC = append(paramsC, member.Type+" "+member.Name)
		case repo.FieldCategoryPointer:
			paramsC = append(paramsC, member.Type+"* "+member.Name)
		case repo.FieldCategoryPointer2:
			paramsC = append(paramsC, member.Type+"** "+member.Name)
		default:
			o.CgoPreamble(fmt.Sprintf(`// // %s.%s is unuspported`, cmd.Name, member.Name))
			return
		}
	}

	o.CgoPreamble(fmt.Sprintf(`PFN_%s ptr_%s;
VKAPI_ATTR %s VKAPI_CALL %s(%s) {
	%sptr_%s(%s);
}`, cmd.Name, cmd.Name, retCType, cmd.Name, strings.Join(paramsC, ", "), retString, cmd.Name, strings.Join(paramCNames, ", ")))

	generated[cmd.Name] = struct{}{}
}

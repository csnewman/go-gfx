package cparse

import (
	"fmt"
	"generator/repo"
	"strings"
)

type ParseConfig struct {
	Path string
}

func Parse(cfg *ParseConfig) *repo.Repo {
	p := &Parser{
		repo: &repo.Repo{
			Types:     make(map[string]*repo.Type),
			Functions: make(map[string]*repo.Function),
		},
	}

	p.ParseFile("", cfg.Path)

	return p.repo
}

func convertEnumName(in string) string {
	name, ok := strings.CutPrefix(in, "Vma")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	// XXX: bit of a hack! Should properly map these.
	name = strings.ReplaceAll(name, "FlagBits", "Flags")

	return name
}

func convertBitmaskName(in string) string {
	name, ok := strings.CutPrefix(in, "Vma")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	return name
}

func convertStructName(in string) string {
	name, ok := strings.CutPrefix(in, "Vma")
	if !ok {
		panic(fmt.Sprintf("struct does not have prefix: %s", in))
	}

	return name
}

package cparse

import (
	"fmt"
	"generator/repo"
	"strings"
)

type ParseConfig struct {
	Path string

	Defines []string

	ConvertEnumName     func(item string) *ConvertEnumOp
	ConvertEnumItemName func(item string) *ConvertEnumItemOp
	ConvertStructName   func(item string) *ConvertStructOp
	ConvertFuncName     func(item string) *ConvertFuncOp
}

func Parse(cfg *ParseConfig) *repo.Repo {
	p := &Parser{
		cfg: cfg,
		repo: &repo.Repo{
			Types:     make(map[string]*repo.Type),
			Functions: make(map[string]*repo.Function),
		},
	}

	p.ParseFile("")

	return p.repo
}

type ConvertEnumOp struct {
	Name string

	Skip bool
}

func (p *Parser) convertEnumName(in string) *ConvertEnumOp {
	return p.cfg.ConvertEnumName(in)
}

type ConvertEnumItemOp struct {
	Name string

	//Skip bool
}

func (p *Parser) convertEnumItem(in string) *ConvertEnumItemOp {
	return p.cfg.ConvertEnumItemName(in)
}

func convertBitmaskName(in string) string {
	name, ok := strings.CutPrefix(in, "Vma")
	if !ok {
		panic(fmt.Sprintf("enum does not have prefix: %s", in))
	}

	return name
}

type ConvertStructOp struct {
	Name string

	Skip bool
}

func (p *Parser) convertStructName(in string) *ConvertStructOp {
	return p.cfg.ConvertStructName(in)
}

type ConvertFuncOp struct {
	Name string

	Skip bool
}

func (p *Parser) convertFuncName(in string) *ConvertFuncOp {
	return p.cfg.ConvertFuncName(in)
}

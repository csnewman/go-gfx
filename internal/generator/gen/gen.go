package gen

import (
	"fmt"
	"generator/repo"
	"log/slog"
	"maps"
	"path/filepath"
	"slices"

	"github.com/dave/jennifer/jen"
)

const ffiPath = "github.com/csnewman/go-gfx/ffi"

type Config struct {
	Repo    *repo.Repo
	PKGName string

	CPreamble string

	Path string
}

type Generator struct {
	Repo *repo.Repo

	OEnums     *jen.File
	OBitmasks  *jen.File
	OHandles   *jen.File
	OStructs   *jen.File
	OFunctions *jen.File

	StructGeneratedOffsets map[string]struct{}
}

func Generate(cfg *Config) {
	g := &Generator{
		Repo: cfg.Repo,

		OEnums:     jen.NewFile(cfg.PKGName),
		OBitmasks:  jen.NewFile(cfg.PKGName),
		OHandles:   jen.NewFile(cfg.PKGName),
		OStructs:   jen.NewFile(cfg.PKGName),
		OFunctions: jen.NewFile(cfg.PKGName),

		StructGeneratedOffsets: make(map[string]struct{}),
	}

	if cfg.CPreamble != "" {
		g.OEnums.CgoPreamble(cfg.CPreamble)
		g.OBitmasks.CgoPreamble(cfg.CPreamble)
		g.OHandles.CgoPreamble(cfg.CPreamble)
		g.OStructs.CgoPreamble(cfg.CPreamble)
		g.OFunctions.CgoPreamble(cfg.CPreamble)
	}

	for _, name := range slices.Sorted(maps.Keys(cfg.Repo.Types)) {
		ty := cfg.Repo.Types[name]

		switch ty.Category {
		case repo.TypeCategoryEnum:
			g.generateEnumType(ty)

		case repo.TypeCategoryBitmask:
			g.generateBitmaskType(ty)

		case repo.TypeCategoryHandle, repo.TypeCategoryHandleNonDispatchable:
			g.generateHandleType(ty)

		case repo.TypeCategoryStruct:
			g.generateStructType(ty)

		default:
			//todo
		}

	}

	for _, name := range slices.Sorted(maps.Keys(cfg.Repo.Functions)) {
		fun := cfg.Repo.Functions[name]

		g.generateFunction(fun)
	}

	if err := g.OEnums.Save(filepath.Join(cfg.Path, "enums.gen.go")); err != nil {
		panic(err)
	}

	if err := g.OBitmasks.Save(filepath.Join(cfg.Path, "bitmasks.gen.go")); err != nil {
		panic(err)
	}

	if err := g.OHandles.Save(filepath.Join(cfg.Path, "handles.gen.go")); err != nil {
		panic(err)
	}

	if err := g.OStructs.Save(filepath.Join(cfg.Path, "structs.gen.go")); err != nil {
		panic(err)
	}

	if err := g.OFunctions.Save(filepath.Join(cfg.Path, "functions.gen.go")); err != nil {
		panic(err)
	}
}

func (g *Generator) generateBitmaskType(ty *repo.Type) {
	slog.Info("Generating bitmask", "name", ty.Name)

	g.OBitmasks.Commentf("%s wraps the bitmask %s.", ty.MappedName, ty.Name)

	var goType string

	switch ty.BitmaskWidth {
	case 32:
		goType = "int32"
	case 64:
		goType = "int64"
	default:
		panic(fmt.Sprintf("bitmask bit width %d", ty.BitmaskWidth))
	}

	g.OBitmasks.Type().Id(ty.MappedName).Id(goType)

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		g.OBitmasks.Commentf("%s wraps the bitmask %s. An alias for %s.", alias.MappedName, alias.Name, ty.MappedName)
		g.OBitmasks.Type().Id(alias.MappedName).Op("=").Id(ty.MappedName)
	}

	if len(ty.EnumValues) > 0 {
		g.OBitmasks.Const().DefsFunc(func(group *jen.Group) {
			for _, valueName := range slices.Sorted(maps.Keys(ty.EnumValues)) {
				value := ty.EnumValues[valueName]

				group.Commentf("%s wraps %s.", value.MappedName, value.Name)
				group.Id(value.MappedName).Id(ty.MappedName).Op("=").Qual("C", value.Name)

				for _, aliasName := range slices.Sorted(maps.Keys(value.Aliases)) {
					alias := value.Aliases[aliasName]

					group.Commentf("%s wraps the bitmask %s. An alias for %s.", alias.MappedName, alias.Name, ty.MappedName)
					group.Id(alias.MappedName).Op("=").Id(value.MappedName)
				}
			}
		})
	}

	g.OBitmasks.Func().
		ParamsFunc(func(group *jen.Group) {
			group.Id("e").Id(ty.MappedName)
		}).Id("String").
		Params().
		Id("string").
		BlockFunc(func(group *jen.Group) {
			group.Return().
				Qual("fmt", "Sprintf").
				Call(
					jen.Lit(fmt.Sprintf("%s(%%b)", ty.Name)),
					jen.Id("e"),
				)
		})
	g.OBitmasks.Line()
}

func (g *Generator) generateEnumType(ty *repo.Type) {
	slog.Info("Generating enum", "name", ty.Name)

	g.OEnums.Commentf("%s wraps the enum %s.", ty.MappedName, ty.Name)
	g.OEnums.Type().Id(ty.MappedName).Id("int32")

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		if alias.NoGeneration {
			continue
		}

		g.OEnums.Commentf("%s wraps the enum %s. An alias for %s.", alias.MappedName, alias.Name, ty.MappedName)
		g.OEnums.Type().Id(alias.MappedName).Op("=").Id(ty.MappedName)
	}

	if len(ty.EnumValues) > 0 {
		g.OEnums.Const().DefsFunc(func(group *jen.Group) {
			for _, valueName := range slices.Sorted(maps.Keys(ty.EnumValues)) {
				value := ty.EnumValues[valueName]

				group.Commentf("%s wraps %s.", value.MappedName, value.Name)
				group.Id(value.MappedName).Id(ty.MappedName).Op("=").Qual("C", value.Name)

				for _, aliasName := range slices.Sorted(maps.Keys(value.Aliases)) {
					alias := value.Aliases[aliasName]

					if alias.NoGeneration {
						continue
					}

					group.Commentf("%s wraps the enum %s. An alias for %s.", alias.MappedName, alias.Name, ty.MappedName)
					group.Id(alias.MappedName).Op("=").Id(value.MappedName)
				}
			}
		})
	}

	g.OEnums.Func().
		ParamsFunc(func(group *jen.Group) {
			group.Id("e").Id(ty.MappedName)
		}).Id("String").
		Params().
		Id("string").
		BlockFunc(func(group *jen.Group) {
			group.Return().
				Qual("fmt", "Sprintf").
				Call(
					jen.Lit(fmt.Sprintf("%s(%%b)", ty.Name)),
					jen.Id("e"),
				)
		})
	g.OEnums.Line()
}

func (g *Generator) generateHandleType(ty *repo.Type) {
	slog.Info("Generating handle", "name", ty.Name)

	cType := "uintptr"

	if ty.Category == repo.TypeCategoryHandleNonDispatchable {
		cType = "uint64"
	}

	g.OHandles.Commentf("%s wraps the handle %s.", ty.MappedName, ty.Name)
	g.OHandles.Type().Id(ty.MappedName).Id(cType)

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		if alias.NoGeneration {
			continue
		}

		g.OHandles.Commentf("%s wraps the handle %s. An alias for %s.", alias.MappedName, alias.Name, ty.MappedName)
		g.OHandles.Type().Id(alias.MappedName).Op("=").Id(ty.MappedName)
	}

	g.OHandles.Commentf("%sNil is a null pointer.", ty.MappedName)
	g.OHandles.Var().Id(ty.MappedName + "Nil").Id(ty.MappedName)

	for _, name := range slices.Sorted(maps.Keys(ty.Aliases)) {
		alias := ty.Aliases[name]

		if alias.NoGeneration {
			continue
		}

		g.OHandles.Commentf("%sNil is a null pointer.", alias.MappedName)
		g.OHandles.Var().Id(alias.MappedName + "Nil").Id(alias.MappedName)
	}
}

package repo

import (
	"encoding/json"
	"os"
)

type TypeCategory string

const (
	TypeCategoryDirect                TypeCategory = "direct"
	TypeCategoryEnum                  TypeCategory = "enum"
	TypeCategoryBitmask               TypeCategory = "bitmask"
	TypeCategoryStruct                TypeCategory = "struct"
	TypeCategoryHandle                TypeCategory = "handle"
	TypeCategoryHandleNonDispatchable TypeCategory = "handle-non-dispatchable"
)

type Type struct {
	Name       string            `json:"name"`
	MappedName string            `json:"mapped_name"`
	Category   TypeCategory      `json:"category"`
	Aliases    map[string]*Alias `json:"aliases,omitempty"`
	Comment    string            `json:"comment,omitempty"`

	// for TypeCategoryBitmask
	BitmaskWidth int `json:"bitmask_width,omitempty"`

	// for TypeCategoryBitmask & TypeCategoryEnum
	EnumValues map[string]*EnumValue `json:"enum_values,omitempty"`

	// for TypeCategoryStruct
	StructFields map[string]*Field `json:"struct_fields,omitempty"`
}

type EnumValueCategory string

type EnumValue struct {
	Name       string `json:"name"`
	MappedName string `json:"mapped_name"`
	Comment    string `json:"comment,omitempty"`

	Aliases map[string]*Alias `json:"aliases,omitempty"`
}

type Alias struct {
	Name         string `json:"name"`
	MappedName   string `json:"mapped_name"`
	NoGeneration bool   `json:"no_generation,omitempty"`
}

type FieldCategory string

const (
	FieldCategoryDirect      FieldCategory = "direct"
	FieldCategoryPointer     FieldCategory = "pointer"
	FieldCategoryPointer2    FieldCategory = "pointer2"
	FieldCategoryArray       FieldCategory = "array"
	FieldCategoryUnsupported FieldCategory = "unsupported"
)

type Field struct {
	Name       string        `json:"name,omitempty"`
	MappedName string        `json:"mapped_name,omitempty"`
	Category   FieldCategory `json:"category"`
	Type       string        `json:"type"`
	Comment    string        `json:"comment,omitempty"`

	// For FieldCategoryArray
	ArraySize int `json:"array_size,omitempty"`
}

type Function struct {
	Name       string            `jen:"name"`
	MappedName string            `jen:"mapped_name"`
	Aliases    map[string]*Alias `jen:"aliases,omitempty"`

	ReturnType *Field   `jen:"return_type,omitempty"`
	Members    []*Field `jen:"members,omitempty"`
	Comment    string   `jen:"comment,omitempty"`
}

type Repo struct {
	Types   map[string]*Type `json:"types"`
	Aliases map[string]*Type `json:"-"`

	Functions       map[string]*Function `json:"functions"`
	FunctionAliases map[string]*Function `json:"-"`
}

func Load(path string) *Repo {
	var out Repo

	enc, err := os.ReadFile(path)
	if err != nil {
		panic(err)

	}

	if err := json.Unmarshal(enc, &out); err != nil {
		panic(err)
	}

	out.PostProcess()

	return &out
}

func (r *Repo) PostProcess() {
	r.Aliases = make(map[string]*Type)

	for _, t := range r.Types {
		for _, alias := range t.Aliases {
			r.Aliases[alias.Name] = t
		}
	}

	r.FunctionAliases = make(map[string]*Function)

	for _, f := range r.Functions {
		for _, alias := range f.Aliases {
			r.FunctionAliases[alias.Name] = f
		}
	}
}

func (r *Repo) LookupType(name string) (*Type, bool) {
	if mapped, ok := r.Aliases[name]; ok {
		return mapped, true
	}

	if mapped, ok := r.Types[name]; ok {
		return mapped, true
	}

	return nil, false
}

func (r *Repo) Write(path string) {
	encRepo, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, encRepo, 0666); err != nil {
		panic(err)
	}
}

package main

type Registry struct {
	Enums    map[string]*Enum
	Types    map[string]*Type
	Aliases  map[string]string
	Features []*Feature
}

type Category string

const (
	CategoryEnum    Category = "enum"
	CategoryBitmask Category = "bitmask"
	CategoryStruct  Category = "struct"
)

type Type struct {
	Name     string
	Category Category
	Aliases  []string

	Feature string

	// Bitmask fields
	BitmaskWidth int

	// Struct fields
	StructReadOnly bool
}

type Enum struct {
	Type     string
	Name     string
	BitWidth int
	Comment  string

	Values []*EnumValue
}

type EnumValue struct {
	Name       string
	BitPos     int
	Value      string
	Alias      string
	Deprecated string
	Comment    string
}

type Feature struct {
	Version  string
	Types    map[string]struct{}
	Commands map[string]struct{}

	EnumExtensions []FeatureEnumExtension
}

type FeatureEnumExtension struct {
	Name    string
	Extends string
	Type    string

	// type=offset
	Offset int
	Ext    int

	// type=bitpos
	Bitpos int

	// type=alias
	Alias string

	// type=value
	Value string
}

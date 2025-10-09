package main

type Registry struct {
	Enums    map[string]*Enum
	Types    map[string]*Type
	Aliases  map[string]string
	Features []*Feature

	Commands map[string]*Command
}

type Category string

const (
	CategoryEnum    Category = "enum"
	CategoryBitmask Category = "bitmask"
	CategoryStruct  Category = "struct"
	CategoryHandle  Category = "handle"
)

type Type struct {
	Name     string
	Category Category
	Aliases  []string

	Feature string

	// Bitmask fields
	BitmaskWidth int
	Requires     string

	// Struct fields
	StructReadOnly bool
	Fields         []StructField

	// Handle fields
	NonDispatchable bool
}

type FieldCategory string

const (
	FieldCategoryDirect      FieldCategory = "direct"
	FieldCategoryPointer     FieldCategory = "pointer"
	FieldCategoryPointer2    FieldCategory = "pointer2"
	FieldCategoryUnsupported FieldCategory = "unsupported"
)

type StructField struct {
	Name     string
	Type     string
	Category FieldCategory
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
	Value      string
	Alias      string
	Deprecated string
	Comment    string

	BitPos    int
	HasBitPos bool
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

type MemberCategory string

const (
	MemberCategoryDirect      MemberCategory = "direct"
	MemberCategoryPointer     MemberCategory = "pointer"
	MemberCategoryPointer2    MemberCategory = "pointer2"
	MemberCategoryUnsupported MemberCategory = "unsupported"
)

type Command struct {
	Name    string
	Aliases []string
	Feature string

	ReturnType string

	Members []CommandMember
}

type CommandMember struct {
	Name     string
	Type     string
	Category MemberCategory
}

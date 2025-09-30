package main

type Registry struct {
	Enums   map[string]*Enum
	Types   map[string]*Type
	Aliases map[string]string
}

type Category string

const (
	CategoryEnum    Category = "enum"
	CategoryBitmask Category = "bitmask"
)

type Type struct {
	Name     string
	Category Category
	Aliases  []string

	// Bitmask fields
	BitmaskWidth int
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

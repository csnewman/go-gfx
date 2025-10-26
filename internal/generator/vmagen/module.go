package main

type Module struct {
	Enums     map[string]*Enum
	Handles   map[string]*Handle
	Structs   map[string]*Struct
	Functions map[string]*Function
}

type Enum struct {
	Name      string
	Constants []*Constant
	Comment   string
}

type Constant struct {
	Name    string
	Comment string
}

type Handle struct {
	Name            string
	NonDispatchable bool
}

type Struct struct {
	Name    string
	Fields  []*Field
	Comment string
}

type Field struct {
	Name    string
	Type    *Type
	Comment string
}

type TypeCategory string

const (
	TypeCategoryDirect   TypeCategory = "direct"
	TypeCategoryPointer  TypeCategory = "pointer"
	TypeCategoryPointer2 TypeCategory = "pointer2"
	TypeCategoryArray    TypeCategory = "array"
)

type Type struct {
	Name     string
	Category TypeCategory

	ArraySize int
}

type Function struct {
	Name       string
	Args       []*Param
	ReturnType *Type
	//Variadic bool
	//Ptr      bool
	Comment string
}

type Param struct {
	Name string
	Type *Type
}

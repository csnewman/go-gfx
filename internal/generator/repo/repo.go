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
	Name       string       `json:"name"`
	MappedName string       `json:"mapped_name"`
	Category   TypeCategory `json:"category"`
}

type Repo struct {
	Types map[string]*Type `json:"types"`
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

	return &out
}

package p

import (
	"github.com/nxpkg/go-jsonschema/jsonschema"
)

type MetaSchemaRefs struct {
	A *jsonschema.Schema `json:"a,omitempty"`
	B *jsonschema.Schema `json:"b,omitempty"`
}

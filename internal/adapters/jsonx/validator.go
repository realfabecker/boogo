package jsonx

import (
	"errors"
	"fmt"
	"github.com/realfabecker/bogo/internal/core/ports"
	"github.com/xeipuuv/gojsonschema"
)

type Validator struct{}

func NewValidator() ports.JsonValidator {
	return &Validator{}
}

func (v Validator) Validate(data []byte, schema string) (bool, error) {
	loader := gojsonschema.NewStringLoader(schema)

	s := gojsonschema.NewSchemaLoader()
	compiled, err := s.Compile(loader)
	if err != nil {
		return false, fmt.Errorf("compile: %w", err)
	}

	document := gojsonschema.NewStringLoader(string(data))
	result, err := compiled.Validate(document)
	if err != nil {
		return false, fmt.Errorf("validate: %w", err)
	}

	if !result.Valid() {
		var e string
		for _, v := range result.Errors() {
			e = e + v.String() + ";"
		}
		return false, errors.New(e)
	}

	return true, nil
}

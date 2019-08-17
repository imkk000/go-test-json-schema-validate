package check

import "github.com/xeipuuv/gojsonschema"

type SchemaCheckUtil struct{}

func (SchemaCheckUtil) Validate(schemaLoader, documentLoader gojsonschema.JSONLoader) (bool, []gojsonschema.ResultError) {
	schemaResult, _ := gojsonschema.Validate(schemaLoader, documentLoader)
	if schemaResult == nil {
		return false, nil
	}
	return schemaResult.Valid(), schemaResult.Errors()
}

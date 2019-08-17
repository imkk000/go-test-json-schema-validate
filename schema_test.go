package schema_test

import (
	. "schema/model"
	. "schema/util/check"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func TestJsonSchemaWithErrorPageNumbersLessThanZero(t *testing.T) {
	documentLoad := gojsonschema.NewGoLoader(MyBook{
		Name:        "",
		AuthorName:  "",
		PageNumbers: -1,
		Price:       0,
		Publisher:   "",
		CreatedDate: "2019-08-01",
	})
	valid, err := Schema.Validate(SchemaLoader(), documentLoad)
	assert.False(t, valid)
	assert.Equal(t, "PageNumbers: Must be greater than or equal to 0/1", err[0].String())
}

func TestJsonSchemaWithSuccess(t *testing.T) {
	documentLoad := gojsonschema.NewGoLoader(MyBook{
		Name:        "Name Example",
		AuthorName:  "Imkk-000",
		PageNumbers: 111,
		Price:       99.55,
		Publisher:   "Medium",
		CreatedDate: "2019-08-01",
	})
	result, err := Schema.Validate(SchemaLoader(), documentLoad)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, result)
}

const MYBOOKSCHEMAURL = "https://github.com/imkk-000/go-test-json-schema-validate/blob/master/schema/mybook.json?raw=true"

var Schema = SchemaCheckUtil{}
var SchemaLoaderBuffer gojsonschema.JSONLoader

func SchemaLoader() gojsonschema.JSONLoader {
	if SchemaLoaderBuffer == nil {
		// Once time download schema
		SchemaLoaderBuffer = gojsonschema.NewReferenceLoader(MYBOOKSCHEMAURL)
	}
	return SchemaLoaderBuffer
}

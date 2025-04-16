package models

import (
	"testing"
)

type NormalStruct struct {
	Field1 string `field:"field1"`
	Field2 string `field:"field2"`
	Field3 string `field:"field3"`
}

type StructWithPrivateFields struct {
	Field1        string `field:"field1"`
	Field2        string `field:"field2"`
	Field3        string `field:"field3"`
	privateField1 string `field:"private1"`
	privateField2 string `field:"private2"`
}

func TestCreateModel(t *testing.T) {

	t.Run("Normal Struct", func(t *testing.T) {

		model, err := CreateModel("normal-struct", &NormalStruct{})

		if err != nil {
			t.Fatalf("Creating the normal struct shouldnt return an error")
		}

		realFields := map[string]string{
			"Field1": "field1",
			"Field2": "field2",
			"Field3": "field3",
		}

		for att, field := range model.fieldMapper {
			realField, ok := realFields[att]

			if !ok {
				t.Fatalf("The attribute %v wasnt found", att)
			}

			if realField != field {
				t.Fatalf("The expected field for attribute %v got %v when %v was expected", att, field, realField)
			}

			t.Logf("Struct attribute: %v -> Field name: %v", att, field)
		}
	})

	t.Run("Struct with private fields", func(t *testing.T) {
		model, err := CreateModel("normal-struct", &NormalStruct{})

		if err != nil {
			t.Fatalf("Creating the normal struct shouldnt return an error")
		}

		realFields := map[string]string{
			"Field1": "field1",
			"Field2": "field2",
			"Field3": "field3",
		}

		privateFields := map[string]string{
			"privateField1": "private1",
			"privateField2": "private2",
		}

		for att, field := range model.fieldMapper {
			realField, ok := realFields[att]

			if !ok {
				t.Fatalf("The attribute %v wasnt found", att)
			}

			if realField != field {
				t.Fatalf("The expected field for attribute %v got %v when %v was expected", att, field, realField)
			}

			t.Logf("Struct attribute: %v -> Field name: %v", att, field)
		}

		for privateAtt, _ := range privateFields {
			_, ok := model.fieldMapper[privateAtt]

			if ok {
				t.Fatalf("Field %v should be supressed because it is private", privateAtt)
			}
		}

	})
}

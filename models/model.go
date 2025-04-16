package models

import (
	i "db/interfaces"
	"db/utils"
)

type Model[T any] struct {
	name        string
	shape       T
	fieldMapper map[string]string
	engine      *i.Engine[T]
}

func CreateModel[T any](name string, atts T) (*Model[T], error) {

	fieldsMapper, err := utils.GetStructFields(atts)

	if err != nil {
		return nil, err
	}

	model := &Model[T]{name: name, shape: atts, fieldMapper: fieldsMapper}

	return model, nil
}

func (m *Model[T]) GetName() string {
	return m.name
}

func (m *Model[T]) GetShape() any {
	return m.shape
}

func (m *Model[T]) GetFieldMapping() map[string]string {
	return m.fieldMapper
}

func (m *Model[T]) FindByPK() {
	engine := *m.engine
	engine.Dosomnt(m)
}

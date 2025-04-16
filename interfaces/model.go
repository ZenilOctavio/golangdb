package interfaces

type Model interface {
	//Meta data
	GetName() string
	GetShape() any
	GetFieldMapping() map[string]string

	//Methods
	FindByPK()
}

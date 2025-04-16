package interfaces

type Engine[T any] interface {
	AddModel(model Model)
	FindByPK(model Model)
	Dosomnt(model Model)
}

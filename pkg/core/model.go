package core

type ModelInterface interface {
	ToDBModel(entity interface{}) interface{}

	ToDomainModel(entity interface{}) interface{}
}

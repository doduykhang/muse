package repositories

type CrudRepository[T any, id any] interface {
	Create(*T) (*T, error)
	Update(id, *T) (*T, error)
	Delete(id) error
	FindById(id) (*T, error)
	FindAll() ([]T, error)
}

type CrudRepositoryImpl[T any, id any] struct{}

func (CrudRepositoryImpl *CrudRepositoryImpl[T, id]) Create(model *T) (*T, error) {
	result := db.Create(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
func (CrudRepositoryImpl *CrudRepositoryImpl[T, id]) Update(i id, model *T) (*T, error) {
	var updateModel T
	result := db.First(&updateModel, i)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Model(&updateModel).Updates(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}
func (CrudRepositoryImpl *CrudRepositoryImpl[T, id]) Delete(i id) error {
	var model T
	result := db.Delete(&model, i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (CrudRepositoryImpl *CrudRepositoryImpl[T, id]) FindById(i id) (*T, error) {
	var model T
	result := db.First(&model, i)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (CrudRepositoryImpl *CrudRepositoryImpl[T, id]) FindAll() ([]T, error) {
	var model []T
	result := db.Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

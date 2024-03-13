package collection

type CollectionRepository interface {
	Save(collection *Collection) error
	FindById(id string) (*Collection, error)
	FindBySlug(slug string) (*Collection, error)
	FindAll(userId string) ([]*Collection, error)
	FindAllByUserId(userId string) ([]*Collection, error)
	Update(collection *Collection) error
	Delete(id string) error
}

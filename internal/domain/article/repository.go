package article

type ArticleRepository interface {
	Save(article *Article) error
	FindById(id string) (*Article, error)
	FindByTitle(title string) (*Article, error)
	FindAll() ([]*Article, error)
	Update(article *Article) error
	Delete(id string) error
}

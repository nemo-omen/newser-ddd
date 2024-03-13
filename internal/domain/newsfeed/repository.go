package newsfeed

type NewsfeedRepository interface {
	Save(newsfeed *Newsfeed) error
	FindById(id string) (*Newsfeed, error)
	FindAll() ([]*Newsfeed, error)
	Update(newsfeed *Newsfeed) error
	Delete(id string) error
}

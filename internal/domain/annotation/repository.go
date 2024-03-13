package annotation

type AnnotationRepository interface {
	Save(annotation *Annotation) error
	FindById(id string) (*Annotation, error)
	FindByTitle(title string) (*Annotation, error)
	FindAll(userId string) ([]*Annotation, error)
	FindAllByUserId(userId string) ([]*Annotation, error)
	Update(annotation *Annotation) error
	Delete(id string) error
}

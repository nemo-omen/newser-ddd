package service

import (
	"newser/internal/domain/collection"
	"newser/internal/domain/value"
	"newser/internal/infra/repository/memory"
)

type CollectionConfiguration func(cs *CollectionService) error

type CollectionService struct {
	collectionRepo collection.CollectionRepository
}

func NewCollectionService(configs ...CollectionConfiguration) (*CollectionService, error) {
	service := &CollectionService{}

	for _, config := range configs {
		err := config(service)
		if err != nil {
			return nil, err
		}
	}

	return service, nil
}

func WithCollectionRepository(repo collection.CollectionRepository) CollectionConfiguration {
	return func(cs *CollectionService) error {
		cs.collectionRepo = repo
		return nil
	}
}

func WithMemoryCollectionRepository() CollectionConfiguration {
	return func(cs *CollectionService) error {
		repo := memory.NewCollectionMemoryRepo()
		cs.collectionRepo = repo
		return nil
	}
}

func (cs *CollectionService) CreateCollection(title string, user value.PersonId) (*collection.Collection, error) {
	c, err := collection.NewCollection(collection.CollectionProps{
		Title: title,
		User:  user,
	})
	if err != nil {
		return nil, err
	}

	err = cs.collectionRepo.Save(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cs *CollectionService) AddArticleToCollection(article value.ArticleId, collectionId value.CollectionId) error {
	c, err := cs.collectionRepo.FindById(collectionId.ID)
	if err != nil {
		return err
	}

	c.AddArticle(article)
	return cs.collectionRepo.Save(c)
}

func (cs *CollectionService) RemoveArticleFromCollection(article value.ArticleId, collectionId value.CollectionId) error {
	c, err := cs.collectionRepo.FindById(collectionId.ID)
	if err != nil {
		return err
	}

	c.RemoveArticle(article)
	return cs.collectionRepo.Save(c)
}

func (cs *CollectionService) FindAllCollectionsByUserId(userId value.PersonId) ([]*collection.Collection, error) {
	return cs.collectionRepo.FindAllByUserId(userId.ID)
}

func (cs *CollectionService) FindCollectionById(id value.CollectionId) (*collection.Collection, error) {
	return cs.collectionRepo.FindById(id.ID)
}

func (cs *CollectionService) FindCollectionBySlug(slug string) (*collection.Collection, error) {
	return cs.collectionRepo.FindBySlug(slug)
}

func (cs *CollectionService) DeleteCollection(id value.CollectionId) error {
	return cs.collectionRepo.Delete(id.ID)
}

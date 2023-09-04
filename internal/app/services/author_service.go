package services

import (
	"zatrano/internal/app/models"
	"zatrano/internal/app/repositories"
)

type AuthorService struct {
	repo *repositories.AuthorRepository
}

func NewAuthorService(repo *repositories.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
	return s.repo.GetAll()
}

func (s *AuthorService) GetAuthorByID(id uint) (models.Author, error) {
	return s.repo.GetByID(id)
}

func (s *AuthorService) CreateAuthor(author models.Author) (models.Author, error) {
	return s.repo.Create(author)
}

func (s *AuthorService) UpdateAuthor(author models.Author) (models.Author, error) {
	return s.repo.Update(author)
}

func (s *AuthorService) DeleteAuthor(id uint) error {
	return s.repo.Delete(id)
}

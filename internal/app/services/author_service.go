package services

import (
	"starter/helpers"
	"starter/internal/app/dtos"
	"starter/internal/app/models"
	"starter/internal/app/repositories"
	"starter/notifiers"

	"go.uber.org/zap"
)

type AuthorService struct {
	repo *repositories.AuthorRepository
}

func NewAuthorService(repo *repositories.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func createAuthorDTOToModel(input dtos.CreateAuthorDTO) models.Author {
	return models.Author{
		Name:    input.Name,
		Surname: input.Surname,
	}
}

func updateAuthorDTOToModel(input dtos.UpdateAuthorDTO, author models.Author) models.Author {
	author.Name = input.Name
	author.Surname = input.Surname
	return author
}

func (s *AuthorService) GetAllAuthors() ([]dtos.AuthorDTO, error) {
	authors, err := s.repo.GetAllAuthors()

	if err != nil {
		return nil, err
	}

	authorDTOs := make([]dtos.AuthorDTO, len(authors))
	for i, author := range authors {
		authorDTOs[i] = dtos.AuthorDTO{
			ID:      author.ID,
			Name:    author.Name,
			Surname: author.Surname,
		}
	}

	return authorDTOs, nil
}

func (s *AuthorService) GetAuthorByID(authorID uint) (*dtos.AuthorDTO, error) {
	author, err := s.repo.GetAuthorByID(authorID)
	if err != nil {
		return nil, err
	}

	authorDTO := &dtos.AuthorDTO{
		ID:      author.ID,
		Name:    author.Name,
		Surname: author.Surname,
	}

	return authorDTO, nil
}

func (s *AuthorService) CreateAuthor(input dtos.CreateAuthorDTO) (*dtos.AuthorDTO, error) {
	author := createAuthorDTOToModel(input)
	err := s.repo.CreateAuthor(&author)
	if err != nil {
		return nil, err
	}

	createdAuthor := dtos.AuthorDTO{
		ID:      author.ID,
		Name:    author.Name,
		Surname: author.Surname,
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": "example@example.com", "title": "Register", "content": "You have successfully registered."}
	es.Send("email", emailContent)

	logger := helpers.GetLogger()
	logger.Info("Author created", zap.String("author_name", helpers.FullName(createdAuthor.Name, createdAuthor.Surname)))

	return &createdAuthor, nil
}

func (s *AuthorService) UpdateAuthor(authorID uint, input dtos.UpdateAuthorDTO) (*dtos.AuthorDTO, error) {
	author, err := s.repo.GetAuthorByID(authorID)
	if err != nil {
		return nil, err
	}

	updatedAuthor := updateAuthorDTOToModel(input, *author)

	err = s.repo.UpdateAuthor(&updatedAuthor)
	if err != nil {
		return nil, err
	}

	updatedAuthorDTO := dtos.AuthorDTO{
		ID:      updatedAuthor.ID,
		Name:    updatedAuthor.Name,
		Surname: updatedAuthor.Surname,
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": "example@example.com", "title": "Register", "content": "You have successfully updated."}
	es.Send("email", emailContent)

	logger := helpers.GetLogger()
	logger.Info("Author updated", zap.String("author_name", helpers.FullName(updatedAuthorDTO.Name, updatedAuthorDTO.Surname)))

	return &updatedAuthorDTO, nil
}

func (s *AuthorService) DeleteAuthor(authorID uint) error {
	return s.repo.DeleteAuthor(authorID)
}

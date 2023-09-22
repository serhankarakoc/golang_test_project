package repositories

import (
	"starter/internal/app/models"
	"starter/internal/database"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{db: database.SetupDatabase()}
}

func (r *AuthorRepository) GetAllAuthors() ([]*models.Author, error) {
	var authors []*models.Author
	if err := r.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepository) GetAuthorByID(id uint) (*models.Author, error) {
	var author models.Author
	if err := r.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) CreateAuthor(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *AuthorRepository) UpdateAuthor(author *models.Author) error {
	return r.db.Save(author).Error
}

func (r *AuthorRepository) DeleteAuthor(id uint) error {
	return r.db.Delete(&models.Author{}, id).Error
}

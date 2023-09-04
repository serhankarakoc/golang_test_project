package repositories

import (
	"zatrano/internal/app/models"
	"zatrano/internal/database"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{db: database.SetupDatabase()}
}

func (r *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	result := r.db.Find(&authors)
	return authors, result.Error
}

func (r *AuthorRepository) GetByID(id uint) (models.Author, error) {
	var author models.Author
	result := r.db.First(&author, id)
	return author, result.Error
}

func (r *AuthorRepository) Create(author models.Author) (models.Author, error) {
	result := r.db.Create(&author)
	return author, result.Error
}

func (r *AuthorRepository) Update(author models.Author) (models.Author, error) {
	result := r.db.Save(&author)
	return author, result.Error
}

func (r *AuthorRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Author{}, id)
	return result.Error
}

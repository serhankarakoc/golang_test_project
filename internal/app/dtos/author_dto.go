package dtos

type AuthorDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
}

type CreateAuthorDTO struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
}

type UpdateAuthorDTO struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
}

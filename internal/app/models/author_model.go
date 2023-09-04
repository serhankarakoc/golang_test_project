package models

type Author struct {
	BaseModel
	Name    string `gorm:"not null"`
	Surname string `gorm:"not null"`
}

func (Author) TableName() string {
	return "authors"
}

func (a *Author) Validate() []string {
	var errs []string

	if len(a.Name) == 0 {
		errs = append(errs, "Name cannot be empty")
	}

	if len(a.Surname) == 0 {
		errs = append(errs, "Surname cannot be empty")
	}

	return errs
}

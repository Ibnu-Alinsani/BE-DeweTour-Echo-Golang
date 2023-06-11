package repository

import (
	"dumbmerch/models"
	"fmt"

	"gorm.io/gorm"
)

type CountryRepository interface {
	GetAllCountry() ([]models.Country, error)
	GetCountryById(id int) (models.Country, error)
	AddCountry(country models.Country) (models.Country, error)
	EditCountry(country models.Country) (models.Country, error)
	DeleteCountry(country models.Country) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllCountry() ([]models.Country, error) {
	var Countries []models.Country

	err := r.db.Find(&Countries).Error

	return Countries, err
}

func (r *repository) GetCountryById(id int) (models.Country, error) {
	var country models.Country

	err := r.db.First(&country, id).Error

	return country, err
}

func (r *repository) AddCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error

	return country, err
}

func (r *repository) EditCountry(country models.Country) (models.Country, error) {
	err := r.db.Save(&country).Error

	return country, err
}

func (r *repository) DeleteCountry(country models.Country) (models.Country, error) {
	err := r.db.Delete(&country).Error
	fmt.Println(err, "ini error repo")

	return country, err
}
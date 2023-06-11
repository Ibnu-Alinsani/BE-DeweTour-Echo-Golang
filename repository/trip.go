package repository

import (
	"dumbmerch/models"
	"fmt"

	"gorm.io/gorm"
)

type TripRepository interface {
	GetAllTrip() ([]models.Trip, error)
	GetTripById(id int) (models.Trip, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	GetCountryByIdForTrip(id int) (models.Country, error)
	DeleteTrip(trip models.Trip) (models.Trip, error)
	EditTrip(trip models.Trip) (models.Trip, error)
	// GetCountryId(trip models.Trip) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) GetAllTrip() ([]models.Trip, error) {
	var Trips []models.Trip

	err := r.db.Preload("Country").Find(&Trips).Error

	return Trips, err
}

func (r *repository) GetTripById(id int) (models.Trip, error) {
	var Trip models.Trip

	err := r.db.Preload("Country").First(&Trip, id).Error
	return Trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Create(&trip).Error
	fmt.Println(trip)

	return trip, err
}

func (r *repository) GetCountryByIdForTrip(id int) (models.Country, error) {
	var country models.Country

	err := r.db.First(&country, id).Error

	return country, err
}

func (r *repository) DeleteTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Delete(&trip).Error
	return trip, err
}

func (r *repository) EditTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Save(&trip).Error
	return trip, err
}
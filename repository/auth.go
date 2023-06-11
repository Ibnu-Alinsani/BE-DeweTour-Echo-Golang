package repository

import (
	"dumbmerch/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(id int) (models.User, error)
	// Logout() (interface{})
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	// err := 
	// r.db.Raw()

	var count int

	err := r.db.Raw("SELECT COUNT(*) FROM users where email = ?", user.Email).Scan(&count).Error

	if err != nil {
		return user, err
	}
	fmt.Println(count)

	if count < 1 {
		err = r.db.Create(&user).Error
	} else {
		err = errors.New("maaf mas email mu sudah dipakai, tapi boleh lah 2000 buat hapus")
	}
	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CheckAuth(id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction.Trip.Country").First(&user, id).Error

	return user, err
}
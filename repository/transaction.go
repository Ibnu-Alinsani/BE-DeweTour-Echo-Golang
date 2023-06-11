package repository

import (
	"dumbmerch/models"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionById(Id int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	EditTransaction(status string, orderId int, orderQty int) (models.Transaction, error)

	GetUserByIdForTrans(id int) (models.UserResponse, error)
	GetTripByIdForTrans(id int) (models.TripResponse, error)
	GetTripByIdForTransUpdate(id int) (models.Trip, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

// Get All Transaction
func (r *repository) GetAllTransactions() ([]models.Transaction, error) {
	var Transaction []models.Transaction

	err := r.db.Preload("User").Preload("Trip.Country").Find(&Transaction).Error

	return Transaction, err
}

// Get transaction by Id
func (r *repository) GetTransactionById(Id int) (models.Transaction, error) {
	var Transaction models.Transaction

	err := r.db.Preload("User").Preload("Trip.Country").First(&Transaction, Id).Error 
	return Transaction, err
}

// Create Transaction
func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Trip.Country").Create(&transaction).Error
	
	return transaction, err
}

// Edit Transaction
func (r *repository) EditTransaction(status string, orderId int, orderQty int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Trip.Country").First(&transaction, orderId)

	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.Preload("Country").First(&trip, transaction.Trip.Id)
		trip.CurrentQuota -= orderQty
		r.db.Save(&trip)
		fmt.Println(trip.CurrentQuota, "ini trip current")
	}

	transaction.Status = status
	fmt.Println(transaction, "ini transaction")
	err := r.db.Save(&transaction).Error

	return transaction, err
}

// Delete Transaction
func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Trip.Country").Delete(&transaction).Error
	return transaction, err
}

// Get User for transaction
func (r *repository) GetUserByIdForTrans(id int) (models.UserResponse, error) {
	var user models.UserResponse
	err := r.db.First(&user, id).Error

	return user, err
}

// Get Trip for Transaction
func (r *repository) GetTripByIdForTrans(id int) (models.TripResponse, error) {
	var Trip models.TripResponse

	err := r.db.Preload("Country").First(&Trip, id).Error
	return Trip, err
}
func (r *repository) GetTripByIdForTransUpdate(id int) (models.Trip, error) {
	var Trip models.Trip

	err := r.db.First(&Trip, id).Error
	return Trip, err
}
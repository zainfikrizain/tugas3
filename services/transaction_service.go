package services

import (
	"Tugas3/models"
	"Tugas3/repositories"
	"context"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, req models.CheckoutRequest) (*models.Transaction, error)
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(r repositories.TransactionRepository) TransactionService {
	return &transactionService{repo: r}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req models.CheckoutRequest) (*models.Transaction, error) {
	var totalAmount int
	var details []models.TransactionDetail

	// Loop items
	for _, item := range req.Items {
		product, err := s.repo.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		subtotal := product.Price * item.Quantity
		totalAmount += subtotal

		details = append(details, models.TransactionDetail{
			ProductID:   product.ID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Create transaction
	tx := &models.Transaction{
		TotalAmount: totalAmount,
		CreatedAt:   time.Now(),
	}

	// Save transaction + details
	err := s.repo.CreateTransaction(ctx, tx, details)
	if err != nil {
		return nil, err
	}

	tx.Details = details
	return tx, nil
}

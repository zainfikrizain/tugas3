package repositories

import (
	"Tugas3/models"
	"context"
	"database/sql"
)

type TransactionRepository interface {
	GetProductByID(ctx context.Context, id int) (*models.Product, error)
	CreateTransaction(ctx context.Context, tx *models.Transaction, details []models.TransactionDetail) error
}

type transcationRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) transcationRepository {
	return transcationRepository{db: db}
}

func (r *transcationRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id, name, price FROM produk WHERE id = $1`

	var p models.Product

	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

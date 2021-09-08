package repository

import "github.com/andynl/my-money/entity"

type TransactionRepository interface {
	FindAll() (transaction []entity.Transaction)
}

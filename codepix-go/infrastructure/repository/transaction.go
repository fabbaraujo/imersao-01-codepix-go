package repository

import (
	"fmt"

	"github.com/fabbaraujo/imersao-01-codepix-go/codepix-go/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (repository TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := repository.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := repository.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	repository.Db.Preload("AccountFrom.Bank").Find(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}

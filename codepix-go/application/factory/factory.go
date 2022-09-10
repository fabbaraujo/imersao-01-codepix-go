package factory

import (
	"github.com/fabbaraujo/imersao-01-codepix-go/codepix-go/application/usecase"
	"github.com/fabbaraujo/imersao-01-codepix-go/codepix-go/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixKeyRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: transactionRepository,
		PixKeyRepository:      pixKeyRepository,
	}

	return transactionUseCase
}

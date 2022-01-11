package transaction

import (
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
)

type UseCase interface {
	Create(m *model.Transaction) error
	Update(m *model.Transaction) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Transaction, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Transactions, error)
}

type Storage interface {
	Create(m *model.Transaction) error
	Update(m *model.Transaction) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Transaction, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Transactions, error)
}

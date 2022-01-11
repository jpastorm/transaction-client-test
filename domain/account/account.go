package account

import (
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
)

type UseCase interface {
	Create(m *model.Account) error
	Update(m *model.Account) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Account, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Accounts, error)
}

type Storage interface {
	Create(m *model.Account) error
	Update(m *model.Account) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Account, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Accounts, error)
}

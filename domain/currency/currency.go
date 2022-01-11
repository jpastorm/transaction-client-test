package currency

import (
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
)

type UseCase interface {
	Create(m *model.Currency) error
	Update(m *model.Currency) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Currency, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Currencies, error)
}

type Storage interface {
	Create(m *model.Currency) error
	Update(m *model.Currency) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Currency, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Currencies, error)
}

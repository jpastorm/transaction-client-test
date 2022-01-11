package client

import (
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
)

type UseCase interface {
	Create(m *model.Client) error
	Update(m *model.Client) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Client, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Clients, error)
}

type Storage interface {
	Create(m *model.Client) error
	Update(m *model.Client) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Client, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Clients, error)
}

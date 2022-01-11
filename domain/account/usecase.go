package account

import (
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/jpastorm/transaction-client-test/model"
)

var allowedFieldsForQuery = []string{
	"id", "client_id", "currency_id", "money", "created_at", "updated_at",
}

// Account implements UseCase
type Account struct {
	storage Storage
}

// New returns a new Account
func New(s Storage) Account {
	return Account{storage: s}
}

// Create creates a model.Account
func (a Account) Create(m *model.Account) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("account: %w", model.ErrNilPointer)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	err := a.storage.Create(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Update updates a model.Account by id
func (a Account) Update(m *model.Account) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("account: %w", model.ErrNilPointer)
	}

	if !m.HasID() {
		return fmt.Errorf("account: %w", model.ErrInvalidID)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	err := a.storage.Update(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

func (a Account) Transfer(m *model.Account) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("account: %w", model.ErrNilPointer)
	}

	if !m.HasID() {
		return fmt.Errorf("account: %w", model.ErrInvalidID)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	err := a.storage.Transfer(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Delete deletes a model.Account by id
func (a Account) Delete(ID uint) error {
	err := a.storage.Delete(ID)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// GetWhere returns a model.Account according to filters and sorts
func (a Account) GetWhere(specification models.FieldsSpecification) (model.Account, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Account{}, fmt.Errorf("account: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Account{}, fmt.Errorf("account: %w", err)
	}

	account, err := a.storage.GetWhere(specification)
	if err != nil {
		return model.Account{}, fmt.Errorf("account: %w", err)
	}

	return account, nil
}

// GetAllWhere returns a model.Account according to filters and sorts
func (a Account) GetAllWhere(specification models.FieldsSpecification) (model.Accounts, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("account: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("account: %w", err)
	}

	account, err := a.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("account: %w", err)
	}

	return account, nil
}

// handleStorageErr handles errors from storage layer
func handleStorageErr(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}

package transaction

import (
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/jpastorm/transaction-client-test/model"
)

var allowedFieldsForQuery = []string{
	"id", "money", "type", "account_holder", "subject", "created_at", "updated_at",
}

// Transaction implements UseCase
type Transaction struct {
	storage Storage
}

// New returns a new Transaction
func New(s Storage) Transaction {
	return Transaction{storage: s}
}

// Create creates a model.Transaction
func (t Transaction) Create(m *model.Transaction) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("transaction: %w", model.ErrNilPointer)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("transaction: %w", err)
	}

	err := t.storage.Create(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Update updates a model.Transaction by id
func (t Transaction) Update(m *model.Transaction) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("transaction: %w", model.ErrNilPointer)
	}

	if !m.HasID() {
		return fmt.Errorf("transaction: %w", model.ErrInvalidID)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("transaction: %w", err)
	}

	err := t.storage.Update(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Delete deletes a model.Transaction by id
func (t Transaction) Delete(ID uint) error {
	err := t.storage.Delete(ID)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// GetWhere returns a model.Transaction according to filters and sorts
func (t Transaction) GetWhere(specification models.FieldsSpecification) (model.Transaction, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Transaction{}, fmt.Errorf("transaction: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Transaction{}, fmt.Errorf("transaction: %w", err)
	}

	transaction, err := t.storage.GetWhere(specification)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("transaction: %w", err)
	}

	return transaction, nil
}

// GetAllWhere returns a model.Transaction according to filters and sorts
func (t Transaction) GetAllWhere(specification models.FieldsSpecification) (model.Transactions, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("transaction: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("transaction: %w", err)
	}

	transaction, err := t.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("transaction: %w", err)
	}

	return transaction, nil
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

package currency

import (
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/jpastorm/transaction-client-test/model"
)

var allowedFieldsForQuery = []string{
	"id", "name", "code", "created_at", "updated_at",
}

// Currency implements UseCase
type Currency struct {
	storage Storage
}

// New returns a new Currency
func New(s Storage) Currency {
	return Currency{storage: s}
}

// Create creates a model.Currency
func (c Currency) Create(m *model.Currency) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("currency: %w", model.ErrNilPointer)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("currency: %w", err)
	}

	err := c.storage.Create(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Update updates a model.Currency by id
func (c Currency) Update(m *model.Currency) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("currency: %w", model.ErrNilPointer)
	}

	if !m.HasID() {
		return fmt.Errorf("currency: %w", model.ErrInvalidID)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("currency: %w", err)
	}

	err := c.storage.Update(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Delete deletes a model.Currency by id
func (c Currency) Delete(ID uint) error {
	err := c.storage.Delete(ID)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// GetWhere returns a model.Currency according to filters and sorts
func (c Currency) GetWhere(specification models.FieldsSpecification) (model.Currency, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Currency{}, fmt.Errorf("currency: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Currency{}, fmt.Errorf("currency: %w", err)
	}

	currency, err := c.storage.GetWhere(specification)
	if err != nil {
		return model.Currency{}, fmt.Errorf("currency: %w", err)
	}

	return currency, nil
}

// GetAllWhere returns a model.Currency according to filters and sorts
func (c Currency) GetAllWhere(specification models.FieldsSpecification) (model.Currencies, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("currency: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("currency: %w", err)
	}

	currency, err := c.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("currency: %w", err)
	}

	return currency, nil
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

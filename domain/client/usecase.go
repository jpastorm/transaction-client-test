package client

import (
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/jpastorm/transaction-client-test/model"
)

var allowedFieldsForQuery = []string{
	"id", "name", "created_at", "updated_at",
}

// Client implements UseCase
type Client struct {
	storage Storage
}

// New returns a new Client
func New(s Storage) Client {
	return Client{storage: s}
}

// Create creates a model.Client
func (c Client) Create(m *model.Client) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("client: %w", model.ErrNilPointer)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("client: %w", err)
	}

	err := c.storage.Create(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Update updates a model.Client by id
func (c Client) Update(m *model.Client) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("client: %w", model.ErrNilPointer)
	}

	if !m.HasID() {
		return fmt.Errorf("client: %w", model.ErrInvalidID)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("client: %w", err)
	}

	err := c.storage.Update(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Delete deletes a model.Client by id
func (c Client) Delete(ID uint) error {
	err := c.storage.Delete(ID)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// GetWhere returns a model.Client according to filters and sorts
func (c Client) GetWhere(specification models.FieldsSpecification) (model.Client, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Client{}, fmt.Errorf("client: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Client{}, fmt.Errorf("client: %w", err)
	}

	client, err := c.storage.GetWhere(specification)
	if err != nil {
		return model.Client{}, fmt.Errorf("client: %w", err)
	}

	return client, nil
}

// GetAllWhere returns a model.Client according to filters and sorts
func (c Client) GetAllWhere(specification models.FieldsSpecification) (model.Clients, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("client: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("client: %w", err)
	}

	client, err := c.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("client: %w", err)
	}

	return client, nil
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

package account

import (
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/jpastorm/transaction-client-test/domain/account"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/request"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/response"
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase  account.UseCase
	response response.Responser
}

func newHandler(useCase account.UseCase, response response.Responser) handler {
	return handler{useCase: useCase, response: response}
}

// Create handles the creation of a model.Account
func (h handler) Create(c echo.Context) error {
	m := model.Account{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

// Update handles the update of a model.Account
func (h handler) Update(c echo.Context) error {
	m := model.Account{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}

	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return h.response.BindFailed(c, err)
	}
	m.ID = uint(ID)

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

// Delete handles the deleting of a model.Account
func (h handler) Delete(c echo.Context) error {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return h.response.BindFailed(c, err)
	}

	err = h.useCase.Delete(uint(ID))
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

// GetWhere handles the search of a model.Account
func (h handler) GetWhere(c echo.Context) error {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return err
	}

	filtersSpecification := models.FieldsSpecification{
		Filters:    models.Fields{{Name: "client_id", Value: ID}},
		Sorts:      models.SortFields{},
		Pagination: models.Pagination{},
	}

	accountData, err := h.useCase.GetWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(accountData))
}

// GetAllWhere handles the search of all model.Account
func (h handler) GetAllWhere(c echo.Context) error {
	filtersSpecification, err := request.GetFiltersSpecification(c)
	if err != nil {
		return err
	}

	account, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(account))
}

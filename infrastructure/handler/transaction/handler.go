package transaction

import (
	"errors"
	"fmt"
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/jpastorm/transaction-client-test/domain/account"
	"github.com/jpastorm/transaction-client-test/domain/currency"
	"github.com/jpastorm/transaction-client-test/domain/transaction"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/request"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/response"
	"github.com/jpastorm/transaction-client-test/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase  transaction.UseCase
	useCaseAccount account.UseCase
	useCaseCurrency currency.UseCase
	response response.Responser
}

func newHandler(useCase transaction.UseCase, useCaseAccount account.UseCase, useCaseCurrency currency.UseCase, response response.Responser) handler {
	return handler{useCase: useCase,useCaseAccount: useCaseAccount, useCaseCurrency: useCaseCurrency, response: response}
}

// Create handles the creation of a model.Transaction
func (h handler) Create(c echo.Context) error {
	m := model.Transaction{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}


	if !model.IsValidTransactionType(m.Type) {
		return h.response.BindFailed(c, fmt.Errorf("invalid type of transaction %s", m.Type))
	}

	accountHolder, err := h.useCaseAccount.GetWhere(models.FieldsSpecification{
		Filters:    models.Fields{{Name: "id", Value: m.AccountHolder}},
		Sorts:      models.SortFields{},
		Pagination: models.Pagination{},
	})
	if err != nil {
		return h.response.Error(c, "useCase.Create().GetWhereAccount()", err)

	}

	err = h.processTransaction(m, accountHolder)
	if err != nil {
		return h.response.BindFailed(c, fmt.Errorf("error proccesing transaction %v", err))
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}


	return c.JSON(h.response.Created(m))
}

// Update handles the update of a model.Transaction
func (h handler) Update(c echo.Context) error {
	m := model.Transaction{}

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

// Delete handles the deleting of a model.Transaction
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

// GetWhere handles the search of a model.Transaction
func (h handler) GetWhere(c echo.Context) error {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return err
	}

	filtersSpecification := models.FieldsSpecification{
		Filters:    models.Fields{{Name: "account_holder", Value: ID}},
		Sorts:      models.SortFields{},
		Pagination: models.Pagination{},
	}

	transactionData, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(transactionData))
}

// GetAllWhere handles the search of all model.Transaction
func (h handler) GetAllWhere(c echo.Context) error {
	filtersSpecification, err := request.GetFiltersSpecification(c)
	if err != nil {
		return err
	}

	transaction, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(transaction))
}

func (h handler) processTransaction(tr model.Transaction, acc model.Account)  error {
	switch tr.Type {
	case "DEPOSIT":
		if acc.Money <= 0 {
			return fmt.Errorf("deposit needs at least 1 coin")
		}
		acc.Money = acc.Money + tr.Money
		err := h.useCaseAccount.Update(&acc)
		if err != nil {
			return err
		}
	case "WITHDRAW":
		if acc.Money < tr.Money {
			return fmt.Errorf("insufficient money for transaction %.2f", tr.Money)
		}

		acc.Money = acc.Money - tr.Money
		fmt.Println(acc.Money)
		err := h.useCaseAccount.Update(&acc)
		if err != nil {
			return err
		}
	case "TRANSFER":
		if !h.isValidCurrency(tr, acc) {
			return errors.New("need the same currency for transaction")
		}

		if acc.Money < tr.Money {
			return fmt.Errorf("insufficient money for transaction %.2f", tr.Money)
		}
		acc.Money = acc.Money - tr.Money
		err := h.useCaseAccount.Update(&acc)
		if err != nil {
			return fmt.Errorf("error extracting money")
		}
		money := acc.Money + tr.Money
		err = h.useCaseAccount.Transfer(&model.Account{ID: tr.Subject, Money: money})
		if err != nil {
			return fmt.Errorf("error transfering money")
		}
	}

	return nil
}

func (h handler) isValidCurrency(tr model.Transaction, acc model.Account) bool {
	subject, err := h.useCaseAccount.GetWhere(models.FieldsSpecification{
		Filters:    models.Fields{{Name: "id", Value: tr.Subject}},
		Sorts:      models.SortFields{},
		Pagination: models.Pagination{},
	})
	if err != nil {
		return false
	}
	subjectCurrency, err := h.useCaseCurrency.GetWhere(models.FieldsSpecification{
		Filters:    models.Fields{{Name: "currency_id", Value: subject.CurrencyID}},
		Sorts:      models.SortFields{},
		Pagination: models.Pagination{},
	})
	if err != nil {
		return false
	}
	accountHolderCurrency, err := h.useCaseCurrency.GetWhere(models.FieldsSpecification{
		Filters:    models.Fields{{Name: "currency_id", Value: acc.CurrencyID}},
		Sorts:      models.SortFields{},
		Pagination: models.Pagination{},
	})
	if err != nil {
		return false
	}

	if subjectCurrency.Code != accountHolderCurrency.Code {
		return false
	}

	return true
}

package account

import (
	"database/sql"

	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"
)

const table = "account"

var fields = []string{
	"client_id",
	"currency_id",
	"money",
}

var constraints = postgres.Constraints{
	// here you will add all constraints that you want to controle, ex:
	// "users_nickname_uk":                model.ErrUsersNicknameUK,
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = "DELETE FROM " + table + " WHERE id = $1"
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
	psqlTransfer = "UPDATE account SET money = $1 WHERE id = $2"
)

// Account struct that implement the interface domain.account.Storage
type Account struct {
	db *sql.DB
}

// New returns a new Account storage
func New(db *sql.DB) Account {
	return Account{db}
}

// Create creates a model.Account
func (a Account) Create(m *model.Account) error {
	stmt, err := a.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.ClientID,
		m.CurrencyID,
		m.Money,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Update this method updates a model.Account by id
func (a Account) Update(m *model.Account) error {
	stmt, err := a.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.ClientID,
		m.CurrencyID,
		m.Money,
		m.ID,
	)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Delete deletes a model.Account by id
func (a Account) Delete(ID uint) error {
	stmt, err := a.db.Prepare(psqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// GetWhere gets an ordered model.Account with filters
func (a Account) GetWhere(specification models.FieldsSpecification) (model.Account, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := a.db.Prepare(query)
	if err != nil {
		return model.Account{}, err
	}
	defer stmt.Close()

	return a.scanRow(stmt.QueryRow(args...))
}

// GetAllWhere gets all model.Accounts with Fields
func (a Account) GetAllWhere(specification models.FieldsSpecification) (model.Accounts, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)
	query += " " + postgres.BuildSQLPagination(specification.Pagination)

	stmt, err := a.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Accounts{}
	for rows.Next() {
		m, err := a.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (a Account) scanRow(s sqlutil.RowScanner) (model.Account, error) {
	m := model.Account{}

	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.ClientID,
		&m.CurrencyID,
		&m.Money,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}


// Transfer this method updates a model.Account by id
func (a Account) Transfer(m *model.Account) error {
	stmt, err := a.db.Prepare(psqlTransfer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Money,
		m.ID,
	)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

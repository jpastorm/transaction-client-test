package currency

import (
	"database/sql"

	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"
)

const table = "currency"

var fields = []string{
	"name",
	"code",
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
)

// Currency struct that implement the interface domain.currency.Storage
type Currency struct {
	db *sql.DB
}

// New returns a new Currency storage
func New(db *sql.DB) Currency {
	return Currency{db}
}

// Create creates a model.Currency
func (c Currency) Create(m *model.Currency) error {
	stmt, err := c.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		m.Code,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Update this method updates a model.Currency by id
func (c Currency) Update(m *model.Currency) error {
	stmt, err := c.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Name,
		m.Code,
		m.ID,
	)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Delete deletes a model.Currency by id
func (c Currency) Delete(ID uint) error {
	stmt, err := c.db.Prepare(psqlDelete)
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

// GetWhere gets an ordered model.Currency with filters
func (c Currency) GetWhere(specification models.FieldsSpecification) (model.Currency, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return model.Currency{}, err
	}
	defer stmt.Close()

	return c.scanRow(stmt.QueryRow(args...))
}

// GetAllWhere gets all model.Currencys with Fields
func (c Currency) GetAllWhere(specification models.FieldsSpecification) (model.Currencies, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)
	query += " " + postgres.BuildSQLPagination(specification.Pagination)

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Currencies{}
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (c Currency) scanRow(s sqlutil.RowScanner) (model.Currency, error) {
	m := model.Currency{}

	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Code,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

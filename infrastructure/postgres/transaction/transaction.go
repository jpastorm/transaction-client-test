package transaction

import (
	"database/sql"

	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"
)

const table = "transaction"

var fields = []string{
	"money",
	"type",
	"account_holder",
	"subject",
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

// Transaction struct that implement the interface domain.transaction.Storage
type Transaction struct {
	db *sql.DB
}

// New returns a new Transaction storage
func New(db *sql.DB) Transaction {
	return Transaction{db}
}

// Create creates a model.Transaction
func (t Transaction) Create(m *model.Transaction) error {
	stmt, err := t.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Money,
		m.Type,
		m.AccountHolder,
		m.Subject,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Update this method updates a model.Transaction by id
func (t Transaction) Update(m *model.Transaction) error {
	stmt, err := t.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Money,
		m.Type,
		m.AccountHolder,
		m.Subject,
		m.ID,
	)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Delete deletes a model.Transaction by id
func (t Transaction) Delete(ID uint) error {
	stmt, err := t.db.Prepare(psqlDelete)
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

// GetWhere gets an ordered model.Transaction with filters
func (t Transaction) GetWhere(specification models.FieldsSpecification) (model.Transaction, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return model.Transaction{}, err
	}
	defer stmt.Close()

	return t.scanRow(stmt.QueryRow(args...))
}

// GetAllWhere gets all model.Transactions with Fields
func (t Transaction) GetAllWhere(specification models.FieldsSpecification) (model.Transactions, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)
	query += " " + postgres.BuildSQLPagination(specification.Pagination)

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Transactions{}
	for rows.Next() {
		m, err := t.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (t Transaction) scanRow(s sqlutil.RowScanner) (model.Transaction, error) {
	m := model.Transaction{}

	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Money,
		&m.Type,
		&m.AccountHolder,
		&m.Subject,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

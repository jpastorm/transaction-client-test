package client

import (
	"database/sql"

	"github.com/jpastorm/transaction-client-test/model"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"
)

const table = "client"

var fields = []string{
	"name",
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

// Client struct that implement the interface domain.client.Storage
type Client struct {
	db *sql.DB
}

// New returns a new Client storage
func New(db *sql.DB) Client {
	return Client{db}
}

// Create creates a model.Client
func (c Client) Create(m *model.Client) error {
	stmt, err := c.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Update this method updates a model.Client by id
func (c Client) Update(m *model.Client) error {
	stmt, err := c.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Name,
		m.ID,
	)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

// Delete deletes a model.Client by id
func (c Client) Delete(ID uint) error {
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

// GetWhere gets an ordered model.Client with filters
func (c Client) GetWhere(specification models.FieldsSpecification) (model.Client, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return model.Client{}, err
	}
	defer stmt.Close()

	return c.scanRow(stmt.QueryRow(args...))
}

// GetAllWhere gets all model.Clients with Fields
func (c Client) GetAllWhere(specification models.FieldsSpecification) (model.Clients, error) {
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

	ms := model.Clients{}
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (c Client) scanRow(s sqlutil.RowScanner) (model.Client, error) {
	m := model.Client{}

	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

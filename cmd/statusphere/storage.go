package main

import (
	"fmt"
	"time"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func dbFileName() string {
	return "statusphere.db"
}

func init() {
	err := Setup()
	if err != nil {
		panic(err)
	}
}

func Setup() error {
	conn, err := sqlite.OpenConn(dbFileName(), sqlite.OpenReadWrite+sqlite.OpenCreate)
	if err != nil {
		return err
	}
	defer conn.Close()
	creates := []string{
		CreateTableStatus,
	}
	for i, create := range creates {
		err = sqlitex.ExecuteTransient(conn, create, &sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				return nil
			},
		})
		if err != nil {
			return fmt.Errorf("%d %q failed: %s", i, create, err)
		}
	}
	return nil
}

const statusTable = "status"

const CreateTableStatus = `
CREATE TABLE IF NOT EXISTS status (
	uri VARCHAR PRIMARY KEY,
	authorDid VARCHAR NOT NULL,
	status VARCHAR NOT NULL,
	createdAt VARCHAR NOT NULL,
	updatedAt VARCHAR NOT NULL
);`

type Status struct {
	Uri       string
	AuthorDid string
	Status    string
	CreatedAt string
	UpdatedAt string
}

func (s *Status) Handle() string {
	return getHandle(s.AuthorDid)
}

func (s *Status) Created() string {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", s.CreatedAt)
	if err != nil {
		return s.CreatedAt
	}
	return t.Format("Mon Jan 02 2006")
}

func saveStatus(status *Status) error {
	conn, err := sqlite.OpenConn(dbFileName(), sqlite.OpenReadWrite)
	if err != nil {
		return err
	}
	defer conn.Close()
	insertStmt, err := conn.Prepare(
		"INSERT INTO " + statusTable + " (uri, authorDid, status, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5);")
	if err != nil {
		return err
	}
	insertStmt.SetText("$1", status.Uri)
	insertStmt.SetText("$2", status.AuthorDid)
	insertStmt.SetText("$3", status.Status)
	insertStmt.SetText("$4", status.CreatedAt)
	insertStmt.SetText("$5", status.UpdatedAt)
	_, err = insertStmt.Step()
	if err != nil {
		return err
	}
	return insertStmt.Reset()
}

func listStatus() ([]*Status, error) {
	conn, err := sqlite.OpenConn(dbFileName(), sqlite.OpenReadOnly)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	values := make([]*Status, 0)
	query := "SELECT uri, authorDid, status, createdAt, updatedAt from " + statusTable + " order by createdAt desc limit 10"
	err = sqlitex.ExecuteTransient(conn,
		query,
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				value := &Status{
					Uri:       stmt.ColumnText(0),
					AuthorDid: stmt.ColumnText(1),
					Status:    stmt.ColumnText(2),
					CreatedAt: stmt.ColumnText(3),
					UpdatedAt: stmt.ColumnText(4),
				}
				values = append(values, value)
				return nil
			},
		})
	if err != nil {
		return nil, err
	}
	return values, nil
}

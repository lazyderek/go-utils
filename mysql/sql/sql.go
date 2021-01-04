package sql

import "database/sql"

func New(driver, dail string) (*sql.DB, error) {
	db, err := sql.Open(driver, dail)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(120)
	db.SetMaxOpenConns(10000)
	return db, nil
}

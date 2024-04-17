package repository

import (
	"meli/internal/core/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateItemSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to open sqlmock database: %v", err)
	}

	defer db.Close()

	itemData := domain.Item{
		ID:          "1",
		Site:        "MLA",
		Price:       "100",
		StartTime:   "2019-12-17 04:00:34 +0000 UTC",
		Name:        "Libros Físicos",
		Description: "Peso argentino",
		Nickname:    "LIBERATE_ARG",
	}

	mock.ExpectExec(`INSERT INTO items`).
		WithArgs(itemData.ID, itemData.Site, itemData.Price, itemData.StartTime, itemData.Name, itemData.Description, itemData.Nickname).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := db.Exec(`
			INSERT INTO items(id, site, price, smart_time, name, description, nickname)
			VALUES (?, ?,?, ?,?, ?,?)`, itemData.ID, itemData.Site, itemData.Price, itemData.StartTime, itemData.Name, itemData.Description, itemData.Nickname)
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	assert.Empty(t, err)
	assert.NotEmpty(t, result)
}

func TestCreateItemFailed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to open sqlmock database: %v", err)
	}

	defer db.Close()

	itemData := domain.Item{
		ID:          "10",
		Site:        "MLM",
		Price:       "500",
		StartTime:   "2020-02-17 04:00:34 +0000 UTC",
		Name:        "Ópticas Delanteras",
		Description: "Peso argentino",
		Nickname:    "FABRICADEROUPAS2030",
	}

	mock.ExpectExec(`It is a bad query`).
		WithArgs(itemData.ID, itemData.Site, itemData.Price, itemData.StartTime, itemData.Name, itemData.Description, itemData.Nickname).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = db.Exec(`
			INSERT INTO items(id, site, price, smart_time, name, description, nickname)
			VALUES (?, ?,?, ?,?, ?,?)`, itemData.ID, itemData.Site, itemData.Price, itemData.StartTime, itemData.Name, itemData.Description, itemData.Nickname)

	assert.NotEmpty(t, err)
}

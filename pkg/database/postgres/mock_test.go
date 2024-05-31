package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockDB(t *testing.T) {
	db, mock, err := NewMockDB()

	assert.NoError(t, err, "expected no error, got %v", err)
	assert.NotNil(t, db, "expected db to be non-nil")
	assert.NotNil(t, mock, "expected mock to be non-nil")

	// Expect a call to Close() on the mock database
	mock.ExpectClose()

	// Retrieve the sql.DB from the gorm.DB and ensure there is no error
	sqlDB, sqlErr := db.DB()
	assert.NoError(t, sqlErr, "expected no error from db.DB(), got %v", sqlErr)

	// Closing the mock database connection to ensure all expectations were met
	err = sqlDB.Close()
	assert.NoError(t, err, "expected no error on db close, got %v", err)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "there were unfulfilled expectations: %v", err)
}

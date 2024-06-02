package repository

import (
	"context"
	"testing"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/model"
	"github.com/hanifmaliki/golang-boilerplate/pkg/database/postgres"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindUsers(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewUserRepository(gormDB)

	t.Run("No Query", func(t *testing.T) {
		ctx := context.Background()

		rows := sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name", "email", "company_id"}).
			AddRow(userID, "creator", "updater", "Hanif Maliki Dewanto", "hanifmaliki97@gmail.com", companyID)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."deleted_at" IS NULL`).
			WillReturnRows(rows)

		request := &model.GetUserRequest{}
		query := &pkg_model.Query{}
		data, pagination, err := repo.Find(ctx, request, query)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Nil(t, pagination)

		assert.Equal(t, userID, data[0].ID)
		assert.Equal(t, "creator", data[0].CreatedBy)
		assert.Equal(t, "updater", data[0].UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data[0].Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data[0].Email)
		assert.Equal(t, companyID, data[0].CompanyID)

		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("With Query", func(t *testing.T) {
		ctx := context.Background()

		rows := sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name", "email", "company_id"}).
			AddRow(userID, "creator", "updater", "Hanif Maliki Dewanto", "hanifmaliki97@gmail.com", companyID).
			AddRow(2, "creator", "updater", "Hanif Maliki Dewanto", "hanifmaliki97@gmail.com", 3)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."deleted_at" IS NULL`).
			WillReturnRows(rows)

		rows = sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name"}).
			AddRow(companyID, "creator", "updater", "Petrosea").
			AddRow(3, "creator", "updater", "Tripatra")
		mock.ExpectQuery(`SELECT \* FROM "companies" WHERE "companies"."id" IN \(\$1,\$2\) AND "companies"."deleted_at" IS NULL`).
			WithArgs(companyID, 3).WillReturnRows(rows)

		rows = sqlmock.NewRows([]string{"id", "created_by", "updated_by", "user_id", "number"}).
			AddRow(ccID, "creator", "updater", userID, "CC2014060101").
			AddRow(uint(11), "creator", "updater", userID, "CC2014060102")
		mock.ExpectQuery(`SELECT \* FROM "credit_cards" WHERE "credit_cards"."user_id" IN \(\$1,\$2\) AND "credit_cards"."deleted_at" IS NULL`).
			WithArgs(userID, 2).WillReturnRows(rows)

		request := &model.GetUserRequest{}
		query := &pkg_model.Query{
			SortBy: "id desc",
			Expand: []string{"Company", "CreditCards"},
		}
		data, pagination, err := repo.Find(ctx, request, query)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Nil(t, pagination)

		assert.Equal(t, userID, data[0].ID)
		assert.Equal(t, "creator", data[0].CreatedBy)
		assert.Equal(t, "updater", data[0].UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data[0].Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data[0].Email)
		assert.Equal(t, companyID, data[0].CompanyID)

		assert.Equal(t, companyID, data[0].Company.ID)
		assert.Equal(t, "creator", data[0].Company.CreatedBy)
		assert.Equal(t, "updater", data[0].Company.UpdatedBy)
		assert.Equal(t, "Petrosea", data[0].Company.Name)

		assert.Equal(t, ccID, data[0].CreditCards[0].ID)
		assert.Equal(t, "creator", data[0].CreditCards[0].CreatedBy)
		assert.Equal(t, "updater", data[0].CreditCards[0].UpdatedBy)
		assert.Equal(t, userID, data[0].CreditCards[0].UserID)
		assert.Equal(t, "CC2014060101", data[0].CreditCards[0].Number)

		assert.Equal(t, uint(11), data[0].CreditCards[1].ID)
		assert.Equal(t, "creator", data[0].CreditCards[1].CreatedBy)
		assert.Equal(t, "updater", data[0].CreditCards[1].UpdatedBy)
		assert.Equal(t, userID, data[0].CreditCards[1].UserID)
		assert.Equal(t, "CC2014060102", data[0].CreditCards[1].Number)

		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

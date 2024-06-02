package repository

import (
	"context"
	"testing"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	"github.com/hanifmaliki/golang-boilerplate/pkg/database/postgres"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	userID    = uint(1)
	companyID = uint(2)
	ccID      = uint(3)
)

func TestCreate(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","created_by","updated_by","deleted_by","name","email","company_id"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8,\$9\) RETURNING "id"`).
		WithArgs(
			sqlmock.AnyArg(),          // created_at
			sqlmock.AnyArg(),          // updated_at
			nil,                       // deleted_at
			"creator",                 // created_by
			"creator",                 // updated_by
			"",                        // deleted_by
			"Hanif Maliki Dewanto",    // name
			"hanifmaliki97@gmail.com", // email
			companyID,                 // company_id
		).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))
	mock.ExpectCommit()

	data := &entity.User{
		Name:      "Hanif Maliki Dewanto",
		Email:     "hanifmaliki97@gmail.com",
		CompanyID: companyID,
	}

	err = repo.Create(ctx, data, "creator")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, userID, data.ID)
	assert.Equal(t, "creator", data.CreatedBy)
	assert.Equal(t, "creator", data.UpdatedBy)
	assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
	assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
	assert.Equal(t, companyID, data.CompanyID)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateOrUpdate(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)

	t.Run("No ID", func(t *testing.T) {
		ctx := context.Background()

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","created_by","updated_by","deleted_by","name","email","company_id"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8,\$9\) RETURNING "id"`).
			WithArgs(
				sqlmock.AnyArg(),          // created_at
				sqlmock.AnyArg(),          // updated_at
				nil,                       // deleted_at
				"creator",                 // created_by
				"creator",                 // updated_by
				"",                        // deleted_by
				"Hanif Maliki Dewanto",    // name
				"hanifmaliki97@gmail.com", // email
				companyID,                 // company_id
			).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))
		mock.ExpectCommit()

		data := &entity.User{
			Name:      "Hanif Maliki Dewanto",
			Email:     "hanifmaliki97@gmail.com",
			CompanyID: companyID,
		}

		err := repo.CreateOrUpdate(ctx, data, "creator")
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, userID, data.ID)
		assert.Equal(t, "creator", data.CreatedBy)
		assert.Equal(t, "creator", data.UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
		assert.Equal(t, companyID, data.CompanyID)

		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("With ID", func(t *testing.T) {
		ctx := context.Background()

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "users" SET "created_at"=\$1,"updated_at"=\$2,"deleted_at"=\$3,"created_by"=\$4,"updated_by"=\$5,"deleted_by"=\$6,"name"=\$7,"email"=\$8,"company_id"=\$9 WHERE "users"."deleted_at" IS NULL AND "id" = \$10`).
			WithArgs(
				sqlmock.AnyArg(),          // created_at
				sqlmock.AnyArg(),          // updated_at
				nil,                       // deleted_at
				"creator",                 // created_by
				"updater",                 // updated_by
				"",                        // deleted_by
				"Hanif Maliki Dewanto",    // name
				"hanifmaliki97@gmail.com", // email
				companyID,                 // company_id
				userID,                    // id
			).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		data := &entity.User{
			Base: pkg_model.Base{
				ID:        userID,
				CreatedBy: "creator",
				UpdatedBy: "creator",
			},
			Name:      "Hanif Maliki Dewanto",
			Email:     "hanifmaliki97@gmail.com",
			CompanyID: companyID,
		}

		err := repo.CreateOrUpdate(ctx, data, "updater")
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, userID, data.ID)
		assert.Equal(t, "creator", data.CreatedBy)
		assert.Equal(t, "updater", data.UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
		assert.Equal(t, companyID, data.CompanyID)

		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "updated_at"=\$1,"updated_by"=\$2,"name"=\$3,"email"=\$4,"company_id"=\$5 WHERE "users"."id" = \$6 AND "users"."deleted_at" IS NULL`).
		WithArgs(
			sqlmock.AnyArg(),          // updated_at
			"updater",                 // updated_by
			"Hanif Maliki Dewanto",    // name
			"hanifmaliki97@gmail.com", // email
			companyID,                 // company_id
			userID,                    // id
		).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	data := &entity.User{
		Name:      "Hanif Maliki Dewanto",
		Email:     "hanifmaliki97@gmail.com",
		CompanyID: companyID,
	}
	conds := &entity.User{Base: pkg_model.Base{ID: 1}}

	err = repo.Update(ctx, data, conds, "updater")
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "deleted_by"=\$1 WHERE "users"."id" = \$2 AND "users"."deleted_at" IS NULL`).
		WithArgs("eraser", userID).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`UPDATE "users" SET "deleted_at"=\$1 WHERE "users"."id" = \$2 AND "users"."deleted_at" IS NULL`).
		WithArgs(sqlmock.AnyArg(), userID).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	conds := &entity.User{Base: pkg_model.Base{ID: 1}}
	err = repo.Delete(ctx, conds, "eraser")
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestHardDelete(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "users" WHERE "users"."id" = \$1`).
		WithArgs(userID).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	conds := &entity.User{Base: pkg_model.Base{ID: 1}}
	err = repo.HardDelete(ctx, conds)
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindOne(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)

	t.Run("No Query", func(t *testing.T) {
		ctx := context.Background()

		rows := sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name", "email", "company_id"}).
			AddRow(userID, "creator", "updater", "Hanif Maliki Dewanto", "hanifmaliki97@gmail.com", companyID)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
			WithArgs(userID, 1).WillReturnRows(rows)

		conds := &entity.User{Base: pkg_model.Base{ID: userID}}
		query := &pkg_model.Query{}
		data, err := repo.FindOne(ctx, conds, query)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		assert.Equal(t, userID, data.ID)
		assert.Equal(t, "creator", data.CreatedBy)
		assert.Equal(t, "updater", data.UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
		assert.Equal(t, companyID, data.CompanyID)

		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("With Query", func(t *testing.T) {
		ctx := context.Background()

		rows := sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name", "email", "company_id"}).
			AddRow(userID, "creator", "updater", "Hanif Maliki Dewanto", "hanifmaliki97@gmail.com", companyID)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 AND "users"."deleted_at" IS NULL ORDER BY id desc,"users"."id" LIMIT \$2`).
			WithArgs(userID, 1).WillReturnRows(rows)

		rows = sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name"}).
			AddRow(companyID, "creator", "updater", "Petrosea")
		mock.ExpectQuery(`SELECT \* FROM "companies" WHERE "companies"."id" = \$1 AND "companies"."deleted_at" IS NULL`).
			WithArgs(companyID).WillReturnRows(rows)

		rows = sqlmock.NewRows([]string{"id", "created_by", "updated_by", "user_id", "number"}).
			AddRow(ccID, "creator", "updater", userID, "CC2014060101").
			AddRow(uint(11), "creator", "updater", userID, "CC2014060102")
		mock.ExpectQuery(`SELECT \* FROM "credit_cards" WHERE "credit_cards"."user_id" = \$1 AND "credit_cards"."deleted_at" IS NULL`).
			WithArgs(userID).WillReturnRows(rows)

		conds := &entity.User{Base: pkg_model.Base{ID: userID}}
		query := &pkg_model.Query{
			SortBy: "id desc",
			Expand: []string{"Company", "CreditCards"},
		}
		data, err := repo.FindOne(ctx, conds, query)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		assert.Equal(t, userID, data.ID)
		assert.Equal(t, "creator", data.CreatedBy)
		assert.Equal(t, "updater", data.UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
		assert.Equal(t, companyID, data.CompanyID)

		assert.Equal(t, companyID, data.Company.ID)
		assert.Equal(t, "creator", data.Company.CreatedBy)
		assert.Equal(t, "updater", data.Company.UpdatedBy)
		assert.Equal(t, "Petrosea", data.Company.Name)

		assert.Equal(t, ccID, data.CreditCards[0].ID)
		assert.Equal(t, "creator", data.CreditCards[0].CreatedBy)
		assert.Equal(t, "updater", data.CreditCards[0].UpdatedBy)
		assert.Equal(t, userID, data.CreditCards[0].UserID)
		assert.Equal(t, "CC2014060101", data.CreditCards[0].Number)

		assert.Equal(t, uint(11), data.CreditCards[1].ID)
		assert.Equal(t, "creator", data.CreditCards[1].CreatedBy)
		assert.Equal(t, "updater", data.CreditCards[1].UpdatedBy)
		assert.Equal(t, userID, data.CreditCards[1].UserID)
		assert.Equal(t, "CC2014060102", data.CreditCards[1].Number)

		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestCount(t *testing.T) {
	gormDB, mock, err := postgres.NewMockDB()
	assert.NoError(t, err)
	repo := NewBaseRepository[entity.User](gormDB)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT count\(\*\) FROM "users" WHERE "users"."deleted_at" IS NULL`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(12))

	conds := &entity.User{}
	count, err := repo.Count(ctx, conds)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), count)

	assert.Nil(t, mock.ExpectationsWereMet())
}

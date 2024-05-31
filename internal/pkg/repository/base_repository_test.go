package repository

import (
	"context"
	"testing"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup function to create a new repository and SQL mock
func setupTestRepository(t *testing.T) (*baseRepository[entity.User], sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	repo := NewBaseRepository[entity.User](gormDB)
	return repo.(*baseRepository[entity.User]), mock
}

func TestCreate(t *testing.T) {
	repo, mock := setupTestRepository(t)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","created_by","updated_by","deleted_by","name","email","company_id","id"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8,\$9,\$10\) RETURNING "id"`).
		WithArgs(
			sqlmock.AnyArg(),          // created_at
			sqlmock.AnyArg(),          // updated_at
			nil,                       // deleted_at
			"creator",                 // created_by
			"creator",                 // updated_by
			"",                        // deleted_by
			"Hanif Maliki Dewanto",    // name
			"hanifmaliki97@gmail.com", // email
			uint(2),                   // company_id
			uint(1),                   // id
		).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	data := &entity.User{
		Base: pkg_model.Base{
			ID: uint(1),
		},
		Name:      "Hanif Maliki Dewanto",
		Email:     "hanifmaliki97@gmail.com",
		CompanyID: uint(2),
	}

	err := repo.Create(ctx, data, "creator")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, uint(1), data.ID)
	assert.Equal(t, "creator", data.CreatedBy)
	assert.Equal(t, "creator", data.UpdatedBy)
	assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
	assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
	assert.Equal(t, uint(2), data.CompanyID)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSave(t *testing.T) {
	repo, mock := setupTestRepository(t)

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
				uint(2),                   // company_id
				uint(1),                   // id
			).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		data := &entity.User{
			Base: pkg_model.Base{
				ID:        uint(1),
				CreatedBy: "creator",
				UpdatedBy: "creator",
			},
			Name:      "Hanif Maliki Dewanto",
			Email:     "hanifmaliki97@gmail.com",
			CompanyID: uint(2),
		}

		err := repo.Save(ctx, data, "updater")
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, uint(1), data.ID)
		assert.Equal(t, "creator", data.CreatedBy)
		assert.Equal(t, "updater", data.UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
		assert.Equal(t, uint(2), data.CompanyID)

		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Without ID", func(t *testing.T) {
		ctx := context.Background()

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","created_by","updated_by","deleted_by","name","email","company_id","id"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8,\$9,\$10\) RETURNING "id"`).
			WithArgs(
				sqlmock.AnyArg(),          // created_at
				sqlmock.AnyArg(),          // updated_at
				nil,                       // deleted_at
				"creator",                 // created_by
				"creator",                 // updated_by
				"",                        // deleted_by
				"Hanif Maliki Dewanto",    // name
				"hanifmaliki97@gmail.com", // email
				uint(2),                   // company_id
				uint(1),                   // id
			).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
		mock.ExpectCommit()

		data := &entity.User{
			Name:      "Hanif Maliki Dewanto",
			Email:     "hanifmaliki97@gmail.com",
			CompanyID: uint(2),
		}

		err := repo.Save(ctx, data, "creator")
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, uint(1), data.ID)
		assert.Equal(t, "creator", data.CreatedBy)
		assert.Equal(t, "creator", data.UpdatedBy)
		assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
		assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
		assert.Equal(t, uint(2), data.CompanyID)

		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	repo, mock := setupTestRepository(t)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "updated_at"=\$1,"updated_by"=\$2,"name"=\$3,"email"=\$4,"company_id"=\$5 WHERE "users"."id" = \$6 AND "users"."deleted_at" IS NULL`).
		WithArgs(
			sqlmock.AnyArg(),          // updated_at
			"updater",                 // updated_by
			"Hanif Maliki Dewanto",    // name
			"hanifmaliki97@gmail.com", // email
			uint(2),                   // company_id
			uint(1),                   // id
		).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	data := &entity.User{
		Name:      "Hanif Maliki Dewanto",
		Email:     "hanifmaliki97@gmail.com",
		CompanyID: uint(2),
	}
	conds := &entity.User{Base: pkg_model.Base{ID: 1}}

	err := repo.Update(ctx, data, conds, "updater")
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	repo, mock := setupTestRepository(t)
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "deleted_by"=\$1 WHERE "users"."id" = \$2 AND "users"."deleted_at" IS NULL`).
		WithArgs("eraser", uint(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`UPDATE "users" SET "deleted_at"=\$1 WHERE "users"."id" = \$2 AND "users"."deleted_at" IS NULL`).
		WithArgs(sqlmock.AnyArg(), uint(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	conds := &entity.User{Base: pkg_model.Base{ID: 1}}
	err := repo.Delete(ctx, conds, "eraser")
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindOne(t *testing.T) {
	repo, mock := setupTestRepository(t)
	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name", "email", "company_id"}).
		AddRow(uint(1), "creator", "updater", "Hanif Maliki Dewanto", "hanifmaliki97@gmail.com", uint(2))

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(uint(1), 1).WillReturnRows(rows)

	conds := &entity.User{Base: pkg_model.Base{ID: uint(1)}}
	query := &pkg_model.Query{}
	data, err := repo.FindOne(ctx, conds, query)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, uint(1), data.ID)
	assert.Equal(t, "creator", data.CreatedBy)
	assert.Equal(t, "updater", data.UpdatedBy)
	assert.Equal(t, "Hanif Maliki Dewanto", data.Name)
	assert.Equal(t, "hanifmaliki97@gmail.com", data.Email)
	assert.Equal(t, uint(2), data.CompanyID)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCount(t *testing.T) {
	repo, mock := setupTestRepository(t)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT count\(\*\) FROM "users" WHERE "users"."deleted_at" IS NULL`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	conds := &entity.User{}
	count, err := repo.Count(ctx, conds)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	assert.Nil(t, mock.ExpectationsWereMet())
}

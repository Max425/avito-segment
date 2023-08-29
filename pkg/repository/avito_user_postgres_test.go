package repository

import (
	"database/sql"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvitoUserPostgres_UpdateUserSegments(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAvitoUserPostgres(db)

	type args struct {
		userID         int
		addSegments    []string
		removeSegments []string
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Success",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM users_segments").
					WithArgs(1, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectExec("INSERT INTO users_segments").
					WithArgs(1, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			input: args{
				userID:         1,
				addSegments:    []string{"segment1", "segment2"},
				removeSegments: []string{"segment3"},
			},
		},
		{
			name: "Error",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM users_segments").
					WithArgs(1, sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)

				mock.ExpectRollback()
			},
			input: args{
				userID:         1,
				addSegments:    []string{"segment1", "segment2"},
				removeSegments: []string{"segment3"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := repo.UpdateUserSegments(tt.input.userID, tt.input.addSegments, tt.input.removeSegments)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

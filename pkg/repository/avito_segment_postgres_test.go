package repository

import (
	"avito-segment/models"
	"database/sql"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvitoSegmentPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAvitoSegmentPostgres(db)

	type args struct {
		segment models.Segment
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"slug"}).AddRow("test-slug")
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO segments").WithArgs("test-slug").WillReturnRows(rows)
				mock.ExpectCommit()
			},
			input: args{
				segment: models.Segment{
					Slug: "test-slug",
				},
			},
			want: "test-slug",
		},
		{
			name: "Error",
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO segments").WithArgs("test-slug").
					WillReturnError(sql.ErrConnDone)
				mock.ExpectRollback()
			},
			input: args{
				segment: models.Segment{
					Slug: "test-slug",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := repo.Create(tt.input.segment)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAvitoSegmentPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAvitoSegmentPostgres(db)

	type args struct {
		slug string
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
				mock.ExpectExec("DELETE FROM segments").WithArgs("test-slug").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				slug: "test-slug",
			},
		},
		{
			name: "Error",
			mock: func() {
				mock.ExpectExec("DELETE FROM segments").WithArgs("test-slug").
					WillReturnError(sql.ErrConnDone)
			},
			input: args{
				slug: "test-slug",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := repo.Delete(tt.input.slug)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAvitoSegmentPostgres_GetUserSegments(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAvitoSegmentPostgres(db)

	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []models.Segment
		wantErr bool
	}{
		{
			name: "Success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"slug"}).AddRow("segment1").AddRow("segment2")
				mock.ExpectQuery("SELECT s.slug FROM segments").WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				userID: 1,
			},
			want: []models.Segment{
				{Slug: "segment1"},
				{Slug: "segment2"},
			},
		},
		{
			name: "Error",
			mock: func() {
				mock.ExpectQuery("SELECT s.slug FROM segments").WithArgs(1).
					WillReturnError(sql.ErrConnDone)
			},
			input: args{
				userID: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := repo.GetUserSegments(tt.input.userID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

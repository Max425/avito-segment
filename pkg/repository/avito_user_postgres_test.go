package repository

import (
	"database/sql"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAvitoUserPostgres_GenerateUserSegmentHistoryReport(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAvitoUserPostgres(db)

	type args struct {
		userID int
		year   int
		month  time.Month
	}
	tests := []struct {
		name     string
		mock     func()
		input    args
		wantErr  bool
		expected string // Expected report content
	}{
		{
			name: "Success",
			mock: func() {
				mock.ExpectQuery("SELECT user_id, segment_slug, action, timestamp FROM users_segments_history").
					WithArgs(1, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"user_id", "segment_slug", "action", "timestamp"}).
					AddRow(1, "segment1", "add", time.Now()).
					AddRow(1, "segment2", "remove", time.Now()))
			},
			input: args{
				userID: 1,
				year:   2023,
				month:  time.August,
			},
			expected: "1;segment1;add;" + time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST") + "\n" +
				"1;segment2;remove;" + time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST") + "\n",
		},
		{
			name: "Error",
			mock: func() {
				mock.ExpectQuery("SELECT user_id, segment_slug, action, timestamp FROM users_segments_history").
					WithArgs(1, sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)
			},
			input: args{
				userID: 1,
				year:   2023,
				month:  time.August,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			reportData, err := repo.GenerateUserSegmentHistoryReport(tt.input.userID, tt.input.year, tt.input.month)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				actualLines := strings.Split(strings.TrimSpace(string(reportData)), "\n")
				expectedLines := strings.Split(strings.TrimSpace(tt.expected), "\n")

				assert.Equal(t, len(expectedLines), len(actualLines))
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

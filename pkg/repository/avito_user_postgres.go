package repository

import (
	"avito-segment/models"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
	"time"
)

type AvitoUserPostgres struct {
	db *sqlx.DB
}

func NewAvitoUserPostgres(db *sqlx.DB) *AvitoUserPostgres {
	return &AvitoUserPostgres{db: db}
}

func (r *AvitoUserPostgres) UpdateUserSegments(userID int, addSegments, removeSegments []string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if len(removeSegments) > 0 {
		query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_slug = ANY($2)", usersSegmentsTable)
		_, err := tx.Exec(query, userID, pq.Array(removeSegments))
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if len(addSegments) > 0 {
		query := fmt.Sprintf("INSERT INTO %s (user_id, segment_slug) SELECT $1, slug FROM %s WHERE slug = ANY($2)",
			usersSegmentsTable, segmentsTable)
		_, err := tx.Exec(query, userID, pq.Array(addSegments))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Сохраняем действия в историю
	if len(removeSegments) > 0 {
		for _, segmentSlug := range removeSegments {
			_, err := tx.Exec(
				fmt.Sprintf("INSERT INTO %s (user_id, segment_slug, action) VALUES ($1, $2, $3)", usersSegmentsHistoryTable),
				userID, segmentSlug, "removal",
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if len(addSegments) > 0 {
		for _, segmentSlug := range addSegments {
			_, err := tx.Exec(
				fmt.Sprintf("INSERT INTO %s (user_id, segment_slug, action) VALUES ($1, $2, $3)", usersSegmentsHistoryTable),
				userID, segmentSlug, "addition",
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

func (r *AvitoUserPostgres) AddUserToSegmentWithTTL(userID int, segmentSlug string, ttl time.Duration) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Вычислите время истечения TTL
	expiresAt := time.Now().Add(ttl)

	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_slug, expires_at) VALUES ($1, $2, $3)", usersSegmentsTable)
	_, err = tx.Exec(query, userID, segmentSlug, expiresAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Сохранение в историю
	historyQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_slug, action, timestamp) VALUES ($1, $2, $3, $4)", usersSegmentsHistoryTable)
	_, err = tx.Exec(historyQuery, userID, segmentSlug, "addition", time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *AvitoUserPostgres) GenerateUserSegmentHistoryReport(userID int, year int, month time.Month) ([]byte, error) {
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	query := fmt.Sprintf("SELECT user_id, segment_slug, action, timestamp FROM %s WHERE user_id = $1 AND timestamp >= $2", usersSegmentsHistoryTable)

	var history []models.UserSegmentHistory
	if err := r.db.Select(&history, query, userID, startDate); err != nil {
		return nil, err
	}

	// Создание CSV файла и запись данных
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	for _, entry := range history {
		record := []string{
			strconv.Itoa(entry.UserID),
			entry.SegmentSlug,
			entry.Action,
			entry.Timestamp.String(),
		}
		_ = writer.Write(record)
	}

	writer.Flush()

	return buf.Bytes(), nil
}

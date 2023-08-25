package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

	return tx.Commit()
}

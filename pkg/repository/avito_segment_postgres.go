package repository

import (
	"avito-segment/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AvitoSegmentPostgres struct {
	db *sqlx.DB
}

func NewAvitoSegmentPostgres(db *sqlx.DB) *AvitoSegmentPostgres {
	return &AvitoSegmentPostgres{db: db}
}

func (r *AvitoSegmentPostgres) Create(segment models.Segment) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var segmentSlug string
	createItemQuery := fmt.Sprintf("INSERT INTO %s (slug) values ($1) RETURNING slug", segmentsTable)

	row := tx.QueryRow(createItemQuery, segment.Slug)
	err = row.Scan(&segmentSlug)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return segmentSlug, tx.Commit()
}

func (r *AvitoSegmentPostgres) Delete(slug string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE slug = $1;", segmentsTable)
	_, err := r.db.Exec(query, slug)

	return err
}

func (r *AvitoSegmentPostgres) GetUserSegments(userID int) ([]models.Segment, error) {
	query := fmt.Sprintf(`
		SELECT s.slug FROM %s s
		INNER JOIN %s us ON s.slug = us.segment_slug
		WHERE us.user_id = $1
	`, segmentsTable, usersSegmentsTable)

	var segments []models.Segment
	err := r.db.Select(&segments, query, userID)
	if err != nil {
		return nil, err
	}
	return segments, nil
}

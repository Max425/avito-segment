package repository

import (
	avito_segment "avito-segment"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AvitoSegmentPostgres struct {
	db *sqlx.DB
}

func NewAvitoSegmentPostgres(db *sqlx.DB) *AvitoSegmentPostgres {
	return &AvitoSegmentPostgres{db: db}
}

func (r *AvitoSegmentPostgres) Create(segment avito_segment.Segment) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var segmentId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (slug) values ($1) RETURNING id", segmentsTable)

	row := tx.QueryRow(createItemQuery, segment.Slug)
	err = row.Scan(&segmentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return segmentId, tx.Commit()
}

func (r *AvitoSegmentPostgres) Delete(slug string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE slug = $1;", segmentsTable)
	_, err := r.db.Exec(query, slug)

	return err
}

func (r *AvitoSegmentPostgres) GetUserSegments(userID int) ([]avito_segment.Segment, error) {
	query := fmt.Sprintf(`
		SELECT s.slug FROM %s s
		INNER JOIN %s us ON s.slug = us.segment_slug
		WHERE us.user_id = $1
	`, segmentsTable, usersSegmentsTable)

	var segments []avito_segment.Segment
	err := r.db.Select(&segments, query, userID)
	if err != nil {
		return nil, err
	}
	return segments, nil
}

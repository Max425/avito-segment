package models

import "time"

type UserSegmentHistory struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	SegmentSlug string    `db:"segment_slug"`
	Action      string    `db:"action"`
	Timestamp   time.Time `db:"timestamp"`
}

package models

type Segment struct {
	Slug string `json:"slug" binding:"required"`
}

package avito_segment

type Segment struct {
	Id   int    `json:"id" db:"id"`
	Slug string `json:"slug" binding:"required"`
}

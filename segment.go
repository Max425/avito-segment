package avito_segment

type Segment struct {
	Slug string `json:"slug" binding:"required"`
}

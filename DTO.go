package avito_segment

type User struct {
	Id int `json:"id" db:"id"`
}

type Segment struct {
	Slug string `json:"slug" binding:"required"`
}

type UserSegmentsRequest struct {
	AddSegments    []string `json:"add_segments"`
	RemoveSegments []string `json:"remove_segments"`
}

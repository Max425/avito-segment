package avito_segment

type Segment struct {
	Slug string `json:"slug" binding:"required"`
}

type User struct {
	Id int `json:"id" db:"id"`
}

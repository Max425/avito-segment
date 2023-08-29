package models

type UserSegmentsRequest struct {
	AddSegments    []string `json:"add_segments"`
	RemoveSegments []string `json:"remove_segments"`
}

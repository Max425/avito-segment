package models

type UserSegmentsRequest struct {
	AddSegments    []string `json:"add_segments"`
	RemoveSegments []string `json:"remove_segments"`
}

type UserToSegmentWithTTLRequest struct {
	SegmentSlug string `json:"segment_slug"`
	TTLMinutes  int    `json:"ttl_minutes"`
}

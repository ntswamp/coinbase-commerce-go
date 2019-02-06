package coinbase

type Pagination struct {
	// User configurable
	Order         string `json:"order,omitempty"`          // asc || desc
	StartingAfter string `json:"starting_after,omitempty"` // resourceID
	EndingBefore  string `json:"ending_before,omitempty"`  // resourceID
	Limit         int    `json:"limit,omitempty"`          // default 25 , max 100

	// Part of response
	Total       int      `json:"total,omitempty"`
	Yielded     int      `json:"yielded,omitempty"`
	PreviousURI string   `json:"previous_uri,omitempty"`
	NextURI     string   `json:"next_uri,omitempty"`
	CursorRange []string `json:"cursor_range,omitempty"` // from resourceID to resourceID
}

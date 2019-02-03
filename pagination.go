package coinbase

type Pagination struct {
	Order         string   `json:"order,omitempty"`
	StartingAfter string   `json:"starting_after,omitempty"`
	EndingBefore  string   `json:"ending_before,omitempty"`
	Total         int      `json:"total,omitempty"`
	Yielded       int      `json:"yielded,omitempty"`
	Limit         int      `json:"limit,omitempty"`
	PreviousURI   string   `json:"previous_uri,omitempty"`
	NextURI       string   `json:"next_uri,omitempty"`
	CursorRange   []string `json:"cursor_range,omitempty"`
}

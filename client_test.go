package coinbase

import (
	"testing"
)

func TestClient(t *testing.T) {
	_ = NewTestClient()
}

func TestURLFromPagination(t *testing.T) {
	basePath := "/checkouts"
	expected := "/checkouts?limit=25&starting_after=1234"
	pagination := Pagination{
		Limit:         25,
		StartingAfter: "1234",
	}
	res := urlWithPagination(basePath, &pagination)
	if res != expected {
		t.Fatalf("expected: %s, got %s", expected, res)
	}
}

func TestURLToPath(t *testing.T) {
	testURL := "https://api.commerce.coinbase.com/checkouts?limit=20&starting_after=fb6721f2-1622-48f0-b713-aac6c819b67a"
	expected := "/checkouts?limit=20&starting_after=fb6721f2-1622-48f0-b713-aac6c819b67a"
	res := uriToPathQuery(testURL)
	if res != expected {
		t.Fatalf("expected: %s, got %s", expected, res)
	}
}

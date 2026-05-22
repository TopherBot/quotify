package main

import "testing"

func TestFetchQuote(t *testing.T) {
    // In a unit test we don't want to hit the network.
    // Instead we verify that the struct can be unmarshaled correctly.
    sample := `{"content":"Be yourself; everyone else is already taken.","author":"Oscar Wilde"}`
    var q quoteResponse
    if err := json.Unmarshal([]byte(sample), &q); err != nil {
        t.Fatalf("unmarshal failed: %v", err)
    }
    if q.Content == "" || q.Author == "" {
        t.Fatalf("unexpected empty fields: %+v", q)
    }
}

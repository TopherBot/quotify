package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
)

type quoteResponse struct {
    Content string `json:"content"`
    Author  string `json:"author"`
}

func fetchQuote() (quoteResponse, error) {
    resp, err := http.Get("https://api.quotable.io/random")
    if err != nil {
        return quoteResponse{}, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return quoteResponse{}, fmt.Errorf("unexpected status: %s", resp.Status)
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return quoteResponse{}, err
    }
    var q quoteResponse
    if err := json.Unmarshal(body, &q); err != nil {
        return quoteResponse{}, err
    }
    return q, nil
}

func main() {
    q, err := fetchQuote()
    if err != nil {
        fmt.Fprintln(os.Stderr, "error fetching quote:", err)
        os.Exit(1)
    }
    fmt.Printf("\"%s\" – %s\n", q.Content, q.Author)
}

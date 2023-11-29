package main

import (
    "net/url"
    "testing"
)

func TestBuildRequestURL(t *testing.T) {
    params := &SearchParams{
        Offset: 0,
        Limit: 10,
        Query: "大根",
        Models: "KN-HW24G",
    }

    result := buildRequestURL(params)
    parsedURL, err := url.Parse(result)
    if err != nil {
        t.Fatalf("Failed to parse URL: %v", err)
    }

    // クエリパラメータのチェック
    values := parsedURL.Query()
    checkQueryParam(t, values, "offset", "0")
    checkQueryParam(t, values, "limit", "10")
    checkQueryParam(t, values, "search", "大根")
    checkQueryParam(t, values, "models", "KN-HW24G")
}

func checkQueryParam(t *testing.T, values url.Values, key, expectedValue string) {
    if values.Get(key) != expectedValue {
        t.Errorf("QueryParam %s: expected %s, got %s", key, expectedValue, values.Get(key))
    }
}


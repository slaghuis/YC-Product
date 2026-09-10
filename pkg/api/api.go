package api

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
  //"os"
  //"path/filepath"
)

type APIServer struct {
  url string
}

func NewAPIServer(url string) *APIServer {
  return &APIServer{
    url: url,
  }
}

func (a *APIServer) APICall(method, requestURL, tokenString string, v any, out any) error {
    // Marshal the request body
    jsonData, err := json.Marshal(v)
    if err != nil {
        return fmt.Errorf("failed to marshal request: %w", err)
    }

    // Create the request with the given method
    req, err := http.NewRequest(method, a.url+requestURL, bytes.NewReader(jsonData))
    if err != nil {
        return fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    if len(tokenString) > 1 {
        req.Header.Set("Authorization", tokenString)
    }

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
        return fmt.Errorf("Unexpected status: %s", resp.Status)
    }

    // Decode response into `out` (must be a pointer)
    if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
        return fmt.Errorf("failed to decode response: %w", err)
    }

    return nil
}


func (a *APIServer) GetUrl() string {
  return a.url
}

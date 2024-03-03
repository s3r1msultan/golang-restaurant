package routers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
    "final/routers"

)
// TO RUN THE TEST 
// IN initTemplates.go UNCOMENT  
func TestHomeRouter(t *testing.T) {
    r := mux.NewRouter()
    routers.HomeRouter(r)

    ts := httptest.NewServer(r) 
    defer ts.Close()

    res, err := http.Get(ts.URL + "/")
    if err != nil {
        t.Fatalf("Failed to send GET request: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
    }

}

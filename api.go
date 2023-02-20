package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items []Item

func main() {
    // Add some example items to the slice
    items = append(items, Item{ID: "1", Name: "Apple"})
    items = append(items, Item{ID: "2", Name: "Banana"})
    items = append(items, Item{ID: "3", Name: "Orange"})

    // Set up the API routes
    http.HandleFunc("/items", getItems)
    http.HandleFunc("/items/add", addItem)

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getItems(w http.ResponseWriter, r *http.Request) {
    // Set the content type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode the items slice as JSON and write it to the response writer
    json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
    // Parse the request body into an Item struct
    var item Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Add the new item to the slice
    items = append(items, item)

    // Set the response status code to 201 Created
    w.WriteHeader(http.StatusCreated)
}
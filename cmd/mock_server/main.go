package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/workflows", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %s request to %s\n", r.Method, r.URL.Path)
		fmt.Printf("Headers: %v\n", r.Header)
		body, _ := io.ReadAll(r.Body)
		fmt.Printf("Body: %s\n", string(body))

		// Return success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"submitted"}`))
	})

	fmt.Println("Mock ColonyOS server listening on :50080")
	log.Fatal(http.ListenAndServe(":50080", nil))
}

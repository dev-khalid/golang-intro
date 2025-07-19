package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func StartServer() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		uptime := time.Since(startTime)
		
		jsonResponse, _ := json.MarshalIndent(map[string]interface{}{
			"status":  "healthy",
			"message": "Server is running smoothly",
			"metadata": map[string]interface{}{
				"version": "1.0.0",
				"uptime":  uptime.String(),
				"serverTime": map[string]string{
					"format":   time.Now().Format(time.RFC3339),
					"location": time.Now().Location().String(),
				},
			},
		}, "", "	")
		w.Write(jsonResponse)
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

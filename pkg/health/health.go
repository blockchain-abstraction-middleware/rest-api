package health

import (
	"encoding/json"
	"fmt"
	"time"
)

// Handler used in the health resource
func Handler() ([]byte, error) {
	payload := struct {
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	}{
		fmt.Sprintf("Healthy response from service at - %s", time.Now()),
		200,
	}

	json, err := json.Marshal(payload)

	return json, err
}

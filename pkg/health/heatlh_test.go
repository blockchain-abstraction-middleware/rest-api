package health

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	_, err := Handler()
	if err != nil {
		log.Fatal("Health handler is broken", err)
	}
}

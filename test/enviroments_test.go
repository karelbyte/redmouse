package test

import (
	"errors"
	"os"
	"testing"
)

func TestEnviroment(t *testing.T) {
	if _, err := os.Stat("../.env"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("file .env not found")
	}
}

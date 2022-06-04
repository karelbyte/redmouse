package test

import (
	"errors"
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	if _, err := os.Stat("../public/assets"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("../public/assets folder not found")
	}

	if _, err := os.Stat("../public/index.html"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("../public/index.html file not found")
	}
}

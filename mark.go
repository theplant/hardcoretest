package hardcoretest

import (
	"os"
	"testing"
)

func Mark(t *testing.T) {
	if len(os.Getenv("HARDCORE")) == 0 {
		t.Skip("Hardcore test, use HARDCORE=1 go test ./... to run this.")
	}
}

package cursediswsl

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestIsWSL(t *testing.T) {
	cmd := exec.Command("uname", "-a")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	expected := err == nil &&
		(strings.Contains(out.String(), "Microsoft") || strings.Contains(out.String(), "microsoft"))

	result := IsWSL()
	if result != expected {
		t.Fatalf("Got %v, expected %v", result, expected)
	}
}

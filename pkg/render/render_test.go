package render_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/madest92/goj2/pkg/render"
)

func TestTemplate(t *testing.T) {
	// Create temporary files for testing
	fromFile := "testdata/template.yml.j2"
	toFile := "testdata/template.yml"
	varsFile := []string{"testdata/var1-template.yml", "testdata/var2-template.yml"}
	os.Setenv("VAR", "v")
	res := "top:\n  a: ${VAR} is v\n  b:\n    a: a\n    b: b2\n"
	defer os.Remove(toFile)

	// Check if the output file was created with correct content
	render.Template(fromFile, toFile, varsFile)
	outputContent, err := os.ReadFile(toFile)
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	expectedOutput := []byte(res)
	if !bytes.Equal(outputContent, expectedOutput) {
		t.Errorf("Expected output content %q, got %q", expectedOutput, outputContent)
	}

	// Check if the STDOUT was created with correct content
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	render.Template(fromFile, "", varsFile)
	w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old

	expectedOutput = []byte(res + "\n")
	if !bytes.Equal(buf.Bytes(), expectedOutput) {
		t.Errorf("Expected output content %q, got %q", expectedOutput, buf.String())
	}
}

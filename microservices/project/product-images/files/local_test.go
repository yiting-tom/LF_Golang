package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// setupLocal creates a new Local struct, a tmp directory, and a cleanup function for these tests.
func setupLocal(t *testing.T) (*Local, string, func()) {
	// Create a temporary directory for testing
	dir, err := ioutil.TempDir("", "test-tmp-files")
	if err != nil {
		t.Fatal(err)
	}

	l, err := NewLocal(dir, 100)
	if err != nil {
		t.Fatal(err)
	}

	return l, dir, func() {
		// Remove the temporary directory
		os.RemoveAll(dir)
	}
}

// TestLocal_NewLocal tests the Save function of the Local struct.
func TestSavesContentsOfReader(t *testing.T) {
	path := "/1/just-a-test.png"
	contents := "just a test"

	// Get a new Local struct, tmp directory, and the cleanup function
	l, dir, cleanup := setupLocal(t)
	defer cleanup()

	// Save the contents of the reader into the local file
	err := l.Save(path, strings.NewReader(contents))
	assert.NoError(t, err)

	// Assert the file exists
	f, err := os.Open(filepath.Join(dir, path))
	assert.NoError(t, err)

	// Assert the contents are correct
	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, contents, string(d))
}

func TestGetsContentsAndWritesToWriter(t *testing.T) {
	path := "1/just-a-test.png"
	contents := "just a test"

	// Get a new Local struct, tmp directory, and the cleanup function
	l, _, cleanup := setupLocal(t)
	defer cleanup()

	// Save the contents of the reader into the local file
	err := l.Save(path, strings.NewReader(contents))
	assert.NoError(t, err)

	// Get the contents of the file
	file, err := l.Get(path)
	assert.NoError(t, err)
	defer file.Close()

	// Read the contents of the file
	d, err := ioutil.ReadAll(file)
	assert.NoError(t, err)
	assert.Equal(t, contents, string(d))
}

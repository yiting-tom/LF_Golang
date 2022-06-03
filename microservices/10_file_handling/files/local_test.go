package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLocal(t *testing.T) (*Local, string, func()) {
	// Create a tmp dir
	dir, err := ioutil.TempDir("", "files-test")
	if err != nil {
		t.Fatal(err)
	}

	l, err := NewLocal(dir, 100)
	if err != nil {
		t.Fatal(err)
	}

	return l, dir, func() {
		// Cleanup
		os.RemoveAll(dir)
	}
}

func TestSaveContentOfReader(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "test"
	l, dir, cleanup := setupLocal(t)
	defer cleanup()

	err := l.Save(savePath, bytes.NewBufferString(fileContents))
	assert.NoError(t, err)

	// Check the file has been correctly written
	f, err := os.Open(filepath.Join(dir, savePath))
	assert.NoError(t, err)

	// Check the file contents
	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}

func TestGetsContentsAndWritesToWriter(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World"
	l, _, cleanup := setupLocal(t)
	defer cleanup()

	// Save a file
	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// Read the file back
	r, err := l.Get(savePath)
	assert.NoError(t, err)
	defer r.Close()

	// read the full contents of the reader
	d, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}

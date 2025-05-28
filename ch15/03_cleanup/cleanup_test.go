package cleanup

import (
	"errors"
	"os"
	"strings"
	"testing"
)

// createFile is a helper function called from multiple tests
func createFile(t *testing.T) (_ string, err error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	// write some data to f

	// Cleanup()関数はテスト終了後に呼び出される
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, don't worry about cleanup
	if !strings.Contains(fName, "tempFile") {
		t.Error("unexpected name")
	}
}

// createFileWithCreateTemp is a helper function called from multiple tests
func createFileWithCreateTemp(tempDir string) (_ string, err error) {
	f, err := os.CreateTemp(tempDir, "tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	// write some data to f
	return f.Name(), nil
}

func TestFileProcessingWithCreateTemp(t *testing.T) {
	// TempDir()メソッドを使うと、テストが終了した時に自動的に削除される
	tempDir := t.TempDir()
	fName, err := createFileWithCreateTemp(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, don't worry about cleanup
	if !strings.Contains(fName, "tempFile") {
		t.Error("unexpected name")
	}
}

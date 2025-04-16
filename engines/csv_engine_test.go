package engines

import (
	"os"
	"testing"
)

func createDirIfNotExists(filemode os.FileMode) {
	_, err := os.Stat(dataDir)

	if os.IsNotExist(err) {
		os.Mkdir(dataDir, filemode)
	}

}

func removeDirIfExists() {
	_, err := os.Stat(dataDir)

	if err == nil {
		os.Remove(dataDir)
	}
}

func TestCreateCSVEngine(t *testing.T) {

	t.Run("Create CSV engine and data directory doesn't exist", func(t *testing.T) {
		removeDirIfExists()

		_, err := CreateCSVEngine()

		if err != nil {
			t.Errorf("Creating a csv engine shouldn't throw an error | %v", err)
		}
	})

	t.Run("Create CSV engine and data directory exists", func(t *testing.T) {
		createDirIfNotExists(0o666)
		_, err := CreateCSVEngine()

		if err != nil {
			t.Errorf("Creating a csv engine with persistent directory functional shouldn't throw an error | %v", err)
		}
	})

}

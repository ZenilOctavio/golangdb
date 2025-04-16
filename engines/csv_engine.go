package engines

import (
	i "db/interfaces"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var dataDir string = "csv_data"

type CSV_Engine struct {
	directory fs.FS
	models    map[string]i.Model
}

var IsNotDir error = errors.New(fmt.Sprintf("There is a file with the same name as the data directory %v", dataDir))

// This function checks for the existence of the data directory for the CSVEngine and then returns it.
func CreateCSVEngine() (*CSV_Engine, error) {

	_, err := os.Stat(dataDir)

	if err != nil {

		if os.IsNotExist(err) {
			os.Mkdir(dataDir, 0o666)
		} else {
			return nil, err
		}
	}

	fileinfo, err := os.Stat(dataDir)

	if !fileinfo.IsDir() {
		return nil, IsNotDir
	}

	directory := os.DirFS(dataDir)

	_, err = os.ReadDir(dataDir)

	if err != nil {
		return nil, err
	}

	csvEngine := new(CSV_Engine)
	csvEngine.models = make(map[string]i.Model)
	csvEngine.directory = directory

	return csvEngine, nil
}

var EngineNotInitialized = errors.New("The engine hasnt been initialized correctly")

func setUpModel(model i.Model) error {
	fmt.Printf("Setting up model %v | %v\n", model.GetName(), model.GetShape())
	return nil
}

func (e *CSV_Engine) AddModel(model i.Model) error {
	if e.directory == nil {
		return EngineNotInitialized
	}

	_, ok := e.models[model.GetName()]

	if !ok {
		err := setUpModel(model)

		if err != nil {
			return err
		}
		e.models[model.GetName()] = model
	}

	log.Printf("A model with the name %v is already registered.", model.GetName())
	return nil
}

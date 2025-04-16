package engines

import (
	i "db/interfaces"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
)

var dataDir string = "csv_data"
var fileNameFormat string = "%v.csv"

type CSV_Engine struct {
	directory     fs.FS
	models        map[string]i.Model
	modelsSources map[string]*os.File
}

var IsNotDir error = errors.New(fmt.Sprintf("There is a file with the same name as the data directory %v", dataDir))

func (e *CSV_Engine) isInitialized() bool {
	return e.directory != nil
}

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
	csvEngine.modelsSources = make(map[string]*os.File)
	csvEngine.directory = directory

	return csvEngine, nil
}

type addingModelErrorMessage uint32

const (
	ENGINE_NOT_INITIALIZED addingModelErrorMessage = iota
	DIR_INSTEAD_OF_FILE
)

var setUpModelsErrMessages = []string{
	"The engine hasnt been initialized correctly when adding model %v",
	"The file %v used as data source for model %v is a directory",
}

func (e *CSV_Engine) setUpModel(modelp *i.Model) error {
	model := *modelp
	filename := path.Join(dataDir, fmt.Sprintf(fileNameFormat, model.GetName()))

	existingFile, err := os.Stat(filename)

	if err == nil {
		if existingFile.IsDir() {
			return errors.New(fmt.Sprintf(setUpModelsErrMessages[DIR_INSTEAD_OF_FILE], filename, model.GetName()))
		}
		log.Printf("Checking if there are new changes in model %v", model.GetName())
	}

	if os.IsNotExist(err) {
		file, err := os.Create(filename)

		if err != nil {
			return err
		}

		e.modelsSources[model.GetName()] = file
		log.Printf("New model registered %v", model.GetName())
		return nil

	}

	return nil
}

func (e *CSV_Engine) AddModel(model i.Model) error {
	if !e.isInitialized() {
		return errors.New(fmt.Sprintf(setUpModelsErrMessages[ENGINE_NOT_INITIALIZED], model.GetName()))
	}

	_, ok := e.models[model.GetName()]

	if !ok {
		err := e.setUpModel(&model)

		if err != nil {
			return err
		}
		e.models[model.GetName()] = model
	}

	return nil
}

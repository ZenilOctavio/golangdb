package main

import (
	e "db/engines"
	m "db/models"
	"log"
)

type Project struct {
	Name        string `field:"name"`
	Description string `field:"description"`
}

func main() {
	engine, err := e.CreateCSVEngine()

	if err != nil {
		log.Fatalf("Couldn't initialize the engine | error: %v", err)
	}

	projectModelName := "project"
	ProjectModel, err := m.CreateModel(projectModelName, &Project{})

	if err != nil {
		log.Fatalf("Couldn't register model %v", projectModelName)
	}

	err = engine.AddModel(ProjectModel)

	if err != nil {
		log.Printf("Couldnt add a model to the db | %v", err)
	}
	log.Printf("Project model was initialized")

}

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

	engine.AddModel(ProjectModel)
	log.Printf("Project model was initialized")

}

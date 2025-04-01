package scene

import (
	"fmt"
	"github.com/fredericobormann/mindlines-web/mindlines-backend/helper"
	"github.com/gin-gonic/gin"
	"github.com/open-spaced-repetition/go-fsrs/v3"
	"log"
	"os"
)

var dummyScene = Scene{
	Name:       "Testscene",
	Identifier: "testscene",
	Index:      42,
	Content: []Line{
		{
			Character: "HAMLET",
			Line:      "To be or not to be",
			Card:      fsrs.NewCard(),
		},
		{
			Character: "ROMEO",
			Line:      "?",
			Card:      fsrs.NewCard(),
		},
	},
}

func setupTest() Module {
	testModule := createTestModule()
	err := testModule.Service.SaveScene(dummyScene, dummyScene.Index)
	if err != nil {
		fmt.Printf("setting up test failed with err: %v", err)
	}
	return testModule
}

func cleanUp() {
	err := os.Remove(helper.GetFilePath("content/scenefiles/scene42.json"))
	if err != nil {
		log.Fatalf("could not clean up test files %v", err)
	}
}

func createTestModule() Module {
	service := Service{
		fsrs: *fsrs.NewFSRS(fsrs.DefaultParam()),
	}
	controller := Controller{
		Service: service,
		FSRS:    *fsrs.NewFSRS(fsrs.DefaultParam()),
	}

	return Module{
		service,
		controller,
	}
}

func setupRouter(module Module) *gin.Engine {
	r := gin.Default()
	api := r.Group("/")
	module.Controller.RegisterRoutes(api)
	return r
}

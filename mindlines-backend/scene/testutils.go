package scene

import (
	"encoding/json"
	"errors"
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

	if _, err := os.Stat(helper.GetFilePath("content/scenelist.json")); !errors.Is(err, os.ErrNotExist) {
		// backup scenelist.json
		err := os.Rename(helper.GetFilePath("content/scenelist.json"), helper.GetFilePath("content/scenelist.json.bak"))
		if err != nil {
			fmt.Printf("could not backup scenelist.json: %v", err)
		}
	}

	// create test scenelist.json
	err := os.MkdirAll(helper.GetFilePath("content"), os.ModePerm)
	if err != nil {
		fmt.Printf("could not create test directory: %v", err)
	}
	content, _ := json.Marshal([]MetaScene{
		{
			Name:       "Testscene",
			Identifier: "testscene",
			Index:      42,
		},
	})
	err = os.WriteFile(helper.GetFilePath("content/scenelist.json"), content, os.ModePerm)
	if err != nil {
		fmt.Printf("could not create test scenelist.json: %v", err)
	}

	err = testModule.Service.SaveScene(dummyScene, dummyScene.Index)
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

	err = os.Remove(helper.GetFilePath("content/scenelist.json"))
	if err != nil {
		log.Fatalf("could not clean up test files %v", err)
	}

	if _, err := os.Stat(helper.GetFilePath("content/scenelist.json.bak")); !errors.Is(err, os.ErrNotExist) {
		// restore scenelist.json
		err := os.Rename(helper.GetFilePath("content/scenelist.json.bak"), helper.GetFilePath("content/scenelist.json"))
		if err != nil {
			fmt.Printf("could not restore scenelist.json: %v", err)
		}
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

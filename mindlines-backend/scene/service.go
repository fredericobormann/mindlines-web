package scene

import (
	"encoding/json"
	"fmt"
	"github.com/fredericobormann/mindlines-web/mindlines-backend/helper"
	"github.com/open-spaced-repetition/go-fsrs/v3"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Service struct {
	fsrs fsrs.FSRS
}

func (s *Service) GetAll() ([]MetaScene, error) {
	file, err := os.ReadFile(helper.GetFilePath("content/scenelist.json"))
	if err != nil {
		return []MetaScene{}, err
	}

	scenes, err := helper.Unmarshal[[]MetaScene](file)
	if err != nil {
		return []MetaScene{}, err
	}
	return *scenes, nil
}

func (s *Service) GetByIndex(index uint8) (Scene, error) {
	file, err := os.ReadFile(helper.GetFilePath(fmt.Sprintf("content/scenefiles/scene%d.json", index)))
	if err != nil {
		return Scene{}, err
	}
	scene, err := helper.Unmarshal[Scene](file)
	if err != nil {
		return Scene{}, err
	}

	return *scene, nil
}

func (s *Service) UpdateLine(index uint16, line Line, sceneNumber uint8) (Scene, error) {
	scene, err := s.GetByIndex(sceneNumber)
	if err != nil {
		return Scene{}, err
	}

	scene.Content[index] = line
	err = s.SaveScene(scene, sceneNumber)
	if err != nil {
		return Scene{}, err
	}
	return scene, nil
}

func (s *Service) SaveScene(scene Scene, sceneNumber uint8) error {
	var err error
	res, err := json.Marshal(scene)
	if err != nil {
		return err
	}

	path := helper.GetFilePath(fmt.Sprintf("content/scenefiles/scene%d.json", sceneNumber))
	log.Printf("%s", path)
	if err = os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		log.Printf("cannot create file %v", err)
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	_, err = file.Write(res)
	if err != nil {
		log.Printf("cannot write scene file %v", err)
	}

	return err
}

func (s *Service) LearnLine(index uint16, rating fsrs.Rating, sceneNumber uint8) (Scene, error) {
	scene, err := s.GetByIndex(sceneNumber)
	if err != nil {
		return Scene{}, err
	}

	scene.Content[index].Card = s.fsrs.Repeat(scene.Content[index].Card, time.Now())[rating].Card
	err = s.SaveScene(scene, sceneNumber)
	if err != nil {
		return Scene{}, err
	}

	return scene, nil
}

package scene

import (
	"github.com/open-spaced-repetition/go-fsrs/v3"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestService_LearnLine(t *testing.T) {
	testModule := setupTest()
	defer cleanUp()
	newScene, err := testModule.Service.LearnLine(0, fsrs.Good, dummyScene.Index)
	if err != nil {
		t.Errorf("LearnLine failed with: %v", err)
	}
	dueIn := newScene.Content[0].Card.Due.Sub(time.Now())
	diff := dueIn - time.Duration(1000*1000*1000*60*10)
	if diff > time.Duration(1000) {
		log.Printf("diff %v", diff)
		t.Errorf("Card learned once (good) should be due in 10 min not in %v", dueIn)
	}

	in10Min := time.Now().Add(time.Duration(1000 * 1000 * 1000 * 60 * 10))
	learnAgain := testModule.Service.fsrs.Repeat(newScene.Content[0].Card, in10Min)[fsrs.Good].Card

	if learnAgain.Due.Sub(in10Min) < time.Duration(1_000_000_000*60*60*24) {
		t.Errorf("Card learned a second time (good) should be due in at least one day not %v", learnAgain.Due.Sub(in10Min))
	}
}

func TestService_GetAll(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	scenes, err := tm.Service.GetAll()
	if err != nil {
		t.Errorf("GetAll failed with: %v", err)
	}
	if len(scenes) == 0 {
		t.Errorf("GetAll returned no scenes")
	}
	assert.Equal(t, scenes[0].Name, "Testscene")
}

func TestService_GetByIndex(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	scene, err := tm.Service.GetByIndex(dummyScene.Index)
	if err != nil {
		t.Errorf("GetByIndex failed with: %v", err)
	}
	assert.Equal(t, scene.Name, dummyScene.Name)
}

func TestService_UpdateLine(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	newLine := Line{
		Character: "ROMEO",
		Line:      "Wherefore art thou Romeo?",
		Card:      fsrs.NewCard(),
	}
	scene, err := tm.Service.UpdateLine(1, newLine, dummyScene.Index)
	if err != nil {
		t.Errorf("UpdateLine failed with: %v", err)
	}
	assert.Equal(t, scene.Content[1].Line, newLine.Line)
}

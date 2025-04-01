package scene

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine_ToDto(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	line := dummyScene.Content[0]
	dto := line.ToDto(tm.Service.fsrs)
	assert.Equal(t, line.Character, dto.Character)
	assert.Equal(t, line.Line, dto.Line)
	assert.Equal(t, line.Card.Due, dto.DueTime)
	assert.Equal(t, "1m0s", dto.ReviewTimes.Again)
	assert.Equal(t, "5m0s", dto.ReviewTimes.Hard)
	assert.Equal(t, "10m0s", dto.ReviewTimes.Good)
	assert.Equal(t, "16 days", dto.ReviewTimes.Easy)
}

func TestSceneDtoFromScene(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	scene := dummyScene
	dto := scene.ToDto(tm.Service.fsrs)
	assert.Equal(t, scene.Name, dto.Name)
	assert.Equal(t, scene.Identifier, dto.Identifier)
	assert.Equal(t, scene.Index, dto.Index)
	assert.Equal(t, len(scene.Content), len(dto.Content))
	assert.Equal(t, scene.Content[0].Character, dto.Content[0].Character)
	assert.Equal(t, scene.Content[0].Line, dto.Content[0].Line)
	assert.Equal(t, scene.Content[0].Card.Due, dto.Content[0].DueTime)
	assert.Equal(t, "1m0s", dto.Content[0].ReviewTimes.Again)
	assert.Equal(t, "5m0s", dto.Content[0].ReviewTimes.Hard)
	assert.Equal(t, "10m0s", dto.Content[0].ReviewTimes.Good)
	assert.Equal(t, "16 days", dto.Content[0].ReviewTimes.Easy)
}

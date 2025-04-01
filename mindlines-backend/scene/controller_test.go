package scene

import (
	"github.com/fredericobormann/mindlines-web/mindlines-backend/helper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_GetSceneList(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	router := setupRouter(tm)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/scenes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res, err := helper.Unmarshal[[]MetaScene](w.Body.Bytes())
	if assert.NoError(t, err) {
		name := (*res)[0].Name
		assert.Equal(t, "Testscene", name)
	}
}

func TestController_GetScene(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	router := setupRouter(tm)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/scenes/42", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res, err := helper.Unmarshal[SceneDto](w.Body.Bytes())
	if assert.NoError(t, err) {
		name := res.Name
		assert.Equal(t, "Testscene", name)
	}
}

func TestController_LearnLine(t *testing.T) {
	tm := setupTest()
	defer cleanUp()
	router := setupRouter(tm)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/scenes/42?lineIndex=0&rating=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res, err := helper.Unmarshal[SceneDto](w.Body.Bytes())
	if assert.NoError(t, err) {
		name := res.Name
		assert.Equal(t, "Testscene", name)
		assert.Equal(t, "4 days", res.Content[0].ReviewTimes.Good)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/scenes/42", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res, err = helper.Unmarshal[SceneDto](w.Body.Bytes())
	if assert.NoError(t, err) {
		assert.Equal(t, "4 days", res.Content[0].ReviewTimes.Good)
	}
}

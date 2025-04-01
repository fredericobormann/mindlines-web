package scene

import (
	"github.com/fredericobormann/mindlines-web/mindlines-backend/helper"
	"github.com/gin-gonic/gin"
	"github.com/open-spaced-repetition/go-fsrs/v3"
	"net/http"
	"strconv"
)

type Controller struct {
	Service Service
	FSRS    fsrs.FSRS
}

func (controller *Controller) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/scenes", controller.GetSceneList)
	router.GET("/scenes/:id", controller.GetScene)
	router.POST("/scenes/:id", controller.LearnLine)
}

func (controller *Controller) GetSceneList(c *gin.Context) {
	scenes, err := controller.Service.GetAll()
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, helper.Map(scenes, func(el MetaScene) MetaSceneDto {
		return el.ToDto()
	}))
}

func (controller *Controller) GetScene(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	scene, err := controller.Service.GetByIndex(uint8(id))
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, scene.ToDto(controller.FSRS))
}

func (controller *Controller) UpdateLine(c *gin.Context) {
	idParam := c.Param("id")
	//lineNumberParam := c.Query("lineNumber")
	_, err := strconv.Atoi(idParam)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//lineNumber, err := strconv.Atoi(lineNumberParam)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//controller.Service.UpdateLine(lineNumber, )
}

func (controller *Controller) LearnLine(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	indexParam := c.Query("lineIndex")
	index, err := strconv.Atoi(indexParam)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ratingParam := c.Query("rating")
	rating, err := strconv.Atoi(ratingParam)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	scene, err := controller.Service.LearnLine(uint16(index), fsrs.Rating(rating), uint8(id))
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, scene.ToDto(controller.FSRS))
}

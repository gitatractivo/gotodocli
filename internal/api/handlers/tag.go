package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitatractivo/gotodocli/internal/models"
)

// create a tag and return tag check if tag already exists and return the tag and an handler to get alll the tags


func CreateTag(c *gin.Context) {
	tag := models.Tag{
		Name: c.PostForm("name"),
	}
	c.JSON(http.StatusOK, tag)
}

func GetTags(c *gin.Context) {
	tag := models.Tag{
		Name: c.PostForm("name"),
	}
	c.JSON(http.StatusOK, tag)
}

func GetTag(c *gin.Context) {
	tag := models.Tag{
		Name: c.PostForm("name"),
	}
	c.JSON(http.StatusOK, tag)
}

func UpdateTag(c *gin.Context) {
	tag := models.Tag{
		Name: c.PostForm("name"),
	}
	c.JSON(http.StatusOK, tag)
}

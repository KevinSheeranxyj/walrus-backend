package api

import (
	"github.com/gin-gonic/gin"
	"kevinsheeran/walrus-backend/model"
	"kevinsheeran/walrus-backend/result"
	"log"
)

// CreateForm
// @router /api/v1/create-form [post]
func CreateForm(c *gin.Context) {
	// Bind json parameters
	var dto model.CreateFormDto
	req := c.BindJSON(&dto)
	log.Printf("request body: %+v", req)
	// TODO: Check form duplication

	// Call walrus API

	result.Success(c, true)

}

func GetFormInfo(address string, epochs string, blob []byte) (string, error) {

	return "", nil
}

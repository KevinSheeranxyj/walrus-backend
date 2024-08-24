package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"kevinsheeran/walrus-backend/model"
	"kevinsheeran/walrus-backend/result"
	"net/http"
	"strings"
)

// CreateForm
// @Summary CreateForm API
// @Tags CreateForm
// @Produce json
// @Param data body model.CreateFormDto true "data"
// @Success 200 {object} result.Result
// @router /api/v1/create-form [post]
func CreateForm(c *gin.Context) {
	// Bind json parameters
	var dto model.CreateFormDto
	if err := c.BindJSON(&dto); err != nil {
		result.Failed(c, http.StatusBadRequest, "Invalid input")
	}
	// TODO: Check form duplication

	// Call walrus publisher API
	response, err := callWalrusPublisher(&dto)
	if err != nil {
		result.Failed(c, http.StatusInternalServerError, err.Error())
		return
	}

	//SurveyHandler()
	SurveyCallbackHandler()

	result.Success(c, response)

}

// GetForm
// @Summary GetForm API
// @Tags GetForm
// @Produce json
// @Param blobId path string true "Blob ID"
// @Success 200 {object} result.Result
// @Failure 400 {string} string "Invalid blob ID"
// @router /api/v1/form/{blobId} [get]
func GetForm(c *gin.Context) {

	// Bind json parameters
	var dto model.CreateFormDto
	if err := c.BindJSON(&dto); err != nil {
		result.Failed(c, http.StatusBadRequest, "Invalid input")
	}
	// TODO: Check form duplication
	// call aggregator API

	response, erro := callWalrusAggregator(&dto)
	if erro != nil {
		result.Failed(c, http.StatusInternalServerError, erro.Error())
		return
	}

	result.Success(c, response)
}

func callWalrusPublisher(data *model.CreateFormDto) (string, error) {
	url := "https://publisher-devnet.walrus.space/v1/store"

	// Convert the DTO data to JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshalling data to JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		return "", fmt.Errorf("error executing request: %v", error)
	}
	defer response.Body.Close()

	// Read the response body
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return "", fmt.Errorf("error reading response: %v", error)
	}

	return string(body), nil
}

func callWalrusAggregator(data *model.CreateFormDto) (string, error) {
	url := fmt.Sprintf("https://aggregator-devnet.walrus.space/v1/%s", data.BlobId)

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return "", fmt.Errorf("error creating request: %v", error)
	}

	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		return "", fmt.Errorf("error executing request: %v", error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return "", fmt.Errorf("error reading response: %v", error)
	}

	return string(body), nil
}

func GetFormInfo(address string, epochs string, blob []byte) (string, error) {

	return "", nil
}

package api

import (
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

	// Call walrus API
	response, err := callWalrusPublisher(&dto)
	if err != nil {
		result.Failed(c, http.StatusInternalServerError, err.Error())
		return
	}
	result.Success(c, response)

}

func callWalrusPublisher(data *model.CreateFormDto) (string, error) {
	url := "https://publisher-devnet.walrus.space/v1/store"

	// Convert the DTO data to application/x-www-form-urlencoded format
	formData := fmt.Sprintf("name=%s&description=%s&value=%s", data.Title, data.Description, data.ItemList)

	req, error := http.NewRequest("PUT", url, strings.NewReader(formData))
	if error != nil {
		return "", fmt.Errorf("error creating request: %v", error)
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

func GetFormInfo(address string, epochs string, blob []byte) (string, error) {

	return "", nil
}

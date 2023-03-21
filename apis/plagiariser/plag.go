package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Detected struct {
	Link    string `json:"link"`
	Count   int    `json:"count"`
	Percent int    `json:"percent"`
}

type Response struct {
	Sources     []Detected `json:"sources"`
	PlagPercent int        `json:"plagPercent"`
}

type data struct {
	Data string `json:"data"`
}

func makeGetRequest(data string) (*Response, error) {
	key := "549dfd3788a41935a7ca8dea30588630"
	URL := "https://www.prepostseo.com/apis/checkPlag"

	body := fmt.Sprintf("key=%s&data=%s", key, data)

	query, err := url.ParseQuery(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.PostForm(URL, query)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func postPlagChecker(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	var content data

	if err := c.BindJSON(&content); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not unmarshal error.")
		return
	}

	resp, err := makeGetRequest(content.Data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Request processing error.")
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func main() {
	router := gin.Default()

	router.POST("/apis/plagiarism", postPlagChecker)

	router.Run("localhost:8083")
}

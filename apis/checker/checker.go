package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type email struct {
	Email string `json:"email"`
}

func makeGetRequest() (string, error) {
	resp, err := http.Get("http://127.0.0.1:3001/api/students")
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func checkStudentEmail(email string) (bool, error) {
	resp, err := makeGetRequest()
	if err != nil {
		return false, err
	}
	return strings.Contains(resp, email), nil
}

func postChecker(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	var email email

	err := c.BindJSON(&email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not unmarshal JSON.")
		return
	}

	isStudent, err := checkStudentEmail(email.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not access service.")
		return
	}
	if !isStudent {
		c.IndentedJSON(http.StatusBadRequest, "Student does not exist.")
		return
	}
	c.IndentedJSON(http.StatusOK, "Student exists.")
}

func main() {
	router := gin.Default()

	router.POST("/apis/check", postChecker)

	router.Run("localhost:8082")
}

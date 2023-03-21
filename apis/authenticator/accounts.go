package authenticator

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/MrAndreiL/go-fullstack-app/database"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func insertAccount(acc *Account) error {
	db := database.GetDbConnection()

	cmd := "INSERT INTO accounts (name, username, email, password) VALUES (?, ?, ?, ?)"

	if _, err := db.Exec(cmd, acc.Name, acc.Username, acc.Email, acc.Password); err != nil {
		return errors.New("Error occurred when inserting.")
	}
	return nil
}

func checkIfExists(username string) (bool, error) {
	db := database.GetDbConnection()

	cmd := "SELECT * FROM accounts WHERE username = ?"

	res := db.QueryRow(cmd, username)

	var acc Account
	err := res.Scan(&acc.ID, &acc.Name, &acc.Username, &acc.Email, &acc.Password)

	if err == nil {
		// user already exists.
		return true, errors.New("User already exists.")
	}
	if err != sql.ErrNoRows {
		return false, errors.New("Error occurred when checking existence.")
	}
	return false, nil
}

func checkPassword(username, password string) (bool, error) {
	db := database.GetDbConnection()

	cmd := "SELECT * FROM accounts WHERE username = ?"

	res := db.QueryRow(cmd, username)

	var acc Account
	err := res.Scan(&acc.ID, &acc.Name, &acc.Username, &acc.Email, &acc.Password)
	if err != nil {
		// user already exists.
		return false, nil
	}
	return acc.Password == password, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PostAuth(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	var account Account

	if err := c.BindJSON(&account); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not unmarshal json.")
		return
	}

	exists, err := checkIfExists(account.Username)
	if exists {
		c.IndentedJSON(http.StatusConflict, "User already exists")
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if err = insertAccount(&account); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, account)
}

func PostLogin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	var account Account

	if err := c.BindJSON(&account); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not unmarshal json.")
		return
	}

	exists, err := checkIfExists(account.Username)
	if !exists && err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if !exists && err == nil {
		c.IndentedJSON(http.StatusBadRequest, "User does not exist")
		return
	}
	// check password
	exists, err = checkPassword(account.Username, account.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, "User does not exist")
		return
	}

	session := sessions.Default(c)
	session.Set("test_cookie", 1)
	session.Save()
	c.IndentedJSON(http.StatusOK, "Login successful")
}

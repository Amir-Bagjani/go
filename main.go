package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ----- User type -----
type User struct {
	Name string
}

// ----- Rich Error -----
type RichError struct {
	Message   string
	MetaData  map[string]string
	Operation string
	Time      time.Time
}

func (r RichError) Error() string {
	return r.Message
}

// ----- Simple Error -----
type SimpleError struct {
	Operation string
	Msg       string
}

func (r SimpleError) Error() string {
	return r.Msg
}

// ----- Custom Logger -----
type CustomLogger struct {
	Errors []RichError
}

func (c *CustomLogger) Append(e error) {
	if err, ok := e.(*RichError); ok {
		c.Errors = append(c.Errors, *err)
	} else if err, ok := e.(*SimpleError); ok {
		c.Errors = append(c.Errors, RichError{
			Operation: err.Operation,
			Message:   err.Msg,
			Time:      time.Now(),
			MetaData:  nil,
		})
	} else {
		c.Errors = append(c.Errors, RichError{
			Message:   e.Error(),
			MetaData:  nil,
			Operation: "unknown",
			Time:      time.Now(),
		})
	}
}
func (c *CustomLogger) AppendAndPrint(e error) {
	c.Append(e)

	fmt.Println(e.Error())
}
func (c *CustomLogger) Save() {
	f, _ := os.OpenFile("errors.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	data, _ := json.Marshal(c.Errors)
	f.Write(data)
}

// ----- main -----

func main() {
	logger := CustomLogger{}

	//case number one
	_, fErr := findByID(0)
	if fErr != nil {
		logger.AppendAndPrint(fErr)
	}

	//case number two
	_, tErr := findUserByID(0)
	if tErr != nil {
		logger.AppendAndPrint(tErr)
	}

	//case number three
	_, oErr := os.OpenFile("da/t/sa/sd", os.O_RDONLY, 0644)
	if oErr != nil {
		logger.AppendAndPrint(oErr)
	}

	logger.Save()
}

func findByID(id int) (User, error) {
	if id <= 0 {
		return User{}, &RichError{
			Message: "id can not be 0",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "findByID",
			Time:      time.Now(),
		}
	}

	return User{}, nil
}

func findUserByID(id int) (User, error) {
	if id <= 0 {
		return User{}, &SimpleError{
			Msg:       "simple error id can not be zero",
			Operation: "findUserByID",
		}
	}

	return User{}, nil
}

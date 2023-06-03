package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func InsertStudentMock() {
	student := models.Student{Name: "Student Test", CPF: "12345678900", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestShouldFindAllStudentsWithSuccess(t *testing.T) {
	database.ConnectDatabase()
	InsertStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.GET("/students", controllers.FindAllStudents)
	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShouldFindStudentByDocumentNumberWithSuccess(t *testing.T) {
	database.ConnectDatabase()
	InsertStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.GET("/students/documentNumber/:cpf", controllers.FindStudentByCPF)
	request, _ := http.NewRequest("GET", "/students/documentNumber/"+"12345678900", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var responseBody models.Student
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, responseBody.CPF, "12345678900")
	assert.Equal(t, responseBody.RG, "123456789")
}

func TestShouldFindStudentByIdWithSuccess(t *testing.T) {
	database.ConnectDatabase()
	InsertStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.GET("/students/:id", controllers.FindStudentById)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var responseBody models.Student
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	assert.Equal(t, responseBody.CPF, "12345678900")
	assert.Equal(t, responseBody.RG, "123456789")
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShouldDeleteStudentByIdWithSuccess(t *testing.T) {
	database.ConnectDatabase()
	InsertStudentMock()

	r := RoutesSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestShouldUpdateStudentByIdWithSuccess(t *testing.T) {
	database.ConnectDatabase()
	InsertStudentMock()
	defer DeleteStudentMock()

	r := RoutesSetup()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	student := models.Student{Name: "Student updated", CPF: "12345678902", RG: "123456789"}
	jsonValue, _ := json.Marshal(student)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var responseBody models.Student
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	assert.NotEqual(t, responseBody.Name, "Student Test")
	assert.Equal(t, responseBody.Name, "Student updated")
	assert.Equal(t, responseBody.CPF, "12345678902")
	assert.Equal(t, responseBody.RG, "123456789")
	assert.Equal(t, http.StatusOK, response.Code)
}

package controller_test

import (
	"encoding/json"
	"net/http/httptest"
	"premium-harbor/controller"
	"premium-harbor/test"
	"strings"
	"testing"
)

var userSuite = NewUserControllerTestSuite()

func TestEmptyUsersList(t *testing.T) {
	userSuite.SetTesting(t)
	userSuite.ExpectUserList(`{"status":"ok","body":{"users":[]}}`)
}

func TestRegisterUser(t *testing.T) {
	userSuite.SetTesting(t)
	userSuite.ExpectUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		getOkResponse(),
	)
	userSuite.ExpectUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1"}]}}`)

	userSuite.ExpectUserRegister(
		`{"email": "testik2@mail.ru","name": "Test The Best 2","password": "ssc-tuatara"}`,
		getOkResponse(),
	)
	userSuite.ExpectUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1"},{"id":2,"email":"testik2@mail.ru","name":"Test The Best 2"}]}}`)
}

func TestRegisterUserWithExistingEmail(t *testing.T) {
	userSuite.ExpectUserRegister(
		`{"email": "testik2@mail.ru","name": "Duplicate Email","password": "ssc-tuatara"}`,
		getErrorResponse("user with email testik2@mail.ru already registered"),
	)
	userSuite.ExpectUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1"},{"id":2,"email":"testik2@mail.ru","name":"Test The Best 2"}]}}`)
}

func TestUserUpdate(t *testing.T) {
	userSuite.SetTesting(t)
	userSuite.ExpectUserUpdate(`{"id":1,"name":"New Name"}`, getOkResponse())
	userSuite.ExpectUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"New Name"},{"id":2,"email":"testik2@mail.ru","name":"Test The Best 2"}]}}`)
}

type UserControllerTestSuite struct {
	test.ControllerTestSuite
	userController *controller.UserController
}

func NewUserControllerTestSuite() *UserControllerTestSuite {
	return &UserControllerTestSuite{
		userController: controller.NewUserController(),
	}
}

func (s UserControllerTestSuite) ExpectUserList(expectedResponse string) {
	s.Request = httptest.NewRequest("GET", controller.ApiV1UserListPath, nil)
	s.Response = httptest.NewRecorder()
	s.userController.HandleUserList(s.Response, s.Request)
	s.TestResponse(expectedResponse)
}

func (s UserControllerTestSuite) ExpectUserRegister(requestBody, expectedResponse string) {
	s.Request = httptest.NewRequest("POST", controller.ApiV1UserRegisterPath, strings.NewReader(requestBody))
	s.Response = httptest.NewRecorder()
	s.userController.HandleUserRegister(s.Response, s.Request)
	s.TestResponse(expectedResponse)
}

func (s UserControllerTestSuite) ExpectUserUpdate(requestBody, expectedResponse string) {
	s.Request = httptest.NewRequest("POST", controller.ApiV1UserUpdatePath, strings.NewReader(requestBody))
	s.Response = httptest.NewRecorder()
	s.userController.HandleUserUpdate(s.Response, s.Request)
	s.TestResponse(expectedResponse)
}

func getOkResponse() string {
	response := controller.APIResponse{
		Status: "ok",
		Body:   map[string]interface{}{},
	}
	result, _ := json.Marshal(response)
	return string(result)
}

func getErrorResponse(message string) string {
	response := controller.APIResponse{
		Status: "error",
		Body: map[string]interface{}{
			"message": message,
		},
	}
	result, _ := json.Marshal(response)
	return string(result)
}
package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ControllerTestSuite struct {
	T        *testing.T
	Request  *http.Request
	Response *httptest.ResponseRecorder
}

func (s *ControllerTestSuite) SetTesting(t *testing.T) {
	s.T = t
}

func (s ControllerTestSuite) TestResponse(expectedResponse string) {
	s.TestResponseStatus()
	s.TestResponseBody(expectedResponse)
}

func (s ControllerTestSuite) TestResponseBody(expectedResponse string) {
	actualBody := s.Response.Body.String()
	if strings.TrimSpace(expectedResponse) != strings.TrimSpace(actualBody) {
		s.T.Errorf("\nexpected response body:\n%v\ngot:\n%v", expectedResponse, actualBody)
	}
}

func (s ControllerTestSuite) TestResponseStatus() {
	if s.Response.Code != http.StatusOK {
		s.T.Errorf("expected status code 200, got %v", s.Response.Code)
	}
}

func (s ControllerTestSuite) GetResponseBody() (map[string]*json.RawMessage, error) {
	var response map[string]*json.RawMessage
	err := json.NewDecoder(s.Response.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	var responseBody map[string]*json.RawMessage
	err = json.Unmarshal(*response["body"], &responseBody)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

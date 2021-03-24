package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"twit/servers"
)

func TestRegisterUser(t *testing.T) {
	testingServer := httptest.NewServer(servers.SetupServer())

	defer testingServer.Close()

	requestBody, err := json.Marshal(map[string]string{
		"username": "user1",
		"email":    "user1@example.com",
		"password": "user1",
	})
	response, err := http.Post(fmt.Sprintf("%s/user/register", testingServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", response.StatusCode)
	}
}

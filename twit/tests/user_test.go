package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"twit/models"
	"twit/servers"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSingleUserPositive(t *testing.T) {
	assert := assert.New(t)

	router, cleanUpExecutor := servers.SetupTestingServer()

	testingServer := httptest.NewServer(router)
	defer testingServer.Close()

	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"email":    "email",
		"password": "password",
	})
	assert.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	assert.Equal(http.StatusOK, response.StatusCode)

	defer response.Body.Close()
	body := models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	assert.Equal("Success to register user", body.Message)
	assert.Equal(true, body.Success)

	cleanUpExecutor.TruncateTable([]string{"users"})
}

// func TestRegisterTheSameUserTwice(t *testing.T) {
// 	testingServer := httptest.NewServer(servers.SetupServer())
// 	defer testingServer.Close()

// 	requestBody, err := json.Marshal(map[string]string{
// 		"username": "username",
// 		"email":    "email",
// 		"password": "password",
// 	})
// 	response, err := http.Post(fmt.Sprintf("%s/user/register", testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
// }

// type userCase struct {
// 	user       map[string]string
// 	expectFail bool
// }

// func generateRegisterUserRequest(username, email, password string) map[string]string {
// 	return map[string]string{
// 		"username": username,
// 		"email":    email,
// 		"password": password,
// 	}
// }

// func registerUserCycle(t *testing.T, url string, userMap map[string]string, expectFail bool) {
// 	requestBody, err := json.Marshal(userMap)
// 	response, err := http.Post(fmt.Sprintf("%s/user/register", url), "application/json", bytes.NewBuffer(requestBody))

// 	if !expectFail {
// 		if err != nil {
// 			t.Fatalf("Input value: %v, Expected no error, got %v", userMap, err)
// 		}

// 		if response.StatusCode != http.StatusOK {
// 			t.Fatalf("Input value: %v, Expected status code 200, got %v", userMap, response.StatusCode)
// 		}
// 	} else {
// 		if response.StatusCode == http.StatusOK {
// 			t.Fatalf("Input value: %v, Expected status code !200, got %v", userMap, response.StatusCode)
// 		}
// 	}
// }

// func TestRegisterUser(t *testing.T) {
// 	testingServer := httptest.NewServer(servers.SetupServer())

// 	defer testingServer.Close()

// 	cases := []userCase{
// 		{user: generateRegisterUserRequest("user1", "user1@example.com", "user1"), expectFail: false},
// 		{user: generateRegisterUserRequest("user1", "user1@example.com", "user1"), expectFail: true},
// 		{user: generateRegisterUserRequest("user2", "user1@example.com", ""), expectFail: true},
// 		{user: generateRegisterUserRequest("user3", "", "user3"), expectFail: true},
// 		{user: generateRegisterUserRequest("user4", "", ""), expectFail: true},
// 		{user: generateRegisterUserRequest("user5", "user5@example.com", "user5"), expectFail: false},
// 	}

// 	for _, cs := range cases {
// 		registerUserCycle(t, testingServer.URL, cs.user, cs.expectFail)
// 	}
// }

// func generateLoginUserRequest(email, password string) map[string]string {
// 	return map[string]string{
// 		"email":    email,
// 		"password": password,
// 	}
// }

// func loginUserCycle(t *testing.T, url string, userMap map[string]string, expectFail bool) {
// 	requestBody, err := json.Marshal(userMap)
// 	target := fmt.Sprintf("%s/user/login", url)
// 	fmt.Println("target", target)
// 	response, err := http.Post(target, "application/json", bytes.NewBuffer(requestBody))

// 	if !expectFail {
// 		if err != nil {
// 			t.Fatalf("Input value: %v, Expected no error, got %v", userMap, err)
// 		}

// 		if response.StatusCode != http.StatusOK {
// 			t.Fatalf("Input value: %v, Expected status code 200, got %v", userMap, response.StatusCode)
// 		}
// 	} else {
// 		if response.StatusCode == http.StatusOK {
// 			t.Fatalf("Input value: %v, Expected status code !200, got %v", userMap, response.StatusCode)
// 		}
// 	}
// }

// func TestLoginUser(t *testing.T) {
// 	testingServer := httptest.NewServer(servers.SetupServer())

// 	defer testingServer.Close()

// 	cases := []userCase{
// 		{user: generateLoginUserRequest("user1@example.com", "user1"), expectFail: false},
// 		{user: generateLoginUserRequest("user1@example.com", ""), expectFail: true},
// 		{user: generateLoginUserRequest("", "user3"), expectFail: true},
// 		{user: generateLoginUserRequest("", ""), expectFail: true},
// 		{user: generateLoginUserRequest("user5@example.com", "user5"), expectFail: false},
// 	}

// 	for _, cs := range cases {
// 		loginUserCycle(t, testingServer.URL, cs.user, cs.expectFail)
// 	}
// }

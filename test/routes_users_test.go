package test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"

)


func TestRoutesUsers(t *testing.T) {

	mockResponse := `{"data":[],"total_rows":0}`
	
	res, err := http.Get("http://localhost:8080/api/users")
	if err != nil {
		t.Errorf("Route /api/users not found")
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, mockResponse, string(responseData))

}

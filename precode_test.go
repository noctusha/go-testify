package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"strconv"
)

type ResponseRecorder struct {
	Code int
	HeaderMap http.Header
	Body *bytes.Buffer
	Flushed bool
}

func GetARespRec (url string) (req *http.Request, responseRecorder *httptest.ResponseRecorder) {
	req = httptest.NewRequest("GET", url, nil)
	responseRecorder = httptest.NewRecorder()
	return req, responseRecorder
}

func TestMainHandlerWhenOK(t *testing.T) {
	req, responseRecorder := GetARespRec("/cafe?count=2&city=moscow")


	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)


	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return
	}
	
	cnt := 1
	for i := 0; i < len(responseRecorder.Body.String()); i++ {
		if responseRecorder.Body.String()[i] == ',' {
			cnt++
		}
	}
	assert.Equal(t, count, cnt)
	require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req, responseRecorder := GetARespRec("/cafe?count=2&city=Bryansk")

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountIsMoreThanMax(t *testing.T) {
	req, responseRecorder := GetARespRec("/cafe?count=11&city=moscow")

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)


	cnt := 1
	for i := 0; i < len(responseRecorder.Body.String()); i++ {
		if responseRecorder.Body.String()[i] == ',' {
			cnt++
		}
	}

	assert.Equal(t, cnt, 4)
}

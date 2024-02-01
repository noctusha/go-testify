package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, status, http.StatusOK)
	require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=Bryansk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountIsMoreThanMax(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=11&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := "Мир кофе, Сладкоежка, Кофе и завтраки, Сытый студент"

	assert.Equal(t, responseRecorder.Body.String(), expected)
}

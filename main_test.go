package main

import (
	"log"
	"testing"
	"net/http"
	"net/http/httptest"
	"product-api/config"

	"github.com/stretchr/testify/assert"
)

var routers = config.Routers

func TestRootIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println(err)
	} else {
		routers.ServeHTTP(rec, req)
		code := http.StatusOK
		assert.Equal(t, code, rec.Code)
		assert.Equal(t, http.StatusText(code), rec.Body.String())
	}
}

func TestHealthCheckIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		log.Println(err)
	} else {
		routers.ServeHTTP(rec, req)
		codeOK := http.StatusOK
		assert.Equal(t, codeOK, rec.Code)
		assert.Equal(t, http.StatusText(codeOK), rec.Body.String())
		codeBadRequest := http.StatusBadRequest
		assert.NotEqual(t, codeBadRequest, rec.Code)
		assert.NotEqual(t, http.StatusText(codeBadRequest), rec.Body.String())
	}
}

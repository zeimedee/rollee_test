package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeimedee/test2/internal/handlers"
	"github.com/zeimedee/test2/internal/services"

	"github.com/gin-gonic/gin"
)

func TestWordHandler_StoreWord(t *testing.T) {
	gin.SetMode(gin.TestMode)

	wordService := services.NewWordService()
	wordHandler := handlers.NewWordHandler(wordService)

	router := gin.New()
	router.POST("/service/store", wordHandler.StoreWord)

	request, _ := http.NewRequest("POST", "/service/store?word=animal", nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
	}

	expectedBody := `{"message":"animal Stored Successfully"}`

	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected response body '%s', but got '%s'", expectedBody, body)
	}

	if count := wordService.GetFrequentWord("animal"); count != "animal" {
		t.Errorf("Expected most frequent word for prefix 'animal' to be 'animal', but got '%s'", count)
	}
}

func TestWordHandler_GetMostFrequentWord(t *testing.T) {
	gin.SetMode(gin.TestMode)

	wordService := services.NewWordService()
	wordService.StoreWord("abc")
	wordService.StoreWord("ab")
	wordService.StoreWord("ab")

	wordHandler := handlers.NewWordHandler(wordService)

	router := gin.New()
	router.GET("/service/retrieve", wordHandler.RetrieveWord)

	request, _ := http.NewRequest("GET", "/service/retrieve?prefix=a", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
	}

	expectedBody := `{"word":"ab"}`

	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected response body '%s', but got '%s'", expectedBody, body)
	}

	request, _ = http.NewRequest("GET", "/service/retrieve?prefix=d", nil)
	response = httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, response.Code)
	}
}

package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Teste para a rota base "/"
func TestBaseRoute(t *testing.T) {
	gin.SetMode(gin.TestMode) // Modo de teste do Gin
	router := gin.Default()
	baseRoutes(router) // Registra a rota de base

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"resp":"ok"`) // Verifica se a resposta contém a chave "resp"
}

// Teste para a rota base "/health"
func TestHealthRoute(t *testing.T) {
	gin.SetMode(gin.TestMode) // Modo de teste do Gin
	router := gin.Default()
	baseRoutes(router) // Registra a rota de base

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"resp":"ok"`) // Verifica se a resposta contém a chave "resp"
}

// Teste para a rota "/api/v1/"
func TestV1Route(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	apiRoutes(router) // Registra as rotas da API

	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"resp":"ok"`)
}

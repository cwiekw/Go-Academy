package movie

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func newMovieWithIDHandler() *MovieWithIDHandler {
	return NewMMovieWithIDHandler(zap.NewNop(), newMockMovieDataBase())
}

func TestMovieWithIDHandler_ServeHTTP_WrongIDFormat(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "wrong")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Endpoint /movies with wrong format ID should return HTTP 400 regardless of method")
}

func TestMovieWithIDHandler_ServeHTTP_GET(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	expectedBody := `{"Id":1,"Name":"Movie1","Year":2001}` + "\n"

	assert.Equal(t, http.StatusOK, rr.Code, "Endpoint /movies/1 should return movie for GET method")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, expectedBody, rr.Body.String())
}

func TestMovieWithIDHandler_ServeHTTP_GET_NotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Endpoint /movies/999 should return HTTP 404 for GET method")
}

func TestMovieWithIDHandler_ServeHTTP_PUT(t *testing.T) {
	body := []byte(`{"Name":"Movie1","Year":2001}`)
	req, err := http.NewRequest(http.MethodPut, "/movies", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "Endpoint /movies/1 should return HTTP 204 for PUT method")
}

func TestMovieWithIDHandler_ServeHTTP_PUT_NotFound(t *testing.T) {
	body := []byte(`{"Name":"Movie1","Year":2001}`)
	req, err := http.NewRequest(http.MethodPut, "/movies", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Endpoint /movies/999 should return HTTP 404 for PUT method")
}

func TestMovieWithIDHandler_ServeHTTP_DELETE(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "Endpoint /movies/1 should return HTTP 204 for DELETE method")
}

func TestMovieWithIDHandler_ServeHTTP_DELETE_NotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Endpoint /movies/999 should return HTTP 404 for DELETE method")
}

func TestMovieWithIDHandler_ServeHTTP_POST_NotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newMovieWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Endpoint /movies/1 should not allow POST method")
}

package movie

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func newMovieHandler() *MovieHandler {
	return NewMovieHandler(zap.NewNop(), newMockMovieDataBase())
}

func TestMovieHandler_ServeHTTP_GET(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newMovieHandler()

	handler.ServeHTTP(rr, req)

	expectedBody := `[{"Id":1,"Name":"Movie1","Year":2001},{"Id":2,"Name":"Movie2","Year":2002}]` + "\n"

	assert.Equal(t, http.StatusOK, rr.Code, "Endpoint /movies should return list of movies for GET method")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, expectedBody, rr.Body.String())
}

func TestMovieHandler_ServeHTTP_POST(t *testing.T) {
	body := []byte(`{"Name":"Movie1","Year":2001}`)
	req, err := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newMovieHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Endpoint /movies should return ID for newly created Movie for POST method")
	assert.Equal(t, "1", rr.Body.String())
}

func TestMovieHandler_ServeHTTP_PUT_NotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPut, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newMovieHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Endpoint /movies should not allow PUT method")
}

func TestMovieHandler_ServeHTTP_DELETE_NotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/movies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newMovieHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Endpoint /movies should not allow DELETE method")
}

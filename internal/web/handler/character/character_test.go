package character

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func newCharacterHandler() *CharacterHandler {
	return NewCharacterHandler(zap.NewNop(), newMockCharacterDataBase())
}

func TestCharacterHandler_ServeHTTP_GET(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newCharacterHandler()

	handler.ServeHTTP(rr, req)

	expectedBody := `[{"Id":1,"Name":"Character1","MovieId":1001},{"Id":2,"Name":"Character2","MovieId":1002}]` + "\n"

	assert.Equal(t, http.StatusOK, rr.Code, "Endpoint /characters should return list of movies for GET method")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, expectedBody, rr.Body.String())
}

func TestCharacterHandler_ServeHTTP_POST(t *testing.T) {
	body := []byte(`{"Name":"Character1","MovieId":1001}`)
	req, err := http.NewRequest(http.MethodPost, "/characters", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newCharacterHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Endpoint /characters should return ID for newly created Movie for POST method")
	assert.Equal(t, "1", rr.Body.String())
}

func TestCharacterHandler_ServeHTTP_PUT_NotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPut, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newCharacterHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Endpoint /characters should not allow PUT method")
}

func TestCharacterHandler_ServeHTTP_DELETE_NotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newCharacterHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Endpoint /characters should not allow DELETE method")
}

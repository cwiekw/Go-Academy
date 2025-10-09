package character

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func newCharacterWithIDHandler() *CharacterWithIDHandler {
	return NewCharacterWithIDHandler(zap.NewNop(), newMockCharacterDataBase())
}

func TestCharacterWithIDHandler_ServeHTTP_WrongIDFormat(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "wrong")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Endpoint /characters with wrong format ID should return HTTP 400 regardless of method")
}

func TestCharacterWithIDHandler_ServeHTTP_GET(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	expectedBody := `{"Id":1,"Name":"Character1","MovieId":1001}` + "\n"

	assert.Equal(t, http.StatusOK, rr.Code, "Endpoint /characters/1 should return character for GET method")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, expectedBody, rr.Body.String())
}

func TestCharacterWithIDHandler_ServeHTTP_GET_NotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Endpoint /characters/999 should return HTTP 404 for GET method")
}

func TestCharacterWithIDHandler_ServeHTTP_UPDATE(t *testing.T) {
	body := []byte(`{"Name":"Character1","MovieId":1001}`)
	req, err := http.NewRequest(http.MethodPut, "/characters", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "Endpoint /characters/1 should return HTTP 204 for PUT method")
}

func TestCharacterWithIDHandler_ServeHTTP_UPDATE_NotFound(t *testing.T) {
	body := []byte(`{"Name":"Character1","MovieId":1001}`)
	req, err := http.NewRequest(http.MethodPut, "/characters", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Endpoint /characters/999 should return HTTP 404 for PUT method")
}

func TestCharacterWithIDHandler_ServeHTTP_DELETE(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "Endpoint /characters/1 should return HTTP 204 for DELETE method")
}

func TestCharacterWithIDHandler_ServeHTTP_DELETE_NotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Endpoint /characters/999 should return HTTP 404 for DELETE method")
}

func TestCharacterWithIDHandler_ServeHTTP_POST_NotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/characters", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := newCharacterWithIDHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Endpoint /characters/1 should not allow POST method")
}

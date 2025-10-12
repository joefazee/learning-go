package user

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(body)
}

func TestHandler(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	Handler(w, req)

	assert.Equal(t, 200, w.Code)
	body, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, `Hello World!`, string(body))

}

func TestEchoHandler(t *testing.T) {

	rr := httptest.NewRecorder()
	reqBody := bytes.NewBufferString(`echo body payload`)
	req := httptest.NewRequest("GET", "/", reqBody)

	EchoHandler(rr, req)

	assert.Equal(t, 200, rr.Code)
	body, err := io.ReadAll(rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, `echo body payload`, string(body))

}

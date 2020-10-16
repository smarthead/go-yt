package rest

import (
	"bytes"
	"encoding/json"
	"go-yt/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestData struct {
	Field string `json:"field"`
}

func Test_ShouldReturn_HttpError_For_BadRequest(t *testing.T) {
	srv := httptest.NewServer(HttpHandler(BadRequest))
	defer srv.Close()

	rest, err := NewRestApiClient(srv.URL, "token")
	if err != nil {
		t.Error(err)
	}

	err = rest.Get("/api", nil, nil)

	if err == nil {
		t.Fatal("Error is nil")
	}

	t.Log(err.(*errors.HttpError).Error())
}

func Test_ShouldReturn_TestData_For_Ok(t *testing.T) {
	srv := httptest.NewServer(HttpHandler(OkRequest))
	defer srv.Close()

	rest, err := NewRestApiClient(srv.URL, "token")
	if err != nil {
		t.Error(err)
	}

	data := &TestData{}
	err = rest.Get("/api", nil, data)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(data.Field)

	if data.Field != "Field" {
		t.Fatal("Wrong response")
	}
}

func HttpHandler(f http.HandlerFunc) http.HandlerFunc {
	return f
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func OkRequest(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(&TestData{Field: "Field"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

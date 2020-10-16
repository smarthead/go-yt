package rest

import (
	"bytes"
	"encoding/json"
	"go-yt/errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

	err = rest.Get("/api", nil, nil, nil)

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
	err = rest.Get("/api", nil, nil, data)

	if err != nil {
		t.Fatal(err)
	}

	if data.Field != "Field" {
		t.Log(data.Field)
		t.Fatal("Wrong response")
	}
}

func Test_ShouldSend_QueryParameters(t *testing.T) {
	srv := httptest.NewServer(HttpHandler(PingPong))
	defer srv.Close()

	rest, err := NewRestApiClient(srv.URL, "token")
	if err != nil {
		t.Error(err)
	}

	expected := map[string]string{
		"key": "value",
	}
	actual := make(map[string]string)
	err = rest.Get("/api", expected, nil, &actual)

	if err != nil {
		t.Fatal(err)
	}

	for k, v := range expected {
		if actual[k] != v {
			t.Log(actual[k], v)
			log.Fatalln("Wrong query parameters")
		}
	}
}

func Test_ShouldSend_Body(t *testing.T) {
	srv := httptest.NewServer(HttpHandler(PingPong))
	defer srv.Close()

	rest, err := NewRestApiClient(srv.URL, "token")
	if err != nil {
		t.Error(err)
	}

	expected := &TestData{Field: "Expected"}
	actual := &TestData{}
	err = rest.Post("/api", expected, nil, actual)

	if err != nil {
		t.Fatal(err)
	}

	if expected.Field != actual.Field {
		t.Log(actual.Field, expected.Field)
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
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

func PingPong(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	var bodyBytes []byte
	if len(query) > 0 {
		queryValues, _ := url.ParseQuery(query)
		queryMap := make(map[string]string)

		for k, v := range queryValues {
			queryMap[k] = strings.Join(v, "")
		}

		buffer := new(bytes.Buffer)
		encoder := json.NewEncoder(buffer)
		if err := encoder.Encode(queryMap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bodyBytes = buffer.Bytes()
	} else {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bodyBytes)
}

package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mengdage/hello-go/apitest/handlers"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		// create a req aka client request
		req, err := http.NewRequest("GET", "/sendjson", nil)

		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotX, err)
		}
		t.Log("\tShould be able to create a request.", checkMark)

		// create a response recorder
		recorder := httptest.NewRecorder()

		// send the request to the server and record response data in recorder
		http.DefaultServeMux.ServeHTTP(recorder, req)

		if recorder.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, recorder.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(recorder.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name != "Meng" || u.Email != "meng@gmail.com" {
			t.Error("\tShould have correct Name and Email.", ballotX)
		} else {
			t.Log("\tShould have correct Name and Email.", checkMark)
		}
	}
}

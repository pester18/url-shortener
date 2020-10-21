package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pester18/url-shortener/config"
	"github.com/pester18/url-shortener/infrastructure/datastore"
	"github.com/pester18/url-shortener/infrastructure/router"
	"github.com/pester18/url-shortener/registry"
)

var handler http.Handler

func init() {
	conf := config.New()

	dialStr := fmt.Sprintf("%s:%s", conf.Mongo.Host, conf.Mongo.Port)
	db := datastore.NewDB(dialStr, conf.Mongo.DB)

	reg := registry.NewRegistry(db)

	handler = router.NewRouter(reg.NewAppController())
}

type input struct {
	url  string
	body string
}

type output struct {
	statusCode int
}

func TestUrlRedirection(t *testing.T) {
	srv := httptest.NewServer(handler)
	defer srv.Close()

	inputs := []input{
		{
			url: "EJb6q6LqTDq5HkKUbBBErn", //redirects to google.com (presaved) so the final status code is 200
		},
		{
			url: "1234",
		},
	}

	expectedOutputs := []output{
		{
			statusCode: http.StatusOK,
		},
		{
			statusCode: http.StatusNotFound,
		},
	}

	for i, inp := range inputs {
		res, err := http.Get(fmt.Sprintf("%s/%s", srv.URL, inp.url))

		if err != nil {
			t.Fatal(err)
		}

		expectedOut := expectedOutputs[i]

		if res.StatusCode != expectedOut.statusCode {
			t.Errorf("Status is: %d and expected to be: %d", res.StatusCode, expectedOut)
		}
	}
}

func TestShortUrlGeneration(t *testing.T) {
	srv := httptest.NewServer(handler)
	defer srv.Close()

	inputs := []input{
		{
			url:  "shorten",
			body: `{"url": "https://www.google.com/"}`,
		},
		{
			url:  "shorten",
			body: `{"abc": "abc"}`,
		},
	}

	expectedOutputs := []output{
		{
			statusCode: http.StatusOK,
		},
		{
			statusCode: http.StatusBadRequest,
		},
	}

	for i, inp := range inputs {
		res, err := http.Post(
			fmt.Sprintf("%s/%s", srv.URL, inp.url),
			"application/json",
			strings.NewReader(inp.body))

		if err != nil {
			t.Fatal(err)
		}

		expectedOut := expectedOutputs[i]

		if res.StatusCode != expectedOut.statusCode {
			t.Errorf("Status is: %d and expected to be: %d", res.StatusCode, expectedOut)
		}
	}
}

func TestShortUrlDeletion(t *testing.T) {
	srv := httptest.NewServer(handler)
	defer srv.Close()

	httpClient := srv.Client()

	inputs := []input{
		{
			url: "f2XfhB2hA4Xfux7Zrss6g5", //previously created url record but it will be deleted after running this test
		},
		{
			url: "1234",
		},
	}

	expectedOutputs := []output{
		{
			statusCode: http.StatusOK,
		},
		{
			statusCode: http.StatusNotFound,
		},
	}

	for i, inp := range inputs {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", srv.URL, inp.url), nil)

		res, err := httpClient.Do(req)

		if err != nil {
			t.Fatal(err)
		}

		expectedOut := expectedOutputs[i]

		if res.StatusCode != expectedOut.statusCode {
			t.Errorf("Status is: %d and expected to be: %d", res.StatusCode, expectedOut)
		}
	}
}

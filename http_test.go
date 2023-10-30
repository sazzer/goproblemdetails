package goproblemdetails_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/sazzer/goproblemdetails"
	"github.com/stretchr/testify/assert"
)

func TestSendEmptyProblem(t *testing.T) {
	t.Parallel()

	problem := goproblemdetails.New(http.StatusBadRequest)

	rec := httptest.NewRecorder()
	err := goproblemdetails.Send(rec, problem)
	assert.NoError(t, err)

	response := rec.Result()
	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	assert.Empty(t, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	assert.Empty(t, body)
}

func TestSendPopulatedProblem(t *testing.T) {
	t.Parallel()

	problem := goproblemdetails.New(http.StatusForbidden,
		goproblemdetails.WithType("https://example.com/probs/out-of-credit", "You do not have enough credit."),
		goproblemdetails.WithDetail("Your current balance is 30, but that costs 50."),
		goproblemdetails.WithInstance("/account/12345/msgs/abc"),
		goproblemdetails.WithValue("balance", 30),
		goproblemdetails.WithValue("accounts", []string{"/account/12345", "/account/67890"}),
	)

	rec := httptest.NewRecorder()
	err := goproblemdetails.Send(rec, problem)
	assert.NoError(t, err)

	response := rec.Result()
	defer response.Body.Close()

	assert.Equal(t, http.StatusForbidden, response.StatusCode)

	assert.Equal(t, []string{"application/problem+json; charset=utf-8"}, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(body), `{
		"type": "https://example.com/probs/out-of-credit",
		"title": "You do not have enough credit.",
		"detail": "Your current balance is 30, but that costs 50.",
		"instance": "/account/12345/msgs/abc",
		"balance": 30,
		"accounts": ["/account/12345", "/account/67890"]
	}`)
}

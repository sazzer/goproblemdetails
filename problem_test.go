package goproblemdetails_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/sazzer/goproblemdetails"
	"github.com/stretchr/testify/assert"
)

func TestRenderEmptyProblem(t *testing.T) {
	t.Parallel()

	problem := goproblemdetails.New(http.StatusBadRequest)

	result, err := json.Marshal(problem)

	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(result), `{}`)
}

func TestRenderPopulatedProblem(t *testing.T) {
	t.Parallel()

	problem := goproblemdetails.New(http.StatusForbidden,
		goproblemdetails.WithType("https://example.com/probs/out-of-credit", "You do not have enough credit."),
		goproblemdetails.WithDetail("Your current balance is 30, but that costs 50."),
		goproblemdetails.WithInstance("/account/12345/msgs/abc"),
		goproblemdetails.WithValue("balance", 30),
		goproblemdetails.WithValue("accounts", []string{"/account/12345", "/account/67890"}),
	)

	result, err := json.Marshal(problem)

	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(result), `{
		"type": "https://example.com/probs/out-of-credit",
		"title": "You do not have enough credit.",
		"detail": "Your current balance is 30, but that costs 50.",
		"instance": "/account/12345/msgs/abc",
		"balance": 30,
		"accounts": ["/account/12345", "/account/67890"]
	}`)
}

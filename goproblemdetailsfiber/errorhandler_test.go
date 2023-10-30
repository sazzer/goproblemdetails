package goproblemdetailsfiber_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kinbiko/jsonassert"
	"github.com/sazzer/goproblemdetails"
	"github.com/sazzer/goproblemdetails/goproblemdetailsfiber"
	"github.com/stretchr/testify/assert"
)

func TestEmptyProblem(t *testing.T) {
	t.Parallel()

	app := fiber.New(fiber.Config{
		ErrorHandler: goproblemdetailsfiber.ErrorHandler,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return goproblemdetails.New(http.StatusBadRequest)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	assert.Empty(t, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	assert.Empty(t, body)
}

func TestPopulatedProblem(t *testing.T) {
	t.Parallel()

	app := fiber.New(fiber.Config{
		ErrorHandler: goproblemdetailsfiber.ErrorHandler,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return goproblemdetails.New(http.StatusForbidden,
			goproblemdetails.WithType("https://example.com/probs/out-of-credit", "You do not have enough credit."),
			goproblemdetails.WithDetail("Your current balance is 30, but that costs 50."),
			goproblemdetails.WithInstance("/account/12345/msgs/abc"),
			goproblemdetails.WithValue("balance", 30),
			goproblemdetails.WithValue("accounts", []string{"/account/12345", "/account/67890"}),
		)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

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

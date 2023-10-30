package goproblemdetailsfiber

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sazzer/goproblemdetails"
)

// ErrorHandler is a custom error handler for Fiber applications that handles HTTP Problem responses.
// It checks if the error is of type goproblemdetails.Problem and, if so, sends a Problem response to the client.
//
// Parameters:
//
//	ctx - The Fiber.Ctx context that represents the current request and response.
//	err - The error that occurred during request processing.
//
// Returns:
//
//	An error if there was an issue encoding and sending the response; otherwise, it returns nil.
//
// Example:
//
//	// Define a Fiber app and set the custom error handler.
//	app := fiber.New(fiber.Config{
//	  ErrorHandler: goproblemdetailsfiber.ErrorHandler,
//	})
//
//	// Use the Problem instance in your application.
//	app.Get("/not-found", func(c *fiber.Ctx) error {
//	    return goproblemdetails.New(http.StatusNotFound, goproblemdetails.WithType("not-found", "Resource Not Found"))
//	})
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var problem goproblemdetails.Problem
	if errors.As(err, &problem) {
		ctx.Response().Header.Del("content-type")
		ctx.Response().ResetBody()

		ctx.Status(problem.StatusCode)

		if len(problem.Payload) > 0 {
			if err := ctx.JSON(problem.Payload); err != nil {
				return err
			}

			ctx.Set("content-type", "application/problem+json; charset=utf-8")
		}

		return nil
	}

	return fiber.DefaultErrorHandler(ctx, err)
}

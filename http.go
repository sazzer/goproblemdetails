package goproblemdetails

import (
	"encoding/json"
	"net/http"
)

// Send sends an HTTP Problem response to the client.
//
// Parameters:
//
//	w - The http.ResponseWriter where the response will be written.
//	problem - The Problem instance containing the details of the problem response.
//
// Returns:
//
//	An error if there was an issue encoding and sending the response; otherwise, it returns nil.
//
// Example:
//
//	// Create a Problem instance with a status code and payload.
//	problem := problem.New(http.StatusNotFound, problem.WithType("not-found", "Resource Not Found"))
//
//	// Send the Problem response to the client.
//	err := goproblemdetails.Send(w, problem)
//	if err != nil {
//	    // Handle the error, e.g., log it or send an alternative response.
//	}
func Send(w http.ResponseWriter, problem Problem) error {
	if len(problem.Payload) > 0 {
		w.Header().Add("content-type", "application/problem+json;charset=utf-8")
		w.WriteHeader(problem.StatusCode)

		return json.NewEncoder(w).Encode(problem)
	}

	w.WriteHeader(problem.StatusCode)

	return nil
}

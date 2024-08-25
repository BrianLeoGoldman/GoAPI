package middleware

import (
	"errors"
	"net/http"

	"github.com/brianleogoldman/goapi/api"
	"github.com/brianleogoldman/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// UnAuthorizedError a custom unauthorized error
var UnAuthorizedError = errors.New("invalid username or token")

// The Authorization function is being used as middleware, so it must have a specific signature (it must receive
// and return a http.Handler)
func Authorization(next http.Handler) http.Handler {
	// ResponseWriter is used to construct the response to the caller (set response body, headers and status code)
	// Request contains all the information about the incoming HTTP request (headers, payload, etc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// We get the username from the request URL
		var username string = r.URL.Query().Get("username")
		// We get the token from the request header
		var token = r.Header.Get("Authorization")
		var err error

		// If either username or token are empty, we return an error
		if username == "" || token == "" {
			// We log an error message
			log.Error(UnAuthorizedError)
			// We create the error and add it to the response: RequestErrorHandler function -> writeError function
			api.RequestErrorHandler(w, UnAuthorizedError)
			// We exit the Authorization function
			return
		}

		// We instantiate a pointer to the database
		var database *tools.DatabaseInterface
		// We call the NewDatabase method
		database, err = tools.NewDatabase()
		// If we cannot get the database
		if err != nil {
			// We return an Internal Error
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		// We query the database
		loginDetails = (*database).GetUserLoginDetails(username)

		// If we didn't find the client with the username or the token doesn't match the one from the database
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			// We log an error message
			log.Error(UnAuthorizedError)
			// We create the error and add it to the response
			api.RequestErrorHandler(w, UnAuthorizedError)
			// We exit the Authorization function
			return
		}

		// Calls the next middleware or the HandlerFunc if there are no more middleware
		// Middleware1 -> next.serveHTTP -> Middleware2 -> ... -> HandlerFunc
		// Authorization (Middleware1) -> next.serveHTPP -> GetCoinBalance (HandlerFunc)
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"encoding/json"
	"net/http"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type CustomRelay struct {
	Schema *graphql.Schema
}

func (c *CustomRelay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := c.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authtoken := *r.Context().Value("authtoken").(*string)
	if len(authtoken) > 1 {
		http.SetCookie(w, &http.Cookie{
			Name:     "authtoken",
			Value:    authtoken,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			MaxAge:   int(24 * time.Hour.Seconds()),
			HttpOnly: true,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

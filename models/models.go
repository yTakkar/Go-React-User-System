package models

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("Shh..It's very secret"))

// MakeTimestamp function
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// GetSession function
func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Fatal(err)
	}
	return session
}

// AllSessions function to return all the sessions
func AllSessions(r *http.Request) (interface{}, interface{}, interface{}, interface{}) {
	session := GetSession(r)
	id := session.Values["id"]
	username := session.Values["username"]
	email := session.Values["email"]
	joined := session.Values["joined"]
	return id, username, email, joined
}

// IsLoggedIn function to check if user is loggedIn
func IsLoggedIn(r *http.Request) bool {
	id, _, _, _ := AllSessions(r)
	var loggedIn bool
	if id == nil {
		loggedIn = false
	} else {
		loggedIn = true
	}
	return loggedIn
}

// JSON function
func JSON(w http.ResponseWriter, r *http.Request, resp interface{}) {
	final, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(final)
}

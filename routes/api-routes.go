package routes

import (
	M "Go-React-User-System/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetSession function
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})

	id, username, email, joined := M.AllSessions(r)
	loggedIn := M.IsLoggedIn(r)

	Response["loggedIn"] = loggedIn
	Response["session"] = &M.MySession{ID: id, Username: username, Email: email, Joined: joined}

	M.JSON(w, r, Response)

}

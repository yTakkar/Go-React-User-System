package routes

import (
	M "Go-React-User-System/models"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/badoux/checkmail"
	"github.com/julienschmidt/httprouter"
)

// Signup function
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})

	username := strings.TrimSpace(r.PostFormValue("username"))
	email := strings.TrimSpace(r.PostFormValue("email"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	passwordAgain := strings.TrimSpace(r.PostFormValue("password_again"))

	db := M.DB()

	mailErr := checkmail.ValidateFormat(email)

	var (
		userCount  int
		emailCount int
	)

	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE username=?", username).Scan(&userCount)
	db.QueryRow("SELECT COUNT(id) AS emailCount FROM users WHERE email=?", email).Scan(&emailCount)

	if username == "" || email == "" || password == "" || passwordAgain == "" {
		Response["mssg"] = "Some values are missing!"
	} else if len(username) < 4 || len(username) > 32 {
		Response["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		Response["mssg"] = "Invalid Format!"
	} else if password != passwordAgain {
		Response["mssg"] = "Passwords don't match"
	} else if userCount > 0 {
		Response["mssg"] = "Username already exists!"
	} else if emailCount > 0 {
		Response["mssg"] = "Email already exists!"
	} else {

		hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if hashErr != nil {
			panic(hashErr)
		}

		_, insErr := db.Exec(
			"INSERT INTO users(username, email, password, joined) VALUES (?, ?, ?, ?)",
			username,
			email,
			hash,
			M.MakeTimestamp(),
		)

		if insErr != nil {
			panic(insErr)
		}

		Response["mssg"] = "You are now registered!!"
		Response["success"] = true

	}

	M.JSON(w, r, Response)

}

// Login function
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})

	rusername := strings.TrimSpace(r.PostFormValue("username"))
	rpassword := strings.TrimSpace(r.PostFormValue("password"))

	db := M.DB()
	var (
		userCount int
		id        int
		username  string
		email     string
		password  string
		joined    string
	)

	db.QueryRow("SELECT COUNT(id) AS userCount, id, username, email, password, joined FROM users WHERE username=?", rusername).Scan(&userCount, &id, &username, &email, &password, &joined)

	encErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(rpassword))

	if rusername == "" || rpassword == "" {
		Response["mssg"] = "Some values are missing!"
	} else if userCount == 0 {
		Response["mssg"] = "Invalid username!"
	} else if encErr != nil {
		Response["mssg"] = "Invalid password!"
	} else {

		session := M.GetSession(r)
		session.Values["id"] = id
		session.Values["username"] = username
		session.Values["email"] = email
		session.Values["joined"] = joined
		session.Save(r, w)

		idS, usernameS, emailS, joinedS := M.AllSessions(r)

		Response["mssg"] = "You are now logged in!!"
		Response["session"] = &M.MySession{ID: idS, Username: usernameS, Email: emailS, Joined: joinedS}
		Response["loggedIn"] = true
		Response["success"] = true

	}

	M.JSON(w, r, Response)

}

// Logout function
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := M.GetSession(r)

	delete(session.Values, "id")
	delete(session.Values, "username")
	delete(session.Values, "email")
	delete(session.Values, "joined")

	session.Save(r, w)
}

package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

// NewUsers creates a newuser controller
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
}

// New renders signup form
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create method processes signup form to create user account
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	signIn(w, &user)
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

type LoginForm struct {
	Email    string `schems:"email"`
	Password string `schema:"password"`
}

// Login is used to verify the email and pw and log the user in if theyre correct
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invlaid email address")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password provided")
	case nil:
		fmt.Println(w, user)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err != nil {
		return
	}

	signIn(w, user)
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

func signIn(w http.ResponseWriter, user *models.User) {
	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}

	http.SetCookie(w, &cookie)
}

// CookieTest is used to display cookies set on current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Email is: ", cookie.Value)
}

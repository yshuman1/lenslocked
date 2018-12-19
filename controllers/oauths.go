package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/gorilla/csrf"
	"golang.org/x/oauth2"
	llctx "lenslocked.com/context"
	"lenslocked.com/models"
)

// NewOAuths is used to create a new OAuths controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewOAuths(os models.OAuthService, configs map[string]*oauth2.Config) *OAuths {
	return &OAuths{
		os:      os,
		configs: configs,
	}
}

type OAuths struct {
	os      models.OAuthService
	configs map[string]*oauth2.Config
}

func (o *OAuths) Connect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	oauthConfig, ok := o.configs[service]
	if !ok {
		http.Error(w, "invalid OAuth2 service", http.StatusBadRequest)
	}
	state := csrf.Token(r)
	cookie := http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	url := oauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}

func (o *OAuths) Callback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	oauthConfig, ok := o.configs[service]
	if !ok {
		http.Error(w, "invalid OAuth2 service", http.StatusBadRequest)
	}
	r.ParseForm()
	state := r.FormValue("state")
	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if cookie == nil || cookie.Value != state {
		http.Error(w, "Invalid state provided", http.StatusBadRequest)
		return
	}
	cookie.Value = ""
	cookie.Expires = time.Now()
	http.SetCookie(w, cookie)

	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := llctx.User(r.Context())
	existing, err := o.os.Find(user.ID, service)
	if err == models.ErrNotFound {
		// noop
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		o.os.Delete(existing.ID)
	}
	userOAuth := models.OAuth{
		UserID:  user.ID,
		Token:   *token,
		Service: service,
	}
	err = o.os.Create(&userOAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%+v", token)
}

func (o *OAuths) DropboxTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	oauthConfig, ok := o.configs[service]
	if !ok {
		http.Error(w, "invalid OAuth2 service", http.StatusBadRequest)
	}
	r.ParseForm()
	path := r.FormValue("path")

	user := llctx.User(r.Context())
	userOAuth, err := o.os.Find(user.ID, service)
	if err != nil {
		panic(err)
	}
	token := userOAuth.Token

	data := struct {
		Path string `json:"path"`
	}{
		Path: path,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	client := oauthConfig.Client(context.TODO(), &token)
	req, err := http.NewRequest(http.MethodPost, "https://api.dropboxapi.com/2/files/list_folder", bytes.NewReader(dataBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

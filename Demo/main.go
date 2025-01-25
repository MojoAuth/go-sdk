package main

import (
	"fmt"
	"html/template"
	"log"
	"mojogodemo/config"
	"mojogodemo/utils"
	"net/http"

	"encoding/json"

	"github.com/gorilla/sessions"
	go_mojoauth "github.com/mojoauth/go-sdk"
	"github.com/mojoauth/go-sdk/api"
	"github.com/mojoauth/go-sdk/mojoerror"
)

type TemplateData struct {
	Title    string
	Subtitle string
	BaseURL  string
	ApiKey   string
}

type Token struct {
	Token string
}

type TokenResponse struct {
	IsValid bool   `json:"isValid"`
	Token   string `json:"access_token"`
}

type UserData struct {
	Title    string
	Subtitle string
	Email    string
}

type Oauth struct {
	AccessToken  string
	ExpiresIn    string
	IdToken      string
	RefreshToken string
	TokenType    string
}
type User struct {
	CreatedAt  string
	Identifier string
	Issuer     string
	UpdatedAt  string
	UserId     string
}

type Response struct {
	Authenticated bool
	Oauth         Oauth
	User          User
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	addr := ":" + fmt.Sprint(config.App.Port)
	mux := http.NewServeMux()

	var errors string

	cfg := go_mojoauth.Config{
		ApiKey: config.App.APIKey,
	}

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("tmpl/index.tmpl"))

		// Initialze a struct storing page data and todo data
		data := TemplateData{
			Title:    "Golang MojoAuth Demo",
			Subtitle: "This is a Email magic link authentication demo",
			BaseURL:  config.App.BaseURL,
			ApiKey:   config.App.APIKey,
		}

		// Render the data and output using standard output
		t.Execute(w, data)
	}

	verifyHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		var reqObj Token
		err := json.NewDecoder(r.Body).Decode(&reqObj)
		if err != nil {
			fmt.Println("Token Missing")
			return
		}

		mojoClient, err := go_mojoauth.NewMojoAuth(&cfg)
		if err != nil {
			errors = errors + err.(mojoerror.Error).OrigErr().Error()
			//      respCode = 500
		}
		response, err := api.Mojoauth{mojoClient}.VerifyToken(reqObj.Token)
		if err != nil {
			errors = errors + err.(mojoerror.Error).OrigErr().Error()
			//		respCode = 500
		}
		if errors != "" {
			log.Printf(errors)

			return
		}
		if response.IsValid == true {
			session, _ := store.Get(r, "GoDemoCookie")
			session.Values["authenticated"] = true
			session.Values["token"] = response.Token
			session.Save(r, w)
		} else {
			session, _ := store.Get(r, "GoDemoCookie")
			session.Values["authenticated"] = false
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		// body, err := ioutil.ReadAll(res.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		utils.WriteResponseJSON(r, w, http.StatusOK, response)
	}

	MagicLinkHandler := func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		var reqObj map[string]string
		err := json.NewDecoder(r.Body).Decode(&reqObj)
		if err != nil {
			fmt.Println("Token Missing")
			return
		}
		language := r.URL.Query().Get("language")
		redirect_url := r.URL.Query().Get("redirect_url")
		mojoClient, err := go_mojoauth.NewMojoAuth(&cfg)
		if err != nil {
			errors = errors + err.(mojoerror.Error).OrigErr().Error()
			//      respCode = 500
		}
		queryParams := map[string]string{
			"language":     language,
			"redirect_url": redirect_url,
		}
		res, err := api.Mojoauth{mojoClient}.SigninWithMagicLink(reqObj, queryParams)
		if err != nil {
			errors = errors + err.(mojoerror.Error).OrigErr().Error()
			//		respCode = 500
		}
		if errors != "" {
			log.Printf(errors)

			return
		}

		// var response TokenResponse

		var jsonMap map[string]string
		json.Unmarshal([]byte(res.Body), &jsonMap)

		fmt.Println(jsonMap)
		// if jsonMap.Authenticated == true {
		// 	session, _ := store.Get(r, "GoDemoCookie")
		// 	session.Values["authenticated"] = true
		// 	session.Values["identifier"] = jsonMap.User.Identifier
		// 	session.Save(r, w)
		// 	http.Redirect(w, r, "/myAccount", http.StatusSeeOther)
		// } else {
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// }

	}

	profileHandler := func(w http.ResponseWriter, r *http.Request) {

		stateId := r.URL.Query().Get("state_id")
		fmt.Println(stateId)
		mojoClient, err := go_mojoauth.NewMojoAuth(&cfg)
		if err != nil {
			errors = errors + err.(mojoerror.Error).OrigErr().Error()
			//      respCode = 500
		}
		queryParams := map[string]string{
			"state_id": stateId,
		}
		res, err := api.Mojoauth{mojoClient}.PingStatus(queryParams)
		if err != nil {
			errors = errors + err.(mojoerror.Error).OrigErr().Error()
			//		respCode = 500
		}
		if errors != "" {
			log.Printf(errors)

			return
		}

		// var response TokenResponse

		var jsonMap Response
		json.Unmarshal([]byte(res.Body), &jsonMap)

		fmt.Println(jsonMap)
		if jsonMap.Authenticated == true {
			session, _ := store.Get(r, "GoDemoCookie")
			session.Values["authenticated"] = true
			session.Values["identifier"] = jsonMap.User.Identifier
			session.Save(r, w)
			http.Redirect(w, r, "/myAccount", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}

	myAccountHandler := func(w http.ResponseWriter, r *http.Request) {

		t := template.Must(template.ParseFiles("tmpl/myaccount.tmpl"))
		session, _ := store.Get(r, "GoDemoCookie")
		// Initialze a struct storing page data and todo data
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			identifier := session.Values["identifier"].(string)
			data := UserData{
				Title:    "User Account",
				Subtitle: "You are successfully logged in",
				Email:    identifier,
			}

			// Render the data and output using standard output
			t.Execute(w, data)
		}
	}
	logoutHandler := func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "GoDemoCookie")
		session.Values["authenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	// Make and parse the HTML template

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/magiclink", MagicLinkHandler)
	mux.HandleFunc("/verify", verifyHandler)
	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/myAccount", myAccountHandler)
	mux.HandleFunc("/logout", logoutHandler)
	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

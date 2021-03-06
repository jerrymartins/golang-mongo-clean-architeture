package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/jsonq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"strings"

	. "go-rest-mongo-clean-architeture/config"
	. "go-rest-mongo-clean-architeture/config/error"
	. "go-rest-mongo-clean-architeture/controller/config"
	. "go-rest-mongo-clean-architeture/usecase"
)

var (
	// oauth2
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
	authenticated    = false
	//oauthUseCase = Oau
	config       = Config{}
	oAuthUseCase = OauthUseCase{}
)

func init() {
	config.Read()

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  config.UrlApi + "/callback",
		ClientID:     config.GoogleClientId,
		ClientSecret: config.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func HandleLoggin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tokenJWT, err := oAuthUseCase.GenerateLocalJWTToken(params["googleToken"])

	if err != nil {
		requestError := err.(*RequestError)
		RespondWithError(w, requestError.StatusCode(), requestError.Error())
	} else {
		RespondWithText(w, http.StatusOK, tokenJWT)
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		err := oAuthUseCase.ValidateToken([]byte(token))

		if err != nil {
			http.Redirect(w, r, config.UrlApi, http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(content)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	authenticated, err = jq.Bool("verified_email")

	// TODO definir tratamento de erro caso a resposta da api do google não contenha a propriedade verified_email
	// lançar um erro e retornar para a tela de login?
	if err != nil {
		fmt.Println("tratar erro caso a propriedade não exista")
	}

	fmt.Fprintf(w, "Content: %s\n", content)
}

func HandleLogoff() {
	//	TODO: implementar logoff com a api do google
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

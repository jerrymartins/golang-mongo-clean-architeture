package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	base64 "encoding/base64"
	"encoding/json"
	"net/http"
)

func init() {

}

type OauthUseCase struct{}

type TokenJWT struct {
	Token string `json:"token"`
}

func (*OauthUseCase) ValidateGoogleToken(token string) (*http.Response, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	// seria uma boa retornar só o status code
	return response, err
}

func (*OauthUseCase) GenerateLocalJWTToken(googleToken string) (TokenJWT, error) {
	headerPayload := configHeader() + "." + configPayload(googleToken)
	signature := configSignature(headerPayload)
	token := headerPayload + "." + signature

	return TokenJWT{Token: token}, nil
}

func configHeader() string {
	header := make(map[string]string)
	header["alg"] = "HS256"
	header["type"] = "JWT"

	headerJson, _ := json.Marshal(header)
	headerBase64 := base64.StdEncoding.EncodeToString([]byte(headerJson))

	return headerBase64
}

func configPayload(googleToken string) string {
	payload := make(map[string]string)
	payload["code"] = base64.StdEncoding.EncodeToString([]byte(googleToken))
	payload["iss"] = "localhost"

	payloadJson, _ := json.Marshal(payload)
	payloadBase64 := base64.StdEncoding.EncodeToString([]byte(payloadJson))

	return payloadBase64
}

func configSignature(headerPayload string) string {
	h := hmac.New(sha256.New, []byte(config.SecretKey))

	h.Write([]byte(headerPayload))

	sha := base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))

	return sha
}

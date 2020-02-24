package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gbrlsnchs/jwt/v3"
	requestError "go-rest-mongo-clean-architeture/config/error"
	"io"
	"log"
	"net/http"
	"time"
)

type OauthUseCase struct{}

type TokenInfo struct {
	Payload Payload
}

type Payload struct {
	jwt.Payload
	GoogleToken string `json:"googleToken,omitempty"`
}

var hs *jwt.HMACSHA

func init() {
	config.Read()
	hs = jwt.NewHS256([]byte(config.SecretKey))
}

func validateGoogleToken(token string) (*http.Response, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)

	return response, err
}

func (*OauthUseCase) GenerateLocalJWTToken(googleToken string) (string, error) {
	tokenValidateResponse, errTokenValidate := validateGoogleToken(googleToken)

	if errTokenValidate != nil {
		return "", errTokenValidate
	}

	if tokenValidateResponse.StatusCode != 200 {

		errorTokenValidate := requestError.ConfigureRequestError(tokenValidateResponse.Status, tokenValidateResponse.StatusCode)

		return "", &errorTokenValidate
	}

	//if encrypted, err := encrypt([]byte(config.SecretKey), googleToken); err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println(encrypted)
	//	if decrypted, err := decrypt([]byte(config.SecretKey), encrypted); err != nil {
	//		log.Println(err)
	//	} else {
	//		log.Printf("DECRYPTED: %s\n", decrypted)
	//	}
	//}

	if googleTokenEncrypted, err := encrypt([]byte(config.SecretKey), googleToken); err != nil {
		log.Println(err)
		return googleTokenEncrypted, errors.New("not crypted")
	} else {
		token, errTokenCreation := createToken(googleTokenEncrypted)

		return token, errTokenCreation
	}

}

func (*OauthUseCase) ValidateToken(token []byte) error {
	var pl Payload
	_, err := jwt.Verify([]byte(token), hs, &pl)
	if err != nil {
		return err
	}
	return nil
}

func createToken(googleToken string) (string, error) {
	now := time.Now()
	pl := Payload{
		Payload: jwt.Payload{
			Issuer:         "borala",
			Subject:        "id do usuario todo",
			Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "152576",
		},
		GoogleToken: googleToken,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		return "", err
	}

	return string(token), err
}

func encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func decrypt(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}

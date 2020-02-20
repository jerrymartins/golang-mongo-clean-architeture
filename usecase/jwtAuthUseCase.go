package usecase

import (
	"fmt"
	"github.com/gbrlsnchs/jwt"
	"time"
)

type CustomPayload struct {
	jwt.Payload
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"bar,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

func (*CustomPayload) GenerateToken() []byte {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer:         "gbrlsnchs",
			Subject:        "someone",
			Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		Foo: "foo",
		Bar: 1337,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		// ...
	}

	return token

}

func (*CustomPayload) VerifyToken(token []byte) {
	var hs2 = jwt.NewHS256([]byte("secret"))
	var pl CustomPayload
	hd, err := jwt.Verify(token, hs2, &pl)
	if err != nil {
		// ...
	}

	fmt.Println(hd)
}

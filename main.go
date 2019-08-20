package main

import (
	"fmt"
	"github.com/gbrlsnchs/jwt/v3"
	"time"
)

type CustomPayload struct {
	jwt.Payload
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"bar,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

func main() {
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
		fmt.Println(err)
	}

	fmt.Println(string(token))
	fmt.Println("-----------")

	hd, err := jwt.Verify(token, hs, &pl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hd.Algorithm, hd.ContentType, hd.KeyID, hd.Type)
}

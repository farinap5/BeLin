package main

import (
	"belin/keys"
	"encoding/base64"
	"fmt"
)

func main() {
	priv, publ, err := keys.Extract("cobaltstrike.beacon_keys")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	base64.StdEncoding.EncodeToString(publ)

	prPem := fmt.Sprintf(`
-----BEGIN RSA PRIVATE KEY-----
%s
-----END RSA PRIVATE KEY-----`,base64.StdEncoding.EncodeToString(priv))

	pbPem := fmt.Sprintf(`
-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----`,base64.StdEncoding.EncodeToString(publ))

	fmt.Println(prPem,pbPem)
}
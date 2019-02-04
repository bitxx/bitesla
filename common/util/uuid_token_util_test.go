package util

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	key, e := GetKey()

	if e != nil {
		t.Error(e)
	}
	t.Log("key:", key)

	secret := GetSecret(key)
	t.Log("secret:", secret)

	token, err := GetToken(key, secret, "aichain", "aichain123456", 99999)
	if err != nil {
		t.Error(e)
	}
	claims, err := ParseToken(token, "aichain123456")
	if err != nil {
		t.Error(err)
	}
	t.Log("result:", claims)

}

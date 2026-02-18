package tiktokads

import (
	"log"
	"testing"
)

func TestGetAdAccounts(t *testing.T) {
	initTestSession()

	accounts, err := GetAdAccounts()

	assertNotEmptyResult(t, accounts, err)
	log.Println(LastResponse)

	for _, acc := range accounts {
		if len(acc.Id) == 0 {
			t.Fatal("empty account id")
		}
		t.Log(acc)
	}
}

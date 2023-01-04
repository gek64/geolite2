package geolite2

import (
	"fmt"
	"testing"
)

func TestAsnDB_Search(t *testing.T) {
	asnDB, err := OpenAsnDB("example_database/ASN-IPv4.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func(asnDB AsnDB) {
		err := asnDB.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(asnDB)

	asnResult, err := asnDB.Search("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(asnResult)
}

package geolite2

import (
	"fmt"
	"testing"
)

func TestCountryBlockDB_Search(t *testing.T) {
	countryBlockDB, err := OpenCountryBlockDB("example_database/Country-Blocks-IPv4.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func(countryBlockDB CountryBlockDB) {
		err := countryBlockDB.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(countryBlockDB)

	countryBlockResult, err := countryBlockDB.Search("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(countryBlockResult)
}

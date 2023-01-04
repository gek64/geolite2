package geolite2

import (
	"fmt"
	"testing"
)

func TestCityBlockDB_Search(t *testing.T) {
	cityBlockDB, err := OpenCityBlockDB("example_database/City-Blocks-IPv4.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func(cityBlockDB CityBlockDB) {
		err := cityBlockDB.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(cityBlockDB)

	cityBlockResult, err := cityBlockDB.Search("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cityBlockResult)
}

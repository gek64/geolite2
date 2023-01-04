package geolite2

import (
	"fmt"
	"testing"
)

func TestCityLocationDB_Search(t *testing.T) {
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

	cityLocationDB, err := OpenCityLocationDB("example_database/City-Locations-en.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func(cityLocationDB CityLocationDB) {
		err := cityLocationDB.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(cityLocationDB)

	cityLocationResult, err := cityLocationDB.Search(cityBlockResult.GeoNameID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cityLocationResult)
}

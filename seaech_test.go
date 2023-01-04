package geolite2

import (
	"fmt"
	"testing"
)

var (
	ipv4 = "1.1.1.1"
	ipv6 = "2606:4700:4700::64"
)

func TestAsnSearch_IPv4(t *testing.T) {
	asnResult, err := AsnSearch(ipv4, "example_database/ASN-IPv4.csv")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(asnResult)
}

func TestCitySearch_IPv4(t *testing.T) {
	cityResult, err := CitySearch(ipv4, "example_database/City-Blocks-IPv4.csv", "example_database/City-Locations-en.csv")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cityResult)
}

func TestCountrySearch_IPv4(t *testing.T) {
	countryResult, err := CountrySearch(ipv4, "example_database/Country-Blocks-IPv4.csv", "example_database/Country-Locations-en.csv")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(countryResult)
}

func TestAsnSearch_IPv6(t *testing.T) {
	asnResult, err := AsnSearch(ipv6, "example_database/ASN-IPv6.csv")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(asnResult)
}

func TestCitySearch_IPv6(t *testing.T) {
	cityResult, err := CitySearch(ipv6, "example_database/City-Blocks-IPv6.csv", "example_database/City-Locations-en.csv")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cityResult)
}

func TestCountrySearch_IPv6(t *testing.T) {
	countryResult, err := CountrySearch(ipv6, "example_database/Country-Blocks-IPv6.csv", "example_database/Country-Locations-en.csv")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(countryResult)
}

package geolite2

import (
	"log"
)

func asnResultSearch(ip string, database string) (asn AsnResult, err error) {
	asnDB, err := OpenAsnDB(database)
	if err != nil {
		return AsnResult{}, err
	}
	defer func(asnDB AsnDB) {
		err := asnDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(asnDB)
	return asnDB.Search(ip)
}

func cityBlockSearch(ip string, database string) (cbResult CityBlockResult, err error) {
	cbDB, err := OpenCityBlockDB(database)
	if err != nil {
		return CityBlockResult{}, err
	}
	defer func(cbDB CityBlockDB) {
		err := cbDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(cbDB)
	return cbDB.Search(ip)
}

func cityLocationSearch(geoNameID int, database string) (clResult CityLocationResult, err error) {
	clDB, err := OpenCityLocationDB(database)
	if err != nil {
		return CityLocationResult{}, err
	}
	defer func(clDB CityLocationDB) {
		err := clDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(clDB)
	return clDB.Search(geoNameID)
}

func countryBlockSearch(ip string, database string) (cbResult CountryBlockResult, err error) {
	cbDB, err := OpenCountryBlockDB(database)
	if err != nil {
		return CountryBlockResult{}, err
	}
	defer func(cbDB CountryBlockDB) {
		err := cbDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(cbDB)
	return cbDB.Search(ip)
}

func countryLocationSearch(geoNameID int, database string) (clResult CountryLocationResult, err error) {
	clDB, err := OpenCountryLocationDB(database)
	if err != nil {
		return CountryLocationResult{}, err
	}
	defer func(clDB CountryLocationDB) {
		err := clDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(clDB)
	return clDB.Search(geoNameID)
}

// 外部可访问函数

func AsnSearch(ip string, database string) (asn AsnResult, err error) {
	return asnResultSearch(ip, database)
}

func CitySearch(ip string, cityBlockDB string, cityLocationDB string) (clResult CityLocationResult, err error) {
	blockResult, err := cityBlockSearch(ip, cityBlockDB)
	if err != nil {
		return CityLocationResult{}, err
	}

	return cityLocationSearch(blockResult.GeoNameID, cityLocationDB)
}

func CountrySearch(ip string, countryBlockDB string, countryLocationDB string) (clResult CountryLocationResult, err error) {
	blockResult, err := countryBlockSearch(ip, countryBlockDB)
	if err != nil {
		return CountryLocationResult{}, err
	}

	return countryLocationSearch(blockResult.GeoNameID, countryLocationDB)
}

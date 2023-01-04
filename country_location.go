package geolite2

import (
	"encoding/csv"
	"os"
	"strconv"
)

type CountryLocationDB struct {
	fs     *os.File
	reader *csv.Reader
}

type CountryLocationResult struct {
	GeoNameID         int
	LocaleCode        string
	ContinentCode     string
	ContinentName     string
	CountryIsoCode    string
	CountryName       string
	IsInEuropeanUnion bool
}

// OpenCountryLocationDB 打开一个Country Location数据库文件
func OpenCountryLocationDB(database string) (countryLocationDB CountryLocationDB, err error) {
	fs, err := os.OpenFile(database, os.O_RDONLY, 0644)
	if err != nil {
		return CountryLocationDB{}, err
	}

	return CountryLocationDB{fs: fs, reader: csv.NewReader(fs)}, nil
}

// Close 关闭
func (clDB CountryLocationDB) Close() (err error) {
	return clDB.fs.Close()
}

// Search 搜索
func (clDB CountryLocationDB) Search(geoNameID int) (clResult CountryLocationResult, err error) {
	// 循环遍历整个数据库,找到对应的项目后返回,或读到EOF后退出
	for {
		// 读一行cbDB记录
		record, err := clDB.reader.Read()
		if err != nil {
			return CountryLocationResult{}, err
		}

		if record[0] == strconv.Itoa(geoNameID) {
			clResult.GeoNameID, err = strconv.Atoi(record[0])
			if err != nil {
				clResult.GeoNameID = -1
			}
			clResult.LocaleCode = record[1]
			clResult.ContinentCode = record[2]
			clResult.ContinentName = record[3]
			clResult.CountryIsoCode = record[4]
			clResult.CountryName = record[5]
			clResult.IsInEuropeanUnion = record[6] == "1"
			break
		}
	}
	return clResult, err
}

package geolite2

import (
	"encoding/csv"
	"os"
	"strconv"
)

type CityLocationDB struct {
	fs     *os.File
	reader *csv.Reader
}

type CityLocationResult struct {
	GeoNameID           int
	LocaleCode          string
	ContinentCode       string
	ContinentName       string
	CountryIsoCode      string
	CountryName         string
	Subdivision1IsoCode string
	Subdivision1Name    string
	Subdivision2IsoCode string
	Subdivision2Name    string
	CityName            string
	MetroCode           string
	TimeZone            string
	IsInEuropeanUnion   bool
}

// OpenCityLocationDB 打开一个City Location数据库文件
func OpenCityLocationDB(database string) (cityLocationDB CityLocationDB, err error) {
	fs, err := os.OpenFile(database, os.O_RDONLY, 0644)
	if err != nil {
		return CityLocationDB{}, err
	}

	return CityLocationDB{fs: fs, reader: csv.NewReader(fs)}, nil
}

// Close 关闭
func (clDB CityLocationDB) Close() (err error) {
	return clDB.fs.Close()
}

// Search 搜索
func (clDB CityLocationDB) Search(geoNameID int) (clResult CityLocationResult, err error) {
	// 循环遍历整个数据库,找到对应的项目后返回,或读到EOF后退出
	for {
		// 读一行cbDB记录
		record, err := clDB.reader.Read()
		if err != nil {
			return CityLocationResult{}, err
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
			clResult.Subdivision1IsoCode = record[6]
			clResult.Subdivision1Name = record[7]
			clResult.Subdivision2IsoCode = record[8]
			clResult.Subdivision2Name = record[9]
			clResult.CityName = record[10]
			clResult.MetroCode = record[11]
			clResult.TimeZone = record[12]
			clResult.IsInEuropeanUnion = record[13] == "1"
			break
		}
	}
	return clResult, err
}

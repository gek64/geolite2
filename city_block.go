package geolite2

import (
	"encoding/csv"
	"net/netip"
	"os"
	"strconv"
)

type CityBlockDB struct {
	fs     *os.File
	reader *csv.Reader
}

type CityBlockResult struct {
	Network                     netip.Addr
	GeoNameID                   int
	RegisteredCountryGeoNameId  int
	RepresentedCountryGeoNameId int
	IsAnonymousProxy            bool
	IsSatelliteProvider         bool
	PostalCode                  string
	Latitude                    float64
	Longitude                   float64
	AccuracyRadius              float64
}

// OpenCityBlockDB 打开一个City Block数据库文件
func OpenCityBlockDB(database string) (cityBlockDB CityBlockDB, err error) {
	fs, err := os.OpenFile(database, os.O_RDONLY, 0644)
	if err != nil {
		return CityBlockDB{}, err
	}

	return CityBlockDB{fs: fs, reader: csv.NewReader(fs)}, nil
}

// Close 关闭
func (cbDB CityBlockDB) Close() (err error) {
	return cbDB.fs.Close()
}

// Search 搜索
func (cbDB CityBlockDB) Search(ip string) (cbResult CityBlockResult, err error) {
	// 循环遍历整个数据库,找到对应的项目后返回,或读到EOF后退出
	for {
		// 解析IP
		addr, err := netip.ParseAddr(ip)
		if err != nil {
			return CityBlockResult{}, err
		}
		// 读一行cbDB记录
		record, err := cbDB.reader.Read()
		if err != nil {
			return CityBlockResult{}, err
		}
		// 解析Network块
		cidr, err := netip.ParsePrefix(record[0])
		if err != nil {
			continue
		}
		// 判断解析后的IP是否属于当前数据库Network块
		if cidr.Contains(addr) {
			cbResult.Network = addr
			cbResult.GeoNameID, err = strconv.Atoi(record[1])
			if err != nil {
				cbResult.RegisteredCountryGeoNameId = -1
			}
			cbResult.RegisteredCountryGeoNameId, err = strconv.Atoi(record[2])
			if err != nil {
				cbResult.RegisteredCountryGeoNameId = -1
			}
			cbResult.RepresentedCountryGeoNameId, err = strconv.Atoi(record[3])
			if err != nil {
				cbResult.RepresentedCountryGeoNameId = -1
			}
			cbResult.IsAnonymousProxy = record[4] == "1"
			cbResult.IsSatelliteProvider = record[5] == "1"
			cbResult.PostalCode = record[6]
			cbResult.Latitude, err = strconv.ParseFloat(record[7], 64)
			if err != nil {
				cbResult.Latitude = -1
			}
			cbResult.Longitude, err = strconv.ParseFloat(record[8], 64)
			if err != nil {
				cbResult.Longitude = -1
			}
			cbResult.AccuracyRadius, err = strconv.ParseFloat(record[9], 64)
			if err != nil {
				cbResult.AccuracyRadius = -1
			}
			break
		}
	}

	return cbResult, err
}

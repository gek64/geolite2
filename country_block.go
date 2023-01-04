package geolite2

import (
	"encoding/csv"
	"net/netip"
	"os"
	"strconv"
)

type CountryBlockDB struct {
	fs     *os.File
	reader *csv.Reader
}

type CountryBlockResult struct {
	Network                     netip.Addr
	GeoNameID                   int
	RegisteredCountryGeoNameId  int
	RepresentedCountryGeoNameId int
	IsAnonymousProxy            bool
	IsSatelliteProvider         bool
}

// OpenCountryBlockDB 打开一个Country Block数据库文件
func OpenCountryBlockDB(database string) (countryBlockDB CountryBlockDB, err error) {
	fs, err := os.OpenFile(database, os.O_RDONLY, 0644)
	if err != nil {
		return CountryBlockDB{}, err
	}

	return CountryBlockDB{fs: fs, reader: csv.NewReader(fs)}, nil
}

// Close 关闭
func (cbDB CountryBlockDB) Close() (err error) {
	return cbDB.fs.Close()
}

// Search 搜索
func (cbDB CountryBlockDB) Search(ip string) (cbResult CountryBlockResult, err error) {
	// 循环遍历整个数据库,找到对应的项目后返回,或读到EOF后退出
	for {
		// 解析IP
		addr, err := netip.ParseAddr(ip)
		if err != nil {
			return CountryBlockResult{}, err
		}
		// 读一行cbDB记录
		record, err := cbDB.reader.Read()
		if err != nil {
			return CountryBlockResult{}, err
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
			break
		}
	}

	return cbResult, err
}

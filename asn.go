package geolite2

import (
	"encoding/csv"
	"net/netip"
	"os"
	"strconv"
)

type AsnDB struct {
	fs     *os.File
	reader *csv.Reader
}

type AsnResult struct {
	Network netip.Addr
	Asn     int
	Name    string
}

// OpenAsnDB 打开一个ASN数据库文件
func OpenAsnDB(database string) (asn AsnDB, err error) {
	fs, err := os.OpenFile(database, os.O_RDONLY, 0644)
	if err != nil {
		return AsnDB{}, err
	}

	return AsnDB{fs: fs, reader: csv.NewReader(fs)}, nil
}

// Close 关闭
func (asnDB AsnDB) Close() (err error) {
	return asnDB.fs.Close()
}

// Search 搜索
func (asnDB AsnDB) Search(ip string) (asnResult AsnResult, err error) {
	// 循环遍历整个ASN数据库,找到对应的项目后返回,或读到EOF后退出
	for {
		// 解析IP
		addr, err := netip.ParseAddr(ip)
		if err != nil {
			return AsnResult{}, err
		}
		// 读一行AsnDB记录
		record, err := asnDB.reader.Read()
		if err != nil {
			return AsnResult{}, err
		}
		// 解析Network块
		cidr, err := netip.ParsePrefix(record[0])
		if err != nil {
			continue
		}
		// 判断解析后的IP是否属于当前数据库Network块
		if cidr.Contains(addr) {
			asnResult.Network = addr
			asnResult.Asn, err = strconv.Atoi(record[1])
			if err != nil {
				return AsnResult{}, err
			}
			asnResult.Name = record[2]
			break
		}
	}

	return asnResult, err
}

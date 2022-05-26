// Code generated by go-bindata.
// sources:
// 1571295324_create_properties_table.down.sql
// 1571295324_create_properties_table.up.sql
// 1571295325_create_models_table.down.sql
// 1571295325_create_models_table.up.sql
// 1571295326_create_abstractions_table.down.sql
// 1571295326_create_abstractions_table.up.sql
// 1571295327_create_activations_table.down.sql
// 1571295327_create_activations_table.up.sql
// 1571295328_create_rules_table.down.sql
// 1571295328_create_rules_table.up.sql
// 1572590132_create_collections_table.down.sql
// 1572590132_create_collections_table.up.sql
// 1572590669_create_collection_items_table.down.sql
// 1572590669_create_collection_items_table.up.sql
// DO NOT EDIT!

package metadata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __1571295324_create_properties_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x28\x28\xca\x2f\x48\x2d\x2a\xc9\x4c\x2d\x4e\xb0\x06\x04\x00\x00\xff\xff\x9d\xed\xb7\x55\x18\x00\x00\x00")

func _1571295324_create_properties_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295324_create_properties_tableDownSql,
		"1571295324_create_properties_table.down.sql",
	)
}

func _1571295324_create_properties_tableDownSql() (*asset, error) {
	bytes, err := _1571295324_create_properties_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295324_create_properties_table.down.sql", size: 24, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295324_create_properties_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xbd\x8e\xd3\x40\x14\x85\xfb\x3c\xc5\xed\x6c\x4b\x34\x41\x44\x42\x42\x29\x26\xf6\x25\x8c\xb0\xc7\x61\x32\x23\x91\xca\x33\xb1\x07\xb0\x14\x3b\x96\x33\x89\x44\x09\x2d\x20\xaa\xb4\x3c\x00\x05\x84\x9a\xe7\x49\x76\xf7\x2d\x56\x76\x76\xb3\x7f\x91\x76\xd3\x9f\xef\xdc\xcf\xc7\xe3\x73\x24\x02\x41\x90\x41\x88\xa0\xaa\x7a\x5e\x99\xda\xe6\x66\xa1\xc0\xed\x00\xa8\x3c\x53\x90\x97\xd6\xed\x76\x3d\x60\xb1\x00\x26\xc3\x10\x88\x14\x71\x42\x99\xcf\x31\x42\x26\x9e\x35\xb9\x52\x17\x46\xc1\x4a\xd7\xe9\x27\x5d\xbb\xcf\x7b\x3d\x0f\xfc\x38\x0c\x9b\xea\xa5\xfd\xf0\xb2\x98\xbe\x48\x96\x65\x9e\xce\x33\x93\xa4\xf9\xa1\xa9\x45\x67\x7a\x6a\x66\x4f\x67\x03\x7c\x4d\x64\x78\x8b\xb7\x9f\x2b\x73\x44\xd2\x8f\xa3\xc6\x0e\x9c\xdd\x7a\xb3\xfb\xf1\xe7\xec\xdf\xff\xed\xaf\x6f\x4e\x4b\xac\xf4\x2c\xcf\xb4\x35\xc9\x63\xe8\xc5\xef\xef\xe7\x7f\xbf\x1c\x45\x75\xfd\x71\x71\xfa\x07\x1f\xec\x1d\xe7\xfe\x95\xed\xcf\xaf\xbb\xf5\x66\x7f\x25\xad\x8d\xb6\x26\x4b\xb4\x55\x60\xf3\xc2\x2c\xac\x2e\xaa\xbb\x05\xbe\xe4\x1c\x99\x48\x04\x8d\x70\x2c\x48\x34\x6a\xc1\x65\x95\x9d\x0c\x42\xcc\x40\x8e\x82\xc6\xfc\x68\xa9\x64\xf4\x9d\x44\x78\x8b\x13\x50\x8c\x44\x98\xd0\xe0\xbd\x02\x77\xff\xcb\x3d\x90\x63\xca\x86\x30\x10\x1c\xb1\x49\x8f\x38\x8d\x08\x9f\xb4\x71\xb7\x79\x3e\x5e\xc7\x03\x64\x43\xca\xb0\x4f\xcb\x72\x1e\x0c\x6e\x4c\xde\x10\x3e\x46\xd1\xbf\xda\xeb\x7a\xbf\xfe\xc3\xfd\x5e\x5d\x06\x00\x00\xff\xff\x05\x99\xa7\xcd\xa6\x02\x00\x00")

func _1571295324_create_properties_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295324_create_properties_tableUpSql,
		"1571295324_create_properties_table.up.sql",
	)
}

func _1571295324_create_properties_tableUpSql() (*asset, error) {
	bytes, err := _1571295324_create_properties_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295324_create_properties_table.up.sql", size: 678, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295325_create_models_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xc8\xcd\x4f\x49\xcd\x29\x4e\xb0\x06\x04\x00\x00\xff\xff\x55\x0c\x38\xcc\x14\x00\x00\x00")

func _1571295325_create_models_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295325_create_models_tableDownSql,
		"1571295325_create_models_table.down.sql",
	)
}

func _1571295325_create_models_tableDownSql() (*asset, error) {
	bytes, err := _1571295325_create_models_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295325_create_models_table.down.sql", size: 20, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295325_create_models_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\xb1\x6e\xf2\x30\x14\x85\x77\x9e\xe2\x6e\x24\xd2\x3f\xfc\x54\x20\x55\xaa\x18\x4c\x72\x4b\xa3\x3a\x0e\x32\xf6\xc0\x14\x9b\xd8\xa5\x96\x88\x41\xc4\xee\xf3\x57\x86\xb6\xa8\x6a\x17\xba\xdf\xef\xd3\x39\xe7\x16\x1c\x89\x40\x10\x64\x41\x11\x54\x7f\x30\x76\x3f\x28\xc8\x46\x00\xca\x19\x05\xce\x87\x6c\x32\xc9\x81\x35\x02\x98\xa4\x14\x88\x14\x4d\x5b\xb1\x82\x63\x8d\x4c\xfc\x4b\x77\x5e\xf7\x56\xc1\x9b\x3e\x75\xaf\xfa\x94\xdd\xcd\x66\x39\x14\x0d\xa5\x49\x1b\xc3\xcb\x7d\xbf\x9d\xb6\xd1\xbb\xee\x60\x6c\xdb\xb9\x2f\xd3\x19\xdd\xeb\xad\xdd\xff\x91\xdd\xc5\x14\xf0\x56\x14\x8a\xa6\x4e\xc9\x61\xbc\x94\x55\x39\x3e\x9b\x86\xa0\x43\x1c\x2e\x65\xa7\x39\x44\x3f\xb8\x9d\xb7\xe6\xca\x94\xf8\x48\x24\x15\x30\xfe\x7f\x01\xba\x93\xd5\xc1\x9a\x56\x07\x05\xc1\xf5\x76\x08\xba\x3f\x7e\x3f\x2d\x24\xe7\xc8\x44\x2b\xaa\x1a\xd7\x82\xd4\xab\x33\x18\x8f\xe6\x66\x10\x1a\x06\x72\x55\xa6\x5e\xbf\x4a\x57\xbc\xaa\x09\xdf\xc0\x33\x6e\x20\x4b\x4f\xcb\x47\x39\x20\x5b\x56\x0c\xe7\x95\xf7\x87\x72\x71\x75\x3f\x11\xbe\x46\x31\xff\xd8\xe7\x73\xaf\xf9\xcf\xbd\x1e\xde\x03\x00\x00\xff\xff\xac\x4d\x14\x89\x18\x02\x00\x00")

func _1571295325_create_models_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295325_create_models_tableUpSql,
		"1571295325_create_models_table.up.sql",
	)
}

func _1571295325_create_models_tableUpSql() (*asset, error) {
	bytes, err := _1571295325_create_models_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295325_create_models_table.up.sql", size: 536, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295326_create_abstractions_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x48\x4c\x2a\x2e\x29\x4a\x4c\x2e\xc9\xcc\xcf\x2b\x4e\xb0\x06\x04\x00\x00\xff\xff\xcc\x37\xc6\x15\x1a\x00\x00\x00")

func _1571295326_create_abstractions_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295326_create_abstractions_tableDownSql,
		"1571295326_create_abstractions_table.down.sql",
	)
}

func _1571295326_create_abstractions_tableDownSql() (*asset, error) {
	bytes, err := _1571295326_create_abstractions_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295326_create_abstractions_table.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295326_create_abstractions_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xcf\x8b\xd3\x40\x1c\xc5\xef\xfb\x57\x7c\x6f\x4d\xc0\x83\x5d\x76\x41\x90\x1e\xd2\x74\x5c\x83\xc9\xa4\xa6\x13\x70\x4f\x93\x69\x32\xad\x03\xf9\x51\x92\x69\xd1\xa3\xa7\x55\x3c\xb8\x17\x29\x88\x0a\x2b\x22\x5e\x14\x84\x75\x0f\x56\x76\xff\x19\x93\x76\xff\x0b\x49\xab\xdd\x6e\x0c\x42\xf7\x3c\xef\x7d\x78\xdf\xc7\x1b\xdd\x41\x1a\x41\x40\xb4\xb6\x89\xc0\x63\xfd\x4c\xa6\xcc\x97\x22\x89\x33\x0f\x94\x1d\x00\x4f\x04\x1e\x88\x58\x2a\xcd\xa6\x0a\xd8\x26\x80\x5d\xd3\x04\xcd\x25\x36\x35\xb0\xee\x20\x0b\x61\x72\xab\xd4\xc5\x2c\xe2\x1e\x4c\x58\xea\x3f\x66\xa9\xb2\xbb\xbf\xaf\x82\x6e\x9b\x66\x09\x1f\xcb\xc1\x9d\xa8\xbf\x47\xc7\xb1\xf0\x93\x80\x53\x5f\xac\x49\x4b\x6b\xc8\xfa\x3c\xbc\xa1\x37\x4a\x02\x1e\xd2\xba\x90\xcb\x67\x36\x1c\xa6\x7c\xc8\x24\xa7\xf2\xe9\x88\xaf\x44\x7b\x1b\x87\xe8\xb6\x55\x5e\x00\x8d\xc5\xb3\x37\xf9\xf1\xf3\xf9\xb7\x59\xfe\xfe\x65\xa3\x62\x1d\x08\x1e\x06\x95\x7c\x6b\x42\x07\xdd\xd3\x5c\x93\x40\xa3\x51\x85\xe5\x5f\xa6\xc5\xd7\xef\x55\x98\x88\x25\x4f\x27\x2c\xdc\x08\x74\xad\xda\x0a\xa4\x98\x9e\x5d\x4e\x4f\xe7\x2f\x8e\xea\xa3\xad\x69\x13\x16\x8e\xb7\xc1\x5d\xbe\xbe\xc8\x7f\x7c\x5a\xe1\x06\x22\x94\x3c\xa5\xfc\xc9\x28\xe5\x59\x26\x92\xf8\xea\xd6\xe6\xed\xdd\xfa\xba\x2e\x8e\x8a\xd9\xc7\xe2\xdd\xc9\xaf\xd9\xd9\xe2\xe4\xf3\xe2\xfc\x3c\xff\xf9\x6a\x45\xcb\x78\xe9\xfd\x7f\x69\x6b\x4e\x71\xfc\x76\x7e\xfa\x61\xb3\x29\x3f\xe5\x4c\xf2\x80\x32\xe9\x81\x14\x11\xcf\x24\x8b\x46\xd7\xab\xd6\x5d\xc7\x41\x98\x50\x62\x58\xa8\x47\x34\xab\xbb\x34\x8e\x47\xc1\xd6\x46\xb0\x31\xb8\xdd\x4e\x39\xb4\x5a\xa8\x8b\x8d\x87\x2e\x82\x07\xe8\x10\x3c\xac\x59\x88\x1a\x9d\x47\x1e\x28\xab\xb1\xab\xe0\xf6\x0c\x7c\x00\x6d\xe2\x20\x54\xaa\xbb\x8e\x61\x69\xce\xe1\x52\xae\x94\x1f\x47\xdd\x51\x01\xe1\x03\x03\xa3\x96\x11\xc7\x49\xa7\x7d\x95\xe4\xbe\xe6\xf4\x10\x69\xfd\x99\xf7\xdf\xb9\xb7\xfe\x9d\xfb\xdd\xdf\x01\x00\x00\xff\xff\xfc\xef\x8e\xef\xa2\x03\x00\x00")

func _1571295326_create_abstractions_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295326_create_abstractions_tableUpSql,
		"1571295326_create_abstractions_table.up.sql",
	)
}

func _1571295326_create_abstractions_tableUpSql() (*asset, error) {
	bytes, err := _1571295326_create_abstractions_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295326_create_abstractions_table.up.sql", size: 930, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295327_create_activations_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x48\x4c\x2e\xc9\x2c\x4b\x2c\xc9\xcc\xcf\x2b\x4e\xb0\x06\x04\x00\x00\xff\xff\xf5\x2f\x37\x8d\x19\x00\x00\x00")

func _1571295327_create_activations_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295327_create_activations_tableDownSql,
		"1571295327_create_activations_table.down.sql",
	)
}

func _1571295327_create_activations_tableDownSql() (*asset, error) {
	bytes, err := _1571295327_create_activations_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295327_create_activations_table.down.sql", size: 25, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295327_create_activations_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xb1\x8e\xd4\x30\x10\x86\xfb\x7b\x8a\xe9\x92\x48\x14\x1c\xe2\x24\x24\xb4\x85\x37\x19\x0e\x8b\xc4\x59\xbc\xb6\xc4\x55\xb6\x37\x31\x87\x45\xe2\x9c\x12\x67\x69\x29\x91\xe8\xb6\xa7\xa5\xe3\x09\xe0\x75\x58\x78\x0c\x94\x00\xbb\x42\x6c\x73\x5b\xfb\xff\x3e\xcd\xf8\x9f\x94\x23\x11\x08\x82\x2c\x73\x04\x6d\xaa\xe0\xb6\x26\xb8\xce\x0f\x1a\xe2\x0b\x00\xed\x6a\x0d\xce\x87\xf8\xf2\x32\x01\x56\x0a\x60\x32\xcf\x81\x48\x51\x2a\xca\x52\x8e\x05\x32\xf1\x60\xca\x79\xd3\x5a\x0d\x5b\xd3\x57\x6f\x4c\x1f\x3f\xba\xba\x4a\x20\x2d\xf3\x7c\x72\x8f\xe1\xf5\x93\x76\xf3\x58\x8d\xde\x55\x5d\x6d\x55\xe5\x0e\xa6\x19\x6d\xcc\xc6\x36\x67\xb2\x6d\x57\xdb\x46\x9d\x1a\x72\x7e\x7e\x67\x7a\xef\xfc\xad\x1a\xaa\xae\xb7\xc7\xcc\xe8\x07\x77\xeb\x6d\x7d\xdc\x28\xc3\x67\x44\xe6\x02\xa2\x87\x11\xa4\x65\x31\xad\x05\xd1\xcf\x2f\x9f\xf7\x1f\x76\xdf\xdf\x7f\x8b\x66\xd9\xa6\xe9\xaa\xb7\x67\xaa\xf6\x1f\x77\x3f\xbe\x7e\x3a\xa8\xaa\xde\x9a\x60\x6b\x65\x82\x86\xe0\x5a\x3b\x04\xd3\xde\xfd\x8b\xa7\x92\x73\x64\x42\x09\x5a\xe0\x5a\x90\x62\x35\x83\xe3\x5d\x7d\x6f\x10\x4a\x06\x72\x95\x4d\xdf\x79\x52\x2a\x19\x7d\x29\x11\x5e\xe0\x0d\x68\x46\x0a\x54\x34\x7b\xa5\x21\xfe\x5d\x69\x02\x72\x4d\xd9\x35\x2c\x05\x47\x9c\xd2\x2b\x4e\x0b\xc2\x6f\xe6\x78\x3c\x9d\x47\x72\x91\x00\xb2\x6b\xca\x70\x41\xbd\xef\xb2\xe5\x71\x92\xe7\x84\xaf\x51\x2c\xfe\x94\xf8\xb7\xd4\xc5\xff\xa5\x3e\xfd\x15\x00\x00\xff\xff\xaa\x14\xd0\x24\x87\x02\x00\x00")

func _1571295327_create_activations_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295327_create_activations_tableUpSql,
		"1571295327_create_activations_table.up.sql",
	)
}

func _1571295327_create_activations_tableUpSql() (*asset, error) {
	bytes, err := _1571295327_create_activations_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295327_create_activations_table.up.sql", size: 647, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295328_create_rules_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x28\x2a\xcd\x49\x2d\x4e\xb0\x06\x04\x00\x00\xff\xff\x11\x1b\x96\xb2\x13\x00\x00\x00")

func _1571295328_create_rules_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295328_create_rules_tableDownSql,
		"1571295328_create_rules_table.down.sql",
	)
}

func _1571295328_create_rules_tableDownSql() (*asset, error) {
	bytes, err := _1571295328_create_rules_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295328_create_rules_table.down.sql", size: 19, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1571295328_create_rules_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xcf\x8b\xd3\x40\x1c\xc5\xef\xfb\x57\x7c\x6f\x6d\x45\x16\x57\x5c\x10\x64\x0f\xd9\x74\xd4\x62\x9a\x94\xec\xe4\xb0\xa7\x64\x9a\x8e\x3a\xd0\x26\x65\x32\x59\x3c\x7a\xd1\x76\xb1\xae\x2b\x2a\x39\x88\x68\x2f\xfe\x00\xb7\x7a\x51\x4a\xcd\x76\xff\x99\xe6\x87\xa7\xfe\x0b\x92\x46\x9b\x8a\x15\x5c\xef\xef\x7d\xde\xbc\xf7\x65\x64\x1d\x49\x18\x01\x96\x76\x15\x04\x16\xf7\xdb\xd4\xb3\xa0\xbc\x01\x60\xb1\x96\x05\xcc\x11\xe5\xad\xad\x0a\xa8\x1a\x06\xd5\x50\x14\x90\x0c\xac\x99\x35\x55\xd6\x51\x1d\xa9\xf8\x62\xa6\x6b\x93\x26\x6d\x5b\x70\x40\xb8\x7d\x97\xf0\xf2\xe5\xed\xed\x0a\xc8\x9a\xa2\x64\x58\x5f\xdc\xbe\xda\x69\x5e\x31\x7d\x87\xd9\x6e\x8b\x9a\x36\x5b\xa2\x16\x5e\x62\x0b\x76\x40\x04\x73\x1d\x73\x5d\x5c\xae\x69\x7a\x82\x67\xc2\xbf\x88\x40\xd6\xea\xd9\x63\xa0\x14\x0f\x7a\xf1\x9b\x5e\x74\x12\xc4\xa3\x2f\xf3\xb0\x3f\x1b\x4f\x92\x0f\x93\x38\xf8\x1a\x0f\x0e\xe3\xfb\xdf\x66\xa7\xcf\xd2\xd1\x30\x19\x05\xc0\x36\xe9\x26\xc8\x9a\xa1\xe2\xf2\x85\xca\x3c\x3c\x2c\x2d\x72\x9a\xc4\xa3\xa6\x67\xbb\x9c\x16\x19\xbe\xe3\xb1\x3b\x0e\x6d\x15\x61\x55\x74\x5d\x32\x14\x0c\xa5\x4b\xa5\x22\x38\x7a\x7a\x3a\x1b\x9f\x44\xfd\x57\xd1\xbb\x47\xd1\x34\x88\xfa\x0f\x57\x90\x8e\xdf\xf9\x5f\xe0\xeb\x49\xfc\xe2\xf3\x3c\xec\x7f\x7f\xf0\x38\x3a\xee\xa7\x67\xc7\xc9\x28\x48\x3e\xbe\x9d\x87\x83\xd9\xf8\x68\xb5\x6d\x7a\xf6\x32\x1d\x0e\xf2\x7a\xcb\x42\x6e\x97\x72\x22\x5c\x5e\xa4\xff\xc3\x59\x8a\x37\x2c\xe3\x72\x1a\xbd\xd7\xe5\xd4\xf3\x98\xeb\x9c\xff\xd6\x2b\xd0\xe1\xfb\x74\x3a\x8d\xc2\x27\x39\x94\x13\x41\xd7\x5c\x74\xed\x26\xf1\xa7\xe7\xc9\x51\x2f\xf7\xd9\x9c\x12\x41\x5b\x26\x11\x16\x08\xd6\xa1\x9e\x20\x9d\xee\xef\x5e\xd9\xd0\x75\xa4\x62\x13\xd7\xea\x68\x0f\x4b\xf5\xc6\xc2\xe8\x77\x5b\xe7\x36\x82\xa6\x82\xd1\xa8\x66\x1d\xd7\x42\x1b\x7a\xad\x2e\xe9\xfb\x70\x0b\xed\x43\x39\xfb\x36\x95\x8d\x0a\x20\xf5\x46\x4d\x45\x3b\x35\xc7\x71\xab\xbb\x05\xfb\xa6\xa4\xef\x21\xbc\xf3\x73\xab\x5f\xdb\xed\xfc\xb9\xdd\xb5\x1f\x01\x00\x00\xff\xff\xc6\xbe\x42\xea\x99\x03\x00\x00")

func _1571295328_create_rules_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1571295328_create_rules_tableUpSql,
		"1571295328_create_rules_table.up.sql",
	)
}

func _1571295328_create_rules_tableUpSql() (*asset, error) {
	bytes, err := _1571295328_create_rules_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1571295328_create_rules_table.up.sql", size: 921, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1572590132_create_collections_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x48\xce\xcf\xc9\x49\x4d\x2e\xc9\xcc\xcf\x2b\x4e\xb0\x06\x04\x00\x00\xff\xff\x3c\x40\x15\xcd\x19\x00\x00\x00")

func _1572590132_create_collections_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1572590132_create_collections_tableDownSql,
		"1572590132_create_collections_table.down.sql",
	)
}

func _1572590132_create_collections_tableDownSql() (*asset, error) {
	bytes, err := _1572590132_create_collections_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1572590132_create_collections_table.down.sql", size: 25, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1572590132_create_collections_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xb1\xae\xd3\x30\x14\x86\xf7\x3e\xc5\xd9\x92\x48\x2c\x45\xad\x84\x84\x3a\xb8\xc9\xa1\x58\x24\x4e\x71\x6d\x89\x4e\x71\x9a\x18\x11\x29\x71\xa2\xd4\x45\xea\x0b\x00\x82\x85\x1d\x89\x99\x15\x31\xf2\x3c\x05\xf5\x2d\x90\x53\x28\xb7\xf7\xf6\x0e\x77\xff\xbe\xdf\x47\xfe\x42\x8e\x44\x20\x08\x32\x8f\x11\x54\xd1\xd6\xb5\x2e\x6c\xd5\x9a\xad\x02\x7f\x04\xa0\xaa\x52\x41\x65\xac\x3f\x1e\x07\xc0\x52\x01\x4c\xc6\x31\x10\x29\xd2\x8c\xb2\x90\x63\x82\x4c\x3c\x72\x9c\xc9\x1b\xad\xe0\x6d\xde\x17\x6f\xf2\xde\x7f\x3c\x9d\x06\x10\xa6\x71\xec\xb6\x77\xf6\xf5\x93\x66\x33\xc9\x76\xa6\x2a\xda\x52\x67\x45\x75\x5e\x1a\xd4\x3a\xdf\xe8\xfa\x96\x7b\x01\x74\x7d\xdb\xe9\xde\xee\xb3\x6b\xc7\x0c\x44\xd1\x36\x8d\x36\xf6\x9e\x11\x88\xf0\x19\x91\xb1\x00\xcf\x83\x30\x4d\xdc\xcd\xe0\xfd\xfa\xf1\xed\xf8\xfe\xa3\x37\xe8\x76\xdf\x69\x05\xb6\x32\x7b\xb7\x3e\xb9\x61\x9e\xf1\xe3\x97\x77\x87\xcf\x1f\x7e\x7f\xff\x79\xf8\xfa\xe9\x24\x15\xbd\xce\xad\x2e\xb3\xdc\x3a\xb5\xd1\x5b\x9b\x37\xdd\xe5\x7b\xa1\xe4\x1c\x99\xc8\x04\x4d\x70\x25\x48\xb2\x1c\xc4\x5d\x57\x3e\x58\x84\x94\x81\x5c\x46\xee\x3b\xaf\x8e\x4a\x46\x5f\x4a\x84\x17\xb8\x06\xc5\x48\x82\x19\x8d\x5e\x29\xf0\x4f\x59\x02\x90\x2b\xca\x16\x30\x17\x1c\xd1\xd1\x4b\x4e\x13\xc2\xd7\x03\xee\xbb\xc4\xc1\x28\x00\x64\x0b\xca\x70\x46\x8d\x69\xa3\xf9\xff\x4b\x9e\x13\xbe\x42\x31\xfb\x1b\xf1\x5f\xd4\xd9\xdd\xa8\x4f\xff\x04\x00\x00\xff\xff\x7a\xe5\xf5\x69\x4b\x02\x00\x00")

func _1572590132_create_collections_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1572590132_create_collections_tableUpSql,
		"1572590132_create_collections_table.up.sql",
	)
}

func _1572590132_create_collections_tableUpSql() (*asset, error) {
	bytes, err := _1572590132_create_collections_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1572590132_create_collections_table.up.sql", size: 587, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1572590669_create_collection_items_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x48\xce\xcf\xc9\x49\x4d\x2e\xc9\xcc\xcf\x8b\xcf\x2c\x49\xcd\x2d\x4e\xb0\x06\x04\x00\x00\xff\xff\x25\x3b\xce\x82\x1e\x00\x00\x00")

func _1572590669_create_collection_items_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1572590669_create_collection_items_tableDownSql,
		"1572590669_create_collection_items_table.down.sql",
	)
}

func _1572590669_create_collection_items_tableDownSql() (*asset, error) {
	bytes, err := _1572590669_create_collection_items_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1572590669_create_collection_items_table.down.sql", size: 30, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1572590669_create_collection_items_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x6a\xeb\x30\x10\x45\xf7\xf9\x8a\x59\xda\xf0\x36\x79\x34\x50\x28\x59\x28\xf6\xb4\x35\x95\x65\xa3\x8c\x17\x59\x49\xaa\xa4\x52\x81\x2d\x07\x47\x0e\xfd\xfc\xe2\x94\xd2\x52\xd2\x45\x97\xc3\x9c\x73\xb9\xdc\x42\x22\x23\x04\x62\x3b\x8e\xa0\xed\xd8\xf7\xde\xa6\x30\x46\x15\x92\x1f\x4e\x1a\xb2\x15\x80\x0e\x4e\x43\x88\x29\x5b\xaf\x73\x10\x0d\x81\xe8\x38\x07\xd6\x51\xa3\x2a\x51\x48\xac\x51\xd0\xbf\x85\xfb\xae\x5f\x51\x2e\x4c\x88\xce\xbf\xfd\xf2\x3b\x9b\x7e\xf6\x1a\xce\x66\xb2\xaf\x66\xca\xfe\x6f\x36\x3f\x00\x3b\x79\x93\xbc\x53\x26\x69\x48\x61\xf0\xa7\x64\x86\xe3\x47\x9d\x12\xef\x59\xc7\x09\x8a\x4e\x4a\x14\xa4\xa8\xaa\x71\x4f\xac\x6e\x2f\xe2\x7c\x74\x7f\x16\xa1\x11\xd0\xb5\xe5\x32\xcf\xd5\xd0\x56\x56\x35\x93\x07\x78\xc2\x03\x64\xcb\x46\xf9\x2a\x07\x14\x0f\x95\xc0\x6d\x15\xe3\x58\xee\xbe\xb2\x1f\x99\xdc\x23\x6d\xe7\xf4\x72\x3b\x3c\xdf\x40\xd1\x70\xce\x08\x3f\x6f\x35\xc7\x60\x47\xe7\x95\x0d\x77\xef\x01\x00\x00\xff\xff\xf6\x66\x16\xd8\x91\x01\x00\x00")

func _1572590669_create_collection_items_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1572590669_create_collection_items_tableUpSql,
		"1572590669_create_collection_items_table.up.sql",
	)
}

func _1572590669_create_collection_items_tableUpSql() (*asset, error) {
	bytes, err := _1572590669_create_collection_items_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1572590669_create_collection_items_table.up.sql", size: 401, mode: os.FileMode(420), modTime: time.Unix(1500000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"1571295324_create_properties_table.down.sql":       _1571295324_create_properties_tableDownSql,
	"1571295324_create_properties_table.up.sql":         _1571295324_create_properties_tableUpSql,
	"1571295325_create_models_table.down.sql":           _1571295325_create_models_tableDownSql,
	"1571295325_create_models_table.up.sql":             _1571295325_create_models_tableUpSql,
	"1571295326_create_abstractions_table.down.sql":     _1571295326_create_abstractions_tableDownSql,
	"1571295326_create_abstractions_table.up.sql":       _1571295326_create_abstractions_tableUpSql,
	"1571295327_create_activations_table.down.sql":      _1571295327_create_activations_tableDownSql,
	"1571295327_create_activations_table.up.sql":        _1571295327_create_activations_tableUpSql,
	"1571295328_create_rules_table.down.sql":            _1571295328_create_rules_tableDownSql,
	"1571295328_create_rules_table.up.sql":              _1571295328_create_rules_tableUpSql,
	"1572590132_create_collections_table.down.sql":      _1572590132_create_collections_tableDownSql,
	"1572590132_create_collections_table.up.sql":        _1572590132_create_collections_tableUpSql,
	"1572590669_create_collection_items_table.down.sql": _1572590669_create_collection_items_tableDownSql,
	"1572590669_create_collection_items_table.up.sql":   _1572590669_create_collection_items_tableUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"1571295324_create_properties_table.down.sql":       &bintree{_1571295324_create_properties_tableDownSql, map[string]*bintree{}},
	"1571295324_create_properties_table.up.sql":         &bintree{_1571295324_create_properties_tableUpSql, map[string]*bintree{}},
	"1571295325_create_models_table.down.sql":           &bintree{_1571295325_create_models_tableDownSql, map[string]*bintree{}},
	"1571295325_create_models_table.up.sql":             &bintree{_1571295325_create_models_tableUpSql, map[string]*bintree{}},
	"1571295326_create_abstractions_table.down.sql":     &bintree{_1571295326_create_abstractions_tableDownSql, map[string]*bintree{}},
	"1571295326_create_abstractions_table.up.sql":       &bintree{_1571295326_create_abstractions_tableUpSql, map[string]*bintree{}},
	"1571295327_create_activations_table.down.sql":      &bintree{_1571295327_create_activations_tableDownSql, map[string]*bintree{}},
	"1571295327_create_activations_table.up.sql":        &bintree{_1571295327_create_activations_tableUpSql, map[string]*bintree{}},
	"1571295328_create_rules_table.down.sql":            &bintree{_1571295328_create_rules_tableDownSql, map[string]*bintree{}},
	"1571295328_create_rules_table.up.sql":              &bintree{_1571295328_create_rules_tableUpSql, map[string]*bintree{}},
	"1572590132_create_collections_table.down.sql":      &bintree{_1572590132_create_collections_tableDownSql, map[string]*bintree{}},
	"1572590132_create_collections_table.up.sql":        &bintree{_1572590132_create_collections_tableUpSql, map[string]*bintree{}},
	"1572590669_create_collection_items_table.down.sql": &bintree{_1572590669_create_collection_items_tableDownSql, map[string]*bintree{}},
	"1572590669_create_collection_items_table.up.sql":   &bintree{_1572590669_create_collection_items_tableUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
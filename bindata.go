// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// tpl/curd.tpl
// tpl/entity.tpl
// tpl/example.tpl
// tpl/markdown.tpl
// tpl/tables.tpl
package gomygen

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

var _tplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5d\x6b\x24\x45\x14\x7d\xaf\x5f\x71\xb7\x59\x42\x77\x32\x54\x56\x10\x1f\x02\xa3\x24\x33\x1d\x77\x30\x99\xc4\x99\x8e\x22\xcb\x22\x35\xd3\x77\x42\xb3\xfd\x31\xa9\xae\xce\x07\x45\x81\x82\x2f\x0b\xae\x3e\x6d\x84\x55\x94\x2c\x0a\xbe\xb8\x9b\x07\xc5\x47\xff\x8c\x19\xf3\x33\xa4\xba\x7b\x66\x7a\x26\xf3\xd1\x81\x28\xfb\xb0\x0f\x21\x4c\x57\xd5\x3d\xb7\xce\xb9\xa7\xee\x15\x67\x7d\x04\x29\x69\x5b\xf0\xa4\x2b\x1c\xd6\xf1\xb1\xc9\x02\x54\x6a\x37\x72\xd1\x87\x38\xfd\x0c\x92\xd4\xb7\x60\x35\x3e\xf2\x69\x7d\x8b\x28\x42\x7a\x49\xd8\x85\x26\x9e\xcc\x3a\x68\xba\x9d\xe1\x56\x0b\x56\xe7\x87\x96\x84\xa3\x48\x78\x08\x2b\x73\xf7\x68\xd8\x0d\xb7\x53\x21\x4a\x83\xae\xaf\x5f\x7f\xf3\xe7\xd5\xb7\xe7\x83\xa7\x5f\x0c\x7e\x78\xfa\xcf\x8b\xaf\xae\x2f\x7e\xbd\xfa\xed\xbb\xc1\xab\x3f\xb2\x7c\xcc\x60\x01\x9c\x05\x87\x28\x6a\x91\x9f\x04\x61\x6c\x5a\xfa\x5e\x5e\x78\x38\xce\xc1\xd0\x1c\x6c\xfa\xfe\xb6\x87\xbe\xbb\xe3\xc5\x42\x29\x30\x8a\xa0\x57\x3f\xbf\xb8\xbe\xf8\x7a\xf0\xfc\x72\xf0\xec\x15\x2d\x8b\xd7\x8a\x4e\x62\x33\x3e\xf2\x9d\x53\x91\x23\x56\xa0\xcf\x38\x0b\x62\xa0\x94\x7a\xa1\x40\xde\x63\x5d\x94\xca\x02\x93\x47\x27\x71\x0b\xe3\xc4\x17\xf0\xe8\xf1\xcc\xb8\x15\x40\xce\xf5\x5f\xc4\x2d\x90\xe4\x28\x41\x7e\x96\x7d\xdb\xa8\x42\x40\xeb\x5b\xf4\x63\xfd\x29\x07\x1c\x22\x51\x4a\x2d\xe2\x62\x0f\x39\xa4\x27\x68\xcd\x8f\x62\x34\x2d\xe2\xf5\xd2\xb3\xf7\xaa\x10\x7a\x63\x35\x88\x22\xbd\x68\xb8\xb5\x89\xa7\xc2\xd4\x58\x3c\x3a\xd1\x20\x52\xd2\x66\xe2\xfb\x37\x32\x93\x8a\xe8\x50\xd5\xfc\x58\xbb\xcb\x42\x93\x48\xc9\x59\x78\x88\x90\x1e\x49\x79\x8d\x1b\x61\x2f\x52\x6a\x85\x47\x27\x54\x4a\xfa\x30\x09\xfa\xf9\xcd\xd6\xd7\xa5\xa4\xb5\x28\x08\x30\x14\x4a\x11\x29\x31\x74\x95\x4a\x73\xd4\xc9\xdd\xab\xa6\xa9\x4a\xd2\x8d\x42\xe1\x85\x09\x12\x45\x0a\x7c\x55\x81\xf5\xfb\x18\xba\x05\x0e\x2b\xb3\xab\x4a\xce\xcf\x4a\x4a\x4d\xc8\x11\xd0\x0f\x23\x47\x7b\xc2\xe8\xf9\x11\x13\xef\xbd\x6b\xe8\xa5\x42\xae\x1b\x37\xb2\xa7\xdb\xd9\xce\x19\xb7\xf0\x63\x84\xa9\xb0\x5e\x58\x2e\x68\x23\x2c\x1d\x52\x78\x01\x52\xc7\x0b\xb0\x44\x58\xbd\x6d\x76\xd4\xe5\x67\xdb\x59\x09\xcf\x16\x2b\xff\xa7\x2c\xad\xcd\xb0\x94\xc6\xf6\x79\xf6\x7c\x64\x9f\x5b\xb8\x47\xd7\x72\x19\xe7\xe4\x85\x50\xda\x36\x93\x8e\xc9\x81\x26\x1c\x33\xd7\x1d\x6f\x8a\x15\x8a\x09\x8d\x7c\xf0\xff\x17\x3d\xbc\x2d\x7a\xa5\x14\x99\x28\xf9\xcf\x07\xdf\xff\x3e\x38\xbf\xbc\x45\xb5\xb7\xd9\x31\x4e\x37\x8a\x63\xe6\x27\x78\xb3\xda\x3b\xd0\x89\x22\x7f\xaa\xa8\x63\x11\x88\xc9\x56\xb0\xcf\xb1\xcf\xf8\x30\xe8\x82\x7a\xce\x5a\x83\x0e\x30\xea\x0c\x3c\x7f\x42\xf3\x78\xe9\x9a\x7d\x8a\x5d\x33\x4d\x69\xb1\x3d\x8e\x19\x07\xd6\xeb\x61\x57\xd4\xa2\x24\x14\x90\x2a\x4f\x0a\x5f\xb2\xb0\x55\xc8\x40\xa8\x6e\x91\x9b\xe9\x2a\xba\x8b\x9a\x52\x47\xbf\xf2\x85\xb8\xef\xc3\x83\x09\xce\x07\xe7\x97\x57\x2f\x7f\xfc\xfb\xaf\x8b\xc1\x97\xaf\x4b\x51\x5e\xe3\xc8\x04\x66\x57\x9a\xbd\xd1\x02\xd3\x67\xb1\x68\xb8\xd9\x25\xa6\x29\x3f\xf2\x1d\x3c\x15\x9a\x20\xa3\xd1\x6c\xdb\x2d\x07\x1a\x4d\x67\x0f\x0c\x58\xd3\x8f\xc3\x41\xbf\x8f\xbc\x10\x0c\xd6\xc0\x00\x53\x4a\xda\x08\x63\xe4\xa2\x30\x67\x58\xf0\xc9\xe6\xce\x81\xdd\x2e\xac\xee\x32\xfe\x44\x29\xcb\x58\xa2\x2b\xde\x99\xb0\xe3\x17\x22\x4b\x20\x7b\x1d\x32\xb9\x4b\x3f\x4f\xb3\xd3\xc8\x28\x9c\x52\x7d\x47\x7f\xcc\x90\x16\xaa\x3e\xa1\xf0\x6d\x4d\x75\xd0\x77\x97\x2b\x3c\xdb\x4e\x05\x6d\x0f\xf6\xeb\x9b\x8e\xbd\x50\xd6\xb6\xed\x64\x6b\x1a\xaf\x38\x41\x7e\xfa\xd0\x6e\xd9\x7a\x69\x9f\x7b\x01\xe3\x67\x1f\xe1\x99\x52\x50\x85\x0f\x0c\x92\x37\x33\x2d\x2b\x7b\x82\xe6\xa3\xc7\x05\x97\x57\xe0\x81\x35\x52\xe4\xbe\x57\x81\xfb\xc7\xcc\xd7\x5b\x73\x08\x1d\x3d\x85\x51\x2a\x0f\x33\x9a\x80\xb2\xdf\x15\x90\x52\x9f\xd1\xc2\xe4\x0a\x0d\x07\xdd\x80\x8e\xde\x1a\x9c\x1a\x12\x33\x96\x7f\xfa\xe5\xfa\xf5\xcb\xe2\xb4\x5b\x8a\xeb\x6d\x2f\x74\x97\x30\xcd\x31\xd6\x89\x97\x9c\x6e\x8b\x12\xb4\xed\x1d\xbb\xe6\x68\x05\x02\x3a\x31\xc5\xaf\x81\xb1\xdd\xda\xdb\x9d\x27\x0e\xc9\x21\x87\xc5\x97\x9e\x1e\xcd\xe4\xa9\x7b\xee\x60\x52\xd9\xf6\x78\x2c\x96\xdf\xbd\xec\x7c\x72\x17\x17\x4f\xab\x72\xa7\xb1\xdb\x70\xe0\x1d\x63\xc2\xf5\x63\x12\x4a\xbc\x20\x77\x40\x8e\x36\xfa\x1b\xc8\xcd\x5e\xab\x6e\xb7\x60\xeb\x33\x68\xd4\xa1\x6e\xb7\x6b\xff\x09\x59\x19\x53\xe5\x9a\x91\xee\x69\xa6\x05\x66\x77\xdc\x34\x97\xde\x1d\x6a\x7b\x07\x4d\xc7\x5c\xb5\x60\xa1\x09\xe6\x0f\xbc\x4b\x6e\x74\x63\x8e\x5d\x49\xb3\x2b\xc3\xc1\xbf\x01\x00\x00\xff\xff\x77\x0b\xb6\x7b\x56\x10\x00\x00")

func tplCurdTplBytes() ([]byte, error) {
	return bindataRead(
		_tplCurdTpl,
		"tpl/curd.tpl",
	)
}

func tplCurdTpl() (*asset, error) {
	bytes, err := tplCurdTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/curd.tpl", size: 4182, mode: os.FileMode(420), modTime: time.Unix(1552876725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplEntityTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xcd\x8a\x83\x30\x10\x80\xef\x79\x8a\x39\x78\x5c\xe2\x7d\x61\x4f\x0b\x1e\x3d\x2c\x3e\x80\x71\x1d\x76\x6d\xf3\x23\x71\x44\x64\x98\x77\x2f\xa9\x8d\x96\xb6\x74\x8e\x5f\x32\xc3\xf7\xa9\xb2\x64\xd6\x8d\xe9\x2c\x7e\x07\xe7\xd0\x93\x88\xa2\x75\x44\xc8\x58\x04\x26\x8a\xf3\x2f\x01\x2b\xe6\x68\xfc\x1f\x42\x71\xfa\x80\x62\x20\x74\xf0\xf9\x05\xba\x1a\xd0\xf6\x93\x08\xf3\x95\xe9\xda\xb8\xb4\xb5\x4d\x86\xcd\x3a\xde\x60\x9b\x51\x15\xa2\x33\x94\xb7\x5b\x38\x26\x49\x6d\x7f\x7e\xd0\x99\x78\x16\x51\xcc\xe8\x7b\x11\x25\xaf\x8c\x81\xfe\x0d\x81\x33\x2b\x74\x08\x7e\xb6\x56\xef\x0d\xf5\x6c\xed\xfb\x8e\x18\x96\xc7\x8c\x18\x96\xbb\x8a\x0c\xd2\xa5\xbd\xe2\xf0\x4c\x6f\xcf\x9a\x97\x00\x00\x00\xff\xff\xc0\xeb\x0a\xff\x59\x01\x00\x00")

func tplEntityTplBytes() ([]byte, error) {
	return bindataRead(
		_tplEntityTpl,
		"tpl/entity.tpl",
	)
}

func tplEntityTpl() (*asset, error) {
	bytes, err := tplEntityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/entity.tpl", size: 345, mode: os.FileMode(420), modTime: time.Unix(1552876725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplExampleTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\x6d\x6b\x1b\x47\x10\xfe\x3e\xbf\x62\xbb\x60\xe7\xce\x5c\xee\x9c\x3a\x29\xe1\x40\xb4\x95\xe4\xd0\x40\x92\xba\xd8\xa5\x1f\x4a\x09\xeb\xbb\x95\xb4\xf8\x6e\xf7\xb4\xbb\x97\xc6\x15\x02\x05\x4a\x93\x90\xc4\x0e\xa4\x49\xa0\x75\xfa\xee\xbe\x50\xa8\xfd\xa1\x34\x46\x49\xe9\x9f\xf1\x9d\xec\x7f\x51\x56\x77\x92\xa5\xb8\x69\x4a\x68\xe9\xb7\xd9\xd9\x99\xd9\x67\x9e\x79\x66\x13\x12\xac\x91\x26\x45\xb1\x08\x69\xa4\x00\x58\x9c\x08\xa9\x91\x05\x38\x24\x9a\xac\x12\x45\x3d\xd5\x8e\x30\xe0\x46\xac\x31\x5c\x46\xb8\xc9\x74\x2b\x5d\x75\x03\x11\x7b\x4d\x71\x52\xb5\xa3\x93\xa1\x64\x57\xa8\xf4\xe2\xf5\x22\x50\x53\xa5\x19\x6f\x62\xc0\x4a\x4b\xc6\x9b\x0a\x03\xe6\x54\x7b\xa9\x8c\x30\xd8\x00\x7a\x3d\xa1\xa8\x5e\xad\x09\xde\x60\xcd\x45\xae\x99\x5e\x47\x4a\xcb\x34\xd0\xa8\x03\x6f\x09\xa5\x11\x42\xa8\x48\x45\x9e\x97\x6d\xed\x66\x8f\x7a\xb0\x64\x40\x21\x84\x18\xd7\xc8\xf3\x06\x3f\xef\x64\x9b\xdf\xc2\x25\x12\xd3\xa9\xe0\xc1\xa7\x3f\xe6\x37\x1e\xc3\x12\x51\x6a\xba\xc8\xce\x27\x83\xaf\xae\x41\xbd\x5a\x64\x1c\xf9\xfb\xf7\xb2\xbb\x77\xa0\xd6\x22\x52\x51\x3d\x51\xe7\xe9\x03\x13\xbf\xc2\x62\xfa\x91\xe0\xf4\xe8\x22\x7f\xf8\x5b\x76\xbb\x0f\x17\xc9\xd5\xf3\x61\x44\x4b\x34\xf9\x56\x2f\xfb\xee\x87\xc1\x4f\xfd\xc3\x87\xbf\x1e\xfc\xf1\x45\xbe\xb1\x6d\x02\xde\x4e\x28\x9f\x0a\x28\xae\xf2\xfb\xbb\xd0\x05\xf0\xbc\xf1\x31\xbf\xf3\x4b\xd6\xbf\x07\x8d\x94\x07\xe8\x3c\x67\xba\x5e\xb5\x82\x46\xf3\x19\x86\x6c\x34\xa7\xda\x91\x5b\xaf\xa2\x0e\xb0\x46\x09\x48\xb9\x8b\xed\x94\x44\xe7\x44\x14\x9a\x14\x77\x84\xd7\x41\x18\xdb\xa8\x03\x93\x3e\x54\x41\xf8\xc4\x9b\x8a\x11\x6f\xb9\x45\x78\xb3\x45\xd8\x09\x0c\x5d\x08\x15\x47\x7e\x05\x35\x62\xed\x2e\x27\x92\x71\xdd\xb0\xf0\x8c\xf2\x67\xd4\x1b\x3a\x48\x2c\x63\x85\xb6\x37\xa3\x5e\x0f\x0a\x8a\x2a\x33\x6a\x36\x31\x96\x29\x5b\xd1\x32\xa5\xb3\x91\x08\x2a\x17\x44\x40\xa2\x59\xcd\x62\x7a\xd9\xbc\x55\x99\x51\xd8\x19\xbe\x6e\xf8\x2e\x2c\x33\x93\xc2\x32\x23\x2e\x7d\x42\x96\x56\x31\x99\xc2\x2e\xa7\xe1\x40\x2a\x23\xf7\x9d\x94\xca\xf5\x45\x15\x90\x84\x4e\xb5\x68\x3b\x60\x43\x20\x38\xa7\x81\x66\x82\x3b\x88\x4a\x69\x1a\x31\x24\x19\xe6\x2d\x5c\xc8\xd1\x41\xa1\xe2\xb6\xa1\xcc\x04\xbc\x52\x41\x9c\x45\xa8\x03\x09\xe1\x2c\xb0\xa8\x94\x36\x74\x27\xca\xb8\xcb\x54\x97\xa3\xab\x09\xce\xd5\xf0\xc9\xd2\x61\x1f\x8f\x33\x1a\x98\x8a\x33\x0e\x1b\x24\xd5\xa9\xe4\xe8\x28\x1c\xba\x70\x85\x48\x54\xaf\x8e\x86\x38\x3c\x76\x3a\x43\x7a\xba\xdd\xfa\x2a\x9a\x1b\x1f\x2e\x9a\x3d\x84\x42\x0d\x8c\x33\x5d\x1b\x57\xb1\xca\x91\x9a\x36\xa7\xc5\x51\xac\x8d\x8f\x23\x33\x86\x96\x50\x1a\x3b\xc3\x85\xf1\x17\x16\xe6\xcf\x3a\xc3\x35\xf1\xb1\x14\x62\xe8\x27\x4a\xf9\xf8\xd4\xab\x0b\xa7\xcf\xbc\x86\x9d\x72\x25\x7c\xbc\xc6\x78\x18\xd1\x10\x3b\xa3\x65\xf0\x71\xaa\x1b\x67\xe3\xd5\xd3\xd8\x19\xa9\xd9\x3f\x35\xef\x8c\xa4\xef\xa3\x33\x0e\x74\xa1\x5e\x45\x95\x09\xd1\xda\x30\xd9\x54\x05\x5d\xa2\x1f\x8e\x1d\x56\xbd\x6a\x17\xca\xcf\xbf\xdc\x3e\xd8\xf9\x26\xbf\xd9\xcb\xb7\x6e\x0e\x3e\xfb\xb8\x58\x81\xa2\xe3\x15\xaa\xf4\x38\xe3\x1c\xe3\xe1\x7b\x2d\x2a\xa9\xa5\xd1\x5c\xf9\xab\xb8\x2b\x86\x85\x67\x89\x81\x55\x21\xd6\x0c\x2f\xe3\xdc\x0e\x78\x9e\x16\xa1\x80\x2e\x48\xaa\xd2\x48\x8f\x15\x32\x81\xd0\x35\x2f\x58\xb3\x26\xd9\x36\xbf\xc1\xf5\xdf\xf3\xfb\xbb\x87\xd7\x6f\x1f\x6e\xf5\x0e\xbe\xbf\x96\xdf\x7a\x9a\x6f\x6c\xef\xef\xf5\xf6\xf7\x6e\xa9\x76\x74\x4c\x43\xda\x5d\x94\x52\xc8\x91\x8a\xcc\x0a\x2d\x99\x0d\x8a\xb8\x55\x3c\xf9\xfe\xfc\x07\xe6\xc2\xf3\x0e\x36\x1e\x67\x9b\x0f\xcc\x2f\x70\x77\x63\x7f\xaf\x97\x3f\xfa\xfa\xb9\x3d\x5f\x20\x4a\xbf\xb8\xdd\xbf\x6d\x49\x2a\x6d\x71\x16\x1d\x17\xfd\x0b\x01\x4f\xa1\xed\x3d\x19\x92\xb1\xf9\x17\x20\x6b\x22\xe5\xff\x37\xca\xec\xc6\xe7\xd9\x93\xfe\x73\x79\xac\x49\x4a\xf4\x3f\x10\x4e\x18\xbe\x84\x6c\xca\xe2\xb3\x61\xf8\x2f\xc3\x5f\xbc\x4a\xe2\x24\xa2\xef\x26\xe1\x7f\x07\xbe\x2c\xfe\xd2\xe0\xff\x0c\x00\x00\xff\xff\xdc\x57\x0d\x95\x31\x08\x00\x00")

func tplExampleTplBytes() ([]byte, error) {
	return bindataRead(
		_tplExampleTpl,
		"tpl/example.tpl",
	)
}

func tplExampleTpl() (*asset, error) {
	bytes, err := tplExampleTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/example.tpl", size: 2097, mode: os.FileMode(420), modTime: time.Unix(1552904525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMarkdownTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\xc1\x6a\xea\x40\x14\x86\xf7\xf3\x14\x07\xe2\x42\xe1\x26\x0f\x20\xdc\xbb\xb9\x6e\xe4\x5e\xac\x0b\xe9\x46\x5c\xc4\x7a\x5a\xa4\x49\x5a\xa2\x81\x86\x99\x01\x17\x85\x16\x4a\x5b\x17\x05\xb3\x2a\x14\x44\xec\xa2\x4a\x69\x0b\x62\xfa\x38\xce\xa8\x6f\x51\x32\x63\x62\x5b\x9b\xc5\x09\xe7\xff\x0f\xe7\x4f\xbe\x19\xc3\x80\xd5\xc3\x58\xf6\x62\x39\x3a\x27\x84\x89\xf9\xad\xb8\x1c\xb0\x44\x13\xfd\x6b\x60\xb0\x9a\xbe\xca\xe8\x06\x18\x88\xe1\x85\x7c\x19\x03\x23\xac\x68\x9a\xa6\x2a\x69\x05\xfd\x22\x94\xfa\xb6\x77\x84\x60\xd5\xec\xa6\x83\xff\xdb\x9d\x2e\xe7\x0c\x28\xb5\xca\x5e\x0b\xcf\x38\x67\xf5\xa4\x51\x66\xc5\x76\x91\xf3\x46\xde\xd8\xba\x5f\xad\x02\xe8\x87\x51\x6a\xfd\x3d\x71\x5d\xf4\xd4\xb2\x24\x04\xbd\x16\xe7\x24\x4b\xcb\x1d\x63\xf8\x0b\x72\xed\x2e\xba\x50\xfc\x0d\x56\x09\x3b\x07\x3a\x9b\x18\x86\x01\x94\x2a\x2b\x4d\xb1\xd2\xfe\x53\x16\xf9\x93\x4d\x65\x51\x84\x98\x0a\xc2\xd3\x40\x4e\xde\x34\x84\x04\x8f\x6a\x81\xc1\xf2\x39\x16\xf7\x57\xc0\x40\x46\x53\xd1\x1f\x2d\x66\xf3\xe5\xe3\x9c\xc1\x3a\x8e\x56\x93\xa1\xe8\xbd\x67\x8e\x8c\xa6\x8b\x59\xbc\xbe\x9b\xec\xb2\xdc\xa0\x34\x77\x81\x7e\x13\xb6\x64\xf5\x47\x6e\xc8\x2a\x32\x4e\xe0\x7a\xfa\x2f\x00\x14\x6b\x2d\xd5\xc2\x53\xac\x04\x6e\x13\x7d\xce\xb5\x5e\xee\x54\x02\xc7\xd9\x1c\x48\x09\x0f\xed\xc0\xe9\xee\xdb\x4e\x80\xe9\x40\xd5\x6f\xbb\xb6\x1f\xfe\xc3\x30\x55\xf4\xaa\x1f\xe9\xd7\x6b\x7b\xd5\x46\xde\xc8\xae\x4e\x21\x75\x3e\x02\x00\x00\xff\xff\xd7\x32\xbc\x15\x50\x02\x00\x00")

func tplMarkdownTplBytes() ([]byte, error) {
	return bindataRead(
		_tplMarkdownTpl,
		"tpl/markdown.tpl",
	)
}

func tplMarkdownTpl() (*asset, error) {
	bytes, err := tplMarkdownTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/markdown.tpl", size: 592, mode: os.FileMode(420), modTime: time.Unix(1552988812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplTablesTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x4a\xce\xcf\x2b\x2e\x51\xd0\xe0\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x50\xc9\x2c\x49\xcd\x55\xb0\xb2\x55\xd0\xab\xad\xe5\xaa\xae\x06\x73\xf5\x42\x0b\x0a\x52\x8b\x42\x12\x93\x72\x52\xfd\x12\x73\x53\x6b\x6b\x15\x14\x14\x14\x6c\x15\x94\x60\xd2\x48\x32\x4a\x0a\xfa\xfa\x30\x61\xe7\xfc\xdc\xdc\xd4\xbc\x92\xda\xda\xea\xea\xd4\xbc\x94\xda\x5a\x2e\x4d\x40\x00\x00\x00\xff\xff\xe2\xac\xd0\x9a\x72\x00\x00\x00")

func tplTablesTplBytes() ([]byte, error) {
	return bindataRead(
		_tplTablesTpl,
		"tpl/tables.tpl",
	)
}

func tplTablesTpl() (*asset, error) {
	bytes, err := tplTablesTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/tables.tpl", size: 114, mode: os.FileMode(420), modTime: time.Unix(1552876725, 0)}
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
	"tpl/curd.tpl":     tplCurdTpl,
	"tpl/entity.tpl":   tplEntityTpl,
	"tpl/example.tpl":  tplExampleTpl,
	"tpl/markdown.tpl": tplMarkdownTpl,
	"tpl/tables.tpl":   tplTablesTpl,
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
	"tpl": &bintree{nil, map[string]*bintree{
		"curd.tpl":     &bintree{tplCurdTpl, map[string]*bintree{}},
		"entity.tpl":   &bintree{tplEntityTpl, map[string]*bintree{}},
		"example.tpl":  &bintree{tplExampleTpl, map[string]*bintree{}},
		"markdown.tpl": &bintree{tplMarkdownTpl, map[string]*bintree{}},
		"tables.tpl":   &bintree{tplTablesTpl, map[string]*bintree{}},
	}},
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

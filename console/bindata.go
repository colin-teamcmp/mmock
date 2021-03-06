// Code generated by go-bindata.
// sources:
// tmpl/css/style.css
// tmpl/index.html
// tmpl/js/mapping.js
// tmpl/js/request_logger.js
// tmpl/js/sorting.js
// tmpl/js/util.js
// tmpl/swagger.json
// DO NOT EDIT!

package console

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

var _tmplCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xdd\x8a\xdb\x30\x10\x85\xaf\xab\xa7\x18\xd8\xdb\xda\x34\xdb\x2c\xdd\x2a\x50\xd8\xfe\xec\x6b\x2c\x23\x69\xec\x08\x2b\x1a\x33\x52\xea\x4d\x43\xde\xbd\xc8\x56\x9a\x16\xea\xbd\x32\x3e\xdf\x9c\x23\x74\x46\x5a\x37\x13\x99\xc1\xe7\x26\x59\xe1\x10\x0c\x0a\x9c\x15\x00\xc0\xe4\x5d\xde\x6b\xd8\xdc\x8f\xaf\x3b\x75\x51\xff\x19\x6c\xb2\xa0\x1d\xea\xb8\x41\x3b\xf4\xc2\xc7\xe8\x1a\xcb\x81\x45\xc3\xdd\xb7\xa7\x1f\xdb\xe7\xa7\xdd\x82\x59\x1c\x49\x13\xa8\xcb\x1a\x36\xe3\x2b\x24\x0e\xde\xc1\x9d\xb5\x76\x35\x7d\x7f\x3c\x98\xf5\xf4\xed\xfd\xa3\xb1\xf8\xa6\x59\xef\xf9\x27\xc9\x7a\xc4\xe6\xfb\xa7\xed\xd7\xe7\x12\xa1\xda\x44\x81\x6c\x26\xf7\x22\x3c\xc1\x19\x3a\x8e\xb9\x99\xc8\xf7\xfb\xac\xc1\x70\x70\xbb\x8b\x6a\x53\x16\x1f\x7b\x38\x43\x0d\xe8\x85\x28\xee\xe0\xa2\xda\x78\x3c\x98\x72\xd2\x95\x38\x94\x81\x05\x63\x4f\x33\x36\xcc\x81\x30\xde\xb8\x09\x47\xaa\xc6\x10\x6e\xf2\x01\x7b\x8a\x19\x67\x32\xd0\xe9\x06\x84\x5c\x11\x55\x46\x13\x08\x5a\xcb\x21\xe0\x98\xa8\xf5\x11\xce\xea\x9d\xf3\x69\x0c\x78\xd2\x33\x6d\x84\xa7\xe5\x4a\xcb\xef\xd2\xc1\x17\xc8\x86\xdd\xa9\x7c\xa5\xde\xb5\x0c\xea\x3f\xd4\xbd\x57\x7f\xe9\xb3\xb2\x5e\x5c\xd7\x75\xb5\x35\x96\xec\x63\xff\xe2\x2d\xc7\x54\xe7\x47\x74\xce\xc7\xbe\xee\xfa\xa1\x3c\x9f\x22\x5f\xbd\x88\xb8\x08\x73\xc5\xc9\xff\x22\xfd\xa1\x7d\xa4\xc3\xbf\x6b\xa8\x59\x57\x93\xfb\xfc\xf0\x71\x5b\xce\xfc\x1d\x00\x00\xff\xff\x65\xd5\x4f\x11\xb0\x02\x00\x00")

func tmplCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_tmplCssStyleCss,
		"tmpl/css/style.css",
	)
}

func tmplCssStyleCss() (*asset, error) {
	bytes, err := tmplCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/css/style.css", size: 688, mode: os.FileMode(420), modTime: time.Unix(1516991133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x3a\xeb\x72\xe3\xb6\xb9\xff\xfd\x14\x58\x26\x39\x96\x67\x97\xa4\x65\xd9\x6b\xc7\xa6\x78\xc6\xb1\x95\xb3\xeb\x91\x76\x9d\xb5\x62\x9f\xa4\xd3\xd9\x01\x09\x48\x84\x0c\x02\x34\x00\xea\x12\x8f\x67\xfa\x04\x7d\x8f\xfe\xe8\x4b\xb4\x6f\xd2\x27\xe9\x80\xa0\x28\x8a\xba\xda\x3b\x49\xd3\xe8\x87\x44\x82\xdf\x15\xdf\x15\x1f\xe5\xbd\xba\xfc\x78\xd1\xfd\xe9\xba\x05\xde\x75\x3b\x6d\x7f\xc7\x8b\x54\x4c\xfd\x9d\x1d\x2f\xc2\x10\xf9\x3b\x00\x00\xe0\xc5\x58\x41\x10\x29\x95\xd8\xf8\x21\x25\xc3\xa6\x75\xc1\x99\xc2\x4c\xd9\xdd\x49\x82\x2d\x10\x9a\xbb\xa6\xa5\xf0\x58\xb9\x1a\xff\x0c\x84\x11\x14\x12\xab\xe6\x8f\xdd\xef\xed\x13\x0b\xb8\x39\x25\x45\x14\xc5\x7e\xa7\xc3\xc3\x7b\x70\xc1\x99\xe4\x14\x7b\xae\x59\x34\x00\x32\x14\x24\x51\x40\x8a\xb0\x69\x69\x8e\xf2\xd4\x75\x43\x8e\xb0\x33\x78\x48\xb1\x98\x38\x21\x8f\x5d\x73\x69\x1f\x38\x75\xe7\xd0\x89\x09\x73\x06\xd2\xf2\x3d\xd7\xa0\xae\xa3\x83\xd8\x40\x3a\x21\xe5\x29\xea\x51\x28\x70\x46\x0c\x0e\xe0\xd8\xa5\x24\x90\x6e\x04\x19\xa2\x38\x80\x42\x3a\x03\xe9\x1e\x38\xfb\xce\xfe\xfc\x5a\x99\x89\xe1\x42\x09\xbb\x07\x02\xd3\xa6\x25\xd5\x84\x62\x19\x61\xac\x2c\x10\x09\xdc\x9b\x71\x8d\xe1\x38\x44\xcc\x09\x38\x57\x52\x09\x98\xe8\x1b\xcd\xb8\x58\x70\x1b\x4e\xc3\x79\xeb\x86\x52\xce\xd6\x32\xb5\x42\x29\x2d\x40\x98\xc2\x7d\x41\xd4\xa4\x69\xc9\x08\x36\x4e\x0e\xed\xfa\xc3\x49\xdc\xbd\xfa\x78\x7e\x33\x3e\x19\xd4\xcf\xd3\xd7\xf0\xe8\xee\xf2\x96\x5d\x93\x03\x7a\xff\x7d\x6f\x34\x6a\x9d\xc3\x93\xe8\xf2\x12\x0d\x7e\xa6\x49\x1b\xf7\xc7\xd1\xe0\xb6\xd3\xaa\xf7\xfa\x83\xbb\xeb\xff\x8b\xef\x7f\x91\xc7\x16\x08\x05\x97\x92\x0b\xd2\x27\xac\x69\x41\xc6\xd9\x24\xe6\xa9\xb4\xfc\x5f\x59\x29\x5b\x45\x38\xc6\xeb\x54\xeb\xb5\xef\x0e\x3e\xec\xd7\x69\xe7\x61\x00\xef\xbf\xbb\x1f\x37\xa8\xdb\xf9\xb6\x05\xa3\x74\x94\xdc\xf4\xf0\x87\xe1\xed\xdb\xc6\xd5\x11\xfe\x85\x35\xd2\x9f\x7f\x81\x49\x77\x3f\x3d\x6e\xfd\x24\xff\xbf\x33\xf8\xe1\xf6\xf5\x7e\x8b\x1d\x89\x4d\xaa\x2d\xf3\x8a\x6d\x55\x19\x54\xcd\x33\x58\xaa\xc2\x7e\x7c\x13\x5c\x5d\xb6\xde\x11\x48\x7b\x71\xfa\xdd\x77\x3f\x5c\xbf\x3d\x3f\xfc\x41\x24\xe2\xe1\xe8\xe3\x6d\xef\xae\x71\x7c\xfd\xe9\x53\x63\x70\xd4\x6a\x3f\x8c\xa5\xac\x4f\x6e\x1f\x3e\x2a\x86\x13\xf6\xee\xf6\xfa\x5b\x78\x75\x3c\xbe\x59\xad\xc2\x52\xe7\x9b\x37\xcb\x06\x0f\x2f\xe4\x6f\xd8\x88\x40\xca\xfb\x6e\xdd\x69\x1c\x3a\xdf\x56\xcc\x64\x9e\xcd\xec\xb4\xe0\x0b\x6a\x92\xe0\x3c\xda\x35\xc0\xcb\x43\x6e\xa5\x40\x83\x15\xf2\x2c\x0b\xc2\x9c\x6f\x49\xa8\x2b\x38\x84\x37\xd9\xaa\x65\xc4\x19\x48\x37\x55\x84\xae\x49\x14\x1b\xb0\x05\x7e\x48\xb1\x54\x9f\x29\xef\xf7\xb1\x78\x39\x9d\x18\x26\x09\x61\xfd\x97\x13\x90\x5c\xa8\x15\x04\x96\xc7\x6d\xc5\x56\xb9\xc3\x68\x83\x67\x50\x4e\xc9\x80\x59\x9a\x67\x30\xc6\x4d\x6b\x48\xf0\x28\xe1\x42\x95\x92\xfb\x88\x20\x15\x35\x11\x1e\x92\x10\xdb\xd9\xcd\x1b\x40\x18\x51\x04\x52\x5b\x86\x90\xe2\x66\xdd\xda\x46\x15\x03\xa3\x3f\x5f\xd7\x00\xe2\x61\x1a\x63\xa6\xc0\x9e\x23\x30\x44\x93\x5a\x2f\x65\xa1\x22\x9c\xd5\xf6\xc0\x23\x28\x20\xf5\x67\xee\x66\x08\x05\xc8\x6d\xd2\xce\x4c\x02\x9a\x80\xe1\x11\xf8\x54\x5e\xab\xed\x9d\x2d\x62\x8d\xe4\xb5\xe0\x8a\x87\x9c\x82\x26\x18\x11\x86\xf8\xc8\xa1\x3c\x84\x9a\xa9\x93\x14\x8f\x9a\x4d\xb0\x6b\x9c\x78\x17\xfc\x2f\xd8\x1d\x49\xed\xcd\xbb\xe0\x54\x5f\xea\xab\xb3\x25\x84\x73\x19\xee\x70\x70\xc3\xc3\x7b\xac\x6a\x25\x5e\xaf\x41\xc1\x24\xe2\x52\x81\xd7\xc0\x72\x71\x18\x71\x6b\x6f\x9e\xd2\x48\x3a\x9c\xc5\x58\x4a\xd8\xc7\xa0\x09\x8a\xed\xc0\x43\xcc\xd4\xc2\x9e\x4c\x79\xcf\x10\xae\x6e\x3e\x7e\x70\x12\x5d\x7b\x0d\x8a\x83\xa0\x82\x15\x1e\xfa\x33\xb7\x79\x0e\xe5\xfd\x16\x53\x62\x52\xcb\x09\x55\x10\x9e\xce\x76\x76\x16\xf4\xcd\x5d\x39\x57\xba\x63\xee\x6a\x73\x5a\x56\xc8\x4c\x9d\xbf\x47\x28\x6d\x13\xa9\x6a\x7b\x67\xf3\x64\x5d\x37\x27\x03\x14\x0c\x28\x06\xb9\xb3\xcf\xc1\x50\xac\xa6\x84\xba\x1a\xe8\x3d\x02\x4d\x60\x7d\x95\x2f\xd9\x19\xa2\x75\xb6\x80\xa2\xbb\x19\x2c\x5a\x14\x6b\x6f\xd3\x96\xfa\xd3\x9f\x2b\xcc\xbf\xae\xcd\x93\xdd\x73\x7a\x84\xa1\xda\xae\x8a\x76\xf7\x1c\xc5\xcf\x85\x80\x93\xda\x9e\xd3\xe3\xa2\x05\xc3\x68\xe6\xa8\xd8\xd0\xdc\x03\x8f\x0b\x9b\xac\x19\x13\x2d\x5f\x0e\xe3\x10\xb4\x68\x09\xd2\x03\x35\x82\xc0\x2b\xed\x71\xbb\x7b\xe0\x71\x5e\x52\x27\x49\x65\x54\xdb\xfd\x6a\x17\xbc\x06\x04\xed\x3d\xcd\xdb\xa5\xba\x81\x9a\x61\xbe\x69\xb9\x5d\x6e\xcc\x5d\x45\xb7\x37\x95\xfd\xa8\x18\x6a\x9a\x64\x74\x78\xd7\x4a\xcf\x0a\x7e\xb3\xcc\xe3\xb9\xa6\x4f\xdc\xf1\x02\x8e\x26\x79\xfc\x47\x22\xbf\x40\x64\x08\x42\x0a\xa5\x6c\x5a\x3a\x8f\x40\xc2\xb0\xb0\x7b\x34\x25\xa8\x94\x06\xbc\xa8\x5e\xed\x08\xa3\x7a\xe9\x71\x89\x88\xe0\xa3\x12\x62\xf6\x14\x4a\x82\xf0\x8c\x09\xb5\x63\x64\x1f\x80\x04\x22\x5b\x90\x7e\xa4\xec\xfd\x0a\x42\x86\x94\xd2\x29\x06\x83\x43\xc0\xe0\xd0\x4e\x08\xa5\x32\xbb\x92\x0a\x86\xf7\x18\x2d\x41\x03\x26\xcb\x4e\x51\x61\xa8\xc8\x10\x5b\xbe\x07\x81\x0e\x32\x5b\xf1\x7e\x9f\xea\x84\x07\x83\x69\x9a\xfd\x2a\x0f\x33\xcb\xcf\x13\x93\xf4\x5c\xb8\x82\xb0\x4b\xc9\x4a\x96\x6b\x79\xe4\xa6\xb5\xfc\x3c\x78\x5e\xce\xc2\x10\x1c\xc0\x21\x34\xe6\x3d\x3d\xb3\xe6\xd9\xc6\x1c\x41\x3a\x5d\x83\xa2\x8f\x95\xe6\x3f\xe9\x64\xcb\xfe\x79\xc0\x53\xf5\x4c\xee\x9e\x9b\xd2\x8a\x45\xdd\xcc\xa4\x95\xc5\x39\x4f\xca\x8c\x5c\xd7\x96\x5d\xa4\x57\x02\x54\x30\xb0\xf3\xf2\xb5\xca\x9a\xaf\x6c\xbb\x0b\x03\x50\xb7\xed\x15\x00\x9a\x1c\x41\x4d\x6b\x6a\xc8\x32\xed\x04\x32\x0c\x7a\x10\x61\x40\x18\x98\x7a\xc3\x52\x32\x53\x5e\x84\x05\x7c\x0c\x14\xe7\x34\x80\x62\x15\xcf\xaa\x1a\x8b\x4e\xbf\x0e\x5a\xef\xce\x58\xda\xf5\x83\x0d\x38\x55\xbc\x24\xa5\xd4\x84\xcc\x16\x88\x19\x32\x9c\xa2\x06\x8a\x81\x40\x31\x1b\xe1\x1e\x4c\xa9\xca\xae\xc7\xd2\x02\xd9\xb9\xae\x69\x7d\xc2\x3d\x81\x65\x64\x65\xfb\x18\x28\x76\x41\x31\x14\x6d\xde\xdf\x82\x4f\x06\xba\x5e\xf5\x55\xee\x36\x0f\x84\xc8\x70\xc3\x16\x6e\x01\xf2\xca\xb6\x5d\xbd\xbd\x6b\xec\x06\x56\x59\x03\xc8\x04\x86\x58\x1c\xe9\xae\x6d\x2d\xa7\x4d\x8f\xb5\x10\x82\x8f\xd6\x3a\x8f\x86\x79\x91\xab\x69\x8f\xa6\x20\xfb\x2e\xac\x39\xeb\x14\xac\x72\x24\xb4\x79\xff\x5a\xc3\x6d\x72\xcd\x57\xb6\x3d\xed\x4e\x28\x91\x6a\xab\xcd\x2b\xb1\xb1\x75\xae\x29\x47\x1d\xc5\xb6\xc0\x32\xe1\x4c\xae\x0f\xb7\x82\xa0\x69\x25\xca\x24\x4d\x8f\x30\x47\xd3\x34\x1c\x76\xc4\x87\x58\x37\x96\x99\xc3\xda\x45\xc9\x02\x59\x70\x20\x3e\x62\x16\xc8\xba\xe6\xa9\x24\x14\x4e\x78\xaa\x4e\x41\x8f\x8c\x71\x5e\xda\x47\x5c\x20\x7b\x24\x60\x72\x0a\x02\x81\xe1\xbd\xad\x17\xce\xb6\x8d\x29\xa5\xab\xe9\x54\xb2\x88\x20\x84\xb3\x50\xda\x0e\xdb\x50\x10\xdb\x03\xe7\x2c\xcb\xce\x2a\x63\xbb\x61\xf9\x9e\x54\x82\xb3\xbe\x7f\x09\x15\xf6\xdc\xfc\xe6\x79\x64\x5d\x15\x7d\xa9\x20\xf5\x99\x20\x17\x1c\xfd\x3e\x04\xe9\x60\x15\x71\xf4\x9f\x13\xe5\x78\x26\xca\x35\x54\xd1\x6f\x21\x88\xe7\x6e\xeb\x53\x9a\x6e\x31\x35\xdc\x08\xac\x66\x1d\xe3\x16\x84\xb7\x03\xf6\xdc\x2c\x30\xbf\x28\xd7\x6f\x93\x81\x4d\x76\x35\x89\x72\xcb\xdc\x3a\xc2\x94\x02\xfd\x65\x4b\xa0\xcf\xc4\xd3\x62\xeb\xe1\xd8\x6f\x43\x59\x24\x5a\x90\x26\x08\x2a\x8c\x4e\x81\x27\x13\xc8\xb2\xcc\x45\xa1\x54\x9f\xf3\x75\xcb\xbf\x83\x24\x6b\xcf\x1d\xc7\x73\x35\x88\xef\xb9\x38\x7e\x91\x3e\xeb\x1e\xbd\xb2\x6d\xa0\x7b\xa3\x83\x8d\xbd\xd1\xb4\x01\x5d\xda\x1b\x6d\xe8\x88\x36\xd6\xb2\x0a\x93\x2f\xaa\x07\xa5\x5a\x30\x7f\x5e\x5c\x56\x0b\xa4\x12\x24\xc1\xa8\x5c\x19\xb6\x2a\x37\xcf\x89\x80\x67\xe4\x6a\x9d\x0a\xcc\x08\xc6\xaa\x1f\x7d\x63\x4a\x71\x2a\x88\xe5\xff\x28\x88\x71\x94\x5c\x87\xfc\xf0\xf6\x99\x84\x9c\x49\x9d\x2d\x4a\xcf\xfa\x74\x92\x44\xfa\x01\x28\xae\xec\x30\xc2\x43\xc1\x99\x29\x6f\xfe\xd4\xa3\xb6\xc5\x4a\x93\x19\x4e\xe5\xe7\x59\x39\x66\xa6\xdd\xc1\xdb\x5c\x3b\x84\x65\x68\xf9\x97\xd8\x9c\x45\x08\x67\xbf\x33\x35\x5f\xa8\x5f\x7d\x3f\xd7\x2f\xce\x4a\x89\x35\x2d\x29\x7f\x08\xe5\x0a\xe3\x25\x50\x45\x96\x29\x51\x7f\x08\xc5\x4e\xbe\xb1\x40\xde\x44\xca\x94\x66\x07\x7a\xfd\xfb\xdf\xad\xdc\xaf\x8a\xb1\x5d\xdb\x50\xb4\x0c\x5b\x64\xd6\xed\x5a\x80\x6d\x58\x6e\xa6\xb4\xb1\x91\x78\x41\x55\x5d\xb2\x5e\x59\x2a\xdd\xe6\x97\xe6\x5a\xd7\xe2\x6c\xc8\x02\xb2\x19\xcb\xb4\x60\x96\xdb\x8b\x6c\x34\x63\x2a\xae\xc9\x2e\xf9\x54\x06\x08\xae\xcf\x2b\xe6\x1d\x8a\xb5\x7c\xac\x96\xe1\xe6\xaf\x59\x80\xb9\x91\x71\x75\xd0\xb6\x00\xbf\x7a\xb4\xb2\x08\x6b\xc6\x8d\xab\xa6\x30\x41\xaa\x14\x67\xf9\xfb\x02\x73\x53\x94\xe4\x90\x72\x89\xf3\x91\x13\x22\x32\x26\x05\x51\xcb\xff\x1f\x45\x62\x2c\xcf\x3c\xd7\xe0\xac\xa0\x1e\x1d\xce\xcb\x92\x8d\x25\xac\x7c\xe4\x78\x87\x83\xd2\xd8\xf1\x70\xd9\x80\x6a\xa9\x31\x17\x55\xd4\x3e\xb5\x4a\xc1\xc4\xef\xc4\x9a\x1b\x91\x00\x32\xf0\xae\xdb\xbd\x06\xfa\x9e\xb0\x3e\x80\x49\x42\x89\x99\x99\x83\x1e\x17\x40\x61\x99\x8d\x6f\x21\x43\xa0\xa7\xdb\xc2\xec\x9d\x84\x9a\x98\xc9\x5e\xb2\x92\x41\x9b\x84\x98\x85\x18\x5c\xf0\x64\x92\x75\x96\xe0\x1f\x7f\xfb\xd7\x5f\xfe\x0a\x0e\xf6\xeb\xc7\x6f\xc0\x15\x17\x88\x80\x0e\x14\xea\x9f\x7f\x67\xa0\x56\xcc\xfa\x22\xa5\x92\x53\xd7\x1d\xe8\xc7\x0e\xe1\x16\x98\x4e\xf5\x3e\x07\x14\xb2\x7b\xcb\xaf\x00\x78\x2e\xf4\xf7\xd6\x89\xf1\x09\x53\x0c\x25\x46\x20\x65\x08\x0b\xd0\x79\xdf\x05\x54\x4b\x26\xf1\x1b\x20\x31\x06\x73\x9c\xe5\xa9\xeb\xf6\x89\x8a\xd2\xc0\xbc\xd7\x8f\xa1\x4e\xa0\x27\x07\x6e\xac\x77\xc7\x0d\x28\x0f\xdc\x18\x4a\x85\x85\xdb\x7e\x7f\xd1\xfa\x70\xd3\x72\x62\xb4\x28\x63\xfe\x6c\xcd\xe4\x31\xf1\xb3\xbd\x45\x58\x41\x42\xa5\xa3\x2d\xae\x04\x09\x52\x85\x4f\x9f\x23\xd1\xf2\xed\x59\x8f\xb3\x54\xaa\xad\x9d\xaa\xc7\xb9\x7a\x59\xdc\x54\x86\x72\x2b\x22\xe8\x42\x47\xd7\xea\xf8\xf9\xb2\x8c\xd5\x6d\x75\xae\xdb\xe7\xdd\xd6\x0d\x28\x12\x56\xfe\x66\xb0\x3c\x8d\xc1\x4c\x89\xc9\xdc\x7b\xca\xb1\x3d\xfb\xdb\x85\xad\x70\x9c\x50\xa8\x70\x79\xd8\xeb\x29\x31\x4b\x0f\x24\xbc\xcf\xba\xf6\xc7\xc7\xe9\x4b\xda\x90\x53\x2e\x9e\x9e\x2a\x73\xea\x62\xa0\x64\x9c\xc0\x9e\xc1\xb3\x34\x7e\x7a\xaa\xa6\x3b\x85\x16\xc7\x22\x33\x14\x7d\x12\x7b\x7a\xf2\x5c\x85\x36\xa1\xd5\x2d\xdf\x0b\xfc\xb2\x70\x28\xc3\x0c\xaa\xdb\xba\x15\xa9\x19\x1d\xd3\x3d\x6e\x27\xc3\x71\x19\x51\x77\x66\x55\xb4\xf9\x5a\xad\xb7\x77\x6e\x04\xb7\x62\xc7\x8a\xbe\x07\x53\x1c\x2a\x5b\xf0\x51\x31\x1e\x43\x44\x26\x14\x4e\x4e\x01\xe3\xac\x7a\x2a\xcb\xe4\xe3\x54\xb7\x2c\x4d\xeb\x70\x55\xf9\xc8\x09\x25\x10\x21\xc2\xfa\xa7\x60\x3f\x19\x83\xfa\x51\xf9\x6b\x55\x54\x44\x0d\xff\xd2\x48\xec\xb9\x51\x63\x35\xd0\x5c\xa7\x45\x61\x80\x29\xc8\xbe\xed\x44\x90\x18\x8a\x49\xf1\xd2\xa6\x68\xaf\x56\x52\x5b\x3f\x30\x48\x44\xf1\x82\x2a\x11\xd8\x96\xa1\xe0\x94\x66\xe7\x4e\xdf\xd3\xee\xe0\x3f\x16\x3b\xfb\xa4\x2d\x93\xad\x79\x6e\x22\x56\xb4\x1f\xeb\x5a\x8f\x75\x7a\x2d\x46\x88\xee\x62\xf5\xd1\x19\xff\x36\x2a\x1a\x5e\xbf\xaa\x8e\xd3\x74\xe7\xb7\x79\xff\x37\x51\x8a\xf2\xfe\x73\x14\xda\x2e\xc5\x2e\x8f\xcd\xd2\x1f\x5f\x16\xb2\xe9\x74\x9e\xb1\x7d\x36\x2d\xe8\x57\xdb\x74\x4f\x21\x93\xb2\x7e\xfc\xf4\xde\x64\xaa\xa5\x39\xc6\x7f\x7c\x44\xb3\xb3\xf9\x8a\x44\xb4\xce\x19\x4d\x06\x9b\xf9\x62\xe1\x9e\xce\x2c\xb7\x4d\xcf\x35\x4b\xd9\x17\xf0\x26\xa5\x3d\x5f\x02\xa9\xa0\x4a\xe5\xbc\x04\xc6\x47\x1d\xf3\xe8\x22\x4f\xd6\x6b\xc4\xd8\xf4\x1a\x2c\xfb\x19\x12\x3c\xb2\x8b\x31\x59\x56\x96\x52\x41\x9a\x56\xbe\xc7\x96\x7f\x4b\xf0\x48\x77\x0a\x5f\xc0\x03\x23\xa2\xd6\xf2\x68\x21\xa2\x16\x79\x2c\x77\x2f\xcf\x35\x47\xa4\x1d\xcf\x35\xff\x07\xfd\x77\x00\x00\x00\xff\xff\x6c\x6f\x7a\x17\x27\x2a\x00\x00")

func tmplIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplIndexHtml,
		"tmpl/index.html",
	)
}

func tmplIndexHtml() (*asset, error) {
	bytes, err := tmplIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/index.html", size: 10791, mode: os.FileMode(420), modTime: time.Unix(1519223099, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsMappingJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x6f\xe3\x36\x10\xbd\xe7\x57\x0c\xd4\x00\x94\x60\x5b\x0e\x50\x14\x05\xe4\x28\x87\x4d\x0b\x2c\x8a\x6e\x5b\x20\xdb\x53\x1a\x2c\x68\x69\x62\x71\x97\x22\x59\x72\x14\x5b\x28\xfc\xdf\x0b\xea\xc3\x2b\x5b\xb2\x37\xc0\x86\x27\x53\x33\xf3\xe6\xe3\xbd\xa1\x9f\x2b\x95\x91\xd0\x0a\x3e\x70\x63\x84\xda\x84\xb9\x2e\xb9\x50\x11\xfc\x77\x75\x05\x00\x40\x85\x70\x71\xfb\x0d\x52\x68\x7f\x34\x86\xe5\x32\xd3\xa5\x11\x12\x81\x0a\x04\xc2\xd2\x48\x4e\xd8\x98\x5e\xb8\x85\xb2\x85\xfb\xe4\x74\x65\x33\x84\x14\xae\xc3\xe0\x87\xee\xe3\x02\x15\xd9\x3a\x88\xe2\x82\x4a\x19\x46\xab\x51\x50\x8f\x06\x29\xbc\xe7\x2a\x97\xb8\xe6\xd6\xc5\x5d\xbe\xf0\x18\x3a\x5a\x5d\x1d\xe2\x9f\x85\x94\xbf\x0b\x47\x90\x42\xdf\x57\xe8\x3b\x81\xee\x5c\xc7\x1b\xa4\xdf\x1e\xfe\xfc\x23\xdc\x0a\x95\xeb\x6d\x2c\x75\xc6\xbd\x57\x6c\xac\x26\x9d\x69\x09\x33\x08\x96\xcb\x00\x66\x5d\xab\xcd\x9d\x1b\xb1\xec\x92\x06\xf3\xaf\xc8\x39\x27\x3e\x44\xef\xab\x10\x84\xa5\x83\x14\x1e\x9f\x56\x47\xb6\xe1\x04\x88\xaf\xfd\xe4\xd6\x3a\xf7\x73\xc0\xd2\x50\xdd\x0f\xe2\x6b\xb1\xc8\xb3\xa2\xc9\x32\x48\xfa\x05\xeb\x79\x3f\xa7\xd3\xe4\xfe\x74\xa6\x47\xe6\x88\x53\xe5\x3e\x65\x5a\x6a\xcb\x9e\x20\x05\xd8\x20\xdd\xfb\xdb\xbb\xfa\xa1\xb1\xf5\x83\x8c\x2d\x3a\xa3\x95\xc3\xb8\x8d\xb9\xd7\x39\x9e\xd4\x72\x84\x5c\x22\x15\x3a\x9f\x46\xfe\xd0\xd8\x06\xc8\xff\x56\xe8\x28\x6e\x43\x26\x40\xfd\xbc\xbc\x0a\x20\x1d\x91\xdf\x83\x4c\x44\x9d\x9f\x24\x37\x06\x55\x1e\x7a\xc8\x93\xb8\xfd\xe0\xbe\xef\x45\xb3\x5f\x0d\x54\x3e\x54\x4f\xf7\xb3\x33\x5f\x87\xec\x90\xcf\xf3\xc1\xa2\x58\xab\x90\x65\x52\x64\x5f\xd8\x1c\x58\xbc\x26\xb5\x78\x11\xb8\x5d\x74\x5e\x6c\x3e\x2d\x40\xdf\x6e\x65\x45\xb3\x0e\x3e\x67\x14\x7b\xb8\x30\xa8\xac\x08\x06\xf5\x7d\xbf\x4e\x1b\x53\x65\xc5\x2b\xe4\x9a\x69\x45\xa8\x7c\xd7\x3e\x63\xec\xc8\x0a\xb5\x11\xcf\x75\xa7\x3c\x55\x49\x39\x87\xe0\x1f\x0a\xa2\xa3\xd0\x77\x5a\x93\x23\xcb\xcd\x2f\x82\x4b\xbd\x89\x5d\xa1\xb7\xe1\x58\x8f\x24\x48\x62\x02\xac\x7b\x5e\x20\xc7\x67\xa1\x84\x2f\x88\xcd\xc7\x12\x43\xe7\xf8\xc6\xbb\xdf\x1a\x8b\x77\x0c\x66\xe0\x6a\x45\x7c\xf7\x5e\x6c\x0a\x29\x36\x05\x85\x5d\xb5\x11\xcc\x80\xdd\x2e\x1b\xaf\x8b\x34\x1f\xe8\x7e\x3d\x91\x98\x0b\x7a\x23\x22\xbd\x1b\xaa\xdc\x68\xd1\x0c\xf8\x7b\xd9\x9c\x52\x48\x0f\xff\x46\x4c\x8f\x89\xf6\x91\xd7\x84\x3b\x6a\x9a\x65\xb7\xfe\x27\xb7\xc8\x41\xe4\x69\xe0\x2f\x8b\xca\xe4\x9c\xb0\x1f\x59\x00\x8e\x6a\x89\x69\x50\x0a\xb5\xd8\x8a\x9c\x8a\xe4\xa7\x9f\x6f\xcc\x6e\xe5\xef\x05\x7a\x16\x93\x1f\x6f\x6e\xcc\x2e\xb8\xbb\x5d\xf6\x60\x77\xec\xf4\xf5\xf3\x96\xf8\x85\xcb\x03\xe3\x47\xe6\xb7\x90\xa2\xe7\xf9\x92\x08\x9b\x12\xc6\xf6\x75\x45\xa4\x95\x4b\xe0\x71\x9c\xc2\x1f\xc9\xd7\x28\x13\x60\xf7\x5c\x65\x28\x27\x12\xf8\xc3\x1b\xa6\x92\x01\x67\x4d\xe9\x53\x2f\x7a\x7f\x5a\x8f\x38\x93\xda\xe1\xe9\x7f\x45\x7f\xf6\xa3\xaf\xfb\xf9\x19\xc8\xbe\xcc\x07\xfe\x82\x67\x8a\xcc\x9c\xbb\x97\xdc\xb9\x04\x98\xdf\x0b\x63\x45\xc9\x6d\xfd\x66\x1d\x1d\x4b\xd2\xaf\xe6\x84\x9a\x58\xd4\x88\xe0\x4c\xbf\xcd\x26\xf0\xcf\x7c\x37\x41\xf7\xf0\x50\x6d\x3c\xf3\x7f\xfd\xfd\xf1\x4c\xf5\xfd\xa9\xac\x4c\x0e\x1b\x7b\xd9\xd5\x2f\x4d\xd2\x37\x70\xd9\xd5\x55\x59\x86\x7e\x8e\xa7\x2b\x3a\x62\xf5\xf0\xcf\x13\x46\xb0\xbf\x0c\x8a\xd6\x6a\x3b\x01\xc9\x25\x5a\x0a\x83\x5f\x5b\x73\x30\x9b\xd8\xf5\x28\x5a\x8d\x12\x7f\x23\x5b\xd7\xe7\xc7\x66\x90\x01\x37\x46\x8a\xf6\x11\x5b\x7e\x76\x5a\x05\xdf\x1e\x55\x1b\xc9\xbc\x37\x3b\xeb\xbc\x7f\xbd\xae\x9f\x5e\xf5\xec\xef\xaf\xfe\x0f\x00\x00\xff\xff\x61\xde\x25\x81\xe1\x0a\x00\x00")

func tmplJsMappingJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsMappingJs,
		"tmpl/js/mapping.js",
	)
}

func tmplJsMappingJs() (*asset, error) {
	bytes, err := tmplJsMappingJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/mapping.js", size: 2785, mode: os.FileMode(420), modTime: time.Unix(1519223695, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsRequest_loggerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xcc\x6a\x17\x30\x59\xbb\x4a\xb6\xe8\x49\x86\x80\xa2\x69\x8a\xa0\x48\xb6\xe8\x26\x45\xd1\x53\x40\x4b\x13\x59\x58\x8a\x54\xc9\x61\xbb\xc6\xc2\xff\xbd\x20\xa9\x6f\x38\x1f\x3a\xe4\x30\x7e\xf3\xf8\xf8\xe6\x85\xf3\xe4\x54\x41\xb5\x56\xf0\x19\xff\x71\x68\xe9\x56\x57\x15\x1a\xc6\xe1\xdb\x0a\x00\xe0\xe2\xa2\xd0\x4d\x5b\x4b\x04\x3a\x20\x10\x36\xad\x14\x84\xe1\xa7\x7f\x85\x01\xab\x9d\x29\x10\x72\xf8\xc0\x92\xf7\x26\x32\x7c\x8f\x8a\xcc\x31\xe1\xe9\x81\x1a\xc9\xf8\x6e\x00\xf7\xdd\x90\xc3\x8d\x50\xa5\xc4\xbd\x30\x36\xed\xf8\x59\xa4\xe2\xbb\x55\xc0\xd3\xa1\xb6\xa9\x72\x0d\xe4\x70\x19\x2b\x83\x50\xd7\x96\x82\xf0\xa1\x26\x89\x83\x4c\xff\x7d\x60\xa5\x2e\x5c\x83\x8a\x78\x2a\x88\x0c\x4b\xc8\x63\x92\x2d\x24\x9f\xae\xff\x82\xcf\xd7\x7f\xfc\x79\x7d\xff\xf0\xee\x5d\xd2\x49\xf2\x9f\x45\x7a\xa8\x1b\xd4\x8e\xd8\xc0\xef\x49\x5f\x22\xbb\xbb\xd3\xc5\x17\xb8\xd2\xca\x6a\x89\x09\xdf\xc1\x69\x0b\x3f\x5c\x5e\x5e\x76\xb4\xa7\x85\xdc\x0a\xa9\xb3\xd6\x9f\xc4\xa8\x6e\xd0\x92\x68\xda\xa9\x74\xef\x8e\x19\x41\x90\x83\xc2\xff\xe0\x17\x41\x13\xfc\x77\x1f\xc7\x33\xfa\x1e\x6f\x04\xc5\x86\x49\x7b\x5a\x21\x85\x5e\x0e\x1b\x48\x2e\x12\xd8\x0c\x4d\xfe\x63\x0b\xe8\x9d\x56\x74\x08\xd8\x8f\xe7\x1b\x16\xf8\x5f\x9d\x94\x7f\xa3\x30\x91\x1e\x7e\x82\x57\xf0\x37\xda\x19\x1b\xc1\xd9\x2b\xd0\xbb\x5a\x39\xc2\x37\x82\xef\xb1\xd0\xaa\xb4\x6c\xe2\x89\x41\x72\x46\x0d\xb6\x9c\x1f\x48\xcc\xcf\xad\xb0\xfd\x5c\xe6\x3e\xcf\x23\xb5\x7e\x2f\x85\xa5\xc7\xd8\x53\xae\x79\x4a\xf8\x95\xd8\xb3\x33\x7d\x3e\x03\x57\x5a\x85\x56\xe5\x9a\xad\x17\x28\x96\xf3\xb7\x24\xc8\x59\xc8\xc3\x8f\xa9\x41\xdb\x6a\x65\x31\x8d\xe5\x2b\x5d\xe2\x7c\xf4\x0d\xd2\x41\x97\x23\x3c\xa8\x49\x63\x75\x8e\x6c\x05\x1d\x96\x38\x5f\xdb\x9d\x8b\x1f\xe4\x60\x8f\x8a\xc4\xd7\x9b\xba\x3a\xc8\xba\x3a\x10\xfb\xed\xfe\xf7\x4f\xa9\x25\x53\xab\xaa\x7e\x3a\xb2\x29\xcf\x16\x9c\x2a\xf1\xa9\x56\x58\x6e\xe1\x47\xce\x97\x94\xf1\x0e\x6f\xe6\x8c\xf0\x97\x49\xa5\xae\xde\xce\xe7\xe4\x2b\x12\x0b\x2d\xb5\x81\x3c\x0e\x48\x6a\xf3\xf3\xf1\x3e\xf8\xcd\xa2\xed\x7c\xb5\x8c\xd6\xb7\xde\xa8\x47\xe5\x9a\x0c\xc2\x34\xbb\x4a\x06\x83\x2b\xfd\x55\x32\x18\x2f\x65\xa4\xae\x32\x2f\x7f\x68\x78\xf4\x99\xca\x96\x0f\x44\xd0\xee\x13\xc5\x47\x60\xa1\x4b\xcc\xba\x84\x8c\xd5\x38\xec\xac\x8b\xc2\x58\xf7\xc3\xcd\xc2\xd8\xa7\x0c\x52\x9b\xac\xbb\xef\x69\x96\xd2\xf0\xd4\x4a\x5d\x5d\xfb\x67\x1b\xf2\x31\xb5\x6c\x19\xd3\xfe\x51\xde\x6c\x96\x2e\x86\x6c\xf7\x3e\xc6\xa0\xf7\xe0\x2e\xed\xf3\x0e\xbf\x1a\x20\x1f\x36\x02\xeb\x18\x26\xa8\xe9\x42\x21\xb1\xf7\x0b\x68\xaf\x4b\xbf\x56\x5a\x83\x2d\xaa\x92\x79\x8e\x49\xc3\x6c\x2d\x2c\xcb\xcb\xff\xf6\xd1\xe4\x6e\xdf\x9c\x76\xab\xd3\x6a\x35\x7d\xf5\x0d\x8a\xf2\x38\xdf\x0a\xab\x5e\xd8\x9e\xd4\x95\x44\x61\x6e\x75\x95\xf0\xb4\x90\x75\xf1\xe5\x0c\xf2\xe5\x6b\x60\xd3\xd2\xb1\x57\x7a\xea\x75\xf8\x17\xa7\xc7\x7b\x91\x6b\x9e\x6a\xc5\xd6\xe1\x88\xf5\x16\xd6\x1d\x87\x59\x6f\xe1\xfc\x81\xde\xdd\xba\x0c\x1b\xd9\x4f\x80\xa7\x9e\x85\x25\x24\x4c\x85\x94\x2c\x0c\x4e\x60\x03\x75\xc9\x53\xd2\x55\x35\xda\x16\xc4\xac\xfc\xdf\xff\x03\x00\x00\xff\xff\x02\x4a\x2c\xba\x1d\x08\x00\x00")

func tmplJsRequest_loggerJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsRequest_loggerJs,
		"tmpl/js/request_logger.js",
	)
}

func tmplJsRequest_loggerJs() (*asset, error) {
	bytes, err := tmplJsRequest_loggerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/request_logger.js", size: 2077, mode: os.FileMode(420), modTime: time.Unix(1516991133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsSortingJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x41\x6f\xdb\x3c\x0c\xbd\xf7\x57\xf0\x10\xc0\xd2\xf7\x39\x42\x7b\x5d\xd7\x01\x69\xd7\xc3\x80\xa1\x97\xac\xbb\x2b\x32\x13\x6b\x95\x25\x4f\x62\xdc\x16\x45\xff\xfb\x20\xc7\x89\x65\x3b\x29\x72\x18\xa6\x4b\x90\x88\x7c\x8f\xef\x91\x54\xd6\x5b\xab\x48\x3b\x0b\x4b\xe7\x49\xdb\x0d\x23\xb9\x32\xb8\x44\x83\x8a\x9c\xcf\xa1\xfd\x7a\x6f\xb0\x42\x4b\x81\xc3\xdb\xc5\x05\x00\x80\x41\x82\xd0\xc5\x3c\xd6\x70\x03\x99\xd8\x98\xd7\xba\xd4\xca\xd9\xb9\x2a\xb1\xf1\xce\xce\xb7\x75\x76\x3d\x89\xfe\xea\x9e\xed\x89\xf8\xc2\x3d\xdb\x49\x06\x16\x77\x46\x86\xf0\x20\x2b\x8c\x69\xfb\x1f\xb3\xeb\x5d\x21\x87\xf2\x83\xf3\xf4\x23\xd6\xca\x9c\x2f\xd0\xe7\xa0\x09\xab\x70\xe7\xb6\x96\xf8\x5b\x1b\xba\xc7\xf5\xee\x39\xc0\x0d\xcc\x86\x4a\xe1\x7f\xc8\x80\x56\xae\x78\x05\x20\x9f\x71\xb1\x41\x62\xbc\x63\x89\x27\xa6\x89\xc8\xc2\xf6\x9c\x4c\xe6\xb0\x3a\x58\x92\x52\x2c\xe0\x06\x36\x48\x3f\xa5\x61\x92\x5f\x4f\xae\x6f\xfb\xeb\x55\x4a\x11\x8f\x5e\xb3\x05\x7c\x86\xdb\x88\x0b\xa3\xe3\x91\xb6\xde\xc2\xfc\xea\xbf\x56\xe3\x10\xf8\x7d\x0a\xf3\xe5\x43\x98\x33\x50\xba\xc8\xcb\x3e\xe6\x3d\xad\xf7\x60\x7e\x27\x06\x4d\xc5\xdf\x26\x6a\x9b\xd6\xec\x78\x27\x54\xa9\x4d\xe1\xd1\xb2\x8c\x8a\x8c\x0b\xfc\xcd\x92\x2e\x09\xc2\x17\x62\x5c\x90\x7b\xac\x6b\xf4\x77\x32\x20\x1b\x99\xa7\xd7\x6c\x26\x74\x78\xd8\x56\xe8\xb5\x62\x0d\xe7\x53\x75\x91\xae\x96\x3e\xe0\x37\x4b\xac\xc9\xaf\x2e\xf9\x19\x0a\x9b\x44\x61\xaf\x6f\x26\x50\xaa\x92\xc5\xce\xe7\x07\xb1\x4c\xdb\x02\x5f\xf2\x38\x0f\x63\x73\x47\x23\x35\xd0\x1b\x27\x2b\xe3\x42\xd6\x35\xda\x22\x42\xf2\xa1\xa9\x09\xf5\xc1\x56\x49\x24\x55\xf9\x5d\x07\x42\x8b\xbe\xf5\x2a\x87\x8e\x1e\xf7\x2b\xd9\x97\x30\x6b\x23\xb8\x58\x6b\x5b\xb0\x7e\x3b\xb9\x70\x96\x65\xca\x68\xf5\x94\x25\x32\x8e\x74\xaa\x6f\x46\xdb\xb2\x1d\x5a\xed\xb1\x59\x18\xc3\xb8\x30\x68\x37\x54\x0e\xdd\xec\xd7\x6e\x7e\x05\x79\xd2\xcd\x61\x58\x25\xfd\xd3\x22\x2c\xbb\xe5\x65\x33\x46\xa5\x0e\xfc\x88\x05\x27\x75\xc4\x77\xe3\x1f\x29\xf9\x2b\x42\xc6\xbd\x1c\x25\xba\xd5\x2f\x54\x94\x0e\x90\x32\x28\xfd\x32\x7d\xf3\xd2\xe9\xdf\xc5\x0b\x59\x74\x57\x93\xc7\xf1\x04\xed\x31\x54\x48\x47\x26\x9d\x13\x8f\x95\x6b\xf0\x63\x82\x61\xd6\xae\x2b\xe7\xe4\x8d\x0b\xd3\x56\xd3\xa0\x94\xc1\x1f\x8d\x58\x3b\x7f\x1f\x77\x6f\xb8\x02\x43\xac\x6e\x77\x7b\x88\x88\x09\x9f\xda\x8f\x2e\xee\xfd\x4f\x00\x00\x00\xff\xff\x8f\xda\xb9\x25\xdf\x06\x00\x00")

func tmplJsSortingJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsSortingJs,
		"tmpl/js/sorting.js",
	)
}

func tmplJsSortingJs() (*asset, error) {
	bytes, err := tmplJsSortingJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/sorting.js", size: 1759, mode: os.FileMode(420), modTime: time.Unix(1516991133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsUtilJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x51\xaf\x93\x30\x14\xc7\xdf\xf9\x14\x27\xc4\x5c\xca\x26\x63\x9a\xbd\xb8\x8d\x11\xbd\x12\x7d\xd0\x68\x72\xaf\x2f\x52\x96\x74\xd0\x31\xb4\x2b\x4b\x5b\xd4\x79\xd9\x77\x37\x65\x64\x63\x5c\x70\x3e\x78\xfb\x40\xd3\xff\x39\xfd\x9d\x3f\xe7\x14\x00\x60\x5d\xf0\x58\x65\x39\x87\x94\xaa\xdb\x9c\xe5\xe2\xcd\xfe\x4e\x11\x55\x48\x24\xab\xed\x36\x4f\xa8\x0d\x0f\x06\xd4\x2b\x5b\x43\x23\x02\x9e\xe7\xc1\xcb\xf1\x18\xca\x12\x1e\xa9\x2f\x9a\xf7\xf4\x12\x54\x15\x82\x83\x29\x8b\x38\xa6\x52\x9a\xb3\x53\xf4\x00\x94\x49\xda\x05\x9f\x8c\x27\x7d\x98\x84\xf0\x94\x8a\xc7\x94\xee\xec\x9f\x44\xf0\x8c\xa7\xcd\x74\xe3\xf8\x35\x7a\xfa\xf0\x91\xaa\x4d\x9e\xa0\x6d\xb5\xb5\x7b\x70\x54\x2b\x8b\xd6\xbb\xe0\xde\xea\x33\x99\xf1\x75\xde\xfd\xa3\x4d\xc2\xe7\x4f\x77\xfd\x88\xbf\xb6\xeb\x82\xf2\xe5\x3f\x40\xde\x06\x1f\x82\xfb\xa0\x97\xf3\x34\x4d\x97\x7b\xae\xc8\xaf\xf7\x59\xba\x61\x59\xba\x51\xe8\x9b\xcc\x79\xd3\x81\x3e\x83\x57\x6d\x23\x41\x77\x8c\xc4\x14\xb9\x37\x6e\xfa\x1c\xac\x1b\xb2\xdd\xcd\x2c\xfb\x2c\xcf\x8f\x32\x53\x17\xea\xe2\xa8\xa6\x5a\x3d\xbb\xa9\x7d\x5e\x72\x91\x89\x30\x2e\x42\xe2\xfc\x7e\xed\x7c\x1d\x3b\xaf\xa2\x87\xc9\xa1\xc4\x38\x5c\x16\x51\x19\x2e\x31\x36\x23\x7b\x60\x22\x2c\x07\x53\xdb\x2f\xf1\x0a\x29\x51\xd0\x72\x4d\x98\xa4\x25\x2f\x18\xb3\xf1\xaa\x74\x7c\x9c\x0c\x91\x3f\xc5\x23\x9c\x0c\x6c\x1f\xf9\xd3\x90\x06\x51\x38\xc4\x4e\xa4\x23\xb6\x6f\x6b\x3b\xa7\xdf\x47\x5b\xa2\xe2\x4d\xbb\xe5\x3f\x88\x80\x98\x49\xf0\xc0\xe2\xc5\x76\x45\x85\x35\xbb\x88\xeb\xc9\xb9\x4b\xd3\x1d\x29\x2a\x55\x8d\x68\x33\x4e\x79\xd3\x67\x57\xf3\xf4\xaa\xeb\x7d\xa7\xfb\x56\xb1\xde\x51\xb7\x6e\x4a\x25\x32\x9e\x76\x5d\x36\x3a\x50\x95\xb5\x73\xff\xae\x5a\xac\x8b\xac\xf2\x9c\x51\xc2\x5b\x55\x9a\x4c\x3d\x88\x7f\xa5\xe9\xdc\x36\xaa\xeb\x35\x5b\x73\xb9\x23\x1c\x62\x46\xa4\xf4\x4c\x0b\x86\x15\x60\x08\x96\xb9\xd0\x87\xaa\x8e\x3e\xce\x5d\x9d\xb7\x68\x20\x0f\xf5\x93\x3b\x18\x7f\x02\x00\x00\xff\xff\x60\x64\x0c\xd1\x73\x05\x00\x00")

func tmplJsUtilJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsUtilJs,
		"tmpl/js/util.js",
	)
}

func tmplJsUtilJs() (*asset, error) {
	bytes, err := tmplJsUtilJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/util.js", size: 1395, mode: os.FileMode(420), modTime: time.Unix(1516991133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4d\x6f\xdb\x46\x13\xbe\xfb\x57\x0c\xf8\xbe\x87\x16\x30\x4c\xd7\xbd\x04\xbe\xe5\x0b\x85\x51\xbb\x48\xe5\xa4\x40\x11\x04\xc1\x88\x1c\x89\x1b\x2f\x77\x99\xdd\xa1\x5d\x36\xd0\x7f\x2f\x96\xa4\x28\x92\xa2\x24\x52\xa2\x1c\x23\xf0\xc9\xd6\x7e\xcc\xce\xcc\xf3\xcc\xc7\x2e\xbf\x9d\x00\x78\xf6\x01\xe7\x73\x32\xde\x25\x78\x17\x67\xe7\xde\xa9\x1b\x13\x6a\xa6\xbd\x4b\x70\xf3\x00\xde\x3d\x19\x2b\xb4\x72\x2b\x7e\x29\x57\x00\x78\x2c\x58\x92\x1b\xbb\xb9\xd1\xc1\xdd\x72\x34\x24\x1b\x18\x91\x70\xb9\xfe\x26\xd6\xc1\x1d\x08\x0b\x08\x4c\x96\x85\x9a\x03\xaa\x10\x66\x68\x19\x12\xa3\x59\x73\x96\xb8\x41\xd6\x5a\xc2\x4c\x1b\x08\xe9\x9e\xa4\x4e\xc8\xd8\xa5\x44\x29\x02\x52\x96\x2a\x75\x00\x3c\x85\x71\x71\xf2\xd5\xfb\x72\x15\x80\x97\x1a\xe9\xc6\x22\xe6\xc4\x5e\xfa\xfe\x5c\x70\x94\x4e\xcf\x02\x1d\xfb\x5f\x62\x34\x2c\xd4\x8b\x0b\x3f\x76\xea\xf8\x53\xa9\xa7\x7e\x8c\x96\xc9\xf8\xd7\x57\xaf\xdf\xfe\x71\xfb\xf6\x2c\x0e\xbd\x5c\xd0\xe2\x04\x60\x91\xfb\x20\xd2\x96\x9d\x40\xa9\x03\x94\xee\xc7\xe5\x8b\xf3\x17\x17\x85\x7f\x18\xe7\xd6\xbb\x84\x8f\xf9\x96\x35\xb5\x0c\x7d\x4d\xc9\xf2\x4a\xb5\x96\x53\xde\x47\x04\x85\x63\x0c\x05\xda\x84\x16\x50\x4a\x28\x77\x59\x10\xec\xc6\x49\xdc\x93\x05\xa1\x20\xa6\x58\x9b\x0c\x7e\x42\x06\x49\xce\x6f\xa9\x62\x21\xdd\x2a\x61\xc1\x90\x25\xfe\xf9\x0c\xde\x47\xc2\x42\x8c\x77\x94\x6f\x4f\xb4\xb5\x62\x2a\x09\x58\xc3\x3d\x19\x31\xcb\x80\x23\x64\xc0\xe5\x21\x10\x23\x07\x51\x8e\x06\xd8\x84\x02\x31\x13\x01\x24\xc8\x4c\x46\xc1\x03\xda\xa5\x02\xe1\x69\x0e\x17\x4a\xab\x9d\xac\x19\x71\x10\x01\x47\xb4\x52\x36\x24\x46\x21\xed\x59\xe9\xbd\xd3\x6e\x8f\xd8\x80\x14\x1a\xa1\x37\xba\xe4\x25\x2c\x97\x38\xab\xc8\x5a\x52\x2c\x50\xca\xcc\x29\xc8\xc8\x04\x31\x3a\x85\x09\x1e\x22\x6d\xa9\x18\xb3\x10\xa0\x82\x29\x01\x9a\xa9\x60\x83\x46\xb8\xf5\xd6\x8a\xb9\xa2\x70\x87\x46\x31\x26\x8e\x78\x1b\x15\xfa\x5b\xa7\xb9\xf4\x18\x15\xce\x9d\xbd\xb1\x66\x92\x19\x64\x3a\x35\x60\x39\x9d\x42\x29\xc1\xc2\x43\x44\x8a\xee\xc9\xb8\x39\x50\x44\x35\x26\x7d\xca\xd9\x32\x45\x4b\xef\x90\x23\x27\xd7\xc7\x44\x14\x1c\xb2\x41\x44\x31\xad\x68\x94\x33\xd7\xab\x36\x05\x5a\xd9\xb4\x31\x8f\x49\x22\x45\x80\x4e\x43\xff\x8b\xd5\x6a\xb5\x36\x31\x3a\x4c\x83\x9e\x6b\x91\x23\xbb\x8a\x6d\xbf\x44\xd2\x47\x29\xeb\x21\x36\x27\xae\xfd\x74\xea\xa6\x71\x8c\x26\x73\x36\xfc\x46\xdc\x60\x6c\xe5\xc4\x76\x64\x94\x63\xcb\x78\xa8\xc6\x3e\xd5\x36\xb8\x58\xcf\xf5\xbc\x0a\x77\xc9\x5e\x33\xb3\x1c\xef\x32\x76\xfd\xa0\x04\x0d\xc6\xc4\x2e\xb3\x5c\xc2\xc7\xfa\x8c\x21\x9b\x68\x65\xc9\x36\x2c\x06\xf0\x2e\xce\xcf\x5b\x43\x1d\xcc\x35\x06\x33\xd0\xb3\x22\xa4\x6a\xea\x16\x5e\x73\x20\xe3\x9a\x10\x00\xef\xff\x86\x66\x6e\xff\xff\xfc\x90\x66\x42\x09\x27\xcf\xfa\x2e\x95\xde\x38\x41\xd7\xa2\xee\xb0\x25\x9f\xd6\xff\x5f\xfe\xb7\xa8\xb3\x7d\x85\x6a\xae\x14\x85\x43\x91\x2d\xb7\x1d\x15\xe1\x6d\x67\x3c\x23\x3d\x18\xe9\xbc\x12\xf4\xc5\x79\xe2\x16\x83\x65\x6d\x8e\x04\xf2\xce\x03\x9e\x24\xc2\x87\x62\xfa\x32\x70\x7f\x27\xa5\x0a\x23\x80\x9a\xaa\x3d\x03\xb8\xda\x78\xd4\x10\xde\x7e\xca\x93\x84\xf8\x7b\x07\x71\xcd\x0e\x4b\x41\x6a\x04\x67\xb9\x7d\xdb\x79\x50\x34\x71\x75\x12\x24\x45\x77\xda\xcd\x82\xbf\x8a\x9e\xef\x18\xd0\x6f\x11\x3d\x26\xde\xb5\xfd\x2d\x3c\x37\x35\xd9\xe5\xb4\xc8\x51\x9e\xea\x30\x6b\xcf\xb8\x0d\xc2\xe4\xa1\xc4\x26\xa5\x03\x91\x9f\xb4\x7d\x06\x1b\x31\xff\x61\xb8\xdb\xcd\xd1\x65\xe3\x6e\x8b\x12\xf4\x79\x40\x23\x59\x54\x09\x97\x49\xca\x4a\x51\xc9\xda\x45\xd9\xea\x42\xd1\xbf\x18\xed\x3a\xe6\x49\xa6\xab\x27\x53\x91\x96\xf7\xa5\xa1\x8d\xa4\xbb\xe5\x2e\x6f\x4a\xbb\x40\x5d\x9e\x31\xa8\x8b\xdc\x74\xc0\x93\x84\xd3\xc5\x1c\xac\xf0\x01\x29\xd6\x33\xd8\x3e\x71\xfc\xa6\xfa\x7d\x60\x30\x97\x9e\xf4\xbf\x39\xbf\x7e\x76\x77\xc5\xc5\x10\xc8\xe3\xa6\x7d\x63\x03\xbe\x45\xfc\xe3\x16\x9f\xca\x3b\xdd\xe5\xa7\x6b\xa6\x8b\x09\x5d\xeb\x38\x4b\x8a\x27\x13\x36\xf5\xf7\x89\x72\xb6\x55\xc4\x8e\x5b\x6f\x8e\xcf\xd6\x41\x4c\xad\x9e\x6a\x92\x74\x33\x0b\x3f\x24\x61\xfe\x5a\x74\x3c\x22\xee\x3e\xe1\x99\x8b\xa7\xbb\xd5\x7e\xd5\xd1\xa2\xed\xdb\xbc\x8d\x5c\x34\x87\x73\x74\xbc\x16\x2f\xaf\xd7\x60\xc8\xa6\xf2\xe0\x58\x3b\xa0\xf8\xd7\x9e\x45\x25\x31\x6d\x0c\xb7\x37\xf9\xf4\x31\xc3\x6d\xf7\x09\xcf\xe1\xf6\xe3\xf2\xf0\x04\x1a\xdf\x64\x6a\x52\x57\x4f\xd8\x1d\x0d\x50\xad\x65\x59\x7d\xa6\x72\x6e\xbf\xae\x97\xb0\xca\xe1\xe8\x2e\x56\xab\x61\xc1\x14\x37\xbd\x36\x28\x51\x34\x5b\xaa\xe6\x35\xab\x53\x31\x37\xfb\x9d\x34\xab\xad\xd9\xe8\xb3\x8e\xa8\xab\xf4\xd3\xd3\x2f\x14\xd4\xf4\x4e\x8c\x0b\x5f\x16\x2d\xd6\x79\x1f\x26\x57\x6d\x1a\x6e\x4b\xda\xed\x50\x58\x11\xa3\x46\xf0\xa6\x80\x71\x65\x2f\x5f\x17\x5a\x72\x7b\xbe\x05\x2c\x3a\xc2\xb0\xaf\xa8\x76\x90\xd4\x65\x05\x5a\xb1\xd1\xb2\x9f\xa8\xd7\xe5\xe2\x2d\xc9\xbd\x96\x52\x3e\x36\xb1\xda\xe0\xe6\x0e\x0f\x75\x99\xda\xa1\x72\x39\xf2\xa9\xc1\xc0\xc9\x9a\x97\x57\xd4\x9b\xb4\x3f\xa2\x0e\xa5\x5c\xd4\x7e\x21\x1b\x83\x17\x31\x71\xa4\xc3\xd1\xc5\x26\xc5\xd7\xc1\x71\x85\x7e\x4d\xc9\x64\xb7\xf9\x82\x77\xf5\x22\xd7\x83\x3b\x7f\x76\x6e\xed\x3c\x25\x22\x0c\x7b\xcb\xfd\x9d\xb2\xfc\x11\x6b\x13\xbd\xf5\x5d\x1b\xc6\x2d\xf4\x76\x8b\xbb\x05\xe5\x1d\xe4\x28\xee\xec\x19\x32\x25\x2d\xda\x88\x76\x92\xbe\xdb\xb7\xb5\x10\xe8\xe3\xba\x55\x04\xad\x65\x97\xc1\x81\x62\x19\x39\xb5\xaf\x75\xb8\x96\xa3\x7a\xf9\x4c\x28\xa6\x39\x99\xe6\xe4\x4c\x9b\x18\xb9\x9c\xfe\xf5\xe2\x99\x3a\x1b\xa9\x53\xf3\x7e\x27\x5d\x2a\xbb\x0f\x40\xf8\x8e\xf6\xb3\xa9\xd9\x7d\x40\x77\x07\xb2\xd5\x0b\xbb\x9f\x9c\x4a\x30\x7a\x58\xb7\x49\xe5\x21\x66\xf7\xc9\x9b\xf7\x28\xd3\xb5\x50\xd8\x09\x73\xcb\xa8\x76\xa1\x1e\x8c\x59\x62\x84\x2e\xbf\x0e\x3d\x66\x4c\x86\x24\xf1\xb1\xcf\x0c\x0c\xfe\xbb\xdf\x99\x53\xad\x25\xa1\xea\x16\x5b\x7d\x23\xe8\x95\x14\x6e\xd7\xbe\x28\x2c\x9a\x77\xcc\x7f\xb2\x57\x68\xe9\xc3\xe4\x7a\xdc\xf4\xd0\xe0\xcd\xed\xba\xca\x83\x89\x53\x5e\x54\xc7\xef\x87\x5d\xee\xba\x65\xe4\xfd\xa4\x8f\x9c\x4b\x6a\xba\x29\x7a\xd8\x5f\xad\x7e\xb8\xb4\x6e\xb0\x07\xa0\x53\xde\xa8\x8f\xa0\xaa\xbb\x24\x9f\x2c\xfe\x0b\x00\x00\xff\xff\x35\x82\x09\x63\xd3\x29\x00\x00")

func tmplSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_tmplSwaggerJson,
		"tmpl/swagger.json",
	)
}

func tmplSwaggerJson() (*asset, error) {
	bytes, err := tmplSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/swagger.json", size: 10707, mode: os.FileMode(420), modTime: time.Unix(1516991133, 0)}
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
	"tmpl/css/style.css":        tmplCssStyleCss,
	"tmpl/index.html":           tmplIndexHtml,
	"tmpl/js/mapping.js":        tmplJsMappingJs,
	"tmpl/js/request_logger.js": tmplJsRequest_loggerJs,
	"tmpl/js/sorting.js":        tmplJsSortingJs,
	"tmpl/js/util.js":           tmplJsUtilJs,
	"tmpl/swagger.json":         tmplSwaggerJson,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{tmplCssStyleCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{tmplIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"mapping.js":        &bintree{tmplJsMappingJs, map[string]*bintree{}},
			"request_logger.js": &bintree{tmplJsRequest_loggerJs, map[string]*bintree{}},
			"sorting.js":        &bintree{tmplJsSortingJs, map[string]*bintree{}},
			"util.js":           &bintree{tmplJsUtilJs, map[string]*bintree{}},
		}},
		"swagger.json": &bintree{tmplSwaggerJson, map[string]*bintree{}},
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

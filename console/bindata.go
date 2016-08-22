// Code generated by go-bindata.
// sources:
// tmpl/css/style.css
// tmpl/index.html
// tmpl/js/script.js
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

var _tmplCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\xdd\x6e\xe2\x30\x10\x85\xaf\xd7\x4f\x31\x12\xda\x1b\xb4\xe6\x6f\x11\xbb\x4a\xae\xe8\x0f\xef\xe1\x9f\x49\x6c\xc5\xb1\xa3\xc1\x29\x20\xc4\xbb\xd7\x86\xd0\x54\x6a\x69\xae\x9c\x39\xdf\x39\x1e\x1d\xcb\xa0\x4f\x70\x66\x90\x3e\x67\x3d\x72\x83\xb6\x36\xb1\x80\xe5\x6c\x83\x6d\x09\xf3\x29\x54\xf6\x08\x55\x20\x88\x46\xd8\x74\xf0\x71\x3a\x67\x17\xc6\x98\x89\xad\xfb\x03\x9f\xfc\x1f\xd6\xc5\xe2\x77\x99\x91\x99\x4a\xb4\x48\xa9\xc4\x2b\xd7\x5b\x7d\x66\xbf\x5a\xeb\xf9\x57\xae\x13\x1e\x1d\xbf\x47\xdd\xf5\xbf\xab\x45\x77\x2c\x21\x13\x45\xc1\x0f\x28\x1b\x1b\xf9\x5e\x51\x70\x4e\x0a\x1a\x2e\x3d\x58\x1d\x4d\xca\x5a\x25\x34\x91\xdf\x80\x3c\x92\x50\xcd\x80\xcb\x74\xac\x29\xf4\x5e\x73\x15\x5c\xa0\x02\x26\xcf\xdb\xd7\xf5\x6e\x5b\xde\xe4\x40\x3a\x6d\xeb\xb0\xca\xfb\x75\x47\xd8\x07\x67\x35\x4c\x94\x52\x0f\xd3\x4d\xdf\xca\xc7\xe9\xeb\xd5\x7f\xa9\xc4\x8f\xe6\xc2\x84\x37\xa4\xc7\x11\xcb\x97\x7f\xeb\xa7\xdd\xb5\x29\xd6\x11\x66\x6e\xa8\x68\xb9\xc9\x0d\xa5\xff\x1c\x50\xb9\x70\x28\x44\x1f\x43\x1e\xdc\x6a\x19\x1b\xde\x47\xb2\xbe\x86\x33\x0c\xa1\x35\x21\xfa\x12\x2e\x6c\xe6\xd3\x06\xf9\xf6\xbb\xa2\x05\x35\x81\x84\xaf\xf1\x2a\xcb\x10\x1c\x0a\x3f\xea\xd2\xf5\x38\x18\x9d\x1b\xc7\xad\xa8\x31\x3d\xf6\x55\x69\xf0\x34\x0a\x84\x3a\x0f\xdf\x03\x00\x00\xff\xff\x58\xaf\x91\x51\x68\x02\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/css/style.css", size: 616, mode: os.FileMode(420), modTime: time.Unix(1465898006, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\x6d\x53\xdb\x38\x10\xfe\x5c\x7e\x85\xea\xbb\x9b\x81\xe1\x6c\x13\x52\x0a\x05\x27\x33\x14\xd2\x63\x18\x52\xd2\x26\x07\xd7\x7e\x53\x6c\xc5\x56\x22\x4b\x46\x92\xf3\xd2\x0e\xff\xfd\x56\x96\x1d\x27\x10\x28\xcc\xcd\x31\x90\xc8\xab\xdd\x67\xdf\x77\x4d\xf0\xf6\xfc\xfa\x6c\xf0\xad\xd7\x41\x17\x83\xee\x55\x7b\x2b\x48\x74\xca\xcc\x17\xc1\x51\x7b\xeb\x4d\x90\x12\x8d\x51\xa2\x75\xe6\x92\xbb\x9c\x4e\x5b\xce\x99\xe0\x9a\x70\xed\x0e\x16\x19\x71\x50\x68\x9f\x5a\x8e\x26\x73\xed\x1b\xd9\x13\x14\x26\x58\x2a\xa2\x5b\x7f\x0f\x3e\xb9\x47\x0e\xf2\x0d\x8c\xa6\x9a\x91\x76\xb7\x2b\xc2\x09\x02\x04\x25\x18\x09\x7c\x4b\xdc\x82\x6b\x15\x4a\x9a\x69\xa4\x64\xd8\x72\x7c\x3f\x14\x11\xf1\xc6\x77\x39\x91\x0b\x2f\x14\xa9\x6f\x8f\x6e\xc3\x6b\xec\x7b\x7b\x5e\x4a\xb9\x37\x56\x4e\x3b\xf0\xad\x54\xfb\xa5\xf2\x29\x8d\x25\xd6\x04\x70\xf6\xbd\xc6\x06\x18\xc0\x79\xeb\xba\xe8\x0a\x78\x94\x06\xcf\xd2\x8c\x32\x12\x21\xcc\x23\x04\xcc\x74\x44\xe1\xe1\xac\xdf\x47\xae\x6b\x54\x32\xca\x27\x48\x12\xd6\x72\x94\x5e\x30\xa2\x12\x42\xb4\x83\x12\x49\x46\x2d\xc7\xc4\x4b\x1d\xfb\x7e\x8a\xe7\x61\xc4\xbd\xa1\x10\x5a\x69\x89\x33\xf3\x60\x0c\x5a\x12\xfc\xa6\xd7\xf4\xde\xfb\xa1\x52\x35\xad\xb0\x0c\x28\x0e\xa2\x10\xdb\x58\x52\xbd\x00\x1d\x09\x6e\x1e\xbd\x73\x1b\x77\x47\xe9\xe0\xf2\xfa\xb4\x3f\x3f\x1a\x37\x4e\xf3\x5d\x7c\x70\x7b\x7e\xc3\x7b\x74\x9f\x4d\x3e\x8d\x66\xb3\xce\x29\x3e\x4a\xce\xcf\xa3\xf1\x77\x96\x5d\x91\x78\x9e\x8c\x6f\xba\x9d\xc6\x28\x1e\xdf\xf6\xfe\x4a\x27\x3f\xd4\x21\x24\x4c\x0a\xa5\x84\xa4\x31\xe5\x2d\x07\x73\xc1\x17\xa9\xc8\x21\x0c\x95\xf7\xd7\x99\xa6\x82\x63\x86\x74\x42\x52\xf2\x7f\xfb\xea\x16\x5a\x9e\xf3\x78\x74\x75\xbb\xff\x79\xaf\xc1\xba\x77\x63\x3c\xf9\x38\x99\x37\x99\xdf\xfd\xd0\xc1\x49\x3e\xcb\xfa\x23\xf2\x79\x7a\xf3\xbe\x79\x79\x40\x7e\xf0\x66\xfe\xfd\x07\xce\x06\x7b\xf9\x61\xe7\x9b\xfa\xa7\x3b\xfe\x72\xb3\xbb\xd7\xe1\x07\xf2\xd7\x1e\x3f\x9b\xef\x4b\x3c\xc5\x7d\x5b\x5c\x36\x14\xab\x95\xf6\x5a\xd7\xc7\x0f\xb3\x3c\xde\xe8\xf2\x5e\xda\x1f\x5e\x9e\x77\x2e\x28\x66\xa3\x34\xff\xf8\xf1\x4b\xef\xfd\xe9\xbb\x2f\x32\x93\x77\x07\xd7\x37\xa3\xdb\xe6\x61\xef\xeb\xd7\xe6\xf8\xa0\x73\x75\x37\x57\xaa\xb1\xb8\xb9\xbb\xd6\x9c\x64\xfc\xe2\xa6\xf7\x01\x5f\x1e\xce\xfb\x4f\xbb\xbc\x56\xeb\xa5\x27\x1a\xfa\xb8\x6c\xdf\xda\x59\xc7\x3a\x08\x06\x5b\xae\xc7\x9d\xb2\xb9\x24\x56\xc0\x8a\x7c\xda\x12\x31\x29\x2f\xb8\x8a\x24\x2f\xa7\x0a\xc7\x29\x30\x4f\x29\x99\x65\x42\xea\x95\x59\x32\xa3\x91\x4e\x5a\x11\x99\xd2\x90\xb8\xc5\xc3\x9f\x10\x26\xaa\x21\x20\xae\x0a\x31\x23\xad\x86\x41\xf9\x85\x0b\xc0\xf1\xfb\x36\x8a\x44\x98\xa7\x80\x8a\x76\x3c\x09\x23\x6d\xb1\x3d\xca\x79\x68\x4a\x7c\x7b\x07\xfd\x04\x0c\x84\xa6\x58\xa2\x99\x42\x2d\xc4\xc9\x0c\xdd\x92\x61\x1f\x66\x14\xd1\xdb\xce\xcc\xa4\xd6\x41\xbb\x88\x89\x10\x1b\x09\x2f\x11\x50\x27\xbb\xc8\x39\xfe\xf9\xd3\xeb\x81\xc9\xf7\xf7\xce\xce\x49\x81\x31\x53\x9e\xe0\x29\x51\x0a\xc7\x04\x90\x96\x3a\x48\xa5\xa4\xd4\x53\xb3\x5c\xf6\xaf\x3f\x7b\x99\x19\x94\xdb\x64\x0a\xf6\x79\x11\xd6\xb8\x44\x83\x5f\x26\xe2\xaf\x30\x70\xa1\x30\xb7\x4b\x99\x9d\x13\x73\x65\x7e\x0a\x9e\x7b\x60\xbd\x2f\xf8\xeb\xb4\x04\xbe\x9d\xda\x5b\x5b\xc1\x50\x44\x8b\x76\x59\xe0\x7d\x8d\x61\xea\x42\xa7\x41\x84\x59\x9e\x72\x85\x04\x47\xa9\x18\x42\xb5\xa3\xe1\x02\xa5\x78\x42\x79\x0c\x34\x02\x76\x33\x66\x03\x5e\x74\x81\x11\x11\xf0\x21\x51\x02\xa5\x58\x5e\xd8\x36\x88\xe8\x14\x85\x0c\x2b\x05\xd9\x85\xb4\x61\xca\x89\x74\x47\x2c\xa7\x51\x95\x46\x28\x21\x93\xf2\x96\x93\xe1\x28\x02\x05\xc7\xa8\x71\x90\xcd\x4f\x6c\xe2\x56\xc4\xa5\x98\x15\xb4\x07\x98\xcc\x4d\x23\xb7\xb1\x6f\x4c\x76\x59\x6c\x4f\xc5\xda\xb0\xcc\x6b\xdc\x19\xe6\x84\xa1\xe2\xd3\xcd\x24\x4d\xb1\x5c\x38\xc8\xb2\x3d\xe6\x73\x4d\x8c\xc0\x9e\x12\x07\x38\x92\xe6\x3a\x43\xb1\x96\x9c\x87\xcb\x2a\x69\x56\x88\x3e\x40\x3e\x89\x6e\x02\xbf\x74\x5d\x4c\x89\x1c\x31\x31\x3b\xc6\xb9\x16\x4b\x85\xe5\xd7\x9b\x20\x67\x95\x2c\xa3\x4a\xbb\xb1\x14\x79\x06\x03\x21\x6a\x39\xc5\xf1\x6c\xcd\xe1\x35\x49\x3f\x67\xeb\x68\xcf\x1b\x35\x82\xa9\x43\x64\xed\x31\xc3\x43\x08\xd8\x7a\x7e\x5c\x2d\x32\xc8\xd1\x1f\x27\xb5\xbe\x80\xf2\x2c\xaf\xda\x2b\x4c\x48\x38\x19\x8a\xb9\x35\x30\x4c\x26\xa7\xe0\x13\xb4\x9a\x60\x0c\x5a\xd7\x5c\x12\xa0\x6b\x99\xc3\x5b\x41\x1b\x99\x4b\xa4\x8a\xdb\x4a\xa9\x5f\x68\x5d\xda\x30\xcc\xb5\x86\x3a\xb4\xe0\xf6\xc1\xa9\xac\x1e\x6a\x8e\xe0\xaf\x4e\xa6\xc6\x43\xca\x23\x32\x6f\x39\x7b\x56\x3f\xdc\x9e\x31\x82\xe5\x95\x88\x97\xd1\x86\x48\x63\x7d\x0c\x53\x2f\xd1\xb5\x13\x81\x82\x10\x54\xb8\x31\x5b\x64\x09\x85\x42\x42\xcb\x93\x2b\x49\x0a\x69\x72\x15\x8d\x79\x31\xe2\x80\xbd\x8d\x0a\xec\x3a\xf7\xd6\xbc\xc7\xf9\xaf\x8f\xf6\xf4\x66\x6b\x85\xf6\xcb\x22\x87\x41\x26\x35\x2a\x3e\x5d\xca\x47\x62\xe9\x48\x44\x55\xc6\xf0\xe2\x18\x06\x37\xb1\xee\x16\x4c\x5d\x18\x04\x84\xc7\x50\x10\xb7\x84\xc1\x7e\x21\x5e\xad\xeb\x29\xa5\xe8\x45\xad\xa5\x34\xd6\xc5\x4a\x7c\x49\x67\xfd\xe7\xc6\xba\x18\x0c\x7a\xe8\x1c\x36\x00\x65\xea\x75\x7d\x65\x36\x4f\x7d\x6f\x02\xa3\xa9\x24\x21\xce\x9e\x08\xdd\x86\x40\xab\x3c\x0c\x89\x59\x4d\x95\x95\x08\x05\xb0\x90\x05\x8f\xdb\xb6\x54\x2c\x6a\x15\x11\x54\x56\x04\x7c\x59\x26\xb4\xca\x25\xed\x88\xae\xd9\x96\xb5\x5e\xbb\x52\x35\xed\x4a\xb7\x73\x3c\x85\xe5\x37\x75\xa1\xaa\x95\xb3\x62\x07\xa3\x4b\x8b\x61\x7d\x4c\x4d\x23\x05\xb8\xdc\xa1\xbf\x01\x73\xa5\x0e\x99\x5d\x01\xfd\x1a\xc7\xc6\x65\xb8\x70\xda\xe5\xb2\x08\x7c\x0c\xa6\x32\xba\x06\xfa\x10\x44\x65\x50\xd7\x64\x33\x8a\xbd\x7b\x01\x0c\x33\x8d\xf7\x18\x01\xfa\xf1\xa1\xb0\x1d\x56\x6b\x99\x2b\xbd\x04\x09\xb7\xda\x16\x2b\xaa\x96\xb9\x5d\xf1\x78\x45\xc2\xd4\x03\x1a\xe1\x88\xc0\x3b\x01\x2a\x03\x55\x4b\x83\x7c\x26\x49\x21\x9f\x44\xcb\xfc\x04\x3e\x10\x57\x54\xac\xe6\xe7\xa1\xc6\x2a\x3c\x9b\x54\x3e\xad\xa8\x94\x7a\x85\xa6\x22\x82\xaf\x51\x62\x04\x36\xe2\xa3\x67\xea\xae\x24\x6d\x1c\x57\x5b\x2b\x93\xa3\xa2\xc0\xbb\x83\x6f\x5f\x1e\x02\xdf\xfe\x2b\xf8\x6f\x00\x00\x00\xff\xff\xab\x00\xb0\xad\x22\x0e\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/index.html", size: 3618, mode: os.FileMode(420), modTime: time.Unix(1471872890, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsScriptJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x6d\x6f\xdb\x36\x10\xfe\xee\x5f\x41\x68\x45\x25\xc5\xf1\x4b\x87\xec\xc3\xec\xbc\xac\x4b\x36\x64\x43\xbb\x02\x4b\xbf\x6c\x96\x03\x30\x12\x6d\x73\xa1\x48\x8d\xa4\x9a\xb9\x75\xff\xfb\x8e\x94\x64\x93\x74\xec\xa6\x85\x63\xf9\xee\x79\xee\x8e\xc7\x87\x27\x7e\xc2\x12\xe5\xa2\xe6\x1a\x5d\xa0\xf1\xb4\x67\x7e\x4a\xf2\x6f\x4d\x94\x56\x60\x99\xcd\xa7\xbd\xde\xab\xa4\x10\x79\x5d\x12\xae\xd3\xa1\x24\xb8\x58\x27\x8b\x9a\xe7\x9a\x0a\x9e\xa4\xe8\x4b\x0f\xc1\xbf\x57\x49\xf4\xdd\x83\xe6\xd7\x8c\x60\xf9\x4e\x2c\xa3\x74\x98\x33\x9a\x3f\xee\x03\x3b\xf0\x52\x8a\xba\xba\x16\x5c\x09\x46\x00\x4d\xca\x4a\xaf\x93\x74\xea\x61\x34\x95\x24\xc7\x15\xb8\x57\xb4\x20\xa1\x77\x55\xb4\x75\x1a\xbf\x2e\x59\x12\x45\xcf\x20\x54\x05\x39\xc8\x11\x08\xb3\xd5\xfa\xde\xaf\xf0\x6d\x3e\xbd\xae\x7c\x44\x79\x2e\x89\xe9\xc0\xb5\x69\xd5\x76\x35\x92\xe8\x5a\xf2\xa6\x7f\xfd\x3e\x90\x1c\x8a\x5a\x89\xa7\x1b\xa2\x31\x65\x2a\xa1\x45\xc7\x80\x74\x9d\xb1\x6b\xf3\x8c\x16\xf3\xd4\xe7\x2e\x09\x24\x62\x42\xfe\xbc\xbe\xd3\x58\xd7\x2a\x51\xf6\xeb\x5a\x14\xc4\x04\xb2\x91\xe8\x02\x39\x66\x74\x71\x81\xbe\x1f\x8f\xd1\x66\x83\x42\xe3\x1b\xb7\xf7\x6d\xc5\x91\xaa\xf3\x9c\x28\x15\xb5\x0b\x46\x84\x29\xf2\x4c\xc8\xb3\xf1\xd9\x73\xec\x02\xf3\x25\x91\x3e\x79\x1f\xf5\x84\x25\xa7\x7c\xd9\xc1\xbc\x15\x3a\x7d\xf8\x47\x09\xde\x25\x71\xe4\x07\xea\xfb\xfd\xee\xc3\x1f\x43\xa5\x25\x04\xa1\x8b\xb5\x05\x0e\x5b\xef\x29\xaa\x79\x41\x16\x94\x93\xe2\x14\x9d\xb5\xfb\xd6\xb0\x9b\x2d\x3f\x48\x6f\xdc\x07\xf9\x50\xd8\x61\x6a\xcd\x0e\x27\x6e\x1a\x07\x5c\x2f\xcf\x70\xd7\xcf\x69\x6f\x5f\xd8\x46\x24\x9d\xb0\x7d\x0f\xd6\x5a\x26\x71\xce\xb0\x52\xf1\x29\x8a\x31\x23\x52\x23\xfb\x77\x10\xa3\xfe\x21\x85\xa4\x5e\xac\xc6\xd6\xc9\xbb\x45\xf8\xc9\xbc\x13\xe4\x36\x78\x58\x12\xbd\x12\x05\xa4\x8a\xe0\x7f\x1f\x79\xbe\x0a\xeb\x95\x13\x68\xef\x28\xaa\x35\xd7\xf8\xbf\x5b\xba\x5c\x31\xf8\xe8\x4e\xea\x69\xc0\xf1\x0f\xe7\x3e\xa9\xf1\xfb\x2c\xe7\xbc\x86\x04\x70\xa5\xcd\x41\xf2\x74\xf6\x67\x93\x7c\x4f\x67\x2f\xdc\x30\x03\xa5\x05\xc0\xc2\x19\xb0\x73\x17\x58\x13\x4d\x4b\xa3\x39\xb3\x2f\xb5\x94\x80\xfa\x08\x06\x17\xb4\xa8\x19\x7b\x67\xc5\xb5\x85\x9b\xde\x9e\x0f\xf6\xdb\xfb\xad\xd6\x4f\xdb\xe1\xb3\x9b\x1f\xed\x2a\x76\x8d\x0a\xe6\x2b\xae\x2a\xc2\x8b\x24\x3e\x67\x14\x59\x4d\x5d\x44\x0c\xe4\x31\xb0\xb0\x01\xd5\xa4\x44\xc1\xef\xa3\x2a\x03\x57\x1c\x21\xc1\xed\x90\xbf\x88\xdc\x51\x67\x68\xd4\xd4\x1e\xa7\xd3\x76\x10\x2c\x30\xcc\x87\xe8\xd2\x78\xba\x26\x80\xfb\x7c\xc4\xe8\x65\xdc\x36\x28\x1c\x07\xd6\x98\x9b\xb7\xc9\x07\x56\x00\x41\x75\x9d\x54\xb9\x14\x36\x84\xba\x11\x4f\x3c\x09\xb7\xdb\xa7\x38\x9b\x0d\x09\x7e\x83\x55\xdd\xd1\xcf\x66\x97\xc2\x0e\xc1\xe2\xcd\x61\x04\x67\x97\xc7\x8c\x42\x97\x73\x89\x7e\x18\xfb\xef\xb0\x38\x8c\x30\x59\x50\xa9\x74\x6c\xde\x91\xa5\xf8\xb4\x8d\xe4\xcf\xbd\xb0\xfe\x36\xa4\x49\x67\x8a\xca\x57\x8f\x6f\x6b\x2d\xee\x2c\x0a\x4a\xa2\xd0\xd0\x49\xbe\x22\xf9\x23\x29\xe2\xf4\x1b\x2f\xd1\x26\xf6\x47\x51\x25\xcf\x38\x61\x2b\x93\x71\x87\xb9\x25\xe6\xc0\x3c\x5b\x60\x28\x60\xa7\x87\x79\x63\x37\xf2\x85\x1e\x72\xf2\x84\x6e\xe0\xf1\xc0\x41\x70\xc0\x26\x77\x83\x34\x92\x1e\x81\xa4\xb7\xab\x48\x02\xd8\x7b\xc1\xf5\xca\xe2\xde\xec\x83\x03\xec\xaf\xa0\xa5\xbf\x60\xbb\x9b\xb0\xe8\x27\x74\x04\x7b\x2b\x6a\xa9\x1a\xe0\xe4\x08\xec\x3d\xe5\xb5\x26\x2f\x00\xde\x91\x5c\xf0\x62\xab\xca\x56\xe8\xdd\xf2\x83\x8b\x40\x30\xa6\xdc\x39\x64\x9e\x77\x13\xa8\x62\x38\x27\xc9\xe8\xf5\x68\x09\x13\xff\x35\x2e\xab\xa9\x55\x53\x6b\x3e\x6f\xcc\x4c\x7b\xd6\xcb\xc6\xba\x34\x56\xaf\x18\x3f\x66\x12\x25\x59\x56\xcf\xf0\xe0\xf3\xdb\xc1\xdf\xe3\xc1\x8f\xf3\x2f\x67\x5f\x37\x59\x36\xbb\xaf\xe7\x9b\xd9\x7d\x96\x45\xf3\xf4\x04\x20\xea\x64\x92\x5e\x6d\xb2\x87\x44\xcb\x9a\x6c\xec\xc1\xdd\x70\xe8\x73\x9a\x3d\x6c\x06\x57\x59\xd1\x4f\xae\x26\xd9\x30\x2b\x4e\xd2\x2b\x78\x9a\x91\x5f\xe6\xb3\x7e\x36\x98\x1b\x4f\x7a\x95\x9a\x52\xb6\x37\xbe\x12\xeb\x7c\xe5\x2a\xd6\x4a\x88\x99\x81\x1b\xf3\xba\x7c\x20\x32\xde\xdd\xc5\x8c\xfe\x47\xf7\xd1\x68\xa8\xcd\xa4\x6e\xa8\x2e\x77\x8b\x99\xbc\x3a\x8a\x69\xc6\x86\xcd\xf1\x48\xd6\x4e\x02\x2b\xf4\xf0\xa2\x12\x30\x9a\xf7\x7d\x48\xea\x05\x74\x5b\xc6\xae\x3d\x47\xcb\x69\x03\x3f\x08\x38\x83\x98\x3b\x91\xdd\x58\xa6\xbf\x2f\x89\x62\x70\x6e\x88\xf0\xbe\x15\x9f\xab\x0a\xf3\x6e\xbc\x9b\x69\x6b\x88\x66\x52\xdb\xd1\x6b\x63\x37\x83\xd7\xe0\x2e\x63\xe7\xae\xfb\x7f\x00\x00\x00\xff\xff\x08\xc1\x31\x2f\x05\x0c\x00\x00")

func tmplJsScriptJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsScriptJs,
		"tmpl/js/script.js",
	)
}

func tmplJsScriptJs() (*asset, error) {
	bytes, err := tmplJsScriptJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/script.js", size: 3077, mode: os.FileMode(420), modTime: time.Unix(1465898006, 0)}
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
	"tmpl/css/style.css": tmplCssStyleCss,
	"tmpl/index.html": tmplIndexHtml,
	"tmpl/js/script.js": tmplJsScriptJs,
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
			"script.js": &bintree{tmplJsScriptJs, map[string]*bintree{}},
		}},
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

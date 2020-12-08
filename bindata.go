// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// static/service.crt
// static/service.key
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticServiceCrt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x95\xb7\x12\xb3\x48\x16\x46\x73\x9e\x62\x73\xd5\x16\x42\x80\x04\x61\xd3\x34\xbe\x11\x56\x98\x4c\x78\x81\x70\xc2\xf3\xf4\x5b\xf3\xef\xd4\x24\x73\xc3\x73\x83\x2f\x39\x55\xe7\xbf\x7f\x9d\x80\x64\xd5\xfc\x0f\x44\x8e\xa7\x4a\x2a\x04\x1e\xfa\x43\x09\xac\xaa\xe8\x7e\x42\x08\x66\x0c\x25\x65\x19\xc6\xeb\xda\x63\x64\x35\x69\xad\x0c\xcb\x9b\x0e\xb6\xa7\x4e\x72\x73\xfb\xc0\xe0\x2a\x43\x77\x94\x5d\x35\xa1\x45\x1b\x09\x70\xf3\x01\x46\xd3\x4e\x20\x0f\x58\x42\x69\xbe\x04\x80\x31\x44\x26\xab\x22\xf3\x88\xc2\xb2\xc4\x8e\xbd\xa1\x32\x12\x5f\xb6\xad\x8b\x60\x16\x73\xff\xff\xdc\x0d\xd8\x5a\x15\x91\x87\x05\x24\x03\xca\x47\x04\xdc\x36\x3d\xa2\xbf\x75\xda\x4a\x8c\xd1\x7e\x57\x55\x44\x06\x06\xcd\x9f\xa7\x50\x61\xf8\x6a\x98\x4d\xa9\x52\x13\x7b\xcd\x66\xd6\xe8\xc4\x5e\xb9\xe1\x13\xb3\xc1\x5f\xac\xc6\x1b\x81\x6b\xc0\xfe\x03\x6b\x10\x63\x27\xdd\x24\xfb\xcf\xb2\x28\x02\xf6\x9a\xdc\xf8\xc9\x68\xe7\x21\x56\x9c\xc1\x68\xcd\x35\xf1\x20\x04\x1f\x75\x13\xed\x48\xd3\xfb\x58\x25\xaa\x35\x35\x81\x8d\x04\xc1\x06\x62\x59\xaa\x16\x10\x21\x04\x65\x0f\xcb\x52\x15\x80\x09\x5b\xed\xa0\x24\xd7\x3a\xd9\xef\x9c\x32\x27\x9b\xa8\x21\xd0\xb1\xb8\xb9\x9f\xa0\xff\x15\xe6\xed\x7d\x21\xa8\x6e\xe2\xbf\xd1\xdc\xd6\x6e\xdf\xaa\xb3\xaf\x4c\x01\xa3\xac\x19\xf9\xcc\xf8\xfc\x35\x2e\xd4\x27\x94\xdf\x24\x4b\xc3\xa9\xfd\x5d\xb4\x40\x1a\x63\x51\xd6\xbe\x2c\x30\xea\xbb\x6f\x49\xbb\x98\x30\x04\xd7\x34\xad\x41\x06\x61\xbb\x33\x1b\x2f\x5b\xcf\x89\xf2\xcd\xe3\x87\x6f\xdb\x52\x54\x9b\x97\x91\xab\xb7\x42\x45\x43\x28\xb6\xdd\xa2\xce\x12\xeb\x6e\x5d\x0f\x30\xfb\x8a\xad\x70\x93\xe3\xe6\x36\x11\xb9\x19\xad\x5d\xad\x5c\xa5\x5c\x46\x92\x84\x2a\x4f\x6a\x2e\x0f\xd5\x3a\xaf\x42\xd5\x96\xf5\x71\x7c\xbc\x78\x71\x8b\x4d\x86\x06\xb5\xd4\x93\xb7\x30\x6f\x2a\x3e\x7e\x46\xc9\xf6\xcc\x55\x3e\x97\x88\xda\x52\x0a\x61\x0d\xa8\xa6\x6d\x96\xe9\xe1\x2d\x24\x1b\xab\x8c\xf1\x78\xf7\x5d\x33\x22\xdd\xe5\xc7\x37\x7e\x39\x4f\x19\x84\xef\x6f\x47\x5d\xe4\x2f\x0b\x17\xfa\xa0\x50\x95\x1d\xcc\xfd\x9b\x82\x92\x70\x3d\x68\xb9\x7b\xbf\x9f\x59\x55\x30\x5f\x51\x7e\x99\xb4\xc6\x9d\x77\xc7\x82\x97\xd2\xc3\x14\x46\xe9\xb3\x5d\x1e\xe4\xf8\x4b\x5b\xea\xd1\xd2\xd7\x81\x9c\x55\x31\x74\x90\xbf\x0f\xe2\x53\x81\x05\x61\xa4\x77\xe6\xe7\x31\x5c\x0f\x5c\xa9\x7e\x7c\x39\x37\x9d\x38\x8d\x7a\x16\x06\xd9\xb6\x4e\x5c\xd3\x0f\xf2\xc7\x1a\x2f\xbf\xf7\x12\x05\xdc\x48\xcb\xb0\x8c\xcd\x4a\xc7\x32\x7e\xe6\xee\x59\x68\xee\x9b\x40\xf2\xdb\x10\xd6\xbb\xc6\xea\xf9\x79\xdb\x4b\xb5\xf9\xe5\xf7\x19\x07\x89\xfe\x68\x45\xe9\xc1\x1c\xef\x95\x44\x9c\x16\xc4\x5c\xd0\x69\xc1\xf3\x17\xdd\xad\x8f\xc1\x27\xf0\x72\xfa\xf4\x5d\x52\x57\x9e\x58\x38\xc4\x3e\x7f\xf8\xad\x87\xee\x62\x3f\x2e\x23\x84\x5d\xb7\xfb\x47\xcc\x06\xa5\xce\x65\x26\x59\x0c\x3f\xea\x19\xf0\xfb\x90\x0b\xa4\xd1\x2b\x1d\xd8\xbb\x72\x9c\x75\xcb\xd9\x98\x60\xfb\x46\x2f\xe2\x26\x5b\x4b\x4b\x31\x2c\x8c\x9a\xa7\x23\x41\x37\x60\x3d\xcb\x53\x28\x83\xa4\x1f\x97\x9a\xfe\x85\xa8\xff\xe1\x8e\xd1\x85\xbe\x7b\x72\x7d\xd5\x58\xb9\x91\xba\x99\xdc\x5f\xdf\x8f\xc5\x8f\x66\xf8\x26\xaa\x1a\x46\x5c\x4c\xde\xc3\x85\xc9\xf1\x84\xa4\x24\x31\x7e\x91\xa7\xaf\xdd\x17\xbf\x84\x28\xbe\x06\xa8\xbe\xfe\xba\xa1\xac\xc9\x0a\x8f\xb3\x53\xc0\xcb\xc9\xc4\x82\xf8\x94\x5f\xdf\xe2\xd0\xef\x1e\xf1\x0e\xfc\x6d\x14\x1a\x01\x94\x58\x00\x00\xfd\xe3\xf8\xdf\x8a\x1b\x7f\x2b\x2e\x00\x58\x2b\xf9\x78\x8f\x0c\xb6\xbe\xb4\xab\x6d\xaa\xfc\xf5\x31\x05\x5b\x42\x3c\xe9\x14\xb5\x86\x6b\x46\xbc\x27\x04\xa1\xc5\xe9\x67\x7a\x7c\x6e\x46\xf3\xfa\xde\x8a\x9e\x34\xa8\xc3\x28\x83\xa3\x2b\x9c\x57\xc9\x7d\xc7\x24\xe4\x9b\xa5\x13\xdf\x12\xf3\xba\x80\x2f\x99\xbd\x68\x85\xb0\x28\x9f\x46\xef\xa4\xb2\xaf\xe7\x32\xfa\x0e\x09\xe6\x32\x50\xf5\xa5\x1c\x76\xff\xbe\xdc\x7b\x72\x78\xe0\x82\x5d\xb6\xa4\xcc\x38\x92\x6e\x5e\x49\x69\xaa\x77\x9e\x0b\xcf\x65\x48\x7d\x64\xcc\x0c\x41\x77\xe1\x0e\x59\x47\x90\x64\xdd\xcd\x8c\xfa\x67\xa0\x5d\xd9\x58\xf3\xb6\xc5\x2e\xb4\x4e\x4e\x00\xe8\xc4\x55\xdf\xde\xe5\x62\x07\xc9\x35\x2f\xf5\x50\xf3\x6a\x65\xba\x5a\xb7\x90\xcf\xa8\xcc\x22\xba\xd0\xba\x5f\x70\x2a\xea\x0f\x53\xe9\xa7\xe3\x02\xb5\x3a\xf2\x6e\xe1\xd2\xad\x6b\x47\x97\x8f\xae\xf2\xe8\x47\xff\x48\xaf\x0e\x95\x77\x5f\x4e\x45\x7b\x76\x4b\x70\xfb\xe9\x2c\xc5\x52\xe6\x26\x87\xc4\xf9\xd1\x9e\x73\x12\xcd\x39\x6b\xa7\xc1\xc8\x75\xe9\x44\x36\x13\x56\xbf\xd8\xd9\x1d\x77\x3e\x52\x10\x36\x5d\x7a\x52\xcf\x4b\x1f\xd3\x6c\x15\xd3\x71\x08\x78\x0b\x7d\x04\x8e\x91\x4c\x24\x41\x95\x80\xaf\x96\xf1\x48\x4f\x1b\x6a\xef\x18\xa8\x21\x65\x7b\xcd\x7d\xe9\x47\xc5\xfb\x87\xb3\x65\x59\x26\x9c\x92\xfc\x4e\x7f\x30\x5f\xc7\xa6\x09\x13\xd6\xa3\xe3\xf6\x04\xc5\x2d\x65\xdf\x9b\x5e\x93\x1e\x51\x14\xac\x56\x45\x22\xfb\x1b\xf8\xc8\xef\xb7\xd7\x9a\xf0\x4a\x14\x07\x9e\x5c\x54\x8b\x2f\x88\xf6\xec\x24\x97\xda\x20\x03\xb0\x67\xf5\x04\xe0\xe0\x09\x65\x3a\xfc\xd2\x1b\x5d\xce\xdd\xec\x68\x03\xe1\xe7\x85\x6d\x06\xf2\x35\x52\xae\xd1\x7b\x32\xa3\xca\xe7\x9b\xd5\xec\x62\xe1\xc7\xd6\x8e\xc9\x99\xd6\xc2\xf4\x0b\x9f\x85\x07\x8b\xc5\xbc\x6f\xeb\x30\x31\xc8\x98\x0b\x82\xfa\x9a\xcc\xe1\xaa\xc8\x44\x49\xd1\xbe\x66\xaf\x7b\xa2\x2d\x13\xc6\xf4\x85\xaf\xd4\x2f\xe4\x8c\x4e\x38\xfd\x21\x07\x2b\x85\x75\x3a\x2f\x50\xc9\xeb\x5c\x6a\x09\x2e\x25\x08\x21\x8d\x0d\x54\xac\xd1\x87\x24\x6d\xfd\x94\x08\x0f\xaf\x29\xe6\x4e\xec\x74\x1c\x7d\x9d\x64\xce\x9d\x9a\xec\x30\x2e\xde\x49\x0e\x6e\x6b\x1b\xb0\x2d\xaf\xb2\xc3\x1e\xe3\x32\x95\x2f\x25\x47\x03\x6b\xc1\x56\x88\x9d\xe0\x6e\x57\xea\xc4\x24\x34\x41\xbd\xaf\xd2\xa4\xed\xcd\xea\xec\xbe\xa6\xf9\x8f\x6f\x2b\x5a\xc4\x9f\x68\x21\x53\xfc\x77\xc8\xfe\x17\x00\x00\xff\xff\x9f\xbd\xd5\x0a\xe5\x06\x00\x00")

func staticServiceCrtBytes() ([]byte, error) {
	return bindataRead(
		_staticServiceCrt,
		"static/service.crt",
	)
}

func staticServiceCrt() (*asset, error) {
	bytes, err := staticServiceCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/service.crt", size: 1765, mode: os.FileMode(438), modTime: time.Unix(1560449017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticServiceKey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x97\xb7\x0e\xac\x8a\x01\x44\x7b\xbe\xe2\xf6\xc8\x02\x96\x5c\xb8\x20\x2f\x39\xc7\x8e\xb4\xe4\x9c\xf9\x7a\xeb\xbd\xdb\x7a\xda\xe9\x8e\x46\x1a\x9d\xff\xfc\x13\x56\x90\x64\xe3\x8f\xe3\x32\x7f\x2c\x47\x0e\x18\x4f\xf8\xa3\x0a\xf1\xbf\x0d\xa0\xcb\xb2\xa2\x32\x32\xcb\x30\x2a\xc7\x54\x02\x03\xab\xf1\xa8\xf9\x81\x0c\x59\x43\x88\xbc\xad\xd5\x6f\x75\xce\x4c\x97\xc5\xaa\x4a\xda\xec\x14\x9a\xac\x64\x58\xbe\x9f\xa0\xf9\xa4\x86\xda\xc4\x0d\x80\xd8\xe5\x9d\x55\x25\x8d\x52\x38\x8a\xe0\xe1\x42\x86\x4a\x9e\xac\xd0\x98\xab\x4f\x7a\x62\x7d\x1c\xce\x9d\x1e\x0f\x51\xc7\x1c\xd6\x6c\x53\x91\x60\x9c\xed\xe3\xba\xf1\x49\x27\x65\xf6\x6d\x79\x18\x88\x29\xe2\x0d\x6c\x54\xdd\xde\x8c\xc3\xc1\x2f\x6b\xa0\x20\x6d\x52\x72\xde\xd9\x4e\xcf\x2a\xa0\x81\x6c\xd0\x04\x79\x32\x87\xd8\x05\xf3\x7b\x6e\x51\xc1\xd9\x5a\x41\xf3\xd1\xa6\xb1\x36\xd0\x2a\x1f\xf0\x25\x11\x3f\xac\x1b\xcf\x1b\x83\x10\xf3\x9a\x89\x97\x65\xdd\x23\xd8\x62\x3b\xb4\xcd\x39\x3d\x6c\xa1\xa4\x20\xdc\x2f\xd2\xe1\xf5\xc9\xe0\x58\x8b\x60\xa1\x4e\x15\x94\x44\x67\xa1\x9b\xba\xc4\x0b\xec\x26\x09\xf5\x5d\x5b\x9d\xfb\xd2\x94\xee\x64\xcf\xda\x67\x9e\xee\x40\xc0\x62\xb6\x58\xc2\x5f\x84\xa5\x51\xa7\x92\x3f\xcd\x77\x45\x54\x6b\x97\xe0\xe2\x58\x45\x97\x69\x4d\x6a\xbf\xc6\x47\x84\x1a\xc0\xd7\x13\x1f\xcd\xc7\xd7\x9a\x05\x4a\x23\x58\xfd\xf5\x2f\xe7\x22\x52\xf2\x04\x97\xfd\xc9\xa2\x23\xfb\xb9\x23\x08\x57\x46\x21\x38\x9e\xd4\xe9\x58\xae\x50\xfb\xbb\x36\xbb\xd5\x3e\x15\x2b\x87\x16\x70\x44\xb7\xf2\xbc\xd7\xe8\xe3\xd4\x09\xa6\x89\x38\x58\xbf\x13\x3c\xbb\x3d\x70\x6b\x63\xcb\xf9\x8c\xa2\x36\xea\x64\xe8\x67\x61\x3b\x1c\xd7\x0c\xaa\x57\x26\x3b\x99\x36\x09\x9a\xc6\x6e\x26\xad\x0c\x90\x98\xc6\xdd\xc9\x05\xbe\x93\x6d\x3d\x62\x3d\x3a\x5b\x65\x81\x28\xef\xea\x93\xf1\x4e\xf3\x3e\x26\xf6\x76\x02\xe5\x13\xd9\x34\xd2\xf0\x7e\x93\xdf\x40\x1f\xf2\xf2\x3a\x62\xbb\xa7\xb9\x50\x71\x1e\x38\x89\x49\x56\xcb\x9f\xe8\x29\x63\x9f\x32\xeb\x0d\xa3\x34\x38\x9c\x3e\x9e\xa1\x52\x8f\x7d\xa9\xa3\x2a\x73\x5e\x2a\xb9\x05\x9e\xe8\x2f\x5f\x67\x5c\x50\x47\x71\x0c\x12\x59\xd4\x76\x4a\x63\xe3\x02\xe0\x2b\x4a\xd6\xeb\xb4\xfe\xf3\x4f\xe8\x77\x12\x56\xb9\xe7\x6e\xb3\x3d\xbf\x6a\x92\x9a\x9c\x7c\x1a\xc9\xa6\xb1\xfd\x56\xe0\xb4\x75\xd6\x03\x77\x1a\x58\x59\x24\x5d\x69\xef\x11\x9a\x4b\xb2\xc6\x1e\xe0\xb2\x83\x7d\x3b\xf6\x5a\x5f\xc0\xd2\xbf\x7d\xb1\x1e\x9d\xd8\xb5\xdc\xa3\x1c\x38\x0b\x14\x9e\x05\x11\x29\x8d\xb4\xea\x4e\xd0\xb1\x24\x0c\x20\x79\x9d\x8d\x39\xf1\xb8\x49\xb2\x05\x8e\xb9\x04\x06\x60\xec\xbf\x13\x3e\x8a\xf0\x65\x1c\x9b\x3c\xcd\x8f\x10\x1c\xe4\x77\xfc\x49\x71\x98\xb3\xab\x0d\x2a\xcd\x22\x9e\xe2\x10\xaa\x0e\x4a\x3a\xab\xf1\x7d\xf9\x5c\xc2\x44\xae\xc6\x4d\x68\xb3\x2f\x20\xeb\x61\xef\xe9\x7d\x79\xfe\x98\x24\xbf\xc3\x56\x75\x71\x32\x42\x7f\xf0\xec\x84\x4e\x78\xaf\xbc\x8c\x3b\x62\xfb\x0b\xee\xdd\xec\xa9\xdf\x57\x05\xe9\xe9\xbc\xe2\xdc\x75\x8b\x87\xd4\x68\x1b\x77\x81\x69\x99\xa8\x4e\x11\x50\x35\x9a\xd9\xb6\xa1\xa4\xe6\x9b\xfb\x6d\xdb\x24\x60\xc8\x7d\x3a\xd9\xbb\xfb\x5d\x2c\xf6\x34\x47\x0a\x23\x8d\xe3\x4c\x78\x57\xcd\xfa\x44\x65\xe9\x86\x82\x38\xbb\x94\x4b\x02\x61\xdf\xf8\xaf\x07\x19\xb9\x61\x62\x52\x89\x8f\xc4\xa0\x7e\xc5\x4b\xc1\x7f\x97\x94\xd7\x0e\x0d\x43\x05\x0d\x3d\x76\x9c\x5c\x82\x8b\xc2\xc8\xd0\x15\x5d\x5c\xdb\x82\xe9\x05\xe9\xd3\x5d\xad\x36\x01\x70\x66\xb7\x2d\x81\xb8\x13\x5e\x21\x7b\x15\x77\xdd\xc4\xf6\x29\x6e\xeb\xc4\x57\x21\xf7\x86\x05\xa5\x27\x1c\x62\x70\xf7\x65\x6c\x2c\xd9\x16\x5e\x02\x66\x0f\xd0\xad\xb3\xf0\xa9\x3c\x63\x23\x0f\x00\x75\x44\x17\x98\x77\x79\xe1\xdd\xfc\xf2\x23\x9a\x4f\x21\x62\x7b\x9b\x9e\x8e\x3d\x2e\x59\xfb\x38\xe1\x8a\xda\x95\xdc\x66\x0a\x35\xca\x92\x0a\xff\x30\xba\xbf\xf3\x15\xdf\x7e\x06\xfa\x43\xeb\xfb\x0b\x10\xdc\x3c\xb1\x45\xfb\x16\xdf\xa9\xcc\x7f\x3c\x84\x99\x66\xb2\x1b\xc3\xd9\x12\x44\x6c\x28\xf7\xf9\xb1\xc8\x81\xb6\x23\x95\x0d\x60\x7a\xd3\x85\x28\x3e\x63\x66\x01\xcf\x0c\x7d\x94\x64\xee\xef\x1d\x90\xd4\xdc\x3d\x83\x39\xa5\x55\x9b\x7f\xca\x16\x5b\x70\x1e\x95\x29\x90\xe6\x73\xee\x93\x95\x11\x3d\x9a\x7a\xb8\x1e\x3f\xb2\xc1\xd3\xd1\x84\xf4\x56\x6e\x7c\x12\x74\x77\xda\xb6\x51\x58\x09\xd6\x1f\xd0\x69\x7d\x40\xcf\x0d\x0b\x19\x5b\x81\x32\xd4\x9c\x4b\x30\xde\x41\x2b\x5d\x15\xa3\xad\xd4\x49\x87\xc1\xe8\xc3\x8a\x41\x85\xbb\xef\xe3\x9b\xaa\x6c\x0c\x12\x93\x04\x4c\x95\xc0\x65\x8a\xef\x71\x06\x04\x0e\xf5\x19\x85\x5d\x22\xcf\x83\xa4\x1c\xaf\xfd\xc5\x90\x51\x61\xc3\x46\xc1\x50\xe5\x50\x62\x94\xd1\x55\xc4\x89\x05\x57\x84\x7c\xa7\x25\xcd\x4a\x8a\xad\x02\x31\x4a\xbb\x48\x37\xe4\xc1\xbf\x01\x48\x27\xe9\x5e\xb1\x1f\x08\x26\x2d\xc5\x5c\xb9\x35\x23\x4a\xb2\xc4\x7d\x20\xbf\x70\x9b\x4e\xc9\xf2\xce\x48\x67\x04\x8d\x75\x52\xa9\xce\x0d\xe0\x90\x54\x30\x57\x55\x02\xcb\x58\xdb\xc0\x89\x61\x01\x7c\xc9\x00\x77\x84\xeb\xf3\xa2\xe3\x18\xae\xe3\x85\x2c\xe1\x4e\x1e\x64\x43\x1f\x0c\x27\xd1\xd5\x05\x09\x3a\xf4\xe8\x88\xc5\x14\x9d\xaf\x9a\xd2\x73\xc0\x53\x73\x3c\x92\xd9\x0c\x55\x21\xf6\xc2\x00\xb8\xc1\xc6\x16\x59\xb3\x9a\x0c\x0e\x0d\xad\xfc\xa8\x81\xa4\x7a\xe9\x2d\xf0\x98\x73\x1e\x83\xe5\x51\xbe\x01\xa5\xab\xc9\x55\xd0\xc8\x6e\x25\x81\x14\xfa\xd2\x1a\x9b\x60\x40\x47\x5b\x1b\xb4\x43\x01\x64\x33\x5f\x31\xfe\x19\x19\xa8\x69\x4d\xf6\xd0\x0b\x76\x2a\x0b\x0e\xa6\xb6\x6c\xd5\xb9\xbb\x73\xcd\xdb\xba\xca\x48\x5b\xe8\x29\x90\xae\xf1\xe0\x7b\x56\xc3\x16\xe9\xb4\xf8\xd3\x9b\x75\xf3\x74\x01\x24\xb1\x6f\x99\x09\xb3\x2e\x35\x95\xfc\x05\x2f\x1f\xf5\xe4\x4d\x47\x4f\xbe\x3f\xfd\xde\xf4\xb8\x9d\xbb\x48\x72\x6e\x45\xd1\xeb\xbd\x6f\xf1\xf8\x2c\x48\x83\x5a\x02\x05\x69\x9c\x7d\xbe\x4b\x20\x00\x49\xc9\xf7\x6a\x82\x6f\xe7\x66\x59\xed\xf6\xf3\xed\xe6\x40\x2e\x08\xb7\x0c\xe2\xf4\xb1\x81\x55\x2f\x4b\xbe\x27\x43\x76\x21\xb8\x55\x16\xe4\x79\x65\x89\x15\x29\x67\xd1\xaf\x80\x7b\x74\xa7\xf5\x01\x4e\x39\xf6\xcf\xfc\x89\x97\x9c\xba\xd3\x88\xfa\x4b\xd8\x70\x77\x9a\x37\x50\x8e\xf6\x08\x91\x7c\x93\x25\x6f\xea\xcb\x0b\x94\x8d\xa5\x4f\x77\xbe\x06\xf2\x5a\x15\xe1\x03\x6a\x2a\xb7\x63\xc0\x71\xc7\xc6\xf7\x20\x5d\xf5\x57\x32\xf1\x5b\x61\x45\x64\x9d\xec\x95\xe6\xcf\x17\xfa\xca\xcf\x6d\xa1\x76\xd5\x8d\x02\x7f\xb7\xd4\x13\x78\xa2\x25\x9b\x88\xfa\xf1\xcf\xf8\x20\xf6\x83\x93\x18\x6c\x02\xc2\x28\x67\xde\x25\xb0\x3c\x8e\xe4\xd2\x11\xdf\x7c\xc5\xb2\x55\x81\xb5\xfa\xa8\x84\x9c\xad\x28\xd8\xe2\xbd\xc7\xef\x58\xfe\xda\x73\x6e\xbc\xae\x4a\x2e\xed\x9c\x09\x19\x23\xcd\x0a\x2f\x8e\x47\x04\xfc\xcf\x40\x27\x17\xa5\x8b\x94\x08\x53\x58\xea\xb0\x23\x01\x4a\x15\x59\x4d\xab\x90\x7c\xee\xef\x03\x7d\xbe\x84\xd0\xbe\x1b\xe5\x5e\x85\x2b\x83\x8b\x84\x78\x0c\x66\xce\xd3\xfa\x8b\xe8\x22\x67\x49\x80\x5a\xcb\xcf\x68\x7c\x93\xba\x4d\x18\xe1\x81\xbe\xb2\xf8\xbd\x27\xc6\x8e\x3b\x86\x94\x87\xc2\x77\x6f\x03\x2b\x11\x26\x98\x11\xad\x97\x5c\x26\x48\xb4\xfb\xa1\x49\x78\xe7\x12\xbb\xb7\x09\x11\x32\x01\xce\xea\x8f\x2f\x6e\xce\x60\x95\xce\x9a\xc5\xf7\x7a\x5e\xf6\x72\xf8\xa9\x33\x1e\x31\x10\x51\x20\x5a\x09\xb9\x79\x89\x03\xcf\xbf\x80\x19\x15\x09\xe5\x69\x86\x03\x04\xc3\x6e\x76\x09\x40\x20\xbb\x71\xd9\xf2\xfd\x82\x15\xc4\x74\xd7\x83\x7c\xc1\xe5\xbb\xa3\x18\x86\x3d\x66\xb0\x85\x71\xb3\x47\xe4\x53\x25\x94\xc8\x99\x8b\xb3\x19\x09\xfc\x39\xfd\x04\x53\x89\xa1\x4a\xfa\x06\x9b\x13\x19\x28\x8b\xdf\xd3\x3f\xc4\xd4\x8c\x3e\xb8\x64\x69\x37\x6f\xbd\x9d\x83\x42\xc3\x80\x34\xe5\x65\x1e\x99\xaa\xbc\x65\x6d\x2f\x56\x94\xf8\xda\x70\x97\xc9\x51\xcd\x55\x87\x08\x06\x6e\x52\xbd\xce\x71\x09\xbc\xfe\x49\x25\x83\x83\x6f\x50\x25\x9a\x64\x8b\xdb\x29\xba\x7a\xc2\xd5\x41\x21\x24\xdc\xc6\x5e\xe0\x65\xb7\xae\xd1\xb5\xf4\xf2\xbe\x2a\x0c\x49\x48\x19\xe3\xbe\x06\x8a\x89\x7c\x14\x9c\xea\x17\x05\xe4\xe7\x89\xed\xa6\x05\x7f\xe1\xe5\x15\x99\x2f\x7e\xbc\x19\xc3\x07\xa8\x86\x16\x68\x1a\x91\x83\x53\x7f\x28\x9e\x39\xdf\xfd\x54\xba\x9c\x6c\x75\x6c\xac\xdf\xba\xd0\x17\xb8\x8c\x42\x47\x34\xf8\x10\x58\x92\x65\x19\xf9\x0a\x9d\x96\xad\xb4\xbb\xdc\x26\x98\x0e\x2d\x5a\xf4\xbc\x8d\xb9\xf5\x83\xcb\x78\xc4\xa0\xbd\x21\xf2\x0a\x9f\xeb\xe3\xe6\x33\x3a\xa9\x7c\x99\xd1\x89\x24\xfe\x52\xe3\x57\x9f\x01\xe0\x5b\x99\xb2\xfc\x73\x6f\x36\x63\x3e\x56\x67\x70\xfb\x27\xdb\xb5\x7c\x3e\xf3\x1f\x89\x2c\x5a\x8b\x46\xa5\x4e\x83\xe6\x4a\x22\xed\xe7\x73\x1e\x0b\x11\x85\x7d\xce\xe8\x77\xe3\x41\xb8\xdd\x6d\x0c\xd0\xe7\xac\xcd\xf3\x0f\xc8\x4f\x56\x53\xe7\x07\x51\xf4\x92\x35\x5f\x9f\x65\x84\xf9\xbd\xdb\x11\xcc\x43\x1b\x82\xe8\xdd\x20\xfb\xd0\x05\x5b\xa2\xc8\xab\xf4\x68\x35\x12\xc1\x06\x3f\x4b\x18\x78\x03\x50\xc2\x21\x0f\x49\xc7\x3e\x1c\xe4\x9d\x26\xd9\x51\xf2\xd2\x98\x0f\xb6\x43\x33\x23\x76\x1e\xeb\x39\x23\x9c\x89\xb0\x68\xa0\x0e\xe7\x45\xde\xd7\x2c\x69\x61\x74\x91\xb0\x70\x19\xb2\xe8\xe5\xd1\x0c\x30\xb6\x93\x83\x13\x23\x8a\xf8\x54\x29\x4b\x38\x5d\x0d\xdd\xf5\xcc\x25\xd0\xc4\x20\x10\x9b\x35\x87\x2e\x64\x6a\x56\x77\x67\x4f\xfe\xce\xf2\xee\xa5\x53\x51\x8b\xed\x6b\xc4\x5a\x63\x6a\x35\xec\x02\xb9\x34\xf6\xbc\x94\x49\x0f\x2e\x56\x06\xe6\x6d\x76\x83\x5a\x6d\x3b\xe3\xc9\xac\x83\xa7\xd3\x60\x87\x12\x3b\xaa\x83\x29\x41\xb4\xce\xc6\xb7\x9e\x97\x50\x06\x95\xad\xa9\xc8\x51\x5b\x6b\x7b\x90\x81\xd9\x96\x87\xd7\xae\x3a\xdc\x1c\x17\xc6\x5d\x1d\x95\xb7\x3e\x71\x5a\x26\x35\x8b\x77\x2a\x82\xcf\xcc\x24\xb3\x0c\xb7\xe6\x7d\xa7\xdd\x8a\xc7\xd0\x77\x66\x11\xd2\x41\x0a\x45\xd1\xd9\x76\xa2\xb7\x00\xac\x25\xbc\x73\xad\x13\x3c\x73\x2d\xcb\xc9\x52\xdf\xe7\x2a\xf7\x5a\xbb\x4e\x1a\x4e\x94\xd5\x3a\x47\xd4\xc4\xa5\xa2\x6a\x95\x49\x58\x14\xcf\x62\xd1\xf5\x77\x05\x51\x1a\x9e\xd9\xa4\x7b\xec\x0a\x40\x7c\xac\xf8\xb6\x4b\x13\xef\x52\xa7\x28\x0b\x56\x12\xcb\x9a\x1a\x65\xeb\xa5\xb7\x20\xab\xde\xcb\xa4\x0c\x18\xd2\x09\x44\x88\x9a\xdd\x1a\x47\xf4\xb2\xea\x4c\x90\xee\x54\xfe\xf4\x4f\x44\xbb\x07\x40\x73\x56\x97\x3d\xf3\x57\x91\xb4\xea\x7b\x7f\x71\xd3\x59\x4d\x1b\x24\x92\x96\x8b\x8b\xe2\x3a\x89\x89\x38\x15\x75\x8d\x39\xee\x70\x90\x79\x6c\xbb\x78\x1f\xcc\xb2\xf4\x1d\x4f\xd3\xc7\xb5\xb3\x45\xc0\x1e\x7e\x34\xa6\xa1\xfa\x2f\x4c\x84\x14\x73\x1f\xac\x36\x7a\x8a\x20\xab\xd6\x99\x3a\x70\x1f\x68\x91\x1f\x72\xee\x27\x06\x50\xb3\xa5\x9b\xfc\xad\xbe\x7e\xf1\x03\x8f\x92\x32\x37\x1f\x8e\x2a\xc4\x00\xdc\x83\x35\xb9\x51\x4e\x07\xff\x1c\x69\x86\x0a\x82\x44\x1f\xe6\xf4\xa0\xae\xe6\x53\xac\x8a\x5a\x57\x68\x93\x19\x5f\x72\x2f\xeb\x5e\xfb\x6d\x02\x67\x74\x8c\x1f\x38\x64\xa1\xbb\xf1\x7f\x81\x7f\x95\x43\x30\xf8\xff\xaf\x22\xff\x0b\x00\x00\xff\xff\x6b\x72\x32\xf6\xab\x0c\x00\x00")

func staticServiceKeyBytes() ([]byte, error) {
	return bindataRead(
		_staticServiceKey,
		"static/service.key",
	)
}

func staticServiceKey() (*asset, error) {
	bytes, err := staticServiceKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/service.key", size: 3243, mode: os.FileMode(438), modTime: time.Unix(1560447831, 0)}
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
	"static/service.crt": staticServiceCrt,
	"static/service.key": staticServiceKey,
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
	"static": &bintree{nil, map[string]*bintree{
		"service.crt": &bintree{staticServiceCrt, map[string]*bintree{}},
		"service.key": &bintree{staticServiceKey, map[string]*bintree{}},
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
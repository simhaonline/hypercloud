// Code generated by go-bindata.
// sources:
// schema/application.json
// DO NOT EDIT!

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

var _schemaApplicationJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x53\x4d\x6f\xa3\x30\x10\xbd\xf3\x2b\x2c\x2b\x87\xdd\x4d\x26\xe4\x4b\xab\x84\x4b\x45\x0f\x3d\x55\x6a\x0f\x3d\x05\xdc\x6a\x1a\x9b\x40\x45\x0c\x35\xce\x21\x41\xfe\xef\x95\x0d\x49\x68\x8a\x2a\xf5\xd6\xd3\xcc\x1b\xfb\xf9\xbd\x99\x81\xda\x23\x84\x0e\xaa\x4d\x2a\x76\x48\x03\x42\x53\xad\xcb\xc0\xf7\xdf\xaa\x42\x42\x53\x1d\x17\x6a\xeb\x73\x85\x89\x86\xc9\xc2\x6f\x6f\x8e\x2c\xad\x54\x45\x29\x94\xce\x44\x45\x03\x62\x1f\x22\x84\x66\xfc\x9c\x13\x42\xf5\xa1\x14\xf6\xd5\x4a\xab\x4c\x6e\x1d\xcb\xd5\x4b\xd4\x5a\x28\x69\x8f\x9e\xa3\x09\xac\x42\xb8\x43\x48\x58\xbd\x34\xd0\x85\x0b\x03\x8b\x2e\x9e\x1b\x88\x96\xab\xf0\x16\x5f\xd9\x75\xb9\x03\xa7\x33\x33\xa0\x4e\xc9\x34\x82\x54\xe2\x4e\xfc\xd8\x56\x08\x6b\x84\xe3\x04\x56\x6c\xf8\xe7\x26\x80\x2e\xfe\xfb\x6f\x70\x21\xed\x32\x79\x2f\xe4\x56\xa7\x34\x20\xf3\x4f\xb2\x1b\x25\x50\x0b\xfe\x82\xba\x47\x3c\x3a\xab\x13\x2a\xf7\x79\x4e\x59\xaf\x8d\x38\xe6\x76\x0c\x36\xcc\x4e\xe1\xa9\x09\x71\x1c\x5c\x25\xd1\x10\x58\x43\xb8\x1a\xc0\xbe\xe4\xbf\xc4\x49\x5a\x54\xba\xea\x35\x81\x4a\xe1\xa1\xc7\x03\xe6\xf9\x43\x62\x6f\xb4\x05\x72\xe6\x7e\xb7\xcb\xaf\xfb\x44\x38\x86\xb0\xb6\xfb\xbb\xa4\xc0\xea\xe9\xe8\xff\xd4\x74\x0e\xe3\x78\xdc\x22\x56\xcf\x46\x27\xfb\xae\x85\x36\x63\xde\x09\xb9\xa6\xa8\x12\xef\xfb\x4c\x09\xee\xba\x70\x9f\x9a\x33\x4f\x91\xf3\x4c\x67\x85\xc4\xfc\xb1\xfb\xa3\x24\x98\x57\xc2\x33\xde\x47\x00\x00\x00\xff\xff\x72\x08\xfd\x8f\x7a\x03\x00\x00")

func schemaApplicationJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaApplicationJson,
		"schema/application.json",
	)
}

func schemaApplicationJson() (*asset, error) {
	bytes, err := schemaApplicationJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/application.json", size: 890, mode: os.FileMode(420), modTime: time.Unix(1515628654, 0)}
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
	"schema/application.json": schemaApplicationJson,
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
	"schema": &bintree{nil, map[string]*bintree{
		"application.json": &bintree{schemaApplicationJson, map[string]*bintree{}},
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


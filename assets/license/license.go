// Code generated by go-bindata.
// sources:
// LICENSE.txt
// DO NOT EDIT!

package assets

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

var _licenseTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcd\x8e\xdb\x36\x14\x85\xf7\x7a\x8a\x83\xac\x92\x42\x9d\xfe\x6d\x8a\x66\x45\x4b\xd7\x16\x01\x59\x54\x49\x6a\x1c\x2f\x35\x12\x1d\x13\xb0\x44\x83\xa2\x67\x30\x6f\x5f\x90\xf6\x60\x9c\x36\x68\xd1\x95\x2f\xac\xfb\x73\xce\xc7\xb3\x52\x25\x7e\xfb\xb1\x38\xf5\x97\xc5\xa0\xb6\x83\x99\x17\x93\x65\x85\x3b\xbf\x7a\xfb\xf5\x18\xf0\x71\xf8\x84\x5f\x7f\xfe\xe5\xf7\x1c\xb5\x3d\x18\xdf\xbf\xe6\xe0\xf3\x90\xb1\xd3\x09\xa9\x61\x81\x37\x8b\xf1\xcf\x66\x7c\xc8\x32\x69\x46\xbb\x04\x6f\x9f\x2e\xc1\xba\x19\xfd\x3c\x22\xae\xb5\x33\x16\x77\xf1\x83\x49\xff\x3c\xd9\xb9\xf7\xaf\x38\x38\x3f\x2d\x39\x5e\x6c\x38\xc2\xf9\xf4\xeb\x2e\x21\x9b\xdc\x68\x0f\x76\xe8\xe3\x82\x1c\xbd\x37\x38\x1b\x3f\xd9\x10\xcc\x88\xb3\x77\xcf\x76\x34\x23\xc2\xb1\x0f\x08\x47\x83\x83\x3b\x9d\xdc\x8b\x9d\xbf\x62\x70\xf3\x68\xe3\xd0\x92\x86\x26\x13\xfe\xc8\xb2\x1f\xf0\xad\xa2\x05\xee\xf0\x26\x65\x70\xa3\xc1\x74\x59\x02\xbc\x09\xbd\x9d\xd3\xbe\xfe\xc9\x3d\xc7\x4f\x6f\xee\x67\x17\xec\x60\x72\x84\xa3\x5d\x32\xe0\x64\x97\x10\x57\xdc\x1f\x9b\xc7\xbf\x29\x19\xed\x32\x9c\x7a\x3b\x19\xff\xf0\x3d\x05\x76\xbe\x27\xf0\xa6\xe0\xec\xdd\x78\x19\xcc\xbf\x89\xc8\x90\x64\xfc\x5f\x11\xb8\x59\x1b\xdd\x70\x99\xcc\x1c\x12\xd9\x0c\x71\xe6\x27\xe7\xe1\xc2\xd1\x78\x4c\x7d\x30\xde\xf6\xa7\xe5\x9d\x71\x7a\x98\x34\x78\x27\x3f\x39\x6a\x8c\x4d\x43\xf1\xe3\xdc\x4f\x26\x8a\x89\xf5\xbb\xe2\xa3\x3b\x8d\xc6\x63\x76\xef\x4d\x09\xbd\x0d\x91\xe2\xe0\xe6\xeb\x42\xe7\x17\x4c\xfd\x2b\x9e\x4c\x8c\xc9\x88\xe0\x60\xe6\xd1\xf9\xc5\xc4\x44\x9c\xbd\x9b\x5c\x30\xb8\xa2\x09\x0b\x46\xe3\xed\xb3\x19\x71\xf0\x6e\x7a\x83\xb1\xb8\x43\x78\x89\x0f\x7e\xcb\x0f\x96\xb3\x19\x62\x80\x70\xf6\x36\xc6\xca\xc7\xe8\xcc\xd7\x10\x2d\xcb\xd5\x82\xae\xb8\x82\x12\x6b\xbd\x63\x92\xc0\x15\x5a\x29\x1e\x79\x49\x25\x56\x7b\xe8\x8a\x50\x88\x76\x2f\xf9\xa6\xd2\xa8\x44\x5d\x92\x54\x60\x4d\x89\x42\x34\x5a\xf2\x55\xa7\x85\x54\xf8\xc0\x14\xb8\xfa\x90\xc5\x0f\xac\xd9\x83\xbe\xb4\x92\x94\x82\x90\xe0\xdb\xb6\xe6\x54\x62\xc7\xa4\x64\x8d\xe6\xa4\x72\xf0\xa6\xa8\xbb\x92\x37\x9b\x1c\xab\x4e\xa3\x11\x1a\x35\xdf\x72\x4d\x25\xb4\xc8\xe3\xd1\xec\x9f\x63\x10\x6b\x6c\x49\x16\x15\x6b\x34\x5b\xf1\x9a\xeb\x7d\x12\xb2\xe6\xba\x89\xb7\xd6\x42\x82\xa1\x65\x52\xf3\xa2\xab\x99\x44\xdb\xc9\x56\x28\x02\x93\x94\x95\x5c\x15\x35\xe3\x5b\x2a\x1f\xc0\x1b\x34\x02\xf4\x48\x8d\x86\xaa\x58\x5d\x7f\xd7\x65\xd4\xfe\x8d\xc7\x15\xa1\xe6\x6c\x55\x53\x96\x2e\x35\x7b\x94\x5c\x52\xa1\xa3\x9d\xf7\xaa\xe0\x25\x35\x9a\xd5\x39\x54\x4b\x05\x8f\x05\x7d\xa1\x6d\x5b\x33\xb9\xcf\x6f\x3b\x15\xfd\xd9\x51\xa3\x39\xab\xb3\x92\x6d\xd9\x86\x14\x3e\xfe\x07\x92\x56\x8a\xa2\x93\xb4\x8d\x9a\xc5\x1a\xaa\x5b\x29\xcd\x75\xa7\x09\x1b\x21\xca\x08\x3a\x53\x24\x1f\x79\x41\xea\x33\x6a\xa1\x12\xad\x4e\x51\x8e\x92\x69\x96\x0e\xb7\x52\xac\xb9\x56\x9f\x63\xbd\xea\x14\x4f\xd0\x78\xa3\x49\xca\xae\xd5\x5c\x34\x9f\x50\x89\x1d\x3d\x92\xcc\x0a\xd6\x29\x2a\x13\x5d\xd1\x24\xab\xba\x22\x21\xf7\x71\x69\x64\x90\xe0\xe7\xd8\x55\xa4\x2b\x92\x11\x68\x22\xc5\x22\x02\xa5\x25\x2f\xf4\x5d\x5b\x26\x24\xb4\x90\xfa\xce\x23\x1a\xda\xd4\x7c\x43\x4d\x41\x51\x8d\x88\x5b\x76\x5c\xd1\x27\x30\xc9\x55\x6c\xe0\xd7\xb3\x3b\xb6\x87\xe8\x92\xe5\xf8\x46\x9d\xa2\x2c\x95\x77\x89\xcd\xd3\x4b\x82\xaf\xc1\xca\x47\x1e\x65\xdf\x9a\x5b\xa1\x14\xbf\xe5\x24\x21\x2b\x2a\x5c\x71\x3f\x64\x7f\x05\x00\x00\xff\xff\x93\x75\x64\x6c\xe8\x05\x00\x00")

func licenseTxtBytes() ([]byte, error) {
	return bindataRead(
		_licenseTxt,
		"LICENSE.txt",
	)
}

func licenseTxt() (*asset, error) {
	bytes, err := licenseTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "LICENSE.txt", size: 1512, mode: os.FileMode(420), modTime: time.Unix(1520776257, 0)}
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
	"LICENSE.txt": licenseTxt,
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
	"LICENSE.txt": &bintree{licenseTxt, map[string]*bintree{}},
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

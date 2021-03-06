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

var _licenseTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x31\x6f\xdb\x30\x10\x85\x77\xfd\x8a\x37\x26\x80\xeb\xba\x59\x3d\xa9\x86\x9d\x08\x50\x65\x43\x92\x11\x78\xa4\xc4\x53\x75\x00\x4d\x0a\xc7\x93\x0d\xfd\xfb\x82\x6e\x8a\x2c\x1d\x32\xf2\x91\xef\xf1\x7b\x77\xbb\x30\x2d\xc2\xbf\x47\xc5\x53\xff\x8c\x97\xcd\x66\xf3\x6d\x12\x8a\xe4\x15\x25\x0f\x24\x66\x59\xa1\xf0\xfd\x1a\xb9\x73\x78\x3c\x8c\x48\xf7\x72\x23\xbb\xce\xb2\x76\xe4\x08\xc7\x9d\x18\x59\xc0\x11\x83\x10\x21\x86\x41\xef\x46\x68\x8b\x25\xcc\xe8\x8d\x87\x90\xe5\xa8\xc2\xdd\xac\x04\x56\x18\x6f\xbf\x07\xc1\x35\x58\x1e\x96\x24\xcc\xde\x92\x64\x3a\x12\x94\xe4\x1a\x11\x06\xa4\xc3\x6b\x75\x46\x49\x31\x92\xe0\x95\x3c\x89\x71\x38\xcd\x9d\xe3\x1e\x25\xf7\xe4\x23\xc1\x44\x4c\x49\x89\x23\x59\x74\xcb\xc3\x75\x10\xa2\xac\xf9\x80\xc0\x21\xcc\xde\x1a\xe5\xe0\xb7\x20\xd6\x91\x04\x37\x92\xc8\xc1\xe3\x65\xfd\xe3\xdf\x4f\x1f\x79\x2b\x04\xc1\x93\xd1\x44\x2e\x08\x53\xb2\x3d\x67\xc6\x2f\x70\x46\x3f\x9d\xff\x69\xfe\x59\xd0\x82\xfd\x23\x73\x0c\x13\x41\x47\xa3\xa9\xe1\x9d\x9d\x43\x47\x98\x23\x0d\xb3\x5b\xa1\x9b\x15\xef\x45\xfb\x76\x3c\xb7\x59\x5e\x5d\xf0\x9e\xd7\x75\x5e\xb5\x97\x2d\xee\xac\x63\x98\x15\x74\xa3\xbf\x39\x7c\x9d\x1c\x93\xc5\xdd\x88\x18\xaf\x4b\x62\xfe\xb5\xaf\x77\x6f\x79\xd5\xe6\x3f\x8b\xb2\x68\x2f\x09\xfb\x50\xb4\xd5\xbe\x69\xb2\xc3\xb1\x46\x8e\x53\x5e\xb7\xc5\xee\x5c\xe6\x35\x4e\xe7\xfa\x74\x6c\xf6\x6b\x34\x44\x5f\x1c\xeb\xf0\xd8\x8e\x50\x66\x49\x0d\xbb\xb8\xce\xfe\x04\x00\x00\xff\xff\xeb\x13\x27\xa8\x29\x02\x00\x00")

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

	info := bindataFileInfo{name: "LICENSE.txt", size: 553, mode: os.FileMode(420), modTime: time.Unix(1594509253, 0)}
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


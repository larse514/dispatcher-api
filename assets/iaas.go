// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ias/cloudformation/SQS.md
// ias/cloudformation/cfn-transformed-template.yml
// ias/cloudformation/index.js
// ias/cloudformation/sam.yml

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _sqsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x41\x6f\xd3\x40\x10\x85\xef\xf9\x15\x43\x2e\x3d\x8d\x5d\xb7\x69\x43\x8a\x7a\x08\x55\x11\x91\x6a\xc0\x31\x3e\x54\x42\xb2\xc6\xbb\x13\xb0\xb2\xeb\x5d\xef\xec\x52\x19\xc4\x7f\x47\x6d\x4a\x0e\x80\x04\xc7\x37\x4f\xf3\xbd\xa7\x19\x7a\x10\x50\xc6\x25\xbd\x73\xc1\x52\xec\xdd\x00\x9e\xd4\x9e\x3e\x33\x20\x46\xb6\xde\x50\x64\xdc\xf5\x86\x41\xc8\x66\x93\x35\x80\xe8\x52\xf4\x29\xfe\x66\xab\xdd\x80\x31\xd0\x20\x8f\x24\xd6\x47\xf7\x79\x47\xce\xb1\x4b\x6a\xcf\x11\x0c\xd9\x4e\x53\x76\x50\x99\x75\x4e\x1b\x9e\xfd\xa5\x87\x66\x6f\xdc\xf4\x47\x8d\x2c\xff\x57\x52\x24\xb5\xc7\x81\x2c\x3f\x47\xa1\x8c\x82\x3e\x38\xc5\x22\x2e\x00\xa2\x22\x4f\x5d\x6f\xfa\xd8\xb3\xc0\xcd\xfa\xc3\xfa\xf5\xe6\x6e\xf3\xf1\xbe\xdd\xac\x4b\x40\xf4\x14\xc8\x72\xe4\x80\xee\x2b\x87\xd0\x6b\x16\xa8\x12\x27\x7e\x47\x96\xaf\x6b\x67\xf9\xa8\x66\xb3\xfb\xf7\xcd\xb6\xad\xab\xba\xad\x9a\xdb\xe6\xb6\x6d\xb6\x77\xd7\x5f\x62\xf4\x72\x95\xe7\x32\x4a\x96\x04\x99\x24\x62\x91\x91\xa5\x6f\x6e\xa0\x07\xc9\x94\xb3\x79\x71\x76\xbe\xb8\xb8\x5c\xbe\x5c\x9d\x16\x67\xb9\x9d\x70\x7c\x24\xbe\x82\x4f\x4f\x77\x90\x51\x40\x78\xd0\x68\x59\xe4\xf0\x89\x27\x1f\x53\x30\xf0\x3f\xf4\x45\xb1\xbc\x2c\x2e\x16\xa7\xab\xd5\x72\x91\xd7\x55\xbd\x2d\xdf\xdc\x6c\x00\x7f\xf1\xb0\x73\x7a\x82\x93\xef\x30\xb7\x53\x79\x18\xcd\xaf\x60\xfe\x96\x8d\x71\x50\xaf\xcb\x17\x73\xf8\x71\xf2\x33\x00\x00\xff\xff\x8e\x89\x60\x53\x19\x02\x00\x00")

func sqsMdBytes() ([]byte, error) {
	return bindataRead(
		_sqsMd,
		"SQS.md",
	)
}

func sqsMd() (*asset, error) {
	bytes, err := sqsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "SQS.md", size: 537, mode: os.FileMode(420), modTime: time.Unix(1531793991, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7b, 0x5d, 0x20, 0x75, 0xaa, 0xcd, 0x16, 0xa5, 0xb6, 0xa4, 0xcd, 0xd4, 0x61, 0x43, 0xef, 0x8, 0x46, 0xa4, 0xc4, 0x8, 0xc, 0xc3, 0x7f, 0x10, 0x1e, 0x28, 0x45, 0xcd, 0x15, 0x6e, 0x6c, 0x1f}}
	return a, nil
}

var _cfnTransformedTemplateYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xc1\x72\xe3\x36\x0c\xbd\xfb\x2b\xb0\x9a\x5e\x95\xb5\xe3\xdd\xa6\xe6\xa5\x23\x37\x4e\xdb\xa9\xdd\xb1\xcd\x6c\x72\x86\x29\xc8\xe5\x54\x22\xb5\x24\x95\xac\xb7\xd3\x7f\xdf\x21\x25\xd9\x74\x2c\x5f\xa2\x93\x08\x3c\x3e\x00\x0f\x00\xd3\x34\x85\x51\xf6\xcc\x1f\xa9\xaa\x4b\x74\xf4\xa0\x4d\x85\xee\x89\x8c\x95\x5a\x31\x48\x6e\xc7\x93\x71\x3a\x9e\xa5\xe3\x59\x32\xba\x27\x2b\x8c\xac\x5d\xf0\xdc\x4b\x5b\xa3\x13\xff\x90\x81\x25\x56\xbb\x1c\x01\x55\x0e\x7c\xc3\xc1\x75\x54\xa3\x35\x1a\xac\xc8\x91\xb1\x0c\x46\xd0\xc1\xb8\x6e\x8c\xa0\x60\x00\x38\x63\x4c\x96\x5a\xa0\xff\x05\x5d\xf4\x9c\x42\xe7\x04\x52\x81\x9d\x26\xe1\xc2\xe3\xa1\x26\x06\xdc\x19\xa9\xf6\x23\x80\x4d\x43\x0d\xfd\x8d\xd5\x30\x9f\x77\x40\xa1\x4d\xc8\xea\xab\x87\x0e\x92\x6c\x75\xe3\xe8\xcb\x76\xc9\x2e\x29\xbe\x6c\x97\x3e\x19\xa1\x95\x6d\x2a\x32\xe0\x34\x18\x8f\x86\x8a\xac\xc5\x3d\x79\xc3\x25\xe5\x96\x6c\x28\xb2\x2d\x7b\x75\xe0\x1b\x1e\x12\x7d\x68\x94\x68\x79\xc3\x9d\xb5\xd1\x35\x19\x27\x3b\x9c\xff\x7e\xd3\x39\x1d\x0f\x00\x7c\x3a\x6f\xc4\xbf\xe4\x18\x94\x41\x8e\x9b\x5d\x38\xde\x54\x5a\xe7\x25\x45\xb0\xbf\xe8\xc0\x60\x37\xbb\xfb\x54\xfc\x8c\x77\xe2\xd3\x9d\x98\x16\xbf\xe4\x45\x3e\xfd\x9c\xef\x6e\x3f\x8f\x71\xb2\x13\x93\xd9\xac\xc3\xff\x81\x2a\x2f\xc9\x30\xa8\x50\xaa\xce\xb6\xd5\x65\x1c\xf7\x57\x48\x1e\x14\x63\xbf\x93\xcb\x9c\x4b\x8e\xe6\x08\x01\x90\x0e\x14\xe6\x69\xce\x20\x99\x39\x46\x68\x94\x93\xbe\x4f\x7b\x3d\xb9\xf9\xd6\x19\x1f\x65\x45\xba\x71\x0c\xa6\x9d\x61\xa1\x5e\xa4\xd1\xaa\x22\xe5\xd8\x91\xe8\x09\x8d\xc4\x5d\x49\x96\x45\xdc\x5b\xda\x07\x25\x3f\xf0\x66\x07\x3f\xfd\x97\x3d\x73\xc6\x5a\xdb\xff\x31\xaa\x6f\x2d\x7c\xd8\x52\x71\x3c\xf6\xe1\x71\x6f\xa3\x9a\xd2\xb8\xbc\xa0\x68\xd2\xca\xce\x84\x21\x74\x94\xcf\x0f\x49\x84\x78\xc2\xb2\xf1\x3d\xcf\x56\xd1\x04\x24\x21\x91\x76\x78\x19\xeb\x75\x49\x06\xc7\x20\x58\x16\x2f\xbe\xd6\xab\x03\x31\xf7\x2b\xc6\xe5\x77\x62\x30\x19\xf7\x22\xf9\x2b\xed\x1a\x65\x46\xbd\xa7\x6d\x5f\x6d\x48\xe4\x4a\xab\xfa\xfc\xa2\xbd\x6a\x15\x2f\xd8\x40\x15\x57\x8b\x8f\xd2\x5c\x61\x5d\x4b\xb5\x1f\x96\x21\x9a\xbd\x01\x01\x32\xeb\x37\xcf\x63\xd6\xba\x94\xe2\x70\xaf\x45\x53\x9d\x24\xf3\x1f\x77\xe8\xe8\x8d\xed\x4d\x37\x01\xb2\x78\xf3\x62\x54\x62\x9d\x65\xa7\x28\xc9\x19\x60\x51\x14\x24\x1c\x83\xac\x2c\xf5\xeb\x99\x67\x6d\xa4\x12\xb2\xc6\xf2\x82\x92\x93\x79\x91\x82\x2e\xec\x3e\x58\xb7\xc6\x58\xe1\x77\xad\xf0\xd5\xde\x08\x5d\x9d\xe6\x3c\x7e\x74\x6f\xd3\xc9\x38\x9d\xdc\xf5\xe9\xac\x50\xe1\x9e\xf2\x56\x83\xcc\xa8\xf3\xb9\x4d\xd0\x28\x86\xaf\x96\x49\xac\x58\xf8\xa9\x03\xf0\xa3\x6d\x73\x49\x8d\x2e\xe9\x63\xf6\xcc\xdb\xde\xcc\xd1\x4a\xb1\xf8\x46\xa2\xe9\xf5\x4f\xde\x4f\xd6\xb7\x73\x80\x2f\x9e\x8a\x3f\xb3\x15\x63\xbd\xeb\x34\x82\xd7\x1b\xff\xf6\x69\x3f\x8e\xe0\xd1\x71\x11\x83\x6f\x38\x63\x9b\xf6\x9d\xff\x11\x00\x00\xff\xff\x07\xc9\x1d\x4a\xd8\x06\x00\x00")

func cfnTransformedTemplateYmlBytes() ([]byte, error) {
	return bindataRead(
		_cfnTransformedTemplateYml,
		"cfn-transformed-template.yml",
	)
}

func cfnTransformedTemplateYml() (*asset, error) {
	bytes, err := cfnTransformedTemplateYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cfn-transformed-template.yml", size: 1752, mode: os.FileMode(420), modTime: time.Unix(1531929316, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0x75, 0xa6, 0x16, 0xc2, 0x44, 0xa6, 0x38, 0x59, 0x36, 0x99, 0x74, 0x32, 0x33, 0xf6, 0x91, 0x6e, 0x5b, 0x54, 0x35, 0xb, 0x90, 0x18, 0x93, 0x26, 0xa4, 0xa, 0x60, 0x7a, 0xc2, 0x2d, 0xef}}
	return a, nil
}

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xb1\x0a\xc2\x40\x0c\x86\xf7\x3e\xc5\x4f\xa7\x16\xe4\xba\x0b\xdd\x1c\xdc\x0a\xe2\x0b\x1c\xd7\x58\x85\x6b\x22\xb9\x54\x2a\xc5\x77\x17\xaf\xd6\xc5\x21\x81\xf0\xfd\x7c\xe4\xf7\xe9\xc9\x01\x97\x89\x83\xdd\x84\x71\xf5\xdc\x47\x52\x54\xf4\x20\xb6\x1d\x82\xb0\xd1\x6c\x35\x96\x02\x00\x9a\x06\xe7\xee\xd0\xed\x71\xcc\x39\x8c\x94\x92\x1f\xc8\x39\x97\x71\x10\x4e\x06\xa5\x20\xda\x27\xb4\xc8\x12\x77\x5a\xef\x5f\x42\x22\xb9\x28\x43\x55\xd2\x6c\xea\x11\x65\x28\xeb\x3f\xf8\x95\xac\x20\x2f\x25\x9b\x94\xb1\xbc\x0a\xe0\x33\x05\x30\x4a\x3f\x45\x72\x34\xdf\x45\x2d\xb9\xed\xf9\x76\xab\xf1\x0e\x00\x00\xff\xff\x39\x06\x8c\xea\xdd\x00\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 221, mode: os.FileMode(420), modTime: time.Unix(1531796665, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x2, 0x84, 0x6c, 0x55, 0x3a, 0xe9, 0xee, 0x72, 0xd8, 0x23, 0xa3, 0xf6, 0x7a, 0x61, 0xa9, 0x46, 0xf, 0x31, 0x6, 0x47, 0xd2, 0x12, 0x6e, 0x5d, 0xd4, 0xce, 0xb9, 0x48, 0x41, 0xca, 0xa8}}
	return a, nil
}

var _samYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x3d\x73\xe2\x30\x10\xed\xfd\x2b\x96\x8a\xca\x8c\x0c\x33\x37\x87\x3a\xdf\x1d\xdc\x15\x5c\x06\x23\x12\x6a\xc5\x5e\x83\x32\x96\x64\x56\x32\x81\xfc\xfa\x8c\x65\x87\x8f\x90\x22\x5b\x79\xdf\xbe\x7d\xfb\xfc\xec\x74\x23\xd6\xa8\xeb\x4a\x7a\x9c\x5b\xd2\xd2\x3f\x21\x39\x65\x0d\x87\xe1\x98\x25\x2c\x66\xd3\x98\x4d\x87\xd1\x9a\xa4\x71\xa5\x25\xcd\x21\xdd\x08\xce\x05\xd2\x01\xa9\x42\xe7\xe2\x31\x4b\x7e\xc4\x09\x8b\x27\x49\xf4\x07\x5d\x4e\xaa\xf6\x61\x7f\x76\x94\xba\xae\x10\x6c\x09\x35\xd9\x1c\x9d\x53\x66\x0b\x1a\x9d\x93\x5b\x74\x60\x0d\x48\x03\x22\x13\xb0\x6f\xb0\x41\x78\x55\x7e\x07\x0b\xa9\x9f\x0b\x19\x2d\x25\x49\x8d\x1e\xc9\xf1\x08\x20\x6b\xe7\x0f\x52\x63\xdb\x00\xac\x4f\x35\x72\x10\x9e\x94\xd9\x06\xe0\xe6\x6a\xcb\x83\xd2\xd2\x45\x39\x82\x5e\x56\xd8\x86\xf2\x6f\x89\x2c\x6c\x2e\xdb\xa7\xd6\x7b\xb7\x0b\xb9\x2d\x10\x94\x01\x37\x89\x56\xe8\x82\x52\x30\xf7\xff\x24\x32\x11\x1c\xce\x1b\x93\x87\xf5\xab\x03\x9f\xb2\xe2\xfc\x83\x14\x38\x4b\xb2\x35\x92\x57\x9d\x52\x5b\xbf\x6d\x81\x8f\xa4\x38\x0c\x56\x58\xde\xd8\xee\x09\xff\xa4\x29\x2a\x24\x0e\xca\x14\x78\x1c\xed\xba\xb6\x1f\xae\x1a\xe3\x95\x46\x0e\xc6\x16\xf8\xe2\x7e\x8e\x12\xd6\x4f\x66\x07\x34\xfe\x7c\xa5\x77\x1d\xc0\x0b\x76\x0e\x25\x13\x57\xd8\xbd\xc7\xae\xc2\x2b\x73\x18\xfc\x45\x9f\x7a\xdf\x0a\xee\x5d\xc0\x46\x29\x99\x1b\xe6\x2f\xe9\xf3\x9d\x50\x6f\xc8\x21\x61\x51\x17\x59\xcf\xbd\x8f\x2a\x13\x9c\x67\xfd\x57\xfb\xea\xf8\xe5\x5f\xe8\x22\x3a\xf7\xef\x01\x00\x00\xff\xff\x69\x13\x77\x43\xc8\x02\x00\x00")

func samYmlBytes() ([]byte, error) {
	return bindataRead(
		_samYml,
		"sam.yml",
	)
}

func samYml() (*asset, error) {
	bytes, err := samYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sam.yml", size: 712, mode: os.FileMode(420), modTime: time.Unix(1531685283, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4f, 0xf7, 0xa3, 0x63, 0xad, 0xfb, 0x1e, 0x24, 0x6c, 0xd7, 0x9f, 0x2f, 0x8, 0x1, 0x66, 0x4, 0x6b, 0x2d, 0x1d, 0x24, 0x25, 0x7b, 0x4f, 0x63, 0x43, 0xa9, 0xcc, 0x2e, 0xf3, 0x7e, 0x63, 0xa7}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"SQS.md": sqsMd,

	"cfn-transformed-template.yml": cfnTransformedTemplateYml,

	"index.js": indexJs,

	"sam.yml": samYml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"SQS.md":                       &bintree{sqsMd, map[string]*bintree{}},
	"cfn-transformed-template.yml": &bintree{cfnTransformedTemplateYml, map[string]*bintree{}},
	"index.js":                     &bintree{indexJs, map[string]*bintree{}},
	"sam.yml":                      &bintree{samYml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

// Code generated by go-bindata.
// sources:
// pkg/manifest/data/pkg-manifest.schema.json
// DO NOT EDIT!

package manifest

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

var _pkgManifestDataPkgManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x6f\x73\xdb\x36\xd2\x7f\xef\x4f\x81\x51\x33\x53\xf9\x49\x24\x39\x7f\x9f\xab\x3b\x1d\x4f\x26\x4d\xe7\x7a\xd3\x26\x99\x4b\x93\x9b\xa9\xa5\x74\x20\x72\x25\xa1\x26\x09\x1e\x00\xca\x71\xef\xfa\xdd\x6f\x00\x90\x14\x09\x02\xa4\x48\x91\xb5\x92\x38\x6f\x62\x13\x58\x10\xfb\xc3\x0f\xd8\xc5\x62\x09\xff\xe7\x04\xa1\xd1\x3d\xee\x6d\x20\xc4\xa3\x73\x34\xda\x08\x11\x9f\xcf\x66\xbf\x73\x1a\x4d\xf4\xd3\x29\x65\xeb\x99\xcf\xf0\x4a\x4c\xce\x9e\xcc\xf4\xb3\xaf\x46\x0f\xa4\x1c\xf1\x33\x11\x7e\x3e\x9b\xd1\x98\xc7\xe0\x4d\x09\x9d\x9d\x4d\x1f\x4e\x9f\xce\xe2\xab\xf5\x24\xc4\x11\x59\x01\x17\xd3\xb4\x2d\xd9\xae\x96\x15\x44\x04\x20\xc5\x63\xec\x5d\xe1\x35\xfc\x9c\xd6\xd4\xa5\x3e\x70\x8f\x91\x58\x10\x1a\xc9\x3a\x59\x21\x5a\x51\x86\x30\x4a\x45\x74\xd5\x98\xd1\x18\x98\x20\xc0\x47\xe7\x48\xaa\x83\xd0\x28\xc2\x21\xe4\xbf\x55\x9b\x7b\x85\x43\x40\x74\x85\xc4\x06\x10\x8d\x55\x33\xaa\x9a\xb8\x89\x55\x97\xb8\x60\x24\x5a\x8f\xd4\xe3\x3f\x75\xa9\xd1\x84\xab\xe5\xef\x77\xbf\xb6\x7d\x01\x89\xe2\x44\xf0\x62\xdb\xf7\x18\xac\x64\xed\xaf\x66\x3e\xac\x48\x44\x64\xab\x7c\x16\x63\x86\xc3\xb2\x28\x4d\x44\x67\x59\x96\x44\xcd\x72\xdc\x33\x3a\xbb\x05\xc6\xeb\x91\x78\xaf\x6b\xd8\x50\x70\xbc\x03\xc2\xf7\xc0\xd2\xd7\x9c\xa4\xaf\x1a\x31\xf8\x77\x42\x18\x48\xa2\x5d\x16\xc6\xf6\x04\xa1\x85\x2a\xc7\xbe\xaf\xe4\x71\xf0\xa6\xc8\x83\x15\x0e\x38\xa4\x4c\xca\x5f\xb1\xe3\x87\x4f\xd8\x1b\x05\x45\xa1\xff\x39\x21\xf3\xc2\x07\xae\x41\x26\x0c\x3c\x41\xd9\x0d\x52\x78\x82\x00\x26\xd5\xc4\x91\x75\xac\xe9\xf2\x77\xf0\xc4\xee\xb9\x85\xaf\x59\x9f\x4a\x0f\xdc\x55\x6b\xf8\x98\x17\xdb\x98\x96\xfd\xfb\xf3\x81\xd9\xd4\x0a\x27\x81\xb0\x35\x53\xa1\xb7\xaa\x89\xb6\x38\x48\xe0\x5b\xa9\x31\x5e\x72\x1a\x24\x02\x50\x8c\xc5\x06\xad\x18\x0d\x11\xa3\x54\x48\x3c\xe2\xab\x35\xa2\x0c\x31\x08\xb0\x20\xdb\xb4\x06\x89\x04\xb0\x98\x81\x00\x5f\xd7\x96\xe4\xf0\x09\x43\x24\x42\xd7\x1b\xe2\x6d\x52\xb6\x20\xc2\x91\xa4\xe6\x83\x43\x14\x23\xfc\x2d\x78\x0c\xf6\xd1\xec\xc7\x55\xd6\x93\x74\x68\x09\x47\x5c\x0b\x3b\xbb\xb0\xa4\x34\x00\x1c\x19\x7d\x38\x71\xf4\xa7\x96\xaa\x27\xa6\x78\x2e\xda\xc8\x70\x55\xa9\x32\x4d\x50\x46\xa9\xf4\xb7\x45\x69\xfa\xae\x48\x00\xee\x09\xb0\x2b\x75\xcd\x80\x1f\x48\x00\xbd\x92\x5f\xbe\xf2\x8e\xfd\xb7\xcd\x7e\x39\x0a\x9f\x05\xf1\x15\x9d\xac\xcc\x8f\x93\x20\x78\xc1\xc0\x2f\x19\x4b\x17\x5b\x0d\x94\xa4\x1c\x44\x82\xe0\x80\xa3\x84\x83\x8f\xfc\x44\x8e\x02\xc2\x89\xd8\xc8\xe7\x1e\x56\x76\xff\x9a\x08\x3d\x8e\x9c\x26\xcc\x83\x74\x76\x90\x10\xaf\x41\x32\xa2\xe8\xbc\xa0\xba\x39\x91\x70\x60\x86\x23\x83\xdc\xe6\x53\x53\x8b\x06\x58\x80\xff\xd6\x60\x47\x61\x38\x46\x31\xe6\xfc\x9a\x32\xbf\x87\x56\x2b\x23\x66\x1f\x8c\x5c\x0f\x5b\x2f\xb2\x41\xda\x6b\xd4\x4b\x23\x19\x25\xe1\x12\xd8\x0b\x1a\x71\xc1\x30\x89\xca\xee\x4f\xbe\x96\x55\x6b\x75\x5e\xa3\x70\x10\xbc\x5e\x99\xb0\x19\x14\xf9\xc7\xdb\xd7\xaf\xd0\x5b\xe5\xeb\xa2\x4b\x25\x80\xae\xe0\x46\x6a\xba\x18\x67\x7e\xb2\xa0\x34\xe0\x53\x02\x62\xa5\x5c\xeb\x8d\x08\x83\xd4\xbf\xbe\x66\x64\xbd\x11\x93\x82\xf3\x3d\xd9\xe2\x80\xf8\x8a\x57\x93\xb3\xb3\xaf\x38\x78\xea\xc7\xa7\xd3\x47\x8f\x4e\x4b\xf3\x33\xd7\x06\x33\x86\x6f\xca\x45\x44\x40\x68\x59\x43\xed\x23\x5e\x05\xac\x38\xaf\xad\x94\xc2\xd1\x4d\x4b\x60\xa4\xc0\x50\xc0\x3c\x3e\x1e\x60\x20\x4a\xc2\x36\xb8\xc8\xfa\x43\xc1\x72\x76\x30\x2c\x99\x84\x06\xa2\x59\xfb\x15\x65\x21\x36\xad\xcf\x88\x46\xa0\xc8\x72\x59\x6a\xbc\x6a\xc1\xb3\x09\x2c\x57\xa0\x35\xb0\xaa\x31\x32\x80\xfc\xa7\x5e\x7b\xb8\x5a\x79\x75\x17\xd1\x12\xd4\xca\xeb\x6a\xc1\x30\xa6\x95\xf2\x74\xf8\x2e\x8d\xe7\x68\xd7\x29\xa3\x64\xe1\x34\x88\x0b\x2b\x42\x21\xfe\x48\xc2\x76\x14\x49\x45\x86\x62\x89\x83\x24\xe6\x90\x97\x94\x20\x51\x6b\x25\xb4\xc8\x40\x4a\x3c\xe9\xa2\x44\x12\x08\x12\x07\xd0\x6e\x1d\xdb\x49\x0d\xa4\xca\xc3\x0e\xaa\x44\xb4\x32\xe7\xea\x74\x88\xa8\x18\x8a\x4c\x4f\x8d\xde\xb7\x5e\x57\x8b\x6a\x65\xeb\xc6\xde\x8a\x29\x81\xa1\x54\x73\x71\xec\xaf\x32\x32\xad\x9c\x65\x8b\xdb\xe4\xde\xfc\x15\xcb\x5d\x8e\xf0\x2b\xbd\xbc\xf6\xb9\x01\x4c\x19\x7d\x74\x5b\x40\xb7\xd1\xeb\x63\xc3\x95\xda\xa9\xce\x5b\x2e\xa3\x03\x9e\xd5\x19\xce\x8b\x3b\x50\xee\xb8\xb6\x75\xe5\x41\x28\x6f\xec\x34\xd1\x9a\xb6\x03\xd5\x5a\x9d\x19\x1b\xe2\x8f\x6f\x5c\xf4\x6c\x32\xe1\x3b\xc1\xa1\x0c\x87\xcb\x0b\xb6\x39\x43\x05\x13\x7e\xe6\x32\xf1\x1d\x55\x2d\x0a\x0e\xa5\xaa\x6b\x2d\xee\xa4\x6a\x81\x78\x7b\x6b\x99\xc9\x0c\xa5\xa0\x69\x47\xbb\xbb\xee\xd5\xc5\xd0\xee\xba\x3b\xd7\xdd\x5a\x1c\xe2\xc1\x87\xfa\x99\x03\x09\x63\xce\xa2\x9a\x75\xc6\x61\x86\xf7\x3d\xf7\x6a\x86\x0e\x0b\x01\xac\xe3\x7c\xa9\x08\x0f\x05\xe4\xff\x1f\x3f\x90\x7b\xbc\xb6\x3e\xc8\x60\x91\x1f\x0a\xce\xbf\xf5\x0e\xa7\x7d\x9b\x5c\xdd\x28\x37\xb9\x08\x15\x27\xc1\xda\xc4\x01\x83\x87\x0c\x1f\xa1\xbc\x0b\xb6\x0f\xad\x0f\x31\x44\x3e\x44\x5e\xcb\x21\x2d\xca\x0d\x35\x94\xdf\x1c\xc3\x50\x5a\x5e\xd6\xd3\x48\xed\x41\x06\xb7\x79\x49\x2b\xd8\x8d\x8c\x21\x6d\xf3\xbb\x51\x85\x2b\xed\xb9\xf3\x59\xc6\xd4\x4c\x6a\x8c\xa2\x24\x08\xaa\x9b\x81\x94\x16\x2d\x11\xfb\x32\xe2\xd6\x55\xcf\xbe\x19\x98\x2f\x22\x6e\xdd\x01\x98\x2f\x23\xd6\xd2\x01\x98\x4f\x2d\xb6\x56\xa3\xe2\x01\x21\x24\xdd\xaa\x3b\x84\x54\x2c\x77\x85\x90\x5e\xab\x3a\xbd\x86\x90\xd2\x7a\x47\x1b\x42\xb2\xad\xde\x87\x87\x90\x74\xab\xb7\x1b\x42\xaa\x9d\x49\xc7\x15\x42\x2a\x0f\x42\x39\x84\xc4\xa9\x77\x05\x35\xbc\x2e\x96\x37\xb2\xd4\x18\xad\xb7\x4a\xb6\x96\xef\x2e\x5e\xeb\xd7\xde\x12\xaf\xdb\x13\x52\x77\xf7\xb3\x48\x23\x49\x91\xb7\x93\x45\x81\xd7\x14\x6f\xac\xd6\xba\x4b\x3f\x48\x1f\x3b\xb2\x40\x2b\x80\xdd\xb9\x71\x5d\x81\xf9\x2c\xb7\x4a\x7b\xc7\x30\xed\xe9\x07\xe6\xb2\x9c\xc4\xc0\x38\xa8\x8c\xbd\x12\x16\x5a\x7a\x10\x34\xcc\xd8\x5b\xdb\x8c\x08\x1f\x0b\x98\x08\x52\x4a\xa8\xda\x63\x74\x73\x31\xa4\x75\xeb\x57\xa7\xe9\x63\xf3\xb0\xda\x36\x68\x2d\x52\x2c\x76\x5a\x1a\x65\x8b\x3a\x7b\x55\x83\x9a\x5c\xcb\xd9\x44\xa5\xe1\x4d\xe4\x0c\x6b\x02\xef\x39\xd2\x22\x69\xe6\x1e\x83\x15\x30\x88\x3c\x40\x98\x23\x35\x31\xc1\x47\xcb\x1b\x74\xb9\x26\x62\x93\x2c\xa7\x1e\x0d\x67\x5a\x60\xe6\x13\xa9\xee\x32\x91\x2d\xcd\x72\xb9\x1d\xde\x0d\x12\x82\x01\x64\x05\x0f\xa7\x0f\x1f\xef\x9a\xe8\x17\x60\x13\x90\x7e\x70\x86\x10\x13\x4b\xe8\xa2\x76\xdd\x91\x22\x43\xb1\xf2\x51\xaf\xa0\x69\xed\xfa\x41\x6a\x43\xb9\x30\xf2\x22\xf7\x00\x2b\x93\x1a\x0a\xaf\xc7\xbd\xe2\x95\xeb\xd8\x0f\x64\x24\xde\x3e\x69\x07\x97\x94\x18\x0a\xaa\x27\xbd\x42\xa5\x74\xeb\x0d\xa6\x67\xad\x61\x7a\x36\x14\x4c\x4f\xfb\x86\xe9\x59\x4f\x30\x25\x8c\xb4\x43\x29\x61\x64\x28\x90\x9e\xf5\x0a\x92\xd4\xac\x1f\x8c\x38\x84\xdb\x3d\xb2\x2f\x9f\x23\x0e\x21\x8e\x04\xf1\x50\xfa\x69\x98\x69\x26\x75\x43\x12\x23\x8d\xdd\xf9\x6c\xb6\x7b\x34\xeb\x55\xfb\xb4\xcf\xf5\x00\x9c\xd8\x4a\x8c\xe4\xcc\x9f\x20\x5a\x8b\x4d\xcb\x84\x07\x2d\x34\x90\x1f\xed\x3a\x00\x6f\xc8\x75\x78\x68\xd7\x30\xeb\xec\x31\x69\xe8\x3a\x99\x6e\xca\xe6\x78\x50\x56\x20\x8b\xc6\xd9\xb3\x3c\x3e\xb5\x38\x6e\xcd\xe6\xef\xcb\x8b\xdb\x77\xd8\x09\xa7\x79\x15\x1d\x52\x31\x06\x02\xc7\x95\x30\x60\x59\xec\x76\x1b\xd9\x11\x83\x35\x7c\xec\x25\x84\xaf\xdf\x53\x13\xea\x2c\x94\xb7\x0e\x75\x2a\xd9\x6e\xa1\x4e\xad\xfe\xd1\x86\xf0\x07\x8a\x98\x6a\xc0\x6e\x35\x84\x5f\x3b\xa9\x8e\x2c\x2a\x5b\x1a\x04\xe3\xf3\x3e\x93\xd1\x06\xe2\x6f\xba\x9c\x37\xd5\x66\x74\x8d\x2e\x27\xbf\x4d\xf1\xe4\x8f\xe7\x93\x5f\xcf\x26\xdf\x2c\xee\x77\xfc\xde\xc5\x3e\x2a\xf9\xa7\xe8\xad\x9c\x37\x7b\x5b\xbb\xaf\x7a\x7b\x68\xac\x98\x25\xde\x43\x73\xc5\x13\xc3\x1e\x9a\x2b\x1e\xd4\xf4\xd1\x5c\x61\x31\xdc\xc7\x89\x3c\x60\x59\xf6\xd6\xf6\xe5\xd8\x5b\xd7\x2c\xb7\x58\xba\xdd\x1e\x0e\x02\xb4\x66\x38\xde\xe4\xd4\xfe\x16\x71\x00\x94\xd9\x2d\x88\xa6\xd7\xe4\x8a\xc4\xe0\x13\x9d\x2b\x24\x7f\x9b\xbd\xc0\x41\xf0\x9b\x12\xdb\xbd\xa0\x4a\xd9\x3d\xbe\x14\xf5\x68\x24\x30\x89\x80\xc9\x16\xad\xee\xc9\x1e\x8d\xd0\xf8\x10\x69\x39\xf9\x83\x00\x82\x43\xda\xe0\xc0\x08\x36\x5b\xb0\x2e\x35\x65\x85\x6d\xa3\x56\xae\xd1\xf9\xe0\x27\x6f\xa6\x8d\x5d\xf4\x42\x33\x85\xd9\xc6\x9d\x17\x34\x0c\x71\xe4\x23\x96\x44\x72\x93\x86\x51\xfe\xae\x6f\x11\xdd\x02\x63\xc4\x07\x8e\x70\x74\x83\x38\x08\x84\x85\x32\x57\x3a\x1e\x1a\xc0\x16\x2c\x71\xbe\xfa\xe4\xb1\x9a\xd4\xb1\x0e\x1f\x20\x97\x47\xa7\x32\xda\x5a\x61\xc2\xac\x56\xb0\x21\xd3\xce\x75\xe5\x07\x01\x8e\x48\xa4\x60\xd8\x0d\x4b\x45\xb8\x29\x07\x38\xad\xf6\x61\x7c\xa9\x0d\xc7\xe2\xfc\xf4\x42\x9a\x91\xf9\x7c\x56\xb0\x24\xf7\x5c\x29\x76\xae\x84\x42\xfd\xcf\x26\x62\x53\x69\x7c\x4d\x82\x00\x2d\x01\x2d\x69\x12\xf9\x48\x50\xc4\x71\x98\xdf\x53\x90\x7d\xa6\x5e\xdd\x98\x57\x20\x54\xb9\x6a\xd6\x4a\x95\x6c\xc3\x76\x3d\xf4\x09\x43\x0c\x56\xfa\xa3\xf9\x52\xaf\x9a\x3b\xe5\x4a\x3f\x44\x96\x14\x44\x54\x89\x0e\xd8\x6a\x55\x94\xd9\xcf\xf5\xb1\x88\x8e\x20\xda\xbe\xc7\xbd\xf0\xf2\x65\xb4\x25\x8c\x46\x21\x44\x02\x6d\x31\x23\x78\x19\xf4\xca\xd0\xcb\x0f\xdf\xdd\x02\x11\x49\x84\xb8\x47\x63\x75\xe8\x82\xae\x67\x9a\x98\x11\x0e\x6f\x95\x8d\x9a\x51\xa9\xe7\x93\xf1\x32\x75\xdd\xd5\x8d\x20\xee\xbe\x75\x5d\xda\xd2\x7e\xdf\x36\x5d\xa5\xe7\xd8\x07\x59\x7f\x20\xfd\x92\xf3\x6e\xf9\xac\xed\xa1\x1c\xb6\x99\x95\xb5\x9f\xd7\x6a\xaa\xfc\x91\x2e\xf4\xac\xf1\x9f\x90\xde\x7e\x9a\xc1\xb4\xbc\xa8\x72\xf9\x40\x76\x38\x2c\x68\x7e\xd9\x8b\x15\xd6\x03\xd6\x02\x0b\x1d\xac\xd7\xd9\xec\xf1\xba\x9d\x58\x87\x81\xb2\xee\xca\x4b\xa0\x19\x4f\x17\xfd\x8d\xb5\xe5\x3e\x1c\x64\x19\x0f\x75\xc1\x5f\x69\x95\x41\x1e\x8e\xe4\x6c\xcd\xcf\xd2\xd5\xa9\x84\xba\x83\x89\x8a\x8d\x2c\xcf\x6a\xf2\xc6\xa3\x88\xda\x1e\xc6\x94\xd9\xe3\x2e\x66\x24\x42\xd6\x4b\x17\x8f\xfc\x2a\xa8\x5d\x77\x05\x55\x0f\x36\x94\xd7\x84\x83\x9c\xc4\xde\x6f\x01\xbd\x54\xeb\xe4\x78\xa2\xff\x3f\xbd\x18\x0b\x2f\xfe\x6f\xe2\xc7\xa7\x17\x7b\xd2\xfe\xef\x94\x0b\x24\x15\x1e\xf3\x53\xd9\xe3\x25\x51\x2b\xa1\x9d\xf8\x0d\xe7\x39\xa8\x1c\xa7\xad\x74\xae\x0b\x53\x3b\xd3\x4c\x07\x10\x3a\xd9\xbc\x7d\xb1\x3f\x77\xc7\x8c\xf2\x4a\x95\xdd\x5a\xc6\x8e\x34\xc7\x12\xfb\x3e\x03\xce\x51\x88\xe3\x18\x94\x0d\xc2\x59\x91\x2d\xc3\x05\xed\xb3\xa4\x0f\x89\xaa\xf0\x5f\x32\x73\x17\xdb\x27\xa8\x1f\xa6\x6e\xab\xef\xc6\x52\xf8\xc0\x18\x8a\x19\xac\xc8\xc7\x32\x94\xda\xb9\x3b\x52\x28\x5f\x27\xfb\x64\x9b\xff\xd5\x50\xd2\x44\x7c\x62\x50\x5e\x53\x76\xf5\x7d\xe5\xee\x4d\x9b\xa2\xff\xa2\xec\x4a\x6a\xe1\x17\xee\xff\x14\x1b\x34\x2e\x47\x4a\x0a\x87\xde\xca\x0d\x68\x3e\xda\x3e\x71\x69\x5a\x8e\xb4\x3b\xed\x6e\xea\x01\x15\x9e\x2d\xfa\x08\xd1\xdb\xa3\xef\x3b\xbf\xfd\xc4\x78\x57\x8b\x0f\x48\x62\x67\xcc\x2c\x2d\xea\x1c\x2c\xa3\xb1\x19\x25\xab\xfb\x80\xb3\x2e\x82\x16\x5f\x99\x07\x51\x4d\xcd\x35\x35\x89\x0e\x71\x2a\x91\xec\x4f\x67\xc3\xba\x3b\x43\x4c\x18\x99\xe4\x4e\xd0\x9d\x93\x69\x79\x7b\xf5\x26\xe6\xbc\xc4\x3c\xbf\x93\xaa\xe2\x80\xfc\x01\x1c\xfd\xf8\xea\xcd\xbb\x5f\x7e\x7b\xf5\xfc\xe7\x97\xda\x9d\x7b\xff\xfc\xa7\x77\x2f\xe5\x26\x2b\xcd\x24\xfe\x7a\x57\xe1\x5c\x17\x7e\x3d\x45\x3f\xae\xb2\x7a\x1c\xc9\x7d\xe0\x03\x44\x04\xfa\xf9\xdd\xdb\x5f\xd4\xd5\x65\x9c\x27\x21\xf8\x69\x8d\xef\xd0\xbd\xf1\xae\x89\x9a\x45\xe5\x50\xc7\xa4\xf6\x38\x2b\xaf\xd6\x71\x13\xdd\xff\xa6\xf7\xc0\xd8\x4a\xcb\xb7\xd5\x5d\x07\x74\x68\x9b\xb6\xef\xc3\xf2\x36\x6f\x7b\x93\x5d\xbd\x62\x3c\x2f\xaa\x99\x14\xbb\xe9\xf0\xfa\xdd\x2f\xf9\xfc\x28\x4c\x0a\x3d\x1d\x0a\x85\x7a\x52\x94\x6a\xd7\x4c\x0d\x55\x41\xce\x8c\x82\xc0\xdd\xd4\x30\x5b\x3c\xae\xe8\x8d\xd3\xbf\xd9\xf3\xa4\xbc\xde\x32\x48\x2b\x69\x5a\x86\x3a\x62\xb7\x6c\xbe\x39\x4f\xfe\xaf\xf1\xc2\x68\x7c\x80\xfb\x55\x3a\x27\xb5\x39\x61\xa5\x0a\x9d\x5d\xb1\xac\x15\x97\x43\x76\x70\x1e\x98\xe7\xfa\x0a\x68\x4f\x10\xf3\x0e\x76\x87\xb2\x70\x5c\x6c\x3d\xb7\xdf\x15\x77\x86\x51\xb7\x71\xb4\x20\xa6\xdd\xeb\x0e\xa1\xc5\x54\xdb\xa0\xb4\x54\x6b\x48\x40\xe3\x49\x1c\x53\x26\xe4\x8f\xf7\xc6\xca\x32\x20\x9e\x2c\xb9\x20\x42\x7d\x67\x83\xf8\x4d\x24\xf0\x47\x74\xbd\x01\x06\x79\x8d\x2c\xaa\xcf\x20\x0e\x70\x1a\x25\x14\x1b\xd0\x27\x3c\x88\xae\xb4\xc9\xc1\x02\xb1\x24\x2a\x7d\x83\x55\xfb\xb7\x3d\xd2\xbf\x6d\x61\xa7\x88\x2a\x72\xe9\xb2\x7d\x34\x3d\x9b\x9e\x55\xb3\xab\xc7\x59\x6a\x47\x39\x8f\x9a\xc7\xe0\xcd\xb4\xcc\x74\x23\xc2\xe0\xd4\xd9\x3f\x33\xbd\x49\x16\x7d\x18\xa7\xd1\xb6\xf9\x7c\x6a\xf9\x71\x7c\x71\x3e\x9e\xcf\x55\x44\xee\xf9\xe4\x57\x3c\xf9\x63\xb2\xb8\x3f\xbe\x38\x9f\xcf\xa7\xa5\x47\xa7\xff\x77\x7a\x7a\xa1\x9e\xdf\x2f\x3c\x9f\xcf\x27\xf3\xf9\x74\x71\xff\xf4\xe2\x5e\xe1\x6f\x7c\x9c\xfc\x79\x72\xf2\xbf\x00\x00\x00\xff\xff\x07\xaa\x9d\x89\x83\x66\x00\x00")

func pkgManifestDataPkgManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_pkgManifestDataPkgManifestSchemaJson,
		"pkg/manifest/data/pkg-manifest.schema.json",
	)
}

func pkgManifestDataPkgManifestSchemaJson() (*asset, error) {
	bytes, err := pkgManifestDataPkgManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/manifest/data/pkg-manifest.schema.json", size: 26243, mode: os.FileMode(420), modTime: time.Unix(1502761439, 0)}
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
	"pkg/manifest/data/pkg-manifest.schema.json": pkgManifestDataPkgManifestSchemaJson,
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
	"pkg": &bintree{nil, map[string]*bintree{
		"manifest": &bintree{nil, map[string]*bintree{
			"data": &bintree{nil, map[string]*bintree{
				"pkg-manifest.schema.json": &bintree{pkgManifestDataPkgManifestSchemaJson, map[string]*bintree{}},
			}},
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

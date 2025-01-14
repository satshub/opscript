// Code generated for package internal by go-bindata DO NOT EDIT. (@generated)
// sources:
// spec.json
package internal

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

var _specJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5d\x5f\x53\x23\x39\x92\x7f\xdf\x4f\x91\x31\x2f\x07\x11\xc6\x63\x43\x43\xc3\xdc\x93\x1b\xcc\xe0\x68\xc0\xac\x6d\x7a\xe7\x6e\x6f\xa3\x43\x55\x25\xbb\x14\x94\xa5\x5a\x49\x85\xcd\x5e\xcc\x77\xbf\x90\x54\x65\x97\x41\x92\x51\xd9\x4d\x5f\xec\x3c\x4c\xcc\x38\xe5\x72\xfd\x32\xa5\xfc\x9f\xe2\x7f\xff\x02\x00\xf0\xcb\xf0\xe1\x7b\xe7\x97\xdf\xc0\xfc\x9f\xfe\x64\xc1\x78\xf2\xcb\x6f\x25\xa5\xb5\xf9\xf9\x77\x94\xc9\x92\x76\xdd\xbb\x1d\xf7\xeb\x74\x96\xc7\x2c\xc1\x8a\xda\x59\x76\x36\xbe\x49\x68\x5e\xe8\xaf\xdd\x33\x99\x12\x3a\x6b\x6f\x7c\xad\x90\x25\xf5\x00\xcf\x73\xf9\x02\xcf\x28\x2b\xf0\x61\x7d\x89\x48\x19\xd7\x2b\x7a\x14\xcc\x1a\xc4\x39\x7a\x01\x36\x85\xe8\x45\x62\x01\x44\x40\x5e\x88\x14\x27\xc0\xa8\x64\x20\x53\x0c\x42\xa2\xf8\xa9\x0d\x07\x93\x94\x68\x3a\x65\x12\x10\x50\x76\xc4\xf2\xdf\x00\x51\x20\x12\xcf\xd5\xe7\x28\x49\x70\x02\x1b\x5f\x3a\xfc\x45\xff\xf4\x9f\xad\x35\x87\xee\x87\x93\xfe\x5f\x1f\x7b\xb7\x6e\x4e\xad\x56\x38\x38\xe6\xe2\xd4\xc5\xb1\x95\x53\x84\xda\x79\xc4\x0a\x69\xe5\xcc\x08\xcb\x82\x53\x01\x1d\x20\x53\x0d\x45\x3f\x49\x01\xec\xb4\xa1\x0b\x4c\xa6\x98\x2f\x88\xc0\xed\x37\xd8\xba\x4e\x4c\x5d\x8f\xf4\x27\xa3\x47\xa7\xf0\x4f\xbb\x0d\x84\xdf\xb5\xc2\x9a\xa4\x18\x68\x31\x8f\x30\x87\xae\x5b\xca\x6f\x31\xb9\xb7\x74\xd7\xb9\xa7\x9d\x70\x50\x13\x38\x9d\x6d\x78\x08\xd5\x08\xd4\x1b\x00\x45\x73\x0c\x07\xc7\x47\xdd\xb3\xc3\x10\x94\x1e\xd1\x39\x65\xe7\x44\x19\x35\x41\xb9\x55\x6a\x7b\x40\x79\xec\x46\x79\x1c\x8c\x32\x6e\x82\xf2\xf8\x03\x50\x9e\xb8\x51\x9e\x04\xa3\x4c\x9a\xa0\x3c\xf9\x00\x94\x9f\xdc\x28\x3f\x05\xa3\xc4\x4d\x50\x7e\xfa\x00\x94\xa7\x6e\x94\xa7\xc1\x28\xa7\x4d\x50\x9e\x7e\x00\xca\x33\x37\xca\xb3\x50\x94\x67\x4d\xfc\x85\xcd\x9f\xf9\x41\x28\x7b\x57\x57\x6e\x9c\x8a\x18\x88\xf4\xdc\xae\x67\x43\xed\x7d\xf7\x8d\xef\xa2\x1f\x65\x01\x70\xdf\xff\xbd\x37\xe9\xbb\x31\x94\xf4\x40\x18\x9f\x9a\x6c\xcb\xa3\xad\xe6\xe2\x28\xc8\xca\x8f\x1f\xbf\xb8\x71\x29\x62\xa8\x6c\xec\xd6\xa1\x91\x6c\x44\x11\x49\x8e\x62\x89\x13\x98\x72\x36\xf7\x89\xc8\x6d\xe0\xc2\xed\x9b\xdd\x9b\xf4\x8b\xe5\x03\xcc\xdb\xf1\x91\x5f\x63\x54\xf4\x60\xb8\x47\x0d\x75\x87\x02\xf0\x01\xb0\xaf\x06\xdf\xdc\x98\x15\x31\x74\x87\xda\x6d\x5e\xe8\x0e\x9d\xd4\x03\x84\x84\x3c\x13\xa5\x47\xa2\x17\x38\x6e\x43\x42\x04\x8a\x32\x9c\xd8\xc0\x8c\x86\x0f\x1e\x34\x8a\x1a\xaa\xf6\xed\x8e\xca\xb2\x0b\xcb\x63\x3b\xa2\x52\xaa\x8e\x18\x68\xce\x9e\xb1\xd0\x02\x91\x2c\x07\xb9\x60\x46\x30\x3a\xd8\x13\x36\x44\x8f\x3e\x40\x8f\xe1\x78\xec\xe2\xf1\xe0\xd1\x24\x78\xb3\x60\x85\xe9\xaa\xc8\x33\x12\x23\x19\x04\xeb\xee\xd1\x1d\xa8\x6a\x62\xe8\xae\xb3\x8b\x69\xa7\x5d\x37\x2f\x32\x49\xf2\x8c\xbc\x67\xe3\x0d\xbf\xf5\x47\x6e\x40\x9a\x1a\x88\xe8\xb3\x5d\x67\x94\xd2\x38\x81\xe5\x27\xaf\xb8\xd4\x02\x8f\xd0\x2e\x59\x4e\x4a\x81\xe5\x88\x70\x60\x53\x23\x2a\x23\xba\x1c\xc5\x58\x40\xa4\x05\x48\xd7\xda\xa3\xb2\xe3\x53\xce\xa8\xcd\x48\x8c\x86\x13\x37\x0f\x14\x31\x94\x05\xf6\x28\x7d\x03\xe1\x29\x2c\xcf\x1c\x8c\x58\x2f\xf0\x30\x42\x49\x7c\x4a\xa6\x32\x05\x44\x13\x10\x64\x29\xd3\x92\x13\x1a\x3e\xe2\x18\xd4\x91\x5d\xf9\x30\x6a\x7b\xb3\xa9\x57\xa3\x8e\xff\xd6\xf3\x9c\x59\x4d\x0d\x65\x84\xdd\x66\x6e\xdf\x0b\x5b\x76\xc1\x78\x81\xf2\xcd\x53\xab\x36\x83\x58\xed\x86\xb7\xe0\xdc\x31\x60\x78\x08\x78\xd2\xc0\x28\x7e\x40\x04\x78\xe2\x53\xb9\x27\x4d\x54\xae\xdd\x11\xad\xa4\xb7\xed\x18\x5b\x96\x79\x75\x6f\xca\x31\xf6\x6b\x5f\x77\x88\x1b\x1e\xe1\x7e\x6a\x20\xc4\x0f\x08\x70\xdd\xf1\x6d\x78\x78\x7b\xda\x00\xe2\x07\x44\xb7\x6e\x57\x35\xdc\x49\x3d\x6b\x00\xf1\x03\x9c\xd3\xcf\x4e\x88\x9f\x83\x21\x7e\x6e\x00\xf1\xf3\x8f\x87\x78\xee\x84\x78\x1e\x0c\xf1\xbc\x01\xc4\xf3\x1f\x0f\xf1\xc2\x09\xf1\x22\x18\xe2\x45\x03\x88\x17\x3f\x1e\x62\xef\xcb\xd8\x09\x52\xd1\x02\x61\x5e\xd8\x7d\xbf\xdd\xbc\x59\x94\x60\xc8\x99\x20\x92\x3c\x5b\x6a\x2c\xbe\x24\x52\x83\x1c\xd2\x85\xdd\xb8\x23\x88\xc2\x20\xa0\x8d\x24\x52\x64\x79\xef\x7b\xcf\x7b\xdf\x87\xe7\xbe\xec\xf6\xcc\x13\x1e\xb9\xde\xfc\x0b\x63\x19\x46\x54\xbb\x95\x11\x96\x0b\x8c\x29\x60\x14\xa7\x10\x11\x59\xed\x38\xfd\x7c\xe1\x0b\x2c\xbe\x0c\x87\xb7\x3e\x88\x15\x3d\x54\x3c\xf6\x82\x51\xb0\x78\x06\x53\x88\x98\xf2\x9d\x0d\x4c\xed\x2e\x53\x26\xa1\xd3\xd2\xf8\xcc\x03\x94\x08\xbb\x6d\x18\x56\x05\x3e\xe8\xd8\x71\x0e\xdd\x11\x54\x49\x0e\x45\x69\x4f\x64\x36\x41\x89\x80\x71\x88\xaa\x2a\x6d\x30\xbc\xcb\x9e\x3b\x32\x52\xb4\xd0\x78\x20\x38\x88\x77\x41\xbb\x64\x54\xb9\x8f\xd4\xb8\x90\x3a\x6c\xe7\x84\xce\xbc\x7b\xf2\xf2\xa6\x7f\xf9\xf5\x76\x78\xf9\x75\x32\xb8\xeb\x7f\xeb\x8f\x06\xd7\xff\xe5\x06\x67\x59\x1b\x08\x36\x72\x44\x81\x0e\xbf\x19\x7e\x85\x29\x22\x99\x15\xed\x1d\xe2\x4f\x02\x24\x47\x54\xa0\x58\x12\x46\x01\x09\x20\xf4\x19\x65\x24\xa9\xaa\xd2\xca\x87\x5e\x7b\xcf\x4a\xb8\x33\x8e\x91\xc4\x1c\x64\x8a\xcc\xb1\xad\x3d\xe0\x3f\x04\xd0\x5b\x16\x3f\x4d\xc8\x5c\xc5\x91\x38\x4b\x5a\xeb\x4a\x36\x88\x98\x93\x5c\x02\x7e\x46\x59\x81\xf4\xef\xc5\x8c\x4a\x42\x0b\x2c\xd4\x2f\xcb\x94\x15\x33\x15\x78\xc2\xf0\xe1\xfb\xfd\xf0\x01\x16\x48\x00\x5e\xe2\xb8\x90\x38\x69\xc3\xa4\xf6\x9e\x4a\x0b\x66\x82\xd5\x5f\xb6\xdb\xae\x45\xe6\x44\x98\xb6\x83\xff\x54\xfb\xf4\xb8\xed\x40\x42\xf1\x0c\x29\xf5\xaf\x57\x9d\xb8\x56\x6d\xe0\x65\x1c\xf0\x3f\x0b\x94\x29\xf5\x7b\xda\x29\xff\x81\x45\x4a\x32\xbc\x9d\x17\xea\x69\x19\x16\xc2\x3c\x6a\xf5\xf5\x96\x7a\xea\x33\x89\x31\x3c\x63\x2e\x90\x7e\x9b\x4f\xed\xb5\x46\x54\x0f\x1a\xe3\x7f\x16\x98\xc6\xb5\x07\xad\x5e\xa3\xb3\x9c\x96\xff\xb4\x41\x99\xb8\x9c\xe3\x58\x73\x1b\xcf\x11\x95\x24\x16\x5a\x0d\x25\x58\x71\x3f\xc2\x89\x52\xb6\x5f\x06\x0f\xd0\xe9\x9c\x9d\x3a\x36\xf3\xdd\xe3\xed\x64\x30\x1e\xfc\xee\xdf\xc6\xab\x55\x81\x1b\x18\x39\x4e\x2b\x08\x32\xeb\xaa\x7f\x1d\x43\xbb\xdd\x86\xff\x29\x3a\x9d\x93\xb8\x74\x49\xd8\x54\x11\x28\x92\x05\xc7\x42\x53\x30\xe4\x45\xd4\x55\xff\x3a\x7e\xb3\x34\x2f\xa2\x8c\xc4\xf0\x84\x5f\xca\xb5\xf6\xb3\x31\xe1\x05\x86\x5f\xe1\x1a\x65\x02\x3b\xb4\xc1\x3c\x47\xbc\x0c\x26\xa7\x84\x0b\xb9\x7e\x0b\x40\x33\x44\xa8\x90\xc6\x88\xad\x7f\x11\x0a\x2a\x49\x06\x44\xc2\x94\xd0\x44\xa8\xdd\xdc\xbf\xbc\x1a\xf7\x60\x8e\x64\x9c\xb6\x61\x2c\x11\x97\x84\xce\x60\x41\x64\x6a\x76\x6c\x11\x09\x2d\x5d\x59\x7b\x4c\x4b\x3d\x21\xae\xff\xbe\xc0\x31\xd3\x29\x19\xeb\x0b\x70\x3c\x47\x84\xaa\xe7\xbe\xff\x55\xcc\x66\x61\xb1\xda\x91\x44\x00\xc7\xb9\xda\xe8\x49\xf9\x35\x94\x65\x35\x9e\x43\x8a\x9e\x31\x44\xca\x6a\xc7\x29\x8e\x9f\x94\x13\xc8\xb5\xf2\xc7\x54\x9f\xda\x1a\xcf\xcb\x97\x51\x7b\x33\xe7\x2c\x29\x62\x0c\x08\x44\x11\xab\x1f\x9a\x16\x19\x70\x2c\x8a\x4c\xb6\xa1\xb7\xf9\x0b\x14\x1b\xaf\x46\xbf\x1d\xa0\xda\x13\xdb\xf0\x05\xc7\xa8\x10\x78\xe3\x57\x2a\xe3\x5a\xbd\x8f\xe6\x47\xa9\xb6\x5e\xb4\xd2\x03\x44\x5f\x6a\x0c\x33\xec\x24\x82\xd1\x56\xfd\x77\xe7\x85\x90\x10\x61\xc8\x33\x14\x9b\xe3\xa1\xd9\xad\x75\xd5\x98\xcc\xa0\x10\x8a\xad\xfa\x33\xe5\x0c\x33\x9e\x60\x6e\xd4\x15\x26\x1c\x62\xc6\x39\x16\x39\xa3\xc9\x26\xf3\x05\x2c\x30\xb7\x3f\xf4\xa1\x88\xbe\xe2\x17\xc5\x3f\x8e\x13\x8c\xe7\x63\xfd\x69\x1b\x94\x5d\xdd\x64\x89\x82\xa8\x15\x5c\xcb\x34\xd3\x70\xdd\x33\x84\x93\x16\x74\x6a\x0d\x42\x70\x55\x28\xbd\x05\x08\xa2\x62\xd6\x02\x46\x31\xe0\xa5\xe4\x08\x0a\x5a\x08\x9c\x98\xf6\x2c\xf3\x75\x93\xbb\x5b\x95\xb7\x1c\xee\xfb\xc6\xf1\x7e\x8f\x45\x7b\xb5\x36\x54\x21\x38\x12\x42\xcd\x15\x82\x6d\xf9\x3b\x95\x42\x19\x3b\xf9\xcc\xe6\x58\x6d\x03\x24\xe0\x35\xfa\x16\x44\x85\x54\x9f\x1a\x36\x68\x25\x5d\x5a\x2f\x40\x53\x89\xf9\x02\x71\x97\xff\x30\xee\xff\xf5\xb1\x7f\x7f\xf9\x2e\xff\xe1\xd5\xda\x50\xff\xc1\x91\x3c\xfd\x41\xfe\x03\xc7\x99\xb6\xb1\x90\xb1\xf8\x09\xa4\xb2\x85\xac\xde\xef\x76\x80\xe9\x94\xf1\xd8\x14\x15\x4a\xbb\x74\x6e\x94\xe3\xca\xe8\x1d\x56\x8e\xe6\xca\xe6\x31\x0e\x19\xa3\xb3\xba\x17\x62\x76\x39\xb3\x79\x2d\x61\x76\xb1\xdb\x3d\x76\x09\x69\x9b\x49\x6c\x62\x0d\xed\x15\x6c\x41\xb4\x32\x79\xc2\x2f\x5e\xdb\x35\x75\xda\x2e\x85\x18\x53\x49\xf8\x6b\x9f\xc4\x3c\x43\xb4\xca\x68\xab\x65\x12\xfd\xc6\x33\x3b\x58\x69\x86\x39\x13\x12\x38\x8e\x31\x95\xd9\xcb\xd1\x6a\x1b\x2b\xa0\xc3\xab\xfe\xb8\xff\xd0\x1b\xf5\x26\xc3\x51\x55\x06\xc0\x34\x39\xd4\xec\x4c\x91\x48\xb5\xb3\xa6\xb4\xcb\x4a\xf3\x6a\x35\x14\xbd\x40\x8d\x4f\x2b\xb5\x8b\x8c\x86\xab\xad\x9e\x32\x25\x56\x22\xf4\xc3\xf4\xeb\xd5\x4d\xc1\x60\xaa\x2c\x1a\x11\x5e\x95\xe8\x94\x9f\x2f\xae\xaf\xaf\x09\x3d\x54\xf6\x00\x52\x49\x91\xaa\xb7\x0f\x4d\x54\x70\x5c\x6e\x68\xb3\x49\x73\x96\xe7\x6f\xf5\x36\x98\x84\x86\xc4\xea\x1c\x68\x9f\x98\xd0\x98\xe3\x39\xa6\xd2\xb0\x5b\x59\x02\xc5\xbe\x8a\x47\x1b\x4d\xb3\xd5\x01\x5d\xf3\x9d\x88\x52\x14\x46\x00\x75\x6b\xab\x1f\x53\xdb\x47\x96\xdf\xb6\xff\x48\x41\xe3\x14\xd1\x19\x4e\xac\x3f\xa7\xb7\x8e\x69\x21\xc6\xb1\x64\x1c\x0e\x86\x0f\xdf\x3b\x87\x6d\x18\x50\x6d\xbb\xb5\x40\x21\x46\x02\xb7\x6a\xf6\xd3\x20\xd5\xaf\xaa\xde\x83\x08\x30\xa2\x50\x9f\x33\x9a\xbd\x00\x7a\x46\x24\x53\xd1\x9a\xb6\xbb\x28\x37\x5f\x6b\xff\xfd\xf8\x1f\xce\x7d\xf1\x2e\xbd\xdb\xd8\xc0\xd9\xab\xb1\xdb\xce\x78\x03\x53\xb4\x83\x15\xaa\x1f\x6b\x37\x1f\x36\x56\x85\xb2\xc1\xdb\xb0\xea\xe5\x81\x15\xbc\xf2\x1f\xd9\xeb\x5d\xa5\xbd\x41\xed\x5f\x33\x9e\x08\x58\x10\xb5\x48\x6d\x0a\xe3\x56\xd6\x9c\xab\x72\x9f\x26\x48\x22\xc3\x98\x10\xc5\xf7\x96\x81\x57\xfd\x87\xc9\x8d\x93\x71\x86\x1a\x9a\xd7\xf0\x56\x93\xec\x0c\x33\x6e\xcf\x58\x1f\x3e\x41\xfe\x85\xdf\xfa\x3a\x2b\xfe\x3d\x14\xd2\x1c\xc1\x9a\x4f\xb5\xae\x93\x6d\x4d\xf2\xfa\x3a\x65\x1a\x34\xca\x5c\xd8\x8b\x2e\xcd\x52\xa4\xb5\x0e\x99\xc8\x97\xbb\xf1\x36\xc8\x34\xe9\x8f\xf9\x6c\xaf\x8e\x39\xbc\xab\x90\xde\x98\x9a\x3b\xf3\x16\x86\xa7\x44\xdb\xa0\x42\xfb\xd9\x2e\x0a\xa7\x8b\xb8\x7c\x6f\x31\xd6\x87\xa1\x7f\x3b\x76\xf7\x3c\xf6\x5f\x4d\xae\xbc\xab\xce\x6c\x2f\x71\xd9\x41\xb8\x12\x9d\xb2\x74\x1c\xb1\x0e\xef\x86\x0f\xdf\x07\xd7\xca\xf9\xd4\x09\xaa\xc9\xea\xbf\xd5\xeb\xe9\x6c\x95\x76\x52\x2b\x9d\x21\x53\xac\x5d\x53\xa1\x8f\x90\xd4\x86\xd9\xd8\x73\x65\x4d\x49\xf0\xc3\xdf\xf1\x60\xca\x2c\x8d\x28\xfd\xfb\xab\xc1\xb5\x9b\xb5\x9a\x1a\xca\x5b\x7b\x6d\x2d\x80\xb7\xfd\x32\x23\x41\xa6\xbf\xe2\x4c\x60\x88\x54\x7c\x60\x72\x02\xfa\x3f\xcb\xb8\x1c\xd3\x44\xa7\xc7\x5e\xe5\xd6\x36\x5c\x80\x9e\x4e\x18\x6a\x1c\x3a\x70\x60\xc6\xfe\x0d\xae\x01\x23\x9e\x11\xcc\x5f\xa7\x0b\x2d\x2c\xf2\x8e\x0a\x35\x1a\x13\x3a\xb7\x6f\x3f\x4f\x52\x7a\xab\x4f\x5f\x8d\x0c\x75\x37\x46\x86\x8c\xe4\xf1\x12\xc5\x32\x7b\x31\x21\xd2\x16\x67\x58\xe3\xd9\xe2\xf1\xd4\xd7\x84\x22\xb7\x6f\x8e\xed\x3d\x82\xef\xf4\x74\xf4\xab\x19\x37\x87\x17\x54\xd4\x7c\x1d\x8f\x7f\x63\x86\xdf\x5c\x68\xdf\x8c\xc6\xbd\x1a\x9e\xea\xb8\xb0\xfe\x9b\x8d\xcd\x5d\x8f\x86\x77\xbd\xdb\xc9\x78\xd2\xbb\xfc\xea\x66\x56\x7d\x51\xa8\xe2\xb0\x87\xbb\x07\x28\x93\x87\xcb\xae\xc3\xc0\xd8\x9b\xd0\x57\xee\x8b\x49\x24\xac\x18\x52\x6b\x49\xd3\x69\xc8\x92\x41\x95\x3d\x25\x72\x1d\x47\xa1\x4c\xba\xbc\x9b\xdf\x47\xfd\xde\xa4\x3f\x9a\xdc\xf4\xee\x9d\x8c\xa8\xaf\x09\xf5\x85\xed\xfb\x26\xd8\xdb\xd9\x50\x09\xe8\x4d\xb9\x22\xda\xa2\x09\x6a\x08\x86\x23\xbf\x12\xb4\x2c\x0d\xc5\x6c\xcf\x3c\xed\x1b\x73\xbd\x44\xb3\x0d\xff\x4d\x6f\x7c\xd3\x3d\x73\xcf\x1e\x56\xf4\x50\xa4\xf6\xbe\x0c\x57\xc3\x42\x8a\x44\xea\x4c\xe2\xac\x3a\x16\x4c\x6a\x05\xe4\x82\xc4\xf8\xb7\xb2\x1c\xa1\xf3\x64\xe3\x9b\xde\xd1\xf1\xe9\x99\x89\xd1\x95\x6b\xa0\x3f\x1d\x0d\x1e\xfa\x77\x57\x47\xdd\x33\x4b\xfd\x55\xc1\x3a\x3e\x75\x77\x4c\x55\xf4\x50\xd8\xf6\x2c\xc8\x9e\x60\x33\x9d\x3d\x14\x1b\x98\xdf\x42\xf3\x78\x3a\x0d\xdc\x1c\x7b\xdb\x46\xb8\x0b\xb9\x76\x7e\x57\xc9\x78\xa5\xa7\x75\xe9\xa9\x55\x69\xe4\xba\x23\x57\x2b\x7b\xda\xbf\x5e\xe6\xf2\x6d\x0c\xf0\x45\x03\x86\x1a\x1a\x0f\xd8\xd9\xe0\x49\x19\xbb\x62\x02\x3f\x3b\x3a\x2d\x48\xaa\xa0\x01\x88\xc5\x93\x1d\xdc\x7f\xeb\xdd\x0e\xae\x86\x0f\x2a\x16\x77\x63\xdc\x58\x15\x88\x75\x6a\x2f\x46\x04\x88\xfc\x0e\xc9\x38\xc5\xc2\x24\xb0\x4c\x62\x4a\xa6\x48\x56\x20\x5f\xb0\x04\x24\x04\x99\x51\xab\xf4\xee\xbf\xf5\x47\xee\xe6\x88\x92\x1c\xea\x90\xd9\xe5\x17\xda\x3f\x75\x9d\x91\x5c\xe8\x12\x55\x69\x5e\x23\x22\xc5\x46\xef\x8e\x2f\xd4\xbe\xed\x5f\xbb\x71\x69\x62\x28\x2a\x57\x57\x98\xce\x7b\x84\x41\xfb\x8a\x71\x5e\x26\x0f\xe3\x14\xe9\x29\x30\x2e\x20\xc3\x53\xb9\xca\x30\xe5\x38\x26\x53\x82\x13\xc8\x19\xa1\xba\x65\x09\x95\xdd\x21\x7e\xd4\xe3\xb1\xd7\x87\x58\x2d\x08\x4d\x97\xd8\xf7\xe9\x1e\x8c\xe9\xba\x43\x61\x9b\xf5\xac\xde\x7d\x9b\xeb\xf0\x7a\x5d\xa8\x59\xb1\x77\xbc\xec\x15\x6a\x88\xd3\x70\x3b\xbe\x19\xf8\xf6\xb2\x21\x87\xca\xd3\x1e\x34\x05\x83\x1c\xa7\x64\xaa\x6c\x88\xd9\xbc\x91\x3e\xa3\x2d\xc8\x39\x16\x98\x3f\xab\x18\x4b\x69\x1e\xdf\x8e\xbd\xeb\xfd\xe1\x84\xa6\x68\xa1\xc2\xb3\xe7\x30\x1b\x0b\x4f\x9d\xc5\x0c\xf1\x99\xc9\x57\x96\x8d\x76\x16\x14\x03\xf7\x91\x53\xb4\x50\x14\x7b\xea\xdf\xac\xa3\x10\x73\x94\x65\xdb\x60\x0c\xdd\xe5\x2a\x45\x0b\xdd\x64\xf6\x9c\xc4\x4e\x30\x4c\xb3\x87\xee\x89\xd0\x49\x74\x9d\x7a\x55\x1b\x0d\x6d\xcd\xbe\xfa\xa6\xde\x1a\x0c\xbd\x5d\xd8\x73\xaf\xcd\x32\xc8\x9b\xd3\x6e\x5e\x18\x5b\xe6\xb5\x9b\x8d\x6b\x9f\xdb\xf5\x7b\x93\x9e\x67\x75\xe2\x37\x8b\xee\x44\xc0\x34\x23\x79\x6e\x45\x33\x70\x7b\x8f\x8a\x16\xea\x3b\x06\xa7\xc1\x1c\xc3\x59\xf5\x7c\xb8\xe9\x85\x3a\x92\xec\x68\x5b\x56\xf9\xde\x93\xdf\xbf\x6f\x30\xfe\xea\xbd\x28\xc7\x9b\xdc\xb2\xe7\xc8\x19\xd6\x5e\xa1\x49\x16\x59\x5e\xde\x7d\x61\x8c\x26\x06\xbe\x7e\xd4\x69\x41\x67\x19\x9d\x1c\x75\x96\x91\x3d\x44\x0d\xf0\x73\x27\xd5\x8c\x00\x11\x40\x66\x94\x71\x15\xad\x54\x80\x60\x8e\xf8\x93\xa3\x35\xc4\x0e\xd4\x1d\x84\x1b\x6a\x28\xd4\xff\x67\xf8\xdc\x43\x66\x9a\x18\x8a\x6e\xe7\xc0\x74\xbf\xe8\xdc\xf3\x65\x9a\x18\x8a\xce\xee\x33\xfc\x34\x74\xee\x44\x89\x26\x86\xa2\xb3\x9b\xa9\x9f\x86\xce\x3d\x55\xa6\x89\xa1\xe8\xec\xb5\xc3\x9f\x86\xce\x3d\x50\xa6\x89\xa1\xe8\x76\xae\x29\xee\x17\x9d\x7b\x96\x4c\x13\x43\xd1\xed\x5c\xd5\xdb\x2b\x3a\x77\x88\x75\x1f\x3e\x3c\x7f\x61\xb7\xdc\xa1\xfe\xd4\xe0\xf5\x5d\x7d\x2a\x72\xec\xb6\x4c\x4b\xda\xca\xa9\xaa\x4d\xbf\xd4\x06\x63\x74\x2f\x4a\x64\x1d\x88\xd1\xf5\x5e\x1f\xdc\x06\xb9\xcb\x9d\x95\xa8\x3b\x59\xf7\x23\xf2\x96\xf7\x8f\x77\xfe\x2c\xc2\x6a\x41\xa8\xe4\xed\x25\xa7\xdd\xd2\x07\xeb\xb6\x99\x12\xf7\x3b\xca\xae\x15\x80\x2d\x95\xd7\x57\xcb\x42\xd1\xda\x7b\xcd\x9c\x68\xc3\x4a\xaf\xd5\xbb\x85\x56\x5f\xef\x1f\xef\xb6\x5e\xc8\x59\x5f\x13\x0a\xda\x3e\x52\xb2\x5f\x11\xaf\x1a\x90\xb7\x88\xd9\xd3\x3e\x17\xde\x33\x77\xee\xe8\x27\x6a\x3e\x80\xc9\xf8\x0e\xf3\x97\xde\x7b\x5d\x1a\x5d\xeb\x12\xdc\x2b\x50\xdd\xbf\x63\x05\x59\xbb\xc7\xc5\x15\x23\xd6\x6e\x2e\x79\x8b\xef\x61\xe0\xa9\x7b\x6b\x62\x28\x3e\x7b\x18\xb2\xa4\x7a\x46\xa1\xbc\x48\xa8\x53\x0d\x2b\xf8\x06\x13\x5e\x7f\x65\x49\x9d\x16\x58\xc3\xa4\x96\x0b\x6b\x88\x80\x58\x71\x28\xf1\x33\xe1\xf1\xcb\xd7\xbe\x5b\x3f\x95\xe4\x40\x46\x4c\xed\x47\x34\xc0\x1c\x8d\xb0\xce\x62\x6a\x4b\x53\x6f\x4d\xd6\x63\x3d\x92\x44\x19\x36\xa5\xc1\x5a\x2f\xac\x0b\xdb\x4d\x6f\xec\x6e\xd5\xac\x2d\x09\xc5\x68\xd7\xbd\x7b\xc0\x58\x16\x41\x2b\x7c\x65\x59\xda\x06\x6f\x7c\x73\xd5\x9b\xf4\xdc\x69\x83\xf5\x8a\x40\x70\x9f\x1c\x9d\x1b\xba\x30\x82\xb2\x43\x3b\xca\x04\x49\xe4\xdc\xa7\x14\x2f\xa5\x6e\x6c\xd1\x73\xa0\x88\xd0\xd7\xed\xa8\xa6\xe9\x45\x32\x3d\x9e\xf5\xbe\xab\x07\x2a\x7c\xee\x2b\xfa\xd6\x2b\x42\x39\x60\x17\xef\xce\x1c\x90\x0b\x56\x22\x2d\xd9\x10\xca\x05\x75\xc8\x33\x22\x65\xa6\x07\x31\x88\x56\xf2\x09\xe6\x6e\xe6\xb8\x93\x11\xeb\x15\xa1\xcc\xb1\x9f\xef\x9d\x99\x33\x65\x05\xff\x20\xee\x8c\xfa\xe3\xfe\xe8\x5b\xdf\x9d\x71\x5f\x2d\x08\xe4\xcd\xa9\xbd\x52\x19\x12\x57\x59\x5b\x1f\xa1\xa0\xba\x82\xc5\xe2\xb8\xe0\xca\x89\x23\x14\x10\x85\x82\xd6\xdb\xd6\x07\xd7\x10\x71\x44\xe3\xd4\x89\xd7\xad\x2c\xd6\x2b\x42\xdd\x96\xdd\xb3\x6f\x3f\x12\xb1\x5b\x39\xac\x57\x84\x22\xb6\xf7\xbe\xfc\x7c\xc4\x93\xc7\x91\xbb\x14\x56\x92\x43\x43\x4b\xef\xfd\xea\x76\xc8\x8d\x86\x08\xdb\x30\x26\x34\xd6\x8d\x06\x31\x23\x14\x3a\xed\x8b\x96\x2e\xbb\x23\x9a\x20\x9e\xc0\xc2\xf4\x46\x22\x29\x51\xac\xa3\x18\x33\xf7\xaa\x47\x3a\x94\x0a\x58\x3f\x54\xb7\x46\x4a\x06\x28\x49\x00\xc1\xbf\x30\x67\x47\xe5\xc0\x60\x15\x9b\xeb\x6b\x3b\x36\x06\x74\x63\x46\x05\x11\x7a\x58\x9b\x4d\x61\xc5\x2d\x98\xb2\x2c\x63\x0b\x53\x18\x52\x3f\xd5\x86\x71\x11\xa7\xd5\x60\x9d\x19\xd8\xe2\xec\x19\x45\xd9\x0b\x14\x54\xe4\x98\x26\x7a\x18\x49\x0f\xda\x19\x75\x98\xbd\x28\x27\x3b\x46\x3c\xa9\xc6\xba\x84\x64\x1c\xcd\x70\xe5\xaf\x3d\x4e\xfe\x18\x82\xc0\xb2\x05\x1c\x27\x45\x5c\x4e\x20\xeb\x79\x63\x21\x2b\xef\x8d\x62\xb9\x60\xfc\xa9\xe2\x52\xa7\xdd\x3d\x6e\xad\x99\xc3\x71\x86\x5e\x80\x17\x19\xd6\x4d\x1b\x6c\xa1\xf0\x11\x3a\xcb\x36\x41\xaf\x70\xb5\x4c\x9b\xca\xca\x14\xeb\xe1\xe9\xea\xda\x01\x3d\x40\x2b\xd2\x7a\xb8\x7f\x60\xda\xd3\xab\x13\xf3\xf7\xee\x3f\x0e\x6b\x73\x34\x6b\x76\x29\x5e\xe8\x61\x0c\xe3\x70\x4a\x94\x6d\xb2\x39\xc3\x74\x26\x53\xdd\xa2\x2a\xcd\xf8\xcd\xf9\x89\xd1\xea\x16\x15\x3d\xf8\xfd\xc6\x9d\x16\x32\xd4\xd0\x83\xeb\x4a\x0c\xed\xad\x8d\x84\x93\x59\xba\x6b\x1f\x89\xe9\xe6\xf3\xb5\x2a\xae\x57\x04\x32\x00\xd9\x33\xb6\x7b\xe9\xda\x33\xb3\xf3\xbe\x56\xc4\xd1\xf0\xd6\x9d\x12\xd0\xc4\xd0\x88\xcb\xae\x9c\x1a\x44\x5c\xfb\x08\xb7\x5e\x5f\x96\x69\x63\x80\x67\x43\x37\xb8\x24\xd4\x3e\x64\xb7\xe5\x76\xc5\xf2\x6a\x45\x27\xc4\x13\x9e\x18\x98\x09\x5b\x84\x23\xf4\xf7\xcb\x8c\x9a\xf5\xcb\xd8\x3d\x8c\xe6\xfd\x32\xe6\x94\x36\x68\x98\x19\xdf\x78\xc2\x2d\x4d\x0c\x3d\x90\xf6\x22\xc3\x1e\x0f\xe4\xf8\xa6\x77\xd4\xb5\x22\xf1\x75\x05\x97\xe4\x50\x34\xf6\x0c\xcf\x9e\xd1\x58\x7b\x81\xc7\x83\xff\x76\xb7\x65\x68\x62\xa8\xad\x08\xfb\xd3\x3f\x16\x1b\x52\x9b\x5e\x10\x69\x95\xa9\xd2\xda\xbf\xb2\x83\xb5\x1b\x0c\x70\xa6\x4d\xed\xc6\x05\xbb\x70\x50\x8d\x38\xe5\x2c\xcf\xb5\x43\x28\x0f\x2d\xd0\x3d\x17\xed\x37\xb8\x67\xff\x62\x4f\x5d\x5c\x91\xed\xa2\x7d\x64\x7d\xfd\xf1\xc4\x9d\x69\x2c\xc9\xa1\xaa\xd1\xd5\x52\x03\x11\x9e\x35\xb2\xf7\x55\xbe\x18\x81\xc0\xc6\x77\xd5\xcd\x5c\xdb\xcd\xb9\xf7\x4e\xe4\x46\x57\x22\xdb\xb3\x33\xde\x36\x1b\x9f\xc2\xaf\x2e\x40\xae\x46\x81\x6b\x1b\x50\xf9\xb7\x62\x81\xec\xdd\x43\x93\xe1\xd6\xf1\xa1\xda\x92\xd0\xb8\xc3\x65\xdb\xec\x08\x2d\x43\x45\x21\xc3\x43\xeb\xe1\x20\xeb\xec\x50\x6d\xb6\xe8\x2d\x17\x46\x8f\x6e\xc5\xf3\xfa\x0f\x71\xbd\x1a\x35\xeb\x3a\xb3\x08\x3f\xfd\x8f\x74\x4d\x1e\x7d\x62\x7d\x6c\x90\x1d\x0f\xfe\x6b\x02\xde\xab\xf7\x57\xce\x18\x92\xd6\xcb\xc9\x6b\xf9\x6f\x3d\x81\x4b\x05\xe6\xfa\xaa\x0c\x3c\x65\x1c\x5b\xaa\x06\xf6\x9e\x32\x5f\x0d\xa4\x41\x09\xe4\xcc\x6e\x54\x7e\x76\xe6\x40\x17\xf5\x7c\x38\x1b\x94\xa4\x77\xef\x7c\xb1\x23\xc5\xcf\x98\xc2\x22\xc5\x74\x27\xb0\xee\x92\x47\xb3\x52\xec\x99\xdd\x45\x75\x8d\xf9\x06\xd5\x64\xb7\x5e\xbd\x64\x9f\xb0\x91\xbc\xc0\xe1\xa5\xf9\x6f\xfd\x91\xbf\x41\x61\xb5\x20\x94\x43\xbb\x37\x0b\xfd\xa0\x0d\xf1\xb7\xc1\xe4\xc6\xd3\x42\x5e\x92\x43\x5d\x61\x47\x31\x17\xe6\x84\xc2\x1c\x39\xe6\xa9\xde\x55\xac\x5e\x2a\xf4\xca\x33\xac\x02\xd0\x55\x9a\x81\x23\x3a\xc3\x70\x90\xe1\xa9\x3c\x22\x34\xce\x0a\x41\x9e\xf1\xe1\x96\x1a\xf6\x1f\x9e\x22\xf6\x1f\x0d\xaa\xd8\x8e\x0b\x25\x9a\x57\xb1\xf1\xb2\x44\xd2\xa4\x9e\xfd\x97\x3f\xff\x2f\x00\x00\xff\xff\xf0\x1d\x60\xbf\x4b\x75\x00\x00")

func specJsonBytes() ([]byte, error) {
	return bindataRead(
		_specJson,
		"spec.json",
	)
}

func specJson() (*asset, error) {
	bytes, err := specJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "spec.json", size: 30027, mode: os.FileMode(420), modTime: time.Unix(1708331577, 0)}
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
	"spec.json": specJson,
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
	"spec.json": &bintree{specJson, map[string]*bintree{}},
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

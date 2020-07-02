// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../resources/events.yaml
package util

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

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x5d\xcd\x72\x1b\xb7\xb2\xde\xfb\x29\x50\x5e\x9c\x4a\xaa\xac\x1c\x8b\x3f\xb1\xe8\x9d\x2c\x4b\x2e\xd5\xb5\x2c\x97\x24\xe7\xdc\xbb\x4a\x81\x33\x4d\x12\xd1\x0c\x30\x01\x30\x94\x79\x52\x79\x17\x3e\x0b\x9f\xec\x16\x80\xf9\x23\x81\x21\x9b\xa0\x2b\x5e\x59\xe4\x0c\xbe\xee\x46\xa3\xff\xd0\x00\xcf\xce\xce\x5e\x7d\x12\x34\x7b\x4f\xae\x97\xc0\xb5\x22\xdf\x38\x9b\xb1\x84\x6a\x26\xf8\xab\x5c\xa4\x65\x06\xea\xfd\x2b\x42\xce\x08\xa7\x39\xbc\x27\xaf\xaf\xee\xef\xee\xee\xbf\xbc\x7e\x45\x08\x21\x89\x28\xb9\x7e\x4f\x26\x93\x89\xfd\x93\xa5\x8f\x9a\x4a\xfd\x9e\xbc\xad\xfe\xbc\xe6\xe9\x7b\x42\xda\xef\xf9\x4c\xbc\xb7\xff\x33\xe3\x25\x22\x85\xfa\x51\xf3\x2f\x83\x25\x64\xef\xc9\xeb\xdb\x2f\x37\xf7\xaf\x9b\x4f\x73\x50\x8a\xce\x0d\xf0\x63\x99\x24\xa0\x54\xfb\x55\x21\xc5\x34\x83\xfc\x3d\x79\xdd\x7e\xa6\x44\x56\x1a\xd2\xab\x0f\x5b\xaa\x3f\xdf\xee\x90\xbc\x4d\xf1\xf9\xdb\xb7\x3b\x44\x9f\xbf\x9d\x5c\xf4\x52\xdd\x3c\x8e\x21\x3c\xc9\x18\x51\x20\x97\x20\x09\xe3\x4c\x33\x9a\xb1\xff\x42\x7a\x14\x1f\x3b\xd8\xe7\x68\x6c\x0e\x2f\x24\xc9\x18\x70\x7d\x0a\xde\x60\x17\xcf\xc0\xf9\x68\x0e\x88\xa4\x4c\x25\x82\x73\x48\x34\xa4\x1e\xa8\x8f\xe8\xc3\x0d\x71\x70\x12\xfe\x2c\x41\x69\x92\xab\x39\x91\x90\x2c\xa3\xb0\x46\x58\x2c\x55\x08\xae\xc0\x82\x29\xe0\x3a\x0a\x6c\x8c\x03\x03\x4e\xa7\x19\x90\x67\x80\xe2\x8c\x66\x6c\x09\x51\x60\xbf\xee\x82\xfd\xe7\xf2\xe1\x0b\x02\x8c\xcc\x28\xcb\x22\x67\xee\xdd\x51\x0c\x4a\x28\x15\x9c\xd1\x34\x95\x51\x60\x17\x47\x31\xd8\x82\x6d\xd6\xf1\x1c\x9e\x7b\xcb\xbe\x07\xf4\x7b\x02\x85\x19\x82\x08\x4e\x12\x91\xe7\x94\xa7\x06\xc7\x58\x30\xc6\xe7\x51\xd0\xde\xaa\xbf\x7e\x78\xb8\x7f\xf0\xb1\x1d\x77\x44\x0b\xa3\xa8\x71\x5c\x22\x17\xfc\xf5\xff\xde\x3e\x45\x8d\xef\xad\xf0\xb0\x14\xcd\xf8\x3b\x5c\xee\x45\x69\xcc\xfd\xed\x07\x71\x73\xff\x48\x1e\x35\xd5\x40\xee\x28\xa7\x73\x90\x5b\xf6\x7f\xd0\x98\xfb\xda\x01\x9c\x7b\x0e\x60\xb0\xc7\x6d\x35\x8f\x77\x78\xf8\x78\xfd\xe1\xdb\x27\x9f\x09\x65\x88\xf8\xbd\x28\xa7\x19\x53\x0b\x88\x51\xf5\x73\xdf\xe0\x87\xe7\xc3\x42\x91\x44\x70\x0d\xdf\x35\xa1\x69\x1a\xa5\xe4\xe7\x58\x7b\xbf\x0d\x27\x21\x17\xcb\x48\x40\xa4\xc5\xdf\x06\x2c\x8b\x94\xc6\xb9\x98\x73\xac\xd9\xaf\x00\x17\x94\xcf\xe3\x80\x7c\x7b\x11\x06\x7a\xfa\xf0\x31\x62\xf8\x81\xaf\x84\x0f\xd7\x5f\xef\x1f\x9e\x7e\x7f\x7a\xb8\xbc\xba\xfe\x31\x30\x3b\x08\x3f\x8c\xf6\xa1\x37\x07\xfb\x69\x67\xd5\x9a\x5e\x29\x0d\x39\x79\x80\x44\x2c\x41\xae\xc8\x2d\x2f\xa4\x98\x4b\x50\xea\x48\x1b\xf1\x9b\xc8\xca\x1c\x42\xc6\x61\xb8\x6b\x1c\x06\x5e\x74\x38\xd8\x67\x1c\x06\x81\xe8\x30\x3c\xed\x8e\x06\x92\x48\x88\xd3\xe4\x41\x20\x16\xdc\x8b\x94\x42\x06\xb1\x48\x48\xab\x50\x21\xe5\x46\x96\x91\x48\x48\x73\x50\x21\x95\xfc\x14\x2c\xa4\x25\xa8\xb0\xac\x4d\x35\xae\x55\x2f\x80\x64\x4c\x45\xc4\x81\x03\x74\x1c\xe8\x30\x37\xeb\xca\xb2\x92\x99\x14\xb9\x01\xde\xac\x63\x91\xb1\xe1\xcb\xd3\x02\xea\xf0\x1a\x52\xb2\xac\x74\x47\x80\x22\x5c\x68\x02\xdf\x43\xf0\xf6\xa5\xea\xd1\x17\xa6\x17\x56\x44\xde\x20\x66\xf1\x11\x21\xeb\x3f\x6f\x3f\xf6\x0d\xdb\x66\x24\xd7\x5c\x83\xb4\xa3\x25\x42\x4a\x48\x74\xff\x58\x74\x66\x1e\x4d\x16\x90\x3c\x33\x3e\xb7\xef\x54\x5f\x76\x24\xd6\x11\x87\xb7\x78\xc2\xe2\xa8\x26\x5f\xb1\xff\x02\x61\x8a\x68\x21\x48\x46\xe5\xdc\x8f\xcb\x9b\x05\x4d\xb9\x65\x69\xda\xac\x6d\x32\x85\x84\x96\x0a\x2c\x4d\x39\xfd\xce\xf2\x32\x27\xbc\xcc\xa7\x20\x89\x98\x55\x54\x2a\xa2\x17\x54\xdb\xb7\x3b\x6f\x32\x65\xc3\x49\xe8\x7a\xf3\x56\x3a\x0f\xa0\xe5\xaa\x62\xdc\x2a\x8a\x61\xbc\x34\xe9\x97\xa1\x5e\xae\x6a\x01\xd0\x5c\xf0\xb9\x93\xb1\x79\xa2\x1e\xbc\x02\xf6\x45\xe3\xad\xf6\x9e\x10\xcd\x52\xe6\x86\xb3\x71\xf6\xb2\x95\x55\x58\x45\x3c\x95\xb0\x62\x75\x1c\x2a\x72\x29\x25\x5d\x91\x84\x16\x34\x59\x05\xb8\xbd\x32\x53\x6b\x65\xa8\x9c\x1f\xb0\x4f\x32\xbd\x22\x26\xc0\xb6\x92\x98\x53\xc6\x3d\x76\x06\x48\xc5\xdf\x99\xbf\x05\x5d\x02\xa1\x64\x9a\x51\xfe\x6c\xd5\xcd\x67\xe9\x83\xfd\x8e\x39\x15\xa6\x59\x26\x5e\xcc\x4a\x6d\xb5\x72\xeb\x2d\xac\x52\x77\xde\xb7\x92\x65\x22\xc0\xd2\x71\xca\x6b\xc6\xdd\xac\x1b\xf5\x15\x81\x24\xa4\xfb\x28\x01\x43\xa0\xd3\x3e\xef\x85\x5d\x36\xe8\x16\xf1\x56\x87\x3b\xab\xda\x2a\xc8\xf7\xef\x26\x8c\x92\x34\xd1\x20\x7d\x6d\x1b\x20\xb5\xad\x4b\x20\x53\x24\x2d\x8b\x8c\x25\x41\xd7\x49\x2e\xb7\x6c\x11\x6d\x9f\x75\x6f\xd3\x4c\x02\x4d\x57\x6e\x41\xa8\x3d\x9c\xa5\x6c\x36\x03\x09\x5c\x07\x26\xb4\xc3\x00\x32\xa3\xf9\xc6\x69\x62\x32\xc3\xee\x52\x09\xeb\x55\x68\x2e\x4c\xe4\x4b\x19\x57\x84\x12\xa5\xa5\xb3\x70\x54\xef\x28\xdf\x5e\x65\xf3\xe6\x29\x07\xd0\xca\xfb\x4a\x96\x59\xc0\x26\x0c\x91\x8b\xe8\xa1\x59\xe1\x76\x69\xef\x31\x97\xfd\xf6\x20\x68\x66\xc3\x56\x6f\xcb\xdc\x17\x42\x29\x16\xb6\x43\x1d\x46\x90\x4b\x27\xc0\x88\xca\x69\x96\x1d\xcf\xc8\x66\xbd\xfd\x62\xc8\xa2\xe5\x8c\x5b\xaf\x60\xe6\x31\xf1\xad\xa9\xb5\x70\xd2\xf0\xed\xf3\x83\x5c\x3e\x3b\xfc\x38\xa5\x61\x73\x1e\x5a\x3f\x21\x8e\x36\x6b\xfb\x5e\xb3\xb6\xed\xcb\x6e\x81\x4d\x33\x91\x3c\x6f\x5b\xfd\x96\xc7\x5b\x5e\x94\x7a\x8b\x17\x2d\x8c\x87\xcb\xcb\x4c\xb3\x22\x03\xe3\x01\xbd\x01\x5a\xf6\x46\xc7\x19\xef\x7a\x69\xf7\x45\x85\xe4\x52\x6b\xc8\x0b\x6d\x88\xb0\xcf\xb4\xf6\xab\x5e\x4e\x7d\x43\xb4\x2c\x7d\x11\x7a\x61\x75\x4e\x90\x54\xf8\x14\x1f\x67\x9b\x6b\xb8\xfe\x48\xb6\x4b\x73\xf5\x54\x3f\xd5\xfe\x30\x58\xba\xc7\x7e\xde\xc2\x67\xc2\xa7\xfb\xaa\xf2\x8f\x2e\x9d\xa8\xc5\x54\xd1\xd3\x4b\xbc\x8b\x9c\xab\x77\x68\xdf\x5b\x2d\xad\x1d\x9e\xab\x97\xea\x58\xc6\xae\xfb\x8a\xcd\xed\x40\xcf\x63\xe9\x57\x8f\xa5\x6b\x29\x85\xf4\x79\xba\xa1\x2c\x33\x50\x2e\x32\xaa\xa1\x72\xd0\x34\x26\xd2\xfe\xd5\xaf\xd6\xed\x87\x15\x05\xf0\x93\x41\x3d\x3b\x70\x00\xd4\xe8\xcb\xc9\xa0\x9e\xeb\x3b\x00\xfa\x22\xd9\x0f\x90\xaf\x97\xb7\xf5\xa0\xfe\xd6\xe2\x6c\xd6\xb6\x74\xc4\x35\x99\x4a\xf1\x0c\x3c\x06\xf7\x1d\x56\x9d\xee\x20\x17\xc6\x45\x39\x63\xce\x04\x77\x65\xe7\x52\x86\xd6\x07\x49\xa8\xaa\xd6\xb1\x5a\x88\x32\x4b\x09\x87\xa5\xc9\x0c\x92\xa4\x94\xe4\x8c\x2c\x80\x16\x9d\xa1\xc8\xee\x48\xed\x9a\x79\xea\x0d\x81\xdf\x61\x35\xf2\x96\x2f\x69\xc6\x52\xc2\x78\x0a\xdf\x09\x6c\x3d\x84\x27\xd9\xbe\xfd\x53\x35\xcb\x2c\xfd\x99\x30\x13\x84\x70\x9a\x65\x2b\x32\x97\x94\x57\x99\x0d\x73\x60\x41\xa7\xe1\x9e\x27\x99\x98\xb3\x64\xb3\xee\x12\xd2\xe1\xea\x58\x95\xb7\x52\xfc\x65\xb3\xe6\xf0\xb2\x59\x37\x99\x63\x04\x83\xdf\xdc\xbe\x82\x16\x64\xce\x96\xd0\x26\xa1\x6f\x48\x0a\xaa\x30\x2a\xde\x89\xaa\x6c\x65\xa9\x0e\xd4\x72\xfa\x3d\x9e\x5f\xec\x6a\x33\xfe\x5b\x51\x97\x16\x57\x44\xec\xc4\xba\x78\x56\x2f\xbd\xbc\xbe\x1e\xf9\x60\x18\x8d\xe3\xea\xc2\xf7\xf0\x54\xf2\xee\x46\x49\xc3\xd7\x7d\xa9\x4d\xb8\xf0\xa7\x50\x44\x52\x1e\x0a\x2a\x2f\xc9\x92\x66\x25\x90\x0c\x94\x4d\xa8\xf9\x76\x74\x55\xd8\x3c\xc0\x4c\x9d\x19\xc3\x3d\xfa\x42\x55\x1d\x64\xa3\x42\xb4\xf6\xcd\x6e\xc2\x5e\x87\xe9\x5b\xa9\xe8\x2f\x1e\xb3\x13\x3c\xb3\x2e\x1b\xb6\x3a\xd3\x17\x18\xb4\x19\xeb\x4e\xc1\x41\x48\x92\x09\x9a\x42\x6a\x67\x4d\x94\x9a\xb0\xa9\x98\x09\xb5\x27\xa8\x69\x8c\x47\xe5\x61\x5d\x9c\xe1\x5e\xf3\xd9\xf0\x63\x9c\x3e\x36\xaa\x9a\xed\x0d\x2d\xb3\x40\xdd\xa8\xe6\xa0\xda\x25\x6b\x39\x29\x40\xce\x84\xcc\x8d\xa1\x70\x73\xf8\xf8\x74\xff\xd5\xd5\xe4\x11\x96\x7a\xe2\x07\xc5\x7d\xf4\x7d\x14\xbc\xd2\xed\x1e\x6b\xf7\x28\xcc\xf2\x31\xdf\x29\x92\xd3\x55\xb5\x30\xd2\x52\x36\x69\x47\xb3\xb5\x67\x94\xb3\x53\xf9\x7a\xe3\xb4\xc1\x2c\x99\x72\xaa\xcc\x67\x5c\x1b\x5f\x2f\x9d\x21\xcf\xab\xe9\x7d\x11\xf2\x99\xbc\x40\x96\xfd\x12\x4a\xdf\x0c\xb0\xad\x09\x58\x12\xc8\x82\xf2\x34\x33\x50\x34\x33\x13\x3b\x5f\x10\xa6\x6b\xb1\x39\xca\x2c\x2f\xa5\x02\x49\x1c\x64\xe2\xc5\x45\x81\xfd\xab\x5e\xf1\x18\xca\xdd\x0c\x1a\x84\xbe\x62\xb5\x1f\x56\x77\x8d\x06\x17\x9d\x61\x76\x87\x68\x59\x7d\xac\x9f\x51\x24\x2f\xd5\x4e\x01\x6d\x26\x64\x15\x6b\x1a\xe6\xdb\xca\x55\x53\xd2\xbf\xfb\xf0\x10\xdc\xed\x3b\xf7\x0a\xfa\x43\xaf\xa0\x3f\xdc\x57\xd0\x1f\x62\x0b\xfa\x97\xce\x47\x9b\x78\xc3\xba\x7f\xe5\x9a\x56\x22\x22\x8d\x21\xb6\xb2\x9f\x4f\x25\x49\xa9\xa6\x75\xe0\x6a\xf4\xcd\x86\x58\x51\xa0\xc8\x22\x7f\x03\x4a\xd3\xf4\x44\x44\x64\xb1\xbf\xa0\x92\xe9\x95\x4b\xf6\x4f\x12\x2b\x76\xeb\xcf\xa9\x6a\x59\xb2\xf4\x74\x50\x6c\xdd\x3d\x85\x25\x4b\x5c\x62\x3e\x13\x65\xcc\xe6\xfd\x60\x88\x2d\x6a\x6f\x09\xd4\x04\x94\x51\x60\xc8\xca\x83\x2f\xcd\x68\x44\x64\xa5\xad\xd1\xd0\xd3\x84\xe9\x69\x4b\x3f\xda\x33\xac\x4e\x03\xf3\xeb\x6b\x3d\x1d\x1e\xbb\x56\x26\x56\x96\x7e\x1d\xac\x07\xb0\x11\xa6\xab\xd5\x07\x9d\x24\x06\xcf\x0f\xd6\xfb\xf1\x8c\xeb\xa7\xfa\x04\x30\x3f\x52\x0e\x83\xcd\x41\x6f\xe9\x27\x5e\x9a\x8d\xe3\xf9\x68\xd7\x2d\xce\xf5\x8c\x3c\xd7\x33\xda\xe7\x7a\x46\x58\xd7\x73\xeb\x7a\x0b\x89\x4a\x68\x4c\x66\x3b\xc2\xfa\x9b\x07\x88\x46\x40\x3a\x97\xcf\x4c\xe9\xca\x14\x46\xc1\x20\x3d\x8a\x85\xa1\x26\xc2\xde\xac\x4f\x40\x0b\xfa\x93\x7f\x87\xcd\xc4\x0d\xe3\xa9\xe1\xec\xa7\x92\xa5\x3f\x47\xa1\x79\x5b\xc7\x3d\x4a\x6d\xe6\x28\x7a\xed\x8c\xfc\xe6\xc1\x1e\x98\xca\x61\xd9\xe2\x55\x3c\x9a\xd7\x36\xb8\x1f\x2d\xc9\x84\x8a\xb7\x42\x23\x6c\xe3\x60\x85\xb6\x59\xf7\x1b\x75\xe2\x16\x7e\xa7\x50\xee\x0a\xdb\xf5\x8e\xbd\xc9\x2b\xfd\xb7\xf7\x90\x36\xc1\x91\xf6\x45\x70\x0b\x61\xd3\xd8\x8f\x61\xdd\x25\xf7\x3c\x5b\x91\x14\x34\x4d\x16\x90\x36\xb2\x73\xa9\xea\x6e\xb3\x56\x3f\x49\xd8\x0e\xa6\x94\x42\x2e\xb8\xc9\xce\x64\x4c\x6b\xc2\xc8\x0f\x5b\x0e\xe0\x88\x22\x0a\x06\x69\x83\x2a\x71\x51\x6d\xc4\x97\xdb\xb6\x66\xd0\x91\x2d\xcd\x83\x91\x1f\xb4\xec\x43\x35\xe6\xe8\xc7\xc0\x22\x63\x5d\xc7\xe5\x09\x26\xf7\x1c\xd9\xd1\xe2\xf8\x3a\x05\x08\xd9\xd5\xbc\xc5\x51\xbc\xa9\x38\xf7\x2c\x53\x9f\xa9\xf8\x31\x70\x48\xcb\x54\x36\xa5\x47\xbb\xdc\x48\x2e\x38\xd3\x42\x86\x1a\x8b\x09\xcd\xb2\xce\xf7\xd5\xf2\x51\x84\xca\x66\xdb\x69\xb3\x96\x25\x37\x99\xff\x1b\x22\xa4\x49\xd4\xab\xc7\x55\x5f\x47\x8e\x4f\x37\xd2\x6c\x75\xe9\x16\xc5\x21\xb2\x37\x6b\x9f\xee\xcd\xba\x43\xb8\x1d\xa5\x80\x34\x9a\x6e\xbf\x2b\xa4\xcf\xf1\xb4\x46\xa7\x80\x94\x94\x1c\xbe\x17\x76\x55\x66\x2b\x9f\x74\xcc\xc3\x7b\x68\xf2\xec\x60\x4f\x5f\x71\x21\xa1\x30\xc2\x88\x5f\x4c\x7e\xd7\x45\x0f\x94\x84\x0c\xa8\x3a\x0e\xaa\x89\x8d\x6d\x0d\xf3\x40\x7f\xe5\xd8\x8b\x89\xdf\xed\x8b\x89\xc7\x7e\x4c\xdc\x97\x29\xd9\x02\x6a\xb5\x37\x9a\xd2\x66\x47\xf7\x97\x08\x79\x8d\xfd\x08\x19\x83\x6a\x77\x64\x4f\x82\xc5\x66\x4d\x0e\xb6\xa7\x24\x7f\x0c\x20\x36\x73\x72\x80\xdb\x0d\x7c\x71\x88\xfe\x6e\x5e\x18\xf1\xe3\xe9\x05\x92\x31\x3a\x86\xd6\x0b\xe8\xb4\xe8\x29\x95\xb6\x45\xc9\x29\xe8\x17\x00\x4e\xfe\xfa\xdb\x56\xbf\xfe\xfa\x3b\x8a\x10\x6c\x94\xcd\x97\x79\xed\x5b\x4e\x4a\xc0\xc7\xe8\x48\xfb\x51\x55\x3d\x22\x35\xc3\x73\x5b\x5c\x94\x6e\x3f\xe5\xaf\xbf\xc9\x74\xa5\x21\x6e\xaa\x3d\x1f\xd7\x4f\x42\x2b\xef\x66\xcf\xc9\x10\x15\x87\xeb\xf9\xa8\x1e\xdc\xbb\x0f\x0f\x9b\xb5\xdd\x18\x8f\x16\xb3\x1f\x33\xf7\x63\x55\xdb\xe1\xf1\x58\x58\xa3\x64\x77\x5f\xce\xb4\xc8\x40\x52\x9e\x80\xb5\xac\xe4\x44\x3e\xb1\x96\xe9\x3f\x52\xf0\xb9\x47\x41\x0e\x7a\x21\x52\xa2\x57\x45\x8c\xfb\x1a\xfb\x41\x75\x0f\xfa\xeb\xbf\xfe\x26\x5f\xa9\xd4\xcc\x6e\xb9\xd4\x0d\x95\x8e\x6d\xff\x58\x24\x06\x19\x6b\xae\x5a\x64\xc1\xed\x6e\xdc\x29\xa0\x58\xb3\x65\x85\xbd\x59\x3b\xe3\x0c\x4b\x93\x44\xb0\x28\x43\xe9\x07\xdb\x3d\x88\xd5\x26\xbe\xdb\xe3\xa5\x19\xa1\x69\x2a\x41\xa9\x13\x14\x0b\x6b\x24\x9c\x62\x19\x0d\xb2\xfb\x63\xd4\x39\xa4\x9e\x84\xf8\x69\x01\xee\xd1\x9f\xcc\xb3\xd3\x72\x36\x33\x96\xdd\xed\xac\xa5\x54\xd3\x33\xa5\x85\xa4\x73\xf8\x99\xa8\x02\x12\x36\x63\x90\x92\xe9\xca\xda\x9e\xee\xc0\xcd\x8e\x1e\x4d\x74\x49\xb3\xfa\x53\x3b\xb2\x0d\xc8\xea\x8e\xd1\xd0\x5e\x5e\xbb\x97\xec\x9e\xef\xeb\xe9\x1b\xfb\xb1\x69\x38\xa5\xb2\x9b\x81\x75\x68\x46\xda\xc4\x31\x46\xec\xd8\x96\x62\x6e\x0c\x47\x4e\x99\x49\x1a\x88\x3a\x2d\x0a\x1d\xfb\x51\xe8\xe1\x73\x8d\x2f\xb4\xa8\x70\xed\xe6\xa2\x99\xbc\x18\x68\xbf\xe0\x1e\x96\xb1\xa7\x01\x0b\xaa\xc8\xd4\xf8\xff\xc8\x33\x76\x83\xb1\x5f\x7a\x3f\x16\x3a\xfa\xbc\xdd\x60\x8c\xed\x17\x6d\x25\x4e\xd3\xd4\x23\xc5\x5f\x60\x1e\xb1\x71\xbd\xd7\x3e\xbd\xc8\x3d\x9f\x96\x5e\x27\x1c\x04\xc9\xc6\x26\x74\x89\xdd\x3e\xb4\x52\x37\x9a\x9b\x78\xb3\xea\x41\xc8\x98\xb2\x0d\x28\xdd\x61\x03\x0b\xdd\xe7\x01\x5b\x8b\xe9\x8c\xeb\xd2\xf4\x24\x03\x2a\x23\x27\x1a\xeb\x33\xb6\x67\x5a\xc8\x5a\x80\x7d\x62\xa3\x4d\x63\x4a\xef\x5e\xbe\x4f\x0c\xd6\x9d\x1c\x4b\x4c\xc9\x9f\xb9\x78\xe1\x5b\xf3\xf8\x62\x1d\x43\xf8\xf0\x85\x4f\x19\x36\x2c\xde\xa6\x6c\xdf\x32\x68\x75\xa8\xa2\xaa\x93\x0d\x56\x9a\x64\x65\x88\xa0\x2e\xd0\x94\xbb\x47\x73\x94\xa6\xba\x54\xf1\x87\x63\x07\xbf\xfa\xc9\xf5\xd5\xc3\xed\xd3\xed\xd5\xe5\xe7\x40\x90\xfe\x7f\x8f\x4f\xd7\x77\xb6\x31\xe7\xeb\x75\xa0\x8b\x6e\x66\x42\x3e\xd2\x86\x7c\xf5\x41\x26\x92\x96\xe0\xfa\x7a\xad\x70\xfa\xfb\x19\x25\x14\x19\xb5\x4f\x98\x81\x9a\x65\x61\xdd\xe6\x59\xdd\xf1\x50\x09\xb2\x29\x39\x3c\xc0\xb4\x64\x59\x7a\x60\x23\xee\xc2\x2b\x3a\x5c\xec\x2b\x3a\x5c\x60\x8f\xed\x49\x87\x4e\x8e\x3b\x97\xba\x83\xe5\xad\x95\x9e\x2e\x05\x17\xa8\x24\xf6\x3c\x48\x61\xb5\x2c\x13\xc9\x73\xd4\xc4\x5f\x60\xaf\x6f\xd8\x01\xdd\xac\x99\x22\x25\x8f\xc7\xc5\x6e\x31\xd4\x82\x4d\x44\x5e\xd8\x76\xf0\xaa\x23\x63\x56\x66\x81\x3a\x1c\x06\x18\x9b\x3b\xbd\xae\xa1\x25\xa8\x32\xd3\x2d\x29\x95\xe2\xc6\x04\xf4\x17\xd8\xde\x89\x5e\xf0\xaa\xc2\x18\x07\x8e\x74\xa9\x36\xd4\x73\x70\x9a\xca\x39\x04\xfa\xf7\xf4\x62\x27\xfc\x74\x9d\xbd\xae\x89\x2f\x13\xee\xb0\x01\xe5\x2b\x52\xd4\x39\x18\x86\x3e\xa4\xbb\xe4\x82\x70\xa8\x7d\xbe\xa5\xd3\x27\xd0\x9d\xe2\xaf\x48\x4a\x61\x2e\x69\xf0\x28\xa5\x4f\x03\xd6\x7b\xde\xb4\x91\xa9\xad\xc7\x57\x94\xbc\xa9\xdb\x0a\x98\x72\x05\xef\xed\xd6\x45\x5c\xf9\xf4\x4a\xf0\x19\x9b\xe3\x5a\x0b\x26\x9e\x45\xdb\x77\xf5\x52\xf3\xf8\x41\x21\x27\x96\x06\x32\x63\xf6\x1e\x15\x9a\x92\x54\xf0\x98\x58\x7f\x82\xed\x31\x98\x83\xae\xfa\x6b\xa3\x91\x90\xfb\x7c\x2e\x71\x75\x58\xf1\x0d\x5e\x13\x6c\xcf\x41\x17\x2e\xb2\x6b\x67\x12\x68\x26\xeb\xeb\xdc\xac\xaf\x48\xaa\x26\xd0\xa6\x9c\xb1\x49\xf9\x24\xd0\x58\x76\x08\xf7\x19\x56\xf1\x78\x81\xdd\x99\x70\xab\x7b\x57\x3d\x6d\x27\x42\xac\x64\xfd\xfc\xf7\x30\x62\x75\x26\x2e\x9e\x4b\xec\x01\x86\x2d\xcc\x93\xea\x76\x13\xff\xb4\x2a\x42\xb2\x27\xf6\x63\x4d\x06\xd8\x13\x3b\x7f\x28\xc1\x49\x2a\x92\xda\x62\x8b\xe9\x1f\x90\x04\xfc\x4e\xad\x64\xad\xb1\xf8\x97\x5b\x5f\xf6\xf8\x96\xfd\xc0\xb5\xb1\x6f\x2b\xc6\x9b\x4e\xd1\x93\xfc\x2b\xcc\x56\xd0\x12\xdf\xd1\xa2\x38\x74\x8d\xd0\xd0\xbb\x29\x64\x78\xbe\x63\x82\xdb\x8d\x31\xd7\xc5\x27\x0e\x0d\x39\xf0\x86\xdc\xed\x55\x6e\x86\xfc\x74\xb5\xdf\x43\x0c\xbd\xe6\xb3\xe1\xbe\xe6\xb3\x21\xba\xf9\xec\xd3\x15\xd1\x92\xcd\xe7\x10\x95\xa8\x0e\xd1\xcd\x67\x9f\xae\x9c\x8b\x8d\x44\x41\x3a\x85\x4f\x57\x9b\x75\x9c\xe3\x19\xa2\xbb\xcf\x3e\x5d\x35\xb7\x0c\x44\xf6\xcc\x0c\xd1\xbd\x39\xbf\xb1\x44\xb3\xbc\x09\xd7\x13\xc1\x95\x96\x65\xa2\x63\x16\xf2\x10\xdd\xa9\xf3\x59\x50\x13\xb3\x2e\x41\x2a\x20\x39\x8d\x68\xd7\x19\xa2\xdb\x75\x2c\x96\x73\xae\xf6\x20\xf1\xb1\x97\x01\xdd\x81\xa6\x37\x8f\x5b\x8b\x67\xbc\xbb\x78\x46\xde\xda\x1e\xed\x5b\x3c\xa3\xc0\x2d\x40\x3d\xfb\xeb\x77\xee\x3f\x6d\x1d\x3b\x85\x69\x39\x9f\x47\x5d\x14\x37\x0a\x9c\xec\x08\xcb\xcc\x43\x35\x5c\x18\x63\xd8\x8d\xd0\xf1\xb8\x81\xcb\xa8\xc2\x29\xc5\x6d\x73\x09\x26\xb9\xbb\x79\xdc\x3a\xad\x78\x1c\x1e\xb2\x5a\x7d\x67\x0f\xa3\x9c\x06\x85\x4c\xd5\xbe\x55\x47\xb3\x4f\x01\x0b\x9c\x35\xe9\xdb\x3d\xbc\x79\x34\x3e\x32\xfa\xf6\xa3\x51\xe0\x8c\xc9\x7e\xa8\x13\xae\x5a\x1a\x05\xce\x96\xf4\x81\xed\x2a\x66\x64\xcc\x31\x0a\x1c\x2e\xe9\x81\xac\xcf\xbc\x16\x54\xd2\x1c\x3a\x77\x94\x1c\x07\x87\xdd\x17\x34\x56\xdf\xd6\x02\xa3\x50\xb0\x69\x69\x5d\x6f\x8c\x47\x3a\xbe\x62\xdb\xde\x70\x1b\x05\x78\x7c\x21\x76\x6a\xdb\xc3\xca\x08\xef\x32\x1a\xa2\x7b\x11\xb8\x20\x39\xa4\x8c\x5a\x6d\xbc\xbb\x79\x8c\x02\xc3\x36\x20\x54\xb7\x74\x6c\xdd\x83\x32\x65\x31\xae\x20\x70\xee\xe8\xa0\x3c\x4d\x9c\x1c\x05\x85\x35\x25\xac\x39\x5e\x7e\x8a\xb7\x09\x9c\x3c\x3a\xc8\x5a\x55\x2e\x36\xe9\x40\x14\x22\xd6\x98\xb4\x88\x33\x56\x6d\x15\x47\x63\x62\x2d\x4a\x0d\xd2\x24\x4d\x05\x84\xae\x75\xc1\x40\x1e\xbf\x67\x74\x3a\x9b\x58\x43\xa3\x25\x6b\xd5\xb4\x65\x3a\xda\xd7\xfa\x0d\xbe\x87\xb5\xc8\x9e\x3e\x38\x89\x5b\xac\xd9\xe9\x48\xd8\x65\xa9\x27\xa1\x62\xed\x4f\x8b\x5a\xdd\xba\x72\x0a\x2a\xba\xc5\xb6\xab\xc0\xb5\xe3\x8a\xdd\x52\x18\x0d\x03\x95\x9c\x43\xcc\xda\xaa\xc0\x49\xac\x1e\x6f\x91\x5c\xef\x57\x75\x9d\x64\x34\x2e\xd6\x2e\xbd\x48\xa1\xdb\x5b\xa9\xdc\x4d\x14\xda\x96\xeb\xed\x96\xa9\xec\xf8\x1d\x8f\x0e\xeb\xf9\x24\xcd\x49\xca\xd4\x33\x86\x28\xac\xe1\xea\xde\x56\x0d\xcf\xfd\x82\x38\x9a\x00\xac\x19\xb3\xf3\xfe\x0f\x09\x05\x6b\xe6\xfe\xd1\x99\xc2\x9a\x40\x97\xda\xf4\xdc\x9d\x7e\x34\x2a\xd6\x06\xd6\xb7\x5d\xfd\x28\x5c\xac\x15\xac\x43\xbe\x93\x11\xd1\xa7\x7a\xbd\x40\x25\x07\x4d\x49\xd5\x93\x16\x63\x18\xd0\xc7\x7b\x77\x9d\xdb\xc9\xc0\xe8\x73\xbe\x1d\xa0\x5a\xd3\xa3\xb3\x15\xf4\x81\xdf\xb0\xa0\x7b\x2e\x2f\xc3\x00\x1f\x6f\xeb\x6c\xf8\x72\x2a\xec\xf1\x81\x5a\x67\x7a\x4f\xc0\x45\xb7\xad\x0b\x4d\x80\xdb\x0b\x4a\x54\x41\x13\x68\x25\x1e\x05\x8b\xb5\x53\x34\xcb\x85\xd2\x64\x56\x06\x2e\x6a\xc4\xe0\x60\x2d\x53\x1d\xa5\x58\xa1\x46\x45\x29\x43\xac\x31\xb2\xf7\xc4\x52\x4d\x33\x31\x0f\x5d\x9a\x75\x04\xa4\x7f\x97\x62\x0f\x24\xf0\x3f\x4b\x28\xe3\x7f\x0d\x64\x34\xf4\x2f\x41\xdc\xaf\x9f\xd6\xb5\x2d\xa9\x64\xa2\x54\xc6\x0a\xa8\xb8\xc4\x70\x14\xd1\x4f\x4a\x97\x40\x18\x17\x69\x94\x62\x8e\x8e\xb7\x3b\x85\x28\xdc\x6d\xde\x94\x58\x29\x47\xc1\x1e\x6f\x75\x8a\x52\x2d\x6c\x33\x46\x34\x6a\x60\x43\xa8\xbf\x39\xec\xaa\xd5\xd8\x63\x2e\xdc\x6b\x8a\xf2\x8f\x2b\x75\x66\xfb\xd7\xb7\xea\xf2\x17\xbb\x65\x79\xef\xf0\xd8\x68\xfc\x6e\x4f\x59\x3e\x70\x78\xac\xa7\x2c\xef\x4c\xe4\xef\x96\x84\x08\x59\x05\x8e\x81\xf4\x00\xb9\xee\x99\xdf\x69\x1a\x8f\xe5\xe7\x59\x7b\x7f\x8e\xa4\xfa\x49\x8b\x63\xf0\x9a\x69\xf9\x2c\xe6\xfe\xcf\xa9\xec\xce\xc9\x85\x37\x27\xfb\xb6\x4a\xc6\x81\x0b\xe0\xc2\x1d\xbe\x2c\xd3\x20\x89\x5a\x71\x4d\xfb\xee\x26\x44\x48\xeb\x02\xdf\xe8\x5e\x21\xba\xd4\xd4\x95\x1c\x0a\xaa\x17\x51\xa8\xc8\x5d\x84\xb6\x77\x34\x13\xf3\x33\xfb\xa8\xcb\x8f\xf4\xd1\xc6\xbf\xfd\x45\x9c\xfb\xaf\x54\x2f\xce\xf8\x32\x9f\xed\xdf\x22\x1e\x7b\xbb\x5c\xe3\xb7\x7d\xdb\xcd\xd5\xa0\x33\x69\xaf\xd5\x3c\xd0\x6f\x39\xf6\x7e\x61\x67\xdc\xbb\x35\x5e\x0d\x3c\xa5\xc9\xf3\xe1\x71\xbd\xfd\xf1\xf1\xee\x8f\x73\x34\xe3\xba\x7e\xac\xdd\xf1\x76\x06\x9c\x78\x12\x98\xec\x6b\xa3\x9a\xa0\x7f\xed\xe3\xf1\xeb\xe5\xd5\xf5\xf6\x77\x68\xed\x99\x04\x7e\xe9\xe3\xa0\xbd\xaf\x73\xa7\xad\x1b\xfc\x3a\x99\x4c\x75\x21\x60\xe8\x7e\xc1\x3d\x74\x1c\xef\x5f\x0f\xd1\xb1\x60\xf3\x05\x48\x52\x48\x26\xec\x9d\x54\x7f\x88\x29\x51\x65\xb2\x20\x54\xd5\xbd\x74\x8c\x5b\x3f\x52\xf7\xd4\x32\x3e\x0f\x9d\xae\xd1\x3b\x17\x18\x52\xf2\xb2\x68\x8b\x09\x1d\x1e\x8e\x77\xd9\x28\x49\xd6\x21\x61\x73\x48\x63\x9f\x58\xdb\xcd\xe7\xc9\xff\x90\x9f\xee\xce\x9e\x84\xc8\xfe\xfd\xf1\xec\x72\x0e\x5c\xff\xec\xe9\xe8\xee\xef\x55\xbd\xf5\x7f\x93\xc6\x7e\xf4\xea\xff\x03\x00\x00\xff\xff\x6a\xd7\x2a\x3c\xd2\x71\x00\x00")

func ResourcesEventsYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesEventsYaml,
		"../resources/events.yaml",
	)
}

func ResourcesEventsYaml() (*asset, error) {
	bytes, err := ResourcesEventsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../resources/events.yaml", size: 29138, mode: os.FileMode(420), modTime: time.Unix(1593680945, 0)}
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
	"../resources/events.yaml": ResourcesEventsYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"events.yaml": &bintree{ResourcesEventsYaml, map[string]*bintree{}},
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

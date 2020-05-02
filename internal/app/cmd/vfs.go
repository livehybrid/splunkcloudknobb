// Code generated by vfsgen; DO NOT EDIT.

// +build release

package cmd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// CmdHelpEmbed statically implements the virtual filesystem provided to vfsgen.
var CmdHelpEmbed = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 5, 2, 15, 8, 57, 125504014, time.UTC),
		},
		"/backup": &vfsgen۰DirInfo{
			name:    "backup",
			modTime: time.Date(2020, 4, 26, 23, 29, 55, 505186519, time.UTC),
		},
		"/backup/backup.go": &vfsgen۰CompressedFileInfo{
			name:             "backup.go",
			modTime:          time.Date(2020, 5, 2, 14, 29, 49, 961200952, time.UTC),
			uncompressedSize: 1497,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x6f\x8b\xe3\x36\x10\xc6\x5f\x5b\x9f\x62\x6a\x38\xce\x3e\x7c\xb6\xf7\x7a\xbb\x85\x40\x5e\xa4\x4e\x96\x2d\x1c\xdb\x25\x9b\x72\xd0\xe3\xe8\xc9\xd2\xc8\x56\x23\x4b\x46\x92\x1b\x96\xf6\xbe\x7b\x91\x95\xec\xa6\x29\x85\x52\xe8\x2b\xeb\xcf\x68\xe6\xf9\x3d\x33\x1e\x29\xdb\xd3\x0e\xa1\xa5\x6c\x3f\x8d\x84\xc8\x61\x34\xd6\x43\x46\x92\xd4\xb8\x94\x24\xa9\x97\x03\xa6\x84\x24\x69\x27\x7d\x3f\xb5\x25\x33\x43\x75\xe8\xd1\xa2\x74\xfb\xc9\xfa\xca\x8d\x6a\xd2\x7b\xa6\xcc\xc4\xf7\xda\xb4\x6d\x35\xee\xbb\x8a\x19\x2d\x64\x77\xf9\xcc\x8d\xe2\xea\xdb\x8a\x99\xd6\xd2\x94\xe4\x84\x54\x15\x34\x03\x87\xde\x28\xee\xc0\xf7\x08\xf1\xd9\x64\xa9\x97\x46\x13\xff\x34\xe2\x1c\xe0\xbc\x9d\x98\x87\xdf\x49\xd2\xcc\x01\xf0\x26\x06\x96\x71\x4b\x92\xf5\xee\x11\x00\x42\x9c\xd4\x1d\xf9\x3a\x67\xbe\xc7\xc3\xf7\x33\x13\x58\xf4\x93\xd5\x0e\x28\x68\x3c\x40\x3c\x6c\x06\x4e\xc4\xa4\xd9\x4b\x58\xc6\xc4\x65\xe2\x1c\x32\x06\x6f\x9a\x81\xe7\xa1\x36\x83\x65\x48\x90\x85\x3d\x49\xd8\x31\x06\x96\xc0\x44\x17\xf6\x41\xc4\x12\x82\x5d\xe5\xbd\x39\x64\x79\x79\x6b\xec\x40\x7d\x96\xbe\xab\xeb\x9b\xfa\xaa\x7e\xb7\xbb\xba\xae\xdf\xd7\xd7\x69\x4e\x92\x28\x09\xd8\x51\xeb\x1d\xaa\xf1\x8e\x6a\xae\xd0\x02\x55\x0a\x8c\x98\xed\xd8\x6b\x73\x50\xc8\x3b\x04\xd3\xfe\x8a\xcc\x3b\xc8\x24\x96\x00\x9c\xba\xbe\x35\xd4\x72\x57\x80\xc5\xd0\x2f\x57\x80\x32\x26\xc0\x0a\xa9\xd0\x15\x50\x96\x79\xe4\x7b\x26\x38\xab\x91\xb1\x81\x07\xd4\xd6\xd2\xb2\x31\xc3\x40\x35\x2f\x80\xda\xce\xc1\xa7\xcf\xd1\xc3\x08\x7c\x44\x2c\x9b\x0f\x3f\x94\x8f\x9e\xa3\xb5\x21\xc9\x0e\x87\x51\x51\x8f\x59\xfa\x93\xa3\x1d\x46\xf7\xd2\x02\xb4\x54\xf9\x91\x67\xa5\xd4\xff\xc6\xb1\x52\xea\x5f\xeb\x5f\x9f\x0a\x84\x17\x31\x64\x6e\xdd\x76\x2e\x76\x71\xf8\x88\xd4\xb2\xfe\x4e\x3a\x6f\xec\xd3\xf9\xdd\xcc\xe4\xcd\xc4\xfa\x5b\xa9\x10\x0e\x52\x29\x98\x46\x4e\x3d\xce\x74\x71\x21\x07\x04\xe7\xe9\x30\x9e\xa0\x03\x00\x8c\xd4\x39\xe4\x17\x04\xcf\xa9\xb2\x10\xa3\xe9\x80\x05\x70\xef\xe0\x4c\xba\x2f\x00\x61\x71\x9c\xa6\x07\x6a\x1d\xc6\x31\x7a\x5b\x5f\xbd\x9d\x07\x69\x51\xbf\x5f\xd4\xd7\x3f\xd7\xdf\x2d\xea\x3a\x9d\x9f\xe7\x24\x91\x02\x10\xbe\x59\x86\x4e\x84\x24\x2f\x0d\xfc\x60\xba\xf2\x96\x7a\xaa\x44\xf6\x45\x84\xef\x02\x18\xd5\xda\x78\x18\x43\xee\x33\x84\xd7\xaf\xdc\xeb\x2f\xa7\x7c\x5f\x49\x82\xb0\x04\xe3\xca\xa6\x0f\xb7\xee\x4c\xb0\x2f\xc0\xff\xb7\x92\x0e\x3d\x28\xea\x3c\x0c\x86\x4b\x21\x91\x57\x94\x31\x86\x2e\x0a\x30\xe2\x6f\x1a\xe2\xcf\x12\xda\xf0\x57\x1f\x0f\x56\x7a\x6c\x8c\xf6\xa8\x7d\xc6\xe2\xf7\x68\x62\x01\x27\xa5\xe7\xae\x8a\x02\xd0\xda\xe0\xab\x71\xe5\x8f\x23\xea\x8b\x26\x84\xd3\x5f\x56\x0f\x0f\x9b\xfb\xf5\x1f\xf3\xba\xd9\x6e\x56\xbb\x4d\x5c\x6f\xd7\x1f\xb7\x05\xd4\x37\x37\x37\x47\x6c\x6b\xff\x11\x7c\x63\xad\xb1\x22\x4b\x31\x7c\xc1\x8c\xa8\xa5\xee\xe2\x44\x04\x38\x10\xc6\xce\xea\xa5\xee\x16\xf0\xea\xb7\xf4\x45\xee\xac\x30\x27\xc9\x09\x3a\xf0\x73\x14\x68\x41\x94\x8d\x32\x0e\xb3\x9c\x90\x44\x94\x1f\x03\x7b\xf6\xe9\x73\xfb\xe4\xf1\xc4\x9e\x87\xab\x67\xb3\xfe\x0c\x00\x00\xff\xff\x90\xe6\xfe\xcf\xd9\x05\x00\x00"),
		},
		"/backup/backup.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "backup.tmpl",
			modTime:          time.Date(2020, 4, 29, 2, 4, 55, 4717301, time.UTC),
			uncompressedSize: 329,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\xcb\x4a\xc5\x30\x10\x06\xe0\xfd\x3c\xc5\x50\x5c\xb6\x0b\x2f\x2b\xc1\x45\xab\x5d\x08\xf6\x02\xb5\x20\x88\x8b\xa4\x19\xdb\xd2\x9a\x84\x24\x2e\x4a\xe8\xbb\x8b\x69\x51\xce\x05\xce\x76\xfe\x6f\x98\x9f\x01\xef\x51\xd0\xe7\x28\x09\xa3\xd6\xb2\x9e\x32\xd6\x4d\xdf\x3a\xc2\x75\x05\xf0\x3e\x41\x47\x5f\x7a\x66\x8e\x30\x4a\xb5\x0e\x22\x64\x8f\x55\x51\xa4\xe5\xd3\x3d\x20\x22\xf2\xb0\x03\xd0\xb4\xd9\xc1\x5c\x30\x3b\x70\xc5\x8c\x88\xd1\x90\x56\xc6\xc5\x38\x8c\xd6\x29\xb3\x00\x54\xf5\xeb\x73\x55\x36\x1b\x4c\x92\x89\x96\x87\x77\xeb\xcc\x28\xfb\x0f\x04\xc8\xdf\xd2\xa2\x7e\xc9\xf7\xf8\x0a\x6d\x37\x49\xc5\x39\xb2\x79\xde\xb1\xa5\xce\x90\xbb\xbe\xb9\xbd\x0b\xe4\xc8\xfd\x5d\x3e\xaf\xff\xe1\xd6\xeb\x92\xda\x5b\x9f\x32\xf0\x9e\xa4\xf8\xfd\xd5\x4f\x00\x00\x00\xff\xff\x53\x21\x38\x73\x49\x01\x00\x00"),
		},
		"/backup/dashboard.go": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.go",
			modTime:          time.Date(2020, 5, 2, 14, 29, 49, 993201438, time.UTC),
			uncompressedSize: 2699,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xd1\x6e\xe3\xb6\x12\x7d\x96\xbe\x62\x56\x80\x61\x69\xa3\x4b\xef\xde\x16\x08\x60\x20\x0f\xa9\xd3\x00\x5b\xb8\xbb\x41\x8c\xa0\x0f\x4d\x51\x50\x14\x25\xb1\xa6\x48\x81\xa4\xa2\x0d\x8a\xbc\xf7\x1b\xfa\x79\xfd\x92\x62\x28\x59\x96\x13\x1b\x2e\xda\x17\x9b\xa2\xc8\x99\x73\xe6\x9c\xa1\xd8\x50\xb6\xa5\x25\x87\x8c\xb2\x6d\xdb\x84\xa1\xa8\x1b\x6d\x1c\xc4\x61\x10\x15\xb5\x8b\xc2\x20\xd2\x16\x7f\x1b\xea\xaa\x45\x21\x24\xc7\x41\x14\x86\x41\x54\x0a\x57\xb5\x19\x61\xba\x5e\x74\x15\x37\x5c\xd8\x6d\x6b\xdc\xc2\x36\xb2\x55\x5b\x26\x75\x9b\x6f\x95\xce\xb2\x45\xb3\x2d\x17\x96\x9b\x27\xc1\xf8\xeb\x7d\xb6\x29\x3e\x7e\xb3\x60\x3a\x33\x34\x0a\x93\x30\x5c\x2c\xe0\x86\xda\x2a\xd3\xd4\xe4\x90\x75\x42\xca\x01\x17\x50\x29\x41\x17\xe0\x2a\x0e\xf9\xb8\xc2\x69\xa0\x20\x35\xa3\x12\x10\x59\x58\xb4\x8a\x41\xcc\xe0\xfd\xaa\xce\x93\x7d\xa4\x98\xd5\x39\xbc\xf7\x59\xc8\x4a\xd7\x35\x55\x79\x0a\xd4\x94\x16\x7e\xfe\xc5\x3a\x23\x54\x99\xc0\xef\x61\xc0\xb4\x2a\x44\x09\xcb\x2b\x60\x64\xe5\xc7\x61\x20\x75\x3f\xe1\x1f\xc9\x5a\x97\x61\x60\x71\x62\xe0\x43\x3e\xf3\x6e\xd3\x0f\x63\xa9\xcb\x24\xdc\x45\x21\xab\xf5\x27\x72\x67\x74\xdd\xb8\x38\x7a\x54\x51\x72\xf4\xc5\xc6\x51\xe3\x84\x2a\x77\x24\x0b\x6d\xf6\xec\x2c\x21\xe4\x51\xf9\xbd\x61\x40\x5b\x57\xa5\xc0\x8d\xf1\xc9\x11\x88\x50\xf1\x10\xf1\xe1\x7e\x9d\xee\x10\x3e\x58\x6e\x14\xad\xf9\x38\x71\x47\xad\xed\xb4\xc9\xc7\x89\x95\xd6\x5b\xc1\xef\xb4\x71\xe9\x84\x56\x12\x06\xa2\xf0\xf1\xdf\x5d\x81\x12\x12\xeb\x81\xe4\xc9\x2d\x75\x54\x16\x71\x54\xe0\xff\x12\x36\x5e\x5c\x58\xa1\xba\x70\xdd\xba\x8a\x2b\x27\x18\x75\x42\xab\x25\xcc\x2e\x9e\x22\x0f\x32\x09\x83\x17\x5f\x3b\xf2\x49\x15\xba\x88\xa3\x4d\xcb\x18\xb7\xb6\x68\x25\x48\x84\x8e\xca\x09\x65\x1d\x55\x8c\x2f\x1f\x15\xc0\xcc\x46\x29\x20\x49\x64\x73\xb4\x58\x45\xed\xc8\xa6\x31\x42\xb9\x22\x8e\xfe\xfa\xe3\x4f\x78\x13\xb3\x13\xae\x82\xf9\xcc\xce\xdf\x44\x3f\x88\x3f\x56\x68\x4c\x07\x17\x10\x0d\x75\x66\x15\x55\x39\xd6\xb8\xa6\x5b\x1e\xe3\xd3\x28\xf5\x68\xa7\x24\x0c\x4a\x0d\xe8\xb5\x38\x19\xeb\x34\x10\xbd\xe7\xce\x08\xfe\x84\x9a\xee\x8d\xbc\x16\xd6\xab\x8c\xf2\x5e\xaf\xd7\x13\x89\x81\x10\x02\x68\x8e\x60\x2f\xad\xb0\x6e\xef\xdc\x5e\x77\x8f\x0a\x57\xbd\xd5\xe8\x98\x48\x4c\xb7\x32\x57\x73\x07\xa6\x47\x33\xb4\x0c\xd5\x98\x51\x0a\xeb\xec\xb2\x2f\x48\x2f\x15\x6a\xf5\x12\x9f\xf4\x2e\xd6\xfa\x96\x3b\x56\xf1\x7c\xd2\x7a\xf2\x2c\xa7\x30\x28\xb4\xcc\xb9\xa7\xb5\x3b\x37\xc8\x0f\x1a\x7d\x3b\xb4\x17\xf9\xd2\xba\xa6\x75\xb7\x7e\x59\x0a\xd1\x18\x03\x2b\x42\x33\x9b\xc2\xaf\x07\x9b\xaf\x33\x1b\xf7\x31\x93\x03\x77\xad\x0c\xa7\xd3\x36\xda\xa3\x1c\x10\xa0\xde\x39\x5a\x63\xc6\xe6\xe8\x03\x0c\xad\x2d\xb9\xa3\xae\xda\xf0\x86\x1a\xea\x34\x86\xd4\x96\xfc\xb8\xcd\x85\xb9\x96\x32\xf6\x6b\x3e\x5c\x5e\x5e\x22\x11\x43\xbb\x7f\xcf\x25\x85\xc8\xd0\x0e\x29\x19\xda\x9d\x60\x35\x26\xf8\x6f\xc4\x76\x09\xce\x70\xdb\x2d\xdb\xd1\x3b\xd3\x6c\xbd\x05\xbe\xe3\xa5\x50\x0a\xc1\x38\x0d\x9d\x11\x8e\x7b\x0a\x76\x07\xa5\x07\x01\xe4\x64\x7d\x31\xd3\x13\x35\xe0\xb4\xa3\x32\x7b\x76\xdc\xc2\x15\x7c\xe8\xe7\x98\x6e\x95\xeb\x1f\xfd\x19\x88\x05\x32\x54\x95\xbc\xb7\xbe\x77\xfa\x11\x73\x12\x6f\xb3\x00\x71\x60\x25\xbe\xd6\xf2\xad\x40\xc5\xa0\xc8\x01\xa5\x99\x25\x5f\x6b\x19\xa5\x90\x93\xcf\xb4\xe6\x09\x76\x01\x23\x9e\xd5\x4a\x2b\xc7\x95\x8b\x73\x32\x8c\x52\x98\xc4\xef\x17\x3a\xdd\xb2\xea\x56\x48\x1e\x4f\x5e\x61\xb0\x87\x26\xa7\x8e\xe7\x07\xa8\x7e\xb3\x5a\x9d\x86\xe5\xbd\xf1\x16\x1d\x6e\x3a\x03\xef\x9e\x76\xdf\x2b\x67\x9e\xf7\xf8\x70\xd3\x29\x80\xf8\xee\x35\xc2\x03\x25\x26\x0f\x17\x20\xb9\x9a\x64\x48\x16\x1f\x3f\xfc\xff\xdb\xe1\xa0\xbb\xe1\x59\x5b\x7a\x90\x29\xcc\xf2\x6d\x96\x82\x1f\xe1\x71\x72\x50\x8c\x63\x21\x10\xc0\x97\x4e\x21\xef\x09\x12\x94\xb6\x97\xbf\xff\xbf\x80\x8f\xf8\xfd\xd8\x77\xc2\xa1\x1d\x7f\x32\xda\x71\x98\xcf\xf2\xf9\xf4\xd0\x71\xfa\xd0\x88\x51\xda\x47\x3b\xed\xc6\x7f\x66\xfb\xe3\xd9\x76\xef\x44\x7f\x06\xfa\x3e\x78\x8d\xe0\x51\x9d\xc7\x70\xd0\xed\x37\x5a\x71\xb2\x4f\xb8\xcd\xe6\xbd\x28\x78\xef\x61\xbd\xec\x04\x96\xff\x4b\xa2\x74\x22\xd6\x79\x1e\x90\x6b\xc5\xdf\xe1\x46\xc0\xeb\xc4\xe4\xd3\xf9\xee\x30\x19\x65\x46\x5b\xfb\x8a\xeb\x78\x35\x41\x8e\x64\x17\x60\xa5\x55\x69\xa8\x6b\xa5\xff\xf2\x5b\x78\xd6\x2d\x28\xdd\x41\x45\x9f\xf8\x78\x25\xeb\x77\xfa\xcc\xfe\x1e\x33\x85\x3d\x54\xc6\x97\xc0\x70\xd7\x1a\x15\xbe\x84\x7f\x07\x00\x00\xff\xff\xbe\x78\x82\x54\x8b\x0a\x00\x00"),
		},
		"/backup/report.go": &vfsgen۰CompressedFileInfo{
			name:             "report.go",
			modTime:          time.Date(2020, 5, 2, 14, 29, 49, 993201438, time.UTC),
			uncompressedSize: 2651,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xc1\x6e\xe3\x36\x10\x3d\x4b\x5f\x31\x11\x60\x58\xda\xa8\x74\xd2\x16\x08\x60\x20\x87\xd4\xdb\x00\x5b\xb8\xbb\x41\x8c\xa0\x87\xa6\x28\x68\x8a\x92\x58\x53\xa4\x40\x52\xd1\x06\x45\xee\xfd\x86\x7e\x5e\xbf\xa4\x18\x52\x76\xa4\x8d\xd3\x14\xed\xc5\x96\x48\xce\xcc\x7b\xf3\x1e\x45\xb6\x94\xed\x68\xc5\x61\x4b\xd9\xae\x6b\xe3\x58\x34\xad\x36\x0e\xd2\x38\x4a\xca\xc6\x25\x71\x94\x68\x8b\xbf\x2d\x75\xf5\xa2\x14\x92\xe3\x43\x12\xc7\x51\x52\x09\x57\x77\x5b\xc2\x74\xb3\xb0\x6d\x79\xfe\xcd\x82\xe9\xad\xa1\xc9\x74\xa6\xaf\xb9\xe1\xc2\xee\x3a\xe3\x16\xb6\x95\x9d\xda\x31\xa9\xbb\x62\xa7\xf4\x76\xbb\x68\x77\xd5\xc2\x72\xf3\x20\x18\x4f\xe2\x2c\x8e\x17\x0b\xb8\xe5\xbe\x7c\x2f\xa4\x1c\x20\x01\x95\x12\x74\x09\xae\xe6\x60\xfc\xac\x05\xa7\x81\x82\xd4\x8c\x4a\x40\x48\x71\xd9\x29\x06\x29\x83\x77\xab\xa6\xc8\x86\x1c\x29\x6b\x0a\x78\xe7\x31\x91\x95\x6e\x1a\xaa\x8a\x1c\xa8\xa9\x2c\xfc\xfc\x8b\x75\x46\xa8\x2a\x83\xdf\xe3\x88\x69\x55\x8a\x0a\x96\x97\xc0\xc8\xca\x3f\xc7\x91\xd4\x61\xc0\xbf\x92\xb5\xae\xe2\xc8\xe2\xc0\x80\x95\x7c\xe4\xfd\x26\x3c\xa6\x52\x57\x59\xbc\xcf\x42\x56\xeb\x0f\xe4\xc6\xe8\xa6\x75\x69\x72\xaf\x92\xec\xe8\xc4\xc6\x51\xe3\x84\xaa\xf6\xfc\x4a\x6d\x06\xc8\x96\x10\x72\xaf\x7c\x60\x1c\xd1\xce\xd5\x39\x70\x63\x7c\x65\x44\x21\x54\x3a\xa4\xbb\xbb\x5d\xe7\x7b\x78\x77\x96\x1b\x45\x1b\x7e\x18\xb8\xa1\xd6\xf6\xda\x14\x87\x81\x95\xd6\x3b\xc1\x6f\xb4\x71\xf9\x88\x53\x16\x47\xa2\xf4\xf9\x4f\x2e\x41\x09\x89\xcd\x40\xe6\xe4\x9a\x3a\x2a\xcb\x34\x29\xf1\x7f\x09\x1b\xaf\x1a\xac\x50\x36\xb8\xea\x5c\xcd\x95\x13\x8c\x3a\xa1\xd5\x12\x66\xa7\x0f\x89\x07\x99\xc5\xd1\x93\x6f\x1c\xf9\xa0\x4a\x5d\xa6\xc9\xa6\x63\x8c\x5b\x5b\x76\x12\x24\x42\x47\xcd\x84\xb2\x8e\x2a\xc6\x97\xf7\x0a\x60\x66\x93\x1c\x90\x24\xb2\x39\xda\xa9\xb2\x71\x64\xd3\x1a\xa1\x5c\x99\x26\x7f\xfd\xf1\x27\xbc\xc8\xd9\x0b\x57\xc3\x7c\x66\xe7\x2f\xb2\x4f\xf2\x1f\x3a\x74\x28\x07\xa7\x90\x0c\x7d\x66\x35\x55\xbe\xc7\x0d\xdd\xf1\x14\xdf\x0e\x3a\x07\x55\xb2\x38\xaa\x34\xa0\xc5\xd2\xec\xd0\xa4\x81\xe5\x2d\x77\x46\xf0\x07\x54\x73\x90\x10\xd6\xc2\x7a\x75\x51\xd6\xab\xf5\xfa\xe0\x59\x42\x08\xa0\x23\xa2\x67\x49\x85\x75\x83\x57\x83\xd8\x1e\x0a\x2e\x79\x29\xcc\x31\x65\x98\xee\x64\xa1\xe6\x0e\x4c\x40\xb1\xdf\x1f\x20\x85\x75\x76\x19\x3a\x10\xb4\x41\x71\x8e\x9b\x14\xfb\x7a\xcd\x1d\xab\x79\x31\x0e\xff\x27\x06\x4f\x29\xf6\xad\xd4\xb2\xe0\x9e\xc9\xfe\xab\x40\x7e\xd0\x68\xd1\x61\x1b\x91\x4f\x9d\x6b\x3b\x77\xed\x97\xe5\x90\x84\x2c\x18\x4f\xb7\x36\x87\x5f\x27\x91\x57\x5b\x9b\x86\x84\xd9\xc4\x45\x2b\xc3\xe9\x78\xaf\x0c\x08\x87\xda\x28\x6a\x81\xfa\xcf\xd8\x1c\xc5\xc6\xbc\xda\x92\x1b\xea\xea\x0d\x6f\xa9\xa1\x4e\x63\x3e\x6d\xc9\x8f\xbb\x42\x98\x2b\x29\x53\xbf\xe6\xec\xe2\xe2\x02\x29\x18\xda\xff\x47\x16\xf8\x44\x7b\x24\x63\x68\xff\x0a\x9f\x43\xf6\xff\x41\x69\x9f\xfd\x0d\x56\xfb\x65\x7b\x62\x6f\xec\xa5\xa0\xfa\x77\xbc\x12\x4a\x21\x12\xa7\xa1\x37\xc2\x71\x8f\xdf\xee\xa1\x04\x10\x40\x5e\xed\x2c\x56\x7a\xa0\x06\x9c\x76\x54\x6e\x1f\x1d\xb7\x70\x09\x67\x61\x8c\xe9\x4e\xb9\xf0\x8a\x36\xf2\x1d\x36\x54\x55\x3c\x98\xdc\x7b\xfa\x88\x1f\x89\xdf\x92\x11\xe2\xc0\x4e\x7c\x6e\xe4\x4b\x69\xca\x41\x8b\x09\xa5\x99\x25\xe1\x78\x21\xee\x33\xca\x63\xc8\x47\xda\xf0\x0c\xad\xcf\x88\x27\xb7\xd2\xca\x71\xe5\x52\x43\x36\x9c\x1a\x56\xe7\x30\xaa\x12\xd6\x39\xdd\xb1\xfa\x5a\x48\x9e\x8e\xa6\x30\xd7\x5d\x5b\x50\xc7\x8b\x09\xb6\xdf\xac\x56\xaf\x83\xf3\xf6\x78\x89\x11\x83\xde\x40\x77\x4b\xfb\xef\x95\x33\x8f\xcf\xf8\x30\xe8\x35\x80\x38\xf7\x25\xc2\x89\x1e\xa3\x97\x53\x90\x5c\x8d\x2a\x64\x8b\xf3\xb3\xaf\xbf\x1d\x3e\x68\xef\xf9\xb6\xab\x3c\xc8\x1c\x66\xc5\x6e\x9b\x83\x7f\xc2\x4f\xc8\xa4\x19\xc7\x52\x20\x80\x4f\xbd\x42\xde\x23\x24\x28\x70\x30\x41\xf8\x3f\x85\x73\x3c\x24\x9e\x37\xc3\xd4\x94\x3f\x19\xed\x38\xcc\x67\xc5\x7c\x7c\xcc\x4f\xbc\x98\xe4\x21\xd5\xeb\x86\xfc\x77\xce\x3f\x52\x6a\x3f\x21\xc2\x97\xcf\xef\x83\x2f\xcb\xdf\xab\xb7\x01\x4c\xb6\xfa\x7b\xad\x38\x79\xae\xb6\xdb\xce\x83\x1c\x78\x9d\x61\x41\x70\x02\xcb\xaf\xb2\x24\x1f\xc9\xf4\x36\x09\x80\x42\x2b\x7e\x82\x91\x80\xd7\x85\xd1\xd1\x78\x32\xad\x46\x99\xd1\xd6\x8e\x99\x1e\x2e\x1d\xc8\x90\xec\xa3\x57\x5a\x55\x86\xba\x4e\xfa\x63\xdd\xc2\xa3\xee\x40\xe9\x1e\x6a\xfa\xc0\x0f\x37\xad\x10\xe9\xcb\xfa\x4b\xca\x18\xf4\xd0\x17\xdf\x00\xc3\x5d\x67\x54\xfc\x14\xff\x1d\x00\x00\xff\xff\x9c\xda\x63\x46\x5b\x0a\x00\x00"),
		},
		"/backup/searchhistory.go": &vfsgen۰CompressedFileInfo{
			name:             "searchhistory.go",
			modTime:          time.Date(2020, 5, 2, 14, 29, 49, 989201378, time.UTC),
			uncompressedSize: 2734,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xcd\x6e\xe3\x36\x10\x3e\x4b\x4f\x31\x2b\xc0\xb0\xb4\x51\xe9\xdd\xb6\xc0\x02\x06\x72\x48\x9d\x06\x4d\x91\xb6\x41\xdc\x45\x0f\x4d\x51\xd0\x14\x25\xb1\xa6\x48\x81\x1c\xc5\x1b\x2c\x72\xef\x33\xf4\xf1\xfa\x24\xc5\x50\xb2\x63\xad\xed\x66\x51\xf4\x62\x4b\xfc\x99\xf9\xbe\x6f\xbe\xa1\xd8\x72\xb1\xe6\x95\x84\x15\x17\xeb\xae\x8d\x63\xd5\xb4\xd6\x21\xa4\x71\x94\x94\x0d\x26\x71\x94\x58\x4f\xbf\x2d\xc7\x7a\x56\x2a\x2d\xe9\x21\x89\xe3\x28\xa9\x14\xd6\xdd\x8a\x09\xdb\xcc\x7c\x5b\xbe\xfd\x6a\x26\xec\xca\xf1\x64\x3c\xb3\xa9\xa5\x93\xca\xaf\x3b\x87\x33\xdf\xea\xce\xac\x85\xb6\x5d\xb1\x36\x76\xb5\x9a\xb5\xeb\x6a\xe6\xa5\x7b\x50\x42\x26\x71\x16\xc7\xb3\x19\x2c\x25\x77\xa2\xfe\x4e\x79\xb4\xee\x11\x36\x4a\x6b\x68\x3b\xad\x03\xbc\xb8\xec\x8c\x80\x54\xc0\xeb\x45\x53\x64\xe3\xa5\xa9\x68\x0a\x78\x1d\x10\xb0\x85\x6d\x1a\x6e\x8a\x1c\xb8\xab\x3c\xfc\xfa\x9b\x47\xa7\x4c\x95\xc1\xc7\x38\x8e\x84\x35\xa5\xaa\x60\x7e\x0e\x82\x2d\xc2\x73\x1c\x69\xdb\x0f\x84\x57\x76\x63\xab\x38\xf2\x34\x30\x40\x63\x3f\xca\xcd\xb2\x7f\x4c\xb5\xad\xb2\x5d\x14\xb6\xb8\xb9\x66\xb7\xce\x36\x2d\xa6\xc9\xbd\x49\xb2\xa3\x13\x4b\xe4\x0e\x95\xa9\x06\x85\xa1\xb4\x6e\x80\x0e\x03\x76\xc6\xd8\xbd\x09\xfb\xe3\x88\x77\x58\xe7\x20\x9d\x0b\x00\x08\x8c\x32\xe9\x10\xf5\xfd\xdd\x4d\xbe\x45\xf9\xde\x4b\x67\x78\x23\x77\x03\xb7\xdc\xfb\x8d\x75\xc5\x6e\x60\x61\xed\x5a\xc9\x5b\xeb\x30\xdf\xa3\x96\xc5\x91\x2a\x43\xfc\x57\xe7\x60\x94\x86\x8f\x71\x44\x02\xb0\x2b\x8e\x5c\x97\x69\x52\xd2\xff\x1c\x96\xa1\x56\xb0\xa0\x62\xc1\x45\x87\xb5\x34\xa8\x04\x47\x65\xcd\x1c\x26\x67\x0f\x49\x00\x99\xc5\xd1\x53\xd0\x8f\x5d\x9b\xd2\x96\x69\xb2\xec\x84\x90\xde\x97\x9d\x06\x4d\xd0\x01\x2d\x28\xe3\x91\x1b\x21\xe7\xf7\x06\x60\xe2\x93\x1c\x88\x24\xb1\x39\x2a\x58\xd9\x20\x5b\xb6\x4e\x19\x2c\xd3\xe4\xef\x3f\xff\x82\x83\x98\x1b\x85\x35\x4c\x27\x7e\x7a\x10\x7d\x14\x7f\xa7\xd0\x2e\x1d\x9c\x41\x32\xe8\x2c\x6a\x6e\x82\xc6\x0d\x5f\xcb\x94\xde\x76\xe5\x1e\xf9\xea\x4e\xfa\x4e\x63\x16\x47\x95\x05\xb2\x5f\x9a\xed\x14\x1b\x28\xdf\x49\x74\x4a\x3e\x50\x85\xc7\x65\x0d\x95\xbe\xb8\xb9\x01\x1f\x86\xa5\x07\xc6\x18\x90\x4b\xa2\xe7\xfa\x2a\x8f\x63\x1f\xf7\x06\x08\xf0\x68\xe5\x61\xb1\x8e\x55\x4b\xd8\x4e\x17\x66\x8a\xe0\x7a\x30\x72\xc8\x09\xf5\x00\x45\x2b\x8f\x7e\xde\xab\xd3\xd7\x8d\x0a\x77\xdc\xc7\xa4\xf9\x95\x44\x51\xcb\xe2\x58\x18\x22\x3a\x66\xb6\x9b\xde\xf2\x7b\x4a\x49\xe2\xd2\xea\x42\x06\x9e\xdb\x63\x83\x7d\x6f\xc9\xcd\x43\xe3\xb1\x9f\x3a\x6c\x3b\xbc\x0a\xcb\x72\x48\xfa\x60\xb4\x9f\xaf\x7c\x0e\xbf\x8f\x76\x5e\xac\x7c\xda\x07\xcc\x46\x86\x5b\x38\xc9\xf7\xbb\xeb\x13\x44\x03\x06\xf2\x41\x41\x96\x99\x88\x29\xf9\x83\xe2\x5b\xcf\x6e\x39\xd6\x4b\xd9\x72\xc7\xd1\x52\x5c\xeb\xd9\x0f\xeb\x42\xb9\x0b\xad\xd3\xb0\xe6\xcd\xbb\x77\xef\x88\x8a\xe3\x9b\xff\xc8\x26\x87\xc4\xf1\x0d\x91\x72\x7c\x73\x82\xd7\x2e\xfa\xff\x40\x6d\x9b\xe5\x05\x76\xdb\x65\x5b\x82\x2f\xb4\x61\x6f\x8a\x6f\x64\xa5\x8c\x21\x44\x68\x61\xe3\x14\xca\xc0\xc3\x6f\xa1\xf4\x20\x80\x9d\x54\x98\x32\x3d\x70\x07\x68\x91\xeb\xd5\x23\x4a\x0f\xe7\xf0\xa6\x1f\x13\xb6\x33\xd8\xbf\x92\xbb\x82\xd2\x8e\x9b\x4a\xf6\xbd\x10\xac\x7f\xc4\xae\x2c\x74\x73\x44\x38\x48\x89\x0f\x8d\x3e\x2c\x51\x39\xd4\x64\x44\x69\xe2\x59\xff\x3d\x62\xf8\x01\x49\xb8\xa1\xef\xaf\x2f\x33\xea\x0e\xc1\x02\xc1\x85\x35\x28\x0d\xa6\xdb\xd9\x1c\xf6\x32\xf5\xeb\xd0\x76\xa2\xbe\x52\x5a\xa6\x7b\x53\x14\xef\x67\xd5\xc8\x11\xb8\x3f\xbc\x35\xa7\xd1\x05\x9f\x1c\x82\xa4\x4d\x9f\x01\xef\x8e\x6f\xbe\x35\xe8\x1e\x9f\x01\xd2\xc6\x53\x08\x69\x6e\x04\x71\x54\x91\xbd\x97\x33\xd0\xd2\xec\x85\xcf\x66\x6f\xdf\x7c\xf9\xf5\x70\x00\x5e\xca\x55\x57\x05\x94\x39\x4c\x8a\xf5\x2a\x87\xf0\x44\x67\xcc\x48\x8a\x63\x21\x0e\x0e\xe9\x3d\x82\xa1\xd2\xbd\x1b\xfa\xff\x33\x78\x4b\x1f\x9a\xe7\xee\x18\xbb\xf3\x17\x67\x51\xc2\x74\x52\x4c\xc7\x3d\xa2\x3e\xb5\x66\x92\xf7\x01\x4f\xfb\xf3\xf3\x1a\xe1\xdf\x12\x6e\x57\xa8\xfe\xa0\x0c\xfd\x81\x76\x8c\xe3\xde\xbc\x8c\x64\x74\x14\x5c\x5a\x23\xd9\x73\xda\xf5\x6a\xda\x17\x09\x6c\x49\xdf\x76\xf2\x00\x83\xf9\x17\x59\x92\xef\x15\xef\x65\x36\x00\x85\x35\xf2\x15\xed\x04\xba\x81\xec\x7d\x6d\x5f\x8d\xb3\x71\xe1\xac\xf7\x27\x28\x6f\x2f\x36\xc4\x95\x6d\xe3\x2c\xac\xa9\x1c\xc7\x4e\x87\x3b\x83\x87\x47\xdb\x81\xb1\x1b\xa8\xf9\x83\x04\x0e\xda\x0a\xae\x87\x9d\x01\x40\xb8\x01\xed\xc3\x1f\x14\x0a\x52\x38\x89\x9d\x33\xf1\x53\xfc\x4f\x00\x00\x00\xff\xff\xe6\x2b\xc8\x6e\xae\x0a\x00\x00"),
		},
		"/restore": &vfsgen۰DirInfo{
			name:    "restore",
			modTime: time.Date(2020, 4, 18, 4, 49, 17, 392213227, time.UTC),
		},
		"/restore/restore.go": &vfsgen۰CompressedFileInfo{
			name:             "restore.go",
			modTime:          time.Date(2020, 5, 2, 14, 29, 49, 993201438, time.UTC),
			uncompressedSize: 164,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcb\x41\x8e\xc3\x20\x0c\x85\xe1\x75\x7c\x0a\x2b\xcb\x59\x84\x43\x64\x6e\xd0\x5e\x20\x18\x97\x20\x02\x46\xc6\xa8\xaa\xaa\xde\xbd\x4a\xb3\x7c\x7a\xff\xd7\x36\xca\x5b\x64\x54\xee\x26\xca\x00\xa9\x34\x51\xc3\x39\x26\xdb\x87\x5f\x48\x8a\x7b\xee\xac\x9c\x7a\x1e\x6a\xae\xb7\x63\xd4\x4c\x87\x8c\x90\xab\x78\xef\x5a\x8e\x8e\xa4\x3e\x52\x9c\x01\x9c\x5b\x4b\xc0\xc0\x9d\x34\x79\xee\x68\x3b\x23\x49\x29\x5b\x0d\x60\xaf\xc6\x78\xde\xdd\x74\x90\xe1\x1b\xa6\xf5\xe7\xf0\xef\xf2\xcb\x35\x61\xfa\xbf\xdf\x10\xf1\xec\x52\x8d\xf0\x81\x6f\x00\x00\x00\xff\xff\x61\xe6\x9c\xf2\xa4\x00\x00\x00"),
		},
		"/restore/restore.tmpl": &vfsgen۰FileInfo{
			name:    "restore.tmpl",
			modTime: time.Date(2020, 4, 18, 4, 49, 17, 392213227, time.UTC),
			content: []byte(""),
		},
		"/scknobb.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "scknobb.tmpl",
			modTime:          time.Date(2020, 4, 29, 1, 46, 35, 677669212, time.UTC),
			uncompressedSize: 1929,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x5b\x6f\x1a\x39\x14\x7e\xf7\xaf\xf8\x52\x55\x4a\x91\x02\xa3\x36\xdd\x17\xb4\x17\x25\xe4\xaa\x4d\x20\x6a\x48\x57\x2b\x40\x23\xc3\x1c\x18\x2f\x1e\x7b\xe2\x0b\x2c\x02\xfe\xfb\xca\x9e\x09\x24\xd9\x54\xed\x43\x8d\x74\x38\xbe\x7c\xdf\x39\xf6\xb9\xcc\x7a\x8d\x8c\xa6\x42\x11\xde\x9d\x94\xe5\x15\xf1\x8c\xcc\x3b\x6c\xb7\xd8\x8d\x14\xaf\x46\xfa\x4c\x32\xa4\x69\x8a\x20\x36\xd8\x44\x3d\x4c\xa3\xb2\xc1\x66\x27\xc3\xc9\x04\x69\xba\x89\x02\x9b\x04\x09\x0e\x53\x0c\x91\x20\xc5\x70\x13\xf5\x4a\x02\xeb\x75\xeb\x0b\x49\xe2\x96\xce\xb8\xa3\xed\x96\x0d\xd3\xb0\xfe\x21\x00\x81\x5f\x83\x9d\xf0\xfb\x90\x36\x22\xf9\x5e\xee\x81\x5f\xc9\x58\xa1\xd5\x76\x0b\xb6\x49\xd3\x34\x19\x06\xff\xd2\xcd\x30\x1d\x06\xe3\x41\x49\xd3\x64\x93\xb6\x76\x32\x5a\xbd\x14\xee\x8a\xdb\x3c\xa0\x5e\x5f\xf9\x9b\x83\xfd\x8e\xfb\x52\x7a\x35\x47\x47\x6a\x9f\xe1\xcf\xae\x5e\x4a\xca\x66\x84\xde\xe9\x3f\x34\x71\x38\xe5\x93\xb9\x2f\xc1\xd8\x89\x42\xe7\xe6\x1a\x42\x39\x32\x53\x3e\xa1\xa0\x69\xb8\x9c\x5e\x12\x78\x2b\xd4\x0c\x97\xfa\x80\xb1\x0b\xa1\x32\x14\xda\x84\xa3\x53\x6d\x0a\xee\x84\x56\xe0\xae\x1d\xdd\xcb\x9d\x2b\x6d\x3b\x49\x66\xc2\xe5\x7e\xdc\x9a\xe8\x22\x59\xe6\x64\x48\xd8\xb9\x37\x2e\xb1\x93\xb9\xd2\xe3\x71\xc2\xd6\x6b\x90\xca\xb0\xdd\x32\xf6\x32\xd6\x0f\x96\xcf\x28\x84\x9a\x45\xad\x62\xad\x61\x18\x5c\xde\xf4\x4e\x4f\x6e\x46\x18\x74\x7a\xb7\xb7\x27\xdd\xb3\x11\x06\xf7\x0f\xa7\xfb\xc9\x49\xa7\x7f\xdd\xeb\xa2\xd5\x6a\x8d\x30\xe8\xdd\x85\xc9\xfd\x88\xb1\x4b\xa9\xc7\x5c\xa2\x57\x06\x5f\x6d\xc5\xf9\x95\xcc\x58\x5b\xe1\x56\xed\xfa\x5d\x9b\x4d\x49\x0b\x92\xbf\x1d\xef\x9e\xf1\x9e\x9c\x8d\x6f\xa1\xbd\x2b\xbd\xc3\xe2\x09\x82\x78\x12\xca\x17\x64\xc4\x84\x4b\xb9\xc2\x20\xa3\x29\xf7\xd2\x8d\x76\x6c\x56\x48\x52\xee\x08\x68\xda\x27\x36\x48\x3d\x9b\x09\x35\x4b\x6a\xc2\x8a\x66\x10\xff\x3e\xee\x91\x8f\x5e\x50\x00\xa2\xf9\xf8\x23\xc8\x4f\x7b\x64\x08\xc9\xd1\xb3\x44\xf8\x0e\xf2\xb8\xf9\x3f\xaf\x33\x1a\xfb\x59\xb4\xbd\xf8\x11\x86\xcf\x7b\xa4\x33\x7c\x42\x15\x32\x42\x7b\xd5\x69\xa7\x71\xdf\x3f\xeb\x3d\xf4\xc1\x55\x16\x66\x52\xcf\x30\x15\x92\x6a\x86\x5f\x46\xdf\x49\x86\xf3\x7f\x79\x51\xca\x2a\x27\x3a\xba\x28\xb8\xca\xea\x10\x8e\xab\x2c\x0e\xa3\x4e\xe8\xb9\x7a\x4a\x74\x3d\x0e\x89\x6e\x19\xab\xe1\x35\xe4\xfd\x2e\x97\x6a\x30\x97\xf2\xed\x8d\x8c\xdb\x7c\xac\xb9\xc9\xde\xde\x36\x54\x6a\xe3\xde\xde\xcb\x85\x75\xda\xac\x18\xbb\xd0\xa6\x2a\x95\x9c\x64\xf9\xda\x81\xb0\x56\x23\xe2\xce\x37\x9e\xe1\xce\xe8\xa2\x74\x5d\x5a\x76\xb4\x52\x53\x31\xf3\x26\x56\x5c\x7c\x8e\xe6\xcf\x1d\xac\xab\xd1\xd1\xcf\x8c\xe0\x22\xc4\xe9\x42\x7b\x95\xb5\x71\xb8\x5e\xb7\xc2\x5c\xf1\x82\xb6\xdb\xc3\x9f\x6e\x7c\x70\x1d\x8a\x4d\x58\xac\xb4\x37\x98\x0a\x63\x1d\x8c\x57\x47\x28\xc9\xe4\xbc\xb4\x7f\xe0\xaf\xf3\x9b\x4e\xef\xf6\xfc\xe0\x00\xed\xc6\x88\xb1\xbe\xc6\x44\x2b\x15\xfa\xd9\x4a\x7b\x14\xde\x3a\x94\x46\x2f\x44\x46\x15\xc7\xc7\xc6\xcb\x2e\x26\x94\x75\x5c\x4d\x08\x0f\x5f\x6e\xf0\xa9\x01\x6f\xc9\x84\xdb\xc4\xcc\x64\xc7\x0d\x94\xdc\xda\xa5\x36\x59\x0b\xfd\x9c\x2c\x61\xc1\xa5\x27\x0b\x6e\x08\xa4\x26\x66\x55\x3a\xca\xb0\x14\x2e\xc7\xe7\x46\x6c\x0c\x71\x4d\x63\x4e\x2b\x2c\x73\x52\x08\x71\x0f\x9d\x52\x28\x16\xb6\x83\x01\x8b\xf7\x57\xba\xa0\x33\x61\x5a\x8c\x75\x7b\xfd\xf3\x36\xfe\xd6\x1e\x4b\x21\x25\x72\xbe\xa0\xaa\xc1\xc4\xce\x14\xaa\x43\x69\x07\x1b\x96\x2b\xea\xc0\x1c\xbc\x7b\xba\x97\x70\xc8\x7c\x34\x51\xd7\x5d\x69\xf4\xcc\xf0\x02\x42\x2d\xf4\xa4\x0a\x5a\x74\xf0\xb0\xd9\x9c\xd3\xea\x10\xdc\x82\xa3\xe4\x86\x17\xe4\x28\xb8\x70\x17\xbf\x45\x3b\xc2\x60\xdd\xd0\xa3\x17\x86\xb2\xfa\xbe\x6d\xb6\x4f\xc5\xff\x02\x00\x00\xff\xff\x1d\x74\x20\xeb\x89\x07\x00\x00"),
		},
		"/vfs_mock.go": &vfsgen۰CompressedFileInfo{
			name:             "vfs_mock.go",
			modTime:          time.Date(2020, 5, 2, 15, 1, 3, 893284788, time.UTC),
			uncompressedSize: 694,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\xcf\x6e\x9c\x30\x1c\x84\xcf\xeb\xa7\x98\x72\xc9\xa2\x12\xb6\xe7\x48\x39\xa5\x8d\xaa\x5e\x9b\x9e\x57\x06\x0f\x60\xc5\x7f\x90\xfd\x63\x57\xab\xaa\xef\x5e\x19\x96\xaa\xb9\xd8\xc6\x1a\x7f\x8c\x3f\x9f\x4e\xf8\xdc\x2d\xd6\x19\x7c\x4a\x74\xd4\x99\x4a\x9d\x4e\x78\x9b\x6c\xc6\xac\xfb\x77\x3d\x12\x36\x23\x06\x77\xc3\x92\x69\x60\x96\x64\xc3\x08\xc3\x0b\x5d\x9c\x3d\x83\x40\x22\x5c\xd4\x06\x32\x11\x89\x39\x2e\xa9\x67\xc6\x90\xa2\x87\xb1\xf9\xbd\x2d\xc0\x5f\xb9\x9c\x1a\x23\x46\x06\x26\x2d\xc4\xd5\x3a\x87\x3e\xb1\xac\x35\x1e\x84\x7e\x76\x5a\x78\xde\x03\xed\x18\x1f\xb6\x50\xa6\xac\xec\xb7\x7b\xe4\x35\x3a\xc3\x04\x63\x87\x81\x89\x41\xdc\x4d\xa9\xbd\x6b\xef\x8d\x52\xd6\xcf\x31\x09\x8e\xea\x50\x05\xca\x69\x12\x99\x2b\x75\xa8\x66\x2d\x53\x99\xd3\x12\xc4\x7a\x56\xaa\x5e\xef\xfa\xe2\xcd\x77\xba\xf9\x9b\xef\x68\x60\xfd\xec\x58\x6e\x95\xd7\x5f\x96\xa3\x18\xac\x63\xbe\x65\xa1\x6f\xd0\x2d\xb2\xfa\xb8\x30\x25\x6b\x0c\x03\xae\x53\x19\x58\x48\x9b\xc8\xab\x95\x09\xa2\xc7\x8c\xe3\x18\xef\x7b\x8f\xeb\xf7\xdd\x70\x0d\x29\x7a\x0b\x16\xd7\x18\x1e\x04\x1d\xd7\x9c\xac\xfc\x42\xda\x6d\xe4\xff\x75\x6c\x36\x3a\xb6\xea\xa2\xd3\xc7\xda\xa5\x67\xfb\x6a\x1d\x7f\xae\x3d\x95\x1a\x96\xd0\xc3\x06\x2b\xc7\x1a\xbf\xd5\x61\x7f\xd2\x40\x9a\x5c\x1e\xac\xe3\xa6\x35\x42\x07\xe8\x2e\x47\xb7\x08\x31\x6c\x66\x8b\xa8\x06\x39\xe2\x4a\x18\x26\x7b\x21\xac\xb4\x78\x7a\xac\xd5\xe1\xdc\xac\xc5\x83\xf6\x6c\x70\x6e\x70\xc6\xd3\x33\xee\x46\xdb\x17\xed\x1c\xd3\xf1\x4b\xad\xd4\xe1\x43\xbf\xe7\xad\xe1\x57\x9b\x8e\x05\xde\xfe\x88\x36\x6c\xab\xb2\xb5\x03\xeb\x06\x55\xdb\x56\xff\x46\x3d\xcf\x65\xea\xbd\xa9\xea\x82\x4c\x94\x25\x05\xf5\x47\xfd\x0d\x00\x00\xff\xff\x0c\xf9\x0c\x68\xb6\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/backup"].(os.FileInfo),
		fs["/restore"].(os.FileInfo),
		fs["/scknobb.tmpl"].(os.FileInfo),
		fs["/vfs_mock.go"].(os.FileInfo),
	}
	fs["/backup"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/backup/backup.go"].(os.FileInfo),
		fs["/backup/backup.tmpl"].(os.FileInfo),
		fs["/backup/dashboard.go"].(os.FileInfo),
		fs["/backup/report.go"].(os.FileInfo),
		fs["/backup/searchhistory.go"].(os.FileInfo),
	}
	fs["/restore"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/restore/restore.go"].(os.FileInfo),
		fs["/restore/restore.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}

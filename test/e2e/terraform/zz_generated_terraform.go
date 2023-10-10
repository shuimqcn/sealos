// Code generated by go-bindata. DO NOT EDIT.
// sources:
// infra/.gitignore
// infra/providers.tf
// infra/terraform.tf
// infra/vars.tf
// infra/vars.tf.amd64
// infra/vars.tf.arm64
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len
	return b, nil
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

var _infraGitignore = "\x2a\x2e\x74\x66\x73\x74\x61\x74\x65\x2a\x0a\x2e\x2e\x2f\x2e\x74\x65\x72\x72\x61\x66\x6f\x72\x6d\x2e\x6c\x6f\x63\x6b\x2e\x68\x63\x6c\x0a\x2e\x74\x65\x72\x72\x61\x66\x6f\x72\x6d\x0a"

func infraGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_infraGitignore,
		"infra/.gitignore",
	)
}

func infraGitignore() (*asset, error) {
	bytes, err := infraGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infra/.gitignore", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infraProvidersTf = "\x70\x72\x6f\x76\x69\x64\x65\x72\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x22\x20\x7b\x0a\x20\x20\x72\x65\x67\x69\x6f\x6e\x20\x3d\x20\x76\x61\x72\x2e\x72\x65\x67\x69\x6f\x6e\x0a\x7d\x0a"

func infraProvidersTfBytes() ([]byte, error) {
	return bindataRead(
		_infraProvidersTf,
		"infra/providers.tf",
	)
}

func infraProvidersTf() (*asset, error) {
	bytes, err := infraProvidersTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infra/providers.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infraTerraformTf = "\x72\x65\x73\x6f\x75\x72\x63\x65\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x76\x70\x63\x22\x20\x22\x76\x70\x63\x22\x20\x7b\x0a\x20\x20\x76\x70\x63\x5f\x6e\x61\x6d\x65\x20\x20\x20\x3d\x20\x22\x24\x7b\x76\x61\x72\x2e\x6e\x61\x6d\x65\x7d\x2d\x76\x70\x63\x22\x0a\x20\x20\x63\x69\x64\x72\x5f\x62\x6c\x6f\x63\x6b\x20\x3d\x20\x22\x31\x37\x32\x2e\x31\x36\x2e\x30\x2e\x30\x2f\x31\x32\x22\x0a\x7d\x0a\x0a\x72\x65\x73\x6f\x75\x72\x63\x65\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x76\x73\x77\x69\x74\x63\x68\x22\x20\x22\x76\x73\x77\x22\x20\x7b\x0a\x20\x20\x76\x73\x77\x69\x74\x63\x68\x5f\x6e\x61\x6d\x65\x20\x3d\x20\x22\x24\x7b\x76\x61\x72\x2e\x6e\x61\x6d\x65\x7d\x2d\x76\x73\x77\x22\x0a\x20\x20\x76\x70\x63\x5f\x69\x64\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x76\x70\x63\x2e\x76\x70\x63\x2e\x69\x64\x0a\x20\x20\x63\x69\x64\x72\x5f\x62\x6c\x6f\x63\x6b\x20\x20\x20\x3d\x20\x22\x31\x37\x32\x2e\x31\x36\x2e\x30\x2e\x30\x2f\x32\x31\x22\x0a\x20\x20\x7a\x6f\x6e\x65\x5f\x69\x64\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x7a\x6f\x6e\x65\x0a\x7d\x0a\x0a\x72\x65\x73\x6f\x75\x72\x63\x65\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x22\x20\x22\x64\x65\x66\x61\x75\x6c\x74\x22\x20\x7b\x0a\x20\x20\x6e\x61\x6d\x65\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x6e\x61\x6d\x65\x0a\x20\x20\x76\x70\x63\x5f\x69\x64\x20\x3d\x20\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x76\x70\x63\x2e\x76\x70\x63\x2e\x69\x64\x0a\x7d\x0a\x0a\x72\x65\x73\x6f\x75\x72\x63\x65\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x5f\x72\x75\x6c\x65\x22\x20\x22\x61\x6c\x6c\x6f\x77\x5f\x61\x6c\x6c\x5f\x74\x63\x70\x22\x20\x7b\x0a\x20\x20\x74\x79\x70\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x69\x6e\x67\x72\x65\x73\x73\x22\x0a\x20\x20\x69\x70\x5f\x70\x72\x6f\x74\x6f\x63\x6f\x6c\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x74\x63\x70\x22\x0a\x20\x20\x6e\x69\x63\x5f\x74\x79\x70\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x69\x6e\x74\x72\x61\x6e\x65\x74\x22\x0a\x20\x20\x70\x6f\x6c\x69\x63\x79\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x61\x63\x63\x65\x70\x74\x22\x0a\x20\x20\x70\x6f\x72\x74\x5f\x72\x61\x6e\x67\x65\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x31\x2f\x36\x35\x35\x33\x35\x22\x0a\x20\x20\x70\x72\x69\x6f\x72\x69\x74\x79\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x31\x0a\x20\x20\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x5f\x69\x64\x20\x3d\x20\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x2e\x64\x65\x66\x61\x75\x6c\x74\x2e\x69\x64\x0a\x20\x20\x63\x69\x64\x72\x5f\x69\x70\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x30\x2e\x30\x2e\x30\x2e\x30\x2f\x30\x22\x0a\x7d\x0a\x0a\x0a\x72\x65\x73\x6f\x75\x72\x63\x65\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x69\x6e\x73\x74\x61\x6e\x63\x65\x22\x20\x22\x73\x65\x61\x6c\x6f\x73\x22\x20\x7b\x0a\x20\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x5f\x6e\x61\x6d\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x24\x7b\x76\x61\x72\x2e\x6e\x61\x6d\x65\x7d\x2d\x73\x65\x61\x6c\x6f\x73\x22\x0a\x20\x20\x63\x6f\x75\x6e\x74\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x31\x0a\x20\x20\x61\x76\x61\x69\x6c\x61\x62\x69\x6c\x69\x74\x79\x5f\x7a\x6f\x6e\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x7a\x6f\x6e\x65\x0a\x20\x20\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x73\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x5b\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x2e\x64\x65\x66\x61\x75\x6c\x74\x2e\x69\x64\x5d\x0a\x20\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x5f\x74\x79\x70\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x65\x63\x73\x5f\x74\x79\x70\x65\x0a\x20\x20\x73\x79\x73\x74\x65\x6d\x5f\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x0a\x20\x20\x73\x79\x73\x74\x65\x6d\x5f\x64\x69\x73\x6b\x5f\x73\x69\x7a\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x35\x30\x0a\x20\x20\x70\x61\x73\x73\x77\x6f\x72\x64\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x65\x63\x73\x5f\x70\x61\x73\x73\x77\x6f\x72\x64\x0a\x20\x20\x69\x6d\x61\x67\x65\x5f\x69\x64\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x69\x6d\x61\x67\x65\x5f\x69\x64\x0a\x20\x20\x76\x73\x77\x69\x74\x63\x68\x5f\x69\x64\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x76\x73\x77\x69\x74\x63\x68\x2e\x76\x73\x77\x2e\x69\x64\x0a\x20\x20\x69\x6e\x74\x65\x72\x6e\x65\x74\x5f\x6d\x61\x78\x5f\x62\x61\x6e\x64\x77\x69\x64\x74\x68\x5f\x6f\x75\x74\x20\x3d\x20\x31\x30\x30\x0a\x20\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x5f\x63\x68\x61\x72\x67\x65\x5f\x74\x79\x70\x65\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x50\x6f\x73\x74\x50\x61\x69\x64\x22\x0a\x20\x20\x73\x70\x6f\x74\x5f\x73\x74\x72\x61\x74\x65\x67\x79\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x53\x70\x6f\x74\x41\x73\x50\x72\x69\x63\x65\x47\x6f\x22\x0a\x7d\x0a\x0a\x72\x65\x73\x6f\x75\x72\x63\x65\x20\x22\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x69\x6e\x73\x74\x61\x6e\x63\x65\x22\x20\x22\x6b\x75\x62\x65\x22\x20\x7b\x0a\x20\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x5f\x6e\x61\x6d\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x24\x7b\x76\x61\x72\x2e\x6e\x61\x6d\x65\x7d\x2d\x6b\x75\x62\x65\x2d\x24\x7b\x63\x6f\x75\x6e\x74\x2e\x69\x6e\x64\x65\x78\x7d\x22\x0a\x20\x20\x63\x6f\x75\x6e\x74\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x36\x0a\x20\x20\x61\x76\x61\x69\x6c\x61\x62\x69\x6c\x69\x74\x79\x5f\x7a\x6f\x6e\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x7a\x6f\x6e\x65\x0a\x20\x20\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x73\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x5b\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x73\x65\x63\x75\x72\x69\x74\x79\x5f\x67\x72\x6f\x75\x70\x2e\x64\x65\x66\x61\x75\x6c\x74\x2e\x69\x64\x5d\x0a\x20\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x5f\x74\x79\x70\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x65\x63\x73\x5f\x74\x79\x70\x65\x0a\x20\x20\x73\x79\x73\x74\x65\x6d\x5f\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x0a\x20\x20\x73\x79\x73\x74\x65\x6d\x5f\x64\x69\x73\x6b\x5f\x73\x69\x7a\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x35\x30\x0a\x20\x20\x70\x61\x73\x73\x77\x6f\x72\x64\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x65\x63\x73\x5f\x70\x61\x73\x73\x77\x6f\x72\x64\x0a\x20\x20\x69\x6d\x61\x67\x65\x5f\x69\x64\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x76\x61\x72\x2e\x69\x6d\x61\x67\x65\x5f\x69\x64\x0a\x20\x20\x76\x73\x77\x69\x74\x63\x68\x5f\x69\x64\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x61\x6c\x69\x63\x6c\x6f\x75\x64\x5f\x76\x73\x77\x69\x74\x63\x68\x2e\x76\x73\x77\x2e\x69\x64\x0a\x20\x20\x69\x6e\x74\x65\x72\x6e\x65\x74\x5f\x6d\x61\x78\x5f\x62\x61\x6e\x64\x77\x69\x64\x74\x68\x5f\x6f\x75\x74\x20\x3d\x20\x30\x0a\x20\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x5f\x63\x68\x61\x72\x67\x65\x5f\x74\x79\x70\x65\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x50\x6f\x73\x74\x50\x61\x69\x64\x22\x0a\x20\x20\x73\x70\x6f\x74\x5f\x73\x74\x72\x61\x74\x65\x67\x79\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x22\x53\x70\x6f\x74\x41\x73\x50\x72\x69\x63\x65\x47\x6f\x22\x0a\x7d\x0a\x0a"

func infraTerraformTfBytes() ([]byte, error) {
	return bindataRead(
		_infraTerraformTf,
		"infra/terraform.tf",
	)
}

func infraTerraformTf() (*asset, error) {
	bytes, err := infraTerraformTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infra/terraform.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infraVarsTf = "\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x6e\x61\x6d\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x6b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x74\x66\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x7a\x6f\x6e\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6e\x2d\x68\x6f\x6e\x67\x6b\x6f\x6e\x67\x2d\x62\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x72\x65\x67\x69\x6f\x6e\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6e\x2d\x68\x6f\x6e\x67\x6b\x6f\x6e\x67\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x69\x6d\x61\x67\x65\x5f\x69\x64\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x75\x62\x75\x6e\x74\x75\x5f\x32\x32\x5f\x30\x34\x5f\x78\x36\x34\x5f\x32\x30\x47\x5f\x61\x6c\x69\x62\x61\x73\x65\x5f\x32\x30\x32\x33\x30\x36\x31\x33\x2e\x76\x68\x64\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x65\x63\x73\x5f\x74\x79\x70\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x65\x63\x73\x2e\x67\x37\x2e\x78\x6c\x61\x72\x67\x65\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x65\x63\x73\x5f\x70\x61\x73\x73\x77\x6f\x72\x64\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x46\x61\x6e\x75\x78\x23\x31\x32\x33\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6c\x6f\x75\x64\x5f\x65\x73\x73\x64\x22\x0a\x7d\x0a"

func infraVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_infraVarsTf,
		"infra/vars.tf",
	)
}

func infraVarsTf() (*asset, error) {
	bytes, err := infraVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infra/vars.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infraVarsTfAmd64 = "\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x6e\x61\x6d\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x6b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x74\x66\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x7a\x6f\x6e\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6e\x2d\x68\x6f\x6e\x67\x6b\x6f\x6e\x67\x2d\x62\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x72\x65\x67\x69\x6f\x6e\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6e\x2d\x68\x6f\x6e\x67\x6b\x6f\x6e\x67\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x69\x6d\x61\x67\x65\x5f\x69\x64\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x75\x62\x75\x6e\x74\x75\x5f\x32\x32\x5f\x30\x34\x5f\x78\x36\x34\x5f\x32\x30\x47\x5f\x61\x6c\x69\x62\x61\x73\x65\x5f\x32\x30\x32\x33\x30\x36\x31\x33\x2e\x76\x68\x64\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x65\x63\x73\x5f\x74\x79\x70\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x65\x63\x73\x2e\x67\x37\x2e\x78\x6c\x61\x72\x67\x65\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x65\x63\x73\x5f\x70\x61\x73\x73\x77\x6f\x72\x64\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x46\x61\x6e\x75\x78\x23\x31\x32\x33\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6c\x6f\x75\x64\x5f\x65\x73\x73\x64\x22\x0a\x7d\x0a"

func infraVarsTfAmd64Bytes() ([]byte, error) {
	return bindataRead(
		_infraVarsTfAmd64,
		"infra/vars.tf.amd64",
	)
}

func infraVarsTfAmd64() (*asset, error) {
	bytes, err := infraVarsTfAmd64Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infra/vars.tf.amd64", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infraVarsTfArm64 = "\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x6e\x61\x6d\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x6b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x74\x66\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x7a\x6f\x6e\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x61\x70\x2d\x73\x6f\x75\x74\x68\x65\x61\x73\x74\x2d\x31\x61\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x72\x65\x67\x69\x6f\x6e\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x61\x70\x2d\x73\x6f\x75\x74\x68\x65\x61\x73\x74\x2d\x31\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x69\x6d\x61\x67\x65\x5f\x69\x64\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x75\x62\x75\x6e\x74\x75\x5f\x32\x32\x5f\x30\x34\x5f\x61\x72\x6d\x36\x34\x5f\x32\x30\x47\x5f\x61\x6c\x69\x62\x61\x73\x65\x5f\x32\x30\x32\x33\x30\x37\x31\x32\x2e\x76\x68\x64\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x65\x63\x73\x5f\x74\x79\x70\x65\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x65\x63\x73\x2e\x67\x36\x72\x2e\x78\x6c\x61\x72\x67\x65\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x65\x63\x73\x5f\x70\x61\x73\x73\x77\x6f\x72\x64\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x46\x61\x6e\x75\x78\x23\x31\x32\x33\x22\x0a\x7d\x0a\x0a\x76\x61\x72\x69\x61\x62\x6c\x65\x20\x22\x64\x69\x73\x6b\x5f\x63\x61\x74\x65\x67\x6f\x72\x79\x22\x20\x7b\x0a\x20\x20\x64\x65\x66\x61\x75\x6c\x74\x20\x3d\x20\x22\x63\x6c\x6f\x75\x64\x5f\x65\x73\x73\x64\x22\x0a\x7d\x0a"

func infraVarsTfArm64Bytes() ([]byte, error) {
	return bindataRead(
		_infraVarsTfArm64,
		"infra/vars.tf.arm64",
	)
}

func infraVarsTfArm64() (*asset, error) {
	bytes, err := infraVarsTfArm64Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infra/vars.tf.arm64", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"infra/.gitignore":    infraGitignore,
	"infra/providers.tf":  infraProvidersTf,
	"infra/terraform.tf":  infraTerraformTf,
	"infra/vars.tf":       infraVarsTf,
	"infra/vars.tf.amd64": infraVarsTfAmd64,
	"infra/vars.tf.arm64": infraVarsTfArm64,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"infra": {nil, map[string]*bintree{
		".gitignore":    {infraGitignore, map[string]*bintree{}},
		"providers.tf":  {infraProvidersTf, map[string]*bintree{}},
		"terraform.tf":  {infraTerraformTf, map[string]*bintree{}},
		"vars.tf":       {infraVarsTf, map[string]*bintree{}},
		"vars.tf.amd64": {infraVarsTfAmd64, map[string]*bintree{}},
		"vars.tf.arm64": {infraVarsTfArm64, map[string]*bintree{}},
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
//go:build web
// +build web

package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/GopeedLab/gopeed/cmd"
	"github.com/GopeedLab/gopeed/pkg/rest/model"
)

//go:embed dist/*
var dist embed.FS

func main() {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

	args := parse()
	var webBasicAuth *model.WebBasicAuth
	if isNotBlank(args.Username) && isNotBlank(args.Password) {
		webBasicAuth = &model.WebBasicAuth{
			Username: *args.Username,
			Password: *args.Password,
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := &model.StartConfig{
		Network:        "tcp",
		Address:        fmt.Sprintf("%s:%d", *args.Address, *args.Port),
		Storage:        model.StorageBolt,
		StorageDir:     filepath.Join(dir, "storage"),
		ApiToken:       *args.ApiToken,
		ProductionMode: true,
		WebEnable:      true,
		WebFS:          sub,
		WebBasicAuth:   webBasicAuth,
	}
	cmd.Start(cfg)
}

func isNotBlank(str *string) bool {
	return str != nil && *str != ""
}

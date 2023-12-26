package fsadmin

import (
	"errors"
	"os"
	"time"

	"github.com/opentdp/go-helper/filer"
)

type FilerParam struct {
	Action string
	Path   string
	File   filer.FileInfo
}

func Filer(data *FilerParam) ([]*filer.FileInfo, error) {

	var (
		err   error
		files []*filer.FileInfo
	)

	switch data.Action {
	case "ls":
		files, err = filer.List(data.Path)
	case "read":
		info := &filer.FileInfo{}
		info, err = filer.Detail(data.Path, true)
		files = []*filer.FileInfo{info}
	case "write":
		err = filer.Write(data.Path, data.File.Data)
		if err == nil && data.File.Mode > 0 {
			err = os.Chmod(data.Path, data.File.Mode)
		}
		if err == nil && data.File.ModTime > 0 {
			mtime := time.Unix(data.File.ModTime, 0)
			err = os.Chtimes(data.Path, mtime, mtime)
		}
	case "chmod":
		err = os.Chmod(data.Path, data.File.Mode)
	case "chtime":
		mtime := time.Unix(data.File.ModTime, 0)
		os.Chtimes(data.Path, mtime, mtime)
	case "mkdir":
		err = os.MkdirAll(data.Path, 0755)
	case "rm":
		err = os.RemoveAll(data.Path)
	case "mv":
		err = os.Rename(data.Path, data.File.Name)
	default:
		err = errors.New("暂不支持该操作")
	}

	return files, err

}

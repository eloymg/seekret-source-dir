package sourcedir 

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/apuigsech/seekret"
	"github.com/apuigsech/seekret/models"
)

var (
	SourceTypeDir = &SourceDir{}
)

const (
	Type = "seekret-source-dir"
)

type SourceDir struct{}

type SourceDirLoadOptions struct {
	Hidden    bool
	Recursive bool
}

func prepareDirLoadOptions(o seekret.LoadOptions) SourceDirLoadOptions {
	opt := SourceDirLoadOptions{
		Hidden:    false,
		Recursive: false,
	}

	if hidden, ok := o["hidden"].(bool); ok {
		opt.Hidden = hidden
	}
	if recursive, ok := o["recursive"].(bool); ok {
		opt.Recursive = recursive
	}

	return opt
}

func (s *SourceDir) LoadObjects(source string, opta seekret.LoadOptions) ([]models.Object, error) {
	var objectList []models.Object

	opt := prepareDirLoadOptions(opta)

	firstPath := true

	filepath.Walk(source, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			if firstPath {
				firstPath = false
				return nil
			}
			if strings.HasPrefix(filepath.Base(path), ".") && !opt.Hidden {
				return filepath.SkipDir
			}

			if !firstPath && !opt.Recursive {
				return filepath.SkipDir
			}
		} else {
			if !strings.HasPrefix(filepath.Base(path), ".") || (strings.HasPrefix(filepath.Base(path), ".")  && opt.Hidden) {
				f, err := os.Open(path)
				if err != nil {
					return err
				}
				
				content, err := ioutil.ReadAll(f)
				if err != nil {
					return err
				}

				o := models.NewObject(path, Type, "file-content", content)
		
				objectList = append(objectList, *o)

				f.Close()
			}
		}

		return nil
	})

	return objectList, nil
}

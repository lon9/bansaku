package libs

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/flosch/pongo2.v3"
)

type Template struct {
	tmplMap map[string]*pongo2.Template
}

type Options struct {
	// Directory to load templates. Default is "templates"
	Directory string
	// Extensions to parse template files from. Defaults to [".tmpl", ".html"]
	Extensions []string
}

type Context echo.Context

func C(c *echo.Context) *Context {
	return (*Context)(c)
}

func prepareOptions(opt Options) Options {
	// Defaults
	if len(opt.Directory) == 0 {
		opt.Directory = "templates"
	}
	if len(opt.Extensions) == 0 {
		opt.Extensions = []string{".tmpl", ".html"}
	}

	return opt
}

func preCompile(opt Options) *Template {
	tmplMap := make(map[string]*pongo2.Template)
	readDirs(opt.Directory, opt.Extensions, tmplMap)
	return &Template{tmplMap}
}

func readDirs(path string, ext []string, templates map[string]*pongo2.Template) *map[string]*pongo2.Template {
	dirPath := filepath.Dir(path)
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		// If fileInfo is directory, recurse the directory
		if fileInfo.IsDir() {
			readDirs(filepath.Join(dirPath, fileInfo.Name())+"/", ext, templates)
		}
		for _, s := range ext {
			if isMatched, _ := regexp.MatchString(".*"+s+"$", fileInfo.Name()); isMatched {
				t, err := pongo2.FromFile(filepath.Join(path, fileInfo.Name()))
				if err != nil {
					log.Fatalf("\"%s\": %v", fileInfo.Name(), err)
				}
				templates[strings.Replace(fileInfo.Name(), s, "", -1)] = t
			}
		}
	}
	return &templates
}

func PrepareTemplates(option Options) *Template {
	return preCompile(prepareOptions(option))
}

func (t *Template) Render(w io.Writer, templateName string, data interface{}) error {
	dataMap := data.(map[string]interface{})
	template, exist := t.tmplMap[templateName]
	if !exist {
		return errors.New("template " + templateName + " not found")
	}
	return template.ExecuteWriter(dataMap, w)
}

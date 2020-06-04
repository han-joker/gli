package gen

import (
	"io/ioutil"
	"strings"
)

func GenGoMod(module, version string) error {
	code := tplGoMod
	file := "go.mod"
	code = strings.Replace(code, "<module>", module, -1)
	code = strings.Replace(code, "<version>", version, -1)
	if err := ioutil.WriteFile(file, []byte(code), 0644); err != nil {
		return err
	}
	return nil

}
var tplGoMod = `module <module>

go <version>
`

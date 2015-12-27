package keepalived

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"regexp"
)

func Load(filename string) []byte {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	dir := path.Dir(filename)

	re := regexp.MustCompile("include (.+)")
	cb := func(s []byte) []byte {
		files, err := filepath.Glob(path.Join(dir, re.FindStringSubmatch(string(s))[1]))
		if err != nil {
			log.Fatal(err)
		}
		var buf []byte
		for _, f := range files {
			buf = append(buf, Load(f)...)
		}
		return buf
	}
	return re.ReplaceAllFunc(body, cb)
}

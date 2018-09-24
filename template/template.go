package template

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"log"
)

func Template() (name string, tpl []byte) {
	name = "Template"
	errCheck := func(err error) {
		if err != nil {
			log.Panic(err)
		}
	}

	raw, err := base64.StdEncoding.DecodeString(``)
	errCheck(err)

	r, err := gzip.NewReader(bytes.NewReader(raw))
	errCheck(err)

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, r)
	errCheck(err)

	tpl = buf.Bytes()

	return
}

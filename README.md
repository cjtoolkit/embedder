# Embedder

Generate constant variable by globbing any kind of text file such as html or sql. The values are
gzip compressed and base64 encoded (using standard encoding, not URL encoding).  This is useful
for keeping web app to a single file, therefore simplifying deployment.

For other frontend stuff you may also want to also consider [ZipFS](https://github.com/cjtoolkit/zipfs)
as it's better suited for code generated by a task runners or other automated system such as grunt, gulp,
webpack, rollup or [taskforce](https://github.com/cjtoolkit/taskforce) 😉.

Have fun! 😄

## Installation

```sh
$ go get github.com/cjtoolkit/embedder
```

## Usage

```sh
$ embedder DestPackageName, DestFilename, GlobPattern...
```

Also add the following to anywhere in your codebase, so you can decode the constant values, make sure the package name
is correct and you can change it to your taste.

```go
package embedderTools

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"log"
)

func DecodeValue(value string) []byte {
	raw, err := base64.StdEncoding.DecodeString(value)
	errCheck(err)

	r, err := gzip.NewReader(bytes.NewReader(raw))
	errCheck(err)

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, r)
	errCheck(err)

	return buf.Bytes()
}

func errCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}
```

## Example

### Input

resources/example1.html
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
</body>
</html>
```

resources/example2.html
```html
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
```

generate.go
```go
//go:generate embedder example const.go resources/*.html

package example
```

and than run `$ go generate`

### Output

const.go
```go
// Code generated by Embedder. DO NOT EDIT.

package example

const (
	// Source: resources/example1.html
	Example1 = `H4sIAAAAAAAC/3yQO08EIRDHez7FSO+S6yxmaTxtvQILSzzGXZI5SBYsDOG7G2A1WxgbHv/Hjwfe
nV8ezdvlCdZ8Yy2wTcA2LLOkIJtA1mkBAIA3yhauq90S5Vm+muf7B7lb2WcmbdqIamwEqtHF9+i+
GumkS5kudqGeqxXVetICP3lAStlsWAgmE11MtXZxGP4DpnMMdBD7sezhyjalWboYSDb+L5u9PhCI
0x/t/wrB7fmfNap2U1T7c1T/se8AAAD//2oeexlBAQAA`
	// Source: resources/example2.html
	Example2 = `H4sIAAAAAAAC/7LJMLSrrtYLSExPDcksyUmtrbXRzzC047IpzbHjUlBQUKiuLkrMS09V0AvJT8kv
rq0FC0IkMtMU9Fzy81KRBEHAJidTITknsbjYViklPy9VCWQ+3OycTDskE1JzirHoxqchLwWqHsa2
0S/NsQMEAAD//wbsY17GAAAA`
)
```

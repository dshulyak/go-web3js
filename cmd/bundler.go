package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	src = flag.String("src", "web3.js", "web3.js source file")
	dst = flag.String("dst", "", "destination file. leave empty to print into stdout")
	pkg = flag.String("pkg", "web3js", "package of the destionation file")
)

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*src)
	must(err)
	b := bytes.NewBuffer([]byte{})
	fmt.Fprintf(b, "// Code is generated. DO NOT EDIT.\n\n")
	fmt.Fprintf(b, "package %s\n\n", *pkg)
	fmt.Fprintf(b, "// Web3CODE is a binary representation of web3js mini.\n")
	fmt.Fprintf(b, "var Web3CODE = []byte(")
	fmt.Fprintf(b, "`%s`", string(data))
	fmt.Fprintf(b, ")\n")
	if *dst == "" {
		io.Copy(os.Stdout, b)
	} else {
		must(ioutil.WriteFile(*dst, b.Bytes(), 0644))
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

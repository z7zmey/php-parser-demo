// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample helloworld is a basic App Engine flexible app.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/visitor"
)

var port *int

func main() {

	var port = flag.Int("port", 8080, "listen port")
	flag.Parse()

	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	http.HandleFunc("/parse", parseHandler)
	http.Handle("/", http.FileServer(http.Dir(binDir+"/www")))
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Printf("Listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func parseHandler(w http.ResponseWriter, r *http.Request) {
	var p parser.Parser
	var comments comment.Comments
	var positions position.Positions

	src := bytes.NewBufferString(r.FormValue("script"))

	if r.FormValue("php5") == "true" {
		p = php5.NewParser(src, "input.php")
	} else {
		p = php7.NewParser(src, "input.php")
	}

	p.Parse()

	if r.FormValue("positions") == "true" {
		positions = p.GetPositions()
	}

	if r.FormValue("comments") == "true" {
		comments = p.GetComments()
	}

	for _, e := range p.GetErrors() {
		io.WriteString(w, e.String()+"\n")
	}

	nsResolver := visitor.NewNamespaceResolver()
	p.GetRootNode().Walk(nsResolver)

	dumper := visitor.Dumper{
		Writer:     w,
		Indent:     "",
		Comments:   comments,
		Positions:  positions,
		NsResolver: nsResolver,
	}
	p.GetRootNode().Walk(dumper)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

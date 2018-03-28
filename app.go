// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample helloworld is a basic App Engine flexible app.
package main

import (
	"os"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {

	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	http.HandleFunc("/parse", parseHandler)
	http.Handle("/", http.FileServer(http.Dir(binDir + "/www")))
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func parseHandler(w http.ResponseWriter, r *http.Request) {
	var nodes node.Node
	var comments comment.Comments
	var positions position.Positions

	src := bytes.NewBufferString(r.FormValue("script"))
	phpVersion := r.FormValue("phpVersion")

	if phpVersion == "php5" {
		nodes, comments, positions = php5.Parse(src, "input.php")
	} else {
		nodes, comments, positions = php7.Parse(src, "input.php")
	}

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := visitor.Dumper{
		Writer:     w,
		Indent:     "",
		Comments:   comments,
		Positions:  positions,
		NsResolver: nsResolver,
	}
	nodes.Walk(dumper)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

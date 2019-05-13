/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"

	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log := logf.Log.WithName("httpserver")

	io.WriteString(w, "Hello remote world!")
	log.Info("[GET /] Hello remote world!")
}

func main() {
	logf.SetLogger(logf.ZapLogger(false))
	log := logf.Log.WithName("httpserver")

	log.Info("Setting up server ...")
	portNumber := "9000"
	http.HandleFunc("/", hello)

	log.Info(fmt.Sprintln("Server listening on port ", portNumber))
	http.ListenAndServe(":"+portNumber, nil)
}

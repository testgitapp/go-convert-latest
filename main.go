// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/drone/go-convert/command"
	"github.com/google/subcommands"
	"net/http"
)

func runBinaryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters for two string inputs
	query := r.URL.Query()
	param1 := query.Get("input1")
	param2 := query.Get("input2")

	if param1 == "" || param2 == "" {
		http.Error(w, "Both input1 and input2 are required", http.StatusBadRequest)
		return
	}
	flag.Parse()
	ctx := context.Background()

	var output []byte
	output = new(command.Drone).ExecuteCommand(ctx, param2)

	// Send the output back to the client
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func main() {
	subcommands.Register(new(command.Drone), "")
	subcommands.Register(new(command.Github), "")
	subcommands.Register(new(command.Gitlab), "")
	subcommands.Register(new(command.Jenkins), "")
	subcommands.Register(new(command.Travis), "")
	subcommands.Register(new(command.Downgrade), "")
	subcommands.Register(new(command.JenkinsJson), "")
	subcommands.Register(new(command.JenkinsXml), "")

	http.HandleFunc("/run", runBinaryHandler)

	// Start the server
	fmt.Println("Server starting on :8990...")
	if err := http.ListenAndServe(":8990", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

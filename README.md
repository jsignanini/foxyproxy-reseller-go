# foxyproxy-reseller-go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/jsignanini/foxyproxy-reseller-go)
[![Build Status](https://travis-ci.org/jsignanini/foxyproxy-reseller-go.svg?branch=master)](https://travis-ci.org/jsignanini/foxyproxy-reseller-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsignanini/foxyproxy-reseller-go)](https://goreportcard.com/report/github.com/jsignanini/foxyproxy-reseller-go)
[![Coverage Status](https://coveralls.io/repos/github/jsignanini/foxyproxy-reseller-go/badge.svg?branch=master)](https://coveralls.io/github/jsignanini/foxyproxy-reseller-go?branch=master)

FoxyProxy Reseller API bindings for Go.


## Installation

Install foxyproxy-reseller-go with:
```sh
go get -u github.com/jsignanini/foxyproxy-reseller-go
```

Then, import it using:
```go
import "github.com/jsignanini/foxyproxy-reseller-go"
```


## Usage

```go
package main

import (
	"fmt"

	"github.com/jsignanini/foxyproxy-reseller-go"
)

func main() {
	client := foxyproxy.NewClient(&foxyproxy.NewClientParams{
		DomainHeader:    "x-domain-header-provided-by-foxyproxy",
		EndpointBaseURL: "https://reseller.test.api.foxyproxy.com",
		Username:        "foxyproxy-username",
		Password:        "foxyproxy-password",
	})

	// get all nodes
	nodes, err := client.GetAllNodes(0, 10)
	if err != nil {
		panic(err)
	}

	// iterate nodes and print active connections
	for _, node := range nodes {
		actives, err := node.GetActiveConnectionTotals()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Node: %s, Total Active Connections: %d\n", node.Name, actives)
	}
}

```

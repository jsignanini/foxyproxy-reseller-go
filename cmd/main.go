package main

import (
	"fmt"

	"github.com/jsignanini/foxyproxy-reseller-go"
)

func main() {
	client := foxyproxy.NewClient(&foxyproxy.NewClientParams{
		Username:        "5kZNZg",
		Password:        "cb7GtGBcGavR5KxBYvgkXjpp5DAE3px3",
		DomainHeader:    "ghostery",
		EndpointBaseURL: "https://reseller.test.api.foxyproxy.com",
	})
	// accounts, err := client.GetAllNodes(0, 10)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, a := range accounts {
	// 	a2, err := a.GetAccountsByNode(0, 10)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%+v %+v\n", *a, a2[0])
	// }
	c, err := client.ActivateAccounts("mdTXgIq4Qo0Mu9Ay")
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func main() {
	sess := session.New()

	account := services.GetAccountService(sess)
	user, err := account.GetCurrentUser()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Acting as %s\n", *user.Username)

	for _, arg := range os.Args[1:] {
		vlanID, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Could not interpret %s as valid VLAN id\n", arg)
			continue
		}

		printSubnet(sess, vlanID)
	}
}

func printSubnet(sess *session.Session, vlanID int) {
	service := services.GetNetworkVlanService(sess)
	network := service.Id(vlanID)
	subnets, _ := network.GetSubnets()

	for _, subnet := range subnets {
		fmt.Printf("%s/%d\n", *subnet.NetworkIdentifier, *subnet.Cidr)
	}
}

func contains(s []string, e int) bool {
	for _, a := range s {
		if a == strconv.Itoa(e) {
			return true
		}
	}
	return false
}

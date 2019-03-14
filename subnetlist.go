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
	user, err := services.GetAccountService(sess).GetCurrentUser()

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
	subnets, err := network.GetSubnets()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing subnets of VLAN %d: %s\n", vlanID, err)
		return
	}

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

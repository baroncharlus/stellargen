package main

import (
	"github.com/baroncharlus/stellargen/seed"
	"github.com/baroncharlus/stellargen/trust"
	"log"
)

func main() {
	// generate a keypair to recieve funds.
	rKey, rAddr := seed.Seed()
	// generate a keypair to send funds
	sKey, sAddr := seed.Seed()

	log.Println("Recieving key: " + rKey)
	log.Println("Recieving address: " + rAddr)
	log.Println("Sending key: " + sKey)
	log.Println("Sending Address: " + sAddr)

	// Fund our addresses by phoning friendbot. We really only need to fund
	// one, but ...
	seed.Fund([]string{sAddr, rAddr}...)

	// In this example we will trust our newly minted recieving address to
	// trust our newly minted sending address.
	trust.Trust(rKey, sAddr, "USD")
}

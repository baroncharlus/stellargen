package seed

import (
	"fmt"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"io/ioutil"
	"log"
	"net/http"
)

// Generate a keypair.
func Seed() (seed, addr string) {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	seed = pair.Seed()
	addr = pair.Address()

	return
}

// Fund our addresses with funds from friendbot.
func Fund(addrs ...string) {
	for _, addr := range addrs {
		resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + addr)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))

		account, err := horizon.DefaultTestNetClient.LoadAccount(addr)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Balances for account:", addr)

		for _, balance := range account.Balances {
			log.Println(balance)
		}
	}
}

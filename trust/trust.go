package trust

import (
	"fmt"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func Trust(seed, issuer, assetType string) {

	// Set seed to trust issuer with assetType.
	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: seed},
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.TestNetwork,
		b.Trust(assetType, issuer),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)

	blob := txeB64

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", resp)
}

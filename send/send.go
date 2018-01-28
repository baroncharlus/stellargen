package send

import (
	"fmt"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func Send(from, to string) {
	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: "0.1"},
		),
	)

	txe, err := tx.Sign(from)
	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)

	blob := txeB64

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}

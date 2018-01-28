package send

import (
	"fmt"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func main() {
	// GBMANC7RRV5NHP3Q5BMPK56Q7OE5ULO3X7DZ5E5O5MERUNCSLEAPKNB5
	from := SCBOAOKX4V75F3KEUKJIZIBT3GPRXAJGDWLUAEOIVP7P2NKO24NDCXT6
	// SCQNAKR6IQJOR2FIIN3NHNZC3AR56TKGKVOBZ2MU5L5PCQUAQWFJOMBJ
	to := GCALT5MCIN6GIRFUANWTO6D76BDAZHRFLRF7R45HWEM2GKK76JMHR4KO

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

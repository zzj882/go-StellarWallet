package publicdefine

import (
	"fmt"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/xdr"
)

type StellarAccountCreateInfo struct {
	SrcInfo    *StellarAccInfoDef
	Amount     float64
	Destinaton string
	signBase64 string
	ResultHash string
}

func (this *StellarAccountCreateInfo) GetSigned(seed string) string {

	_, spriv, _ := stellarbase.GenerateKeyFromSeed(seed)

	tx := build.TransactionBuilder{}

	ca := build.CreateAccountBuilder{}
	ca.Mutate(build.Destination{this.Destinaton})
	ca.Mutate(build.SourceAccount{this.SrcInfo.ID})
	ca.Mutate(build.NativeAmount{fmt.Sprintf("%f", this.Amount)})

	tx.Mutate(build.Sequence{xdr.SequenceNumber(this.SrcInfo.NextSequence())})
	if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
		tx.Mutate(build.TestNetwork)
	} else {
		tx.Mutate(build.PublicNetwork)
	}
	tx.Mutate(ca)
	tx.Mutate(build.SourceAccount{this.SrcInfo.ID})
	tx.TX.Fee = BASEMENT_FEE
	result := tx.Sign(&spriv)

	var err error

	this.signBase64, err = result.Base64()
	// fmt.Printf("tx base64: %s\r\n", this.signBase64)

	if err == nil {
		return this.signBase64
	}

	fmt.Println(err)
	return ""
}

func (this *StellarAccountCreateInfo) PutResult(ret map[string]interface{}) {
	hash, ok := ret["hash"]
	this.ResultHash = ""
	if ok {
		this.ResultHash = hash.(string)
		return
	}
	fmt.Println(ret)
}

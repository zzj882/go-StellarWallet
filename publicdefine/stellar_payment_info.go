package publicdefine

import (
	"fmt"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/xdr"
)

const (
	BASEMENT_FEE = 100
)

type StellarPaymentInfo struct {
	SrcInfo    *StellarAccInfoDef
	Amount     float64
	Destinaton string
	signBase64 string
	ResultHash string
}

func (this *StellarPaymentInfo) GetSigned(seed string) string {

	_, spriv, _ := stellarbase.GenerateKeyFromSeed(seed)

	tx := build.TransactionBuilder{}
	pb := build.PaymentBuilder{}
	des := build.Destination{this.Destinaton}
	na := build.NativeAmount{fmt.Sprintf("%f", this.Amount)}

	pb.Mutate(des)
	pb.Mutate(na)

	tx.Mutate(build.Sequence{xdr.SequenceNumber(this.SrcInfo.NextSequence())})
	tx.Mutate(build.DefaultNetwork)
	tx.Mutate(pb)
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

func (this *StellarPaymentInfo) PutResult(ret map[string]interface{}) {
	hash, ok := ret["hash"]
	this.ResultHash = ""
	if ok {
		this.ResultHash = hash.(string)
		return
	}
	fmt.Println(ret)
}

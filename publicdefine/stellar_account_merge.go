package publicdefine

import (
	"fmt"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/xdr"
)

type StellarAccountMerge struct {
	SrcInfo        *StellarAccInfoDef
	DestPublicAddr string
	signBase64     string
	ResultHash     string
}

func (this *StellarAccountMerge) GetSigned(seed string) string {

	tx := build.TransactionBuilder{}
	tx.TX = &xdr.Transaction{}
	opt := xdr.Operation{}
	srcAccID, _ := stellarbase.AddressToAccountId(this.SrcInfo.ID)
	destAccID, _ := stellarbase.AddressToAccountId(this.DestPublicAddr)

	opt.SourceAccount = &srcAccID
	opt.Body, _ = xdr.NewOperationBody(xdr.OperationTypeAccountMerge,
		destAccID)
	tx.TX.Operations = append(tx.TX.Operations, opt)

	tx.Mutate(build.Sequence{xdr.SequenceNumber(this.SrcInfo.NextSequence())})
	if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
		tx.Mutate(build.TestNetwork)
	} else {
		tx.Mutate(build.PublicNetwork)
	}
	tx.Mutate(build.SourceAccount{this.SrcInfo.ID})
	tx.TX.Fee = BASEMENT_FEE
	result := tx.Sign(seed)

	var err error

	this.signBase64, err = result.Base64()
	// fmt.Printf("tx base64: %s\r\n", this.signBase64)

	if err == nil {
		return this.signBase64
	}

	fmt.Println(err)
	return ""
}

func (this *StellarAccountMerge) PutResult(ret map[string]interface{}) {
	hash, ok := ret["hash"]
	this.ResultHash = ""
	if ok {
		this.ResultHash = hash.(string)
		return
	}
	fmt.Println(ret)
}

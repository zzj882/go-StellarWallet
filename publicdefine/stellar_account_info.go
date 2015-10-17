package publicdefine

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	STELLAR_ONE_UNIT float64 = 10000000.0
)

type StellarAccInfoDef struct {
	ID         string
	Balance    string
	Sequence   string
	Seq_Acc    string
	Seq_Number string
	Type       string
	Title      string
	Status     string
}

func (this *StellarAccInfoDef) ToString() (ret string) {
	ret = fmt.Sprintf("  Public ID\t: %s\n", this.ID)
	if this.IsExist() {
		ret += fmt.Sprintf("  Balance\t: %s\n", this.Balance)
		ret += fmt.Sprintf("  Sequence\t: %s\n", this.Sequence)
		ret += fmt.Sprintf("  Seq_Number\t: %s\n", this.Seq_Number)
	} else {
		ret += fmt.Sprintf("  Type\t\t: %s\n", this.Type)
		ret += fmt.Sprintf("  Title\t\t: %s\n", this.Title)
	}
	return
}

func (this *StellarAccInfoDef) IsExist() bool {
	return this.Status != "404"
}

func (this *StellarAccInfoDef) PutMapBody(idAddr string, mbody map[string]interface{}) {
	id, okid := mbody["id"]
	sequence, oksequence := mbody["sequence"]
	balances, okbalances := mbody["balances"]

	stype, oktype := mbody["type"]
	title, oktitle := mbody["title"]
	status, okstauts := mbody["status"]
	if oktitle && oktype && okstauts {
		this.ID = idAddr
		this.Title = title.(string)
		this.Type = stype.(string)
		this.Status = fmt.Sprintf("%d", int(status.(float64)))
	} else if okid && oksequence && okbalances {
		this.ID = fmt.Sprint(id)
		this.Sequence = fmt.Sprintf("%15.0f", sequence)
		this.Sequence = strings.Trim(this.Sequence, " ")

		tmp, _ := strconv.ParseUint(this.Sequence, 10, 64)
		tmp_low := tmp & 0xFFFFFFFF
		tmp_high := (tmp >> 32) & 0xFFFFFFFF
		this.Seq_Acc = fmt.Sprintf("%08x", tmp_high)
		this.Seq_Number = fmt.Sprintf("%08x", tmp_low)

		balance, _ := balances.([]interface{})[0].(map[string]interface{})["balance"]
		this.Balance = balance.(string)
	}
}

func (this *StellarAccInfoDef) NextSequence() uint64 {
	seqHigh, _ := strconv.ParseUint(this.Seq_Acc, 16, 64)
	seqLow, _ := strconv.ParseUint(this.Seq_Number, 16, 64)
	seqLow += 1
	var ret uint64
	ret = (seqHigh << 32) | (seqLow & 0xFFFFFFFF)
	// fmt.Println("NextSequence : ret = ", ret)
	// fmt.Println("NextSequence : seqHigh = ", seqHigh)
	// fmt.Println("NextSequence : seqLow = ", seqLow)
	return ret
}

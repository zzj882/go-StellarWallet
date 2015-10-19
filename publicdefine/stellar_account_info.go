package publicdefine

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	STELLAR_ONE_UNIT float64 = 10000000.0
)

type StellarAsset struct {
	AssetCode string
	Balance   string
}

type StellarAccInfoDef struct {
	ID         string
	Balance    string
	Sequence   string
	Seq_Acc    string
	Seq_Number string
	Type       string
	Title      string
	Status     string
	AssetMap   map[string][]StellarAsset
}

func (this *StellarAccInfoDef) ToString() (ret string) {
	ret = fmt.Sprintf("  Public ID\t: %s\n", this.ID)
	if this.IsExist() {
		ret += fmt.Sprintf("  Balance\t: %s\n", this.Balance)
		ret += fmt.Sprintf("  Sequence\t: %s\n", this.Sequence)
		ret += fmt.Sprintf("  Seq_Number\t: %s\n", this.Seq_Number)
		if this.AssetMap != nil {
			ret += fmt.Sprintf("   **** Trust List ****\r\n")
			for key, val := range this.AssetMap {
				ret += fmt.Sprintf("    Issuer is [ %s ] \r\n", key)
				for i := 0; i < len(val); i++ {
					ret += fmt.Sprintf("    \tAssetCode -> %s \t   Balance -> %s\r\n", val[i].AssetCode, val[i].Balance)
				}
			}
		}
	} else {
		ret += fmt.Sprintf("  Type\t\t: %s\n", this.Type)
		ret += fmt.Sprintf("  Title\t\t: %s\n", this.Title)
	}
	return
}

func (this *StellarAccInfoDef) IsExist() bool {
	return this.Status != "404"
}

/*
	[ httpsget().body ]
		 {
	  "_links": {
	    "effects": {
	      "href": "/accounts/GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN/effects{?cursor,limit,order}",
	      "templated": true
	    },
	    "offers": {
	      "href": "/accounts/GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN/offers{?cursor,limit,order}",
	      "templated": true
	    },
	    "operations": {
	      "href": "/accounts/GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN/operations{?cursor,limit,order}",
	      "templated": true
	    },
	    "self": {
	      "href": "/accounts/GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN"
	    },
	    "transactions": {
	      "href": "/accounts/GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN/transactions{?cursor,limit,order}",
	      "templated": true
	    }
	  },
	  "id": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
	  "paging_token": "412767831986177",
	  "address": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
	  "sequence": 412767831982081,
	  "subentry_count": 0,
	  "inflation_destination": null,
	  "home_domain": "",
	  "thresholds": {
	    "low_threshold": 0,
	    "med_threshold": 0,
	    "high_threshold": 0
	  },
	  "flags": {
	    "auth_required": false,
	    "auth_revocable": false
	  },
	  "balances": [
	    {
	      "asset_type": "native",
	      "balance": "70099.1999200"
	    }
	  ],
	  "signers": [
	    {
	      "address": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
	      "weight": 1
	    }
	  ]
	}

=====================================================================================================
	[ httpsget().body ]
		 {
	  "_links": {
	    "effects": {
	      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/effects{?cursor,limit,order}",
	      "templated": true
	    },
	    "offers": {
	      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/offers{?cursor,limit,order}",
	      "templated": true
	    },
	    "operations": {
	      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations{?cursor,limit,order}",
	      "templated": true
	    },
	    "self": {
	      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X"
	    },
	    "transactions": {
	      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/transactions{?cursor,limit,order}",
	      "templated": true
	    }
	  },
	  "id": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
	  "paging_token": "413278933094401",
	  "address": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
	  "sequence": 413278933090320,
	  "subentry_count": 2,
	  "inflation_destination": null,
	  "home_domain": "",
	  "thresholds": {
	    "low_threshold": 0,
	    "med_threshold": 0,
	    "high_threshold": 0
	  },
	  "flags": {
	    "auth_required": false,
	    "auth_revocable": false
	  },
	  "balances": [
	    {
	      "asset_type": "credit_alphanum4",
	      "balance": "0.0000000",
	      "asset_code": "USD",
	      "issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
	      "limit": "922337203685.4775807"
	    },
	    {
	      "asset_type": "credit_alphanum4",
	      "balance": "0.0000000",
	      "asset_code": "XLM",
	      "issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
	      "limit": "922337203685.4775807"
	    },
	    {
	      "asset_type": "native",
	      "balance": "9900.7994600"
	    }
	  ],
	  "signers": [
	    {
	      "address": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
	      "weight": 1
	    }
	  ]
	}
*/

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

		for _, itm := range balances.([]interface{}) {
			asset_type, _ := itm.(map[string]interface{})["asset_type"]
			balance, _ := itm.(map[string]interface{})["balance"]
			if "native" == asset_type.(string) {
				this.Balance = balance.(string)
			} else if "credit_alphanum4" == asset_type.(string) {
				if this.AssetMap == nil {
					this.AssetMap = make(map[string][]StellarAsset)
				}
				issuer, _ := itm.(map[string]interface{})["issuer"]
				asset_code, _ := itm.(map[string]interface{})["asset_code"]

				sas, ok := this.AssetMap[issuer.(string)]
				if !ok {
					sas = make([]StellarAsset, 0)
				}
				sas = append(sas, StellarAsset{asset_code.(string), balance.(string)})
				this.AssetMap[issuer.(string)] = sas
			}
		}
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

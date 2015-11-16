package publicdefine

import (
	"fmt"
	"regexp"
	"strings"
)

type SubAccOperQuaryRecordInterface interface {
	ToString() string
	GetType() string
	DecodeBody(b map[string]interface{})
}

type AssetInfoBaseItem struct {
	Asset_Code   string
	Asset_Type   string
	Asset_Issuer string
}

type SubAccOperRecordItemBase struct {
	OpType          string
	SourceAccount   string
	TransactionHash string
}

type SubCreateAccountItem struct {
	SubAccOperRecordItemBase
	Account         string
	Funder          string
	StartingBalance string
}

type SubPaymentItem struct {
	SubAccOperRecordItemBase
	Amount      string
	FromAccount string
	ToAccount   string
	AssetInfo   *AssetInfoBaseItem
}

type SubChangeTrustItem struct {
	SubAccOperRecordItemBase
	AssetCode   string
	AssetIssuer string
	AssetType   string
	Trustee     string
	Trustor     string
}

type SubAccountMergeItem struct {
	SubAccOperRecordItemBase
	MergeSource string
	MergeInto   string
}

type SubSetOptionsItem struct {
	SubAccOperRecordItemBase
	HomeDomain    string
	InflationDest string
	SignerKey     string
	SignerWeight  string
	SetFlags      []string
	ClearFlags    []string
}

type SubManageOfferItem struct {
	SubAccOperRecordItemBase
	Amount  string
	OfferID string
	Price   string
	Buying  *AssetInfoBaseItem
	Selling *AssetInfoBaseItem
}

type StellarAccOperationQuary struct {
	QuaryCursor string
	IsEnd       bool
	Records     []SubAccOperQuaryRecordInterface
}

func (this *SubAccOperRecordItemBase) ToString() string {
	return fmt.Sprintf("          Type = [%s]\r\n SourceAccount = [%s]\r\n          Hash = [%s]\r\n",
		this.OpType, this.SourceAccount, this.TransactionHash)
}

func (this *SubAccOperRecordItemBase) GetType() string {
	return this.OpType
}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/413278933094401/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=413278933094401\u0026order=asc"
       },
       "self": {
         "href": "/operations/413278933094401"
       },
       "succeeds": {
         "href": "/operations?cursor=413278933094401\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/dcef180a209b3dab35791a56b175c18a3a9ee1c57062f74a4a885a1b7a8b8067"
       }
     },
     "account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "funder": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
     "id": 413278933094401,
     "paging_token": "413278933094401",
     "source_account": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
     "starting_balance": "10000.0",
     "type": "create_account",
     "type_i": 0
   },
*/
func (this *SubAccOperRecordItemBase) DecodeBody(b map[string]interface{}) {
	_links, _linksok := b["_links"]
	source_account, source_account_ok := b["source_account"]
	if _linksok && source_account_ok {
		this.SourceAccount = source_account.(string)

		transaction, _ := _links.(map[string]interface{})["transaction"]
		href, _ := transaction.(map[string]interface{})["href"]
		hrefurl := href.(string)
		this.TransactionHash = strings.Trim(hrefurl, "/transactions/")
	}
}

func (this *SubCreateAccountItem) ToString() (ret string) {
	ret = this.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("        Funder = [%s]\r\n       Account = [%s]\r\n       Balance = [%s]\r\n",
		this.Funder, this.Account, this.StartingBalance)
	return
}

func (this *SubCreateAccountItem) DecodeBody(b map[string]interface{}) {
	this.SubAccOperRecordItemBase.DecodeBody(b)
	account, accountok := b["account"]
	funder, funderok := b["funder"]
	starting_balance, starting_balanceok := b["starting_balance"]
	if accountok && funderok && starting_balanceok {
		this.Account = account.(string)
		this.Funder = funder.(string)
		this.StartingBalance = starting_balance.(string)
	}
}

func (this *SubPaymentItem) ToString() (ret string) {
	ret = this.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("          From = [%s]\r\n            To = [%s]\r\n        Amount = [%s]\r\n",
		this.FromAccount, this.ToAccount, this.Amount)
	if this.AssetInfo != nil {
		ret += fmt.Sprintf("          Code = [%s]\r\n        Issuer = [%s]\r\n",
			this.AssetInfo.Asset_Code, this.AssetInfo.Asset_Issuer)
	}
	return
}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/477574593515521/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=477574593515521\u0026order=asc"
       },
       "self": {
         "href": "/operations/477574593515521"
       },
       "succeeds": {
         "href": "/operations?cursor=477574593515521\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/d3642254d90547a67a7f25827c61f79ca57010521615b5f391b5ac664aa42028"
       }
     },
     "amount": "10.0",
     "asset_type": "native",
     "from": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "id": 477574593515521,
     "paging_token": "477574593515521",
     "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "to": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "type": "payment",
     "type_i": 1
   },

*/
func (this *SubPaymentItem) DecodeBody(b map[string]interface{}) {
	this.SubAccOperRecordItemBase.DecodeBody(b)
	amount, amountok := b["amount"]
	from, fromok := b["from"]
	to, took := b["to"]
	if amountok && fromok && took {
		this.Amount = amount.(string)
		this.FromAccount = from.(string)
		this.ToAccount = to.(string)
	}
	asset_issuer, asset_issuerok := b["asset_issuer"]
	asset_type, asset_typeok := b["asset_type"]
	asset_code, asset_codeok := b["asset_code"]
	if asset_issuerok && asset_typeok && asset_codeok {
		if this.AssetInfo == nil {
			this.AssetInfo = &AssetInfoBaseItem{}
		}

		this.AssetInfo.Asset_Code = asset_code.(string)
		this.AssetInfo.Asset_Issuer = asset_issuer.(string)
		this.AssetInfo.Asset_Type = asset_type.(string)
	}
}

func (this *SubChangeTrustItem) ToString() (ret string) {
	ret = this.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("     AssetCode = [%s]\r\n       Trustee = [%s]\r\n       Trustor = [%s]\r\n",
		this.AssetCode, this.Trustee, this.Trustor)
	return
}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/777758447767553/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=777758447767553\u0026order=asc"
       },
       "self": {
         "href": "/operations/777758447767553"
       },
       "succeeds": {
         "href": "/operations?cursor=777758447767553\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/973bed257adf83d4ffe4b9693a2ce7ffb91cbe5afaf4734bc1b7ef8f782f498b"
       }
     },
     "asset_code": "XLM",
     "asset_issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "asset_type": "credit_alphanum4",
     "id": 777758447767553,
     "limit": "922337203685.4775807",
     "paging_token": "777758447767553",
     "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "trustee": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "trustor": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "type": "change_trust",
     "type_i": 6
   }

*/
func (this *SubChangeTrustItem) DecodeBody(b map[string]interface{}) {
	this.SubAccOperRecordItemBase.DecodeBody(b)
	asset_code, asset_codeok := b["asset_code"]
	asset_issuer, asset_issuerok := b["asset_issuer"]
	asset_type, asset_typeok := b["asset_type"]
	trustee, trusteeok := b["trustee"]
	trustor, trustorok := b["trustor"]
	if asset_codeok && asset_issuerok && asset_typeok && trusteeok && trustorok {
		this.AssetCode = asset_code.(string)
		this.AssetIssuer = asset_issuer.(string)
		this.AssetType = asset_type.(string)
		this.Trustee = trustee.(string)
		this.Trustor = trustor.(string)
	}
}

func (this *SubAccountMergeItem) ToString() (ret string) {
	ret = this.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("   MergeSource = [%s]\r\n     MergeInto = [%s]\r\n",
		this.MergeSource, this.MergeInto)
	return

}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/496962075889665/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=496962075889665\u0026order=asc"
       },
       "self": {
         "href": "/operations/496962075889665"
       },
       "succeeds": {
         "href": "/operations?cursor=496962075889665\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/c9819aa9d497279c69d49f5fa24942cea2312a0c46002148e51bb98b90d83a20"

       }
     },
     "account": "GBRFZNZB3RDJHBWEUDGFMZEE6OTTZXHOGEQLBZL22RXW7VOH2NHOS4X6",
     "id": 496962075889665,
     "into": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "paging_token": "496962075889665",
     "source_account": "GBRFZNZB3RDJHBWEUDGFMZEE6OTTZXHOGEQLBZL22RXW7VOH2NHOS4X6",
     "type": "account_merge",
     "type_i": 8
   },
*/
func (this *SubAccountMergeItem) DecodeBody(b map[string]interface{}) {
	this.SubAccOperRecordItemBase.DecodeBody(b)
	account, accountok := b["account"]
	into, intook := b["into"]
	if accountok && intook {
		this.MergeSource = account.(string)
		this.MergeInto = into.(string)
	}
}

func (this *SubSetOptionsItem) ToString() (ret string) {
	ret = this.SubAccOperRecordItemBase.ToString()
	if len(this.SignerKey) > 0 {
		ret += fmt.Sprintf("     SignerKey = [%s]\r\n", this.SignerKey)
	}
	if len(this.SignerWeight) > 0 {
		ret += fmt.Sprintf("  SignerWeight = [%s]\r\n", this.SignerWeight)
	}
	if len(this.HomeDomain) > 0 {
		ret += fmt.Sprintf("    HomeDomain = [%s]\r\n", this.HomeDomain)
	}
	if len(this.InflationDest) > 0 {
		ret += fmt.Sprintf(" InflationDest = [%s]\r\n", this.InflationDest)
	}
	if this.SetFlags != nil && len(this.SetFlags) > 0 {
		ret += fmt.Sprintf("      SetFlags = %s\r\n", this.SetFlags)
	}
	if this.ClearFlags != nil && len(this.ClearFlags) > 0 {
		ret += fmt.Sprintf("    ClearFlags = %s\r\n", this.ClearFlags)
	}
	return
}

/*
{
  "_links": {
    "effects": {
      "href": "/operations/3317677552570369/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=3317677552570369\u0026order=asc"
    },
    "self": {
      "href": "/operations/3317677552570369"
    },
    "succeeds": {
      "href": "/operations?cursor=3317677552570369\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/be63c2d5c010711b9946b0363b85a43d514edca9a691eec229fa2108359d2115"
    }
  },
  "home_domain": "www.ledgercn.com",
  "id": 3317677552570369,
  "inflation_dest": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
  "paging_token": "3317677552570369",
  "set_flags": [
    1,
    2
  ],
  "set_flags_s": [
    "auth_required_flag",
    "auth_revocable_flag"
  ],
  "signer_key": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
  "signer_weight": 255,
  "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
  "type": "set_options",
  "type_i": 5
},
*/
func (this *SubSetOptionsItem) DecodeBody(b map[string]interface{}) {
	this.SubAccOperRecordItemBase.DecodeBody(b)

	home_domain, ok := b["home_domain"]
	if ok {
		this.HomeDomain = home_domain.(string)
	}

	inflation_dest, ok := b["inflation_dest"]
	if ok {
		this.InflationDest = inflation_dest.(string)
	}

	signer_key, ok := b["signer_key"]
	if ok {
		this.SignerKey = signer_key.(string)
	}

	signer_weight, ok := b["signer_weight"]
	if ok {
		this.SignerWeight = fmt.Sprintf("%d", (int64)(signer_weight.(float64)))
	}

	set_flags_s, ok := b["set_flags_s"]
	if ok {
		for _, v := range set_flags_s.([]interface{}) {
			if this.SetFlags == nil {
				this.SetFlags = make([]string, 0)
			}
			this.SetFlags = append(this.SetFlags, v.(string))
		}
	}

	clear_flags_s, ok := b["clear_flags_s"]
	if ok {
		for _, v := range clear_flags_s.([]interface{}) {
			if this.ClearFlags == nil {
				this.ClearFlags = make([]string, 0)
			}
			this.ClearFlags = append(this.ClearFlags, v.(string))
		}
	}
}

func (this *SubManageOfferItem) ToString() (ret string) {
	ret = this.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("        Amount = [%s]\r\n       OfferID = [%s]\r\n         Price = [%s]\r\n",
		this.Amount, this.OfferID, this.Price)
	if this.Buying != nil {
		ret += fmt.Sprintf("        Buying = [\r\n\t\t  Code : %s\r\n\t\t  Type : %s\r\n\t\t  Issuer : %s\r\n\t\t ]\r\n ",
			this.Buying.Asset_Code, this.Buying.Asset_Type, this.Buying.Asset_Issuer)
	}
	if this.Selling != nil {
		ret += fmt.Sprintf("      Selling = [\r\n\t\t  Code : %s\r\n\t\t  Type : %s\r\n\t\t  Issuer : %s\r\n\t\t ]\r\n ",
			this.Selling.Asset_Code, this.Selling.Asset_Type, this.Selling.Asset_Issuer)
	}
	return
}

func (this *SubManageOfferItem) DecodeBody(b map[string]interface{}) {
	this.SubAccOperRecordItemBase.DecodeBody(b)
	amount, ok := b["amount"]
	if ok {
		this.Amount = amount.(string)
	}

	offer_id, ok := b["offer_id"]
	if ok {
		this.OfferID = fmt.Sprintf("%d", (int64)(offer_id.(float64)))
	}

	price, ok := b["price"]
	if ok {
		this.Price = price.(string)
	}

	buying_asset_type, ok_buy_type := b["buying_asset_type"]
	buying_asset_issuer, ok_buy_iusser := b["buying_asset_issuer"]
	buying_asset_code, ok_buy_code := b["buying_asset_code"]
	selling_asset_type, ok_sell_type := b["selling_asset_type"]
	selling_asset_issuer, ok_sell_iusser := b["selling_asset_issuer"]
	selling_asset_code, ok_sell_code := b["selling_asset_code"]

	if ok_buy_type {
		if this.Buying == nil {
			this.Buying = &AssetInfoBaseItem{}
		}
		this.Buying.Asset_Type = buying_asset_type.(string)
		if this.Buying.Asset_Type != "native" {
			if ok_buy_code {
				this.Buying.Asset_Code = buying_asset_code.(string)
			}

			if ok_buy_iusser {
				this.Buying.Asset_Issuer = buying_asset_issuer.(string)
			}
		} else {
			this.Buying.Asset_Code = "XLM"

			if ok_sell_iusser {
				this.Buying.Asset_Issuer = selling_asset_issuer.(string)
			}
		}
	}

	if ok_sell_type {
		if this.Selling == nil {
			this.Selling = &AssetInfoBaseItem{}
		}
		this.Selling.Asset_Type = selling_asset_type.(string)
		if this.Selling.Asset_Type != "native" {
			if ok_sell_code {
				this.Selling.Asset_Code = selling_asset_code.(string)
			}

			if ok_sell_iusser {
				this.Selling.Asset_Issuer = selling_asset_issuer.(string)
			}
		} else {
			this.Selling.Asset_Code = "XLM"

			if ok_buy_iusser {
				this.Selling.Asset_Issuer = buying_asset_issuer.(string)
			}
		}
	}
}

/*
{
  "_embedded": {
    "records": [
      {
        "_links": {
          "effects": {
            "href": "/operations/413278933094401/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=413278933094401\u0026order=asc"
          },
          "self": {
            "href": "/operations/413278933094401"
          },
          "succeeds": {
            "href": "/operations?cursor=413278933094401\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/dcef180a209b3dab35791a56b175c18a3a9ee1c57062f74a4a885a1b7a8b8067"
          }
        },
        "account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "funder": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
        "id": 413278933094401,
        "paging_token": "413278933094401",
        "source_account": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
        "starting_balance": "10000.0",
        "type": "create_account",
        "type_i": 0
      },
      {
        "_links": {
          "effects": {
            "href": "/operations/477063492407297/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=477063492407297\u0026order=asc"
          },
          "self": {
            "href": "/operations/477063492407297"
          },
          "succeeds": {
            "href": "/operations?cursor=477063492407297\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/104af896c4a9e1fb4d5825626ff5da35eb106e6bb7eb61d97d79c618b59f4ec5"
          }
        },
        "amount": "1000.0",
        "asset_type": "native",
        "from": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "id": 477063492407297,
        "paging_token": "477063492407297",
        "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "to": "GBRFZNZB3RDJHBWEUDGFMZEE6OTTZXHOGEQLBZL22RXW7VOH2NHOS4X6",
        "type": "payment",
        "type_i": 1
      },
      {
        "_links": {
          "effects": {
            "href": "/operations/777827167244289/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=777827167244289\u0026order=asc"
          },
          "self": {
            "href": "/operations/777827167244289"
          },
          "succeeds": {
            "href": "/operations?cursor=777827167244289\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/791e3575cddec3e07ed52ef46fa134b9d7acbd0563cfc9ecc908db66017082a6"
          }
        },
        "asset_code": "USD",
        "asset_issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "asset_type": "credit_alphanum4",
        "id": 777827167244289,
        "limit": "922337203685.4775807",
        "paging_token": "777827167244289",
        "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "trustee": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "trustor": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "type": "change_trust",
        "type_i": 6
      },
      {
        "_links": {
          "effects": {
            "href": "/operations/777758447767553/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=777758447767553\u0026order=asc"
          },
          "self": {
            "href": "/operations/777758447767553"
          },
          "succeeds": {
            "href": "/operations?cursor=777758447767553\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/973bed257adf83d4ffe4b9693a2ce7ffb91cbe5afaf4734bc1b7ef8f782f498b"
          }
        },
        "asset_code": "XLM",
        "asset_issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "asset_type": "credit_alphanum4",
        "id": 777758447767553,
        "limit": "922337203685.4775807",
        "paging_token": "777758447767553",
        "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "trustee": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "trustor": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "type": "change_trust",
        "type_i": 6
      }
    ]
  },

  "_links": {
    "next": {
      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations?order=asc\u0026limit=10\u0026cursor=487208205160449"
    },
    "prev": {
      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations?order=desc\u0026limit=10\u0026cursor=413278933094401"
    },
    "self": {
      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations?order=asc\u0026limit=10\u0026cursor="
    }
  }
}
*/
func (this *StellarAccOperationQuary) PutMapBody(mbody map[string]interface{}) {
	_embedded, _embeddedok := mbody["_embedded"]
	if _embeddedok {
		records, recordsok := _embedded.(map[string]interface{})["records"]
		if recordsok {
			recordsSlice := records.([]interface{})
			length := len(recordsSlice)

			this.Records = make([]SubAccOperQuaryRecordInterface, 0)

			for i := 0; i < length; i++ {
				subRecord := recordsSlice[i]
				this.decodeSubrecord(subRecord.(map[string]interface{}))
			}
		}
	}

	_links, _ := mbody["_links"]
	this.decodeCursor(_links.(map[string]interface{}))

	if len(this.Records) < 10 {
		this.IsEnd = true
	}
}

func (this *StellarAccOperationQuary) decodeSubrecord(itm map[string]interface{}) {
	stype, stypeok := itm["type"]
	if stypeok {
		switch stype {
		case "create_account":
			subitm := &SubCreateAccountItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		case "payment":
			subitm := &SubPaymentItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		case "change_trust":
			subitm := &SubChangeTrustItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		case "account_merge":
			subitm := &SubAccountMergeItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		case "set_options":
			subitm := &SubSetOptionsItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		case "manage_offer":
			subitm := &SubManageOfferItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		default:
			subitm := &SubAccOperRecordItemBase{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			this.Records = append(this.Records, subitm)
		}
	}
}

func (this *StellarAccOperationQuary) decodeCursor(b map[string]interface{}) {
	prev, prevok := b["next"]
	if prevok {
		href, _ := prev.(map[string]interface{})["href"]
		hrefurl := href.(string)
		reg := regexp.MustCompile(`cursor=[\d]*`)
		cStr := reg.FindString(hrefurl)
		if len(cStr) > len("cursor=") {
			this.QuaryCursor = strings.Trim(cStr, "cursor=")
		} else {
			this.QuaryCursor = ""
		}
	}
}

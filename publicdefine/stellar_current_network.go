package publicdefine

const (
	STELLAR_TEST_NETWORK         = "https://horizon-testnet.stellar.org/"
	STELLAR_LIVE_NETWORK         = "https://horizon.stellar.org/"
	STELLAR_NETWORK_ACCOUNTS     = "accounts"
	STELLAR_NETWORK_TRANSACTIONS = "transactions"
)

var STELLAR_DEFAULT_NETWORK string = STELLAR_TEST_NETWORK

func GetDefaultNWString() string {
	if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
		return "Test network"
	}
	return "Live network"
}

package publicdefine

import (
	"github.com/stellar/go-stellar-base/strkey"
)

func VerifyGAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteAccountID, addr)
	return err
}

func VerifySAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteSeed, addr)
	return err
}

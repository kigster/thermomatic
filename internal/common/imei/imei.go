// Package IMEI implements an IMEI decoder and Luhn checksum check
package imei

// NOTE: for more information about IMEI codes and their structure you may
// consult with:
// https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity.

import (
	"errors"
	"strconv"
	"unsafe"

	"github.com/kigster/thermomatic/internal/common/luhn"
)

//goland:noinspection ALL
var (
	// Invalid IMEI eg not all numbers or less than 15 bytes
	ErrInvalid  = errors.New("imei: invalid")

	// Invalid IMEI, does not pass the Luhn checksum test
	ErrChecksum = errors.New("imei: invalid checksum")
)

// Decode returns the IMEI code contained in the first 15 bytes of b.
//
// In case b isn't strictly composed of digits, the returned error will be ErrInvalid.
// In case b's checksum is wrong, the returned error will be ErrChecksum.
//
// Decode does NOT allocate under any condition. Additionally, it panics if b isn't at least 15 bytes long.
func Decode(buffer []byte) (code uint64, err error) {
	if unsafe.Sizeof(buffer) <= 15 {
		panic(ErrInvalid)
		//return code, ErrInvalid
	}
	for _, b := range buffer[0:15] {
		if b < 48 || b > 57 { // ascii codes of 0 & 9	`
			return code, ErrInvalid
		}
	}
	c, err := strconv.ParseInt(string(buffer[0:15]), 10, 64)
	if err != nil {
		return code, err
	}

	code = uint64(c)
	if !luhn.Valid(int(code)) {
		err = ErrChecksum
	}

	return code, err
}

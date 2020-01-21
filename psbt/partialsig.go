package psbt

import (
	"bytes"

	"github.com/DiviProject/divid/diviec"
)

// PartialSig encapsulate a (DIVI public key, ECDSA signature)
// pair, note that the fields are stored as byte slices, not
// diviec.PublicKey or diviec.Signature (because manipulations will
// be with the former not the latter, here); compliance with consensus
// serialization is enforced with .checkValid()
type PartialSig struct {
	PubKey    []byte
	Signature []byte
}

// PartialSigSorter implements sort.Interface for PartialSig.
type PartialSigSorter []*PartialSig

func (s PartialSigSorter) Len() int { return len(s) }

func (s PartialSigSorter) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s PartialSigSorter) Less(i, j int) bool {
	return bytes.Compare(s[i].PubKey, s[j].PubKey) < 0
}

// validatePubkey checks if pubKey is *any* valid pubKey serialization in a
// Bitcoin context (compressed/uncomp. OK).
func validatePubkey(pubKey []byte) bool {
	_, err := diviec.ParsePubKey(pubKey, diviec.S256())
	if err != nil {
		return false
	}
	return true
}

// validateSignature checks that the passed byte slice is a valid DER-encoded
// ECDSA signature, including the sighash flag.  It does *not* of course
// validate the signature against any message or public key.
func validateSignature(sig []byte) bool {
	_, err := diviec.ParseDERSignature(sig, diviec.S256())
	if err != nil {
		return false
	}
	return true
}

// checkValid checks that both the pbukey and sig are valid. See the methods
// (PartialSig, validatePubkey, validateSignature) for more details.
//
// TODO(waxwing): update for Schnorr will be needed here if/when that
// activates.
func (ps *PartialSig) checkValid() bool {
	return validatePubkey(ps.PubKey) && validateSignature(ps.Signature)
}

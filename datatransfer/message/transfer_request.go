package message

import (
	"io"

	"github.com/filecoin-project/go-fil-components/datatransfer"
	"github.com/ipfs/go-cid"
)

//go:generate cbor-gen-for transferRequest

// transferRequest is a struct that fulfills the DataTransferRequest interface.
// its members are exported to be used by cbor-gen
type transferRequest struct {
	BCid   string
	Canc   bool
	PID    []byte
	Part   bool
	Pull   bool
	Stor   []byte
	Vouch  []byte
	VTyp   string
	XferID uint64
}

// IsRequest always returns true in this case because this is a transfer request
func (trq *transferRequest) IsRequest() bool {
	return true
}

func (trq *transferRequest) TransferID() datatransfer.TransferID {
	return datatransfer.TransferID(trq.XferID)
}

// ========= DataTransferRequest interface
// IsPull returns true if this is a data pull request
func (trq *transferRequest) IsPull() bool {
	return trq.Pull
}

// VoucherType returns the Voucher ID
func (trq *transferRequest) VoucherType() string {
	return trq.VTyp
}

// Voucher returns the Voucher bytes
func (trq *transferRequest) Voucher() []byte {
	return trq.Vouch
}

// BaseCid returns the Base CID
func (trq *transferRequest) BaseCid() cid.Cid {
	res, err := cid.Decode(trq.BCid)
	if err != nil {
		return cid.Undef
	}
	return res
}

// Selector returns the message Selector bytes
func (trq *transferRequest) Selector() []byte {
	return trq.Stor
}

// IsCancel returns true if this is a cancel request
func (trq *transferRequest) IsCancel() bool {
	return trq.Canc
}

// IsPartial returns true if this is a partial request
func (trq *transferRequest) IsPartial() bool {
	return trq.Part
}

// Cancel cancels a transfer request
func (trq *transferRequest) Cancel() error {
	// do other stuff ?
	trq.Canc = true
	return nil
}

// ToNet serializes a transfer request. It's a wrapper for MarshalCBOR to provide
// symmetry with FromNet
func (trq *transferRequest) ToNet(w io.Writer) error {
	msg := transferMessage{
		IsRq:     true,
		Request:  trq,
		Response: nil,
	}
	return msg.MarshalCBOR(w)
}
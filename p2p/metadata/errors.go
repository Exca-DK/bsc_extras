package metadata

import "errors"

type PacketError error

var (
	ErrPacketNotFound   PacketError = errors.New("packet not found")
	ErrPackedRegistered PacketError = errors.New("packet already registered")
)

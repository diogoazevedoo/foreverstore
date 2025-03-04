package p2p

import "net"

// RPC holds any arbitrary data that is being sent over
// each transport between two nodes inthe network.
type RPC struct {
	From    net.Addr
	Payload []byte
}

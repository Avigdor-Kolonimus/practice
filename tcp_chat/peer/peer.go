package peer

import (
	"net"
	"time"
)

type Peer struct {
	Conn        net.Conn
	ConnectedAt time.Time
}
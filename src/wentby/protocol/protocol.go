package protocol

import "net"

// PacketReader is used to unmarshal a complete packet from buff
type ReaderInter interface {
	// Read data from conn and build a complete packet.
	ReadPacket(conn net.Conn, buff []byte) (interface{}, []byte, error)
}

// PacketWriter is used to marshal packet into buff
type WriterInter interface {
	WritePacket(conn net.Conn, buff []byte) error
}

// PacketProtocol just a composite interface
type ProtocolInter interface {
	ReaderInter
	WriterInter
}

type ProtocolImpl struct {
}

func (pi *ProtocolImpl) ReadPacket(conn net.Conn, buff []byte) (interface{}, []byte, error) {
	if cap(buff) < 12 {

	}
}

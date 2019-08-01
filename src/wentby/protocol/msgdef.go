package protocol

type MsgHead struct {
	id  uint16
	len uint16
}
type MsgBody struct {
	data []byte
}
type MsgPacket struct {
	head MsgHead
	body MsgBody
}

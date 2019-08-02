package protocol

/*
-----------------------------------------------
               msgpacket
-----------------------------------------------
      msghead     |  msgbody
-----------------------------------------------
id      |   len   |   data
-----------------------------------------------
*/
type MsgHead struct {
	Id  uint16
	Len uint16
}
type MsgBody struct {
	Data []byte
}
type MsgPacket struct {
	Head MsgHead
	Body MsgBody
}

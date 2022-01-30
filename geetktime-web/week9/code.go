package week9

import (
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
)

type receive_message struct {
	verison uint16
	op uint32
	seq uint32
	body []byte
}

/*
goim 协议结构
4bytes PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
2bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
2bytes Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
4bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
4bytes Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码。
*/
func (rm * receive_message)decoder(data []byte) error {
	if len(data) <= 16 {
		return errors.New("illegal package length")
	}
	packetLen := binary.BigEndian.Uint32(data[: 4])
	fmt.Printf("packetLen: %v\n", packetLen)

	headerLen := uint32(binary.BigEndian.Uint16(data[4: 6]))
	fmt.Printf("headerLen: %v\n", headerLen)

	rm.verison = binary.BigEndian.Uint16(data[6: 8])
	fmt.Printf("version: %v\n", rm.verison)

	rm.op = binary.BigEndian.Uint32(data[8: 12])
	fmt.Printf("operation: %v\n", rm.op)

	rm.seq = binary.BigEndian.Uint32(data[12: 16])
	fmt.Printf("sequence: %v\n", rm.seq)

	rm.body = data[16 : packetLen-headerLen]
	fmt.Printf("body: %v\n", rm.body)

	return nil
}

// Encode
func (rm *receive_message) encoder() ([]byte, error) {
	packLen := uint32(16 + len(rm.body))
	pos := 0
	raw := make([]byte, packLen)

	binary.BigEndian.PutUint32(raw[pos:], packLen)
	pos += 4

	binary.BigEndian.PutUint16(raw[pos:], 16)
	pos += 2

	binary.BigEndian.PutUint16(raw[pos:], rm.verison)
	pos += 2

	binary.BigEndian.PutUint32(raw[pos:], rm.op)
	pos += 4

	binary.BigEndian.PutUint32(raw[pos:], rm.seq)
	pos += 4

	copy(raw[pos:], rm.body)
	return raw, nil
}

func main() {
	rm := receive_message{}
	rm.decoder([]byte("Hello world"))
	rm.encoder()
}
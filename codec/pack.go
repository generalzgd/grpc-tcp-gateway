/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: pack.go
 * @time: 2019/9/23 21:20
 */
package codec

import (
	`bytes`
	`encoding/binary`
	`errors`
	`fmt`
)

const (
	PACK_CODEC_PROTO = 0
	PACK_CODEC_JSON  = 1

	PACK_HEAD_SIZE = 8
	PACK_MAX_SIZE  = 1024 * 32
)

const (
	// 心跳id
	ID_Heartbeat = 1 + iota
	// todo maintain other logical id
	// ...
)

var (
	TooLargeError = errors.New("package too large")
	NilError = errors.New("nil pointer")
	UnserializeError = errors.New("unserialize package failure")
	IdFieldError = errors.New("id field value error")
	EndpointError = errors.New("endpoint error")
	// TransmitError = errors.New("transmit error")
)

// /
func EncodePacket(pack *GateClientPack) []byte {
	if pack == nil {
		return nil
	}

	buf := make([]byte, PACK_HEAD_SIZE+pack.Length)
	if _, err := pack.SerializeWithBuf(buf); err != nil {
		return nil
	}
	return buf
}

func DecodePacket(data []byte, pack *GateClientPack) error {
	if pack == nil {
		return NilError
	}
	if err := pack.Unserialize(data); err != nil {
		return UnserializeError
	}

	if int(pack.Length) > PACK_MAX_SIZE {
		return TooLargeError
	}

	if pack.Id < 1 {
		return IdFieldError
	}
	return nil
}
// /

type GateClientPackHead struct {
	Length uint16 // body的长度，65535/1024 ~ 63k
	Seq    uint16 // 序列号
	Id     uint16 // 协议id，可以映射到对应的service:method（兼容字段，后期考虑，把房间聊天网关迁移过来）
	Codec  uint16 // 0:proto  1:json
}

// 网关包, 小端
type GateClientPack struct {
	GateClientPackHead
	Body []byte // protobuf or json
}

func (p *GateClientPack) Serialize() []byte {
	out, _ := p.SerializeWithBuf(nil)
	return out
}

func (p *GateClientPack) SerializeWithBuf(out []byte) ([]byte, error) {
	// val := int(unsafe.Sizeof(p.GateClientPackHead))
	if cap(out) < 1 {
		out = make([]byte, 0, PACK_HEAD_SIZE+p.Length)
	}

	buf := bytes.NewBuffer(out)
	if err := binary.Write(buf, binary.LittleEndian, p.GateClientPackHead); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, p.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *GateClientPack) Unserialize(in []byte) error {
	r := bytes.NewReader(in)
	if err := binary.Read(r, binary.LittleEndian, &p.GateClientPackHead); err != nil {
		return err
	}
	p.Body = make([]byte, p.Length)
	copy(p.Body, in[PACK_HEAD_SIZE:])
	return nil
}

func (p *GateClientPack) String() string {
	return fmt.Sprintf("head:%v,Body:%s", p.GateClientPackHead, string(p.Body))
}

package uqholder

import (
	"bytes"
	"neo-omega-kernel/minecraft"
	"neo-omega-kernel/minecraft/protocol/packet"
	"neo-omega-kernel/neomega"
	"neo-omega-kernel/neomega/encoding/binary_read_write"
	binary_read_write2 "neo-omega-kernel/neomega/encoding/binary_read_write"
	"neo-omega-kernel/neomega/encoding/little_endian"
)

func init() {
	if false {
		func(neomega.BotBasicInfoHolder) {}(&BotBasicInfoHolder{})
	}
}

type BotBasicInfoHolder struct {
	BotName      string
	BotRuntimeID uint64
	BotUniqueID  int64
	BotIdentity  string
}

func (b *BotBasicInfoHolder) Marshal() (data []byte, err error) {
	basicWriter := bytes.NewBuffer(nil)
	writer := binary_read_write2.WrapBinaryWriter(basicWriter)
	err = little_endian.WriteString(writer, b.BotName)
	if err != nil {
		return nil, err
	}
	err = little_endian.WriteUint64(writer, b.BotRuntimeID)
	if err != nil {
		return nil, err
	}
	err = little_endian.WriteInt64(writer, b.BotUniqueID)
	if err != nil {
		return nil, err
	}
	err = little_endian.WriteString(writer, b.BotIdentity)
	if err != nil {
		return nil, err
	}
	return basicWriter.Bytes(), err
}

func (b *BotBasicInfoHolder) Unmarshal(data []byte) (err error) {
	basicReader := bytes.NewReader(data)
	reader := binary_read_write.WrapBinaryReader(basicReader)
	b.BotName, err = little_endian.String(reader)
	if err != nil {
		return err
	}
	b.BotRuntimeID, err = little_endian.Uint64(reader)
	if err != nil {
		return err
	}
	b.BotUniqueID, err = little_endian.Int64(reader)
	if err != nil {
		return err
	}
	b.BotIdentity, err = little_endian.String(reader)
	if err != nil {
		return err
	}
	return nil
}

func (b *BotBasicInfoHolder) UpdateFromPacket(packet packet.Packet) {
}

func (b *BotBasicInfoHolder) GetBotName() string {
	return b.BotName
}

func (b *BotBasicInfoHolder) GetBotRuntimeID() uint64 {
	return b.BotRuntimeID
}

func (b *BotBasicInfoHolder) GetBotUniqueID() int64 {
	return b.BotUniqueID
}

func (b *BotBasicInfoHolder) GetBotIdentity() string {
	return b.BotIdentity
}

func (b *BotBasicInfoHolder) GetBotUUIDStr() string {
	return b.BotIdentity
}

func NewBotInfoHolder(conn *minecraft.Conn) neomega.BotBasicInfoHolder {
	h := &BotBasicInfoHolder{}
	gd := conn.GameData()
	h.BotRuntimeID = gd.EntityRuntimeID
	h.BotUniqueID = gd.EntityUniqueID
	h.BotName = conn.IdentityData().DisplayName
	h.BotIdentity = conn.IdentityData().Identity
	if DEBUG {
		println("BotRuntimeID:", h.BotRuntimeID)
		println("BotUniqueID:", h.BotUniqueID)
		println("BotName:", h.BotName)
		println("BotIdentity:", h.BotIdentity)
	}
	return h
}

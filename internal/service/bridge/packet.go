package bridge

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Packet struct {
	capacity uint
	cursize  uint
	buffer   []byte
	pointer  uint
}

// i need to make a global typedef for capacity
func CreatePacket(capacity uint) *Packet {
	return &Packet{
		capacity: capacity,
		cursize:  0,
		buffer:   make([]byte, capacity),
		pointer:  0,
	}
}

// might need to check for overflow errors
func ConvertToPacket(packet []byte, capacity uint) *Packet {
	return &Packet{
		capacity: capacity,
		cursize:  uint(len(packet)),
		buffer:   packet,
		pointer:  0,
	}
}

func (p *Packet) GetBuffer() []byte {
	return p.buffer
}

func (p *Packet) ReadInt32() (int32, error) {
	var offset uint = 4
	if p.pointer+offset >= p.cursize {
		return 0, errors.New("you are out of bounds of the packet")
	}
	byte_data := p.buffer[p.pointer : p.pointer+offset]
	var data int32
	err := binary.Read(bytes.NewReader(byte_data), binary.BigEndian, &data)
	if err != nil {
		return 0, err
	}
	p.pointer += offset
	return data, nil
}
func (p *Packet) ReadInt64() (int64, error) {
	var offset uint = 8
	if p.pointer+offset >= p.cursize {
		return 0, errors.New("you are out of bounds of the packet")
	}
	byte_data := p.buffer[p.pointer : p.pointer+offset]
	var data int64
	err := binary.Read(bytes.NewReader(byte_data), binary.BigEndian, &data)
	if err != nil {
		return 0, err
	}
	p.pointer += offset
	return data, nil
}

func (p *Packet) ReadFloat32() (float32, error) {
	var offset uint = 4
	if p.pointer+offset >= p.cursize {
		return 0, errors.New("you are out of bounds of the packet")
	}
	byte_data := p.buffer[p.pointer : p.pointer+offset]
	var data float32
	err := binary.Read(bytes.NewReader(byte_data), binary.BigEndian, &data)
	if err != nil {
		return 0, err
	}
	p.pointer += offset
	return data, nil
}

func (p *Packet) ReadFloat64() (float64, error) {
	var offset uint = 8
	if p.pointer+offset >= p.cursize {
		return 0, errors.New("you are out of bounds of the packet")
	}
	byte_data := p.buffer[p.pointer : p.pointer+offset]
	var data float64
	err := binary.Read(bytes.NewReader(byte_data), binary.BigEndian, &data)
	if err != nil {
		return 0, err
	}
	p.pointer += offset
	return data, nil
}

func (p *Packet) ReadString() (string, error) {
	string_size, err := p.ReadInt32()
	if err != nil {
		return "", err
	}
	var offset uint = uint(string_size)
	if p.pointer+offset >= p.cursize {
		return "", errors.New("you are out of bounds of the packet")
	}
	byte_data := p.buffer[p.pointer : p.pointer+offset]
	if byte_data[len(byte_data)-1] != 0x00 {
		return "", errors.New("missing term byte, invalid string")
	}
	data := string(byte_data)
	p.pointer += offset
	return data, nil
}

// need to finish writing to packet LFMAO
func (p *Packet) WriteInt32(data int32) error {
	if !p.HasCapcity(4) {
		return errors.New("packet out of space")
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err == nil {
		p.buffer = append(p.buffer, buf.Bytes()...)
		p.cursize += 4
	}

	return err
}

// need to finish writing to packet LFMAO
func (p *Packet) WriteInt64(data int64) error {
	if !p.HasCapcity(8) {
		return errors.New("packet out of space")
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err == nil {
		p.buffer = append(p.buffer, buf.Bytes()...)
		p.cursize += 8
	}

	return err
}

func (p *Packet) WriteFloat32(data float32) error {
	if !p.HasCapcity(4) {
		return errors.New("packet out of space")
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err == nil {
		p.buffer = append(p.buffer, buf.Bytes()...)
		p.cursize += 4
	}

	return err
}

func (p *Packet) WriteFloat64(data float64) error {
	if !p.HasCapcity(8) {
		return errors.New("packet out of space")
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)

	if err == nil {
		p.buffer = append(p.buffer, buf.Bytes()...)
		p.cursize += 8
	}

	return err
}

func (p *Packet) WriteString(data string) error {
	buf := []byte(data)
	buf_len := len(buf)
	if !p.HasCapcity(uint(buf_len) + 4) {
		return errors.New("packet out of space")
	}
	err := p.WriteInt32(int32(buf_len))
	p.buffer = append(p.buffer, buf...)
	p.cursize += uint(buf_len)

	return err
}
func (p *Packet) HasCapcity(offset uint) bool {
	return p.capacity <= p.cursize+offset

}

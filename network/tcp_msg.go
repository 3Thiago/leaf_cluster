package network

import (
	"encoding/binary"
	"errors"
	"io"
)

// ------------------------
// | len | type |data | # |
// ------------------------

type MsgParser struct {
	minMsgLen    uint32
	maxMsgLen    uint32
	littleEndian bool
}

const (
	DefaultMaxLen = 4096
	DefaultMinLen = 4
)

func NewMsgParser() *MsgParser {
	p := new(MsgParser)
	p.minMsgLen = DefaultMinLen
	p.maxMsgLen = DefaultMaxLen

	return p
}

// It's dangerous to call the method on reading or writing
func (p *MsgParser) SetMsgLen(minMsgLen uint32, maxMsgLen uint32) {
	if minMsgLen != 0 {
		p.minMsgLen = minMsgLen
	}
	if maxMsgLen != 0 {
		p.maxMsgLen = maxMsgLen
	}
}

// It's dangerous to call the method on reading or writing
func (p *MsgParser) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// goroutine safe
func (p *MsgParser) Read(conn *TCPConn) ([]byte, error) {
	var bufMsgLen = make([]byte, 2)
	// read len
	if _, err := io.ReadFull(conn, bufMsgLen[0:]); err != nil {
		return nil, err
	}

	// parse len
	var msgLen uint32
	msgLen = uint32(binary.LittleEndian.Uint16(bufMsgLen[0:2]))
	// check len
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil, errors.New("message too short")
	}
	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		return nil, err
	}
	if msgData[msgLen] != '#' {
		return nil, errors.New("invalid message end")
	}
	return msgData, nil
}

// goroutine safe
func (p *MsgParser) Write(conn *TCPConn, args ...[]byte) error {
	// get len
	var msgLen uint32
	for i := 0; i < len(args); i++ {
		msgLen += uint32(len(args[i]))
	}

	// check len
	if msgLen > p.maxMsgLen {
		return errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return errors.New("message too short")
	}

	msg := make([]byte, 2+msgLen+1)
	// write len
	binary.LittleEndian.PutUint16(msg, uint16(msgLen))
	// write data
	l := 2
	for i := 0; i < len(args); i++ {
		copy(msg[l:], args[i])
		l += len(args[i])
	}
	msg[l] = '#'
	conn.Write(msg)

	return nil
}

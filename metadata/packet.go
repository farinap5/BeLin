package metadata

import (
	"bytes"
	"encoding/binary"
)

func ParseCmd(buf *bytes.Buffer, packLen *uint32) (uint32, []byte, error) {
	cmdTypeB := make([]byte, 4)
	_, err := buf.Read(cmdTypeB)
	if err != nil {return 0, nil, err}
	cmdType := binary.BigEndian.Uint32(cmdTypeB)

	cmdLenB := make([]byte, 4)
	_, err = buf.Read(cmdLenB)
	if err != nil {return 0, nil, err}

	cmdLen := binary.BigEndian.Uint32(cmdLenB)
	cmdBuf := make([]byte, cmdLen)
	_, err = buf.Read(commandBuf)
	if err != nil {return 0, nil, err}
	
	*totalLen = *totalLen - (4 + 4 + commandLen)

	return cmdType, 
}
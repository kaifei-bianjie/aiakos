package aiakos

import (
	"testing"
	"github.com/tendermint/tendermint/libs/log"
)

type cusLogger struct {
	log.Logger
}

func TestLoadPriValKeyToHsm(t *testing.T)  {
	filePathPriValidator := "/Users/kaifei/.iris/config/priv_validator.json.bak"
	pv := AiakosPV{
		hsmURL: "192.168.22.129:12345",
		authKeyID: uint16(1),
		password: "password",
		signingKeyID: uint16(203),
	}
	cusLogger := cusLogger{}

	hsm, err := NewAiakosPV(pv.hsmURL, pv.signingKeyID, pv.authKeyID, pv.password, cusLogger)
	if err != nil {
		t.Fatalf("init hsm fail, err is %v\n", err)
	}
	err = hsm.OnStart()
	if err != nil {
		t.Fatalf("start hsm fail, err is %v\n", err)
	}

	err = loadPriValKeyToHsm(filePathPriValidator, hsm)
	if err != nil {
		t.Fatalf("load private validator fail, err is %v\n", err)
	} else {
		t.Log("import success")
	}
}

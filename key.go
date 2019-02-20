package aiakos

import (
	pvm "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func loadPriValKeyToHsm(FilePathPriValidation string, hsm *AiakosPV) error {
	filePv := pvm.LoadFilePV(FilePathPriValidation)
	key := filePv.PrivKey.(ed25519.PrivKeyEd25519)
	err := hsm.ImportKey(uint16(hsm.signingKeyID), key[:32])

	if err != nil {
		return err
	}
	return nil
}

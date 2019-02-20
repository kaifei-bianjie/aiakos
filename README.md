aiakos
==

Aiakos is a [Tendermint](https://github.com/tendermint/tendermint) PrivValidator (signer) implementation using the YubiHSM2.

It uses the [yubihsm-go](https://github.com/certusone/yubihsm-go) library to communicate with the YubiHSM connector which needs to be running on the HSM host.

## Usage

To implement the PV interface in your Tendermint application you need to initialize it appropriately. To simplify the state management we implemented Tendermint's Service system.

The following example shows you how to implement a simple Aiakos service using env variables (please note that some errors are unhandled for simplicity).
We are using the tendermint `log` package for the service initialization and further log output.

**Step 1: import validator private key into HSM**

This operation should be done manual,
there have a example for how to import validator private key into HSM.

[key_test.go](./key_test.go)

**Step 2: Use HSM to sign vote and proposal**

```
if os.Getenv("AIAKOS_URL") == "" {
  return nil, errors.New("no Aiakos hsm url specified. Please set AIAKOS_URL in the format host:port")
}
aiakosUrl := os.Getenv("AIAKOS_URL")

if os.Getenv("AIAKOS_SIGNING_KEY") == "" {
  return nil, errors.New("no Aiakos signing key ID specified. Please set AIAKOS_SIGNING_KEY")
}
aiakosSigningKey, err := strconv.ParseUint(os.Getenv("AIAKOS_SIGNING_KEY"), 10, 16)
if err != nil {
  return nil, errors.New("invalid Aiakos signing key ID.")
}

if os.Getenv("AIAKOS_AUTH_KEY") == "" {

  return nil, errors.New("no Aiakos auth key ID specified. Please set AIAKOS_AUTH_KEY")
}
aiakosAuthKey, err := strconv.ParseUint(os.Getenv("AIAKOS_AUTH_KEY"), 10, 16)
if err != nil {
  return nil, errors.New("invalid Aiakos auth key ID.")
}

if os.Getenv("AIAKOS_AUTH_KEY_PASSWORD") == "" {
  return nil, errors.New("no Aiakos auth key password specified. Please set AIAKOS_AUTH_KEY_PASSWORD")
}
aiakosAuthPassword := os.Getenv("AIAKOS_AUTH_KEY_PASSWORD")

// Init Aiakos module
hsm, err := aiakos.NewAiakosPV(aiakosUrl, uint16(aiakosSigningKey), uint16(aiakosAuthKey), aiakosAuthPassword, log.NewNopLogger())
if err != nil {
  return nil, err
}

// Start Aiakos
err = hsm.Start()
if err != nil {
  return nil, err
}

```

Now you can use `hsm` as the PV in your Tendermint App initializer.

## Need improving

- Not handle `last_height`, `last_round`, `last_step` when implement SignVote and SignProposal method
- key can't be get wrapped and put wrapped which created use `yubihsm-go`, because those key haven't capabilities
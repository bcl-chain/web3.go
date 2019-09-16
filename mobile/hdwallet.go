package web3go

import (
	native "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	wallet *native.Wallet
}

func NewMnemonic(bits int) (string, error) {
	return native.NewMnemonic(bits)
}

func NewFromMnemonic(mnemonic string) (*Wallet, error) {
	wallet, err := native.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return &Wallet{wallet}, nil
}

func (w *Wallet) Derive(path string, pin bool) (*Account, error) {
	parsed := native.MustParseDerivationPath(path)
	account, err := w.wallet.Derive(parsed, pin)
	if err != nil {
		return nil, err
	}
	return &Account{account}, nil
}

func (w *Wallet) AddressHex(account *Account) (string, error) {
	return w.wallet.AddressHex(account.account)
}

func (w *Wallet) PrivateKeyHex(account *Account) (string, error) {
	return w.wallet.PrivateKeyHex(account.account)
}

func (w *Wallet) PublicKeyHex(account *Account) (string, error) {
	return w.wallet.PublicKeyHex(account.account)
}

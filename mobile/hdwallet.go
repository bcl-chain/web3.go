package web3go

import (
	"github.com/ethereum/go-ethereum/accounts"
	native "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	wallet *native.Wallet
}

type NewAccount struct {
	account *accounts.Account
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

func (w *Wallet) Derive(path string, pin bool) (*NewAccount, error) {
	parsed := native.MustParseDerivationPath(path)
	account, err := w.wallet.Derive(parsed, pin)
	if err != nil {
		return nil, err
	}
	return &NewAccount{&account}, nil
}

func (w *Wallet) AddressHex(account *NewAccount) (string, error) {
	return w.wallet.AddressHex(*account.account)
}

func (w *Wallet) PrivateKeyHex(account *NewAccount) (string, error) {
	return w.wallet.PrivateKeyHex(*account.account)
}

func (w *Wallet) PublicKeyHex(account *NewAccount) (string, error) {
	return w.wallet.PublicKeyHex(*account.account)
}

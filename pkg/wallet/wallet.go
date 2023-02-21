package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

//10^9 wei in a gwei
//10^9 gwei in a ether -> 10^18 wei in a ether

// Wallet holds all information to transact with a wallet
type Wallet struct {
	HolderName string

	AddressHex    string
	PublicKeyHex  string
	PrivateKeyHex string
	ChainID       string

	Balance *big.Int

	client     *ethclient.Client
	privatekey *ecdsa.PrivateKey
	publickey  *ecdsa.PublicKey
	address    common.Address
}

func NewFromPrivateKey(c *ethclient.Client, ownername, privatekey string) (Wallet, error) {
	wallet := Wallet{client: c}
	if private, err := crypto.HexToECDSA(privatekey); err != nil {
		return Wallet{}, err
	} else {
		wallet.HolderName = ownername

		wallet.publickey = private.Public().(*ecdsa.PublicKey)
		wallet.PublicKeyHex = hexutil.Encode(crypto.FromECDSAPub(wallet.publickey))[4:]

		wallet.privatekey = private
		wallet.PrivateKeyHex = privatekey

		wallet.address = crypto.PubkeyToAddress(private.PublicKey)
		wallet.AddressHex = wallet.address.Hex()
	}

	if balance, err := c.BalanceAt(context.Background(), wallet.address, nil); err != nil {
		return wallet, err
	} else {
		wallet.Balance = balance
	}

	fmt.Printf("%+v", wallet)
	return wallet, nil
}

// NewWallet generates a new private key and creates wallet
func New(c *ethclient.Client, ownername string) (Wallet, error) {
	if pk, err := crypto.GenerateKey(); err != nil {
		return Wallet{}, err
	} else {
		privateKeyBytes := crypto.FromECDSA(pk)
		return NewFromPrivateKey(c, ownername, hexutil.Encode(privateKeyBytes)[2:])
	}
}

func (w Wallet) ToKeyedTransactor(wei int64) (auth *bind.TransactOpts, err error) {
	chainid := big.NewInt(0)
	gasprice := big.NewInt(0)
	nonce := uint64(0)

	if chainid, err = w.client.ChainID(context.Background()); err != nil {
		return nil, err
	} else if auth, err = bind.NewKeyedTransactorWithChainID(w.privatekey, chainid); err != nil {
		return nil, err
	} else if nonce, err = w.client.PendingNonceAt(context.Background(), w.address); err != nil {
		return nil, err
	} else if gasprice, err = w.client.SuggestGasPrice(context.Background()); err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(wei) // in wei
	auth.GasPrice = gasprice
	auth.GasLimit = uint64(300000) // in units

	return auth, nil
}

func (w Wallet) ToContractCaller() *bind.CallOpts {
	return &bind.CallOpts{
		Pending: false,
		From:    w.address,
	}
}

func (w Wallet) ToContractTransactor() {

}

func (w Wallet) String() string {
	return fmt.Sprintf("Owner: %s\nAddress: %s\nBalance: `%d wei`\nPublic Key: %s\nPrivate Key: %s\n\n\n", w.HolderName, w.AddressHex, w.Balance, w.PublicKeyHex, w.PrivateKeyHex)
}

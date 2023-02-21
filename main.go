package main

import (
	"ethereum/pkg/wallet"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	RPCMainnet       = "https://eth.llamarpc.com"
	RPCGoerliTestnet = "https://rpc.ankr.com/eth_goerli"
	RPCGanacheLocal  = "HTTP://127.0.0.1:7545"
	client           *ethclient.Client
	defaultWallet    wallet.Wallet
)

/*
solc --abi Store.sol
solc --bin Store.sol
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go
*/

func init() {
	if c, err := ethclient.Dial(RPCGoerliTestnet); err != nil {
		panic(err)
	} else {
		client = c
	}

	if w, err := wallet.NewFromPrivateKey(client, "Devesh", "689ba212c2a5259ec3e6dc9adfed82cc08c811849eabd9e0672ce3dfbd1352c2"); err != nil {
		panic(err)
	} else {
		defaultWallet = w
	}

	// deployContract()
}

func deployContract() {
	if auth, err := defaultWallet.ToKeyedTransactor(0); err != nil {
		panic(err)
	} else if address, tx, _, err := DeployMain(auth, client); err != nil {
		panic(err)
	} else {

		fmt.Println(tx.Gas(), address)
	}
}

func main() {
	// var (
	// 	auth *bind.TransactOpts
	// 	err  error
	// )

	// if auth, err = defaultWallet.ToKeyedTransactor(10e9); err != nil {
	// 	panic(err)
	// }

	address := common.HexToAddress("0xf887CC1D615C6C189C6652107c554881ab9F4986")
	if contract, err := NewMain(address, client); err != nil {
		panic(err)
	} else {
		fmt.Println(contract.AccruementCycleLength(defaultWallet.ToContractCaller()))
		// if tx, err := (contract.(auth, big.NewInt(10e9))); err != nil {
		// 	panic(err)
		// } else {
		// 	fmt.Println(string(tx.Data()))
		// 	fmt.Println(contract.Balance(defaultWallet.ToContractCaller()))
		// }

	}

}

/*
contract Stake {
    address creator;
    uint256 annualYield;

    //StakingRecord keeps track of user stakes
    struct StakingRecord {
        uint256 StakedAmount; // in wei
        uint256 TotalExpectedYield; // in wei, basically StakedAmount * AnnualInterest
        uint256 PayoutSize; //in wei
        uint256 PaidOutTillNow; //in wei
        uint256 PayoutPeriod; //in seconds, the time between two payments
        uint256 FinalPaymentTimestamp;
        uint256 CreatedAtTimestamp; //timestamp in seconds
    }

    function CreateStake(uint256 cycles) {
        StakingRecord record = (msg.value, msg.)
    }
}

*/

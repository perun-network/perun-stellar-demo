package client

import (
	"errors"
	"log"
	"math/big"
	"perun.network/perun-stellar-backend/channel/env"
	"perun.network/perun-stellar-backend/channel/types"
	"strconv"
	"time"
)

func FormatBalance(bal *big.Int) string {
	log.Printf("balance: %s", bal.String())
	balStellar := bigIntToFloat64(bal)
	return strconv.FormatFloat(balStellar, 'f', 6, 64) + " Stellar Token"
}

func bigIntToFloat64(bi *big.Int) float64 {
	bf := new(big.Float).SetInt(bi)
	f64, _ := bf.Float64()
	return f64
}

func (p *PaymentClient) PollBalances() {
	defer log.Println("PollBalances: stopped")
	pollingInterval := time.Second

	log.Println("PollBalances")
	updateBalance := func() {

		balance := p.GetOwnBalance()

		p.balanceMutex.Lock()
		if balance.Cmp(p.balance) != 0 {
			p.balance = balance
			bal := p.balance.Int64()
			p.balanceMutex.Unlock()
			p.NotifyAllBalance(bal)
		} else {
			p.balanceMutex.Unlock()
		}
	}
	// Poll the balance every 5 seconds.
	for {
		updateBalance()
		time.Sleep(pollingInterval)
	}
}

func (p *PaymentClient) GetOwnBalance() *big.Int {

	kp := p.stellarClient.GetKeyPair()
	tokenAddr := p.tokenAddr
	// here find out how to get xdr.Address from kp
	balanceOf, err := types.MakeAccountAddress(kp)
	if err != nil {
		panic(err)
	}

	GetTokenBalanceArgs, err := env.BuildGetTokenBalanceArgs(balanceOf)
	if err != nil {
		panic(err)
	}
	txMeta, err := p.stellarClient.InvokeAndProcessHostFunction("balance", GetTokenBalanceArgs, tokenAddr, kp)
	if err != nil {
		panic(err)
	}

	bal := txMeta.V3.SorobanMeta.ReturnValue.I128

	if bal.Hi != 0 {
		panic(errors.New("balance too large - cannot be mapped to uint64"))
	} else {
		loInt64 := int64(bal.Lo)
		bigIntBal := big.NewInt(loInt64)
		return bigIntBal
	}

}

// }

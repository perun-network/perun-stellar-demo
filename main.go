// Copyright 2023 PolyCrypt GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/stellar/go/keypair"
	"log"
	"os"
	"perun.network/go-perun/wire"
	vc "perun.network/perun-demo-tui/client"
	"perun.network/perun-demo-tui/view"
	"perun.network/perun-stellar-demo/client"
	"perun.network/perun-stellar-demo/util"
)

const PerunContractPath = "./testdata/perun_soroban_contract.wasm"
const StellarAssetContractPath = "./testdata/perun_soroban_token.wasm"

func SetLogFile(path string) {
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
}
func main() {
	SetLogFile("demo.log")
	wAlice, accAlice, kpAlice := util.MakeRandPerunWallet()
	wBob, accBob, kpBob := util.MakeRandPerunWallet()
	_, _, kpDepToken := util.MakeRandPerunWallet()
	_, _, kpDepPerun := util.MakeRandPerunWallet()
	kps := []*keypair.Full{kpAlice, kpBob, kpDepToken, kpDepPerun}

	checkErr(util.CreateFundStellarAccounts(kps, len(kps), "1000000"))

	tokenAddr, _ := util.Deploy(kpDepToken, StellarAssetContractPath)
	checkErr(util.InitTokenContract(kpDepToken, tokenAddr))

	aliceAddr, err := util.MakeAccountAddress(kpAlice)
	checkErr(err)
	bobAddr, err := util.MakeAccountAddress(kpBob)
	checkErr(err)

	checkErr(util.MintToken(kpDepToken, tokenAddr, 1000000, aliceAddr))
	checkErr(util.MintToken(kpDepToken, tokenAddr, 1000000, bobAddr))

	perunAddr, _ := util.Deploy(kpDepPerun, PerunContractPath)

	bus := wire.NewLocalBus()
	alice, err := client.SetupPaymentClient("alice", wAlice, accAlice, kpAlice, tokenAddr, perunAddr, bus)
	checkErr(err)
	bob, err := client.SetupPaymentClient("bob", wBob, accBob, kpBob, tokenAddr, perunAddr, bus)
	checkErr(err)

	clients := []vc.DemoClient{alice, bob}
	_ = view.RunDemo("Perun Payment Channel on Stellar", clients)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

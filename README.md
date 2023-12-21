<h1 align="center"><br>
    <a href="https://perun.network/"><img src=".assets/go-perun.png" alt="Perun" width="196"></a>
<br></h1>

<h2 align="center">Perun Stellar Demo</h2>

This repository contains the Perun Stellar demo client, which utilizes the [go-perun](https://github.com/perun-network/go-perun) channel library, and also the [Stellar payment channel backend](https://github.com/perun-network/perun-stellar-backend)

# Setup

Clone this repository:

```
  $ git clone https://github.com/perun-network/perun-stellar-demo.git
```


Spin up the local Stellar blockchain, serving as a local testnet for demonstration purposes.

```
  $ ./quickstart.sh standalone
```

This will start the Stellar, Horizon and Soroban nodes in the background. This is the platform on which we deploy the Stellar Asset Contract (SAC), and the Perun Payment Channel contract. This allows us to create and utilize L2 channels on Stellar for any customized Stellar asset tokens,

# Using the demo

You can start the demo by simply running

```
  $ go run main.go
```
The accounts used in the demo are generated randomly and funded at the initialization stage of the demo. 

At the beginning, you will be greeted with a demo window that is split into two panes. You can play around with the different functions of the demo by using the keybinds listed below.

## Keybinds

* `ctrl+a`: Select left pane
* `ctrl+b`: Select right pane
* `tab`: Cycle through selectable fields
* `Enter`: Select or confirm highlighted field
* `r`: Go back to parent page
* `q`: Close the demo


## Exiting the demo

You can exit the demo by pressing `q` at any time. Afterwards, you can simply stop the ```kickstart.sh``` script with ctrl+c.

## Copyright

Copyright 2023 PolyCrypt GmbH. Use of the source code is governed by the Apache 2.0 license that can be found in the [LICENSE file](LICENSE).
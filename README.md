# Dharani
Selling fractional  property aka Real Estate tokenization

## Overview
Real estate is still an illiquid asset class with high barriers to entry.  The momentum to enter for smaller investors to participate is not possible. By tokenization we able to provide fractional ownership, everyone could able to buy a stake in the property. Along with that issuers can automate many middlemen functions and bring liquidity to an asset class that has been starved of it.


## Install
**Note**: required [Go 1.15+](https://golang.org/dl/)
```bash=
go get github.com/saiSunkari19/dharani

cd $GOPATH/src/github.com/saiSunkari19/dharani

make install
```
### Create Account
```bash=
dharanicli keys add validator
```
#### Init Chain
```bash
bash entry-point.sh <chain-id> <validator-address>

bash entry-point.sh dharani $(dharanicli keys show validator -a)
```

## Demo

we hosted [Dharani](https://dharani.multiverse.tk/), to test.

### Signup & Faucet
![](docs/assets/1SignupAndFaucet.gif)

More [Videos](./docs/assets/1.%20Assets%20.md)

## Documentation
This documenation provide you brief about project, with technical documentation


#### [1. Security Tokens](docs/1.%20Dharani%20(Real%20Estate%20tokenization).md)
#### [2. Real Estate Tokenization](docs/2.%20Real%20Estate%20Tokenization%20.md)
#### [3. Technical Documentation](docs/3.%20Technical%20Overview.md)
#### [4. API Documentation](docs/4.%20API%20Docs.md)



While using application, if you have any questions or find a bug, feel free to open an issue in the repo.
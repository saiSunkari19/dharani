# Dharani

Selling fractional property aka Real Estate tokenization

## Overview

Real estate is still an illiquid asset class with high barriers to entry. The momentum to enter for smaller investors to participate is not possible. By tokenization we able to provide fractional ownership, everyone could able to buy a stake in the property. Along with that issuers can automate many middlemen functions and bring liquidity to an asset class that has been starved of it.

## Install

**Note**: required [Go 1.15+](https://golang.org/dl/)

```text
go get github.com/saiSunkari19/dharani

cd $GOPATH/src/github.com/saiSunkari19/dharani

make install
```

### Create Account

```text
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

![](.gitbook/assets/1SignupAndFaucet.gif)

More [Videos](docs/assets/1.-assets.md)

##   👉  [Documentation](https://docs.dharani.multiverse.tk) 

### Highlights 🌠 

Winners of [HackAtom India](https://www.hackerearth.com/challenges/hackathon/hackatom-india), grab the  **Cosmonaut Award** 🏆 




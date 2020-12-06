# Dharani

[Nextdocs](https://app.gitbook.com/@saisunkari19/s/prithvidevs/~/drafts/-MNqqUeR4VU6QhsT5Ohk/docs)Selling fractional property aka Real Estate tokenization

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

## Documentation

This documentation provides you brief about the project, with technical documentation

#### [1. Security Tokens](docs/1.-dharani-real-estate-tokenization.md)

#### [2. Real Estate Tokenization](docs/2.-real-estate-tokenization.md)

#### [3. Technical Documentation](docs/3.-technical-overview.md)

#### [4. API Documentation](docs/4.-api-docs.md)

While using the application, if you have any questions or find a bug, feel free to open an issue in the repo.


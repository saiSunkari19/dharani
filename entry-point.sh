#!/bin/sh

CHAINID=$1
GENACCT=$2

if [ -z "$1" ]; then
  echo "Need to input chain id..."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Need to input genesis account address..."
  exit 1
fi

# Build genesis file incl account for passed address
coins="1000000000000000cent,1000000000000000mdm"
Dharanid init --chain-id $CHAINID $CHAINID --stake-denom cent
Dharanicli keys add validator
Dharanid add-genesis-account validator $coins
Dharanid add-genesis-account $GENACCT $coins
Dharanid gentx --name validator --amount 100000000cent
Dharanid collect-gentxs

#Dharanid start --pruning=nothing

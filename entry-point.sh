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
dharanid init --chain-id $CHAINID $CHAINID --stake-denom cent
dharanicli keys add validator
dharanid add-genesis-account validator $coins
dharanid add-genesis-account $GENACCT $coins
dharanid gentx --name validator --amount 100000000cent
dharanid collect-gentxs

#dharanid start --pruning=nothing

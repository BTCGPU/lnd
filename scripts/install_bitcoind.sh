#!/usr/bin/env bash

set -ev

export BITCOIND_VERSION=0.15.2

if sudo cp ~/bitcoin-gold/bitcoin-gold-$BITCOIND_VERSION/bin/bgoldd /usr/local/bin/bgoldd
then
        echo "found cached bgoldd"
else
        mkdir -p ~/bitcoin-gold && \
        pushd ~/bitcoin-gold && \
        wget https://github.com/BTCGPU/BTCGPU/releases/download/v$BITCOIND_VERSION/bitcoin-gold-$BITCOIND_VERSION-x86_64-linux-gnu.tar.gz && \
        tar xvfz bitcoin-gold-$BITCOIND_VERSION-x86_64-linux-gnu.tar.gz && \
        sudo cp ./bitcoin-gold-$BITCOIND_VERSION/bin/bgoldd /usr/local/bin/bgoldd && \
        popd
fi


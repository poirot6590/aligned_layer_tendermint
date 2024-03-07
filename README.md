# Bootcamp Verifying Lambchain

This repository contains a WIP zkSNARK verifier blockchain using Cosmos SDK and CometBFT and created with Ignite CLI.

The application interacts with zkSNARK verifiers built in Rust through FFI.

## Requirements

- Go
- Ignite

## Single Node Usage

To run a single node blockchain, run:

```sh
ignite chain serve
```

This command installs dependencies, builds, initializes, and starts your blockchain in development.

To send verify message (transaction), run:

```sh
lambchaind tx lambchain verify --from alice --chain-id lambchain <proof>
```

To get the transaction result, run:

```sh
lambchaind query tx <txhash>
```

## Configure

The blockchain in development can be configured with `config.yml`.

## Setting up multiple nodes using docker

Build docker image:
```sh
docker build -f Dockerfile . -t lambchaind_i
```

Run script:
```sh
bash multi_node_setup.sh
```
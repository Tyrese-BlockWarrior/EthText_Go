# Overview
Solidity and ÐApp experiments with [Go contract bindings](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts).

# Setup

### [Install solc](https://github.com/ethereum/solidity/releases) 
Needed only if generating Go directly from `sol` rather than from `ABI`.

### OSX
```
brew tap ethereum/ethereum
brew install solidity
brew linkapps solidity
```

### [Install Geth](https://ethereum.github.io/go-ethereum/downloads/)
The Windows installer puts it in the `path`. For other operating systems, move the binary to an existing path location such as `$GOPATH/bin`.

Alternatively, build from the source downloaded below.

### Install Solidity language support
In VSCode, the [JuanBlanco.solidity](https://marketplace.visualstudio.com/items?itemName=JuanBlanco.solidity) plugin provides syntax support and compilation.

### Install ABI to Go generator

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum
go install ./cmd/abigen
```

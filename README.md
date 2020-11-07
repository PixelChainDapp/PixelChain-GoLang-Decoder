# PixelChain GoLang Decoder

This is an open source implementation of the PixelChain decoder written in GoLang. The intention of it is to demostrate how the data is stored 100% on the blockchain and how it can be converted back to an image, in this case a PNG file.

## Pre requirements

Have GoLang installed and configured in your local environment. More information can be found here https://golang.org/doc/install

## Install (Linux/OSX)

Clone the repository and run the following commands

`cd PixelChain-GoLang-Decoder`

`cp config.json.dist config.json`

Register in https://infura.io/ and create a project. Copy the `MAINNET` https endpoint assigned to your project an replace it in your `config.json` file under the `NetworkAddress` key.

To run the application without build

`go run pixelchain.go 0`

To build the application run

`go build pixelchain.go`

## Run

To run the application execute the following in command line

`./pixelchain 0`

Replace `0` with any other ethID that you wish to retrieve the image. This ID can be found in the OpenSea URL. 

FE: https://opensea.io/assets/0xbc0e164ee423b7800e355b012c06446e28b1a29d/0

The image will be stored in the path defined in the `config.json` file. By default will be stored under `images`

## Ethereum and GoLang

If you wish to read more about the libraries used to communicate with the blockchain read https://github.com/miguelmota/ethereum-development-with-go-book

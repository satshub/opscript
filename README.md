## OP_SCRIPT

A viewer and debugger of Bitcoin scripts. **Early development.**

![Screenshot](./screenshot.png)


## Features
1. Can fetch transactions from:
    1. A Bitcoin node (requires a full node with `txindex=1`). Default.
    1. [Blockstream.info](https://blockstream.info) JSON API. Use `--blockstream` flag.
1. Automatically finds related output.
1. Allows to navigate forward and backward.
1. Shows stack per line of code.
1. Shows opcodes information (hex code, input, output, and description).
1. Supports witness data (SegWit).
1. Uses [`btcd/txscript`](https://github.com/btcsuite/btcd/tree/master/txscript) under the hood.


## Usage

1. `go get github.com/Jeiwan/opscript`
1. `opscript help`
    ```shell
    Usage:
    opscript [flags] transactionHash:inputIndex
    opscript [command]

    Available Commands:
    buildspec
    help        Help about any command

    Flags:
        --blockstream        Use blockstream.info API to get transactions.
    -h, --help               help for opscript
		--hex                decode raw transaction from hex
        --node               Use Bitcoin node to get transactions (requires 'txindex=1'). (default true)
        --node-addr string   Bitcoin node address. (default "127.0.0.1:8332")
        --rpc-pass string    Bitcoin JSON-RPC password.
        --rpc-user string    Bitcoin JSON-RPC username.

    Use "opscript [command] --help" for more information about a command.
    ```

## Key bindings

* `q` – quit
* `j`/`k` – navigate between lines of code(vimer customized)


## Examples
* Using Blockstream.info API:
    ```shell
    opscript --blockstream 70fde4687efab8dae09737f87e30042030288fec42fd9e12f34c435cdeb7812c
    ```
* Specifying input index:
    ```shell
    opscript --blockstream 70fde4687efab8dae09737f87e30042030288fec42fd9e12f34c435cdeb7812c:0
    ```
* Using a Bitcoin node:
    ```shell
    opscript --rpc-user=woot --rpc-pass=woot 70fde4687efab8dae09737f87e30042030288fec42fd9e12f34c435cdeb7812c
    ```
* Using a command line hex input:
    ```shell
    opscript --hex 01000000000101c4ac2944eb1bf09cd61fdfdb74fd3130256d1ad741f832c13f53713597f465200100000000ffffffff0fd08d3d000000000017a914bcdbf9d5189f59a6b437ba9e3e039058c6f1838787b2c2ce0300000000160014dc6bf86354105de2fcd9868a2b0376d6731cb92f5d42040000000000160014344f054feaa82bbf22c1f33ccbb3732a12af899f4756010000000000160014dfbaa9776f03a310cb81d7dd2b6442a401ad7569109e01000000000016001400dd4e69dcebc34edd8d5d65066d2034c64718c2b72518000000000017a91497f50406c4f79b3c19599c7194174a261edec39a877ef301000000000017a914a6d174870c8a72c7818a1aeb1f99cab83cad2b55876055040000000000220020096ece310718f2988b8ffd6b11246d61355dd4f0f995183b0465a00bd692f662c301060000000000160014b2e22e7d0e9576c3a6066ed4818f35673f63d3f7d66cad000000000017a914d29afacc14cfddf60d10341cbe6b93c60663308c87ec392f000000000017a9149f59d4888b86dfb22ea4cfb7eec34c44b78dccb087f818030000000000160014889af39b595949d5018bf39fa69413bdeb890f1b5b59120c0000000016001454776b66a31287f2ad67d928af0280f9bfaa482160a40100000000001976a914a58fa3b6a328f1ef2f4593d728797b885c3a869c88ac80a30400000000001976a91424bd0da14b9ee9840a3ea6b9bfce049044ed6ae988ac02483045022100b874342569d98dc2dab4760a2d97ca0e1deedae27fe43109b5aad61388be110902201648db39959a5a85c7f6186ee215b69e16aac6016c96a352a7fad0d4ae750d20012102174ee672429ff94304321cdae1fc1e487edf658b34bd1d36da03761658a2bb0900000000
    ```

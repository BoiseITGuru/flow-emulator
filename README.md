# EmeraldCity-CLI

:warning: THIS PROJECT IS UNDER ACTIVE DEVELOPMENT AND NOT READY FOR PRODUCTION USAGE!

The EmeraldCity-CLI project was developed to allow the [Emerald City Playground](https://github.com/BoiseITGuru/tx-script-utility/tree/dev-local-emulator) to start and manage a local Flow Emulator. "Phase 2" of the project will utlize the existing WebSocket connection to allow developers to utilize their local workstation as a file system for the Emerald City Playground.

## Challenges

The ultimate goal is to utilize [Overflow](https://github.com/bjartek/overflow) to start and manage a Flow Emulator. Currently however, there is an issue with with the emulator Overflow creates that does not expose the needed ports for the gRPC, REST API, and admin servers. To get around this you will need to temporarily have the FLow CLI installed as we use this to start the emulator before activating Overflow to allow the Emeralid City PLayground to manage it over WebSockets.

## Current Status

The playground and CLI are still undergoing heavy development, the CLI is written in GO and requires manual compiling still. Once we have finished development it will be available as an NPM package. Currently, the CLI connects to the playground and starts the emulators as well as executes scripts.

## How to use the EmeraldCity-CLI

Current Commands **propose we create the shortname "ecd" to invoke the CLI **

1. ```EmeraldCity-CLI emulator```
    * Starts a WebSocket server to await connections from the Emerald City Playground and upon connection start the Flow Emulator and Dev Wallet. The defualt confiugration starts the Flow Emulator with the ```--contracts``` flag set to automatically deploy all base contracts.
        * Flags - all flags for the Flow Emulator are supported with the following exceptions/addtions:
            1. The ```--contracts``` flag has been renamed to ```--no-contracts``` and will start the emulator without the NonFungibleToken, MetadataViews, ExampleNFT, or NFTStorefront contracts.
            2. ```--no-ws-server``` will start the Flow Emulator without creating a WebSocket server to wait for the Emerald City Playground to connect.
            3. ```--no-dev-wallet``` will start the Flow Emulator without start the Flow Dev Wallet
                * Currently none of the dev-wallet flags are supported but we are working on adding them
2. ```EmeraldCity-CLI create-fcl-app <nextjs/svelte>``` **COMING SOON - IN DEVELOPMENT**
    * Deploys FCL Quick Start project code sets chossing either the fcl-nextjs-quickstart or the fcl-svelte-quickstart project templates
        1. ```--name <name>```
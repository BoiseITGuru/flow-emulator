# EmeraldCity-CLI

:warning: THIS PROJECT IS UNDER ACTIVE DEVELOPMENT AND NOT READY FOR PRODUCTION USAGE!

The EmeraldCity-CLI project was developed to allow the [Emerald City Playground](https://github.com/BoiseITGuru/tx-script-utility/tree/dev-local-emulator) to start and manage a local Flow Emulator. "Phase 2" of the project will utlize the existing WebSocket connection to allow developers to utilize their local workstation as a file system for the Emerald City Playground.

## Challenges

The ultimate goal is to utilize [Overflow](https://github.com/bjartek/overflow) to start and manage a Flow Emulator. Currently however, there is an issue with with the emulator Overflow creates that does not expose the needed ports for the gRPC, REST API, and admin servers. To get around this you will need to temporarily have the FLow CLI installed as we use this to start the emulator before activating Overflow to allow the Emeralid City PLayground to manage it over WebSockets.

## Current Status

The playground and CLI are still undergoing heavy development, the CLI is written in GO and requires manual compiling still. Once we have finished development it will be available as an NPM package. Currently, the CLI connects to the playground and starts the emulators as well as executes scripts.

## How to use the EmeraldCity-CLI

Current Commands

1. ecd emulator
    * Stars the Flow Emulator and WebSocket server to await connections from the Emerald City Playground.
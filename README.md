# cosmos-app
Testing cosmos blockchain by building an app specific blockchain
#### Key concepts:
      1. Tendermint for consensus.
      2. Cosmos SDK is using for routing like ISP.
      3. IBC protocol is used for interchain communication like TCP.
      4. ABCI is the main interface by which cosmos sdk interact with tendermint. and app port should only to open to tendermint.
      5. Cosmos SDK comes with build modules and any custom modules can built and integrated with the base modules.
      6. Module's key's and keepers are for managing intermodule communication
      7. Module manager is used for registering msg service and grpc services.
      8. SDK also has appCodec to handle data serde and convert to []bytes
      9. Starport CLI is used for installing dependencies, type and proto definition and compiling to tendermint binary. It can also be used for tesing simulation of the blockchain app.
  

# blog
**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## 1. Explain what does it mean by breaking consensus.
Breaking consensus means making changes to the blockchain that alters the rules of how nodes in a blockchain network agrees on the state of the blockchain. For example, changing the way nodes validate transactions or blocks on the network leads to nodes disagreeing on the state of the blockchain.

## 2. Explain why your change would break the consensus.
This change adds a "Timestamp" field to the "Post" type. Changing the data structure of "Post" modifies the way that data is stored, processed, and validated across the network. Nodes running the updated code will know to expect the "Timestamp" field while nodes running the old code will not recognize the new data.
## BNB Smart Chain - fork

This repo aims to introduce new improvements over the default geth. 

The public version should be considered more as of an guide for new users. 

Private version contact: Discord: Exca#0775

## Instalation
```
git clone https://github.com/Exca-DK/bsc_extras.git
cd bsc_extras
make geth
cd ..
mv ./bsc_extras/build/bin/geth ./geth
```

## Key changes/features

- ### Metadata

Bnb currently is swarmed with fake nodes which have negative impact for overall performance of the chain. Metadata can be used to filter the most basic nodes due to their inactivity or spamming actions.
 
The change is reflected in method admin_peers as extra field called metadata. Ref: https://geth.ethereum.org/docs/rpc/ns-admin Example:
```
curl --data '{"method":"admin_peers","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:PORT
```

| stat | description |
| ------ | ------ |
| value | total amount of received packets from peer |
| cache | hashes of recent transactions/blocks |
| lastActivity | timestamp of packet 
| connectedAt | timestamp of peer connection |

By default only last 10 actions are being cached. The value can be specified with metadata.cache.size flag

- ## Censorship
  
Fake nodes and toxic bots on the network can be additionaly censored by forbiding specific accounts from propagation. As node operation you will be still receiving the data of such accounts but the transactions won't be relayed to connected peers.
 
| function | description |
| ------ | ------ |
| blacklistedPropagation | returns currently blocked addressess |
| unblacklistPropagation | unblocks address if blocked |
| blacklistPropagation | blocks address | 

```
Unblock address:
curl --data '{"method":"mev_unblacklistPropagation","params":["address"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:PORT

Block address:
curl --data '{"method":"mev_blacklistPropagation","params":["address"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:PORT

Check addresses:
curl --data '{"method":"mev_blacklistedPropagation","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:PORT
```

##### Note: Those functions have been added to "mev" module. If you have http/ws.api flag enabled then you need to also speficy that module, eg. --http.api "eth,admin,mev".

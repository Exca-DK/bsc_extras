## BNB Smart Chain - fork

This repo aims to introduce new improvements over the default geth. 

The public version should be considered more as of an guide for new users.

Private version contact: Discord: Exca#0775

## Key changes/features

### P2P layer
Bnb currently is swarmed with fake nodes which have negative impact for overall performance of the chain.

As of new I decided to share only one change which is extra metadata for each peer. Such metadata can be used to filter the most basic nodes due to their inactivity or spamming actions.
 
The change is reflected in method admin_peers as extra field called metadata. Ref: https://geth.ethereum.org/docs/rpc/ns-admin 

| stat | description |
| ------ | ------ |
| value | total amount of received packets from peer |
| cache | hashes of transactions/blocks |
| lastActivity | timestamp of packet 
| connectedAt | timestamp of peer connection |

By default only last 10 actions are being cached. The value can be specified with metadata.cache.size flag
 


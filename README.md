# discord-vanity-reclaim


**discord-vanity-reclaim** is a high-performance tool for reclaiming Discord vanity URLs during server transfers. 

**NOTE: It is recommended to first test on a free/claimable vanity before attempting to transfer a high-tier vanity.**

## **Features** :
- Automatically send notifications to webhook
- Supports sniping a lot of vanities simultaneously
- Advanced proxy switching mechanism
- Multiple modes used to claim the vanities (fasthttp and sockets)


## Basic Usage
1) Building the program by source
2) Configure your "config.yml" file ( webhook, guildid, amplify and token)
3) Add your proxies to the "proxies.txt" file
4) Add your vanity list to the "vanities.txt" file
5) Run the binary and wait for the program to claim


## How to build from source
1) First download the suitable golang version from the following link: https://go.dev/dl/
2) Clone / download the code from this github repository
3) Either cd to the directory where the code is or open a terminal in the directory
4) Type go build .
5) Then you should have a executable which you can then use to claim


## Configuration

| Name | Description | 
| ---  | ---  |
| `amplify` | Amount of goroutines per vanity, if there are 5 vanities and amplify is 6 it means there will be 6 goroutines per vanity
| `token` | Discord token which has access to changing the vanity of server
| `webhook` | Webhook for sending notifications
| `guildid` | Guildid of server which has access to the vanity feature
| `usesockets` | It is recommended to leave this option disabled unless you know how the code works, as it is currently experimental
| `socketchannels` | Amount of socket channels, ignore this if you are not using the usesockets option

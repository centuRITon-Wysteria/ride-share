<a href="https://github.com/centuRITon-Wysteria/ride-share/blob/master/LICENSE">![LICENSE](https://img.shields.io/github/license/centuRITon-Wysteria/ride-share)</a>
<a href="https://github.com/centuRITon-Wysteria/ride-share">![GitHub last commit](https://img.shields.io/github/last-commit/centuRITon-Wysteria/ride-share?label=GitHub)</a>

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/open-source.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/contains-tasty-spaghetti-code.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-coffee.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/not-a-bug-a-feature.svg)](https://forthebadge.com)

# ride-share[baby-chain]

#### repo on GitHub - [Centuriton-Wysteria-Rideshare](https://github.com/centuRITon-Wysteria/ride-share)

#### generated from our template - [baby-chain](https://github.com/centuRITon-Wysteria/baby-chain)

### A Decentralized

- 1

**Details**:

- **Technology stack**:
  > **Status**: **`alpha`** focusing on middleware integration<br>
  > ###### _Tested on go 1.19, Windows 10_
- **Future work**:
  > P2P network on external IP<br>
  > writing **unit tests** for gpp<br>
  > expanding announcement protocols<br>
  > writing in-code docs<br>
  > writing **ReadTheDocs**<br>
  > **blogging**<br>

---

**Screenshots**:

<div align="center">

</div>

---

## Dependencies

- `go>=1.19`
- `github.com/gin-gonic/gin>=1.8.1`
- `github.com/ubiq/g0-ubiq>=3.0.1`

for production dependencies see [Requirements](gpp/go.sum)<br>

## Installation

```shell
$ go mod tidy
```

<a id="usage"></a>

## Usage

**Setup**

```go
```

**Basic usage**


- ###### 1

  > To run the blockchain run the main.go file in gpp - This initializes all the blockchain protocols required and starts the server

    ```go
  package main
    
    import (
    "blockchain/block"
    "blockchain/blockchain"
    "fmt"
    "github.com/gin-gonic/gin"
    "gpp/chain"
    "gpp/services"
    "log"
    )
    
    func Init() {
    sd := chain.LoadStateData()
    stts := chain.LoadStates()
    stts.Init(&sd)
    bc := blockchain.New(block.Data{})
    if err := chain.SaveStateData(&sd); err != nil {
    fmt.Println(err)
    }
    if err := chain.SaveBlockchain(&bc); err != nil {
    fmt.Println(err)
    }
    }
    
    func main() {
    if false {
    Init()
    } else {
    chainName := "baby_chain"
    server := gin.Default()
    basePath := server.Group("/" + chainName)
    services.RegisterClientRoutes(basePath)
    log.Fatalln(server.Run(":9090"))
    }
    }
    ```

for detailed usage see [EXAMPLE](scripts/main.py)

## API
**1. REGISTER NEW NODE : Registers a new node with the blockchain**

**[URL] : ip_address_of_node:9090/baby_chain/service/newnode**

Request Body
```json
    {
      "ip_address": <ip address of node to be registered>
    }
```
Response Body
```json
    {
      "public_key": <128 characters long unique public key>,
      "private_key": <64 characters long unique private key>
    }
```


**2. REGISTER PROFILE : Registers user with given profile parameters**

**[URL] : ip_address_of_node:9090/baby_chain/service/publicinfo**

Request Body
```json
   {
      "ip_address":<ip address of node to be registered>,
      "name":<name of the registering user>,
      "driver":<true if user is a driver,false otherwise>,
      "license":<licence plate number of the user or driver>,
      "email":<email id of the user>,
      "occupation":<occupation of the user>,
      "hobbies":<hobbies of the user>,
      "skills":<skills of the user>,
      "interests":<interests of the user>,
      "others":<any other details that the user wants to share>
   }
```
Response Body
```json
    {
      "public_key": <128 characters long unique public key>,
      "private_key": <64 characters long unique private key>
    }
```


**3. ANNOUNCE TRAVEL : Announces ride request details to all the driver nodes**

**[URL] : ip_address_of_node:9090/baby_chain/service/announcetravel**

Request Body
```json
   {
      "from_lat":<latitude of from location>,
      "from_lon":<longitude of from location>,
      "to_lat":<latitude of to location>,
      "to_lon":<longitude of to location>,
      "time":<proposed ride timings>,
      "public_key": <public key of the user>,
      "private_key": <private key of the user>
   }
```
Response Body
```json
    {
      "public_key": <128 characters long unique public key>,
      "private_key": <64 characters long unique private key>
    }
```


**4. MESSAGING : Sends message to the given public address**

**[URL] : ip_address_of_node:9090/baby_chain/service/messaging**

Request Body
```json
   {
      "type":<send if the user wants to send,receive otherwise>,
      "public_key":<public key of the sender>,
      "key":<private key of the sender>,
      "to_public_key":<public key of the receiver>,
      "message":<message to be sent>
   }
```
Response Body
```json
    {
      "message": "ok"
    }
```
## How to test the software

```shell
go test blockchain
```

```shell
go test gpp
```

## Known issues

- **_`Can't connect to devices outside the network`_**

## Getting help

- Check [main](gpp/main.go) for detailed usage
- Check [Usage](#usage) Section

If you have questions, concerns, bug reports, etc.
please file an [issue](https://github.com/centuRITon-Wysteria/baby-chain/issues) in this repository's Issue Tracker or
open a [discussion](https://github.com/centuRITon-Wysteria/baby-chain/discussions) in this repository's Discussion
section.

## Getting involved

<a id="contribute"></a>

- `Looking for Contributors for feature additions`
- `Looking for Contributors for optimization`
- `Looking for Contributors for unit testing`
- `Looking for Contributors for ReadTheDocs`
- `Looking for Contributors for WebApp`
- `Looking for Contributors for API `
- [Fork](https://github.com/centuRITon-Wysteria/baby-chain/fork) the repository
  and issue a [PR](https://github.com/centuRITon-Wysteria/baby-chain/pulls) to contribute

General instructions on _how_ to contribute [CONTRIBUTING](CONTRIBUTING.md)
and [CODE OF CONDUCT](CODE_OF_CONDUCT.md)

----

## Open source licensing info

1. [TERMS](TERMS.md)
2. [LICENSE](LICENSE)
3. [CFPB Source Code Policy](https://github.com/cfpb/source-code-policy/)

----

## Credits and references

1. 1

**Maintainer**:

|        <a href="https://github.com/Amith225"><img src="https://avatars.githubusercontent.com/u/75326634?v=4" height=250></a>        |
|:-----------------------------------------------------------------------------------------------------------------------------------:|
|                                    **[Amith M](https://www.linkedin.com/in/amith-m-17088b246/)**                                    |
| [![Instagram](https://img.shields.io/badge/Instagram-%23E4405F.svg?logo=Instagram&logoColor=white)](https://instagram.com/amithm3 ) |

| <a href="https://github.com/hari-vishwanath"><img src=https://avatars.githubusercontent.com/u/108003385?v=4" height=250></a> |
|:----------------------------------------------------------------------------------------------------------------------------:|
|                        **[Hari Vishwanath](https://www.linkedin.com/in/hari-vishwanath-694aaa243/)**                         |

**App and UI Devs**

| <a href="https://github.com/ScaraTB"><img src="https://avatars.githubusercontent.com/u/66898000?v=4" height=250></a> |
|:--------------------------------------------------------------------------------------------------------------------:|
|                                                  **[Sumukha G]()**                                                   |

| <a href="https://github.com/yashpapa6969"><img src="https://avatars.githubusercontent.com/u/108118711?v=4" height=250></a> |
|:--------------------------------------------------------------------------------------------------------------------------:|
|                            **[Yash Gupta](https://www.linkedin.com/in/yash-gupta-7070aa214/)**                             |


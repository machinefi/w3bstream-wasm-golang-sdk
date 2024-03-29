# How to use
1. build a http.Request, must have a header `eventType`
2. create a new `Event Routing` on W3bstream
3. call the api.Call() function, and get a http.Response
4. get data from the http.Response
5. handle the `eventType` result

# Hello World

Get a hello.

**example** : examples/hello/main.go

**URL** : `/system/hello`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Response examples**

```json
{
	"message": "hello w3bstream"
}
```

**Async Response examples**

```json
{
	"async message": "hello w3bstream"
}
```

# Read Eth-Compatible Transaction

Read a transaction from blockchains which follow [Ethereum JSON-RPC specification](https://ethereum.github.io/execution-apis/api-documentation/). The supported blockchains include:

| chainName | chainID |
| --------- | ------- |
| iotex-mainnet | 4689 |
| iotex-testnet | 4690 |
| ethereum-mainnet | 1 |
| goerli | 5 |
| polygon-mainnet | 137 |
| mumbai | 80001 |
| arbitrum-goerli | 421613 |
| arbitrum-one | 42161 |
| op-goerli | 420 |
| op-mainnet | 10 |
| base-goerli | 84531 |
| base-mainnet | 8453 |

**example** : examples/read_tx/main.go

**URL** : `/system/read_tx`

**Method** : `GET`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`  

**Request examples**

read transaction from iotex-testnet
```json
{
    "chainName": "iotex-testnet",
    "hash": "fcaf377ff3cc785d60c58de7e121d6a2e79e1c58c189ea8641f3ea61f7605285"
}
```

**Response examples**

```json
{
   "transaction":{
      "type":"0x0",
      "nonce":"0x53",
      "gasPrice":"0xe8d4a51000",
      "maxPriorityFeePerGas":null,
      "maxFeePerGas":null,
      "gas":"0x997f",
      "value":"0x0",
      "input":"0xf3fef3a3000000000000000000000000fdff3eafde9a0cc42d18aab2a7454b1105f19edf00000000000000000000000000000000000000000000002086ac351052600000",
      "v":"0x1b",
      "r":"0xadff1da88c93f4e80c27bab0d613147fb7aeeed6e976231695de52cd9ac5aa8a",
      "s":"0x3094e02759b838514f8376e05ceb266badc791ac2e7045ee7c15e58fc626980b",
      "to":"0xd313b3131e238c635f2fe4a84eadad71b3ed25fa",
      "hash":"0x441ab0b3af3b4de256fe95efcdbf73b7652cf260b89cf8746e0aea486c0467bd"
   }
}
```

# Send Eth-Compatible Transaction

Send transaction to blockchains which follow Ethereum JSON-RPC specification.

**example** : examples/send_eth_tx/main.go

**URL** : `/system/send_tx`

**Method** : `POST`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`  

**Request examples**

```json
{
    "chainName": "iotex-testnet",
    "operatorName": "default",
    "to": "0x1ED83F5AD999262eC06Ed8f3B801e108024b3e9c",
    "value": "0",
    "data": "40c10f1900000000000000000000000097186a21fa8e7955c0f154f960d588c3aca44f140000000000000000000000000000000000000000000000000de0b6b3a7640000"
}
```

**Response examples**

```json
{
	"transactionID": "11288089962365952"
}
```

**Async Response examples**

```json
{
	"transactionID": "11288089962365952",
	"state": "CONFIRMED"
}
```

# Read Solana Transaction

Read solana chain transaction.

**example** : examples/read_solana_tx/main.go

**URL** : `/system/read_tx`

**Method** : `GET`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`  

**Request examples**

```json
{
   "chainName":"solana-devnet",
   "hash":"5FYQ3TEG56TaYfUC7UohgAmNR6cuSQCJA9eiwzMxwXXpxsi2FvAK2bxm9oCZLLRzZiYEVwozn2MBmHnmR2mXqd99"
}
```

**Response examples**

```json
{
   "result":{
      "Slot":235318864,
      "Meta":{
         "Err":null,
         "Fee":5000,
         "PreBalances":[
            11999990000,
            1
         ],
         "PostBalances":[
            11999985000,
            1
         ],
         "PreTokenBalances":[
            
         ],
         "PostTokenBalances":[
            
         ],
         "LogMessages":[
            "Program 11111111111111111111111111111111 invoke [1]",
            "Program 11111111111111111111111111111111 success"
         ],
         "InnerInstructions":[
            
         ],
         "LoadedAddresses":{
            "writable":[
               
            ],
            "readonly":[
               
            ]
         },
         "ReturnData":null,
         "ComputeUnitsConsumed":0
      },
      "Transaction":{
         "Signatures":[
            "1JuedY5d+oYTEcSXz2WkluB9/caDX4czUf9iI9FrtW/GaswfkyO6GuIL+o2K57oY1qsthuvq/zDuunXKE8WCCA=="
         ],
         "Message":{
            "Version":"legacy",
            "Header":{
               "NumRequireSignatures":1,
               "NumReadonlySignedAccounts":0,
               "NumReadonlyUnsignedAccounts":1
            },
            "Accounts":[
               "79VcKNkFH9zgHxRmy8skRbWVaV4LCbkjZjVTrDPCcBf6",
               "11111111111111111111111111111111"
            ],
            "RecentBlockHash":"CXv7Cjf7VeFXzdY4hzTMsndU587eA5C7jmb9XexoHRoe",
            "Instructions":[
               {
                  "ProgramIDIndex":1,
                  "Accounts":[
                     0,
                     0
                  ],
                  "Data":"AgAAAAEAAAAAAAAA"
               }
            ],
            "AddressLookupTables":[
               
            ]
         }
      },
      "BlockTime":1691479501,
      "AccountKeys":[
         "79VcKNkFH9zgHxRmy8skRbWVaV4LCbkjZjVTrDPCcBf6",
         "11111111111111111111111111111111"
      ]
   }
}
```

# Send Solana Transaction

Send solana chain transaction.

**example** : examples/send_solana_tx/main.go

**URL** : `/system/send_tx`

**Method** : `POST`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`  

**Request examples**

```json
{
   "chainName":"solana-devnet",
   "operatorName":"solana-key",
   "data":"[{\"ProgramID\":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],\"Accounts\":[{\"PubKey\":[91,83,29,193,46,31,234,109,208,211,168,16,189,248,144,184,82,206,5,207,47,237,60,0,252,70,215,201,95,8,82,113],\"IsSigner\":true,\"IsWritable\":true},{\"PubKey\":[91,83,29,193,46,31,234,109,208,211,168,16,189,248,144,184,82,206,5,207,47,237,60,0,252,70,215,201,95,8,82,113],\"IsSigner\":false,\"IsWritable\":true}],\"Data\":\"AgAAAAEAAAAAAAAA\"}]"
}
```

**Response examples**

```json
{
	"transactionID": "11288089962365953"
}
```

**Async Response examples**

```json
{
	"transactionID": "11288089962365953",
	"state": "CONFIRMED"
}
```

# Send Zero-knowledge proof 

Send Zero-knowledge proof.

> Get imageID and params information from this table

| model name     | imageID                                                                                         | params description                                                                                              |
|----------------|-------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| addition       | 3863567922, 2404777232, 2693943780, 4207360975, 1147152290, 467959105, 69766310, 1627741011     | two integer strings, separated by `,`, like `"23" "17"`                                                         |
| multiplication | 3585919858, 3547600316, 2236080185, 1080621844, 212293629, 3683389566, 71769817, 967938330      | two integer strings, separated by `,`, like `"23" "17"`                                                         |
| subtraction    | 1753901548, 3541261194, 4025921735, 2828974925, 4201143231, 1100387519, 4193452787, 3597377440  | two integer strings, separated by `,`. first one is larger than the second one, like `"23" "17"`           |
| range          | 3145991386, 3471678490, 3632776032, 2595288688, 1478438623, 4259749138, 987879707, 1456846509   | three integer strings, separated by `,`. first one is between the second and the third, like `"15" "10" "23"`   |                 
| greater        | 123160077, 2552749566, 572037666, 1916298776, 3529235965, 4225778306, 449212140, 1038195090     | two integer strings, separated by `,`. first one is larger than the second one, like `"23" "17"`           |

**example** : examples/risc_zkp/main.go

**URL** : `/system/gen_zk_proof`

**Method** : `POST`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`

**Request examples**

```json
{
   "imageID":"3145991386, 3471678490, 3632776032, 2595288688, 1478438623, 4259749138, 987879707, 1456846509",
   "privateInput":"16",
   "publicInput":"4,30",
   "receiptType": "Stark"
}
```

**Response examples**
the size of proof big, so here just show the fields. 
```json
{
  "inner":{
    "Flat":[
      {
        "seal":[0,0,0,0,0,0,0,0,0,0,2194396090,...,3731194972,1340921636,1146046463,2264386654,1031639317],
        "index":0,
        "hashfn":"sha-256"
      }
    ]
  },
  "journal":[44,0,0,0,73,32,107,110,111,119,32,116,104,101,32,114,101,115,117,108,116,32,105,115,32,49,51,44,32,97,110,100,32,73,32,99,97,110,32,112,114,111,118,101,32,105,116,33]}

```

# Send Eth-Compatible Transaction With Account Abstraction Paymaster

Send userop to blockchains.  
Need create a operator with paymaster key on W3bstream studio first.  
Just support `iotex-testnet` now.  

**example** : examples/send_eth_userop_transfer/main.go  
**example** : examples/send_eth_userop_contract_call/main.go  

**URL** : `/system/send_tx`

**Method** : `POST`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`  

**Request examples**

```json
{
   "chainName":"iotex-testnet",
   "operatorName":"aa-op",
   "to":"0x065e1164818487818E6BA714E8d80B91718ad758",
   "value":"0.001",
   "data":"0x"
}
```

**Response examples**

```json
{
	"transactionID": "13540418344423426"
}
```

**Async Response examples**

```json
{
	"transactionID": "13540418344423426",
	"state": "CONFIRMED"
}
```

# How to use
1. build a http.Request, must have a header `eventType`
2. create a new `Event Routing`
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

```string
hello w3bstream
```

# Read ETH Transaction

Read a evm-compatible blockchain's transaction.

**example** : examples/read_tx/main.go

**URL** : `/system/read_tx`

**Method** : `GET`

## Success Response

**Code** : `200 OK`  
**Code** : `400 Bad Request`  
**Code** : `500 Internal Server Error`  

**Request examples**

```json
{
    "chainID": 4690,
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
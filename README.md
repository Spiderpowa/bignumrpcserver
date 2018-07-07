# Big Number RPC Server
[![CircleCI](https://circleci.com/gh/Spiderpowa/bignumrpcserver.svg?style=shield&circle-token=9ed65ceb38d8216b07240eacb2e2582a26025765)](https://circleci.com/gh/Spiderpowa/bignumrpcserver)

The server allows creating/updating/deleting named number objects that the server manages for clients.
The server also supports basic arithmetic operations on these named number objects: addition, subtraction, multiplication, and division.

## API Reference
All of the parameters are strings.

If the number of parameters are two or more, you have to wrap the parameters into an additional array. That is, the JRON-RPC request shuold look like this.

```
{
    "jsonrpc":"1.0",
    "method":"BigNumber.Create",
    "params":[
        ["AnswerToLife", "42"]
    ],
    "id":1
}
```

### BigNumber.Create
Create a named number objects

#### Parameters
1. `NAME`
2. `NUMBER`

#### Returns
none

---
### BigNumber.Update
Update a named number objects

#### Parameters
1. `NAME`
2. `NUMBER`

#### Returns
none

---
### BigNumber.Delete
Delete a named number objects

#### Parameters
1. `NAME`

#### Returns
none

---
### BigNumber.Add
Add two named number objects or constant numbers
### Parameters
1. `NAME|NUMBER`
2. `NAME|NUMBER`

#### Returns
1. `NUMBER`

---
### BigNumber.Substract
Substract two named number objects or constant numbers
### Parameters
1. `NAME|NUMBER`
2. `NAME|NUMBER`

#### Returns
1. `NUMBER`

---
### BigNumber.Multiply
Multiply two named number objects or constant numbers
### Parameters
1. `NAME|NUMBER`
2. `NAME|NUMBER`

#### Returns
1. `NUMBER`

---
### BigNumber.Division
Calculate division of two named number objects or constant numbers
### Parameters
1. `NAME|NUMBER`
2. `NAME|NUMBER`

#### Returns
1. `NUMBER`

### TOOL:verify which pubkey keys contribute to multisig

#### Notice:

- just only apply to BTC transaction multisig
- you will get nothing when input P2PKH type transaction
- do not support segwit type transaction so far


#### Usage:

- git clone 
- cd this repository path
- go build main.go
- run `./main` And you will fetch help document of this tool
- ./main rawtx-string And you will get your wanted result

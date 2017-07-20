# make-my-life-easier
A couple scripts I use on a daily basis to make my life easier.
They are all ready to go command line tools :)

## Price encryption / decryption (Google Private Data specs)
### Usage examples
#### Decrypt only
The command line below will try to decrypt the encrypted price passed in command line.
```
go run google-private-data-encryption-test.go --encryptionkey=[YOUR_EKEY] --integritykey=[YOUR_IKEY] --price=[ENCRYPTED_PRICE] --mode=decrypt --keyDecodingMode=utf-8
```
#### Encrypt & Decrypt
The command line below will try to encrypt and then decrypt the encrypted price passed in command line.
```
go run google-private-data-encryption-test.go --encryptionkey=[YOUR_EKEY] --integritykey=[YOUR_IKEY] --price=[CLEAR_PRICE] --mode=all --keyDecodingMode=hexa
```

## Hexa keys generation
### Usage examples
The command line below will generate an hexa key of the desired length (by default, 64 bytes).
```
go run hex-key-generator.go
go run hex-key-generator.go --keylength=32
```

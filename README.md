# make-my-life-easier
A couple scripts I use on a daily basis to make my life easier.
They are all ready to go command line tools :)

## Usage examples
### Decrypt only
The command line below will try to decrypt the encrypted price passed in command line.
```
go run google-private-data-encryption-test.go --encryptionkey=[YOUR_EKEY] --integritykey=[YOUR_IKEY] --price=[ENCRYPTED_PRICE] --mode=decrypt --keyDecodingMode=utf-8
```
### Encrypt & Decrypt
The command line below will try to encrypt and then decrypt the encrypted price passed in command line.
```
go run google-private-data-encryption-test.go --encryptionkey=[YOUR_EKEY] --integritykey=[YOUR_IKEY] --price=[CLEAR_PRICE] --mode=all --keyDecodingMode=hexa

# make-my-life-easier
A couple scripts I use on a daily basis to make my life easier.
They are all ready to go command line tools :)

## Price encryption / decryption (Google Private Data specs)
### Usage examples
#### Help
```
go run mmle.go help
usage: mmle [<flags>] <command> [<args> ...]

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose  Verbose mode.

Commands:
  help [<command>...]
    Show help.

  price-encryption [<flags>] <ekey> <ikey> <price>
    Encrypt a price using a given algorithm.

  price-decryption [<flags>] <ekey> <ikey> <price>
    Decrypt a price using a given algorithm.

  hex-key-generator [<keylength>]
    Generate a random hex key of the specified length (by default: 64)
```
#### Decrypt only
The command line below will try to decrypt the encrypted price passed in command line.
```
go run mmle.go price-decryption [YOUR_EKEY] [YOUR_IKEY] [ENCRYPTED_PRICE]
```
#### Encrypt only
The command line below will try to encrypt the clear price passed in command line.
```
go run mmle.go price-encryption [YOUR_EKEY] [YOUR_IKEY] [CLEAR_PRICE]
```
### Verbose mode
The verbose mode describle the process step by steps.
```
go run mmle.go price-encryption [YOUR_EKEY] [YOUR_IKEY] [CLEAR_PRICE] --verbose
Verbose activated
Keys decoding mode :  utf-8
Encryption key :  6356770B3C111C07F778AFD69F16643E9110090FD4C479D91181EED2523788F1
Encryption key (bytes) :  [99 86 119 11 60 17 28 7 247 120 175 214 159 22 100 62 145 16 9 15 212 196 121 217 17 129 238 210 82 55 136 241]
Integrity key :  3588BF6D387E8AEAD4EEC66798255369AF47BFD48B056E8934CEFEF3609C469E
Integrity key (bytes) :  [53 136 191 109 56 126 138 234 212 238 198 103 152 37 83 105 175 71 191 212 139 5 110 137 52 206 254 243 96 156 70 158]
ERROR: logging before flag.Parse: I0329 16:58:01.062251   19950 helpers.go:109] Micro price bytes: [0 0 0 0 0 22 90 168]
Seed :
Initialization vector :  [212 29 140 217 143 0 178 4 233 128 9 152 236 248 66 126]
// pad = hmac(e_key, iv), first 8 bytes
Pad :  [125 170 63 123 246 12 99 173]
// enc_data = pad <xor> data
Encoded price bytes :  [125 170 63 123 246 26 57 5]
// signature = hmac(i_key, data || iv), first 4 bytes
Signature :  [74 124 216 232]
Encrypted price:  1B2M2Y8AsgTpgAmY7PhCfn2qP3v2GjkFSnzY6A==
```

## Hexa keys generation
### Usage examples
The command line below will generate an hexa key of the desired length (by default, 64 bytes).
```
go run mmle.go hex-key-generator
go run mmle.go hex-key-generator --keylength=32
```

## Base64 to Base64 web safe and Base64 web safe to Base64
### Usage examples
TODO

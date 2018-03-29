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

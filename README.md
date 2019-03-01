## Simple Substitution Encryption

### Installation

`go get github.com/eknowlton/substenc`

Example usage:

Decrypt:

```cat plain-text.txt | ./substenc```


Decrypt:

```cat cipher-text.txt | ./substenc -decrypt```

### TODO:
- Add optional iterations
- Add optional custom alphabet

### Disclaimer

I created this completely out of curiosity and wanted to mess around with golang. This is not a secure encryption algorithm and should not be used as so.

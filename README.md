# go-base64

Simple CLI tool that encodes stdin to base64 and vice-versa.

## Usage
Encoding:
```bash
$ echo "lorem ipsum" | ./go-base64
bG9yZW0gaXBzdW0=
```  

Decoding:
```bash
$ echo "bG9yZW0gaXBzdW0=" | ./go-base64 -d
lorem ipsum
```
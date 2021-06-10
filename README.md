# X

Simply converts from a string to hex or binary
Can do the opposite if asked nicely

## Usage

### Hex

```sh
x Ohai
```
will return `4F686169`

```sh
x -r 4F686169
```
will return `Ohai`

### Bin

```sh
x -b Ohai
```
will return `01001111011010000110000101101001`

```sh
x -b -r 01001111011010000110000101101001
```
will return `Ohai`

# uritool

small tool to help with uri specific elements while working on a command line.

## Usage

### Query

Escapes the given first argument to a valid query paramter value:

```sh
uritool query encode --no-newline "hello / world!%%%"
# > hello+%2F+world%25%25
```

Unescapes the given first argument to the original value:

```sh
uritool query decode "hello+%2F+world%25%25"
# > hello / world!%%%
```

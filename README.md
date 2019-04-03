# uritool

small tool to help with uri specific elements while working on a command line.

## Usage

### Query

Escapes the given first argument to a valid query paramter value:

```sh
uirtool query encode --no-newline "hello / world!%%%"
```

Unescapes the given first argument to the original value:

```sh
uirtool query decode "hello+%2F+world%25%25"
```

# uritool

Tool that helps working with URIs on the command line, processing parts or preparing information.

## Usage

### Query commands

Encodes the given value to a valid query parameter value:

```sh
uritool query escape --no-newline "hello / world!%%"

# > hello+%2F+world%25%25
```

Decodes the given escaped query value:

```sh
uritool query unescape "hello+%2F+world%25%25"

# > hello / world!%%
```

### Path commands

Escape the given value to a valid escaped path value:

```sh
uritool path encode "hello world"

# > hello%20world
``` 

Unescape the given escaped path value:

```sh
uritool path decode "hello%20world"

# > hello world
``` 

### Parse commands

Parse a given URI and return all information as JSON:

```sh
uritool parse uri "https://my:pass@the.example.com:8080/what/ ever?this=is&this=isnot#doing"

# > {
# >   "Scheme": "https",
# >   "Opaque": "",
# >   "Username": "my",
# >   "Password": "pass",
# >   "PasswordIsGiven": true,
# >   "Host": "the.example.com:8080",
# >   "Hostname": "the.example.com",
# >   "Port": 8080,
# >   "Path": "/what/ ever",
# >   "PathEscaped": "/what/%20ever",
# >   "RawQuery": "this=is\u0026this=isnot",
# >   "Fragment": "doing",
# >   "Query": {
# >     "this": [
# >       "is",
# >       "isnot"
# >     ]
# >   }
# > }
```

But sometimes you just want a specific part or combination of it, so use can use the [go template](https://golang.org/pkg/text/template/) language:

```sh
uritool parse uri --format="Welcome {{.Username}} from {{.Hostname}}" "https://adrian:something@the.example.com:8080/?uh=oh"

# > Welcome adrian from the.example.com

uritool parse uri --format="Second entry is {{index .Query.things 1}}" "https://adrian:something@the.example.com:8080/?things=one&things=two"

# > Second entry is two
```

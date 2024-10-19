# goj2

CLI tool for Jinja2 template rendering based on [gonja](https://github.com/noirbizarre/gonja).

## Installation

To install the CLI tool, you can build it from source or download pre-built binaries for your operating system.

```
go install github.com/madest92/goj2
```

## Usage

```bash
Usage:
  goj2 [flags]

Flags:
  -f, --from string   Input template file
  -h, --help          help for goj2
  -t, --to string     Output file
  -v, --vars strings  Variables file(s) in YAML format
```

Examples:
```bash
# cat examples/hello.j2
Hello {{ name }}!
# cat examples/vars-hello.yml
name: World

# goj2 --vars examples/vars-hello.yml --from examples/hello.j2
Hello World!
```

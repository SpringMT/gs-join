# gs-join
Google Spreadsheet like join.
https://support.google.com/docs/answer/3094077

# Description

```
Example1 echo "test\ntest\ntest | gs-join -d ","
test,test,test

Example2 $ gs-join -d "," file1 file2
```

## Installation

Executable binaries are available at [releases](https://github.com/SpringMT/gs-join/releases).

For example, for linux x86_64, 

```bash
$ wget https://github.com/SpringMT/gs-join/releases/download/v0.0.1/gs-join_linux_amd64 -O gs-join 
$ chmod a+x gs-join 
```

If you have the go runtime installed, you may use go get. 

```bash
$ go get github.com/SpringMT/gs-join
```

## Usage

```
% gs-join -h
NAME:
   gs-join - Google Spreadsheet like join

USAGE:
   gs-join [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --delimiter value, -d value  Use character char as a field delimiter
   --help, -h                   show help
   --version, -v                print the version
```

## ToDo

1. write tests

## Build

To build, use go get and make

```
$ go get -d github.com/SpringMT/gs-join
$ cd $GOPATH/src/github.com/SpringMT/gs-join
$ make
```

To release binaries, I use [gox](https://github.com/mitchellh/gox) and [ghr](https://github.com/tcnksm/ghr)

```
go get github.com/mitchellh/gox
go get github.com/tcnksm/ghr

mkdir -p pkg && cd pkg && gox ../...
ghr vX.Y.Z .
```

## Contribution

1. Fork (https://github.com/SpringMT/gs-join)
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the go test ./... command and confirm that it passes
6. Run gofmt -s
7. Create new Pull Request

## Copyright

See [LICENSE](./LICENSE)

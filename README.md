# go-stdin-sleep-stdout

Build `go-stdin-sleep-stdout-M.m.P-I.x86_64.rpm`
and   `go-stdin-sleep-stdout_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

`go-stdin-sleep-stdout` is a tool 
to receive input from `stdin`, sleep,
then push the input to `stdout`.

The main use of `go-stdin-sleep-stdout` is to delay the piping of stdout / stdin.  Example:

 ```console
program-1 | go-stdin-sleep-stdout --sleep 5 | program-2
```

This example introduces a 5-second delay from the output of program-1 to the input into program-2.

### Invocation

```console
go-stdin-sleep-stdout
```

Example:

```console
cat test.txt | go-stdin-sleep-stdout --sleep 5
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-stdin-sleep-stdout"
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-stdin-sleep-stdout.git
```

#### Download dependencies

```console
cd ${REPOSITORY_DIR}
make dependencies
```

### Build

#### Local build

```console
cd ${REPOSITORY_DIR}
make build-local
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${REPOSITORY_DIR}
make build
```

The results will be in the `.../target` directory.

### Test

```console
cd ${REPOSITORY_DIR}
make test-local
```

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh go-stdin-sleep-stdout-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh go-stdin-sleep-stdout-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i go-stdin-sleep-stdout_M.m.P-I_amd64.deb
```

### Cleanup

```console
cd ${REPOSITORY_DIR}
make clean
```

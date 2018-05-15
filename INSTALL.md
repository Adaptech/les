# Installation

## Docker

Build the les validation tool image:

```(cd cmd/les ; docker build . -t les)```

Build the les-node image:

```(cd cmd/les-node ; docker build . -t les-node)```

## Linux AMD64

Install the 'les' validation tool:

```sudo curl -L https://github.com/Adaptech/les/releases/download/release-0.10.5-test/les-linux-amd64 -o /usr/local/bin/les && sudo chmod +x /usr/local/bin/les```

Install 'les-node':

```sudo curl -L https://github.com/Adaptech/les/releases/download/les-node/0.10.5-test/les-node-linux-amd64 -o /usr/local/bin/les-node && sudo chmod +x /usr/local/bin/les-node```

## Darwin AMD64

Install the 'les' validation tool:

```sudo curl -L https://github.com/Adaptech/les/releases/download/release-0.10.5-test/les-darwin-amd64 -o /usr/local/bin/les && sudo chmod +x /usr/local/bin/les```

Install 'les-node':

```sudo curl -L https://github.com/Adaptech/les/releases/download/les-node/0.10.5-test/les-node-darwin-amd64 -o /usr/local/bin/les-node && sudo chmod +x /usr/local/bin/les-node```

## Windows AMD64

* [Download](https://github.com/Adaptech/les/releases/download/release-0.10.5-test/les-windows-amd64.exe) the 'les' validation tool .exe

* [Download](https://github.com/Adaptech/les/releases/download/release-0.10.5-test/les-node-windows-amd64.exe) 'les-node' .exe


# hw

A webapp for Hans Wehr's Dictionary of Modern Written Arabic.

## Intstall

### complile

```bash
go install -ldflags "-s -w" github.com/wizsk/hw@latest
hw help
```

### or see releases

```bash
# linux
cd /tmp
wget "https://github.com/wizsk/hw/releases/latest/download/hw_Linux_$(uname -m).tar.gz"
tar xf "hw_Linux_$(uname -m).tar.gz"
sudo mv hw /usr/local/bin/ # or mv hw ~/.local/bin/
```

## Usages

```
>> hw help
hw: [port] [COMMANDS...]
PORT:
        Just the port number. (default: 8001)

COMMANDS:
        nobrowser, nb
                don't open browser
        version
                print version number
```

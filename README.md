# hw

A webapp for Hans Wehr's Dictionary of Modern Written Arabic.

## Intstall

### complile

```bash
go install -ldflags "-s -w" github.com/wizsk/hw@latest
hw help
```

### linux

```bash
# linux
cd /tmp
wget "https://github.com/wizsk/hw/releases/latest/download/hw_Linux_$(uname -m).tar.gz"
tar xf "hw_Linux_$(uname -m).tar.gz"
sudo mv hw /usr/local/bin/ # or mv hw ~/.local/bin/
```

### Windows

Open an `Administrator PowerShell` prompt and paste the following command

Go to Windows Search, type `PowerShell`, then right-click on the PowerShell app in the search results or click the small arrow (>) next to it, and select Run as Administrator.


```ps1
irm https://raw.githubusercontent.com/wizsk/hw/refs/heads/main/install.ps1 | iex
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

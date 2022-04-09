# grvemu

RISC-V emulator for CLI written in Go

## Install

```bash
go install github.com/kinpoko/grvemu@latest
```

## Usage

```bash
grvemu -h
RISC-V emulator for cli written in Go

Usage:
  grvemu [binary file] [flags]

Flags:
  -d, --debug         debug mode
  -e, --eof uint32    end of binary
  -h, --help          help for grvemu
  -i, --inst string   instruction (default "rv32i")
  -t, --test          display global pointer
```
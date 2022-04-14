# grvemu

RISC-V emulator for CLI written in Go

For now, grvemu supports only rv32i.

grvemu can pass some [riscv-tests](https://github.com/riscv-software-src/riscv-tests) and run c program. If you want to try them, you can use [riscv-tools-and-tests-docker-for-grvemu](https://github.com/kinpoko/riscv-tools-and-tests-docker-for-grvemu).

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
  -a, --arch string   architecture (default "rv32i")
  -d, --debug         debug mode
  -e, --eof uint32    end of binary
  -h, --help          help for grvemu
  -t, --test          display global pointer
```

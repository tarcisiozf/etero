# ETERO

A tiny EVM (Ethereum Virtual Machine) build from scratch for study purposes.

This is a golang implementation of the tutorial [Building an EVM from scratch](https://www.notion.so/Building-an-EVM-from-scratch-part-1-the-execution-context-c28ebb4200c94f6fb75948a5feffc686) by @karmacoma-eth (originally implemented in python).

Work is still in progress...

### Examples:

Run a simple example using stack and memory opcodes:

```shell
go run example/main.go 600660070260005360016000f3
```

### TODO

- [ ] improve debuggability of code execution
- [ ] add error wrappers
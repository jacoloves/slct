# SLCT
It is `slct` for `select`.

Simple command line selector.

## Demo
```sh
ls -a | slct
```

[![ls-aslct](https://asciinema.org/a/AIV5gUgisQSdHU0TwroGUVAnK.svg)](https://asciinema.org/a/AIV5gUgisQSdHU0TwroGUVAnK)

```sh
cd $(ghq list --full-path | slct)
```

[![cd-ghq-list](https://asciinema.org/a/xFTzNG7oktfCuT9Idbl3xa2oA.svg)](https://asciinema.org/a/xFTzNG7oktfCuT9Idbl3xa2oA)

## Usage
Just lik [cho](https://github.com/mattn/cho)

Simple command line selector.

## Installation
```sh
go get github.com/jacoloves/slct
```

## Keys
|Key   |Behavior            |
|------|--------------------|
|j or ↓|Next                |
|k or ↑|Previous            |
|Enter |Decide              |
|q     |Exit                |

## In the future
- [ ] add incremental searcha
- [ ] add multi select
- [ ] add config yaml


## LICENSE
Distributed under MIT License. See LICENSE.
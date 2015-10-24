## yosage v0.2.0

![eupho.gif](https://raw.githubusercontent.com/hico-horiuchi/yosage/master/eupho.gif)

`yosage` (良さげ) is LGTM gif generator by Golang.  
This is stand-alone (not require ImageMagick, include [lgtm.png](https://github.com/hico-horiuchi/yosage/blob/master/lgtm.png)).  
I get ideas from [8398a7/lgtm_creater](https://github.com/8398a7/lgtm_creater) and [syohex/speedline](https://github.com/syohex/speedline).

#### Requirements

  - [Golang](https://golang.org/) >= 1

#### Installation

```sh
$ git clone git://github.com/hico-horiuchi/yosage.git
$ cd yosage
$ make gom link
$ make build
$ bin/yosage -i inpunt.gif -o output.gif
```

#### Usage

    Stabd-alone LGTM gif generator by Golang
    https://github.com/hico-horiuchi/yosage
    
    Usage:
      yosage [flags]
    
    Flags:
      -i, --input="input.gif": input gif file path
      -l, --lgtm="": LGTM text png file path
      -o, --output="output.gif": output gif file path
      -v, --version[=false]: print and check the version

#### License

yosage is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/yosage/master/LICENSE).

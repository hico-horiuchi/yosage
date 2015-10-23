## yosage v0.1.0

![eupho.gif](https://raw.githubusercontent.com/hico-horiuchi/yosage/master/eupho.gif)

`yosage` (良さげ) is LGTM gif generator by Golang.  
This is stand-alone (include ImageMagick and [lgtm.png](https://github.com/hico-horiuchi/yosage/blob/master/lgtm.png)) impremented with [gographics/imagick](https://github.com/gographics/imagick).  
I get ideas from [8398a7/lgtm_creater](https://github.com/8398a7/lgtm_creater) and [syohex/speedline](https://github.com/syohex/speedline).

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [ImageMagick](http://www.imagemagick.org/) <= 6.8.8

#### Installation

```sh
$ brew install imagemagick # for OS X
$ sudo apt-get install libmagickwand-dev # for Ubuntu / Debian
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
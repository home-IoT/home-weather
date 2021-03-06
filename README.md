# Home Weather CLI (Jupiter CLI)

 `master`   | `develop`     |
| ---------- | --------------|
| [![Build Status](https://travis-ci.org/home-IoT/home-weather.svg?branch=master)](https://travis-ci.org/home-IoT/home-weather) | [![Build Status](https://travis-ci.org/home-IoT/home-weather.svg?branch=develop)](https://travis-ci.org/home-IoT/home-weather) |

**In Progress**: This repository is a work in progress. Should be completed in a week from now. ;)

A simple CLI to query the current status of weather sensors around the house based on the 
[Jupiter](https://github.com/home-IoT/jupiter) gateway service.

### Build 

Make sure you that
* you have `dep` installed. Visit https://github.com/golang/dep 
* your `GOPATH` and `GOROOT` environments are set properly.

#### Makefile
There is a [`Makefile`](Makefile) provided that offers a number of targets for preparing, building and running the CLI. To build the binary, run:
```
make clean go-dep go-build
```

## License
The code is published under an [MIT license](LICENSE.md). 

## Contributions
Please report issues or feature requests using Github issues. Code contributions can be done using pull requests. 

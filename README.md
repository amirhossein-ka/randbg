# Randbg, Awesome background changer for wm !

## Note
 
this project depends on `feh` to set backgrounds (at least for now),
so this app won't work on most DEs and windows.

## Prerequisites
make sure you have these installed:
* `feh` to set backgrounds
* a working `go` installation, see [here](https://go.dev/doc/install) to install go
* a wm that uses Xorg (sway and other wayland WMs not supported due to wayland)


## Install
1. First method:
   1. clone this repo ```git clone https://github.com/amirhossein-ka/randbg.git```
   2. make sure `$GOPATH/bin/` is available in `$PATH`
   3. Run `make && make install`
   4. Run `randbgd` in background
   5. control daemon with `randbgctl`
2. second method:
   1. make sure `$GOPATH/bin/` is available in `$PATH` 
   2. run ```go install github.com/amirhossein-ka/randbg/cmd/daemon@latest``` to install daemon
   3. run ```go install github.com/amirhossein-ka/randbg/cmd/controller@latest``` to install controller
   4. run `randbgd` in background
   5. control daemon with `randbgctl`
   

## Todo
- [ ] use unix sockets 
- [ ] use pid files to store PID of `randbgd` 
- [ ] add support for `nitrogen`
- [ ] use urls to fetch pics
- [ ] rename folders and change structure of project

If you have issues, inform me with issues or make a pull request 

### _**enjoy !**_


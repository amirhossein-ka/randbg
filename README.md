# Randbg, Awesome background changer for wm !

## Note
 
this project depends on `feh` to set backgrounds (at least for now),
so this app won't work on most DEs and windows.

another reason that this app will not work in windows is that it heavily using unix signals so this app can't be used on Windows
## Prerequisites
make sure you have these installed:
* `feh` to set backgrounds
* a working `go` installation, see [here](https://go.dev/doc/install) to install go
* a wm that uses Xorg (sway and other wayland WMs not supported due to wayland)


## Installation
1. make sure have a working go installation
2. run `go install github.com/amirhossein-ka/randbg@latest`
3. add `$GOPATH/bin` to `$PATH`
4. run `randbg daemon -config /path/to/config.{yml.json}`
   

## Config example
you can put your config file in `$HOME/.config/randbg/config.yaml` to avoid specifying config path with -config

json configuration:
```json
{
  "daemon_config": {
    "image_directory": "/absolute/path/to/images/dir/",
    "interval": "10h20m30s"
  }
}
```

yaml configuration:
```yaml
daemon_config:
  image_directory: "/absolute/path/to/images/dir/",
  interval: 10h20m30s
```

## Todo
- [x] use pid files to store PID of `randbgd` 
- [x] rename folders and change structure of project
- [ ] add support for `nitrogen`
- [ ] use urls to fetch pics
- [ ] (Maybeee) add support for DEs

If you have issues, inform me with issues or make a pull request 

### _**enjoy !**_


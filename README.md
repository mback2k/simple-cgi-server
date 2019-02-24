simple-cgi-server
=================
This Go program is a simple CGI server to host legacy CGI scripts.
This server is intended to be running behind a reverse proxy or
loadbalancer, like [traefik](https://traefik.io/).

[![Build Status](https://travis-ci.org/mback2k/simple-cgi-server.svg?branch=master)](https://travis-ci.org/mback2k/simple-cgi-server)
[![GoDoc](https://godoc.org/github.com/mback2k/simple-cgi-server?status.svg)](https://godoc.org/github.com/mback2k/simple-cgi-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/mback2k/simple-cgi-server)](https://goreportcard.com/report/github.com/mback2k/simple-cgi-server)

Dependencies
------------
The following awesome Go libraries are dependencies:

- https://github.com/sirupsen/logrus
- https://github.com/spf13/viper

Installation
------------
You basically have two options to install this Go program package:

1. If you have Go installed and configured on your PATH, just do the following go get inside your GOPATH to get the latest version:

```
go get -u github.com/mback2k/simple-cgi-server
```

2. If you do not have Go installed and just want to use a released binary,
then you can just go ahead and download a pre-compiled Linux amd64 binary from the [Github releases](https://github.com/mback2k/simple-cgi-server/releases).

Finally put the simple-cgi-server binary onto your PATH and make sure it is executable.

Configuration
-------------
The following YAML file is an example configuration to serve Python scripts:

```
Address: "localhost:8080"
HandlerList:
  - FileExt: .py
    Handler: python
Logging:
  Level: info
```

Save this file in one of the following locations and run `./simple-cgi-server`:

- /etc/simple-cgi-server/simple-cgi-server.yaml
- $HOME/.simple-cgi-server.yaml
- $PWD/simple-cgi-server.yaml

License
-------
Copyright (C) 2019  Marc Hoersken <info@marc-hoersken.de>

This software is licensed as described in the file LICENSE, which
you should have received as part of this software distribution.

All trademarks are the property of their respective owners.

/*
	simple-cgi-server - A simple CGI server to host legacy CGI scripts.
	Copyright (C) 2019  Marc Hoersken <info@marc-hoersken.de>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"github.com/mback2k/simple-cgi-server/cgiserver"
	"github.com/spf13/viper"
)

type configAlias struct {
	Source string
	Target string
}

type configHandler struct {
	FileExt string
	Handler string
}

type configLogging struct {
	Level string
}

type config struct {
	Address        string
	DocumentRoot   string
	InheritEnv     []string
	DirectoryIndex []string
	DefaultHandler string
	AliasList      []*configAlias
	HandlerList    []*configHandler
	Logging        *configLogging
}

func loadConfig(s *cgiserver.Server) (*config, error) {
	vpr := viper.GetViper()
	vpr.SetDefault("Address", s.Address)
	vpr.SetDefault("DocumentRoot", s.DocumentRoot)
	vpr.SetDefault("InheritEnv", s.InheritEnv)
	vpr.SetDefault("DirectoryIndex", s.DirectoryIndex)
	vpr.SetDefault("DefaultHandler", s.DefaultHandler)
	vpr.SetConfigName("simple-cgi-server")
	vpr.AddConfigPath("/etc/simple-cgi-server/")
	vpr.AddConfigPath("$HOME/.simple-cgi-server")
	vpr.AddConfigPath(".")
	err := vpr.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg config
	err = vpr.UnmarshalExact(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

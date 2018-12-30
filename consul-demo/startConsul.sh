#!/bin/bash

consul agent -dev -ui -server -node=consul-dev -client=127.0.0.1 -dns-port=5533 -domain=consul.crazywolf.com

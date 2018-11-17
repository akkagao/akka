#!/bin/bash
rm -rf *
cobra init
cobra add dev
cobra add test
cobra add prod

mkdir conf resources handler doc route common
touch resources/conf-dev.yaml resources/conf-test.yaml resources/conf-prod.yaml

#!/bin/bash

echo "##teamcity[blockOpened name='keys' description='set up credentials']"
make up-login
echo "##teamcity[blockClosed name='keys']"

make publish REGISTRY_ORGS=xpkg.upbound.io/yandexcloud XPKG_REG_ORGS=xpkg.upbound.io/yandexcloud

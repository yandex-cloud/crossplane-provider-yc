#!/bin/bash

echo "##teamcity[blockOpened name='keys' description='set up credentials']"
make up-login
echo "##teamcity[blockClosed name='keys']"

if [[ -z "$VERSION" ]]; then
  echo "VERSION is not set"
  exit 1
fi

echo "Publishing version $VERSION"
make publish REGISTRY_ORGS=xpkg.upbound.io/yandexcloud XPKG_REG_ORGS=xpkg.upbound.io/yandexcloud

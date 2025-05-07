#!/bin/bash

echo "##teamcity[blockOpened name='keys' description='set up credentials']"
echo "$UP_TOKEN" | docker login xpkg.upbound.io -u dbb80626-2a02-4eb6-87d7-b6414272916e --password-stdin
echo "##teamcity[blockClosed name='keys']"

if [[ -z "$VERSION" ]]; then
  echo "VERSION is not set"
  exit 1
fi

echo "Publishing version $VERSION"
make publish REGISTRY_ORGS=xpkg.upbound.io/yandexcloud XPKG_REG_ORGS=xpkg.upbound.io/yandexcloud

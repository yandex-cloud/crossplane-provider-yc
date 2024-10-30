#!/bin/bash

echo "##teamcity[blockOpened name='keys' description='set up credentials']"

yc config profile create robot
yc config set service-account-key $SA_KEY_FILE


yc container registry configure-docker

echo "##teamcity[blockClosed name='keys']"

make publish

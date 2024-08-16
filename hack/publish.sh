#!/bin/bash

[ -n "${SA_KEY_JSON}" ] || { echo SA_KEY_JSON env var not set, can not proceed; exit 1; }
[ -n "${SECRET_ID}" ] || { echo SECRET_ID env var not set, can not proceed; exit 1; }
[ -n "${SECRET_KEY}" ] || { echo SECRET_KEY env var not set, can not proceed; exit 1; }
[ -n "${CLOUD_ID}" ] || { echo CLOUD_ID env var not set, can not proceed; exit 1; }
[ -n "${FOLDER_ID}" ] || { echo FOLDER_ID env var not set, can not proceed; exit 1; }

echo "##teamcity[blockOpened name='keys' description='set up credentials']"

echo $SA_KEY_JSON > sa_key.json

yc config profile create robot
yc config set service-account-key sa_key.json

yc container registry configure-docker

yc config set folder-id ${FOLDER_ID}
yc config set cloud-id ${CLOUD_ID}

UP_TOKEN=$(yc lockbox payload get --id ${SECRET_ID} --key ${SECRET_KEY}) make up-login

echo "##teamcity[blockClosed name='keys']"

make publish

#!/bin/bash

exitcode=0

function delete_all {
    for id in $(yc "${@}" list --folder-id "${FOLDER_ID}" --format json | jq -r '(map({ id: .id}) | .[].id)'); do
        echo Deleting "${@}" $id...
        yc "${@}" delete $id || exitcode=1
    done
}

# this needs to be first, so that Crossplane doesn't attempt to recreate resources as we delete them
delete_all managed-kubernetes cluster

delete_all compute instance
delete_all kms symmetric-key
delete_all dns zone
delete_all datatransfer transfer
delete_all datatransfer endpoint
delete_all managed-redis cluster
delete_all managed-mongodb cluster
delete_all managed-mysql cluster
delete_all managed-postgresql cluster
delete_all container image
delete_all container registry
delete_all managed-kubernetes cluster
delete_all iam service-account
delete_all ydb db
delete_all managed-kafka cluster
delete_all vpc addr
delete_all vpc subnet
delete_all vpc net

exit $exitcode
#!/bin/bash

function delete_all {
    for id in $(yc "${@}" list --folder-id "${FOLDER_ID}" --format json | jq -r '(map({ id: .id}) | .[].id)'); do
        echo Deleting "${@}" $id...
        yc "${@}" delete $id
    done
}

delete_all container image
delete_all container registry
delete_all managed-kubernetes cluster
delete_all iam service-account
delete_all ydb db
delete_all vpc addr
delete_all vpc subnet
delete_all vpc net
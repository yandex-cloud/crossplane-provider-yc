#!/bin/sh

echo $pwd
cd .work/terraform-provider-yandex/website/docs/r
for file in $(ls ".")
do
    numRows=$(cat $file | grep 'subcategory' | wc|awk '{ print $1 }')
    if [ $numRows -gt 0 ]; then
        continue
    fi

    newFile="C${file}"

    if [ ! -f $newFile ]; then
        cp $file $newFile
        sed '1a\subcategory: "unknown"' $file > $newFile
        rm $file
    fi
done
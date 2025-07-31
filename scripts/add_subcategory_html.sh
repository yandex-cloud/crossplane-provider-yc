#!/bin/sh

echo $pwd

DIRECTORY=".work/terraform-provider-yandex/docs/resources"

if [ ! -d "$DIRECTORY" ]; then
  exit 2
fi

cd $DIRECTORY
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
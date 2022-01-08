#!/bin/bash
function read_dir(){
    for file in `ls -a $1`
    do
        if [ -d $1"/"$file ]
        then
            if [[ $file != '.' && $file != '..' && $file != '.git' ]]
            then
                read_dir $1"/"$file $1
            fi
        else
          if [[ "${file##*.}"x == 'proto'x ]]
          then
              protoc --proto_path=$(dirname $2) --go_out=$2
          fi
        fi
    done
}

root_dir=$(cd `dirname $0`; pwd)
read_dir $root_dir $root_dir

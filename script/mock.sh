#! /bin/bash
# shellcheck disable=SC1068
# shellcheck disable=SC1066
# shellcheck disable=SC2006
generated=`pwd`"/test/mock"
function read_dir(){
    # shellcheck disable=SC2045
    for file in `ls $1`
    do
        if [ -d $1"/"$file ] && [ $file != "test" ]
        then
            read_dir $1"/"$file
        else
            if [ "${file##*.}"x = "go"x ] && [[ $file != *_test.go ]] && [[ $file != options.go ]]
            then
                inputFile=$1"/"$file
                outputFile=$1"/mock_"$file
                outputFile=${outputFile/"internal"/"internal-mock"}
                destination=$generated${outputFile#*mgcache}
                echo $destination
                mockgen -source=$inputFile -destination=$destination
            fi
        fi
    done
}
read_dir `pwd`
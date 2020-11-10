#!/bin/bash
# Reloads the server on file change
# Works within docker containers. Uses polling since inotify events aren't
# propagated to docker containers from the host machine.

# kill server on CTRL+C
trap 'kill $pid; exit' SIGINT
app_name="zi-wei-dou-shu-gin"
app_path="/go/src/github.com/zi-wei-dou-shu-gin"

function build_and_execute {
    echo -e "\e[32mBuilding...\e[0m"
    go install $app_path
    ps aux | grep $app_name | awk '{print $2; exit}' | xargs kill -9
    if [[ $? -eq 0 ]]; then
        echo -e "\e[32mExecuting...\e[0m"
        $app_name &
        pid=$!
    fi
}

prev_checksum=""
while true; do
    checksum=$(find . -type f -exec md5sum {} \;)
    if [[ $prev_checksum != $checksum ]] ; then
        if [[ ! -z $pid ]]; then
            if ps -p $pid > /dev/null; then
                echo -e "\e[32mReloading...\e[0m"
                kill $pid
                sleep 0.5
            fi
        fi
        build_and_execute
        prev_checksum=$checksum
    fi
    sleep 2
done
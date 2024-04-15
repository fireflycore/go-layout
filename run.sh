#!/bin/bash

project_name="project_name"

# shellcheck disable=SC2009
pid=$(ps -ef | grep "/opt/app/lhdht/microservice/${project_name}/main" | grep -v grep | awk '{print $2}')
if [ -n "$pid" ]
then
   kill -9 "$pid"
fi

sleep 5
nohup "/opt/app/lhdht/microservice/${project_name}/main" > ./nohup.log 2>&1 &
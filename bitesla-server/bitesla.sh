#!/usr/bin/env bash

MODE=$1
STATE=$2

# 项目根路径
ROOT_PATH=$(pwd)

# docker和项目文件映射地址
RUN_PATH="/data/bitesla-server"

# service image
IMAGE_SERVICE_API=bitesla-service-api
IMAGE_SERVICE_USER=bitesla-service-user
IMAGE_SERVICE_EXCHANGE=bitesla-service-exchange
IMAGE_SERVICE_STRATEGY=bitesla-service-strategy
IMAGE_SERVICE_TRADER=bitesla-service-trader

## docker操作
function docker_state(){
    case ${STATE} in
        "dep")
            docker_dep ;;
        "start")
            docker_start ;;
        "push")
            docker_push ;;
        *)
            printHelp
            exit 1
    esac
}

# 依赖到到一些docker环境
function docker_dep() {
    # 该步骤不要启动aichain service
    # 文件的映射地址直接指向了运行地址，RUN_PATH
    RUN_PATH=${RUN_PATH} docker-compose -f ${RUN_PATH}/docker-compose.yml up -d bitesla-consul bitesla-mysql bitesla-redis bitesla-nsqlookupd bitesla-nsqd bitesla-nsqadmin
}

# 启动所有项目
function docker_start() {
    RUN_PATH=${RUN_PATH} docker-compose -f ${RUN_PATH}/docker-compose.yml up -d
}

function release_state(){
    if [[ ${STATE} == "" ]]; then
        printHelp
        exit 1
    elif [[ ${STATE} == "all" ]]; then
        release_all
    else
        release_one ${STATE}
    fi
}

# 释放所有docker环境,数据库等基建镜像不会被删除
function release_all() {
    docker-compose stop
    docker-compose rm -f

    docker rmi -f wujason/${IMAGE_SERVICE_API}
    docker rmi -f wujason/${IMAGE_SERVICE_USER}
    docker rmi -f wujason/${IMAGE_SERVICE_EXCHANGE}
    docker rmi -f wujason/${IMAGE_SERVICE_STRATEGY}
    docker rmi -f wujason/${IMAGE_SERVICE_TRADER}

    # 删除为none的镜像
    docker images|grep none|awk '{print $3}'|xargs docker rmi
}

function release_one(){
    name=$1
    RUN_PATH=${RUN_PATH} docker-compose stop ${name}
    docker-compose rm -f ${name}
    docker rmi -f ${name}
    echo $1
}

function printHelp() {
    echo "./bitesla.sh docker [+操作码]：用于测试环境和生产环境，是手动操作"
    echo "          [操作码]"
    echo "               dep：启动项目依赖的镜像(mysql redis nsq等)"
    echo "               start：启动项目"
    echo "./bitesla.sh release [+操作码]：用于释放项目和其余容器"
    echo "          [操作码]"
    echo "               all：释放项目所有内容，包括各种容器"
    echo "               指定容器名：释放指定容器，主要是用来释放项目所在的容器"
    echo "其余操作将触发此说明"
}

#启动模式
case ${MODE} in
    "docker")
        docker_state ;;
    "release")
        release_state ;;
    *)
        printHelp
        exit 1
esac
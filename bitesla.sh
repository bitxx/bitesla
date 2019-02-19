#!/usr/bin/env bash
# !/usr/bin/env bash

MODE=$1
STATE=$2

# 项目名
PROJECT_NAME=bitesla

# 项目根路径
ROOT_PATH=$(pwd)

# docker和项目文件映射地址
RUN_PATH=${ROOT_PATH}

# service image
IMAGE_SERVICE_API=bitesla-service-api
IMAGE_SERVICE_USER=bitesla-service-user
IMAGE_SERVICE_EXCHANGE=bitesla-service-exchange
IMAGE_SERVICE_STRATEGY=bitesla-service-strategy
IMAGE_SERVICE_TRADER=bitesla-service-trader

## 本地操作
function local_state(){
    case ${STATE} in
        "proto")
            local_proto ;;
        "doc")
            local_doc ;;
        *)
            printHelp
            exit 1

    esac
}

# 生成服务,用不到了，留作记录
#function local_service(){
#    echo "not need!"
    # service-strategy
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-strategy --namespace=${PROJECT_NAME} --alias algorithm --type=srv

    # service-exchange
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-exchange --namespace=${PROJECT_NAME} --alias exchange --type=srv

    # service-trader
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-trader --namespace=${PROJECT_NAME} --alias trader --type=srv

    # service-user
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-user --namespace=${PROJECT_NAME} --alias=user --type=srv
#}

# 生成接口文档
function local_doc(){
    # 删掉旧的doc文档
    rm -rf ${ROOT_PATH}/service/service-api/docs/*
    # 防止异常
    cd  ${ROOT_PATH}/service/service-api/
    swag init
}

# proto生成
function local_proto(){
    TMP_PATH=${GOPATH}/src/github.com/jason-wj/${PROJECT_NAME}/service
    TMP_S1=${TMP_PATH}/service-strategy/
    TMP_S2=${TMP_PATH}/service-trader/
    TMP_S3=${TMP_PATH}/service-exchange/
    TMP_S4=${TMP_PATH}/service-user/
    protoc --proto_path=${TMP_S1} --micro_out=${TMP_S1} --go_out=${TMP_S1} proto/strategy.proto
    protoc --proto_path=${TMP_S2} --micro_out=${TMP_S2} --go_out=${TMP_S2} proto/trader.proto
    protoc --proto_path=${TMP_S3} --micro_out=${TMP_S3} --go_out=${TMP_S3} proto/exchange.proto
    protoc --proto_path=${TMP_S4} --micro_out=${TMP_S4} --go_out=${TMP_S4} proto/user.proto
}

# 用于本地执行，但感觉不方便
#function local_start(){
#    # 配置
#    cp -rf ./bitesla-config.ini ./service/service-user/conf
#    cp -rf ./bitesla-config.ini ./service/service-api/conf
#    docker_dep
##    cd ${ROOT_PATH}/service/service-user/
##    nohup go run ${ROOT_PATH}/service/service-user/main.go &
##    echo "------------"
#    BUILD_PATH=${RUN_PATH} docker-compose -f ${ROOT_PATH}/docker-compose.yml up -d bitesla-service-user
#
#    cd ${ROOT_PATH}/service/service-api/
#    go run ${ROOT_PATH}/service/service-api/main.go &
#}

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
    # 程序配置文件的正常读取是在该目录下进行的
    cd  ${ROOT_PATH}

    # 该步骤不要启动aichain service
    # 文件的映射地址直接指向了运行地址，RUN_PATH
    RUN_PATH=${RUN_PATH} docker-compose -f ${ROOT_PATH}/docker-compose.yml up -d bitesla-consul bitesla-mysql bitesla-redis bitesla-nsqlookupd bitesla-nsqd bitesla-nsqadmin
}

# 启动所有项目
function docker_start() {
    # 将执行环境复制到不同的服务中，统一管理
    cp -rf ./bitesla-config.ini ./service/service-user/conf
    cp -rf ./bitesla-config.ini ./service/service-api/conf
    cp -rf ./bitesla-config.ini ./service/service-strategy/conf
    cp -rf ./bitesla-config.ini ./service/service-trader/conf
    cp -rf ./bitesla-config.ini ./service/service-exchange/conf

    # 开始执行
    make build
    # 再启动
    RUN_PATH=${RUN_PATH} docker-compose -f ${ROOT_PATH}/docker-compose.yml up -d
}

function docker_push(){

    #docker login --username=${DOCKER_USERNAME} --password ${DOCKER_PASSWORD}
    docker login

    docker tag ${IMAGE_SERVICE_API} wujason/${IMAGE_SERVICE_API}:latest
    docker tag ${IMAGE_SERVICE_USER} wujason/${IMAGE_SERVICE_USER}:latest
    docker tag ${IMAGE_SERVICE_EXCHANGE} wujason/${IMAGE_SERVICE_EXCHANGE}:latest
    docker tag ${IMAGE_SERVICE_STRATEGY} wujason/${IMAGE_SERVICE_STRATEGY}:latest
    docker tag ${IMAGE_SERVICE_TRADER} wujason/${IMAGE_SERVICE_TRADER}:latest

    docker push wujason/${IMAGE_SERVICE_API}:latest
    docker push wujason/${IMAGE_SERVICE_USER}:latest
    docker push wujason/${IMAGE_SERVICE_EXCHANGE}:latest
    docker push wujason/${IMAGE_SERVICE_STRATEGY}:latest
    docker push wujason/${IMAGE_SERVICE_TRADER}:latest
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
    rm -rf ${ROOT_PATH}/service/service-api/api-srv
    rm -rf ${ROOT_PATH}/service/service-user/user-srv
    rm -rf ${ROOT_PATH}/service/service-exchange/exchange-srv
    rm -rf ${ROOT_PATH}/service/service-strategy/strategy-srv
    rm -rf ${ROOT_PATH}/service/service-trader/trader-srv

    docker-compose stop
    docker-compose rm -f
    docker rmi -f ${IMAGE_SERVICE_API}
    docker rmi -f ${IMAGE_SERVICE_USER}
    docker rmi -f ${IMAGE_SERVICE_EXCHANGE}
    docker rmi -f ${IMAGE_SERVICE_STRATEGY}
    docker rmi -f ${IMAGE_SERVICE_TRADER}

    docker rmi -f wujason/${IMAGE_SERVICE_API}
    docker rmi -f wujason/${IMAGE_SERVICE_USER}
    docker rmi -f wujason/${IMAGE_SERVICE_EXCHANGE}
    docker rmi -f wujason/${IMAGE_SERVICE_STRATEGY}
    docker rmi -f wujason/${IMAGE_SERVICE_TRADER}

    # 删除为none的镜像
    docker images|grep none|awk '{print $3}'|xargs docker rmi -f
}

function release_one(){
    name=$1
    RUN_PATH=${ROOT_PATH} docker-compose stop ${name}
    docker-compose rm -f ${name}
    docker rmi -f ${name}
    echo $1
}

function printHelp() {
    echo "./aichain.sh local [+操作码]：仅能用于本地环境开发"
    echo "          [操作码]"
    echo "               doc：生成对外restful api接口文档"
    echo "               proto：各服务生成protobuf文件"
    echo "./bitesla.sh docker [+操作码]：用于测试环境和生产环境，是手动操作"
    echo "          [操作码]"
    echo "               dep：启动项目依赖的镜像(mysql redis nsq等)"
    echo "               start：启动项目"
    echo "               push：将各服务项目镜像推送等docker hub"
    echo "./bitesla.sh release [+操作码]：用于释放项目和其余容器"
    echo "          [操作码]"
    echo "               all：释放项目所有内容，包括各种容器"
    echo "               指定容器名：释放指定容器，主要是用来释放项目所在的容器"
    echo "其余操作将触发此说明"
}

#启动模式
case ${MODE} in
    "local")
        local_state ;;
    "docker")
        docker_state ;;
    "release")
        release_state ;;
    *)
        printHelp
        exit 1
esac
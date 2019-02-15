#!/usr/bin/env bash
# !/usr/bin/env bash

# 输入模式：start release
MODE=$1
STATE=$2

# 项目名
PROJECT_NAME=bitesla

# 项目根路径
ROOT_PATH=$(pwd)

# 测试、生产环境下，docker和项目文件映射地址
RUN_PATH=${ROOT_PATH}

## 本地操作
## 需要local_main执行后才能执行local_nsq
function local_state(){
    case ${STATE} in
        "service")
            local_service ;;
        "proto")
            local_proto ;;
        "doc")
            local_doc ;;
        *)
            printHelp
            exit 1

    esac
}

# 生成服务
function local_service(){
    echo "not need!"
    # service-strategy
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-strategy --namespace=${PROJECT_NAME} --alias algorithm --type=srv

    # service-exchange
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-exchange --namespace=${PROJECT_NAME} --alias exchange --type=srv

    # service-trader
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-trader --namespace=${PROJECT_NAME} --alias trader --type=srv

    # service-user
    #micro new github.com/jason-wj/${PROJECT_NAME}/service/service-user --namespace=${PROJECT_NAME} --alias=user --type=srv
}

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
        "move")
            docker_move ;;
        "start")
            docker_start ;;
        *)
            printHelp
            exit 1
    esac
}

# 依赖到到一些docker环境
function docker_dep() {
    # 先释放
    # release

    # 程序配置文件的正常读取是在该目录下进行的
    cd  ${ROOT_PATH}

    # 该步骤不要启动aichain service
    # 文件的映射地址直接指向了运行地址，RUN_PATH
    BUILD_PATH=${RUN_PATH} docker-compose -f ${ROOT_PATH}/docker-compose.yml up -d bitesla-consul bitesla-mysql bitesla-redis
}

# 发布，将编译好的go文件和配置，拷贝到发布地址中
function docker_move() {
    cp -rf ${ROOT_PATH}/server-aichain/server ${RUN_PATH}/aichain-server/
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
    BUILD_PATH=${RUN_PATH} docker-compose -f ${ROOT_PATH}/docker-compose.yml up -d
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
    docker rmi -f bitesla-service-api
    docker rmi -f bitesla-service-user
    docker rmi -f bitesla-service-exchange
    docker rmi -f bitesla-service-strategy
    docker rmi -f bitesla-service-trader
    # 删除为none的镜像
    docker images|grep none|awk '{print $3}'|xargs docker rmi
}

function release_one(){
    name=$1
    BUILD_PATH=${ROOT_PATH} docker-compose stop ${name}
    docker-compose rm -f ${name}
    docker rmi -f ${name}
    echo $1
}

function printHelp() {
    echo "./aichain.sh local [+操作码]：仅能用于本地环境开发"
    echo "          [操作码]"
    echo "               service：生成服务，切记仅第一次生成可用，之后再不可用"
    echo "               proto：生成protobuf文件"
    echo "               conf：环境配置文件"
    echo "./bitesla.sh docker [+操作码]：用于测试环境和生产环境，是手动操作"
    echo "          [操作码]"
    echo "               dep：启动项目依赖的镜像"
    echo "               move：将项目移动到运行目录"
    echo "               start：启动项目"
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
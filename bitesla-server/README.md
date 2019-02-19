# bitesla-server
此处直接运行项目

## 使用方法
进入bitesla.sh，将`RUN_PATH`设置为当前`bitesla-server`所在位置，比如：
```sh
RUN_PATH=/data/bitesla-server
```

## 启动运行
```bash
./bitesla.sh docker start
```
浏览器打开：`http://localhost:8090/swagger/index.html`

## 停止并清理镜像
```bash
./bitesla.sh release all
```
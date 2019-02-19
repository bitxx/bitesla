# Bitesla-比特斯拉
基于微服务(go-micro)的代币量化交易框架

* 本项目旨在追逐以下3个目标：
    1. 提高个人技术开发能力
    2. 深化个人对金融行业的理解
    3. 通过简单的策略编辑，实现代币自动化交易

* 愿景：  
    为下一步实现股市量化交易框架做准备（毕竟总使用第三方公司的服务，并不那么自由）

* 本项目注意事项：
    1. 暂不提供前端web展示页面，仅提供接口页面展示(swagger)
    2. 目前仅提供火币交易所的对接，其余交易所待本项目稳定后再添加
    3. 本项目还处在初期阶段，安全性、稳定性并无法绝对保证，使用者若急需，可根据需要自行完善
    4. 目前仅支持使用golang编辑策略
    5. 本项目[demo](http://106.12.77.114:8090/swagger/index.html)，其中的接口调用使用的域名：`www.bitesla.com`，因此，需要在自己操作系统中设置一下hosts来解析。

## 框架图与解析
`ps:` 本节所展示的结构图是根据本项目已实现内容绘制的。

### 全局框架
![image](/img/1.png)  
使用了go-micro微服务框架，划分了5大服务，每个服务拥有各自独立的职责。所有服务均使用一个全局的关系型数据库mysql，当前项目没有加入redis缓存。而终端用户，是通过restful api来访问的。
接着来解释下这五大服务的功能和细节。

### service-api(对外接口服务)
该服务统一对外提供restful接口通信（后续会新增socket通信），就是说，微服务中的其余每个服务都是通过该服务对终端提供接口的。
该服务集成了gin框架，通过该框架来进行web端操作。
该服务在整个框架中的定位如下图所示：
![image](/img/2.png)

### service-user（用户管理服务）
该服务主要负责项目全局账户管理，如：用户注册、用户登录、token令牌生成等。
该服务相比较要简单很多，这里就不画图介绍了。

### service-exchange（交易所管理服务）
该服务主要是用来管理各大交易所，将不同交易所的接口统一，能够提供每个交易所的k线、订单、交易等内容。
目前只集成了火币交易所，火币交易所受制于国内政策，不能直接访问，需要通过代理来访问。
整体结构如下图所示：
![image](/img/3.png)

### service-strategy（策略管理服务）
该服务很简单，就不画图了，主要是如下内容：
* 该服务主要是用来管理策略的，编写、删除、查看策略等操作。
* 注意，这是用来简单管理策略的，该服务并不提供执行策略功能。
* 目前仅仅支持使用`golang`编写策略，策略编写好后，会被存放在全局数据库mysql中

### service-trader（策略执行服务）
该服务主要是运行终端用户指定的策略，该模块比较复杂，涉及到策略执行队列，策略执行语言、编译等问题。
关于编辑策略的开发语言选择问题，我考虑了很久，首先参考的是[Samaritan](https://github.com/jason-wj/samaritan)，它是使用`JavaScript`来编辑策略的，
对应的需要[otto](https://github.com/robertkrimen/otto)来支撑，这个库有一段时间没更新了，实在有点勉强，最终`JavaScript`我选择了放弃。
我另外一个备选方案是`Python`，因为在股票量化中是非常受欢迎的（计算库很强大），我本人用过一段时间这个语言，但分精力去想办法集成一种新的语言到本项目是比较消耗时间的，因此我决定暂时放弃该方案，待后续项目稳定后再来支持该语言。
最终选择了使用golang来编写策略，这个对于自己来说，要容易很多，毕竟先实现功能才是重要的。策略的编写与应用我会在后面章节中详细讲解。
该服务的整体结构如下图所示：
![image](/img/4.png)
就是说，用户发起的所有策略执行请求，都依次加入到队列，然后依次取出，并发执行策略。
当然，这个服务仅仅是完成了上述所说的基本内容，而策略中断，策略重复执行等问题暂时还未处理，后续将会逐步解决这些问题。

## 项目部署及使用
我是在macos上开发的，windows应该也类似，可参考下述内容，最终整个项目都是在docker容器中运行。

### 环境
需要预先安装以下内容:
    1. golang 1.11.1
    2. docker 18.09.0
    3. docker-compose 1.23.2
后面版本号是我开发时使用的版本，具体安装过程请自行百度。

### 下载项目
在$GPATH/github.com/jason-wj/目录下执行如下命令：
```sh
git clone https://github.com/jason-wj/bitesla
```

#### 主要目录介绍
.  
|\_\_\_\_bitesla-server         //这是项目运行时候，mysql、redis、nsq等镜像的volume默认映射位置，我测试时候的所有数据均在这里，可直接使用  
|\_\_\_\_common                 //golang的通用内容  
|\_\_\_\_service                //微服务，其中有不同的服务  
| |\_\_\_\_service-api  
| |\_\_\_\_service-exchange  
| |\_\_\_\_service-strategy  
| |\_\_\_\_service-trader  
| |\_\_\_\_service-user  
|\_\_\_\_bitesla-config.ini     //用于配置当前项目处于哪个环境下，目前来说，哪个环境都一样  
|\_\_\_\_Makefile               //统一编译各服务  
|\_\_\_\_docker-compose.yml     //镜像部署目录  
|\_\_\_\_bitesla.sh             //脚本，便捷操作编译运行项目或者清理项目

### 项目启动
在根目录下，执行：
```sh
./bitesla.sh docker start
```
`备注`:
bitesla.sh脚本的其余命令的使用方式，可执行以下命令来参考：
```sh
./bitesla.sh help
```
![image](/img/5.jpg)

### 项目效果展示
项目启动后，在浏览器中输入：
`http://localhost:8090/swagger/index.html`
可以看到如下内容:
![image](/img/6.jpg)
![image](/img/7.jpg)
![image](/img/8.jpg)

### 项目接口基本使用
1. 获取token
上述页面中，使用`账户操作`部分的`邮箱注册接口`或者`邮箱登录接口`生成token，如下图所示：
![image](/img/9.jpg)

2. 将token复制到页面顶部的
![image](/img/10.jpg)

3. 然后就可以操作其余接口了

## 策略编写及运行
这个过程我新开了一个项目来单独描述，请参考这里：[bitesla策略编写(golang版)](https://github.com/jason-wj/bitesla-strategy-exec)
简单来说，就是在`strategy.go`文件中编辑好自己的策略后，将其中的内容复制出来，通过如下两个接口来实现运行：
1. 先将策略添加到数据库
2. 另一个接口执行策略
![image](/img/11.jpg)


## 特别感谢
>[QuantBot](https://github.com/jason-wj/QuantBot)

>[Samaritan](https://github.com/jason-wj/samaritan)

>[GoEx](https://github.com/nntaoli-project/GoEx)
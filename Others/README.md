## Go依赖管理

### 1 依赖管理演进

GOPATH -> Go Vendor -> Go Module

控制依赖库的版本问题

#### 1.1 GOPATH

gopath是Go语言支持的一个环境变量，vaule是Go项目的工作区

- bin	项目编译的二进制文件
- pkg   项目编译的中间产物，加速编译
- src    项目源码

项目代码直接依赖src下的代码

`go get `下载的最新版本包到src目录下

**弊端**

- 无法实现package的多版本控制



#### 1.2 GO Vendor

项目目录增加vendor文件，所有依赖包副本形式放在项目根目录的/vendor下

- 依赖寻址方式：vendor => gopath

- 通过每个项目引入一份依赖的副本，解决多项目需要同一个package依赖的冲突问题

**弊端**

- 无法控制依赖的版本
- 更新项目可能出现依赖冲突，导致编译错误

归根到底是不能很清晰的标识依赖的版本概念



#### 1.3 GO Module

- 通过go.mod文件管理依赖包版本
- 通过`go get/go mod`指令工具管理依赖包

终极目标：**定义版本规则和管理项目依赖关系**



### 2 依赖管理三要素

1. 配置文件，描述依赖	go.mod
2. 中心仓库管理依赖库    proxy
3. 本地工具                       go get/mod



### 3 依赖配置

#### 3.1 go.mod

![image-20230124170019094](https://images.ltd7.ltd/img/studygo/gomod.png)

- 依赖管理基本单元：如果是github前缀则表示可以从github找到这个模块
- 原生库：标识原生Go sdk版本
- 单元依赖：使用模块路径和版本来唯一标示



#### 3.2 version

gomod为了方便管理定义了版本规则

- 语义化版本

  `${MAJOR}.${MINOR}.${PATCH}`

​		**MAJOR**版本表示是不兼容的API，所以即使是同一库，不同MAJOR版本认为是不同的模块

​		**MINOR**版本通常是新增函数或功能，向后兼容

​		**PATCH**版本一般是修复bug

​		v1.1.0

- 基于commit的伪版本

​		`vx.0.0-yyyymmddhhmmss-abcdefgh1234`

​		基础版本前缀和语义化版本一样，时间戳是commit的时间，校验码是包含12位的哈希值



#### 3.3 indirect

依赖单元中的特殊标识符

- 直接依赖
- 间接依赖

goland的自动格式化会将两类分开



#### 3.4 incompatible

对于引入没有gomod文件但tag在v2及以上的库会在版本号后面加上`+incompatible`后缀



#### 3.5 依赖分发

直接使用版本管理仓库下载依赖存在多个问题

- 无法保证构建稳定性
- 无法保证依赖可用性
- 增加第三方压力（github...）
- 网络问题

**Proxy**

go proxy是一个服务站点，会缓存源站的内容，而版本不会改变，且在源站内容删除后依旧可用。使用后构建时会直接从proxy站点拉取依赖

GOPROXY环境变量设置代理站点列表



#### 3.6 工具

- `go get`

  1.17版本后是只能在项目目录下进行（含gomod文件的地方）下载更新，全局则需使用`go install`

- `go mod`
    - init	初始化，创建go.mod文件
    - download     下载模块到本地缓存
    - tidy	对当前项目所需的依赖自动添加删除
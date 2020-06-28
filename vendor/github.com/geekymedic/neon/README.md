# neon
[![Go Report Card](https://goreportcard.com/badge/github.com/geekymedic/neon)](https://goreportcard.com/report/github.com/geekymedic/neon)

## 系统架构

### 名词解释
* System 
	
	指一组使用同一套存储的、功能边界明确的、分层的(BFF/Service)一组微服务提供的系统
* BFF(Backend for Frontend)

	为客户端提供HTTP接口组装的服务层
* Service

	对后端其他service/bff提供Rpc接口的服务，服务本身不对外暴露
	
### neon
neon是一套使用本架构开发理念的基础开发套件、该库作为其他BFF/Service开发的基础依赖库，使用该库可以减少业务服务的开发负担
#### 目录结构
```
+ neon
  + bff     //bff层服务基础库
  + config  //配置文件处理
  + doc     //文档
  + errors  //带StackTrace的错误处理
  + logger  //统一log处理
  + plugin  //插件
  	+ db
  	+ es
  	+ metrics
  	+ redis
  	+ rpc
  + service //rpc层服务基础库
  + utils   //工具型类库
  	+ validator //参数校验依赖库
```

## 开发前准备
### 安装命令行工具

```shell
go get github.com/geekymedic/neon-cli
```

或者手动下载二进制包: [neon-cli](https://github.com/geekymedic/neon-cli/releases)

### IDE插件支持
* IDE插件支持
	* [GoLand](https://github.com/geekymedic/neon-idea-plugin)

## Doc

- [:cn:-zh](doc/README-zh.md)
- [:us:-en]()


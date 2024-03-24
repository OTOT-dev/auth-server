# auth-server

## 简介

基于属性的权限验证（ABAC，Attribute-Based Access Control）的微服务，为其他业务服务提供权限认证。



## 快速开始

- 启动项目-开发（含热加载）

  ```sh
  air
  ```

  > air启动后需要过一段时间才可以终止程序（立即`CTRL+C`后，程序可能仍然运行）。

- 获取/查看接口文档

  拉取项目后，安装[go swag](https://github.com/swaggo/gin-swagger)组件，执行以下命令生成swag文档。
  
  ```sh
  swag init
  ```
  
  然后可以**在debug模式下**前往http://<BaseUrl>/swagger/index.html查看。



## 参考

- [企业级的 Go 语言实战项目：认证和授权系统](https://github.com/marmotedu/iam/tree/v1.0.0)

- [go swag api文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#%E5%A3%B0%E6%98%8E%E5%BC%8F%E6%B3%A8%E9%87%8A%E6%A0%BC%E5%BC%8F)

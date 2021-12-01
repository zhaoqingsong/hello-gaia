# Hellobike Container Platform Service

## 配置（环境变量说明）

AppContainerPlatformService 的所有的配置为了适应云原生体系，均以环境变量方式存在。

| Key                       | Description           | Default | Example                                                       |
|---------------------------|-----------------------|---------|---------------------------------------------------------------|
| WEB\_LISTEN\_ADDR         | 监听端口                  |         | :8080                                                         |
| DB\_DRIVER                | 数据库 Driver            |         | postgres                                                      |
| DB\_CONNECTION            | 数据库连接                 |         | postgres://username:password@127\.0\.0\.1/cps?sslmode=disable |
| DB\_CONN\_MAX\_LIFETIME   | 数据库???                |         | 30                                                            |
| DB\_MAX\_OPEN\_CONNS      | 数据库最大连接数              |         | 100                                                           |
| DB\_MAX\_IDLE\_CONNS      | 数据库最大空闲连接             |         | 30                                                            |
| JENKINS\_URL              | Jenkins URL           |         | http://jenkins\-master                                        |
| JENKINS\_USERNAME         | Jenkins 用户名           |         | admin                                                         |
| JENKINS\_APITOKEN         | Jenkins token            |         |                                                               |
| JENKINS\_GITCREDENTIALSID | Jenkins 的 Git Credentials ID | |                                                               |
| DB\_DEBUG                 | 是否启用调试模式              | false   | true                                                          |
| CMP\_URL                  | 配置管理平台 URL            |         | http://cmp\.hellobike\.cn                                     |
| CMP\_TOKEN                | 配置管理平台 token（系统右上角可看） |         |                                                               |
| CRP\_URL                  | 发布平台 URL              |         | http://crp\.hellobike\.cn                                     |
| HARBOR\_HOST              | Harbor 主机名            |         | harbor\.hellobike\.cn                                         |
| HARBOR\_BASE\_URL         | Harbor API base URL   |         | http://$\{HARBOR\_HOST\}/api                                  |
| HARBOR\_USERNAME          | Harbor 系统用户名          |         | admin                                                         |
| HARBOR\_PASSWORD          | Harbor 系统密码           |         | Harbor12345                                                   |
| GITLAB\_URL               | Gitlab URL            |         | https://gitlab.hellobike.cn                                    |
| GITLAB\_TOKEN             | Gitlab Token          |         | xxxxxxxxxxxxxxxxxxxxxxxxxxx                                    |

*注：生产系统可以直接注入环境变量；开发环境可以在项目根目录中放置一个 `.env` 文件中以方便调试。*

## 初始化数据库

1. 新建一个数据库，如 `cps`，然后执行 `cps-cicd migrate` 就导入了数据库的表结构。
2. 数据的导入：TODO

## 运行

执行 `cps-cicd` 启动

## 项目结构

该项目遵循 kubernetes owners 规范：<https://go.k8s.io/owners>

## 提交代码

每次提交代码需要执行 `make check` 以检查代码依赖项、代码格式、代码漏洞、lint code，后才能提交。

## 相关文档

- Gin <https://github.com/gin-gonic/gin>
- gorm <http://gorm.io/docs/models.html>
- swaggo <https://github.com/swaggo/swag>

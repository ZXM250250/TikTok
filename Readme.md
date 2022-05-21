.


├── api
├── cmd
│   └── main.go
├── gen
│   └── go
├── deployments
├── integration_test
├── scripts
├── config
├── internal
│   ├── mock
│   ├── handler
│   ├── model
│   ├── dao
│   └── router
├── pkg
├── vendor
├── README.md
├── Makefile
├── .gitlab-ci.yml
├── go.mod
└── go.sum















api：存放项目相关的api/proto定义。
cmd：程序的主入口，也就是main.go所在的地方，如果项目有多个应用的话则可以再次划分多个应用入口目录。
gen：存放proto文件生成的目标代码。
deployments：存放项目部署的模板和配置文件。
integration_test：存放项目的集成测试文件（包括sql，docker-compose依赖等）。
scripts：存放项目build/lint/githooks等脚本，提供给Makefile使用。
config：存放项目配置定义以及注册方法（配置中心，db，redis等）。
internal：存放所有的内部实现（不对外暴露）的代码模块。
pkg：存放可被外部项目引入的组件和模块。
vendor：项目依赖管理是一个令人头痛的问题，对我来说把vendor提交到仓库是一个可以接受的方案。
README.md：项目的说明文件。
Makefile：用来构建和启动Go程序。
.gitlab-ci.yml：gitlab的CI定义文件，在其中定义各种测试任务（lint/build/unit_test）。
go.mod：go module可能是目前最好的依赖管理工具。

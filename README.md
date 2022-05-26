

### 代码结构
- controller

web接口,与http相关的代码都要放在这里,如参数合法性、格式转换等
- service

业务逻辑层,处于应用环境(如web、grpc等)与dao之间,负责复杂业务逻辑处理. 如业务逻辑十分简单,可省略service层,直接在dao层实现. *它不应关心运行环境*
- dao

负责数据库entity定义,及简单数据库操作.若涉及多表操作或较复杂的业务逻辑,应放在service层
- dao/dao_event.go

可集中定义dao中事件相关业务逻辑,如增加评论后增加反馈的评论计数这种非严格强制性业务逻辑(失败也不影响功能)
- conf

配置生成逻辑.配置可能通过两种途径获得:
     1. 默认配置文件为可执行文件所在目录的`config.toml`文件,也可通过`-c`参数传参,改变默认配置文件位置
     2. 通过acm动态配置.当可执行文件所在目录存在`acm.toml`文件且未通过`-c`参数改变默认配置文件位置时,此模式会启动,它通过拉取阿里云acm服务中保存的配置作为应用配置文件.同时,它会监测配置是否变更,如发生变更,则会退出当前服务,由其他进程管理软件重新拉起该服务
- route

web路由配置
- utils

工具类

#### xorm 目前不支持全文索引配置,因此自己执行sql.
> ALTER TABLE feedback_business ADD FULLTEXT INDEX search_text (search_text) WITH PARSER ngram;

一、清理项目（没用的踢掉）

二、导入、两个接口

三、完成简单的增删改

四、优化（schema、微服务解耦、组织用户解耦、市场解耦、元数据、前后端关系）

1、人员、账号要与单位建立对应关系表，而不是加入一个单位就新建一个人员、一个账号
2、表间关联统一用id关联，之前存在部分id关联，部分code关联的情况，比较混乱
3、数据删除时应该物理删除，而不是逻辑删除（等保要求）
4、应用开发商按照前端查找平台域
5、平台租户（1.角色拉取所有,2.不能分配应用菜单）
6、租户类型设置，按照租户类型过滤（租户分组[树状]）
7、优化代码字典、元数据（租户定义审核）

五、做一中心接口<=>平台发送中心注册一下（大市场）

六、与华为能力、复杂美能力（IM）、安恒能力对接

七、平台分布式对齐（-IPFS）

八、graphql 实现数据库开放给第三方（权限控制、参考微信公众号后台）
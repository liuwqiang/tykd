#### gateway.api 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | id字段,服务端自动生成 | varchar(20) | PRI | NO |  |  |
| 2 | name | api的名字 | varchar(50) |  | NO |  |  |
| 3 | description | api的描述内容 | varchar(100) |  | YES |  |  |
| 4 | listen_path | 监听的地址 | varchar(50) |  | NO |  |  |
| 5 | target_url | 目标服务的地址 | varchar(100) |  | NO |  |  |
| 6 | onwer_id | api拥有者的id | varchar(20) |  | NO |  |  |
| 7 | group_id | api分组的id | varchar(20) |  | NO |  |  |
| 8 | group_name | api分组的名称 | varchar(50) |  | NO |  |  |
| 9 | create_date | api创建日期 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 10 | update_date | api的修改日期 | datetime |  | YES |  |  |
| 11 | auth_mode | 认证模式 [OPEN(无认证),APPID(简单认证),SIGNATURE(签名认证)] | varchar(20) |  | NO |  |  |

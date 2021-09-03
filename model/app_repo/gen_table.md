#### gateway.app 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | id字段,服务端自动生成 | varchar(20) | PRI | NO |  |  |
| 2 | name | app的名字 | varchar(50) |  | NO |  |  |
| 3 | description | app的描述内容 | varchar(100) |  | NO |  |  |
| 4 | app_secret | app的密钥,可重置 | varchar(20) |  | NO |  |  |
| 5 | owner_id | app拥有者的id | varchar(20) |  | NO |  |  |
| 6 | create_date | app创建日期 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 7 | update_date | app的修改日期 | datetime |  | YES |  |  |

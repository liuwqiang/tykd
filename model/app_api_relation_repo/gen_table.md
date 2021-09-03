#### gateway.app_api_relation 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | api_id | api的id字段 | varchar(20) |  | NO |  |  |
| 2 | app_id | app的id字段 | varchar(20) |  | NO |  |  |
| 3 | expired_date | 过期时间 | datetime |  | YES |  |  |
| 4 | rate | 速率 | int(8) |  | YES |  |  |
| 5 | per | 和rate配合使用,每per秒允许rate个请求通过 | int(10) |  | YES |  |  |
| 6 | quota | 配额 | int(10) |  | YES |  |  |
| 7 | quota_renewal_rate | 和quota配合使用,每quota_renewal_rate个周期更新一次配额 | int(10) |  | YES |  |  |
| 8 | extra_data | 额外数据,传递给后端api | json |  | YES |  |  |
| 9 | create_date | api和app映射的创建日期 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 10 | update_date |  | datetime |  | YES |  |  |

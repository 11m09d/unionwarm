# unionwarm
## 什么是 Unionwarm

Unionwarm 是openfalcon中的Portal的Go版本，目前大的功能欠缺expressions和nodata 其它基本完成。
其中部分API的方式自行改写适用于目前的场景

### 1. 待办事宜 [Todo 列表]
### 2. Restful接口介绍

|序号 | 内容                                    | 作用   |  输入  | 输出 |
|--------|:--------|:--------|:--------|:--------|
|01  | /                                             | 首页        |    | |
|02  | /zjz_work_cmdb/cmdb/logout                    | 登出        |    | |
|03  | /zjz_work_cmdb/cmdb/falcontpl                 | 模版页      |    | |
|04  | /zjz_work_cmdb/cmdb/falconnewtpl              | 创建模版页   |        | |
|05  | /zjz_work_cmdb/cmdb/falconaddtpl              | 增加模版     |{Tpl_name: TemplateName,Create_user: Username} | {"success":true,"message":id}|
|06  | /zjz_work_cmdb/cmdb/falcondeltpl              | 删除模版     |{Id:Id} |{"success":true,"message":id} |
|07  | /zjz_work_cmdb/cmdb/falconnewstrategy         | 创建策略页   |    | |
|08  | /zjz_work_cmdb/cmdb/falconshowstrategy        | 展示策略       | {Id:Id}  | {"data":strategyall} |
|09  | /zjz_work_cmdb/cmdb/falconaddstrategy         | 增加策略      |  'sid': sid,'metric':metric,'tags': tag,'max_step':max_step ,'priority': priority,'note': note,'func': func,'op': op,'right_value': right_value,'run_begin': run_begin,'run_end': run_end,'tpl_id': tpl_id  | {"success":true,"message":id} |
|10  | /zjz_work_cmdb/cmdb/falconnewgrp              |        |    | |
|11  | /zjz_work_cmdb/cmdb/falconaddgrp              |        |    | |
|12  | /zjz_work_cmdb/cmdb/falcondelgrp              |        |    | |
|13  | /zjz_work_cmdb/cmdb/falconbindgphost          |        |    | |
|14  | /zjz_work_cmdb/cmdb/falconhg                  |        |    | |
|15  | /zjz_work_cmdb/cmdb/falcongphost              |        |    | |
|16  | /zjz_work_cmdb/cmdb/falconaddhosts            |        |    | |
|17  | /zjz_work_cmdb/cmdb/falconhgdel               |        |    | |
|18  | /zjz_work_cmdb/cmdb/falconshowgrptpl          |        |    | |
|19  | /zjz_work_cmdb/cmdb/falcongptpl               |        |    | |
|20  | /zjz_work_cmdb/cmdb/falconquerytpl            |        |    | |
|21  | /zjz_work_cmdb/cmdb/falconbindgrptpl          |        |    | |
|22  | /zjz_work_cmdb/cmdb/falcongrptpldel           |        |    | |
|23  | /zjz_work_cmdb/cmdb/falconparenttpl           |        |    | |
|24  | /zjz_work_cmdb/cmdb/falcongetactioncount      |        |    | |
|25  | /zjz_work_cmdb/cmdb/falcongetuiccount         |        |    | |
|26  | /zjz_work_cmdb/cmdb/falcongetactiondetail     |        |    | |
|27  | /zjz_work_cmdb/cmdb/falconactionmodify        |        |    | |
|28  | /zjz_work_cmdb/cmdb/api/action/:id([0-9]+)    |        |    | |
|29  | /zjz_work_cmdb/cmdb/falcondelstrategy         |        |    | |
|30  | /zjz_work_cmdb/cmdb/falconhostmaintain        |        |    | |
|31  | /zjz_work_cmdb/cmdb/falconucintf              |        |    | |
|32  | /zjz_work_cmdb/cmdb/falconupdatestrategy      |        |    | |
|33  | /zjz_work_cmdb/cmdb/template/view/:id([0-9]+) |        |    | |
|34  | /zjz_work_cmdb/cmdb/falconagg                 |        |    | |
|35  | /zjz_work_cmdb/cmdb/falconnewcls              |        |    | |
|36  | /zjz_work_cmdb/cmdb/falconaddcls              |        |    | |
|37  | /zjz_work_cmdb/cmdb/falcondelcls              |        |    | |
|38  | /zjz_work_cmdb/cmdbapp/falconeditcls          |        |    | |

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
|08  | /zjz_work_cmdb/cmdb/falconshowstrategy        | 展示策略     | {Id:Id}  | {"data":strategyall} |
|09  | /zjz_work_cmdb/cmdb/falconaddstrategy         | 增加策略     |  {'sid': sid,'metric':metric,'tags': tag,'max_step':max_step ,'priority': priority,'note': note,'func': func,'op': op,'right_value': right_value,'run_begin': run_begin,'run_end': run_end,'tpl_id': tpl_id}  | {"success":true,"message":id} |
|10  | /zjz_work_cmdb/cmdb/falconnewgrp              | 创建监控组页  |    | |
|11  | /zjz_work_cmdb/cmdb/falconaddgrp              | 增加监控组    | {Grp_name: GroupName,Create_user: Username}   | {"success":true,"message":id}|
|12  | /zjz_work_cmdb/cmdb/falcondelgrp              | 删除监控组    | {Id:Id}  | {"success":true,"message":id} |
|13  | /zjz_work_cmdb/cmdb/falconhg                  | 监控组页      |    | |
|14  | /zjz_work_cmdb/cmdb/falcongphost              | 展示监控组设备 |  {Id:Id}  | {"data":hostlists}|
|15  | /zjz_work_cmdb/cmdb/falconaddhosts            | 增加监控设备   | {'hosts': hosts,'grp_id':grp_id} | {"success":true}|
|16  | /zjz_work_cmdb/cmdb/falconhgdel               | 删除监控组设备 |  {Grp_id:grp_id,Id:Id}  | {"success":true,"message":id} |
|17  | /zjz_work_cmdb/cmdb/falcongptpl               | 展示关联模版      | {Id:Id}   | {"data": templateslists}|
|18  | /zjz_work_cmdb/cmdb/falconbindgrptpl          | 监控组绑定模版       |{'tpl_id': tpl_id,'grp_id':grp_id}   | {"success":true,"message":id} |
|19  | /zjz_work_cmdb/cmdb/falconparenttpl           | 设置父模版       |   {'tpl_id': tpl_id,'parent_tpl_id': parent_tpl_id} | {"success":true}|
|20  | /zjz_work_cmdb/cmdb/falcongetactiondetail     | 展示告警对象详细信息    |{tpl_id:Id}    | {"success":true,"message":actions} |
|21  | /zjz_work_cmdb/cmdb/falconactionmodify        | 修改告警对象信息  | {'action_id': action_id,'tpl_id':tpl_id,'uic': uic,'callback_url': callback_url,'callback': callback,'after_callback_mail': after_callback_mail_num,'after_callback_sms': after_callback_sms_num,'before_callback_mail': before_callback_mail_num,'before_callback_sms': before_callback_sms_num}  | {"success":true}|
|22  | /zjz_work_cmdb/cmdb/api/action/:id([0-9]+)    | ALARM调用 根据id查action       | {:id}   | {"data":action,"msg":msg} |
|23  | /zjz_work_cmdb/cmdb/falcondelstrategy         | 删除策略     |   {'Id': Id} |{"success":true} |
|24  | /zjz_work_cmdb/cmdb/falconhostmaintain        | 设置监控设备维护时间       |   {'Id': Id,'host_begin': host_begin,'host_end': host_end,'grp_id': grp_id} | {"data":hostlists} |
|25  | /zjz_work_cmdb/cmdb/falconucintf              |   统一采集     |    | |
|26  | /zjz_work_cmdb/cmdb/falconupdatestrategy      |      |    | |
|27  | /zjz_work_cmdb/cmdb/template/view/:id([0-9]+) | Warm告警页 根据tpl_id查详细信息      | {:id}   | {'action':action,'tpl':tpl,'parent_tpl':parent_tpl,'strategy':strategy}|
|28  | /zjz_work_cmdb/cmdb/falconagg                 |  集群告警页      |    | |
|29  | /zjz_work_cmdb/cmdb/falconnewcls              |  创建集群告警页     |    | |
|30  | /zjz_work_cmdb/cmdb/falconaddcls              |  增加集群告警   |  {Numerator: Numerator,Denominator: Denominator,Endpoint: Endpoint,Metric: Metric,Tags: Tags,Grp_id: Grp_id1}  | {"success":true} |
|31  | /zjz_work_cmdb/cmdb/falcondelcls              |  删除集群告警      |  {Id:Id}  | {"success":true}|
|32  | /zjz_work_cmdb/cmdbapp/falconeditcls          |  编辑集群告警      |   {id:Id,algorithm:algorithmName,endpoint:endpointName,metric:metricName,tags:tagsName} | {"success":true}|

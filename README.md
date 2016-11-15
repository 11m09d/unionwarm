# unionwarm
## 什么是 Unionwarm

Unionwarm 是openfalcon中的Portal的Go版本，目前大的功能欠缺expressions和nodata 其它基本完成。
其中部分API的方式自行改写适用于目前的场景

### 1. 内容介绍

#### 新增策略模版、模版包含多种策略，可以设置告警对象组（UIC）; 可以关联父模版; 使用统一UIC登陆，创建者为UIC用户

![告警模版](https://github.com/chenqi123/unionwarm/blob/master/png/1.png?raw=true =600x250)

#### 新增监控组,能够自动/手动添加监控设备，关联策略模版

![告警组](https://github.com/chenqi123/unionwarm/blob/master/png/2.png?raw=true =600x250)

#### 新增自定义监控，绑定监控组，设置集群算法，重新生成新监控指标回送统一采集

![告警组](https://github.com/chenqi123/unionwarm/blob/master/png/3.png?raw=true =600x250)

### 2. 待办事宜 [Todo 列表]
[TAB][TAB]删除告警组时，集群自定义监控没有删除
[TAB][TAB]open-falcon 中的HBS 当策略组和主机都成倍增加的时候，策略输出可能是个问题，目前所有的judge 拿到一份相同的总体策略，当策略1000*主机1000就是100万的Map,量太大(不知道是不是我的理解有误)

### 3. Restful接口介绍
|序号 | 内容                                    | 作用   |  输入  | 输出 |
|--------|:-----|:--------|:--------|:--------|
|01  | /                                             | 首页        |    | |
|02  | /logout                    | 登出        |    | |
|03  | /falcontpl                 | 模版页      |    | |
|04  | /falconnewtpl              | 创建模版页   |        | |
|05  | /falconaddtpl              | 增加模版     |{Tpl_name: TemplateName, | {"success":true,|
||||Create_user: Username}|"message":id}|
|06  | /falcondeltpl              | 删除模版     |{Id:Id} |{"success":true, |
|||||"message":id}|
|07  | /falconnewstrategy         | 创建策略页   |    | |
|08  | /falconshowstrategy        | 展示策略     | {Id:Id}  | {"data":strategyall} |
|09  | /falconaddstrategy         | 增加策略     |  {'sid': sid,  | {"success":true, |
||||'metric':metric,|"message":id}|
||||'tags': tag,||
||||'max_step':max_step ,||
||||'priority': priority,||
||||'note': note,||
||||'func': func,||
||||'op': op,||
||||'right_value': right_value,||
||||'run_begin': run_begin,||
||||'run_end': run_end,||
||||'tpl_id': tpl_id}||
|10  | /falconnewgrp              | 创建监控组页  |    | |
|11  | /falconaddgrp              | 增加监控组    | {Grp_name: GroupName,   | {"success":true,|
||||Create_user: Username}|"message":id}|
|12  | /falcondelgrp              | 删除监控组    | {Id:Id}  | {"success":true |
|||||,"message":id}|
|13  | /falconhg                  | 监控组页      |    | |
|14  | /falcongphost              | 展示监控组设备 |  {Id:Id}  | {"data":hostlists}|
|15  | /falconaddhosts            | 增加监控设备   | {'hosts': hosts, | {"success":true}|
||||'grp_id':grp_id}||
|16  | /falconhgdel               | 删除监控组设备 |  {Grp_id:grp_id,  | {"success":true, |
||||Id:Id}|"message":id}|
|17  | /falcongptpl               | 展示关联模版      | {Id:Id}   | {"data": templateslists}|
|18  | /falconbindgrptpl          | 监控组绑定模版       |{'tpl_id': tpl_id, | {"success":true,|
||||'grp_id':grp_id}|"message":id}|
|19  | /falconparenttpl           | 设置父模版       |   {'tpl_id': tpl_id, | {"success":true}|
||||'parent_tpl_id': parent_tpl_id}||
|20  | /falcongetactiondetail     | 展示告警对象详细信息    |{tpl_id:Id}    | {"success":true,|
|||||"message":actions} |
|21  | /falconactionmodify        | 修改告警对象信息  | {'action_id': action_id,'callback': callback,  | {"success":true}|
||||'tpl_id':tpl_id,||
||||'uic': uic,||
||||'callback_url': callback_url,||
||||'after_callback_mail': after_callback_mail_num,||
||||'after_callback_sms': after_callback_sms_num,||
||||'before_callback_mail': before_callback_mail_num,||
||||'before_callback_sms': before_callback_sms_num}||
|22  | /api/action/:id([0-9]+)    | ALARM调用 根据id查action       | {:id}   | {"data":action, |
|||||"msg":msg}|
|23  | /falcondelstrategy         | 删除策略     |   {'Id': Id} |{"success":true} |
|24  | /falconhostmaintain        | 设置监控设备维护时间       |   {'Id': Id, | {"data":hostlists} |
||||'host_begin': host_begin,||
||||'host_end': host_end,||
||||'grp_id': grp_id}||
|25  | /falconucintf              |   统一采集     |    | |
|26  | /falconupdatestrategy      |      |    | |
|27  | /template/view/:id([0-9]+) | Warm告警页 根据tpl_id查详细信息      | {:id}   | {'action':action,|
|||||'tpl':tpl,|
|||||'parent_tpl':parent_tpl,|
|||||'strategy':strategy}|
|28  | /falconagg                 |  集群告警页      |    | |
|29  | /falconnewcls              |  创建集群告警页     |    | |
|30  | /falconaddcls              |  增加集群告警   |  {Numerator: Numerator,Denominator: Denominator,  | {"success":true} |
||||Endpoint: Endpoint,||
||||Metric: Metric,||
||||Tags: Tags,||
||||Grp_id: Grp_id1}||
|31  | /falcondelcls              |  删除集群告警      |  {Id:Id}  | {"success":true}|
|32  | app/falconeditcls          |  编辑集群告警      |   {id:Id, | {"success":true}|
||||algorithm:algorithmName,||
||||endpoint:endpointName,||
||||metric:metricName,||
||||tags:tagsName}||




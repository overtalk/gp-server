# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/announcement.proto](#proto/announcement.proto)
    - [AddAnnouncementReq](#protocol.AddAnnouncementReq)
    - [AddAnnouncementResp](#protocol.AddAnnouncementResp)
    - [AnnouncementDetailReq](#protocol.AnnouncementDetailReq)
    - [AnnouncementDetailResp](#protocol.AnnouncementDetailResp)
    - [AnnouncementsReq](#protocol.AnnouncementsReq)
    - [AnnouncementsResp](#protocol.AnnouncementsResp)
    - [DelAnnouncementReq](#protocol.DelAnnouncementReq)
    - [DelAnnouncementResp](#protocol.DelAnnouncementResp)
    - [EditAnnouncementReq](#protocol.EditAnnouncementReq)
    - [EditAnnouncementResp](#protocol.EditAnnouncementResp)
  
  
  
  

- [proto/class_manage.proto](#proto/class_manage.proto)
    - [AddClassReq](#protocol.AddClassReq)
    - [AddClassResp](#protocol.AddClassResp)
    - [EditClassReq](#protocol.EditClassReq)
    - [EditClassResp](#protocol.EditClassResp)
    - [GetClassByIDReq](#protocol.GetClassByIDReq)
    - [GetClassByIDResp](#protocol.GetClassByIDResp)
    - [GetClassesReq](#protocol.GetClassesReq)
    - [GetClassesResp](#protocol.GetClassesResp)
    - [MemberManageReq](#protocol.MemberManageReq)
    - [MemberManageResp](#protocol.MemberManageResp)
  
    - [MemberManageReq.ManageType](#protocol.MemberManageReq.ManageType)
  
  
  

- [proto/common.proto](#proto/common.proto)
    - [Announcement](#protocol.Announcement)
    - [Class](#protocol.Class)
    - [Match](#protocol.Match)
    - [Paper](#protocol.Paper)
    - [Problem](#protocol.Problem)
    - [ProblemExample](#protocol.ProblemExample)
    - [RankItem](#protocol.RankItem)
    - [SubmitRecord](#protocol.SubmitRecord)
    - [UserInfo](#protocol.UserInfo)
  
    - [ProblemDifficluty](#protocol.ProblemDifficluty)
    - [Role](#protocol.Role)
  
  
  

- [proto/login.proto](#proto/login.proto)
    - [LoginReq](#protocol.LoginReq)
    - [LoginResp](#protocol.LoginResp)
    - [LogoutResp](#protocol.LogoutResp)
  
  
  
  

- [proto/match.proto](#proto/match.proto)
    - [EditMatchReq](#protocol.EditMatchReq)
    - [EditMatchResp](#protocol.EditMatchResp)
    - [GetMatchByIDReq](#protocol.GetMatchByIDReq)
    - [GetMatchByIDResp](#protocol.GetMatchByIDResp)
    - [GetMatchesReq](#protocol.GetMatchesReq)
    - [GetMatchesResp](#protocol.GetMatchesResp)
    - [GetPaperByIDReq](#protocol.GetPaperByIDReq)
    - [GetPaperByIDResp](#protocol.GetPaperByIDResp)
    - [NewMatchReq](#protocol.NewMatchReq)
    - [NewMatchResp](#protocol.NewMatchResp)
  
  
  
  

- [proto/problem_manage.proto](#proto/problem_manage.proto)
    - [AddProblemReq](#protocol.AddProblemReq)
    - [AddProblemResp](#protocol.AddProblemResp)
    - [EditProblemReq](#protocol.EditProblemReq)
    - [EditProblemResp](#protocol.EditProblemResp)
    - [GetProblemByIDReq](#protocol.GetProblemByIDReq)
    - [GetProblemByIDResp](#protocol.GetProblemByIDResp)
    - [GetProblemsReq](#protocol.GetProblemsReq)
    - [GetProblemsResp](#protocol.GetProblemsResp)
  
  
  
  

- [proto/rank.proto](#proto/rank.proto)
    - [RankListReq](#protocol.RankListReq)
    - [RankListResp](#protocol.RankListResp)
  
  
  
  

- [proto/status.proto](#proto/status.proto)
    - [Status](#protocol.Status)
  
    - [Code](#protocol.Code)
  
  
  

- [proto/user_manage.proto](#proto/user_manage.proto)
    - [AddUsersReq](#protocol.AddUsersReq)
    - [AddUsersResp](#protocol.AddUsersResp)
    - [DelUsersReq](#protocol.DelUsersReq)
    - [DelUsersResp](#protocol.DelUsersResp)
    - [GetUsersReq](#protocol.GetUsersReq)
    - [GetUsersResp](#protocol.GetUsersResp)
    - [UpdateUsersReq](#protocol.UpdateUsersReq)
    - [UpdateUsersResp](#protocol.UpdateUsersResp)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/announcement.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/announcement.proto



<a name="protocol.AddAnnouncementReq"></a>

### AddAnnouncementReq
增加新公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| announcement | [Announcement](#protocol.Announcement) |  |  |






<a name="protocol.AddAnnouncementResp"></a>

### AddAnnouncementResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.AnnouncementDetailReq"></a>

### AnnouncementDetailReq
获取公告具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.AnnouncementDetailResp"></a>

### AnnouncementDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| announcement | [Announcement](#protocol.Announcement) |  |  |






<a name="protocol.AnnouncementsReq"></a>

### AnnouncementsReq
获取所有公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.AnnouncementsResp"></a>

### AnnouncementsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| announcements | [Announcement](#protocol.Announcement) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.DelAnnouncementReq"></a>

### DelAnnouncementReq
删除公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.DelAnnouncementResp"></a>

### DelAnnouncementResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditAnnouncementReq"></a>

### EditAnnouncementReq
修改公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| announcement | [Announcement](#protocol.Announcement) |  |  |






<a name="protocol.EditAnnouncementResp"></a>

### EditAnnouncementResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 

 

 

 



<a name="proto/class_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/class_manage.proto



<a name="protocol.AddClassReq"></a>

### AddClassReq
新增班级


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [Class](#protocol.Class) |  |  |






<a name="protocol.AddClassResp"></a>

### AddClassResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditClassReq"></a>

### EditClassReq
修改班级信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [Class](#protocol.Class) |  | 只需要填充id以及改变的field |






<a name="protocol.EditClassResp"></a>

### EditClassResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetClassByIDReq"></a>

### GetClassByIDReq
根据ID获得班级具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetClassByIDResp"></a>

### GetClassByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| class | [Class](#protocol.Class) |  |  |






<a name="protocol.GetClassesReq"></a>

### GetClassesReq
批量获取班级信息(只包括基础的信息，班级名称)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetClassesResp"></a>

### GetClassesResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| classes | [Class](#protocol.Class) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.MemberManageReq"></a>

### MemberManageReq
小组成员管理


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| manage_type | [MemberManageReq.ManageType](#protocol.MemberManageReq.ManageType) |  |  |
| class_id | [int64](#int64) |  |  |
| member_id | [int64](#int64) |  |  |






<a name="protocol.MemberManageResp"></a>

### MemberManageResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 


<a name="protocol.MemberManageReq.ManageType"></a>

### MemberManageReq.ManageType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DELETE | 0 | 删除小组成员 |
| SET_ADMINISTRATOR | 1 | 设置成管理员 |
| CANCEL_ADMINISTRATOR | 2 | 取消管理员 |


 

 

 



<a name="proto/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/common.proto



<a name="protocol.Announcement"></a>

### Announcement
Announcement : 公告，包括班级公告和全局公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| publisher | [string](#string) |  | 发布人的姓名 |
| title | [string](#string) |  | 公告标题 |
| detail | [string](#string) |  |  |
| create_time | [int64](#int64) |  |  |
| last_update_time | [int64](#int64) |  |  |






<a name="protocol.Class"></a>

### Class
Class : 班级信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 班级id |
| tutor | [string](#string) |  | 导师的姓名 |
| name | [string](#string) |  | 班级名称 |
| introduction | [string](#string) |  | 班级简介 |
| number | [int64](#int64) |  | 班级人数 |
| is_check | [bool](#bool) |  | 加入班级设置：false（无需审核，运行任何人进入），true（需要导师审核） |
| create_time | [int64](#int64) |  | 班级创建时间 |
| announcements | [Announcement](#protocol.Announcement) | repeated | 班级公告 |






<a name="protocol.Match"></a>

### Match
Match : 比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| is_public | [bool](#bool) |  | 是否是公开比赛 |
| start_time | [int64](#int64) |  | 开始时间 |
| duration | [int64](#int64) |  | 时间长度 |
| is_over | [bool](#bool) |  | 是否结束 |
| name | [string](#string) |  | 比赛名称 |
| intriduction | [string](#string) |  | 比赛简介 |
| paper_id | [int64](#int64) |  | 试卷id |






<a name="protocol.Paper"></a>

### Paper
Paper : 试卷


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| problems | [Problem](#protocol.Problem) | repeated | 题目 |
| difficulty | [int64](#int64) |  | 组卷的参数 |
| problem_num | [int64](#int64) |  |  |
| knowledge_points | [int64](#int64) | repeated |  |






<a name="protocol.Problem"></a>

### Problem
Problem : 题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 题目id |
| title | [string](#string) |  | 题目标题 |
| description | [string](#string) |  | 题目描述 |
| in | [string](#string) |  | 输入 |
| out | [string](#string) |  | 输出 |
| hint | [string](#string) |  | 题目提示 |
| in_out_examples | [ProblemExample](#protocol.ProblemExample) | repeated | 输入输出样例 |
| judge_limit_time | [int64](#int64) |  | 判题时间限制 |
| judge_limit_mem | [int64](#int64) |  | 判题内存限制 |
| tags | [int64](#int64) | repeated | 题目标签 |
| difficluty | [ProblemDifficluty](#protocol.ProblemDifficluty) |  | 难度 |
| submit_time | [int64](#int64) |  | 提交次数 |
| accept_time | [int64](#int64) |  | 通过次数 |






<a name="protocol.ProblemExample"></a>

### ProblemExample
ProblemExample : 题目输入输出样例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| input | [string](#string) |  |  |
| output | [string](#string) |  |  |






<a name="protocol.RankItem"></a>

### RankItem
RankItem : 排名item


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ranking | [int64](#int64) |  | 排名 |
| user_id | [int64](#int64) |  | 用户id |
| name | [string](#string) |  | 用户姓名 |
| pass_num | [int64](#int64) |  | 通过题目数量 |
| submit_num | [int64](#int64) |  | 提交次数 |






<a name="protocol.SubmitRecord"></a>

### SubmitRecord
SubmitRecord : 提交情况


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  | 题目 |
| submit_time | [int64](#int64) |  | 提交时间戳 |
| is_pass | [bool](#bool) |  | 是否通过 |






<a name="protocol.UserInfo"></a>

### UserInfo
UserInfo : 用户基本信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| role | [Role](#protocol.Role) |  |  |
| name | [string](#string) |  |  |
| sex | [bool](#bool) |  |  |
| phone | [string](#string) |  |  |
| email | [string](#string) |  |  |
| school | [string](#string) |  |  |
| last_login | [int64](#int64) |  |  |
| create | [int64](#int64) |  |  |
| account | [string](#string) |  | 这两个字段只有在用户管理中的新增用户才会用到, 客户端向服务端发送数据是填充 |
| password | [string](#string) |  |  |





 


<a name="protocol.ProblemDifficluty"></a>

### ProblemDifficluty
ProblemDifficluty : 题目难度

| Name | Number | Description |
| ---- | ------ | ----------- |
| EASY | 0 |  |
| MEDIUM | 1 |  |
| HARD | 2 |  |



<a name="protocol.Role"></a>

### Role
Role : 用户角色（学生/老师...）

| Name | Number | Description |
| ---- | ------ | ----------- |
| STUDENT | 0 |  |
| TEACHER | 1 |  |
| MANAGER | 2 |  |


 

 

 



<a name="proto/login.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login.proto



<a name="protocol.LoginReq"></a>

### LoginReq
登陆 (post)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="protocol.LoginResp"></a>

### LoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| token | [string](#string) |  |  |
| user | [UserInfo](#protocol.UserInfo) |  | 用户信息 |
| submit_records | [SubmitRecord](#protocol.SubmitRecord) | repeated | submit记录 （ TODO: 可以考虑提到新的协议中） |






<a name="protocol.LogoutResp"></a>

### LogoutResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |





 

 

 

 



<a name="proto/match.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/match.proto



<a name="protocol.EditMatchReq"></a>

### EditMatchReq
编辑比赛题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match | [Match](#protocol.Match) |  |  |






<a name="protocol.EditMatchResp"></a>

### EditMatchResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_ok | [bool](#bool) |  |  |






<a name="protocol.GetMatchByIDReq"></a>

### GetMatchByIDReq
拿到所有的比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetMatchByIDResp"></a>

### GetMatchByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| match | [Match](#protocol.Match) |  |  |






<a name="protocol.GetMatchesReq"></a>

### GetMatchesReq
拿到所有的比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetMatchesResp"></a>

### GetMatchesResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |
| matches | [Match](#protocol.Match) | repeated |  |






<a name="protocol.GetPaperByIDReq"></a>

### GetPaperByIDReq
获取比赛试卷


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetPaperByIDResp"></a>

### GetPaperByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| paper | [Paper](#protocol.Paper) |  |  |






<a name="protocol.NewMatchReq"></a>

### NewMatchReq
创建比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paper | [Paper](#protocol.Paper) |  |  |
| match | [Match](#protocol.Match) |  |  |






<a name="protocol.NewMatchResp"></a>

### NewMatchResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| result | [bool](#bool) |  |  |





 

 

 

 



<a name="proto/problem_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/problem_manage.proto



<a name="protocol.AddProblemReq"></a>

### AddProblemReq
新增题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  |  |






<a name="protocol.AddProblemResp"></a>

### AddProblemResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditProblemReq"></a>

### EditProblemReq
编辑题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  | 需要id |






<a name="protocol.EditProblemResp"></a>

### EditProblemResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetProblemByIDReq"></a>

### GetProblemByIDReq
根据ID获得题目具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetProblemByIDResp"></a>

### GetProblemByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| problem | [Problem](#protocol.Problem) |  |  |






<a name="protocol.GetProblemsReq"></a>

### GetProblemsReq
获取全部题目信息（只下发 id &amp; title &amp; difficulty &amp; pass_rate）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [int64](#int64) |  |  |
| get_all | [bool](#bool) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetProblemsResp"></a>

### GetProblemsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| problems | [Problem](#protocol.Problem) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/rank.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/rank.proto



<a name="protocol.RankListReq"></a>

### RankListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.RankListResp"></a>

### RankListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| rank_items | [RankItem](#protocol.RankItem) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |
| pos | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/status.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/status.proto



<a name="protocol.Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [Code](#protocol.Code) |  |  |
| message | [string](#string) |  |  |





 


<a name="protocol.Code"></a>

### Code


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | ok |
| INTERNAL | 1 | 服务端内部错误 |
| DATA_LOSE | 2 | 数据序列化错误 |
| NO_TOKEN | 3 | 没有token |
| UNAUTHORIZATED | 4 | token错误 |
| PERMISSION_DENIED | 5 | 没有权限 |


 

 

 



<a name="proto/user_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/user_manage.proto



<a name="protocol.AddUsersReq"></a>

### AddUsersReq
批量增加用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserInfo](#protocol.UserInfo) | repeated |  |






<a name="protocol.AddUsersResp"></a>

### AddUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [UserInfo](#protocol.UserInfo) | repeated |  |
| fail | [UserInfo](#protocol.UserInfo) | repeated |  |






<a name="protocol.DelUsersReq"></a>

### DelUsersReq
批量删除用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users_id | [int64](#int64) | repeated |  |






<a name="protocol.DelUsersResp"></a>

### DelUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [int64](#int64) | repeated |  |
| fail | [int64](#int64) | repeated |  |






<a name="protocol.GetUsersReq"></a>

### GetUsersReq
批量获取用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [Role](#protocol.Role) |  |  |
| get_all | [bool](#bool) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetUsersResp"></a>

### GetUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| users | [UserInfo](#protocol.UserInfo) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.UpdateUsersReq"></a>

### UpdateUsersReq
批量修改用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserInfo](#protocol.UserInfo) | repeated | 只需要填充id以及改变的field |






<a name="protocol.UpdateUsersResp"></a>

### UpdateUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [UserInfo](#protocol.UserInfo) | repeated |  |
| fail | [UserInfo](#protocol.UserInfo) | repeated |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |


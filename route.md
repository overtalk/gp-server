# 路由

## 登陆模块


| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 登陆 | /login  | POST |  login.proto --> LoginReq & LoginResp |
| 登出 | /logout | GET | login.proto --> LogoutResp |
 



## 管理员模块

### 用户管理

| 功能描述 | 路由 | HTTP方法 |protobuf |  
| ----- | ---- | ----- |  ----- |
| 获得用户 | /getUsers  | POST |  user_manage.proto --> GetUsersReq & GetUsersResp |
| 新增用户 | /addUsers | POST |  user_manage.proto --> AddUsersReq & AddUsersResp |
| 更新用户 | /updateUsers | POST |  user_manage.proto --> UpdateUsersReq & UpdateUsersResp |
| 删除用户 | /delUsers | POST |  user_manage.proto --> DelUsersReq & DelUsersResp |

### 题目管理 （目前没有判题文件的部分）

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 获得所有题目 （只有部分信息，包括题目标题，难度...） | /getProblems  | POST |  problem_manage.proto --> GetProblemsReq & GetProblemsResp |
| 获得题目具体信息 | /getProblemByID | POST | problem_manage.proto --> GetProblemByIDReq & GetProblemByIDResp |
| 新增题目 | /addProblem | POST | problem_manage.proto --> AddProblemReq & AddProblemResp |
| 编辑题目 | /editProblem | POST | problem_manage.proto --> EditProblemReq & EditProblemResp |


### 班级管理 

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 获得所有班级 （只有部分信息，包括班级名称，创建时间，总人数...） | /getClasses  | POST | class_manage.proto --> GetClassesReq & GetClassesResp |
| 获得班级具体信息 | /getClassByID | POST | class_manage.proto --> GetClassByIDReq & GetClassByIDResp |
| 新增班级 | /addClass | POST | class_manage.proto --> AddClassReq & AddClassResp |
| 编辑班级 | /editClass | POST | class_manage.proto --> EditClassReq & EditClassResp |
| 班级成员管理 | /memberManage | POST | class_manage.proto --> MemberManageReq & MemberManageResp |


## 排行榜模块


| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 排行榜信息 | /rank  | POST |  rank.proto --> RankListReq & RankListResp |
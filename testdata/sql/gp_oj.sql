DROP database IF EXISTS gp_oj;
CREATE DATABASE IF NOT EXISTS gp_oj;
USE gp_oj;

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(20) NOT NULL auto_increment,
  `account` varchar(50) NOT NULL UNIQUE,  -- 用户登陆的账号密码
  `password` varchar(100) NOT NULL,
  `role` tinyint(4) NOT NULL DEFAULT 0,  -- 0 ： 学生,
  `name` varchar(50) NOT NULL,
  `sex` boolean NOT NULL DEFAULT 0,
  
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `school` varchar(50) DEFAULT NULL,      
  
  `create` int(64) NOT NULL,        -- 创建时间 ： 时间戳
  `last_login` int(64) NOT NULL,    -- 上次登陆时间 ： 时间戳

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `announcement` (
  `id` int(20) NOT NULL auto_increment,
  `publisher` int(20) NOT NULL,  -- 发布人
  `detail` text NOT NULL,
  
  `create_time` int(64) NOT NULL,        -- 创建时间 ： 时间戳
  `disable_time` int(64) NOT NULL,       -- 失效时间 ： 时间戳

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `match` (
  `id` int(20) NOT NULL auto_increment,
  `paper_id` int(20) NOT NULL,  -- 试卷id
  `is_public` boolean NOT NULL DEFAULT 1, -- 比赛默认是公开的，所有人都可以参加

  `start_time`  int(64) NOT NULL,  -- 开始比赛时间 ： 时间戳
  `duration` int(20) NOT NULL, -- 比赛时长， 默认单位是秒

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `class` (
  `id` int(20) NOT NULL auto_increment,
  `tutor` int(20) NOT NULL,  -- 导师id

  `create_time`  int(64) NOT NULL,  -- 创建班级时间 ： 时间戳
  `announcement` json NOT NULL,   -- 班级公告

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `problem` (
  `id` int(20) NOT NULL auto_increment,
  `title` varchar(300) NOT NULL,
  `description` text NOT NULL,
  `hint` text,         -- 题目提示（可为空）
  `example` text NOT NULL,      -- 输入输出样例
  `judge_file` varchar(100) NOT NULL,
  `judge_limit` json,

  `tags` json NOT NULL,                   -- 题目分类
  `difficulty` tinyint(4) NOT NULL DEFAULT 0,
  `last_used`  int(64) NOT NULL,          -- 上次使用时间
  `used_time` int(20) NOT NULL DEFAULT 0,   -- 使用次数

  `submit_time` int(64) NOT NULL DEFAULT 0,   -- 提交总次数
  `ac` int(64) NOT NULL DEFAULT 0,            -- 通过次数
  `wa` int(64) NOT NULL DEFAULT 0,            -- 答案错误次数
  `tle` int(64) NOT NULL DEFAULT 0,           -- 超时次数
  `mle` int(64) NOT NULL DEFAULT 0,           -- 超内存次数
  `pe` int(64) NOT NULL DEFAULT 0,            -- 格式错误次数
  `ce` int(64) NOT NULL DEFAULT 0,            -- 编译次数

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `paper` (
  `id` int(20) NOT NULL auto_increment,

  `difficulty` tinyint(4) NOT NULL,  
  `knowledge_point` text NOT NULL,  -- 考察的知识点，由出题人自己填写/由程序自动生成
  -- 其他组卷所需要填写的限制也需要记录

  `problems` json NOT NULL,         -- 题目id

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_match` (
  `user_id` int(20) NOT NULL,
  `match_id` int(20) NOT NULL,

  `result` tinyint(4) NOT NULL,  
  `rank` smallint(4) NOT NULL,  
  
	PRIMARY KEY(`user_id`, `match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_class` (
  `user_id` int(20) NOT NULL,
  `class_id` int(20) NOT NULL,

  `announcement` blob DEFAULT NULL,  
  
	PRIMARY KEY(`user_id`, `class_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_problem` (
  `id` int(20) NOT NULL auto_increment,
  `user_id` int(20) NOT NULL,
  `problem_id` int(20) NOT NULL, 
  `submit_time` int(64) NOT NULL,               -- 提交时间

  `isPass` boolean NOT NULL,
  `running_langurage` tinyint(4) NOT NULL,      -- 运行语言
  `running_time` int(20),                       -- 运行时间 
  `running_mem`  int(20),                       -- 运行内存

  `code` text NOT NULL,  -- 提交的代码
  
	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
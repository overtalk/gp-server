DROP database IF EXISTS gp_oj;
CREATE DATABASE IF NOT EXISTS gp_oj;
USE gp_oj;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(64) NOT NULL auto_increment,
  `account` varchar(50) NOT NULL UNIQUE,  -- 用户登陆的账号密码
  `password` varchar(100) NOT NULL,
  `role` tinyint(4) NOT NULL DEFAULT 0,  -- 0 ： 学生,
  `name` varchar(50) NOT NULL,
  `sex` boolean NOT NULL DEFAULT 0,
  
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `school` varchar(50) DEFAULT NULL,      
  
  `create` bigint(64) NOT NULL,        -- 创建时间 ： 时间戳
  `last_login` bigint(64) NOT NULL,    -- 上次登陆时间 ： 时间戳

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 比赛表
CREATE TABLE IF NOT EXISTS `match` (
  `id` bigint(64) NOT NULL auto_increment,
  `paper_id` bigint(64) NOT NULL,  -- 试卷id
  `is_public` boolean NOT NULL DEFAULT 1, -- 比赛默认是公开的，所有人都可以参加

  `start_time` bigint(64) NOT NULL,  -- 开始比赛时间 ： 时间戳
  `duration` int(20) NOT NULL, -- 比赛时长， 默认单位是秒

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 班级表
CREATE TABLE IF NOT EXISTS `class` (
  `id` bigint(64) NOT NULL auto_increment,
  `tutor` bigint(64) NOT NULL,  -- 导师id

  `name` text NOT NULL,   -- 班级名称
  `introduction` text,    -- 班级简介
  `number` int(11) NOT NULL DEFAULT 0,
  `is_check`  boolean NOT NULL DEFAULT false,  -- 加入班级设置：false（无需审核，运行任何人进入），true（需要导师审核）
  `create_time`  bigint(64) NOT NULL,  -- 创建班级时间 ： 时间戳

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 全局公告表
CREATE TABLE IF NOT EXISTS `announcement` (
  `id` bigint(64) NOT NULL auto_increment,
  `publisher` bigint(64) NOT NULL,  -- 发布人
  `title` text NOT NULL,
  `detail` text NOT NULL,
  `class_id` bigint(64),          -- 如果是班级公告，则填写班级id
  
  `create_time` bigint(64) NOT NULL,        -- 创建时间 ： 时间戳
  `last_update_time` bigint(64) NOT NULL,   -- 上次更新时间 ： 时间戳

	PRIMARY KEY(`id`),
  foreign key(`class_id`) references `class`(`id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 题目表
CREATE TABLE IF NOT EXISTS `problem` (
  `id` bigint(64) NOT NULL auto_increment,
  `title` text NOT NULL,
  `description` text NOT NULL,
  `in_description` text NOT NULL,
  `out_description` text NOT NULL,
  `hint` text,         -- 题目提示（可为空）
  `judge_limit_time` int(11) NOT NULL,    -- 时间限制
  `judge_limit_mem` int(11) NOT NULL,     -- 内存限制

  `difficulty` tinyint(4) NOT NULL DEFAULT 0,
  `last_used`  bigint(64) NOT NULL DEFAULT 0,  -- 上次使用时间
  `used_time` int(64) NOT NULL DEFAULT 0,   -- 使用次数

  `submit_time` int(64) NOT NULL DEFAULT 0,   -- 提交总次数
  `ac` int(64) NOT NULL DEFAULT 0,            -- 通过次数
  `wa` int(64) NOT NULL DEFAULT 0,            -- 答案错误次数
  `tle` int(64) NOT NULL DEFAULT 0,           -- 超时次数
  `mle` int(64) NOT NULL DEFAULT 0,           -- 超内存次数
  `pe` int(64) NOT NULL DEFAULT 0,            -- 格式错误次数
  `ce` int(64) NOT NULL DEFAULT 0,            -- 编译次数

  `judge_file` varchar(100) NOT NULL,         -- 判题目文件的路径

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 试卷表
CREATE TABLE IF NOT EXISTS `paper` (
  `id` bigint(64) NOT NULL auto_increment,

  `difficulty` tinyint(4) NOT NULL,  
  `knowledge_point` text NOT NULL,  -- 考察的知识点，由出题人自己填写/由程序自动生成
  -- 其他组卷所需要填写的限制也需要记录

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- tag 标签表
CREATE TABLE IF NOT EXISTS `tag` (
  `id` int(11) NOT NULL auto_increment,
  `detail` varchar(100) NOT NULL,
  
	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_match` (
  `user_id` bigint(64) NOT NULL,
  `match_id` bigint(64) NOT NULL,

  `result` tinyint(4) NOT NULL,  
  `rank` smallint(4) NOT NULL,  
  
	PRIMARY KEY(`user_id`, `match_id`),
  foreign key(`user_id`) references `user`(`id`),
  foreign key(`match_id`) references `match`(`id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_class` (
  `user_id` bigint(64) NOT NULL,
  `class_id` bigint(64) NOT NULL,

  `is_checked` boolean NOT NULL DEFAULT false,          -- 是否被管理员审核通过
  `is_administrator` boolean NOT NULL DEFAULT false,    -- 是否为管理员
  
	PRIMARY KEY(`user_id`, `class_id`),
  foreign key(`user_id`) references `user`(`id`),
  foreign key(`class_id`) references `class`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_problem` (
  `id` bigint(64) NOT NULL auto_increment,
  `user_id` bigint(64) NOT NULL,
  `problem_id` bigint(64) NOT NULL, 
  `submit_time` int(64) NOT NULL,               -- 提交时间

  `isPass` boolean NOT NULL,
  `running_langurage` tinyint(4) NOT NULL,      -- 运行语言
  `running_time` int(64),                       -- 运行时间 
  `running_mem`  int(64),                       -- 运行内存

  `code` text NOT NULL,  -- 提交的代码
  
	PRIMARY KEY(`id`),
  foreign key(`user_id`) references `user`(`id`),
  foreign key(`problem_id`) references `problem`(`id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 输入输出样例表
CREATE TABLE IF NOT EXISTS `test_data` (
  `id` bigint(64) NOT NULL auto_increment,
  `problem_id` bigint(64) NOT NULL, 
  `in` text,
  `out` text,
  
	PRIMARY KEY(`id`),
  foreign key(`problem_id`) references `problem`(`id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- problem_tag 题目标签对应表
CREATE TABLE IF NOT EXISTS `problem_tag` (
  `id` bigint(64) NOT NULL auto_increment,
  `problem_id` bigint(64) NOT NULL, 
  `tag_id` int(11) NOT NULL,

	PRIMARY KEY(`id`),
  foreign key(`problem_id`) references `problem`(`id`),
  foreign key(`tag_id`) references `tag`(`id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- paper_problem 试卷题目
CREATE TABLE IF NOT EXISTS `paper_problem` (
  `problem_id` bigint(64) NOT NULL, 
  `paper_id` bigint(64) NOT NULL, 
  `index` int(11) NOT NULL,

	PRIMARY KEY(`problem_id`,`paper_id`,`index`),
  foreign key(`problem_id`) references `problem`(`id`),
  foreign key(`paper_id`) references `paper`(`id`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
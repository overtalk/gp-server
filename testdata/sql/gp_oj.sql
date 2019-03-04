DROP database IF EXISTS gp_oj;
CREATE DATABASE IF NOT EXISTS gp_oj;
USE gp_oj;

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(20) NOT NULL auto_increment,
  `auth_code` varchar(50) NOT NULL UNIQUE,
  `operation_auth` tinyint(4) NOT NULL DEFAULT 0,
  `role` tinyint(4) NOT NULL DEFAULT 0, -- 0 ： 学生

  `name` varchar(50) NOT NULL,
  `sex` boolean NOT NULL,
  `email` varchar(50) NOT NULL,
  `academy` varchar(50) NOT NULL, -- 暂定string， 以后可改成枚举
  `major` varchar(50) NOT NULL,  -- 暂定string， 以后可改成枚举
  
  `last_login` int(64) NOT NULL ,   -- 上次登陆时间 ： 时间戳

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `match` (
  `id` int(20) NOT NULL auto_increment,
  `paper_id` int(20) NOT NULL,  -- 试卷id
  `invitation_code` varchar(50) NOT NULL UNIQUE,
  `is_public` boolean NOT NULL DEFAULT 1, -- 比赛默认是公开的，所有人都可以参加

  `start_time` int(64) NOT NULL,  -- 开始比赛时间 ： 时间戳
  `duration` int(20) NOT NULL, -- 比赛时长， 默认单位是秒

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `class` (
  `id` int(20) NOT NULL auto_increment,
  `tutor` int(20) NOT NULL,  -- 导师id

  `name` int(64) NOT NULL,  
  `created_time` int(64) NOT NULL,  -- 创建班级时间 ： 时间戳

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `paper` (
  `id` int(20) NOT NULL auto_increment,

  `difficulty` tinyint(4) NOT NULL,  
  `knowledge_point` text NOT NULL,  -- 考察的知识点，由出题人自己填写/由程序自动生成

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

CREATE TABLE IF NOT EXISTS `user_oj` (
  `user_id` int(20) NOT NULL,
  `oj_id` varchar(100) NOT NULL,   -- 暂时还不知道从第三方数据中拿到的题目id是什么格式，暂定未varchar（100）

  `pass_time` int(20) DEFAULT NULL,
  `refused_time` int(20) DEFAULT NULL,  
  
	PRIMARY KEY(`user_id`, `oj_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
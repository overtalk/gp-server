DROP database IF EXISTS test_db;
CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;

CREATE TABLE IF NOT EXISTS `player` (
  `id` varchar(20) NOT NULL,
  `nickname` varchar(50) NOT NULL,
  `url` varchar(200) NOT NULL DEFAULT 'url',
  `created_at` timestamp NOT NULL ,
  `last_login` timestamp NOT NULL,
  `heartbeat_time` timestamp NOT NULL ,
  `secret` blob DEFAULT NULL,
  `last_update_strength` timestamp NOT NULL ,
  `strength` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `gold` bigint(20) NOT NULL DEFAULT 0,
  `diamond` bigint(20) NOT NULL DEFAULT 0,
  `level` int(11) NOT NULL DEFAULT 1,
  `experience` int(11) NOT NULL DEFAULT 0,
  `treasure` blob DEFAULT NULL,
  `achievement` blob DEFAULT NULL,
  `daily_task` blob DEFAULT NULL,
  `weapon` blob DEFAULT NULL,
  `weapon_equipped` blob DEFAULT NULL,
  `fashion` blob DEFAULT NULL,
  `fashion_equipped` blob DEFAULT NULL,
  `pve` blob DEFAULT NULL,
  `daily` blob DEFAULT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT 0,
  `global_mail_time` bigint(20) NOT NULL DEFAULT 0,

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

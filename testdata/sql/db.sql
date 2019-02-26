DROP database IF EXISTS test_db;
CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;

CREATE TABLE IF NOT EXISTS `test` (
  `id` int(20) NOT NULL auto_increment,
  `nickname` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL ,
  `achievement` blob DEFAULT NULL,
  `level` int(11) NOT NULL DEFAULT 1,

	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

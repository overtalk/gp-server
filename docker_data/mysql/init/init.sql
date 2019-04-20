-- MySQL dump 10.13  Distrib 8.0.12, for macos10.13 (x86_64)
--
-- Host: localhost    Database: gp_oj
-- ------------------------------------------------------
-- Server version	8.0.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP database IF EXISTS gp_oj;
CREATE DATABASE IF NOT EXISTS gp_oj;
USE gp_oj;

--
-- Table structure for table `announcement`
--

DROP TABLE IF EXISTS `announcement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `announcement` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `publisher` bigint(64) NOT NULL,
  `title` text NOT NULL,
  `detail` text NOT NULL,
  `class_id` bigint(64) DEFAULT NULL,
  `create_time` bigint(64) NOT NULL,
  `last_update_time` bigint(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `class_id` (`class_id`),
  CONSTRAINT `announcement_ibfk_1` FOREIGN KEY (`class_id`) REFERENCES `class` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `announcement`
--

LOCK TABLES `announcement` WRITE;
/*!40000 ALTER TABLE `announcement` DISABLE KEYS */;
INSERT INTO `announcement` VALUES (1,1,'测试全局公告0','测试全局公告0的内容',NULL,1555404092,1555404092),(2,1,'测试全局公告1','测试全局公告1的内容',NULL,1555404092,1555404092),(3,1,'测试全局公告2','测试全局公告2的内容',NULL,1555404092,1555404092),(4,1,'测试全局公告3','测试全局公告3的内容',NULL,1555404092,1555404092),(5,1,'测试全局公告4','测试全局公告4的内容',NULL,1555404092,1555404092),(6,1,'测试全局公告5','测试全局公告5的内容',NULL,1555404092,1555404092),(7,1,'测试全局公告6','测试全局公告6的内容',NULL,1555404092,1555404092),(8,1,'测试全局公告7','测试全局公告7的内容',NULL,1555404092,1555404092),(9,1,'测试全局公告8','测试全局公告8的内容',NULL,1555404092,1555404092),(10,1,'测试全局公告9','测试全局公告9的内容',NULL,1555404092,1555404092),(11,1,'公告1011','大家进入班级请改名',1,1555404093,1555404093),(12,1,'公告101','大家进入班级请改名111',1,1555404093,1555404093),(13,1,'公告1011','大家进入班级请改名',2,1555404093,1555404093),(14,1,'公告101','大家进入班级请改名111',2,1555404093,1555404093),(15,1,'公告1011','大家进入班级请改名',3,1555404093,1555404093),(16,1,'公告101','大家进入班级请改名111',3,1555404093,1555404093),(17,1,'公告1011','大家进入班级请改名',4,1555404093,1555404093),(18,1,'公告101','大家进入班级请改名111',4,1555404093,1555404093),(19,1,'公告1011','大家进入班级请改名',5,1555404093,1555404093),(20,1,'公告101','大家进入班级请改名111',5,1555404093,1555404093),(21,1,'公告1011','大家进入班级请改名',6,1555404093,1555404093),(22,1,'公告101','大家进入班级请改名111',6,1555404093,1555404093),(23,1,'公告1011','大家进入班级请改名',7,1555404093,1555404093),(24,1,'公告101','大家进入班级请改名111',7,1555404093,1555404093),(25,1,'公告1011','大家进入班级请改名',8,1555404093,1555404093),(26,1,'公告101','大家进入班级请改名111',8,1555404093,1555404093),(27,1,'公告1011','大家进入班级请改名',9,1555404093,1555404093),(28,1,'公告101','大家进入班级请改名111',9,1555404093,1555404093),(29,1,'公告1011','大家进入班级请改名',10,1555404093,1555404093),(30,1,'公告101','大家进入班级请改名111',10,1555404093,1555404093);
/*!40000 ALTER TABLE `announcement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `class`
--

DROP TABLE IF EXISTS `class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `class` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `tutor` bigint(64) NOT NULL,
  `name` text NOT NULL,
  `introduction` text,
  `number` int(11) NOT NULL DEFAULT '0',
  `is_check` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` bigint(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `class`
--

LOCK TABLES `class` WRITE;
/*!40000 ALTER TABLE `class` DISABLE KEYS */;
INSERT INTO `class` VALUES (1,1,'测试班级','这个一个测试的班级',0,0,1555404093),(2,1,'测试班级','这个一个测试的班级',0,0,1555404093),(3,1,'测试班级','这个一个测试的班级',0,0,1555404093),(4,1,'测试班级','这个一个测试的班级',0,0,1555404093),(5,1,'测试班级','这个一个测试的班级',0,0,1555404093),(6,1,'测试班级','这个一个测试的班级',0,0,1555404093),(7,1,'测试班级','这个一个测试的班级',0,0,1555404093),(8,1,'测试班级','这个一个测试的班级',0,0,1555404093),(9,1,'测试班级','这个一个测试的班级',0,0,1555404093),(10,1,'测试班级','这个一个测试的班级',0,0,1555404093);
/*!40000 ALTER TABLE `class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `difficulty`
--

DROP TABLE IF EXISTS `difficulty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `difficulty` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `detail` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `difficulty`
--

LOCK TABLES `difficulty` WRITE;
/*!40000 ALTER TABLE `difficulty` DISABLE KEYS */;
INSERT INTO `difficulty` VALUES (1,'很简单'),(2,'不那么简单'),(3,'中等'),(4,'有点困难'),(5,'很困难');
/*!40000 ALTER TABLE `difficulty` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `language`
--

DROP TABLE IF EXISTS `language`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `language` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `detail` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `language`
--

LOCK TABLES `language` WRITE;
/*!40000 ALTER TABLE `language` DISABLE KEYS */;
INSERT INTO `language` VALUES (1,'C'),(2,'C_PLUS'),(3,'JAVA'),(4,'PYTHON2'),(5,'PYTHON3');
/*!40000 ALTER TABLE `language` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `match`
--

DROP TABLE IF EXISTS `match`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `match` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `paper_id` bigint(64) NOT NULL,
  `is_public` tinyint(1) NOT NULL DEFAULT '1',
  `title` text NOT NULL,
  `introduction` text,
  `start_time` bigint(64) NOT NULL,
  `end_time` bigint(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `match`
--

LOCK TABLES `match` WRITE;
/*!40000 ALTER TABLE `match` DISABLE KEYS */;
INSERT INTO `match` VALUES (1,1,1,'比赛00','测试的比赛',1555404093,1555414093),(2,2,1,'比赛01','测试的比赛',1555404093,1555414093),(3,3,1,'比赛02','测试的比赛',1555404093,1555414093),(4,4,1,'比赛03','测试的比赛',1555404093,1555414093),(5,5,1,'比赛04','测试的比赛',1555404093,1555414093),(6,6,1,'比赛05','测试的比赛',1555404093,1555414093),(7,7,1,'比赛06','测试的比赛',1555404093,1555414093),(8,8,1,'比赛07','测试的比赛',1555404093,1555414093),(9,9,1,'比赛08','测试的比赛',1555404093,1555414093),(10,10,1,'比赛09','测试的比赛',1555404093,1555414093),(11,11,0,'呵呵呵考试','asdfasdf',1555423598,1555433598);
/*!40000 ALTER TABLE `match` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paper`
--

DROP TABLE IF EXISTS `paper`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `paper` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `difficulty` int(11) NOT NULL,
  `problem_num` int(11) NOT NULL,
  `tags` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paper`
--

LOCK TABLES `paper` WRITE;
/*!40000 ALTER TABLE `paper` DISABLE KEYS */;
INSERT INTO `paper` VALUES (1,1,3,'[1,2]'),(2,1,3,'[1,2]'),(3,1,3,'[1,2]'),(4,1,3,'[1,2]'),(5,1,3,'[1,2]'),(6,1,3,'[1,2]'),(7,1,3,'[1,2]'),(8,1,3,'[1,2]'),(9,1,3,'[1,2]'),(10,1,3,'[1,2]'),(11,1,3,'[1,2,3]');
/*!40000 ALTER TABLE `paper` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paper_problem`
--

DROP TABLE IF EXISTS `paper_problem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `paper_problem` (
  `problem_id` bigint(64) NOT NULL,
  `paper_id` bigint(64) NOT NULL,
  `index` int(11) NOT NULL,
  PRIMARY KEY (`problem_id`,`paper_id`,`index`),
  KEY `paper_id` (`paper_id`),
  CONSTRAINT `paper_problem_ibfk_1` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`id`),
  CONSTRAINT `paper_problem_ibfk_2` FOREIGN KEY (`paper_id`) REFERENCES `paper` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paper_problem`
--

LOCK TABLES `paper_problem` WRITE;
/*!40000 ALTER TABLE `paper_problem` DISABLE KEYS */;
INSERT INTO `paper_problem` VALUES (1,1,1),(2,1,2),(4,1,3),(1,2,1),(2,2,2),(4,2,3),(1,3,1),(2,3,2),(4,3,3),(1,4,1),(2,4,2),(4,4,3),(1,5,1),(2,5,2),(4,5,3),(1,6,1),(2,6,2),(4,6,3),(1,7,1),(2,7,2),(4,7,3),(1,8,1),(2,8,2),(4,8,3),(1,9,1),(2,9,2),(4,9,3),(1,10,1),(2,10,2),(4,10,3),(1,11,1),(1,11,2);
/*!40000 ALTER TABLE `paper_problem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problem`
--

DROP TABLE IF EXISTS `problem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `problem` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  `description` text NOT NULL,
  `in_description` text NOT NULL,
  `out_description` text NOT NULL,
  `hint` text,
  `judge_limit_time` int(11) NOT NULL,
  `judge_limit_mem` int(11) NOT NULL,
  `tags` text NOT NULL,
  `difficulty` int(11) NOT NULL DEFAULT '1',
  `create_time` bigint(64) NOT NULL DEFAULT '0',
  `publisher` bigint(64) NOT NULL,
  `last_used` bigint(64) NOT NULL DEFAULT '0',
  `used_time` int(64) NOT NULL DEFAULT '0',
  `submit_time` int(64) NOT NULL DEFAULT '0',
  `ac` int(64) NOT NULL DEFAULT '0',
  `wa` int(64) NOT NULL DEFAULT '0',
  `tle` int(64) NOT NULL DEFAULT '0',
  `mle` int(64) NOT NULL DEFAULT '0',
  `pe` int(64) NOT NULL DEFAULT '0',
  `ce` int(64) NOT NULL DEFAULT '0',
  `judge_file` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `publisher` (`publisher`),
  KEY `difficulty` (`difficulty`),
  CONSTRAINT `problem_ibfk_1` FOREIGN KEY (`publisher`) REFERENCES `user` (`id`),
  CONSTRAINT `problem_ibfk_2` FOREIGN KEY (`difficulty`) REFERENCES `difficulty` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problem`
--

LOCK TABLES `problem` WRITE;
/*!40000 ALTER TABLE `problem` DISABLE KEYS */;
INSERT INTO `problem` VALUES (1,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(2,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(3,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(4,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(5,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(6,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(7,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(8,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(9,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404092,1,1555404092,0,0,0,0,0,0,0,0,''),(10,'求和问题','求两个数的和','输入两个int型数','输出两个数的和','无提示',1000,134217728,'[1,2,3]',1,1555404093,1,1555404093,0,0,0,0,0,0,0,0,'');
/*!40000 ALTER TABLE `problem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `detail` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'学生'),(2,'老师'),(3,'管理员');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `detail` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag`
--

LOCK TABLES `tag` WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` VALUES (1,'tag0'),(2,'tag1'),(3,'tag2'),(4,'tag3'),(5,'tag4'),(6,'tag5'),(7,'tag6'),(8,'tag7'),(9,'tag8'),(10,'tag9');
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_data`
--

DROP TABLE IF EXISTS `test_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_data` (
  `problem_id` bigint(64) NOT NULL,
  `index` bigint(64) NOT NULL,
  `in` text,
  `out` text,
  PRIMARY KEY (`index`,`problem_id`),
  KEY `problem_id` (`problem_id`),
  CONSTRAINT `test_data_ibfk_1` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_data`
--

LOCK TABLES `test_data` WRITE;
/*!40000 ALTER TABLE `test_data` DISABLE KEYS */;
INSERT INTO `test_data` VALUES (1,1,'in 1','out 1'),(2,1,'in 1','out 1'),(3,1,'in 1','out 1'),(4,1,'in 1','out 1'),(5,1,'in 1','out 1'),(6,1,'in 1','out 1'),(7,1,'in 1','out 1'),(8,1,'in 1','out 1'),(9,1,'in 1','out 1'),(10,1,'in 1','out 1'),(1,2,'in 2','out 2'),(2,2,'in 2','out 2'),(3,2,'in 2','out 2'),(4,2,'in 2','out 2'),(5,2,'in 2','out 2'),(6,2,'in 2','out 2'),(7,2,'in 2','out 2'),(8,2,'in 2','out 2'),(9,2,'in 2','out 2'),(10,2,'in 2','out 2');
/*!40000 ALTER TABLE `test_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `role` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT '0',
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `school` varchar(50) DEFAULT NULL,
  `create` bigint(64) NOT NULL,
  `last_login` bigint(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`),
  KEY `role` (`role`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`role`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'tom0','a0cdca52112c51a89c328ac7ac51af64',3,'tom0',1,'tom0','tom0','tom0',1555404092,1555404092),(2,'tom1','5a041fd60714e48104e86ef43250c2a3',1,'tom1',1,'tom1','tom1','tom1',1555404092,1555404092),(3,'tom2','28d3047f0ed0c91841069c34863bdfab',2,'tom2',0,'tom2','tom2','tom2',1555404092,1555404092),(4,'tom3','8acc18b1b85fef612a8ce6ed7369a5bb',3,'tom3',1,'tom3','tom3','tom3',1555404092,1555404092),(5,'tom4','125eed4f307bbdee2378cf2d5e6c5edd',3,'tom4',1,'tom4','tom4','tom4',1555404092,1555404092),(6,'tom5','4bb1073250c20742f156db5d7a3b1ac3',1,'tom5',0,'tom5','tom5','tom5',1555404092,1555404092),(7,'tom6','8296f1180854286ce14c19c56ab076b9',2,'tom6',0,'tom6','tom6','tom6',1555404092,1555404092),(8,'tom7','8f9be3a0410c8d8e6a98061d3c5a2c51',1,'tom7',0,'tom7','tom7','tom7',1555404092,1555404092),(9,'tom8','3e954ed96119dcbdc6d869ffe674820b',1,'tom8',0,'tom8','tom8','tom8',1555404092,1555404092),(10,'tom9','8e26ad372c98acdde6e76ae5889e41b2',2,'tom9',1,'tom9','tom9','tom9',1555404092,1555404092),(11,'1','c4ca4238a0b923820dcc509a6f75849b',2,'lyx',0,'111','xxx@xxx','shu',1555404092,1555404092);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_class`
--

DROP TABLE IF EXISTS `user_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_class` (
  `user_id` bigint(64) NOT NULL,
  `class_id` bigint(64) NOT NULL,
  `is_checked` tinyint(1) NOT NULL DEFAULT '0',
  `is_administrator` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`class_id`),
  KEY `class_id` (`class_id`),
  CONSTRAINT `user_class_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_class_ibfk_2` FOREIGN KEY (`class_id`) REFERENCES `class` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_class`
--

LOCK TABLES `user_class` WRITE;
/*!40000 ALTER TABLE `user_class` DISABLE KEYS */;
INSERT INTO `user_class` VALUES (1,1,1,0),(1,2,1,0),(1,3,1,0),(1,4,1,0),(1,5,1,0),(1,6,1,0),(1,7,1,0),(1,8,1,0),(1,9,1,0),(1,10,1,0),(2,1,1,0),(2,2,1,0),(2,3,1,0),(2,4,1,0),(2,5,1,0),(2,6,1,0),(2,7,1,0),(2,8,1,0),(2,9,1,0),(2,10,1,0),(3,1,1,0),(3,2,1,0),(3,3,1,0),(3,4,1,0),(3,5,1,0),(3,6,1,0),(3,7,1,0),(3,8,1,0),(3,9,1,0),(3,10,1,0),(4,1,1,0),(4,2,1,0),(4,3,1,0),(4,4,1,0),(4,5,1,0),(4,6,1,0),(4,7,1,0),(4,8,1,0),(4,9,1,0),(4,10,1,0),(5,1,1,0),(5,2,1,0),(5,3,1,0),(5,4,1,0),(5,5,1,0),(5,6,1,0),(5,7,1,0),(5,8,1,0),(5,9,1,0),(5,10,1,0),(6,1,1,0),(6,2,1,0),(6,3,1,0),(6,4,1,0),(6,5,1,0),(6,6,1,0),(6,7,1,0),(6,8,1,0),(6,9,1,0),(6,10,1,0),(7,1,1,0),(7,2,1,0),(7,3,1,0),(7,4,1,0),(7,5,1,0),(7,6,1,0),(7,7,1,0),(7,8,1,0),(7,9,1,0),(7,10,1,0),(8,1,1,0),(8,2,1,0),(8,3,1,0),(8,4,1,0),(8,5,1,0),(8,6,1,0),(8,7,1,0),(8,8,1,0),(8,9,1,0),(8,10,1,0);
/*!40000 ALTER TABLE `user_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_match`
--

DROP TABLE IF EXISTS `user_match`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_match` (
  `user_id` bigint(64) NOT NULL,
  `match_id` bigint(64) NOT NULL,
  `result` tinyint(4) NOT NULL,
  `rank` smallint(4) NOT NULL,
  PRIMARY KEY (`user_id`,`match_id`),
  KEY `match_id` (`match_id`),
  CONSTRAINT `user_match_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_match_ibfk_2` FOREIGN KEY (`match_id`) REFERENCES `match` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_match`
--

LOCK TABLES `user_match` WRITE;
/*!40000 ALTER TABLE `user_match` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_match` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_problem`
--

DROP TABLE IF EXISTS `user_problem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_problem` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(64) NOT NULL,
  `problem_id` bigint(64) NOT NULL,
  `submit_time` int(64) NOT NULL,
  `isPass` tinyint(1) NOT NULL,
  `running_language` int(11) NOT NULL,
  `running_time` int(64) DEFAULT NULL,
  `running_mem` int(64) DEFAULT NULL,
  `code` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `problem_id` (`problem_id`),
  KEY `running_language` (`running_language`),
  CONSTRAINT `user_problem_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_problem_ibfk_2` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`id`),
  CONSTRAINT `user_problem_ibfk_3` FOREIGN KEY (`running_language`) REFERENCES `language` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=301 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_problem`
--

LOCK TABLES `user_problem` WRITE;
/*!40000 ALTER TABLE `user_problem` DISABLE KEYS */;
INSERT INTO `user_problem` VALUES (1,4,4,1555404093,1,2,1,10,'木有代码哦'),(2,6,6,1555404093,0,2,1,10,'木有代码哦'),(3,7,7,1555404093,1,1,1,10,'木有代码哦'),(4,3,3,1555404093,1,3,1,10,'木有代码哦'),(5,7,7,1555404093,0,4,1,10,'木有代码哦'),(6,2,2,1555404093,1,1,1,10,'木有代码哦'),(7,6,6,1555404093,1,1,1,10,'木有代码哦'),(8,8,8,1555404093,1,4,1,10,'木有代码哦'),(9,6,6,1555404093,0,4,1,10,'木有代码哦'),(10,1,1,1555404093,0,2,1,10,'木有代码哦'),(11,9,9,1555404093,1,2,1,10,'木有代码哦'),(12,8,8,1555404093,1,1,1,10,'木有代码哦'),(13,5,5,1555404093,0,4,1,10,'木有代码哦'),(14,6,6,1555404093,0,1,1,10,'木有代码哦'),(15,4,4,1555404093,0,2,1,10,'木有代码哦'),(16,4,4,1555404093,0,4,1,10,'木有代码哦'),(17,10,10,1555404093,0,1,1,10,'木有代码哦'),(18,8,8,1555404093,1,1,1,10,'木有代码哦'),(19,8,8,1555404093,1,2,1,10,'木有代码哦'),(20,9,9,1555404093,0,3,1,10,'木有代码哦'),(21,10,10,1555404093,0,2,1,10,'木有代码哦'),(22,5,5,1555404093,1,1,1,10,'木有代码哦'),(23,2,2,1555404093,0,3,1,10,'木有代码哦'),(24,9,9,1555404093,1,4,1,10,'木有代码哦'),(25,10,10,1555404093,1,2,1,10,'木有代码哦'),(26,3,3,1555404093,0,2,1,10,'木有代码哦'),(27,7,7,1555404093,0,4,1,10,'木有代码哦'),(28,3,3,1555404093,0,4,1,10,'木有代码哦'),(29,4,4,1555404093,0,1,1,10,'木有代码哦'),(30,3,3,1555404093,1,1,1,10,'木有代码哦'),(31,7,7,1555404093,1,3,1,10,'木有代码哦'),(32,6,6,1555404093,1,4,1,10,'木有代码哦'),(33,7,7,1555404093,1,2,1,10,'木有代码哦'),(34,6,6,1555404093,0,2,1,10,'木有代码哦'),(35,2,2,1555404093,1,2,1,10,'木有代码哦'),(36,4,4,1555404093,1,3,1,10,'木有代码哦'),(37,6,6,1555404093,1,4,1,10,'木有代码哦'),(38,5,5,1555404093,1,1,1,10,'木有代码哦'),(39,6,6,1555404093,1,3,1,10,'木有代码哦'),(40,5,5,1555404093,0,2,1,10,'木有代码哦'),(41,7,7,1555404093,0,1,1,10,'木有代码哦'),(42,3,3,1555404093,1,2,1,10,'木有代码哦'),(43,9,9,1555404093,1,2,1,10,'木有代码哦'),(44,10,10,1555404093,0,3,1,10,'木有代码哦'),(45,9,9,1555404093,0,2,1,10,'木有代码哦'),(46,6,6,1555404093,1,3,1,10,'木有代码哦'),(47,1,1,1555404093,0,4,1,10,'木有代码哦'),(48,8,8,1555404093,1,1,1,10,'木有代码哦'),(49,10,10,1555404093,1,1,1,10,'木有代码哦'),(50,3,3,1555404093,0,4,1,10,'木有代码哦'),(51,5,5,1555404093,1,2,1,10,'木有代码哦'),(52,9,9,1555404093,1,1,1,10,'木有代码哦'),(53,1,1,1555404093,0,2,1,10,'木有代码哦'),(54,4,4,1555404093,0,4,1,10,'木有代码哦'),(55,7,7,1555404093,1,4,1,10,'木有代码哦'),(56,6,6,1555404093,0,2,1,10,'木有代码哦'),(57,4,4,1555404093,0,2,1,10,'木有代码哦'),(58,8,8,1555404093,0,2,1,10,'木有代码哦'),(59,9,9,1555404093,1,2,1,10,'木有代码哦'),(60,3,3,1555404093,0,2,1,10,'木有代码哦'),(61,4,4,1555404093,1,2,1,10,'木有代码哦'),(62,4,4,1555404093,0,3,1,10,'木有代码哦'),(63,3,3,1555404093,1,4,1,10,'木有代码哦'),(64,6,6,1555404093,0,3,1,10,'木有代码哦'),(65,1,1,1555404093,0,4,1,10,'木有代码哦'),(66,5,5,1555404093,1,4,1,10,'木有代码哦'),(67,9,9,1555404093,1,4,1,10,'木有代码哦'),(68,5,5,1555404093,0,3,1,10,'木有代码哦'),(69,3,3,1555404093,1,1,1,10,'木有代码哦'),(70,2,2,1555404093,1,1,1,10,'木有代码哦'),(71,3,3,1555404093,0,2,1,10,'木有代码哦'),(72,6,6,1555404093,1,2,1,10,'木有代码哦'),(73,4,4,1555404093,1,1,1,10,'木有代码哦'),(74,4,4,1555404093,1,2,1,10,'木有代码哦'),(75,6,6,1555404093,1,1,1,10,'木有代码哦'),(76,1,1,1555404093,0,2,1,10,'木有代码哦'),(77,7,7,1555404093,1,3,1,10,'木有代码哦'),(78,3,3,1555404093,0,4,1,10,'木有代码哦'),(79,9,9,1555404093,0,4,1,10,'木有代码哦'),(80,10,10,1555404093,0,3,1,10,'木有代码哦'),(81,8,8,1555404093,0,4,1,10,'木有代码哦'),(82,5,5,1555404093,0,4,1,10,'木有代码哦'),(83,6,6,1555404093,1,1,1,10,'木有代码哦'),(84,6,6,1555404093,0,2,1,10,'木有代码哦'),(85,9,9,1555404093,1,1,1,10,'木有代码哦'),(86,3,3,1555404093,1,1,1,10,'木有代码哦'),(87,7,7,1555404093,1,4,1,10,'木有代码哦'),(88,3,3,1555404093,0,3,1,10,'木有代码哦'),(89,10,10,1555404093,1,3,1,10,'木有代码哦'),(90,3,3,1555404093,0,2,1,10,'木有代码哦'),(91,3,3,1555404093,0,2,1,10,'木有代码哦'),(92,7,7,1555404093,1,2,1,10,'木有代码哦'),(93,2,2,1555404093,1,3,1,10,'木有代码哦'),(94,10,10,1555404093,1,4,1,10,'木有代码哦'),(95,9,9,1555404093,1,3,1,10,'木有代码哦'),(96,3,3,1555404093,0,4,1,10,'木有代码哦'),(97,2,2,1555404093,0,3,1,10,'木有代码哦'),(98,8,8,1555404093,0,4,1,10,'木有代码哦'),(99,2,2,1555404093,1,3,1,10,'木有代码哦'),(100,5,5,1555404093,0,2,1,10,'木有代码哦'),(101,10,10,1555404093,0,3,1,10,'木有代码哦'),(102,10,10,1555404093,1,3,1,10,'木有代码哦'),(103,4,4,1555404093,1,4,1,10,'木有代码哦'),(104,10,10,1555404093,1,3,1,10,'木有代码哦'),(105,3,3,1555404093,0,4,1,10,'木有代码哦'),(106,5,5,1555404093,1,4,1,10,'木有代码哦'),(107,10,10,1555404093,0,1,1,10,'木有代码哦'),(108,4,4,1555404093,0,3,1,10,'木有代码哦'),(109,4,4,1555404093,0,4,1,10,'木有代码哦'),(110,5,5,1555404093,1,4,1,10,'木有代码哦'),(111,10,10,1555404093,1,1,1,10,'木有代码哦'),(112,8,8,1555404093,0,2,1,10,'木有代码哦'),(113,5,5,1555404093,0,2,1,10,'木有代码哦'),(114,4,4,1555404093,1,3,1,10,'木有代码哦'),(115,6,6,1555404093,1,4,1,10,'木有代码哦'),(116,4,4,1555404093,0,4,1,10,'木有代码哦'),(117,6,6,1555404093,0,1,1,10,'木有代码哦'),(118,5,5,1555404093,1,1,1,10,'木有代码哦'),(119,10,10,1555404093,0,4,1,10,'木有代码哦'),(120,4,4,1555404093,0,3,1,10,'木有代码哦'),(121,6,6,1555404093,1,1,1,10,'木有代码哦'),(122,2,2,1555404093,1,1,1,10,'木有代码哦'),(123,2,2,1555404093,1,1,1,10,'木有代码哦'),(124,2,2,1555404093,1,1,1,10,'木有代码哦'),(125,2,2,1555404093,1,3,1,10,'木有代码哦'),(126,8,8,1555404093,0,2,1,10,'木有代码哦'),(127,4,4,1555404093,0,1,1,10,'木有代码哦'),(128,2,2,1555404093,0,3,1,10,'木有代码哦'),(129,1,1,1555404093,1,4,1,10,'木有代码哦'),(130,9,9,1555404093,0,2,1,10,'木有代码哦'),(131,9,9,1555404093,1,4,1,10,'木有代码哦'),(132,8,8,1555404093,0,4,1,10,'木有代码哦'),(133,9,9,1555404093,1,3,1,10,'木有代码哦'),(134,5,5,1555404093,1,4,1,10,'木有代码哦'),(135,7,7,1555404093,0,4,1,10,'木有代码哦'),(136,1,1,1555404093,1,1,1,10,'木有代码哦'),(137,6,6,1555404093,1,1,1,10,'木有代码哦'),(138,5,5,1555404093,0,4,1,10,'木有代码哦'),(139,2,2,1555404093,0,1,1,10,'木有代码哦'),(140,2,2,1555404093,0,1,1,10,'木有代码哦'),(141,1,1,1555404093,1,1,1,10,'木有代码哦'),(142,2,2,1555404093,1,1,1,10,'木有代码哦'),(143,8,8,1555404093,0,1,1,10,'木有代码哦'),(144,5,5,1555404093,1,2,1,10,'木有代码哦'),(145,1,1,1555404093,0,3,1,10,'木有代码哦'),(146,7,7,1555404093,0,2,1,10,'木有代码哦'),(147,4,4,1555404093,1,3,1,10,'木有代码哦'),(148,7,7,1555404093,0,2,1,10,'木有代码哦'),(149,8,8,1555404093,1,1,1,10,'木有代码哦'),(150,2,2,1555404093,1,3,1,10,'木有代码哦'),(151,1,1,1555404093,1,3,1,10,'木有代码哦'),(152,9,9,1555404093,0,2,1,10,'木有代码哦'),(153,5,5,1555404093,0,4,1,10,'木有代码哦'),(154,10,10,1555404093,1,1,1,10,'木有代码哦'),(155,5,5,1555404093,1,1,1,10,'木有代码哦'),(156,7,7,1555404093,0,1,1,10,'木有代码哦'),(157,5,5,1555404093,0,2,1,10,'木有代码哦'),(158,9,9,1555404093,0,4,1,10,'木有代码哦'),(159,9,9,1555404093,0,2,1,10,'木有代码哦'),(160,6,6,1555404093,1,3,1,10,'木有代码哦'),(161,2,2,1555404093,0,3,1,10,'木有代码哦'),(162,4,4,1555404093,1,3,1,10,'木有代码哦'),(163,5,5,1555404093,0,1,1,10,'木有代码哦'),(164,8,8,1555404093,1,1,1,10,'木有代码哦'),(165,4,4,1555404093,1,1,1,10,'木有代码哦'),(166,7,7,1555404093,0,2,1,10,'木有代码哦'),(167,7,7,1555404093,0,4,1,10,'木有代码哦'),(168,3,3,1555404093,1,1,1,10,'木有代码哦'),(169,5,5,1555404093,0,2,1,10,'木有代码哦'),(170,7,7,1555404093,0,1,1,10,'木有代码哦'),(171,3,3,1555404093,0,3,1,10,'木有代码哦'),(172,5,5,1555404093,1,4,1,10,'木有代码哦'),(173,8,8,1555404093,0,4,1,10,'木有代码哦'),(174,4,4,1555404093,1,3,1,10,'木有代码哦'),(175,3,3,1555404093,0,2,1,10,'木有代码哦'),(176,8,8,1555404093,0,1,1,10,'木有代码哦'),(177,5,5,1555404093,0,2,1,10,'木有代码哦'),(178,1,1,1555404093,0,2,1,10,'木有代码哦'),(179,1,1,1555404093,0,1,1,10,'木有代码哦'),(180,5,5,1555404093,0,2,1,10,'木有代码哦'),(181,2,2,1555404093,0,2,1,10,'木有代码哦'),(182,5,5,1555404093,1,4,1,10,'木有代码哦'),(183,8,8,1555404093,0,1,1,10,'木有代码哦'),(184,5,5,1555404093,0,3,1,10,'木有代码哦'),(185,4,4,1555404093,0,1,1,10,'木有代码哦'),(186,5,5,1555404093,0,2,1,10,'木有代码哦'),(187,10,10,1555404093,1,3,1,10,'木有代码哦'),(188,3,3,1555404093,1,4,1,10,'木有代码哦'),(189,5,5,1555404093,1,4,1,10,'木有代码哦'),(190,8,8,1555404093,1,1,1,10,'木有代码哦'),(191,3,3,1555404093,0,1,1,10,'木有代码哦'),(192,6,6,1555404093,1,1,1,10,'木有代码哦'),(193,5,5,1555404093,1,2,1,10,'木有代码哦'),(194,10,10,1555404093,0,4,1,10,'木有代码哦'),(195,9,9,1555404093,1,2,1,10,'木有代码哦'),(196,2,2,1555404093,1,3,1,10,'木有代码哦'),(197,9,9,1555404093,0,3,1,10,'木有代码哦'),(198,2,2,1555404093,1,4,1,10,'木有代码哦'),(199,9,9,1555404093,0,2,1,10,'木有代码哦'),(200,2,2,1555404093,0,3,1,10,'木有代码哦'),(201,6,6,1555404093,1,3,1,10,'木有代码哦'),(202,9,9,1555404093,0,4,1,10,'木有代码哦'),(203,2,2,1555404093,1,4,1,10,'木有代码哦'),(204,5,5,1555404093,1,4,1,10,'木有代码哦'),(205,6,6,1555404093,1,1,1,10,'木有代码哦'),(206,6,6,1555404093,1,1,1,10,'木有代码哦'),(207,10,10,1555404093,1,1,1,10,'木有代码哦'),(208,9,9,1555404093,1,2,1,10,'木有代码哦'),(209,5,5,1555404093,1,1,1,10,'木有代码哦'),(210,2,2,1555404093,0,1,1,10,'木有代码哦'),(211,6,6,1555404093,1,3,1,10,'木有代码哦'),(212,10,10,1555404093,0,3,1,10,'木有代码哦'),(213,5,5,1555404093,1,3,1,10,'木有代码哦'),(214,8,8,1555404093,1,1,1,10,'木有代码哦'),(215,7,7,1555404093,1,4,1,10,'木有代码哦'),(216,8,8,1555404093,1,1,1,10,'木有代码哦'),(217,2,2,1555404093,1,1,1,10,'木有代码哦'),(218,7,7,1555404093,0,4,1,10,'木有代码哦'),(219,5,5,1555404093,1,1,1,10,'木有代码哦'),(220,6,6,1555404093,1,2,1,10,'木有代码哦'),(221,3,3,1555404093,1,1,1,10,'木有代码哦'),(222,3,3,1555404093,0,2,1,10,'木有代码哦'),(223,1,1,1555404093,0,3,1,10,'木有代码哦'),(224,1,1,1555404093,1,2,1,10,'木有代码哦'),(225,5,5,1555404093,1,2,1,10,'木有代码哦'),(226,4,4,1555404093,1,2,1,10,'木有代码哦'),(227,10,10,1555404093,1,1,1,10,'木有代码哦'),(228,4,4,1555404093,1,4,1,10,'木有代码哦'),(229,5,5,1555404093,0,3,1,10,'木有代码哦'),(230,6,6,1555404093,0,3,1,10,'木有代码哦'),(231,5,5,1555404093,0,2,1,10,'木有代码哦'),(232,5,5,1555404093,1,3,1,10,'木有代码哦'),(233,5,5,1555404093,0,1,1,10,'木有代码哦'),(234,10,10,1555404093,1,3,1,10,'木有代码哦'),(235,9,9,1555404093,1,4,1,10,'木有代码哦'),(236,9,9,1555404093,0,4,1,10,'木有代码哦'),(237,10,10,1555404093,0,1,1,10,'木有代码哦'),(238,9,9,1555404093,1,4,1,10,'木有代码哦'),(239,3,3,1555404093,0,2,1,10,'木有代码哦'),(240,10,10,1555404093,1,4,1,10,'木有代码哦'),(241,5,5,1555404093,0,3,1,10,'木有代码哦'),(242,1,1,1555404093,0,4,1,10,'木有代码哦'),(243,9,9,1555404093,0,4,1,10,'木有代码哦'),(244,4,4,1555404093,1,4,1,10,'木有代码哦'),(245,5,5,1555404093,1,2,1,10,'木有代码哦'),(246,8,8,1555404093,1,3,1,10,'木有代码哦'),(247,3,3,1555404093,1,2,1,10,'木有代码哦'),(248,4,4,1555404093,0,2,1,10,'木有代码哦'),(249,9,9,1555404093,1,3,1,10,'木有代码哦'),(250,6,6,1555404093,0,4,1,10,'木有代码哦'),(251,4,4,1555404093,1,2,1,10,'木有代码哦'),(252,2,2,1555404093,0,1,1,10,'木有代码哦'),(253,4,4,1555404093,0,4,1,10,'木有代码哦'),(254,5,5,1555404093,1,4,1,10,'木有代码哦'),(255,9,9,1555404093,1,4,1,10,'木有代码哦'),(256,4,4,1555404093,0,2,1,10,'木有代码哦'),(257,9,9,1555404093,0,1,1,10,'木有代码哦'),(258,10,10,1555404093,0,3,1,10,'木有代码哦'),(259,6,6,1555404093,1,1,1,10,'木有代码哦'),(260,6,6,1555404093,1,2,1,10,'木有代码哦'),(261,1,1,1555404093,1,1,1,10,'木有代码哦'),(262,9,9,1555404093,0,2,1,10,'木有代码哦'),(263,1,1,1555404093,1,2,1,10,'木有代码哦'),(264,6,6,1555404093,1,3,1,10,'木有代码哦'),(265,7,7,1555404093,0,2,1,10,'木有代码哦'),(266,6,6,1555404093,0,3,1,10,'木有代码哦'),(267,5,5,1555404093,0,4,1,10,'木有代码哦'),(268,3,3,1555404093,0,1,1,10,'木有代码哦'),(269,2,2,1555404093,1,3,1,10,'木有代码哦'),(270,9,9,1555404093,0,4,1,10,'木有代码哦'),(271,2,2,1555404093,0,3,1,10,'木有代码哦'),(272,7,7,1555404093,0,2,1,10,'木有代码哦'),(273,6,6,1555404093,1,1,1,10,'木有代码哦'),(274,9,9,1555404093,0,4,1,10,'木有代码哦'),(275,5,5,1555404093,0,3,1,10,'木有代码哦'),(276,10,10,1555404093,1,1,1,10,'木有代码哦'),(277,2,2,1555404093,0,1,1,10,'木有代码哦'),(278,8,8,1555404093,0,2,1,10,'木有代码哦'),(279,9,9,1555404093,1,2,1,10,'木有代码哦'),(280,3,3,1555404093,0,2,1,10,'木有代码哦'),(281,6,6,1555404093,0,1,1,10,'木有代码哦'),(282,9,9,1555404093,1,2,1,10,'木有代码哦'),(283,5,5,1555404093,0,3,1,10,'木有代码哦'),(284,4,4,1555404093,1,1,1,10,'木有代码哦'),(285,3,3,1555404093,0,2,1,10,'木有代码哦'),(286,9,9,1555404093,1,2,1,10,'木有代码哦'),(287,4,4,1555404093,1,3,1,10,'木有代码哦'),(288,9,9,1555404093,1,3,1,10,'木有代码哦'),(289,7,7,1555404093,1,4,1,10,'木有代码哦'),(290,6,6,1555404093,1,4,1,10,'木有代码哦'),(291,7,7,1555404093,0,4,1,10,'木有代码哦'),(292,1,1,1555404093,0,1,1,10,'木有代码哦'),(293,7,7,1555404093,0,4,1,10,'木有代码哦'),(294,7,7,1555404093,1,2,1,10,'木有代码哦'),(295,2,2,1555404093,0,2,1,10,'木有代码哦'),(296,9,9,1555404093,0,2,1,10,'木有代码哦'),(297,7,7,1555404093,1,4,1,10,'木有代码哦'),(298,10,10,1555404093,0,3,1,10,'木有代码哦'),(299,1,1,1555404093,0,3,1,10,'木有代码哦'),(300,9,9,1555404093,1,1,1,10,'木有代码哦');
/*!40000 ALTER TABLE `user_problem` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-04-20  9:51:58

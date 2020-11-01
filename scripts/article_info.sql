CREATE DATABASE /*!32312 IF NOT EXISTS*/ `article_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

CREATE TABLE `account_tab` (
  `article_id` BIGINT NOT NULL AUTO_INCREMENT,
  `title` varchar(1024) NOT NULL,
  `author` varchar(1024) NOT NULL,
  `content` varchar(1048576) NOT NULL,
  `updateTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`article_id`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
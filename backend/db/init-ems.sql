CREATE DATABASE IF NOT EXISTS ems;
use ems;

CREATE TABLE IF NOT EXISTS device_table (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `real_device_uniq_id` bigint(20) NOT NULL DEFAULT 0,
  `last_visible_date` datetime DEFAULT NULL,
  `friendly_name` varchar(256) DEFAULT NULL,
  `ipv4` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `REAL_DEVICE_UNIQ_ID` (`real_device_uniq_id`)
) AUTO_INCREMENT=1;


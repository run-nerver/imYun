/*
 Navicat Premium Data Transfer

 Source Server         : docker本地
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost
 Source Database       : PrintYun

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : utf-8

 Date: 02/01/2021 15:16:03 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `orders`
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `printer_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `file_name` varchar(128) COLLATE utf8mb4_german2_ci NOT NULL,
  `re_name` varchar(128) COLLATE utf8mb4_german2_ci NOT NULL,
  `ds` int(8) DEFAULT NULL,
  `status` int(8) NOT NULL DEFAULT '0',
  `paper_format` varchar(64) COLLATE utf8mb4_german2_ci DEFAULT '',
  `num` int(8) DEFAULT '0',
  `color` int(8) DEFAULT '0',
  `direction` int(8) DEFAULT '0',
  `single_side` int(8) DEFAULT '0',
  `remarks` text COLLATE utf8mb4_german2_ci,
  `code` varchar(128) COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `fk_printers_order` (`printer_id`),
  KEY `fk_users_orders` (`user_id`),
  CONSTRAINT `fk_printers_order` FOREIGN KEY (`printer_id`) REFERENCES `printers` (`id`),
  CONSTRAINT `fk_users_orders` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

-- ----------------------------
--  Table structure for `printer_icons`
-- ----------------------------
DROP TABLE IF EXISTS `printer_icons`;
CREATE TABLE `printer_icons` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `printer_id` bigint(20) unsigned NOT NULL,
  `image_avater` text COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_printer_icons_deleted_at` (`deleted_at`),
  KEY `fk_printers_printer_icon` (`printer_id`),
  CONSTRAINT `fk_printers_printer_icon` FOREIGN KEY (`printer_id`) REFERENCES `printers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

-- ----------------------------
--  Table structure for `printers`
-- ----------------------------
DROP TABLE IF EXISTS `printers`;
CREATE TABLE `printers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nick_name` varchar(64) COLLATE utf8mb4_german2_ci NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_german2_ci NOT NULL,
  `pass_word` varchar(64) COLLATE utf8mb4_german2_ci NOT NULL,
  `printer_shop_name` varchar(64) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `introduction` varchar(512) COLLATE utf8mb4_german2_ci NOT NULL,
  `avatar` text COLLATE utf8mb4_german2_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_printers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

-- ----------------------------
--  Records of `printers`
-- ----------------------------
BEGIN;
INSERT INTO `printers` VALUES ('1', '2021-01-18 21:16:12.699', '2021-01-18 21:16:12.699', null, 'admin', 'admin', '$2a$10$fOOqRbPM2H44.bkMgD1HleNwfSfy2tJ14l7adbwlN2RUt5gj3/iyO', '', 'This is admin', ''), ('2', '2021-01-18 21:16:12.840', '2021-01-18 21:16:12.840', null, 'admin', 'admin1', '$2a$10$qjo.OUu.2ThDJ5nqYxjLmuEPelo2K2zeDGd3Y05pqsbrm9A3lV.FC', '', 'This is admin', '');
COMMIT;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nick_name` varchar(64) COLLATE utf8mb4_german2_ci NOT NULL,
  `introduction` varchar(512) COLLATE utf8mb4_german2_ci NOT NULL,
  `avatar` text COLLATE utf8mb4_german2_ci,
  `city` varchar(32) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `gender` int(8) DEFAULT NULL,
  `province` varchar(32) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `unionid` varchar(128) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `openid` varchar(128) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `invite_code` varchar(64) COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

SET FOREIGN_KEY_CHECKS = 1;

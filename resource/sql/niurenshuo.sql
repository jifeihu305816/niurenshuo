/*
Navicat MySQL Data Transfer

Source Server         : Vmware-Centos
Source Server Version : 50717
Source Host           : 192.168.78.128:60975
Source Database       : niurenshuo

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2018-10-24 20:43:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for nrs_comment
-- ----------------------------
DROP TABLE IF EXISTS `nrs_comment`;
CREATE TABLE `nrs_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `topic_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主题id',
  `topic_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主题类型 0为评论 其他为modelId',
  `content` varchar(5000) NOT NULL DEFAULT '',
  `from_uid` int(11) unsigned NOT NULL DEFAULT '0',
  `web_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '网站id',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  `to_uid` int(11) NOT NULL DEFAULT '0' COMMENT '目标用户',
  `comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '评论id，当存在id时，为回复',
  PRIMARY KEY (`id`),
  KEY `idx_web_topic` (`web_id`,`topic_type`,`topic_id`) USING BTREE,
  KEY `idx_cid` (`comment_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='评论';

-- ----------------------------
-- Records of nrs_comment
-- ----------------------------
INSERT INTO `nrs_comment` VALUES ('1', '2018-10-23 20:40:18', '2018-10-24 16:29:36', '2018-10-23 21:31:24', '1', '1', 'haha', '1', '1', '1', '0', '0');
INSERT INTO `nrs_comment` VALUES ('2', '2018-10-23 20:57:10', '2018-10-24 16:29:37', null, '1', '1', 'haha', '1', '1', '0', '0', '0');
INSERT INTO `nrs_comment` VALUES ('3', '2018-10-23 20:39:23', '2018-10-24 16:29:39', null, '1', '1', 'haha', '1', '1', '1', '0', '0');
INSERT INTO `nrs_comment` VALUES ('4', '2018-10-24 15:52:57', '2018-10-24 16:29:40', null, '1', '1', 'haha', '1', '1', '1', '2', '0');
INSERT INTO `nrs_comment` VALUES ('5', '2018-10-24 15:57:00', '2018-10-24 16:35:03', null, '1', '1', 'haha', '2', '1', '1', '1', '2');

-- ----------------------------
-- Table structure for nrs_reply
-- ----------------------------
DROP TABLE IF EXISTS `nrs_reply`;
CREATE TABLE `nrs_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '评论id',
  `reply_id` int(11) NOT NULL DEFAULT '0' COMMENT '回复id',
  `reply_type` int(11) NOT NULL DEFAULT '0' COMMENT '回复类型',
  `content` varchar(5000) NOT NULL DEFAULT '' COMMENT '回复内容',
  `from_uid` int(11) NOT NULL DEFAULT '0' COMMENT '回复uid',
  `to_uid` int(11) NOT NULL DEFAULT '0' COMMENT '目标uid',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_commentid` (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nrs_reply
-- ----------------------------

-- ----------------------------
-- Table structure for nrs_tag
-- ----------------------------
DROP TABLE IF EXISTS `nrs_tag`;
CREATE TABLE `nrs_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

-- ----------------------------
-- Records of nrs_tag
-- ----------------------------

-- ----------------------------
-- Table structure for nrs_user
-- ----------------------------
DROP TABLE IF EXISTS `nrs_user`;
CREATE TABLE `nrs_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `locked` int(11) DEFAULT NULL,
  `last_visit_time` timestamp NULL DEFAULT NULL,
  `register_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nrs_user
-- ----------------------------

-- ----------------------------
-- Table structure for nrs_user_auth
-- ----------------------------
DROP TABLE IF EXISTS `nrs_user_auth`;
CREATE TABLE `nrs_user_auth` (
  `auth_id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT '0',
  `identity_type` varchar(255) DEFAULT NULL COMMENT 'phone手机 email邮箱 name用户名 third第三方',
  `identifier` varchar(255) DEFAULT NULL,
  `credential` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`auth_id`),
  KEY `uid` (`uid`),
  KEY `identity_type` (`identity_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nrs_user_auth
-- ----------------------------
INSERT INTO `nrs_user_auth` VALUES ('1', '1', 'name', 'test', '123456');

-- ----------------------------
-- Table structure for nrs_user_detail
-- ----------------------------
DROP TABLE IF EXISTS `nrs_user_detail`;
CREATE TABLE `nrs_user_detail` (
  `uid` int(11) NOT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `real_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nrs_user_detail
-- ----------------------------

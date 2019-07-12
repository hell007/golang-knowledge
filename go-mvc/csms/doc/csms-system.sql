/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50719
Source Host           : localhost:3306
Source Database       : csms-system

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2019-07-12 15:32:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for jie_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `jie_casbin_rule`;
CREATE TABLE `jie_casbin_rule` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_jie_casbin_rule_v5` (`v5`),
  KEY `IDX_jie_casbin_rule_p_type` (`p_type`),
  KEY `IDX_jie_casbin_rule_v0` (`v0`),
  KEY `IDX_jie_casbin_rule_v1` (`v1`),
  KEY `IDX_jie_casbin_rule_v2` (`v2`),
  KEY `IDX_jie_casbin_rule_v3` (`v3`),
  KEY `IDX_jie_casbin_rule_v4` (`v4`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of jie_casbin_rule
-- ----------------------------
INSERT INTO `jie_casbin_rule` VALUES ('1', 'p', 'superadmin', '/*', 'ANY', '.*', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('2', 'g', 'superadmin', 'admin', '', '', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('3', 'g', 'superadmin', 'goods', '', '', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('4', 'p', 'admin', '/sys/*', 'GET|POST|DELETE|PUT', '.*', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('5', 'p', 'goods', '/goods*', 'GET|POST|DELETE|PUT', '.*', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('6', 'g', 'admin', 'goods', '', '', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('7', 'g', 'admin', 'news', '', '', '', '');
INSERT INTO `jie_casbin_rule` VALUES ('8', 'g', 'superadmin', 'news', '', '', '', '');

-- ----------------------------
-- Table structure for jie_menu
-- ----------------------------
DROP TABLE IF EXISTS `jie_menu`;
CREATE TABLE `jie_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) DEFAULT '',
  `url` varchar(64) DEFAULT '',
  `modular` varchar(64) DEFAULT '' COMMENT '模块名',
  `component` varchar(64) DEFAULT '',
  `name` varchar(64) DEFAULT '',
  `icon` varchar(64) DEFAULT '',
  `keep_alive` varchar(64) DEFAULT '',
  `require_auth` varchar(64) DEFAULT '',
  `parent_id` int(10) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of jie_menu
-- ----------------------------
INSERT INTO `jie_menu` VALUES ('1', '/', null, '', null, '所有', null, null, null, null, '1');
INSERT INTO `jie_menu` VALUES ('2', '/', '/home', '', 'Home', '权限管理', 'fa fa-cog', null, null, '1', '1');
INSERT INTO `jie_menu` VALUES ('3', '/user', '/a', 'admin', 'User', '用户管理', 'fa fa-user', null, null, '2', '1');
INSERT INTO `jie_menu` VALUES ('4', '/role', '/a', 'admin', 'Role', '角色管理', 'fa fa-user-secret', '', '', '2', '1');
INSERT INTO `jie_menu` VALUES ('5', '/menu', '/a', 'admin', 'Menu', '菜单管理', 'fa fa-quora', '', '', '2', '1');
INSERT INTO `jie_menu` VALUES ('14', '/user', '/', '7', '7', 'as', '7', '', '', '3', '7');
INSERT INTO `jie_menu` VALUES ('15', '/user', '/', '8', '8', '大萨达撒の21321', '8', '', '', '3', '8');
INSERT INTO `jie_menu` VALUES ('16', '/menu', '/', '9', '9', 'eqwewqedsads', '9', '', '', '5', '9');
INSERT INTO `jie_menu` VALUES ('17', '/menu', '/', '1', '1', 'kjhjhgjghjgh炬华科技好看', '1', '', '', '5', '1');
INSERT INTO `jie_menu` VALUES ('18', '/menu', '/', '计划国际化', '1', 'ss', '1', '', '', '5', '1');

-- ----------------------------
-- Table structure for jie_role
-- ----------------------------
DROP TABLE IF EXISTS `jie_role`;
CREATE TABLE `jie_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `pid` int(10) unsigned zerofill NOT NULL DEFAULT '0000000001' COMMENT '角色id的父级id',
  `role_name` varchar(20) NOT NULL COMMENT '角色名称',
  `role_note` varchar(100) DEFAULT NULL COMMENT '角色职责描述',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态：启用=1/停用=2',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id` (`id`),
  UNIQUE KEY `name` (`role_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of jie_role
-- ----------------------------
INSERT INTO `jie_role` VALUES ('1', '0000000000', 'superadmin', '超级管理员作为预留最高管理者', '1');
INSERT INTO `jie_role` VALUES ('2', '0000000001', 'admin', '系统管理员', '1');
INSERT INTO `jie_role` VALUES ('3', '0000000001', 'goods', '商品管理员,负责商品维护', '1');
INSERT INTO `jie_role` VALUES ('4', '0000000001', 'news', '文章维护管理人员', '1');

-- ----------------------------
-- Table structure for jie_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `jie_role_menu`;
CREATE TABLE `jie_role_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `rid` int(10) NOT NULL COMMENT '角色id',
  `mid` int(10) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of jie_role_menu
-- ----------------------------
INSERT INTO `jie_role_menu` VALUES ('1', '2', '2');
INSERT INTO `jie_role_menu` VALUES ('2', '2', '3');
INSERT INTO `jie_role_menu` VALUES ('3', '2', '4');
INSERT INTO `jie_role_menu` VALUES ('4', '2', '5');

-- ----------------------------
-- Table structure for jie_user
-- ----------------------------
DROP TABLE IF EXISTS `jie_user`;
CREATE TABLE `jie_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统id',
  `role_id` int(10) NOT NULL COMMENT '角色id',
  `username` varchar(20) NOT NULL COMMENT '系统用户',
  `password` char(32) NOT NULL COMMENT '密码',
  `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：启用=1/停用=0',
  `salt` varchar(64) DEFAULT NULL COMMENT '盐值',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(11) NOT NULL COMMENT '手机号码',
  `ip` varchar(20) DEFAULT NULL COMMENT '登录ip',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `name` (`username`),
  UNIQUE KEY `mobile` (`mobile`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of jie_user
-- ----------------------------
INSERT INTO `jie_user` VALUES ('1', '1', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', '', '', '', '2019-07-11 17:40:44', null);
INSERT INTO `jie_user` VALUES ('2', '2', '曹操', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', 'cc@sina.com', '13888888887', '', '2019-07-05 14:35:26', null);
INSERT INTO `jie_user` VALUES ('3', '3', '刘备', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', 'lb@sina.com', '13888888886', '', '2019-07-05 14:36:37', null);

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo` (
  `uid` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `department` varchar(64) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `UQE_userinfo_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of userinfo
-- ----------------------------

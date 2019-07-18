/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50719
Source Host           : localhost:3306
Source Database       : csms-system

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2019-07-17 11:29:29
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
-- Table structure for jie_dep
-- ----------------------------
DROP TABLE IF EXISTS `jie_dep`;
CREATE TABLE `jie_dep` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `dep_name` varchar(64) NOT NULL COMMENT '部门名称',
  `leader` varchar(64) NOT NULL COMMENT '部门领导人uid',
  `tell` varchar(20) DEFAULT NULL COMMENT '部门电话',
  `email` varchar(64) DEFAULT NULL COMMENT '部门邮箱',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：启用=1/停用=0',
  `pid` int(10) NOT NULL,
  `dep_note` varchar(255) DEFAULT '' COMMENT '部门描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of jie_dep
-- ----------------------------
INSERT INTO `jie_dep` VALUES ('1', '稷下之学', '荀子', '17830895300', '614143260@qq.com', '0', '0', '');
INSERT INTO `jie_dep` VALUES ('2', '儒家', '孔子', '1000', 'laozi@163.com', '0', '1', '崇尚\"礼乐\"和\"仁义\"，主张\"德治\"和\"仁政\"，教育了很多人为人处世所遵循的准则规范。');
INSERT INTO `jie_dep` VALUES ('3', '道家', '老子', '1001', 'laozi@163.com', '0', '1', '天道运行的原理。天道轮回各种哲学上的思想都能在道家找到答案。世界上很多发生的事情都有其规律，人力不可更改。');
INSERT INTO `jie_dep` VALUES ('4', '法家', '李斯', '1002', 'lisi@163.com', '0', '1', '依法治国，天子犯法与庶民同罪。法律视为一种有利于社会统治的强制性工具，体现法制建设的思想，一直被沿用至今。');
INSERT INTO `jie_dep` VALUES ('5', '墨家', '墨子', '1003', 'mozi@163.com', '0', '1', '兼爱非攻，反对强大的王国去攻击弱小的王国，在思想上尊天事鬼，一切以保持社会现状的稳定为主。');
INSERT INTO `jie_dep` VALUES ('6', '名家', '公孙龙', '1004', 'gongsl@163.com', '0', '1', '以辩论出名，著名的白马非马也是名家的思想。以逻辑推理来证明事物的好与坏、真实与否。');
INSERT INTO `jie_dep` VALUES ('7', '阴阳家', '邹衍', '1005', 'zhouyan@163.com', '0', '1', '五行学说，从天文和地理方面来判断事物的凶吉。');
INSERT INTO `jie_dep` VALUES ('8', '纵横家', '苏秦|张仪', '1006', 'sz@163.com', '0', '1', '合纵连横，捭阖之术，阴阳之变化也。');
INSERT INTO `jie_dep` VALUES ('9', '农家', '许行', '1007', 'xuxing@163.com', '0', '1', '注重生产，研究植物生长和产出的学派。');
INSERT INTO `jie_dep` VALUES ('10', '兵家', '孙膑', '1008', 'sunbing@163.com', '0', '1', '讲究利用武力，最大化的夺取敌方的利益从而赢得战争的胜利。');
INSERT INTO `jie_dep` VALUES ('11', '医家', '扁鹊', '1009', 'bianque@163.com', '0', '1', '医者仁心');
INSERT INTO `jie_dep` VALUES ('12', '礼乐', '孟子', '1010', 'mengzi@163.com', '0', '2', '');
INSERT INTO `jie_dep` VALUES ('13', '武当', '张三丰', '1011', 'zsf@163.com', '0', '3', '');
INSERT INTO `jie_dep` VALUES ('14', '庄子游', '庄子', '1012', 'zz@163.com', '0', '3', '');
INSERT INTO `jie_dep` VALUES ('24', 'test', '', '', '', '1', '5', '');

-- ----------------------------
-- Table structure for jie_dep_user
-- ----------------------------
DROP TABLE IF EXISTS `jie_dep_user`;
CREATE TABLE `jie_dep_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `dep_ip` int(10) NOT NULL COMMENT '部门id',
  `uid` int(10) NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of jie_dep_user
-- ----------------------------

-- ----------------------------
-- Table structure for jie_menu
-- ----------------------------
DROP TABLE IF EXISTS `jie_menu`;
CREATE TABLE `jie_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '权限id',
  `name` varchar(32) NOT NULL COMMENT '权限名称',
  `pid` int(10) unsigned NOT NULL COMMENT '权限父id',
  `path` varchar(100) DEFAULT NULL COMMENT '权限路径',
  `redirect` varchar(100) DEFAULT NULL COMMENT 'url',
  `code` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `level` tinyint(3) unsigned NOT NULL COMMENT '权限级别',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：启用=1/停用=0',
  `sort` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '排序',
  `hidden` tinyint(3) NOT NULL DEFAULT '2' COMMENT '显示： 显示=1/隐藏=2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='节点表';

-- ----------------------------
-- Records of jie_menu
-- ----------------------------
INSERT INTO `jie_menu` VALUES ('1', '系统管理', '0', '', '', '', '1', 'xa-icon-config', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('2', '权限管理', '1', 'permission', 'permission/list', 'permission:list', '2', 'xa-icon-verify', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('3', '权限列表', '2', 'list', 'permission/list', 'permission:list', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('4', '权限添加', '2', 'form', 'permission/form', 'permission:add', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('5', '权限修改', '2', 'form', 'permission/form', 'permission:edit', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('6', '权限保存', '2', '', 'permission/save', 'permission:save', '3', '', '1', '0', '2');
INSERT INTO `jie_menu` VALUES ('7', '权限删除', '2', '', 'permission/delete', 'permission:delete', '3', '', '1', '0', '2');
INSERT INTO `jie_menu` VALUES ('8', '权限预留', '2', '', null, '', '3', '', '1', '0', '2');
INSERT INTO `jie_menu` VALUES ('9', '权限预留', '2', '', null, '', '3', '', '1', '0', '2');
INSERT INTO `jie_menu` VALUES ('24', '用户管理', '1', 'user', 'user/list', 'user:list', '2', 'xa-icon-member', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('25', '用户列表', '24', 'list', 'user/list', 'user:list', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('26', '用户添加', '24', 'form', 'user/form', 'user:add', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('27', '用户修改', '24', 'form', 'user/form', 'user:edit', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('35', '角色管理', '1', 'role', 'role/list', 'role:list', '2', 'xa-icon-password', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('36', '角色列表', '35', 'list', 'role/list', 'role:list', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('37', '角色添加', '35', 'form', 'role/form', 'role:add', '3', '', '1', '0', '1');
INSERT INTO `jie_menu` VALUES ('38', '角色修改', '35', 'form', 'role/form', 'role:edit', '3', '', '1', '0', '1');

-- ----------------------------
-- Table structure for jie_menu_bak
-- ----------------------------
DROP TABLE IF EXISTS `jie_menu_bak`;
CREATE TABLE `jie_menu_bak` (
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
-- Records of jie_menu_bak
-- ----------------------------
INSERT INTO `jie_menu_bak` VALUES ('1', '/', null, '', null, '所有', null, null, null, null, '1');
INSERT INTO `jie_menu_bak` VALUES ('2', '/', '/home', '', 'Home', '权限管理', 'fa fa-cog', null, null, '1', '1');
INSERT INTO `jie_menu_bak` VALUES ('3', '/user', '/a', 'admin', 'User', '用户管理', 'fa fa-user', null, null, '2', '1');
INSERT INTO `jie_menu_bak` VALUES ('4', '/role', '/a', 'admin', 'Role', '角色管理', 'fa fa-user-secret', '', '', '2', '1');
INSERT INTO `jie_menu_bak` VALUES ('5', '/menu', '/a', 'admin', 'Menu', '菜单管理', 'fa fa-quora', '', '', '2', '1');
INSERT INTO `jie_menu_bak` VALUES ('14', '/user', '/', '7', '7', 'as', '7', '', '', '3', '7');
INSERT INTO `jie_menu_bak` VALUES ('15', '/user', '/', '8', '8', '大萨达撒の21321', '8', '', '', '3', '8');
INSERT INTO `jie_menu_bak` VALUES ('16', '/menu', '/', '9', '9', 'eqwewqedsads', '9', '', '', '5', '9');
INSERT INTO `jie_menu_bak` VALUES ('17', '/menu', '/', '1', '1', 'kjhjhgjghjgh炬华科技好看', '1', '', '', '5', '1');
INSERT INTO `jie_menu_bak` VALUES ('18', '/menu', '/', '计划国际化', '1', 'ss', '1', '', '', '5', '1');

-- ----------------------------
-- Table structure for jie_role
-- ----------------------------
DROP TABLE IF EXISTS `jie_role`;
CREATE TABLE `jie_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `pid` int(10) unsigned zerofill NOT NULL DEFAULT '0000000001' COMMENT '角色id的父级id',
  `role_name` varchar(20) NOT NULL COMMENT '角色名称',
  `role_note` varchar(100) DEFAULT NULL COMMENT '角色职责描述',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：启用=1/停用=0',
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of jie_role_menu
-- ----------------------------
INSERT INTO `jie_role_menu` VALUES ('1', '2', '1');
INSERT INTO `jie_role_menu` VALUES ('2', '2', '24');
INSERT INTO `jie_role_menu` VALUES ('3', '2', '25');
INSERT INTO `jie_role_menu` VALUES ('4', '2', '26');
INSERT INTO `jie_role_menu` VALUES ('5', '2', '27');

-- ----------------------------
-- Table structure for jie_user
-- ----------------------------
DROP TABLE IF EXISTS `jie_user`;
CREATE TABLE `jie_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统id',
  `role_id` int(10) NOT NULL COMMENT '角色id',
  `username` varchar(20) NOT NULL COMMENT '系统用户',
  `password` char(32) NOT NULL COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：启用=1/停用=0',
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
INSERT INTO `jie_user` VALUES ('1', '1', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', 'root@sina.com', '13888888888', '127.0.0.1', '2019-07-11 17:40:44', null);
INSERT INTO `jie_user` VALUES ('2', '2', '曹操', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', 'cc@sina.com', '13888888887', '127.0.0.1', '2019-07-05 14:35:26', null);
INSERT INTO `jie_user` VALUES ('3', '3', '刘备', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', 'lb@sina.com', '13888888886', '127.0.0.1', '2019-07-05 14:36:37', null);

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

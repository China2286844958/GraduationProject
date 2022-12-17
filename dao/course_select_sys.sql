/*
 Navicat Premium Data Transfer

 Source Server         : MySql
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : course_select_sys

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 09/11/2022 14:13:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `a_id` bigint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '管理员ID',
  `a_email` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员邮箱',
  `a_passwd` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员密码',
  `a_power` int(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '管理员权限',
  `a_loginCode` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '--Locked' COMMENT '管理员登录验证码',
  PRIMARY KEY (`a_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'admin1@1234', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 1, '--Locked');
INSERT INTO `admin` VALUES (2, 'admin2@1234', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 1, '--Locked');
INSERT INTO `admin` VALUES (3, 'admin3@1234', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 1, '--Locked');
INSERT INTO `admin` VALUES (4, 'admin4@1234', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 1, '--Locked');
INSERT INTO `admin` VALUES (5, 'admin5@1234', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 1, '--Locked');
INSERT INTO `admin` VALUES (20020228, 'root@1234', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', 10, '07caip07');

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `c_id` bigint(0) NOT NULL COMMENT '课程ID',
  `c_name` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '课程名',
  `c_teach` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '授课老师',
  `c_Tscore` double(5, 2) NULL DEFAULT NULL COMMENT '课程总分',
  PRIMARY KEY (`c_id`) USING BTREE,
  INDEX `c_Tscore`(`c_Tscore`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course
-- ----------------------------
INSERT INTO `course` VALUES (20227281602843848, 'Python', '吉多·范罗苏姆', 150.00);
INSERT INTO `course` VALUES (202262719856461836, '计算机导论', '鼠标老师', 150.00);
INSERT INTO `course` VALUES (202262719918605121, '汽车构造', '引擎老师', 120.00);
INSERT INTO `course` VALUES (202262719945695528, '数字经济', '眼睛老师', 120.00);
INSERT INTO `course` VALUES (202262915827428533, 'Go Web', '王老师', 150.00);
INSERT INTO `course` VALUES (202272816335793413, 'Vue3+TypeScript', '尤雨溪', 150.00);
INSERT INTO `course` VALUES (2022627191124774753, '机械动力学', '燃油老师', 120.00);
INSERT INTO `course` VALUES (2022627191151750856, '航空礼仪', '旗袍老师', 120.00);
INSERT INTO `course` VALUES (2022728155917827259, 'Java Spring MVC', '詹姆斯·戈士林', 150.00);

-- ----------------------------
-- Table structure for stud_course
-- ----------------------------
DROP TABLE IF EXISTS `stud_course`;
CREATE TABLE `stud_course`  (
  `sc_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '时间戳id',
  `c_id` bigint(0) NULL DEFAULT NULL COMMENT '外键课程表id',
  `sd_id` bigint(0) NULL DEFAULT NULL COMMENT '外键学生id',
  `sd_Sscore` float(5, 2) UNSIGNED ZEROFILL NULL DEFAULT 00.00 COMMENT '学生选课的分数',
  PRIMARY KEY (`sc_id`) USING BTREE,
  INDEX `fk_sid`(`sd_id`) USING BTREE,
  INDEX `fk_cid`(`c_id`) USING BTREE,
  CONSTRAINT `fk_cid` FOREIGN KEY (`c_id`) REFERENCES `course` (`c_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sid` FOREIGN KEY (`sd_id`) REFERENCES `stud_del` (`sd_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 20227281620171641 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of stud_course
-- ----------------------------
INSERT INTO `stud_course` VALUES (2022101923545881, 2022728155917827259, 2022724112219199495, 100.00);
INSERT INTO `stud_course` VALUES (2022101923549885, 20227281602843848, 2022724112219199495, 150.00);
INSERT INTO `stud_course` VALUES (2022102085432505, 202262719918605121, 2022724112219199495, 120.00);
INSERT INTO `stud_course` VALUES (2022716195015470, 202262719918605121, 2022716184157325679, 99.00);
INSERT INTO `stud_course` VALUES (2022716195042156, 202262719856461836, 2022716184245917440, 90.00);
INSERT INTO `stud_course` VALUES (2022716195049806, 202262915827428533, 2022716184348470304, 90.00);
INSERT INTO `stud_course` VALUES (2022716195059154, 202262915827428533, 2022716184157325679, 60.00);
INSERT INTO `stud_course` VALUES (2022716195092966, 202262719945695528, 2022716184157325679, 60.00);
INSERT INTO `stud_course` VALUES (2022716195107633, 2022627191151750856, 2022716184348470304, 80.00);
INSERT INTO `stud_course` VALUES (2022716195111997, 202262915827428533, 202271618441462122, 70.00);
INSERT INTO `stud_course` VALUES (2022716195125637, 2022627191124774753, 202271618441462122, 66.00);
INSERT INTO `stud_course` VALUES (2022716195136282, 202262719856461836, 2022716184348470304, 80.00);
INSERT INTO `stud_course` VALUES (2022832254349226, 202262915827428533, 20228322395254129, 90.00);
INSERT INTO `stud_course` VALUES (20221019235413735, 202272816335793413, 2022724112219199495, 150.00);
INSERT INTO `stud_course` VALUES (20221019235417617, 202262719945695528, 2022724112219199495, 120.00);
INSERT INTO `stud_course` VALUES (20221019235422310, 2022627191124774753, 2022724112219199495, 119.00);
INSERT INTO `stud_course` VALUES (20221019235424902, 2022627191151750856, 2022724112219199495, 120.00);
INSERT INTO `stud_course` VALUES (20221019235428142, 202262719856461836, 2022724112219199495, 150.00);
INSERT INTO `stud_course` VALUES (20227161950122472, 2022627191124774753, 2022716184157325679, 80.00);
INSERT INTO `stud_course` VALUES (20227161950177995, 2022627191151750856, 2022716184157325679, 80.00);
INSERT INTO `stud_course` VALUES (20227161950203777, 202262719856461836, 2022716184157325679, 80.00);
INSERT INTO `stud_course` VALUES (20227161950276812, 202262915827428533, 2022716184245917440, 58.00);
INSERT INTO `stud_course` VALUES (20227161950306769, 202262719945695528, 2022716184245917440, 80.00);
INSERT INTO `stud_course` VALUES (20227161950339411, 2022627191124774753, 2022716184245917440, 80.00);
INSERT INTO `stud_course` VALUES (20227161950369310, 202262719918605121, 2022716184245917440, 80.00);
INSERT INTO `stud_course` VALUES (20227161950396130, 2022627191151750856, 2022716184245917440, 80.00);
INSERT INTO `stud_course` VALUES (20227161950519791, 202262719945695528, 2022716184348470304, 80.00);
INSERT INTO `stud_course` VALUES (20227161950549673, 2022627191124774753, 2022716184348470304, 80.00);
INSERT INTO `stud_course` VALUES (20227161950577289, 202262719918605121, 2022716184348470304, 80.00);
INSERT INTO `stud_course` VALUES (20227161951161780, 202262719856461836, 202271618441462122, 80.00);
INSERT INTO `stud_course` VALUES (20227161951189962, 202262719945695528, 202271618441462122, 80.00);
INSERT INTO `stud_course` VALUES (20227161951216453, 2022627191151750856, 202271618441462122, 80.00);
INSERT INTO `stud_course` VALUES (20227161951282775, 202262719918605121, 202271618441462122, 80.00);
INSERT INTO `stud_course` VALUES (20227161951401780, 202262915827428533, 2022716184411764691, 59.00);
INSERT INTO `stud_course` VALUES (20227161951435941, 202262719856461836, 2022716184411764691, 35.00);
INSERT INTO `stud_course` VALUES (20227161951457332, 202262719945695528, 2022716184411764691, 33.00);
INSERT INTO `stud_course` VALUES (20227161951487299, 2022627191151750856, 2022716184411764691, 30.00);
INSERT INTO `stud_course` VALUES (20227161951513942, 2022627191124774753, 2022716184411764691, 21.00);
INSERT INTO `stud_course` VALUES (20227161951545434, 202262719918605121, 2022716184411764691, 75.00);
INSERT INTO `stud_course` VALUES (20227241124102887, 202262915827428533, 2022724112219199495, 147.00);

-- ----------------------------
-- Table structure for stud_del
-- ----------------------------
DROP TABLE IF EXISTS `stud_del`;
CREATE TABLE `stud_del`  (
  `sd_id` bigint(0) NOT NULL,
  `sd_relname` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '真实姓名',
  `sd_gender` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '性别',
  `sd_age` int(0) NULL DEFAULT NULL COMMENT '年龄',
  `sd_address` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '住址',
  `sd_sys` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '院系',
  PRIMARY KEY (`sd_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of stud_del
-- ----------------------------
INSERT INTO `stud_del` VALUES (20228322395254129, '凯丽', '女', 22, '美国纽约', '信息传媒系');
INSERT INTO `stud_del` VALUES (202271618441462122, '金', '男', 20, '山东', '信息传媒系');
INSERT INTO `stud_del` VALUES (2022716184157325679, '木', '男', 20, '山东', '信息传媒系');
INSERT INTO `stud_del` VALUES (2022716184245917440, '水', '女', 20, '山东', '信息传媒系');
INSERT INTO `stud_del` VALUES (2022716184348470304, '火', '女', 20, '山东', '信息传媒系');
INSERT INTO `stud_del` VALUES (2022716184411764691, '土', '女', 20, '山东', '信息传媒系');
INSERT INTO `stud_del` VALUES (2022724112219199495, '薛智敏', '男', 120, '山东菏泽', '信息传媒系');

-- ----------------------------
-- Table structure for stud_login
-- ----------------------------
DROP TABLE IF EXISTS `stud_login`;
CREATE TABLE `stud_login`  (
  `sl_id` bigint(0) NOT NULL COMMENT '学生ID',
  `sl_email` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '绑定邮箱',
  `sl_passwd` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `sl_loginCode` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '--Locked' COMMENT '登录的验证码',
  `sl_logo` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '/UploadFiles/logo.gif' COMMENT '[Bate]头像地址',
  PRIMARY KEY (`sl_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of stud_login
-- ----------------------------
INSERT INTO `stud_login` VALUES (202271618441462122, 'student4@qq.com', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', '--Locked', '/UploadFiles/logo.gif');
INSERT INTO `stud_login` VALUES (2022716184157325679, 'student1@qq.com', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', '--Locked', '/UploadFiles/logo.gif');
INSERT INTO `stud_login` VALUES (2022716184245917440, 'student2@qq.com', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', '--Locked', '/UploadFiles/logo.gif');
INSERT INTO `stud_login` VALUES (2022716184348470304, 'student3@qq.com', '0ffe1abd1a08215353c233d6e009613e95eec4253832a761af28ff37ac5a150c', '--Locked', '/UploadFiles/logo.gif');
INSERT INTO `stud_login` VALUES (2022724112219199495, '2286844958@qq.com', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', 'zpra', '/UploadFiles/2022724112219199495_20221191482117310.jpg');

-- ----------------------------
-- Procedure structure for Demo1
-- ----------------------------
DROP PROCEDURE IF EXISTS `Demo1`;
delimiter ;;
CREATE PROCEDURE `Demo1`(IN `X` int)
BEGIN
	#Routine body goes here...
	

END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table course
-- ----------------------------
DROP TRIGGER IF EXISTS `选课表同步删除`;
delimiter ;;
CREATE TRIGGER `选课表同步删除` BEFORE DELETE ON `course` FOR EACH ROW BEGIN
delete from stud_course where stud_course.c_id=old.c_id;

END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table stud_del
-- ----------------------------
DROP TRIGGER IF EXISTS `删除学生登录表`;
delimiter ;;
CREATE TRIGGER `删除学生登录表` BEFORE DELETE ON `stud_del` FOR EACH ROW begin 
-- 删除学生登录表
DELETE from stud_login  where stud_login.sl_id = old.sd_id;

end
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table stud_del
-- ----------------------------
DROP TRIGGER IF EXISTS `删除选课表`;
delimiter ;;
CREATE TRIGGER `删除选课表` BEFORE DELETE ON `stud_del` FOR EACH ROW begin

-- 删除课程表的id

Delete from stud_course  where stud_course.sd_id =old.sd_id;

end
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table stud_login
-- ----------------------------
DROP TRIGGER IF EXISTS `in_syc_stud_del`;
delimiter ;;
CREATE TRIGGER `in_syc_stud_del` AFTER INSERT ON `stud_login` FOR EACH ROW begin 

insert stud_del (stud_del.sd_id) value(new.sl_id);

end
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;

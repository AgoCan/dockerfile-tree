-- dockerfile表，主要就是记录dockerfile，可以被创建和修改
drop TABLE IF EXISTS `dockerfile`;
CREATE TABLE `dockerfile` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `dockerfile` text NOT NULL,
  `level_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_dockerfile_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 层级表
/*
  1. cuda
  2. py
      2.1 py36
      2.2 py37
  3. opencv
  4. vidio
  5. ice-tools
  6. framework
      6.1 tensorflow1.13
          6.1.1 openvino
  7. lab
  8. coding-env
*/
drop TABLE IF EXISTS `level`;
CREATE TABLE `level` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(64) NOT NULL COMMENT '当前层级的名称',
  `comment` varchar(256) COMMENT '当前层级的注释',
  `order_id` int NOT NULL COMMENT '顺序',
  `parent_id` int NOT NULL COMMENT '没有父节点的为顶层，没有对应的dockerfile',
  PRIMARY KEY (`id`),
  KEY `idx_level_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- 插入level
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (1, 'cuda', 1, 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (2, 'python', 2, 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `comment`, `parent_id`) VALUES (3, 'video', 3, 'ffmpeg相关组件', 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (4, 'protobuf', 4, 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (5, 'opencv', 5, 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `comment`, `parent_id`) VALUES (6, 'ice', 6, 'ice是rpc框架，全称 zeroc ice', 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (7, 'framework', 7, 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (8, 'tensorflow1.13', 1, 7);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (9, 'openvino', 1, 8);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (10, 'openvinor2', 2, 8);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (11, 'lab', 8, 0);
INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (12, 'coding-env', 9, 0);

INSERT INTO `dockerfile_tree`.`level`(`id`, `name`, `order_id`, `parent_id`) VALUES (13, '10.0-cudnn7.6.5', 1, 1);

-- 层级和组合任务表的联合表，并且记录构建的dockerfile内容，该内容是不允许手动页面修改，只能程序修改
drop TABLE IF EXISTS `level_combination_task`;
CREATE TABLE `level_combination_task` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `record_dockerfile` text NOT NULL COMMENT '记录构建成功的dockerfile',
  `combination_task_id` int NOT NULL COMMENT 'task id',
  `level_id` int NOT NULL COMMENT 'level id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 组合表，记录构建的组合关系
drop TABLE IF EXISTS `combination_task`;
CREATE TABLE `combination_task` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `last_level` varchar(64) NOT NULL,
  `build_status` int NOT NULL DEFAULT 1 COMMENT '是否构建成功，0 失败， 1成功',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 记录表，跟组合表一一对应，主要记录组合构建成功后的数据
drop TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `image_info` varchar(512) NOT NULL COMMENT '构建信息展示',
  `image_name` varchar(512) NOT NULL COMMENT '镜像名',
  `push_status` int NOT NULL DEFAULT 1 COMMENT '是否push镜像成功，0 失败， 1成功',
  `combination_task_id` int NOT NULL COMMENT '一对一task表',
  UNIQUE KEY `image_name` (`image_name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 资源表，构建镜像的时候  FROM resource:v1 as resource 不被删除，主要就是 copy使用
drop TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `image_name` varchar(64) NOT NULL,
  `dockerfile_url_path` varchar(64) DEFAULT NULL COMMENT '记录资源镜像的dockerfile制作位置',
  UNIQUE KEY `image_name` (`image_name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 配置表，配置文件
drop TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  `key` varchar(64) NOT NULL COMMENT '配置文件的key',
  `value` varchar(64) NOT NULL COMMENT '配置文件的value',
  UNIQUE KEY `key` (`key`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- 创建数据库
CREATE
DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

-- 进入 blog_service
use
blog_service;

/* 公共字段
`created_on`  int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
`deleted_on`  int(10) unsigned DEFAULT '0' COMMENT '删除时间',
`is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0未删除，1已删除',
 */

-- 创建认证表
CREATE TABLE `blog_auth`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key`     varchar(20)  DEFAULT '' COMMENT 'Key',
    `app_secret`  varchar(50)  DEFAULT '' COMMENT 'Secret',
    `created_on`  int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0未删除，1已删除',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态，0禁用，1启用',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

-- 插入一个账户
INSERT INTO `blog_service`.`blog_auth` (`id`, `app_Key`, `app_secret`, `created_on`, `created_by`, `modified_on`,
                                        `modified_by`, `deleted_on`, `is_del`, `state`)
VALUES (1, 'creator', 'creator_secret', 0, 'creator', 0, '', 0, 0, 1);

-- 创建标签表
CREATE TABLE `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_on`  int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0未删除，1已删除',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态，0禁用，1启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

-- 创建文章表
CREATE TABLE `blog_article`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`           varchar(100) DEFAULT '' COMMENT '文章标题',
    `desc`            varchar(255) DEFAULT '' COMMENT '文章描述',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `content`         longtext COMMENT '文章内容',
    `created_on`      int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by`      varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on`     int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by`     varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`      int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`          tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0未删除，1已删除',
    `state`           tinyint(3) unsigned DEFAULT '1' COMMENT '状态，0禁用，1启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

-- 创建文章标签关联表
CREATE TABLE `blog_article_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id`  int(11) unsigned NOT NULL COMMENT '文章ID',
    `tag_id`      int(11) unsigned NOT NULL COMMENT '标签ID',
    `created_on`  int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0未删除，1已删除',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态，0禁用，1启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';

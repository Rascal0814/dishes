CREATE TABLE dishes (
    `id` bigint unsigned not null AUTO_INCREMENT COMMENT '主键',
    `name` varchar(255) not null default '' COMMENT '菜名',
    `logo` varchar(255) not null default '' COMMENT 'logo',
    `remark` varchar(255) not null default '' COMMENT '备注',
    `created_at` datetime not null default now() comment '创建时间',
    `updated_at` datetime not null default now() on update now() comment '更新时间',
    `deleted_at` datetime default null comment '是否已删除',
    PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
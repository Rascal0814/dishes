CREATE TABLE orders (
    `id` bigint unsigned not null AUTO_INCREMENT COMMENT '主键',
    `dish_id` bigint unsigned not null comment '菜品ID',
    `status` int(1) not null default 0 comment '订单状态 0未开始 1制作中 2已完成',
    `creator` bigint unsigned not null comment '创建人',
    `remark` varchar(255) not null default '' COMMENT '备注',
    `created_at` datetime not null default now() comment '创建时间',
    `updated_at` datetime not null default now() on update now() comment '更新时间',
    `deleted_at` datetime default null comment '是否已删除',
    PRIMARY KEY ( `id` ),
    FOREIGN KEY (`dish_id`) references dishes (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
create database blog charset=utf8mb4;

use blog;


create table `article`(
    `id` bigint(20) not null auto_increment comment `文章id`,
    `category_id` bigint(20) unsignded not null comment `分类id`,
    `content` longtext not null comment `文章内容`,
    `title` varchar(256) unsigned not null comment `文章标题`,
    `view_count` int(255) comment `阅读次数`,
    `comment_count` int(255) comment `评论次数`,
    `username` varchar(256) not null comment `作者`,
    `status` int(10) default 1 comment `评审状态`,
    `summary` varchar(255) not null comment `文章摘要`,
    `create_time` timestamp not null default current_timestamp comment `创建时间`,
    `update_time` timestamp null default current_timestamp on update current_timestamp comment `更新时间`,
    primary key (`id`) using btree
    key `idx_view_count` (`view_count`) using btree comment `阅读次数索引`,
    key `idx_comment_count` (`comment_count`) using btree comment `评论次数索引`,
    key `idx_category_id` (`category_id`) using btree comment `分类id索引`
) engine=InnoDB auto_increment=6 default charset=utfmb4 row_format=dynamic;

begin;
insert into `article` values (1,1,'go真简单','go真简单',10,0,'娜扎','1','go真简单','2019-07-05 20:54:32','2019-07-05 20:55:32')
commit ;

create table `category`(
    `id` bigint(20) not null auto_increment comment `分类id`,
    `category_name` varchar(255)  not null comment `分类名字`,
    `category_no` int(10)  not null comment `分类排序`,
    `create_time` timestamp not null default current_timestamp comment `创建时间`,
    `update_time` timestamp null default current_timestamp on update current_timestamp comment `更新时间`,
    primary key (`id`) using btree
)engine=InnoDB auto_increment=6 default charset=utfmb4 row_format=dynamic;

begin ;
insert into `category` values (1,'go develop',1,'2019-07-05 20:54:32','2019-07-05 20:55:32')
commit ;

create table `comment`(
    `id` bigint(20) not null auto_increment comment `评论id`,
    `content` text  not null comment `评论内容`,
    `username` varchar(64) not null comment `评论作者`,
    `create_time` timestamp not null default current_timestamp comment `创建时间`,
    `status` int(10) unsigned comment `评论状态，0，删除，1，正常`,
    `article_id` bigint(20) unsigned default null ,
    primary key (`id`) using btree
)engine=InnoDB  default charset=utfmb4 row_format=dynamic;

create table `person` (
    `id` bigint(20) not null auto_increment ,
    `person_id` bigint(20) not null ,
    `username` varchar(64) not null ,
    `nickname` varchar(64) not null ,
    `password` varchar(64) not null ,
    `gender` tinyint(4) not null default 0,
    `create_time` timestamp not null default current_timestamp comment `创建时间`,
    `update_time` timestamp null default current_timestamp on update current_timestamp comment `更新时间`,
    primary key (`id`),
    unique key `idx_username` (`username`) using btree,
    unique key `idx_user_id` (`user_id`) using btree
)engine=InnoDB  default charset=utfmb4;





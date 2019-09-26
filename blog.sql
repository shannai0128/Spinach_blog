create database spinachBlog charset=utf8mb4;

use spinachBlog;


create table `article`(
    `id` bigint(20) not null auto_increment,
    `category_id` bigint(20) not null,
    `content` longtext not null,
    `title` varchar(256) not null,
    `view_count` int(255),
    `person_name` varchar(256) not null,
    `status` int(10) default 0,
    `violat_reason`  text,
    `summary` varchar(255) not null,
    `origin`  int,
    `praise` bigint(20) default 0,
    `create_time` TIMESTAMP default CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP default CURRENT_TIMESTAMP on update current_timestamp,
    primary key (`id`) using btree,
    key `idx_view_count` (`view_count`) using btree,
    key `idx_category_id` (`category_id`) using btree
) engine=InnoDB auto_increment=6 default charset=utf8mb4 row_format=dynamic;

begin;
insert into `article` values (1,1,'go真简单','go真简单',10,'娜扎','1','demo demo','go summary',1,100,'2019-07-05 20:54:32','2019-07-05 20:55:32');
commit;

create table `category`(
    `id` bigint(20) not null auto_increment,
    `category_name` varchar(255)  not null,
    primary key (`id`) using btree
)engine=InnoDB auto_increment=6 default charset=utf8mb4 row_format=dynamic;

begin;
insert into `category` values (1,'go语言');
commit;

# create table `comment`(
#     `id` bigint(20) not null auto_increment,
#     `content` text  not null,
#     `person_id` varchar(64) not null,
#     `create_time` timestamp default current_timestamp,
#     `status` int(10) unsigned,
#     `article_id` bigint(20) unsigned default null ,
#     primary key (`id`) using btree,
#     CONSTRAINT`comment_ibfk_1`FOREIGNKEY(`person_id`)REFERENCES`person` (`id`)
# )engine=InnoDB  default charset=utf8mb4 row_format=dynamic;

create table `person` (
    `id` bigint(20) not null auto_increment ,
    `person_name` varchar(64) not null,
    `id_card`    varchar(64) not null,
    `password` varchar(64) not null,
    `gender` tinyint(4) not null default 0,
    `login_ip` varchar(20),
    `login_time` timestamp default current_timestamp,
    `is_active`  tinyint(4) default 0,
    `create_time` timestamp default current_timestamp,
    `update_time` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_username` (`person_name`) using btree
)engine=InnoDB  default charset=utf8mb4;


create table `person_info`(
    `id` bigint(20) not null auto_increment,
    `person_email` varchar(128),
    `person_addr` varchar(255),
    `nick_name` varchar(50),
    primary key (`id`) using btree,
    foreign key (`id`) REFERENCES `person` (`id`)
)engine=InnoDB default charset=utf8mb4 row_format=dynamic;


create table `follower`(
    `id` bigint(20) not null auto_increment,
    `followed_id` text,
    `followed` text,
    `status` int ,
    primary key (`id`) using btree,
    foreign key (`id`)REFERENCES `person` (`id`)
)engine=InnoDB default charset=utf8mb4 row_format=dynamic;

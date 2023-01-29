create table `carts` (
	`user_id` int not null,
    `food_id` int not null,
    `quantity` int not null,
    `status` int not null default 1,
	`create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key(`user_id`, `food_id`),
    key `food_id` (`food_id`)
);

create table `categories` (
	`id` int not null auto_increment,
    `name` varchar(100) not null,
    `desc` text,
    `icon` json default null,
    `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`)
);

create table `cities` (
	`id` int not null auto_increment,
    `title` varchar(100) not null,
     `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`)
);

create table `food_likes` (
	`user_id` int not null,
    `food_id` int not null,
     `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`user_id`, `food_id`),
    key `food_id` (`food_id`)
);

drop table if exists `food_ratings`;
create table `food_ratings`(
	`id` int not null auto_increment,
    `user_id` int not null,
    `food_id` int not null,
    `point` float default 0,
    `comment` text,
    `status` int not null default 1,
     `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
	primary key (`id`),
    key `food_id` (`food_id`) using btree
);

create table `foods`(
	`id` int not null auto_increment,
    `restaurant_id` int not null,
    `category_id` int default null, 
    `name` varchar(255) not null,
    `desc` text,
    `price` float not null,
    `images` json not null,
    `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key(`id`),
    key `restaurant_id` (`restaurant_id`) using btree,
    key `category_id` (`category_id`) using btree,
    key `status` (`status`) using btree
    
);

create table `images` (
	`id` int not null auto_increment,
    `file_name` varchar(100) character set utf8mb4 collate utf8mb4_0900_ai_ci not null,
    `width` int not null,
    `height` int not null,
    `status` int not null default 1, 
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key(`id`)
);


create table `order_details` (
	`id` int not null auto_increment,
    `order_id` int not null,
    `food_origin` json default null,
    `price` float not null,
    `quantity` int not null,
    `discount` float default 0,
     `status` int not null default 1, 
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key(`id`),
    key `order_id` (`order_id`) using btree
);

create table `order_trackings` (
	`id` int not null auto_increment,
    `order_id` int not null,
    `state` enum("waiting_for_shipper", "preparing", "on_the_way", "delivered", "cancel")not null,
     `status` int not null default 1, 
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`),
    key `order_id` (`order_id`) using btree
);

create table orders(
	`id` int not null auto_increment,
    `user_id` int not null,
    `total_price` float not null, 
    `shipper_id` int default null,
    `status` int not null default 1, 
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`),
    key `user_id` (`user_id`) using btree,
    key `shipper_id` (`shipper_id`) using btree
);

drop table if exists `restaurant_foods`;
create table `restaurant_foods` (
	`restaurant_id` int not null,
    `food_id` int not null,
    `status` int not null default 1, 
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`restaurant_id`,`food_id` ),
    key `food_id` (`food_id`)
);categories

create table restaurant_likes(
	`restaurant_id` int not null,
    `user_id` int not null,
    `status` int not null default 1, 
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`restaurant_id`,`user_id` ),
    key `user_id` (`user_id`)
);

create table restaurant_ratings(
	`id` int not null auto_increment,
    `user_id` int not null,
    `restaurant_id` int not null,
    `point` float not null default 0,
    `comment` text,
    `status` int not null default 1, 
	`create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key(`id`),
    key `user_id` (`user_id`) using btree,
    key `restaurant_id` (`restaurant_id`) using btree
);

create table restaurants(
	`id` int not null auto_increment,
    `owner_id` int not null,
    `name` varchar(50) not null,
    `addr` varchar(255) not null,
    `city_id` int default null,
    `lat` double default null,
    `lng` double default null,
    `cover` json not null,
    `logo` json not null,
    `shipping_fee_per_km` double default 0,
    `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`),
    key `owner_id` (`owner_id`) using btree,
    key `city_id` (`city_id`) using btree,
    key `status` (`status`) using btree
);

create table user_addresses (
	`id` int not null auto_increment,
    `user_id` int not null,
    `city_id` int not null,
    `title` varchar(100) default null,
    `icon` json default null,
    `addr` varchar(255) not null,
    `lat` double default null,
    `lng` double default null,
    `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`),
	key `user_id` (`user_id`) using btree,
    key `city_id` (`city_id`) using btree
);

create table `user_device_tokens` (
	`id` int unsigned not null auto_increment,
    `user_id` int unsigned default null,
    `is_production` tinyint(1) default 0,
    `os` enum("ios", "android", "web") default "ios" comment "1: iOS, 2: Android",
    `token` varchar(255) default null,
    `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key (`id`), 
    key `user_id` (`user_id`) using btree,
    key `os` (`os`) using btree
);

create table users(
	`id` int not null auto_increment,
    `email` varchar(50) not null,
    `fb_id` varchar(50) default null,
    `gg_id` varchar(50) default null,
    `password` varchar(50) not null,
    `salt` varchar(50) default null,
    `last_name` varchar(50) not null,
    `first_name` varchar(50) not null,
    `phone` varchar(20) default null,
    `role` enum("user", "admin", "shipper") not null default "user",
    `avatar` json default null,
    `status` int not null default 1,
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    primary key(`id`),
    unique key `email` (`email`)
)









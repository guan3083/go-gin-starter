CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL,
  `password` varchar(128) NOT NULL,
  `nick_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `phone_number` varchar(32) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  `secret_key` varchar(64) NOT NULL,
  `status` tinyint NOT NULL COMMENT '1：暂停，2：启用',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

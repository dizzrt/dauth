CREATE TABLE IF NOT EXISTS `dauth`.`identity_users` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'uid,user_id',
  `username` varchar(64) NOT NULL,
  `nickname` varchar(64) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `password` char(60) NOT NULL,
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '1-ENABLED 2-DISABLED 3-LOCKED',
  `last_login_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  UNIQUE KEY `uk_phone` (`phone`),
  UNIQUE KEY `uk_email` (`email`),
  KEY `k_status` (`status`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

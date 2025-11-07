CREATE TABLE IF NOT EXISTS `dauth`.`identity_users` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'uid,user_id',
  `email` varchar(127) NOT NULL,
  `username` varchar(127) NOT NULL,
  `password` char(60) NOT NULL,
	`status` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `last_login_time` datetime NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  KEY `i_status` (`status`),
	UNIQUE INDEX `ui_email`(`email`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS `dauth`.`identity_roles` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(127) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
	`created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS `dauth`.`identity_user_role_associations` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL COMMENT 'identity.users.id',
  `role_id` int UNSIGNED NOT NULL COMMENT 'identity.roles.id',
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  KEY `i_user_id` (`user_id`),
  KEY `i_role_id` (`role_id`),
  UNIQUE KEY `ui_user_role` (`user_id`,`role_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;
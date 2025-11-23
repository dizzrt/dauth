CREATE TABLE IF NOT EXISTS `dauth`.`client_clients` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'cid,client_id',
  `name` varchar(127) NOT NULL,
  `description` varchar(255) NULL,
  `redirect_uri` varchar(255) NOT NULL,
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS `dauth`.`client_scopes` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(127) NOT NULL,
  `description` varchar(255) NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

INSERT INTO `dauth`.`client_scopes` (`name`, `description`, `status`, `created_at`, `updated_at`) VALUES
	('openid', '', NOW(), NOW()),
	('profile', '', NOW(), NOW()),
	('email', '', NOW(), NOW()),
	('phone', '', NOW(), NOW()),
	('address', '', NOW(), NOW());

CREATE TABLE IF NOT EXISTS `dauth`.`client_client_scope_associations` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `client_id` varchar(127) NOT NULL,
  `scope_id` varchar(255) NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_client_scope` (`client_id`, `scope_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

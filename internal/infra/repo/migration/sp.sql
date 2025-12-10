CREATE TABLE IF NOT EXISTS `dauth`.`sp_service_providers` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'sp_id,service_provider_id',
  `name` varchar(127) NOT NULL,
  `description` varchar(255) NULL,
	`secret` varchar(127) NOT NULL,
  `redirect_uri` varchar(255) NOT NULL,
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS `dauth`.`sp_scopes` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(127) NOT NULL,
  `description` varchar(255) NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

INSERT INTO `dauth`.`sp_scopes` (`name`, `description`, `created_at`, `updated_at`) VALUES
	('openid', '', NOW(), NOW()),
	('profile', '', NOW(), NOW()),
	('email', '', NOW(), NOW()),
	('phone', '', NOW(), NOW()),
	('address', '', NOW(), NOW());

CREATE TABLE IF NOT EXISTS `dauth`.`sp_sp_scope_associations` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `sp_id` int UNSIGNED NOT NULL,
  `scope_id` int UNSIGNED NOT NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_sp_scope` (`sp_id`, `scope_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

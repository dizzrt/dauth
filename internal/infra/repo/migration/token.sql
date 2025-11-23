CREATE TABLE IF NOT EXISTS `dauth`.`token_token_blacklist` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `token_id` varchar(64) NOT NULL COMMENT 'JWT jti',
  `revoked_at` datetime NOT NULL,
  `expires_at` datetime NOT NULL,
  `revoke_reason` varchar(255) DEFAULT NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_token_id` (`token_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

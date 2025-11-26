CREATE TABLE IF NOT EXISTS `dauth`.`auth_authorization_codes` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` varchar(127) NOT NULL,
  `user_id` int UNSIGNED NOT NULL,
  `client_id` int UNSIGNED NOT NULL,
  `redirect_uri` varchar(255) NOT NULL,
  `scope` varchar(255) NOT NULL,
  `issued_at` datetime NOT NULL,
  `expires_at` datetime NOT NULL,
  `used` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_code` (`code`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;
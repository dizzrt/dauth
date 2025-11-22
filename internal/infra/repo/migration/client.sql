CREATE TABLE IF NOT EXISTS `dauth`.`client_clients` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'cid,client_id',
  `name` varchar(127) NOT NULL,
  `redirect_uri` varchar(255) NOT NULL,
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  KEY `i_name` (`name`),
  KEY `i_status` (`status`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;
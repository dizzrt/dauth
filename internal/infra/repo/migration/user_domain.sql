CREATE TABLE IF NOT EXISTS `dauth`.`users`  (
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
  UNIQUE INDEX `ui_email`(`email`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4;

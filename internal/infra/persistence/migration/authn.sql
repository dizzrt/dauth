CREATE TABLE IF NOT EXISTS `dauth`.`authn_attempts` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int UNSIGNED DEFAULT NULL COMMENT 'identity_users.id, nullable when user not found',
  `did` varchar(128) DEFAULT NULL COMMENT 'device unique id',
  `authn_type` tinyint NOT NULL COMMENT '1-PWD 2-PWD+MFA 3-THIRD_PARTY',
  `account` varchar(64) NOT NULL COMMENT 'identity_users.username|phone|email',
  `client_ip` varchar(64) DEFAULT NULL,
  `session_id` varchar(64) DEFAULT NULL COMMENT 'session id',
  `status` tinyint NOT NULL COMMENT 'attempt status: 1-success 2-failure 3-blocked',
  `fail_reason` varchar(128) DEFAULT NULL COMMENT 'failure reason: password error/MFA error/account locked etc.',
  `user_agent` text DEFAULT NULL COMMENT 'Client UA (device/browser info)',
  `created_at` datetime NOT NULL COMMENT 'attempt time',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_account_status_created` (`account`, `status`, `created_at`),
  KEY `idx_ip_status_created` (`client_ip`, `status`, `created_at`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4 COMMENT='authentication attempt record table';

CREATE TABLE IF NOT EXISTS `dauth`.`authn_policies` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `policy_type` tinyint NOT NULL COMMENT '1-global 2-user 3-role',
  `pwd_fail_limit` tinyint NOT NULL DEFAULT 5 COMMENT 'password fail limit (default 5 times)',
  `pwd_lock_duration` int NOT NULL DEFAULT 1800 COMMENT 'password lock duration (seconds, default 30 minutes)',
  `mfa_force` tinyint NOT NULL DEFAULT 0 COMMENT '0-false 1-true',
  `mfa_fail_limit` tinyint NOT NULL DEFAULT 3 COMMENT 'MFA fail limit (default 3 times)',
  `ip_whitelist` json DEFAULT NULL COMMENT 'IP whitelist (JSON array)',
  `ip_blacklist` json DEFAULT NULL COMMENT 'IP blacklist (JSON array)',
  `session_timeout` int NOT NULL DEFAULT 86400 COMMENT 'session timeout (seconds, default 24 hours)',
  `is_enabled` tinyint NOT NULL DEFAULT 1 COMMENT '0-disabled 1-enabled',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4 COMMENT='Login security policy table';

CREATE TABLE IF NOT EXISTS `dauth`.`authn_policy_targets` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `policy_id` int UNSIGNED NOT NULL COMMENT 'authn_policies.id',
  `target_id` int UNSIGNED NOT NULL COMMENT 'identity_users.id|identity_roles.id',
  `target_type` tinyint NOT NULL COMMENT '1-user 2-role',
  `is_enabled` tinyint NOT NULL DEFAULT 1 COMMENT '0-disabled 1-enabled',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_policy_id_target_id_target_type` (`policy_id`,`target_id`,`target_type`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4 COMMENT='Login security policy target table';

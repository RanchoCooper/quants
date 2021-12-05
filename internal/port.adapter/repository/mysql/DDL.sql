CREATE DATABASE IF NOT EXISTS `quant` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `quant_user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(100) DEFAULT '' COMMENT 'user name',
    `user_email` VARCHAR(100) DEFAULT '' COMMENT 'user email',
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT 'state 0: enable, 1: unable',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE KEY unique_key_user_email(`user_email`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = 'user table';

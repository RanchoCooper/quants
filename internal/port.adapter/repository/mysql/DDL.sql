DROP TABLE IF EXISTS `example`;

CREATE TABLE `example` (
    `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '名称',
    `alias` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '别称',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Hexagonal示例表';

CREATE TABLE IF NOT EXISTS `quant_user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(100) DEFAULT '' COMMENT 'user name',
    `user_email` VARCHAR(100) DEFAULT '' COMMENT 'user email',
    `asset` DECIMAL NOT NULL DEFAULT '0.0' COMMENT 'asset',
    `profit` DECIMAL NOT NULL DEFAULT '0.0' COMMENT 'profit',
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT 'state 0: enable, 1: unable',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE KEY unique_key_user_email(`user_email`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = 'user table';

CREATE TABLE IF NOT EXISTS `quant_trade` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT 'user email',
    `symbol` VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'symbol',
    `order_id` VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'order id',
    `type` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT 'state 0: default, 1: buy, 2: sell',
    `price` DECIMAL NOT NULL DEFAULT '0.0' COMMENT 'price',
    `quantity` DECIMAL NOT NULL DEFAULT '0.0' COMMENT 'quantity',
    `is_simulate` BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'is simulate',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE KEY unique_key_order_id(`order_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = 'user table';
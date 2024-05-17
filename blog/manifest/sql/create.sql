CREATE TABLE `blog` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'Blog ID',
    `title` varchar(255) NOT NULL COMMENT 'Title',
    `content` varchar(500) NOT NULL COMMENT 'Content',
    `nickname` varchar(45) NOT NULL COMMENT 'Nickname',
    `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    `delete_at` datetime DEFAULT NULL COMMENT 'Deleted Time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
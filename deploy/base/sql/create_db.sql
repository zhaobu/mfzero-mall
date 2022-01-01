-- 给数据库新用户必要权限(权限分配可以参考 https://www.yuque.com/docs/share/eec500f1-a331-4668-a565-1a7437a62f08?# 《MySQL修改用户权限》)
-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'amigo'
use mysql;
UPDATE user SET `host` = '%' WHERE `user` = 'lw' LIMIT 1;
flush privileges;
GRANT ALL ON *.* TO 'lw'@'%';
flush privileges;
-- 创建数据库
CREATE DATABASE IF NOT EXISTS `mf_admin` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_bin';
CREATE DATABASE IF NOT EXISTS `mf_demo` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_bin';

-- 建demo表
use mf_demo;
CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `number` varchar(255) NOT NULL DEFAULT '' COMMENT '学号',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender` char(5)  NOT NULL COMMENT '男｜女｜未公开',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `number_unique` (`number`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;

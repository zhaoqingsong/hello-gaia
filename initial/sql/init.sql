-- -- AppContainerUploaderService
-- DROP TABLE IF EXISTS `file_list`;
-- CREATE TABLE IF NOT EXISTS `file_list` (
--   `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID,和业务无关',
--   `name` varchar(8) NOT NULL DEFAULT '' COMMENT '文件名',
--   `url` varchar(512) NOT NULL DEFAULT '' COMMENT 'url',
--   PRIMARY KEY (`id`) COMMENT '自增主键，和业务无关',
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='file_list文件列表';
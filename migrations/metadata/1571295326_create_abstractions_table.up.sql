CREATE TABLE `abstractions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `label` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` int(11) NOT NULL,
  `aggregate_type` int(4) NOT NULL COMMENT '聚合类型',
  `aggregate_field` varchar(255) NOT NULL DEFAULT '' COMMENT '聚合字段',
  `aggregate_interval_type` int(11) NOT NULL COMMENT '聚合时间片类型',
  `aggregate_interval_value` int(11) NOT NULL COMMENT '聚合时间片长度',
  `filter_expression` varchar(1024) NOT NULL COMMENT '过滤条件表达式',
  `search_field` varchar(255) NOT NULL COMMENT '搜索字段',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `NAME_IDX` (`name`) USING BTREE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
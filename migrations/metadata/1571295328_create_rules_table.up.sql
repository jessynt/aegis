CREATE TABLE `rules` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `activation_id` int(11) NOT NULL,
  `abstraction_id` int(11) NOT NULL COMMENT '指标字段（为空时按总体计算 i.e. COUNT(*)）',
  `base_score` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '命中初始得分',
  `base_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '命中基数（配合运算符，与指标字段进行计算）',
  `operator` int(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '运算符',
  `expression` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '表达式',
  `rate` int(11) NOT NULL DEFAULT '0' COMMENT '比率',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
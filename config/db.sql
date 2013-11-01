DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
      `id` bigint(11) NOT NULL AUTO_INCREMENT,
      `name` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
      `salt` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
      `passwd` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
      `email` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
      `created_at` bigint(20) unsigned NOT NULL DEFAULT '0',
      `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`),
      UNIQUE KEY `UNIQ_8D93D649E7927C74` (`email`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

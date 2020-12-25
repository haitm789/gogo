-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ac_user_id` int(10) unsigned DEFAULT NULL,
  `email` varchar(64) NOT NULL,
  `encrypted_password` varchar(512) NOT NULL,
  `card_status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `confirmed_unixtime` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `users`;
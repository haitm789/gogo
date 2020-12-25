-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `mails` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mail_template_id` int(10) unsigned NOT NULL DEFAULT '0',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `params` json DEFAULT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `send_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `done_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `mails`;
-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `mail_templates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subject` varchar(512) NOT NULL DEFAULT '',
  `content` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `mail_templates`;
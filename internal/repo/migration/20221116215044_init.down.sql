ALTER TABLE `worker_bot` DROP FOREIGN KEY `worker_bot_fk0`;
ALTER TABLE `session` DROP FOREIGN KEY `session_fk0`;
ALTER TABLE `session` DROP FOREIGN KEY `session_fk1`;


DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `worker_bot`;
DROP TABLE IF EXISTS `session`;
DROP TABLE IF EXISTS `application`;
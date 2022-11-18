CREATE TABLE `user` (
	`id` int NOT NULL AUTO_INCREMENT,
	`username` varchar(255) NOT NULL,
	`sessions_count` int DEFAULT '0',
	PRIMARY KEY (`id`)
);

CREATE TABLE `worker_bot` (
	`id` int NOT NULL AUTO_INCREMENT,
	`botname` varchar(255),
	`dialog_data` json,
	`ref_user_id` int,
	PRIMARY KEY (`id`)
);

CREATE TABLE `session` (
	`id` int NOT NULL AUTO_INCREMENT,
	`phone` varchar(255),
	`confirmed` BOOLEAN DEFAULT '0',
	`is2fa` BOOLEAN DEFAULT '0',
	`ref_user_id` int,
	`ref_worker_bot_id` int,
	PRIMARY KEY (`id`)
);

CREATE TABLE `application` (
	`id` int NOT NULL AUTO_INCREMENT,
	`key_id` varchar(255) NOT NULL UNIQUE,
	`tag` varchar(255),
	`invited` BOOLEAN DEFAULT '0',
	PRIMARY KEY (`id`)
);

ALTER TABLE `worker_bot` ADD CONSTRAINT `worker_bot_fk0` FOREIGN KEY (`ref_user_id`) REFERENCES `user`(`id`);

ALTER TABLE `session` ADD CONSTRAINT `session_fk0` FOREIGN KEY (`ref_user_id`) REFERENCES `user`(`id`);

ALTER TABLE `session` ADD CONSTRAINT `session_fk1` FOREIGN KEY (`ref_worker_bot_id`) REFERENCES `worker_bot`(`id`);





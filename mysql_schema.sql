-- At Ubuntu 22, install workbench as follow
-- sudo snap install mysql-workbench-community
-- sudo snap connect mysql-workbench-community:password-manager-service :password-manager-service
-- Then paste below stmt into query box

create	schema `users_db` default character set utf8;
CREATE TABLE `users_db`.`users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NOT NULL,
  `date_created` VARCHAR(45) NULL,
  `status` VARCHAR(45) NOT NULL,
  `password` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);

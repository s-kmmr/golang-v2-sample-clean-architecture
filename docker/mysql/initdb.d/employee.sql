DROP DATABASE IF EXISTS pracv2;
CREATE DATABASE IF NOT EXISTS pracv2;
use pracv2;

DROP TABLE IF EXISTS `employee`;


CREATE TABLE IF NOT EXISTS `employee`
(
 `id`               INT(20) AUTO_INCREMENT,
 `employee_num`     VARCHAR(20) NOT NULL,
 `department_num`   INT(20) NOT NULL,
 `first_name`       VARCHAR(20) NOT NULL,
 `last_name`        VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO employee (employee_num,department_num,first_name,last_name) VALUES ("E001",1,'ken','ito');
INSERT INTO employee (employee_num,department_num,first_name,last_name) VALUES ("E002",2,'taro','kato');
INSERT INTO employee (employee_num,department_num,first_name,last_name) VALUES ("E003",1,'jiro','isida');
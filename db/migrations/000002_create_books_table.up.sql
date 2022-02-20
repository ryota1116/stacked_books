CREATE TABLE IF NOT EXISTS books (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT UNIQUE NOT NULL,
    `google_books_id` VARCHAR(20) UNIQUE NOT NULL,
    `title` VARCHAR (255) UNIQUE NOT NULL,
    `description` VARCHAR (255),
    `image` VARCHAR(255),
    `isbn_10` VARCHAR(10),
    `isbn_13` VARCHAR(13),
    `page_count` INT(11),
    `published_year` INT(4) DEFAULT NULL,
    `published_month` INT(2) DEFAULT NULL,
    `published_date` INT(2) DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8;
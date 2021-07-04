CREATE TABLE IF NOT EXISTS user_books (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT UNIQUE NOT NULL,
    `user_id` INT(11) UNSIGNED NOT NULL,
    `book_id` INT(11) UNSIGNED NOT NULL,
    `status` INT  NOT NULL DEFAULT 0,
    `memo` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    UNIQUE KEY `unique_of_user_id_and_book_id` (`user_id`, `book_id`),
    KEY `index_of_user_id` (`user_id`),
    KEY `index_of_book_id` (`book_id`),
    CONSTRAINT `fk_user_id_of_user_books` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_book_id_of_user_books` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8;
-- user_booksテーブルのstatusカラムのデフォルト値を0から1に変更する
-- GoのEnum定義に合わせる
ALTER TABLE `user_books`
    MODIFY `status` INT(11) NOT NULL DEFAULT 1;

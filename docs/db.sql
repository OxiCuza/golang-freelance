CREATE TABLE `users` (
  `id` varchar(255) PRIMARY KEY,
  `name` varchar(255),
  `email` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `blog_posts` (
  `id` varchar(255) PRIMARY KEY,
  `title` varchar(255),
  `content` text,
  `user_id` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `blog_posts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

-- ginblog.article definition

CREATE TABLE `article` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                           `created_at` datetime(3) DEFAULT NULL,
                           `updated_at` datetime(3) DEFAULT NULL,
                           `deleted_at` datetime(3) DEFAULT NULL,
                           `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `cid` bigint unsigned NOT NULL,
                           `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                           `img` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `comment_count` bigint NOT NULL DEFAULT '0',
                           `read_count` bigint NOT NULL DEFAULT '0',
                           PRIMARY KEY (`id`) USING BTREE,
                           KEY `idx_article_deleted_at` (`deleted_at`) USING BTREE,
                           KEY `fk_article_category` (`cid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=579 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- ginblog.category definition

CREATE TABLE `category` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- ginblog.comment definition

CREATE TABLE `comment` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                           `created_at` datetime(3) DEFAULT NULL,
                           `updated_at` datetime(3) DEFAULT NULL,
                           `deleted_at` datetime(3) DEFAULT NULL,
                           `user_id` bigint unsigned DEFAULT NULL,
                           `article_id` bigint unsigned DEFAULT NULL,
                           `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `status` tinyint DEFAULT '2',
                           `article_title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                           `username` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                           `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                           PRIMARY KEY (`id`) USING BTREE,
                           KEY `idx_comment_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- ginblog.profile definition

CREATE TABLE `profile` (
                           `id` bigint NOT NULL AUTO_INCREMENT,
                           `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `qqchat` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `wechat` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `weibo` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `bili` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `img` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=574 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- ginblog.`user` definition

CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `created_at` datetime(3) DEFAULT NULL,
                        `updated_at` datetime(3) DEFAULT NULL,
                        `deleted_at` datetime(3) DEFAULT NULL,
                        `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                        `password` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                        `role` bigint DEFAULT '2',
                        PRIMARY KEY (`id`) USING BTREE,
                        KEY `idx_user_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
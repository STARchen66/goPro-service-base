DROP TABLE IF EXISTS `base_users`;
CREATE TABLE `base_users`  (
                               `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                               `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                               `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                               `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                               `created_at` datetime NULL DEFAULT NULL,
                               `updated_at` datetime NULL DEFAULT NULL,
                               `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '头像',
                               `oauth_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '第三方id',
                               `bound_oauth` tinyint(1) NULL DEFAULT 0 COMMENT '1\\github 2\\gitee',
                               `oauth_type` tinyint(1) NULL DEFAULT NULL COMMENT '1.微博 2.github',
                               `status` tinyint(1) NULL DEFAULT 0 COMMENT '0 离线 1 在线',
                               `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户简介',
                               `sex` tinyint(1) NULL DEFAULT 0 COMMENT '0 未知 1.男 2.女',
                               `client_type` tinyint(1) NULL DEFAULT NULL COMMENT '1.web 2.pc 3.app',
                               `age` int(3) NULL DEFAULT NULL,
                               `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
                               `uid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT 'uid 关联',
                               `user_json` json NULL,
                               `gitee` int(3) NULL DEFAULT NULL,
                               `github` int(3) NULL DEFAULT NULL,
                               `user_type` int(3) NULL DEFAULT NULL,

                               `github_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                               `gitee_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                               `gitee_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                               `github_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,


                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

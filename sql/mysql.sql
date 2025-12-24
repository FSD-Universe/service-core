-- Copyright (c) 2025 Half_nothing
-- SPDX-License-Identifier: MIT

CREATE TABLE `users`
(
    `id`               INT UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `username`         VARCHAR(64)     NOT NULL COMMENT '用户名',
    `email`            VARCHAR(128)    NOT NULL COMMENT '用户邮箱',
    `cid`              INT UNSIGNED    NOT NULL COMMENT '用户CID',
    `password`         VARCHAR(128)    NOT NULL COMMENT '用户密码(bcrypt加密)',
    `image_id`         INT UNSIGNED    NULL     DEFAULT NULL COMMENT '用户头像索引',
    `qq`               VARCHAR(16)     NULL     DEFAULT NULL COMMENT '用户QQ',
    `banned`           TINYINT(1)      NOT NULL DEFAULT FALSE COMMENT '用户被封禁',
    `banned_until`     DATETIME        NULL     DEFAULT NULL COMMENT '用户封禁到期时间',
    `rating`           INT             NOT NULL DEFAULT 0 COMMENT '用户管制权限',
    `permission`       BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户飞控权限',
    `total_pilot_time` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户连线时间',
    `last_login_time`  DATETIME        NULL     DEFAULT NULL COMMENT '上次登录时间',
    `last_login_ip`    VARCHAR(128)    NULL     DEFAULT NULL COMMENT '上次登录IP',
    `created_at`       DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
    `updated_at`       DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '信息更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_users_username` (`username`) USING BTREE COMMENT '用户名索引',
    UNIQUE INDEX `idx_users_email` (`email`) USING BTREE COMMENT '用户邮箱索引',
    UNIQUE INDEX `idx_users_cid` (`cid`) USING BTREE COMMENT '用户CID索引'
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '用户信息表';

CREATE TABLE `images`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '图片id',
    `user_id`    INT UNSIGNED NOT NULL COMMENT '上传本图片的用户id',
    `hashcode`   VARCHAR(128) NOT NULL COMMENT '哈希值',
    `filename`   VARCHAR(128) NOT NULL COMMENT '文件名',
    `url`        TEXT         NOT NULL COMMENT '访问路径',
    `size`       BIGINT       NOT NULL DEFAULT 0 COMMENT '文件大小',
    `mime_type`  VARCHAR(128) NOT NULL COMMENT '文件元类型',
    `comment`    TEXT         NOT NULL COMMENT '文件备注',
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME     NULL     DEFAULT NULL COMMENT '软删除',
    PRIMARY KEY (`id`),
    INDEX `idx_images_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    CONSTRAINT `fk_images_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '图片存储表';

CREATE TABLE `histories`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '连线历史id',
    `user_id`     INT UNSIGNED NOT NULL COMMENT '用户id',
    `callsign`    VARCHAR(16)  NOT NULL COMMENT '呼号',
    `start_time`  DATETIME     NOT NULL COMMENT '上线时间',
    `end_time`    DATETIME     NOT NULL COMMENT '下线时间',
    `online_time` INT          NOT NULL COMMENT '在线时间(秒)',
    `is_atc`      TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '是否为管制',
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_histories_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    CONSTRAINT `fk_histories_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '用户连线历史表';

CREATE TABLE `roles`
(
    `id`         INT UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '角色id',
    `name`       VARCHAR(64)     NOT NULL COMMENT '角色名',
    `permission` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色权限',
    `comment`    TEXT            NOT NULL COMMENT '角色备注',
    `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '角色表';

CREATE TABLE `user_roles`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`    INT UNSIGNED NOT NULL COMMENT '用户id',
    `role_id`    INT UNSIGNED NOT NULL COMMENT '角色id',
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_user_roles_user_role` (`user_id`, `role_id`) USING BTREE COMMENT '用户角色索引',
    INDEX `idx_user_roles_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_user_roles_role_id` (`role_id`) USING BTREE COMMENT '角色id索引',
    CONSTRAINT `fk_user_roles_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_user_roles_role_id` FOREIGN KEY (`role_id`)
        REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '用户角色授予表';

CREATE TABLE `announcements`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '公告id',
    `user_id`    INT UNSIGNED NOT NULL COMMENT '发布用户id',
    `title`      TEXT         NOT NULL COMMENT '公告标题',
    `content`    TEXT         NOT NULL COMMENT '公告内容',
    `type`       INT          NOT NULL DEFAULT 0 COMMENT '公告类型',
    `important`  TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '重要公告',
    `force_show` TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '主动弹出',
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at` DATETIME     NULL     DEFAULT NULL COMMENT '软删除',
    PRIMARY KEY (`id`),
    INDEX `idx_announcements_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    CONSTRAINT `fk_announcements_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '公告表';

CREATE TABLE `audit_logs`
(
    `id`             INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `subject`        TEXT         NOT NULL COMMENT '操作来源实体',
    `object`         TEXT         NOT NULL COMMENT '操作对象实体',
    `event`          VARCHAR(64)  NOT NULL COMMENT '事件名',
    `ip`             VARCHAR(128) NOT NULL COMMENT '操作来源IP',
    `user_agent`     TEXT         NOT NULL COMMENT '操作来源用户代理',
    `change_details` TEXT         NULL COMMENT '更改明细',
    `created_at`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
    PRIMARY KEY (`id`),
    INDEX `idx_audit_logs_event` (`event`) USING BTREE COMMENT '事件索引'
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '审计日志表';

CREATE TABLE `flight_plans`
(
    `id`                  INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`             INT UNSIGNED NOT NULL COMMENT '用户id',
    `callsign`            VARCHAR(16)  NOT NULL COMMENT '呼号',
    `flight_rule`         VARCHAR(4)   NOT NULL COMMENT '飞行规则',
    `aircraft`            VARCHAR(128) NOT NULL COMMENT '机型',
    `tas`                 INT          NOT NULL DEFAULT 0 COMMENT '巡航真空速',
    `departure_airport`   VARCHAR(8)   NOT NULL COMMENT '离场机场',
    `plan_departure_time` VARCHAR(8)   NOT NULL DEFAULT '0000' COMMENT '计划离场时间',
    `atc_departure_time`  VARCHAR(8)   NOT NULL DEFAULT '0000' COMMENT 'ATC分配离场时间',
    `cruise_altitude`     VARCHAR(8)   NOT NULL COMMENT '巡航高度',
    `arrival_airport`     VARCHAR(8)   NOT NULL COMMENT '到达机场',
    `route_time_hour`     VARCHAR(2)   NOT NULL DEFAULT '00' COMMENT '航路耗时(小时)',
    `route_time_minute`   VARCHAR(2)   NOT NULL DEFAULT '00' COMMENT '航路耗时(分钟)',
    `air_time_hour`       VARCHAR(2)   NOT NULL DEFAULT '00' COMMENT '滞空时间(小时)',
    `air_time_minute`     VARCHAR(2)   NOT NULL DEFAULT '00' COMMENT '滞空时间(分钟)',
    `alternate_airport`   VARCHAR(8)   NOT NULL COMMENT '备降机场',
    `remarks`             TEXT         NOT NULL COMMENT '备注',
    `route`               TEXT         NOT NULL COMMENT '航路',
    `locked`              TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '计划锁定',
    `from_web`            TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '从网页提交',
    `created_at`          DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_flight_plans_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_flight_plans_callsign` (`callsign`) USING BTREE COMMENT '呼号索引',
    CONSTRAINT `fk_flight_plans_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '飞行计划表';

CREATE TABLE `instructors`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`    INT UNSIGNED NOT NULL COMMENT '用户id',
    `email`      VARCHAR(128) NOT NULL COMMENT '教员邮箱',
    `twr`        INT          NOT NULL DEFAULT 0 COMMENT '塔台教员:0-无,1-教师,2-教员',
    `app`        INT          NOT NULL DEFAULT 0 COMMENT '进近教员:0-无,1-教师,2-教员',
    `ctr`        INT          NOT NULL DEFAULT 0 COMMENT '区域教员:0-无,1-教师,2-教员',
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '信息更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_instructors_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    CONSTRAINT `fk_instructors_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '教员信息表';

CREATE TABLE `controller_applications`
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`       INT UNSIGNED NOT NULL COMMENT '用户id',
    `reason`        TEXT         NOT NULL COMMENT '申请理由',
    `record`        TEXT         NOT NULL COMMENT '管制经历',
    `is_guest`      TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '是否为客座',
    `platform`      VARCHAR(16)  NOT NULL COMMENT '客座平台',
    `image_id`      INT UNSIGNED NULL     DEFAULT NULL COMMENT '客座证明资料',
    `status`        INT          NOT NULL DEFAULT 0 COMMENT '申请状态',
    `message`       TEXT         NULL     DEFAULT NULL COMMENT '回复消息',
    `instructor_id` INT UNSIGNED NULL     DEFAULT NULL COMMENT '面试教员id',
    `created_at`    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_controller_applications_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_controller_applications_image_id` (`image_id`) USING BTREE COMMENT '图片id索引',
    INDEX `idx_controller_applications_instructor_id` (`instructor_id`) USING BTREE COMMENT '教员id索引',
    CONSTRAINT `fk_controller_applications_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_controller_applications_image_id` FOREIGN KEY (`image_id`)
        REFERENCES `images` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_controller_applications_instructor_id` FOREIGN KEY (`instructor_id`)
        REFERENCES `instructors` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '管制员申请表';

CREATE TABLE `controller_application_times`
(
    `id`             INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `application_id` INT UNSIGNED NOT NULL COMMENT '申请id',
    `time`           DATETIME     NOT NULL COMMENT '申请时间',
    `selected`       TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '是否被选中',
    PRIMARY KEY (`id`),
    INDEX `idx_controller_application_times_application_id` (`application_id`) USING BTREE COMMENT '申请id索引',
    CONSTRAINT `fk_controller_application_times_application_id` FOREIGN KEY (`application_id`)
        REFERENCES `controller_applications` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '管制员申请时间表';

CREATE TABLE `controllers`
(
    `id`                    INT UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`               INT UNSIGNED    NOT NULL COMMENT '用户id',
    `instructor_id`         INT UNSIGNED    NULL     DEFAULT NULL COMMENT '教员id',
    `guest`                 TINYINT(1)      NOT NULL DEFAULT FALSE COMMENT '是否为客座管制员',
    `under_monitor`         TINYINT(1)      NOT NULL DEFAULT FALSE COMMENT '是否为实习管制员',
    `under_solo`            TINYINT(1)      NOT NULL DEFAULT FALSE COMMENT '是否SOLO',
    `solo_until`            DATETIME        NULL     DEFAULT NULL COMMENT 'SOLO时限',
    `tier2`                 TINYINT(1)      NOT NULL DEFAULT FALSE COMMENT '是否有程序塔权限',
    `total_controller_time` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总管制时长',
    `created_at`            DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
    `updated_at`            DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '信息更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_controllers_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_controllers_instructor_id` (`instructor_id`) USING BTREE COMMENT '教员id索引',
    CONSTRAINT `fk_controllers_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_controllers_instructor_id` FOREIGN KEY (`instructor_id`)
        REFERENCES `instructors` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '管制员表';

CREATE TABLE `controller_records`
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`       INT UNSIGNED NOT NULL COMMENT '用户id',
    `instructor_id` INT UNSIGNED NOT NULL COMMENT '教员id',
    `type`          INT          NOT NULL DEFAULT 0 COMMENT '履历类型',
    `content`       TEXT         NOT NULL COMMENT '履历内容',
    `created_at`    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
    `deleted_at`    DATETIME     NULL     DEFAULT NULL COMMENT '软删除',
    PRIMARY KEY (`id`),
    INDEX `idx_controller_records_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_controller_records_instructor_id` (`instructor_id`) USING BTREE COMMENT '教员id索引',
    CONSTRAINT `fk_controller_records_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_controller_records_instructor_id` FOREIGN KEY (`instructor_id`)
        REFERENCES `instructors` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '管制员履历表';

CREATE TABLE `tickets`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`    INT UNSIGNED NOT NULL COMMENT '用户id',
    `type`       INT          NOT NULL DEFAULT 0 COMMENT '工单类型',
    `title`      TEXT         NOT NULL COMMENT '工单标题',
    `content`    TEXT         NOT NULL COMMENT '工单内容',
    `reply`      TEXT         NULL     DEFAULT NULL COMMENT '工单回复',
    `replier`    INT UNSIGNED NULL     DEFAULT NULL COMMENT '工单回复人id',
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `closed_at`  DATETIME     NULL     DEFAULT NULL COMMENT '工单关闭时间',
    PRIMARY KEY (`id`),
    INDEX `idx_tickets_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_tickets_replier_id` (`replier`) USING BTREE COMMENT '工单回复人id索引',
    CONSTRAINT `fk_tickets_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_tickets_replier_id` FOREIGN KEY (`replier`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '工单表';

CREATE TABLE `activities`
(
    `id`                INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `publisher_id`      INT UNSIGNED NOT NULL COMMENT '发布者id',
    `type`              INT          NOT NULL DEFAULT 0 COMMENT '活动类型',
    `title`             TEXT         NOT NULL COMMENT '活动标题',
    `image_id`          INT UNSIGNED NULL     DEFAULT NULL COMMENT '活动图片id',
    `active_time`       DATETIME     NOT NULL COMMENT '活动时间',
    `departure_airport` VARCHAR(64)  NOT NULL COMMENT '离场机场',
    `arrival_airport`   VARCHAR(64)  NOT NULL COMMENT '到达机场',
    `route`             TEXT         NULL     DEFAULT NULL COMMENT '活动航路',
    `distance`          INT          NULL     DEFAULT NULL COMMENT '活动距离',
    `second_route`      TEXT         NULL     DEFAULT NULL COMMENT '第二活动航路',
    `second_distance`   INT          NULL     DEFAULT NULL COMMENT '第二活动距离',
    `open_fir`          VARCHAR(128) NULL     DEFAULT NULL COMMENT '空域开放日',
    `status`            INT          NOT NULL DEFAULT 0 COMMENT '活动状态',
    `notams`            TEXT         NOT NULL COMMENT '航行通告',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        DATETIME     NULL     DEFAULT NULL COMMENT '软删除',
    PRIMARY KEY (`id`),
    INDEX `idx_activities_publisher_id` (`publisher_id`) USING BTREE COMMENT '发布者id索引',
    INDEX `idx_activities_image_id` (`image_id`) USING BTREE COMMENT '图片id索引',
    CONSTRAINT `fk_activities_publisher_id` FOREIGN KEY (`publisher_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_activities_image_id` FOREIGN KEY (`image_id`)
        REFERENCES `images` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '活动表';

CREATE TABLE `activity_pilots`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     INT UNSIGNED NOT NULL COMMENT '飞行员id',
    `activity_id` INT UNSIGNED NOT NULL COMMENT '活动id',
    `callsign`    VARCHAR(64)  NOT NULL COMMENT '呼号',
    `aircraft`    VARCHAR(64)  NOT NULL COMMENT '机型',
    `status`      INT          NOT NULL DEFAULT 0 COMMENT '飞行员状态',
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '报名时间',
    `updated_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_activity_pilots_activity_id` (`activity_id`) USING BTREE COMMENT '活动id索引',
    INDEX `idx_activity_pilots_user_id` (`user_id`) USING BTREE COMMENT '飞行员id索引',
    CONSTRAINT `fk_activity_pilots_activity_id` FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_activity_pilots_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '活动飞行员表';

CREATE TABLE `activity_facilities`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `activity_id` INT UNSIGNED NOT NULL COMMENT '活动id',
    `min_rating`  INT          NOT NULL DEFAULT 0 COMMENT '最低等级',
    `callsign`    VARCHAR(32)  NOT NULL COMMENT '呼号',
    `frequency`   VARCHAR(32)  NOT NULL COMMENT '频率',
    `tier2`       TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT '是否为程序塔台',
    `sort_index`  INT          NOT NULL DEFAULT 0 COMMENT '排序索引',
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_activity_facilities_activity_id` (`activity_id`) USING BTREE COMMENT '活动id索引',
    CONSTRAINT `fk_activity_facilities_activity_id` FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '活动席位表';

CREATE TABLE `activity_controllers`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     INT UNSIGNED NOT NULL COMMENT '用户id',
    `activity_id` INT UNSIGNED NOT NULL COMMENT '活动id',
    `facility_id` INT UNSIGNED NOT NULL COMMENT '活动席位id',
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_activity_controllers_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    INDEX `idx_activity_controllers_activity_id` (`activity_id`) USING BTREE COMMENT '活动id索引',
    INDEX `idx_activity_controllers_facility_id` (`facility_id`) USING BTREE COMMENT '活动席位id索引',
    CONSTRAINT `fk_activity_controllers_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_activity_controllers_activity_id` FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_activity_controllers_facility_id` FOREIGN KEY (`facility_id`)
        REFERENCES `activity_facilities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = `utf8mb4` COLLATE = `utf8mb4_general_ci` COMMENT = '活动管制员表';

CREATE TABLE `activity_records`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `activity_id` INT UNSIGNED NOT NULL COMMENT '活动id',
    `user_id`     INT UNSIGNED NOT NULL COMMENT '用户id',
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_activity_records_activity_id` (`activity_id`) USING BTREE COMMENT '活动id索引',
    INDEX `idx_activity_records_user_id` (`user_id`) USING BTREE COMMENT '用户id索引',
    CONSTRAINT `fk_activity_records_activity_id` FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_activity_records_user_id` FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = innodb CHARACTER SET `utf8mb4` COLLATE `utf8mb4_general_ci` COMMENT = '活动记录表';

CREATE TABLE `activity_coordinations`
(
    `id`                INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `activity_id`       INT UNSIGNED NOT NULL COMMENT '活动id',
    `facility_id`       INT UNSIGNED NOT NULL COMMENT '活动席位id',
    `logon_code`        VARCHAR(8)   NULL     DEFAULT NULL COMMENT 'CPDLC/DCL识别码',
    `logon_network`     VARCHAR(64)  NULL     DEFAULT NULL COMMENT 'CPDLC/DCL登陆网络',
    `pdc_available`     TINYINT(1)   NOT NULL DEFAULT FALSE COMMENT 'PDC是否可用',
    `transform`         VARCHAR(64)  NULL     DEFAULT NULL COMMENT '移交席位',
    `transform_remarks` TEXT         NULL     DEFAULT NULL COMMENT '移交备注',
    `procedure`         VARCHAR(64)  NULL     DEFAULT NULL COMMENT '程序',
    `runway`            VARCHAR(64)  NULL     DEFAULT NULL COMMENT '跑道',
    `runway_remarks`    TEXT         NULL     DEFAULT NULL COMMENT '跑道运行备注',
    `altitude`          VARCHAR(64)  NULL     DEFAULT NULL COMMENT '起始高度',
    `remarks`           TEXT         NULL     DEFAULT NULL COMMENT '其他备注信息',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_activity_coordinations_activity_id` (`activity_id`) USING BTREE COMMENT '活动id索引',
    INDEX `idx_activity_coordinations_facility_id` (`facility_id`) USING BTREE COMMENT '活动席位id索引',
    CONSTRAINT `fk_activity_coordinations_activity_id` FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_activity_coordinations_facility_id` FOREIGN KEY (`facility_id`)
        REFERENCES `activity_facilities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = innodb CHARACTER SET `utf8mb4` COLLATE `utf8mb4_general_ci` COMMENT = '活动协调表';
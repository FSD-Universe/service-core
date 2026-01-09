-- Copyright (c) 2025-2026 Half_nothing
-- SPDX-License-Identifier: MIT

-- 用户信息表
CREATE TABLE `users`
(
    `id`               INTEGER PRIMARY KEY AUTOINCREMENT,          -- 用户id
    `username`         TEXT    NOT NULL,                           -- 用户名
    `email`            TEXT    NOT NULL,                           -- 用户邮箱
    `cid`              INTEGER NOT NULL,                           -- 用户CID
    `password`         TEXT    NOT NULL,                           -- 用户密码(bcrypt加密)
    `image_id`         INTEGER,                                    -- 用户头像索引
    `qq`               TEXT,                                       -- 用户QQ
    `banned`           TEXT,
    `rating`           INTEGER NOT NULL DEFAULT 0,                 -- 用户管制权限
    `permission`       INTEGER NOT NULL DEFAULT 0,                 -- 用户飞控权限
    `total_pilot_time` INTEGER NOT NULL DEFAULT 0,                 -- 用户连线时间
    `created_at`       TEXT             DEFAULT CURRENT_TIMESTAMP, -- 注册时间
    `updated_at`       TEXT             DEFAULT CURRENT_TIMESTAMP, -- 信息更新时间
    `last_login_time`  TEXT,                                       -- 上次登录时间
    `last_login_ip`    TEXT                                        -- 上次登录IP
);

CREATE UNIQUE INDEX `idx_users_username` ON `users` (`username`);
CREATE UNIQUE INDEX `idx_users_email` ON `users` (`email`);
CREATE UNIQUE INDEX `idx_users_cid` ON `users` (`cid`);

-- 图片存储表
CREATE TABLE `images`
(
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,          -- 图片id
    `user_id`    INTEGER NOT NULL,                           -- 上传本图片的用户id
    `hashcode`   TEXT    NOT NULL,                           -- 哈希值
    `filename`   TEXT    NOT NULL,                           -- 文件名
    `url`        TEXT    NOT NULL,                           -- 访问路径
    `size`       INTEGER NOT NULL DEFAULT 0,                 -- 文件大小
    `mime_type`  TEXT    NOT NULL,                           -- 文件元类型
    `comment`    TEXT    NOT NULL,                           -- 文件备注
    `created_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 上传时间
    `updated_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `deleted_at` TEXT,                                       -- 软删除
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_images_user_id` ON `images` (`user_id`);

-- 用户连线历史表
CREATE TABLE `histories`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT,          -- 连线历史id
    `user_id`     INTEGER NOT NULL,                           -- 用户id
    `callsign`    TEXT    NOT NULL,                           -- 呼号
    `start_time`  TEXT    NOT NULL,                           -- 上线时间
    `end_time`    TEXT    NOT NULL,                           -- 下线时间
    `online_time` INTEGER NOT NULL,                           -- 在线时间(秒)
    `is_atc`      INTEGER NOT NULL DEFAULT FALSE,             -- 是否为管制
    `created_at`  TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_histories_user_id` ON `histories` (`user_id`);

-- 角色表
CREATE TABLE `roles`
(
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,          -- 角色id
    `name`       TEXT    NOT NULL,                           -- 角色名
    `permission` INTEGER NOT NULL DEFAULT 0,                 -- 角色权限
    `comment`    TEXT    NOT NULL,                           -- 角色备注
    `created_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` TEXT             DEFAULT CURRENT_TIMESTAMP  -- 更新时间
);

-- 用户角色授予表
CREATE TABLE `user_roles`
(
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键
    `user_id`    INTEGER NOT NULL,                  -- 用户id
    `role_id`    INTEGER NOT NULL,                  -- 角色id
    `created_at` TEXT DEFAULT CURRENT_TIMESTAMP,    -- 创建时间
    UNIQUE (`user_id`, `role_id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`role_id`)
        REFERENCES `roles` (`id`) ON DELETE RESTRICT
);

CREATE UNIQUE INDEX `idx_user_roles_user_role` ON `user_roles` (`user_id`, `role_id`);
CREATE INDEX `idx_user_roles_user_id` ON `user_roles` (`user_id`);
CREATE INDEX `idx_user_roles_role_id` ON `user_roles` (`role_id`);

-- 公告表
CREATE TABLE `announcements`
(
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,          -- 公告id
    `user_id`    INTEGER NOT NULL,                           -- 发布用户id
    `title`      TEXT    NOT NULL,                           -- 公告标题
    `content`    TEXT    NOT NULL,                           -- 公告内容
    `type`       INTEGER NOT NULL DEFAULT 0,                 -- 公告类型
    `important`  INTEGER NOT NULL DEFAULT FALSE,             -- 重要公告
    `force_show` INTEGER NOT NULL DEFAULT FALSE,             -- 主动弹出
    `created_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 修改时间
    `deleted_at` TEXT,                                       -- 软删除
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_announcements_user_id` ON `announcements` (`user_id`);

-- 审计日志表
CREATE TABLE `audit_logs`
(
    `id`             INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键
    `subject`        TEXT NOT NULL,                     -- 操作来源实体
    `object`         TEXT NOT NULL,                     -- 操作对象实体
    `event`          TEXT NOT NULL,                     -- 事件名
    `ip`             TEXT NOT NULL,                     -- 操作来源IP
    `user_agent`     TEXT NOT NULL,                     -- 操作来源用户代理
    `change_details` TEXT,                              -- 更改明细
    `created_at`     TEXT DEFAULT CURRENT_TIMESTAMP     -- 记录时间
);

CREATE INDEX `idx_audit_logs_event` ON `audit_logs` (`event`);

-- 飞行计划表
CREATE TABLE `flight_plans`
(
    `id`                  INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `user_id`             INTEGER NOT NULL,                           -- 用户id
    `callsign`            TEXT    NOT NULL,                           -- 呼号
    `flight_rule`         TEXT    NOT NULL,                           -- 飞行规则
    `aircraft`            TEXT    NOT NULL,                           -- 机型
    `tas`                 INTEGER NOT NULL DEFAULT 0,                 -- 巡航真空速
    `departure_airport`   TEXT    NOT NULL,                           -- 离场机场
    `plan_departure_time` TEXT    NOT NULL DEFAULT '0000',            -- 计划离场时间
    `atc_departure_time`  TEXT    NOT NULL DEFAULT '0000',            -- ATC分配离场时间
    `cruise_altitude`     TEXT    NOT NULL,                           -- 巡航高度
    `arrival_airport`     TEXT    NOT NULL,                           -- 到达机场
    `route_time_hour`     TEXT    NOT NULL DEFAULT '00',              -- 航路耗时(小时)
    `route_time_minute`   TEXT    NOT NULL DEFAULT '00',              -- 航路耗时(分钟)
    `air_time_hour`       TEXT    NOT NULL DEFAULT '00',              -- 滞空时间(小时)
    `air_time_minute`     TEXT    NOT NULL DEFAULT '00',              -- 滞空时间(分钟)
    `alternate_airport`   TEXT    NOT NULL,                           -- 备降机场
    `remarks`             TEXT    NOT NULL,                           -- 备注
    `route`               TEXT    NOT NULL,                           -- 航路
    `locked`              INTEGER NOT NULL DEFAULT FALSE,             -- 计划锁定
    `from_web`            INTEGER NOT NULL DEFAULT FALSE,             -- 从网页提交
    `created_at`          TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at`          TEXT             DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE UNIQUE INDEX `idx_flight_plans_user_id` ON `flight_plans` (`user_id`);
CREATE INDEX `idx_flight_plans_callsign` ON `flight_plans` (`callsign`);

-- 教员信息表
CREATE TABLE `instructors`
(
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `user_id`    INTEGER NOT NULL,                           -- 用户id
    `email`      TEXT    NOT NULL,                           -- 教员邮箱
    `twr`        INTEGER NOT NULL DEFAULT 0,                 -- 塔台教员:0-无,1-教师,2-教员
    `app`        INTEGER NOT NULL DEFAULT 0,                 -- 进近教员:0-无,1-教师,2-教员
    `ctr`        INTEGER NOT NULL DEFAULT 0,                 -- 区域教员:0-无,1-教师,2-教员
    `created_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 加入时间
    `updated_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 信息更新时间
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE UNIQUE INDEX `idx_instructors_user_id` ON `instructors` (`user_id`);

-- 管制员申请表
CREATE TABLE `controller_applications`
(
    `id`            INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `user_id`       INTEGER NOT NULL,                           -- 用户id
    `reason`        TEXT    NOT NULL,                           -- 申请理由
    `record`        TEXT    NOT NULL,                           -- 管制经历
    `is_guest`      INTEGER NOT NULL DEFAULT FALSE,             -- 是否为客座
    `platform`      TEXT    NOT NULL,                           -- 客座平台
    `image_id`      INTEGER,                                    -- 客座证明资料
    `status`        INTEGER NOT NULL DEFAULT 0,                 -- 申请状态
    `message`       TEXT,                                       -- 回复消息
    `instructor_id` INTEGER,                                    -- 教员id
    `created_at`    TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at`    TEXT             DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`image_id`)
        REFERENCES `images` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`instructor_id`)
        REFERENCES `instructors` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_controller_applications_user_id` ON `controller_applications` (`user_id`);
CREATE INDEX `idx_controller_applications_image_id` ON `controller_applications` (`image_id`);
CREATE INDEX `idx_controller_applications_status` ON `controller_applications` (`status`);

-- 管制申请时间表
CREATE TABLE `controller_application_times`
(
    `id`             INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键
    `application_id` INTEGER NOT NULL,                  -- 申请id
    `time`           TEXT    NOT NULL,                  -- 申请时间
    `selected`       INTEGER NOT NULL DEFAULT FALSE,    -- 是否被选中
    FOREIGN KEY (`application_id`)
        REFERENCES `controller_applications` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_controller_application_times_application_id` ON `controller_application_times` (`application_id`);

-- 管制员表
CREATE TABLE `controllers`
(
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `user_id`               INTEGER NOT NULL,                           -- 用户id
    `instructor_id`         INTEGER,                                    -- 教员id
    `guest`                 INTEGER NOT NULL DEFAULT FALSE,             -- 是否为客座管制员
    `under_monitor`         INTEGER NOT NULL DEFAULT FALSE,             -- 是否为实习管制员
    `under_solo`            INTEGER NOT NULL DEFAULT FALSE,             -- 是否SOLO
    `solo_until`            TEXT,                                       -- SOLO时限
    `tier2`                 INTEGER NOT NULL DEFAULT FALSE,             -- 是否有程序塔权限
    `total_controller_time` INTEGER NOT NULL DEFAULT 0,                 -- 总管制时长
    `created_at`            TEXT             DEFAULT CURRENT_TIMESTAMP, -- 加入时间
    `updated_at`            TEXT             DEFAULT CURRENT_TIMESTAMP, -- 信息更新时间
    UNIQUE (`user_id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`instructor_id`)
        REFERENCES `instructors` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_controllers_instructor_id` ON `controllers` (`instructor_id`);

-- 管制员履历表
CREATE TABLE `controller_records`
(
    `id`            INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `user_id`       INTEGER NOT NULL,                           -- 用户id
    `instructor_id` INTEGER,                                    -- 教员id
    `type`          INTEGER NOT NULL DEFAULT 0,                 -- 履历类型
    `content`       TEXT    NOT NULL,                           -- 履历内容
    `created_at`    TEXT             DEFAULT CURRENT_TIMESTAMP, -- 记录时间
    `deleted_at`    TEXT,                                       -- 软删除
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`instructor_id`)
        REFERENCES `instructors` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_controller_records_user_id` ON `controller_records` (`user_id`);
CREATE INDEX `idx_controller_records_instructor_id` ON `controller_records` (`instructor_id`);

-- 工单表
CREATE TABLE `tickets`
(
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `user_id`    INTEGER NOT NULL,                           -- 用户id
    `type`       INTEGER NOT NULL DEFAULT 0,                 -- 工单类型
    `title`      TEXT    NOT NULL,                           -- 工单标题
    `content`    TEXT    NOT NULL,                           -- 工单内容
    `reply`      TEXT,                                       -- 工单回复
    `replier`    INTEGER,                                    -- 工单回复人id
    `created_at` TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `closed_at`  TEXT,                                       -- 工单关闭时间
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`replier`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_tickets_user_id` ON `tickets` (`user_id`);
CREATE INDEX `idx_tickets_replier_id` ON `tickets` (`replier`);

-- 活动表
CREATE TABLE `activities`
(
    `id`                INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `publisher_id`      INTEGER NOT NULL,                           -- 发布者id
    `type`              INTEGER NOT NULL DEFAULT 0,                 -- 活动类型
    `title`             TEXT    NOT NULL,                           -- 活动标题
    `image_id`          INTEGER,                                    -- 活动图片id
    `active_time`       TEXT    NOT NULL,                           -- 活动时间
    `departure_airport` TEXT    NOT NULL,                           -- 离场机场
    `arrival_airport`   TEXT    NOT NULL,                           -- 到达机场
    `route`             TEXT,                                       -- 活动路线
    `distance`          INTEGER,                                    -- 活动距离
    `second_route`      TEXT,                                       -- 活动路线2
    `second_distance`   INTEGER,                                    -- 活动距离2
    `open_fir`          TEXT,                                       -- 空域开放日
    `status`            INTEGER NOT NULL DEFAULT 0,                 -- 活动状态
    `notams`            TEXT    NOT NULL,                           -- 航行通告
    `created_at`        TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at`        TEXT             DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `deleted_at`        TEXT,                                       -- 软删除
    FOREIGN KEY (`publisher_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`image_id`)
        REFERENCES `images` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_activities_publisher_id` ON `activities` (`publisher_id`);
CREATE INDEX `idx_activities_image_id` ON `activities` (`image_id`);

-- 活动飞行员表
CREATE TABLE `activity_pilots`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `activity_id` INTEGER NOT NULL,                           -- 活动id
    `user_id`     INTEGER NOT NULL,                           -- 飞行员id
    `callsign`    TEXT    NOT NULL,                           -- 呼号
    `aircraft`    TEXT    NOT NULL,                           -- 机型
    `status`      INTEGER NOT NULL DEFAULT 0,                 -- 飞行员状态
    `created_at`  TEXT             DEFAULT CURRENT_TIMESTAMP, -- 报名时间
    `updated_at`  TEXT             DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_activity_pilots_activity_id` ON `activity_pilots` (`activity_id`);
CREATE INDEX `idx_activity_pilots_user_id` ON `activity_pilots` (`user_id`);

-- 活动席位表
CREATE TABLE `activity_facilities`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT,          -- 主键
    `activity_id` INTEGER NOT NULL,                           -- 活动id
    `min_rating`  INTEGER NOT NULL DEFAULT 0,                 -- 最低等级
    `callsign`    TEXT    NOT NULL,                           -- 呼号
    `frequency`   TEXT    NOT NULL,                           -- 频率
    `tier2`       INTEGER NOT NULL DEFAULT FALSE,             -- 是否为程序塔台
    `sort_index`  INTEGER NOT NULL DEFAULT 0,                 -- 排序索引
    `created_at`  TEXT             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at`  TEXT             DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_activity_facilities_activity_id` ON `activity_facilities` (`activity_id`);

-- 活动管制员表
CREATE TABLE `activity_controllers`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键
    `user_id`     INTEGER NOT NULL,                  -- 用户id
    `activity_id` INTEGER NOT NULL,                  -- 活动id
    `facility_id` INTEGER NOT NULL,                  -- 活动席位id
    `created_at`  TEXT DEFAULT CURRENT_TIMESTAMP,    -- 创建时间
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`facility_id`)
        REFERENCES `activity_facilities` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_activity_controllers_user_id` ON `activity_controllers` (`user_id`);
CREATE INDEX `idx_activity_controllers_activity_id` ON `activity_controllers` (`activity_id`);
CREATE INDEX `idx_activity_controllers_facility_id` ON `activity_controllers` (`facility_id`);

-- 活动记录表
CREATE TABLE `activity_records`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键
    `activity_id` INTEGER NOT NULL,                  -- 活动id
    `user_id`     INTEGER NOT NULL,                  -- 用户id
    `created_at`  TEXT DEFAULT CURRENT_TIMESTAMP,    -- 创建时间
    FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_activity_records_activity_id` ON `activity_records` (`activity_id`);
CREATE INDEX `idx_activity_records_user_id` ON `activity_records` (`user_id`);

-- 活动协调表
CREATE TABLE `activity_coordinations`
(
    `id`                INTEGER PRIMARY KEY AUTOINCREMENT,           --主键
    `activity_id`       INTEGER  NOT NULL,                           --活动id
    `facility_id`       INTEGER  NOT NULL,                           --活动席位id
    `logon_code`        TEXT,                                        --CPDLC/DCL识别码
    `logon_network`     TEXT,                                        --CPDLC/DCL登陆网络
    `pdc_available`     INTEGER  NOT NULL DEFAULT FALSE,             --PDC是否可用
    `transform`         TEXT,                                        --移交席位
    `transform_remarks` TEXT,                                        --移交备注
    `procedure`         TEXT,                                        --程序',
    `runway`            TEXT,                                        --跑道
    `runway_remarks`    TEXT,                                        --跑道运行备注
    `altitude`          TEXT,                                        --起始高度
    `remarks`           TEXT,                                        --其他备注信息
    `created_at`        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, --创建时间
    `updated_at`        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, --更新时间
    FOREIGN KEY (`activity_id`)
        REFERENCES `activities` (`id`) ON DELETE RESTRICT,
    FOREIGN KEY (`facility_id`)
        REFERENCES `activity_facilities` (`id`) ON DELETE RESTRICT
);

CREATE INDEX `idx_activity_coordinations_activity_id` ON `activity_coordinations` (`activity_id`);
CREATE INDEX `idx_activity_coordinations_facility_id` ON `activity_coordinations` (`facility_id`);

-- Copyright (c) 2025-2026 Half_nothing
-- SPDX-License-Identifier: MIT

-- 用户信息表
CREATE TABLE "users"
(
    "id"               BIGSERIAL PRIMARY KEY,
    "username"         VARCHAR(64)  NOT NULL,
    "email"            VARCHAR(128) NOT NULL,
    "cid"              INTEGER      NOT NULL,
    "password"         VARCHAR(128) NOT NULL,
    "image_id"         BIGINT       NULL,
    "qq"               VARCHAR(16)  NULL,
    "banned"           TIMESTAMP    NULL,
    "rating"           INTEGER      NOT NULL DEFAULT 0,
    "permission"       BIGINT       NOT NULL DEFAULT 0,
    "total_pilot_time" BIGINT       NOT NULL DEFAULT 0,
    "created_at"       TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "updated_at"       TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "last_login_time"  TIMESTAMP    NULL,
    "last_login_ip"    VARCHAR(128) NULL
);

COMMENT ON TABLE "users" IS '用户信息表';
COMMENT ON COLUMN "users"."id" IS '用户id';
COMMENT ON COLUMN "users"."username" IS '用户名';
COMMENT ON COLUMN "users"."email" IS '用户邮箱';
COMMENT ON COLUMN "users"."cid" IS '用户CID';
COMMENT ON COLUMN "users"."password" IS '用户密码(bcrypt加密)';
COMMENT ON COLUMN "users"."image_id" IS '用户头像索引';
COMMENT ON COLUMN "users"."qq" IS '用户QQ';
COMMENT ON COLUMN "users"."rating" IS '用户管制权限';
COMMENT ON COLUMN "users"."permission" IS '用户飞控权限';
COMMENT ON COLUMN "users"."total_pilot_time" IS '用户连线时间';
COMMENT ON COLUMN "users"."created_at" IS '注册时间';
COMMENT ON COLUMN "users"."updated_at" IS '信息更新时间';
COMMENT ON COLUMN "users"."last_login_time" IS '上次登录时间';
COMMENT ON COLUMN "users"."last_login_ip" IS '上次登录IP';

CREATE UNIQUE INDEX "idx_users_username" ON "users" ("username");
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
CREATE UNIQUE INDEX "idx_users_cid" ON "users" ("cid");

-- 图片存储表
CREATE TABLE "images"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "user_id"    BIGINT       NOT NULL,
    "hashcode"   VARCHAR(128) NOT NULL,
    "filename"   VARCHAR(128) NOT NULL,
    "url"        TEXT         NOT NULL,
    "size"       BIGINT       NOT NULL DEFAULT 0,
    "mime_type"  VARCHAR(128) NOT NULL,
    "comment"    TEXT         NOT NULL,
    "created_at" TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP    NULL
);

COMMENT ON TABLE "images" IS '图片存储表';
COMMENT ON COLUMN "images"."id" IS '图片id';
COMMENT ON COLUMN "images"."user_id" IS '上传本图片的用户id';
COMMENT ON COLUMN "images"."hashcode" IS '哈希值';
COMMENT ON COLUMN "images"."filename" IS '文件名';
COMMENT ON COLUMN "images"."url" IS '访问路径';
COMMENT ON COLUMN "images"."size" IS '文件大小';
COMMENT ON COLUMN "images"."mime_type" IS '文件元类型';
COMMENT ON COLUMN "images"."comment" IS '文件备注';
COMMENT ON COLUMN "images"."created_at" IS '上传时间';
COMMENT ON COLUMN "images"."updated_at" IS '更新时间';
COMMENT ON COLUMN "images"."deleted_at" IS '软删除';

CREATE INDEX "idx_images_user_id" ON "images" ("user_id");

ALTER TABLE "images"
    ADD CONSTRAINT "fk_images_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 用户连线历史表
CREATE TABLE "histories"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "user_id"     BIGINT      NOT NULL,
    "callsign"    VARCHAR(16) NOT NULL,
    "start_time"  TIMESTAMP   NOT NULL,
    "end_time"    TIMESTAMP   NOT NULL,
    "online_time" INTEGER     NOT NULL,
    "is_atc"      BOOLEAN     NOT NULL DEFAULT FALSE,
    "created_at"  TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "histories" IS '用户连线历史表';
COMMENT ON COLUMN "histories"."id" IS '连线历史id';
COMMENT ON COLUMN "histories"."user_id" IS '用户id';
COMMENT ON COLUMN "histories"."callsign" IS '呼号';
COMMENT ON COLUMN "histories"."start_time" IS '上线时间';
COMMENT ON COLUMN "histories"."end_time" IS '下线时间';
COMMENT ON COLUMN "histories"."online_time" IS '在线时间(秒)';
COMMENT ON COLUMN "histories"."is_atc" IS '是否为管制';
COMMENT ON COLUMN "histories"."created_at" IS '创建时间';

CREATE INDEX "idx_histories_user_id" ON "histories" ("user_id");

ALTER TABLE "histories"
    ADD CONSTRAINT "fk_histories_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 角色表
CREATE TABLE "roles"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "name"       VARCHAR(64) NOT NULL,
    "permission" BIGINT      NOT NULL DEFAULT 0,
    "comment"    TEXT        NOT NULL,
    "created_at" TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "roles" IS '角色表';
COMMENT ON COLUMN "roles"."id" IS '角色id';
COMMENT ON COLUMN "roles"."name" IS '角色名';
COMMENT ON COLUMN "roles"."permission" IS '角色权限';
COMMENT ON COLUMN "roles"."comment" IS '角色备注';
COMMENT ON COLUMN "roles"."created_at" IS '创建时间';
COMMENT ON COLUMN "roles"."updated_at" IS '更新时间';

-- 用户角色授予表
CREATE TABLE "user_roles"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "user_id"    BIGINT NOT NULL,
    "role_id"    BIGINT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "user_roles" IS '用户角色授予表';
COMMENT ON COLUMN "user_roles"."id" IS '主键';
COMMENT ON COLUMN "user_roles"."user_id" IS '用户id';
COMMENT ON COLUMN "user_roles"."role_id" IS '角色id';
COMMENT ON COLUMN "user_roles"."created_at" IS '创建时间';

CREATE UNIQUE INDEX "idx_user_roles_user_id_role_id" ON "user_roles" ("user_id", "role_id");
CREATE INDEX "idx_user_roles_user_id" ON "user_roles" ("user_id");
CREATE INDEX "idx_user_roles_role_id" ON "user_roles" ("role_id");

ALTER TABLE "user_roles"
    ADD CONSTRAINT "fk_user_role_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "user_roles"
    ADD CONSTRAINT "fk_user_role_role_id" FOREIGN KEY ("role_id")
        REFERENCES "roles" ("id") ON DELETE RESTRICT;

-- 公告表
CREATE TABLE "announcements"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "user_id"    BIGINT    NOT NULL,
    "title"      TEXT      NOT NULL,
    "content"    TEXT      NOT NULL,
    "type"       INTEGER   NOT NULL DEFAULT 0,
    "important"  BOOLEAN   NOT NULL DEFAULT FALSE,
    "force_show" BOOLEAN   NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL
);

COMMENT ON TABLE "announcements" IS '公告表';
COMMENT ON COLUMN "announcements"."id" IS '公告id';
COMMENT ON COLUMN "announcements"."user_id" IS '发布用户id';
COMMENT ON COLUMN "announcements"."title" IS '公告标题';
COMMENT ON COLUMN "announcements"."content" IS '公告内容';
COMMENT ON COLUMN "announcements"."type" IS '公告类型';
COMMENT ON COLUMN "announcements"."important" IS '重要公告';
COMMENT ON COLUMN "announcements"."force_show" IS '主动弹出';
COMMENT ON COLUMN "announcements"."created_at" IS '创建时间';
COMMENT ON COLUMN "announcements"."updated_at" IS '修改时间';
COMMENT ON COLUMN "announcements"."deleted_at" IS '软删除';

CREATE INDEX "idx_announcements_user_id" ON "announcements" ("user_id");

ALTER TABLE "announcements"
    ADD CONSTRAINT "fk_announcement_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 审计日志表
CREATE TABLE "audit_logs"
(
    "id"             BIGSERIAL PRIMARY KEY,
    "subject"        TEXT         NOT NULL,
    "object"         TEXT         NOT NULL,
    "event"          VARCHAR(64)  NOT NULL,
    "ip"             VARCHAR(128) NOT NULL,
    "user_agent"     TEXT         NOT NULL,
    "change_details" TEXT         NULL,
    "created_at"     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "audit_logs" IS '审计日志表';
COMMENT ON COLUMN "audit_logs"."id" IS '主键';
COMMENT ON COLUMN "audit_logs"."subject" IS '操作来源实体';
COMMENT ON COLUMN "audit_logs"."object" IS '操作对象实体';
COMMENT ON COLUMN "audit_logs"."event" IS '事件名';
COMMENT ON COLUMN "audit_logs"."ip" IS '操作来源IP';
COMMENT ON COLUMN "audit_logs"."user_agent" IS '操作来源用户代理';
COMMENT ON COLUMN "audit_logs"."change_details" IS '更改明细';
COMMENT ON COLUMN "audit_logs"."created_at" IS '记录时间';

CREATE INDEX "idx_audit_logs_event" ON "audit_logs" ("event");

-- 飞行计划表
CREATE TABLE "flight_plans"
(
    "id"                  BIGSERIAL PRIMARY KEY,
    "user_id"             BIGINT       NOT NULL,
    "callsign"            VARCHAR(16)  NOT NULL,
    "flight_rule"         VARCHAR(4)   NOT NULL,
    "aircraft"            VARCHAR(128) NOT NULL,
    "tas"                 INTEGER      NOT NULL DEFAULT 0,
    "departure_airport"   VARCHAR(8)   NOT NULL,
    "plan_departure_time" VARCHAR(8)   NOT NULL DEFAULT '0000',
    "atc_departure_time"  VARCHAR(8)   NOT NULL DEFAULT '0000',
    "cruise_altitude"     VARCHAR(8)   NOT NULL,
    "arrival_airport"     VARCHAR(8)   NOT NULL,
    "route_time_hour"     VARCHAR(2)   NOT NULL DEFAULT '00',
    "route_time_minute"   VARCHAR(2)   NOT NULL DEFAULT '00',
    "air_time_hour"       VARCHAR(2)   NOT NULL DEFAULT '00',
    "air_time_minute"     VARCHAR(2)   NOT NULL DEFAULT '00',
    "alternate_airport"   VARCHAR(8)   NOT NULL,
    "remarks"             TEXT         NOT NULL,
    "route"               TEXT         NOT NULL,
    "locked"              BOOLEAN      NOT NULL DEFAULT FALSE,
    "from_web"            BOOLEAN      NOT NULL DEFAULT FALSE,
    "created_at"          TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "updated_at"          TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "flight_plans" IS '飞行计划表';
COMMENT ON COLUMN "flight_plans"."id" IS '主键';
COMMENT ON COLUMN "flight_plans"."user_id" IS '用户id';
COMMENT ON COLUMN "flight_plans"."callsign" IS '呼号';
COMMENT ON COLUMN "flight_plans"."flight_rule" IS '飞行规则';
COMMENT ON COLUMN "flight_plans"."aircraft" IS '机型';
COMMENT ON COLUMN "flight_plans"."tas" IS '巡航真空速';
COMMENT ON COLUMN "flight_plans"."departure_airport" IS '离场机场';
COMMENT ON COLUMN "flight_plans"."plan_departure_time" IS '计划离场时间';
COMMENT ON COLUMN "flight_plans"."atc_departure_time" IS 'ATC分配离场时间';
COMMENT ON COLUMN "flight_plans"."cruise_altitude" IS '巡航高度';
COMMENT ON COLUMN "flight_plans"."arrival_airport" IS '到达机场';
COMMENT ON COLUMN "flight_plans"."route_time_hour" IS '航路耗时(小时)';
COMMENT ON COLUMN "flight_plans"."route_time_minute" IS '航路耗时(分钟)';
COMMENT ON COLUMN "flight_plans"."air_time_hour" IS '滞空时间(小时)';
COMMENT ON COLUMN "flight_plans"."air_time_minute" IS '滞空时间(分钟)';
COMMENT ON COLUMN "flight_plans"."alternate_airport" IS '备降机场';
COMMENT ON COLUMN "flight_plans"."remarks" IS '备注';
COMMENT ON COLUMN "flight_plans"."route" IS '航路';
COMMENT ON COLUMN "flight_plans"."locked" IS '计划锁定';
COMMENT ON COLUMN "flight_plans"."from_web" IS '从网页提交';
COMMENT ON COLUMN "flight_plans"."created_at" IS '创建时间';
COMMENT ON COLUMN "flight_plans"."updated_at" IS '更新时间';

CREATE UNIQUE INDEX "idx_flight_plans_user_id" ON "flight_plans" ("user_id");
CREATE INDEX "idx_flight_plans_callsign" ON "flight_plans" ("callsign");

ALTER TABLE "flight_plans"
    ADD CONSTRAINT "fk_flight_plans_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 教员信息表
CREATE TABLE "instructors"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "user_id"    BIGINT       NOT NULL,
    "email"      VARCHAR(128) NOT NULL,
    "twr"        INTEGER      NOT NULL DEFAULT 0,
    "app"        INTEGER      NOT NULL DEFAULT 0,
    "ctr"        INTEGER      NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "instructors" IS '教员信息表';
COMMENT ON COLUMN "instructors"."id" IS '主键';
COMMENT ON COLUMN "instructors"."user_id" IS '用户id';
COMMENT ON COLUMN "instructors"."email" IS '教员邮箱';
COMMENT ON COLUMN "instructors"."twr" IS '塔台教员:0-无,1-教师,2-教员';
COMMENT ON COLUMN "instructors"."app" IS '进近教员:0-无,1-教师,2-教员';
COMMENT ON COLUMN "instructors"."ctr" IS '区域教员:0-无,1-教师,2-教员';
COMMENT ON COLUMN "instructors"."created_at" IS '加入时间';
COMMENT ON COLUMN "instructors"."updated_at" IS '信息更新时间';

CREATE UNIQUE INDEX "idx_instructors_user_id" ON "instructors" ("user_id");

ALTER TABLE "instructors"
    ADD CONSTRAINT "fk_instructor_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;


-- 管制员申请表
CREATE TABLE "controller_applications"
(
    "id"            BIGSERIAL PRIMARY KEY,
    "user_id"       BIGINT      NOT NULL,
    "reason"        TEXT        NOT NULL,
    "record"        TEXT        NOT NULL,
    "is_guest"      BOOLEAN     NOT NULL DEFAULT FALSE,
    "platform"      VARCHAR(16) NOT NULL,
    "image_id"      BIGINT      NULL     DEFAULT NULL,
    "status"        INTEGER     NOT NULL DEFAULT 0,
    "message"       TEXT        NULL     DEFAULT NULL,
    "instructor_id" BIGINT      NULL,
    "created_at"    TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "controller_applications" IS '管制员申请表';
COMMENT ON COLUMN "controller_applications"."id" IS '主键';
COMMENT ON COLUMN "controller_applications"."user_id" IS '用户id';
COMMENT ON COLUMN "controller_applications"."reason" IS '申请理由';
COMMENT ON COLUMN "controller_applications"."record" IS '管制经历';
COMMENT ON COLUMN "controller_applications"."is_guest" IS '是否为客座';
COMMENT ON COLUMN "controller_applications"."platform" IS '客座平台';
COMMENT ON COLUMN "controller_applications"."image_id" IS '客座证明资料';
COMMENT ON COLUMN "controller_applications"."status" IS '申请状态';
COMMENT ON COLUMN "controller_applications"."message" IS '回复消息';
COMMENT ON COLUMN "controller_applications"."created_at" IS '创建时间';
COMMENT ON COLUMN "controller_applications"."updated_at" IS '更新时间';

CREATE INDEX "idx_controller_applications_user_id" ON "controller_applications" ("user_id");
CREATE INDEX "idx_controller_applications_image_id" ON "controller_applications" ("image_id");
CREATE INDEX "idx_controller_applications_instructor_id" ON "controller_applications" ("instructor_id");

ALTER TABLE "controller_applications"
    ADD CONSTRAINT "fk_controller_application_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "controller_applications"
    ADD CONSTRAINT "fk_controller_application_image_id" FOREIGN KEY ("image_id")
        REFERENCES "images" ("id") ON DELETE RESTRICT;
ALTER TABLE "controller_applications"
    ADD CONSTRAINT "fk_controller_application_instructor_id" FOREIGN KEY ("instructor_id")
        REFERENCES "instructors" ("id") ON DELETE RESTRICT;

CREATE TABLE "controller_application_times"
(
    "id"             BIGSERIAL PRIMARY KEY,
    "application_id" BIGINT    NOT NULL,
    "time"           TIMESTAMP NOT NULL,
    "selected"       BOOLEAN   NOT NULL DEFAULT FALSE
);

COMMENT ON TABLE "controller_application_times" IS '管制员申请时间表';
COMMENT ON COLUMN "controller_application_times"."id" IS '主键';
COMMENT ON COLUMN "controller_application_times"."application_id" IS '申请id';
COMMENT ON COLUMN "controller_application_times"."time" IS '申请时间';
COMMENT ON COLUMN "controller_application_times"."selected" IS '是否被选中';

CREATE INDEX "idx_controller_application_times_application_id" ON "controller_application_times" ("application_id");

ALTER TABLE "controller_application_times"
    ADD CONSTRAINT "fk_controller_application_times_application_id" FOREIGN KEY ("application_id")
        REFERENCES "controller_applications" ("id") ON DELETE RESTRICT;

-- 管制员表
CREATE TABLE "controllers"
(
    "id"                    BIGSERIAL PRIMARY KEY,
    "user_id"               BIGINT    NOT NULL,
    "instructor_id"         BIGINT    NULL,
    "guest"                 BOOLEAN   NOT NULL DEFAULT FALSE,
    "under_monitor"         BOOLEAN   NOT NULL DEFAULT FALSE,
    "under_solo"            BOOLEAN   NOT NULL DEFAULT FALSE,
    "solo_until"            TIMESTAMP NULL,
    "tier2"                 BOOLEAN   NOT NULL DEFAULT FALSE,
    "total_controller_time" BIGINT    NOT NULL DEFAULT 0,
    "created_at"            TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "updated_at"            TIMESTAMP          DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "controllers" IS '管制员表';
COMMENT ON COLUMN "controllers"."id" IS '主键';
COMMENT ON COLUMN "controllers"."user_id" IS '用户id';
COMMENT ON COLUMN "controllers"."instructor_id" IS '教员id';
COMMENT ON COLUMN "controllers"."guest" IS '是否为客座管制员';
COMMENT ON COLUMN "controllers"."under_monitor" IS '是否为实习管制员';
COMMENT ON COLUMN "controllers"."under_solo" IS '是否SOLO';
COMMENT ON COLUMN "controllers"."solo_until" IS 'SOLO时限';
COMMENT ON COLUMN "controllers"."tier2" IS '是否有程序塔权限';
COMMENT ON COLUMN "controllers"."total_controller_time" IS '总管制时长';
COMMENT ON COLUMN "controllers"."created_at" IS '加入时间';
COMMENT ON COLUMN "controllers"."updated_at" IS '信息更新时间';

CREATE UNIQUE INDEX "idx_controllers_user_id" ON "controllers" ("user_id");
CREATE INDEX "idx_controllers_instructor_id" ON "controllers" ("instructor_id");

ALTER TABLE "controllers"
    ADD CONSTRAINT "fk_controller_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "controllers"
    ADD CONSTRAINT "fk_controller_instructor_id" FOREIGN KEY ("instructor_id")
        REFERENCES "instructors" ("id") ON DELETE RESTRICT;

-- 管制员履历表
CREATE TABLE "controller_records"
(
    "id"            BIGSERIAL PRIMARY KEY,
    "user_id"       BIGINT    NOT NULL,
    "instructor_id" BIGINT    NULL,
    "type"          INTEGER   NOT NULL DEFAULT 0,
    "content"       TEXT      NOT NULL,
    "created_at"    TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"    TIMESTAMP NULL
);

COMMENT ON TABLE "controller_records" IS '管制员履历表';
COMMENT ON COLUMN "controller_records"."id" IS '主键';
COMMENT ON COLUMN "controller_records"."user_id" IS '用户id';
COMMENT ON COLUMN "controller_records"."instructor_id" IS '教员id';
COMMENT ON COLUMN "controller_records"."type" IS '履历类型';
COMMENT ON COLUMN "controller_records"."content" IS '履历内容';
COMMENT ON COLUMN "controller_records"."created_at" IS '记录时间';
COMMENT ON COLUMN "controller_records"."deleted_at" IS '软删除';

CREATE INDEX "idx_controller_records_user_id" ON "controller_records" ("user_id");
CREATE INDEX "idx_controller_records_instructor_id" ON "controller_records" ("instructor_id");

ALTER TABLE "controller_records"
    ADD CONSTRAINT "fk_controller_record_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "controller_records"
    ADD CONSTRAINT "fk_controller_record_instructor_id" FOREIGN KEY ("instructor_id")
        REFERENCES "instructors" ("id") ON DELETE RESTRICT;

-- 工单表
CREATE TABLE "tickets"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "user_id"    BIGINT    NOT NULL,
    "type"       INTEGER   NOT NULL DEFAULT 0,
    "title"      TEXT      NOT NULL,
    "content"    TEXT      NOT NULL,
    "reply"      TEXT      NULL,
    "replier"    BIGINT    NULL,
    "created_at" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "closed_at"  TIMESTAMP NULL
);

COMMENT ON TABLE "tickets" IS '工单表';
COMMENT ON COLUMN "tickets"."id" IS '主键';
COMMENT ON COLUMN "tickets"."user_id" IS '用户id';
COMMENT ON COLUMN "tickets"."type" IS '工单类型';
COMMENT ON COLUMN "tickets"."title" IS '工单标题';
COMMENT ON COLUMN "tickets"."content" IS '工单内容';
COMMENT ON COLUMN "tickets"."reply" IS '工单回复';
COMMENT ON COLUMN "tickets"."replier" IS '工单回复人id';
COMMENT ON COLUMN "tickets"."created_at" IS '创建时间';
COMMENT ON COLUMN "tickets"."closed_at" IS '工单关闭时间';

CREATE INDEX "idx_tickets_user_id" ON "tickets" ("user_id");
CREATE INDEX "idx_tickets_replier_id" ON "tickets" ("replier");

ALTER TABLE "tickets"
    ADD CONSTRAINT "fk_ticket_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "tickets"
    ADD CONSTRAINT "fk_ticket_replier_id" FOREIGN KEY ("replier")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 活动表
CREATE TABLE "activities"
(
    "id"                BIGSERIAL PRIMARY KEY,
    "publisher_id"      BIGINT       NOT NULL,
    "type"              INTEGER      NOT NULL DEFAULT 0,
    "title"             TEXT         NOT NULL,
    "image_id"          BIGINT       NULL,
    "active_time"       TIMESTAMP    NOT NULL,
    "departure_airport" VARCHAR(64)  NOT NULL,
    "arrival_airport"   VARCHAR(64)  NOT NULL,
    "route"             TEXT         NULL,
    "distance"          INTEGER      NULL,
    "route2"            TEXT         NULL,
    "distance2"         INTEGER      NULL,
    "open_fir"          VARCHAR(128) NULL,
    "status"            INTEGER      NOT NULL DEFAULT 0,
    "notams"            TEXT         NOT NULL,
    "created_at"        TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "updated_at"        TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"        TIMESTAMP    NULL
);

COMMENT ON TABLE "activities" IS '活动表';
COMMENT ON COLUMN "activities"."id" IS '主键';
COMMENT ON COLUMN "activities"."publisher_id" IS '发布者id';
COMMENT ON COLUMN "activities"."type" IS '活动类型';
COMMENT ON COLUMN "activities"."title" IS '活动标题';
COMMENT ON COLUMN "activities"."image_id" IS '活动图片id';
COMMENT ON COLUMN "activities"."active_time" IS '活动时间';
COMMENT ON COLUMN "activities"."departure_airport" IS '离场机场';
COMMENT ON COLUMN "activities"."arrival_airport" IS '到达机场';
COMMENT ON COLUMN "activities"."route" IS '活动路线';
COMMENT ON COLUMN "activities"."distance" IS '活动距离';
COMMENT ON COLUMN "activities"."route2" IS '活动路线2';
COMMENT ON COLUMN "activities"."distance2" IS '活动距离2';
COMMENT ON COLUMN "activities"."open_fir" IS '空域开放日';
COMMENT ON COLUMN "activities"."status" IS '活动状态';
COMMENT ON COLUMN "activities"."notams" IS '航行通告';
COMMENT ON COLUMN "activities"."created_at" IS '创建时间';
COMMENT ON COLUMN "activities"."updated_at" IS '更新时间';
COMMENT ON COLUMN "activities"."deleted_at" IS '软删除';

CREATE INDEX "idx_activities_publisher_id" ON "activities" ("publisher_id");
CREATE INDEX "idx_activities_image_id" ON "activities" ("image_id");

ALTER TABLE "activities"
    ADD CONSTRAINT "fk_activity_publisher_id" FOREIGN KEY ("publisher_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "activities"
    ADD CONSTRAINT "fk_activity_image_id" FOREIGN KEY ("image_id")
        REFERENCES "images" ("id") ON DELETE RESTRICT;

-- 活动飞行员表
CREATE TABLE "activity_pilots"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "activity_id" BIGINT      NOT NULL,
    "user_id"     BIGINT      NOT NULL,
    "callsign"    VARCHAR(64) NOT NULL,
    "aircraft"    VARCHAR(64) NOT NULL,
    "status"      INTEGER     NOT NULL DEFAULT 0,
    "created_at"  TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "activity_pilots" IS '活动飞行员表';
COMMENT ON COLUMN "activity_pilots"."id" IS '主键';
COMMENT ON COLUMN "activity_pilots"."activity_id" IS '活动id';
COMMENT ON COLUMN "activity_pilots"."user_id" IS '飞行员id';
COMMENT ON COLUMN "activity_pilots"."callsign" IS '呼号';
COMMENT ON COLUMN "activity_pilots"."aircraft" IS '机型';
COMMENT ON COLUMN "activity_pilots"."status" IS '飞行员状态';
COMMENT ON COLUMN "activity_pilots"."created_at" IS '报名时间';
COMMENT ON COLUMN "activity_pilots"."updated_at" IS '更新时间';

CREATE INDEX "idx_activity_pilots_activity_id" ON "activity_pilots" ("activity_id");
CREATE INDEX "idx_activity_pilots_user_id" ON "activity_pilots" ("user_id");

ALTER TABLE "activity_pilots"
    ADD CONSTRAINT "fk_activity_pilot_activity_id" FOREIGN KEY ("activity_id")
        REFERENCES "activities" ("id") ON DELETE RESTRICT;
ALTER TABLE "activity_pilots"
    ADD CONSTRAINT "fk_activity_pilot_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 活动席位表
CREATE TABLE "activity_facilities"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "activity_id" BIGINT      NOT NULL,
    "min_rating"  INTEGER     NOT NULL DEFAULT 0,
    "callsign"    VARCHAR(32) NOT NULL,
    "frequency"   VARCHAR(32) NOT NULL,
    "tier2"       BOOLEAN     NOT NULL DEFAULT FALSE,
    "sort_index"  INTEGER     NOT NULL DEFAULT 0,
    "created_at"  TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "activity_facilities" IS '活动席位表';
COMMENT ON COLUMN "activity_facilities"."id" IS '主键';
COMMENT ON COLUMN "activity_facilities"."activity_id" IS '活动id';
COMMENT ON COLUMN "activity_facilities"."min_rating" IS '最低等级';
COMMENT ON COLUMN "activity_facilities"."callsign" IS '呼号';
COMMENT ON COLUMN "activity_facilities"."frequency" IS '频率';
COMMENT ON COLUMN "activity_facilities"."tier2" IS '是否为程序塔台';
COMMENT ON COLUMN "activity_facilities"."sort_index" IS '排序索引';
COMMENT ON COLUMN "activity_facilities"."created_at" IS '创建时间';
COMMENT ON COLUMN "activity_facilities"."updated_at" IS '更新时间';

CREATE INDEX "idx_activity_facilities_activity_id" ON "activity_facilities" ("activity_id");

ALTER TABLE "activity_facilities"
    ADD CONSTRAINT "fk_activity_facility_activity_id" FOREIGN KEY ("activity_id")
        REFERENCES "activities" ("id") ON DELETE RESTRICT;

-- 活动管制员表
CREATE TABLE "activity_controllers"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "user_id"     BIGINT NOT NULL,
    "activity_id" BIGINT NOT NULL,
    "facility_id" BIGINT NOT NULL,
    "created_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "activity_controllers" IS '活动管制员表';
COMMENT ON COLUMN "activity_controllers"."id" IS '主键';
COMMENT ON COLUMN "activity_controllers"."user_id" IS '用户id';
COMMENT ON COLUMN "activity_controllers"."activity_id" IS '活动id';
COMMENT ON COLUMN "activity_controllers"."facility_id" IS '活动席位id';
COMMENT ON COLUMN "activity_controllers"."created_at" IS '创建时间';

CREATE INDEX "idx_activity_controllers_user_id" ON "activity_controllers" ("user_id");
CREATE INDEX "idx_activity_controllers_activity_id" ON "activity_controllers" ("activity_id");
CREATE INDEX "idx_activity_controllers_facility_id" ON "activity_controllers" ("facility_id");

ALTER TABLE "activity_controllers"
    ADD CONSTRAINT "fk_activity_controller_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;
ALTER TABLE "activity_controllers"
    ADD CONSTRAINT "fk_activity_controller_activity_id" FOREIGN KEY ("activity_id")
        REFERENCES "activities" ("id") ON DELETE RESTRICT;
ALTER TABLE "activity_controllers"
    ADD CONSTRAINT "fk_activity_controller_facility_id" FOREIGN KEY ("facility_id")
        REFERENCES "activity_facilities" ("id") ON DELETE RESTRICT;

-- 活动记录表
CREATE TABLE "activity_records"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "activity_id" BIGINT NOT NULL,
    "user_id"     BIGINT NOT NULL,
    "created_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "activity_records" IS '活动记录表';
COMMENT ON COLUMN "activity_records"."id" IS '主键';
COMMENT ON COLUMN "activity_records"."activity_id" IS '活动id';
COMMENT ON COLUMN "activity_records"."user_id" IS '用户id';
COMMENT ON COLUMN "activity_records"."created_at" IS '创建时间';

CREATE INDEX "idx_activity_records_activity_id" ON "activity_records" ("activity_id");
CREATE INDEX "idx_activity_records_user_id" ON "activity_records" ("user_id");

ALTER TABLE "activity_records"
    ADD CONSTRAINT "fk_activity_record_activity_id" FOREIGN KEY ("activity_id")
        REFERENCES "activities" ("id") ON DELETE RESTRICT;
ALTER TABLE "activity_records"
    ADD CONSTRAINT "fk_activity_record_user_id" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id") ON DELETE RESTRICT;

-- 活动协调表
CREATE TABLE "activity_coordinations"
(
    "id"                BIGSERIAL PRIMARY KEY,
    "activity_id"       BIGINT      NOT NULL,
    "facility_id"       BIGINT      NOT NULL,
    "logon_code"        VARCHAR(8)  NULL     DEFAULT NULL,
    "logon_network"     VARCHAR(64) NULL     DEFAULT NULL,
    "pdc_available"     BOOLEAN     NOT NULL DEFAULT FALSE,
    "transform"         VARCHAR(64) NULL     DEFAULT NULL,
    "transform_remarks" TEXT        NULL     DEFAULT NULL,
    "procedure"         VARCHAR(64) NULL     DEFAULT NULL,
    "runway"            VARCHAR(64) NULL     DEFAULT NULL,
    "runway_remarks"    TEXT        NULL     DEFAULT NULL,
    "altitude"          VARCHAR(64) NULL     DEFAULT NULL,
    "remarks"           TEXT        NULL     DEFAULT NULL,
    "created_at"        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "activity_coordinations" IS '活动协调表';
COMMENT ON COLUMN "activity_coordinations"."id" IS '主键';
COMMENT ON COLUMN "activity_coordinations"."activity_id" IS '活动id';
COMMENT ON COLUMN "activity_coordinations"."facility_id" IS '活动席位id';
COMMENT ON COLUMN "activity_coordinations"."logon_code" IS 'CPDLC/DCL识别码';
COMMENT ON COLUMN "activity_coordinations"."logon_network" IS 'CPDLC/DCL登陆网络';
COMMENT ON COLUMN "activity_coordinations"."pdc_available" IS 'PDC是否可用';
COMMENT ON COLUMN "activity_coordinations"."transform" IS '移交席位';
COMMENT ON COLUMN "activity_coordinations"."transform_remarks" IS '移交备注';
COMMENT ON COLUMN "activity_coordinations"."procedure" IS '程序';
COMMENT ON COLUMN "activity_coordinations"."runway" IS '跑道';
COMMENT ON COLUMN "activity_coordinations"."runway_remarks" IS '跑道运行备注';
COMMENT ON COLUMN "activity_coordinations"."altitude" IS '起始高度';
COMMENT ON COLUMN "activity_coordinations"."remarks" IS '其他备注信息';
COMMENT ON COLUMN "activity_coordinations"."created_at" IS '创建时间';
COMMENT ON COLUMN "activity_coordinations"."updated_at" IS '更新时间';

CREATE INDEX "idx_activity_coordinations_activity_id" ON "activity_coordinations" ("activity_id");
CREATE INDEX "idx_activity_coordinations_facility_id" ON "activity_coordinations" ("facility_id");

ALTER TABLE "activity_coordinations"
    ADD CONSTRAINT "fk_activity_coordinator_activity_id" FOREIGN KEY ("activity_id")
        REFERENCES "activities" ("id") ON DELETE RESTRICT;

ALTER TABLE "activity_coordinations"
    ADD CONSTRAINT "fk_activity_coordinator_facility_id" FOREIGN KEY ("facility_id")
        REFERENCES "activity_facilities" ("id") ON DELETE RESTRICT;
create table if not exists editor
(
    `id`          int auto_increment comment 'ID link'  primary key,
    `dt_last_use` datetime default CURRENT_TIMESTAMP    not null comment 'DATETIME create link',
    `table`       varchar(64)                           not null comment 'Table name',
    `row_name`    varchar(64)                           not null comment 'Row name',
    `id_line`     int                                   not null comment 'ID Line',
    `type`        enum ('single-line-text', 'checkbox', 'input-text') not null comment 'Edit type',
    `key`         varchar(32)                           not null comment 'Key of the link'
) comment 'editor links' charset = utf8mb4;

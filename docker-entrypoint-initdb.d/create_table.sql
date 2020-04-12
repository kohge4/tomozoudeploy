CREATE DATABASE tomozou;

use tomozou;

CREATE TABLE `user` (
    `id` int(11),
    `social_id` varchar(255),
    `name` varchar(255),
    `auth` varchar(255),
    `image` varchar(255),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

CREATE TABLE `artist` (
    `id` int(11),
    `social_id` varchar(255),
    `name` varchar(255),
    `image` varchar(255),
    `created_at` datetime(6),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

CREATE TABLE `user_artist_tag` (
    `id` int(11),
    `user_id` varchar(255),
    `artist_id` varchar(255),
    `tag_name` varchar(255),
    `created_at` datetime(6),
    `artist_comment` varchar(255),
    `image` varchar(255),
    `artist_name` varchar(255),
    `url` varchar(255),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

CREATE TABLE `track` (
    `id` int(11),
    `social_id` varchar(255),
    `name` varchar(255),
    `artist_id` int(11),
    `artist_name` varchar(255),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;


CREATE TABLE `user_track_tag` (
    `id` int(11),
    `user_id` int(11),
    `artist_id` int(11),
    `track_id` int(11),
    `tag_name` varchar(255),
    `created_at` datetime(6),
    `track_name` varchar(255),
    `track_comment` varchar(255),
    `track_social_id` varchar(255),
    `artist_name` varchar(255),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;


CREATE TABLE `token` (
    `id` int(11),
    `user_id` varchar(255),
    `auth_type` varchar(255),
    `access_token` varchar(255),
    `token_type` varchar(255),
    `refresh_token` varchar(255),
    `expiry` datetime(6),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;


CREATE TABLE `user_chat` (
    `id` int(11),
    `user_id` int(11),
    `artist_id` int(11),
    `comment` varchar(255),
    `created_at` datetime(6),
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
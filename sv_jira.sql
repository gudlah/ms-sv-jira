
CREATE TABLE `activity_log` (
  `id` varchar(100) NOT NULL,
  `id_request` varchar(100) DEFAULT NULL,
  `id_user` varchar(100) DEFAULT NULL,
  `client_ip` varchar(100) DEFAULT NULL,
  `endpoint` varchar(100) DEFAULT NULL,
  `response_code` varchar(100) DEFAULT NULL,
  `body_request` text DEFAULT NULL,
  `body_response` longtext DEFAULT NULL,
  `created_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `ip_blockeds` (
  `id` varchar(255) NOT NULL,
  `client_ip` varchar(255) DEFAULT NULL,
  `created_at` varchar(100) DEFAULT NULL,
  `deleted_at` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_attachments` (
  `attachment_id` varchar(10) NOT NULL,
  `card_id` varchar(10) DEFAULT NULL,
  `attachment_file_name` varchar(100) DEFAULT NULL,
  `attachment_url` text DEFAULT NULL,
  `author_id` varchar(100) DEFAULT NULL,
  `attachment_created` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_boards` (
  `board_id` int(100) NOT NULL,
  `project_id` varchar(100) DEFAULT NULL,
  `board_name` varchar(100) DEFAULT NULL,
  `board_type` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_cards` (
  `card_id` varchar(10) NOT NULL,
  `column_id` varchar(10) DEFAULT NULL,
  `card_title` varchar(255) DEFAULT NULL,
  `card_key` varchar(10) DEFAULT NULL,
  `card_description` longtext DEFAULT NULL,
  `card_type_id` varchar(10) DEFAULT NULL,
  `priority_id` int(11) DEFAULT NULL,
  `assignee_id` varchar(100) DEFAULT NULL,
  `creator_id` varchar(100) DEFAULT NULL,
  `reporter_id` varchar(100) DEFAULT NULL,
  `card_created` datetime DEFAULT NULL,
  `card_updated` datetime DEFAULT NULL,
  `card_resolved` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_columns` (
  `column_id` varchar(10) NOT NULL,
  `sprint_id` varchar(10) DEFAULT NULL,
  `column_name` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_comments` (
  `comment_id` varchar(10) NOT NULL,
  `card_id` varchar(10) DEFAULT NULL,
  `author_id` varchar(100) DEFAULT NULL,
  `comment_body` longtext DEFAULT NULL,
  `comment_created` datetime DEFAULT NULL,
  `comment_updated` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_priorities` (
  `priority_id` int(11) NOT NULL,
  `priority_name` varchar(100) DEFAULT NULL,
  `priority_description` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_projects` (
  `project_id` varchar(50) NOT NULL,
  `project_key` varchar(50) DEFAULT NULL,
  `project_name` varchar(100) DEFAULT NULL,
  `project_type_key` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_sprints` (
  `sprint_id` int(11) NOT NULL,
  `board_id` int(11) DEFAULT NULL,
  `sprint_state` varchar(100) DEFAULT NULL,
  `sprint_name` varchar(100) DEFAULT NULL,
  `sprint_start_date` datetime DEFAULT NULL,
  `sprint_end_date` datetime DEFAULT NULL,
  `sprint_created_date` datetime DEFAULT NULL,
  `sprint_goal` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_sub_tasks` (
  `sub_task_id` varchar(10) NOT NULL,
  `card_key` varchar(50) DEFAULT NULL,
  `sub_task_key` varchar(10) DEFAULT NULL,
  `sub_task_title` varchar(100) DEFAULT NULL,
  `sub_task_description` longtext DEFAULT NULL,
  `status_id` varchar(10) DEFAULT NULL,
  `priority_id` int(11) DEFAULT NULL,
  `creator_id` varchar(100) DEFAULT NULL,
  `reporter_id` varchar(100) DEFAULT NULL,
  `assignee_id` varchar(100) DEFAULT NULL,
  `sub_task_created` datetime DEFAULT NULL,
  `sub_task_updated` datetime DEFAULT NULL,
  `sub_task_resolved` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jira_users` (
  `user_id` varchar(100) NOT NULL,
  `user_email` varchar(100) DEFAULT NULL,
  `user_display_name` varchar(100) DEFAULT NULL,
  `user_active` tinyint(1) DEFAULT NULL,
  `user_locale` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `service_users` (
  `id` varchar(255) NOT NULL,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `is_active` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `symbols` (
  `id` int(11) NOT NULL,
  `symbol` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `upstream_service_request_log` (
  `id` varchar(100) NOT NULL,
  `id_request` varchar(100) DEFAULT NULL,
  `request_payload` longtext DEFAULT NULL,
  `response_payload` longtext DEFAULT NULL,
  `is_success` int(11) DEFAULT NULL,
  `response_timestamp` datetime DEFAULT NULL,
  `request_timestamp` datetime DEFAULT NULL,
  `url` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

ALTER TABLE `activity_log` ADD PRIMARY KEY (`id`);

ALTER TABLE `ip_blockeds` ADD PRIMARY KEY (`id`);

ALTER TABLE `jira_attachments` ADD PRIMARY KEY (`attachment_id`);

ALTER TABLE `jira_boards` ADD PRIMARY KEY (`board_id`);

ALTER TABLE `jira_cards` ADD PRIMARY KEY (`card_id`);

ALTER TABLE `jira_columns` ADD PRIMARY KEY (`column_id`);

ALTER TABLE `jira_comments` ADD PRIMARY KEY (`comment_id`);

ALTER TABLE `jira_priorities` ADD PRIMARY KEY (`priority_id`);

ALTER TABLE `jira_projects` ADD PRIMARY KEY (`project_id`);

ALTER TABLE `jira_sprints` ADD PRIMARY KEY (`sprint_id`);

ALTER TABLE `jira_sub_tasks` ADD PRIMARY KEY (`sub_task_id`);

ALTER TABLE `jira_users` ADD PRIMARY KEY (`user_id`);

ALTER TABLE `service_users` ADD PRIMARY KEY (`id`);

ALTER TABLE `symbols` ADD PRIMARY KEY (`id`);

ALTER TABLE `upstream_service_request_log` ADD PRIMARY KEY (`id`);

INSERT INTO `service_users` (`id`, `username`, `password`, `is_active`, `created_at`) VALUES
('b0577d97-3a3d-44c9-a82f-0ff348ca951d', 'svi', '$2a$12$HUUsrBqls5p8Sh3TxmKV0uPQjTsXa2p.O9Uob2G6iJFa.g/z5mAJu', 1, '2024-09-09 12:32:01');

INSERT INTO `symbols` (`id`, `symbol`) VALUES
(1, '`'),
(2, '--'),
(6, '%'),
(7, ';'),
(8, '+'),
(9, '||'),
(10, '='),
(11, '>'),
(12, '<'),
(13, '<='),
(14, '>='),
(15, '=='),
(16, '<>'),
(17, '!='),
(19, '^^'),
(20, '&&'),
(21, '{'),
(22, '}'),
(23, '('),
(24, ')'),
(27, "'"),
(28, '"'),
(29, '?'),
(30, '!'),
(31, '['),
(32, ']'),
(33, '*');
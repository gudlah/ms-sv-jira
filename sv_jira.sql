-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 10, 2024 at 09:05 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sv_jira`
--

-- --------------------------------------------------------

--
-- Table structure for table `activity_log`
--

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

-- --------------------------------------------------------

--
-- Table structure for table `ip_blockeds`
--

CREATE TABLE `ip_blockeds` (
  `id` varchar(255) NOT NULL,
  `client_ip` varchar(255) DEFAULT NULL,
  `created_at` varchar(100) DEFAULT NULL,
  `deleted_at` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `symbols`
--

CREATE TABLE `symbols` (
  `id` int(11) NOT NULL,
  `symbol` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `symbols`
--

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

-- --------------------------------------------------------

--
-- Table structure for table `upstream_service_request_log`
--

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

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `is_active` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `is_active`, `created_at`) VALUES
('b0577d97-3a3d-44c9-a82f-0ff348ca951d', 'svi', '$2a$12$HUUsrBqls5p8Sh3TxmKV0uPQjTsXa2p.O9Uob2G6iJFa.g/z5mAJu', 1, '2024-09-09 19:32:01');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `activity_log`
--
ALTER TABLE `activity_log`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ip_blockeds`
--
ALTER TABLE `ip_blockeds`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `symbols`
--
ALTER TABLE `symbols`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `upstream_service_request_log`
--
ALTER TABLE `upstream_service_request_log`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

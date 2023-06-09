--resource表的创建SQL--

CREATE TABLE `resource` (
  `id` char(64) CHARACTER SET latin1 NOT NULL,
  `vendor` tinyint(1) NOT NULL,
  `region` varchar(64) CHARACTER SET latin1 NOT NULL,
  `zone` varchar(64) CHARACTER SET latin1 NOT NULL,
  `create_at` bigint(13) NOT NULL,
  `expire_at` bigint(13) DEFAULT NULL,
  `category` varchar(64) CHARACTER SET latin1 NOT NULL,
  `type` varchar(120) CHARACTER SET latin1 NOT NULL,
  `instance_id` varchar(120) CHARACTER SET latin1 NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) CHARACTER SET latin1 NOT NULL,
  `update_at` bigint(13) DEFAULT NULL,
  `sync_at` bigint(13) DEFAULT NULL,
  `sync_accout` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
  `public_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
  `private_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
  `pay_type` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
  `describe_hash` varchar(255) NOT NULL,
  `resource_hash` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`) USING BTREE,
  KEY `status` (`status`) USING HASH,
  KEY `private_ip` (`public_ip`) USING BTREE,
  KEY `public_ip` (`public_ip`) USING BTREE,
  KEY `instance_id` (`instance_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4


--host表SQL如下:--

CREATE TABLE `host` (
  `resource_id` varchar(64) NOT NULL,
  `cpu` tinyint(4) NOT NULL,
  `memory` int(13) NOT NULL,
  `gpu_amount` tinyint(4) DEFAULT NULL,
  `gpu_spec` varchar(255) DEFAULT NULL,
  `os_type` varchar(255) DEFAULT NULL,
  `os_name` varchar(255) DEFAULT NULL,
  `serial_number` varchar(120) DEFAULT NULL,
  `image_id` char(64) DEFAULT NULL,
  `internet_max_bandwidth_out` int(10) DEFAULT NULL,
  `internet_max_bandwidth_in` int(10) DEFAULT NULL,
  `key_pair_name` varchar(255) DEFAULT NULL,
  `security_groups` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



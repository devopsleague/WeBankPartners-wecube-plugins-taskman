SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `attach_file`;
CREATE TABLE IF NOT EXISTS `attach_file` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `s3_url` varchar(255)  NOT NULL COMMENT 's3url',
  `s3_bucket_name` varchar(255)  DEFAULT NULL COMMENT 's3_bucket',
  `s3_key_name` varchar(255)  DEFAULT NULL COMMENT 's3_key',
  `created_by` varchar(255)  NULL,
  `created_time` datetime NULL,
  `updated_by` varchar(255)  NULL,
  `updated_time` datetime NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
  `id` varchar(64)  NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `updated_time` datetime,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `form_template`;
CREATE TABLE IF NOT EXISTS `form_template` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `description` varchar(512)  DEFAULT NULL,
  `version` varchar(64) DEFAULT NULL,
  `confirm_time` datetime,
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` datetime,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` datetime,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `form_item_template`;
CREATE TABLE IF NOT EXISTS `form_item_template` (
  `id` varchar(64) NOT NULL,
  `record_id` varchar(64) DEFAULT NULL,
  `version` varchar(64) DEFAULT NULL,
  `form_template` varchar(64) DEFAULT NULL,
  `name` varchar(64) NOT NULL,
  `description` varchar(255)  DEFAULT NULL,
  `default_value` varchar(255) DEFAULT NULL,
  `sort` int NOT NULL DEFAULT '0',
  `package_name` varchar(255) DEFAULT NULL,
  `entity` varchar(255) DEFAULT NULL,
  `attr_def_id` varchar(255) DEFAULT NULL,
  `attr_def_name` varchar(255) DEFAULT NULL,
  `attr_def_data_type` varchar(255) DEFAULT NULL,
  `element_type` varchar(64) NOT NULL DEFAULT 'text',
  `title` varchar(64) DEFAULT NULL,
  `width` int(11) DEFAULT 80,
  `ref_package_name` varchar(255) DEFAULT NULL,
  `ref_entity` varchar(255) DEFAULT NULL,
  `data_options` text,
  `required` varchar(16) NOT NULL DEFAULT 'no',
  `regular` varchar(255) DEFAULT NULL,
  `is_edit` varchar(16) NOT NULL DEFAULT 'yes',
  `is_view` varchar(16) NOT NULL DEFAULT 'yes',
  `is_output` varchar(16) NOT NULL DEFAULT 'no',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_form_item_template` FOREIGN KEY (`form_template`) REFERENCES `form_template` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';


INSERT INTO `form_item_template`
(`id`, `form_template`, `name`, `default_value`, `sort`, `package_name`, `entity`, `attr_def_id`, `attr_def_data_type`, `element_type`, `title`, `width`, `ref_package_name`, `ref_entity`, `data_options`)
 VALUES
    ('1333228970392473601', NULL, 'reporter', NULL, 0, NULL, NULL, NULL, NULL, 'text', '上报人', 80, NULL, NULL, NULL),
    ('1333303089673617409', NULL, 'emergency', NULL, 0, NULL, NULL, NULL, NULL, 'select', '紧急程度', 80, NULL, NULL, NULL),
    ('1333304415006547970', NULL, 'attach_file_id', NULL, 0, NULL, NULL, NULL, NULL, 'file', '附件', 80, NULL, NULL, NULL),
    ('1333319171714420738', NULL, 'result', NULL, 0, NULL, NULL, NULL, NULL, 'text', '处理结果', 80, NULL, NULL, NULL),
    ('1333324857626169346', NULL, 'report_time', NULL, 0, NULL, NULL, NULL, NULL, 'date', '上报时间', 80, NULL, NULL, NULL),
    ('1333324857626169347', NULL, 'due_data', NULL, 0, NULL, NULL, NULL, NULL, 'date', '过期时间', 80, NULL, NULL, NULL);

DROP TABLE IF EXISTS `form`;
CREATE TABLE IF NOT EXISTS `form` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `description` varchar(255)  DEFAULT NULL,
  `form_template` varchar(64)  DEFAULT NULL,
  `form_version` varchar(64)  DEFAULT NULL,
  `created_time` datetime NULL,
  `created_by` varchar(255)   NULL,
  `updated_by` varchar(255)   NULL,
  `updated_time` datetime NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_form_form_template` FOREIGN KEY (`form_template`) REFERENCES `form_template` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `form_item`;
CREATE TABLE IF NOT EXISTS `form_item` (
  `id` varchar(64)  NOT NULL,
  `form` varchar(64)  NOT NULL,
  `form_item_template` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `value` varchar(255)  DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_form_item_form` FOREIGN KEY (`form`) REFERENCES `form` (`id`),
  CONSTRAINT `fore_form_item_ref_template` FOREIGN KEY (`form_item_template`) REFERENCES `form_item_template` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `request_template_group`;
CREATE TABLE IF NOT EXISTS `request_template_group` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `description` varchar(255)  DEFAULT NULL,
  `manage_role` varchar(64)  DEFAULT NULL,
  `created_by` varchar(255)   NULL,
  `created_time` datetime NULL,
  `updated_by` varchar(255)   NULL,
  `updated_time` datetime NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_template_group_role` FOREIGN KEY (`manage_role`) REFERENCES `role` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `request_template`;
CREATE TABLE IF NOT EXISTS `request_template` (
  `id` varchar(64)  NOT NULL,
  `group` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `description` varchar(255)  DEFAULT NULL,
  `form_template` varchar(64) DEFAULT NULL,
  `tags` varchar(255)  DEFAULT NULL,
  `status` varchar(32) NOT NULL DEFAULT 'created',
  `package_name` varchar(255)  NULL,
  `entity_name` varchar(255)  NULL,
  `proc_def_key` varchar(255)  NOT NULL COMMENT 'key',
  `proc_def_id` varchar(255)  NOT NULL,
  `proc_def_name` varchar(255)  NOT NULL,
  `created_by` varchar(255)   NULL,
  `created_time` datetime NULL,
  `updated_by` varchar(255)   NULL,
  `updated_time` datetime NULL,
  `entity_attrs` text DEFAULT NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_request_template_form` FOREIGN KEY (`form_template`) REFERENCES `form_template` (`id`),
  CONSTRAINT `fore_request_template_group` FOREIGN KEY (`group`) REFERENCES `request_template_group` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `request`;
CREATE TABLE IF NOT EXISTS `request` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `form` varchar(64) DEFAULT NULL,
  `request_template` varchar(64)  NOT NULL,
  `proc_instance_id` varchar(64)  DEFAULT NULL,
  `reporter` varchar(64)  DEFAULT NULL,
  `report_time` datetime DEFAULT NULL,
  `emergency` varchar(16)  DEFAULT 'no',
  `report_role` varchar(64)  DEFAULT NULL,
  `attach_file` varchar(64)  DEFAULT NULL,
  `status` varchar(64)  NOT NULL DEFAULT 'created',
  `result` mediumtext ,
  `created_by` varchar(255)   NULL,
  `created_time` datetime NULL,
  `updated_by` varchar(255)   NULL,
  `updated_time` datetime NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_request_form` FOREIGN KEY (`form`) REFERENCES `form` (`id`),
  CONSTRAINT `fore_request_ref_template` FOREIGN KEY (`request_template`) REFERENCES `request_template` (`id`),
  CONSTRAINT `fore_request_report_role` FOREIGN KEY (`report_role`) REFERENCES `role` (`id`),
  CONSTRAINT `fore_request_attach_file` FOREIGN KEY (`attach_file`) REFERENCES `attach_file` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `request_template_role`;
CREATE TABLE IF NOT EXISTS `request_template_role` (
  `id` varchar(160)  NOT NULL,
  `request_template` varchar(64) NOT NULL,
  `role` varchar(64) NOT NULL,
  `role_type` varchar(64) NOT NULL DEFAULT 'MGMT',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_request_template` FOREIGN KEY (`request_template`) REFERENCES `request_template` (`id`),
  CONSTRAINT `fore_request_template_role` FOREIGN KEY (`role`) REFERENCES `role` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `task_template`;
CREATE TABLE IF NOT EXISTS `task_template` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `description` varchar(255)  DEFAULT NULL,
  `form_template` varchar(64) DEFAULT NULL,
  `request_template` varchar(64) DEFAULT NULL,
  `node_def_id` varchar(255)  DEFAULT NULL,
  `node_name` varchar(255)  DEFAULT NULL,
  `created_by` varchar(255)   NULL,
  `created_time` datetime NULL,
  `updated_by` varchar(255)   NULL,
  `updated_time` datetime NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_task_template_form` FOREIGN KEY (`form_template`) REFERENCES `form_template` (`id`),
  CONSTRAINT `fore_task_template_request` FOREIGN KEY (`request_template`) REFERENCES `request_template` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `task_template_role`;
CREATE TABLE IF NOT EXISTS `task_template_role` (
  `id` varchar(160)  NOT NULL,
  `task_template` varchar(64) NOT NULL,
  `role` varchar(64) NOT NULL,
  `role_type` varchar(64) NOT NULL DEFAULT 'MGMT',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_task_role_ref_template` FOREIGN KEY (`task_template`) REFERENCES `task_template` (`id`),
  CONSTRAINT `fore_task_template_ref_role` FOREIGN KEY (`role`) REFERENCES `role` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

DROP TABLE IF EXISTS `task`;
CREATE TABLE IF NOT EXISTS `task` (
  `id` varchar(64)  NOT NULL,
  `name` varchar(255)  NOT NULL,
  `description` varchar(255)  DEFAULT NULL,
  `attach_file` varchar(64)  DEFAULT NULL,
  `status` varchar(64) DEFAULT 'created',
  `version` varchar(64)  DEFAULT NULL,
  `request` varchar(64)  DEFAULT NULL,
  `parent` varchar(64)  DEFAULT NULL,
  `task_template` varchar(64)  DEFAULT NULL,
  `package_name` varchar(255)  DEFAULT NULL,
  `entity_name` varchar(255)  DEFAULT NULL,
  `proc_def_id` varchar(255)  DEFAULT NULL,
  `proc_def_key` varchar(255)  DEFAULT NULL,
  `proc_def_name` varchar(255)  DEFAULT NULL,
  `node_def_id` varchar(64)  DEFAULT NULL,
  `node_name` varchar(64)  DEFAULT NULL,
  `callback_url` varchar(255)  DEFAULT NULL,
  `callback_parameter` text  DEFAULT NULL,
  `reporter` varchar(64)  DEFAULT NULL,
  `report_time` datetime DEFAULT NULL,
  `report_role` varchar(64)  DEFAULT NULL,
  `emergency` varchar(16)  DEFAULT 'no',
  `result` text ,
  `created_by` varchar(255)   NULL,
  `created_time` datetime NULL,
  `updated_by` varchar(255)   NULL,
  `updated_time` datetime NULL,
  `del_flag` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT `fore_task_ref_request` FOREIGN KEY (`request`) REFERENCES `request` (`id`),
  CONSTRAINT `fore_task_attach_file` FOREIGN KEY (`attach_file`) REFERENCES `attach_file` (`id`),
  CONSTRAINT `fore_task_ref_template` FOREIGN KEY (`task_template`) REFERENCES `task_template` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT=' ';

SET FOREIGN_KEY_CHECKS = 1;
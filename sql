CREATE TABLE `company` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `address` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telephone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `logo` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `customer_uuid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `company_un` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `customer` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `provider` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `nick_name` varchar(100) NOT NULL,
  `description` varchar(100) NOT NULL,
  `user_id` varchar(100) NOT NULL,
  `avatar_url` varchar(300) NOT NULL,
  `location` varchar(300) NOT NULL,
  `access_token` varchar(300) NOT NULL,
  `access_token_secret` varchar(300) NOT NULL,
  `refresh_token` varchar(300) NOT NULL,
  `expires_at` varchar(100) NOT NULL,
  `id_token` varchar(100) NOT NULL,
  `jwt_token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `uuid` varchar(100) NOT NULL,
  `isAdmin` tinyint(1) NOT NULL DEFAULT '1',
  `last_update` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- piendb.payment_product definition

CREATE TABLE `payment_product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `product_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `payment_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `payment_product_FK` (`payment_code`),
  CONSTRAINT `payment_product_FK` FOREIGN KEY (`payment_code`) REFERENCES `orders` (`payment_code`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- piendb.product definition

CREATE TABLE `product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `product_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_name` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_size` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_type` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `uuid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_un` (`product_code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- piendb.orders2 definition

CREATE TABLE `orders` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `payment_code` varchar(50) NOT NULL,
  `company_code` varchar(50) NOT NULL,
  `customer_uuid` varchar(50) NOT NULL,
  `order_no` varchar(20) NOT NULL,
  `payment_due_date` date NOT NULL,
  `tax_invoice_no` varchar(20) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `days` varchar(2) NOT NULL,
  `sub_total_baht` varchar(10) NOT NULL,
  `vat` varchar(10) NOT NULL,
  `total_baht` varchar(10) NOT NULL,
  `create_date` timestamp NOT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `orders_un_payment_code` (`payment_code`),
  UNIQUE KEY `orders_un_order_no` (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `tb_menu` (
  `menu_id` varchar(36) NOT NULL,
  `menu_type_id` int(11) NOT NULL,
  `warteg_id` varchar(36) DEFAULT NULL,
  `menu_name` varchar(255) NOT NULL,
  `menu_detail` varchar(2000) DEFAULT NULL,
  `menu_picture` varchar(2000) DEFAULT NULL,
  `menu_price` int(11) NOT NULL,
  `updated_date` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- foodmenu.tb_menu_type definition

CREATE TABLE `tb_menu_type` (
  `menu_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_type_name` varchar(255) NOT NULL,
  `updated_date` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`menu_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

INSERT INTO tb_menu_type (menu_type_id,menu_type_name,updated_date) VALUES
	 (1,'Makanan','2021-02-10 22:43:22.957'),
	 (2,'Minuman','2021-02-10 22:43:56.606');
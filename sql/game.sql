
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '玩家ID',
  `acct_name` varchar(25) NOT NULL DEFAULT '' COMMENT '账号名字',
  `lord_name` varchar(25) NOT NULL DEFAULT '' COMMENT '玩家领主名',
  `sign` varchar(100) NOT NULL DEFAULT '' COMMENT '签名',
  `x` int(11) NOT NULL DEFAULT '0' COMMENT '出生点X坐标',
  `y` int(11) NOT NULL DEFAULT '0' COMMENT '出生点y坐标',
  `country` int(11) NOT NULL DEFAULT '0' COMMENT '所属国家',
  `level` int(11) NOT NULL DEFAULT '1' COMMENT '等级',
  `exp` int(11) NOT NULL DEFAULT '0' COMMENT '经验',
  `wood` int(11) NOT NULL DEFAULT '0' COMMENT '木材',
  `iron` int(11) NOT NULL DEFAULT '0' COMMENT '铁矿',
  `stone` int(11) NOT NULL DEFAULT '0' COMMENT '石料',
  `forage` int(11) NOT NULL DEFAULT '0' COMMENT '粮草',
  `gold` int(11) NOT NULL DEFAULT '0' COMMENT '金币',
  `diamond` int(11) NOT NULL DEFAULT '0' COMMENT '钻石',
  `bind_diamond` int(11) NOT NULL DEFAULT '0' COMMENT '绑定钻石',
  `decree` int(11) NOT NULL DEFAULT '0' COMMENT '政令',
  `army_order` int(11) NOT NULL DEFAULT '0' COMMENT '军令',
  `power` int(11) NOT NULL DEFAULT '0' COMMENT '势力值',
  `domain` int(11) NOT NULL DEFAULT '0' COMMENT '领地个数',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;


DROP TABLE IF EXISTS `t_user_mail`;
CREATE TABLE `t_user_mail` (
  `mail_id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '邮件id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '对应的玩家ID',
  `get_flag` smallint(3) NOT NULL DEFAULT '0' COMMENT '领取标志位，0未领取，1已领取，2已删除',
  `type` smallint(3) NOT NULL DEFAULT '0' COMMENT '邮件类型',
  `time` int(11) NOT NULL DEFAULT '0' COMMENT '创建邮件时间戳',
  `title` text COMMENT '邮件标题',
  `context` text COMMENT '邮件内容',
  `data_info` text COMMENT '附件',
  PRIMARY KEY (`mail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `t_user_task`;
CREATE TABLE `t_user_task` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '玩家ID',
  `task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务类型',
  `sub_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '子类型',
  `target` text NOT NULL COMMENT '任务目标',
  `task_state` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务状态',
  `other_data` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`user_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

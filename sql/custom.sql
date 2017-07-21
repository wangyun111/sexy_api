#抽奖记录
DROP TABLE IF EXISTS `lottery_record`;
CREATE TABLE `lottery_record` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键编号',
	`uid` int(11) NULL DEFAULT 0 COMMENT '用户编号',
	`phone` VARCHAR(11) NULL DEFAULT "" COMMENT '用户电话',
	`gift_name` VARCHAR(256) NULL DEFAULT "" COMMENT '礼物名称',
	`create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	PRIMARY KEY (`id`)
)comment="抽奖记录" ENGINE=InnoDB DEFAULT CHARSET=utf8;



#话费充值记录
DROP TABLE IF EXISTS `recharge_record`;
CREATE TABLE `recharge_record` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `order_no` VARCHAR(50) NULL DEFAULT "" COMMENT '充值订单号',
  `stream_id` VARCHAR(50) NULL DEFAULT 0 COMMENT '流水号',
  `recharge_price` VARCHAR(50) NULL DEFAULT 0 COMMENT '充值金额',
  `recharge_phone` VARCHAR(11) NULL DEFAULT "" COMMENT '充值号码',
  `status` ENUM('SUCCESS','FAILED','PROCESS') NULL DEFAULT "PROCESS" COMMENT '充值状态:SUCCESS 成功,FAILED 失败,PROCESS 处理中',
  `pay_money` VARCHAR(50) NULL DEFAULT 0 COMMENT '实付金额',
  `source_code` TINYINT(1) NULL DEFAULT 0 COMMENT '来源:0. 抽奖 1.夺宝 2.其他',
  `return_msg` VARCHAR(512) NULL DEFAULT "" COMMENT '失败时返回错误原因',
  `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
)comment="话费充值记录" ENGINE=InnoDB DEFAULT CHARSET=utf8;

#公告信息
DROP TABLE IF EXISTS `inform_bar`;
CREATE TABLE `inform_bar` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `title` VARCHAR(255) NULL DEFAULT "" COMMENT '公告标题',
  `content` TEXT COMMENT  '公告内容',
  `description` VARCHAR(1024) NULL DEFAULT "" COMMENT '公告描述',
  `skip_url` VARCHAR(255) NULL DEFAULT "" COMMENT '跳转地址',
  `begin_time` DATETIME NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` DATETIME NULL DEFAULT NULL COMMENT '结束时间',
  `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
)comment="公告信息" ENGINE=InnoDB DEFAULT CHARSET=utf8;

#查询电话号(空号,号码不支持)
DROP TABLE IF EXISTS `nonuse_phone`;
CREATE TABLE `nonuse_phone` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `uid` INT(11) NULL DEFAULT NULL COMMENT '用户id',
  `phone` VARCHAR(11) NULL DEFAULT NULL COMMENT '手机号',
  `is_normal` TINYINT(1) NULL DEFAULT '0' COMMENT '号码是否异常 0正常,1异常,2点数不足',
  `abnormal_result` VARCHAR(255) NULL DEFAULT NULL COMMENT '异常原因',
  `create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `phone`(`phone`),
)comment="查询电话号(空号,号码不支持)" ENGINE=InnoDB;


#录音记录
DROP TABLE IF EXISTS `tape_record`;
CREATE TABLE `tape_record` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `name` VARCHAR(11) NULL DEFAULT "" COMMENT '姓名',
  `phone` VARCHAR(11) NULL DEFAULT "" COMMENT '手机号',
  `loan_money` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '借款金额',
  `over_day` int(5) NULL DEFAULT 0 COMMENT '逾期天数',
  `wait_money` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '待还金额',
  `call_type` VARCHAR(50) NULL DEFAULT "" COMMENT '呼叫类型(实时状态)',
  `call_class` VARCHAR(50) NULL DEFAULT "" COMMENT '呼入呼出类型',
  `begin_time` DATETIME NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` DATETIME NULL DEFAULT NULL COMMENT '结束时间',
  `duration_second` int(5) NULL DEFAULT 0 COMMENT '呼叫时长(秒)',
  `user` VARCHAR(50) NULL DEFAULT "" COMMENT '坐席账号',
  `password` VARCHAR(50) NULL DEFAULT "" COMMENT '坐席密码',
  `session_id` VARCHAR(255) NULL DEFAULT "" COMMENT '通话标识',
  `create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `phone`(`phone`)
)comment="录音记录" ENGINE=InnoDB;



#录音记录
DROP TABLE IF EXISTS `tape_record`;
CREATE TABLE `tape_record` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `name` VARCHAR(11) NULL DEFAULT "" COMMENT '姓名',
  `phone` VARCHAR(11) NULL DEFAULT "" COMMENT '手机号',
  `loan_money` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '借款金额',
  `over_day` int(5) NULL DEFAULT 0 COMMENT '逾期天数',
  `wait_money` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '待还金额',
  `call_type` VARCHAR(50) NULL DEFAULT "" COMMENT '呼叫类型(实时状态)',
  `call_class` VARCHAR(50) NULL DEFAULT "" COMMENT '呼入呼出类型',
  `begin_time` DATETIME NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` DATETIME NULL DEFAULT NULL COMMENT '结束时间',
  `duration_second` int(5) NULL DEFAULT 0 COMMENT '呼叫时长(秒)',
  `user` VARCHAR(50) NULL DEFAULT "" COMMENT '坐席账号',
  `password` VARCHAR(50) NULL DEFAULT "" COMMENT '坐席密码',
  `session_id` VARCHAR(255) NULL DEFAULT "" COMMENT '通话标识',
  `create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `phone`(`phone`)
)comment="录音记录" ENGINE=InnoDB;



#今日头条回调数据
DROP TABLE IF EXISTS `toutiao_data`;
CREATE TABLE `toutiao_data` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `aid` VARCHAR(50) NULL DEFAULT "" COMMENT '广告计划id',
  `cid` VARCHAR(50) NULL DEFAULT "" COMMENT '广告创意id',
  `imei` VARCHAR(50) NULL DEFAULT "" COMMENT '用户终端的imei',
  `androidid` VARCHAR(50) NULL DEFAULT "" COMMENT '用户终端的androidid',
  `timestamp`  VARCHAR(50) NULL DEFAULT "" COMMENT '时间戳',
  `idfa` VARCHAR(50) NULL DEFAULT "" COMMENT 'ios适用ios6及以上',
  `phone_type` ENUM('IOS','ANDROID','') NULL DEFAULT "" COMMENT '手机类型',
  `create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
)comment="今日头条回调数据" ENGINE=InnoDB;









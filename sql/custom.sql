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

#录音地址
DROP TABLE IF EXISTS `tape_record_files`;
CREATE TABLE `tape_record_files` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `record_url` VARCHAR(11) NULL DEFAULT "" COMMENT '',
  `tape_id` int(11) NULL DEFAULT 0 COMMENT '关联录音记录id',
  `sys_id` int(11) NULL DEFAULT 0 COMMENT '关联坐席id',
  `create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `tapeId`(`tape_id`)
)comment="录音地址" ENGINE=InnoDB;


#费用减免审批
DROP TABLE IF EXISTS `cost_relief_approve`;
CREATE TABLE `cost_relief_approve` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `uid` INT(11) NULL DEFAULT 0 COMMENT '用户id',
  `phone` VARCHAR(20) NULL DEFAULT "" COMMENT '手机号',
  `state` ENUM('NONE','AGREE','REFUSE','DONE') NULL DEFAULT "NONE" COMMENT '审批状态',
  `rep_sch_id` INT(11) NULL DEFAULT 0 COMMENT '还款计划id',
  `loan_id` INT(11) NULL DEFAULT 0 COMMENT '借款id',
  `money` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '减免金额',
  `reason` VARCHAR(1024) NULL DEFAULT "" COMMENT '减免原因',
  `approve_result` VARCHAR(1024) NULL DEFAULT "" COMMENT '审批结果',
  `sys_id` int(11) NULL DEFAULT 0 COMMENT '审批人id',
  `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `uid`(`uid`)
)comment="费用减免审批" ENGINE=InnoDB;

#审批处理记录
DROP TABLE IF EXISTS `cost_relief_approve_record`;
CREATE TABLE `cost_relief_approve_record` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `cra_id` INT(11) NULL DEFAULT 0 COMMENT '审批id',
  `old_approve_id` INT(11) NULL DEFAULT 0 COMMENT '原审批人id',
  `new_approve_id` INT(11) NULL DEFAULT 0 COMMENT '新审批人id',
  `source` ENUM('APPROVE','DISPOSE') NULL DEFAULT "APPROVE" COMMENT '来源于审批,处理', 
  `opinion` VARCHAR(1024) NULL DEFAULT "" COMMENT '审批意见',
  `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
)comment="审批处理记录" ENGINE=InnoDB;


#借款续期日志
DROP TABLE IF EXISTS `loan_renew_logs`;
CREATE TABLE `loan_renew_logs` (
 `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
 `money` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '本金',
 `interest_fee` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '利息',
 `service_fee` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '服务费',
 `procedure_fee` DECIMAL(10,2) NULL DEFAULT 0 COMMENT '手续费', 
 `state` ENUM('NONE','FAILURE','SUCCESS') NULL DEFAULT "NONE" COMMENT '记录状态',
 `flow` TINYINT(1) NULL DEFAULT 0 COMMENT '记录流程(走到哪一步了)',
 `loan_id` INT(11) NULL DEFAULT 0 COMMENT '借款id',
 `schedule_id` INT(11) NULL DEFAULT 0 COMMENT '还款记录id',
 `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', PRIMARY KEY (`id`)
) COMMENT="借款续期日志" ENGINE=InnoDB;

#工作信息
DROP TABLE IF EXISTS `user_work_info`;
CREATE TABLE `user_work_info` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键,自增',
  `uid` INT(11) NULL DEFAULT 0 COMMENT '用户id',
  `company_name` VARCHAR(1024) NULL DEFAULT 0 COMMENT '利息',
  `work_station` VARCHAR(1024) NULL DEFAULT 0 COMMENT '服务费',
  `income` VARCHAR(1024) NULL DEFAULT 0 COMMENT '手续费',
  `state` VARCHAR(1024) NULL DEFAULT "NONE" COMMENT '记录状态',
  `flow` VARCHAR(1024) NULL DEFAULT 0 COMMENT '记录流程(走到哪一步了)',
  `company_phone` INT(11) NULL DEFAULT 0 COMMENT '借款id',
) COMMENT="借款续期日志" ENGINE=InnoDB;


package include

type PlayerAttr struct {
	UserId 			int32
	AcctName 		string		// 账号名
	LordName		string		// 玩家领主名
	Sign			string		// 签名
	X				int32		// 出生点X
	Y				int32		// 出生点y
	Country			int32		// 所属国家
	Level 			int32		// 等级
	Exp				int32		// 经验
	Wood			int32		// 木材
	Iron			int32		// 铁矿
	Stone			int32		// 石料
	Forage			int32		// 粮草
	Gold			int32		// 金币
	Diamond			int32		// 钻石
	BindDiamond		int32		// 绑定钻石
	Decree			int32		// 政令
	ArmyOrder		int32		// 军令
	Power			int32		// 势力值
	Domain			int32		// 领地个数

}

type PlayerMail struct {
	userId 			int32
	mailId			int32
}

type PlayerTask struct {
	userId 			int32
	taskId			int32
}
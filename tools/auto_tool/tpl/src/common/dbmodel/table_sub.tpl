package dbmodel

/*
	分表, 绑定分库
*/
// 例子
// // 账户表
// func (tb *AccountUser) SubTableNum() int { // 分表数量
// 	return 1
// }
// 
// func (tb *AccountUser) SubTable(baseValue int64) string { // 获取表名
// 	return subName(tb.SubName(), baseValue, tb.SubTableNum())
// }
// 
// func (tb *AccountUser) BindSubTreasury() DbSubTreasury { // 绑定表和分类库
// 	return DbSubTreasury_Account
// }

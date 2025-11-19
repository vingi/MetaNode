package task3

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance float64
}

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func CreateTable2() {
	db := ConnectDB()
	// 创建 accounts 表
	db.AutoMigrate(&Account{})
	// 创建 transactions 表
	db.AutoMigrate(&Transaction{})
}

func InsertAccount() {
	db := ConnectDB()
	// 插入账户 A 和账户 B
	db.Create(&Account{Balance: 1000})
	db.Create(&Account{Balance: 500})
}

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func Transfer(fromAccountID, toAccountID uint, amount float64) {
	db := ConnectDB()
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		panic("事务开启失败")
	}

	// 检查账户 A 的余额是否足够
	var fromAccount Account
	tx.Where("id = ?", fromAccountID).First(&fromAccount)
	if fromAccount.Balance < amount {
		// 余额不足，回滚事务
		tx.Rollback()
		panic("账户余额不足")
	}
	// 从账户 A 扣除 100 元
	fromAccount.Balance -= amount
	tx.Save(&fromAccount)

	// 向账户 B 增加 100 元
	var toAccount Account
	tx.Where("id = ?", toAccountID).First(&toAccount)
	toAccount.Balance += amount
	tx.Save(&toAccount)

	// 在 transactions 表中记录该笔转账信息
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	result := tx.Create(&transaction)

	if result.Error != nil {
		// 回滚事务
		tx.Rollback()
		panic("数据库操作失败")
	}
	// 提交事务
	tx.Commit()
}

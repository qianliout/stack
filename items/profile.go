package items

type NameCode struct {
	ID         int64
	Name       string
	Code       string `gorm:"uniqueIndex:uidex_code,priority:1"`
	Profile    int64
	CashFlow   int64
	Balance    int64
	StockPrice float64 // 当前股价
	CrawlDate  int64
	SHSZ       string `gorm:"column:shsz"`
}

type Industry struct {
	Name string `gorm:"uniqueIndex:uidex_indeustry_code,priority:2"`
	Code string `gorm:"uniqueIndex:uidex_indeustry_code,priority:1"`
}

// 利润表
type Profile struct {
	ID               int64
	Name             string
	Code             string `gorm:"uniqueIndex:uidex_profile,priority:1"`
	ReportingPeriod  string `gorm:"uniqueIndex:uidex_profile,priority:2"`
	OperateAllIncome float64
	OperateIncome    float64
	OperateAllCost   float64
	OperateCost      float64
	Tax              float64
	SalesExpense     float64
	ManageExpense    float64
	FinancialExpense float64
	DilutedEarn      float64 // 每股收益
}

// 上海股票的代码爬取
type NubSh struct {
	Result []DataSH `json:"result"`
}
type DataSH struct {
	Code string `json:"SECURITY_CODE_A"`
	Name string `json:"COMPANY_ABBR"`
}

type CashFlow struct {
	ID              int64
	Code            string `gorm:"uniqueIndex:uidex_cash,priority:1"`
	Name            string
	ReportingPeriod string `gorm:"uniqueIndex:uidex_cash,priority:2"`
	SalesCash       float64
	SumInFow        float64
	BuyCash         float64
	SumOutFow       float64
	NetCashFlow     float64
}

type Balance struct {
	ID              int64
	Name            string
	Code            string `gorm:"uniqueIndex:uidex_balance,priority:1"`
	ReportingPeriod string `gorm:"uniqueIndex:uidex_balance,priority:2"`
	MoneyFunds      float64
	TransFinance    float64
	Stock           float64
	ShortLoan       float64
	LongLoan        float64
}

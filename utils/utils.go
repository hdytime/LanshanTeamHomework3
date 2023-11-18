package utils

import "fmt"

type Goods struct {
	Name    string
	Price   int
	Num     int
	Loop    bool
	Details string
	key     int
	Direct  string
}

type ElecGoods struct {
	Goods
	Brand string
	Model string
}

type Manage interface {
	Check()
	Print()
	Buy()
	Exit()
	Menu()
}

func (g *ElecGoods) Check() {
	fmt.Println("当前库存数量有:", g.Num)
}

func (g *ElecGoods) Print() {
	g.Details += fmt.Sprintf("\n %v\t%v\t%v\t%v\t%v", g.Name, g.Price, g.Num, g.Brand, g.Model)
	fmt.Println(g.Details)
}

func (g *ElecGoods) Buy() {
	fmt.Println("当前库存数量为:", g.Num)
	fmt.Println("请输入购买数量")
	fmt.Scanln(&g.key)
	if g.key > g.Num {
		fmt.Println("购买数量大于库存数量")
		return
	}
	g.Num -= g.key
	fmt.Println("购买成功！")
}

func (g *ElecGoods) Exit() {
	fmt.Println("您确定要退出蓝山电子商务平台吗？ 请输入yes/no")

	for {
		fmt.Scanln(&g.Direct)
		if g.Direct != "yes" && g.Direct != "no" {
			fmt.Println("输入错误指令，请重新输入yes/no")
		}
		break
	}
	if g.Direct == "yes" {
		g.Loop = false
	}

}

func (g *ElecGoods) Menu() {
	for {
		fmt.Println("欢迎来到蓝山电子商务平台")
		fmt.Println("1.检查商品库存数量")
		fmt.Println("2.打印商品详细信息")
		fmt.Println("3.购买商品")
		fmt.Println("4.退出蓝山电子商务平台")
		fmt.Println("请选择上述功能(1-4)")
		fmt.Scanln(&g.key)
		switch g.key {
		case 1:
			g.Check()
		case 2:
			g.Print()
		case 3:
			g.Buy()
		case 4:
			g.Exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !g.Loop {
			break
		}
	}
}

func Initialinformation() *ElecGoods {
	return &ElecGoods{
		Goods: Goods{
			Name:    "手机",
			Price:   5200,
			Num:     200,
			Loop:    true,
			Details: " 商品名\t价格\t库存数量\t品牌\t型号",
			key:     0,
			Direct:  "",
		},
		Brand: "小米",
		Model: "14",
	}

}

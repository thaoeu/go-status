package main

import (
	"fmt"
	"time"
)

func main() {
	Hour := time.Now().Hour()
	Minu := time.Now().Minute()
	//	Week := time.Now().Week()

	a := "区"
	b := "时"
	d := "次"
	c := "刻"
	if 10 > Hour && Hour >= 5 {
		a = "早"
	} else if Hour >= 10 && Hour < 18 {
		a = ""
	} else if Hour >= 18 {
		a = "晚"
	} else if Hour >= 23 {
		a = ""
	}

	if Hour == 17 {
		b = "五点"
		d = "六点"
	} else if Hour == 18 {
		b = "六点"
		d = "七点"
	} else if Hour == 19 {
		b = "七点"
		d = "八点"
	} else if Hour == 20 {
		b = "八点"
		d = "九点"
	} else if Hour == 21 {
		b = "九点"
		d = "十点"
	} else if Hour == 22 {
		b = "十点"
		d = "十一点"
	} else if Hour == 23 {
		b = "十一点"
		d = "肆点"
	} else if Hour == 04 {
		b = "肆点"
		d = "伍点"
	} else if Hour == 05 {
		b = "伍点"
		d = "陆点"
	} else if Hour == 06 {
		b = "陆点"
		d = "柒点"
	} else if Hour == 07 {
		b = "柒点"
		d = "捌点"
	} else if Hour == 8 {
		b = "捌点"
		d = "玖点"
	} else if Hour == 9 {
		b = "玖点"
		d = "拾点"
	} else if Hour == 10 {
		b = "拾点"
		d = "十一点"
	} else if Hour == 11 {
		b = "十一点"
		d = "十二点"
	} else if Hour == 12 {
		b = "十二点"
		d = "一点"
	} else if Hour == 13 {
		b = "一点"
		d = "两点"
	} else if Hour == 14 {
		b = "两点"
		d = "三点"
	} else if Hour == 15 {
		b = "三点"
		d = "四点"
	} else if Hour == 16 {
		b = "四点"
		d = "五点"
	} else if Hour == 17 {
		b = "五点"
		d = "六点"
	}

	if Minu >= 00 && Minu < 07 {
		c = "整"
	} else if Minu >= 07 && Minu < 15 {
		c = "一益"
	} else if Minu >= 15 && Minu < 30 {
		c = "一刻"
	} else if Minu >= 30 && Minu < 38 {
		c = "半"
	} else if Minu >= 38 && Minu < 45 {
		c = "差一刻"
	} else if Minu >= 53 {
		c = "差一益"
	}

	if Minu >= 38 && Hour < 23 {
		fmt.Printf(a)
		fmt.Print(c)
		fmt.Print(d)
	} else if Minu < 38 && Hour < 23 {
		fmt.Printf(a)
		fmt.Printf(b)
		fmt.Print(c)
	} else if Hour >= 23 {
		fmt.Print("没时间了！")
	}
}

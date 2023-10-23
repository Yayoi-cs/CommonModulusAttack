package main

import (
	"fmt"
	"math/big"
)

func egcd(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	zero := big.NewInt(0)
	one := big.NewInt(1)

	if a.Cmp(zero) == 0 {
		return b, zero, one
	} else {
		g, y, x := egcd(new(big.Int).Mod(b, a), a)
		temp := new(big.Int)
		temp.Div(b, a)
		temp.Mul(temp, y)
		x.Sub(x, temp)
		return g, x, y
	}
}

func modinv(a, m *big.Int) *big.Int {
	_, x, _ := egcd(a, m)
	temp := new(big.Int)
	temp.Mod(x, m)
	return temp
}

func commonWorld(c1, c2, e1, e2, n *big.Int) *big.Int {
	gcd, s1, s2 := egcd(e1, e2)
	fmt.Println("GCD:", gcd)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	zero := big.NewInt(0)
	if s1.Cmp(zero) < 0 {
		s1.Neg(s1)
		c1.ModInverse(c1, n)
	} else if s2.Cmp(zero) < 0 {
		s2.Neg(s2)
		c2.ModInverse(c2, n)
	}

	v := new(big.Int).Exp(c1, s1, n)
	w := new(big.Int).Exp(c2, s2, n)
	x := new(big.Int)
	x.Mul(v, w)
	x.Mod(x, n)
	return x
}

func main() {
	c1, _ := new(big.Int).SetString("", 10) //crypt strings1
	c2, _ := new(big.Int).SetString("", 10) //crypt strings2
	e1 := big.NewInt(11)                    //e1
	e2 := big.NewInt(13)                    //e2
	n, _ := new(big.Int).SetString("", 10)  //common N
	result := commonWorld(c1, c2, e1, e2, n)
	fmt.Println("Result:", result)
}

package gommit

import (
	"fmt"
)

func ShowType(){
	fmt.Println("You can use those types :")
	for _, commitType := range o.Type {
		fmt.Println("	-", commitType)
	}
}

func ShowExample(){
	fmt.Println("The right commit format is :")
	fmt.Println("	"+string(colorBlue)+"<type>"+string(colorMagenta)+"(<scope>)"+string(colorReset)+": "+string(colorCyan)+"<short summary>"+string(colorReset))
	fmt.Println(string(colorBlue),"type : the commit type"+string(colorReset))
	fmt.Println(string(colorMagenta),"scope : to give more contect to your commit (optional)"+string(colorReset))
	fmt.Println(string(colorCyan),"short summary : a short summary of the code changes"+string(colorReset))
	fmt.Println("Exemples :")
	fmt.Println(string(colorGreen)+"	fix: change payload of login route	is a good commit message"+string(colorReset))
	fmt.Println(string(colorRed)+"	fix: fix				is not a good commit message"+string(colorReset))

}
func ShowReq(commit string) {
	ShowType()

	ShowExample()
}
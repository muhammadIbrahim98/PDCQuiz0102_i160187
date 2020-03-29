package PDC_Task

import(
	"fmt"
	"log"
)

func PrintCoronaStats(){

	var option int
	fmt.Println("Please Select and option:\n1. Print COVID-19 cases in Pakistan.\n2. Print COVID-19 cases in SouthKorea.\n3. Print COVID-19 cases in France.\n4. Print my personalized message to Coronavirus.\n0. Exit.")
	_, err := fmt.Scanf("%d", &option)
	if err != nil{
		log.Fatal(err)
	}
	flag := false
	if option == 0 || option == 1 || option == 2 || option == 3 || option == 4{
		flag = true
	}
	
	for !flag {
		fmt.Println("Please Select and option:\n1. Print COVID-19 cases in Pakistan.\n2. Print COVID-19 cases in SouthKorea.\n3. Print COVID-19 cases in France.\n4. Print my personalized message to Coronavirus.\n0. Exit.")
		_, err := fmt.Scanf("%d", &option)
		if err != nil{
			log.Fatal(err)
		}
		if option == 0 || option == 1 || option == 2 || option == 3 || option == 4{
			flag = true
		}
	}
	
	if option == 0{
		fmt.Println("GoodBye.")
	}else if option == 1{
		fmt.Println("Pakistan -> 1526")
	}else if option == 2{
		fmt.Println("South Korea -> 9583")
	}else if option == 3{
		fmt.Println("France -> 40174")
	}else if option == 4{
		fmt.Println("Corona, Please go away.")
	}
}

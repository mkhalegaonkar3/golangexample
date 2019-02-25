package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
func main(){
	truck1:=truck{
		vehicle:vehicle{
			doors:2,
			color:"red",
		},
		fourWheel:true,
	}
	sedan1:=sedan{
		vehicle:vehicle{
			doors:4,
			color:"black",
		},
		luxury:false,
	}
	fmt.Println("***TRUCK INFORMATION***")
	fmt.Println("Doors :",truck1.doors,"\nColor:",truck1.color,"\nis it Four Wheeler:",truck1.fourWheel)


	fmt.Println("----------\n****SEDAN INFORMATION****")
	fmt.Println("Doors :",sedan1.doors,"\nColor:",sedan1.color,"\nis it luxury:",sedan1.luxury)
	//checking boolean values
	if sedan1.luxury{
		fmt.Println("it is a luxury car")
	}else{
		fmt.Println("it is not a luxurius car")
	}
}

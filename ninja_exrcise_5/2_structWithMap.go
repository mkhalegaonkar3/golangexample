package main

import "fmt"

type student struct {
	fname   string
	lname   string
	hobbies []string
}

func main() {
	s1:=student{fname:"mayuresh",
		lname:"khalegaonkar",
		hobbies: []string{"dancing","cricket"},
	}
	s2:=student{fname:"chinmay",
		lname:"joshi",
		hobbies: []string{"guitar","poetry"},
	}
	m:=map[string]student{
		"khalegaonkar":s1,
			"joshi":s2,
	}
	//adding key and value into MAP explicitly
	m["phand"]=student{
		fname:"shon",
		lname:"phand",
		hobbies:[]string{"singing","art"},
	}
	for k,v:=range m{
		fmt.Println("values for key",k,"are")

		fmt.Println("\t",v.fname,"",v.lname)

		for j,val :=range v.hobbies{
			fmt.Println("\t",j," \tHobby",val)
		}
	}

	if n,ok:=m["phand"];ok{//delete function in map
		delete(m,"phand")
		fmt.Println(n.fname,"deleted from map")
	}
	fmt.Println("****AFTER DELETION THE KEY FROM MAP****")
	for k,v:=range m{
		fmt.Println("values for key",k,"are")

		fmt.Println("\t",v.fname,"",v.lname)

		for j,val :=range v.hobbies{
			fmt.Println("\t",j," \tHobby",val)
		}
	}

}


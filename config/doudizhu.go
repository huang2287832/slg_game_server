package config

type Doudizhu struct{
	Pool_id int
	Round_time int
	Max_money int
	Match_time []int
	Desc string
	
}

func Get_attr(key int) *Doudizhu {
	
	if key == 1 {
 		return &Doudizhu{Pool_id:1, Round_time:12, Max_money:10000, Match_time:[]int{3,4,5,6}, Desc:"fuck", }		
 
	}	
	if key == 2 {
 		return &Doudizhu{Pool_id:2, Round_time:12, Max_money:50000, Match_time:[]int{3,4,5,7}, Desc:"fuck", }		
 
	}	
	if key == 3 {
 		return &Doudizhu{Pool_id:3, Round_time:12, Max_money:100000, Match_time:[]int{3,4,5,8}, Desc:"fuck", }		
 
	}	
	if key == 4 {
 		return &Doudizhu{Pool_id:4, Round_time:12, Max_money:300000, Match_time:[]int{3,4,5,9}, Desc:"fuck", }		
 
	}
	return nil 
}	
package cache

import (
	"math"
	"strconv"
	"userauth/models"
)

const prefix  = "cache_"

type getListFromDb func()[]models.Links

type Hobby struct {
	List map[string]int
}


func GetList(clientId string,method getListFromDb)[]models.Links{
	 list := method()
	 hobbyList := getHobbyList(clientId)

	 hobby := Hobby{
		 List: hobbyList.(map[string]int),
	 }
	count := Count(clientId)
	return  makeSelfList(hobby,list,count)
}

func Set(key ,value string)bool{
	 _,err :=  RedisClient.Do("set",joinPrefix(key),value)
	 if err != nil{
	 	return false
	 }
	 return true
}

func Increment(clientId string,category string){
	RedisClient.Do("ZINCRBY",joinPrefix(clientId),1,category)
	incrementCount(clientId)
}

func Decrement(clientId string,category string){
	RedisClient.Do("ZINCRBY",joinPrefix(clientId),1,category)

}
//
func makeSelfList(hobby Hobby,list []models.Links,count int)[]models.Links{
	 res  := make([]models.Links,0)
	 for k,v:= range hobby.List{
	 		num := 0
		 breakNum := int(math.Ceil(float64(count/v)))
		 for _,vv := range list{
			if(vv.CategoryName == k){
				res = append(res,vv)
				num ++
			}
			if(num >breakNum){
				break;
			}
		 }
	 }
	 if(len(res) <60 ){
	 	for _,v := range list{
	 		canAppend := true
	 		for _,vv := range res{
	 			if(v.Name == vv.Name ){
					canAppend = false
				}
			}
	 		if(canAppend){
				res = append(res,v)
			}
	 		if(len(res)>=60){
	 			break;
			}
		}
	 }
	 return res
}


func Count(clientId string)int{
	res ,_  := RedisClient.Do("get",joinPrefix(clientId+"_count"))
	if res == nil{
		return 0
	}
	return int((res.([]uint8))[0])
}

func incrementCount(clientId string)interface{}{
	res,_ := RedisClient.Do("INCR",joinPrefix(clientId+"_count"))
	return res
}

func getHobbyList(client string)interface{}{
	res ,_ := RedisClient.Do("ZREVRANGE",joinPrefix(client),0,10,"WITHSCORES")
	tplmap := make(map[string]int)
	tplName := ""
	for k,v := range res.([]interface{}){
		if(k % 2 == 0){
			tplmap[string(v.([]byte))] = 0;
			tplName = string(v.([]byte))
		}else{
			 clientNum ,_ := strconv.Atoi(string(v.([]byte)))
			tplmap[tplName] = clientNum
		}
	}
	return tplmap
}


func joinPrefix(str string)string{
	return prefix+str
}




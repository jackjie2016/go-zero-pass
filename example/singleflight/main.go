package main

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
)

var errorNotExist = errors.New("not exist")

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	//模拟10个并发
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			data, err := getData("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

// 获取数据
//func getData(key string) (string, error) {
//	data, err := getDataFromCache(key)
//	if err == errorNotExist {
//		//模拟从db中获取数据
//		data, err = getDataFromDB(key)
//		if err != nil {
//			log.Println(err)
//			return "", err
//		}
//
//		//TOOD: set cache
//	} else if err != nil {
//		return "", err
//	}
//	return data, nil
//}

var gsf singleflight.Group

// 获取数据
func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		//模拟从db中获取数据
		v, err, _ := gsf.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
			//set cache
		})
		if err != nil {
			log.Println(err)
			return "", err
		}

		//TOOD: set cache
		data = v.(string)
	} else if err != nil {
		return "", err
	}
	return data, nil
}

// 模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

// 模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from database", key)
	return "data", nil
}

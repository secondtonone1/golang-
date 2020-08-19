package main

import (
	"context"
	"fmt"
	"golang-/gomongo/constants"
	model "golang-/gomongo/model"
	"golang-/gomongo/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	//初始化连接
	mongodb.Init()
	//insertOne()
	//insertMany()
	//findOne()
	//findMany()
	//updateOne()
	//updateMany()
	//findGroup()
	limitPage()
	//断开连接
	mongodb.Close()
}

//单条插入
func insertOne() {
	client := mongodb.DB.Mongo
	// 获取数据库和集合
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)
	userdata := model.UserData{}
	userdata.Age = 13
	userdata.BirthMonth = 11
	userdata.Number = 3
	userdata.Name = "zack"
	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()
	// 插入一条数据
	insertOneResult, err := collection.InsertOne(ctx, &userdata)
	if err != nil {
		fmt.Println("insert one error is ", err)
		return
	}
	log.Println("collection.InsertOne: ", insertOneResult.InsertedID)
	//将objectid转换为string
	docId := insertOneResult.InsertedID.(primitive.ObjectID)
	recordId := docId.Hex()
	fmt.Println("insert one ID str is :", recordId)
}

//多条插入
func insertMany() {
	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()

	client := mongodb.DB.Mongo
	// 获取数据库和集合
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)

	userdata1 := model.UserData{}
	userdata1.Age = 20
	userdata1.BirthMonth = 11
	userdata1.Number = 4
	userdata1.Name = "Lilei"

	userdata2 := model.UserData{}
	userdata2.Age = 20
	userdata2.BirthMonth = 12
	userdata2.Number = 5
	userdata2.Name = "HanMeiMei"

	var list []interface{}
	list = append(list, &userdata1)
	list = append(list, &userdata2)
	result, err := collection.InsertMany(ctx, list)
	if err != nil {
		fmt.Println("insert many error is ", err)
	}
	fmt.Println("insert many success, res is ", result.InsertedIDs)
}

//查找单个
func findOne() {

	client := mongodb.DB.Mongo
	// 获取数据库和集合
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)
	filter := bson.M{"name": "HanMeiMei"}

	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()
	singleResult := collection.FindOne(ctx, filter)
	if singleResult == nil || singleResult.Err() != nil {
		fmt.Println("find one error is ", singleResult.Err().Error())
		return
	}

	userData := &model.UserData{}
	err := singleResult.Decode(userData)
	if err != nil {
		fmt.Println("find one failed error is ", err)
		return
	}

	fmt.Println("find one success, res is ", userData)
}

//查询多个结果集，用cursor
func findMany() {
	client := mongodb.DB.Mongo
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()
	filter := bson.M{"birthMonth": bson.M{"$lte": 12}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("find res failed , error is ", err)
		return
	}
	defer cursor.Close(context.Background())

	result := make(map[string]*model.UserData)
	for cursor.Next(context.Background()) {
		ud := &model.UserData{}
		err := cursor.Decode(ud)
		if err != nil {
			fmt.Println("decode error is ", err)
			continue
		}
		result[ud.Name] = ud
	}

	fmt.Println("success is ", result)
	return
}

//更新
func updateOne() {
	client := mongodb.DB.Mongo
	collection, _ := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION).Clone()
	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()
	/*
		oid, err := primitive.ObjectIDFromHex(obj.RecordId)
		if err != nil {
			logging.Logger.Info("convert string from object failed")
			return err
		}
		filter := bson.M{"_id": oid}
	*/
	filter := bson.M{"name": "zack"}
	value := bson.M{"$set": bson.M{
		"number": 1024}}

	_, err := collection.UpdateOne(ctx, filter, value)
	if err != nil {
		fmt.Println("update user data failed, err is ", err)
		return
	}
	fmt.Println("update success !")
	return
}

//更新多条记录
func updateMany() {
	client := mongodb.DB.Mongo
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()
	var names = []string{"zack", "HanMeiMei"}
	filter := bson.M{"name": bson.M{"$in": names}}

	value := bson.M{"$set": bson.M{"birthMonth": 3}}
	result, err := collection.UpdateMany(ctx, filter, value)
	if err != nil {
		fmt.Println("update many failed error is ", err)
		return
	}
	fmt.Println("update many success !, result is ", result)
	return
}

//分组查询
func findGroup() {
	client := mongodb.DB.Mongo
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()
	//复杂查询，先匹配后分组
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"birthMonth": 3},
		},
		bson.M{"$group": bson.M{
			"_id":        bson.M{"birthMonthUid": "$birthMonth"},
			"totalCount": bson.M{"$sum": 1},
			"nameG":      bson.M{"$min": "$name"},
			"ageG":       bson.M{"$min": "$age"},
		},
		},

		//bson.M{"$sort": bson.M{"time": 1}},
	}
	fmt.Println("pipeline is ", pipeline)

	cursor, err := collection.Aggregate(ctx, pipeline)
	fmt.Println("findGroup cursor is ", cursor)
	if err != nil {
		fmt.Printf("dao.findGroup collection.Aggregate() error=[%s]\n", err)
		return
	}

	for cursor.Next(context.Background()) {
		doc := cursor.Current

		totalCount, err_2 := doc.LookupErr("totalCount")
		if err_2 != nil {
			fmt.Printf("dao.findGroup totalCount err_2=[%s]\n", err_2)
			return
		}

		nameData, err_4 := doc.LookupErr("nameG")
		if err_4 != nil {
			fmt.Printf("dao.findGroup insertDateG err_4=[%s]\n", err_4)
			return
		}

		ageData, err_5 := doc.LookupErr("ageG")
		if err_5 != nil {
			fmt.Printf("dao.findGroup ageG err_5=[%s]\n", err_5)
			continue
		}
		fmt.Println("totalCount is ", totalCount)
		fmt.Println("nameData is ", nameData)
		fmt.Println("ageData is ", ageData)
	}
}

//分页查询
func limitPage() {
	client := mongodb.DB.Mongo
	collection := client.Database(constants.DB_DATABASES).Collection(constants.DB_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), constants.QUERY_TIME_OUT)
	defer cancel()

	filter := bson.M{"age": bson.M{"$gte": 0}}

	SORT := bson.D{{"number", -1}}
	findOptions := options.Find().SetSort(SORT)
	//从第1页获取，每次获取10条
	skipTmp := int64((1 - 1) * 10)
	limitTmp := int64(10)
	findOptions.Skip = &skipTmp
	findOptions.Limit = &limitTmp
	cursor, err := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(context.Background())
	if err != nil {
		fmt.Println("limit page error is ", err)
		return
	}

	for cursor.Next(context.Background()) {
		ud := &model.UserData{}
		err := cursor.Decode(ud)
		if err != nil {
			fmt.Println("user data decode error is ", err)
			continue
		}

		fmt.Println("user data is ", ud)

	}
}

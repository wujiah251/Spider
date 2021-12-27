package dict

type DictRecord struct {
	Word string  `bson:"word"` //任务名
	Idf  float64 `bson:"idf"`  //shell命令
}

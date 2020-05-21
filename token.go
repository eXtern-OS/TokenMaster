package TokenMaster

type Token struct {
	TokenId string `bson:"id"`
	Expires int64  `bson:"expires"`
	Ip      string `bson:"ip"`
	Uid     string `bson:"uid"`
}

package global
import(
	"user-web/proto"

	
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
)
var(
	UserSrvClient proto.UserClient
	Trans ut.Translator
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 密码
		DB:       0,                // 默认数据库
	})
)
//rdc.Set(context.Background(), mobile, sms, time.Duration(global.SmsConfig.RedisInFo.Expire) * time.Second)//手机验证码有效时间
/*cachedValue, err := getCache(ctx, rdb, key)
	if err != nil {
		log.Fatalf("Failed to get cache: %v", err)
	}
	fmt.Printf("Cached Value: %s\n", cachedValue)/*
	/*err = deleteCache(ctx, rdb, key)
	if err != nil {
		log.Fatalf("Failed to delete cache: %v", err)
	}
	fmt.Println("Cache deleted successfully")*/

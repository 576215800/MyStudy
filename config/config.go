package config
import(
	"fmt"
	"github.com/spf13/viper"
)
type DBConfig struct{
	Host  		string
	Port  		string
	User     	string
	Password 	string
	DBType      string

}
func Init(){
	// viper默认的路径是在项目根目录
	viper.SetConfigFile("./config/server.yaml")
	//viper.Set("Address", "127.0.0.1:9090")  //統一把key处理成小写
	//err:=viper.WriteConfig()
	// 读取配置文件
	err := viper.ReadInConfig()
	if err!=nil{
		panic(fmt.Errorf("Fatal err read config file: %s\n",err))
	}
	//配置文件中主只要保持缩进一样就ok, 不强求是一个Tab的缩进
	username := viper.GetString("username")
	password := viper.GetString("password")
	DBInfo :=viper.GetStringMapString("DB")
	fmt.Println(DBInfo)
	//viper把key都统一转换为小写处理
	fmt.Println(username,":",password)
	db:=&DBConfig{}
	viper.UnmarshalKey("DB",db)
	fmt.Println("......")
	fmt.Println(db)
	cm := viper.Get("companies")
	fmt.Println(cm)


}
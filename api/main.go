package main

import (
	"net/http"
	"os/exec"
	"time"

	sqlite "profileyou/api/config/database"
	controllers "profileyou/api/controllers"
	"profileyou/api/infrastructure/persistance"
	"profileyou/api/usecase"

	// "profileyou/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
)

func main() {
	// connect to the database
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer connect.Close()

	// DI
	keywordRepository := persistance.NewKeywordPersistance(db)
	keywordUseCase := usecase.NewKeywordUseCase(keywordRepository)
	keywordController := controllers.NewKeywordController(keywordUseCase)

	r := gin.Default()
	r.LoadHTMLGlob("api/view/*html")

	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// Shell CommandからPython実行 ---------------
	// sentence := "a photo of an astronaut riding a horse on a swimming pool"
	// // command_line := "python3 api/create.py \"twon\\nwild\\ sleeping\\ dog\""
	// command_line := "python3 api/create.py"
	// // command_line := "python3 api/api.py test attr"
	// command := strings.Fields(command_line)
	// shell := os.Getenv("SHELL")
	// status, output := getstatusoutput(sentence, command...)
	// fmt.Printf("--- Result ---------------\n")
	// fmt.Printf("Shell        : %s\n", shell)
	// fmt.Printf("Command      : %s\n", command)
	// fmt.Printf("StatusCode   : %d\n", status)
	// fmt.Printf("ResultMessage: %s\n", output)
	// fmt.Printf("--------------------------\n")

	// ここまでShell CommandからPython実行 ---------------

	// list all the keywords
	r.GET("/", keywordController.Index)
	r.GET("/keywords", keywordController.GetAllKeywordsGin)
	// list one keyword
	r.GET("/keywords/:id", keywordController.DetailKeyword)
	// create a new keyword
	r.POST("/keyword/create/", keywordController.CreateKeyword)
	r.POST("/keyword/update/", keywordController.UpdateKeyword)
	r.POST("/keyword/delete", keywordController.DeleteKeyword)
	r.GET("/message", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.Run(":8080")

	// out, err := exec.Command("/bin/bash", "python3 api/api.py").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(out))

}

func getstatusoutput(cmd string, args ...string) (status int, output string) {
	exec_command := exec.Command(args[0], args[1], cmd)
	std_out, std_err := exec_command.Output()
	status = exec_command.ProcessState.ExitCode()
	if std_err != nil {
		output = std_err.Error()
	} else {
		output = string(std_out)
	}
	return
}

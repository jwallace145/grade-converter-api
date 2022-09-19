package api

/*
IMPORTS
*/
import (
	"example/grade-converter-api/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
CONSTANTS
*/
const AWS_REGION = "us-east-1"

/*
STRUCTS
*/
type GradeConverterAPI struct {
	database *database.DynamoDBGateway
	router   *gin.Engine
}

type Env struct {
	db *database.DynamoDBGateway
}

/*
METHODS
*/
func (env *Env) GetHealth(context *gin.Context) {
	fmt.Printf("Getting health status of Grade Converter API\n")
	context.String(http.StatusOK, "healthy")
}

func (env *Env) GetGradeByFrench(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		env.db.GetGradeByFrench(context.Param("grade")),
	)
}

func (env *Env) GetGradeByYDS(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		env.db.GetGradeByYDS(context.Param("grade")),
	)
}

func (env *Env) PutGrade(context *gin.Context) {
	grade := database.Grade{}
	if err := context.BindJSON(&grade); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	env.db.PutGrade(grade)
	context.IndentedJSON(http.StatusCreated, grade)
}

func (env *Env) PutGrades(context *gin.Context) {
	grades := []database.Grade{}
	if err := context.BindJSON(&grades); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i := 0; i < len(grades); i++ {
		env.db.PutGrade(grades[i])
	}
	context.IndentedJSON(http.StatusCreated, grades)
}

func New(host string, port int) GradeConverterAPI {
	hostport := host + ":" + strconv.Itoa(port)

	fmt.Printf("Running Grade Converter API on %s", hostport)

	database := database.New(AWS_REGION)
	router := gin.New()
	env := &Env{db: database}

	router.GET("/health", env.GetHealth)
	router.GET("/grade/french/:grade", env.GetGradeByFrench)
	router.GET("/grade/yds/:grade", env.GetGradeByYDS)
	router.POST("/grade", env.PutGrade)
	router.POST("/grades", env.PutGrades)

	router.Run(hostport)

	return GradeConverterAPI{
		database: database,
		router:   router,
	}
}

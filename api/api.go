package api

/*
IMPORTS
*/
import (
	"example/grade-converter-api/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
CONSTANTS
*/
const AWS_REGION string = "us-east-1"

/*
STRUCTS
*/
type GradeConverterAPI struct {
	db     *database.DynamoDBGateway
	router *gin.Engine
}

/*
METHODS
*/
func (api *GradeConverterAPI) GetHealth(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"status": "healthy",
		},
	)
}

func (api *GradeConverterAPI) GetGradeByFrench(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		api.db.GetGradeByFrench(context.Param("grade")),
	)
}

func (api *GradeConverterAPI) GetGradeByYDS(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		api.db.GetGradeByYDS(context.Param("grade")),
	)
}

func (api *GradeConverterAPI) PutGrade(context *gin.Context) {
	grade := database.Grade{}
	if err := context.BindJSON(&grade); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	api.db.PutGrade(grade)
	context.IndentedJSON(http.StatusCreated, grade)
}

func (api *GradeConverterAPI) PutGrades(context *gin.Context) {
	grades := []database.Grade{}
	if err := context.BindJSON(&grades); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i := 0; i < len(grades); i++ {
		api.db.PutGrade(grades[i])
	}
	context.IndentedJSON(http.StatusCreated, grades)
}

func New(host string, port int) *GradeConverterAPI {
	hostport := host + ":" + strconv.Itoa(port)

	log.Printf("Running Grade Converter API on %s", hostport)

	api := &GradeConverterAPI{
		db:     database.New(AWS_REGION),
		router: gin.New(),
	}

	api.router.GET("/health", api.GetHealth)
	api.router.GET("/grade/french/:grade", api.GetGradeByFrench)
	api.router.GET("/grade/yds/:grade", api.GetGradeByYDS)
	api.router.POST("/grade", api.PutGrade)
	api.router.POST("/grades", api.PutGrades)

	api.router.Run(hostport)

	return api
}

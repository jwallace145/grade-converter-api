package database

/*
IMPORTS
*/
import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davecgh/go-spew/spew"
)

/*
CONSTANTS
*/
const DDB_GRADES_TABLE = "ClimbingGrades"

/*
STRUCTS
*/
type DynamoDBGateway struct {
	client *dynamodb.DynamoDB
}

type Grade struct {
	YDS     string `json:"yds"`
	French  string `json:"french"`
	British string `json:"british"`
}

/*
METHODS
*/
func New(region string) *DynamoDBGateway {
	fmt.Printf("Creating DynamoDB Gateway in region %s\n", region)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		fmt.Printf("Unexpected error occurred creating a new AWS session: %s", err)
	}

	ddb := DynamoDBGateway{
		client: dynamodb.New(sess),
	}

	return &ddb
}

func (ddb DynamoDBGateway) GetGradeByFrench(french string) Grade {
	fmt.Printf("Getting '%s' grade from table '%s'\n", french, DDB_GRADES_TABLE)

	result, err := ddb.client.Query(
		&dynamodb.QueryInput{
			TableName: aws.String(DDB_GRADES_TABLE),
			IndexName: aws.String("french-index"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":french": {
					S: aws.String(french),
				},
			},
			KeyConditionExpression: aws.String("french = :french"),
		},
	)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	grade := Grade{
		YDS:     *result.Items[0]["yds"].S,
		French:  *result.Items[0]["french"].S,
		British: *result.Items[0]["british"].S,
	}

	spew.Dump(grade)

	return grade
}

func (ddb DynamoDBGateway) GetGradeByYDS(yds string) Grade {
	fmt.Printf("Getting '%s' grade from table '%s'\n", yds, DDB_GRADES_TABLE)

	result, err := ddb.client.GetItem(
		&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"yds": {
					S: aws.String(yds),
				},
			},
			TableName: aws.String(DDB_GRADES_TABLE),
		})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	grade := Grade{
		YDS:     *result.Item["yds"].S,
		French:  *result.Item["french"].S,
		British: *result.Item["british"].S,
	}

	spew.Dump(grade)

	return grade
}

func (ddb DynamoDBGateway) PutGrade(grade Grade) {
	fmt.Printf("Adding grade to table '%s':\n", DDB_GRADES_TABLE)
	spew.Dump(grade)

	_, err := ddb.client.PutItem(
		&dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"yds": {
					S: aws.String(grade.YDS),
				},
				"french": {
					S: aws.String(grade.French),
				},
				"british": {
					S: aws.String(grade.British),
				},
			},
			TableName: aws.String(DDB_GRADES_TABLE),
		})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				fmt.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

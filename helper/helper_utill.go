package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"net/http"
)

func GetModelFromBody(c *gin.Context) model_other.ReqModel {
	var reqModel model_other.ReqModel
	reqModel.Language = c.Query("lang")

	if err := c.Bind(&ReqModel); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//fmt.Printf("%+v/n", ReqModel)
	return ReqModel
}

func ConvertToJson(data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(string(result))
}

func ConcatJoin(typeJoin, leftTable, rightTable, joinKey string) string {
	if typeJoin == "" || leftTable == "" || rightTable == "" || joinKey == "" {
		return ""
	}
	sql := fmt.Sprintf("%s join %s on %s.%s = %s.%s",
		typeJoin, rightTable, leftTable, joinKey, rightTable, joinKey)
	fmt.Println(sql)
	return sql
}

func ConvertUUIDtoStringMap(uuidList []uuid.UUID) map[string]bool {
	strMap := make(map[string]bool)
	for _, wa := range uuidList {
		strMap[wa.UUID.String()] = true
	}
	return strMap
}
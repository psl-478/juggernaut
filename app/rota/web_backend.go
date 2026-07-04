package stand

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"io"
	logmgr "judgement/config/log"
	"net/http"
	"path/filepath"
	"strconv"
)

import (
	"encoding/json"
)

type UpdateValue struct {
	Row      int    `json:"row"`
	Header   string `json:"header"`
	NewValue string `json:"newValue"`
}

func Data(c *gin.Context) {
	// 打开Excel文件
	f, err := excelize.OpenFile("./rota/standby.xlsx")
	if err != nil {
		logmgr.Log.Errorf("cannot open file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"cannot open file:": err})
		return
	}
	defer f.Close()

	// 获取第一个工作表的名称
	sheetName := f.GetSheetName(0)

	// 获取工作表的所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		logmgr.Log.Errorf("cannot get rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"cannot get rows:": err})
		return
	}

	// 用于存储JSON数据的切片
	var data []map[string]interface{}

	// 读取表头
	headers := []string{"mouth", "day", "week", "stand", "backupone", "backuptwo", "restone", "resttwo", "weekday"}
	// 读取数据行
	for _, row := range rows[1:] {
		rowMap := make(map[string]interface{})
		for i, header := range headers {
			rowMap[header] = row[i]
		}
		data = append(data, rowMap)
	}

	c.JSON(http.StatusOK, data)

}

func Update(c *gin.Context) {
	var u UpdateValue
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logmgr.Log.Errorf("cannot read body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"cannot get rows:": err})
		return
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &u)
	if err != nil {
		logmgr.Log.Errorf("cannot unmarshal body: %v", err)
		return
	}
	//
	// 打开一个已存在的 Excel 文件
	f, err := excelize.OpenFile("./standby.xlsx")
	if err != nil {
		logmgr.Log.Errorf("无法打开文件: %s", err)
		return
	}

	// 获取工作表的名称
	sheetName := f.GetSheetName(0)

	// 修改单元格的值
	var cell string
	switch u.Header {
	case "stand":
		cell = "D" + strconv.Itoa(u.Row+2)
	case "backupone":
		cell = "E" + strconv.Itoa(u.Row+2)
	case "backuptwo":
		cell = "F" + strconv.Itoa(u.Row+2)
	case "restone":
		cell = "G" + strconv.Itoa(u.Row+2)
	case "resttwo":
		cell = "H" + strconv.Itoa(u.Row+2)
	}
	fmt.Printf("实际单元格的值是%v", cell)
	f.SetCellValue(sheetName, cell, u.NewValue)

	// 保存修改后的文件
	if err := f.Save(); err != nil {
		logmgr.Log.Errorf("无法保存文件: %s", err)
		return
	}

	logmgr.Log.Info("文件已成功修改并保存")
}

func Upload(c *gin.Context) {

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建文件保存的目录
	saveDir := "./rota/"

	// 生成文件保存的路径
	savePath := filepath.Join(saveDir, file.Filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": file.Filename})
}

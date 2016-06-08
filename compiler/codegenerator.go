package compiler

import "ExcelConvert/def"
import "fmt"
import "strings"
import "bytes"

type TypeIndex int

var fieldList []string



type FieldInfo struct {
    fieldType string
    typeIndex TypeIndex
    fieldName string
    isArray bool
    num int
}

func init()  {
    fieldList = []string{
        "int",
        "bool",
        "float",
        "string",
    }
}

func getTypeIndex(fieldType string) (TypeIndex, error) {
    
    
    return 0, nil
}

// date format int string int[]
func analyzeFieldInfo(data [][]string) ([]FieldInfo, error) {
    m := make(map[string]FieldInfo)
    countMap := make(map[string]int)
    if len(data) < def.NameLine {
        return nil, fmt.Errorf("invalid data data row = %d", len(data))
    }
    fInfos := make([]FieldInfo, len(data[0]))
    var index int
    for i := 0; i < len(data[0]); i++ {
        var isArray bool
        fieldType := data[0][i] 
        if strings.Contains(data[0][i], "[]") {
            isArray = true
            fieldType = strings.Split(data[0][i], "[]")[0]
        }
               
        _, ok := m[data[1][i]]
        if ok && !isArray{
            return nil, fmt.Errorf("Error ! duplicate fieldname %s ", data[1][i])
        }
        var fInfo FieldInfo
        fInfo.fieldName = data[1][i]
        fInfo.fieldType = fieldType
        fInfo.isArray = isArray
        countMap[data[1][i]] = countMap[data[1][i]] + 1   
        m[fInfo.fieldName] = fInfo
        if countMap[data[1][i]] == 1 {
            fInfos[index] = fInfo
            index++
        }
    }
    
    for i, v := range(fInfos) {
        if v.isArray {
            fInfos[i].num = countMap[v.fieldName]
        }
    }    
    return fInfos[0:index], nil
}

func generatorCsharpStruct(fInfos []FieldInfo, basename string) string {
    var buff bytes.Buffer
    
    buff.WriteString("using System;\n")
    buff.WriteString("using System.Colletions.Generic;\n")
    buff.WriteString("using System.Linq;\n")
    buff.WriteString("using System.Text;\n")
    buff.WriteString("\n")
    buff.WriteString("namespace GenTable \n")
    buff.WriteString("{")
    buff.WriteString("\tpublic class ")
    buff.WriteString(basename + "Data\n")
    buff.WriteString("\t{")
                        
       
    //for i, v := range(fInfos) {
        
    //}
    
    return buff.String()
}


func AutoGenerator(data [][]string, path, basename string, lang string) error {
    fInfos, err := analyzeFieldInfo(data)
    if nil != err {
        return err
    }
    
    fmt.Print(fInfos)
    
    return nil
}
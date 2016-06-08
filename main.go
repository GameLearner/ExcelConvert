package main

import (
    "fmt"
    "flag"
    "os"
    "io/ioutil"
    "strings"
    "bytes"
    "ExcelConvert/def"
    "ExcelConvert/compiler"
)

var isServer *bool
var languageType *string
var outputDir *string
var dirName string

func init()  {
    isServer = flag.Bool("server", false, "export server/client txt default client")
    languageType = flag.String("lang", "csharp", "export language type, support: csharp or go default csharp")
    outputDir = flag.String("output", "./", "output directory")
}

func isDirExists(path string) bool {
    fi, err := os.Stat(path)
    if err != nil {
        return os.IsExist(err)
    } else {
        return fi.IsDir()
    }
}

func clearDir(path string)  {
    files, err := ioutil.ReadDir(dirName) 
    if nil != err {
        fmt.Println(err)
        return
    }
    
    for _, f := range(files) {
        err := os.Remove(path + "/" + f.Name())
        if nil != err {
            fmt.Println(err)
        }
    }
}

func exportTxt(data [][]string, path, basename string, isserver bool) error {
    filename := path + "/" + basename + ".txt"
    file, err := os.Create(filename)
    
    if nil != err {
        return err
    }
    
    defer func(){
        file.Close()
    }()
    
    if len(data) <= def.ClassLine {
        return fmt.Errorf("Invalid Excel row < %d", def.ClassLine)
    }
    var buff bytes.Buffer
    
    filterstr := "client"
    
    if isserver {
        filterstr = "client"
    } else {
        filterstr = "server"
    }
    
    for i := def.ClassLine + 1; i < len(data); i++ {
        colnum := len(data[i])
        for j := 0; j < colnum - 1; j++ {
            if data[def.ClassLine][j] != filterstr {
                buff.WriteString(data[i][j] + "\t")  
            }
        }
        if colnum >= 1 {
            buff.WriteString(data[i][colnum - 1])
        }
        buff.WriteString("\n")
    }
    file.Write(buff.Bytes())
    return nil  
}

func AutoConvert(dirName string) error {
    files, err := ioutil.ReadDir(dirName) 
    if nil != err {
        fmt.Println(err)
    }
    
    if false == isDirExists(*outputDir) {
        err := os.Mkdir(*outputDir, os.ModePerm)
        if nil != err {
            fmt.Println(err)
            return err
        }
    } else {
        //clearDir(*outputDir)
    }
    
    for _, finfo := range(files) {
        pathname := dirName + "/" + finfo.Name()
        _, err := os.OpenFile(pathname, os.O_RDONLY, 0666)
        fmt.Println("open file" + finfo.Name())
        if nil != err {
            fmt.Println(err)
            continue
        }
        data, err := readExcelData(pathname)
        if nil != err {
            fmt.Println(err)
            continue
        }
        splits := strings.Split(finfo.Name(), ".")
        basename := splits[0]
        exportTxt(data, *outputDir, basename, *isServer)
        
    }    
    return nil
}


func main()  {
     //excelTest()
     flag.Parse()
     fmt.Printf("%v\n", flag.Args())
     
     data := make([][]string, 5)
     fmt.Println(data)
     
     if len(flag.Args()) < 1 {
         fmt.Println("No Input Directory")
         return 
     }
     
     dirName = flag.Args()[0]
     
     AutoConvert(dirName)
}
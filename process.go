package main

import (
    "fmt"
    "os"
   "path/filepath"
    "io"
)

func get_args() {
    args := os.Args[1:]
    fmt.Println(args)
}

func read_file(target string) ( []byte, int){
    
    file, err := os.Open(target)
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
    data := make([]byte, 64)
    var n_read int
    for{
        n, err := file.Read(data)
        if err == io.EOF{  
            break  
        }
        n_read = n
    }
    return data, n_read
}

func count_lines(data []byte, amount int) int {
    i := 0
    for _, b := range data[:amount] {
            if b == '\n' {
                i++
            }
        }
        i++
        return i
}
func main() {
    /* get args from command line */
    args := os.Args[1:]
    input_dir := args[0]
    //output_dir := args[1]
   
   /* get list of files in dir */
    var files[] string
    root := input_dir
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    
    if err != nil {
        panic(err)
    }

   /* iterate through files */
    for _, file := range files[1:] {

        /* defining variables */
        data := make([]byte, 64)
        var amount int
        var line_num int

        /* file name */
        fmt.Println(file)

        /* get lines amount */
        data, amount = read_file(file)
        line_num = count_lines(data, amount)
        fmt.Println(line_num)
    }
}
package main

import (
    "fmt"
    "os"
   "path/filepath"
    "io"
    "strings"
    "sync"
)

func get_args() {
    args := os.Args[1:]
    fmt.Println(args)
}

func file_size(target string) (int64){

    file, err := os.Open(target)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    fi, err := file.Stat()
    return fi.Size()
}

func read_file(target string) ( []byte, int){

    file, err := os.Open(target)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    fi, err := file.Stat()
    data := make([]byte, fi.Size())
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

func write_file(file_name string, target_dir string, data string) {

    file_path := target_dir + "\\" + file_name
    if _, err_existance := os.Stat(target_dir)
    os.IsNotExist(err_existance){
        os.MkdirAll(target_dir, 0755)
    }
    file, err := os.Create(file_path)
    if err != nil{
        fmt.Println("Unable to create file:", err)
        os.Exit(1)
    }
    defer file.Close()
    file.WriteString(data)
}

func name_generator(file_path string, extension string) string {
    string_arr := strings.Split(file_path, "\\")
    file_name := strings.Split(string_arr[len(string_arr)-1], ".")
    full_file_name := file_name[0] + "."+ extension
    return full_file_name
}

func worker(wg *sync.WaitGroup, cs chan string, file string, output_dir string) {

        defer wg.Done()
        cs <- "Worker:" +" "+ file

        /* defining variables */
        data := make([]byte, file_size(file))
        var amount int
        var line_num int

        /* get lines amount */
        data, amount = read_file(file)
        line_num = count_lines(data, amount)

        /* write output */
        res_file_name := name_generator(file,"res")
        write_file(res_file_name, output_dir, fmt.Sprintf("%d",line_num))

}

func monitorWorker(wg *sync.WaitGroup, cs chan string) {
    wg.Wait()
    close(cs)
}

func printWorkerLog(cs <-chan string, done chan<- bool) {
   for i := range cs {
         fmt.Println(i)
    }
    done <- true
}

func main() {
    /* get args from command line */
    args := os.Args[1:]
    input_dir := args[0]
    output_dir := args[1]

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

    wg := &sync.WaitGroup{}
    cs := make(chan string)

   /* iterate through files */
    for _, file := range files[1:] {
        wg.Add(1)
        go worker(wg, cs, file, output_dir)

    }
    go monitorWorker(wg, cs)
    done := make(chan bool, 1)
    go printWorkerLog(cs, done)
    <-done
    fmt.Println("Total number of processed files: " + fmt.Sprintf("%d", len(files) - 1))

}

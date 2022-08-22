// package main

// import (
// 	"fmt"
//     "reflect"
//     "log"
//     "net/http"
//     "database/sql"
//     "time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
//     test_query_1()
//     // fmt.Println(secret.db_info)
// }

// func test_query_1() {
//     db, err := sql.Open("mysql", "")

//     if err != nil {
//         panic(err)
//     }
//     db.SetConnMaxLifetime(time.Minute * 3)
//     db.SetMaxOpenConns(10)
//     db.SetMaxIdleConns(10)
    
//     fmt.Println("connect success", db)
    
//     defer db.Close()
// }

// func test_server_1() {
//     http.HandleFunc("/", test_index_1)
//     log.Println("Start Server")
//     log.Fatal(http.ListenAndServe(":9000", nil))
// }

// func test_index_1(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "안녕하세요, %s", r.URL.Path[1:])
//     log.Println("Received Request")
// }

// func test_for_2() {

//     sum := 0

//     for i := 1; i <= 10; i++ {
//         sum += i
//     }

//     fmt.Printf("1부터 10까지의 합 : %d\n", sum)
// }

// func test_for_1() {

//     for i := 0; i < 10; i++ {
//         fmt.Println(i, i + i)
//     }
// }

// func test_print_1() {

//     fmt.Print("줄 넘김 없음")
//     fmt.Print("줄 넘기려면 이걸 해야돼.\n")

//     fmt.Println("이건 줄 넘김 있음")

//     name := "멍멍멍뭉"
//     fmt.Printf("내 이름은 %s이다.", name)
//     fmt.Printf("근데 이건 줄넘김 없음.")
//     fmt.Printf("줄 넘기려면 이걸 해야돼.\n")
// }

// func test_array_5() {
//     var array_1 [5]int = [5]int{1, 2, 3, 4, 5}
//     var array_2 [5]int
//     var array_3 []int

//     array_2 = array_1
//     array_3 = array_1[:]

//     array_1[0] = 100
//     array_2[0] = 200

//     fmt.Println("array_1 Type :", reflect.TypeOf(array_1)) // [5]int
//     fmt.Println("array_2 Type :", reflect.TypeOf(array_2)) // [5]int
//     fmt.Println("array_3 Type :", reflect.TypeOf(array_3)) // []int
//     fmt.Println("array_1 Data :", array_1) // [100 2 3 4 5]
//     fmt.Println("array_2 Data :", array_2) // [200 2 3 4 5]
//     fmt.Println("array_3 Data :", array_3) // [100 2 3 4 5]
//     fmt.Println("array_1 Length :", len(array_1)) // 5
//     fmt.Println("array_2 Length :", len(array_2)) // 5
//     fmt.Println("array_3 Length :", len(array_3)) // 5
// }


// func test_array_4() {
//     var array_1 [5]int = [5]int{1, 2, 3, 4, 5}
//     var array_2 = [...]int{1, 2, 3, 4, 5}
    
//     fmt.Println("array_1 Type :", reflect.TypeOf(array_1)) // [5]int
//     fmt.Println("array_2 Type :", reflect.TypeOf(array_2)) // [5]int
//     fmt.Println("array_1 Data :", array_1) // [1 2 3 4 5]
//     fmt.Println("array_2 Data :", array_2) // [1 2 3 4 5]
//     fmt.Println("array_1 Length :", len(array_1)) // 5
//     fmt.Println("array_2 Length :", len(array_2)) // 5
// }

// func test_array_3() {
//     var array_1 [5]int = [5]int{1, 2, 3, 4, 5}
//     var array_2 [5]int
//     array_2 = array_1
//     var array_3 [5]int
    
//     fmt.Println("array_1 Type :", reflect.TypeOf(array_1)) // [5]int
//     fmt.Println("array_2 Type :", reflect.TypeOf(array_2)) // [5]int
//     fmt.Println("array_3 Type :", reflect.TypeOf(array_3)) // [5]int
//     fmt.Println("array_1 Data :", array_1) // [1 2 3 4 5]
//     fmt.Println("array_2 Data :", array_2) // [1 2 3 4 5]
//     fmt.Println("array_3 Data :", array_3) // [0 0 0 0 0]
//     fmt.Println("array_1 Length :", len(array_1)) // 5
//     fmt.Println("array_2 Length :", len(array_2)) // 5
//     fmt.Println("array_3 Length :", len(array_3)) // 5
// }

// func test_array_2() {
//     var array_1 [5]int = [5]int{1, 2, 3, 4, 5}
//     var array_2 [5]int
//     var array_3 []int
    
//     fmt.Println("array_1 Type :", reflect.TypeOf(array_1)) // [5]int
//     fmt.Println("array_2 Type :", reflect.TypeOf(array_2)) // [5]int
//     fmt.Println("array_3 Type :", reflect.TypeOf(array_3)) // []int
//     fmt.Println("array_1 Data :", array_1) // [1 2 3 4 5]
//     fmt.Println("array_2 Data :", array_2) // [0 0 0 0 0]
//     fmt.Println("array_3 Data :", array_3) // []
//     fmt.Println("array_1 Length :", len(array_1)) // 5
//     fmt.Println("array_2 Length :", len(array_2)) // 5
//     fmt.Println("array_3 Length :", len(array_3)) // 0
// }

// func test_array_1() {
//     var array_1 [5]int = [5]int{1, 2, 3, 4, 5}
//     var array_2 [5]int
//     var array_3 []int

//     array_2 = array_1
//     array_3 = array_1[:]
    
//     fmt.Println("array_1 Type :", reflect.TypeOf(array_1)) // [5]int
//     fmt.Println("array_2 Type :", reflect.TypeOf(array_2)) // [5]int
//     fmt.Println("array_3 Type :", reflect.TypeOf(array_3)) // []int
//     fmt.Println("array_1 Data :", array_1) // [1 2 3 4 5]
//     fmt.Println("array_2 Data :", array_2) // [1 2 3 4 5]
//     fmt.Println("array_3 Data :", array_3) // [1 2 3 4 5]
//     fmt.Println("array_1 Length :", len(array_1)) // 5
//     fmt.Println("array_2 Length :", len(array_2)) // 5
//     fmt.Println("array_3 Length :", len(array_3)) // 5
// }

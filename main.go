package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// _ "github.com/go-sql-driver/mysql"
)

// var MySQLdb *sql.DB

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NEXTBEST.SI APIs by SNAILWALK.CO (info@snailwalk.co)\r")
	fmt.Fprintf(w, "============================================================\r")
	fmt.Fprintf(w, "/                                     [GET] ==> Here\r")
	fmt.Fprintf(w, "/Coupons                              [GET] ==> List all coupons\r")
	fmt.Fprintf(w, "/Add/{id}/{grp}/{active}              [POST] ==> Add new coupon\r")
	fmt.Fprintf(w, "/Search/{id}                          [GET] ==> Search coupon\r")
	fmt.Fprintf(w, "/Use/{id}/{terminalid}/{sku}/{ref}    [POST] ==> Use coupon\r")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")

	// myRouter.HandleFunc("/gigest", AllCoupons).Methods("GET")
	//	myRouter.HandleFunc("/Search/{id}", GetCoupon).Methods("GET")
	myRouter.HandleFunc("/GenerateDigest", GenerateDigest).Methods("POST")
	// myRouter.HandleFunc("/Use/Staff/{id}/{terminalid}/{sku}/{ref}", UseCouponStaff).Methods("POST")
	// myRouter.HandleFunc("/Use/Customer/{id}/{terminalid}/{sku}/{ref}", UseCouponCustomer).Methods("POST")
	// myRouter.HandleFunc("/Use/Staff/{id}/{terminalid}/{sku}/{ref}/{amount}", UseCouponStaffMySQL).Methods("POST")
	// myRouter.HandleFunc("/Use/Customer/{id}/{terminalid}/{sku}/{ref}/{amount}", UseCouponCustomerMySQL).Methods("POST")
	// myRouter.HandleFunc("/Report/{year}/{month}/{day}/{lcid}", DailyReport).Methods("GET")
	// myRouter.HandleFunc("/Sendmail/{year}/{month}/{day/{lcid}}", SendMail).Methods("GET")
	// myRouter.HandleFunc("/Sendmail/{lcid}", SendMail_1).Methods("GET")

	// myRouter.HandleFunc("/Report/{year}/{month}", MonthlyReport).Methods("GET")

	// Admin only
	// myRouter.HandleFunc("/Add/{id}/{grp}/{active}", NewCoupon).Methods("POST")

	log.Fatal(http.ListenAndServe(":7767", myRouter))

}

func main() {

	fmt.Println("Connecting to NEXTBEST.AI APIs by CIAOKIT.COM")

	// MySQLdb, err := sql.Open("mysql", "vmc:vmc$n@il@tcp(aisecoupons.snailwalk.co:3306)/AISDB")

	// if err != nil {

	// 	panic(err.Error())

	// } else {

	// 	fmt.Println("MySQL Database CONNECTED !!!")
	// }

	// defer MySQLdb.Close()

	fmt.Println("============================================")
	fmt.Println(" NEXTBEST.AI FEE CALCULATOR Service STARTED")
	fmt.Println("============================================")
	//	fmt.Println(strconv.Itoa(int(time.Now().Unix())))

	//	InitialMigration()

	handleRequests()
}

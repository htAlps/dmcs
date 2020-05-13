//  ._______.___________________.___________________.___________________.___________________._______;
//  JWT 
    package jwt

    import ( "os"; "fmt"; "encoding/json"; "net/http"; "io/ioutil"; )

    type JRecordT   struct {
        Name, Addr, Email, Phone string
    }


//  ._______.___________________.___________________.___________________.___________________._______;
//  UTILITY FNCTIONS 


    func checkErr(ee error, desc string)) {
        if ee != nil {
            fmt.Printf("Error: %c - %s\n", ee, desc)
            panic (ee)
        }
    }

//  ._______.___________________.___________________.___________________.___________________._______;
//  CLIENT FNCTIONS 

    func GetValFromKey(url, key string) (val string) {

        res, err := http.Get(url)
        checkErr (err)
        


        mux.Main_svc()
    }

//  ._______.___________________.___________________.___________________.___________________._______;
//  SERVER FNCTIONS 



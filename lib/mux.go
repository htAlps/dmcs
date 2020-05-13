//  ._______.___________________.___________________.___________________.___________________.___________________.___________________.___________________.                   .                   .                   .                   .                   ;
//  svc.go  ▻20‹05‹07τ10›25›21›     
    package lib

    import ("os"; "io/ioutil"; "time"; "fmt"; "net/http";  mux "github.com/julienschmidt/httprouter"; )

    const msg1      = `Hello World`
    const url0      = `http://localhost:8080/key/`
    const url1      = `http://localhost:8081/key/`
    const url2      = `http://localhost:8082/key/`
    const z_string  = ``
    const TTL_7sec  = 7   * time.Second
    const path_pk1  = `/usr/local/sys/bin/key/pk1.crt`
    const path_sk1  = `/usr/local/sys/bin/key/sk1.key`


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  Basic http handlers

    func handle_index(ww http.ResponseWriter, rr *http.Request, _ mux.Params) {
        fmt.Fprint(ww, `Welcome... use:iam to tell me your name, e.g., /iam/John `)
    }

    func handle_key0(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := cli_get(url1, params.ByName("key") + "B")
        if err != nil {
            fmt.Fprintf(ww, "From svc0: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc0: [ %s ] ", body)
    }

    func handle_key1(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := cli_get(url2, params.ByName("key") + "C")
        if err != nil {
            fmt.Fprintf(ww, "From svc1: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc1: [ %s ]", body)
    }

    func handle_key2(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        fmt.Fprintf(ww, "from svc2: [ %s ]", params.ByName("key"))
    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  cli_get sends a key and gets a corresponding value
    func cli_get(url, key string) (string, error) {

        req, err := http.NewRequest(`GET`, url + key, nil)
        if err != nil {
            fmt.Printf("\n%v \nExiting ... \n", err.Error())
            return z_string, err
        }
        cli := &http.Client { Timeout: TTL_7sec, }

        resp, err := cli.Do(req)
        if err != nil {
            fmt.Printf("\nERROR: %v \nRESP: %v \nExiting ... \n", err.Error(), resp)
            return z_string, err
        }
        defer resp.Body.Close()
        fmt.Printf("\nRESP: \n%v\n\n", resp)

        body, err := ioutil.ReadAll(resp.Body)
        return string(body), err
    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  MAIN http clients

    func Cli0_main(key string) {

        body, err := cli_get(url0, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli1_main(key string) {

        body, err := cli_get(url1, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli2_main(key string) {

        body, err := cli_get(url2, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  MAIN http server 
    func Svc0_main() {

        rr := mux.New()
        rr.GET("/",         handle_index)
        rr.GET("/key/:key", handle_key0)

        http.ListenAndServe(":8080", rr)
    }


    func Svc1_main() {

        rr := mux.New()
        rr.GET("/key/:key", handle_key1)

        http.ListenAndServe(":8081", rr)
    }


    func Svc2_main() {

        rr := mux.New()
        rr.GET("/key/:key", handle_key2)

        http.ListenAndServe(":8082", rr)
    }


/*  CODE PIT
    const url       = `https://localhost:8080/hello/`
        http.ListenAndServeTLS("localhost:8080", path_pk1, path_sk1, rr)
        //  fmt.Printf("\n%s\n\n", msg1)
        
*/

//  ._______.___________________.___________________.___________________.___________________.___________________.___________________.___________________.                   .                   .                   .                   .                   ;
//  mux.go  ▻20‹05‹07τ10›25›21›     
    package lib

    import ("os"; "io/ioutil"; "time"; "fmt"; "net/http";  mux "github.com/julienschmidt/httprouter"; )

    const msg1          = `Hello World`
    const url1          = `http://localhost:8071/key/`
    const url2          = `http://localhost:8072/key/`
    const url3          = `http://localhost:8073/key/`
    const rest1_port    = `:8071`
    const rest2_port    = `:8072`
    const rest3_port    = `:8073`
    const z_string      = ``
    const TTL_7sec      = 7 * time.Second
    const path_pk1      = `/usr/local/sys/bin/key/pk1.crt`
    const path_sk1      = `/usr/local/sys/bin/key/sk1.key`


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  Basic http handlers

    func handle_index(ww http.ResponseWriter, rr *http.Request, _ mux.Params) {
        fmt.Fprint(ww, `Welcome... use:iam to tell me your name, e.g., /iam/John `)
    }

    func handle_key1(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := cli_get(url2, params.ByName("key") + "B")
        if err != nil {
            fmt.Fprintf(ww, "From svc1: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc1: [ %s ] ", body)
    }

    func handle_key2(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := cli_get(url3, params.ByName("key") + "C")
        if err != nil {
            fmt.Fprintf(ww, "From svc1: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc2: [ %s ]", body)
    }

    func handle_key3(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        fmt.Fprintf(ww, "from svc3: [ %s ]", params.ByName("key"))
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
//  main http1.1 clients

    func Cli1_rest_main(key string) {

        body, err := cli_get(url1, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli2_rest_main(key string) {

        body, err := cli_get(url2, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli3_rest_main(key string) {

        body, err := cli_get(url3, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  main http servers
    func Svc1_rest_main() {

        rr := mux.New()
        rr.GET("/",         handle_index)
        rr.GET("/key/:key", handle_key1)

        http.ListenAndServe(rest1_port, rr)
    }


    func Svc2_rest_main() {

        rr := mux.New()
        rr.GET("/key/:key", handle_key2)

        http.ListenAndServe(rest2_port, rr)
    }


    func Svc3_rest_main() {

        rr := mux.New()
        rr.GET("/key/:key", handle_key3)

        http.ListenAndServe(rest3_port, rr)
    }


/*  CODE PIT
    const url       = `https://localhost:8080/hello/`
        http.ListenAndServeTLS("localhost:8080", path_pk1, path_sk1, rr)
        //  fmt.Printf("\n%s\n\n", msg1)
        
*/

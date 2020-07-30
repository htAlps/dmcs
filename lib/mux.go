//  ._______.___________________.___________________.___________________.___________________.___________________.___________________.___________________.                   .                   .                   .                   .                   ;
//  mux.go  ▻20‹07‹28τ23›56›05› ▻20‹05‹07τ10›25›21›
    package lib

    import ( "os"; "io/ioutil"; "time"; "fmt"; "net/http";  mux "github.com/julienschmidt/httprouter"; )
    import ( "github.com/hashicorp/consul/api"; "github.com/hashicorp/consul/connect"; )

    const msg1          = `Hello World`
    // const http_url      = `http://localhost:8071/key/`
    const http_url      = `http://localhost`
    const http_url1     = `http://localhost:8071/key/`
    const http_url2     = `http://localhost:8072/key/`
    const http_url3     = `http://localhost:8073/key/`
    const http_url4     = `http://localhost:8074/key/`
    const http_url5     = `http://localhost:8075/key/`
    const http_url6     = `http://localhost:8076/key/`
    const http_url7     = `http://localhost:8077/key/`
    const http_url8     = `http://localhost:8078/key/`
    const http_url9     = `http://localhost:8079/key/`

    const http_pt1      = `:8071`
    const http_pt2      = `:8072`
    const http_pt3      = `:8073`
    const http_pt4      = `:8074`
    const http_pt5      = `:8075`
    const http_pt6      = `:8076`
    const http_pt7      = `:8077`
    const http_pt8      = `:8078`
    const http_pt9      = `:8079`

    const rest_key      = `key`

    const tls_url7      = `https://svc7native.service.consul/`
    const tls_url8      = `https://svc8native.service.consul/`
    const tls_url9      = `https://svc9native.service.consul/`

    const z_string      = ``
    const TTL_7sec      = 7 * time.Second
    const path_pk1      = `/usr/local/sys/bin/key/pk1.crt`
    const path_sk1      = `/usr/local/sys/bin/key/sk1.key`

    //    resp, err := cli.Get(`https://svc7native.service.consul/key/AAAA`)  // dial svc7native service. Must use Consul DNS.

//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  Basic http handlers

    func handle_index(ww http.ResponseWriter, rr *http.Request, _ mux.Params) {
        fmt.Fprint(ww, `Welcome... use:iam to tell me your name, e.g., /iam/John `)
    }


    func handle1(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := http_get(http_url + http_pt2 + rest_key + `/`, params.ByName(rest_key) + "B")
        // body, err := http_get(http_url2, params.ByName(rest_key) + "B")
        if err != nil {
            fmt.Fprintf(ww, "From svc1: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc1: [ %s ] ", body)
    }


    func handle2(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := http_get(http_url + http_pt2 + rest_key + `/`, params.ByName(rest_key) + "C")
        if err != nil {
            fmt.Fprintf(ww, "From svc2: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc2: [ %s ]", body)
    }


    func handle3(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        val := params.ByName(rest_key) + "D"
        fmt.Fprintf(ww, "From svc3: [ %s ]", val)
    }


    func handle4(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := http_get(http_url + http_pt5 + rest_key + `/`, params.ByName(rest_key) + "B")
        if err != nil {
            fmt.Fprintf(ww, "From svc4: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc4: [ %s ] ", body)
    }


    func handle5(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        body, err := http_get(http_url + http_pt6 + rest_key + `/`, params.ByName(rest_key) + "C")
        if err != nil {
            fmt.Fprintf(ww, "From svc5: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc5: [ %s ]", body)
    }


    func handle6(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        val := params.ByName(rest_key) + "D"
        fmt.Fprintf(ww, "From svc6: [ %s ]", val)
    }


    func handle7(ww http.ResponseWriter, rr *http.Request, params mux.Params) {

        body, err := tls_get(`svc7native`, `svc8native`, params.ByName(rest_key) + "B")
        if err != nil {
            fmt.Fprintf(ww, "From svc7: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc7: [ %s ] ", body)
    }


    func handle8(ww http.ResponseWriter, rr *http.Request, params mux.Params) {

        body, err := tls_get(`svc8native`, `svc9native`, params.ByName(rest_key) + "C")
        if err != nil {
            fmt.Fprintf(ww, "From svc8: [ ERROR ] ")
        }
        fmt.Fprintf(ww, "From svc8: [ %s ] ", body)
    }


    func handle9(ww http.ResponseWriter, rr *http.Request, params mux.Params) {
        val := params.ByName(rest_key) + "D"
        fmt.Fprintf(ww, "From svc9: [ %s ]", val)
    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  tls_get sends a key and gets a corresponding value via TLS
    func tls_get(source_svc, dest_svc, key string) (string, error) {


        cApiCli, _ := api.NewClient(api.DefaultConfig())                // cApiCli is a pointer to the Consul API client

        thisSvc, _ := connect.NewService(source_svc, cApiCli)           // thisSvc is a pointer to the Consul Connect agent service
        defer thisSvc.Close()

        cli := thisSvc.HTTPClient()                                     // instantiate an HTTP client via Consul Connect

        dest_url := `https://` + dest_svc + `.service.consul/key/` + key

        resp, err := cli.Get(dest_url)                                  // dial target URL. Must use Consul DNS.
        if err != nil {
            return z_string, err
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return z_string, err
        }
        return string(body), err
    }
    // resp, err := cli.Get(`https://svc8native.service.consul/key/AAAA`)
    //  urlStr := fmt.Sprintf("%v", u)


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  http_get sends a key and gets a corresponding value
    func http_get(url, key string) (string, error) {

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
//  main http1.1 Integration Tests (ITest)

    // ITest1_rest_dmcs: Integration Test #1, REST, Main Entry Point.
    func ITest1_rest_dmcs(api_key string) {

        body, err := http_get(http_url1, api_key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    // ITest2_rest_dmcs: Integration Test #2, REST, Main Entry Point.
    func ITest2_rest_dmcs(api_key string) {

        body, err := http_get(http_url2, api_key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    // ITest3_rest_dmcs: Integration Test #3, REST, Main Entry Point.
    func ITest3_rest_dmcs(api_key string) {

        body, err := http_get(http_url3, api_key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


    // ITest4_rest_dmcs: Integration Test #4, Consul Connect, SideCar, REST, Main Entry Point.
    func ITest4_cConn_sidecar_rest_dmcs(api_key string) {

        body, err := http_get(http_url4, api_key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


    // ITest5_rest_dmcs: Integration Test #5, Consul Connect, SideCar, REST, Main Entry Point.
    func ITest5_cConn_sidecar_rest_dmcs(api_key string) {

        body, err := http_get(http_url5, api_key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


    // ITest6_rest_dmcs: Integration Test #6, Consul Connect, SideCar, REST, Main Entry Point.
    func ITest6_cConn_sidecar_rest_dmcs(api_key string) {

        body, err := http_get(http_url6, api_key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


    // ITest7_cConn_native_rest_dmcs: Integration Test #7, Consul Connect, Golang Native, REST, Main Entry Point.
    // this function interfaces with the Consul and Connect APIs to get mTLS keys so that it can access svc7native API
    func ITest7_cConn_native_rest_dmcs(api_key string) {

        body, err := tls_get(`itest7native`, `svc7native`, api_key)
        if err != nil {
            fmt.Printf("From ITest7: ERROR: %s \n", err.Error())
        }
        fmt.Printf("\nBODY: \n%s\n\n", body)
    }


    // ITest8_cConn_native_rest_dmcs: Integration Test #8, Consul Connect, Golang Native, REST, Main Entry Point.
    // this function interfaces with the Consul and Connect APIs to get mTLS keys so that it can access svc8native API
    func ITest8_cConn_native_rest_dmcs(api_key string) {

        body, err := tls_get(`itest8native`, `svc8native`, api_key)
        if err != nil {
            fmt.Printf("From ITest8: ERROR: %s \n", err.Error())
        }
        fmt.Printf("\nBODY: \n%s\n\n", body)
    }
    //  fmt.Printf("\nERROR: %v \nRESP: %v \nExiting ... \n", err.Error(), resp)


    // ITest9_cConn_native_rest_dmcs: Integration Test #9, Consul Connect, Golang Native, REST, Main Entry Point.
    // this function interfaces with the Consul and Connect APIs to get mTLS keys so that it can access svc9native API
    func ITest9_cConn_native_rest_dmcs(api_key string) {

        body, err := tls_get(`itest9native`, `svc9native`, api_key)
        if err != nil {
            fmt.Printf("From ITest9: ERROR: %s \n", err.Error())
        }
        fmt.Printf("\nBODY: \n%s\n\n", body)
    }
    // fmt.Printf("\nERROR: %v \nRESP: %v \nExiting ... \n", err.Error(), resp)


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  main http servers

    // Svc1_rest means: web service#1, REST, Main Entry Point.
    func Svc1_rest_dmcs() {

        rr := mux.New()
        rr.GET("/",         handle_index)
        rr.GET("/key/:key", handle1)

        http.ListenAndServe(http_pt1, rr)
    }


    // Svc2_rest means: web service#2, REST, Main Entry Point.
    func Svc2_rest_dmcs() {

        rr := mux.New()
        rr.GET("/key/:key", handle2)

        http.ListenAndServe(http_pt2, rr)
    }


    // Svc3_rest means: web service#3, REST, Main Entry Point.
    func Svc3_rest_dmcs() {

        rr := mux.New()
        rr.GET("/key/:key", handle3)

        http.ListenAndServe(http_pt3, rr)
    }


    // Svc4_cConn_sidecar_rest_dmcs means: web service #4, Consul Connect, SideCar, REST, Main Entry Point.
    func Svc4_cConn_sidecar_rest_dmcs() {

        rr := mux.New()
        rr.GET("/key/:key", handle4)

        http.ListenAndServe(http_pt4, rr)
    }


    // Svc5_cConn_sidecar_rest_dmcs means: web service #5, Consul Connect, SideCar, REST, Main Entry Point.
    func Svc5_cConn_sidecar_rest_dmcs() {

        rr := mux.New()
        rr.GET("/key/:key", handle5)

        http.ListenAndServe(http_pt5, rr)
    }


    // Svc6_cConn_sidecar_rest_dmcs means: web service #6, Consul Connect, SideCar, REST, Main Entry Point.
    func Svc6_cConn_sidecar_rest_dmcs() {

        rr := mux.New()
        rr.GET("/key/:key", handle6)

        http.ListenAndServe(http_pt6, rr)
    }


    // Svc7_cConn_native_rest means: web service#7, Consul Connect, Golang Native, REST, Main Entry Point.
    // this function interfaces with the Consul and Connect APIs to get mTLS keys so that only authorized services can access it
    func Svc7_cConn_native_rest_dmcs() {

        cApiCli, _ := api.NewClient(api.DefaultConfig())            // instantiate a Consul API client

        thisSvc, _ := connect.NewService("svc7native", cApiCli)     // instantiate a Consul-Connect HTTPS server
        defer thisSvc.Close()                                       // clean up an close thisSvc before exiting this function

        rr := mux.New()                                             // instantiate a new http-router
        rr.GET("/key/:key", handle7)                                // with a simple REST key, handled by function: handle7

        server := &http.Server{                                     // instantiate an HTTPS Server via Consul Connect
            Addr:      http_pt7,
            Handler:   rr,
            TLSConfig: thisSvc.ServerTLSConfig(),                   // here we get specific the certs for this server for mTLS
        }
        server.ListenAndServeTLS("", "")                            // Listen and Serve.
    }


    // Svc8_cConn_native_rest means: web service#8, Consul Connect, Golang Native, REST, Main Entry Point.
    // this function interfaces with the Consul and Connect APIs to get mTLS keys so that only authorized services can access it
    func Svc8_cConn_native_rest_dmcs() {

        cApiCli, _ := api.NewClient(api.DefaultConfig())            // instantiate a Consul API client

        thisSvc, _ := connect.NewService("svc8native", cApiCli)     // instantiate a Consul-Connect HTTPS server
        defer thisSvc.Close()                                       // clean up an close thisSvc before exiting this function

        rr := mux.New()                                             // instantiate a new http-router
        rr.GET("/key/:key", handle8)                                // with a simple REST key, handled by function: handle8

        server := &http.Server{                                     // instantiate an HTTPS Server via Consul Connect
            Addr:      http_pt8,
            Handler:   rr,
            TLSConfig: thisSvc.ServerTLSConfig(),                   // here we get specific the certs for this server for mTLS
        }
        server.ListenAndServeTLS("", "")                            // Listen and Serve.
    }


    // Svc9_cConn_native_rest means: web service#9, Consul Connect, Golang Native, REST, Main Entry Point.
    // this function interfaces with the Consul and Connect APIs to get mTLS keys so that only authorized services can access it
    func Svc9_cConn_native_rest_dmcs() {

        cApiCli, _ := api.NewClient(api.DefaultConfig())            // instantiate a Consul API client

        thisSvc, _ := connect.NewService("svc9native", cApiCli)     // instantiate a Consul-Connect HTTPS server
        defer thisSvc.Close()                                       // clean up an close thisSvc before exiting this function

        rr := mux.New()                                             // instantiate a new http-router
        rr.GET("/key/:key", handle9)                                // with a simple REST key, handled by function: handle9

        server := &http.Server{                                     // instantiate an HTTPS Server via Consul Connect
            Addr:      http_pt9,
            Handler:   rr,
            TLSConfig: thisSvc.ServerTLSConfig(),                   // here we get specific the certs for this server for mTLS
        }
        server.ListenAndServeTLS("", "")                            // Listen and Serve.
    }


/*  CODE PIT
    const url       = `https://localhost:8080/hello/`
        http.ListenAndServeTLS("localhost:8080", path_pk1, path_sk1, rr)
        //  fmt.Printf("\n%s\n\n", msg1)

*/

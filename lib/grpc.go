//  ._______.___________________.___________________.___________________.___________________.___________________.___________________.___________________.                   .                   .                   .                   .                   ;
//  grpc.go    ▻20‹06‹06τ19›05›23›
    package lib

    import ("os"; "fmt"; "log"; "net"; "net/http"; "google.golang.org/grpc";  mux "github.com/julienschmidt/httprouter"; )

    const tcp       = `tcp`
    const grpc_port = `:9090`

//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  main programs
    func Hello_main() {
        var ii int = 32
        jj := 33;
        xx := 43.2345

        fmt.Println("Hello World!", ii, jj, xx)
        fmt.Println(jj-ii)

        fmt.Println()
    }

//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  main http2.0 clients

    func Cli0_grpc_main(key string) {

        body, err := cli_get(url0, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli1_grpc_main(key string) {

        body, err := cli_get(url1, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli2_grpc_main(key string) {

        body, err := cli_get(url2, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  main http servers
    func Svc0_grpc_main() {

        conn, err := net.Listen(tcp, grpc_port)
        if err != nil {
            log.Fatalf("Failed to Listen on port %s - %v", grpc_port, err)
        }

        grpcSvc := grpc.NewServer()

        if err := grpcSvc.Serve(conn); err != nil {
            log.Fatalf("Failed to Serve on port %s - %v", grpc_port, err)
        }

    }


    func Svc1_grpc_main() {

        rr := mux.New()
        rr.GET("/key/:key", handle_key1)

        http.ListenAndServe(":8081", rr)
    }


    func Svc2_grpc_main() {

        rr := mux.New()
        rr.GET("/key/:key", handle_key2)

        http.ListenAndServe(":8082", rr)
    }




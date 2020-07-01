//  ._______.___________________.___________________.___________________.___________________.___________________.___________________.___________________.                   .                   .                   .                   .                   ;
//  grpc.go    ▻20‹06‹06τ19›05›23›
    package lib

    import ("os"; "fmt"; "log"; "net"; "google.golang.org/grpc"; "context"; )

    const tcp           = `tcp`
    const grpc1_port    = `:9091`
    const grpc2_port    = `:9092`
    const grpc3_port    = `:9093`

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

    func Cli1_grpc_main(key string) {

        body, err := cli_get(url3, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli2_grpc_main(key string) {

        body, err := cli_get(url1, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }

    func Cli3_grpc_main(key string) {

        body, err := cli_get(url2, key)
        if err != nil {
            fmt.Printf("\nERROR: %v \nExiting ... \n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("\nRESP:\n%v\n\n", body)
    }


    type Svc struct {
    }

    func (ss *Svc) GetValue(ctx context.Context, msg *Msg) (*Msg, error) {
        log.Printf("Got message from client: %s", msg.Body)
        return &Msg{Body: "Value 111"}, nil
    }

//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  main http servers
    func Svc1_grpc_main() {

        conn, err := net.Listen(tcp, grpc1_port)
        if err != nil {
            log.Fatalf("Listen Failed - Port: %s - Error: %v", grpc1_port, err)
        }

        grpc1svc := grpc.NewServer()

        if err := grpc1svc.Serve(conn); err != nil {
            log.Fatalf("Server Failed - Port: %s - Error: %v", grpc1_port, err)
        }
    }


    func Svc2_grpc_main() {

        conn, err := net.Listen(tcp, grpc1_port)
        if err != nil {
            log.Fatalf("Listen Failed - Port: %s - Error: %v", grpc2_port, err)
        }

        grpc2svc := grpc.NewServer()

        if err := grpc2svc.Serve(conn); err != nil {
            log.Fatalf("Server Failed - Port: %s - Error: %v", grpc2_port, err)
        }
    }


    func Svc3_grpc_main() {

        conn, err := net.Listen(tcp, grpc3_port)
        if err != nil {
            log.Fatalf("Listen Failed - Port: %s - Error: %v", grpc3_port, err)
        }

        grpc3svc := grpc.NewServer()

        if err := grpc3svc.Serve(conn); err != nil {
            log.Fatalf("Server Failed - Port: %s - Error: %v", grpc3_port, err)
        }
    }


    func Hello_main2() {
        var ii int = 32
        jj := 33;
        xx := 43.2345

        fmt.Println("Hello World!", ii, jj, xx)
        fmt.Println(jj-ii)

        fmt.Println()
    }




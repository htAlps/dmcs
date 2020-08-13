//  svc9native.hcl
service {
    name = "svc9native",
    tags = ["svc"],
    port = 8079,
    connect {
        native = true,
    }
}

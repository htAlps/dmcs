//  svc7native.hcl
service {
    name = "svc7native",
    tags = ["svc"],
    port = 8077,
    connect {
        native = true,
    }
}

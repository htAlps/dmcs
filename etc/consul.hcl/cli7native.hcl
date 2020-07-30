//  cli7native.hcl
service {
    name = "cli7native",
    tags = ["cli"],
    port = 8057,
    connect {
        native = true,
    }
}

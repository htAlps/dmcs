//  cli9native.hcl
service {
    name = "cli9native",
    tags = ["cli"],
    port = 8059,
    connect {
        native = true,
    }
}

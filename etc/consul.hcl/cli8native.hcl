//  cli8native.hcl
service {
    name = "cli8native",
    tags = ["cli"],
    port = 8058,
    connect {
        native = true,
    }
}

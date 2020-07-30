//  svc3rest.hcl
service = {
    check = {
        args = [
            "curl",
            "localhost:8073/key/TEST3->",
        ],
        interval = "10s",
    },
    name = "svc3rest",
    tags = ["svc"],
    port = 8073,
}

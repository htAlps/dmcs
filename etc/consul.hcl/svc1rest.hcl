//  svc1rest.hcl
service = {
    check = {
        args = [
            "curl",
            "localhost:8071/key/TEST1->",
        ],
        interval = "10s",
    },
    name = "svc1rest",
    tags = ["svc"],
    port = 8071,
}

//  svc2rest.hcl
service = {
    check = {
        args = [
            "curl",
            "localhost:8072/key/TEST2->",
        ],
        interval = "10s",
    },
    name = "svc2rest",
    tags = ["svc"],
    port = 8072,
}

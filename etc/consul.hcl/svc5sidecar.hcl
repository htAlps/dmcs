//  svc5sidecar.hcl
service = {
    name = "svc5sidecar", 
    tags = ["svc"], 
    port = 8075,
    connect = {
        sidecar_service = {}
    }
} 

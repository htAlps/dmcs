//  svc6sidecar.hcl
service = {
    name = "svc6sidecar", 
    tags = ["svc"], 
    port = 8076,
    connect = {
        sidecar_service = {}
    }
} 

//  cli6sidecar.hcl
service {
    name = "cli6sidecar", 
    tags = ["cli"], 
    port = 8056,
    connect = {
        sidecar_service {}
    }
} 

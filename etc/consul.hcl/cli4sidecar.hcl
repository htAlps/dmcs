//  cli4sidecar.hcl
service {
    name = "cli4sidecar", 
    tags = ["cli"], 
    port = 8054,
    connect = {
        sidecar_service {}
    }
} 

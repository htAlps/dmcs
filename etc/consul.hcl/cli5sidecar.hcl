//  cli5sidecar.hcl
service {
    name = "cli5sidecar", 
    tags = ["cli"], 
    port = 8055,
    connect = {
        sidecar_service {}
    }
} 

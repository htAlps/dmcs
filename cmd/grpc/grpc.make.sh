go install $sysroot/src/dmcs/cmd/grpc/intest1.grpc.dmcs.go
go install $sysroot/src/dmcs/cmd/grpc/intest_grpc.dmcs.go
go install $sysroot/src/dmcs/cmd/grpc/intest_grpc.dmcs.go

go install $sysroot/src/dmcs/cmd/grpc/svc1_grpc.dmcs.go
go install $sysroot/src/dmcs/cmd/grpc/svc2_grpc.dmcs.go
go install $sysroot/src/dmcs/cmd/grpc/svc3_grpc.dmcs.go

go install $sysroot/src/dmcs/cmd/grpc/hello.go

mv  intest_grpc.dmcs.go     intest1.grpc.dmcs.go
mv  intest_grpc.dmcs.go     intest.grpc.dmcs.go
mv  intest_grpc.dmcs.go     intest.grpc.dmcs.go
                                                
mv  svc1_grpc.dmcs.go       svc1.grpc.dmcs.go
mv  svc2_grpc.dmcs.go       svc2.grpc.dmcs.go
mv  svc3_grpc.dmcs.go       svc3.grpc.dmcs.go




-rwx------  1 sys7  wheel  335 Jul 15 07:16 cli1.rest.dmcs.go
-rwx------  1 sys7  wheel  335 Jul 15 07:57 cli2.rest.dmcs.go
-rwx------  1 sys7  wheel  335 Jul 15 07:57 cli3.rest.dmcs.go
-rw-r--r--  1 sys7  wheel  383 Jul 15 07:57 cli4.cConn_sidecar.rest.dmcs.go     
-rw-r--r--  1 sys7  wheel  378 Jul 15 07:57 cli7.cConn_native.rest.dmcs.go
-rw-r--r--  1 sys7  wheel  336 Jul 15 07:18 cliA.mongo.rest.dmcs.go

mv  cli1.rest.dmcs.go                   intest1.rest.dmcs.go
mv  cli2.rest.dmcs.go                   intest2.rest.dmcs.go
mv  cli3.rest.dmcs.go                   intest3.rest.dmcs.go

mv  cli4.cConn_sidecar.rest.dmcs.go     intest4.cConn_sidecar.rest.dmcs.go     
mv  cli7.cConn_native.rest.dmcs.go      intest7.cConn_native.rest.dmcs.go
mv  cliA.mongo.rest.dmcs.go             intestA.mongo.rest.dmcs.go

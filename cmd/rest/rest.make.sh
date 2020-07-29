#!/bin/zsh
p1=$( printf '%d' $1 )
echo                           

installAll() {
    echo Installing All
    installRest1-3
    installSideCar4-6
    installNative7-9
    installMongoA-C
}

installRest1-3() {
    echo Installing REST 1-3
    go install $sysroot/src/dmcs/cmd/rest/intest1.rest.dmcs.go
go install $sysroot/src/dmcs/cmd/rest/intest2.rest.dmcs.go
go install $sysroot/src/dmcs/cmd/rest/intest3.rest.dmcs.go

go install $sysroot/src/dmcs/cmd/rest/svc1.rest.dmcs.go
go install $sysroot/src/dmcs/cmd/rest/svc2.rest.dmcs.go
go install $sysroot/src/dmcs/cmd/rest/svc3.rest.dmcs.go
}

installSideCar4-6() {
    echo $p1 - Installing SideCar 4-6
go install $sysroot/src/dmcs/cmd/rest/intest4.cConn_sidecar.rest.dmcs.go
go install $sysroot/src/dmcs/cmd/rest/svc4.cConn_sidecar.rest.dmcs.go
}

installNative7-9() {
    echo $p1 - Installing Native 7-9
    go install $sysroot/src/dmcs/cmd/rest/intest7.cConn_native.rest.dmcs.go
    go install $sysroot/src/dmcs/cmd/rest/svc7.cConn_native.rest.dmcs.go
}

installMongoA-C() {
    echo $p1 - Installing Mongo A-C
    go install $sysroot/src/dmcs/cmd/rest/intestA.mongo.rest.dmcs.go
    go install $sysroot/src/dmcs/cmd/rest/svcA.mongo.rest.dmcs.go
}

installSideCar4() {
    echo $p1 - Installing SideCar 4
    go install $sysroot/src/dmcs/cmd/rest/intest4.cConn_sidecar.rest.dmcs.go
    go install $sysroot/src/dmcs/cmd/rest/svc4.cConn_sidecar.rest.dmcs.go
}

installNative7() {
    echo $p1 - Installing Native 7
    go install $sysroot/src/dmcs/cmd/rest/intest7.cConn_native.rest.dmcs.go
    go install $sysroot/src/dmcs/cmd/rest/svc7.cConn_native.rest.dmcs.go
}

installMongoA() {
    echo $p1 - Installing Native 7
    go install $sysroot/src/dmcs/cmd/rest/intestA.mongo.rest.dmcs.go
    go install $sysroot/src/dmcs/cmd/rest/svcA.mongo.rest.dmcs.go
}

intestA.mongo.rest.dmcs.go  
svcA.mongo.rest.dmcs.go

other() {
    echo $p1 is Not an Option, USAGE:
    echo "1:   Install REST    1-3 "
    echo "2:   Install SideCar 4-6 "
    echo "3:   Install Native  7-9 "
    echo "4:   Install Mongo   A-C "
    echo "9:   Install ALL         "
    echo "104: Install SideCar 4   "
    echo "107: Install Native  7   "
    echo "110: Install Mongo   A   "
}


if [ $1 -eq 0 ]; then
    other;

elif [ $1 -eq 1 ]; then
    installRest1-3;

elif [ $1 -eq 2 ]; then
    installSideCar4-6;

elif [ $1 -eq 3 ]; then
    installNative7-9;

elif [ $1 -eq 4 ]; then
    installMongoA-C

elif [ $1 -eq 9 ]; then
    installAll;

elif [ $1 -eq 104 ]; then
    installSideCar4;

elif [ $1 -eq 107 ]; then
    installNative7;

else
    other;
fi


# 🀄️🀄️🀄️🀄️🀄️··🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️··🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️🀄️··🀄️🀄️🀄️🀄️🀄️·
# λ,REST.MAKE.SH                                    ▻20‹07‹31τ13›51›10› 
# Installs some or all the services in a node of DMCS
p1=$( printf '%d' $1 )
echo                           

installAll() {
    echo $p1 - Installing All
    installRest;
    installSideCar;
    installNative;
    installMongo;
}
installRest() {
    echo $p1 - Installing REST
    installRest101;
    installRest102;
    installRest103;
}
installSideCar() {
    echo $p1 - Installing SideCar
    installSideCar104;
    installSideCar105;
    installSideCar106;
}
installNative() {
    echo $p1 - Installing Native
    installNative107;
    installNative108;
    installNative109;
}
installMongo() {
    echo $p1 - Installing Mongo
    installMongo110;
}

installRest101() {
    echo $p1 - Installing REST 101
    go install $sysroot/src/dmcs/cmd/rest/itest1.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc1.rest.dmcs.go;
}
installRest102() {
    echo $p1 - Installing REST 102
    go install $sysroot/src/dmcs/cmd/rest/itest2.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc2.rest.dmcs.go;
}
installRest103() {
    echo $p1 - Installing REST 103
    go install $sysroot/src/dmcs/cmd/rest/itest3.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc3.rest.dmcs.go;
}

installSideCar104() {
    echo $p1 - Installing SideCar 104
    go install $sysroot/src/dmcs/cmd/rest/itest4.cConn_sidecar.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc4.cConn_sidecar.rest.dmcs.go;
}
installSideCar105() {
    echo $p1 - Installing SideCar 105
    go install $sysroot/src/dmcs/cmd/rest/itest5.cConn_sidecar.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc5.cConn_sidecar.rest.dmcs.go;
}
installSideCar106() {
    echo $p1 - Installing SideCar 106
    go install $sysroot/src/dmcs/cmd/rest/itest6.cConn_sidecar.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc6.cConn_sidecar.rest.dmcs.go;
}

installNative107() {
    echo $p1 - Installing Native 107
    go install $sysroot/src/dmcs/cmd/rest/itest7.cConn_native.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc7.cConn_native.rest.dmcs.go;
}
installNative108() {
    echo $p1 - Installing Native 108
    go install $sysroot/src/dmcs/cmd/rest/itest8.cConn_native.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc8.cConn_native.rest.dmcs.go;
}
installNative109() {
    echo $p1 - Installing Native 109
    go install $sysroot/src/dmcs/cmd/rest/itest9.cConn_native.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svc9.cConn_native.rest.dmcs.go;
}

installMongo110() {
    echo $p1 - Installing Mongo 110
    go install $sysroot/src/dmcs/cmd/rest/itestA.mongo.rest.dmcs.go;
    go install $sysroot/src/dmcs/cmd/rest/svcA.mongo.rest.dmcs.go;
}

other() {
    echo $p1 is Not an Option, USAGE:
    echo "1:    Install REST        "
    echo "2:    Install SideCar     "
    echo "3:    Install Native      "
    echo "4:    Install Mongo       "
    echo "9:    Install ALL         "
    echo "101:  Install Rest-1      "
    echo "102:  Install Rest-2      "
    echo "103:  Install Rest-3      "
    echo "104:  Install SideCar-4   "
    echo "105:  Install SideCar-5   "
    echo "106:  Install SideCar-6   "
    echo "107:  Install Native-7    "
    echo "108:  Install Native-8    "
    echo "109:  Install Native-9    "
    echo "110:  Install Mongo-A     "
}

if   [[ $1 = 0 ]]; then
    other;
elif [[ $1 = 1 ]]; then
    installRest;
elif [[ $1 = 2 ]]; then
    installSideCar;
elif [[ $1 = 3 ]]; then
    installNative;
elif [[ $1 = 4 ]]; then
    installMongo;
elif [[ $1 = 9 ]]; then
    installAll;
elif [[ $1 = 101 ]]; then
    installRest101;
elif [[ $1 = 102 ]]; then
    installRest102;
elif [[ $1 = 103 ]]; then
    installRest103;
elif [[ $1 = 104 ]]; then
    installSideCar104;
elif [[ $1 = 105 ]]; then
    installSideCar105;
elif [[ $1 = 106 ]]; then
    installSideCar106;
elif [[ $1 = 107 ]]; then
    installNative107;
elif [[ $1 = 108 ]]; then
    installNative108;
elif [[ $1 = 109 ]]; then
    installNative109;
elif [[ $1 = 110 ]]; then
    installNative110;
else
    other;
fi


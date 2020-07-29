#!/bin/zsh
p1=$( printf '%d' $1 )
echo                           

installAll() {
    echo Installing All
    installRest1-3
    installSideCar4-6
    installNative7-9
}

installRest1-3() {
    echo Installing REST 1-3
}

installSideCar4-6() {
    echo $p1 - Installing SideCar 4-6
}

installNative7-9() {
    echo $p1 - Installing Native 7-9
}

installSideCar4() {
    echo $p1 - Installing SideCar 4
}

installNative7() {
    echo $p1 - Installing Native 7
}


other() {
    echo $p1 is Not an Option, USAGE:
    echo "1:   Install REST    1-3 "
    echo "2:   Install SideCar 4-6 "
    echo "3:   Install Native  7-9 "
    echo "4:   Install ALL         "
    echo "104: Install SideCar 4   "
    echo "107: Install Native  7   "
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
    installAll;

elif [ $1 -eq 104 ]; then
    installSideCar4;

elif [ $1 -eq 107 ]; then
    installNative7;

else
    other;
fi



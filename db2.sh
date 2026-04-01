# Path to where you extracted clidriver
export DB2_HOME=/home/genaro/Downloads/clidriver

# Point the compiler to the header files
export CGO_CFLAGS="-I$DB2_HOME/include"

# Point the linker to the library files
export CGO_LDFLAGS="-L$DB2_HOME/lib"

# Ensure the OS can find the libraries at runtime
export LD_LIBRARY_PATH=$DB2_HOME/lib:$LD_LIBRARY_PATH

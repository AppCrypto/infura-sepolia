rm -rf ../build/*.bin
rm -rf ../build/*.abi
rm -rf ../gen/*.go
solc --bin --abi hello.sol -o ../build
abigen --bin ../build/Hello.bin --abi ../build/Hello.abi --pkg hello --out ../gen/hello.go
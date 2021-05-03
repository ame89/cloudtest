
name=$1
echo "compiling $name ..."
go build -o $name
echo "run $name ..."
./$name
rm ./$name
echo "$name done."


# PathTraverser : a go binary which traveses the given path tree and write the result into timestamp.txt file with fileheaders: filename,fileextension,size,hash



The following table shows the possible combinations of GOOS and GOARCH you can use:
![image](https://user-images.githubusercontent.com/22313972/127993993-ec631488-1f43-431c-8dd3-0677933d931b.png)


# Script to build OS respective binaries
cmd:  ./go-build.bash PathTraverser

# Usage while executing binary pass first argument as string filepath to traverse 

e.g.: ./PathTraverser_linux_amd64 "/home/sajeed/Dev"



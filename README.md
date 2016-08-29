# microservicesInGoWorkshop
Repo for the Google Developer Group Sardegna workshop

#From scratch

####Install GO
https://golang.org/
or use  Go Version manager (gvm)
https://github.com/moovweb/gvm

####GO ENV
export GOROOT=<path to your go installation>
export GOPATH=<path_to your_project>
export PATH=$PATH:$GOPATH/bin

(with gvm)
gvm use go1.7
export GOPATH=<path_to your_project>
export PATH=$PATH:$GOPATH/bin

####Create project structure 
(src already present)
mkdir bin pkg


####Install Glide (dependency manager)
https://github.com/Masterminds/glide
install dependencies with
glide install

####ENV VARS
KEYFILE=<path>/GDGWorshop/key.pem;PORT=8081;DEVELOPMENT=true;ADDRESS=127.0.0.1;APPNAME=gdg;CERTFILE=<path>/GDGWorshop/cert.pem;LOGPATH=<path>/gdg.log;LOGSERVER=localhost:1902

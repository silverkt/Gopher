export GOROOT=/usr/local/opt/go/libexec
# GOPAT为上面创建的目录路径
export GOPATH=/Users/tsing/Documents/Gopher
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home
PATH=$JAVA_HOME/bin:$PATH:.
CLASSPATH=$JAVA_HOME/lib:.

GRADLE_HOME=/usr/local/Caskroom/gradle
PATH=$GRADLE_HOME/bin:$PATH:.

export GRADLE_HOME
export JAVA_HOME
export PATH
export CLASSPATH
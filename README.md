![gopher](https://golang.org/doc/gopher/project.png)
![gopher](https://golang.org/doc/gopher/pkg.png)
![gopehr](https://golang.org/doc/gopher/run.png)
![gopher](https://golang.org/doc/gopher/ref.png)
![gopher](https://golang.org/doc/gopher/talks.png)


# HelloGolang
My Golang-Self-Study repository

## Installing Go

Install Golang by either downloading it [manually](https://golang.org/dl/) or use [Homebrew](http://brew.sh/).

### For Example :

```
mkdir /work/go
cd /work/go/
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
mkdir ~/go
tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz
vi /etc/profile
  ` Add The Following:
  export PATH=$PATH:/usr/local/go/bin
  `
source /etc/profile
go version
mkdir $HOME/go
mkdir $HOME/go/{bin,pkg,src}
vi ~/ .bash_profile
  ` Add The Following:
  # Golang
  export GOPATH=$HOME/go
  PATH=$PATH:$HOME/bin:$GOPATH/bin
  export PATH
  `
source ~/.bash_profile
```



## Helpful Links

- [A Tour of Go](https://tour.golang.org/list)
  A good place to start for a new Gopher.
- [Go By Example](https://gobyexample.com)
  A cheatsheet of how to do this or that particular thing in Golang. I still check this when I forget one detail or another.
- [Golang wiki](https://github.com/golang/go/wiki) Lots of good resources here.
- [Effective Go](https://golang.org/doc/effective_go.html)
  A reference of best practices when writing Golang. As you get more comfortable with the language, return to read this again and again to absorb all the good advice.
- [Go in 5 Minutes](https://github.com/arschles/go-in-5-minutes)
  A collection of short tutorials to solve basic problems. Somewhat advanced, but fun to return to now and then.
- [Go Newsletter](http://golangweekly.com)
  A weekly collection of interesting articles from the web. Don't worry if all the articles make sense to you. Just focus on whatever you find interesting.
- [Go Documentation](https://golang.org/doc/)
  The official collection of Golang-related documentation. Take a look!
- [Golang Packages](https://golang.org/pkg/)
  The complete list of the packages in the standard library with links to documentation about each. Once you become comfortable with Golang, you'll spend a lot of time reading through these pages.

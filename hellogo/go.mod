module github.com/jsilryan/Golang/hellogo

go 1.23.1

// replace github.com/jsilryan/Golang/My_Strings v0.0.0 => ../My_Strings // Makes things work locally instead of checking github but resolve it to the path ../My_Strings

require (
    github.com/jsilryan/Golang/My_Strings v0.0.0
)

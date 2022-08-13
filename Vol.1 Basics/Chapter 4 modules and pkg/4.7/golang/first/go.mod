module golang/first

go 1.19

replace golang/second => ../second // путь, относительно текущей директории проекта

require golang/second v0.0.0-00010101000000-000000000000

require github.com/google/uuid v1.3.0 // indirect

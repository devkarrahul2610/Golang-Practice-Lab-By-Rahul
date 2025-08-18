# Golang-Practice-Lab-By-Rahul

An array is a collection of elements of the same type and has a fixed length defined at the time of declaration. Once created its size cannot be changed. Arrays are value types in GO,which means assigning one array to another copies all its elements. 


eg. var arr[5] int // default values are zero.
    arr[0]=10

    * short declaration:
    a:=[3]string{"Go","Java","python"}

    Even go can infer the size.
    b:=[...]int{1,2,3,4,5}
# go_arrays
A collection of array methods for golang that should be included in the std  
*Note: This package was developed for golang 1.20, as of 1.21, some of this method have, indeed, been added to the std*

## Examples

```golang
arr := int[]{1,2,3,4,5}
array.Contains(arr, 2) // true
array.Contains(arr, 10) //false
```

```golang
arr := int[]{1,2,2,2,5}
array.Count(arr, 2) //2
array.Count(arr, 5) //1
array.Count(arr, 10)//10 
```

```golang
arr := int[]{1,2,3,4,5}
array.IndexOf(arr, 2) // *1
array.IndexOf(arr, 10) // null
```

```golang
arr := int[]{1,2,3,4,5}
array.Map(arr, func(int i) int {return i * 2}) // [2,4,6,8,10]
```

```golang
arr := int[]{1,2,3,4,5}
array.Filter(arr, func(int i) bool {return i % 2 == 0}) // [2,4]
```

```golang
arr := int[]{1,2,3,4,5}

array.SliceEquals(arr, []int{1,2,3,4,5}) //true
array.SliceEquals(arr, []int{5,4,3,2,1}) //false
```

## Instalation
```bash
//Run on the terminal
go get github.com/AlcalinoGitHub/go_arrays
```

```golang
//On your golang file
import (
  //Your other imports
  array "github.com/AlcalinoGitHub/go_arrays"
)
``` 


package ReflectionDemos

import "reflect"

type RefStudent struct {
   name string
   age int
}
func reflectDemo(payload ...interface{}) interface{}{
   args := make([]reflect.Value, len(payload))
   for i, arg := range payload {
       args[i] = reflect.ValueOf(arg)
   }
   var a interface{}
   a = args
   return a
}

package Data_Nullable

import "gopurs/output/gopurs_runtime"

var Null interface{} = nil

func Nullable(a interface{}, r interface{}, f func(interface{}) interface{}) interface{} {
    val := a.(gopurs_runtime.Value)
    if val.Type == gopurs_runtime.TypeAny {
        if val.UnsafePtr == nil {
            return r
        }
        if *(*any)(val.UnsafePtr) == nil {
            return r
        }
    }
    return f(a)
}

func NotNull(x interface{}) interface{} {
    return x
}

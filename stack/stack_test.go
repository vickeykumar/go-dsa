package stack


import (
    "fmt"
    "testing"
)

func TestStack(t *testing.T) {
    st := NewStack()
    st.Push(21)
    st.Push(23)
    st.Push(25)
    st.Push(27)
    st.Display(func(data interface{}) {
         fmt.Printf("%d\n",data.(int))
    })

    data:=st.Pop()
    t.Log("data deleted: ",data)
    t.Log("after delete:")
    st.Display(func(data interface{}) {
        fmt.Printf("%d\n",data.(int))
    })

}


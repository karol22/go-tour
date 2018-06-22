package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
 
func (x rot13Reader) Read(b []byte) (int,error) {  
    n, err := x.r.Read(b)  
    if err != io.EOF {  
        for i := 0; i<n; i++ {    
            if b[i]>='A' && b[i]<='A'+12 || b[i]>='a'&& b[i]<='a'+12 {  
                b[i]+=13  
            } else if b[i]>'Z'-13 && b[i]<='Z' || b[i]>'z'-13 && b[i]<= 'z' {  
                b[i]-= 13  
            }  
        }  
    }  
    return n, err  
}  

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

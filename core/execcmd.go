package core

/*
int system(const char *command);
void free(void *ptr);
*/
import "C"
import "unsafe"

func CallSystem(cmd_line string) int {
    cs := C.CString(cmd_line)
    ret := int(C.system(cs))
    C.free(unsafe.Pointer(cs))

    return ret
}

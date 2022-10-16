#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>


#include "../cdata/arrow/c/abi.h"

int  from_chunks_ffi(const ArrowArray *arrptr, const ArrowSchema *schptr, uintptr_t l);
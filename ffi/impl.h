#include <stdint.h>
#include "../cdata/arrow/c/abi.h"
int from_chunks_ffi(const struct ArrowArray *arrptr, const struct  ArrowSchema *schptr, uintptr_t l);
int call_with_ffi_voidptr(void* schema, void* array, uintptr_t l) { return from_chunks_ffi(array, schema, l); }


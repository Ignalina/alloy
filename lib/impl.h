// NOTE: You could use https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate
#include <stdint.h>
#include <stdlib.h>
#include "../ARROW_C_DATA_INTERFACE.h"
void call_with_ffi(ArrowArray* array, ArrowSchema* schema);

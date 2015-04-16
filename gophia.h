#include "sophia.h"

SP_API void *g_ctl(void* s) {
	return sp_ctl(s);
}

SP_API int g_destroy(void* s) {
	return sp_destroy(s);
}

SP_API void *g_async(void* s) {
	return sp_async(s);
}

SP_API void *g_object(void* s) {
	return sp_object(s);
}

SP_API int g_open(void* s) {
	return sp_open(s);
}

SP_API int g_error(void* s) {
	return sp_error(s);
}

SP_API int g_set__cfg(void* s, char* i, char* v) {
	return sp_set(s, i, v);
}

SP_API void *g_get__key(void* s, char* k) {
	return sp_get(s, k);
}

SP_API void *g_get__keyv(void* s, char* k) {
	return sp_get(s, k);
}

SP_API int g_delete(void* s) {
	return sp_delete(s);
}

SP_API int g_drop(void* s) {
	return sp_drop(s);
}

SP_API void *g_begin(void* s) {
	return sp_begin(s);
}

SP_API int g_prepare(void* s) {
	return sp_prepare(s);
}

SP_API int g_commit(void* s) {
	return sp_commit(s);
}

SP_API void *g_cursor(void* s) {
	return sp_cursor(s);
}

SP_API void *g_type(void* s) {
	return sp_type(s);
}

char* g__vtoc(void *v) {
	return (char *)v;
}

#ifndef CLIBRARY_H
#define CLIBRARY_H

#ifdef __cplusplus
extern "C" {
#endif


void ACFunction();

void init();

void deinit();

//! Submit query to engine implementation and return identification
int submitQuery(const char *query, char *uuid);


#ifdef __cplusplus
}
#endif

#endif
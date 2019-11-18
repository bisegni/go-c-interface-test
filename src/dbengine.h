#ifndef CLIBRARY_H
#define CLIBRARY_H

#ifdef __cplusplus
extern "C" {
#endif

#include "SqlTypes.h"

#include <stdint.h>

void ACFunction();

void init();

void deinit();

//! Submit query to engine implementation and return identification
int submitQuery(const char *query, char *uuid);
int columnCount(const char *uuid, int *column_count);
// int getColumnType(const char *uuid, SQlType type[]);
// int hasRow(const char *uuid, bool **has);

#ifdef __cplusplus
}
#endif

#endif
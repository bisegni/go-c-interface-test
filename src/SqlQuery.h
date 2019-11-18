#ifndef SQL_QUERY_H
#define SQL_QUERY_H

#include "SqlTypes.h"
#include "SqlColumnValue.h"

#include <string>
#include <vector>

class SQlQuery
{
    std::string query_id;
    int row_counter;
    inline std::string rndString(unsigned int size);
public:
    SQlQuery();
    ~SQlQuery();
    const std::string& getID();
    std::vector<SQlType> getSchema();
    bool hasNext() const;
    std::vector<SqlColumnValueShrdPtr> getRow();
};

#endif
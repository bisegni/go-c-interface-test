#ifndef SQL_QUERY_H
#define SQL_QUERY_H

#include <string>

class SQlQuery
{
    std::string query_id;
public:
    SQlQuery();
    ~SQlQuery();

    const std::string& getID();
};

#endif
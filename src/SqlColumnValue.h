#ifndef SQL_COLUMN_VALUE_H
#define SQL_COLUMN_VALUE_H

#include "SqlTypes.h"

#include <string>
#include <boost/variant.hpp>

class SqlColumnValue
{
    boost::variant<int32_t, int64_t, double, std::string> value;
    SQlType type;
public:
    explicit SqlColumnValue(int32_t int32_value);
    explicit SqlColumnValue(int64_t int64_value);
    explicit SqlColumnValue(double double_value);
    explicit SqlColumnValue(const std::string &string_value);
    ~SqlColumnValue();
    SQlType getType();
    
    template<typename T>
    T* getValue() {
        return boost::get<std::string>(value);
    }
};


typedef std::shared_ptr<SqlColumnValue> SqlColumnValueShrdPtr;
#endif
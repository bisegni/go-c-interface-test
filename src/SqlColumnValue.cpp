#include "SqlColumnValue.h"

SqlColumnValue::SqlColumnValue(int32_t int32_value) : type(SQlTypeInt32),
                                                      value(int32_value) {}

SqlColumnValue::SqlColumnValue(int64_t int64_value) : type(SQlTypeInt64),
                                                      value(int64_value) {}

SqlColumnValue::SqlColumnValue(double double_value) : type(SQlTypeDouble),
                                                      value(double_value) {}

SqlColumnValue::SqlColumnValue(const std::string &string_value) : type(SQLTypeString),
                                                                  value(string_value) {}

SqlColumnValue::~SqlColumnValue() {}

SQlType SqlColumnValue::getType() {return type;}
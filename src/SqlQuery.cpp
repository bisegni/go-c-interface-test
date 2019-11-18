#include "SqlQuery.h"

#include <boost/uuid/uuid.hpp>
#include <boost/uuid/uuid_io.hpp>
#include <boost/uuid/uuid_generators.hpp>

#include <boost/lexical_cast.hpp>

#include <random>

using namespace boost::uuids;

static std::random_device rd;
static std::uniform_int_distribution<int32_t> rnd_int32;
static std::uniform_int_distribution<int64_t> rnd_int64;
static std::uniform_int_distribution<double> rnd_double;
static std::uniform_int_distribution<int> distribution{'a', 'z'};

SQlQuery::SQlQuery() : row_counter(1)
{
    random_generator rnd_gen;
    uuid _uuid = rnd_gen();
    query_id = boost::uuids::to_string(_uuid);
}

SQlQuery::~SQlQuery() {}

const std::string &SQlQuery::getID() { return query_id; }

std::vector<SQlType> SQlQuery::getSchema()
{
    return {SQlTypeInt32, SQlTypeInt64, SQlTypeDouble, SQLTypeString};
}

bool SQlQuery::hasNext() const
{
    return (row_counter % 10) ? true : false;
}

std::vector<SqlColumnValueShrdPtr> SQlQuery::getRow()
{
    std::vector<SqlColumnValueShrdPtr> res;
    res.push_back(std::make_shared<SqlColumnValue>(rnd_int32(rd)));
    res.push_back(std::make_shared<SqlColumnValue>(rnd_int64(rd)));
    res.push_back(std::make_shared<SqlColumnValue>(rnd_double(rd)));
    res.push_back(std::make_shared<SqlColumnValue>(rndString(10)));
    return res;
}

//private methods
inline std::string SQlQuery::rndString(unsigned int size)
{
    std::string rand_str(size, '\0');
    for (auto &dis : rand_str)
    {
        dis = distribution(rd);
    };
    return rand_str;
}